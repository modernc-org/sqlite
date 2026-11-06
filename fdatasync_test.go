// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"database/sql"
	"path/filepath"
	"runtime"
	"testing"
)

// TestFdatasync checks the switch: off by default, settable and queryable
// where the library syncs with fdatasync, unavailable and inert on macOS and
// Windows. Which system call the library then issues is proven by
// TestFdatasyncSyscall in modernc.org/libsqlite3, whose transpiles lib/
// carries unchanged.
func TestFdatasync(t *testing.T) {
	if FdatasyncEnabled() {
		t.Fatal("fdatasync on by default")
	}
	switch runtime.GOOS {
	case "darwin", "windows":
		for _, on := range []bool{true, false} {
			if _, err := Fdatasync(on); err != ErrFdatasyncUnavailable {
				t.Fatalf("Fdatasync(%v): err = %v, want ErrFdatasyncUnavailable", on, err)
			}
		}
		if FdatasyncEnabled() {
			t.Fatal("FdatasyncEnabled() = true where fdatasync is unavailable")
		}
		return
	}

	defer Fdatasync(false)
	for i, v := range []struct{ on, prev bool }{{true, false}, {true, true}, {false, true}, {false, false}} {
		if prev, err := Fdatasync(v.on); err != nil || prev != v.prev {
			t.Fatalf("%d: Fdatasync(%v): prev = %v, err = %v, want %v, nil", i, v.on, prev, err, v.prev)
		}
		if g, e := FdatasyncEnabled(), v.on; g != e {
			t.Fatalf("%d: FdatasyncEnabled() = %v, want %v", i, g, e)
		}
	}

	// A connection opened before the switch commits through it: the setting
	// is read at every sync.
	db, err := sql.Open(driverName, filepath.Join(t.TempDir(), "fdatasync.db")+"?_pragma=journal_mode(WAL)&_pragma=synchronous(FULL)")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if _, err := db.Exec("CREATE TABLE t(x)"); err != nil {
		t.Fatal(err)
	}
	for _, on := range []bool{true, false, true} {
		if _, err := Fdatasync(on); err != nil {
			t.Fatal(err)
		}
		if _, err := db.Exec("INSERT INTO t VALUES(1)"); err != nil {
			t.Fatalf("commit with fdatasync %v: %v", on, err)
		}
	}
	var n int
	if err := db.QueryRow("SELECT count(*) FROM t").Scan(&n); err != nil || n != 3 {
		t.Fatalf("count = %d, err = %v, want 3, nil", n, err)
	}
}
