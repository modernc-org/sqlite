// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"math/bits"
	"sync"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libc/sys/types"
	sqlite3 "modernc.org/sqlite/lib"
)

// The pcache binding owns the sqlite3_pcache_page stub that SQLite sees
// and forwards user-facing calls to the registered PageCache and
// per-database Cache instances. The user-facing surface in
// pagecache.go is idiomatic Go; this file holds the package-internal
// trampolines that bridge to the cgo-free C ABI.
//
// Ownership model:
//
//   - Each xCreate on the C side mints a fresh pcacheBinding and a
//     uintptr handle from pcacheBindings.gen. SQLite stores the handle
//     as its opaque sqlite3_pcache*; trampolines recover the binding
//     via pcacheBindings.m under RLock.
//   - Each xFetch on the C side maps a (binding, key) pair to a
//     sqlite3_pcache_page stub allocated via libc.Xcalloc. The stub
//     lives until Unpin(discard=true), a Truncate sweep covering the
//     key, or Destroy on the owning cache. SQLite re-fetches for the
//     same key return the same stub; this is the binding-side fix for
//     the "stale pPgHdr->pPage" issue called out in MR #126's
//     §2 design note.
//   - The Page object the user returns from Cache.Fetch is held
//     verbatim in the binding's byKey map. Buf() and Extra() are
//     called only at stub-mint time; SQLite reads pBuf/pExtra
//     directly from the stub on subsequent fetches, so the user impl
//     does not pay for repeated interface dispatch on the hot path.
type pcacheBinding struct {
	mu     sync.Mutex
	cache  Cache
	byKey  map[uint32]uintptr      // key -> sqlite3_pcache_page stub uintptr
	byStub map[uintptr]pcacheEntry // stub uintptr -> {Page, key}
}

type pcacheEntry struct {
	page Page
	key  uint32
}

// pcacheBindings is the package-global registry mapping the opaque
// sqlite3_pcache* handles SQLite sees to the *pcacheBinding they
// represent. Handles are minted from a bitset-backed allocator so they
// can be reclaimed on Destroy and stay small across the process
// lifetime (otherwise a long-running process that opens and closes many
// caches would grow the map without bound).
var pcacheBindings = struct {
	mu  sync.RWMutex
	m   map[uintptr]*pcacheBinding
	gen pcacheIDGen
}{
	m: map[uintptr]*pcacheBinding{},
}

// lookupPcacheBinding returns the binding for the given C handle, or
// nil if the handle is unknown. It is called from every trampoline
// other than Init/Shutdown.
func lookupPcacheBinding(h uintptr) *pcacheBinding {
	pcacheBindings.mu.RLock()
	defer pcacheBindings.mu.RUnlock()
	return pcacheBindings.m[h]
}

// registerPcacheBinding installs b into the registry and returns the
// opaque handle SQLite will see.
func registerPcacheBinding(b *pcacheBinding) uintptr {
	pcacheBindings.mu.Lock()
	defer pcacheBindings.mu.Unlock()
	h := pcacheBindings.gen.next()
	pcacheBindings.m[h] = b
	return h
}

// unregisterPcacheBinding removes h from the registry and reclaims its
// ID slot. Called from pcacheTrampolineDestroy after the user's
// Destroy returns.
func unregisterPcacheBinding(h uintptr) {
	pcacheBindings.mu.Lock()
	defer pcacheBindings.mu.Unlock()
	delete(pcacheBindings.m, h)
	pcacheBindings.gen.reclaim(h)
}

// pcacheIDGen is a bitset-backed allocator over uintptr handles
// starting at 1 (0 is reserved as a sentinel for "not registered").
// Identical in shape to the vtab idGen at sqlite.go:808-824; reproduced
// here to keep the pcache subsystem self-contained.
type pcacheIDGen struct {
	bitset []uint64
}

func (g *pcacheIDGen) next() uintptr {
	base := uintptr(1)
	for i := 0; i < len(g.bitset); i, base = i+1, base+64 {
		b := g.bitset[i]
		if b != 1<<64-1 {
			n := uintptr(bits.TrailingZeros64(^b))
			g.bitset[i] |= 1 << n
			return base + n
		}
	}
	g.bitset = append(g.bitset, 1)
	return base
}

func (g *pcacheIDGen) reclaim(id uintptr) {
	bit := id - 1
	g.bitset[bit/64] &^= 1 << (bit % 64)
}

// Sentinel returned by the binding when SQLite asks for a stub but the
// user impl returned nil or the binding could not allocate. Translating
// to uintptr(0) tells SQLite to fall back to its stress path.
const pcacheNullStub uintptr = 0

// pcacheStubSize is the size of the sqlite3_pcache_page struct that
// SQLite reads: two uintptr fields (pBuf, pExtra). Computed at startup
// to keep the trampolines free of unsafe.Sizeof calls on the hot path.
var pcacheStubSize = types.Size_t(unsafe.Sizeof(uintptr(0)) * 2)

// pcacheTrampolineInit is wired to sqlite3_pcache_methods2.xInit. The
// binding has no per-process state to set up at sqlite3_initialize
// time; PageCache.Create is invoked per database, not per
// initialize, and the package-global pcacheBindings registry is ready
// from package init. Returns SQLITE_OK unconditionally.
func pcacheTrampolineInit(tls *libc.TLS, arg uintptr) int32 {
	return sqlite3.SQLITE_OK
}

// pcacheTrampolineShutdown is wired to xShutdown. The binding does not
// call sqlite3_shutdown internally, so in normal operation this
// trampoline never fires; it exists so the C methods table never
// holds a null xShutdown slot when paired with custom user code that
// might trigger sqlite3_shutdown out of band.
func pcacheTrampolineShutdown(tls *libc.TLS, arg uintptr) {}

// pcacheTrampolineCreate is wired to xCreate. SQLite asks the binding
// to create a new per-database cache. The trampoline calls
// PageCache.Create on the registered module and registers the
// returned Cache in pcacheBindings, handing SQLite back the opaque
// handle.
//
// szPage and szExtra are the sizes SQLite wants for each page's Buf
// and Extra buffers; szExtra is NOT just the application-visible
// extraSize, it already includes SQLite's PgHdr overhead. The
// Cache implementation must honour it as opaque.
//
// bPurgeable is advisory: 0 for in-memory databases (Unpin will only
// ever be called with discard=true), non-zero otherwise.
//
// A return of 0 tells SQLite the cache could not be created (treated
// as SQLITE_NOMEM by the caller in _sqlite3PcacheSetPageSize).
func pcacheTrampolineCreate(tls *libc.TLS, szPage, szExtra, bPurgeable int32) uintptr {
	module := pcacheState.registered
	if module == nil {
		// Defensive: should be impossible because SQLite only consults
		// the methods table after _sqlite3Config.Fpcache2 is installed,
		// which happens only after registered is set.
		return pcacheNullStub
	}
	cache, err := module.Create(int(szPage), int(szExtra), bPurgeable != 0)
	if err != nil || cache == nil {
		return pcacheNullStub
	}
	b := &pcacheBinding{
		cache:  cache,
		byKey:  map[uint32]uintptr{},
		byStub: map[uintptr]pcacheEntry{},
	}
	return registerPcacheBinding(b)
}

// pcacheTrampolineCachesize forwards SQLite's xCachesize to
// Cache.SetSize. The size is advisory.
func pcacheTrampolineCachesize(tls *libc.TLS, pCache uintptr, nCachesize int32) {
	b := lookupPcacheBinding(pCache)
	if b == nil {
		return
	}
	b.cache.SetSize(int(nCachesize))
}

// pcacheTrampolinePagecount forwards xPagecount to Cache.PageCount.
func pcacheTrampolinePagecount(tls *libc.TLS, pCache uintptr) int32 {
	b := lookupPcacheBinding(pCache)
	if b == nil {
		return 0
	}
	return int32(b.cache.PageCount())
}

// pcacheTrampolineFetch forwards xFetch and owns the
// sqlite3_pcache_page stub identity. If the binding already has a
// stub for this key it returns it without calling the user impl
// (SQLite's contract requires identical pointer on re-fetch for the
// same key). Otherwise it asks the user impl for a Page and mints a
// fresh stub via libc.Xcalloc, copying Buf and Extra into the stub
// once.
func pcacheTrampolineFetch(tls *libc.TLS, pCache uintptr, key uint32, createFlag int32) uintptr {
	b := lookupPcacheBinding(pCache)
	if b == nil {
		return pcacheNullStub
	}
	b.mu.Lock()
	if stub, ok := b.byKey[key]; ok {
		b.mu.Unlock()
		return stub
	}
	b.mu.Unlock()
	page := b.cache.Fetch(key, FetchMode(createFlag))
	if page == nil {
		return pcacheNullStub
	}
	stub := libc.Xcalloc(tls, 1, pcacheStubSize)
	if stub == 0 {
		return pcacheNullStub
	}
	stubFields := (*[2]uintptr)(unsafe.Pointer(stub))
	stubFields[0] = uintptr(page.Buf())
	stubFields[1] = uintptr(page.Extra())
	b.mu.Lock()
	// Re-check: another goroutine may have raced ahead while we were
	// outside the lock. If so, drop our stub and return theirs.
	if existing, ok := b.byKey[key]; ok {
		b.mu.Unlock()
		libc.Xfree(tls, stub)
		return existing
	}
	b.byKey[key] = stub
	b.byStub[stub] = pcacheEntry{page: page, key: key}
	b.mu.Unlock()
	return stub
}

// pcacheTrampolineUnpin forwards xUnpin to Cache.Unpin. When
// discard is non-zero the entry is evicted: the user's Unpin is
// invoked first, then the binding frees the stub and drops it from
// both maps. When discard is zero the binding only forwards to the
// user (the cache MAY retain the entry); the stub stays alive and
// will be returned again on the next Fetch for the same key.
func pcacheTrampolineUnpin(tls *libc.TLS, pCache, pPage uintptr, discard int32) {
	b := lookupPcacheBinding(pCache)
	if b == nil {
		return
	}
	b.mu.Lock()
	entry, ok := b.byStub[pPage]
	b.mu.Unlock()
	if !ok {
		return
	}
	b.cache.Unpin(entry.page, discard != 0)
	if discard == 0 {
		return
	}
	// Re-check under the second lock acquisition before freeing:
	// a concurrent Truncate covering entry.key could have already
	// freed and deleted this stub between the first Unlock and now.
	// Freeing twice would corrupt the libc allocator.
	b.mu.Lock()
	_, stillThere := b.byStub[pPage]
	if stillThere {
		delete(b.byKey, entry.key)
		delete(b.byStub, pPage)
	}
	b.mu.Unlock()
	if stillThere {
		libc.Xfree(tls, pPage)
	}
}

// pcacheTrampolineRekey forwards xRekey to Cache.Rekey, then
// updates the binding's byKey map so a subsequent Fetch for newKey
// returns the same stub. If newKey already has a stub the spec
// guarantees the colliding entry is unpinned; we evict it via libc.Xfree.
func pcacheTrampolineRekey(tls *libc.TLS, pCache, pPage uintptr, oldKey, newKey uint32) {
	b := lookupPcacheBinding(pCache)
	if b == nil {
		return
	}
	b.mu.Lock()
	entry, ok := b.byStub[pPage]
	b.mu.Unlock()
	if !ok {
		return
	}
	b.cache.Rekey(entry.page, oldKey, newKey)
	// Re-check under the second lock acquisition before mutating: a
	// concurrent Truncate covering oldKey could have already freed
	// pPage and dropped it from byStub. Without the re-check we would
	// resurrect the freed pointer into byKey/byStub and a subsequent
	// Fetch would return a stub pointing at libc-freed memory.
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, stillThere := b.byStub[pPage]; !stillThere {
		return
	}
	if collider, ok := b.byKey[newKey]; ok && collider != pPage {
		delete(b.byKey, newKey)
		delete(b.byStub, collider)
		libc.Xfree(tls, collider)
	}
	delete(b.byKey, oldKey)
	b.byKey[newKey] = pPage
	entry.key = newKey
	b.byStub[pPage] = entry
}

// pcacheTrampolineTruncate forwards xTruncate to Cache.Truncate
// and then sweeps every binding-owned stub whose key is greater than
// or equal to limit. SQLite documents Truncate as the only callback
// permitted to evict a pinned page, so the binding does the eviction
// unconditionally.
func pcacheTrampolineTruncate(tls *libc.TLS, pCache uintptr, iLimit uint32) {
	b := lookupPcacheBinding(pCache)
	if b == nil {
		return
	}
	b.cache.Truncate(iLimit)
	b.mu.Lock()
	for key, stub := range b.byKey {
		if key >= iLimit {
			delete(b.byKey, key)
			delete(b.byStub, stub)
			libc.Xfree(tls, stub)
		}
	}
	b.mu.Unlock()
}

// pcacheTrampolineDestroy forwards xDestroy to Cache.Destroy,
// frees every remaining stub, and removes the binding from the
// registry. After this trampoline returns the SQLite handle is no
// longer valid and the binding instance is unreachable.
func pcacheTrampolineDestroy(tls *libc.TLS, pCache uintptr) {
	b := lookupPcacheBinding(pCache)
	if b == nil {
		return
	}
	b.cache.Destroy()
	b.mu.Lock()
	for _, stub := range b.byKey {
		libc.Xfree(tls, stub)
	}
	b.byKey = nil
	b.byStub = nil
	b.mu.Unlock()
	unregisterPcacheBinding(pCache)
}

// pcacheTrampolineShrink forwards xShrink to Cache.Shrink. SQLite
// calls this as a memory-pressure hint; the user impl is free to
// ignore it.
func pcacheTrampolineShrink(tls *libc.TLS, pCache uintptr) {
	b := lookupPcacheBinding(pCache)
	if b == nil {
		return
	}
	b.cache.Shrink()
}
