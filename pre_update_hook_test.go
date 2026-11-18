package sqlite_test

import (
	"context"
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func TestPreUpdateHook(t *testing.T) {
	connStr := ":memory:"
	driverName := "sqlite_pre_update_hook_test"

	var (
		insertCount     int
		insertNewValues []any
		updateCount     int
		updateOldValues []any
		updateNewValues []any
		deleteCount     int
		deleteOldValues []any
		commitCount     int
		rollbackCount   int
	)
	var testDriver sqlite.Driver
	testDriver.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if hooker, ok := conn.(sqlite.HookRegisterer); ok {
			hooker.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
				switch data.Op {
				case sqlite3.SQLITE_INSERT:
					insertCount++
					insertNewValues = make([]any, data.Count())
					err := data.New(insertNewValues...)
					if err != nil {
						t.Fatal(err)
					}
				case sqlite3.SQLITE_UPDATE:
					updateCount++
					updateOldValues = make([]any, data.Count())
					err := data.Old(updateOldValues...)
					if err != nil {
						t.Fatal(err)
					}
					updateNewValues = make([]any, data.Count())
					err = data.New(updateNewValues...)
					if err != nil {
						t.Fatal(err)
					}
				case sqlite3.SQLITE_DELETE:
					deleteCount++
					deleteOldValues = make([]any, data.Count())
					err := data.Old(deleteOldValues...)
					if err != nil {
						t.Fatal(err)
					}
				}
			})
			hooker.RegisterCommitHook(func() int32 {
				commitCount++
				return 0
			})
			hooker.RegisterRollbackHook(func() {
				rollbackCount++
			})
		}
		return nil
	})

	sql.Register(driverName, &testDriver)

	db, err := sql.Open(driverName, connStr)
	if err != nil {
		t.Fatal(err)
	}

	expectInsertValues := []any{int64(42), 3.1415, "Test", "will be nil"}
	expectUpdateValues := []any{int64(43), 1.5, "Test update", nil}
	_, err = db.Exec(`
	CREATE TABLE in_memory_test(id INTEGER PRIMARY KEY, f FLOAT, t TEXT, x ANY);
	INSERT INTO in_memory_test VALUES(42, 3.1415, 'Test', 'will be nil');
	UPDATE in_memory_test SET id = 43, f = 1.5, t = 'Test update', x = null;
	DELETE FROM in_memory_test;
	SELECT 1;
	`)
	if err != nil {
		t.Fatal(err)
	}
	if insertCount != 1 {
		t.Errorf("pre update hook: expect %d inserts call, got %d", 1, insertCount)
	}
	if !slices.Equal(insertNewValues, expectInsertValues) {
		t.Errorf("pre update hook: expect %v as inserted new values, got %v", expectInsertValues, insertNewValues)
	}
	if updateCount != 1 {
		t.Errorf("pre update hook: expect %d updates call, got %d", 1, updateCount)
	}
	if !slices.Equal(updateOldValues, expectInsertValues) {
		t.Errorf("pre update hook: expect %v as updated old values, got %v", expectInsertValues, updateOldValues)
	}
	if !slices.Equal(updateNewValues, expectUpdateValues) {
		t.Errorf("pre update hook: expect %v as updated new values, got %v", expectUpdateValues, updateNewValues)
	}
	if deleteCount != 1 {
		t.Errorf("pre update hook: expect %d deletes call, got %d", 1, deleteCount)
	}
	if !slices.Equal(deleteOldValues, expectUpdateValues) {
		t.Errorf("pre update hook: expect %v as deleted old values, got %v", expectUpdateValues, deleteOldValues)
	}
	if commitCount != 4 {
		t.Errorf("commit hook: expect %d, got %d", 4, commitCount)
	}
	if rollbackCount != 0 {
		t.Errorf("rollback hook: expect %d, got %d", 0, rollbackCount)
	}

}

func TestPreUpdateHookLargeRowID(t *testing.T) {
	driverName := fmt.Sprintf("sqlite_preupdate_large_rowid_%p", t)
	var gotOldRowID, gotNewRowID int64
	var gotOp int32

	var testDriver sqlite.Driver
	testDriver.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if hooker, ok := conn.(sqlite.HookRegisterer); ok {
			hooker.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
				gotOp = data.Op
				gotOldRowID = data.OldRowID
				gotNewRowID = data.NewRowID
			})
		}
		return nil
	})
	sql.Register(driverName, &testDriver)

	db, err := sql.Open(driverName, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE t(id INTEGER PRIMARY KEY, v TEXT)`)
	if err != nil {
		t.Fatal(err)
	}

	const largeRowID = int64(1) << 40 // 1 TiB — well beyond int32 range
	_, err = db.Exec(`INSERT INTO t VALUES(?, 'hello')`, largeRowID)
	if err != nil {
		t.Fatal(err)
	}

	if gotOp != sqlite3.SQLITE_INSERT {
		t.Fatalf("expected INSERT op, got %d", gotOp)
	}
	if gotNewRowID != largeRowID {
		t.Errorf("NewRowID: got %d, want %d", gotNewRowID, largeRowID)
	}

	// Also test UPDATE with large rowids
	const newLargeRowID = int64(1)<<40 + 1
	gotOldRowID = 0
	gotNewRowID = 0
	_, err = db.Exec(`UPDATE t SET id = ?, v = 'updated' WHERE id = ?`, newLargeRowID, largeRowID)
	if err != nil {
		t.Fatal(err)
	}

	if gotOp != sqlite3.SQLITE_UPDATE {
		t.Fatalf("expected UPDATE op, got %d", gotOp)
	}
	if gotOldRowID != largeRowID {
		t.Errorf("OldRowID after UPDATE: got %d, want %d", gotOldRowID, largeRowID)
	}
	if gotNewRowID != newLargeRowID {
		t.Errorf("NewRowID after UPDATE: got %d, want %d", gotNewRowID, newLargeRowID)
	}
}

// sqliteVarint reads a variable-length integer of the database file format.
func sqliteVarint(b []byte) (v uint64, n int) {
	for n < 8 {
		v = v<<7 | uint64(b[n]&0x7f)
		n++
		if b[n-1] < 0x80 {
			return v, n
		}
	}
	return v<<8 | uint64(b[8]), 9
}

// firstOverflowPointer returns where in the database file db the first cell
// of page 2 has the number of its first overflow page. Page 2 must be the
// leaf that is the root of a table, and the payload of the cell must not fit
// on it. How much of the payload stays in the cell is the arithmetic of the
// file format ("B-tree Cell Format").
func firstOverflowPointer(t *testing.T, db []byte) int {
	t.Helper()
	pageSize := int(binary.BigEndian.Uint16(db[16:]))
	if pageSize == 1 {
		pageSize = 65536
	}

	usable := pageSize - int(db[20])
	page := db[pageSize : 2*pageSize]
	if page[0] != 13 {
		t.Fatalf("page 2 has type %d, want 13, a table leaf", page[0])
	}

	cell := int(binary.BigEndian.Uint16(page[8:]))
	payload, n := sqliteVarint(page[cell:])
	_, m := sqliteVarint(page[cell+n:]) // the rowid
	size, most := int(payload), usable-35
	if size <= most {
		t.Fatalf("a payload of %d bytes has no overflow page", size)
	}

	least := (usable-12)*32/255 - 23
	local := least + (size-least)%(usable-4)
	if local > most {
		local = least
	}
	return pageSize + cell + n + m + local
}

// TestPreUpdateHookRowUnreadable has a pre-update hook ask for the old row
// when SQLite cannot read it. The row has a value that goes on to overflow
// pages, and in the damaged copy of the database the number of the first of
// them points past the end of the file. SQLite reads the old row only when
// the hook asks for it and then returns SQLITE_CORRUPT without a value: Old
// must return that error. The statement fails with the same code on its own
// account.
func TestPreUpdateHookRowUnreadable(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "intact.db")
	text := strings.Repeat("A", 20000)
	db, err := sql.Open("sqlite", "file:"+path)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec("PRAGMA page_size=4096; CREATE TABLE t(a TEXT); INSERT INTO t VALUES(?)", text); err != nil {
		t.Fatal(err)
	}

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	intact, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	at := firstOverflowPointer(t, intact)
	pages := uint32(len(intact) / 4096)
	if next := binary.BigEndian.Uint32(intact[at:]); next < 3 || next > pages {
		t.Fatalf("the bytes at offset %d are %d, not the number of a page of this file of %d pages", at, next, pages)
	}

	damaged := slices.Clone(intact)
	binary.BigEndian.PutUint32(damaged[at:], pages+1000)

	for _, stmt := range []string{"DELETE FROM t", "UPDATE t SET a = 'x'"} {
		for _, c := range []struct {
			name    string
			image   []byte
			corrupt bool
		}{
			{"intact", intact, false},
			{"damaged", damaged, true},
		} {
			t.Run(stmt+", "+c.name, func(t *testing.T) {
				path := filepath.Join(t.TempDir(), "test.db")
				if err := os.WriteFile(path, c.image, 0o644); err != nil {
					t.Fatal(err)
				}

				db, err := sql.Open("sqlite", "file:"+path)
				if err != nil {
					t.Fatal(err)
				}
				defer db.Close()

				ctx := context.Background()
				conn, err := db.Conn(ctx)
				if err != nil {
					t.Fatal(err)
				}
				defer conn.Close()

				var (
					calls    int
					old      any
					oldErr   error
					panicked any
				)
				if err := conn.Raw(func(driverConn any) error {
					driverConn.(sqlite.HookRegisterer).RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
						// A crash in Old is what this test is about: report
						// it as a failure and let the others run.
						defer func() { panicked = recover() }()
						calls++
						dest := make([]any, data.Count())
						oldErr = data.Old(dest...)
						old = dest[0]
					})
					return nil
				}); err != nil {
					t.Fatal(err)
				}

				_, stmtErr := conn.ExecContext(ctx, stmt)
				if panicked != nil {
					t.Fatalf("Old panicked: %v", panicked)
				}

				if calls != 1 {
					t.Fatalf("the hook ran %d time(s), want once", calls)
				}

				if !c.corrupt {
					if oldErr != nil || stmtErr != nil || old != text {
						t.Fatalf("Old: %v, the statement: %v, and the old value has the wanted %d bytes: %v", oldErr, stmtErr, len(text), old == text)
					}

					return
				}

				for what, err := range map[string]error{"Old": oldErr, "the statement": stmtErr} {
					var se *sqlite.Error
					if !errors.As(err, &se) || se.Code()&0xff != sqlite3.SQLITE_CORRUPT {
						t.Errorf("%s returned %v, want an error with SQLITE_CORRUPT", what, err)
					}
				}
				if old != nil {
					t.Errorf("Old returned an error and still stored a value of type %T", old)
				}
			})
		}
	}
}

// TestPreUpdateHookTextNUL checks that the pre-update hook gets a TEXT value
// whole when it holds a NUL. It used to end at the first one.
func TestPreUpdateHookTextNUL(t *testing.T) {
	db, err := sql.Open("sqlite", "file::memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	conn, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	var got []string
	if err := conn.Raw(func(driverConn any) error {
		driverConn.(sqlite.HookRegisterer).RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
			record := func(what string, read func(...any) error) {
				dest := make([]any, data.Count())
				if err := read(dest...); err != nil {
					got = append(got, what+": "+err.Error())
					return
				}

				got = append(got, fmt.Sprintf("%s: %T %q", what, dest[0], dest[0]))
			}
			if data.Op != sqlite3.SQLITE_INSERT {
				record("old", data.Old)
			}
			if data.Op != sqlite3.SQLITE_DELETE {
				record("new", data.New)
			}
		})
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	for _, stmt := range []string{
		"CREATE TABLE t(a TEXT)",
		"INSERT INTO t VALUES(CAST(x'6100620063' AS TEXT)), (''), (CAST(x'00' AS TEXT))",
		"UPDATE t SET a = CAST(x'7A007A' AS TEXT) WHERE rowid = 1",
		"DELETE FROM t",
	} {
		if _, err := conn.ExecContext(ctx, stmt); err != nil {
			t.Fatalf("%s: %v", stmt, err)
		}
	}

	want := []string{
		`new: string "a\x00b\x00c"`, `new: string ""`, `new: string "\x00"`,
		`old: string "a\x00b\x00c"`, `new: string "z\x00z"`,
		`old: string "z\x00z"`, `old: string ""`, `old: string "\x00"`,
	}
	if !slices.Equal(got, want) {
		t.Fatalf("the hook got\n%q\nwant\n%q", got, want)
	}
}
