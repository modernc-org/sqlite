// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"sync"
	"testing"
	"unsafe"

	"modernc.org/libc"
	sqlite3 "modernc.org/sqlite/lib"
)

// hookVFS is a VFS for the tests that passes everything on to SQLite's
// default one and lets a test see and change what goes by: it is what the
// encrypting VFS is put on top of when a test needs to know what reached
// the storage, or to stop it.
type hookVFS struct {
	name string
	vfs  uintptr
	base uintptr

	mu sync.Mutex
	// written is everything written, by the name the file was opened
	// under; files without a name are under "".
	written map[string][]byte
	opens   map[string]int // opens by name
	// beforeDelete and afterDelete are called around the deletion of a file.
	beforeDelete, afterDelete func(name string)
	// failOpen, if it returns a result code other than SQLITE_OK, is the
	// result of the open; with leave set the file of the base stays open
	// with its methods, as a VFS may leave it.
	failOpen func(name string, flags int32) (rc int32, leave bool)
	// onName is called with the name of every open, as the pointer the
	// VFS was given.
	onName func(zName uintptr, flags int32)
	locks  int // locks asked for
	opened int // files the VFS gave its methods to
	closes int // files closed
	gone   bool
	tls    *libc.TLS
}

// unregister takes the VFS away while the test is still running.
func (h *hookVFS) unregister() {
	if !h.gone {
		h.gone = true
		sqlite3.Xsqlite3_vfs_unregister(h.tls, h.vfs)
	}
}

type hookFile struct {
	v    *hookVFS
	name string
	p    uintptr
}

var hookMethods = sqlite3.Tsqlite3_io_methods{FiVersion: 1}

func init() {
	set(&hookMethods.FxClose, hClose)
	set(&hookMethods.FxRead, hRead)
	set(&hookMethods.FxWrite, hWrite)
	set(&hookMethods.FxTruncate, hTruncate)
	set(&hookMethods.FxSync, hSync)
	set(&hookMethods.FxFileSize, hFileSize)
	set(&hookMethods.FxLock, hLock)
	set(&hookMethods.FxUnlock, hUnlock)
	set(&hookMethods.FxCheckReservedLock, hCheckReservedLock)
	set(&hookMethods.FxFileControl, hFileControl)
	set(&hookMethods.FxSectorSize, hSectorSize)
	set(&hookMethods.FxDeviceCharacteristics, hDeviceCharacteristics)
}

// newHookVFS registers a hookVFS for the time of the test.
func newHookVFS(t testing.TB) *hookVFS {
	t.Helper()
	tls := libc.NewTLS()
	h := &hookVFS{written: map[string][]byte{}, opens: map[string]int{}, tls: tls}
	h.base = sqlite3.Xsqlite3_vfs_find(tls, 0)
	handle := addObject(h)
	h.name = "hook" + string(rune('a'+handle%26)) + string(rune('a'+handle/26%26)) + string(rune('a'+handle/676%26))
	h.vfs = alloc(tls, int(unsafe.Sizeof(sqlite3.Tsqlite3_vfs{})))
	b, v := vfsOf(h.base), vfsOf(h.vfs)
	*v = *b // the methods that take no file and no name ignore the VFS they are given
	v.FiVersion = 1
	v.FpNext = 0
	v.FszOsFile = b.FszOsFile + int32(hdrSize)
	v.FzName = cstring(tls, h.name)
	v.FpAppData = handle
	set(&v.FxOpen, hOpen)
	set(&v.FxDelete, hDelete)
	set(&v.FxAccess, hAccess)
	set(&v.FxFullPathname, hFullPathname)
	if rc := sqlite3.Xsqlite3_vfs_register(tls, h.vfs, 0); rc != sqlite3.SQLITE_OK {
		t.Fatalf("registering the hook VFS: %d", rc)
	}

	t.Cleanup(func() {
		h.unregister()
		sqlite3.Xsqlite3_free(tls, v.FzName)
		sqlite3.Xsqlite3_free(tls, h.vfs)
		objects.Delete(handle)
		tls.Close()
	})
	return h
}

func hookOf(pVfs uintptr) *hookVFS {
	o, _ := objects.Load(vfsOf(pVfs).FpAppData)
	return o.(*hookVFS)
}

func hookFileOf(pFile uintptr) *hookFile {
	o, _ := objects.Load((*fileHdr)(unsafe.Pointer(pFile)).handle)
	return o.(*hookFile)
}

func hOpen(tls *libc.TLS, pVfs, zName, pFile uintptr, flags int32, pOut uintptr) int32 {
	h := hookOf(pVfs)
	hdr := (*fileHdr)(unsafe.Pointer(pFile))
	hdr.methods, hdr.handle = 0, 0
	name := ""
	if zName != 0 {
		name = libc.GoString(zName)
	}

	h.mu.Lock()
	onName := h.onName
	h.mu.Unlock()
	if onName != nil && zName != 0 {
		onName(zName, flags)
	}

	base := pFile + hdrSize
	rc := fn[func(*libc.TLS, uintptr, uintptr, uintptr, int32, uintptr) int32](vfsOf(h.base).FxOpen)(tls, h.base, zName, base, flags, pOut)
	if rc != sqlite3.SQLITE_OK {
		return rc
	}

	f := &hookFile{v: h, name: name, p: base}
	h.mu.Lock()
	fail := h.failOpen
	h.opens[name]++
	h.mu.Unlock()
	if fail != nil {
		if rc, leave := fail(name, flags); rc != sqlite3.SQLITE_OK {
			if leave {
				// The methods stay set: whoever opened the file closes it.
				h.mu.Lock()
				h.opened++
				h.mu.Unlock()
				hdr.handle = addObject(f)
				hdr.methods = uintptr(unsafe.Pointer(&hookMethods))
				return rc
			}

			fn[func(*libc.TLS, uintptr) int32](methodsOf(base).FxClose)(tls, base)
			return rc
		}
	}

	h.mu.Lock()
	h.opened++
	h.mu.Unlock()
	hdr.handle = addObject(f)
	hdr.methods = uintptr(unsafe.Pointer(&hookMethods))
	return rc
}

func hDelete(tls *libc.TLS, pVfs, zName uintptr, syncDir int32) int32 {
	h := hookOf(pVfs)
	name := libc.GoString(zName)
	h.mu.Lock()
	before, after := h.beforeDelete, h.afterDelete
	h.mu.Unlock()
	if before != nil {
		before(name)
	}

	rc := fn[func(*libc.TLS, uintptr, uintptr, int32) int32](vfsOf(h.base).FxDelete)(tls, h.base, zName, syncDir)
	if after != nil {
		after(name)
	}

	return rc
}

func hAccess(tls *libc.TLS, pVfs, zName uintptr, flags int32, pRes uintptr) int32 {
	h := hookOf(pVfs)
	return fn[func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32](vfsOf(h.base).FxAccess)(tls, h.base, zName, flags, pRes)
}

func hFullPathname(tls *libc.TLS, pVfs, zName uintptr, nOut int32, zOut uintptr) int32 {
	h := hookOf(pVfs)
	return fn[func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32](vfsOf(h.base).FxFullPathname)(tls, h.base, zName, nOut, zOut)
}

func hClose(tls *libc.TLS, pFile uintptr) int32 {
	hdr := (*fileHdr)(unsafe.Pointer(pFile))
	f := hookFileOf(pFile)
	f.v.mu.Lock()
	f.v.closes++
	f.v.mu.Unlock()
	rc := fn[func(*libc.TLS, uintptr) int32](methodsOf(f.p).FxClose)(tls, f.p)
	objects.Delete(hdr.handle)
	hdr.methods, hdr.handle = 0, 0
	return rc
}

func hRead(tls *libc.TLS, pFile, zBuf uintptr, n int32, off int64) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr, uintptr, int32, int64) int32](methodsOf(f.p).FxRead)(tls, f.p, zBuf, n, off)
}

func hWrite(tls *libc.TLS, pFile, zBuf uintptr, n int32, off int64) int32 {
	f := hookFileOf(pFile)
	f.v.mu.Lock()
	f.v.written[f.name] = append(f.v.written[f.name], unsafe.Slice((*byte)(unsafe.Pointer(zBuf)), n)...)
	f.v.mu.Unlock()
	return fn[func(*libc.TLS, uintptr, uintptr, int32, int64) int32](methodsOf(f.p).FxWrite)(tls, f.p, zBuf, n, off)
}

func hTruncate(tls *libc.TLS, pFile uintptr, size int64) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr, int64) int32](methodsOf(f.p).FxTruncate)(tls, f.p, size)
}

func hSync(tls *libc.TLS, pFile uintptr, flags int32) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr, int32) int32](methodsOf(f.p).FxSync)(tls, f.p, flags)
}

func hFileSize(tls *libc.TLS, pFile, pSize uintptr) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr, uintptr) int32](methodsOf(f.p).FxFileSize)(tls, f.p, pSize)
}

func hLock(tls *libc.TLS, pFile uintptr, level int32) int32 {
	f := hookFileOf(pFile)
	f.v.mu.Lock()
	f.v.locks++
	f.v.mu.Unlock()
	return fn[func(*libc.TLS, uintptr, int32) int32](methodsOf(f.p).FxLock)(tls, f.p, level)
}

func hUnlock(tls *libc.TLS, pFile uintptr, level int32) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr, int32) int32](methodsOf(f.p).FxUnlock)(tls, f.p, level)
}

func hCheckReservedLock(tls *libc.TLS, pFile, pRes uintptr) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr, uintptr) int32](methodsOf(f.p).FxCheckReservedLock)(tls, f.p, pRes)
}

func hFileControl(tls *libc.TLS, pFile uintptr, op int32, pArg uintptr) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr, int32, uintptr) int32](methodsOf(f.p).FxFileControl)(tls, f.p, op, pArg)
}

func hSectorSize(tls *libc.TLS, pFile uintptr) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr) int32](methodsOf(f.p).FxSectorSize)(tls, f.p)
}

func hDeviceCharacteristics(tls *libc.TLS, pFile uintptr) int32 {
	f := hookFileOf(pFile)
	return fn[func(*libc.TLS, uintptr) int32](methodsOf(f.p).FxDeviceCharacteristics)(tls, f.p)
}
