// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"
)

// dbStatusConn pins a single connection to a fresh on-disk database and
// returns it along with a helper that runs op through the DBStatus escape
// hatch on that exact connection.
func dbStatusConn(t *testing.T) (*sql.Conn, func(op DBStatusOp, reset bool) (cur, high int)) {
	t.Helper()
	name := filepath.Join(t.TempDir(), "dbstatus.db")
	db, err := sql.Open(driverName, fmt.Sprintf("file:%s", name))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })

	conn, err := db.Conn(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() })

	status := func(op DBStatusOp, reset bool) (cur, high int) {
		t.Helper()
		if err := conn.Raw(func(driverConn any) error {
			s, ok := driverConn.(DBStatus)
			if !ok {
				return fmt.Errorf("driver connection didn't implement DBStatus")
			}
			c, h, err := s.Status(op, reset)
			cur, high = c, h
			return err
		}); err != nil {
			t.Fatalf("Status(%d, %v): %v", op, reset, err)
		}
		return cur, high
	}
	return conn, status
}

// TestDBStatusSchemaUsedGrows exercises a memory high-water op: schema memory
// in use rises after new tables are added to the connection's schema.
func TestDBStatusSchemaUsedGrows(t *testing.T) {
	conn, status := dbStatusConn(t)

	before, _ := status(DBStatusSchemaUsed, false)
	for i := 0; i < 8; i++ {
		stmt := fmt.Sprintf("create table t%d (a int, b text, c blob, d real)", i)
		if _, err := conn.ExecContext(context.TODO(), stmt); err != nil {
			t.Fatal(err)
		}
	}
	after, _ := status(DBStatusSchemaUsed, false)

	if after <= before {
		t.Errorf("SchemaUsed did not grow after adding tables: before=%d after=%d", before, after)
	}
}

// TestDBStatusCacheHitReset exercises a running-counter op: cache hits
// accumulate as pages are re-read, and the reset flag zeroes the counter so a
// subsequent read sees a fresh, smaller count.
func TestDBStatusCacheHitReset(t *testing.T) {
	conn, status := dbStatusConn(t)

	if _, err := conn.ExecContext(context.TODO(),
		`create table t (k integer primary key, v blob)`); err != nil {
		t.Fatal(err)
	}
	blob := make([]byte, 256)
	for i := 0; i < 500; i++ {
		if _, err := conn.ExecContext(context.TODO(),
			`insert into t(v) values (?)`, blob); err != nil {
			t.Fatalf("insert[%d]: %v", i, err)
		}
	}
	// Re-read the same pages several times so the page cache records hits.
	for i := 0; i < 5; i++ {
		var n int
		if err := conn.QueryRowContext(context.TODO(),
			`select count(*) from t`).Scan(&n); err != nil {
			t.Fatalf("count[%d]: %v", i, err)
		}
	}

	// Read-and-reset, then read again with no intervening SQL. The counter is
	// a cumulative total with a 0 high-water; the reset zeroes current, so the
	// immediate re-read must be strictly smaller than the value we reset.
	hit, high := status(DBStatusCacheHit, true)
	if hit <= 0 {
		t.Fatalf("CacheHit current = %d after a re-reading workload, want > 0", hit)
	}
	if high != 0 {
		t.Errorf("CacheHit high = %d, want 0 (running-counter ops report 0 high-water)", high)
	}
	afterReset, _ := status(DBStatusCacheHit, false)
	if afterReset >= hit {
		t.Errorf("CacheHit after reset = %d, want < pre-reset value %d", afterReset, hit)
	}
}

// TestDBStatusLookasideHitInHighwater pins the documented quirk that the
// lookaside event ops report their count in the high-water out-param, not
// current.
func TestDBStatusLookasideHitInHighwater(t *testing.T) {
	conn, status := dbStatusConn(t)

	if _, err := conn.ExecContext(context.TODO(),
		`create table t (a int, b int, c int)`); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 50; i++ {
		if _, err := conn.ExecContext(context.TODO(),
			`insert into t values (?, ?, ?)`, i, i*2, i*3); err != nil {
			t.Fatalf("insert[%d]: %v", i, err)
		}
	}

	cur, high := status(DBStatusLookasideHit, false)
	if cur != 0 {
		t.Errorf("LookasideHit current = %d, want 0 (the value is reported in high-water)", cur)
	}
	if high <= 0 {
		t.Errorf("LookasideHit high = %d, want > 0 after a statement workload", high)
	}
}

// TestDBStatusUnknownOpErrors confirms an out-of-range op surfaces an error
// rather than silently returning zero values.
func TestDBStatusUnknownOpErrors(t *testing.T) {
	conn, _ := dbStatusConn(t)

	err := conn.Raw(func(driverConn any) error {
		s := driverConn.(DBStatus)
		_, _, e := s.Status(DBStatusOp(9999), false)
		return e
	})
	if err == nil {
		t.Fatal("expected an error from an out-of-range DBStatusOp, got nil")
	}
}
