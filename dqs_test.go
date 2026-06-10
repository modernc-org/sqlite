// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"database/sql"
	"strings"
	"testing"

	sqlite3 "modernc.org/sqlite/lib"
)

// TestDQSConfigCallVaList pins the mixed-vararg FFI shape of
// sqlite3_db_config's (db, op, int onoff, int *pRes) form. The hand-laid
// VaList carries one int and one pointer; mixing widths is the easy thing
// to get subtly wrong in the cgo-free transpilation, so the test asserts
// the call returns SQLITE_OK for every boolean-toggle DQS op in both
// directions.
func TestDQSConfigCallVaList(t *testing.T) {
	c, err := newConn(":memory:")
	if err != nil {
		t.Fatalf("newConn: %v", err)
	}
	defer c.Close()

	for _, op := range []int32{
		sqlite3.SQLITE_DBCONFIG_DQS_DDL,
		sqlite3.SQLITE_DBCONFIG_DQS_DML,
	} {
		if rc := c.dbConfigBool(op, false); rc != sqlite3.SQLITE_OK {
			t.Errorf("dbConfigBool(op=%d, off) = %d, want SQLITE_OK", op, rc)
		}
		if rc := c.dbConfigBool(op, true); rc != sqlite3.SQLITE_OK {
			t.Errorf("dbConfigBool(op=%d, on) = %d, want SQLITE_OK", op, rc)
		}
	}
}

// TestDQSOptIn exercises the _dqs DSN parameter end-to-end through
// database/sql. With DQS enabled (default or _dqs=1) a non-resolving
// double-quoted identifier silently falls back to a string literal; with
// _dqs=0 the same expression must fail to parse instead.
func TestDQSOptIn(t *testing.T) {
	cases := []struct {
		name      string
		dsn       string
		wantError bool
	}{
		{"default", "file::memory:?cache=shared&mode=memory", false},
		{"explicit_on", "file::memory:?cache=shared&mode=memory&_dqs=1", false},
		{"off", "file::memory:?cache=shared&mode=memory&_dqs=0", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			db, err := sql.Open("sqlite", tc.dsn)
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			defer db.Close()
			var got string
			err = db.QueryRow(`SELECT "double-quoted-literal"`).Scan(&got)
			switch {
			case tc.wantError && err == nil:
				t.Errorf("expected error for %q, got value %q", tc.dsn, got)
			case !tc.wantError && err != nil:
				t.Errorf("unexpected error for %q: %v", tc.dsn, err)
			case !tc.wantError && got != "double-quoted-literal":
				t.Errorf("got %q, want %q", got, "double-quoted-literal")
			}
		})
	}
}

// TestDQSInvalid verifies that an unparseable _dqs value surfaces an
// error from newConn rather than being silently ignored.
func TestDQSInvalid(t *testing.T) {
	_, err := newConn(":memory:?_dqs=not-a-bool")
	if err == nil {
		t.Fatal("expected error for invalid _dqs value")
	}
	if !strings.Contains(err.Error(), "_dqs") {
		t.Errorf("error %q should mention _dqs", err)
	}
}
