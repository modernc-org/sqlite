// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vfs

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"testing"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libc/sys/types"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// The tests in this file call the methods of the VFS as SQLite does, where
// SQL cannot show what a method returned.

// fakeVFS and fakeFile do what their fields say.
type fakeVFS struct {
	open     func(name string, flags OpenFlag) (File, OpenFlag, error)
	del      func(name string, syncDir bool) error
	access   func(name string, flags AccessFlag) (bool, error)
	pathname func(name string) (string, error)
}

func (v *fakeVFS) Open(name string, flags OpenFlag) (File, OpenFlag, error) {
	if v.open == nil {
		return &fakeFile{}, flags, nil
	}

	return v.open(name, flags)
}

func (v *fakeVFS) Delete(name string, syncDir bool) error {
	if v.del == nil {
		return nil
	}

	return v.del(name, syncDir)
}

func (v *fakeVFS) Access(name string, flags AccessFlag) (bool, error) {
	if v.access == nil {
		return false, nil
	}

	return v.access(name, flags)
}

func (v *fakeVFS) FullPathname(name string) (string, error) {
	if v.pathname == nil {
		return name, nil
	}

	return v.pathname(name)
}

type fakeFile struct {
	readAt   func(p []byte, off int64) (int, error)
	writeAt  func(p []byte, off int64) (int, error)
	truncate func(size int64) error
	sync     func(flags SyncFlag) error
	size     func() (int64, error)
	close    func() error
}

func (f *fakeFile) ReadAt(p []byte, off int64) (int, error) {
	if f.readAt == nil {
		return 0, io.EOF
	}

	return f.readAt(p, off)
}

func (f *fakeFile) WriteAt(p []byte, off int64) (int, error) {
	if f.writeAt == nil {
		return len(p), nil
	}

	return f.writeAt(p, off)
}

func (f *fakeFile) Truncate(size int64) error {
	if f.truncate == nil {
		return nil
	}

	return f.truncate(size)
}

func (f *fakeFile) Sync(flags SyncFlag) error {
	if f.sync == nil {
		return nil
	}

	return f.sync(flags)
}

func (f *fakeFile) Size() (int64, error) {
	if f.size == nil {
		return 0, nil
	}

	return f.size()
}

func (f *fakeFile) Close() error {
	if f.close == nil {
		return nil
	}

	return f.close()
}

// bridge calls a registered VFS the way SQLite does.
type bridge struct {
	t    *testing.T
	tls  *libc.TLS
	fsys *FS
}

func newBridge(t *testing.T, v VFS) *bridge {
	t.Helper()
	_, fsys, err := Register(v)
	if err != nil {
		t.Fatalf("Register: %v", err)
	}

	b := &bridge{t: t, tls: libc.NewTLS(), fsys: fsys}
	t.Cleanup(func() {
		if err := fsys.Close(); err != nil {
			t.Errorf("closing the VFS: %v", err)
		}
		b.tls.Close()
	})
	return b
}

const mainDB = sqlite3.SQLITE_OPEN_READWRITE | sqlite3.SQLITE_OPEN_CREATE | sqlite3.SQLITE_OPEN_MAIN_DB

// open calls xOpen and returns the sqlite3_file, which the caller closes with
// xClose if the result code is SQLITE_OK and frees in any case.
func (b *bridge) open(name string, flags int32) (pFile uintptr, rc int32) {
	b.t.Helper()
	pFile = libc.Xmalloc(b.tls, types.Size_t(unsafe.Sizeof(VFSFile{})))
	for i := range unsafe.Slice((*byte)(unsafe.Pointer(pFile)), unsafe.Sizeof(VFSFile{})) {
		*(*byte)(unsafe.Pointer(pFile + uintptr(i))) = 0xff
	}
	b.t.Cleanup(func() { libc.Xfree(b.tls, pFile) })
	var zName uintptr
	if name != "" {
		var err error
		if zName, err = libc.CString(name); err != nil {
			b.t.Fatal(err)
		}
		defer libc.Xfree(b.tls, zName)
	}
	return pFile, rwOpen(b.tls, b.fsys.cvfs, zName, pFile, flags, 0)
}

func (b *bridge) mustOpen(name string, flags int32) uintptr {
	b.t.Helper()
	pFile, rc := b.open(name, flags)
	if rc != sqlite3.SQLITE_OK {
		b.t.Fatalf("xOpen(%q): rc %d", name, rc)
	}

	return pFile
}

func (b *bridge) buf(n int, fill byte) (uintptr, []byte) {
	p := libc.Xmalloc(b.tls, types.Size_t(n))
	b.t.Cleanup(func() { libc.Xfree(b.tls, p) })
	s := unsafe.Slice((*byte)(unsafe.Pointer(p)), n)
	for i := range s {
		s[i] = fill
	}
	return p, s
}

// TestRegisterLocks walks SQLite's lock protocol over files open on one
// database: readers next to a writer that holds RESERVED, PENDING kept after
// a failed EXCLUSIVE and shutting new readers out, EXCLUSIVE once the last
// reader is gone, the direct step from SHARED to EXCLUSIVE that the rollback
// of a hot journal takes, and a lock released by closing the file.
func TestRegisterLocks(t *testing.T) {
	const ok, busy = sqlite3.SQLITE_OK, sqlite3.SQLITE_BUSY
	b := newBridge(t, &fakeVFS{})
	files := map[string]uintptr{}
	for _, n := range []string{"A", "B", "C", "D"} {
		files[n] = b.mustOpen("test.db", mainDB)
	}
	other := b.mustOpen("other.db", mainDB)
	journal := b.mustOpen("test.db-journal", sqlite3.SQLITE_OPEN_READWRITE|sqlite3.SQLITE_OPEN_CREATE|sqlite3.SQLITE_OPEN_MAIN_JOURNAL)

	reserved := func(f string) int32 {
		p, s := b.buf(4, 0xff)
		if rc := rwCheckReservedLock(b.tls, files[f], p); rc != ok {
			t.Fatalf("xCheckReservedLock(%s): rc %d", f, rc)
		}
		return *(*int32)(unsafe.Pointer(&s[0]))
	}
	steps := []struct {
		f      string
		lock   int32 // level asked for, or -1-level to unlock to level
		want   int32
		writer int32 // what xCheckReservedLock says after the step
	}{
		{"A", lockShared, ok, 0},
		{"B", lockShared, ok, 0},
		{"A", lockReserved, ok, 1},
		{"B", lockReserved, busy, 1},  // one writer
		{"C", lockShared, ok, 1},      // readers next to RESERVED
		{"A", lockExclusive, busy, 1}, // B and C still read: A holds PENDING
		{"D", lockShared, busy, 1},    // PENDING shuts new readers out
		{"B", -1 - lockNone, ok, 1},
		{"A", lockExclusive, busy, 1}, // C still reads
		{"C", -1 - lockNone, ok, 1},
		{"A", lockExclusive, ok, 1}, // alone
		{"B", lockShared, busy, 1},
		{"A", -1 - lockShared, ok, 0}, // back to a reader
		{"B", lockShared, ok, 0},
		// B goes from SHARED to EXCLUSIVE, as a connection does that found a
		// hot journal. A still reads, so B is refused, and it must stay a
		// plain reader: were it counted as a writer, A would take the
		// journal for one in progress and read the database as it is.
		{"B", lockExclusive, busy, 0},
		{"C", lockShared, ok, 0}, // no PENDING either
		{"C", -1 - lockNone, ok, 0},
		{"A", -1 - lockNone, ok, 0},
		{"B", lockExclusive, ok, 1}, // alone: the rollback can go on
		{"B", -1 - lockNone, ok, 0},
		{"A", lockReserved, sqlite3.SQLITE_IOERR_LOCK, 0}, // not from NONE
		{"A", lockPending, sqlite3.SQLITE_IOERR_LOCK, 0},  // never asked for
		{"A", lockShared, ok, 0},
		{"A", -1 - lockReserved, sqlite3.SQLITE_IOERR_UNLOCK, 0},
		{"A", lockReserved, ok, 1},
		{"B", lockShared, ok, 1},
		{"B", lockExclusive, busy, 1}, // A writes: refused outright, and no PENDING for B
		{"C", lockShared, ok, 1},      // so readers still come in
		{"C", -1 - lockNone, ok, 1},
		{"B", -1 - lockNone, ok, 1},
		{"A", lockExclusive, ok, 1}, // and A is still the writer
		{"A", -1 - lockShared, ok, 0},
		{"A", lockReserved, ok, 1},
	}
	for i, s := range steps {
		var rc int32
		if s.lock >= 0 {
			rc = rwLock(b.tls, files[s.f], s.lock)
		} else {
			rc = rwUnlock(b.tls, files[s.f], -1-s.lock)
		}
		if rc != s.want {
			t.Fatalf("step %d, %s %d: rc %d, want %d", i, s.f, s.lock, rc, s.want)
		}

		for _, f := range []string{"A", "B", "C", "D"} {
			if got := reserved(f); got != s.writer {
				t.Fatalf("step %d, %s %d: xCheckReservedLock(%s) says %d, want %d", i, s.f, s.lock, f, got, s.writer)
			}
		}
	}

	// A holds RESERVED: every file on the database sees a writer, the files
	// of another database and the journal do not take part.
	for _, f := range []string{"A", "B", "C"} {
		if got := reserved(f); got != 1 {
			t.Fatalf("xCheckReservedLock(%s) with A at RESERVED: %d, want 1", f, got)
		}
	}
	for _, p := range []uintptr{other, journal} {
		res, s := b.buf(4, 0xff)
		if rc := rwCheckReservedLock(b.tls, p, res); rc != ok || *(*int32)(unsafe.Pointer(&s[0])) != 0 {
			t.Fatalf("xCheckReservedLock of an unrelated file: rc %d, result %d", rc, s[0])
		}
		if rc := rwLock(b.tls, p, lockShared); rc != ok {
			t.Fatalf("xLock of an unrelated file: rc %d", rc)
		}
	}
	if rc := rwLock(b.tls, journal, lockExclusive); rc != ok {
		t.Fatalf("xLock on a journal: rc %d", rc)
	}

	// Closing A gives its lock up.
	if rc := rwClose(b.tls, files["A"]); rc != ok {
		t.Fatalf("xClose: rc %d", rc)
	}
	if got := reserved("B"); got != 0 {
		t.Fatalf("xCheckReservedLock(B) after A was closed: %d, want 0", got)
	}
	if rc := rwLock(b.tls, files["B"], lockShared); rc != ok {
		t.Fatalf("B to SHARED after A was closed: rc %d", rc)
	}
	if rc := rwLock(b.tls, files["B"], lockExclusive); rc != ok {
		t.Fatalf("B to EXCLUSIVE after A was closed: rc %d", rc)
	}

	for _, p := range []uintptr{files["B"], files["C"], files["D"], other, journal} {
		if rc := rwClose(b.tls, p); rc != ok {
			t.Fatalf("xClose: rc %d", rc)
		}
	}
	e := getObject(b.fsys.fsHandle).(*rwEntry)
	if n := len(e.locks); n != 0 {
		t.Fatalf("%d lock(s) left after every file was closed", n)
	}
	if n := e.open.Load(); n != 0 {
		t.Fatalf("%d file(s) counted open after every file was closed", n)
	}
}

// TestRegisterRead checks the three answers of xRead on a buffer filled with
// 0xff: a full read, a read past the end with SQLITE_IOERR_SHORT_READ and
// every unread byte zeroed, and SQLITE_IOERR_READ for a ReadAt that fails or
// breaks the io.ReaderAt contract.
func TestRegisterRead(t *testing.T) {
	data := bytes.Repeat([]byte("0123456789"), 10)
	var f fakeFile
	b := newBridge(t, &fakeVFS{open: func(string, OpenFlag) (File, OpenFlag, error) { return &f, 0, nil }})
	pFile := b.mustOpen("test.db", mainDB)
	defer rwClose(b.tls, pFile)

	exact := func(p []byte, off int64) (int, error) {
		if off >= int64(len(data)) {
			return 0, io.EOF
		}

		n := copy(p, data[off:])
		if n < len(p) {
			return n, io.EOF
		}

		return n, nil
	}
	for _, c := range []struct {
		name   string
		readAt func(p []byte, off int64) (int, error)
		off    int64
		amt    int
		want   int32
		have   int // bytes of data expected in the buffer, the rest zero
	}{
		{"inside", exact, 10, 50, sqlite3.SQLITE_OK, 50},
		{"across the end", exact, 60, 512, sqlite3.SQLITE_IOERR_SHORT_READ, 40},
		{"past the end", exact, 1000, 512, sqlite3.SQLITE_IOERR_SHORT_READ, 0},
		{"unexpected EOF", func(p []byte, off int64) (int, error) { n, _ := exact(p, off); return n, io.ErrUnexpectedEOF }, 60, 512, sqlite3.SQLITE_IOERR_SHORT_READ, 40},
		{"wrapped EOF", func(p []byte, off int64) (int, error) {
			n, _ := exact(p, off)
			return n, fmt.Errorf("reading: %w", io.EOF)
		}, 60, 512, sqlite3.SQLITE_IOERR_SHORT_READ, 40},
		{"all read and EOF", func(p []byte, off int64) (int, error) { n, _ := exact(p, off); return n, io.EOF }, 50, 50, sqlite3.SQLITE_OK, 50},
		{"all read and an error", func(p []byte, off int64) (int, error) { n, _ := exact(p, off); return n, errors.New("checksum") }, 10, 50, sqlite3.SQLITE_IOERR_READ, -1},
		{"short without an error", func(p []byte, off int64) (int, error) { n, _ := exact(p, off); return n, nil }, 60, 512, sqlite3.SQLITE_IOERR_READ, -1},
		{"error", func(p []byte, off int64) (int, error) { return 0, errors.New("no medium") }, 0, 50, sqlite3.SQLITE_IOERR_READ, -1},
		{"more than asked", func(p []byte, off int64) (int, error) { return len(p) + 1, io.EOF }, 0, 50, sqlite3.SQLITE_IOERR_READ, -1},
		{"less than nothing", func(p []byte, off int64) (int, error) { return -1, io.EOF }, 0, 50, sqlite3.SQLITE_IOERR_READ, -1},
	} {
		f.readAt = c.readAt
		zBuf, buf := b.buf(c.amt, 0xff)
		if rc := rwRead(b.tls, pFile, zBuf, int32(c.amt), c.off); rc != c.want {
			t.Fatalf("%s: rc %d, want %d", c.name, rc, c.want)
		}

		if c.have < 0 {
			continue
		}

		if c.have > 0 && !bytes.Equal(buf[:c.have], data[c.off:c.off+int64(c.have)]) {
			t.Fatalf("%s: the bytes read differ from the file", c.name)
		}

		for i, x := range buf[c.have:] {
			if x != 0 {
				t.Fatalf("%s: byte %d of the unread tail is %#x, want 0", c.name, c.have+i, x)
			}
		}
	}
}

// TestRegisterResultCodes checks the result code of every method for every
// kind of answer a VFS or a File can give. Errors from the data path become
// the method's own I/O error or SQLITE_FULL and nothing else: any other code
// would bypass the error state SQLite's pager keeps for them.
func TestRegisterResultCodes(t *testing.T) {
	boom := errors.New("boom")
	var (
		f fakeFile
		v fakeVFS
	)
	v.open = func(string, OpenFlag) (File, OpenFlag, error) { return &f, 0, nil }
	b := newBridge(t, &v)
	pFile := b.mustOpen("test.db", mainDB)
	zBuf, _ := b.buf(64, 0xaa)
	zName, err := libc.CString("test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer libc.Xfree(b.tls, zName)
	res, _ := b.buf(8, 0xff)

	check := func(what string, got, want int32) {
		t.Helper()
		if got != want {
			t.Fatalf("%s: rc %d, want %d", what, got, want)
		}
	}

	for _, c := range []struct {
		name string
		n    int
		err  error
		want int32
	}{
		{"all written", 64, nil, sqlite3.SQLITE_OK},
		{"short without an error", 63, nil, sqlite3.SQLITE_IOERR_WRITE},
		{"error", 0, boom, sqlite3.SQLITE_IOERR_WRITE},
		{"all written and an error", 64, boom, sqlite3.SQLITE_IOERR_WRITE},
		{"full", 10, ErrFull, sqlite3.SQLITE_FULL},
		{"full, wrapped", 0, fmt.Errorf("volume 3: %w", ErrFull), sqlite3.SQLITE_FULL},
		{"permission", 0, fs.ErrPermission, sqlite3.SQLITE_IOERR_WRITE},
	} {
		f.writeAt = func([]byte, int64) (int, error) { return c.n, c.err }
		check("xWrite, "+c.name, rwWrite(b.tls, pFile, zBuf, 64, 0), c.want)
	}

	f.truncate = func(int64) error { return boom }
	check("xTruncate", rwTruncate(b.tls, pFile, 0), sqlite3.SQLITE_IOERR_TRUNCATE)
	f.truncate = func(int64) error { return ErrFull }
	check("xTruncate, full", rwTruncate(b.tls, pFile, 0), sqlite3.SQLITE_IOERR_TRUNCATE)
	f.sync = func(SyncFlag) error { return boom }
	check("xSync", rwSync(b.tls, pFile, sqlite3.SQLITE_SYNC_NORMAL), sqlite3.SQLITE_IOERR_FSYNC)
	f.size = func() (int64, error) { return 0, boom }
	check("xFileSize", rwFileSize(b.tls, pFile, res), sqlite3.SQLITE_IOERR_FSTAT)
	f.size = func() (int64, error) { return -1, nil }
	check("xFileSize, negative", rwFileSize(b.tls, pFile, res), sqlite3.SQLITE_IOERR_FSTAT)
	check("xFileControl", rwFileControl(b.tls, pFile, sqlite3.SQLITE_FCNTL_SIZE_HINT, res), sqlite3.SQLITE_NOTFOUND)
	check("xSectorSize", rwSectorSize(b.tls, pFile), 4096)
	check("xDeviceCharacteristics", rwDeviceCharacteristics(b.tls, pFile), 0)

	// xClose gives the handle and the count back even when Close fails.
	h := (*VFSFile)(unsafe.Pointer(pFile)).fsFile
	f.close = func() error { return boom }
	check("xClose", rwClose(b.tls, pFile), sqlite3.SQLITE_IOERR_CLOSE)
	objectMu.Lock()
	_, live := objects[h]
	objectMu.Unlock()
	if live {
		t.Fatal("the handle of a file whose Close failed is still registered")
	}

	for _, c := range []struct {
		err  error
		want int32
	}{
		{nil, sqlite3.SQLITE_OK},
		{boom, sqlite3.SQLITE_IOERR_DELETE},
		{fs.ErrNotExist, sqlite3.SQLITE_IOERR_DELETE_NOENT},
		{&fs.PathError{Op: "remove", Path: "test.db", Err: fs.ErrNotExist}, sqlite3.SQLITE_IOERR_DELETE_NOENT},
	} {
		v.del = func(string, bool) error { return c.err }
		check(fmt.Sprintf("xDelete, %v", c.err), rwDelete(b.tls, b.fsys.cvfs, zName, 0), c.want)
	}

	v.access = func(string, AccessFlag) (bool, error) { return true, boom }
	check("xAccess", rwAccess(b.tls, b.fsys.cvfs, zName, sqlite3.SQLITE_ACCESS_EXISTS, res), sqlite3.SQLITE_IOERR_ACCESS)
	out, _ := b.buf(64, 0xff)
	v.pathname = func(string) (string, error) { return "", boom }
	check("xFullPathname", rwFullPathname(b.tls, b.fsys.cvfs, zName, 64, out), sqlite3.SQLITE_CANTOPEN)

	// A failed xOpen leaves no methods, so SQLite does not call xClose, and
	// no handle.
	objectMu.Lock()
	before := len(objects)
	objectMu.Unlock()
	for _, open := range []func(string, OpenFlag) (File, OpenFlag, error){
		func(string, OpenFlag) (File, OpenFlag, error) { return nil, 0, boom },
		func(string, OpenFlag) (File, OpenFlag, error) { return nil, 0, nil },
		func(string, OpenFlag) (File, OpenFlag, error) { return nil, 0, fs.ErrPermission },
	} {
		v.open = open
		p, rc := b.open("test.db", mainDB)
		check("xOpen", rc, sqlite3.SQLITE_CANTOPEN)
		if m := (*VFSFile)(unsafe.Pointer(p)).base.pMethods; m != 0 {
			t.Fatalf("a failed xOpen left pMethods %#x", m)
		}
	}
	// A File that comes with an error is closed: SQLite will not do it.
	closed := 0
	v.open = func(string, OpenFlag) (File, OpenFlag, error) {
		return &fakeFile{close: func() error { closed++; return boom }}, 0, boom
	}
	p, rc := b.open("test.db", mainDB)
	check("xOpen, a File and an error", rc, sqlite3.SQLITE_CANTOPEN)
	if m := (*VFSFile)(unsafe.Pointer(p)).base.pMethods; m != 0 || closed != 1 {
		t.Fatalf("an xOpen that got a File and an error left pMethods %#x and closed the File %d time(s), want once", m, closed)
	}

	objectMu.Lock()
	after := len(objects)
	objectMu.Unlock()
	if after != before {
		t.Fatalf("failed opens left %d handle(s)", after-before)
	}
	if n := getObject(b.fsys.fsHandle).(*rwEntry).open.Load(); n != 0 {
		t.Fatalf("%d file(s) counted open", n)
	}
}

// TestRegisterLargeOffsets passes offsets and sizes above 4 GiB through every
// method that carries one. They are 64 bits wide in SQLite on every target,
// the 32-bit ones included.
func TestRegisterLargeOffsets(t *testing.T) {
	const big = int64(5)<<32 + 4096
	var got []int64
	f := &fakeFile{
		readAt:   func(p []byte, off int64) (int, error) { got = append(got, off); return len(p), nil },
		writeAt:  func(p []byte, off int64) (int, error) { got = append(got, off); return len(p), nil },
		truncate: func(size int64) error { got = append(got, size); return nil },
		size:     func() (int64, error) { return big + 3, nil },
	}
	b := newBridge(t, &fakeVFS{open: func(string, OpenFlag) (File, OpenFlag, error) { return f, 0, nil }})
	pFile := b.mustOpen("test.db", mainDB)
	defer rwClose(b.tls, pFile)

	zBuf, _ := b.buf(16, 0)
	res, _ := b.buf(8, 0)
	if rc := rwRead(b.tls, pFile, zBuf, 16, big); rc != sqlite3.SQLITE_OK {
		t.Fatalf("xRead: rc %d", rc)
	}
	if rc := rwWrite(b.tls, pFile, zBuf, 16, big+1); rc != sqlite3.SQLITE_OK {
		t.Fatalf("xWrite: rc %d", rc)
	}
	if rc := rwTruncate(b.tls, pFile, big+2); rc != sqlite3.SQLITE_OK {
		t.Fatalf("xTruncate: rc %d", rc)
	}
	if rc := rwFileSize(b.tls, pFile, res); rc != sqlite3.SQLITE_OK {
		t.Fatalf("xFileSize: rc %d", rc)
	}
	got = append(got, *(*int64)(unsafe.Pointer(res)))
	if want := []int64{big, big + 1, big + 2, big + 3}; fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("offsets and sizes seen: %v, want %v", got, want)
	}
}

// TestRegisterFullPathname checks that the name FullPathname returns is the
// one the database is opened and locked under, so that two names of one
// database meet in one lock, and that a result too long for SQLite fails the
// open.
func TestRegisterFullPathname(t *testing.T) {
	s := newMemStore()
	s.canon = func(name string) string { return strings.TrimPrefix(name, "alias-of-") }
	vfsName, fsys, err := Register(s)
	if err != nil {
		t.Fatal(err)
	}
	defer fsys.Close()

	open := func(name string) *sql.DB {
		db, err := sql.Open("sqlite", "file:"+name+"?vfs="+vfsName)
		if err != nil {
			t.Fatal(err)
		}
		db.SetMaxOpenConns(1)
		return db
	}
	a, b := open("test.db"), open("alias-of-test.db")
	defer a.Close()
	defer b.Close()
	mustExec(t, a, "CREATE TABLE t(v)", "BEGIN IMMEDIATE", "INSERT INTO t VALUES(1)")
	_, err = b.Exec("INSERT INTO t VALUES(2)")
	var se *sqlite.Error
	if !errors.As(err, &se) || se.Code()&0xff != sqlite3.SQLITE_BUSY {
		t.Fatalf("writing through the alias while the other name holds the write lock: %v, want SQLITE_BUSY", err)
	}

	mustExec(t, a, "COMMIT")
	mustExec(t, b, "INSERT INTO t VALUES(2)")
	if got := dump(t, a, "SELECT v FROM t ORDER BY v"); got != "1|\n2|\n" {
		t.Fatalf("got %q", got)
	}

	if got := s.names(); len(got) != 1 || got[0] != "test.db" {
		t.Fatalf("files in the storage: %q, want only test.db", got)
	}

	s.canon = func(string) string { return strings.Repeat("d", 5000) }
	long := open("long.db")
	defer long.Close()
	if err := long.Ping(); !errors.As(err, &se) || se.Code()&0xff != sqlite3.SQLITE_CANTOPEN {
		t.Fatalf("opening a database whose full pathname has 5000 bytes: %v, want SQLITE_CANTOPEN", err)
	}
}

// TestRegisterMethodTable checks that the table of file methods is complete
// for its version and promises nothing beyond it: SQLite calls xShmMap and
// xFetch on a table of version 2 or 3 without asking.
func TestRegisterMethodTable(t *testing.T) {
	if vfsioRW.iVersion != 1 {
		t.Fatalf("iVersion %d, want 1", vfsioRW.iVersion)
	}

	for name, p := range map[string]uintptr{
		"xClose": vfsioRW.xClose, "xRead": vfsioRW.xRead, "xWrite": vfsioRW.xWrite, "xTruncate": vfsioRW.xTruncate,
		"xSync": vfsioRW.xSync, "xFileSize": vfsioRW.xFileSize, "xLock": vfsioRW.xLock, "xUnlock": vfsioRW.xUnlock,
		"xCheckReservedLock": vfsioRW.xCheckReservedLock, "xFileControl": vfsioRW.xFileControl,
		"xSectorSize": vfsioRW.xSectorSize, "xDeviceCharacteristics": vfsioRW.xDeviceCharacteristics,
	} {
		if p == 0 {
			t.Errorf("%s is not set", name)
		}
	}
	for name, p := range map[string]uintptr{
		"xShmMap": vfsioRW.xShmMap, "xShmLock": vfsioRW.xShmLock, "xShmBarrier": vfsioRW.xShmBarrier,
		"xShmUnmap": vfsioRW.xShmUnmap, "xFetch": vfsioRW.xFetch, "xUnfetch": vfsioRW.xUnfetch,
	} {
		if p != 0 {
			t.Errorf("%s is set in a table of version 1", name)
		}
	}

	db := openRW(t, newMemStore(), "test.db", 1)
	mustExec(t, db, "PRAGMA mmap_size=1048576", "CREATE TABLE t(v)", "INSERT INTO t VALUES(zeroblob(100000))")
	if got := dump(t, db, "SELECT length(v) FROM t"); got != "100000|\n" {
		t.Fatalf("with mmap_size set: %q", got)
	}
}

// TestRegisterClose checks the life cycle of a registration: it cannot be
// closed under an open database, temporary files count while they are open,
// and when everything is closed no handle is left.
func TestRegisterClose(t *testing.T) {
	objectMu.Lock()
	before := len(objects)
	objectMu.Unlock()

	if _, _, err := Register(nil); err == nil {
		t.Fatal("Register(nil) succeeded")
	}

	s := newMemStore()
	vfsName, fsys, err := Register(s)
	if err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("sqlite", "file:test.db?vfs="+vfsName)
	if err != nil {
		t.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	// A temporary table too large for its cache is a temporary file.
	mustExec(t, db, "CREATE TABLE t(v)", "CREATE TEMP TABLE tt AS WITH RECURSIVE n(i) AS (SELECT 1 UNION ALL SELECT i+1 FROM n WHERE i<6000) SELECT randomblob(1000) AS v FROM n")
	if err := fsys.Close(); !errors.Is(err, ErrInUse) {
		t.Fatalf("FS.Close with a database open: %v, want ErrInUse", err)
	}

	if n := fsys.open.Load(); n < 2 {
		t.Fatalf("%d file(s) counted open with a database and a temporary table, want at least 2", n)
	}

	mustExec(t, db, "INSERT INTO t VALUES(1)")
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	if err := fsys.Close(); err != nil {
		t.Fatalf("FS.Close after the database was closed: %v", err)
	}

	if err := fsys.Close(); err != nil {
		t.Fatalf("second FS.Close: %v", err)
	}

	objectMu.Lock()
	after := len(objects)
	objectMu.Unlock()
	if after != before {
		t.Fatalf("%d handle(s) left", after-before)
	}
}

// TestRegisterWriteErrors takes an error of the storage up to database/sql:
// a full storage is SQLITE_FULL, any other failed write an I/O error, the
// transaction is rolled back and the database stays sound.
func TestRegisterWriteErrors(t *testing.T) {
	for _, c := range []struct {
		name string
		err  error
		want int
	}{
		{"full", fmt.Errorf("quota: %w", ErrFull), sqlite3.SQLITE_FULL},
		{"error", errors.New("no medium"), sqlite3.SQLITE_IOERR},
	} {
		t.Run(c.name, func(t *testing.T) {
			s := newMemStore()
			db := openRW(t, s, "test.db", 1)
			mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, v)", "INSERT INTO t VALUES(1, 'kept')")
			s.mu.Lock()
			injected := false
			s.fail = func(_ int, op, name string) error {
				if op == "write" && kindOf(name) == "db" && !injected {
					injected = true
					return c.err
				}
				return nil
			}
			s.mu.Unlock()

			_, err := db.Exec("INSERT INTO t VALUES(2, zeroblob(50000))")
			var se *sqlite.Error
			if !errors.As(err, &se) || se.Code()&0xff != c.want {
				t.Fatalf("got %v, want result code %d", err, c.want)
			}

			s.mu.Lock()
			s.fail = nil
			s.mu.Unlock()
			if !injected {
				t.Fatal("no write to the database was made")
			}

			if got := dump(t, db, "SELECT id, v FROM t"); got != "1|kept|\n" {
				t.Fatalf("after the failed insert: %q", got)
			}

			integrity(t, db)
			mustExec(t, db, "INSERT INTO t VALUES(2, 'after')")
		})
	}
}

// TestRegisterWriteOrder checks on the trace of a transaction the order SQLite
// relies on for a rollback journal: the journal is synced before the database
// is written, the database is synced before the journal is deleted, and the
// directory is synced with the deletion under synchronous=EXTRA only.
func TestRegisterWriteOrder(t *testing.T) {
	for _, extra := range []bool{false, true} {
		s := newMemStore()
		db := openRW(t, s, "test.db", 1)
		mustExec(t, db, "CREATE TABLE t(v)")
		if extra {
			mustExec(t, db, "PRAGMA synchronous=EXTRA")
		}
		s.mu.Lock()
		s.trace = nil
		s.mu.Unlock()
		mustExec(t, db, "INSERT INTO t VALUES(zeroblob(20000))")
		s.mu.Lock()
		trace := append([]string(nil), s.trace...)
		s.mu.Unlock()

		index := func(prefix string, last bool) int {
			r := -1
			for i, l := range trace {
				if strings.HasPrefix(l, prefix) {
					if r = i; !last {
						break
					}
				}
			}
			return r
		}
		journalSync, dbWrite, dbLastWrite := index("sync journal", false), index("write db", false), index("write db", true)
		dbSync, del := index("sync db", true), index("delete journal", false)
		if journalSync < 0 || dbWrite < 0 || dbSync < 0 || del < 0 {
			t.Fatalf("trace lacks a step:\n%s", strings.Join(trace, "\n"))
		}

		if !(journalSync < dbWrite && dbLastWrite < dbSync && dbSync < del) {
			t.Fatalf("out of order (journal sync %d, database writes %d..%d, database sync %d, journal deletion %d):\n%s",
				journalSync, dbWrite, dbLastWrite, dbSync, del, strings.Join(trace, "\n"))
		}

		if want := fmt.Sprintf("delete journal syncDir=%v", extra); trace[del] != want {
			t.Fatalf("synchronous=EXTRA %v: %q, want %q", extra, trace[del], want)
		}

		for _, l := range trace {
			if strings.HasPrefix(l, "sync ") && strings.HasSuffix(l, " 0x0") {
				t.Fatalf("a sync without flags: %q", l)
			}
		}
	}
}
