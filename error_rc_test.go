// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	sqlite3 "modernc.org/sqlite/lib"
)

// TestErrorRcOpenTimeUnopenable pins the per-#230 invariants for an
// open-time failure exercised across both modes. Whether the legacy
// formatter appends the temporary handle's errmsg via ": " depends on
// the host's specific failure path (the original #230 reproducer used
// "/sys/repro.db" on Linux where the temp handle carried a stale
// "out of memory" errmsg; other platforms may converge on a canonical
// errmsg already). The structural assertions hold on every platform:
//
//   - Code() returns SQLITE_CANTOPEN unchanged between modes per cznic's
//     review refinement that only the human-readable message changes.
//   - The canonical errstr is present in both modes.
//   - In _error_rc=1 mode, no spurious "out of memory" string is
//     appended even on platforms where the legacy form would have
//     surfaced it.
func TestErrorRcOpenTimeUnopenable(t *testing.T) {
	// Build a path under an existing-but-not-a-directory parent so
	// open_v2 fails with SQLITE_CANTOPEN on every platform without
	// requiring read-only mount permissions.
	dir := t.TempDir()
	notDir := filepath.Join(dir, "notdir")
	if err := mustCreateFile(notDir); err != nil {
		t.Fatalf("setup: %v", err)
	}
	bad := filepath.Join(notDir, "child", "x.db")

	for _, mode := range []string{"legacy_default", "explicit_off", "error_rc_on"} {
		t.Run(mode, func(t *testing.T) {
			dsn := bad
			switch mode {
			case "explicit_off":
				dsn += "?_error_rc=0"
			case "error_rc_on":
				dsn += "?_error_rc=1"
			}
			db, err := sql.Open("sqlite", dsn)
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			defer db.Close()
			err = db.Ping()
			if err == nil {
				t.Fatal("expected error pinging an unopenable database")
			}
			var sqlErr *Error
			if !errors.As(err, &sqlErr) {
				t.Fatalf("expected *sqlite.Error, got %T (%v)", err, err)
			}
			// Cznic refinement #1: Code() is unchanged across modes.
			if sqlErr.Code() != sqlite3.SQLITE_CANTOPEN {
				t.Errorf("Code() = %d, want SQLITE_CANTOPEN (%d)",
					sqlErr.Code(), sqlite3.SQLITE_CANTOPEN)
			}
			if !strings.Contains(sqlErr.Error(), "unable to open database file") {
				t.Errorf("error %q missing canonical errstr", sqlErr)
			}
			if mode == "error_rc_on" &&
				strings.Contains(strings.ToLower(sqlErr.Error()), "out of memory") {
				t.Errorf("_error_rc=1 mode should not surface stale errmsg, got %q", sqlErr)
			}
		})
	}
}

// TestErrorRcSyntaxErrorPreservesErrmsg ensures the _error_rc=1 mode does
// NOT regress the helpful cases: when sqlite3_extended_errcode(db) is
// consistent with the operation rc (the common case for syntax and
// constraint errors), the detailed errmsg is preserved.
func TestErrorRcSyntaxErrorPreservesErrmsg(t *testing.T) {
	for _, mode := range []string{"", "_error_rc=0", "_error_rc=1"} {
		t.Run(mode, func(t *testing.T) {
			dsn := "file::memory:?cache=shared&mode=memory"
			if mode != "" {
				dsn += "&" + mode
			}
			db, err := sql.Open("sqlite", dsn)
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			defer db.Close()
			_, err = db.Exec("SELECT not-a-column FROM not-a-table")
			if err == nil {
				t.Fatal("expected syntax error")
			}
			// The helpful detail ("no such table" or similar) should
			// be present in all modes because the extended_errcode
			// on the db handle matches rc for this path.
			low := strings.ToLower(err.Error())
			if !strings.Contains(low, "no such table") &&
				!strings.Contains(low, "syntax error") {
				t.Errorf("mode=%q: expected helpful detail, got %q", mode, err)
			}
		})
	}
}

// TestErrstrForDBSuppressOnMismatch pins the _error_rc=1 conditional
// suppress branch deterministically by calling errstrForDB with a
// healthy db handle (sqlite3_extended_errcode = SQLITE_OK,
// sqlite3_errmsg = "not an error") and a deliberately mismatched
// rc = SQLITE_CANTOPEN, so the new behavior is verified without
// relying on a host-specific open-time failure path. In legacy mode
// the formatter appends the handle's stale "not an error" as the
// helpful detail; in _error_rc=1 mode the conditional suppress
// branch fires and the canonical errstr(rc) is used alone.
// Code() returns SQLITE_CANTOPEN in both modes.
func TestErrstrForDBSuppressOnMismatch(t *testing.T) {
	c, err := newConn(":memory:")
	if err != nil {
		t.Fatalf("newConn: %v", err)
	}
	defer c.Close()

	const mismatchedRc = int32(sqlite3.SQLITE_CANTOPEN)

	cases := []struct {
		name        string
		errorRcMode bool
		mustHave    []string
		mustNotHave []string
	}{
		{
			name:        "legacy_mode_appends_errmsg",
			errorRcMode: false,
			mustHave:    []string{"unable to open database file", "not an error"},
			mustNotHave: nil,
		},
		{
			name:        "error_rc_mode_suppresses_errmsg",
			errorRcMode: true,
			mustHave:    []string{"unable to open database file"},
			mustNotHave: []string{"not an error"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := errstrForDB(c.tls, mismatchedRc, c.db, tc.errorRcMode)
			var sqlErr *Error
			if !errors.As(got, &sqlErr) {
				t.Fatalf("expected *Error, got %T (%v)", got, got)
			}
			if sqlErr.Code() != sqlite3.SQLITE_CANTOPEN {
				t.Errorf("Code() = %d, want SQLITE_CANTOPEN (%d)",
					sqlErr.Code(), sqlite3.SQLITE_CANTOPEN)
			}
			low := strings.ToLower(sqlErr.Error())
			for _, s := range tc.mustHave {
				if !strings.Contains(low, s) {
					t.Errorf("error %q missing %q", sqlErr, s)
				}
			}
			for _, s := range tc.mustNotHave {
				if strings.Contains(low, s) {
					t.Errorf("error %q must not contain %q", sqlErr, s)
				}
			}
		})
	}
}

// TestErrorRcInvalidValue surfaces an unparseable _error_rc value at
// newConn rather than silently ignoring it.
func TestErrorRcInvalidValue(t *testing.T) {
	_, err := newConn(":memory:?_error_rc=not-a-bool")
	if err == nil {
		t.Fatal("expected error for invalid _error_rc value")
	}
	if !strings.Contains(err.Error(), "_error_rc") {
		t.Errorf("error %q should mention _error_rc", err)
	}
}

// mustCreateFile is a tiny helper that lets the test stage a regular
// file at a path that the connector will then try to descend into as
// if it were a directory.
func mustCreateFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return f.Close()
}
