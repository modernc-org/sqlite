// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcache_test

import (
	"context"
	"database/sql"
	"errors"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"modernc.org/sqlite"
)

// TestSharedCacheTwoConns_Integrity is the canonical shared-cache test
// named in the MR-131 (pcache-shared-cache-draft) scaffold: two sql.Conn
// share one cache=shared database, both write to it concurrently, and
// the database must pass PRAGMA integrity_check at the end. It runs
// against the STOCK, unlocked pcache impl.
//
// Findings (see README_PROBE.md for the full reproduction):
//
//   - WITHOUT the race detector it is FUNCTIONALLY CLEAN: integrity_check
//     is always "ok" and it survived 20 back-to-back runs with no
//     corruption and no "concurrent map writes" fatal. SQLite's BtShared
//     mutex serialises every pcache2 callback, so the shared *pCache is
//     shared across the two connections but never entered concurrently.
//
//   - UNDER `go test -race` it reports a data race on a shared page's
//     pinned/lruElem fields (cache.Fetch reading vs cache.Unpin writing,
//     during one connection's BEGIN vs the other's COMMIT of page 1).
//     That report is a FALSE POSITIVE: the companion probe
//     (shared_cache_probe_test.go) measures max-in-flight = 1, i.e. the
//     two callbacks never actually overlap. The Go race detector simply
//     cannot model libc's pthread-mutex-backed BtShared serialisation as
//     a happens-before edge.
//
// So this test PASSES without -race and FAILS (spuriously) under -race.
//
// poolUnderTest is the Pool registered process-globally in TestMain (see
// e2e_test.go).
func TestSharedCacheTwoConns_Integrity(t *testing.T) {
	dsn := "file:" + filepath.Join(t.TempDir(), "shared.db") + "?cache=shared"
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)

	ctx := context.Background()

	// Two explicit connections, both bound to the one shared cache.
	c1, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer c1.Close()
	c2, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close()

	if _, err := c1.ExecContext(ctx,
		`CREATE TABLE t(k INTEGER PRIMARY KEY, v BLOB)`); err != nil {
		t.Fatal(err)
	}

	// Concurrent writers on the two connections. Shared-cache mode uses
	// table-level write locks, so the loser of a write race gets
	// SQLITE_LOCKED / SQLITE_BUSY; retry until it lands.
	const iters = 1500
	writer := func(conn *sql.Conn, base int) error {
		for i := 0; i < iters; i++ {
			for {
				_, err := conn.ExecContext(ctx,
					"INSERT OR REPLACE INTO t(k, v) VALUES(?, randomblob(80))", base+i)
				if err == nil {
					break
				}
				if isBusyOrLocked(err) {
					time.Sleep(50 * time.Microsecond)
					continue
				}
				return err
			}
		}
		return nil
	}

	var wg sync.WaitGroup
	errc := make(chan error, 2)
	wg.Add(2)
	go func() { defer wg.Done(); errc <- writer(c1, 0) }()
	go func() { defer wg.Done(); errc <- writer(c2, 1_000_000) }()
	wg.Wait()
	close(errc)
	for e := range errc {
		if e != nil {
			t.Fatal(e)
		}
	}

	var res string
	if err := c1.QueryRowContext(ctx, "PRAGMA integrity_check").Scan(&res); err != nil {
		t.Fatal(err)
	}
	if res != "ok" {
		t.Fatalf("integrity_check = %q, want \"ok\"", res)
	}

	st := poolUnderTest.Stats()
	t.Logf("shared-cache integrity ok: pool lifetime caches=%d hits=%d misses=%d allocs=%d evictions=%d",
		st.Caches, st.Hits, st.Misses, st.Allocs, st.Evictions)
}

// isBusyOrLocked reports whether err is a SQLITE_BUSY or SQLITE_LOCKED
// result (primary code, ignoring the extended shared-cache variants).
func isBusyOrLocked(err error) bool {
	var se *sqlite.Error
	if !errors.As(err, &se) {
		return false
	}
	switch se.Code() & 0xff {
	case 5, 6: // SQLITE_BUSY, SQLITE_LOCKED
		return true
	}
	return false
}
