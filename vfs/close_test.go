// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vfs

import (
	"database/sql"
	"errors"
	"os"
	"sync/atomic"
	"testing"
)

// TestFSCloseInUse checks that closing a file system with a database still
// open through it is refused. It used to go ahead, free the sqlite3_vfs the
// open connection still called through, and crash the next query.
func TestFSCloseInUse(t *testing.T) {
	const dbname = "inuse_canary.db"
	dir := t.TempDir()
	seedDB(t, dir, dbname,
		"CREATE TABLE t (v TEXT NOT NULL)",
		"INSERT INTO t (v) VALUES ('open')",
	)
	name, fsvfs, err := New(os.DirFS(dir))
	if err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("sqlite", "file:"+dbname+"?vfs="+name)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	db.SetMaxOpenConns(1)
	var got string
	if err := db.QueryRow("SELECT v FROM t").Scan(&got); err != nil {
		t.Fatalf("query: %s", E(err))
	}

	if err := fsvfs.Close(); !errors.Is(err, ErrInUse) {
		t.Fatalf("Close with a database open: got %v, want ErrInUse", err)
	}

	// Still registered and usable.
	if err := db.QueryRow("SELECT v FROM t").Scan(&got); err != nil || got != "open" {
		t.Fatalf("query after refused Close: %q, %s", got, E(err))
	}

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	if err := fsvfs.Close(); err != nil {
		t.Fatalf("Close after the database closed: %v", err)
	}

	if err := fsvfs.Close(); err != nil {
		t.Fatalf("second Close: %v", err)
	}
}

// TestAddObjectSkipsLiveHandles checks the handle allocator across a
// wraparound, which 32-bit targets reach after 2^32 opens: a handle still in
// use is skipped rather than overwritten, and so is 0.
func TestAddObjectSkipsLiveHandles(t *testing.T) {
	saved := atomic.LoadUintptr(&fToken)
	defer atomic.StoreUintptr(&fToken, saved)

	live := addObject("live")
	defer removeObject(live)

	atomic.StoreUintptr(&fToken, live-1) // the next token() is live
	h := addObject("next")
	if h == live {
		t.Fatalf("addObject reused live handle %#x", live)
	}

	if got := getObject(live); got != "live" {
		t.Fatalf("live handle now holds %v", got)
	}

	removeObject(h)

	atomic.StoreUintptr(&fToken, ^uintptr(0)) // the next token() is 0
	h = addObject("wrapped")
	if h == 0 {
		t.Fatal("addObject handed out handle 0")
	}

	removeObject(h)
}
