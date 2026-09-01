// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package sqlite // import "modernc.org/sqlite"

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
)

const (
	fOfdSetlk = 37

	// ofdChildEnvVar marks the child process a test re-executed with OFD
	// locking enabled; see inOFDChild.
	ofdChildEnvVar = "MODERNC_SQLITE_TEST_OFD_CHILD"
)

// reexecTest re-runs the calling test alone in a child process with the
// given environment and propagates the child's outcome: a skipped child
// skips the caller, a failed or absent run fails it.
func reexecTest(t *testing.T, env []string) {
	t.Helper()
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(exe, "-test.run=^"+t.Name()+"$", "-test.v")
	cmd.Env = env
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("child process: %v\n%s", err, out)
	}

	switch {
	case bytes.Contains(out, []byte("--- SKIP")):
		t.Skipf("child process skipped:\n%s", out)
	case !bytes.Contains(out, []byte("--- PASS")):
		t.Fatalf("child process did not run the test:\n%s", out)
	}
}

// inOFDChild makes the calling test run with OFD locking enabled. The mode
// is process-wide and frozen at the process's first database file lock, so a
// test cannot enable it inside the test binary the rest of the suite runs
// in; instead the test re-executes itself in a child process with
// MODERNC_SQLITE_OFD_LOCK=1 set.
//
// In the child (recognized by ofdChildEnvVar) it verifies the environment
// variable really switched the mode on — the positive control that keeps
// this coverage from silently reverting to POSIX mode — and returns true:
// the caller proceeds with the scenario. In the parent it runs the child,
// propagates its outcome, and returns false; a caller whose scenario holds
// under both locking modes then runs it in the parent process too, in
// whatever mode that process inherited.
func inOFDChild(t *testing.T) bool {
	t.Helper()
	if os.Getenv(ofdChildEnvVar) != "" {
		if !OFDLockingEnabled() {
			t.Fatal("child process: MODERNC_SQLITE_OFD_LOCK=1 did not enable OFD locking")
		}

		return true
	}

	reexecTest(t, append(os.Environ(), "MODERNC_SQLITE_OFD_LOCK=1", ofdChildEnvVar+"=1"))
	return false
}

// procLocks returns the /proc/locks lines describing the locks held on
// path's inode, skipping the calling test where /proc/locks cannot be read.
func procLocks(t *testing.T, path string) string {
	t.Helper()
	var st syscall.Stat_t
	if err := syscall.Stat(path, &st); err != nil {
		t.Fatal(err)
	}

	b, err := os.ReadFile("/proc/locks")
	if err != nil {
		t.Skipf("cannot read /proc/locks: %v", err)
	}

	ino := fmt.Sprint(st.Ino)
	var sb strings.Builder
	for _, ln := range strings.Split(string(b), "\n") {
		f := strings.Fields(ln)
		if len(f) < 6 {
			continue
		}

		if seg := strings.Split(f[5], ":"); seg[len(seg)-1] == ino {
			sb.WriteString(ln)
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

// TestOFDLockSurvivesOSClose verifies that closing an unrelated os.File
// descriptor pointing to the same inode does not strip the active SQLite
// database lock when OFD locking is enabled. Under standard POSIX inode
// locks (F_SETLK), close() on any descriptor drops all locks for that inode
// across the entire process — which is the default behavior and the hazard
// OFD locking exists to close. With Open File Description (OFD) locking
// (F_OFD_SETLK), locks are attached to the open file description,
// preventing accidental unlock on os.Close().
func TestOFDLockSurvivesOSClose(t *testing.T) {
	if !inOFDChild(t) {
		// Under POSIX locks the close() does strip the lock; the scenario
		// only holds in the OFD-enabled child.
		return
	}

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
// a read lock (which would cause SQLITE_BUSY). The invariant must hold under
// both locking modes: the scenario runs in an OFD-enabled child process and
// once more in this process's inherited mode.
//
// Authored by Jan Mercl (@cznic) in https://gitlab.com/cznic/libsqlite3/-/merge_requests/3#note_3726270793.
func TestOFDLockInterleavedReadersWrite(t *testing.T) {
	inOFDChild(t)

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

	// 1. Begin read transaction on c1 (acquires the kernel read lock).
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
	// If c1's kernel read lock leaked because c2's unlock was a no-op on c1's fd,
	// this write will fail with SQLITE_BUSY.
	if _, err := c2.ExecContext(ctx, "BEGIN IMMEDIATE; INSERT INTO t VALUES(2); COMMIT;"); err != nil {
		t.Fatalf("write after interleaved reads failed (leaked read lock): %v", err)
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
// Note that the probe descriptor's own close() strips POSIX locks, so probing
// is only meaningful with OFD locking enabled.
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
	if !inOFDChild(t) {
		// probeSharedRange's own close() would strip POSIX locks, so the
		// probe sequence only holds in the OFD-enabled child.
		return
	}

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
// alive then reads. The read must not go through the closed descriptor. The
// invariant must hold under both locking modes: the scenario runs in an
// OFD-enabled child process and once more in this process's inherited mode.
func TestOFDLockFailedFirstLock(t *testing.T) {
	inOFDChild(t)

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

// TestOFDLockingTooLate: once a database file lock has been attempted in
// this process the locking mode is frozen — OFDLocking refuses to change it,
// while no-change calls and queries keep working.
func TestOFDLockingTooLate(t *testing.T) {
	// Freeze the mode by taking a lock ourselves rather than relying on the
	// rest of the suite having run first.
	dbPath := filepath.Join(t.TempDir(), "frozen.db")
	db, err := sql.Open(driverName, dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if _, err := db.Exec("CREATE TABLE t(x)"); err != nil {
		t.Fatal(err)
	}

	cur := OFDLockingEnabled()
	switch _, err := OFDLocking(!cur); err {
	case ErrOFDLockingTooLate:
		// ok
	case ErrOFDLockingUnavailable:
		// The suite ran with OFD locking requested on a kernel or
		// filesystem that rejected it.
		t.Skip("OFD locks unavailable on this kernel/filesystem")
	default:
		t.Fatalf("OFDLocking(%v) after first lock: err = %v, want ErrOFDLockingTooLate", !cur, err)
	}

	if prev, err := OFDLocking(cur); prev != cur || err != nil {
		t.Fatalf("no-change OFDLocking(%v): prev = %v, err = %v, want %v, nil", cur, prev, err, cur)
	}
	if got := OFDLockingEnabled(); got != cur {
		t.Fatalf("OFDLockingEnabled() = %v, want %v", got, cur)
	}
}

// TestOFDLockingSetter exercises the Go call path end to end in a fresh
// child process with no MODERNC_SQLITE_OFD_LOCK in the environment: enabling
// before the first connection, the kernel-visible OFDLCK record as the
// positive control, and the freeze after the first lock.
func TestOFDLockingSetter(t *testing.T) {
	const childEnvVar = "MODERNC_SQLITE_TEST_OFD_SETTER_CHILD"
	if os.Getenv(childEnvVar) == "" {
		env := []string{childEnvVar + "=1"}
		for _, kv := range os.Environ() {
			if !strings.HasPrefix(kv, "MODERNC_SQLITE_OFD_LOCK=") {
				env = append(env, kv)
			}
		}
		reexecTest(t, env)
		return
	}

	if OFDLockingEnabled() {
		t.Fatal("OFD locking on by default")
	}
	if prev, err := OFDLocking(true); prev || err != nil {
		t.Fatalf("OFDLocking(true): prev = %v, err = %v, want false, nil", prev, err)
	}
	if !OFDLockingEnabled() {
		t.Fatal("OFDLockingEnabled() = false after OFDLocking(true)")
	}

	dbPath := filepath.Join(t.TempDir(), "setter.db")
	db, err := sql.Open(driverName, dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if _, err := db.Exec("CREATE TABLE t(x)"); err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	if _, err := tx.Exec("INSERT INTO t VALUES(1)"); err != nil {
		t.Fatal(err)
	}

	if !OFDLockingEnabled() {
		t.Skip("kernel or filesystem rejected OFD locks; POSIX fallback in effect")
	}

	// The write transaction must be holding OFDLCK, and no POSIX, locks.
	switch locks := procLocks(t, dbPath); {
	case !strings.Contains(locks, "OFDLCK"):
		t.Fatalf("no OFDLCK lock on the database file:\n%s", locks)
	case strings.Contains(locks, "POSIX"):
		t.Fatalf("unexpected POSIX lock on the database file:\n%s", locks)
	}

	if _, err := OFDLocking(false); err != ErrOFDLockingTooLate {
		t.Fatalf("OFDLocking(false) after first lock: err = %v, want ErrOFDLockingTooLate", err)
	}
	if prev, err := OFDLocking(true); !prev || err != nil {
		t.Fatalf("no-change OFDLocking(true): prev = %v, err = %v, want true, nil", prev, err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
