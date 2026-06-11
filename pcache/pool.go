// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pcache provides a bounded, LRU-evicting page cache suitable
// for registration with [modernc.org/sqlite] via
// [sqlite.RegisterPageCache]. It is the reference implementation that
// accompanies the SQLITE_CONFIG_PCACHE2 wrapper defined in the parent
// package.
//
// The contract this package implements is documented on
// [sqlite.PageCache], [sqlite.Cache], and [sqlite.Page]. In brief:
// SQLite calls [Pool.Create] once per database connection; each
// returned cache holds at most cache_size pages (the soft cap set by
// PRAGMA cache_size), evicts the least-recently-unpinned page when it
// has to make room, and stores page memory off the Go heap so SQLite's
// interior pointer arithmetic on the page extras does not trip the
// race detector's checkptr enforcement.
//
// Typical use:
//
//	import (
//	    "modernc.org/sqlite"
//	    "modernc.org/sqlite/pcache"
//	)
//
//	func init() {
//	    sqlite.MustRegisterPageCache(pcache.New())
//	}
//
// Call [New] once and save the result; two distinct *Pool values
// compare unequal and would trip [sqlite.ErrPageCacheConflict] if both
// were registered.
//
// Cross-connection sharing (one Pool serving multiple databases from
// a shared page set) is not in scope for this package today; each
// [sqlite.PageCache.Create] returns a fresh per-database cache. The
// hook for that lives at the binding level and is the subject of a
// separate MR.
package pcache

import (
	"container/list"
	"sync/atomic"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libc/sys/types"
	"modernc.org/sqlite"
)

// Pool is the factory side of the [sqlite.PageCache] contract. It is
// safe for concurrent [Pool.Create] calls (one per opening connection);
// each returned [sqlite.Cache] runs single-goroutine per the parent
// package's SQLITE_OPEN_FULLMUTEX + database/sql contract.
//
// Pool aggregates summary counters across every cache it has created
// so a long-running process can observe hit/miss/eviction behaviour
// through [Pool.Stats] without instrumenting individual caches.
type Pool struct {
	hits      atomic.Int64
	misses    atomic.Int64
	allocs    atomic.Int64
	evictions atomic.Int64
	rekeys    atomic.Int64
	truncates atomic.Int64
	caches    atomic.Int64
}

// New returns a Pool with empty counters. Multiple New values compare
// unequal; call once and reuse.
func New() *Pool { return &Pool{} }

// Stats is a snapshot of a Pool's lifetime counters. The counters are
// monotonically non-decreasing across the lifetime of the Pool.
type Stats struct {
	Hits      int64 // Fetches that found an existing page
	Misses    int64 // Fetches that did not find one
	Allocs    int64 // Fresh page allocations from libc
	Evictions int64 // LRU-driven page releases
	Rekeys    int64 // Rekey calls forwarded to the cache
	Truncates int64 // Truncate calls forwarded to the cache
	Caches    int64 // Caches created over the Pool's lifetime
}

// Stats returns a snapshot of the Pool's lifetime counters.
func (p *Pool) Stats() Stats {
	return Stats{
		Hits:      p.hits.Load(),
		Misses:    p.misses.Load(),
		Allocs:    p.allocs.Load(),
		Evictions: p.evictions.Load(),
		Rekeys:    p.rekeys.Load(),
		Truncates: p.truncates.Load(),
		Caches:    p.caches.Load(),
	}
}

// Create allocates a fresh per-database cache. It is the
// [sqlite.PageCache] entry point; SQLite calls it through the binding
// each time a new database connection's page cache is wired up.
func (p *Pool) Create(pageSize, extraSize int, purgeable bool) (sqlite.Cache, error) {
	p.caches.Add(1)
	return &cache{
		pool:      p,
		pageSize:  pageSize,
		extraSize: extraSize,
		purgeable: purgeable,
		pages:     map[uint32]*page{},
		lru:       list.New(),
		tls:       libc.NewTLS(),
	}, nil
}

// cache is one Pool-issued per-database cache. All methods run
// single-goroutine under the parent package's invariant; no
// per-instance synchronisation is required.
type cache struct {
	pool                *Pool
	pageSize, extraSize int
	purgeable           bool
	targetSize          int
	tls                 *libc.TLS

	pages map[uint32]*page
	// lru holds the unpinned pages with most-recently-unpinned at
	// Front and least-recently-unpinned at Back. Pinned pages are
	// absent from the list and have a nil lruElem.
	lru       *list.List
	destroyed bool

	inflight int32 // PROBE ONLY: goroutines currently inside a method of this cache
}

// PROBE ONLY (experiment, not for merge). The probe uses atomics only —
// it adds NO mutual exclusion — so it measures whether two goroutines
// actually execute this cache's callbacks at the same time, independent
// of the race detector and of any lock. ProbeMaxInflight stays 1 iff the
// callbacks are genuinely serialised (e.g. by SQLite's BtShared mutex);
// a value >1 would be hard proof of real concurrent entry.
//
// Note: because these atomics touch a shared location on entry/exit,
// they incidentally give the race detector a happens-before edge it
// otherwise lacks for libc's pthread mutex — so running the probe build
// under -race SILENCES the (false-positive) report. Read ProbeMaxInflight
// without -race; the number, not the detector, is the evidence.
var (
	ProbeMaxInflight      int32
	ProbeConcurrentEvents int64
)

// ResetProbe zeroes the probe counters; call at the start of a probe test.
func ResetProbe() {
	atomic.StoreInt32(&ProbeMaxInflight, 0)
	atomic.StoreInt64(&ProbeConcurrentEvents, 0)
}

func (c *cache) probeEnter() func() {
	n := atomic.AddInt32(&c.inflight, 1)
	if n > 1 {
		atomic.AddInt64(&ProbeConcurrentEvents, 1)
	}
	for {
		m := atomic.LoadInt32(&ProbeMaxInflight)
		if n <= m || atomic.CompareAndSwapInt32(&ProbeMaxInflight, m, n) {
			break
		}
	}
	return func() { atomic.AddInt32(&c.inflight, -1) }
}

// page is one cached entry. It is comparable via *page identity, which
// the parent package's binding uses to detect retain-vs-replace
// across Fetches.
//
// buf and extra are stored as uintptr because the underlying memory is
// allocated by libc.Xmalloc / libc.Xcalloc and lives outside the Go
// heap, so the Go GC has no pointer to track. Buf and Extra convert
// to unsafe.Pointer only at the package boundary the binding requires.
// go vet flags the conversion as a "possible misuse" because its
// pattern-matcher does not distinguish off-heap C addresses from Go
// pointers; the same warning is tolerated at pagecache.go:347 and
// pagecache_trampolines.go:272 in the parent package.
type page struct {
	owner   *cache
	buf     uintptr
	extra   uintptr
	key     uint32
	pinned  bool
	lruElem *list.Element
}

func (p *page) Buf() unsafe.Pointer   { return unsafe.Pointer(p.buf) }
func (p *page) Extra() unsafe.Pointer { return unsafe.Pointer(p.extra) }

// SetSize records the cache_size advisory and evicts down to it if the
// new target is below the current count. A zero or negative target is
// treated as "unbounded" and disables eviction-on-trim.
func (c *cache) SetSize(n int) {
	defer c.probeEnter()()
	if n < 0 {
		n = 0
	}
	c.targetSize = n
	c.trimToTarget()
}

// PageCount returns the total number of pages held, pinned plus
// unpinned.
func (c *cache) PageCount() int { defer c.probeEnter()(); return len(c.pages) }

// Fetch implements the [sqlite.Cache] Fetch contract.
//
// On a hit (key already present) the page is marked pinned and
// returned regardless of mode; this is safe even for FetchLookup
// because re-pinning a page that was already in cache is a free
// operation.
//
// On a miss the behaviour depends on mode: FetchLookup returns nil,
// FetchCreateEasy allocates only if doing so requires no eviction
// (the cache is below targetSize), and FetchCreateForce allocates
// unconditionally, evicting the LRU tail when at or above targetSize.
// If every page is pinned and the cache is at target, FetchCreateForce
// overcommits and returns a fresh page; pcache1 behaves identically
// and SQLite needs the page to make forward progress.
func (c *cache) Fetch(key uint32, mode sqlite.FetchMode) sqlite.Page {
	defer c.probeEnter()()
	if c.destroyed {
		return nil
	}
	if existing, ok := c.pages[key]; ok {
		c.pool.hits.Add(1)
		if !existing.pinned {
			existing.pinned = true
			if existing.lruElem != nil {
				c.lru.Remove(existing.lruElem)
				existing.lruElem = nil
			}
		}
		return existing
	}
	c.pool.misses.Add(1)
	if mode == sqlite.FetchLookup {
		return nil
	}
	if mode == sqlite.FetchCreateEasy {
		if c.targetSize > 0 && len(c.pages) >= c.targetSize {
			return nil
		}
	}
	if mode == sqlite.FetchCreateForce && c.targetSize > 0 && len(c.pages) >= c.targetSize {
		c.evictOneLRU()
	}
	p := c.allocPage(key)
	if p == nil {
		return nil
	}
	p.pinned = true
	c.pages[key] = p
	return p
}

// Unpin implements the [sqlite.Cache] Unpin contract.
//
// discard=true: page is released immediately and its memory freed.
//
// discard=false on a purgeable cache: page is parked at the LRU front
// for reuse; the cache then trims to targetSize, which may evict the
// LRU tail (possibly this very page if the cache is already at cap).
//
// discard=false on a non-purgeable cache (in-memory database): the
// page is freed. SQLite documents that Unpin on a non-purgeable cache
// is only ever called with discard=true in practice, so this branch
// is a safety net rather than a hot path.
func (c *cache) Unpin(pg sqlite.Page, discard bool) {
	defer c.probeEnter()()
	if c.destroyed {
		return
	}
	p, ok := pg.(*page)
	if !ok || p.owner != c || !p.pinned {
		return
	}
	p.pinned = false
	if discard || !c.purgeable {
		c.freePage(p)
		c.pool.evictions.Add(1)
		return
	}
	p.lruElem = c.lru.PushFront(p)
	c.trimToTarget()
}

// Rekey implements the [sqlite.Cache] Rekey contract. The SQLite spec
// guarantees the colliding entry at newKey, if any, is unpinned at the
// moment of the call; the cache discards it here so the binding's
// stale-stub bookkeeping has nothing left to point at.
func (c *cache) Rekey(pg sqlite.Page, oldKey, newKey uint32) {
	defer c.probeEnter()()
	if c.destroyed {
		return
	}
	p, ok := pg.(*page)
	if !ok || p.owner != c {
		return
	}
	if collider, ok := c.pages[newKey]; ok && collider != p {
		c.freePage(collider)
	}
	delete(c.pages, oldKey)
	p.key = newKey
	c.pages[newKey] = p
	c.pool.rekeys.Add(1)
}

// Truncate implements the [sqlite.Cache] Truncate contract: every
// entry whose key is greater than or equal to limit is released,
// including pinned entries.
func (c *cache) Truncate(limit uint32) {
	defer c.probeEnter()()
	if c.destroyed {
		return
	}
	for key, p := range c.pages {
		if key >= limit {
			c.freePage(p)
		}
	}
	c.pool.truncates.Add(1)
}

// Destroy releases every page (pinned or unpinned), closes the
// per-cache libc TLS, and marks the cache so subsequent callbacks
// from the binding become no-ops. The parent binding will not call
// any other method after Destroy.
func (c *cache) Destroy() {
	defer c.probeEnter()()
	if c.destroyed {
		return
	}
	for _, p := range c.pages {
		c.freeRaw(p)
	}
	c.pages = nil
	if c.lru != nil {
		c.lru.Init()
		c.lru = nil
	}
	c.destroyed = true
	if c.tls != nil {
		c.tls.Close()
		c.tls = nil
	}
}

// Shrink drops every unpinned page, regardless of targetSize. SQLite
// uses this as a memory-pressure hint.
func (c *cache) Shrink() {
	defer c.probeEnter()()
	if c.destroyed {
		return
	}
	for c.lru.Len() > 0 {
		elem := c.lru.Back()
		p := elem.Value.(*page)
		c.lru.Remove(elem)
		p.lruElem = nil
		delete(c.pages, p.key)
		c.freeRaw(p)
		c.pool.evictions.Add(1)
	}
}

// allocPage returns a fresh *page with libc-backed buf and extra. The
// extra buffer is zeroed via Xcalloc because SQLite reads PgHdr.pPage
// from the head of Extra as a null check on first use.
func (c *cache) allocPage(key uint32) *page {
	buf := libc.Xmalloc(c.tls, types.Size_t(c.pageSize))
	if buf == 0 {
		return nil
	}
	extra := libc.Xcalloc(c.tls, 1, types.Size_t(c.extraSize))
	if extra == 0 {
		libc.Xfree(c.tls, buf)
		return nil
	}
	c.pool.allocs.Add(1)
	return &page{owner: c, buf: buf, extra: extra, key: key}
}

// freePage removes p from the cache's bookkeeping and releases its
// off-heap memory. Safe whether p is pinned or in lru.
func (c *cache) freePage(p *page) {
	if p.lruElem != nil {
		c.lru.Remove(p.lruElem)
		p.lruElem = nil
	}
	delete(c.pages, p.key)
	c.freeRaw(p)
}

// freeRaw releases p's off-heap memory without touching the maps.
// Used by Destroy where the maps are torn down in bulk.
func (c *cache) freeRaw(p *page) {
	if p.buf != 0 {
		libc.Xfree(c.tls, p.buf)
		p.buf = 0
	}
	if p.extra != 0 {
		libc.Xfree(c.tls, p.extra)
		p.extra = 0
	}
}

// evictOneLRU drops the least-recently-unpinned page and returns
// whether anything was evicted. Returns false when every page is
// pinned (forces an overcommit on the caller's allocation).
func (c *cache) evictOneLRU() bool {
	elem := c.lru.Back()
	if elem == nil {
		return false
	}
	p := elem.Value.(*page)
	c.freePage(p)
	c.pool.evictions.Add(1)
	return true
}

// trimToTarget evicts LRU-tail entries until len(pages) <=
// targetSize or no more unpinned pages remain. A zero targetSize
// disables trimming.
func (c *cache) trimToTarget() {
	if c.targetSize <= 0 {
		return
	}
	for len(c.pages) > c.targetSize {
		if !c.evictOneLRU() {
			return
		}
	}
}

// Compile-time interface assertions.
var (
	_ sqlite.PageCache = (*Pool)(nil)
	_ sqlite.Cache     = (*cache)(nil)
	_ sqlite.Page      = (*page)(nil)
)
