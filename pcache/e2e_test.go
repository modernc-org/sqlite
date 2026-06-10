// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcache_test

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"modernc.org/sqlite"
	"modernc.org/sqlite/pcache"
)

// poolUnderTest is the singleton Pool installed in TestMain. The
// SQLITE_CONFIG_PCACHE2 wiring is process-global, so every test in
// this binary observes its counters cumulatively. Tests that read
// Stats() snapshot the baseline first and compare deltas.
var poolUnderTest *pcache.Pool

func TestMain(m *testing.M) {
	poolUnderTest = pcache.New()
	if err := sqlite.RegisterPageCache(poolUnderTest); err != nil {
		fmt.Fprintf(os.Stderr, "RegisterPageCache: %v\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

// openTestDB opens a fresh on-disk SQLite database in the test's
// temporary directory. On-disk is required for purgeable=true to take
// effect; in-memory caches are non-purgeable and SQLite only calls
// Unpin with discard=true on them, which would bypass the LRU codepath
// the e2e test is exercising.
func openTestDB(t *testing.T) *sql.DB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "e2e.db")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	if _, err := db.Exec(`PRAGMA journal_mode=WAL`); err != nil {
		t.Fatalf("journal_mode: %v", err)
	}
	return db
}

// TestPoolRoundTripIntegrity is the headline e2e assertion: a real
// SQLite workload with a small bounded cache and a churn pattern
// (insert, delete, incremental_vacuum) round-trips through the pool
// and still passes integrity_check. This is the workload cznic ran
// against the binding during the MR !126 review; reproducing it here
// pins the Pool implementation against the same evidence.
func TestPoolRoundTripIntegrity(t *testing.T) {
	db := openTestDB(t)

	// PRAGMA cache_size is a per-connection setting. Take a single
	// connection out of the pool so the cache_size we set applies to
	// every Exec we issue.
	conn, err := db.Conn(t.Context())
	if err != nil {
		t.Fatalf("Conn: %v", err)
	}
	defer conn.Close()

	if _, err := conn.ExecContext(t.Context(), `PRAGMA cache_size=16`); err != nil {
		t.Fatalf("cache_size: %v", err)
	}
	if _, err := conn.ExecContext(t.Context(), `PRAGMA auto_vacuum=incremental`); err != nil {
		t.Fatalf("auto_vacuum: %v", err)
	}
	if _, err := conn.ExecContext(t.Context(), `CREATE TABLE t (k INTEGER PRIMARY KEY, v BLOB)`); err != nil {
		t.Fatalf("CREATE TABLE: %v", err)
	}

	baseline := poolUnderTest.Stats()

	// Push enough rows through a small cache that the LRU has to
	// recycle pages many times over. 4000 rows × ~256 B BLOB exceeds
	// 16 cache pages × 4096 B by ~16x.
	const N = 4000
	tx, err := conn.BeginTx(t.Context(), nil)
	if err != nil {
		t.Fatalf("BeginTx: %v", err)
	}
	blob := make([]byte, 256)
	for i := 0; i < N; i++ {
		blob[0] = byte(i)
		if _, err := tx.ExecContext(t.Context(), `INSERT INTO t(v) VALUES (?)`, blob); err != nil {
			t.Fatalf("INSERT[%d]: %v", i, err)
		}
	}
	if err := tx.Commit(); err != nil {
		t.Fatalf("Commit insert: %v", err)
	}

	// Delete half the rows and run incremental_vacuum so the b-tree
	// rearranges. This was the workload that previously exercised the
	// xRekey path 15 times on cznic's harness.
	if _, err := conn.ExecContext(t.Context(), `DELETE FROM t WHERE k % 2 = 0`); err != nil {
		t.Fatalf("DELETE: %v", err)
	}
	if _, err := conn.ExecContext(t.Context(), `PRAGMA incremental_vacuum`); err != nil {
		t.Fatalf("incremental_vacuum: %v", err)
	}

	// Read half the surviving rows back, mixing pin/unpin cycles.
	rows, err := conn.QueryContext(t.Context(), `SELECT k, v FROM t WHERE k <= ? ORDER BY k`, N/2)
	if err != nil {
		t.Fatalf("SELECT: %v", err)
	}
	count := 0
	for rows.Next() {
		var k int
		var v []byte
		if err := rows.Scan(&k, &v); err != nil {
			t.Fatalf("Scan: %v", err)
		}
		count++
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows.Err: %v", err)
	}
	rows.Close()

	// Integrity check.
	var integrity string
	if err := conn.QueryRowContext(t.Context(), `PRAGMA integrity_check`).Scan(&integrity); err != nil {
		t.Fatalf("integrity_check: %v", err)
	}
	if integrity != "ok" {
		t.Fatalf("integrity_check = %q, want \"ok\"", integrity)
	}

	delta := statsDelta(baseline, poolUnderTest.Stats())
	t.Logf("workload delta: %+v", delta)

	// The workload churns ~%dx the bounded cache size so evictions
	// MUST have fired; if they didn't, the binding is leaking stubs
	// or the impl is silently overcommitting.
	if delta.Evictions == 0 {
		t.Errorf("Stats.Evictions delta = 0, want > 0 (bounded cache must have evicted)")
	}
	if delta.Allocs == 0 {
		t.Errorf("Stats.Allocs delta = 0, want > 0 (workload should have allocated pages)")
	}
	if delta.Caches == 0 {
		t.Errorf("Stats.Caches delta = 0, want >= 1 (DB open should create at least one cache)")
	}
}

// TestPoolMultipleDatabases opens several distinct on-disk databases
// through the same Pool. Each sql.Open issues its own physical
// connection, which fires PageCache.Create against the Pool. The test
// asserts that the per-cache Stats aggregation reports every fresh
// Cache the Pool was asked to mint.
//
// A single sql.DB cannot drive this assertion because database/sql
// pools and reuses connections across Conn() calls; only a fresh
// sql.Open guarantees a fresh underlying SQLite handle and so a
// fresh PageCache.Create.
func TestPoolMultipleDatabases(t *testing.T) {
	baseline := poolUnderTest.Stats()

	const dbs = 4
	for i := 0; i < dbs; i++ {
		path := filepath.Join(t.TempDir(), fmt.Sprintf("multi-%d.db", i))
		db, err := sql.Open("sqlite", path)
		if err != nil {
			t.Fatalf("Open[%d]: %v", i, err)
		}
		if _, err := db.Exec(`PRAGMA cache_size=8`); err != nil {
			db.Close()
			t.Fatalf("cache_size[%d]: %v", i, err)
		}
		if _, err := db.Exec(`CREATE TABLE m (k INTEGER PRIMARY KEY, v INT)`); err != nil {
			db.Close()
			t.Fatalf("CREATE TABLE[%d]: %v", i, err)
		}
		for j := 0; j < 100; j++ {
			if _, err := db.Exec(`INSERT INTO m(v) VALUES (?)`, i*1000+j); err != nil {
				db.Close()
				t.Fatalf("INSERT[%d/%d]: %v", i, j, err)
			}
		}
		if err := db.Close(); err != nil {
			t.Fatalf("Close[%d]: %v", i, err)
		}
	}

	delta := statsDelta(baseline, poolUnderTest.Stats())
	t.Logf("multi-db delta: %+v", delta)
	if delta.Caches < int64(dbs) {
		t.Errorf("Stats.Caches delta = %d, want >= %d (one Create per opened database)",
			delta.Caches, dbs)
	}
}

// statsDelta returns the field-wise difference between two Stats
// snapshots. Counters are monotonic, so every field is >= 0.
func statsDelta(before, after pcache.Stats) pcache.Stats {
	return pcache.Stats{
		Hits:      after.Hits - before.Hits,
		Misses:    after.Misses - before.Misses,
		Allocs:    after.Allocs - before.Allocs,
		Evictions: after.Evictions - before.Evictions,
		Rekeys:    after.Rekeys - before.Rekeys,
		Truncates: after.Truncates - before.Truncates,
		Caches:    after.Caches - before.Caches,
	}
}
