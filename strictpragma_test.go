// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"database/sql"
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"unsafe"

	"modernc.org/libc"
)

// singleStatementCases are _pragma values, without the PRAGMA keyword, and
// whether they are one statement.
var singleStatementCases = []struct {
	v      string
	single bool
}{
	{"foreign_keys(1)", true},
	{"foreign_keys = 1", true},
	{"foreign_keys(1);", true},
	{"foreign_keys(1);;;", true},
	{"foreign_keys(1); \t\r\n\f", true},
	{"foreign_keys(1); -- trailing comment", true},
	{"foreign_keys(1); /* trailing comment */", true},
	{"foreign_keys(1); /* unterminated comment", true},
	{"foreign_keys(1) -- ; select 1", true},
	{"foreign_keys(1) /* ; select 1 */", true},
	{"application_id = 'a;b'", true},
	{"application_id = 'it''s; fine'", true},
	{`application_id = "a;b"`, true},
	{`application_id = "a"";b"`, true},
	{"application_id = `a;b`", true},
	{"application_id = `a``;b`", true},
	{"application_id = [a;b]", true},
	{"application_id = 'unterminated; select 1", true},
	{"", true},

	{"foreign_keys(1);select 1", false},
	{"foreign_keys(1); select 1", false},
	{"foreign_keys(1);\nselect 1", false},
	{"foreign_keys(1); -- comment\nselect 1", false},
	{"foreign_keys(1); /* comment */ select 1", false},
	{"foreign_keys(1);ATTACH 'x.db' AS x", false},
	{"application_id = 'a';select 1", false},
	{"application_id = 'a'';b';select 1", false},
	{"application_id = [a]]; select 1", false},
	{"foreign_keys(1);;select 1", false},
	{"foreign_keys(1); /", false},
	{"foreign_keys(1); -", false},
	{"foreign_keys(1);\v", false},
	{"foreign_keys(1);/*", false},
	{"foreign_keys(1); /*", false},
	{"foreign_keys(1);/**", true},
	{"foreign_keys(1);/*\x00 comment", false},
	{"foreign_keys(1)\x00;select 1", true},
}

func TestSingleStatement(t *testing.T) {
	for _, tc := range singleStatementCases {
		if g, e := singleStatement(tc.v), tc.single; g != e {
			t.Errorf("singleStatement(%q) = %v, want %v", tc.v, g, e)
		}
	}

	// SQLite accepts a \v inside a run of other whitespace; singleStatement
	// rejects it everywhere, deliberately. It is the one listed disagreement
	// with SQLite, so it is kept out of singleStatementCases.
	if singleStatement("foreign_keys(1); \v") {
		t.Error("a \\v after the statement was accepted")
	}
}

// sqliteTailIsEmpty reports what SQLite's own parser says about
// "pragma "+v: whether it prepares its first statement, and whether
// everything after it prepares to no statement at all.
func sqliteTailIsEmpty(t testing.TB, c *conn, v string) (firstOK, tailEmpty bool) {
	p, err := libc.CString("pragma " + v)
	if err != nil {
		t.Fatal(err)
	}

	defer c.free(p)

	psql := p
	pstmt, err := c.prepareV2(&psql)
	if err != nil || pstmt == 0 {
		return false, false
	}

	c.finalize(pstmt)
	for *(*byte)(unsafe.Pointer(psql)) != 0 {
		prev := psql
		pstmt, err := c.prepareV2(&psql)
		if err != nil {
			return true, false
		}

		if pstmt != 0 {
			c.finalize(pstmt)
			return true, false
		}

		if psql == prev {
			break
		}
	}
	return true, true
}

// TestSingleStatementAgainstSQLite holds the lexical check to SQLite's own
// parser: nothing it accepts may leave SQLite a second statement to run.
func TestSingleStatementAgainstSQLite(t *testing.T) {
	c, err := newConn(":memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

	for _, tc := range singleStatementCases {
		firstOK, tailEmpty := sqliteTailIsEmpty(t, c, tc.v)
		if !firstOK {
			continue
		}

		if singleStatement(tc.v) && !tailEmpty {
			t.Errorf("%q: accepted as one statement, but SQLite finds more", tc.v)
		}

		if !tc.single && tailEmpty {
			t.Errorf("%q: listed as more than one statement, but SQLite finds one", tc.v)
		}
	}
}

func FuzzSingleStatement(f *testing.F) {
	for _, tc := range singleStatementCases {
		f.Add(tc.v)
	}
	c, err := newConn(":memory:")
	if err != nil {
		f.Fatal(err)
	}

	defer c.Close()

	f.Fuzz(func(t *testing.T, v string) {
		if !singleStatement(v) {
			return
		}

		if firstOK, tailEmpty := sqliteTailIsEmpty(t, c, v); firstOK && !tailEmpty {
			t.Fatalf("%q: accepted as one statement, but SQLite finds more", v)
		}
	})
}

func TestStrictPragmas(t *testing.T) {
	if StrictPragmasEnabled() {
		t.Fatal("StrictPragmas is on by default")
	}

	open := func(t *testing.T, pragma string) (dir string, err error) {
		dir = t.TempDir()
		v := url.Values{"_pragma": {pragma}}
		db, err := sql.Open("sqlite", filepath.Join(dir, "a.db")+"?"+v.Encode())
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		return dir, db.Ping()
	}

	attach := func(dir string) string {
		return "foreign_keys(1);ATTACH '" + filepath.Join(dir, "x.db") + "' AS x"
	}

	t.Run("off", func(t *testing.T) {
		dir := t.TempDir()
		v := url.Values{"_pragma": {attach(dir)}}
		db, err := sql.Open("sqlite", filepath.Join(dir, "a.db")+"?"+v.Encode())
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		if err := db.Ping(); err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(filepath.Join(dir, "x.db")); err != nil {
			t.Fatalf("without StrictPragmas the ATTACH should have run: %v", err)
		}
	})

	if prev := StrictPragmas(true); prev {
		t.Fatal("StrictPragmas(true) reported it was already on")
	}

	defer StrictPragmas(false)

	if !StrictPragmasEnabled() {
		t.Fatal("StrictPragmasEnabled = false after StrictPragmas(true)")
	}

	t.Run("rejected", func(t *testing.T) {
		dir := t.TempDir()
		v := url.Values{"_pragma": {"user_version=7", attach(dir)}}
		db, err := sql.Open("sqlite", filepath.Join(dir, "a.db")+"?"+v.Encode())
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		if err := db.Ping(); !errors.Is(err, ErrMultiStatementPragma) {
			t.Fatalf("got %v, want ErrMultiStatementPragma", err)
		}

		if _, err := os.Stat(filepath.Join(dir, "x.db")); err == nil {
			t.Fatal("x.db was created")
		}

		// Rejected in the validation phase: the valid _pragma before it
		// was not applied either.
		db2, err := sql.Open("sqlite", filepath.Join(dir, "a.db"))
		if err != nil {
			t.Fatal(err)
		}

		defer db2.Close()

		var uv int
		if err := db2.QueryRow("pragma user_version").Scan(&uv); err != nil {
			t.Fatal(err)
		}

		if uv != 0 {
			t.Fatalf("user_version = %d, want 0", uv)
		}
	})

	t.Run("accepted", func(t *testing.T) {
		for _, p := range []string{"foreign_keys(1)", "foreign_keys(1);", "foreign_keys(1); -- note"} {
			if _, err := open(t, p); err != nil {
				t.Errorf("_pragma=%q: %v", p, err)
			}
		}
	})
}
