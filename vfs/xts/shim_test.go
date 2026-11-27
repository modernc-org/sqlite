// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func testKey(b byte) []byte { return bytes.Repeat([]byte{b}, KeySize) }

// register makes a VFS over SQLite's default one for the time of the test.
func register(t testing.TB, key []byte) string {
	t.Helper()
	name, closer, err := Register("", key)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := closer.Close(); err != nil {
			t.Errorf("closing the VFS: %v", err)
		}
	})
	return name
}

func open(t testing.TB, path, vfs, query string) *sql.DB {
	t.Helper()
	dsn := "file:" + path + "?_pragma=busy_timeout(10000)"
	if vfs != "" {
		dsn += "&vfs=" + vfs
	}

	if query != "" {
		dsn += "&" + query
	}

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { db.Close() })
	return db
}

func mustExec(t testing.TB, db *sql.DB, q string, args ...any) {
	t.Helper()
	if _, err := db.Exec(q, args...); err != nil {
		t.Fatalf("%s: %v", q, err)
	}
}

func queryString(t testing.TB, db *sql.DB, q string, args ...any) string {
	t.Helper()
	var s string
	if err := db.QueryRow(q, args...).Scan(&s); err != nil {
		t.Fatalf("%s: %v", q, err)
	}

	return s
}

// codeOf is the primary result code of an error of the driver, or -1.
func codeOf(err error) int {
	var e *sqlite.Error
	if errors.As(err, &e) {
		return e.Code() & 0xff
	}

	return -1
}

// marker is written through SQL and must be in no file as it is.
const marker = "PLAINTEXT-MARKER-0123456789-abcdefghijklmnopqrstuvwxyz"

// filesOf returns every file of dir with the hash of its content.
func filesOf(t testing.TB, dir string) map[string][32]byte {
	t.Helper()
	ents, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	m := map[string][32]byte{}
	for _, e := range ents {
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatal(err)
		}

		m[e.Name()] = sha256.Sum256(b)
	}

	return m
}

func sameFiles(t testing.TB, what string, before, after map[string][32]byte) {
	t.Helper()
	for name, h := range before {
		switch got, ok := after[name]; {
		case !ok:
			t.Errorf("%s: %s is gone", what, name)
		case got != h:
			t.Errorf("%s: %s was changed", what, name)
		}
	}

	for name := range after {
		if _, ok := before[name]; !ok {
			t.Errorf("%s: %s appeared", what, name)
		}
	}
}

func noMarker(t testing.TB, dir string) {
	t.Helper()
	ents, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range ents {
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatal(err)
		}

		if bytes.Contains(b, []byte(marker)) || bytes.Contains(b, []byte("SQLite format 3")) {
			t.Errorf("%s holds plain text", e.Name())
		}
	}
}

// A database is written, closed and read again, by a pool; nothing of it is
// on the disk as SQLite wrote it.
func TestRoundTrip(t *testing.T) {
	vfs := register(t, testKey(1))
	dir := t.TempDir()
	path := filepath.Join(dir, "a.db")
	db := open(t, path, vfs, "")
	db.SetMaxOpenConns(4)
	mustExec(t, db, "create table t(a integer primary key, b text)")
	var wg sync.WaitGroup
	for g := range 4 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := range 100 {
				if _, err := db.Exec("insert into t(b) values(?)", fmt.Sprintf("%s %d %d", marker, g, i)); err != nil {
					t.Error(err)
					return
				}

				var n int
				if err := db.QueryRow("select count(*) from t").Scan(&n); err != nil {
					t.Error(err)
					return
				}
			}
		}()
	}

	wg.Wait()
	noMarker(t, dir)
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	noMarker(t, dir)
	db = open(t, path, vfs, "")
	if got := queryString(t, db, "select count(*) || ' ' || min(b) from t"); got != "400 "+marker+" 0 0" {
		t.Fatalf("read back %q", got)
	}

	if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
		t.Fatalf("integrity_check: %s", got)
	}
}

func TestJournalModes(t *testing.T) {
	// Over a VFS that keeps everything written, also to journals SQLite
	// deletes before the test could look at them.
	h := newHookVFS(t)
	vfs := registerOver(t, h.name, testKey(2))
	for _, mode := range []string{"delete", "truncate", "persist", "memory", "off"} {
		for _, pageSize := range []int{512, 4096, 65536} {
			t.Run(fmt.Sprintf("%s/%d", mode, pageSize), func(t *testing.T) {
				dir := t.TempDir()
				path := filepath.Join(dir, "a.db")
				q := fmt.Sprintf("_pragma=page_size(%d)&_pragma=journal_mode(%s)&_pragma=cache_size(10)", pageSize, mode)
				db := open(t, path, vfs, q)
				db.SetMaxOpenConns(1)
				if got := queryString(t, db, "pragma journal_mode"); got != mode {
					t.Fatalf("journal_mode is %s", got)
				}

				mustExec(t, db, "create table t(a integer primary key, b text)")
				for range 3 {
					mustExec(t, db, "with recursive s(value) as (select 1 union all select value+1 from s where value < 300) insert into t(b) select ? || value from s", marker)
				}

				if mode != "off" {
					// A transaction too large for the cache, rolled back:
					// the pages come back from the journal.
					tx, err := db.Begin()
					if err != nil {
						t.Fatal(err)
					}

					if _, err := tx.Exec("update t set b = b || 'x'"); err != nil {
						t.Fatal(err)
					}

					if _, err := tx.Exec("delete from t where a % 3 = 0"); err != nil {
						t.Fatal(err)
					}

					noMarker(t, dir)
					if err := tx.Rollback(); err != nil {
						t.Fatal(err)
					}
				}

				if got := queryString(t, db, "select count(*) || ' ' || max(length(b)) from t"); got != fmt.Sprintf("900 %d", len(marker)+3) {
					t.Fatalf("after the rollback: %s", got)
				}

				mustExec(t, db, "vacuum")
				if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
					t.Fatalf("integrity_check: %s", got)
				}

				if got := queryString(t, db, "pragma page_size"); got != fmt.Sprint(pageSize) {
					t.Fatalf("page_size is %s", got)
				}

				noMarker(t, dir)
				h.mu.Lock()
				defer h.mu.Unlock()

				journals := 0
				for name, b := range h.written {
					if strings.HasSuffix(name, "-journal") {
						journals++
					}

					if bytes.Contains(b, []byte(marker)) || bytes.Contains(b, []byte("SQLite format 3")) {
						t.Errorf("%s was written with plain text", filepath.Base(name))
					}
				}

				if journals == 0 && mode != "memory" && mode != "off" {
					t.Error("no journal was written")
				}

				clear(h.written)
			})
		}
	}
}

// No write-ahead log: the pragma is answered with an error and the database
// goes on as it was.
func TestWALRefused(t *testing.T) {
	vfs := register(t, testKey(3))
	for _, q := range []string{"", "_pragma=locking_mode(exclusive)"} {
		for _, stmt := range []string{
			"pragma journal_mode=wal", "pragma journal_mode=WAL", "pragma journal_mode='wa'",
			"pragma journal_mode=w", "pragma main.journal_mode=wal",
		} {
			t.Run(q+"/"+stmt, func(t *testing.T) {
				dir := t.TempDir()
				db := open(t, filepath.Join(dir, "a.db"), vfs, q)
				db.SetMaxOpenConns(1)
				mustExec(t, db, "create table t(a); insert into t values(1)")
				var mode string
				err := db.QueryRow(stmt).Scan(&mode)
				if err == nil || !strings.Contains(err.Error(), "no write-ahead log") {
					t.Fatalf("%s: mode %q, error %v", stmt, mode, err)
				}

				mustExec(t, db, "insert into t values(2)")
				if got := queryString(t, db, "select count(*) || ' ' || (select * from pragma_journal_mode) from t"); got != "2 delete" {
					t.Fatalf("after the refusal: %s", got)
				}

				if _, err := os.Stat(filepath.Join(dir, "a.db-wal")); err == nil {
					t.Fatal("there is a WAL file")
				}
			})
		}
	}
}

// hotJournal leaves in a new directory a database of 2000 rows and the
// journal of a transaction that was changing it when its writer died: the
// files are copied while the transaction is open, after it has spilled.
func hotJournal(t *testing.T, vfs string) (dir string) {
	t.Helper()
	src := t.TempDir()
	db := open(t, filepath.Join(src, "a.db"), vfs, "_pragma=cache_size(10)")
	db.SetMaxOpenConns(1)
	mustExec(t, db, "create table t(a integer primary key, b text)")
	mustExec(t, db, "with recursive s(value) as (select 1 union all select value+1 from s where value < 2000) insert into t(b) select ? || value from s", marker)
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	// The first statement changes page 1, so its old image is in the journal.
	if _, err := tx.Exec("create table u(x)"); err != nil {
		t.Fatal(err)
	}

	if _, err := tx.Exec("update t set b = 'changed ' || b"); err != nil {
		t.Fatal(err)
	}

	dir = t.TempDir()
	for _, name := range []string{"a.db", "a.db-journal"} {
		b, err := os.ReadFile(filepath.Join(src, name))
		if err != nil {
			t.Fatal(err)
		}

		if err := os.WriteFile(filepath.Join(dir, name), b, 0o600); err != nil {
			t.Fatal(err)
		}
	}

	return dir
}

// The journal of a writer that died is rolled back by the next connection
// that has the key.
func TestHotJournal(t *testing.T) {
	vfs := register(t, testKey(4))
	dir := hotJournal(t, vfs)
	noMarker(t, dir)
	db := open(t, filepath.Join(dir, "a.db"), vfs, "")
	if got := queryString(t, db, "select count(*) || ' ' || count(*) filter (where b like 'changed%') from t"); got != "2000 0" {
		t.Fatalf("after the recovery: %s", got)
	}

	if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
		t.Fatalf("integrity_check: %s", got)
	}

	if _, err := os.Stat(filepath.Join(dir, "a.db-journal")); err == nil {
		t.Fatal("the journal is still there")
	}
}

// Another key, no key, and a plain database through the VFS: the statement
// fails with SQLITE_NOTADB and no file is changed, also when there is a hot
// journal, which SQLite on its own would take for one without a header and
// delete.
func TestWrongKey(t *testing.T) {
	right, wrong := register(t, testKey(5)), register(t, testKey(6))
	for _, hot := range []bool{false, true} {
		for _, vfs := range []string{wrong, ""} {
			t.Run(fmt.Sprintf("hot=%v/vfs=%q", hot, vfs), func(t *testing.T) {
				var dir string
				if hot {
					dir = hotJournal(t, right)
				} else {
					dir = t.TempDir()
					db := open(t, filepath.Join(dir, "a.db"), right, "")
					mustExec(t, db, "create table t(a, b); insert into t values(1, ?)", marker)
					db.Close()
				}

				before := filesOf(t, dir)
				db := open(t, filepath.Join(dir, "a.db"), vfs, "")
				db.SetMaxOpenConns(1)
				for _, q := range []string{
					"select count(*) from t", "create table u(a)", "pragma journal_mode=persist",
					"pragma journal_mode=delete", "vacuum",
				} {
					if _, err := db.Exec(q); codeOf(err) != sqlite3.SQLITE_NOTADB {
						t.Errorf("%s: %v", q, err)
					}
				}

				db.Close()
				sameFiles(t, "after the wrong key", before, filesOf(t, dir))

				// And the right key still finds everything.
				db = open(t, filepath.Join(dir, "a.db"), right, "")
				if got := queryString(t, db, "select count(*) from t where b like 'changed%'"); got != "0" {
					t.Fatalf("with the right key: %s changed rows", got)
				}

				if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
					t.Fatalf("integrity_check: %s", got)
				}
			})
		}
	}

	t.Run("plain database through the VFS", func(t *testing.T) {
		dir := t.TempDir()
		db := open(t, filepath.Join(dir, "a.db"), "", "")
		mustExec(t, db, "create table t(a); insert into t values(1)")
		db.Close()
		before := filesOf(t, dir)
		db = open(t, filepath.Join(dir, "a.db"), right, "")
		if _, err := db.Exec("select count(*) from t"); codeOf(err) != sqlite3.SQLITE_NOTADB {
			t.Errorf("select: %v", err)
		}

		db.Close()
		sameFiles(t, "after the VFS", before, filesOf(t, dir))
	})
}

func TestRegisterErrors(t *testing.T) {
	if _, _, err := Register("", make([]byte, 16)); err == nil {
		t.Error("a key of 16 bytes was taken")
	}

	if _, _, err := Register("no-such-vfs", testKey(1)); err == nil {
		t.Error("a base that does not exist was taken")
	}

	name, closer, err := Register("", testKey(1))
	if err != nil {
		t.Fatal(err)
	}

	db := open(t, filepath.Join(t.TempDir(), "a.db"), name, "")
	mustExec(t, db, "create table t(a)")
	if err := closer.Close(); !errors.Is(err, ErrInUse) {
		t.Errorf("Close with a database open: %v", err)
	}

	db.Close()
	if err := closer.Close(); err != nil {
		t.Errorf("Close: %v", err)
	}

	if err := closer.Close(); err != nil {
		t.Errorf("second Close: %v", err)
	}

	db = open(t, filepath.Join(t.TempDir(), "b.db"), name, "")
	if _, err := db.Exec("create table t(a)"); err == nil {
		t.Error("a database was opened through a closed VFS")
	}
}

// A crash while page 1 was being written leaves a database whose first bytes
// no key reads. The journal restores the page, and the prefix of the journal
// is what tells the right key from a wrong one.
func TestTornPageOne(t *testing.T) {
	right, wrong := register(t, testKey(7)), register(t, testKey(8))
	dir := hotJournal(t, right)
	path := filepath.Join(dir, "a.db")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	copy(b, "torn by a crash, torn by a crash")
	if err := os.WriteFile(path, b, 0o600); err != nil {
		t.Fatal(err)
	}

	before := filesOf(t, dir)
	for _, vfs := range []string{wrong, ""} {
		db := open(t, path, vfs, "")
		if _, err := db.Exec("select count(*) from t"); codeOf(err) != sqlite3.SQLITE_NOTADB {
			t.Errorf("vfs %q: %v", vfs, err)
		}

		db.Close()
		sameFiles(t, fmt.Sprintf("after vfs %q", vfs), before, filesOf(t, dir))
	}

	db := open(t, path, right, "")
	if got := queryString(t, db, "select count(*) || ' ' || count(*) filter (where b like 'changed%') from t"); got != "2000 0" {
		t.Fatalf("after the recovery: %s", got)
	}

	if got := queryString(t, db, "select count(*) from sqlite_schema where name = 'u'"); got != "0" {
		t.Fatal("the table of the transaction that died is there")
	}

	if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
		t.Fatalf("integrity_check: %s", got)
	}
}

// The methods of a file as SQLite calls them: a read past the end is a
// short read with zeros and what there is before them; what the file says
// of itself.
func TestFileMethods(t *testing.T) {
	vfs := register(t, testKey(9))
	dir := t.TempDir()
	path := filepath.Join(dir, "a.db")
	f, rc := rawOpen(t, vfs, path, rawDB)
	if rc != sqlite3.SQLITE_OK {
		t.Fatalf("open: %d", rc)
	}

	defer f.close()

	if f.m().FiVersion != 1 || f.m().FxShmMap != 0 || f.m().FxFetch != 0 {
		t.Fatalf("the methods are version %d, with shared memory or mapping", f.m().FiVersion)
	}

	if dc := f.characteristics(); dc&^keptCharacteristics != 0 {
		t.Fatalf("device characteristics %#x", dc)
	}

	if s := f.sectorSize(); s < 512 || s&(s-1) != 0 {
		t.Fatalf("sector size %d", s)
	}

	if b, rc := f.read(100, 0); rc != sqlite3.SQLITE_IOERR_SHORT_READ || !bytes.Equal(b, make([]byte, 100)) {
		t.Fatalf("read of an empty file: %d, % x", rc, b[:8])
	}

	// It begins as a database does: the size of a file that does not is
	// refused, see TestWrongKey.
	page := append([]byte("SQLite format 3\x00"), bytes.Repeat([]byte(marker), 20)...)[:1024]
	if rc := f.write(page, 0); rc != sqlite3.SQLITE_OK {
		t.Fatalf("write: %d", rc)
	}

	if b, rc := f.read(1024, 0); rc != sqlite3.SQLITE_OK || !bytes.Equal(b, page) {
		t.Fatalf("read: %d", rc)
	}

	// Across the end: what there is, then zeros, and the code that says so.
	b, rc := f.read(4096, 512)
	if want := append(bytes.Clone(page[512:]), make([]byte, 4096-512)...); rc != sqlite3.SQLITE_IOERR_SHORT_READ || !bytes.Equal(b, want) {
		t.Fatalf("read across the end: %d, begins % x", rc, b[:16])
	}

	if n, rc := f.size(); rc != sqlite3.SQLITE_OK || n != 1024 {
		t.Fatalf("size %d, %d", n, rc)
	}

	stored, err := os.ReadFile(path)
	if err != nil || len(stored) != 1024 || bytes.Contains(stored, []byte(marker[:16])) {
		t.Fatalf("stored: %d bytes, %v", len(stored), err)
	}

	// A header that says WAL is not stored.
	wal := bytes.Clone(page[:512])
	wal[18], wal[19] = 2, 2
	if rc := f.write(wal, 0); rc != sqlite3.SQLITE_IOERR_WRITE {
		t.Fatalf("write of a WAL header: %d", rc)
	}

	if b, rc := f.read(512, 0); rc != sqlite3.SQLITE_OK || !bytes.Equal(b, page[:512]) {
		t.Fatalf("after the refused write: %d", rc)
	}

	// The name of the VFS, as SQLite's shims give it: theirs before the
	// name of what is under them.
	out := alloc(f.tls, 8)
	defer sqlite3.Xsqlite3_free(f.tls, out)

	if rc := f.fileControl(sqlite3.SQLITE_FCNTL_VFSNAME, out); rc != sqlite3.SQLITE_OK {
		t.Fatalf("VFSNAME: %d", rc)
	}

	zName := *(*uintptr)(unsafe.Pointer(out))
	if got := libc.GoString(zName); got != vfs+"/"+defaultVFS(t) {
		t.Errorf("VFSNAME is %q", got)
	}

	sqlite3.Xsqlite3_free(f.tls, zName)

	// The kinds of file the VFS does not open.
	for what, flags := range map[string]int32{
		"a WAL":                 rawWAL,
		"a file without a kind": rawRW,
	} {
		if _, rc := rawOpen(t, vfs, filepath.Join(dir, "x"), flags); rc != sqlite3.SQLITE_CANTOPEN {
			t.Errorf("%s: open: %d", what, rc)
		}
	}

	if _, err := os.Stat(filepath.Join(dir, "x")); err == nil {
		t.Error("a refused open left a file")
	}
}

// A connection that saw the file while it was empty, under another key,
// and kept its journal mode: when the file has become a database with a hot
// journal, a change of the journal mode on that connection, which SQLite
// answers by deleting the journal, deletes nothing.
func TestStaleHandleOtherKey(t *testing.T) {
	right, wrong := register(t, testKey(10)), register(t, testKey(11))
	for _, torn := range []bool{false, true} {
		t.Run(fmt.Sprintf("torn=%v", torn), func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "a.db")
			stale := open(t, path, wrong, "")
			stale.SetMaxOpenConns(1)
			if got := queryString(t, stale, "pragma journal_mode=persist"); got != "persist" {
				t.Fatalf("journal_mode: %s", got)
			}

			if got := queryString(t, stale, "select count(*) from sqlite_schema"); got != "0" {
				t.Fatalf("the empty file has %s tables", got)
			}

			src := hotJournal(t, right)
			for _, name := range []string{"a.db", "a.db-journal"} {
				b, err := os.ReadFile(filepath.Join(src, name))
				if err != nil {
					t.Fatal(err)
				}

				if torn && name == "a.db" {
					copy(b, "torn by a crash, torn by a crash")
				}

				if err := os.WriteFile(filepath.Join(dir, name), b, 0o600); err != nil {
					t.Fatal(err)
				}
			}

			before := filesOf(t, dir)
			// SQLite does not pass the failed lock on here; what counts is
			// that the journal stays.
			stale.Exec("pragma journal_mode=delete")
			stale.Exec("pragma journal_mode=truncate")
			sameFiles(t, "after the change of journal mode", before, filesOf(t, dir))
			if _, err := stale.Exec("create table x(a)"); codeOf(err) != sqlite3.SQLITE_NOTADB {
				t.Errorf("a write through the stale connection: %v", err)
			}

			sameFiles(t, "after the write", before, filesOf(t, dir))

			// The stale connection holds no lock: the owner recovers and
			// writes while it is open.
			db := open(t, path, right, "")
			mustExec(t, db, "insert into t(b) values('after')")
			if got := queryString(t, db, "select count(*) || ' ' || count(*) filter (where b like 'changed%') from t"); got != "2001 0" {
				t.Fatalf("after the recovery: %s", got)
			}

			if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
				t.Fatalf("integrity_check: %s", got)
			}
		})
	}
}

// A lock that the key check refuses is not held: SQLite is told the lock
// failed, and the next connection gets every lock it asks for.
func TestLockRefused(t *testing.T) {
	right, wrong := register(t, testKey(12)), register(t, testKey(13))
	path := filepath.Join(t.TempDir(), "a.db")
	db := open(t, path, right, "")
	mustExec(t, db, "create table t(a); insert into t values(1)")
	db.Close()

	other, rc := rawOpen(t, wrong, path, rawDB)
	if rc != sqlite3.SQLITE_OK {
		t.Fatalf("open: %d", rc)
	}

	defer other.close()

	if rc := other.lock(sqlite3.SQLITE_LOCK_SHARED); rc != sqlite3.SQLITE_NOTADB {
		t.Fatalf("lock with the wrong key: %d", rc)
	}

	owner, rc := rawOpen(t, right, path, rawDB)
	if rc != sqlite3.SQLITE_OK {
		t.Fatalf("open: %d", rc)
	}

	defer owner.close()

	for _, level := range []int32{sqlite3.SQLITE_LOCK_SHARED, sqlite3.SQLITE_LOCK_RESERVED, sqlite3.SQLITE_LOCK_EXCLUSIVE} {
		if rc := owner.lock(level); rc != sqlite3.SQLITE_OK {
			t.Fatalf("lock %d with the right key: %d", level, rc)
		}
	}

	if rc := owner.unlock(0); rc != sqlite3.SQLITE_OK {
		t.Fatalf("unlock: %d", rc)
	}
}
