// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcache_test

import (
	"database/sql"
	"path/filepath"
	"runtime"
	"testing"

	"modernc.org/sqlite/pcache"
)

// BenchmarkPoolBoundedCache exercises the workload from the SQLite
// issue #204 conversation: a small bounded cache (cache_size in pages)
// driven by a write/read mix that exceeds it many times over. The
// reportable numbers are
//
//   - allocs/op + bytes/op from the Go-side runtime, which are the
//     wrapper's footprint per insert,
//   - HeapInuse before/after, which captures whether the wrapper or
//     the Pool itself accumulates Go-heap garbage proportional to the
//     workload size,
//   - easy-refusals/op, the number of FetchCreateEasy refusals at cap
//     per insert. SQLite handles a refusal by spilling dirty pages and
//     retrying with FetchCreateForce, so this is a direct proxy for the
//     I/O pressure the strict Easy contract adds vs pcache1's
//     recycle-without-spill behavior (raised by cznic in the !127
//     review),
//   - the Pool.Stats counters at the end, which show that the bounded
//     cache stayed bounded (Allocs and Evictions are within an order
//     of magnitude of each other, never one growing without the other).
//
// To produce a side-by-side memory comparison vs the default pcache1,
// run the same benchmark in a sibling working tree that does not
// import this package; the parent binary then falls back to the
// in-engine pcache1. The MR description carries the comparison
// numbers; this benchmark is the reproducer.
func BenchmarkPoolBoundedCache(b *testing.B) {
	path := filepath.Join(b.TempDir(), "bench.db")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		b.Fatalf("Open: %v", err)
	}
	defer db.Close()
	if _, err := db.Exec(`PRAGMA cache_size=16`); err != nil {
		b.Fatalf("cache_size: %v", err)
	}
	if _, err := db.Exec(`PRAGMA journal_mode=WAL`); err != nil {
		b.Fatalf("journal_mode: %v", err)
	}
	if _, err := db.Exec(`CREATE TABLE t (k INTEGER PRIMARY KEY, v BLOB)`); err != nil {
		b.Fatalf("CREATE TABLE: %v", err)
	}

	blob := make([]byte, 256)
	stmt, err := db.Prepare(`INSERT INTO t(v) VALUES (?)`)
	if err != nil {
		b.Fatalf("Prepare: %v", err)
	}
	defer stmt.Close()

	runtime.GC()
	var memBefore runtime.MemStats
	runtime.ReadMemStats(&memBefore)
	baseline := poolUnderTest.Stats()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blob[0] = byte(i)
		if _, err := stmt.Exec(blob); err != nil {
			b.Fatalf("INSERT[%d]: %v", i, err)
		}
	}
	b.StopTimer()

	runtime.GC()
	var memAfter runtime.MemStats
	runtime.ReadMemStats(&memAfter)
	delta := statsDelta(baseline, poolUnderTest.Stats())

	b.ReportMetric(float64(delta.Allocs)/float64(b.N), "page-allocs/op")
	b.ReportMetric(float64(delta.Evictions)/float64(b.N), "page-evictions/op")
	b.ReportMetric(float64(delta.EasyRefusals)/float64(b.N), "easy-refusals/op")
	b.ReportMetric(float64(memAfter.HeapInuse-memBefore.HeapInuse), "go-heap-inuse-delta-bytes")
	b.Logf("pool delta over %d inserts: %+v", b.N, delta)
}

// BenchmarkPoolEvictionChurn drives a rotating-residue
// delete + vacuum + matching-batch insert cycle against a small
// bounded cache so the LRU has to evict many times the cache_size
// over the run. Cycle i deletes rows where k % 3 == i % 3, then
// re-inserts a small batch of rows whose keys map back to that
// residue so the next visit (three cycles later) finds rows to
// scan. Without the rotation + re-insert the workload converges to
// a no-op after the seed's k % 3 = 0 partition is drained, and
// easy-refusals/op collapses to a fixed first-cycle cost divided
// by b.N. The reportable evictions/op and easy-refusals/op are the
// steady-state churn the binding handles per SQL statement under
// cache pressure. The rotating-residue + matching-batch INSERT
// pattern also reliably triggers the b-tree rebalance paths that
// emit xRekey through the SQL surface (~13 Rekeys per cycle,
// scaling linearly with b.N); the benchmark therefore complements
// the dedicated xRekey unit tests (TestRekey, TestRekeyEvictsCollider)
// rather than deferring to them.
func BenchmarkPoolEvictionChurn(b *testing.B) {
	path := filepath.Join(b.TempDir(), "churn.db")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		b.Fatalf("Open: %v", err)
	}
	defer db.Close()
	if _, err := db.Exec(`PRAGMA cache_size=16; PRAGMA auto_vacuum=incremental`); err != nil {
		b.Fatalf("pragmas: %v", err)
	}
	if _, err := db.Exec(`CREATE TABLE r (k INTEGER PRIMARY KEY, v BLOB)`); err != nil {
		b.Fatalf("CREATE TABLE: %v", err)
	}

	const batchPerCycle = 200 // rows re-inserted into the just-cleared residue each cycle

	// Seed three batches of batchPerCycle rows, one per residue class
	// (k % 3 in {0, 1, 2}), so the first three cycles of the
	// rotating-residue DELETE each remove a same-size partition. From
	// cycle 0 forward, every cycle removes batchPerCycle rows and
	// re-inserts batchPerCycle rows, keeping per-cycle work constant
	// and the reported per-op metrics independent of b.N. Without
	// the constant per-cycle work, the seed's k % 3 = 0 partition
	// would drain on cycle 0 and easy-refusals/op would collapse to
	// a fixed first-cycle cost divided by b.N.
	blob := make([]byte, 256)
	var seedKey int64
	for residue := int64(0); residue < 3; residue++ {
		for j := 0; j < batchPerCycle; j++ {
			seedKey++
			for seedKey%3 != residue {
				seedKey++
			}
			if _, err := db.Exec(`INSERT INTO r(k, v) VALUES (?, ?)`, seedKey, blob); err != nil {
				b.Fatalf("seed[%d,%d]: %v", residue, j, err)
			}
		}
	}

	baseline := poolUnderTest.Stats()
	nextKey := seedKey + 1
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		residue := int64(i % 3)
		if _, err := db.Exec(`DELETE FROM r WHERE k % 3 = ?`, residue); err != nil {
			b.Fatalf("DELETE[%d]: %v", i, err)
		}
		if _, err := db.Exec(`PRAGMA incremental_vacuum`); err != nil {
			b.Fatalf("vacuum[%d]: %v", i, err)
		}
		// Re-insert batchPerCycle rows whose keys land in the
		// just-cleared residue, in a fresh disjoint range. The next
		// time this residue rotates around (three cycles later) the
		// DELETE finds these rows and the spill pressure recurs,
		// keeping EasyRefusals scaling with b.N rather than capping
		// at the seed's one-time first-cycle cost.
		for j := 0; j < batchPerCycle; j++ {
			for nextKey%3 != residue {
				nextKey++
			}
			if _, err := db.Exec(`INSERT INTO r(k, v) VALUES (?, ?)`, nextKey, blob); err != nil {
				b.Fatalf("INSERT[%d,%d]: %v", i, j, err)
			}
			nextKey += 3
		}
	}
	b.StopTimer()
	delta := statsDelta(baseline, poolUnderTest.Stats())
	b.ReportMetric(float64(delta.Allocs)/float64(b.N), "page-allocs/op")
	b.ReportMetric(float64(delta.Evictions)/float64(b.N), "page-evictions/op")
	b.ReportMetric(float64(delta.EasyRefusals)/float64(b.N), "easy-refusals/op")
	b.ReportMetric(float64(delta.Truncates)/float64(b.N), "truncates/op")
	b.Logf("eviction-churn delta over %d cycles: %+v", b.N, delta)
}

// statsDelta is defined in e2e_test.go; same package.

var _ = pcache.Stats{}
