// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcache

// This file documents how the pcache pool behaves under SQLite's
// shared-cache mode (cache=shared URI parameter or the deprecated
// sqlite3_enable_shared_cache flag). The implementation lives in
// pool.go on the cache type's mu sync.Mutex; this file exists as the
// design-record companion to that lock.
//
// Background. The !127 pcache implementation was deliberately scoped
// to one pCache per connection: SQLITE_OPEN_FULLMUTEX + no shared
// cache + database/sql's one-conn-one-goroutine invariant meant every
// callback on a given cache ran single-goroutine. cznic flagged on
// the !127 merge that this assumption would need to be revisited
// when shared-cache support landed.
//
// What changes under cache=shared. The engine reuses one pCache*
// across every Btree opened against the same shared-cache scope, so
// our pcache callbacks may be invoked from different goroutines in
// quick succession. They do not actually overlap: SQLite serialises
// every callback through sqlite3BtreeEnter on the BtShared mutex
// (SQLITE_THREADSAFE=1, the default for modernc.org/sqlite). The
// pcache pool is therefore correct under cache=shared by virtue of
// SQLite's own serialisation; verified empirically with a lock-free
// in-flight probe (max-in-flight = 1 on the canonical two-connection
// shared-cache workload, 4 on a positive control where goroutines
// hit the cache directly without SQLite's mutex).
//
// Why the per-cache sync.Mutex is taken anyway. The Go race detector
// does not recognise SQLite's libc mutex as a happens-before edge.
// Without an additional Go-level synchronisation primitive it reports
// the Fetch vs Unpin reads/writes on cache.pages / page.pinned /
// page.lruElem as DATA RACE, even though the SQLite-level mutex
// guarantees they never actually overlap. The per-cache sync.Mutex
// on the cache type, always taken, hands the race detector the
// happens-before edge it needs and turns the false-positive into
// a clean run.
//
// Cost. On the common (non-shared-cache) path every callback runs
// single-goroutine anyway, so the lock is always uncontended:
// roughly one atomic CAS per Lock/Unlock pair, negligible next to
// the SQLite work the callback bookends. On the shared-cache path
// the lock is also nearly uncontended because SQLite's BtShared
// mutex has already serialised the callers; the Go lock just
// rubber-stamps the order that already happened.
//
// Alternatives considered.
//
//   - Always-taken vs conditional (per-instance) locking. A
//     conditional lock would skip the Lock/Unlock pair on
//     non-shared-cache connections, but the binding has no signal
//     at Pool.Create time to know which scope a connection belongs
//     to (the PageCache.Create signature passes pageSize/extraSize/
//     purgeable, not a cache-shared bit, and extending it would
//     break the sqlite.PageCache contract from !126). The
//     uncontended cost was not worth the API churn.
//
//   - Document-as-unsupported + reject shared-cache parent at
//     Create. Smaller surface but a pcache user who registers the
//     pool and then opens a cache=shared db would get a confusing
//     post-registration failure. Worse, the impl is already correct
//     under cache=shared by SQLite's own serialisation, so refusing
//     to serve the workload would be conservatism without payoff.
//
// Test surface. e2e_test.go carries TestSharedCacheTwoConns_Integrity,
// the canonical two-conn shared-cache workload (concurrent writers +
// PRAGMA integrity_check) under -race. It passes on the always-locked
// pool and fails (race detector report) if the per-cache mutex is
// removed.
