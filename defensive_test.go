// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

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

func TestDefensiveOffPoolIsANegativeControl(t *testing.T) {
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

func TestDefensiveAbsentOrFalseIsBaseline(t *testing.T) {
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

// TestDefensiveRejectsJournalModeOff covers the one shorthand DSN key
// defensive mode silently neuters. SQLite turns PRAGMA journal_mode=OFF into
// a no-op that still reports success, so accepting the combination would
// honour neither parameter without saying so; applyQueryParams rejects it in
// the validation phase, before any statement runs.
func TestDefensiveRejectsJournalModeOff(t *testing.T) {
	for _, tc := range []struct {
		query   string
		wantErr string
	}{
		{"?_defensive=1&_journal_mode=OFF", "cannot take effect under _defensive"},
		{"?_defensive=1&_journal_mode=off", "cannot take effect under _defensive"},
		{"?_defensive=1&_journal=OFF", "cannot take effect under _defensive"},
		{"?_defensive=true&_journal_mode=Off", "cannot take effect under _defensive"},
		// The combination is only rejected when defensive mode would
		// actually suppress the mode change.
		{"?_defensive=0&_journal_mode=OFF", ""},
		{"?_journal_mode=OFF", ""},
		{"?_defensive=1&_journal_mode=WAL", ""},
	} {
		t.Run(tc.query, func(t *testing.T) {
			dbPath := filepath.Join(t.TempDir(), "journal.db")
			db, err := sql.Open("sqlite", dbPath+tc.query)
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			err = db.Ping()
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("Ping = %v, want success", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("Ping error = %v, want substring %q", err, tc.wantErr)
			}
			// The rejection belongs to the validation phase, so it must
			// not have run the PRAGMAs that precede it. journal_mode is
			// persistent, so a database converted here would outlive the
			// failed Open.
			var mode string
			check, err := sql.Open("sqlite", dbPath)
			if err != nil {
				t.Fatal(err)
			}
			defer check.Close()
			if err := check.QueryRow("PRAGMA journal_mode").Scan(&mode); err != nil {
				t.Fatal(err)
			}
			if strings.EqualFold(mode, "off") {
				t.Fatalf("rejected DSN still converted the database: journal_mode=%q", mode)
			}
		})
	}
}

// TestDefensiveDSNForms pins the option to the DSN shapes newConn handles
// differently. A file: DSN keeps its query string all the way into
// sqlite3_open_v2, so SQLite parses _defensive too and must ignore it; any
// other DSN has the query stripped before the open.
func TestDefensiveDSNForms(t *testing.T) {
	for _, tc := range []struct{ name, dsn string }{
		{"path", filepath.Join(t.TempDir(), "path.db") + "?_defensive=1"},
		{"file_uri", "file:" + filepath.Join(t.TempDir(), "uri.db") + "?_defensive=1"},
		{"memory", ":memory:?_defensive=1"},
		{"file_memory", "file::memory:?cache=shared&mode=memory&_defensive=1"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			db, err := sql.Open("sqlite", tc.dsn)
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			db.SetMaxOpenConns(1)
			if _, err := db.Exec("CREATE TABLE t(a)"); err != nil {
				t.Fatal(err)
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
				t.Fatalf("defensive mode not applied for %q: before=%d after=%d", tc.dsn, before, after)
			}
		})
	}
}

// TestDefensiveLeavesOrdinaryUseIntact is the other half of the contract: the
// option restricts the operations documented on Driver.Open and nothing else.
// It runs alongside every other DSN parameter that could plausibly collide
// with it and exercises the features whose implementation touches the
// machinery defensive mode guards.
func TestDefensiveLeavesOrdinaryUseIntact(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "workload.db")
	db, err := sql.Open("sqlite", dbPath+"?_defensive=1&_journal_mode=WAL&_auto_vacuum=FULL"+
		"&_foreign_keys=1&_busy_timeout=1000&_synchronous=NORMAL&_txlock=immediate&_dqs=0"+
		"&_pragma=cache_size(-2000)")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	for _, q := range []struct{ pragma, want string }{
		{"journal_mode", "wal"},
		{"auto_vacuum", "1"},
		{"foreign_keys", "1"},
		{"busy_timeout", "1000"},
		{"synchronous", "1"},
	} {
		var got string
		if err := db.QueryRow("PRAGMA " + q.pragma).Scan(&got); err != nil {
			t.Fatalf("PRAGMA %s: %v", q.pragma, err)
		}
		if !strings.EqualFold(got, q.want) {
			t.Errorf("PRAGMA %s = %q, want %q", q.pragma, got, q.want)
		}
	}

	// _dqs is applied after _defensive; it must still have taken effect.
	if _, err := db.Exec(`CREATE TABLE t(a TEXT)`); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO t VALUES("bare-string")`); err == nil {
		t.Error(`_dqs=0 did not take effect alongside _defensive=1`)
	}

	if _, err := db.Exec(`
		CREATE TABLE parent(id INTEGER PRIMARY KEY);
		CREATE TABLE child(id INTEGER PRIMARY KEY, p INTEGER REFERENCES parent(id));
		INSERT INTO parent VALUES(1);
	`); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO child VALUES(1, 99)`); err == nil {
		t.Error("foreign key enforcement lost under defensive mode")
	}

	// Ordinary use of a virtual table that owns shadow tables keeps working;
	// only direct writes to the shadow tables are refused.
	if _, err := db.Exec(`CREATE VIRTUAL TABLE ft USING fts5(body)`); err != nil {
		t.Fatalf("CREATE VIRTUAL TABLE fts5: %v", err)
	}
	if _, err := db.Exec(`INSERT INTO ft(body) VALUES('hello world')`); err != nil {
		t.Fatalf("fts5 INSERT: %v", err)
	}
	var n int
	if err := db.QueryRow(`SELECT count(*) FROM ft WHERE ft MATCH 'hello'`).Scan(&n); err != nil {
		t.Fatalf("fts5 MATCH: %v", err)
	}
	if n != 1 {
		t.Errorf("fts5 MATCH count = %d, want 1", n)
	}
	if _, err := db.Exec(`INSERT INTO ft_data(id, block) VALUES(999, x'00')`); err == nil {
		t.Error("direct fts5 shadow-table write succeeded under defensive mode")
	}
	if err := db.QueryRow(`SELECT count(*) FROM sqlite_dbpage`).Scan(&n); err != nil {
		t.Errorf("reading sqlite_dbpage: %v", err)
	}
	if _, err := db.Exec(`UPDATE sqlite_dbpage SET data=zeroblob(4096) WHERE pgno=1`); err == nil {
		t.Error("sqlite_dbpage write succeeded under defensive mode")
	}

	if _, err := db.Exec(`VACUUM`); err != nil {
		t.Errorf("VACUUM under defensive mode: %v", err)
	}
	var integrity string
	if err := db.QueryRow(`PRAGMA integrity_check`).Scan(&integrity); err != nil {
		t.Fatal(err)
	}
	if integrity != "ok" {
		t.Errorf("integrity_check = %q, want ok", integrity)
	}
}

// TestDefensiveIsPerConnection documents the scope of the option: it is a
// property of the connection, not of the database file, so a second handle
// opened without it is unrestricted. Driver.Open says so; this pins it.
func TestDefensiveIsPerConnection(t *testing.T) {
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "scope.db")

	hard, err := sql.Open("sqlite", dbPath+"?_defensive=1")
	if err != nil {
		t.Fatal(err)
	}
	defer hard.Close()
	hard.SetMaxOpenConns(1)
	if _, err := hard.ExecContext(ctx, "CREATE TABLE protected(id INTEGER PRIMARY KEY)"); err != nil {
		t.Fatal(err)
	}
	if _, err := hard.ExecContext(ctx, "PRAGMA writable_schema=ON"); err != nil {
		t.Fatal(err)
	}
	if _, err := hard.ExecContext(ctx, "DELETE FROM sqlite_schema WHERE name='protected'"); err == nil {
		t.Fatal("defensive connection allowed a direct sqlite_schema write")
	}

	soft, err := sql.Open("sqlite", dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer soft.Close()
	soft.SetMaxOpenConns(1)
	if _, err := soft.ExecContext(ctx, "PRAGMA writable_schema=ON"); err != nil {
		t.Fatal(err)
	}
	if _, err := soft.ExecContext(ctx, "DELETE FROM sqlite_schema WHERE name='protected'"); err != nil {
		t.Fatalf("a handle opened without _defensive should be unrestricted, got %v", err)
	}
}
