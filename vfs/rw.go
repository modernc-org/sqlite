// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vfs

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"sync"
	"sync/atomic"
	"unsafe"

	"modernc.org/libc"
	sqlite3 "modernc.org/sqlite/lib"
)

// OpenFlag is the set of SQLITE_OPEN_ flags SQLite passes to [VFS.Open]. The
// values are SQLite's.
type OpenFlag int32

// Flags of [VFS.Open]. One of OpenMainDB, OpenMainJournal, OpenSuperJournal
// and OpenWAL says what kind of file SQLite is opening.
const (
	OpenReadOnly      OpenFlag = sqlite3.SQLITE_OPEN_READONLY
	OpenReadWrite     OpenFlag = sqlite3.SQLITE_OPEN_READWRITE
	OpenCreate        OpenFlag = sqlite3.SQLITE_OPEN_CREATE
	OpenDeleteOnClose OpenFlag = sqlite3.SQLITE_OPEN_DELETEONCLOSE
	OpenExclusive     OpenFlag = sqlite3.SQLITE_OPEN_EXCLUSIVE
	OpenMainDB        OpenFlag = sqlite3.SQLITE_OPEN_MAIN_DB
	OpenMainJournal   OpenFlag = sqlite3.SQLITE_OPEN_MAIN_JOURNAL
	OpenSuperJournal  OpenFlag = sqlite3.SQLITE_OPEN_SUPER_JOURNAL
	OpenWAL           OpenFlag = sqlite3.SQLITE_OPEN_WAL
)

// AccessFlag is the question SQLite asks through [VFS.Access]. The values are
// SQLite's.
type AccessFlag int32

// Values of [AccessFlag].
const (
	AccessExists    AccessFlag = sqlite3.SQLITE_ACCESS_EXISTS
	AccessReadWrite AccessFlag = sqlite3.SQLITE_ACCESS_READWRITE
	AccessRead      AccessFlag = sqlite3.SQLITE_ACCESS_READ
)

// SyncFlag says how much [File.Sync] has to make durable. The values are
// SQLite's.
type SyncFlag int32

// Values of [SyncFlag]. SyncDataOnly is or-ed into one of the other two when
// only the content of the file, and not its metadata, has to reach the medium.
const (
	SyncNormal   SyncFlag = sqlite3.SQLITE_SYNC_NORMAL
	SyncFull     SyncFlag = sqlite3.SQLITE_SYNC_FULL
	SyncDataOnly SyncFlag = sqlite3.SQLITE_SYNC_DATAONLY
)

// ErrFull is what [File.WriteAt] returns, alone or wrapped, when the storage
// has no room left. SQLite then reports SQLITE_FULL instead of an I/O error.
var ErrFull = errors.New("vfs: storage is full")

// VFS is a file system SQLite can read and write through, registered with
// [Register].
//
// SQLite opens the database, its rollback journal, its super-journals and its
// write-ahead log through Open, each under the name FullPathname returned for
// the database with a suffix added: "-journal", "-wal", or "-mj" and nine more
// characters. Temporary files are not opened through the VFS: the package
// keeps them in memory, see [Register].
//
// The names come from SQLite and not all of them from the DSN: the name of
// a super-journal is read from the journal found next to a database, so
// whoever can put a journal there chooses a name that Open, Access and Delete
// are then called with. A VFS that maps names to something wider than its
// own storage has to confine them.
//
// The methods are called on the goroutine that runs the statement, so the
// connections of a pool call them, and the methods of the Files of one
// database, at the same time. The package serializes nothing but the locking
// of SQLite: a VFS and whatever its Files share must be safe for concurrent
// use. A method must not run SQL on a connection that goes through the same
// VFS, and a panic in a method is not recovered: it leaves the connection,
// and with it the database, in an undefined state.
type VFS interface {
	// Open opens the named file. Flags holds OpenReadWrite or OpenReadOnly;
	// with OpenCreate the file is to be created if it does not exist, and
	// with OpenCreate and OpenExclusive together the open must fail if it
	// does. The returned flags say how the file was opened in fact: return
	// OpenReadOnly for a file that cannot be written although OpenReadWrite
	// was asked for, and flags otherwise. No other bit of the result is used.
	// A File returned together with an error is closed.
	Open(name string, flags OpenFlag) (File, OpenFlag, error)

	// Delete removes the named file. With syncDir, the removal itself must be
	// durable when Delete returns. Deleting a file that does not exist is an
	// error for which errors.Is(err, fs.ErrNotExist) holds.
	Delete(name string, syncDir bool) error

	// Access reports whether the named file exists (AccessExists), exists
	// and can be read (AccessRead), or exists and can be read and written
	// (AccessReadWrite). A file that does not exist is not an error.
	Access(name string, flags AccessFlag) (bool, error)

	// FullPathname returns the canonical name of the database called name in
	// the DSN. Two names of one database must give the same result: it is
	// the name the other files are derived from, and the name under which the
	// connections of this process are kept from one another.
	FullPathname(name string) (string, error)
}

// File is an open file of a [VFS].
//
// SQLite owns the buffers it passes to ReadAt and WriteAt: a File must not
// keep them, and must not change the one given to WriteAt.
//
// SQLite is told what it assumes of a disk: a write cut short by a loss of
// power may damage the 4096 bytes around it and nothing further, and what
// Sync returned from is safe.
type File interface {
	// ReadAt is [io.ReaderAt]. A read that ends past the end of the file
	// returns what there is and io.EOF.
	io.ReaderAt

	// WriteAt is [io.WriterAt]. A write past the end of the file extends it,
	// and the bytes between the old end and the write read as zeros.
	io.WriterAt

	// Close releases the file. SQLite does not look at the error.
	io.Closer

	// Truncate sets the size of the file.
	Truncate(size int64) error

	// Sync returns when everything written to the file is durable.
	Sync(flags SyncFlag) error

	// Size returns the size of the file in bytes.
	Size() (int64, error)
}

// SQLite's lock levels, see https://www.sqlite.org/lockingv3.html.
const (
	lockNone      = int32(0)
	lockShared    = int32(sqlite3.SQLITE_LOCK_SHARED)
	lockReserved  = int32(sqlite3.SQLITE_LOCK_RESERVED)
	lockPending   = int32(sqlite3.SQLITE_LOCK_PENDING)
	lockExclusive = int32(sqlite3.SQLITE_LOCK_EXCLUSIVE)
)

// lockState is the lock on one database of a registered VFS, shared by every
// file open on its name. It does within one process what the POSIX locks of
// SQLite's own VFS do between processes.
type lockState struct {
	refs   int     // files open on the name
	shared int     // of those, the ones holding SHARED or more
	owner  *rwFile // the one holding RESERVED, PENDING or EXCLUSIVE, if any
	level  int32   // what owner holds
}

// rwEntry is what the pAppData of a VFS made by Register refers to.
type rwEntry struct {
	v    VFS
	open atomic.Int32 // files currently open through v, temporary ones included

	mu    sync.Mutex
	locks map[string]*lockState // by the name of the database
}

// rwFile is what the fsFile of a VFSFile opened through a registered VFS
// refers to.
type rwFile struct {
	f File
	e *rwEntry

	ls   *lockState // nil unless the file is a database
	name string     // key of ls in e.locks
	lock int32      // level this file holds, under e.mu
}

var vfsioRW = sqlite3_io_methods{
	iVersion: 1, // no shared memory, no memory mapping
}

func init() {
	*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&vfsioRW.xClose)) = rwClose
	*(*func(*libc.TLS, uintptr, uintptr, int32, sqlite_int64) int32)(unsafe.Pointer(&vfsioRW.xRead)) = rwRead
	*(*func(*libc.TLS, uintptr, uintptr, int32, sqlite_int64) int32)(unsafe.Pointer(&vfsioRW.xWrite)) = rwWrite
	*(*func(*libc.TLS, uintptr, sqlite_int64) int32)(unsafe.Pointer(&vfsioRW.xTruncate)) = rwTruncate
	*(*func(*libc.TLS, uintptr, int32) int32)(unsafe.Pointer(&vfsioRW.xSync)) = rwSync
	*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&vfsioRW.xFileSize)) = rwFileSize
	*(*func(*libc.TLS, uintptr, int32) int32)(unsafe.Pointer(&vfsioRW.xLock)) = rwLock
	*(*func(*libc.TLS, uintptr, int32) int32)(unsafe.Pointer(&vfsioRW.xUnlock)) = rwUnlock
	*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&vfsioRW.xCheckReservedLock)) = rwCheckReservedLock
	*(*func(*libc.TLS, uintptr, int32, uintptr) int32)(unsafe.Pointer(&vfsioRW.xFileControl)) = rwFileControl
	*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&vfsioRW.xSectorSize)) = rwSectorSize
	*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&vfsioRW.xDeviceCharacteristics)) = rwDeviceCharacteristics
}

func rwEntryOf(pVfs uintptr) *rwEntry {
	return getObject((*sqlite3_vfs)(unsafe.Pointer(pVfs)).pAppData).(*rwEntry)
}

func rwFileOf(pFile uintptr) *rwFile {
	return getObject((*VFSFile)(unsafe.Pointer(pFile)).fsFile).(*rwFile)
}

// rwOpen is the xOpen of a VFS made by Register.
func rwOpen(tls *libc.TLS, pVfs uintptr, zName uintptr, pFile uintptr, flags int32, pOutFlags uintptr) int32 {
	// SQLite calls xClose after a failed xOpen only if pMethods is set.
	*(*VFSFile)(unsafe.Pointer(pFile)) = VFSFile{}
	e := rwEntryOf(pVfs)
	rf := &rwFile{e: e}
	access := flags & (sqlite3.SQLITE_OPEN_READONLY | sqlite3.SQLITE_OPEN_READWRITE)
	if zName == 0 {
		// A temporary file: private to the connection and deleted when it is
		// closed, so there is nothing for the VFS to store.
		rf.f = &tempFile{}
	} else {
		name := libc.GoString(zName)
		f, out, err := e.v.Open(name, OpenFlag(flags))
		if err != nil || f == nil {
			if f != nil {
				f.Close() // it came with an error: SQLite will not close it
			}
			return sqlite3.SQLITE_CANTOPEN
		}

		rf.f = f
		if out&OpenReadOnly != 0 && out&OpenReadWrite == 0 {
			access = sqlite3.SQLITE_OPEN_READONLY
		}
		if flags&sqlite3.SQLITE_OPEN_MAIN_DB != 0 {
			rf.name = name
			e.mu.Lock()
			ls := e.locks[name]
			if ls == nil {
				ls = &lockState{}
				e.locks[name] = ls
			}
			ls.refs++
			rf.ls = ls
			e.mu.Unlock()
		}
	}

	e.open.Add(1)
	(*VFSFile)(unsafe.Pointer(pFile)).fsFile = addObject(rf)
	if pOutFlags != 0 {
		// SQLite acts on more bits of the out-flags than the access mode, so
		// only that is taken from the VFS.
		*(*int32)(unsafe.Pointer(pOutFlags)) = flags&^(sqlite3.SQLITE_OPEN_READONLY|sqlite3.SQLITE_OPEN_READWRITE) | access
	}
	(*VFSFile)(unsafe.Pointer(pFile)).base.pMethods = uintptr(unsafe.Pointer(&vfsioRW))
	return sqlite3.SQLITE_OK
}

// rwClose releases the handle, the lock and the open count whatever the
// file's Close does: SQLite ignores the result of xClose and never calls it
// twice.
func rwClose(tls *libc.TLS, pFile uintptr) int32 {
	h := (*VFSFile)(unsafe.Pointer(pFile)).fsFile
	rf := getObject(h).(*rwFile)
	defer func() {
		if rf.ls != nil {
			rf.e.mu.Lock()
			rf.unlock(lockNone)
			if rf.ls.refs--; rf.ls.refs == 0 {
				delete(rf.e.locks, rf.name)
			}
			rf.e.mu.Unlock()
		}
		removeObject(h)
		rf.e.open.Add(-1)
	}()

	if err := rf.f.Close(); err != nil {
		return sqlite3.SQLITE_IOERR_CLOSE
	}

	return sqlite3.SQLITE_OK
}

func rwRead(tls *libc.TLS, pFile uintptr, zBuf uintptr, iAmt int32, iOfst sqlite_int64) int32 {
	b := (*libc.RawMem)(unsafe.Pointer(zBuf))[:iAmt:iAmt]
	n, err := rwFileOf(pFile).f.ReadAt(b, iOfst)
	if n == len(b) && (err == nil || errors.Is(err, io.EOF)) {
		return sqlite3.SQLITE_OK
	}

	if n < 0 || n >= len(b) || !(errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)) {
		return sqlite3.SQLITE_IOERR_READ
	}

	// SQLite reads past the end of a file on purpose and expects zeros there.
	clear(b[n:])
	return sqlite3.SQLITE_IOERR_SHORT_READ
}

func rwWrite(tls *libc.TLS, pFile uintptr, zBuf uintptr, iAmt int32, iOfst sqlite_int64) int32 {
	b := (*libc.RawMem)(unsafe.Pointer(zBuf))[:iAmt:iAmt]
	n, err := rwFileOf(pFile).f.WriteAt(b, iOfst)
	switch {
	case errors.Is(err, ErrFull):
		return sqlite3.SQLITE_FULL
	case err != nil || n != len(b):
		return sqlite3.SQLITE_IOERR_WRITE
	}

	return sqlite3.SQLITE_OK
}

func rwTruncate(tls *libc.TLS, pFile uintptr, size sqlite_int64) int32 {
	if err := rwFileOf(pFile).f.Truncate(size); err != nil {
		return sqlite3.SQLITE_IOERR_TRUNCATE
	}

	return sqlite3.SQLITE_OK
}

func rwSync(tls *libc.TLS, pFile uintptr, flags int32) int32 {
	if err := rwFileOf(pFile).f.Sync(SyncFlag(flags)); err != nil {
		return sqlite3.SQLITE_IOERR_FSYNC
	}

	return sqlite3.SQLITE_OK
}

func rwFileSize(tls *libc.TLS, pFile uintptr, pSize uintptr) int32 {
	n, err := rwFileOf(pFile).f.Size()
	if err != nil || n < 0 {
		return sqlite3.SQLITE_IOERR_FSTAT
	}

	*(*sqlite_int64)(unsafe.Pointer(pSize)) = n
	return sqlite3.SQLITE_OK
}

// rwLock is SQLite's lock protocol between the files open on one database
// through one registered VFS, as unixLock runs it between the files of one
// process. SQLite asks for locks on database files only, and never for
// PENDING itself.
func rwLock(tls *libc.TLS, pFile uintptr, eLock int32) int32 {
	rf := rwFileOf(pFile)
	if rf.ls == nil {
		return sqlite3.SQLITE_OK
	}

	rf.e.mu.Lock()
	defer rf.e.mu.Unlock()

	ls := rf.ls
	if rf.lock >= eLock {
		return sqlite3.SQLITE_OK
	}

	switch eLock {
	case lockShared:
		// Readers are let in next to a writer that holds RESERVED, and kept
		// out once it holds PENDING, or it would never get the file to itself.
		if ls.owner != nil && ls.level >= lockPending {
			return sqlite3.SQLITE_BUSY
		}

		ls.shared++
		rf.lock = lockShared
	case lockReserved:
		if rf.lock != lockShared {
			return sqlite3.SQLITE_IOERR_LOCK
		}

		if ls.owner != nil {
			return sqlite3.SQLITE_BUSY
		}

		ls.owner, ls.level, rf.lock = rf, lockReserved, lockReserved
	case lockExclusive:
		// From RESERVED, or from PENDING on a retry, to commit; or straight
		// from SHARED, to roll a hot journal back.
		if rf.lock == lockNone {
			return sqlite3.SQLITE_IOERR_LOCK
		}

		if ls.shared > 1 {
			// Others are still in: readers, or a writer, which is a reader
			// too. A writer holds PENDING from here on, which keeps new
			// readers out, and SQLite asks again. A reader that came to
			// roll a journal back stays a plain reader: counted as a
			// writer, it would make the others take the journal for one of
			// a transaction in progress and read the database as it is.
			if ls.owner == rf {
				ls.level, rf.lock = lockPending, lockPending
			}
			return sqlite3.SQLITE_BUSY
		}

		ls.owner, ls.level, rf.lock = rf, lockExclusive, lockExclusive
	default:
		return sqlite3.SQLITE_IOERR_LOCK
	}

	return sqlite3.SQLITE_OK
}

func rwUnlock(tls *libc.TLS, pFile uintptr, eLock int32) int32 {
	rf := rwFileOf(pFile)
	if rf.ls == nil {
		return sqlite3.SQLITE_OK
	}

	if eLock != lockNone && eLock != lockShared {
		return sqlite3.SQLITE_IOERR_UNLOCK
	}

	rf.e.mu.Lock()
	rf.unlock(eLock)
	rf.e.mu.Unlock()
	return sqlite3.SQLITE_OK
}

// unlock lowers the lock rf holds to level, which is lockShared or lockNone.
// The caller holds rf.e.mu.
func (rf *rwFile) unlock(level int32) {
	if rf.lock <= level {
		return
	}

	ls := rf.ls
	if ls.owner == rf {
		ls.owner, ls.level = nil, lockNone
	}
	if level == lockNone {
		ls.shared--
	}
	rf.lock = level
}

// rwCheckReservedLock tells SQLite whether some connection is writing the
// database. It decides whether a journal found next to the database is a hot
// one to be rolled back, or the journal of a transaction in progress.
func rwCheckReservedLock(tls *libc.TLS, pFile uintptr, pResOut uintptr) int32 {
	rf := rwFileOf(pFile)
	var r int32
	if rf.ls != nil {
		rf.e.mu.Lock()
		if rf.ls.owner != nil {
			r = 1
		}
		rf.e.mu.Unlock()
	}
	*(*int32)(unsafe.Pointer(pResOut)) = r
	return sqlite3.SQLITE_OK
}

// No file control verb is implemented. SQLite takes SQLITE_NOTFOUND for an
// answer to every one it sends.
func rwFileControl(tls *libc.TLS, pFile uintptr, op int32, pArg uintptr) int32 {
	return sqlite3.SQLITE_NOTFOUND
}

// rwSectorSize answers what SQLite's own VFS answers. SQLite takes it for the
// unit a write cut short by a loss of power may damage, and journals a page
// together with the rest of its sector. Zero would be read as 512.
func rwSectorSize(tls *libc.TLS, pFile uintptr) int32 {
	return sqlite3.SQLITE_DEFAULT_SECTOR_SIZE
}

// No property of the storage is claimed.
func rwDeviceCharacteristics(tls *libc.TLS, pFile uintptr) int32 { return 0 }

func rwDelete(tls *libc.TLS, pVfs uintptr, zPath uintptr, dirSync int32) int32 {
	switch err := rwEntryOf(pVfs).v.Delete(libc.GoString(zPath), dirSync != 0); {
	case err == nil:
		return sqlite3.SQLITE_OK
	case errors.Is(err, fs.ErrNotExist):
		return sqlite3.SQLITE_IOERR_DELETE_NOENT
	default:
		return sqlite3.SQLITE_IOERR_DELETE
	}
}

func rwAccess(tls *libc.TLS, pVfs uintptr, zPath uintptr, flags int32, pResOut uintptr) int32 {
	ok, err := rwEntryOf(pVfs).v.Access(libc.GoString(zPath), AccessFlag(flags))
	if err != nil {
		return sqlite3.SQLITE_IOERR_ACCESS
	}

	*(*int32)(unsafe.Pointer(pResOut)) = libc.Bool32(ok)
	return sqlite3.SQLITE_OK
}

func rwFullPathname(tls *libc.TLS, pVfs uintptr, zPath uintptr, nPathOut int32, zPathOut uintptr) int32 {
	s, err := rwEntryOf(pVfs).v.FullPathname(libc.GoString(zPath))
	if err != nil {
		return sqlite3.SQLITE_CANTOPEN
	}

	return putPathname(s, nPathOut, zPathOut)
}

// tempChunk is the size of the pieces a tempFile is kept in.
const tempChunk = 64 << 10

// tempFile is a temporary file of SQLite: a statement journal, the spill of a
// sort, a temporary table, the copy VACUUM makes. It lives in memory and goes
// with its last reference. It is kept in pieces of tempChunk bytes, made
// when they are first written: the file takes about its size and never has
// to be copied to grow, and a hole in it costs nothing.
type tempFile struct {
	chunks map[int64][]byte // by offset / tempChunk; a missing one reads as zeros
	size   int64
}

func (t *tempFile) ReadAt(p []byte, off int64) (int, error) {
	if off < 0 {
		return 0, fs.ErrInvalid
	}

	if off >= t.size {
		return 0, io.EOF
	}

	n := int(min(int64(len(p)), t.size-off))
	for done := 0; done < n; {
		at := off + int64(done)
		in := int(at % tempChunk)
		part := p[done : done+min(n-done, tempChunk-in)]
		if c := t.chunks[at/tempChunk]; c != nil {
			copy(part, c[in:])
		} else {
			clear(part)
		}
		done += len(part)
	}
	if n < len(p) {
		return n, io.EOF
	}

	return n, nil
}

func (t *tempFile) WriteAt(p []byte, off int64) (int, error) {
	if off < 0 {
		return 0, fs.ErrInvalid
	}

	end := off + int64(len(p))
	if end < off {
		return 0, ErrFull
	}

	for done := 0; done < len(p); {
		at := off + int64(done)
		in := int(at % tempChunk)
		c := t.chunks[at/tempChunk]
		if c == nil {
			if t.chunks == nil {
				t.chunks = map[int64][]byte{}
			}
			c = make([]byte, tempChunk)
			t.chunks[at/tempChunk] = c
		}
		done += copy(c[in:], p[done:])
	}
	if len(p) != 0 && end > t.size {
		t.size = end
	}
	return len(p), nil
}

func (t *tempFile) Close() error { t.chunks, t.size = nil, 0; return nil }

func (t *tempFile) Truncate(size int64) error {
	if size < 0 {
		return fs.ErrInvalid
	}

	if size < t.size {
		// What is cut off must read as zeros if the file grows again.
		last := size / tempChunk
		for i, c := range t.chunks {
			switch {
			case i > last || size%tempChunk == 0 && i == last:
				delete(t.chunks, i)
			case i == last:
				clear(c[size%tempChunk:])
			}
		}
	}
	t.size = size
	return nil
}

func (t *tempFile) Sync(SyncFlag) error { return nil }

func (t *tempFile) Size() (int64, error) { return t.size, nil }

// Register registers v with SQLite as a VFS a database can be read and
// written through, and returns the name it was registered under, to be given
// as the vfs DSN query parameter, and the handle to unregister it with:
//
//	name, fsvfs, err := vfs.Register(v)
//	if err != nil {
//		return err
//	}
//
//	defer fsvfs.Close()
//
//	db, err := sql.Open("sqlite", "file:test.db?vfs="+name)
//
// The connections of this process that open one database through the returned
// VFS are kept from one another by the package, which runs SQLite's file
// locking between them: a [File] has no lock methods, and a database/sql pool
// of any size works on a VFS that is safe for concurrent use. Nothing keeps a
// second process, or a second registration of the same storage, away from
// the database. Two names for one database are told apart only by
// [VFS.FullPathname].
//
// The VFS has no shared memory, so the write-ahead log works only for a
// connection that holds the database exclusively. With "PRAGMA
// locking_mode=EXCLUSIVE" run before the database is first read, "PRAGMA
// journal_mode=WAL" works; without it the pragma leaves the journal mode as
// it was, and a database that already is in WAL mode cannot be opened at all.
// A WAL database therefore needs the pragma in the DSN and a pool of one
// connection.
//
// SQLite's temporary files are kept in memory and take that much of it for
// as long as they are open. VACUUM builds a copy of the whole database in
// one; a sort or a CREATE INDEX too large for SQLite's cache, a large
// temporary table and a large statement journal are others. Nothing bounds
// them: where SQLite's own VFS would fill a disk, SQL run through this one
// can make the process run out of memory.
//
// Registration is process-global, and [FS.Close] has the same contract as for
// a file system made by [New].
func Register(v VFS) (name string, _ *FS, _ error) {
	if v == nil {
		return "", nil, fmt.Errorf("vfs argument cannot be nil")
	}

	mu.Lock()

	defer mu.Unlock()

	tls := libc.NewTLS()
	e := &rwEntry{v: v, locks: map[string]*lockState{}}
	h := addObject(e)

	name = fmt.Sprintf("vfs%x", h)
	cname, err := libc.CString(name)
	if err != nil {
		removeObject(h)
		tls.Close()
		return "", nil, err
	}

	vfs := Xsqlite3_fsFS(tls, cname, h)
	if vfs == 0 {
		removeObject(h)
		libc.Xfree(tls, cname)
		tls.Close()
		return "", nil, fmt.Errorf("out of memory")
	}

	// The C side wires the methods of the read-only VFS. These are top-level
	// functions: the struct is C memory, where the garbage collector would
	// not see a closure.
	p := (*sqlite3_vfs)(unsafe.Pointer(vfs))
	*(*func(*libc.TLS, uintptr, uintptr, uintptr, int32, uintptr) int32)(unsafe.Pointer(&p.xOpen)) = rwOpen
	*(*func(*libc.TLS, uintptr, uintptr, int32) int32)(unsafe.Pointer(&p.xDelete)) = rwDelete
	*(*func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32)(unsafe.Pointer(&p.xAccess)) = rwAccess
	*(*func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32)(unsafe.Pointer(&p.xFullPathname)) = rwFullPathname
	setCurrentTimeInt64(p)

	if rc := sqlite3.Xsqlite3_vfs_register(tls, vfs, libc.Bool32(false)); rc != sqlite3.SQLITE_OK {
		removeObject(h)
		libc.Xfree(tls, cname)
		libc.Xfree(tls, vfs)
		tls.Close()
		return "", nil, fmt.Errorf("registering VFS %s: %d", name, rc)
	}

	return name, &FS{name: name, cname: cname, cvfs: vfs, open: &e.open, fsHandle: h, tls: tls}, nil
}
