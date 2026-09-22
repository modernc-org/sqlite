// Copyright 2022 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vfs exposes a Go [fs.FS] to SQLite as a read-only VFS.
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
// can only be read from.
//
// Registration is process-global, as SQLite's own VFS registry is. Each [New]
// registers a separate VFS under a fresh name, and [FS.Close] unregisters it
// again. Close every database opened through an [FS] before closing it: while
// any is open, [FS.Close] refuses with an error wrapping [ErrInUse].
package vfs

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"sync"
	"sync/atomic"
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

func vfsOpen(tls *libc.TLS, pVfs uintptr, zName uintptr, pFile uintptr, flags int32, pOutFlags uintptr) int32 {
	if zName == 0 {
		return sqlite3.SQLITE_IOERR
	}

	if flags&sqlite3.SQLITE_OPEN_MAIN_JOURNAL != 0 {
		return sqlite3.SQLITE_NOMEM
	}

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
		*(*int32)(unsafe.Pointer(pOutFlags)) = int32(os.O_RDONLY)
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

// FS represents a SQLite read only file system backed by Go's fs.FS.
type FS struct {
	cname    uintptr
	cvfs     uintptr
	entry    *fsEntry
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

	if rc := sqlite3.Xsqlite3_vfs_register(tls, vfs, libc.Bool32(false)); rc != sqlite3.SQLITE_OK {
		removeObject(h)
		libc.Xfree(tls, cname)
		libc.Xfree(tls, vfs)
		tls.Close()
		return "", nil, fmt.Errorf("registering VFS %s: %d", name, rc)
	}

	return name, &FS{name: name, cname: cname, cvfs: vfs, entry: e, fsHandle: h, tls: tls}, nil
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

	if n := f.entry.open.Load(); n != 0 {
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
