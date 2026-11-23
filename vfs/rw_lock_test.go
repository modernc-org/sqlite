// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vfs

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// lockTable stands in for what keeps the Lockers of a database apart, as the
// kernel does for locks on a file. It also checks that the package uses a
// Locker the way the documentation of Locker says.
type lockTable struct {
	t  testing.TB
	mu sync.Mutex
	in map[string]map[*tableLock]bool // the Lockers of each database
}

func newLockTable(t testing.TB) *lockTable {
	return &lockTable{t: t, in: map[string]map[*tableLock]bool{}}
}

// locker returns a new Locker of the named database.
func (tab *lockTable) locker(name string) *tableLock {
	tab.mu.Lock()
	defer tab.mu.Unlock()
	l := &tableLock{tab: tab, name: name}
	if tab.in[name] == nil {
		tab.in[name] = map[*tableLock]bool{}
	}
	tab.in[name][l] = true
	return l
}

type tableLock struct {
	tab  *lockTable
	name string

	running atomic.Bool // a call is in progress

	// Under tab.mu.
	level  LockLevel
	closed bool
	calls  []string // every call with its result, until take

	// fail, if set, is asked before a call takes effect. With a non-nil
	// error the call returns it, after taking effect all the same if effect
	// is true.
	fail func(call string) (err error, effect bool)
	// enter, if set, runs first thing in every call.
	enter func(call string)
}

var levelNames = map[LockLevel]string{LockNone: "none", LockShared: "shared", LockReserved: "reserved", LockPending: "pending", LockExclusive: "exclusive"}

func (l *tableLock) do(call string, effect func() error) (err error) {
	if !l.running.CompareAndSwap(false, true) {
		l.tab.t.Errorf("%s: %s overlaps another call of the same Locker", l.name, call)
	}
	defer l.running.Store(false)
	if l.enter != nil {
		l.enter(call)
	}

	l.tab.mu.Lock()
	defer l.tab.mu.Unlock()
	defer func() {
		res := "ok"
		switch {
		case errors.Is(err, ErrBusy):
			res = "busy"
		case err != nil:
			res = "error"
		}
		l.calls = append(l.calls, call+"="+res)
	}()
	if l.closed {
		l.tab.t.Errorf("%s: %s after Close", l.name, call)
	}
	if l.fail != nil {
		if ferr, eff := l.fail(call); ferr != nil {
			if eff {
				effect()
			}
			return ferr
		}
	}
	return effect()
}

// over returns the highest level another Locker of the database holds.
func (l *tableLock) over() (max LockLevel) {
	for o := range l.tab.in[l.name] {
		if o != l && o.level > max {
			max = o.level
		}
	}
	return max
}

func (l *tableLock) Lock(level LockLevel) error {
	return l.do("Lock("+levelNames[level]+")", func() error {
		var free bool
		switch {
		case level == LockShared && l.level == LockNone:
			free = l.over() < LockPending
		case level == LockReserved && l.level == LockShared:
			free = l.over() < LockReserved
		case level == LockPending && l.level == LockReserved:
			free = true
		case level == LockExclusive && (l.level == LockPending || l.level == LockShared):
			free = l.over() == LockNone
		default:
			l.tab.t.Errorf("%s: Lock(%s) by a Locker that holds %s", l.name, levelNames[level], levelNames[l.level])
			return errors.New("lock table: not a step of the protocol")
		}
		if !free {
			return ErrBusy
		}

		l.level = level
		return nil
	})
}

func (l *tableLock) Unlock(level LockLevel) error {
	return l.do("Unlock("+levelNames[level]+")", func() error {
		if level != LockNone && (level != LockShared || l.level < LockShared) {
			l.tab.t.Errorf("%s: Unlock(%s) by a Locker that holds %s", l.name, levelNames[level], levelNames[l.level])
			return errors.New("lock table: not a step of the protocol")
		}

		l.level = level
		return nil
	})
}

func (l *tableLock) CheckReserved() (held bool, err error) {
	err = l.do("CheckReserved", func() error {
		held = l.over() >= LockReserved
		return nil
	})
	return held, err
}

func (l *tableLock) Close() error {
	return l.do("Close", func() error {
		l.level, l.closed = LockNone, true
		delete(l.tab.in[l.name], l)
		return nil
	})
}

// take returns the calls made since the last take.
func (l *tableLock) take() []string {
	l.tab.mu.Lock()
	defer l.tab.mu.Unlock()
	c := l.calls
	l.calls = nil
	return c
}

func (l *tableLock) holds() LockLevel {
	l.tab.mu.Lock()
	defer l.tab.mu.Unlock()
	return l.level
}

// lockVFS adds the Lockers of a lockTable to a VFS.
type lockVFS struct {
	VFS
	tab      *lockTable
	openLock func(name string) (Locker, error) // in place of the table, if set

	mu    sync.Mutex
	locks map[string][]*tableLock // every Locker given out, by database
}

func (v *lockVFS) OpenLock(name string) (Locker, error) {
	if v.openLock != nil {
		return v.openLock(name)
	}

	l := v.tab.locker(name)
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.locks == nil {
		v.locks = map[string][]*tableLock{}
	}
	v.locks[name] = append(v.locks[name], l)
	return l, nil
}

// last returns the Locker most recently given out for the database.
func (v *lockVFS) last(name string) *tableLock {
	v.mu.Lock()
	defer v.mu.Unlock()
	if n := len(v.locks[name]); n != 0 {
		return v.locks[name][n-1]
	}

	return nil
}

// lockRig drives the lock methods of files open on test.db through one
// registration whose VFS is a LockOpener, and compares the calls each step
// makes on the Locker.
type lockRig struct {
	t     *testing.T
	b     *bridge
	v     *lockVFS
	files map[string]uintptr

	readAt func() // runs inside every ReadAt of a file of the rig, if set
}

func newLockRig(t *testing.T, tab *lockTable, names ...string) *lockRig {
	t.Helper()
	r := &lockRig{t: t, files: map[string]uintptr{}}
	r.v = &lockVFS{tab: tab, VFS: &fakeVFS{open: func(name string, flags OpenFlag) (File, OpenFlag, error) {
		return &fakeFile{readAt: func(p []byte, off int64) (int, error) {
			if r.readAt != nil {
				r.readAt()
			}
			return 0, io.EOF
		}}, flags, nil
	}}}
	r.b = newBridge(t, r.v)
	for _, n := range names {
		r.files[n] = r.b.mustOpen("test.db", mainDB)
	}
	t.Cleanup(func() {
		for _, f := range r.files {
			rwClose(r.b.tls, f)
		}
	})
	return r
}

func (r *lockRig) calls(step string, want []string) {
	r.t.Helper()
	l := r.v.last("test.db")
	if l == nil {
		if len(want) != 0 {
			r.t.Fatalf("%s: no Locker, want calls %v", step, want)
		}
		return
	}

	if got := l.take(); strings.Join(got, " ") != strings.Join(want, " ") {
		r.t.Fatalf("%s: calls on the Locker %v, want %v", step, got, want)
	}
}

func (r *lockRig) lock(f string, level, want int32, calls ...string) {
	r.t.Helper()
	step := fmt.Sprintf("xLock(%s, %d)", f, level)
	if rc := rwLock(r.b.tls, r.files[f], level); rc != want {
		r.t.Fatalf("%s: rc %d, want %d", step, rc, want)
	}
	r.calls(step, calls)
}

func (r *lockRig) unlock(f string, level, want int32, calls ...string) {
	r.t.Helper()
	step := fmt.Sprintf("xUnlock(%s, %d)", f, level)
	if rc := rwUnlock(r.b.tls, r.files[f], level); rc != want {
		r.t.Fatalf("%s: rc %d, want %d", step, rc, want)
	}
	r.calls(step, calls)
}

func (r *lockRig) reserved(f string, wantRC, want int32, calls ...string) {
	r.t.Helper()
	step := fmt.Sprintf("xCheckReservedLock(%s)", f)
	p, s := r.b.buf(4, 0xff)
	if rc := rwCheckReservedLock(r.b.tls, r.files[f], p); rc != wantRC {
		r.t.Fatalf("%s: rc %d, want %d", step, rc, wantRC)
	}
	if got := *(*int32)(unsafe.Pointer(&s[0])); got != want {
		r.t.Fatalf("%s: %d, want %d", step, got, want)
	}
	r.calls(step, calls)
}

func (r *lockRig) close(f string, calls ...string) {
	r.t.Helper()
	l := r.v.last("test.db")
	if rc := rwClose(r.b.tls, r.files[f]); rc != sqlite3.SQLITE_OK {
		r.t.Fatalf("xClose(%s): rc %d", f, rc)
	}
	delete(r.files, f)
	if got := l.take(); strings.Join(got, " ") != strings.Join(calls, " ") {
		r.t.Fatalf("xClose(%s): calls on the Locker %v, want %v", f, got, calls)
	}
}

// touch checks whether the database can be read and written through f, by
// the result codes of xFileSize, xRead, xWrite, xTruncate and xSync.
func (r *lockRig) touch(f string, usable bool) {
	r.t.Helper()
	p, _ := r.b.buf(8, 0)
	got := [5]int32{
		rwFileSize(r.b.tls, r.files[f], p),
		rwRead(r.b.tls, r.files[f], p, 8, 0),
		rwWrite(r.b.tls, r.files[f], p, 8, 0),
		rwTruncate(r.b.tls, r.files[f], 0),
		rwSync(r.b.tls, r.files[f], sqlite3.SQLITE_SYNC_NORMAL),
	}
	// The fake file is empty, so a read that is let through is a short one.
	want := [5]int32{sqlite3.SQLITE_OK, sqlite3.SQLITE_IOERR_SHORT_READ, sqlite3.SQLITE_OK, sqlite3.SQLITE_OK, sqlite3.SQLITE_OK}
	if !usable {
		want = [5]int32{sqlite3.SQLITE_IOERR_FSTAT, sqlite3.SQLITE_IOERR_READ, sqlite3.SQLITE_IOERR_WRITE, sqlite3.SQLITE_IOERR_TRUNCATE, sqlite3.SQLITE_IOERR_FSYNC}
	}
	if got != want {
		r.t.Fatalf("size, read, write, truncate, sync through %s: rc %v, want %v", f, got, want)
	}
}

const (
	lkOK    = sqlite3.SQLITE_OK
	lkBusy  = sqlite3.SQLITE_BUSY
	lkIOERR = sqlite3.SQLITE_IOERR_LOCK
	ulIOERR = sqlite3.SQLITE_IOERR_UNLOCK
	toNone  = lockNone
)

// TestRegisterLocker walks the lock protocol over files of one registration
// whose VFS is a LockOpener and checks, step by step, what reaches the
// Locker: one lock for all the readers, the steps of a writer in SQLite's
// order, and nothing when only the files of the registration are concerned.
func TestRegisterLocker(t *testing.T) {
	t.Run("readers", func(t *testing.T) {
		r := newLockRig(t, newLockTable(t), "A", "B")
		r.lock("A", lockShared, lkOK, "Lock(shared)=ok")
		r.lock("B", lockShared, lkOK) // the registration is in already
		r.lock("B", lockShared, lkOK)
		r.unlock("A", toNone, lkOK) // B stays
		r.unlock("A", toNone, lkOK)
		r.unlock("B", toNone, lkOK, "Unlock(none)=ok")
		r.lock("B", lockShared, lkOK, "Lock(shared)=ok")
		r.close("B", "Unlock(none)=ok") // closing lets go
		r.close("A", "Close=ok")
	})

	t.Run("writer", func(t *testing.T) {
		r := newLockRig(t, newLockTable(t), "A", "B", "C")
		r.lock("A", lockShared, lkOK, "Lock(shared)=ok")
		r.reserved("B", lkOK, 0, "CheckReserved=ok")
		r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
		r.reserved("B", lkOK, 1) // known without asking
		r.lock("B", lockShared, lkOK)
		r.lock("B", lockReserved, lkBusy) // refused here, the Locker is not asked
		// PENDING goes out before the readers of the registration are counted.
		r.lock("A", lockExclusive, lkBusy, "Lock(pending)=ok")
		r.lock("C", lockShared, lkBusy)
		r.lock("A", lockExclusive, lkBusy) // PENDING is held, B still reads
		r.unlock("B", toNone, lkOK)
		r.lock("A", lockExclusive, lkOK, "Lock(exclusive)=ok")
		r.unlock("A", lockShared, lkOK, "Unlock(shared)=ok")
		r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
		r.unlock("A", toNone, lkOK, "Unlock(none)=ok") // a rollback, straight down
		// A writer that leaves while a reader stays takes the registration
		// down to SHARED.
		r.lock("B", lockShared, lkOK, "Lock(shared)=ok")
		r.lock("A", lockShared, lkOK)
		r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
		r.unlock("A", toNone, lkOK, "Unlock(shared)=ok")
		r.close("A")
		r.close("B", "Unlock(none)=ok")
		r.close("C", "Close=ok")
	})

	t.Run("another locker", func(t *testing.T) {
		tab := newLockTable(t)
		r := newLockRig(t, tab, "A", "B")
		other := tab.locker("test.db") // another process
		step := func(err error) {
			t.Helper()
			if err != nil {
				t.Fatalf("the other locker: %v", err)
			}
		}

		// A reader outside: the writer gets as far as PENDING and stays there.
		step(other.Lock(LockShared))
		r.lock("A", lockShared, lkOK, "Lock(shared)=ok")
		r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
		r.lock("A", lockExclusive, lkBusy, "Lock(pending)=ok", "Lock(exclusive)=busy")
		r.lock("B", lockShared, lkBusy) // PENDING, in the registration too
		r.lock("A", lockExclusive, lkBusy, "Lock(exclusive)=busy")
		if got := r.v.last("test.db").holds(); got != LockPending {
			t.Fatalf("after a refused EXCLUSIVE the Locker holds %s, want pending", levelNames[got])
		}
		step(other.Unlock(LockNone))
		r.lock("A", lockExclusive, lkOK, "Lock(exclusive)=ok")
		if err := other.Lock(LockShared); !errors.Is(err, ErrBusy) {
			t.Fatalf("the other locker reads next to EXCLUSIVE: %v", err)
		}
		r.unlock("A", toNone, lkOK, "Unlock(none)=ok")

		// A writer outside: no second writer, and its journal is not hot.
		step(other.Lock(LockShared))
		step(other.Lock(LockReserved))
		r.lock("A", lockShared, lkOK, "Lock(shared)=ok")
		r.reserved("A", lkOK, 1, "CheckReserved=ok")
		r.lock("A", lockReserved, lkBusy, "Lock(reserved)=busy")
		r.lock("B", lockShared, lkOK) // A is a reader as before
		r.lock("A", lockExclusive, lkBusy)
		r.unlock("B", toNone, lkOK)
		// Alone in the registration, A goes for EXCLUSIVE as the rollback of
		// a hot journal does: no PENDING, and SHARED kept when refused.
		r.lock("A", lockExclusive, lkBusy, "Lock(exclusive)=busy")
		r.lock("B", lockShared, lkOK)
		r.unlock("B", toNone, lkOK)
		r.reserved("A", lkOK, 1, "CheckReserved=ok")

		// The writer outside closes in: readers are refused from PENDING on.
		r.unlock("A", toNone, lkOK, "Unlock(none)=ok")
		step(other.Lock(LockPending))
		r.lock("A", lockShared, lkBusy, "Lock(shared)=busy")
		r.lock("A", lockShared, lkBusy, "Lock(shared)=busy") // asked again each time
		step(other.Lock(LockExclusive))
		r.lock("A", lockShared, lkBusy, "Lock(shared)=busy")
		step(other.Unlock(LockNone))

		// The other one died with its journal behind: nobody holds RESERVED,
		// so the first to come rolls it back under EXCLUSIVE.
		r.lock("A", lockShared, lkOK, "Lock(shared)=ok")
		r.reserved("A", lkOK, 0, "CheckReserved=ok")
		r.lock("A", lockExclusive, lkOK, "Lock(exclusive)=ok")
		r.reserved("B", lkOK, 1)
		r.unlock("A", lockShared, lkOK, "Unlock(shared)=ok")
		r.unlock("A", toNone, lkOK, "Unlock(none)=ok")
	})

	// With nolock=1 SQLite takes no lock and still asks whether a journal it
	// finds is in use.
	t.Run("no lock held", func(t *testing.T) {
		tab := newLockTable(t)
		r := newLockRig(t, tab, "A")
		r.reserved("A", lkOK, 0, "CheckReserved=ok")
		other := tab.locker("test.db")
		if err := errors.Join(other.Lock(LockShared), other.Lock(LockReserved)); err != nil {
			t.Fatal(err)
		}

		r.reserved("A", lkOK, 1, "CheckReserved=ok")
	})

	// Two registrations of one storage are kept apart like two processes,
	// and without a LockOpener they are not, as Register says.
	t.Run("two registrations", func(t *testing.T) {
		for _, withLocker := range []bool{true, false} {
			tab := newLockTable(t)
			var v1, v2 VFS = &fakeVFS{}, &fakeVFS{}
			if withLocker {
				v1, v2 = &lockVFS{VFS: v1, tab: tab}, &lockVFS{VFS: v2, tab: tab}
			}
			b1, b2 := newBridge(t, v1), newBridge(t, v2)
			a, x := b1.mustOpen("test.db", mainDB), b2.mustOpen("test.db", mainDB)
			want := map[bool]int32{true: lkBusy, false: lkOK}[withLocker]
			for i, s := range []struct {
				b     *bridge
				f     uintptr
				level int32
				want  int32
			}{
				{b1, a, lockShared, lkOK},
				{b1, a, lockReserved, lkOK},
				{b2, x, lockShared, lkOK},
				{b2, x, lockReserved, want}, // a second writer
				{b1, a, lockExclusive, want},
			} {
				if rc := rwLock(s.b.tls, s.f, s.level); rc != s.want {
					t.Fatalf("LockOpener %v, step %d: rc %d, want %d", withLocker, i, rc, s.want)
				}
			}
			for _, c := range []struct {
				b *bridge
				f uintptr
			}{{b1, a}, {b2, x}} {
				if rc := rwClose(c.b.tls, c.f); rc != lkOK {
					t.Fatalf("xClose: rc %d", rc)
				}
			}
		}
	})
}

// TestRegisterLockerErrors checks what an error of the Locker does. ErrBusy
// means busy from Lock only. After any other failure the package does not
// know what the Locker holds: it refuses locks and I/O for as long as files
// of the registration think they hold the database, and it has the Locker
// let go before it relies on it again.
func TestRegisterLockerErrors(t *testing.T) {
	boom := errors.New("boom")
	// failing makes the named call of the rig's Locker fail once.
	failing := func(r *lockRig, call string, err error, effect bool) {
		l := r.v.last("test.db")
		l.fail = func(c string) (error, bool) {
			if c != call {
				return nil, false
			}

			l.fail = nil
			return err, effect
		}
	}

	t.Run("first reader", func(t *testing.T) {
		r := newLockRig(t, newLockTable(t), "A", "B")
		failing(r, "Lock(shared)", boom, true) // it took the lock and failed
		r.lock("A", lockShared, lkIOERR, "Lock(shared)=error")
		r.touch("A", true) // nobody thinks to hold the database
		// SQLite unlocks after a failed lock: the Locker lets go then, and
		// not only when the next reader comes.
		r.unlock("A", toNone, lkOK, "Unlock(none)=ok")
		if got := r.v.last("test.db").holds(); got != LockNone {
			t.Fatalf("after the unlock the Locker holds %s", levelNames[got])
		}
		r.unlock("A", toNone, lkOK) // once
		r.lock("B", lockShared, lkOK, "Lock(shared)=ok")
		r.unlock("B", toNone, lkOK, "Unlock(none)=ok")
		// Without that unlock, the Locker lets go before it is asked again.
		failing(r, "Lock(shared)", boom, true)
		r.lock("A", lockShared, lkIOERR, "Lock(shared)=error")
		r.lock("B", lockShared, lkOK, "Unlock(none)=ok", "Lock(shared)=ok")
		r.unlock("B", toNone, lkOK, "Unlock(none)=ok")
		failing(r, "Lock(shared)", boom, false)
		r.lock("A", lockShared, lkIOERR, "Lock(shared)=error")
		failing(r, "Unlock(none)", boom, false)
		r.lock("A", lockShared, lkIOERR, "Unlock(none)=error")
		r.lock("A", lockShared, lkOK, "Unlock(none)=ok", "Lock(shared)=ok")
	})

	for _, call := range []string{"Lock(reserved)", "Lock(pending)", "Lock(exclusive)"} {
		t.Run("writer, "+call, func(t *testing.T) {
			r := newLockRig(t, newLockTable(t), "A", "B")
			r.lock("A", lockShared, lkOK, "Lock(shared)=ok")
			r.lock("B", lockShared, lkOK)
			failing(r, call, boom, false)
			switch call {
			case "Lock(reserved)":
				r.lock("A", lockReserved, lkIOERR, "Lock(reserved)=error")
			case "Lock(pending)":
				r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
				r.lock("A", lockExclusive, lkIOERR, "Lock(pending)=error")
			default:
				r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
				r.unlock("B", toNone, lkOK)
				r.lock("A", lockExclusive, lkIOERR, "Lock(pending)=ok", "Lock(exclusive)=error")
				r.lock("B", lockShared, lkIOERR)
			}
			// Readers are in and what protects them is not known.
			r.touch("A", false)
			r.touch("B", false)
			r.lock("A", lockExclusive, lkIOERR)
			if call == "Lock(reserved)" {
				r.lock("A", lockReserved, lkIOERR)
				r.unlock("A", toNone, lkOK) // B is in still
				r.touch("B", false)
				r.unlock("B", toNone, lkOK, "Unlock(none)=ok")
			} else {
				// The writer steps down first. The Locker is not asked for
				// SHARED while nothing is known of what it holds: only the
				// last file out has it let go.
				r.unlock("A", lockShared, ulIOERR)
				r.touch("A", false)
				r.unlock("B", toNone, lkOK)
				r.unlock("A", toNone, lkOK, "Unlock(none)=ok")
			}
			r.touch("A", true)
			r.lock("B", lockShared, lkOK, "Lock(shared)=ok")
			r.touch("B", true)
		})
	}

	t.Run("unlock to shared", func(t *testing.T) {
		tab := newLockTable(t)
		r := newLockRig(t, tab, "A", "B")
		r.lock("A", lockShared, lkOK, "Lock(shared)=ok")
		r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
		r.lock("B", lockShared, lkOK)
		// The downgrade fails and the Locker has let go altogether, as one
		// does that unlocks and locks again: another writer can come in.
		l := r.v.last("test.db")
		l.fail = func(c string) (error, bool) {
			if c != "Unlock(shared)" {
				return nil, false
			}

			l.fail, l.level = nil, LockNone
			return boom, false
		}
		r.unlock("A", lockShared, ulIOERR, "Unlock(shared)=error")
		other := tab.locker("test.db")
		if err := errors.Join(other.Lock(LockShared), other.Lock(LockReserved), other.Lock(LockPending), other.Lock(LockExclusive)); err != nil {
			t.Fatalf("the other locker: %v", err)
		}

		// A and B think they read under SHARED. They must not get a byte.
		r.touch("A", false)
		r.touch("B", false)
		r.lock("A", lockReserved, lkIOERR)
		r.unlock("A", toNone, lkOK)
		r.touch("B", false)
		r.unlock("B", toNone, lkOK, "Unlock(none)=ok")
		r.touch("B", true)
		r.lock("B", lockShared, lkBusy, "Lock(shared)=busy") // the other one writes
		if err := other.Unlock(LockNone); err != nil {
			t.Fatal(err)
		}

		r.lock("B", lockShared, lkOK, "Lock(shared)=ok")

		// A writer that leaves while its state is not known does not ask
		// for SHARED: nothing is known until everybody is out.
		r.lock("A", lockShared, lkOK)
		r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
		failing(r, "Unlock(shared)", boom, true)
		r.unlock("A", lockShared, ulIOERR, "Unlock(shared)=error")
		r.touch("B", false)
		r.unlock("A", toNone, lkOK)
		failing(r, "Unlock(none)", boom, false)
		r.unlock("B", toNone, ulIOERR, "Unlock(none)=error")
		r.touch("B", true) // nobody is in
		r.lock("A", lockShared, lkOK, "Unlock(none)=ok", "Lock(shared)=ok")
	})

	t.Run("busy is busy from Lock only", func(t *testing.T) {
		r := newLockRig(t, newLockTable(t), "A")
		wrapped := fmt.Errorf("lock server: %w", ErrBusy)
		failing(r, "Lock(shared)", wrapped, false)
		r.lock("A", lockShared, lkBusy, "Lock(shared)=busy")
		r.lock("A", lockShared, lkOK, "Lock(shared)=ok") // a refusal leaves nothing to clear up
		failing(r, "CheckReserved", wrapped, false)
		r.reserved("A", sqlite3.SQLITE_IOERR_CHECKRESERVEDLOCK, 0, "CheckReserved=busy")
		r.lock("A", lockReserved, lkOK, "Lock(reserved)=ok")
		failing(r, "Unlock(shared)", wrapped, false)
		r.unlock("A", lockShared, ulIOERR, "Unlock(shared)=busy")
		r.touch("A", false)
		failing(r, "Unlock(none)", wrapped, false)
		r.unlock("A", toNone, ulIOERR, "Unlock(none)=busy")
		// Closing asks once more before the Locker is closed.
		r.close("A", "Unlock(none)=ok", "Close=ok")
	})

	t.Run("OpenLock", func(t *testing.T) {
		var closedFiles, closedLocks, opens int
		tab := newLockTable(t)
		v := &lockVFS{tab: tab, VFS: &fakeVFS{open: func(name string, flags OpenFlag) (File, OpenFlag, error) {
			return &fakeFile{close: func() error { closedFiles++; return nil }}, flags, nil
		}}}
		v.openLock = func(name string) (Locker, error) {
			if opens++; name != "test.db" {
				t.Errorf("OpenLock(%q), want the name FullPathname gave", name)
			}
			return closerLocker{func() { closedLocks++ }}, fmt.Errorf("no lock: %w", ErrBusy)
		}
		b := newBridge(t, v)
		if _, rc := b.open("test.db", mainDB); rc != sqlite3.SQLITE_CANTOPEN {
			t.Fatalf("xOpen with a failing OpenLock: rc %d, want SQLITE_CANTOPEN", rc)
		}

		if closedFiles != 1 || closedLocks != 1 {
			t.Fatalf("after a failed OpenLock %d file(s) and %d Locker(s) were closed, want 1 and 1", closedFiles, closedLocks)
		}

		e := getObject(b.fsys.fsHandle).(*rwEntry)
		if len(e.locks) != 0 || e.open.Load() != 0 {
			t.Fatalf("after a failed OpenLock: %d lock state(s), %d open file(s)", len(e.locks), e.open.Load())
		}

		// The next open asks again; a nil Locker means the registration is
		// on its own, and other files than databases are not asked about.
		v.openLock = func(name string) (Locker, error) { opens++; return nil, nil }
		f1, f2 := b.mustOpen("test.db", mainDB), b.mustOpen("test.db", mainDB)
		j := b.mustOpen("test.db-journal", sqlite3.SQLITE_OPEN_READWRITE|sqlite3.SQLITE_OPEN_CREATE|sqlite3.SQLITE_OPEN_MAIN_JOURNAL)
		if opens != 2 {
			t.Fatalf("OpenLock was called %d times, want 2: once failing, once for the two files of the database", opens)
		}

		for i, s := range []struct {
			f           uintptr
			level, want int32
		}{
			{f1, lockShared, lkOK}, {f1, lockReserved, lkOK}, {f2, lockShared, lkOK}, {f2, lockReserved, lkBusy}, {f1, lockExclusive, lkBusy},
		} {
			if rc := rwLock(b.tls, s.f, s.level); rc != s.want {
				t.Fatalf("nil Locker, step %d: rc %d, want %d", i, rc, s.want)
			}
		}
		for _, f := range []uintptr{f1, f2, j} {
			rwClose(b.tls, f)
		}

		// One Locker per database for all its files, closed after the last.
		v.openLock = nil
		f1, f2 = b.mustOpen("test.db", mainDB), b.mustOpen("test.db", mainDB)
		o := b.mustOpen("other.db", mainDB)
		if l := v.last("test.db"); len(v.locks["test.db"]) != 1 || len(v.locks["other.db"]) != 1 || l == v.last("other.db") {
			t.Fatalf("Lockers given out: %d for test.db, %d for other.db, want one each", len(v.locks["test.db"]), len(v.locks["other.db"]))
		}

		l := v.last("test.db")
		rwLock(b.tls, f2, lockShared)
		rwClose(b.tls, f1)
		if got := strings.Join(l.take(), " "); got != "Lock(shared)=ok" {
			t.Fatalf("calls after closing one of two files: %q", got)
		}

		rwClose(b.tls, f2)
		if got := strings.Join(l.take(), " "); got != "Unlock(none)=ok Close=ok" {
			t.Fatalf("calls when the last file closes: %q", got)
		}

		// Opened again, the database gets a new Locker.
		f1 = b.mustOpen("test.db", mainDB)
		if len(v.locks["test.db"]) != 2 {
			t.Fatalf("%d Locker(s) for test.db after reopening, want 2", len(v.locks["test.db"]))
		}

		rwClose(b.tls, f1)
		rwClose(b.tls, o)
	})
}

// closerLocker is a Locker of which only Close is expected to be called.
type closerLocker struct{ closed func() }

func (closerLocker) Lock(LockLevel) error         { panic("Lock") }
func (closerLocker) Unlock(LockLevel) error       { panic("Unlock") }
func (closerLocker) CheckReserved() (bool, error) { panic("CheckReserved") }
func (c closerLocker) Close() error               { c.closed(); return nil }

// TestRegisterLockerPanic checks what a panic of a LockOpener or of a Locker
// leaves behind. The package does not recover it, and the database it
// happened on is in no known state; but no mutex of the package stays locked,
// so a program that recovers the panic can go on using the registration.
func TestRegisterLockerPanic(t *testing.T) {
	// panics runs f and reports whether it panicked.
	panics := func(f func()) (panicked bool) {
		defer func() { panicked = recover() != nil }()
		f()
		return false
	}
	// free fails the test if a mutex is still locked, and then uses the
	// registration: another database is opened, locked and closed on it.
	free := func(t *testing.T, b *bridge, ls *lockState) {
		t.Helper()
		e := getObject(b.fsys.fsHandle).(*rwEntry)
		if !e.mu.TryLock() {
			e.mu.Unlock() // or the clean-up of the test would wait for it
			t.Fatal("the mutex of the registration is still locked")
		}

		e.mu.Unlock()
		if ls != nil {
			if !ls.mu.TryLock() {
				ls.mu.Unlock() // as above
				t.Fatal("the mutex of the database is still locked")
			}

			ls.mu.Unlock()
			if !ls.io.TryLock() {
				t.Fatal("reads and writes of the database are still kept out")
			}

			ls.io.Unlock()
		}
		o := b.mustOpen("other.db", mainDB)
		if rc := rwLock(b.tls, o, lockShared); rc != lkOK {
			t.Fatalf("xLock on another database: rc %d", rc)
		}

		if rc := rwClose(b.tls, o); rc != lkOK {
			t.Fatalf("xClose of another database: rc %d", rc)
		}
	}

	t.Run("OpenLock", func(t *testing.T) {
		tab := newLockTable(t)
		v := &lockVFS{VFS: &fakeVFS{}, tab: tab}
		first := true
		v.openLock = func(name string) (Locker, error) {
			if first {
				first = false
				panic("OpenLock")
			}

			return tab.locker(name), nil
		}
		b := newBridge(t, v)
		if !panics(func() { b.open("test.db", mainDB) }) {
			t.Fatal("xOpen did not panic")
		}

		free(t, b, nil)
		f := b.mustOpen("test.db", mainDB) // and the database itself, again
		if rc := rwClose(b.tls, f); rc != lkOK {
			t.Fatalf("xClose: rc %d", rc)
		}
	})

	for _, call := range []string{"Lock(shared)", "Lock(reserved)", "Unlock(shared)", "Unlock(none)", "CheckReserved", "Close"} {
		t.Run(call, func(t *testing.T) {
			r := newLockRig(t, newLockTable(t), "A", "B")
			l := r.v.last("test.db")
			ls := getObject(r.b.fsys.fsHandle).(*rwEntry).locks["test.db"]
			a := r.files["A"]
			p, _ := r.b.buf(4, 0)
			var before []int32 // levels A takes first
			var step func()
			switch call {
			case "Lock(shared)":
				step = func() { rwLock(r.b.tls, a, lockShared) }
			case "Lock(reserved)":
				before, step = []int32{lockShared}, func() { rwLock(r.b.tls, a, lockReserved) }
			case "Unlock(shared)":
				before, step = []int32{lockShared, lockReserved}, func() { rwUnlock(r.b.tls, a, lockShared) }
			case "Unlock(none)":
				before, step = []int32{lockShared}, func() { rwUnlock(r.b.tls, a, lockNone) }
			case "CheckReserved":
				step = func() { rwCheckReservedLock(r.b.tls, a, p) }
			case "Close":
				rwClose(r.b.tls, r.files["B"])
				delete(r.files, "B")
				delete(r.files, "A")
				step = func() { rwClose(r.b.tls, a) }
			}
			for _, level := range before {
				if rc := rwLock(r.b.tls, a, level); rc != lkOK {
					t.Fatalf("xLock(A, %d): rc %d", level, rc)
				}
			}
			l.enter = func(c string) {
				if c == call {
					l.enter = nil
					panic(call)
				}
			}
			if !panics(step) {
				t.Fatalf("the step that calls %s did not panic", call)
			}

			free(t, r.b, ls)
		})
	}
}

// TestRegisterLockerConcurrent holds calls of a Locker open and checks what
// goes on meanwhile. The calls for one database wait for each other, so they
// never overlap; reads of the database wait for a call that can lose its
// lock; another database of the registration is opened, locked and closed;
// and a database is not given a new Locker while its old one is being closed.
//
// That something waits shows in what it does after the call is let go, and
// in the lock table, which reports a call that starts while another runs.
// Before a call is let go the test gives the others time to run into it: if
// they are slow the test proves less, it does not fail.
func TestRegisterLockerConcurrent(t *testing.T) {
	tab := newLockTable(t)
	r := newLockRig(t, tab, "A", "B")
	l := r.v.last("test.db")
	// hold blocks the next call named call inside the Locker until the
	// returned function is called, which is after the Locker was entered.
	hold := func(l *tableLock, call string) (entered chan struct{}, release func()) {
		entered = make(chan struct{})
		gate := make(chan struct{})
		l.enter = func(c string) {
			if c == call {
				l.enter = nil
				close(entered)
				<-gate
			}
		}
		return entered, func() {
			time.Sleep(20 * time.Millisecond)
			close(gate)
		}
	}
	// start runs f on a goroutine and returns when it has begun.
	var wg sync.WaitGroup
	start := func(f func()) {
		begun := make(chan struct{})
		wg.Add(1)
		go func() {
			defer wg.Done()
			close(begun)
			f()
		}()
		<-begun
	}
	check := func(what string, rc, want int32) {
		t.Helper()
		if rc != want {
			t.Errorf("%s: rc %d, want %d", what, rc, want)
		}
	}
	p, _ := r.b.buf(8, 0)

	// The first reader is inside Lock.
	entered, release := hold(l, "Lock(shared)")
	start(func() { check("xLock(A)", rwLock(r.b.tls, r.files["A"], lockShared), lkOK) })
	<-entered
	o := r.b.mustOpen("other.db", mainDB) // another database does not wait
	check("xLock on another database", rwLock(r.b.tls, o, lockShared), lkOK)
	check("xClose of another database", rwClose(r.b.tls, o), lkOK)
	c := r.b.mustOpen("test.db", mainDB) // nor does an open of this one
	start(func() { check("xCheckReservedLock(B)", rwCheckReservedLock(r.b.tls, r.files["B"], p), lkOK) })
	start(func() { check("xLock(B)", rwLock(r.b.tls, r.files["B"], lockShared), lkOK) })
	release()
	wg.Wait()
	if got := l.take(); len(got) != 2 || got[0] != "Lock(shared)=ok" || got[1] != "CheckReserved=ok" {
		t.Fatalf("calls: %v, want Lock(shared) to finish before CheckReserved, and one Lock for both readers", got)
	}

	// The writer is inside Lock(reserved): a read by the other reader waits
	// for it, since the call could lose the lock the read relies on.
	entered, release = hold(l, "Lock(reserved)")
	l.fail = func(call string) (error, bool) {
		if call != "Lock(reserved)" {
			return nil, false
		}

		l.fail = nil
		return errors.New("boom"), false
	}
	start(func() { check("xLock(A, RESERVED)", rwLock(r.b.tls, r.files["A"], lockReserved), lkIOERR) })
	<-entered
	start(func() {
		// Let through before the call, the read would be answered.
		check("xRead(B) across a failing Lock", rwRead(r.b.tls, r.files["B"], p, 8, 0), sqlite3.SQLITE_IOERR_READ)
	})
	release()
	wg.Wait()
	for _, f := range []string{"A", "B"} { // everybody out: the state is known again
		check("xUnlock("+f+")", rwUnlock(r.b.tls, r.files[f], lockNone), lkOK)
	}
	if got := strings.Join(l.take(), " "); got != "Lock(reserved)=error Unlock(none)=ok" {
		t.Fatalf("calls: %q", got)
	}

	// The other way round: a read is under way when the writer asks for
	// RESERVED. The Locker is not called before the read is over.
	for _, f := range []string{"A", "B"} {
		check("xLock("+f+")", rwLock(r.b.tls, r.files[f], lockShared), lkOK)
	}
	l.take()
	var events []string // under evMu
	var evMu sync.Mutex
	event := func(s string) {
		evMu.Lock()
		events = append(events, s)
		evMu.Unlock()
	}
	reading, readOn := make(chan struct{}), make(chan struct{})
	r.readAt = func() {
		close(reading)
		<-readOn
		event("read over")
	}
	l.enter = func(call string) { event(call) }
	start(func() { check("xRead(B)", rwRead(r.b.tls, r.files["B"], p, 8, 0), sqlite3.SQLITE_IOERR_SHORT_READ) })
	<-reading
	r.readAt = nil
	start(func() { check("xLock(A, RESERVED)", rwLock(r.b.tls, r.files["A"], lockReserved), lkOK) })
	time.Sleep(20 * time.Millisecond)
	close(readOn)
	wg.Wait()
	l.enter = nil
	if got := strings.Join(events, ", "); got != "read over, Lock(reserved)" {
		t.Fatalf("order of events: %s, want the read to be over before Lock(reserved)", got)
	}

	l.take()
	check("xUnlock(A)", rwUnlock(r.b.tls, r.files["A"], lockNone), lkOK)
	l.take()

	// The last reader is inside Unlock: the next reader waits and then
	// takes the lock anew.
	entered, release = hold(l, "Unlock(none)")
	start(func() { check("xUnlock(B)", rwUnlock(r.b.tls, r.files["B"], lockNone), lkOK) })
	<-entered
	start(func() { check("xLock(A) again", rwLock(r.b.tls, r.files["A"], lockShared), lkOK) })
	release()
	wg.Wait()
	if got := strings.Join(l.take(), " "); got != "Unlock(none)=ok Lock(shared)=ok" {
		t.Fatalf("calls: %q, want the Unlock to finish before the next Lock", got)
	}

	// The last file closes and is inside the Close of the Locker: an open
	// of the database waits and gets a new Locker after it.
	check("xUnlock(A)", rwUnlock(r.b.tls, r.files["A"], lockNone), lkOK)
	l.take()
	for _, f := range []string{"A", "B"} {
		r.close(f)
	}
	entered, release = hold(l, "Close")
	start(func() { check("xClose of the last file", rwClose(r.b.tls, c), lkOK) })
	<-entered
	var reopened uintptr
	start(func() { reopened = r.b.mustOpen("test.db", mainDB) })
	time.Sleep(20 * time.Millisecond)
	if n := len(r.v.locks["test.db"]); n != 1 {
		// Read without the mutex of the VFS on purpose: OpenLock must not
		// be running.
		t.Errorf("%d Lockers were given out while the first was being closed", n)
	}
	release()
	wg.Wait()
	if l2 := r.v.last("test.db"); l2 == l || len(l.take()) != 1 {
		t.Fatalf("after the Close: the database has no new Locker, or the old one was called again")
	}

	check("xClose", rwClose(r.b.tls, reopened), lkOK)
}

// transfer is a movement of money between two accounts that a worker
// committed.
type transfer struct{ w, i, src, dst, amt int }

// runTransfers has workers move money between the accounts of the database,
// each through the handle dbs[w%len(dbs)], and returns the transfers that
// were committed. Worker w draws its transfers from seed+w.
func runTransfers(t *testing.T, dbs []*sql.DB, accounts, workers, transfers int, seed int64) []transfer {
	t.Helper()
	var (
		wg        sync.WaitGroup
		mu        sync.Mutex
		committed []transfer
	)
	errc := make(chan error, workers)
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			db := dbs[w%len(dbs)]
			rng := rand.New(rand.NewSource(seed + int64(w)))
			for i := 0; i < transfers; i++ {
				tr := transfer{w, i, rng.Intn(accounts), rng.Intn(accounts), 1 + rng.Intn(50)}
				for attempt := 0; ; attempt++ {
					if attempt == 100000 {
						errc <- fmt.Errorf("worker %d, transfer %d: still SQLITE_BUSY after %d attempts", w, i, attempt)
						return
					}

					err := func() error {
						tx, err := db.Begin()
						if err != nil {
							return err
						}
						defer tx.Rollback()

						// A read first, so that the write upgrades a shared lock.
						var bal int
						if err := tx.QueryRow("SELECT bal FROM acct WHERE id=?", tr.src).Scan(&bal); err != nil {
							return err
						}
						if _, err := tx.Exec("UPDATE acct SET bal=bal-? WHERE id=?", tr.amt, tr.src); err != nil {
							return err
						}
						if _, err := tx.Exec("UPDATE acct SET bal=bal+? WHERE id=?", tr.amt, tr.dst); err != nil {
							return err
						}
						if _, err := tx.Exec("INSERT INTO done VALUES(?,?,?,?,?)", tr.w, tr.i, tr.src, tr.dst, tr.amt); err != nil {
							return err
						}
						return tx.Commit()
					}()
					if err == nil {
						mu.Lock()
						committed = append(committed, tr)
						mu.Unlock()
						break
					}

					var se *sqlite.Error
					if errors.As(err, &se) && se.Code()&0xff == sqlite3.SQLITE_BUSY {
						continue
					}

					errc <- fmt.Errorf("worker %d, transfer %d: %s", w, i, E(err))
					return
				}
			}
		}(w)
	}
	wg.Wait()
	close(errc)
	for err := range errc {
		t.Error(err)
	}
	if t.Failed() {
		t.FailNow()
	}

	return committed
}

const transferSchema = "CREATE TABLE acct(id INTEGER PRIMARY KEY, bal INTEGER NOT NULL); CREATE TABLE done(w INTEGER, i INTEGER, src INTEGER, dst INTEGER, amt INTEGER)"

// checkTransfers compares the database with what the committed transfers add
// up to: every one of them is there, and nothing else.
func checkTransfers(t *testing.T, db *sql.DB, accounts int, committed []transfer) {
	t.Helper()
	want := make([]int, accounts)
	for i := range want {
		want[i] = 1000
	}
	for _, tr := range committed {
		want[tr.src] -= tr.amt
		want[tr.dst] += tr.amt
	}
	var wantDump strings.Builder
	for i, b := range want {
		fmt.Fprintf(&wantDump, "%d|%d|\n", i, b)
	}
	if got := dump(t, db, "SELECT id, bal FROM acct ORDER BY id"); got != wantDump.String() {
		t.Fatalf("balances:\n%s\nwant, from the %d transfers that committed:\n%s", got, len(committed), wantDump.String())
	}

	lines := make([]string, len(committed))
	for i, tr := range committed {
		lines[i] = fmt.Sprintf("%d|%d|%d|\n", tr.src, tr.dst, tr.amt)
	}
	sort.Strings(lines)
	if got, want := dump(t, db, "SELECT src, dst, amt FROM done ORDER BY CAST(src AS TEXT)||'|'||dst||'|'||amt||'|'"), strings.Join(lines, ""); got != want {
		t.Fatalf("transfers recorded:\n%s\nwant the %d that committed:\n%s", got, len(committed), want)
	}

	integrity(t, db)
}

// TestRegisterLockerTwoRegistrations writes one database through two
// registrations of one storage at the same time, as two processes would. The
// Lockers of a lockTable keep them apart, and the table checks every call it
// gets against the protocol.
func TestRegisterLockerTwoRegistrations(t *testing.T) {
	const accounts, workers, transfers = 8, 8, 40
	s := newMemStore()
	tab := newLockTable(t)
	a := openRW(t, &lockVFS{VFS: s, tab: tab}, "test.db", workers/2, "_pragma=busy_timeout(10000)")
	mustExec(t, a, strings.Split(transferSchema, "; ")...)
	for i := 0; i < accounts; i++ {
		mustExec(t, a, fmt.Sprintf("INSERT INTO acct VALUES(%d, 1000)", i))
	}
	b := openRW(t, &lockVFS{VFS: s, tab: tab}, "test.db", workers/2, "_pragma=busy_timeout(10000)")

	// One writer at a time, across the registrations: a transaction open
	// through one keeps the other from writing and lets it read.
	nobusy := openRW(t, &lockVFS{VFS: s, tab: tab}, "test.db", 1)
	tx, err := a.Begin()
	if err != nil {
		t.Fatal(err)
	}

	if _, err := tx.Exec("UPDATE acct SET bal=bal WHERE id=0"); err != nil {
		t.Fatal(E(err))
	}

	var se *sqlite.Error
	if _, err := nobusy.Exec("UPDATE acct SET bal=bal WHERE id=1"); !errors.As(err, &se) || se.Code() != sqlite3.SQLITE_BUSY {
		t.Fatalf("a write through a second registration next to an open write transaction: %v, want SQLITE_BUSY", E(err))
	}

	if got := dump(t, nobusy, "SELECT count(*) FROM acct"); got != fmt.Sprintf("%d|\n", accounts) {
		t.Fatalf("a read through a second registration next to an open write transaction: %q", got)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(E(err))
	}

	mustExec(t, nobusy, "UPDATE acct SET bal=bal WHERE id=1")

	committed := runTransfers(t, []*sql.DB{a, b}, accounts, workers, transfers, 0)
	if len(committed) != workers*transfers {
		t.Fatalf("%d transfers committed, want %d", len(committed), workers*transfers)
	}

	checkTransfers(t, a, accounts, committed)
	checkTransfers(t, b, accounts, committed)
}
