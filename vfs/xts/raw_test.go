// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"testing"
	"unsafe"

	"modernc.org/libc"
	sqlite3 "modernc.org/sqlite/lib"
)

// rawFile is a file opened through a registered VFS the way SQLite opens
// one, for what cannot be reached through SQL: the result codes of the
// methods themselves.
type rawFile struct {
	t    testing.TB
	tls  *libc.TLS
	vfs  uintptr
	p    uintptr
	name uintptr
	buf  uintptr
}

const rawBuf = 1 << 17

// rawOpen opens name through the VFS called vfs with the flags of SQLite and
// returns the file, or nil, and the result code.
func rawOpen(t testing.TB, vfs, name string, flags int32) (*rawFile, int32) {
	t.Helper()
	tls := libc.NewTLS()
	t.Cleanup(tls.Close)
	cv := cstring(tls, vfs)
	defer sqlite3.Xsqlite3_free(tls, cv)

	f := &rawFile{t: t, tls: tls, vfs: sqlite3.Xsqlite3_vfs_find(tls, cv)}
	if f.vfs == 0 {
		t.Fatalf("no VFS %q", vfs)
	}

	v := vfsOf(f.vfs)
	f.p = alloc(tls, int(v.FszOsFile))
	f.buf = alloc(tls, rawBuf)
	var block uintptr
	if name != "" {
		// SQLite gives xOpen a name out of a block a VFS may walk through
		// in both directions: four zero bytes, the name of the database,
		// its query parameters, the names of its journal and of its
		// write-ahead log, three zero bytes.
		layout := "\x00\x00\x00\x00" + name + "\x00\x00" + name + "-journal\x00" + name + "-wal\x00\x00\x00\x00"
		block = alloc(tls, len(layout))
		copy(unsafe.Slice((*byte)(unsafe.Pointer(block)), len(layout)), layout)
		f.name = block + 4
	}

	t.Cleanup(func() {
		sqlite3.Xsqlite3_free(tls, f.p)
		sqlite3.Xsqlite3_free(tls, f.buf)
		sqlite3.Xsqlite3_free(tls, block)
	})
	rc := fn[func(*libc.TLS, uintptr, uintptr, uintptr, int32, uintptr) int32](v.FxOpen)(tls, f.vfs, f.name, f.p, flags, 0)
	if (*sqlite3.Tsqlite3_file)(unsafe.Pointer(f.p)).FpMethods == 0 {
		if rc == sqlite3.SQLITE_OK {
			t.Fatal("xOpen succeeded without methods")
		}

		return nil, rc
	}

	if rc != sqlite3.SQLITE_OK {
		// As SQLite does: a file with methods is closed, also after a
		// failed open.
		f.close()
		return nil, rc
	}

	return f, rc
}

func (f *rawFile) m() *sqlite3.Tsqlite3_io_methods { return methodsOf(f.p) }

func (f *rawFile) close() int32 {
	if (*sqlite3.Tsqlite3_file)(unsafe.Pointer(f.p)).FpMethods == 0 {
		return sqlite3.SQLITE_OK
	}

	rc := fn[func(*libc.TLS, uintptr) int32](f.m().FxClose)(f.tls, f.p)
	(*sqlite3.Tsqlite3_file)(unsafe.Pointer(f.p)).FpMethods = 0
	return rc
}

func (f *rawFile) read(n int, off int64) ([]byte, int32) {
	b := unsafe.Slice((*byte)(unsafe.Pointer(f.buf)), n)
	for i := range b {
		b[i] = 0xee
	}

	rc := fn[func(*libc.TLS, uintptr, uintptr, int32, int64) int32](f.m().FxRead)(f.tls, f.p, f.buf, int32(n), off)
	return append([]byte(nil), b...), rc
}

func (f *rawFile) write(p []byte, off int64) int32 {
	copy(unsafe.Slice((*byte)(unsafe.Pointer(f.buf)), len(p)), p)
	return fn[func(*libc.TLS, uintptr, uintptr, int32, int64) int32](f.m().FxWrite)(f.tls, f.p, f.buf, int32(len(p)), off)
}

func (f *rawFile) truncate(size int64) int32 {
	return fn[func(*libc.TLS, uintptr, int64) int32](f.m().FxTruncate)(f.tls, f.p, size)
}

func (f *rawFile) size() (int64, int32) {
	rc := fn[func(*libc.TLS, uintptr, uintptr) int32](f.m().FxFileSize)(f.tls, f.p, f.buf)
	return *(*int64)(unsafe.Pointer(f.buf)), rc
}

func (f *rawFile) lock(level int32) int32 {
	return fn[func(*libc.TLS, uintptr, int32) int32](f.m().FxLock)(f.tls, f.p, level)
}

func (f *rawFile) unlock(level int32) int32 {
	return fn[func(*libc.TLS, uintptr, int32) int32](f.m().FxUnlock)(f.tls, f.p, level)
}

func (f *rawFile) characteristics() int32 {
	return fn[func(*libc.TLS, uintptr) int32](f.m().FxDeviceCharacteristics)(f.tls, f.p)
}

func (f *rawFile) sectorSize() int32 {
	return fn[func(*libc.TLS, uintptr) int32](f.m().FxSectorSize)(f.tls, f.p)
}

const (
	rawRW      = sqlite3.SQLITE_OPEN_READWRITE | sqlite3.SQLITE_OPEN_CREATE
	rawDB      = rawRW | sqlite3.SQLITE_OPEN_MAIN_DB
	rawJournal = rawRW | sqlite3.SQLITE_OPEN_MAIN_JOURNAL
	rawSuper   = rawRW | sqlite3.SQLITE_OPEN_SUPER_JOURNAL | sqlite3.SQLITE_OPEN_EXCLUSIVE
	rawTemp    = rawRW | sqlite3.SQLITE_OPEN_TEMP_DB | sqlite3.SQLITE_OPEN_DELETEONCLOSE | sqlite3.SQLITE_OPEN_EXCLUSIVE
	rawWAL     = rawRW | sqlite3.SQLITE_OPEN_WAL
)

func (f *rawFile) fileControl(op int32, arg uintptr) int32 {
	return fn[func(*libc.TLS, uintptr, int32, uintptr) int32](f.m().FxFileControl)(f.tls, f.p, op, arg)
}

// journal opens, through the same VFS, the journal of the database f is:
// under the name SQLite has for it in the block of names f was opened from,
// as the pager of a connection does.
func (f *rawFile) journal() *rawFile {
	f.t.Helper()
	j := &rawFile{t: f.t, tls: f.tls, vfs: f.vfs, buf: f.buf}
	j.name = sqlite3.Xsqlite3_filename_journal(f.tls, f.name)
	j.p = alloc(f.tls, int(vfsOf(f.vfs).FszOsFile))
	f.t.Cleanup(func() { sqlite3.Xsqlite3_free(f.tls, j.p) })
	if rc := fn[func(*libc.TLS, uintptr, uintptr, uintptr, int32, uintptr) int32](vfsOf(f.vfs).FxOpen)(f.tls, f.vfs, j.name, j.p, rawJournal, 0); rc != sqlite3.SQLITE_OK {
		f.t.Fatalf("opening the journal: %d", rc)
	}

	return j
}
