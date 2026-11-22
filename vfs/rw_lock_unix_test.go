// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package vfs

import (
	"bufio"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"testing"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// The tests in this file put real processes on one database file. Some of
// them go through a VFS over the host's files whose Locker takes the locks
// SQLite's own VFS takes, and some through SQLite's own VFS.

// SQLite's lock bytes, see os.h.
const (
	pendingByte  = 0x40000000
	reservedByte = pendingByte + 1
	sharedFirst  = pendingByte + 2
	sharedSize   = 510
)

// hostVFS is a VFS over files of the host, with locks across processes taken
// as POSIX record locks: one descriptor per database, shared by its Files
// and its Locker and closed with the last of them. It is for one
// registration in a process and for tests: POSIX locks cannot tell two
// Lockers of one process apart, so it gives a database one Locker and
// refuses a second, and a second hostVFS on the same files in one process
// would lose the locks of the first at its first close. With ownFD every
// File opens a descriptor of its own, which is the mistake the documentation
// of Locker warns of.
type hostVFS struct {
	dir     string
	ownFD   bool
	onWrite func() // runs before each write to a database, if set

	mu  sync.Mutex
	dbs map[string]*hostDB
}

type hostDB struct {
	name   string
	f      *os.File
	refs   int
	locker bool // its Locker is given out
}

// hostNoLocker hides the OpenLock of a hostVFS.
type hostNoLocker struct{ VFS }

func (v *hostVFS) db(name string, create bool) (*hostDB, error) {
	v.mu.Lock()
	defer v.mu.Unlock()
	d := v.dbs[name]
	if d == nil {
		flag := os.O_RDWR
		if create {
			flag |= os.O_CREATE
		}
		f, err := os.OpenFile(name, flag, 0o644)
		if err != nil {
			return nil, err
		}

		d = &hostDB{name: name, f: f}
		if v.dbs == nil {
			v.dbs = map[string]*hostDB{}
		}
		v.dbs[name] = d
	}
	d.refs++
	return d, nil
}

func (v *hostVFS) release(d *hostDB) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	if d.refs--; d.refs > 0 {
		return nil
	}

	delete(v.dbs, d.name)
	return d.f.Close()
}

func (v *hostVFS) Open(name string, flags OpenFlag) (File, OpenFlag, error) {
	if flags&OpenMainDB != 0 && !v.ownFD {
		d, err := v.db(name, flags&OpenCreate != 0)
		if err != nil {
			return nil, 0, err
		}

		return &hostFile{v: v, d: d, f: d.f}, flags, nil
	}

	flag := os.O_RDWR
	if flags&OpenCreate != 0 {
		flag |= os.O_CREATE
	}
	if flags&OpenExclusive != 0 {
		flag |= os.O_EXCL
	}
	f, err := os.OpenFile(name, flag, 0o644)
	if err != nil {
		return nil, 0, err
	}

	return &hostFile{v: v, f: f, db: flags&OpenMainDB != 0}, flags, nil
}

func (v *hostVFS) Delete(name string, syncDir bool) error { return os.Remove(name) }

func (v *hostVFS) Access(name string, flags AccessFlag) (bool, error) {
	switch _, err := os.Stat(name); {
	case err == nil:
		return true, nil
	case errors.Is(err, fs.ErrNotExist):
		return false, nil
	default:
		return false, err
	}
}

func (v *hostVFS) FullPathname(name string) (string, error) {
	if filepath.IsAbs(name) {
		return filepath.Clean(name), nil
	}

	return filepath.Join(v.dir, name), nil
}

func (v *hostVFS) OpenLock(name string) (Locker, error) {
	d, err := v.db(name, false)
	if err != nil {
		return nil, err
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	if d.locker {
		d.refs--
		return nil, fmt.Errorf("hostVFS: %s is open through another registration of this process", name)
	}

	d.locker = true
	return &hostLock{v: v, d: d}, nil
}

type hostFile struct {
	v  *hostVFS
	d  *hostDB // the shared descriptor of a database, or nil
	f  *os.File
	db bool // a database on a descriptor of its own
}

func (f *hostFile) ReadAt(p []byte, off int64) (int, error) { return f.f.ReadAt(p, off) }

func (f *hostFile) WriteAt(p []byte, off int64) (int, error) {
	if (f.d != nil || f.db) && f.v.onWrite != nil {
		f.v.onWrite()
	}
	return f.f.WriteAt(p, off)
}

func (f *hostFile) Close() error {
	if f.d != nil {
		return f.v.release(f.d)
	}

	return f.f.Close()
}

func (f *hostFile) Truncate(size int64) error { return f.f.Truncate(size) }
func (f *hostFile) Sync(SyncFlag) error       { return f.f.Sync() }

func (f *hostFile) Size() (int64, error) {
	fi, err := f.f.Stat()
	if err != nil {
		return 0, err
	}

	return fi.Size(), nil
}

// hostLock takes SQLite's byte range locks with fcntl, the way unixLock and
// posixUnlock of SQLite's os_unix.c do.
type hostLock struct {
	v *hostVFS
	d *hostDB
}

func (l *hostLock) set(typ int16, start, n int64) error {
	lk := syscall.Flock_t{Type: typ, Whence: io.SeekStart, Start: start, Len: n}
	err := syscall.FcntlFlock(l.d.f.Fd(), syscall.F_SETLK, &lk)
	if err == syscall.EAGAIN || err == syscall.EACCES {
		return ErrBusy
	}

	return err
}

func (l *hostLock) Lock(level LockLevel) error {
	switch level {
	case LockShared:
		if err := l.set(syscall.F_RDLCK, pendingByte, 1); err != nil {
			return err
		}

		err := l.set(syscall.F_RDLCK, sharedFirst, sharedSize)
		if uerr := l.set(syscall.F_UNLCK, pendingByte, 1); uerr != nil {
			// Not a refusal any more, whatever err is: a lock is left behind.
			return fmt.Errorf("releasing the pending byte: %v", uerr)
		}

		return err
	case LockReserved:
		return l.set(syscall.F_WRLCK, reservedByte, 1)
	case LockPending:
		return l.set(syscall.F_WRLCK, pendingByte, 1)
	case LockExclusive:
		return l.set(syscall.F_WRLCK, sharedFirst, sharedSize)
	}

	return fmt.Errorf("Lock(%d)", level)
}

func (l *hostLock) Unlock(level LockLevel) error {
	if level == LockNone {
		return l.set(syscall.F_UNLCK, 0, 0)
	}

	// A write lock turns into a read lock in one step.
	if err := l.set(syscall.F_RDLCK, sharedFirst, sharedSize); err != nil {
		return err
	}

	return l.set(syscall.F_UNLCK, pendingByte, 2)
}

func (l *hostLock) CheckReserved() (bool, error) {
	pid, err := lockHolder(l.d.f, syscall.F_WRLCK, reservedByte, 1)
	return pid != 0, err
}

func (l *hostLock) Close() error {
	// Whatever a failed call left behind goes now: the descriptor may live
	// on with a File, and the next Locker of the database must start clean.
	err := l.set(syscall.F_UNLCK, 0, 0)
	l.v.mu.Lock()
	l.d.locker = false
	l.v.mu.Unlock()
	return errors.Join(err, l.v.release(l.d))
}

// lockHolder returns the process that holds a lock which keeps the calling
// process from a lock of the given type on the bytes, or 0.
func lockHolder(f *os.File, typ int16, start, n int64) (pid int, err error) {
	lk := syscall.Flock_t{Type: typ, Whence: io.SeekStart, Start: start, Len: n}
	if err := syscall.FcntlFlock(f.Fd(), syscall.F_GETLK, &lk); err != nil {
		return 0, err
	}

	if lk.Type == syscall.F_UNLCK {
		return 0, nil
	}

	return int(lk.Pid), nil
}

const lockChildEnv = "MODERNC_SQLITE_VFS_TEST_LOCK_CHILD"

// lockChild is a second process: the test binary running
// TestRegisterLockerChild in the role its arguments say. Parent and child
// talk in lines: the child's start with "> " on its standard output.
type lockChild struct {
	t   *testing.T
	cmd *exec.Cmd
	in  io.WriteCloser
	out *bufio.Scanner

	said   bytes.Buffer // its standard output so far
	stderr bytes.Buffer // not to be read before the child has ended
}

func startLockChild(t *testing.T, args ...string) *lockChild {
	t.Helper()
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	c := &lockChild{t: t, cmd: exec.Command(exe, "-test.run=^TestRegisterLockerChild$", "-test.v")}
	c.cmd.Env = append(os.Environ(), lockChildEnv+"="+strings.Join(args, "\x1f"))
	c.cmd.Stderr = &c.stderr
	if c.in, err = c.cmd.StdinPipe(); err != nil {
		t.Fatal(err)
	}

	out, err := c.cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}

	c.out = bufio.NewScanner(out)
	if err := c.cmd.Start(); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		c.cmd.Process.Kill()
		c.cmd.Wait()
	})
	return c
}

// line returns the next line the child says, without the mark, and false
// when the child has closed its output.
func (c *lockChild) line() (string, bool) {
	for c.out.Scan() {
		s := c.out.Text()
		fmt.Fprintln(&c.said, s)
		if s, ok := strings.CutPrefix(s, "> "); ok {
			return s, true
		}
	}
	return "", false
}

// expect waits for the child to say a line that starts with want.
func (c *lockChild) expect(want string) string {
	c.t.Helper()
	s, ok := c.line()
	if !ok || !strings.HasPrefix(s, want) {
		c.t.Fatalf("child %v said %q, want %q; its output:\n%s", c.cmd.Args, s, want, c.output())
	}

	return s
}

func (c *lockChild) say(s string) {
	c.t.Helper()
	if _, err := fmt.Fprintln(c.in, s); err != nil {
		c.t.Fatalf("writing to the child: %v", err)
	}
}

// done waits for the child to end and fails the test if it failed.
func (c *lockChild) done() {
	c.t.Helper()
	for {
		if _, ok := c.line(); !ok {
			break
		}
	}
	if err := c.cmd.Wait(); err != nil {
		c.t.Fatalf("child %v: %v\n%s", c.cmd.Args, err, c.output())
	}
}

// output ends the child, if it still runs, and returns what it wrote.
func (c *lockChild) output() string {
	c.cmd.Process.Kill()
	c.cmd.Wait()
	return c.said.String() + c.stderr.String()
}

// openKind opens the database file at path through SQLite's own VFS
// ("default"), through a hostVFS ("host"), through a hostVFS without a
// LockOpener ("host-nolocker") or through a hostVFS that opens a descriptor
// for every File ("host-ownfd").
func openKind(t *testing.T, path, kind string, params ...string) *sql.DB {
	t.Helper()
	if kind == "default" {
		db, err := sql.Open("sqlite", "file:"+path+"?"+strings.Join(params, "&"))
		if err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() { db.Close() })
		return db
	}

	var v VFS = &hostVFS{dir: filepath.Dir(path), ownFD: kind == "host-ownfd"}
	if kind == "host-nolocker" {
		v = hostNoLocker{v}
	}
	return openRW(t, v, path, 1, params...)
}

func isBusy(err error) bool {
	var se *sqlite.Error
	return errors.As(err, &se) && se.Code()&0xff == sqlite3.SQLITE_BUSY
}

const (
	lockAccounts = 8
	bigRows      = 300
)

// TestRegisterLockerChild is the other process of TestRegisterLockerProcesses
// and does nothing when run on its own.
func TestRegisterLockerChild(t *testing.T) {
	env := os.Getenv(lockChildEnv)
	if env == "" {
		return
	}

	args := strings.Split(env, "\x1f")
	say := func(format string, a ...any) { fmt.Printf("> "+format+"\n", a...) }
	stdin := bufio.NewScanner(os.Stdin)
	hear := func() string {
		if !stdin.Scan() {
			t.Fatal("the parent is gone")
		}

		return stdin.Text()
	}
	switch role, path := args[0], args[1]; role {
	case "transfers": // kind, seed, count: move money when told to
		seed, _ := strconv.ParseInt(args[3], 10, 64)
		count, _ := strconv.Atoi(args[4])
		db := openKind(t, path, args[2], "_pragma=busy_timeout(10000)")
		if err := db.Ping(); err != nil {
			t.Fatal(E(err))
		}

		say("READY")
		hear()
		for _, tr := range runTransfers(t, []*sql.DB{db}, lockAccounts, 1, count, seed) {
			say("T %d %d %d", tr.src, tr.dst, tr.amt)
		}
	case "hold": // kind: begin to write, then open and close a second connection
		var v VFS = &hostVFS{dir: filepath.Dir(path), ownFD: args[2] == "host-ownfd"}
		if args[2] == "host-nolocker" {
			v = hostNoLocker{v}
		}
		name, fsys, err := Register(v)
		if err != nil {
			t.Fatal(err)
		}

		defer fsys.Close()
		dsn := "file:" + path + "?vfs=" + name
		db, err := sql.Open("sqlite", dsn)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()
		tx, err := db.Begin()
		if err != nil {
			t.Fatal(E(err))
		}

		if _, err := tx.Exec("UPDATE acct SET bal=bal+1 WHERE id=0"); err != nil {
			t.Fatal(E(err))
		}

		// A second connection of the registration reads and goes away. The
		// File it had is closed; the lock of the writer must stay.
		db2, err := sql.Open("sqlite", dsn)
		if err != nil {
			t.Fatal(err)
		}

		var n int
		if err := db2.QueryRow("SELECT count(*) FROM acct").Scan(&n); err != nil {
			t.Fatal(E(err))
		}

		if err := db2.Close(); err != nil {
			t.Fatal(err)
		}

		say("READY %d", os.Getpid())
		if hear() == "commit" {
			err = tx.Commit()
		} else {
			err = tx.Rollback()
		}
		if err != nil {
			t.Fatal(E(err))
		}

		say("DONE")
	case "crash": // write more than the cache holds, and wait to be killed
		db := openKind(t, path, "host")
		conn, err := db.Conn(t.Context())
		if err != nil {
			t.Fatal(err)
		}

		for _, s := range []string{"PRAGMA cache_size=5", "BEGIN", "UPDATE big SET v=v+1, pad=zeroblob(1000)"} {
			if _, err := conn.ExecContext(t.Context(), s); err != nil {
				t.Fatalf("%s: %s", s, E(err))
			}
		}
		say("READY %d", os.Getpid())
		hear()
	case "read": // one read through SQLite's own VFS, without waiting
		db := openKind(t, path, "default")
		var sum int
		switch err := db.QueryRow("SELECT sum(v) FROM big").Scan(&sum); {
		case err == nil:
			say("R %d", sum)
		case isBusy(err):
			say("R BUSY")
		default:
			say("R ERROR %s", E(err))
		}
	default:
		t.Fatalf("unknown role %q", role)
	}
}

// TestRegisterLockerProcesses puts several processes on one database file.
func TestRegisterLockerProcesses(t *testing.T) {
	// newDB makes the database through a hostVFS and returns its path.
	newDB := func(t *testing.T) string {
		t.Helper()
		path := filepath.Join(t.TempDir(), "test.db")
		name, fsys, err := Register(&hostVFS{dir: filepath.Dir(path)})
		if err != nil {
			t.Fatal(err)
		}

		defer fsys.Close()
		db, err := sql.Open("sqlite", "file:"+path+"?vfs="+name)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()
		mustExec(t, db, strings.Split(transferSchema, "; ")...)
		for i := 0; i < lockAccounts; i++ {
			mustExec(t, db, fmt.Sprintf("INSERT INTO acct VALUES(%d, 1000)", i))
		}
		mustExec(t, db, "CREATE TABLE big(id INTEGER PRIMARY KEY, v INTEGER, pad BLOB)",
			fmt.Sprintf("WITH RECURSIVE n(i) AS (SELECT 1 UNION ALL SELECT i+1 FROM n WHERE i<%d) INSERT INTO big SELECT i, i, randomblob(1000) FROM n", bigRows))
		return path
	}

	// Writers in several processes at once, on the hostVFS alone and mixed
	// with SQLite's own VFS: every committed transfer is in the database.
	for _, kinds := range [][]string{{"host", "host", "host"}, {"host", "default", "host", "default"}} {
		t.Run("writers "+strings.Join(kinds, " "), func(t *testing.T) {
			const count = 25
			path := newDB(t)
			var children []*lockChild
			for i, kind := range kinds {
				c := startLockChild(t, "transfers", path, kind, strconv.Itoa(100*(i+1)), strconv.Itoa(count))
				c.expect("READY")
				children = append(children, c)
			}
			for _, c := range children {
				c.say("go")
			}

			var (
				wg        sync.WaitGroup
				mu        sync.Mutex
				committed []transfer
			)
			for _, c := range children {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for {
						s, ok := c.line()
						if !ok {
							return
						}

						var tr transfer
						if _, err := fmt.Sscanf(s, "T %d %d %d", &tr.src, &tr.dst, &tr.amt); err == nil {
							mu.Lock()
							committed = append(committed, tr)
							mu.Unlock()
						}
					}
				}()
			}
			wg.Wait()
			for _, c := range children {
				c.done()
			}
			if len(committed) != len(kinds)*count {
				t.Fatalf("%d transfers committed, want %d", len(committed), len(kinds)*count)
			}

			checkTransfers(t, openKind(t, path, "default"), lockAccounts, committed)
		})
	}

	// immediate reports what BEGIN IMMEDIATE through the given VFS answers
	// while the child holds its write transaction open.
	immediate := func(t *testing.T, path, kind string) error {
		t.Helper()
		var v VFS = &hostVFS{dir: filepath.Dir(path)}
		dsn := "file:" + path
		if kind != "default" {
			name, fsys, err := Register(v)
			if err != nil {
				t.Fatal(err)
			}

			defer fsys.Close()
			dsn += "?vfs=" + name
		}
		db, err := sql.Open("sqlite", dsn)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()
		var n int
		if err := db.QueryRow("SELECT count(*) FROM acct").Scan(&n); err != nil || n != lockAccounts {
			t.Fatalf("reading next to the writer through the %s VFS: %d rows, %v", kind, n, E(err))
		}

		_, err = db.Exec("BEGIN IMMEDIATE; ROLLBACK")
		return err
	}

	// A process that writes through a hostVFS holds SQLite's RESERVED lock
	// where SQLite's own VFS looks for it, and keeps it when a File of the
	// database is closed in the same process.
	t.Run("excludes", func(t *testing.T) {
		path := newDB(t)
		probe, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}

		defer probe.Close()
		c := startLockChild(t, "hold", path, "host")
		var pid int
		fmt.Sscanf(c.expect("READY"), "READY %d", &pid)
		if got, err := lockHolder(probe, syscall.F_WRLCK, reservedByte, 1); err != nil || got != pid {
			t.Fatalf("SQLite's RESERVED byte is locked by process %d (%v), want the child %d", got, err, pid)
		}

		for _, kind := range []string{"default", "host"} {
			if err := immediate(t, path, kind); !isBusy(err) {
				t.Fatalf("BEGIN IMMEDIATE through the %s VFS next to the writer: %v, want SQLITE_BUSY", kind, E(err))
			}
		}
		c.say("commit")
		c.expect("DONE")
		c.done()
		for _, kind := range []string{"default", "host"} {
			if err := immediate(t, path, kind); err != nil {
				t.Fatalf("BEGIN IMMEDIATE through the %s VFS after the writer: %v", kind, E(err))
			}
		}
		if got := dump(t, openKind(t, path, "default"), "SELECT bal FROM acct WHERE id=0"); got != "1001|\n" {
			t.Fatalf("the writer's change: %q", got)
		}
	})

	// What Register and Locker warn of, shown on the same steps. Without a
	// LockOpener nothing is locked outside the process. With a descriptor
	// for every File, the POSIX locks of the process are gone as soon as
	// one of them is closed.
	for _, kind := range []string{"host-nolocker", "host-ownfd"} {
		t.Run("does not exclude, "+kind, func(t *testing.T) {
			path := newDB(t)
			probe, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}

			defer probe.Close()
			c := startLockChild(t, "hold", path, kind)
			c.expect("READY")
			if got, err := lockHolder(probe, syscall.F_WRLCK, reservedByte, 1); err != nil || got != 0 {
				t.Fatalf("SQLite's RESERVED byte is locked by process %d (%v), want nobody", got, err)
			}

			if err := immediate(t, path, "default"); err != nil {
				t.Fatalf("BEGIN IMMEDIATE next to a writer that holds no lock: %v", E(err))
			}

			c.say("rollback")
			c.expect("DONE")
			c.done()
		})
	}

	// A writer is killed after it has changed the database file and before
	// it commits. The next process to open the database finds nobody at
	// RESERVED, rolls the journal back under EXCLUSIVE, and keeps a reader
	// of SQLite's own VFS out while it does.
	t.Run("killed writer", func(t *testing.T) {
		path := newDB(t)
		before, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		wantSum := bigRows * (bigRows + 1) / 2
		probe, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}

		defer probe.Close()
		c := startLockChild(t, "crash", path)
		var pid int
		fmt.Sscanf(c.expect("READY"), "READY %d", &pid)
		if got, err := lockHolder(probe, syscall.F_RDLCK, sharedFirst, sharedSize); err != nil || got != pid {
			t.Fatalf("SQLite's shared bytes are write-locked by process %d (%v), want the child %d", got, err, pid)
		}

		during, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		if bytes.Equal(before, during) {
			t.Fatal("the writer has not changed the database file yet: nothing to roll back")
		}

		if fi, err := os.Stat(path + "-journal"); err != nil || fi.Size() == 0 {
			t.Fatalf("the writer's journal: %v", err)
		}

		if err := c.cmd.Process.Kill(); err != nil {
			t.Fatal(err)
		}

		c.cmd.Wait()
		if got, err := lockHolder(probe, syscall.F_RDLCK, sharedFirst, sharedSize); err != nil || got != 0 {
			t.Fatalf("after the kill the shared bytes are locked by process %d (%v)", got, err)
		}

		var (
			once   sync.Once
			reader string
		)
		v := &hostVFS{dir: filepath.Dir(path)}
		v.onWrite = func() {
			once.Do(func() {
				r := startLockChild(t, "read", path)
				reader, _ = r.line()
				r.done()
			})
		}
		db := openRW(t, v, path, 1)
		if got, want := dump(t, db, "SELECT sum(v), count(*) FROM big"), fmt.Sprintf("%d|%d|\n", wantSum, bigRows); got != want {
			t.Fatalf("after the rollback: %q, want %q", got, want)
		}

		once.Do(func() {})
		if reader != "R BUSY" {
			t.Fatalf("a reader during the rollback got %q, want R BUSY", reader)
		}

		integrity(t, db)
		if _, err := os.Stat(path + "-journal"); !errors.Is(err, fs.ErrNotExist) {
			t.Fatalf("the journal after the rollback: %v", err)
		}

		r := startLockChild(t, "read", path)
		if got, want := r.expect("R"), fmt.Sprintf("R %d", wantSum); got != want {
			t.Fatalf("a reader after the rollback got %q, want %q", got, want)
		}

		r.done()
	})
}
