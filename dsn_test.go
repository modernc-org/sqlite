// Copyright 2024 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"database/sql"
	"fmt"
	"path/filepath"
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

		// invalid settings should keep the default value
		verifyPragma(t, "_foreign_keys=x", "foreign_keys", "0")
		verifyPragma(t, "_fk=x", "foreign_keys", "0")
	})

	t.Run("Test _busy_timeout | _timeout", func(t *testing.T) {
		verifyPragma(t, "", "busy_timeout", "0")
		verifyPragma(t, "_timeout=5000", "busy_timeout", "5000")
		verifyPragma(t, "_busy_timeout=5000", "busy_timeout", "5000")

		// invalid timeouts should keep the default value
		verifyPragma(t, "_busy_timeout=-1", "busy_timeout", "0")
		verifyPragma(t, "_busy_timeout=true", "busy_timeout", "0")
		verifyPragma(t, "_busy_timeout=on", "busy_timeout", "0")
		verifyPragma(t, "_busy_timeout=x", "busy_timeout", "0")
	})

	t.Run("Test _journal_mode | _journal", func(t *testing.T) {
		verifyPragma(t, "", "journal_mode", "delete")
		verifyPragma(t, "_journal_mode=wal", "journal_mode", "wal")
		verifyPragma(t, "_journal=wal", "journal_mode", "wal")

		// invalid journal mode should keep the default value
		verifyPragma(t, "_journal=x", "journal_mode", "delete")
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

		// invalid journal mode selects NORMAL
		verifyPragma(t, "_synchronous=x", "synchronous", "1")
		verifyPragma(t, "_sync=x", "synchronous", "1")
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

		// invalid keeps it disabled
		verifyPragma(t, "_auto_vacuum=x", "auto_vacuum", "0")
		verifyPragma(t, "_vacuum=x", "auto_vacuum", "0")
	})

	t.Run("Test _query_only", func(t *testing.T) {
		// default is disabled
		verifyPragma(t, "", "query_only", "0")

		verifyPragma(t, "_query_only=0", "query_only", "0")

		// invalid keeps it disabled
		verifyPragma(t, "_query_only=x", "query_only", "0")
	})

	t.Run("Test combined DSN", func(t *testing.T) {
		// The mattn-compat keys must coexist in a single DSN. busy_timeout is
		// applied first and query_only last (see applyQueryParams); the DSN lists
		// them in a different order on purpose, so this also confirms the apply
		// order is fixed by the driver, not by the order the keys appear in the
		// DSN.
		verifyPragma(t, "_query_only=1&_synchronous=NORMAL&_journal_mode=wal&_foreign_keys=on&_busy_timeout=5000",
			"busy_timeout", "5000",
			"foreign_keys", "1",
			"journal_mode", "wal",
			"synchronous", "1",
			"query_only", "1",
		)
	})
}
