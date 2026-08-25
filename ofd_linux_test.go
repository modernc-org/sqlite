// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package sqlite

import (
	"context"
	"database/sql"
	"path/filepath"
	"syscall"
	"testing"
)

const (
	fOfdSetlk = 37
)

// TestOFDLockSurvivesOSClose verifies that closing an unrelated os.File descriptor
// pointing to the same inode does not strip the active SQLite database lock on Linux.
// Under standard POSIX inode locks (F_SETLK), close() on any descriptor drops
// all locks for that inode across the entire process. With Open File Description
// (OFD) locking (F_OFD_SETLK), locks are attached to the open file description,
// preventing accidental unlock on os.Close().
func TestOFDLockSurvivesOSClose(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "ofd_test.db")

	db, err := sql.Open(driverName, dbPath)
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

	if _, err := conn.ExecContext(ctx, "CREATE TABLE t(x);"); err != nil {
		t.Fatal(err)
	}

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "INSERT INTO t VALUES(1);"); err != nil {
		t.Fatal(err)
	}

	// Open and close an independent descriptor to the same database file.
	// On standard POSIX locking, this close() unconditionally strips the lock on the inode.
	fd, err := syscall.Open(dbPath, syscall.O_RDONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	if err := syscall.Close(fd); err != nil {
		t.Fatal(err)
	}

	// Now verify from a raw OS file descriptor that the OFD lock is still held by SQLite.
	// We use syscall.FcntlFlock so that 32-bit platforms (linux/386, linux/arm) route through fcntl64.
	probeFd, err := syscall.Open(dbPath, syscall.O_RDWR, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(probeFd)

	fl := syscall.Flock_t{
		Type:   syscall.F_WRLCK,
		Whence: 0,
		Start:  0,
		Len:    0,
	}

	err = syscall.FcntlFlock(uintptr(probeFd), fOfdSetlk, &fl)
	if err == syscall.EINVAL {
		t.Skipf("F_OFD_SETLK unsupported on this kernel/filesystem (errno %v)", err)
	}
	if err == nil {
		t.Fatal("kernel allowed conflicting write lock on probeFd; SQLite lock was stripped by close()")
	}
	if err != syscall.EAGAIN && err != syscall.EACCES {
		t.Fatalf("expected EAGAIN or EACCES when probing lock, got: %v", err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("tx.Commit failed: %v", err)
	}
}
