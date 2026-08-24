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

// TestOFDLockInterleavedReadersWrite verifies that after two interleaved read
// transactions on separate connections in the same process have both committed,
// a subsequent write transaction on either connection succeeds without leaking
// an OFD read lock (which would cause SQLITE_BUSY).
//
// Authored by Jan Mercl (@cznic) in https://gitlab.com/cznic/libsqlite3/-/merge_requests/3#note_3726270793.
func TestOFDLockInterleavedReadersWrite(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "interleaved_ofd_test.db")

	db, err := sql.Open(driverName, dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	c1, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer c1.Close()

	c2, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close()

	if _, err := c1.ExecContext(ctx, "CREATE TABLE t(x); INSERT INTO t VALUES(1);"); err != nil {
		t.Fatal(err)
	}

	// 1. Begin read transaction on c1 (acquires kernel OFD read lock via c1's fd).
	var count int
	if _, err := c1.ExecContext(ctx, "BEGIN;"); err != nil {
		t.Fatal(err)
	}
	if err := c1.QueryRowContext(ctx, "SELECT count(*) FROM t;").Scan(&count); err != nil {
		t.Fatal(err)
	}

	// 2. Begin read transaction on c2 while c1 is still reading.
	if _, err := c2.ExecContext(ctx, "BEGIN;"); err != nil {
		t.Fatal(err)
	}
	if err := c2.QueryRowContext(ctx, "SELECT count(*) FROM t;").Scan(&count); err != nil {
		t.Fatal(err)
	}

	// 3. Commit c1 first, then commit c2.
	if _, err := c1.ExecContext(ctx, "COMMIT;"); err != nil {
		t.Fatal(err)
	}
	if _, err := c2.ExecContext(ctx, "COMMIT;"); err != nil {
		t.Fatal(err)
	}

	// 4. Now attempt a write transaction on c2.
	// If c1's kernel OFD read lock leaked because c2's unlock was a no-op on c1's fd,
	// this write will fail with SQLITE_BUSY.
	if _, err := c2.ExecContext(ctx, "BEGIN IMMEDIATE; INSERT INTO t VALUES(2); COMMIT;"); err != nil {
		t.Fatalf("write after interleaved reads failed (leaked OFD read lock): %v", err)
	}
}

// The SHARED range of SQLite's locking protocol: SHARED_FIRST = PENDING_BYTE+2,
// SHARED_SIZE = 510, with the default PENDING_BYTE = 0x40000000.
const (
	sharedFirst = 0x40000000 + 2
	sharedSize  = 510
)

// probeSharedRange attempts a conflicting F_OFD_SETLK write lock on the SHARED
// range from an independent descriptor and returns the fcntl error: nil means
// the kernel granted it, i.e. no connection in this process holds a read lock.
func probeSharedRange(t *testing.T, dbPath string) error {
	t.Helper()
	fd, err := syscall.Open(dbPath, syscall.O_RDWR, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(fd)
	fl := syscall.Flock_t{Type: syscall.F_WRLCK, Whence: 0, Start: sharedFirst, Len: sharedSize}
	err = syscall.FcntlFlock(uintptr(fd), fOfdSetlk, &fl)
	if err == syscall.EINVAL {
		t.Skipf("F_OFD_SETLK unsupported on this kernel/filesystem (errno %v)", err)
	}
	return err
}

// TestOFDLockReadOnlyFirstLocker: a read-only connection takes the first lock
// on the inode, then a read-write connection joins at SHARED, which moves the
// designated locking descriptor from the read-only to the read-write one. The
// process must keep a kernel read lock on the SHARED range throughout, and the
// read-write connection must still be able to upgrade.
func TestOFDLockReadOnlyFirstLocker(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "ofd_ro_first.db")
	ctx := context.Background()

	rw, err := sql.Open(driverName, dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rw.Close()
	if _, err := rw.ExecContext(ctx, "CREATE TABLE t(x); INSERT INTO t VALUES(1);"); err != nil {
		t.Fatal(err)
	}

	ro, err := sql.Open(driverName, "file:"+dbPath+"?mode=ro")
	if err != nil {
		t.Fatal(err)
	}
	defer ro.Close()

	roConn, err := ro.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer roConn.Close()
	rwConn, err := rw.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rwConn.Close()

	var n int
	// 1. The read-only connection is the first locker on the inode.
	if _, err := roConn.ExecContext(ctx, "BEGIN"); err != nil {
		t.Fatal(err)
	}
	defer roConn.ExecContext(ctx, "ROLLBACK")
	if err := roConn.QueryRowContext(ctx, "SELECT count(*) FROM t").Scan(&n); err != nil {
		t.Fatal(err)
	}
	if err := probeSharedRange(t, dbPath); err == nil {
		t.Fatal("no kernel read lock while the read-only connection is inside a read transaction")
	}

	// 2. The read-write connection joins at SHARED.
	if _, err := rwConn.ExecContext(ctx, "BEGIN"); err != nil {
		t.Fatal(err)
	}
	defer rwConn.ExecContext(ctx, "ROLLBACK")
	if err := rwConn.QueryRowContext(ctx, "SELECT count(*) FROM t").Scan(&n); err != nil {
		t.Fatal(err)
	}
	if err := probeSharedRange(t, dbPath); err == nil {
		t.Fatal("kernel granted a write lock on the SHARED range while two connections hold read transactions")
	}

	// 3. The read-write connection can still upgrade (RESERVED through the
	// migrated descriptor).
	if _, err := rwConn.ExecContext(ctx, "INSERT INTO t VALUES(2)"); err != nil {
		t.Fatal(err)
	}
}

// TestOFDLockFailedFirstLock: the first SHARED attempt on an inode fails after
// its PENDING lock succeeded (a foreign write lock covers the SHARED range
// only), that connection is closed, and another connection that kept the inode
// alive then reads. The read must not go through the closed descriptor.
func TestOFDLockFailedFirstLock(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "ofd_failed_first.db")
	ctx := context.Background()

	setup, err := sql.Open(driverName, dbPath)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := setup.ExecContext(ctx, "CREATE TABLE t(x); INSERT INTO t VALUES(1);"); err != nil {
		t.Fatal(err)
	}
	setup.Close()

	// B is open and idle for the whole test: it keeps the inode entry alive.
	dbB, err := sql.Open(driverName, dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer dbB.Close()
	connB, err := dbB.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer connB.Close()
	if err := connB.PingContext(ctx); err != nil {
		t.Fatal(err)
	}

	dbA, err := sql.Open(driverName, dbPath)
	if err != nil {
		t.Fatal(err)
	}
	connA, err := dbA.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// A foreign OFD write lock on the SHARED range only (no PENDING byte), so
	// A's SHARED attempt fails only after its PENDING lock succeeded.
	fd, err := syscall.Open(dbPath, syscall.O_RDWR, 0)
	if err != nil {
		t.Fatal(err)
	}
	fl := syscall.Flock_t{Type: syscall.F_WRLCK, Whence: 0, Start: sharedFirst, Len: sharedSize}
	if err := syscall.FcntlFlock(uintptr(fd), fOfdSetlk, &fl); err != nil {
		if err == syscall.EINVAL {
			t.Skipf("F_OFD_SETLK unsupported on this kernel/filesystem (errno %v)", err)
		}
		t.Fatal(err)
	}

	var n int
	if err := connA.QueryRowContext(ctx, "SELECT count(*) FROM t").Scan(&n); err == nil {
		t.Fatal("expected SQLITE_BUSY: the foreign write lock must block the SHARED lock")
	}
	connA.Close()
	dbA.Close()       // closes A's descriptor
	syscall.Close(fd) // releases the foreign lock

	if err := connB.QueryRowContext(ctx, "SELECT count(*) FROM t").Scan(&n); err != nil {
		t.Fatalf("read on the surviving connection failed: %v", err)
	}
}
