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
// SQLite is told a sector of 4096 bytes and no property of the storage: a
// write cut short by a loss of power may damage the 4096-byte sectors it
// touches and nothing further, and what Sync returned from is safe.
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

// LockLevel is a level of SQLite's file locking, described in
// https://www.sqlite.org/lockingv3.html. The values are SQLite's.
type LockLevel int32

// The lock levels. A database is read under LockShared and written under
// LockExclusive; LockReserved announces a writer while others still read, and
// LockPending keeps new readers out while the writer waits for the ones that
// are in.
const (
	LockNone      = LockLevel(lockNone)
	LockShared    = LockLevel(lockShared)
	LockReserved  = LockLevel(lockReserved)
	LockPending   = LockLevel(lockPending)
	LockExclusive = LockLevel(lockExclusive)
)

// ErrBusy is what [Locker.Lock] returns, alone or wrapped, when another
// Locker holds a lock that is in the way. SQLite then reports SQLITE_BUSY,
// and its busy handler decides whether to ask again.
var ErrBusy = errors.New("vfs: database is locked")

// LockOpener is implemented by a [VFS] whose databases can be opened from
// outside the registration: by another process, or through a second
// registration of the same storage. Without it the package keeps apart only
// the connections that go through one registration, see [Register].
type LockOpener interface {
	// OpenLock returns the lock of the database called name, a name
	// [VFS.FullPathname] returned. It is called after the Open of the first
	// file the registration has on the database, whatever the mode of that
	// Open, and the Locker serves every connection that opens the database
	// from then on; it is closed after the last file on the database. A nil
	// Locker with a nil error means that nothing outside the registration
	// can open this database. A Locker returned together with an error is
	// closed.
	//
	// OpenLock and the Close of a Locker are never called at the same time,
	// for any two databases of the registration: a database is not given a
	// second Locker before its first one is closed.
	OpenLock(name string) (Locker, error)
}

// Locker is the lock one registration holds on one database against every
// other Locker of that database, in this process or in another.
//
// The package runs SQLite's locking between the connections of the
// registration itself and calls the Locker only when the lock of all of them
// together has to change: Lock(LockShared) for the first reader,
// Unlock(LockNone) after the last, and the steps of a writer in between. The
// calls for one database never overlap.
//
// What a Locker must keep to, where held means held by another Locker:
//
//   - LockShared is refused while LockPending or LockExclusive is held. The
//     decision is atomic: a reader that was let in is in before a writer can
//     take LockExclusive.
//   - LockReserved is refused while LockReserved or more is held.
//   - LockPending is asked for only by the holder of LockReserved. It may be
//     refused for a moment, as it is by the way SQLite's own VFS lets a
//     reader in.
//   - LockExclusive is refused while any lock is held.
//   - A refusal changes nothing: the lock is what it was before the call.
//   - Unlock(LockShared) keeps the database locked for reading all the time:
//     it must not release the lock and take it again.
//   - A Locker whose process died holds nothing.
//   - Two Lockers of one database are two lockers, also in one process.
//
// The last rule is what POSIX record locks do not give. They belong to the
// process: two descriptors of one process never refuse each other, and
// closing any descriptor of the file drops every lock the process has on it.
// A Locker built on them needs one descriptor and one state per database for
// the whole process, shared by its registrations and by its Files, and not
// closed while a lock is held. Open file description locks on Linux and
// LockFileEx on Windows are held by the handle and need none of that. To
// share a database with programs that use SQLite's own VFS on unix, take the
// locks SQLite takes: the byte at offset 0x40000000 for LockPending, the next
// byte for LockReserved, and the 510 bytes after it for LockShared (a read
// lock) and LockExclusive (a write lock), entering LockShared under a read
// lock on the pending byte.
//
// The methods are called on the goroutine that runs the statement. They must
// not wait for a lock and must not run SQL on a connection that goes through
// the same VFS, and a panic in one is not recovered.
type Locker interface {
	// Lock raises the lock to level: to LockShared from LockNone, to
	// LockReserved from LockShared, to LockPending from LockReserved, and to
	// LockExclusive from LockPending or, to roll back the journal a dead
	// writer left, straight from LockShared. It returns ErrBusy when another
	// Locker is in the way. After any other error the package does not know
	// what the Locker holds, see Unlock.
	Lock(level LockLevel) error

	// Unlock lowers the lock to LockShared or to LockNone.
	//
	// After an error from Unlock, or one other than ErrBusy from Lock, the
	// package does not know what the Locker holds. While connections of the
	// registration still hold the database, it fails every request for a
	// higher lock and every read, write and sync of the database with an
	// I/O error, and no read or write of the database is under way while
	// Lock or Unlock runs. Once the connections have let go, it calls
	// Unlock(LockNone), again if need be before the next Lock, and goes on
	// when that succeeds.
	Unlock(level LockLevel) error

	// CheckReserved reports whether another Locker has taken LockReserved
	// and not given it up: whether a writer is at work whose journal must be
	// left alone. A Locker that went straight from LockShared to
	// LockExclusive need not be reported. CheckReserved may be called while
	// this Locker holds nothing, and it changes nothing of what the Locker
	// holds, also when it fails.
	CheckReserved() (bool, error)

	// Close releases the Locker and what it still holds. SQLite does not
	// look at the error.
	io.Closer
}

// lockState is the lock on one database of a registered VFS, shared by every
// file open on its name. It does within one registration what the POSIX locks
// of SQLite's own VFS do between processes, and like that VFS it takes the
// lock of them all together from outside, if the VFS has one to give.
type lockState struct {
	refs int    // files open on the name, under rwEntry.mu
	lk   Locker // the lock against other registrations and processes, or nil

	mu     sync.Mutex
	shared int     // of the files, the ones holding SHARED or more
	owner  *rwFile // the one holding RESERVED, PENDING or EXCLUSIVE, if any
	level  int32   // what owner holds

	// unsure is set when a call of lk failed and what lk holds is not known.
	// broken is set with it for as long as files still hold the lock: the
	// database is not touched then.
	//
	// io keeps reads and writes of the database apart from the calls that
	// can lose its lock. A read or write holds it shared; Lock and Unlock
	// of lk are called, and broken is changed, with io held exclusively and
	// under mu. So no byte is read under a lock that is being lost, and
	// broken is read under either of the two.
	unsure bool
	broken bool
	io     sync.RWMutex
}

// call makes a call of lk that can change what it holds: with reads and
// writes of the database kept out, and with the state that follows from the
// result set before they come back. The caller holds ls.mu.
func (ls *lockState) call(f func() error, releasesAll bool) error {
	ls.io.Lock()
	defer ls.io.Unlock()
	err := f()
	switch {
	case err == nil && releasesAll:
		ls.unsure, ls.broken = false, false
	case err != nil && !errors.Is(err, errRefused):
		ls.unsure, ls.broken = true, ls.shared > 0
	}
	return err
}

// errRefused marks, inside the package, an ErrBusy that Lock returned: the
// one error of a Locker after which what it holds is known.
var errRefused = errors.New("refused")

// raise takes the lock outside the registration up to level, which the
// in-process state already allows. It returns SQLITE_OK, also when there is
// no Locker, SQLITE_BUSY when another Locker is in the way and nothing has
// changed, or SQLITE_IOERR_LOCK. The caller holds ls.mu.
func (ls *lockState) raise(level int32) int32 {
	if ls.lk == nil {
		return sqlite3.SQLITE_OK
	}

	if ls.unsure {
		// No file holds the lock, or broken would have stopped the caller:
		// have the Locker let go of whatever a failed call left it with.
		if err := ls.call(func() error { return ls.lk.Unlock(LockNone) }, true); err != nil {
			return sqlite3.SQLITE_IOERR_LOCK
		}
	}

	err := ls.call(func() error {
		err := ls.lk.Lock(LockLevel(level))
		if errors.Is(err, ErrBusy) {
			return errRefused
		}

		return err
	}, false)
	switch {
	case err == nil:
		return sqlite3.SQLITE_OK
	case err == errRefused:
		return sqlite3.SQLITE_BUSY
	default:
		return sqlite3.SQLITE_IOERR_LOCK
	}
}

// rwEntry is what the pAppData of a VFS made by Register refers to.
type rwEntry struct {
	v    VFS
	open atomic.Int32 // files currently open through v, temporary ones included

	// mu guards locks, the refs of each lockState and the life of its
	// Locker: OpenLock and the Close of a Locker are called under it, so
	// that a database is never given a second Locker before its first one
	// is closed. It is not held while a lock is taken or released.
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
	lock int32      // level this file holds, under ls.mu
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

// lockStateOf returns the lock state of the named database with one more
// file counted on it. For the first file it asks a VFS that is a LockOpener
// for the Locker, and it returns nil if that fails.
func (e *rwEntry) lockStateOf(name string) *lockState {
	e.mu.Lock()
	defer e.mu.Unlock() // also behind a panic of OpenLock
	ls := e.locks[name]
	if ls == nil {
		ls = &lockState{}
		if lo, ok := e.v.(LockOpener); ok {
			lk, err := lo.OpenLock(name)
			if err != nil {
				if lk != nil {
					lk.Close()
				}
				return nil
			}

			ls.lk = lk
		}
		e.locks[name] = ls
	}
	ls.refs++
	return ls
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
			if rf.ls = e.lockStateOf(name); rf.ls == nil {
				f.Close()
				return sqlite3.SQLITE_CANTOPEN
			}
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
		// The handle and the count go whatever a Locker does, and no mutex
		// stays locked behind a panic of one.
		defer func() {
			removeObject(h)
			rf.e.open.Add(-1)
		}()
		ls := rf.ls
		if ls == nil {
			return
		}

		func() {
			ls.mu.Lock()
			defer ls.mu.Unlock()
			rf.unlock(lockNone)
		}()
		rf.e.mu.Lock()
		defer rf.e.mu.Unlock()
		if ls.refs--; ls.refs == 0 {
			delete(rf.e.locks, rf.name)
			if ls.lk != nil {
				ls.lk.Close()
			}
		}
	}()

	if err := rf.f.Close(); err != nil {
		return sqlite3.SQLITE_IOERR_CLOSE
	}

	return sqlite3.SQLITE_OK
}

// begin is called before a read or write of rf and reports whether it may go
// on: not when rf is a database whose lock against other registrations and
// processes is not known any more. After true, end must be called.
func (rf *rwFile) begin() bool {
	ls := rf.ls
	if ls == nil || ls.lk == nil {
		return true
	}

	ls.io.RLock()
	if ls.broken {
		ls.io.RUnlock()
		return false
	}

	return true
}

func (rf *rwFile) end() {
	if ls := rf.ls; ls != nil && ls.lk != nil {
		ls.io.RUnlock()
	}
}

func rwRead(tls *libc.TLS, pFile uintptr, zBuf uintptr, iAmt int32, iOfst sqlite_int64) int32 {
	rf := rwFileOf(pFile)
	if !rf.begin() {
		return sqlite3.SQLITE_IOERR_READ
	}

	defer rf.end()
	b := (*libc.RawMem)(unsafe.Pointer(zBuf))[:iAmt:iAmt]
	n, err := rf.f.ReadAt(b, iOfst)
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
	rf := rwFileOf(pFile)
	if !rf.begin() {
		return sqlite3.SQLITE_IOERR_WRITE
	}

	defer rf.end()
	b := (*libc.RawMem)(unsafe.Pointer(zBuf))[:iAmt:iAmt]
	n, err := rf.f.WriteAt(b, iOfst)
	switch {
	case errors.Is(err, ErrFull):
		return sqlite3.SQLITE_FULL
	case err != nil || n != len(b):
		return sqlite3.SQLITE_IOERR_WRITE
	}

	return sqlite3.SQLITE_OK
}

func rwTruncate(tls *libc.TLS, pFile uintptr, size sqlite_int64) int32 {
	rf := rwFileOf(pFile)
	if !rf.begin() {
		return sqlite3.SQLITE_IOERR_TRUNCATE
	}

	defer rf.end()
	if err := rf.f.Truncate(size); err != nil {
		return sqlite3.SQLITE_IOERR_TRUNCATE
	}

	return sqlite3.SQLITE_OK
}

func rwSync(tls *libc.TLS, pFile uintptr, flags int32) int32 {
	rf := rwFileOf(pFile)
	if !rf.begin() {
		return sqlite3.SQLITE_IOERR_FSYNC
	}

	defer rf.end()
	if err := rf.f.Sync(SyncFlag(flags)); err != nil {
		return sqlite3.SQLITE_IOERR_FSYNC
	}

	return sqlite3.SQLITE_OK
}

func rwFileSize(tls *libc.TLS, pFile uintptr, pSize uintptr) int32 {
	rf := rwFileOf(pFile)
	if !rf.begin() {
		return sqlite3.SQLITE_IOERR_FSTAT
	}

	defer rf.end()
	n, err := rf.f.Size()
	if err != nil || n < 0 {
		return sqlite3.SQLITE_IOERR_FSTAT
	}

	*(*sqlite_int64)(unsafe.Pointer(pSize)) = n
	return sqlite3.SQLITE_OK
}

// rwLock is SQLite's lock protocol between the files open on one database
// through one registered VFS, as unixLock runs it between the files of one
// process. Like unixLock it asks outside, here the Locker of the database,
// only when the lock of all the files together has to rise, and in the same
// order: a level is recorded for a file once the Locker has given it. SQLite
// asks for locks on database files only, and never for PENDING itself.
func rwLock(tls *libc.TLS, pFile uintptr, eLock int32) int32 {
	rf := rwFileOf(pFile)
	ls := rf.ls
	if ls == nil {
		return sqlite3.SQLITE_OK
	}

	ls.mu.Lock()
	defer ls.mu.Unlock()

	if rf.lock >= eLock {
		return sqlite3.SQLITE_OK
	}

	if ls.broken {
		return sqlite3.SQLITE_IOERR_LOCK
	}

	switch eLock {
	case lockShared:
		// Readers are let in next to a writer that holds RESERVED, and kept
		// out once it holds PENDING, or it would never get the file to itself.
		if ls.owner != nil && ls.level >= lockPending {
			return sqlite3.SQLITE_BUSY
		}

		if ls.shared == 0 {
			if rc := ls.raise(lockShared); rc != sqlite3.SQLITE_OK {
				return rc
			}
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

		if rc := ls.raise(lockReserved); rc != sqlite3.SQLITE_OK {
			return rc
		}

		ls.owner, ls.level, rf.lock = rf, lockReserved, lockReserved
	case lockExclusive:
		// From RESERVED, or from PENDING on a retry, to commit; or straight
		// from SHARED, to roll a hot journal back.
		if rf.lock == lockNone {
			return sqlite3.SQLITE_IOERR_LOCK
		}

		if ls.owner == rf && rf.lock == lockReserved {
			// A writer takes PENDING first, which keeps new readers out,
			// here and outside, and holds it when the next step is refused:
			// SQLite asks again.
			if rc := ls.raise(lockPending); rc != sqlite3.SQLITE_OK {
				return rc
			}

			ls.level, rf.lock = lockPending, lockPending
		}

		if ls.shared > 1 {
			// Others are still in: readers, or a writer, which is a reader
			// too. A reader that came to roll a journal back stays a plain
			// reader: counted as a writer, it would make the others take the
			// journal for one of a transaction in progress and read the
			// database as it is.
			return sqlite3.SQLITE_BUSY
		}

		if rc := ls.raise(lockExclusive); rc != sqlite3.SQLITE_OK {
			return rc
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

	rf.ls.mu.Lock()
	defer rf.ls.mu.Unlock()
	return rf.unlock(eLock)
}

// unlock lowers the lock rf holds to level, which is lockShared or lockNone,
// and with it the lock outside when rf was the writer or the last reader. The
// file is lowered whatever the Locker answers, as SQLite's pager takes it to
// be. The caller holds rf.ls.mu.
func (rf *rwFile) unlock(level int32) int32 {
	ls := rf.ls
	if rf.lock <= level {
		if level == lockNone && ls.unsure && ls.shared == 0 {
			// A Lock that failed for the first reader may have taken the
			// lock all the same. This is the unlock SQLite sends after it:
			// the next reader may be long in coming.
			if err := ls.call(func() error { return ls.lk.Unlock(LockNone) }, true); err != nil {
				return sqlite3.SQLITE_IOERR_UNLOCK
			}
		}
		return sqlite3.SQLITE_OK
	}

	writer := ls.owner == rf
	if writer {
		ls.owner, ls.level = nil, lockNone
	}
	if level == lockNone {
		ls.shared--
	}
	rf.lock = level
	if ls.lk == nil {
		return sqlite3.SQLITE_OK
	}

	out := LockShared
	switch {
	case ls.shared == 0:
		out = LockNone
	case !writer:
		return sqlite3.SQLITE_OK // a reader left and others stay
	case ls.unsure:
		return sqlite3.SQLITE_IOERR_UNLOCK // not known before, no better now
	}

	if err := ls.call(func() error { return ls.lk.Unlock(out) }, out == LockNone); err != nil {
		return sqlite3.SQLITE_IOERR_UNLOCK
	}

	return sqlite3.SQLITE_OK
}

// rwCheckReservedLock tells SQLite whether some connection is writing the
// database. It decides whether a journal found next to the database is a hot
// one to be rolled back, or the journal of a transaction in progress.
func rwCheckReservedLock(tls *libc.TLS, pFile uintptr, pResOut uintptr) int32 {
	r, rc := rwFileOf(pFile).reserved()
	*(*int32)(unsafe.Pointer(pResOut)) = r
	return rc
}

// reserved reports whether a writer holds the database of rf: a file of the
// registration, or, by the word of the Locker, somebody outside it.
func (rf *rwFile) reserved() (r, rc int32) {
	ls := rf.ls
	if ls == nil {
		return 0, sqlite3.SQLITE_OK
	}

	ls.mu.Lock()
	defer ls.mu.Unlock() // also behind a panic of CheckReserved
	switch {
	case ls.owner != nil:
		r = 1
	case ls.lk != nil:
		held, err := ls.lk.CheckReserved()
		switch {
		case err != nil:
			rc = sqlite3.SQLITE_IOERR_CHECKRESERVEDLOCK
		case held:
			r = 1
		}
	}
	return r, rc
}

// No file control verb is implemented. SQLite takes SQLITE_NOTFOUND for an
// answer to every one it sends.
func rwFileControl(tls *libc.TLS, pFile uintptr, op int32, pArg uintptr) int32 {
	return sqlite3.SQLITE_NOTFOUND
}

// SQLite takes the sector size for the unit a write cut short by a loss of
// power may damage, and journals a page together with the rest of its
// sector. Zero would be read as 512. SQLite's own VFS answers 4096 as well,
// but by default it also claims SQLITE_IOCAP_POWERSAFE_OVERWRITE, and with
// that the pager works with 512: this VFS claims nothing, so it is the more
// cautious of the two. It costs journal I/O: a journal header of 4096 bytes
// instead of 512, and with pages smaller than 4096 bytes the neighbours of a
// changed page.
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
// of any size works on a VFS that is safe for concurrent use. A second
// process, or a second registration of the same storage, is kept away from
// the database only by a VFS that implements [LockOpener]: without it, two
// of them that write lose each other's transactions and can damage the
// database. Two names for one database are told apart only by
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
