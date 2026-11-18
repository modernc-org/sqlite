// Copyright 2022 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vfs

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"testing/iotest"
	"time"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libc/sys/types"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func E(err error) string {
	if err == nil {
		return ""
	}
	if err, ok := err.(*sqlite.Error); ok {
		return sqlite.ErrorCodeString[err.Code()]
	}
	return err.Error()
}

func TestVFS(t *testing.T) {
	const dbname = "canary.db"

	tmpdbdir := t.TempDir()

	func() {
		db, err := sql.Open("sqlite", "file:"+filepath.Join(tmpdbdir, dbname))
		if err != nil {
			t.Fatalf("unexpected failure to open database, %s", err.Error())
		}
		defer db.Close()

		_, err = db.Exec("create table 'test' ('name' varchar(32) not null, primary key('name') )")
		if err != nil {
			t.Fatalf("unexpected create table error, %s", E(err))
		}

		_, err = db.Exec("insert into 'test' (name) values ('foobar')")
		if err != nil {
			t.Fatalf("unexpected insert error, %s", E(err))
		}
	}()

	vfsid, fs, err := New(os.DirFS(tmpdbdir))
	if err != nil {
		t.Fatalf("unexpected failure to register new vfs, %s", err.Error())
	}
	defer fs.Close()

	db, err := sql.Open("sqlite", "file:"+dbname+"?vfs="+vfsid)
	if err != nil {
		t.Fatalf("unexpected failure to open database, %s", err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from 'test'")
	if err != nil {
		t.Fatalf("unexpected select error, %s", E(err))
	}
	defer rows.Close()
	var got string
	if !rows.Next() {
		t.Fatalf("unexpected empty select result")
	}
	err = rows.Scan(&got)
	if err != nil {
		t.Fatalf("unexpected scan error, %s", E(err))
	}
	want := "foobar"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

// seedDB opens filepath.Join(dir, name) as a SQLite database on the OS
// filesystem, executes stmts against it, and closes the connection
// before returning so a subsequent reopen through a custom VFS sees a
// settled on-disk file.
func seedDB(t *testing.T, dir, name string, stmts ...string) {
	t.Helper()
	db, err := sql.Open("sqlite", "file:"+filepath.Join(dir, name))
	if err != nil {
		t.Fatalf("open database: %s", err)
	}
	defer db.Close()
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			t.Fatalf("exec %q: %s", s, E(err))
		}
	}
}

// openDBViaVFS registers fsys as a SQLite VFS and opens name through
// it. The VFS and *sql.DB are cleaned up via t.Cleanup.
func openDBViaVFS(t *testing.T, name string, fsys fs.FS) *sql.DB {
	t.Helper()
	vfsid, vfsFS, err := New(fsys)
	if err != nil {
		t.Fatalf("register vfs: %s", err)
	}
	t.Cleanup(func() { vfsFS.Close() })
	db, err := sql.Open("sqlite", "file:"+name+"?vfs="+vfsid)
	if err != nil {
		t.Fatalf("open database via vfs: %s", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

// wrappedReadFile delegates Read to an io.Reader adapter — typically
// one of the testing/iotest wrappers — while preserving Seek and the
// rest of the embedded fs.File. After every Seek the adapter is
// rebuilt from the underlying file so stateful adapters (e.g.
// iotest.DataErrReader, which buffers internally) see the new
// position. Stateless adapters like iotest.HalfReader are rebuilt
// unnecessarily but harmlessly.
type wrappedReadFile struct {
	fs.File
	seeker io.Seeker
	r      io.Reader
	wrap   func(io.Reader) io.Reader
}

func (f *wrappedReadFile) Read(p []byte) (int, error) {
	return f.r.Read(p)
}

func (f *wrappedReadFile) Seek(offset int64, whence int) (int64, error) {
	n, err := f.seeker.Seek(offset, whence)
	f.r = f.wrap(f.File)
	return n, err
}

// wrappedReadFS wraps every file opened from the underlying FS in a
// wrappedReadFile built from wrap.
type wrappedReadFS struct {
	fs.FS
	wrap func(io.Reader) io.Reader
}

func (fsys *wrappedReadFS) Open(name string) (fs.File, error) {
	f, err := fsys.FS.Open(name)
	if err != nil {
		return nil, err
	}
	seeker, ok := f.(io.Seeker)
	if !ok {
		return f, nil
	}
	return &wrappedReadFile{
		File:   f,
		seeker: seeker,
		r:      fsys.wrap(f),
		wrap:   fsys.wrap,
	}, nil
}

// TestVFSWriteRefused checks that a database opened through the VFS is
// read-only to SQLite. Writes must fail with SQLITE_READONLY, including with
// the journal off, where a file reported as writable reaches xWrite, which
// this VFS does not have.
func TestVFSWriteRefused(t *testing.T) {
	const dbname = "readonly_canary.db"
	dir := t.TempDir()
	seedDB(t, dir, dbname,
		"CREATE TABLE t (v TEXT NOT NULL)",
		"INSERT INTO t (v) VALUES ('hello')",
	)

	db := openDBViaVFS(t, dbname, os.DirFS(dir))
	db.SetMaxOpenConns(1)

	for _, q := range []string{
		"UPDATE t SET v = 'changed'",
		"PRAGMA journal_mode=OFF",
		"UPDATE t SET v = 'changed'",
		"DELETE FROM t",
	} {
		_, err := db.Exec(q)
		if q == "PRAGMA journal_mode=OFF" {
			if err != nil {
				t.Fatalf("%s: %s", q, E(err))
			}
			continue
		}
		var se *sqlite.Error
		if !errors.As(err, &se) || se.Code()&0xff != sqlite3.SQLITE_READONLY {
			t.Fatalf("%s: got %s, want SQLITE_READONLY", q, E(err))
		}
	}

	var got string
	if err := db.QueryRow("SELECT v FROM t").Scan(&got); err != nil || got != "hello" {
		t.Fatalf("after the refused writes: %q, %s", got, E(err))
	}
}

// TestVFSReadEOF verifies that vfsRead handles readers that return
// (n > 0, io.EOF) in a single call — valid per io.Reader, and produced
// here by iotest.DataErrReader — by mapping the short tail to
// SQLITE_IOERR_SHORT_READ rather than SQLITE_IOERR_READ.
func TestVFSReadEOF(t *testing.T) {
	const dbname = "eof_canary.db"
	dir := t.TempDir()
	seedDB(t, dir, dbname,
		"CREATE TABLE t (v TEXT NOT NULL)",
		"INSERT INTO t (v) VALUES ('hello')",
	)

	fsys := &wrappedReadFS{FS: os.DirFS(dir), wrap: iotest.DataErrReader}
	db := openDBViaVFS(t, dbname, fsys)

	var got string
	if err := db.QueryRow("SELECT v FROM t").Scan(&got); err != nil {
		t.Fatalf("query: %s", E(err))
	}
	if got != "hello" {
		t.Fatalf("got %q, want %q", got, "hello")
	}

	// The query shows only that SQLite accepted what vfsRead returned. The
	// VFS contract for a read that runs past the end of the file is the
	// return code SQLITE_IOERR_SHORT_READ and a zeroed unread tail, so both
	// are checked on vfsRead itself, reading 512 bytes of which the file
	// holds the last 100.
	const amt, have = 512, 100
	data, err := os.ReadFile(filepath.Join(dir, dbname))
	if err != nil {
		t.Fatal(err)
	}
	f, err := fsys.Open(dbname)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	tls := libc.NewTLS()
	defer tls.Close()
	pFile := libc.Xcalloc(tls, 1, types.Size_t(unsafe.Sizeof(VFSFile{})))
	defer libc.Xfree(tls, pFile)
	zBuf := libc.Xmalloc(tls, amt)
	defer libc.Xfree(tls, zBuf)
	buf := unsafe.Slice((*byte)(unsafe.Pointer(zBuf)), amt)
	for i := range buf {
		buf[i] = 0xff
	}
	h := addObject(&fileEntry{f: f})
	defer removeObject(h)
	(*VFSFile)(unsafe.Pointer(pFile)).fsFile = h

	if rc := vfsRead(tls, pFile, zBuf, amt, int64(len(data)-have)); rc != sqlite3.SQLITE_IOERR_SHORT_READ {
		t.Fatalf("vfsRead past the end of the file: rc %d, want SQLITE_IOERR_SHORT_READ (%d)", rc, sqlite3.SQLITE_IOERR_SHORT_READ)
	}
	if !bytes.Equal(buf[:have], data[len(data)-have:]) {
		t.Fatal("vfsRead past the end of the file: the bytes it read differ from the end of the file")
	}
	for i, c := range buf[have:] {
		if c != 0 {
			t.Fatalf("vfsRead past the end of the file: byte %d of the unread tail is %#x, want 0", have+i, c)
		}
	}
}

// TestVFSReadPartial verifies that vfsRead retries partial mid-stream
// reads (Read returning fewer bytes than requested with nil error),
// synthesized here via iotest.HalfReader.
func TestVFSReadPartial(t *testing.T) {
	const dbname = "partial_canary.db"
	dir := t.TempDir()
	seedDB(t, dir, dbname,
		"CREATE TABLE t (v TEXT NOT NULL)",
		"INSERT INTO t (v) VALUES ('world')",
	)

	fsys := &wrappedReadFS{FS: os.DirFS(dir), wrap: iotest.HalfReader}
	db := openDBViaVFS(t, dbname, fsys)

	var got string
	if err := db.QueryRow("SELECT v FROM t").Scan(&got); err != nil {
		t.Fatalf("query: %s", E(err))
	}
	if got != "world" {
		t.Fatalf("got %q, want %q", got, "world")
	}
}

// truncatedReadFS wraps an fs.FS so reads past a byte threshold return
// EOF while Stat still reports the real file size. This makes SQLite
// compute a page count that exceeds the readable data, forcing the
// zero-fill path in vfsRead on the partially-readable page. Without
// correct zero-fill, stale buffer contents could corrupt SQLite's page
// interpretation. testing/iotest has no equivalent.
type truncatedReadFS struct {
	fs.FS
	readLimit int64
}

func (fsys *truncatedReadFS) Open(name string) (fs.File, error) {
	f, err := fsys.FS.Open(name)
	if err != nil {
		return nil, err
	}
	seeker, ok := f.(io.Seeker)
	if !ok {
		return f, nil
	}
	return &truncatedReadFile{File: f, seeker: seeker, readLimit: fsys.readLimit}, nil
}

type truncatedReadFile struct {
	fs.File
	seeker    io.Seeker
	readLimit int64
}

func (f *truncatedReadFile) Read(p []byte) (int, error) {
	pos, err := f.seeker.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}
	if pos >= f.readLimit {
		return 0, io.EOF
	}
	if avail := f.readLimit - pos; int64(len(p)) > avail {
		p = p[:avail]
	}
	return f.File.Read(p)
}

func (f *truncatedReadFile) Seek(offset int64, whence int) (int64, error) {
	return f.seeker.Seek(offset, whence)
}

// TestVFSReadEOFZeroFill verifies that vfsRead zero-fills unread tail
// bytes when a page read encounters EOF before the buffer is full.
func TestVFSReadEOFZeroFill(t *testing.T) {
	const dbname = "zerofill_canary.db"
	dir := t.TempDir()
	seedDB(t, dir, dbname,
		"CREATE TABLE t (v TEXT NOT NULL)",
		`WITH RECURSIVE cnt(x) AS (VALUES(1) UNION ALL SELECT x+1 FROM cnt WHERE x < 100)
		 INSERT INTO t SELECT hex(zeroblob(50)) FROM cnt`,
	)

	fi, err := os.Stat(filepath.Join(dir, dbname))
	if err != nil {
		t.Fatal(err)
	}
	size := fi.Size()
	if size <= 8192 {
		t.Fatalf("expected database > 2 pages (8192 bytes), got %d", size)
	}

	// Truncate reads 2048 bytes before EOF so the last page is only
	// partially readable; earlier pages remain intact.
	fsys := &truncatedReadFS{FS: os.DirFS(dir), readLimit: size - 2048}
	db := openDBViaVFS(t, dbname, fsys)

	// Rows on intact pages should come back; the short-read page
	// exercises vfsRead's zero-fill. Query itself must succeed —
	// it only touches intact pages. Mid-iteration scan or Err
	// failures are tolerated because they may reflect corruption
	// detected on the truncated page.
	rows, err := db.Query("SELECT v FROM t")
	if err != nil {
		t.Fatalf("db.Query on intact pages: %v", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			t.Logf("rows.Scan (tolerated, truncated page): %v", err)
			break
		}
		count++
	}
	if err := rows.Err(); err != nil {
		t.Logf("rows.Err (tolerated, truncated page): %v", err)
	}
	if count == 0 {
		t.Fatal("expected at least one readable row from intact pages")
	}
}

// seedWALDB creates filepath.Join(srcDir, name) in WAL mode, runs stmts on it
// and copies the database and its -wal file to servedDir while the connection
// is still open: closing it would checkpoint the WAL and remove the file. The
// copy is a WAL database shipped together with its WAL, as one lands in an
// embed.FS when it is copied from a live application.
func seedWALDB(t *testing.T, srcDir, servedDir, name string, stmts ...string) {
	t.Helper()
	db, err := sql.Open("sqlite", "file:"+filepath.Join(srcDir, name)+"?_pragma=journal_mode(WAL)")
	if err != nil {
		t.Fatalf("open database: %s", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			t.Fatalf("exec %q: %s", s, E(err))
		}
	}
	for _, n := range []string{name, name + "-wal"} {
		b, err := os.ReadFile(filepath.Join(srcDir, n))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(servedDir, n), b, 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

// TestVFSWALReadOnlyClose opens a WAL database shipped together with its -wal
// file. Without xShmMap SQLite accepts WAL only in exclusive locking mode,
// where the wal-index lives in heap memory, and closing a WAL database
// checkpoints it unless the file is read-only to SQLite: reported writable,
// the checkpoint reached the xWrite this VFS did not have and the process
// panicked. The read goes through the WAL, the close leaves the files alone,
// and the database opens again.
func TestVFSWALReadOnlyClose(t *testing.T) {
	const dbname = "wal_canary.db"
	srcDir, servedDir := t.TempDir(), t.TempDir()
	seedWALDB(t, srcDir, servedDir, dbname,
		"CREATE TABLE t (v TEXT NOT NULL)",
		"INSERT INTO t (v) VALUES ('from the wal')",
	)
	walBefore, err := os.ReadFile(filepath.Join(servedDir, dbname+"-wal"))
	if err != nil {
		t.Fatal(err)
	}
	if len(walBefore) == 0 {
		t.Fatal("the seeded WAL is empty")
	}

	vfsid, vfsFS, err := New(os.DirFS(servedDir))
	if err != nil {
		t.Fatalf("register vfs: %s", err)
	}
	defer vfsFS.Close()
	for i := 0; i < 2; i++ {
		db, err := sql.Open("sqlite", "file:"+dbname+"?vfs="+vfsid+"&_pragma=locking_mode(EXCLUSIVE)")
		if err != nil {
			t.Fatal(err)
		}
		var got string
		if err := db.QueryRow("SELECT v FROM t").Scan(&got); err != nil || got != "from the wal" {
			t.Fatalf("open %d: got %q, %s", i, got, E(err))
		}
		if err := db.Close(); err != nil {
			t.Fatalf("open %d: close: %s", i, E(err))
		}
		walAfter, err := os.ReadFile(filepath.Join(servedDir, dbname+"-wal"))
		if err != nil {
			t.Fatalf("open %d: after the close: %s", i, err)
		}
		if !bytes.Equal(walAfter, walBefore) {
			t.Fatalf("open %d: the close changed the WAL file", i)
		}
	}
}

// TestVFSDeleteRefused checks that nothing is deleted through the VFS. SQLite
// asks xDelete to remove a journal it believes it no longer needs, without
// consulting the read-only flag: when PRAGMA journal_mode leaves PERSIST or
// TRUNCATE, and on the first read of a 0-byte database that has a -journal
// file next to it. The xDelete transpiled from vfs.c aborted the process; the
// one New installs deletes nothing and fails with SQLITE_IOERR_DELETE, which
// both callers tolerate. The database names are relative, as the path SQLite
// asks to delete then is, so a file of that name in the working directory
// must survive, and so must the journal in the served file system.
func TestVFSDeleteRefused(t *testing.T) {
	const dbname, emptyname = "delete_canary.db", "empty_canary.db"
	served := t.TempDir()
	seedDB(t, served, dbname,
		"CREATE TABLE t (v TEXT NOT NULL)",
		"INSERT INTO t (v) VALUES ('hello')",
	)
	if err := os.WriteFile(filepath.Join(served, emptyname), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(served, emptyname+"-journal"), []byte("not a journal"), 0o644); err != nil {
		t.Fatal(err)
	}
	host := t.TempDir()
	t.Chdir(host)
	decoys := []string{dbname + "-journal", emptyname + "-journal"}
	for _, n := range decoys {
		if err := os.WriteFile(filepath.Join(host, n), []byte("a host file"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	vfsid, vfsFS, err := New(os.DirFS(served))
	if err != nil {
		t.Fatalf("register vfs: %s", err)
	}
	defer vfsFS.Close()

	// Leaving persist and truncate asks for the journal to be deleted
	db, err := sql.Open("sqlite", "file:"+dbname+"?vfs="+vfsid)
	if err != nil {
		t.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	for _, mode := range []string{"persist", "delete", "truncate", "off"} {
		var got string
		if err := db.QueryRow("PRAGMA journal_mode=" + mode).Scan(&got); err != nil || got != mode {
			t.Fatalf("PRAGMA journal_mode=%s: got %q, %s", mode, got, E(err))
		}
	}
	var v string
	if err := db.QueryRow("SELECT v FROM t").Scan(&v); err != nil || v != "hello" {
		t.Fatalf("after the journal_mode changes: %q, %s", v, E(err))
	}
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	// A 0-byte database next to a journal: the first read asks for the
	// journal to be deleted
	edb, err := sql.Open("sqlite", "file:"+emptyname+"?vfs="+vfsid)
	if err != nil {
		t.Fatal(err)
	}
	var n int
	if err := edb.QueryRow("SELECT count(*) FROM sqlite_master").Scan(&n); err != nil || n != 0 {
		t.Fatalf("the empty database: %d, %s", n, E(err))
	}
	if err := edb.Close(); err != nil {
		t.Fatal(err)
	}

	for _, n := range decoys {
		if _, err := os.Stat(filepath.Join(host, n)); err != nil {
			t.Errorf("the file in the working directory: %v", err)
		}
	}
	if _, err := os.Stat(filepath.Join(served, emptyname+"-journal")); err != nil {
		t.Errorf("the journal in the served file system: %v", err)
	}
}

// TestPutPathname checks the tail of both xFullPathname methods: a name that
// fits is copied with its terminator and nothing after it is touched, a name
// that does not fit is refused and nothing is written. The read-only VFS
// used to copy such a name without a terminator.
func TestPutPathname(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	const n = 16
	p := libc.Xmalloc(tls, n+8)
	defer libc.Xfree(tls, p)
	buf := unsafe.Slice((*byte)(unsafe.Pointer(p)), n+8)
	for _, c := range []struct {
		s    string
		want int32
	}{
		{"", sqlite3.SQLITE_OK},
		{"test.db", sqlite3.SQLITE_OK},
		{strings.Repeat("a", n-1), sqlite3.SQLITE_OK},
		{strings.Repeat("a", n), sqlite3.SQLITE_CANTOPEN},
		{strings.Repeat("a", 5000), sqlite3.SQLITE_CANTOPEN},
		{"a\x00b", sqlite3.SQLITE_CANTOPEN},
	} {
		for i := range buf {
			buf[i] = 0xff
		}
		if rc := putPathname(c.s, n, p); rc != c.want {
			t.Fatalf("putPathname(%d bytes): rc %d, want %d", len(c.s), rc, c.want)
		}

		written := 0
		if c.want == sqlite3.SQLITE_OK {
			written = len(c.s) + 1
			if string(buf[:len(c.s)]) != c.s || buf[len(c.s)] != 0 {
				t.Fatalf("putPathname(%q) wrote %q", c.s, buf[:written])
			}
		}
		for i, x := range buf[written:] {
			if x != 0xff {
				t.Fatalf("putPathname(%d bytes) touched byte %d", len(c.s), written+i)
			}
		}
	}

	// Through the read-only VFS, which this guards as well.
	zPath, err := libc.CString(strings.Repeat("a", 5000))
	if err != nil {
		t.Fatal(err)
	}
	defer libc.Xfree(tls, zPath)
	for i := range buf {
		buf[i] = 0xff
	}
	if rc := vfsFullPathname(tls, 0, zPath, n, p); rc != sqlite3.SQLITE_CANTOPEN {
		t.Fatalf("vfsFullPathname of a name longer than the buffer: rc %d", rc)
	}
	if !bytes.Equal(buf, bytes.Repeat([]byte{0xff}, n+8)) {
		t.Fatal("vfsFullPathname of a name longer than the buffer wrote to it")
	}
}

// sqliteNow returns what SQLite takes for the current time on db, in
// milliseconds since 1970.
func sqliteNow(t *testing.T, db *sql.DB) int64 {
	t.Helper()
	var (
		sec  int64
		frac string // SS.SSS
	)
	if err := db.QueryRow("SELECT CAST(strftime('%s','now') AS INTEGER), strftime('%f','now')").Scan(&sec, &frac); err != nil {
		t.Fatal(err)
	}

	i := strings.IndexByte(frac, '.')
	if i < 0 {
		t.Fatalf("strftime('%%f','now'): %q", frac)
	}

	ms, err := strconv.Atoi(frac[i+1:])
	if err != nil {
		t.Fatalf("strftime('%%f','now'): %q", frac)
	}

	return sec*1000 + int64(ms)
}

// checkNow fails unless SQLite's clock on db is the wall clock and moves in
// steps smaller than a second. A clock of whole seconds takes two values at
// most in the 140 ms the samples span. It cannot be told by its fraction: it
// reaches SQLite as a floating point Julian day, and one second in nine comes
// back a millisecond short, as .999.
func checkNow(t *testing.T, db *sql.DB) {
	t.Helper()
	seen := map[int64]bool{}
	for i := 0; i < 20; i++ {
		before := time.Now()
		now := sqliteNow(t, db)
		after := time.Now()
		seen[now] = true
		// If the wall clock did not move as the monotonic one did, it was set
		// while the sample was taken, and there is nothing to hold now against.
		if d := after.UnixMilli() - before.UnixMilli() - after.Sub(before).Milliseconds(); d >= -1 && d <= 1 {
			if now < before.UnixMilli() || now > after.UnixMilli() {
				t.Fatalf("SQLite's 'now' is %d ms, the clock around the query said %d to %d", now, before.UnixMilli(), after.UnixMilli())
			}
		}
		time.Sleep(7 * time.Millisecond)
	}
	if len(seen) < 3 {
		t.Fatalf("20 samples of 'now', 7 ms apart, took %d value(s): the clock moves in whole seconds", len(seen))
	}
}

// TestVFSCurrentTime checks that 'now' has milliseconds on a connection
// through a VFS made by New. It moved in steps of one second: the sqlite3_vfs
// was version 1, whose xCurrentTime is time(0).
func TestVFSCurrentTime(t *testing.T) {
	const dbname = "now.db"
	served := t.TempDir()
	seedDB(t, served, dbname, "CREATE TABLE t (v TEXT NOT NULL)")
	vfsid, vfsFS, err := New(os.DirFS(served))
	if err != nil {
		t.Fatalf("register vfs: %s", err)
	}
	defer vfsFS.Close()

	db, err := sql.Open("sqlite", "file:"+dbname+"?vfs="+vfsid)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	checkNow(t, db)
}
