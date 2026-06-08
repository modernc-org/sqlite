// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libc/sys/types"
	sqlite3 "modernc.org/sqlite/lib"
)

// ErrPageCacheTooLate is returned by RegisterPageCache when a SQLite
// connection has already been opened in this process. SQLITE_CONFIG_PCACHE2
// must be installed before sqlite3_initialize, which is called implicitly
// by the first sqlite3_open_v2. After that point SQLite returns
// SQLITE_MISUSE and the engine cannot switch its page cache backend.
var ErrPageCacheTooLate = errors.New(
	"sqlite: RegisterPageCache called after first Open; " +
		"SQLITE_CONFIG_PCACHE2 must be installed before any connection is opened")

// ErrPageCacheConflict is returned when a different PageCache has
// already been registered in this process. The same module value may be
// re-registered without error, which lets multiple library imports share
// a singleton without coordination.
var ErrPageCacheConflict = errors.New(
	"sqlite: a different page cache module is already registered")

// FetchMode tells Cache.Fetch how aggressively to allocate when the
// requested key is absent. It matches the createFlag of SQLite's xFetch
// (https://sqlite.org/c3ref/pcache_methods2.html).
type FetchMode int32

const (
	// FetchLookup looks up an existing entry. Fetch returns nil if the
	// key is not in cache; no allocation is performed.
	FetchLookup FetchMode = 0
	// FetchCreateEasy allocates only if it can be done without effort
	// (no eviction, no memory pressure). Fetch may return nil; SQLite
	// will then spill dirty pages and retry with FetchCreateForce.
	FetchCreateEasy FetchMode = 1
	// FetchCreateForce allocates unconditionally and may evict to make
	// room. Fetch should return nil only on a genuine out-of-memory
	// condition.
	FetchCreateForce FetchMode = 2
)

// PageCache is the factory for per-database Cache instances.
// SQLite calls Create once per open database; each call must return a
// fresh Cache with the given pageSize and extraSize. The extraSize
// includes SQLite's private PgHdr overhead and must be honoured as the
// opaque size of every Page's Extra buffer.
//
// purgeable is advisory: when false (in-memory databases), SQLite will
// only call Unpin with discard=true and the cache is permitted to free
// every page on Unpin. When true, the cache may retain unpinned pages
// for re-use.
type PageCache interface {
	Create(pageSize, extraSize int, purgeable bool) (Cache, error)
}

// Cache is one database's worth of cached pages. Methods may be
// invoked concurrently from multiple goroutines and MUST be threadsafe
// (this build assumes SQLITE_THREADSAFE=1).
//
// Implementations should NOT call RegisterPageCache directly or
// transitively. Callbacks run under the openGate read lock that the
// Open path holds, and a re-entrant Register would deadlock on the
// gate's write lock.
type Cache interface {
	// SetSize advises the cache of the new target page count
	// (PRAGMA cache_size). The cache is free to ignore the hint.
	SetSize(n int)

	// PageCount returns the number of pages currently held (pinned and
	// unpinned combined).
	PageCount() int

	// Fetch returns the Page for key or nil per FetchMode. SQLite will
	// call Fetch many times for the same key; the implementation MUST
	// return a Page whose Buf() and Extra() pointer addresses are
	// identical across calls for the same key until the binding signals
	// eviction via Unpin(discard=true), a Truncate sweep covering the
	// key, or Destroy on this cache.
	Fetch(key uint32, mode FetchMode) Page

	// Unpin tells the cache that the engine is finished using the
	// Page for now. If discard is true the implementation may release
	// the underlying memory immediately; if discard is false the
	// implementation MUST keep Buf() and Extra() valid at the same
	// addresses until a subsequent eviction signal arrives (Unpin
	// with discard=true, a Truncate covering the key, or Destroy).
	//
	// This is stricter than SQLite's xUnpin contract, which permits
	// the cache to evict unpinned pages on its own. The binding
	// caches the sqlite3_pcache_page stub per key and returns the
	// same stub to SQLite on every subsequent Fetch without
	// re-invoking Fetch, so an implementation that silently evicts
	// on Unpin(discard=false) will hand SQLite a stub pointing at
	// freed memory. SQLite never refcounts: one Unpin call is final
	// regardless of how many Fetches preceded.
	Unpin(p Page, discard bool)

	// Rekey changes the key under which p is filed from oldKey to
	// newKey. If an entry already exists at newKey it must be discarded
	// in the same call; SQLite guarantees the colliding entry is not
	// pinned at the moment of Rekey.
	Rekey(p Page, oldKey, newKey uint32)

	// Truncate discards every entry whose key is greater than or equal
	// to limit, including pinned entries. This is the only callback
	// permitted to evict a pinned page.
	Truncate(limit uint32)

	// Destroy releases every page and any resources owned by this
	// Cache. After Destroy returns the binding will not call any
	// other method on this Cache.
	Destroy()

	// Shrink hints the cache to release as much heap as possible. The
	// implementation is not obligated to free anything; this is purely
	// a memory-pressure advisory.
	Shrink()
}

// Page is one cache entry. Buf and Extra return pointers into
// implementation-owned memory that MUST remain valid and at the same
// address from the Fetch that first returned the Page until the cache
// signals eviction via Unpin(discard=true) on the same Page, a
// Truncate(limit) covering its key, or Destroy on the owning cache.
//
// The memory MUST be C-stable (libc.Xmalloc, runtime.Pinner-pinned
// slices, or an off-heap allocator). Go-heap memory the GC can relocate
// is forbidden: SQLite stores Buf and Extra addresses inside its own C
// structures and dereferences them without GC barriers.
//
// Buf must be at least pageSize bytes and is where SQLite stores the
// database page contents. Extra must be at least extraSize bytes (the
// extraSize passed to PageCache.Create, which already includes
// SQLite's PgHdr overhead) and is treated by SQLite as opaque scratch
// space. Implementations should zero Extra on a freshly-allocated Page
// so SQLite's PgHdr backpointer is read as null; the binding does not
// touch Extra contents.
type Page interface {
	Buf() unsafe.Pointer
	Extra() unsafe.Pointer
}

// pcacheState holds the package-global state shared between
// RegisterPageCache and the Open path.
//
// Locking discipline:
//
//   - openGate.RLock is held for the body of withOpenGate (called from
//     the Driver.Open path). Many opens may proceed concurrently; what
//     is forbidden is registering a page cache while any open is in
//     flight.
//   - openGate.Lock is held for the body of RegisterPageCache.
//     The write lock drains all in-flight opens and blocks all
//     subsequent opens until Xsqlite3_config completes.
//   - opened is set with an unconditional Store on every Open. The
//     Store is intentional and cheap; using CompareAndSwap to only
//     mutate on the first Open buys nothing because the read-side
//     under the write lock is uncontended. The atomic.Bool lets the
//     hot Open path read under RLock without paying for a full mutex
//     acquisition.
//   - configOnce guarantees Xsqlite3_config runs at most once per
//     process. A non-OK return code, an OOM during the methods-table
//     allocation, or a panic inside the once body leaves configErr
//     set and registered nil; every subsequent Register call returns
//     the sticky configErr. Reload is not supported in this MR.
//   - registered holds the canonical PageCache for idempotency
//     comparison; it is non-nil only after a successful install.
//   - cMethods is the libc.Xcalloc-owned C struct SQLite reads at
//     sqlite3_initialize time. Allocated once and lives until process
//     exit; allocating via libc avoids tripping Go's checkptr when
//     the transpiled C code reads the struct (the same reasoning that
//     vtab.go:130-141 uses for sqlite3_module).
var pcacheState struct {
	openGate   sync.RWMutex
	opened     atomic.Bool
	configOnce sync.Once
	configErr  error
	registered PageCache
	cMethods   uintptr
}

// markConnectionOpened is called from the Open path under
// pcacheState.openGate.RLock before sqlite3_open_v2. The RLock-side
// store happens-before RUnlock; RegisterPageCache's Lock
// acquisition waits for all readers to drain, so its subsequent Load
// observes every prior store.
func markConnectionOpened() {
	pcacheState.opened.Store(true)
}

// withOpenGate runs fn while holding the openGate read lock. The Open
// path wraps its entire body in this so a concurrent
// RegisterPageCache cannot squeeze in between the opened-flag
// store and sqlite3_open_v2.
func withOpenGate(fn func() error) error {
	pcacheState.openGate.RLock()
	defer pcacheState.openGate.RUnlock()
	markConnectionOpened()
	return fn()
}

// pcacheMethods2 is the Go view of the C sqlite3_pcache_methods2 struct
// as cznic transpiles it. Every supported GOOS/GOARCH pair exports the
// type with the same FiVersion/FpArg/FxInit/... field names; the per-arch
// Go compiler emits the correct byte layout. The old-generator arches
// (freebsd/386, freebsd/arm, netbsd/amd64) are not in build_all_targets
// and do not currently build for unrelated upstream reasons, so no shim
// is needed.
type pcacheMethods2 = sqlite3.Tsqlite3_pcache_methods2

// RegisterPageCache installs m as the process-global SQLite page
// cache via SQLITE_CONFIG_PCACHE2. It MUST be called before the first
// sql.Open or driver.Open in the program.
//
// Concurrency contract:
//
//   - Safe to call concurrently with itself and with other Register*
//     entry points.
//   - Blocks until any sql.Open calls currently in progress complete.
//     Trade-off: a Register call may block for the duration of an
//     in-flight Open. WAL recovery or cold-file-lock contention can
//     make that wait visible.
//   - Once any connection has been opened, returns ErrPageCacheTooLate
//     without mutating the global module slot.
//   - Calling twice with the same module value is a no-op success.
//     Calling twice with a different value returns ErrPageCacheConflict.
//   - A failed first install is sticky: every subsequent Register call
//     returns the same error. Mutating the module fields after the
//     first successful Register is silently ignored because SQLite has
//     already copied the C methods table.
func RegisterPageCache(m PageCache) error {
	if m == nil {
		return errors.New("sqlite: RegisterPageCache(nil)")
	}

	pcacheState.openGate.Lock()
	defer pcacheState.openGate.Unlock()

	// Idempotency / conflict / too-late checks before we touch the
	// once. pcacheState.registered is non-nil only after a SUCCESSFUL
	// install, so observing it here means a prior Register completed
	// without returning configErr.
	if pcacheState.registered != nil {
		if pcacheState.registered == m {
			return nil
		}
		if pcacheState.opened.Load() {
			return ErrPageCacheTooLate
		}
		return ErrPageCacheConflict
	}
	if pcacheState.opened.Load() {
		return ErrPageCacheTooLate
	}

	// First-time install. The once body commits
	// pcacheState.registered = m only after Xsqlite3_config succeeds.
	// On OOM, SQLite error, or panic, configErr is set and registered
	// stays nil; every subsequent Register returns configErr because
	// the once is already fired.
	pcacheState.configOnce.Do(func() {
		defer func() {
			if r := recover(); r != nil {
				pcacheState.configErr = fmt.Errorf(
					"sqlite: panic during PCACHE2 install: %v", r)
				pcacheState.registered = nil
				pcacheState.cMethods = 0
			}
		}()

		tls := libc.NewTLS()
		defer tls.Close()

		methodsPtr := libc.Xcalloc(tls, 1, types.Size_t(unsafe.Sizeof(pcacheMethods2{})))
		if methodsPtr == 0 {
			pcacheState.configErr = errors.New("sqlite: out of memory allocating pcache_methods2")
			return
		}
		populateCMethods(methodsPtr)

		varArgs := libc.Xmalloc(tls, types.Size_t(unsafe.Sizeof(uintptr(0))))
		if varArgs == 0 {
			libc.Xfree(tls, methodsPtr)
			pcacheState.configErr = errors.New("sqlite: out of memory allocating va_list")
			return
		}
		defer libc.Xfree(tls, varArgs)

		rc := sqlite3.Xsqlite3_config(tls,
			int32(sqlite3.SQLITE_CONFIG_PCACHE2),
			libc.VaList(varArgs, methodsPtr))
		if rc != sqlite3.SQLITE_OK {
			libc.Xfree(tls, methodsPtr)
			pcacheState.configErr = fmt.Errorf(
				"sqlite: Xsqlite3_config(SQLITE_CONFIG_PCACHE2) returned %d", rc)
			return
		}

		// Commit only after every fallible step succeeded.
		pcacheState.cMethods = methodsPtr
		pcacheState.registered = m
	})

	return pcacheState.configErr
}

// MustRegisterPageCache is like RegisterPageCache but
// panics on any error. Intended for init() use where a missing page
// cache is fatal. Mirrors the precedent set by
// MustRegisterDeterministicScalarFunction.
func MustRegisterPageCache(m PageCache) {
	if err := RegisterPageCache(m); err != nil {
		panic(err)
	}
}

// populateCMethods fills the libc-owned SQLite C methods table with the
// addresses of the package-internal trampolines defined in
// pagecache_trampolines.go. The user's PageCache is reached
// through pcacheState.registered, which is set after this call returns
// successfully. Using top-level trampolines avoids passing user
// function values to cFuncPointer, which is undefined for closures and
// method values.
func populateCMethods(ptr uintptr) {
	dst := (*pcacheMethods2)(unsafe.Pointer(ptr))
	dst.FiVersion = 1
	dst.FxInit = cFuncPointer(pcacheTrampolineInit)
	dst.FxShutdown = cFuncPointer(pcacheTrampolineShutdown)
	dst.FxCreate = cFuncPointer(pcacheTrampolineCreate)
	dst.FxCachesize = cFuncPointer(pcacheTrampolineCachesize)
	dst.FxPagecount = cFuncPointer(pcacheTrampolinePagecount)
	dst.FxFetch = cFuncPointer(pcacheTrampolineFetch)
	dst.FxUnpin = cFuncPointer(pcacheTrampolineUnpin)
	dst.FxRekey = cFuncPointer(pcacheTrampolineRekey)
	dst.FxTruncate = cFuncPointer(pcacheTrampolineTruncate)
	dst.FxDestroy = cFuncPointer(pcacheTrampolineDestroy)
	dst.FxShrink = cFuncPointer(pcacheTrampolineShrink)
}
