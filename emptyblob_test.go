// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"database/sql"
	"testing"

	sqlite3 "modernc.org/sqlite/lib"
)

// TestEmptyBlob checks that a zero-length BLOB reads as an empty []byte and
// NULL as nil.
func TestEmptyBlob(t *testing.T) {
	db, err := sql.Open(driverName, "file::memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`
		CREATE TABLE t(i INTEGER, b BLOB);
		INSERT INTO t VALUES (1, x''), (2, ?), (3, NULL);
	`, []byte{}); err != nil {
		t.Fatal(err)
	}

	rows, err := db.Query("SELECT i, b, typeof(b) FROM t ORDER BY i")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i int
		var b []byte
		var typ string
		if err := rows.Scan(&i, &b, &typ); err != nil {
			t.Fatal(err)
		}

		switch typ {
		case "blob":
			if b == nil || len(b) != 0 {
				t.Errorf("row %d: got %#v, want an empty, non-nil []byte", i, b)
			}
		case "null":
			if b != nil {
				t.Errorf("row %d: got %#v, want nil", i, b)
			}
		default:
			t.Errorf("row %d: typeof(b) = %q, want blob or null", i, typ)
		}
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}

	var v any
	if err := db.QueryRow("SELECT x''").Scan(&v); err != nil {
		t.Fatal(err)
	}

	if b, ok := v.([]byte); !ok || b == nil {
		t.Errorf("SELECT x'': got %#v, want an empty, non-nil []byte", v)
	}
}

// TestEmptyBlobPreUpdateHook checks the same for the pre-update hook.
func TestEmptyBlobPreUpdateHook(t *testing.T) {
	var got []any
	var hookErr error
	var testDriver Driver
	testDriver.RegisterConnectionHook(func(conn ExecQuerierContext, dsn string) error {
		conn.(HookRegisterer).RegisterPreUpdateHook(func(data SQLitePreUpdateData) {
			if data.Op != sqlite3.SQLITE_INSERT {
				return
			}

			row := make([]any, data.Count())
			if err := data.New(row...); err != nil {
				hookErr = err
				return
			}

			got = append(got, row[0])
		})
		return nil
	})
	sql.Register("sqlite_empty_blob_pre_update_hook_test", &testDriver)

	db, err := sql.Open("sqlite_empty_blob_pre_update_hook_test", "file::memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`
		CREATE TABLE t(b BLOB);
		INSERT INTO t VALUES (x''), (NULL);
	`); err != nil {
		t.Fatal(err)
	}

	if hookErr != nil {
		t.Fatal(hookErr)
	}

	if len(got) != 2 {
		t.Fatalf("the hook saw %d rows, want 2", len(got))
	}

	if b, ok := got[0].([]byte); !ok || b == nil {
		t.Errorf("x'': got %#v, want an empty, non-nil []byte", got[0])
	}

	if got[1] != nil {
		t.Errorf("null: got %#v, want nil", got[1])
	}
}
