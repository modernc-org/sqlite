// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"path/filepath"
	"testing"

	sqlite3 "modernc.org/sqlite/lib"
)

// TestNoMutex checks that connections are opened with SQLITE_OPEN_NOMUTEX,
// so SQLite allocates no connection mutex. Under SQLITE_OPEN_FULLMUTEX that
// mutex admitted every caller on the connection's TLS, so it kept nothing
// apart and only cost a lock and an unlock per column read. See
// https://github.com/modernc-org/sqlite/issues/13.
func TestNoMutex(t *testing.T) {
	for _, dsn := range []string{
		":memory:",
		filepath.Join(t.TempDir(), "test.db"),
		"file:nomutex?mode=memory&cache=shared",
	} {
		c, err := newConn(dsn)
		if err != nil {
			t.Fatalf("%s: %v", dsn, err)
		}

		if m := sqlite3.Xsqlite3_db_mutex(c.tls, c.db); m != 0 {
			t.Errorf("%s: the connection has a mutex (%#x), want none", dsn, m)
		}
		c.Close()
	}
}
