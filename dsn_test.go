// Copyright 2024 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestDSNPragmas(t *testing.T) {
	verifyPragma := func(t *testing.T, query string, kvs ...string) {
		// some PRAGMAs may require a physical database so using a real one even though it's never created
		dbPath := fmt.Sprintf("%s?%s", filepath.Join(t.TempDir(), "tmp.db"), query)
		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		if len(kvs) == 0 || len(kvs)%2 > 0 {
			t.Fatal("verifyPragma needs to be called with at least one pragma key-value pair")
		}

		for i := 0; i < len(kvs); i += 2 {
			pragmaName := kvs[i]
			expectedValue := kvs[i+1]

			var actualValue string
			err = db.QueryRow("PRAGMA " + pragmaName).Scan(&actualValue)
			if err != nil {
				t.Fatal(err)
			}

			if actualValue != expectedValue {
				t.Fatalf("PRAGMA %s returned '%s' but expected '%s'", pragmaName, actualValue, expectedValue)
			}
		}
	}

	// verifyPragmaError asserts that opening a connection with the given DSN
	// query fails with an error containing wantSubstr. applyQueryParams runs when
	// the connection is established, so the error surfaces on first use.
	verifyPragmaError := func(t *testing.T, query, wantSubstr string) {
		dbPath := fmt.Sprintf("%s?%s", filepath.Join(t.TempDir(), "tmp.db"), query)
		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		var x int
		err = db.QueryRow("SELECT 1").Scan(&x)
		if err == nil {
			t.Fatalf("query %q: expected an error, got none", query)
		}
		if !strings.Contains(err.Error(), wantSubstr) {
			t.Fatalf("query %q: error %q does not contain %q", query, err.Error(), wantSubstr)
		}
	}

	t.Run("Test _pragma", func(t *testing.T) {
		verifyPragma(t, "", "foreign_keys", "0", "busy_timeout", "0")
		verifyPragma(t, "_pragma=foreign_keys=on", "foreign_keys", "1")
		verifyPragma(t, "_pragma=foreign_keys=on&_pragma=busy_timeout=5000", "foreign_keys", "1", "busy_timeout", "5000")
		verifyPragma(t, "_pragma=foreign_keys%3Don&_pragma=busy_timeout%3D5000", "foreign_keys", "1", "busy_timeout", "5000")
	})

	t.Run("Test _foreign_keys | _fk", func(t *testing.T) {
		verifyPragma(t, "", "foreign_keys", "0")
		verifyPragma(t, "_foreign_keys=1", "foreign_keys", "1")
		verifyPragma(t, "_foreign_keys=on", "foreign_keys", "1")
		verifyPragma(t, "_foreign_keys=true", "foreign_keys", "1")
		verifyPragma(t, "_foreign_keys=", "foreign_keys", "0")
		verifyPragma(t, "_foreign_keys=0", "foreign_keys", "0")
		verifyPragma(t, "_foreign_keys=off", "foreign_keys", "0")
		verifyPragma(t, "_foreign_keys=false", "foreign_keys", "0")

		verifyPragma(t, "_fk=1", "foreign_keys", "1")
		verifyPragma(t, "_fk=on", "foreign_keys", "1")
		verifyPragma(t, "_fk=true", "foreign_keys", "1")
		verifyPragma(t, "_fk=", "foreign_keys", "0")
		verifyPragma(t, "_fk=0", "foreign_keys", "0")
		verifyPragma(t, "_fk=off", "foreign_keys", "0")
		verifyPragma(t, "_fk=false", "foreign_keys", "0")

		// invalid settings are a hard error (mattn-compatible)
		verifyPragmaError(t, "_foreign_keys=x", "invalid _foreign_keys")
		verifyPragmaError(t, "_fk=x", "invalid _fk")

		// when a key and its alias disagree, the alias wins (mattn-compatible)
		verifyPragma(t, "_foreign_keys=off&_fk=on", "foreign_keys", "1")
		verifyPragma(t, "_foreign_keys=on&_fk=off", "foreign_keys", "0")
	})

	t.Run("Test _busy_timeout | _timeout", func(t *testing.T) {
		verifyPragma(t, "", "busy_timeout", "0")
		verifyPragma(t, "_timeout=5000", "busy_timeout", "5000")
		verifyPragma(t, "_busy_timeout=5000", "busy_timeout", "5000")

		// a negative integer is valid input; SQLite clamps it to 0
		verifyPragma(t, "_busy_timeout=-1", "busy_timeout", "0")

		// non-integer values are a hard error (mattn-compatible)
		verifyPragmaError(t, "_busy_timeout=true", "invalid _busy_timeout")
		verifyPragmaError(t, "_busy_timeout=on", "invalid _busy_timeout")
		verifyPragmaError(t, "_busy_timeout=x", "invalid _busy_timeout")
		verifyPragmaError(t, "_timeout=x", "invalid _timeout")
	})

	t.Run("Test _journal_mode | _journal", func(t *testing.T) {
		verifyPragma(t, "", "journal_mode", "delete")
		verifyPragma(t, "_journal_mode=wal", "journal_mode", "wal")
		verifyPragma(t, "_journal=wal", "journal_mode", "wal")

		// invalid journal mode is a hard error (mattn-compatible)
		verifyPragmaError(t, "_journal=x", "invalid _journal")
		verifyPragmaError(t, "_journal_mode=WAL2", "invalid _journal_mode")
	})

	t.Run("Test _synchronous | _sync", func(t *testing.T) {
		// default is FULL
		verifyPragma(t, "", "synchronous", "2")

		verifyPragma(t, "_synchronous=1", "synchronous", "1")
		verifyPragma(t, "_synchronous=NORMAL", "synchronous", "1")

		verifyPragma(t, "_synchronous=2", "synchronous", "2")
		verifyPragma(t, "_synchronous=FULL", "synchronous", "2")

		verifyPragma(t, "_synchronous=3", "synchronous", "3")
		verifyPragma(t, "_synchronous=EXTRA", "synchronous", "3")

		verifyPragma(t, "_sync=1", "synchronous", "1")
		verifyPragma(t, "_sync=NORMAL", "synchronous", "1")

		verifyPragma(t, "_sync=2", "synchronous", "2")
		verifyPragma(t, "_sync=FULL", "synchronous", "2")

		verifyPragma(t, "_sync=3", "synchronous", "3")
		verifyPragma(t, "_sync=EXTRA", "synchronous", "3")

		// invalid synchronous mode is a hard error (mattn-compatible)
		verifyPragmaError(t, "_synchronous=x", "invalid _synchronous")
		verifyPragmaError(t, "_sync=x", "invalid _sync")
	})

	t.Run("Test _auto_vacuum | _vacuum", func(t *testing.T) {
		// default is disabled
		verifyPragma(t, "", "auto_vacuum", "0")

		verifyPragma(t, "_auto_vacuum=0", "auto_vacuum", "0")
		verifyPragma(t, "_auto_vacuum=1", "auto_vacuum", "1")
		verifyPragma(t, "_auto_vacuum=2", "auto_vacuum", "2")

		verifyPragma(t, "_vacuum=0", "auto_vacuum", "0")
		verifyPragma(t, "_vacuum=1", "auto_vacuum", "1")
		verifyPragma(t, "_vacuum=2", "auto_vacuum", "2")

		// invalid auto_vacuum is a hard error (mattn-compatible)
		verifyPragmaError(t, "_auto_vacuum=x", "invalid _auto_vacuum")
		verifyPragmaError(t, "_vacuum=x", "invalid _vacuum")
	})

	t.Run("Test _query_only", func(t *testing.T) {
		// default is disabled
		verifyPragma(t, "", "query_only", "0")

		verifyPragma(t, "_query_only=0", "query_only", "0")

		// invalid query_only is a hard error (mattn-compatible)
		verifyPragmaError(t, "_query_only=x", "invalid _query_only")
	})

	t.Run("Test combined DSN", func(t *testing.T) {
		// The mattn-compat keys must coexist in a single DSN. busy_timeout and
		// auto_vacuum are applied before the rest and query_only last (see
		// applyQueryParams); the DSN lists them in a different order on purpose,
		// so this also confirms the apply order is fixed by the driver, not by
		// the order the keys appear in the DSN. _auto_vacuum is included
		// deliberately: it only takes effect if applied before _journal_mode
		// materialises page 1, so reading it back as 1 proves the fixed order.
		verifyPragma(t, "_query_only=1&_synchronous=NORMAL&_journal_mode=wal&_auto_vacuum=1&_foreign_keys=on&_busy_timeout=5000",
			"busy_timeout", "5000",
			"auto_vacuum", "1",
			"foreign_keys", "1",
			"journal_mode", "wal",
			"synchronous", "1",
			"query_only", "1",
		)
	})

	t.Run("Test empty alias suppresses the PRAGMA", func(t *testing.T) {
		// Selection between a key and its alias is by presence, not by value, so
		// an alias supplied empty selects the alias and suppresses the PRAGMA
		// rather than falling back to the primary key. Matches mattn.
		verifyPragma(t, "_foreign_keys=on&_fk=", "foreign_keys", "0")
		verifyPragma(t, "_fk=on&_foreign_keys=", "foreign_keys", "1")
		verifyPragma(t, "_journal_mode=wal&_journal=", "journal_mode", "delete")
		verifyPragma(t, "_busy_timeout=5000&_timeout=", "busy_timeout", "0")

		// An empty value is not validated, so it is not an error either.
		verifyPragma(t, "_synchronous=", "synchronous", "2")
	})
}
