// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"modernc.org/libc"
	sqlite3 "modernc.org/sqlite/lib"
)

// ErrInUse is what the Close of a registration returns, wrapped, while files
// are open through it.
var ErrInUse = errors.New("xts: the VFS is in use")

// rcError is a result code of SQLite, from the base VFS or for it.
type rcError int32

func (e rcError) Error() string { return fmt.Sprintf("xts: SQLite result code %d", int32(e)) }

// rcOf is the result code for err, with def for an error that has none.
func rcOf(err error, def int32) int32 {
	var rc rcError
	switch {
	case err == nil:
		return sqlite3.SQLITE_OK
	case errors.As(err, &rc):
		return int32(rc)
	}

	return def
}

// fn makes a Go function of a function pointer of the transpiled code, and
// set stores one where the transpiled code calls it. Only top-level functions
// are stored: the structs are C memory, where the garbage collector would not
// see a closure.
func fn[T any](p uintptr) T                 { return *(*T)(unsafe.Pointer(&struct{ uintptr }{p})) }
func set[T any](dst *uintptr, f T)          { *(*T)(unsafe.Pointer(dst)) = f }
func vfsOf(p uintptr) *sqlite3.Tsqlite3_vfs { return (*sqlite3.Tsqlite3_vfs)(unsafe.Pointer(p)) }
func methodsOf(pFile uintptr) *sqlite3.Tsqlite3_io_methods {
	return (*sqlite3.Tsqlite3_io_methods)(unsafe.Pointer((*sqlite3.Tsqlite3_file)(unsafe.Pointer(pFile)).FpMethods))
}

// alloc returns n zeroed bytes of SQLite's memory, or 0.
func alloc(tls *libc.TLS, n int) uintptr {
	p := sqlite3.Xsqlite3_malloc64(tls, sqlite3.Tsqlite3_uint64(n))
	if p != 0 {
		clear(unsafe.Slice((*byte)(unsafe.Pointer(p)), n))
	}

	return p
}

// cstring returns s as a C string in SQLite's memory, or 0.
func cstring(tls *libc.TLS, s string) uintptr {
	p := alloc(tls, len(s)+1)
	if p != 0 {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), len(s)), s)
	}

	return p
}

// The objects the C structs refer to, by number: a registration in the
// pAppData of its sqlite3_vfs, a file in the header of its sqlite3_file.
var (
	objects sync.Map // uintptr -> *registration or *shimFile
	lastObj atomic.Uintptr

	// journals are the rollback journals that are open, by the address of
	// the name SQLite opened them under. SQLite keeps the names of a
	// database and of its journal in one block, so the journal a connection
	// has open for a database is found from the name of the database.
	journals sync.Map // uintptr -> *shimFile
)

// addObject gives o a number that no object has. The counter wraps on a
// 32-bit target after 2^32 files, and a number in use, or zero, is passed
// over.
func addObject(o any) uintptr {
	for {
		h := lastObj.Add(1)
		if h == 0 {
			continue
		}

		if _, taken := objects.LoadOrStore(h, o); !taken {
			return h
		}
	}
}

// names are the super-journals SQLite created through the package and has
// not deleted through it, with the base VFS they were created on. See
// journalFile for what it is for.
var names = nameTable{m: map[string]uintptr{}}

const maxNames = 65536

type nameTable struct {
	mu    sync.Mutex
	m     map[string]uintptr
	order []string
}

func (t *nameTable) add(name string, base uintptr) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.m[name]; !ok {
		t.order = append(t.order, name)
	}

	t.m[name] = base
	for len(t.order) > maxNames {
		delete(t.m, t.order[0])
		t.order = t.order[1:]
	}
}

func (t *nameTable) remove(name string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.m[name]; !ok {
		return
	}

	delete(t.m, name)
	for i, n := range t.order {
		if n == name {
			t.order = append(t.order[:i], t.order[i+1:]...)
			break
		}
	}
}

func (t *nameTable) has(name string, base uintptr) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	b, ok := t.m[name]
	return ok && b == base
}

// registration is one VFS made by Register.
type registration struct {
	name     string
	vfs      uintptr // our sqlite3_vfs
	cname    uintptr
	base     uintptr // the sqlite3_vfs under it
	baseName uintptr // its name, to find it again
	handle   uintptr
	open     atomic.Int32

	// mu is held shared by every call that comes in through the VFS and
	// exclusively by Close, so that a call is either done before the VFS is
	// closed or finds it closed.
	mu     sync.RWMutex
	closed bool

	db, journal, length *xts
	check               []byte

	// The methods of the base that take no file and no name, copied when
	// the registration was made.
	misc sqlite3.Tsqlite3_vfs
}

func regOf(pVfs uintptr) *registration {
	o, _ := objects.Load(vfsOf(pVfs).FpAppData)
	r, _ := o.(*registration)
	return r
}

// enter is called at the beginning of a method of the VFS that takes a
// name. It reports whether the call may go on, and then leave must follow.
func (r *registration) enter(tls *libc.TLS) bool {
	if r == nil {
		return false
	}

	r.mu.RLock()
	if r.closed || !r.alive(tls) {
		r.mu.RUnlock()
		return false
	}

	return true
}

func (r *registration) leave() { r.mu.RUnlock() }

// alive reports whether the base VFS is still the one the registration was
// made on. A base that was unregistered is not touched again.
func (r *registration) alive(tls *libc.TLS) bool {
	return sqlite3.Xsqlite3_vfs_find(tls, r.baseName) == r.base
}

// fileHdr is what a file of the VFS begins with; the file of the base VFS
// follows it.
type fileHdr struct {
	methods uintptr
	handle  uintptr
}

const hdrSize = unsafe.Sizeof(fileHdr{})

type fileKind int

const (
	kindRaw     fileKind = iota // passed through: opened as a super-journal
	kindDB                      // a database
	kindTemp                    // a temporary file, under a key of its own
	kindJournal                 // a rollback journal
)

// shimFile is an open file of a registration.
type shimFile struct {
	reg   *registration
	kind  fileKind
	f     file       // the content; nil for kindRaw
	st    *baseStore // the file of the base VFS
	zName uintptr    // the name SQLite opened it under, or 0
}

func fileOf(pFile uintptr) *shimFile {
	o, _ := objects.Load((*fileHdr)(unsafe.Pointer(pFile)).handle)
	f, _ := o.(*shimFile)
	return f
}

// baseStore is a file of the base VFS as a store. Go memory is not handed to
// the base: what is read or written goes through a buffer of SQLite's.
type baseStore struct {
	tls  *libc.TLS // of the call in progress
	p    uintptr   // the sqlite3_file of the base
	buf  uintptr
	nbuf int
	out  uintptr // where the base puts a size
}

func (b *baseStore) scratch(n int) (uintptr, error) {
	if b.nbuf < n {
		sqlite3.Xsqlite3_free(b.tls, b.buf)
		b.buf, b.nbuf = alloc(b.tls, n), n
		if b.buf == 0 {
			b.nbuf = 0
			return 0, rcError(sqlite3.SQLITE_NOMEM)
		}
	}

	return b.buf, nil
}

func (b *baseStore) free() {
	sqlite3.Xsqlite3_free(b.tls, b.buf)
	sqlite3.Xsqlite3_free(b.tls, b.out)
	b.buf, b.nbuf, b.out = 0, 0, 0
}

func (b *baseStore) readAt(p []byte, off int64) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	buf, err := b.scratch(len(p))
	if err != nil {
		return 0, err
	}

	n := len(p)
	switch rc := fn[func(*libc.TLS, uintptr, uintptr, int32, int64) int32](methodsOf(b.p).FxRead)(b.tls, b.p, buf, int32(len(p)), off); rc {
	case sqlite3.SQLITE_OK:
	case sqlite3.SQLITE_IOERR_SHORT_READ:
		// The base says only that the file ended, not where.
		size, err := b.size()
		if err != nil {
			return 0, err
		}

		n = int(min(max(size-off, 0), int64(len(p))))
	default:
		return 0, rcError(rc)
	}

	copy(p[:n], unsafe.Slice((*byte)(unsafe.Pointer(buf)), n))
	return n, nil
}

func (b *baseStore) writeAt(p []byte, off int64) error {
	if len(p) == 0 {
		return nil
	}

	buf, err := b.scratch(len(p))
	if err != nil {
		return err
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(buf)), len(p)), p)
	if rc := fn[func(*libc.TLS, uintptr, uintptr, int32, int64) int32](methodsOf(b.p).FxWrite)(b.tls, b.p, buf, int32(len(p)), off); rc != sqlite3.SQLITE_OK {
		return rcError(rc)
	}

	return nil
}

func (b *baseStore) truncate(size int64) error {
	if rc := fn[func(*libc.TLS, uintptr, int64) int32](methodsOf(b.p).FxTruncate)(b.tls, b.p, size); rc != sqlite3.SQLITE_OK {
		return rcError(rc)
	}

	return nil
}

func (b *baseStore) size() (int64, error) {
	// Not in the scratch: readAt asks for the size after a short read,
	// with what it read still there.
	if b.out == 0 {
		if b.out = alloc(b.tls, 8); b.out == 0 {
			return 0, rcError(sqlite3.SQLITE_NOMEM)
		}
	}

	if rc := fn[func(*libc.TLS, uintptr, uintptr) int32](methodsOf(b.p).FxFileSize)(b.tls, b.p, b.out); rc != sqlite3.SQLITE_OK {
		return 0, rcError(rc)
	}

	return *(*int64)(unsafe.Pointer(b.out)), nil
}

// Register registers with SQLite a VFS that keeps what it stores encrypted
// under key, on top of the VFS called base, and returns the name to give as
// the vfs query parameter of a DSN and what unregisters it again. The empty
// base is SQLite's default VFS; a VFS made by [modernc.org/sqlite/vfs.Register]
// is given by the name that returned. The base must stay registered for as
// long as the returned VFS is.
//
// The key is [KeySize] bytes that nobody else knows, from a random source or
// a key derivation function. It is not kept in a DSN.
//
// Close fails with [ErrInUse] while a file is open through the VFS.
func Register(base string, key []byte) (name string, _ io.Closer, err error) {
	r := &registration{}
	if r.db, err = derive(key, labelDatabase); err != nil {
		return "", nil, err
	}

	if r.journal, err = derive(key, labelJournal); err != nil {
		return "", nil, err
	}

	if r.length, err = derive(key, labelJournalLength); err != nil {
		return "", nil, err
	}

	c, err := derive(key, labelKeyCheck)
	if err != nil {
		return "", nil, err
	}

	r.check = keyCheck(c)
	tls := libc.NewTLS()
	defer tls.Close()

	var cbase uintptr
	if base != "" {
		if cbase = cstring(tls, base); cbase == 0 {
			return "", nil, rcError(sqlite3.SQLITE_NOMEM)
		}

		defer sqlite3.Xsqlite3_free(tls, cbase)
	}

	if r.base = sqlite3.Xsqlite3_vfs_find(tls, cbase); r.base == 0 {
		return "", nil, fmt.Errorf("xts: no VFS %q", base)
	}

	b := vfsOf(r.base)
	r.misc = *b
	r.handle = addObject(r)
	r.name = fmt.Sprintf("xts%x", r.handle)
	r.vfs = alloc(tls, int(unsafe.Sizeof(sqlite3.Tsqlite3_vfs{})))
	r.cname = cstring(tls, r.name)
	// The default VFS can change; the one found now is known by its name.
	r.baseName = cstring(tls, libc.GoString(b.FzName))
	if r.vfs == 0 || r.cname == 0 || r.baseName == 0 {
		r.release(tls)
		return "", nil, rcError(sqlite3.SQLITE_NOMEM)
	}

	v := vfsOf(r.vfs)
	v.FiVersion = min(b.FiVersion, 3)
	v.FszOsFile = b.FszOsFile + int32(hdrSize)
	v.FmxPathname = b.FmxPathname
	v.FzName = r.cname
	v.FpAppData = r.handle
	set(&v.FxOpen, xOpen)
	set(&v.FxDelete, xDelete)
	set(&v.FxAccess, xAccess)
	set(&v.FxFullPathname, xFullPathname)
	set(&v.FxDlOpen, xDlOpen)
	set(&v.FxDlError, xDlError)
	set(&v.FxDlSym, xDlSym)
	set(&v.FxDlClose, xDlClose)
	set(&v.FxRandomness, xRandomness)
	set(&v.FxSleep, xSleep)
	set(&v.FxCurrentTime, xCurrentTime)
	set(&v.FxGetLastError, xGetLastError)
	if v.FiVersion >= 2 && b.FxCurrentTimeInt64 != 0 {
		set(&v.FxCurrentTimeInt64, xCurrentTimeInt64)
	}

	if v.FiVersion >= 3 && b.FxSetSystemCall != 0 && b.FxGetSystemCall != 0 && b.FxNextSystemCall != 0 {
		set(&v.FxSetSystemCall, xSetSystemCall)
		set(&v.FxGetSystemCall, xGetSystemCall)
		set(&v.FxNextSystemCall, xNextSystemCall)
	} else if v.FiVersion >= 3 {
		v.FiVersion = 2
	}

	if rc := sqlite3.Xsqlite3_vfs_register(tls, r.vfs, 0); rc != sqlite3.SQLITE_OK {
		r.release(tls)
		return "", nil, fmt.Errorf("xts: registering the VFS: %w", rcError(rc))
	}

	return r.name, r, nil
}

func (r *registration) release(tls *libc.TLS) {
	objects.Delete(r.handle)
	sqlite3.Xsqlite3_free(tls, r.vfs)
	sqlite3.Xsqlite3_free(tls, r.cname)
	sqlite3.Xsqlite3_free(tls, r.baseName)
	r.vfs, r.cname, r.baseName = 0, 0, 0
}

// Close unregisters the VFS. The few hundred bytes SQLite was given for it
// stay allocated: a call that SQLite had already begun when the VFS was
// closed must find it closed, not freed.
func (r *registration) Close() error {
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil
	}

	if n := r.open.Load(); n != 0 {
		r.mu.Unlock()
		return fmt.Errorf("%w: closing %s with %d file(s) open", ErrInUse, r.name, n)
	}

	r.closed = true
	r.db, r.journal, r.length, r.check = nil, nil, nil, nil
	r.mu.Unlock()
	tls := libc.NewTLS()
	defer tls.Close()

	if rc := sqlite3.Xsqlite3_vfs_unregister(tls, r.vfs); rc != sqlite3.SQLITE_OK {
		return fmt.Errorf("xts: unregistering %s: %w", r.name, rcError(rc))
	}

	return nil
}

var ioMethods = sqlite3.Tsqlite3_io_methods{FiVersion: 1}

func init() {
	set(&ioMethods.FxClose, fClose)
	set(&ioMethods.FxRead, fRead)
	set(&ioMethods.FxWrite, fWrite)
	set(&ioMethods.FxTruncate, fTruncate)
	set(&ioMethods.FxSync, fSync)
	set(&ioMethods.FxFileSize, fFileSize)
	set(&ioMethods.FxLock, fLock)
	set(&ioMethods.FxUnlock, fUnlock)
	set(&ioMethods.FxCheckReservedLock, fCheckReservedLock)
	set(&ioMethods.FxFileControl, fFileControl)
	set(&ioMethods.FxSectorSize, fSectorSize)
	set(&ioMethods.FxDeviceCharacteristics, fDeviceCharacteristics)
}

// The query parameters the VFS looks at, as C strings.
var cNolock, cImmutable = mustCString("nolock"), mustCString("immutable")

func mustCString(s string) uintptr {
	p, err := libc.CString(s)
	if err != nil {
		panic(err)
	}

	return p
}

// uriBool reads a boolean query parameter of the name a database was opened
// under, as SQLite itself reads it.
func uriBool(tls *libc.TLS, zName, param uintptr) bool {
	return sqlite3.Xsqlite3_uri_boolean(tls, zName, param, 0) != 0
}

// randRead is where the keys of temporary files come from.
var randRead = rand.Read

const tempKinds = sqlite3.SQLITE_OPEN_TEMP_DB | sqlite3.SQLITE_OPEN_TRANSIENT_DB |
	sqlite3.SQLITE_OPEN_TEMP_JOURNAL | sqlite3.SQLITE_OPEN_SUBJOURNAL

func xOpen(tls *libc.TLS, pVfs, zName, pFile uintptr, flags int32, pOutFlags uintptr) int32 {
	// SQLite calls xClose after a failed xOpen only if the methods are set.
	hdr := (*fileHdr)(unsafe.Pointer(pFile))
	hdr.methods, hdr.handle = 0, 0
	r := regOf(pVfs)
	if !r.enter(tls) {
		return sqlite3.SQLITE_CANTOPEN
	}

	defer r.leave()

	f := &shimFile{reg: r, zName: zName}
	switch {
	case flags&sqlite3.SQLITE_OPEN_SUPER_JOURNAL != 0:
		f.kind = kindRaw
	case zName == 0 || flags&tempKinds != 0:
		f.kind = kindTemp
	case flags&sqlite3.SQLITE_OPEN_MAIN_JOURNAL != 0:
		f.kind = kindJournal
	case flags&sqlite3.SQLITE_OPEN_MAIN_DB != 0:
		f.kind = kindDB
		// With nolock SQLite locks nothing and still handles journals:
		// there is no moment at which the key could be checked. An
		// immutable database is not locked either, but SQLite neither
		// writes it nor looks at a journal of it.
		if uriBool(tls, zName, cNolock) && !uriBool(tls, zName, cImmutable) {
			return sqlite3.SQLITE_CANTOPEN
		}
	default:
		// A write-ahead log, which this format does not have, also not the
		// one SQLite keeps for a connection in exclusive locking mode; or a
		// kind of file the package does not know.
		return sqlite3.SQLITE_CANTOPEN
	}

	base := pFile + hdrSize
	(*sqlite3.Tsqlite3_file)(unsafe.Pointer(base)).FpMethods = 0
	rc := fn[func(*libc.TLS, uintptr, uintptr, uintptr, int32, uintptr) int32](vfsOf(r.base).FxOpen)(tls, r.base, zName, base, flags, pOutFlags)
	opened := (*sqlite3.Tsqlite3_file)(unsafe.Pointer(base)).FpMethods != 0
	if rc == sqlite3.SQLITE_OK && !opened {
		rc = sqlite3.SQLITE_CANTOPEN
	}

	f.st = &baseStore{tls: tls, p: base}
	if rc == sqlite3.SQLITE_OK {
		rc = f.setup(tls)
	}

	if rc != sqlite3.SQLITE_OK {
		if opened {
			fn[func(*libc.TLS, uintptr) int32](methodsOf(base).FxClose)(tls, base)
			(*sqlite3.Tsqlite3_file)(unsafe.Pointer(base)).FpMethods = 0
		}

		f.st.free()
		return rc
	}

	if f.kind == kindRaw && flags&sqlite3.SQLITE_OPEN_CREATE != 0 && zName != 0 {
		names.add(libc.GoString(zName), r.base)
	}

	if f.kind == kindJournal {
		journals.Store(zName, f)
	}

	hdr.handle = addObject(f)
	hdr.methods = uintptr(unsafe.Pointer(&ioMethods))
	r.open.Add(1)
	return sqlite3.SQLITE_OK
}

// setup makes the layout of a file the base has opened.
func (f *shimFile) setup(tls *libc.TLS) int32 {
	r := f.reg
	switch f.kind {
	case kindDB:
		f.f = &dbFile{st: f.st, c: r.db}
	case kindTemp:
		// What a temporary file holds is not read again after the process
		// that wrote it: every one gets a key of its own.
		key := make([]byte, 2*KeySize)
		if _, err := randRead(key); err != nil {
			return sqlite3.SQLITE_IOERR
		}

		c, err := newXTS(key)
		clear(key)
		if err != nil {
			return sqlite3.SQLITE_IOERR
		}

		f.f = &dbFile{st: f.st, c: c}
	case kindJournal:
		base := r.base
		f.f = &journalFile{
			st:      f.st,
			c:       r.journal,
			cl:      r.length,
			check:   r.check,
			shift:   max(minShift, int64(sectorSize(tls, f.st.p))),
			maxName: int64(vfsOf(r.base).FmxPathname),
			known:   func(name []byte) bool { return names.has(string(name), base) },
		}
	}

	return sqlite3.SQLITE_OK
}

func xDelete(tls *libc.TLS, pVfs, zName uintptr, syncDir int32) int32 {
	r := regOf(pVfs)
	if !r.enter(tls) {
		return sqlite3.SQLITE_IOERR_DELETE
	}

	defer r.leave()

	rc := fn[func(*libc.TLS, uintptr, uintptr, int32) int32](vfsOf(r.base).FxDelete)(tls, r.base, zName, syncDir)
	if rc == sqlite3.SQLITE_OK && zName != 0 {
		names.remove(libc.GoString(zName))
	}

	return rc
}

func xAccess(tls *libc.TLS, pVfs, zName uintptr, flags int32, pResOut uintptr) int32 {
	r := regOf(pVfs)
	if !r.enter(tls) {
		return sqlite3.SQLITE_IOERR_ACCESS
	}

	defer r.leave()

	return fn[func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32](vfsOf(r.base).FxAccess)(tls, r.base, zName, flags, pResOut)
}

func xFullPathname(tls *libc.TLS, pVfs, zName uintptr, nOut int32, zOut uintptr) int32 {
	r := regOf(pVfs)
	if !r.enter(tls) {
		return sqlite3.SQLITE_CANTOPEN
	}

	defer r.leave()

	return fn[func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32](vfsOf(r.base).FxFullPathname)(tls, r.base, zName, nOut, zOut)
}

// The methods below take no file and no name. They go to the functions the
// base had when the registration was made, which do not look at the base
// after it is gone.

func xDlOpen(tls *libc.TLS, pVfs, zName uintptr) uintptr {
	r := regOf(pVfs)
	if r == nil || r.misc.FxDlOpen == 0 {
		return 0
	}

	return fn[func(*libc.TLS, uintptr, uintptr) uintptr](r.misc.FxDlOpen)(tls, r.base, zName)
}

func xDlError(tls *libc.TLS, pVfs uintptr, n int32, zErr uintptr) {
	if r := regOf(pVfs); r != nil && r.misc.FxDlError != 0 {
		fn[func(*libc.TLS, uintptr, int32, uintptr)](r.misc.FxDlError)(tls, r.base, n, zErr)
	}
}

func xDlSym(tls *libc.TLS, pVfs, p, zSym uintptr) uintptr {
	r := regOf(pVfs)
	if r == nil || r.misc.FxDlSym == 0 {
		return 0
	}

	return fn[func(*libc.TLS, uintptr, uintptr, uintptr) uintptr](r.misc.FxDlSym)(tls, r.base, p, zSym)
}

func xDlClose(tls *libc.TLS, pVfs, p uintptr) {
	if r := regOf(pVfs); r != nil && r.misc.FxDlClose != 0 {
		fn[func(*libc.TLS, uintptr, uintptr)](r.misc.FxDlClose)(tls, r.base, p)
	}
}

func xRandomness(tls *libc.TLS, pVfs uintptr, n int32, zOut uintptr) int32 {
	r := regOf(pVfs)
	if r == nil || r.misc.FxRandomness == 0 {
		return 0
	}

	return fn[func(*libc.TLS, uintptr, int32, uintptr) int32](r.misc.FxRandomness)(tls, r.base, n, zOut)
}

func xSleep(tls *libc.TLS, pVfs uintptr, micro int32) int32 {
	r := regOf(pVfs)
	if r == nil || r.misc.FxSleep == 0 {
		return 0
	}

	return fn[func(*libc.TLS, uintptr, int32) int32](r.misc.FxSleep)(tls, r.base, micro)
}

func xCurrentTime(tls *libc.TLS, pVfs, pOut uintptr) int32 {
	r := regOf(pVfs)
	if r == nil || r.misc.FxCurrentTime == 0 {
		return sqlite3.SQLITE_ERROR
	}

	return fn[func(*libc.TLS, uintptr, uintptr) int32](r.misc.FxCurrentTime)(tls, r.base, pOut)
}

func xGetLastError(tls *libc.TLS, pVfs uintptr, n int32, zOut uintptr) int32 {
	r := regOf(pVfs)
	if r == nil || r.misc.FxGetLastError == 0 {
		return 0
	}

	return fn[func(*libc.TLS, uintptr, int32, uintptr) int32](r.misc.FxGetLastError)(tls, r.base, n, zOut)
}

func xCurrentTimeInt64(tls *libc.TLS, pVfs, pOut uintptr) int32 {
	r := regOf(pVfs)
	if r == nil || r.misc.FxCurrentTimeInt64 == 0 {
		return sqlite3.SQLITE_ERROR
	}

	return fn[func(*libc.TLS, uintptr, uintptr) int32](r.misc.FxCurrentTimeInt64)(tls, r.base, pOut)
}

func xSetSystemCall(tls *libc.TLS, pVfs, zName, p uintptr) int32 {
	r := regOf(pVfs)
	if r == nil || r.misc.FxSetSystemCall == 0 {
		return sqlite3.SQLITE_NOTFOUND
	}

	return fn[func(*libc.TLS, uintptr, uintptr, uintptr) int32](r.misc.FxSetSystemCall)(tls, r.base, zName, p)
}

func xGetSystemCall(tls *libc.TLS, pVfs, zName uintptr) uintptr {
	r := regOf(pVfs)
	if r == nil || r.misc.FxGetSystemCall == 0 {
		return 0
	}

	return fn[func(*libc.TLS, uintptr, uintptr) uintptr](r.misc.FxGetSystemCall)(tls, r.base, zName)
}

func xNextSystemCall(tls *libc.TLS, pVfs, zName uintptr) uintptr {
	r := regOf(pVfs)
	if r == nil || r.misc.FxNextSystemCall == 0 {
		return 0
	}

	return fn[func(*libc.TLS, uintptr, uintptr) uintptr](r.misc.FxNextSystemCall)(tls, r.base, zName)
}

func fClose(tls *libc.TLS, pFile uintptr) int32 {
	hdr := (*fileHdr)(unsafe.Pointer(pFile))
	f := fileOf(pFile)
	if f == nil {
		return sqlite3.SQLITE_OK
	}

	f.st.tls = tls
	if f.kind == kindJournal {
		journals.CompareAndDelete(f.zName, f)
	}

	rc := fn[func(*libc.TLS, uintptr) int32](methodsOf(f.st.p).FxClose)(tls, f.st.p)
	f.st.free()
	objects.Delete(hdr.handle)
	hdr.methods, hdr.handle = 0, 0
	f.reg.open.Add(-1)
	return rc
}

func fRead(tls *libc.TLS, pFile, zBuf uintptr, n int32, off int64) int32 {
	f := fileOf(pFile)
	f.st.tls = tls
	if f.kind == kindRaw {
		return fn[func(*libc.TLS, uintptr, uintptr, int32, int64) int32](methodsOf(f.st.p).FxRead)(tls, f.st.p, zBuf, n, off)
	}

	p := unsafe.Slice((*byte)(unsafe.Pointer(zBuf)), n)
	got, err := f.f.read(p, off)
	switch {
	case err != nil:
		return rcOf(err, sqlite3.SQLITE_IOERR_READ)
	case got < int(n):
		clear(p[got:])
		return sqlite3.SQLITE_IOERR_SHORT_READ
	}

	return sqlite3.SQLITE_OK
}

func fWrite(tls *libc.TLS, pFile, zBuf uintptr, n int32, off int64) int32 {
	f := fileOf(pFile)
	f.st.tls = tls
	if f.kind == kindRaw {
		return fn[func(*libc.TLS, uintptr, uintptr, int32, int64) int32](methodsOf(f.st.p).FxWrite)(tls, f.st.p, zBuf, n, off)
	}

	p := unsafe.Slice((*byte)(unsafe.Pointer(zBuf)), n)
	if f.kind == kindDB && off <= 18 && off+int64(n) >= 20 && p[18-off] == 2 && p[19-off] == 2 {
		// A header that says the database is in WAL mode. Stored, it
		// would make the database one that cannot be opened here.
		return sqlite3.SQLITE_IOERR_WRITE
	}

	return rcOf(f.f.write(p, off), sqlite3.SQLITE_IOERR_WRITE)
}

func fTruncate(tls *libc.TLS, pFile uintptr, size int64) int32 {
	f := fileOf(pFile)
	f.st.tls = tls
	if f.kind == kindRaw {
		return fn[func(*libc.TLS, uintptr, int64) int32](methodsOf(f.st.p).FxTruncate)(tls, f.st.p, size)
	}

	return rcOf(f.f.truncate(size), sqlite3.SQLITE_IOERR_TRUNCATE)
}

func fSync(tls *libc.TLS, pFile uintptr, flags int32) int32 {
	f := fileOf(pFile)
	return fn[func(*libc.TLS, uintptr, int32) int32](methodsOf(f.st.p).FxSync)(tls, f.st.p, flags)
}

func fFileSize(tls *libc.TLS, pFile, pSize uintptr) int32 {
	f := fileOf(pFile)
	f.st.tls = tls
	if f.kind == kindRaw {
		return fn[func(*libc.TLS, uintptr, uintptr) int32](methodsOf(f.st.p).FxFileSize)(tls, f.st.p, pSize)
	}

	size, err := f.f.size()
	if err != nil {
		return rcOf(err, sqlite3.SQLITE_IOERR_FSTAT)
	}

	*(*int64)(unsafe.Pointer(pSize)) = size
	return sqlite3.SQLITE_OK
}

func fLock(tls *libc.TLS, pFile uintptr, level int32) int32 {
	f := fileOf(pFile)
	f.st.tls = tls
	rc := fn[func(*libc.TLS, uintptr, int32) int32](methodsOf(f.st.p).FxLock)(tls, f.st.p, level)
	if rc != sqlite3.SQLITE_OK {
		return rc
	}

	if f.kind == kindDB && level == sqlite3.SQLITE_LOCK_SHARED {
		// SQLite asks for a shared lock when it holds none: the beginning
		// of a transaction. Before it looks at a journal, the key is known
		// to be the one of this database. Nothing is remembered from one
		// such lock to the next.
		if rc := f.checkKey(tls); rc != sqlite3.SQLITE_OK {
			fn[func(*libc.TLS, uintptr, int32) int32](methodsOf(f.st.p).FxUnlock)(tls, f.st.p, 0)
			return rc
		}
	}

	return rc
}

func fUnlock(tls *libc.TLS, pFile uintptr, level int32) int32 {
	f := fileOf(pFile)
	return fn[func(*libc.TLS, uintptr, int32) int32](methodsOf(f.st.p).FxUnlock)(tls, f.st.p, level)
}

func fCheckReservedLock(tls *libc.TLS, pFile, pResOut uintptr) int32 {
	f := fileOf(pFile)
	return fn[func(*libc.TLS, uintptr, uintptr) int32](methodsOf(f.st.p).FxCheckReservedLock)(tls, f.st.p, pResOut)
}

// sectorSize is the sector of a file of the base as SQLite will use it: a
// power of two of at least 512 and at most 65536.
func sectorSize(tls *libc.TLS, base uintptr) int32 {
	n := fn[func(*libc.TLS, uintptr) int32](methodsOf(base).FxSectorSize)(tls, base)
	s := int32(512)
	for s < n && s < maxShift {
		s <<= 1
	}

	return s
}

func fSectorSize(tls *libc.TLS, pFile uintptr) int32 {
	return sectorSize(tls, fileOf(pFile).st.p)
}

// What SQLite may assume of the file. A write that is not whole blocks
// rewrites the blocks at its two ends, so nothing is claimed that says a
// write changes only its own bytes or that an append needs no care; SQLite
// then takes a write to be able to damage the sectors it touches, and the
// blocks rewritten lie in those. What is passed on says nothing about
// writes.
const keptCharacteristics = sqlite3.SQLITE_IOCAP_UNDELETABLE_WHEN_OPEN |
	sqlite3.SQLITE_IOCAP_IMMUTABLE | sqlite3.SQLITE_IOCAP_SUBPAGE_READ

func fDeviceCharacteristics(tls *libc.TLS, pFile uintptr) int32 {
	f := fileOf(pFile)
	dc := fn[func(*libc.TLS, uintptr) int32](methodsOf(f.st.p).FxDeviceCharacteristics)(tls, f.st.p)
	if f.kind == kindRaw {
		return dc
	}

	return dc & keptCharacteristics
}

func fFileControl(tls *libc.TLS, pFile uintptr, op int32, pArg uintptr) int32 {
	f := fileOf(pFile)
	f.st.tls = tls
	forward := func() int32 {
		return fn[func(*libc.TLS, uintptr, int32, uintptr) int32](methodsOf(f.st.p).FxFileControl)(tls, f.st.p, op, pArg)
	}

	if f.kind != kindDB {
		return forward()
	}

	switch op {
	case sqlite3.SQLITE_FCNTL_PRAGMA:
		// char *azArg[]: the place of the answer, the name, the value.
		arg := unsafe.Slice((*uintptr)(unsafe.Pointer(pArg)), 3)
		if arg[1] != 0 && arg[2] != 0 && strings.EqualFold(libc.GoString(arg[1]), "journal_mode") {
			// As SQLite reads the value: the first mode it is a prefix of.
			if v := strings.ToLower(libc.GoString(arg[2])); v != "" && strings.HasPrefix("wal", v) {
				arg[0] = sqlite3.Xsqlite3_mprintf(tls, walMessage, 0)
				return sqlite3.SQLITE_ERROR
			}
		}

	case sqlite3.SQLITE_FCNTL_SYNC:
		if pArg != 0 {
			// The commit of a transaction over several databases, with
			// the name of its super-journal.
			if rc := f.checkSuper(tls, libc.GoString(pArg)); rc != sqlite3.SQLITE_OK {
				return rc
			}
		}

	case sqlite3.SQLITE_FCNTL_VFSNAME:
		rc := forward()
		if rc == sqlite3.SQLITE_OK {
			out := (*uintptr)(unsafe.Pointer(pArg))
			name := f.reg.name
			if *out != 0 {
				name += "/" + libc.GoString(*out)
				sqlite3.Xsqlite3_free(tls, *out)
			}

			*out = cstring(tls, name)
		}

		return rc
	}

	return forward()
}

// walMessage is the answer to PRAGMA journal_mode=WAL, a C string.
var walMessage = mustCString("the encrypting VFS has no write-ahead log")

// journalOf opens, through the base and without the cipher, the journal of
// the database this file is. The caller closes what it gets with done; a nil
// store with SQLITE_OK says that there is no journal.
//
// The name is the one SQLite itself has for the journal of this database,
// in the same block of names the database was opened from. The flags are
// those SQLite uses when it looks into the journal of another database
// (pager_delsuper): a read-only open as a super-journal.
func (f *shimFile) journalOf(tls *libc.TLS) (st *baseStore, done func(), rc int32) {
	r := f.reg
	b := vfsOf(r.base)
	name := sqlite3.Xsqlite3_filename_journal(tls, f.zName)
	res := alloc(tls, 8)
	file := alloc(tls, int(b.FszOsFile))
	free := func() {
		sqlite3.Xsqlite3_free(tls, res)
		sqlite3.Xsqlite3_free(tls, file)
	}

	if res == 0 || file == 0 {
		free()
		return nil, nil, sqlite3.SQLITE_NOMEM
	}

	if name == 0 || *(*byte)(unsafe.Pointer(name)) == 0 {
		free()
		return nil, func() {}, sqlite3.SQLITE_OK
	}

	if rc := fn[func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32](b.FxAccess)(tls, r.base, name, sqlite3.SQLITE_ACCESS_EXISTS, res); rc != sqlite3.SQLITE_OK {
		free()
		return nil, nil, rc
	}

	if *(*int32)(unsafe.Pointer(res)) == 0 {
		free()
		return nil, func() {}, sqlite3.SQLITE_OK
	}

	flags := int32(sqlite3.SQLITE_OPEN_READONLY | sqlite3.SQLITE_OPEN_SUPER_JOURNAL)
	rc = fn[func(*libc.TLS, uintptr, uintptr, uintptr, int32, uintptr) int32](b.FxOpen)(tls, r.base, name, file, flags, 0)
	opened := (*sqlite3.Tsqlite3_file)(unsafe.Pointer(file)).FpMethods != 0
	st = &baseStore{tls: tls, p: file}
	done = func() {
		if opened {
			fn[func(*libc.TLS, uintptr) int32](methodsOf(file).FxClose)(tls, file)
		}

		st.free()
		free()
	}

	if rc != sqlite3.SQLITE_OK || !opened {
		done()
		if rc == sqlite3.SQLITE_OK {
			rc = sqlite3.SQLITE_CANTOPEN
		}

		return nil, nil, rc
	}

	return st, done, sqlite3.SQLITE_OK
}

func (f *shimFile) journal(st *baseStore) *journalFile {
	r := f.reg
	base := r.base
	return &journalFile{
		st:      st,
		c:       r.journal,
		cl:      r.length,
		check:   r.check,
		shift:   minShift,
		maxName: int64(vfsOf(r.base).FmxPathname),
		known:   func(name []byte) bool { return names.has(string(name), base) },
	}
}

// checkKey decides whether the key of the registration is the key of this
// database, before SQLite acts on what it reads through it: with another
// key a journal reads as one whose writer died before it had a header, and
// SQLite would end the recovery there and delete it.
func (f *shimFile) checkKey(tls *libc.TLS) int32 {
	size, err := f.st.size()
	if err != nil {
		return rcOf(err, sqlite3.SQLITE_IOERR_FSTAT)
	}

	if size == 0 {
		return sqlite3.SQLITE_OK
	}

	var head [blockSize]byte
	n, err := f.f.read(head[:], 0)
	if err != nil {
		return rcOf(err, sqlite3.SQLITE_IOERR_READ)
	}

	if n == blockSize && bytes.Equal(head[:], []byte("SQLite format 3\x00")) {
		return sqlite3.SQLITE_OK
	}

	// Page 1 may have been torn by a crash. Then there is a journal that
	// restores it, and its prefix reached the storage before the page was
	// touched.
	if f.zName == 0 {
		return sqlite3.SQLITE_NOTADB
	}

	st, done, rc := f.journalOf(tls)
	if rc != sqlite3.SQLITE_OK {
		return rc
	}

	defer done()

	if st == nil {
		return sqlite3.SQLITE_NOTADB
	}

	r, err := st.size()
	if err != nil {
		return rcOf(err, sqlite3.SQLITE_IOERR_FSTAT)
	}

	if s, err := f.journal(st).prefix(r); err != nil {
		return rcOf(err, sqlite3.SQLITE_IOERR_READ)
	} else if s == 0 {
		return sqlite3.SQLITE_NOTADB
	}

	return sqlite3.SQLITE_OK
}

// checkSuper is asked before a transaction over several databases commits,
// with the name of its super-journal. The transaction may go on only if the
// journal the connection has open for this database shows the super-journal
// to readers that do not have the key, which journalFile.mayCommit says: a
// record of a super-journal the package did not see being created has no
// copy. A connection that has no journal open for the database keeps none
// on the storage, in journal_mode OFF or MEMORY: SQLite has written no
// record for it, and it takes part in the transaction as it does without
// the cipher.
func (f *shimFile) checkSuper(tls *libc.TLS, name string) int32 {
	if f.zName == 0 {
		return sqlite3.SQLITE_OK
	}

	// SQLite keeps the names of a database and of its journal in one block:
	// the journal this connection has open is the one opened under the
	// journal name of that block.
	o, ok := journals.Load(sqlite3.Xsqlite3_filename_journal(tls, f.zName))
	if !ok {
		return sqlite3.SQLITE_OK
	}

	own := o.(*shimFile)
	own.st.tls = tls
	switch ok, err := own.f.(*journalFile).mayCommit(name); {
	case err != nil:
		return rcOf(err, sqlite3.SQLITE_IOERR_READ)
	case !ok:
		return sqlite3.SQLITE_IOERR
	}

	return sqlite3.SQLITE_OK
}
