// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

const fill = "with recursive s(v) as (select 1 union all select v+1 from s where v < %d) "

// defaultVFS is the name of SQLite's default VFS.
func defaultVFS(t testing.TB) string {
	t.Helper()
	tls := libc.NewTLS()
	defer tls.Close()

	return libc.GoString(vfsOf(sqlite3.Xsqlite3_vfs_find(tls, 0)).FzName)
}

// backup copies the database of db into the database at uri with SQLite's
// backup interface.
func backup(t testing.TB, db *sql.DB, uri string) error {
	t.Helper()
	c, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

	return c.Raw(func(dc any) error {
		b, err := dc.(interface {
			NewBackup(string) (*sqlite.Backup, error)
		}).NewBackup(uri)
		if err != nil {
			return err
		}

		for more := true; more; {
			if more, err = b.Step(-1); err != nil {
				b.Finish()
				return err
			}
		}

		return b.Finish()
	})
}

// A plain database becomes an encrypted one and back, with VACUUM INTO and
// with the backup interface. A source in WAL mode goes through VACUUM INTO,
// which writes a rollback-mode header; the backup interface copies its
// header as it is, and that is refused.
func TestConvert(t *testing.T) {
	vfs := register(t, testKey(30))
	for _, mode := range []string{"delete", "wal"} {
		t.Run(mode, func(t *testing.T) {
			dir := t.TempDir()
			plain := open(t, filepath.Join(dir, "plain.db"), "", "_pragma=journal_mode("+mode+")")
			plain.SetMaxOpenConns(1)
			mustExec(t, plain, "create table t(a integer primary key, b text)")
			mustExec(t, plain, fmt.Sprintf(fill, 500)+"insert into t(b) select ? || v from s", marker)

			enc := filepath.Join(dir, "enc.db")
			mustExec(t, plain, fmt.Sprintf("vacuum into 'file:%s?vfs=%s'", enc, vfs))
			if b, err := os.ReadFile(enc); err != nil || bytes.Contains(b, []byte(marker)) || bytes.Contains(b, []byte("SQLite format 3")) {
				t.Fatalf("the copy holds plain text, or: %v", err)
			}

			db := open(t, enc, vfs, "")
			db.SetMaxOpenConns(1)
			if got := queryString(t, db, "select count(*) || ' ' || (select * from pragma_journal_mode) from t"); got != "500 delete" {
				t.Fatalf("the encrypted copy: %s", got)
			}

			mustExec(t, db, "insert into t(b) values('more')")

			// And back. The copy goes through the VFS of the connection
			// unless another is named: without one it is encrypted too.
			again := filepath.Join(dir, "again.db")
			mustExec(t, db, fmt.Sprintf("vacuum into 'file:%s'", again))
			if b, err := os.ReadFile(again); err != nil || bytes.Contains(b, []byte(marker)) {
				t.Fatalf("a copy made without naming a VFS holds plain text, or: %v", err)
			}

			back := filepath.Join(dir, "back.db")
			mustExec(t, db, fmt.Sprintf("vacuum into 'file:%s?vfs=%s'", back, defaultVFS(t)))
			if b, err := os.ReadFile(back); err != nil || !bytes.HasPrefix(b, []byte("SQLite format 3")) || !bytes.Contains(b, []byte(marker)) {
				t.Fatalf("the plain copy is not plain, or: %v", err)
			}

			// The backup interface, into an encrypted database that holds
			// something already.
			dst := filepath.Join(dir, "dst.db")
			d := open(t, dst, vfs, "")
			mustExec(t, d, "create table old(x); insert into old values('kept')")
			d.Close()
			before := filesOf(t, dir)
			err := backup(t, plain, fmt.Sprintf("file:%s?vfs=%s", dst, vfs))
			d = open(t, dst, vfs, "")
			if mode == "wal" {
				if err == nil {
					t.Fatal("the backup of a WAL-mode source succeeded")
				}

				sameFiles(t, "after the refused backup", before, filesOf(t, dir))
				if got := queryString(t, d, "select x from old"); got != "kept" {
					t.Fatalf("the destination after the refused backup: %s", got)
				}
			} else {
				if err != nil {
					t.Fatalf("backup: %v", err)
				}

				if got := queryString(t, d, "select count(*) from t"); got != "500" {
					t.Fatalf("the destination after the backup: %s", got)
				}
			}

			if got := queryString(t, d, "pragma integrity_check"); got != "ok" {
				t.Fatalf("integrity_check: %s", got)
			}
		})
	}
}

// The temporary files of a connection go through the VFS of its main
// database. With an encrypted main database they hold no plain text; with a
// plain main database and an encrypted one attached they do, which the
// documentation says, unless they are kept in memory.
func TestTemporaryFiles(t *testing.T) {
	h := newHookVFS(t)
	enc := registerOver(t, h.name, testKey(31))
	work := func(db *sql.DB, schema string) {
		mustExec(t, db, fmt.Sprintf("create table %s.t(a integer primary key, b text)", schema))
		mustExec(t, db, fmt.Sprintf(fill, 4000)+fmt.Sprintf("insert into %s.t(b) select ? || hex(randomblob(300)) || v from s", schema), marker)
		// A sort and a temporary table larger than what SQLite keeps in
		// memory, and a statement journal.
		var n int
		if err := db.QueryRow(fmt.Sprintf("select count(*) from (select b from %s.t order by b desc)", schema)).Scan(&n); err != nil || n != 4000 {
			t.Fatalf("sort: %d rows, %v", n, err)
		}

		mustExec(t, db, fmt.Sprintf("create temp table x as select * from %s.t", schema))
		mustExec(t, db, fmt.Sprintf("update %s.t set b = b || 'y'", schema))
	}

	temporary := func() []byte {
		h.mu.Lock()
		defer h.mu.Unlock()

		b := bytes.Clone(h.written[""])
		delete(h.written, "")
		return b
	}

	t.Run("encrypted main", func(t *testing.T) {
		db := open(t, filepath.Join(t.TempDir(), "a.db"), enc, "_pragma=cache_size(20)&_pragma=temp_store(1)")
		db.SetMaxOpenConns(1)
		temporary()
		work(db, "main")
		tmp := temporary()
		if len(tmp) < 1<<20 {
			t.Fatalf("only %d bytes of temporary files were written", len(tmp))
		}

		if bytes.Contains(tmp, []byte(marker)) || bytes.Contains(tmp, []byte("SQLite format 3")) {
			t.Fatal("a temporary file holds plain text")
		}

		h.mu.Lock()
		defer h.mu.Unlock()

		for name, b := range h.written {
			if bytes.Contains(b, []byte(marker)) {
				t.Errorf("%s was written with plain text", name)
			}
		}
	})

	for _, store := range []string{"1", "2"} {
		t.Run("plain main, temp_store="+store, func(t *testing.T) {
			dir := t.TempDir()
			db := open(t, filepath.Join(dir, "plain.db"), h.name, "_pragma=cache_size(20)&_pragma=temp_store("+store+")")
			db.SetMaxOpenConns(1)
			mustExec(t, db, fmt.Sprintf("attach database 'file:%s?vfs=%s' as e", filepath.Join(dir, "e.db"), enc))
			temporary()
			work(db, "e")
			tmp := temporary()
			if got := bytes.Contains(tmp, []byte(marker)); got != (store == "1") {
				t.Fatalf("plain text in temporary files: %v (%d bytes written)", got, len(tmp))
			}

			if b, err := os.ReadFile(filepath.Join(dir, "e.db")); err != nil || bytes.Contains(b, []byte(marker)) {
				t.Fatalf("the attached database holds plain text, or: %v", err)
			}
		})
	}
}

// nolock=1 is refused: SQLite would lock nothing and still handle journals.
// immutable=1 reads, and still tells a wrong key. What a value of such a
// parameter means is for SQLite to say, and it says so here: the same
// parameters on a plain database show whether SQLite locks it and whether
// it writes it.
func TestNolockImmutable(t *testing.T) {
	h := newHookVFS(t)
	right, wrong := registerOver(t, h.name, testKey(32)), registerOver(t, h.name, testKey(33))
	dir := t.TempDir()
	path, plain := filepath.Join(dir, "a.db"), filepath.Join(dir, "plain.db")
	for p, vfs := range map[string]string{path: right, plain: h.name} {
		db := open(t, p, vfs, "")
		mustExec(t, db, "create table t(a); insert into t values(42)")
		db.Close()
	}

	var refused, accepted int
	for _, v := range []string{
		"nolock=1", "nolock=true", "nolock=on", "nolock=yes", "nolock=0x1", "nolock=257", "nolock=256",
		"nolock=0", "nolock=false", "nolock=off", "nolock=no", "nolock=4294967297", "nolock=maybe",
		"nolock=1&immutable=1", "nolock=1&immutable=256", "nolock=1&immutable=2147483648", "immutable=1",
	} {
		// What SQLite makes of it.
		h.mu.Lock()
		h.locks = 0
		h.mu.Unlock()
		db := open(t, plain, h.name, v)
		if got := queryString(t, db, "select a from t"); got != "42" {
			t.Fatalf("%s on the plain database: %s", v, got)
		}

		_, err := db.Exec("update t set a = 42")
		writes := err == nil
		db.Close()
		h.mu.Lock()
		locks := h.locks != 0
		h.mu.Unlock()

		db = open(t, path, right, v)
		_, err = db.Exec("select a from t")
		db.Close()
		switch want := !locks && writes; {
		case want && codeOf(err) != sqlite3.SQLITE_CANTOPEN:
			t.Errorf("%s: SQLite locks nothing and writes; the open: %v", v, err)
		case !want && err != nil:
			t.Errorf("%s: SQLite locks: %v, writes: %v; the open: %v", v, locks, writes, err)
		case want:
			refused++
		default:
			accepted++
		}
	}

	if refused < 5 || accepted < 5 {
		t.Fatalf("%d values were refused and %d accepted", refused, accepted)
	}

	db := open(t, path, right, "immutable=1")
	if _, err := db.Exec("insert into t values(1)"); err == nil {
		t.Error("an immutable database was written")
	}

	db.Close()
	db = open(t, path, wrong, "immutable=1")
	if _, err := db.Exec("select * from t"); codeOf(err) != sqlite3.SQLITE_NOTADB {
		t.Errorf("immutable with the wrong key: %v", err)
	}
}

// Whatever fails while a file is being opened, the file of the base VFS is
// closed, once, and the VFS counts no file as open.
func TestOpenFailures(t *testing.T) {
	h := newHookVFS(t)
	name, closer, err := Register(h.name, testKey(34))
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "a.db")
	db := open(t, path, name, "_pragma=cache_size(10)")
	db.SetMaxOpenConns(1)
	mustExec(t, db, "create table t(a integer primary key, b text)")
	mustExec(t, db, fmt.Sprintf(fill, 4000)+"insert into t(b) select hex(randomblob(300)) from s")

	for _, leave := range []bool{false, true} {
		h.mu.Lock()
		h.failOpen = func(n string, flags int32) (int32, bool) {
			if strings.HasSuffix(n, "-journal") {
				return sqlite3.SQLITE_IOERR, leave
			}

			return sqlite3.SQLITE_OK, false
		}

		h.mu.Unlock()
		if _, err := db.Exec("insert into t(b) values('x')"); err == nil {
			t.Errorf("leave=%v: a write without a journal succeeded", leave)
		}
	}

	h.mu.Lock()
	h.failOpen = nil
	h.mu.Unlock()

	// The key of a temporary file cannot be made.
	entropy := randRead
	randRead = func([]byte) (int, error) { return 0, errors.New("no entropy") }
	// A sort too large for memory.
	_, err = db.Exec("create temp table x as select b from t order by b desc")
	randRead = entropy
	if err == nil {
		t.Error("a temporary file was opened without a key")
	}

	mustExec(t, db, "insert into t(b) values('y')")
	if got := queryString(t, db, "select count(*) from t"); got != "4001" {
		t.Errorf("rows: %s", got)
	}

	db.Close()
	h.mu.Lock()
	opened, closes := h.opened, h.closes
	h.mu.Unlock()
	if opened != closes || opened == 0 {
		t.Errorf("the base opened %d files and closed %d", opened, closes)
	}

	if err := closer.Close(); err != nil {
		t.Errorf("Close: %v", err)
	}
}

// A base VFS that is unregistered under an idle encrypting VFS is not
// touched again: opens fail.
func TestBaseGone(t *testing.T) {
	h := newHookVFS(t)
	name, closer, err := Register(h.name, testKey(35))
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(t.TempDir(), "a.db")
	db := open(t, path, name, "")
	mustExec(t, db, "create table t(a)")
	db.Close()

	h.unregister()
	db = open(t, path, name, "")
	if _, err := db.Exec("select * from t"); err == nil {
		t.Error("a database was opened through a VFS whose base is gone")
	}

	db.Close()
	if _, rc := rawOpen(t, name, path, rawDB); rc != sqlite3.SQLITE_CANTOPEN {
		t.Errorf("xOpen: %d", rc)
	}

	if err := closer.Close(); err != nil {
		t.Errorf("Close: %v", err)
	}
}

// Every temporary file has a key of its own: the same page at the same
// place of two of them is stored differently, and not under the key of the
// database either.
func TestTemporaryKeys(t *testing.T) {
	h := newHookVFS(t)
	enc := registerOver(t, h.name, testKey(36))
	page := bytes.Repeat([]byte("the same page!! "), 256)
	var stored [][]byte
	for _, flags := range []int32{rawTemp, rawTemp, rawDB} {
		name := ""
		if flags == rawDB {
			name = filepath.Join(t.TempDir(), "a.db")
		}

		f, rc := rawOpen(t, enc, name, flags)
		if rc != sqlite3.SQLITE_OK {
			t.Fatalf("open: %d", rc)
		}

		if rc := f.write(page, 0); rc != sqlite3.SQLITE_OK {
			t.Fatalf("write: %d", rc)
		}

		if b, rc := f.read(len(page), 0); rc != sqlite3.SQLITE_OK || !bytes.Equal(b, page) {
			t.Fatalf("read back: %d", rc)
		}

		f.close()
		h.mu.Lock()
		stored = append(stored, bytes.Clone(h.written[name]))
		delete(h.written, name)
		h.mu.Unlock()
	}

	for i, a := range stored {
		if len(a) != len(page) || bytes.Contains(a, page[:16]) {
			t.Fatalf("file %d: %d bytes stored, or plain text", i, len(a))
		}

		for j := range i {
			if bytes.Equal(a, stored[j]) {
				t.Errorf("files %d and %d are stored alike", j, i)
			}
		}
	}
}

// The VFS opens the journal of a database itself, to look at its prefix.
// The name it gives the VFS underneath is laid out as the names SQLite
// gives: SQLite's functions for such names find the database from it, and
// its query parameters, none.
func TestJournalNameForTheBase(t *testing.T) {
	h := newHookVFS(t)
	enc := registerOver(t, h.name, testKey(60))
	src := hotJournal(t, enc)
	path := filepath.Join(src, "a.db")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	copy(b, "torn by a crash, torn by a crash") // so that the journal is looked at
	if err := os.WriteFile(path, b, 0o600); err != nil {
		t.Fatal(err)
	}

	tls := libc.NewTLS()
	defer tls.Close()

	param := cstring(tls, "nolock")
	defer sqlite3.Xsqlite3_free(tls, param)

	var seen int
	h.mu.Lock()
	h.onName = func(zName uintptr, flags int32) {
		name := libc.GoString(zName)
		if !strings.HasSuffix(name, "-journal") || flags&sqlite3.SQLITE_OPEN_SUPER_JOURNAL == 0 {
			return
		}

		// The open of the key check: as SQLite looks into a journal that
		// is not the one of the connection, and under the name SQLite has
		// for it, from which the database and its file are found.
		seen++
		if db := libc.GoString(sqlite3.Xsqlite3_filename_database(tls, zName)); db+"-journal" != name {
			t.Errorf("the database of %q is found as %q", name, db)
		}

		if sqlite3.Xsqlite3_database_file_object(tls, zName) == 0 {
			t.Errorf("%q has no database file", name)
		}

		if p := sqlite3.Xsqlite3_uri_parameter(tls, zName, param); p != 0 {
			t.Errorf("%q has a query parameter %q", name, libc.GoString(p))
		}
	}
	h.mu.Unlock()
	db := open(t, path, enc, "")
	if got := queryString(t, db, "select count(*) from t"); got != "2000" {
		t.Fatalf("rows: %s", got)
	}

	h.mu.Lock()
	h.onName = nil
	h.mu.Unlock()
	if seen == 0 {
		t.Fatal("the key check did not open the journal")
	}
}

// Close while a file is being opened through the VFS waits for the open,
// and then finds the file open.
func TestCloseDuringOpen(t *testing.T) {
	h := newHookVFS(t)
	name, closer, err := Register(h.name, testKey(61))
	if err != nil {
		t.Fatal(err)
	}

	entered, release := make(chan struct{}), make(chan struct{})
	h.mu.Lock()
	h.failOpen = func(n string, flags int32) (int32, bool) {
		if strings.HasSuffix(n, "a.db") {
			close(entered)
			<-release
		}

		return sqlite3.SQLITE_OK, false
	}
	h.mu.Unlock()

	path := filepath.Join(t.TempDir(), "a.db")
	opened := make(chan *rawFile, 1)
	go func() {
		f, _ := rawOpen(t, name, path, rawDB)
		opened <- f
	}()

	<-entered
	h.mu.Lock()
	h.failOpen = nil
	h.mu.Unlock()
	closed := make(chan error, 1)
	go func() { closed <- closer.Close() }()

	select {
	case err := <-closed:
		t.Fatalf("Close returned while a file was being opened: %v", err)
	case <-time.After(100 * time.Millisecond):
	}

	close(release)
	f := <-opened
	if f == nil {
		t.Fatal("the open that was under way failed")
	}

	if err := <-closed; !errors.Is(err, ErrInUse) {
		t.Fatalf("Close after the open: %v", err)
	}

	f.close()
	if err := closer.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
}

// A call SQLite had begun when the VFS was closed still comes in, with the
// VFS it found before. It finds the VFS closed.
func TestCallAfterClose(t *testing.T) {
	name, closer, err := Register("", testKey(62))
	if err != nil {
		t.Fatal(err)
	}

	tls := libc.NewTLS()
	defer tls.Close()

	cname := cstring(tls, name)
	defer sqlite3.Xsqlite3_free(tls, cname)

	pVfs := sqlite3.Xsqlite3_vfs_find(tls, cname)
	if pVfs == 0 {
		t.Fatal("the VFS is not registered")
	}

	if err := closer.Close(); err != nil {
		t.Fatal(err)
	}

	if sqlite3.Xsqlite3_vfs_find(tls, cname) != 0 {
		t.Fatal("the VFS is still registered")
	}

	path := filepath.Join(t.TempDir(), "a.db")
	block := alloc(tls, len(path)+8)
	defer sqlite3.Xsqlite3_free(tls, block)

	copy(unsafe.Slice((*byte)(unsafe.Pointer(block+4)), len(path)), path)
	file := alloc(tls, int(vfsOf(pVfs).FszOsFile))
	defer sqlite3.Xsqlite3_free(tls, file)

	if rc := fn[func(*libc.TLS, uintptr, uintptr, uintptr, int32, uintptr) int32](vfsOf(pVfs).FxOpen)(tls, pVfs, block+4, file, rawDB, 0); rc != sqlite3.SQLITE_CANTOPEN {
		t.Errorf("xOpen after Close: %d", rc)
	}

	res := alloc(tls, 8)
	defer sqlite3.Xsqlite3_free(tls, res)

	if rc := fn[func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32](vfsOf(pVfs).FxAccess)(tls, pVfs, block+4, 0, res); rc == sqlite3.SQLITE_OK {
		t.Error("xAccess after Close succeeded")
	}

	if _, err := os.Stat(path); err == nil {
		t.Error("a file was made through the closed VFS")
	}
}

// The number an open file is known by is never one in use, also after the
// counter has gone round, which it does on a 32-bit target after 2^32
// files, and never zero.
func TestHandlesAreNotReused(t *testing.T) {
	saved := lastObj.Load()
	defer lastObj.Store(saved)

	first := addObject("first")
	defer objects.Delete(first)

	lastObj.Store(first - 1)
	second := addObject("second")
	defer objects.Delete(second)

	if second == first {
		t.Fatalf("the handle %d was given out twice", first)
	}

	if o, _ := objects.Load(first); o != "first" {
		t.Fatalf("the object at %d is now %v", first, o)
	}

	lastObj.Store(^uintptr(0))
	third := addObject("third")
	defer objects.Delete(third)

	if third == 0 || third == first || third == second {
		t.Fatalf("after the counter went round: handle %d", third)
	}
}

// PRAGMA journal_mode=WAL without the name of a database goes to every
// database of the connection, and SQLite asks only the main one's VFS
// about it. An encrypted database attached to a plain main one stays out of
// WAL mode all the same: the header that says WAL is not written.
func TestWALRefusedAttached(t *testing.T) {
	enc := register(t, testKey(63))
	for _, exclusive := range []bool{false, true} {
		t.Run(fmt.Sprintf("exclusive=%v", exclusive), func(t *testing.T) {
			dir := t.TempDir()
			pb := filepath.Join(dir, "b.db")
			db := open(t, filepath.Join(dir, "a.db"), "", "")
			db.SetMaxOpenConns(1)
			mustExec(t, db, fmt.Sprintf("attach database 'file:%s?vfs=%s' as b", pb, enc))
			mustExec(t, db, "create table b.t(a); insert into b.t values(1)")
			if exclusive {
				mustExec(t, db, "pragma locking_mode=exclusive")
				mustExec(t, db, "select * from b.t")
			}

			_, err := db.Exec("pragma journal_mode=wal")
			if exclusive && err == nil {
				t.Error("under exclusive locking the pragma succeeded")
			}

			db.Close()
			if _, err := os.Stat(pb + "-wal"); err == nil {
				t.Error("there is a WAL file")
			}

			db = open(t, pb, enc, "")
			if got := queryString(t, db, "select count(*) || ' ' || (select * from pragma_journal_mode) from t"); got != "1 delete" {
				t.Errorf("the attached database afterwards: %s", got)
			}

			if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
				t.Errorf("integrity_check: %s", got)
			}
		})
	}
}
