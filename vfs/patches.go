// Copyright 2022 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vfs gives SQLite a VFS written in Go: a read-only one over a Go
// [fs.FS], made by [New], or one a database can be written through, made by
// [Register] from a [VFS].
//
// [New] registers the file system with SQLite and returns the name it was
// registered under. Passing that name as the vfs DSN query parameter opens a
// database read from the wrapped file system instead of from the host one.
// Paths are resolved by the wrapped [fs.FS], so the database is named relative
// to its root:
//
//	name, fsvfs, err := vfs.New(os.DirFS(dir))
//	if err != nil {
//		return err
//	}
//
//	defer fsvfs.Close()
//
//	db, err := sql.Open("sqlite", "file:test.db?vfs="+name)
//
// Any [fs.FS] will do, an embed.FS included, which is what makes this useful
// for shipping a database inside the binary. The VFS is read only: it has no
// write path and refuses to create a journal, so a database opened through it
// can only be read from. A hot journal next to the database, one SQLite would
// have to roll back first, makes reading it fail with SQLITE_READONLY_ROLLBACK,
// as a database opened with mode=ro does.
//
// [Register] takes a [VFS], whose files SQLite reads, writes, truncates, syncs
// and deletes: the database, its journal and the rest. It is the way to keep a
// database somewhere a file system does not reach. See [Register] for what
// the package does on top of the VFS, and for what it does not do yet.
//
// Registration is process-global, as SQLite's own VFS registry is. Each [New]
// and each [Register] registers a separate VFS under a fresh name, and
// [FS.Close] unregisters it again. Close every database opened through an [FS]
// before closing it: while any is open, [FS.Close] refuses with an error
// wrapping [ErrInUse].
package vfs

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"modernc.org/libc"
	sqlite3 "modernc.org/sqlite/lib"
)

// ErrInUse is returned, wrapped, by [FS.Close] while a database opened through
// the file system is still open.
var ErrInUse = errors.New("vfs: file system is in use")

var (
	fToken uintptr

	// New, FS.Close
	mu sync.Mutex

	objectMu sync.Mutex
	objects  = map[uintptr]interface{}{}
)

// fsEntry is what a registered VFS's pAppData refers to.
type fsEntry struct {
	fsys fs.FS
	open atomic.Int32 // files currently open through fsys
}

// fileEntry is what an open VFSFile's fsFile refers to.
type fileEntry struct {
	f  fs.File
	fs *fsEntry
}

func token() uintptr { return atomic.AddUintptr(&fToken, 1) }

// addObject files o under a fresh handle. Handles come from one counter for
// file systems and files alike, so on 32-bit targets it wraps after 2^32
// opens; a handle still in use, such as that of a file system registered at
// start-up, is skipped rather than overwritten. So is 0.
func addObject(o interface{}) uintptr {
	objectMu.Lock()
	defer objectMu.Unlock()

	for {
		t := token()
		if _, busy := objects[t]; t != 0 && !busy {
			objects[t] = o
			return t
		}
	}
}

func getObject(t uintptr) interface{} {
	objectMu.Lock()
	o := objects[t]
	objectMu.Unlock()
	if o == nil {
		panic(fmt.Sprintf("vfs: unknown handle %#x", t))
	}

	return o
}

func removeObject(t uintptr) {
	objectMu.Lock()
	_, ok := objects[t]
	delete(objects, t)
	objectMu.Unlock()
	if !ok {
		panic(fmt.Sprintf("vfs: removing unknown handle %#x", t))
	}
}

var vfsio = sqlite3_io_methods{
	iVersion: 1, // iVersion
}

// The methods a read-only VFS has no use for are implemented all the same: a
// nil method in the table is a crash should SQLite ever call it, a method that
// fails is an error. SQLite consults the SQLITE_OPEN_READONLY out-flag that
// vfsOpen reports before every write path, so these are not expected to run.
// The other methods are wired in patches32.go and patches64.go.
func init() {
	*(*func(*libc.TLS, uintptr, uintptr, int32, sqlite_int64) int32)(unsafe.Pointer(&vfsio.xWrite)) = vfsWriteReadOnly
	*(*func(*libc.TLS, uintptr, sqlite_int64) int32)(unsafe.Pointer(&vfsio.xTruncate)) = vfsTruncateReadOnly
	*(*func(*libc.TLS, uintptr, int32) int32)(unsafe.Pointer(&vfsio.xSync)) = vfsSyncReadOnly
}

func vfsWriteReadOnly(tls *libc.TLS, pFile uintptr, zBuf uintptr, iAmt int32, iOfst sqlite_int64) int32 {
	return sqlite3.SQLITE_READONLY
}

func vfsTruncateReadOnly(tls *libc.TLS, pFile uintptr, size sqlite_int64) int32 {
	return sqlite3.SQLITE_READONLY
}

// A file nothing can write to is always in sync.
func vfsSyncReadOnly(tls *libc.TLS, pFile uintptr, flags int32) int32 {
	return sqlite3.SQLITE_OK
}

// vfsDeleteReadOnly is the xDelete of a VFS made by New: nothing is deleted
// through a read-only file system. SQLite calls xDelete without consulting the
// read-only flag, to remove a journal it believes it no longer needs: when
// PRAGMA journal_mode leaves PERSIST or TRUNCATE, and on the first read of a
// 0-byte database that has a -journal file next to it. Both callers tolerate
// the failure. The vfsDelete transpiled from vfs.c is a stub that aborts the
// process, and would unlink the path on the host file system if it did not;
// New replaces it.
func vfsDeleteReadOnly(tls *libc.TLS, pVfs uintptr, zPath uintptr, dirSync int32) int32 {
	return sqlite3.SQLITE_IOERR_DELETE
}

func vfsOpen(tls *libc.TLS, pVfs uintptr, zName uintptr, pFile uintptr, flags int32, pOutFlags uintptr) int32 {
	if zName == 0 {
		return sqlite3.SQLITE_IOERR
	}

	// A journal opens read-only like any other file. SQLite opens one it finds
	// next to the database to see whether it is hot: one journal_mode=PERSIST
	// or TRUNCATE leaves behind is not, and the database reads; a hot one fails
	// the read with SQLITE_READONLY_ROLLBACK, as with mode=ro. Every journal
	// used to be refused with SQLITE_NOMEM, which failed each read of such a
	// database as out of memory.
	p := pFile
	*(*VFSFile)(unsafe.Pointer(p)) = VFSFile{}
	e := getObject((*sqlite3_vfs)(unsafe.Pointer(pVfs)).pAppData).(*fsEntry)
	f, err := e.fsys.Open(libc.GoString(zName))
	if err != nil {
		return sqlite3.SQLITE_CANTOPEN
	}

	e.open.Add(1)
	h := addObject(&fileEntry{f: f, fs: e})
	(*VFSFile)(unsafe.Pointer(p)).fsFile = h
	if pOutFlags != 0 {
		// SQLite reads the out-flags for SQLITE_OPEN_READONLY, and this VFS
		// cannot write: the xWrite, xTruncate and xSync of vfsio only fail.
		// Reported as writable, a write with the journal off went to an
		// xWrite vfsio did not have then, and panicked.
		*(*int32)(unsafe.Pointer(pOutFlags)) = sqlite3.SQLITE_OPEN_READONLY
	}
	(*VFSFile)(unsafe.Pointer(p)).base.pMethods = uintptr(unsafe.Pointer(&vfsio))
	return sqlite3.SQLITE_OK
}

func vfsRead(tls *libc.TLS, pFile uintptr, zBuf uintptr, iAmt int32, iOfst sqlite_int64) int32 {
	p := pFile
	f := getObject((*VFSFile)(unsafe.Pointer(p)).fsFile).(*fileEntry).f
	seeker, ok := f.(io.Seeker)
	if !ok {
		return sqlite3.SQLITE_IOERR_READ
	}

	if n, err := seeker.Seek(iOfst, io.SeekStart); err != nil || n != iOfst {
		return sqlite3.SQLITE_IOERR_READ
	}

	b := (*libc.RawMem)(unsafe.Pointer(zBuf))[:iAmt]
	n, err := io.ReadFull(f, b)
	if err == nil {
		return sqlite3.SQLITE_OK
	}

	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		clear(b[n:])
		return sqlite3.SQLITE_IOERR_SHORT_READ
	}

	return sqlite3.SQLITE_IOERR_READ
}

func vfsAccess(tls *libc.TLS, pVfs uintptr, zPath uintptr, flags int32, pResOut uintptr) int32 {
	if flags == sqlite3.SQLITE_ACCESS_READWRITE {
		*(*int32)(unsafe.Pointer(pResOut)) = 0
		return sqlite3.SQLITE_OK
	}

	fn := libc.GoString(zPath)
	fsys := getObject((*sqlite3_vfs)(unsafe.Pointer(pVfs)).pAppData).(*fsEntry).fsys
	if _, err := fs.Stat(fsys, fn); err != nil {
		*(*int32)(unsafe.Pointer(pResOut)) = 0
		return sqlite3.SQLITE_OK
	}

	*(*int32)(unsafe.Pointer(pResOut)) = 1
	return sqlite3.SQLITE_OK
}

// vfsFullPathname is the xFullPathname of a VFS made by New: names are
// resolved by the wrapped fs.FS, so the name is its own full pathname. A
// name that does not fit the buffer is refused; it used to be copied
// without its terminator.
func vfsFullPathname(tls *libc.TLS, pVfs uintptr, zPath uintptr, nPathOut int32, zPathOut uintptr) int32 {
	return putPathname(libc.GoString(zPath), nPathOut, zPathOut)
}

// putPathname is the tail of an xFullPathname: it copies s, terminated, into
// the nPathOut bytes at zPathOut, or refuses a name that does not fit there
// or cannot be a C string.
func putPathname(s string, nPathOut int32, zPathOut uintptr) int32 {
	if int64(len(s)) >= int64(nPathOut) {
		return sqlite3.SQLITE_CANTOPEN
	}

	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			return sqlite3.SQLITE_CANTOPEN
		}
	}

	b := (*libc.RawMem)(unsafe.Pointer(zPathOut))[: len(s)+1 : len(s)+1]
	copy(b, s)
	b[len(s)] = 0
	return sqlite3.SQLITE_OK
}

func vfsFileSize(tls *libc.TLS, pFile uintptr, pSize uintptr) int32 {
	p := pFile
	f := getObject((*VFSFile)(unsafe.Pointer(p)).fsFile).(*fileEntry).f
	fi, err := f.Stat()
	if err != nil {
		return sqlite3.SQLITE_IOERR_FSTAT
	}

	*(*sqlite_int64)(unsafe.Pointer(pSize)) = fi.Size()
	return sqlite3.SQLITE_OK
}

func vfsClose(tls *libc.TLS, pFile uintptr) int32 {
	p := pFile
	h := (*VFSFile)(unsafe.Pointer(p)).fsFile
	e := getObject(h).(*fileEntry)
	e.f.Close()
	removeObject(h)
	e.fs.open.Add(-1)
	return sqlite3.SQLITE_OK
}

// FS is a VFS registered with SQLite: a read-only one backed by a Go fs.FS,
// made by [New], or one made by [Register] from a [VFS].
type FS struct {
	cname    uintptr
	cvfs     uintptr
	open     *atomic.Int32 // files currently open through the VFS
	fsHandle uintptr
	name     string
	tls      *libc.TLS

	closed bool // under mu
}

// New creates a new sqlite VFS and registers it. If successful, the
// file system can be used with the URI parameter `?vfs=<returned name>`.
func New(fs fs.FS) (name string, _ *FS, _ error) {
	if fs == nil {
		return "", nil, fmt.Errorf("fs argument cannot be nil")
	}

	mu.Lock()

	defer mu.Unlock()

	tls := libc.NewTLS()
	e := &fsEntry{fsys: fs}
	h := addObject(e)

	name = fmt.Sprintf("vfs%x", h)
	cname, err := libc.CString(name)
	if err != nil {
		return "", nil, err
	}

	vfs := Xsqlite3_fsFS(tls, cname, h)
	if vfs == 0 {
		removeObject(h)
		libc.Xfree(tls, cname)
		tls.Close()
		return "", nil, fmt.Errorf("out of memory")
	}

	// The C side wires the vfsDelete transpiled from vfs.c, see vfsDeleteReadOnly
	*(*func(*libc.TLS, uintptr, uintptr, int32) int32)(unsafe.Pointer(&(*sqlite3_vfs)(unsafe.Pointer(vfs)).xDelete)) = vfsDeleteReadOnly
	setCurrentTimeInt64((*sqlite3_vfs)(unsafe.Pointer(vfs)))

	if rc := sqlite3.Xsqlite3_vfs_register(tls, vfs, libc.Bool32(false)); rc != sqlite3.SQLITE_OK {
		removeObject(h)
		libc.Xfree(tls, cname)
		libc.Xfree(tls, vfs)
		tls.Close()
		return "", nil, fmt.Errorf("registering VFS %s: %d", name, rc)
	}

	return name, &FS{name: name, cname: cname, cvfs: vfs, open: &e.open, fsHandle: h, tls: tls}, nil
}

// Close unregisters f and releases its resources.
//
// Every database opened through f must be closed first. SQLite keeps a
// pointer to the VFS in each connection and calls through it for as long as
// the connection is open, so while any file is open through f, Close does
// nothing and returns an error wrapping [ErrInUse]; close the databases and
// call it again. Close must also not run concurrently with opening a database
// through f: SQLite looks the VFS up before it opens the file, and nothing
// can tell Close about an open that is between the two.
func (f *FS) Close() error {
	mu.Lock()

	defer mu.Unlock()

	if f.closed {
		return nil
	}

	if n := f.open.Load(); n != 0 {
		return fmt.Errorf("%w: closing VFS %s with %d file(s) still open", ErrInUse, f.name, n)
	}

	f.closed = true
	rc := sqlite3.Xsqlite3_vfs_unregister(f.tls, f.cvfs)
	libc.Xfree(f.tls, f.cname)
	libc.Xfree(f.tls, f.cvfs)
	f.tls.Close()
	removeObject(f.fsHandle)
	if rc != 0 {
		return fmt.Errorf("unregistering VFS %s: %d", f.name, rc)
	}

	return nil
}

// setCurrentTimeInt64 makes p a version 2 sqlite3_vfs with xCurrentTimeInt64.
// The one vfs.c builds is version 1, so SQLite took the time from its
// xCurrentTime, which is time(0) as a Julian day: 'now' moved in steps of one
// second, in strftime('%f', 'now'), julianday('now') and
// unixepoch('now', 'subsec') alike. SQLite prefers xCurrentTimeInt64 where
// there is one.
func setCurrentTimeInt64(p *sqlite3_vfs) {
	p.iVersion = 2
	*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&p.xCurrentTimeInt64)) = vfsCurrentTimeInt64
}

// vfsCurrentTimeInt64 is unixCurrentTimeInt64: milliseconds since noon of
// 24 November 4714 BC, the start of the Julian day count.
func vfsCurrentTimeInt64(tls *libc.TLS, pVfs uintptr, piNow uintptr) int32 {
	const unixEpoch = 24405875 * 8640000 // 1 January 1970 in those milliseconds
	*(*int64)(unsafe.Pointer(piNow)) = unixEpoch + time.Now().UnixMilli()
	return sqlite3.SQLITE_OK
}
