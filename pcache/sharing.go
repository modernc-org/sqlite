// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcache

// This file is a scaffold for cross-connection / shared-cache support.
// It compiles but does not yet wire any new behavior. See the MR-C
// description for the three open design questions that gate the
// implementation.
//
// Background: the !127 implementation deliberately scoped out
// cross-connection sharing. Each Pool.Create returns an independent
// per-database cache, and the per-cache state is touched by exactly
// one goroutine at a time because every connection is opened
// SQLITE_OPEN_FULLMUTEX, no shared cache is enabled at the engine
// level, and database/sql never uses one driver.Conn from two
// goroutines.
//
// SQLite's shared-cache mode (cache=shared URI, or the deprecated
// sqlite3_enable_shared_cache flag) breaks that invariant: the engine
// reuses one pCache* across every Btree opened against the same
// shared-cache scope, so calls into our Cache may overlap. The current
// implementation would race in that case. cznic flagged this on the
// !127 merge as "the assumption MR-C will need to revisit."
//
// What this MR is asking for:
//
// Q1. Concurrency primitive. sync.Mutex around every Cache method is
//     the simplest correct shape. sync.RWMutex buys very little because
//     Fetch is read-modify (it touches the LRU position or allocates
//     on a miss). A finer-grained per-bucket lock would not pay off
//     either because every callback runs short. Current lean:
//     sync.Mutex on cache. Open to a different direction.
//
// Q2. Locking surface. Three plausible shapes:
//       (a) the lock lives on the existing cache struct in pool.go and
//           is always taken. Zero behavioral difference for non-shared
//           callers, but the uncontended-lock cost is paid by every
//           connection regardless of mode.
//       (b) a sharedCache wrapper type that adds the lock and is
//           returned from Pool.Create only when the engine asks for a
//           shared cache. Cheaper for the common (non-shared) case but
//           requires the binding to forward a shared-cache signal to
//           the PageCache, which the PageCache.Create signature does
//           not expose today.
//       (c) make sqlite.Cache optionally implement a concurrency hint
//           (e.g. ConcurrentSafe() bool) and let the binding wrap when
//           it returns false. Pushes the choice down to the impl and
//           keeps the wrapper-vs-not split per impl, not per cache
//           instance.
//
// Q3. Discovery. The PageCache.Create call does not currently receive
//     an indicator of shared-cache scope. Three options for signalling
//     it:
//       (a) extend PageCache.Create with a new parameter. Breaks the
//           sqlite.PageCache interface from !126.
//       (b) add a separate PageCache.CreateShared entry point.
//           Backward-compatible but doubles the API surface.
//       (c) detect at the binding level (read the URI on the parent
//           conn at xCreate time) and call a different impl method.
//           Keeps the user-facing PageCache interface clean and
//           confines the new code to the binding.
//
// Test surface this MR is also asking about:
//
//   - A canonical e2e test for shared-cache safety. The shape is two
//     connections to the same database with cache=shared, concurrent
//     writes, PRAGMA integrity_check at the end. Easy to write once
//     the locking shape is settled; left as a stub here so the test
//     name is reserved.
//
// No tests are added by this MR yet because the locking surface
// (Q1/Q2) determines whether they belong on cache, on a sharedCache
// wrapper, or on the binding. Once a direction is picked the
// implementation, the test, the CHANGELOG entry, and any binding
// changes follow as one focused commit.

// sharedCacheStub is a placeholder type. It exists so the file is
// non-empty and the compiler reaches the docs above; the real type
// (or the decision that no new type is needed) follows the resolution
// of Q1/Q2.
type sharedCacheStub struct{}

// TODO(MR-C): once Q1/Q2/Q3 are settled:
//
//   - if Q2(a): drop sharedCacheStub, add sync.Mutex to cache in
//     pool.go, lock all six callback methods, document the lock in
//     the cache struct comment.
//   - if Q2(b): replace sharedCacheStub with a real wrapper type that
//     embeds *cache and adds sync.Mutex; teach Pool.Create to return
//     the wrapper when xCreate signals shared-cache scope (per Q3).
//   - if Q2(c): drop sharedCacheStub, add a ConcurrentSafe optional
//     method to the sqlite.Cache contract in the parent package and
//     have the binding wrap when the impl returns false.
//
// In every branch, add a TestSharedCacheTwoConns_Integrity that opens
// two sql.Conn against the same cache=shared db, runs concurrent
// writes, and asserts PRAGMA integrity_check is ok at the end.
