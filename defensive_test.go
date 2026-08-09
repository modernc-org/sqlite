// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	sqlite3 "modernc.org/sqlite/lib"
)

func TestDefensiveConfigCallVaList(t *testing.T) {
	c, err := newConn(":memory:")
	if err != nil {
		t.Fatalf("newConn: %v", err)
	}
	defer c.Close()

	if rc := c.dbConfigBool(sqlite3.SQLITE_DBCONFIG_DEFENSIVE, true); rc != sqlite3.SQLITE_OK {
		t.Fatalf("enable defensive mode = %d, want SQLITE_OK", rc)
	}
	if rc := c.dbConfigBool(sqlite3.SQLITE_DBCONFIG_DEFENSIVE, false); rc != sqlite3.SQLITE_OK {
		t.Fatalf("disable defensive mode = %d, want SQLITE_OK", rc)
	}
}

func TestDefensiveDSNValidation(t *testing.T) {
	for _, tc := range []struct {
		name  string
		query string
		want  string
	}{
		{"invalid", "?_defensive=not-a-bool&_pragma=schema_version(41)", "invalid _defensive"},
		{"empty", "?_defensive=&_pragma=schema_version(41)", "invalid _defensive"},
		{"duplicate_same", "?_defensive=1&_defensive=1&_pragma=schema_version(41)", "exactly once"},
		{"duplicate_conflicting", "?_defensive=1&_defensive=0&_pragma=schema_version(41)", "exactly once"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dbPath := filepath.Join(t.TempDir(), "invalid.db")
			db, err := sql.Open("sqlite", dbPath+tc.query)
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			err = db.Ping()
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("Ping error = %v, want substring %q", err, tc.want)
			}
			if _, statErr := os.Stat(dbPath); !os.IsNotExist(statErr) {
				t.Fatalf("invalid DSN created database before rejection: stat error=%v", statErr)
			}
		})
	}
}

func TestDefensiveModeBlocksDangerousOperations(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "defensive.db")
	db, err := sql.Open("sqlite", dbPath+"?_defensive=1")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE protected(id INTEGER PRIMARY KEY)"); err != nil {
		t.Fatal(err)
	}

	var journalMode string
	if err := db.QueryRow("PRAGMA journal_mode=OFF").Scan(&journalMode); err != nil {
		t.Fatal(err)
	}
	if strings.EqualFold(journalMode, "off") {
		t.Fatalf("defensive connection entered forbidden journal_mode=%q", journalMode)
	}

	var before, after int64
	if err := db.QueryRow("PRAGMA schema_version").Scan(&before); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(fmt.Sprintf("PRAGMA schema_version=%d", before+41)); err != nil {
		t.Fatal(err)
	}
	if err := db.QueryRow("PRAGMA schema_version").Scan(&after); err != nil {
		t.Fatal(err)
	}
	if after != before {
		t.Fatalf("schema_version changed under defensive mode: before=%d after=%d", before, after)
	}

	if _, err := db.Exec("PRAGMA writable_schema=ON"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec("DELETE FROM sqlite_schema WHERE name='protected'"); err == nil {
		t.Fatal("direct sqlite_schema write unexpectedly succeeded under defensive mode")
	}
	var protectedCount int
	if err := db.QueryRow("SELECT count(*) FROM sqlite_schema WHERE type='table' AND name='protected'").Scan(&protectedCount); err != nil {
		t.Fatal(err)
	}
	if protectedCount != 1 {
		t.Fatalf("protected schema row changed after rejected write: count=%d", protectedCount)
	}
}

func TestDefensiveModeAppliesBeforeUserPragma(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "ordering.db")
	db, err := sql.Open("sqlite", dbPath+"?_defensive=1&_pragma=journal_mode(OFF)")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var journalMode string
	if err := db.QueryRow("PRAGMA journal_mode").Scan(&journalMode); err != nil {
		t.Fatal(err)
	}
	if strings.EqualFold(journalMode, "off") {
		t.Fatalf("user PRAGMA ran before defensive mode: journal_mode=%q", journalMode)
	}
}

func TestDefensiveModeAppliesToEveryPhysicalConnection(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "pool.db")
	db, err := sql.Open("sqlite", dbPath+"?_defensive=1")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(4)

	ctx := context.Background()
	connections := make([]*sql.Conn, 0, 4)
	defer func() {
		for _, c := range connections {
			_ = c.Close()
		}
	}()
	for i := 0; i < 4; i++ {
		c, err := db.Conn(ctx)
		if err != nil {
			t.Fatalf("Conn(%d): %v", i, err)
		}
		connections = append(connections, c)
	}
	for i, c := range connections {
		var before, after int64
		if err := c.QueryRowContext(ctx, "PRAGMA schema_version").Scan(&before); err != nil {
			t.Fatalf("connection %d: %v", i, err)
		}
		if _, err := c.ExecContext(ctx, fmt.Sprintf("PRAGMA schema_version=%d", before+41)); err != nil {
			t.Fatalf("connection %d: %v", i, err)
		}
		if err := c.QueryRowContext(ctx, "PRAGMA schema_version").Scan(&after); err != nil {
			t.Fatalf("connection %d: %v", i, err)
		}
		if after != before {
			t.Fatalf("connection %d changed schema_version under defensive mode: before=%d after=%d", i, before, after)
		}
	}
}

func TestUnprotectedPoolIsANegativeControl(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "pool-control.db")
	db, err := sql.Open("sqlite", dbPath+"?_defensive=0")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(4)

	ctx := context.Background()
	connections := make([]*sql.Conn, 0, 4)
	defer func() {
		for _, c := range connections {
			_ = c.Close()
		}
	}()
	for i := 0; i < 4; i++ {
		c, err := db.Conn(ctx)
		if err != nil {
			t.Fatalf("Conn(%d): %v", i, err)
		}
		connections = append(connections, c)
	}
	for i, c := range connections {
		var before, after int64
		if err := c.QueryRowContext(ctx, "PRAGMA schema_version").Scan(&before); err != nil {
			t.Fatalf("connection %d: %v", i, err)
		}
		if _, err := c.ExecContext(ctx, fmt.Sprintf("PRAGMA schema_version=%d", before+1)); err != nil {
			t.Fatalf("connection %d negative control: %v", i, err)
		}
		if err := c.QueryRowContext(ctx, "PRAGMA schema_version").Scan(&after); err != nil {
			t.Fatalf("connection %d: %v", i, err)
		}
		if after != before+1 {
			t.Fatalf("connection %d negative control did not mutate schema_version: before=%d after=%d", i, before, after)
		}
	}
}

func TestDefensiveBooleanForms(t *testing.T) {
	for _, value := range []string{"1", "true", "TRUE", "t"} {
		t.Run("on_"+value, func(t *testing.T) {
			db, err := sql.Open("sqlite", filepath.Join(t.TempDir(), "on.db")+"?_defensive="+value)
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			var before, after int64
			if err := db.QueryRow("PRAGMA schema_version").Scan(&before); err != nil {
				t.Fatal(err)
			}
			if _, err := db.Exec(fmt.Sprintf("PRAGMA schema_version=%d", before+41)); err != nil {
				t.Fatal(err)
			}
			if err := db.QueryRow("PRAGMA schema_version").Scan(&after); err != nil {
				t.Fatal(err)
			}
			if after != before {
				t.Fatalf("true form %q failed to protect schema_version", value)
			}
		})
	}
}

func TestUnmodifiedBehaviorIsNotAnADEDefensiveControl(t *testing.T) {
	for _, suffix := range []string{"", "?_defensive=0", "?_defensive=false", "?_defensive=FALSE", "?_defensive=f"} {
		t.Run(suffix, func(t *testing.T) {
			dbPath := filepath.Join(t.TempDir(), "negative-control.db")
			db, err := sql.Open("sqlite", dbPath+suffix)
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			var before, after int64
			if err := db.QueryRow("PRAGMA schema_version").Scan(&before); err != nil {
				t.Fatal(err)
			}
			if _, err := db.Exec(fmt.Sprintf("PRAGMA schema_version=%d", before+41)); err != nil {
				t.Fatalf("negative control unexpectedly rejected schema_version write: %v", err)
			}
			if err := db.QueryRow("PRAGMA schema_version").Scan(&after); err != nil {
				t.Fatal(err)
			}
			if after != before+41 {
				t.Fatalf("negative control did not demonstrate baseline behavior: before=%d after=%d", before, after)
			}
		})
	}
}
