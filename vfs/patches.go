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
// again.
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

var (
	fToken uintptr

	// New, FS.Close
	mu sync.Mutex

	objectMu sync.Mutex
	objects  = map[uintptr]interface{}{}
)

func token() uintptr { return atomic.AddUintptr(&fToken, 1) }

func addObject(o interface{}) uintptr {
	t := token()
	objectMu.Lock()
	objects[t] = o
	objectMu.Unlock()
	return t
}

func getObject(t uintptr) interface{} {
	objectMu.Lock()
	o := objects[t]
	if o == nil {
		panic("internal error")
	}

	objectMu.Unlock()
	return o
}

func removeObject(t uintptr) {
	objectMu.Lock()
	if _, ok := objects[t]; !ok {
		panic("internal error")
	}

	delete(objects, t)
	objectMu.Unlock()
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
	fsys := getObject((*sqlite3_vfs)(unsafe.Pointer(pVfs)).pAppData).(fs.FS)
	f, err := fsys.Open(libc.GoString(zName))
	if err != nil {
		return sqlite3.SQLITE_CANTOPEN
	}

	h := addObject(f)
	(*VFSFile)(unsafe.Pointer(p)).fsFile = h
	if pOutFlags != 0 {
		*(*int32)(unsafe.Pointer(pOutFlags)) = int32(os.O_RDONLY)
	}
	(*VFSFile)(unsafe.Pointer(p)).base.pMethods = uintptr(unsafe.Pointer(&vfsio))
	return sqlite3.SQLITE_OK
}

func vfsRead(tls *libc.TLS, pFile uintptr, zBuf uintptr, iAmt int32, iOfst sqlite_int64) int32 {
	p := pFile
	f := getObject((*VFSFile)(unsafe.Pointer(p)).fsFile).(fs.File)
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
	fsys := getObject((*sqlite3_vfs)(unsafe.Pointer(pVfs)).pAppData).(fs.FS)
	if _, err := fs.Stat(fsys, fn); err != nil {
		*(*int32)(unsafe.Pointer(pResOut)) = 0
		return sqlite3.SQLITE_OK
	}

	*(*int32)(unsafe.Pointer(pResOut)) = 1
	return sqlite3.SQLITE_OK
}

func vfsFileSize(tls *libc.TLS, pFile uintptr, pSize uintptr) int32 {
	p := pFile
	f := getObject((*VFSFile)(unsafe.Pointer(p)).fsFile).(fs.File)
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
	f := getObject(h).(fs.File)
	f.Close()
	removeObject(h)
	return sqlite3.SQLITE_OK
}

// FS represents a SQLite read only file system backed by Go's fs.FS.
type FS struct {
	cname    uintptr
	cvfs     uintptr
	fsHandle uintptr
	name     string
	tls      *libc.TLS

	closed int32
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
	h := addObject(fs)

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

	return name, &FS{name: name, cname: cname, cvfs: vfs, fsHandle: h, tls: tls}, nil
}

// Close unregisters f and releases its resources.
func (f *FS) Close() error {
	if atomic.SwapInt32(&f.closed, 1) != 0 {
		return nil
	}

	mu.Lock()

	defer mu.Unlock()

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
