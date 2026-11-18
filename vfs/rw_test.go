// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vfs

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"math"
	"math/rand"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// memStore is the storage behind the VFS the tests register. Next to what a
// running process sees of each file it keeps what would survive a loss of
// power, which is what the last Sync of the file made durable, so that a test
// can reopen the storage as it would be after a crash.
type memStore struct {
	mu      sync.Mutex
	files   map[string]*memData // the files as the process sees them
	durable map[string][]byte   // the files as the medium holds them

	live  int      // files open through Open and not closed yet
	ops   int      // writes, truncates, syncs, creations and deletions so far
	trace []string // one line for each of them

	// fail, if set, is asked before each counted operation and its error is
	// returned in place of doing it.
	fail func(n int, op, name string) error
	// after, if set, runs after each counted operation, with mu held.
	after func(n int)

	canon    func(name string) string      // FullPathname
	outFlags func(flags OpenFlag) OpenFlag // what Open reports
	opened   []string                      // every name passed to Open
	readOnly bool                          // Open refuses to write
}

type memData struct {
	b       []byte
	base    []byte     // b at the last Sync of this file, nil if never synced
	pending []memWrite // what was done to the file since
}

// memWrite is a write, or with data == nil a truncation to off.
type memWrite struct {
	off  int64
	data []byte
}

func newMemStore() *memStore {
	return &memStore{files: map[string]*memData{}, durable: map[string][]byte{}}
}

// reopened returns a new storage holding the given files, all of them durable.
func reopened(image map[string][]byte) *memStore {
	s := newMemStore()
	for name, b := range image {
		b = append([]byte(nil), b...)
		s.files[name] = &memData{b: b, base: append([]byte(nil), b...)}
		s.durable[name] = append([]byte(nil), b...)
	}
	return s
}

// count records a counted operation and returns the error to inject, if any.
// The caller holds mu.
func (s *memStore) count(op, name string, format string, a ...any) error {
	s.ops++
	s.trace = append(s.trace, fmt.Sprintf("%s %s", op, kindOf(name))+fmt.Sprintf(format, a...))
	if s.fail != nil {
		if err := s.fail(s.ops, op, name); err != nil {
			return err
		}
	}
	return nil
}

func (s *memStore) done() {
	if s.after != nil {
		s.after(s.ops)
	}
}

// kindOf names a file by its role, so that traces do not depend on names.
func kindOf(name string) string {
	switch {
	case strings.HasSuffix(name, "-journal"):
		return "journal"
	case strings.HasSuffix(name, "-wal"):
		return "wal"
	case strings.Contains(name, "-mj"):
		return "super"
	}
	return "db"
}

func (s *memStore) Open(name string, flags OpenFlag) (File, OpenFlag, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.opened = append(s.opened, name)
	d := s.files[name]
	switch {
	case d == nil && flags&OpenCreate == 0:
		return nil, 0, fs.ErrNotExist
	case d != nil && flags&OpenCreate != 0 && flags&OpenExclusive != 0:
		return nil, 0, fs.ErrExist
	case d == nil && s.readOnly:
		return nil, 0, fs.ErrPermission
	case d == nil:
		if err := s.count("create", name, ""); err != nil {
			return nil, 0, err
		}

		d = &memData{}
		s.files[name] = d
		s.done()
	}

	out := flags
	if s.readOnly {
		out = out&^OpenReadWrite | OpenReadOnly
	}
	if s.outFlags != nil {
		out = s.outFlags(out)
	}
	s.live++
	return &memFile{s: s, name: name, d: d}, out, nil
}

func (s *memStore) Delete(name string, syncDir bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.count("delete", name, " syncDir=%v", syncDir); err != nil {
		return err
	}

	if s.files[name] == nil {
		return fs.ErrNotExist
	}

	delete(s.files, name)
	if syncDir {
		delete(s.durable, name)
	}
	s.done()
	return nil
}

func (s *memStore) Access(name string, flags AccessFlag) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.files[name] != nil && !(flags == AccessReadWrite && s.readOnly), nil
}

func (s *memStore) FullPathname(name string) (string, error) {
	if s.canon != nil {
		return s.canon(name), nil
	}

	return name, nil
}

// content returns a copy of the named file as the process sees it.
func (s *memStore) content(name string) []byte {
	s.mu.Lock()
	defer s.mu.Unlock()

	if d := s.files[name]; d != nil {
		return append([]byte(nil), d.b...)
	}

	return nil
}

func (s *memStore) names() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var r []string
	for name := range s.files {
		r = append(r, name)
	}
	sort.Strings(r)
	return r
}

// What a loss of power may leave of the storage.
const (
	crashSynced  = iota // only what was synced
	crashWritten        // everything written, synced or not
	crashTorn           // of what was not synced, some sectors
	crashSector         // of what was not synced, some writes, others spoiling their sectors
)

// image returns the files a loss of power at this moment may leave. The caller
// holds mu. The files are gone through by name, so that one seed gives one
// image.
func (s *memStore) image(how int, rng *rand.Rand) map[string][]byte {
	const sector = 512
	img := map[string][]byte{}
	files, durable := slices.Sorted(maps.Keys(s.files)), slices.Sorted(maps.Keys(s.durable))
	switch how {
	case crashSynced:
		for name, b := range s.durable {
			img[name] = append([]byte(nil), b...)
		}
	case crashWritten:
		for name, d := range s.files {
			img[name] = append([]byte(nil), d.b...)
		}
	case crashTorn:
		for _, name := range durable {
			b := s.durable[name]
			// A file deleted without a directory sync, or not yet synced
			// after it was created again, may be the old one still.
			if s.files[name] == nil || s.files[name].base == nil {
				if rng.Intn(2) == 0 {
					img[name] = append([]byte(nil), b...)
				}
			}
		}
		for _, name := range files {
			d := s.files[name]
			if d.base == nil {
				if _, ok := img[name]; ok || rng.Intn(2) == 0 {
					continue // the old file, or no file at all
				}
			}

			b := append([]byte(nil), d.base...)
			for _, w := range d.pending {
				if w.data == nil {
					if rng.Intn(2) == 0 {
						b = resize(b, w.off)
					}
					continue
				}

				for i := int64(0); i < int64(len(w.data)); i += sector {
					if rng.Intn(2) == 0 {
						continue
					}

					part := w.data[i:min(i+sector, int64(len(w.data)))]
					if end := w.off + i + int64(len(part)); end > int64(len(b)) {
						b = resize(b, end)
					}
					copy(b[w.off+i:], part)
				}
			}
			img[name] = b
		}
	case crashSector:
		// SQLite is told a sector of 4096 bytes: a write the power cut short
		// may leave anything in the sectors it touches.
		const block = 4096
		for _, name := range files {
			d := s.files[name]
			if d.base == nil {
				if b, ok := s.durable[name]; ok {
					img[name] = append([]byte(nil), b...) // the old file
				}
				continue
			}

			b := append([]byte(nil), d.base...)
			for _, w := range d.pending {
				switch what := rng.Intn(3); {
				case w.data == nil && what == 0:
					b = resize(b, w.off)
				case w.data == nil || what == 1:
					// not done
				case what == 0:
					if end := w.off + int64(len(w.data)); end > int64(len(b)) {
						b = resize(b, end)
					}
					copy(b[w.off:], w.data)
				default:
					// The file is as long as the write would have left
					// it: the damage is to what the sectors hold.
					end := w.off + int64(len(w.data))
					if end > int64(len(b)) {
						b = resize(b, end)
					}
					from, to := w.off/block*block, min((end+block-1)/block*block, int64(len(b)))
					for i := from; i < to; i++ {
						b[i] = 0xee
					}
				}
			}
			img[name] = b
		}
		for _, name := range durable {
			if s.files[name] == nil && rng.Intn(2) == 0 {
				img[name] = append([]byte(nil), s.durable[name]...)
			}
		}
	}
	return img
}

func resize(b []byte, size int64) []byte {
	if size <= int64(len(b)) {
		return b[:size]
	}

	return append(b, make([]byte, size-int64(len(b)))...)
}

type memFile struct {
	s    *memStore
	name string
	d    *memData
}

func (f *memFile) ReadAt(p []byte, off int64) (int, error) {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	if off >= int64(len(f.d.b)) {
		return 0, io.EOF
	}

	n := copy(p, f.d.b[off:])
	if n < len(p) {
		return n, io.EOF
	}

	return n, nil
}

func (f *memFile) WriteAt(p []byte, off int64) (int, error) {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	if err := f.s.count("write", f.name, " %d+%d", off, len(p)); err != nil {
		return 0, err
	}

	if f.s.readOnly {
		return 0, fs.ErrPermission
	}

	if end := off + int64(len(p)); end > int64(len(f.d.b)) {
		f.d.b = resize(f.d.b, end)
	}
	copy(f.d.b[off:], p)
	f.d.pending = append(f.d.pending, memWrite{off, append([]byte{}, p...)})
	f.s.done()
	return len(p), nil
}

func (f *memFile) Truncate(size int64) error {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	if err := f.s.count("truncate", f.name, " %d", size); err != nil {
		return err
	}

	f.d.b = resize(f.d.b, size)
	f.d.pending = append(f.d.pending, memWrite{off: size})
	f.s.done()
	return nil
}

func (f *memFile) Sync(flags SyncFlag) error {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	if err := f.s.count("sync", f.name, " %#x", int32(flags)); err != nil {
		return err
	}

	// A file deleted while it is open is synced to nowhere.
	if f.s.files[f.name] == f.d {
		f.d.base = append([]byte{}, f.d.b...)
		f.d.pending = nil
		f.s.durable[f.name] = append([]byte(nil), f.d.b...)
	}
	f.s.done()
	return nil
}

func (f *memFile) Size() (int64, error) {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	return int64(len(f.d.b)), nil
}

func (f *memFile) Close() error {
	f.s.mu.Lock()
	f.s.live--
	f.s.mu.Unlock()
	return nil
}

// openRW registers s and opens name through it with the given pool size and
// further DSN parameters. The database and the registration are closed when
// the test ends.
func openRW(t *testing.T, s VFS, name string, conns int, params ...string) *sql.DB {
	t.Helper()
	vfsName, fsys, err := Register(s)
	if err != nil {
		t.Fatalf("Register: %v", err)
	}

	db, err := sql.Open("sqlite", "file:"+name+"?"+strings.Join(append([]string{"vfs=" + vfsName}, params...), "&"))
	if err != nil {
		fsys.Close()
		t.Fatalf("sql.Open: %v", err)
	}

	db.SetMaxOpenConns(conns)
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("closing the database: %v", err)
		}
		if err := fsys.Close(); err != nil {
			t.Errorf("closing the VFS: %v", err)
		}
	})
	return db
}

func mustExec(t *testing.T, db interface {
	Exec(string, ...any) (sql.Result, error)
}, stmts ...string) {
	t.Helper()
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			t.Fatalf("%s: %s", s, E(err))
		}
	}
}

// dump returns the result of a query as one string, a line per row.
func dump(t *testing.T, db *sql.DB, query string) string {
	t.Helper()
	rows, err := db.Query(query)
	if err != nil {
		t.Fatalf("%s: %s", query, E(err))
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		t.Fatal(err)
	}

	var sb strings.Builder
	vals := make([]any, len(cols))
	ptrs := make([]any, len(cols))
	for i := range vals {
		ptrs[i] = &vals[i]
	}
	for rows.Next() {
		if err := rows.Scan(ptrs...); err != nil {
			t.Fatalf("%s: %s", query, E(err))
		}
		for i, v := range vals {
			if b, ok := v.([]byte); ok {
				v = fmt.Sprintf("blob(%d,%x)", len(b), b[:min(len(b), 8)])
			}
			fmt.Fprintf(&sb, "%v|", v)
			_ = i
		}
		sb.WriteByte('\n')
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("%s: %s", query, E(err))
	}
	return sb.String()
}

func integrity(t *testing.T, db *sql.DB) {
	t.Helper()
	if got := dump(t, db, "PRAGMA integrity_check"); got != "ok|\n" {
		t.Fatalf("integrity_check: %s", got)
	}
}

// workload exercises the paths a VFS serves: growth and shrinking of the
// file, a blob larger than a page, a savepoint, a statement that fails half
// way and is undone from its statement journal, a temporary table, a sort
// too large for the cache, and VACUUM.
var workload = []string{
	"PRAGMA cache_size=20",
	"CREATE TABLE t(id INTEGER PRIMARY KEY, k TEXT UNIQUE, v BLOB)",
	"CREATE INDEX t_v ON t(length(v))",
	"WITH RECURSIVE n(i) AS (SELECT 1 UNION ALL SELECT i+1 FROM n WHERE i<400) INSERT INTO t(id,k,v) SELECT i, 'k'||i, randomblob(40+i%60) FROM n",
	"INSERT INTO t(id,k,v) VALUES(1000,'big',zeroblob(300000))",
	"BEGIN",
	"UPDATE t SET k='x'||k WHERE id%7=0",
	"SAVEPOINT s",
	"DELETE FROM t WHERE id%3=0",
	"ROLLBACK TO s",
	"COMMIT",
	"CREATE TEMP TABLE tt AS SELECT id, k FROM t WHERE id%2=0",
	"DELETE FROM t WHERE id IN (SELECT id FROM tt WHERE id%10=0)",
	"VACUUM",
}

var workloadQueries = []string{
	"SELECT count(*), sum(id), sum(length(v)) FROM t",
	"SELECT id, k, length(v) FROM t ORDER BY k DESC",
	"SELECT count(*) FROM tt",
	"SELECT id FROM t ORDER BY length(v), k LIMIT 50",
}

// TestRegister runs a workload through a registered VFS and through SQLite's
// own, and requires the same answers from both and a sound database behind
// each, also when the one behind the VFS is opened again.
func TestRegister(t *testing.T) {
	s := newMemStore()
	db := openRW(t, s, "test.db", 1)
	ref, err := sql.Open("sqlite", "file:"+filepath.Join(t.TempDir(), "ref.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer ref.Close()
	ref.SetMaxOpenConns(1)

	// randomblob differs between the two, so the blobs are compared by size.
	mustExec(t, db, workload...)
	mustExec(t, ref, workload...)
	// The update of a UNIQUE column to a value another row has fails after
	// it changed rows, and those changes are rolled back from the journal.
	for _, d := range []*sql.DB{db, ref} {
		if _, err := d.Exec("UPDATE t SET k = CASE WHEN id < 300 THEN k||'y' ELSE 'k299y' END"); err == nil {
			t.Fatal("the update that breaks UNIQUE succeeded")
		}
	}
	for _, q := range workloadQueries {
		if got, want := dump(t, db, q), dump(t, ref, q); got != want {
			t.Fatalf("%s through the VFS:\n%s\nthrough the default VFS:\n%s", q, got, want)
		}
	}
	integrity(t, db)
	integrity(t, ref)

	s.mu.Lock()
	for _, name := range s.opened {
		if name == "" {
			t.Error("a temporary file was opened through the VFS")
		}
	}
	s.mu.Unlock()
	if got := s.names(); len(got) != 1 || got[0] != "test.db" {
		t.Fatalf("files left in the storage: %q, want only the database", got)
	}

	// The same storage behind a second registration.
	again := openRW(t, reopened(map[string][]byte{"test.db": s.content("test.db")}), "test.db", 1)
	integrity(t, again)
	q := workloadQueries[1]
	if got, want := dump(t, again, q), dump(t, ref, q); got != want {
		t.Fatalf("%s after reopening:\n%s\nwant:\n%s", q, got, want)
	}
}

// TestRegisterJournalModes commits and rolls back in every journal mode a
// version 1 VFS supports, and reads the result from a copy of the storage
// each time. journal_mode=OFF cannot roll back, so there only the commit is
// checked.
func TestRegisterJournalModes(t *testing.T) {
	for _, mode := range []string{"delete", "truncate", "persist", "memory", "off"} {
		t.Run(mode, func(t *testing.T) {
			s := newMemStore()
			db := openRW(t, s, "test.db", 1)
			if got := dump(t, db, "PRAGMA journal_mode="+mode); got != mode+"|\n" {
				t.Fatalf("journal_mode=%s answered %q", mode, got)
			}

			mustExec(t, db,
				"CREATE TABLE t(id INTEGER PRIMARY KEY, v TEXT)",
				"INSERT INTO t VALUES(1,'one'),(2,'two')",
			)
			if mode != "off" {
				mustExec(t, db, "BEGIN", "UPDATE t SET v='lost'", "INSERT INTO t VALUES(3,'lost')", "ROLLBACK")
			}
			mustExec(t, db, "BEGIN", "UPDATE t SET v=v||'!' WHERE id=2", "COMMIT")
			const want = "1|one|\n2|two!|\n"
			if got := dump(t, db, "SELECT * FROM t ORDER BY id"); got != want {
				t.Fatalf("got %q, want %q", got, want)
			}

			image := map[string][]byte{}
			for _, name := range s.names() {
				image[name] = s.content(name)
			}
			again := openRW(t, reopened(image), "test.db", 1)
			if got := dump(t, again, "SELECT * FROM t ORDER BY id"); got != want {
				t.Fatalf("after reopening: got %q, want %q", got, want)
			}
			integrity(t, again)
		})
	}
}

// TestRegisterWAL pins what a VFS without shared memory does with the
// write-ahead log: nothing under normal locking, the real thing under
// exclusive locking, and a database left in WAL mode that only a connection
// in exclusive locking mode can open.
func TestRegisterWAL(t *testing.T) {
	s := newMemStore()
	db := openRW(t, s, "test.db", 1)
	mustExec(t, db, "CREATE TABLE t(v)")
	if got := dump(t, db, "PRAGMA journal_mode=WAL"); got != "delete|\n" {
		t.Fatalf("journal_mode=WAL under normal locking answered %q, want the mode unchanged", got)
	}

	if got := dump(t, db, "PRAGMA locking_mode=EXCLUSIVE"); got != "exclusive|\n" {
		t.Fatalf("locking_mode=EXCLUSIVE answered %q", got)
	}

	if got := dump(t, db, "PRAGMA journal_mode=WAL"); got != "wal|\n" {
		t.Fatalf("journal_mode=WAL under exclusive locking answered %q", got)
	}

	mustExec(t, db, "INSERT INTO t VALUES('in the log')")
	if s.content("test.db-wal") == nil {
		t.Fatal("no write-ahead log in the storage")
	}

	image := map[string][]byte{}
	for _, name := range s.names() {
		image[name] = s.content(name)
	}

	plain := openRW(t, reopened(image), "test.db", 1)
	var v string
	err := plain.QueryRow("SELECT v FROM t").Scan(&v)
	var se *sqlite.Error
	if !errors.As(err, &se) || se.Code()&0xff != sqlite3.SQLITE_CANTOPEN {
		t.Fatalf("a WAL database under normal locking: %v, want SQLITE_CANTOPEN", err)
	}

	excl := openRW(t, reopened(image), "test.db", 1)
	mustExec(t, excl, "PRAGMA locking_mode=EXCLUSIVE")
	if err := excl.QueryRow("SELECT v FROM t").Scan(&v); err != nil || v != "in the log" {
		t.Fatalf("a WAL database under exclusive locking: %q, %v", v, err)
	}
}

// TestRegisterPool has a pool of connections move amounts between rows of
// one database. The package's locking has to keep them apart: every transfer
// that committed must be in the final balances, the journal of a writer must
// not be taken for a hot one, and the file must stay sound. The connections
// have a busy timeout, so a transaction waits for a lock inside SQLite where
// it can; without one the retries below spin, a few hundred thousand of them.
func TestRegisterPool(t *testing.T) {
	const accounts, workers, transfers = 8, 6, 60
	s := newMemStore()
	db := openRW(t, s, "test.db", workers, "_pragma=busy_timeout(10000)")
	mustExec(t, db, "CREATE TABLE acct(id INTEGER PRIMARY KEY, bal INTEGER NOT NULL)", "CREATE TABLE done(w INTEGER, i INTEGER, src INTEGER, dst INTEGER, amt INTEGER)")
	for i := 0; i < accounts; i++ {
		mustExec(t, db, fmt.Sprintf("INSERT INTO acct VALUES(%d, 1000)", i))
	}

	type transfer struct{ w, i, src, dst, amt int }
	var (
		wg        sync.WaitGroup
		mu        sync.Mutex
		committed []transfer
		busy      int

		// Every worker holds its first read open until all of them have one,
		// so their first writes meet for certain.
		reading sync.WaitGroup
		release = make(chan struct{})
	)
	reading.Add(workers)
	go func() {
		reading.Wait()
		close(release)
	}()
	errc := make(chan error, workers)
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(int64(w)))
			first := true
			defer func() {
				if first {
					reading.Done() // failed before its first read: do not hold the others
				}
			}()
			for i := 0; i < transfers; i++ {
				tr := transfer{w, i, rng.Intn(accounts), rng.Intn(accounts), 1 + rng.Intn(50)}
				for attempt := 0; ; attempt++ {
					if attempt == 100000 {
						errc <- fmt.Errorf("worker %d, transfer %d: still SQLITE_BUSY after %d attempts", w, i, attempt)
						return
					}

					err := func() error {
						tx, err := db.Begin()
						if err != nil {
							return err
						}
						defer tx.Rollback()

						// A read first, so that the write upgrades a shared lock,
						// as two writers that deadlock in SQLite do.
						var bal int
						if err := tx.QueryRow("SELECT bal FROM acct WHERE id=?", tr.src).Scan(&bal); err != nil {
							return err
						}
						if first {
							first = false
							reading.Done()
							<-release
						}
						if _, err := tx.Exec("UPDATE acct SET bal=bal-? WHERE id=?", tr.amt, tr.src); err != nil {
							return err
						}
						if _, err := tx.Exec("UPDATE acct SET bal=bal+? WHERE id=?", tr.amt, tr.dst); err != nil {
							return err
						}
						if _, err := tx.Exec("INSERT INTO done VALUES(?,?,?,?,?)", tr.w, tr.i, tr.src, tr.dst, tr.amt); err != nil {
							return err
						}
						return tx.Commit()
					}()
					if err == nil {
						mu.Lock()
						committed = append(committed, tr)
						mu.Unlock()
						break
					}

					var se *sqlite.Error
					if errors.As(err, &se) && se.Code()&0xff == sqlite3.SQLITE_BUSY {
						mu.Lock()
						busy++
						mu.Unlock()
						continue
					}

					errc <- fmt.Errorf("worker %d, transfer %d: %w", w, i, err)
					return
				}
			}
		}(w)
	}
	wg.Wait()
	close(errc)
	for err := range errc {
		t.Error(err)
	}
	if t.Failed() {
		t.FailNow()
	}

	want := make([]int, accounts)
	for i := range want {
		want[i] = 1000
	}
	for _, tr := range committed {
		want[tr.src] -= tr.amt
		want[tr.dst] += tr.amt
	}
	var wantDump strings.Builder
	for i, b := range want {
		fmt.Fprintf(&wantDump, "%d|%d|\n", i, b)
	}
	if got := dump(t, db, "SELECT id, bal FROM acct ORDER BY id"); got != wantDump.String() {
		t.Fatalf("balances:\n%s\nwant, from the %d transfers that committed:\n%s", got, len(committed), wantDump.String())
	}

	if got, want := dump(t, db, "SELECT count(*) FROM done"), fmt.Sprintf("%d|\n", workers*transfers); got != want {
		t.Fatalf("transfers recorded: %s, want %s", got, want)
	}

	integrity(t, db)
	if busy < workers-1 {
		t.Fatalf("%d transaction(s) met a lock, want at least %d: all %d held a read lock when they went to write", busy, workers-1, workers)
	}
	t.Logf("%d transaction(s) met a lock", busy)
}

// TestRegisterOutFlags checks that of what Open reports only the access mode
// reaches SQLite, which would take the bit for an in-memory database at its
// word and stop journalling, and that a file reported read-only is one.
func TestRegisterOutFlags(t *testing.T) {
	s := newMemStore()
	s.outFlags = func(f OpenFlag) OpenFlag { return f | OpenFlag(sqlite3.SQLITE_OPEN_MEMORY) }
	db := openRW(t, s, "test.db", 1)
	mustExec(t, db, "CREATE TABLE t(v)")
	if got := dump(t, db, "PRAGMA journal_mode"); got != "delete|\n" {
		t.Fatalf("journal_mode is %q after Open reported SQLITE_OPEN_MEMORY, want delete", got)
	}

	ro := reopened(map[string][]byte{"test.db": s.content("test.db")})
	ro.readOnly = true
	rdb := openRW(t, ro, "test.db", 1)
	if got := dump(t, rdb, "SELECT count(*) FROM t"); got != "0|\n" {
		t.Fatalf("reading a read-only database: %q", got)
	}

	_, err := rdb.Exec("INSERT INTO t VALUES(1)")
	var se *sqlite.Error
	if !errors.As(err, &se) || se.Code()&0xff != sqlite3.SQLITE_READONLY {
		t.Fatalf("writing a database opened read-only: %v, want SQLITE_READONLY", err)
	}
}

// tempFiles returns how many files are open through the registration that
// the storage did not open: SQLite's temporary files, which the package keeps
// in memory.
func tempFiles(fsys *FS, s *memStore) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return int(fsys.open.Load()) - s.live
}

// TestRegisterTempFiles makes SQLite open temporary files, which it does
// through the VFS of the connection and with no name: a temporary table too
// large for its cache, a sort too large for memory, and a statement journal
// past its spill limit, used to undo an update that fails half way.
func TestRegisterTempFiles(t *testing.T) {
	s := newMemStore()
	vfsName, fsys, err := Register(s)
	if err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("sqlite", "file:test.db?vfs="+vfsName)
	if err != nil {
		t.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
		if n := tempFiles(fsys, s); n != 0 {
			t.Errorf("%d temporary file(s) open after the database was closed", n)
		}
		if err := fsys.Close(); err != nil {
			t.Error(err)
		}
	}()

	const rows = "WITH RECURSIVE n(i) AS (SELECT 1 UNION ALL SELECT i+1 FROM n WHERE i<6000) "
	mustExec(t, db,
		"PRAGMA cache_size=20",
		"CREATE TABLE t(id INTEGER PRIMARY KEY, k INTEGER UNIQUE, v BLOB)",
		rows+"INSERT INTO t SELECT i, i, zeroblob(900) FROM n",
	)
	if n := tempFiles(fsys, s); n != 0 {
		t.Fatalf("%d temporary file(s) open before any was needed", n)
	}

	mustExec(t, db, "CREATE TEMP TABLE big AS "+rows+"SELECT i AS id, randomblob(1000) AS v FROM n")
	if n := tempFiles(fsys, s); n < 1 {
		t.Fatalf("a temporary table of 6 MB opened %d temporary file(s)", n)
	}

	if got := dump(t, db, "SELECT count(*), sum(length(v)), sum(id) FROM big"); got != "6000|6000000|18003000|\n" {
		t.Fatalf("the temporary table reads as %q", got)
	}

	// A sort of 6 MB, against the same rows sorted through the index.
	if got, want := dump(t, db, "SELECT id FROM (SELECT id, v FROM big ORDER BY v) ORDER BY id LIMIT 3"), "1|\n2|\n3|\n"; got != want {
		t.Fatalf("after the sort: %q", got)
	}

	var prev []byte
	sorted, err := db.Query("SELECT v FROM big ORDER BY v")
	if err != nil {
		t.Fatal(err)
	}
	n := 0
	for sorted.Next() {
		var v []byte
		if err := sorted.Scan(&v); err != nil {
			t.Fatal(err)
		}
		if string(v) < string(prev) {
			t.Fatalf("row %d of the sort is out of order", n)
		}
		prev, n = v, n+1
	}
	if err := sorted.Close(); err != nil || n != 6000 {
		t.Fatalf("the sort returned %d rows, %v", n, err)
	}

	// A statement journal holds the pages a statement changes that the
	// transaction had changed before it. So the first update rewrites every
	// row, and the second does it again and fails on its last row, which it
	// gives a key the first row has by then: megabytes of pages go to the
	// statement journal, a temporary file next to the one of the table above,
	// and come back from it.
	mustExec(t, db, "BEGIN", "UPDATE t SET v=zeroblob(901)")
	peak := 0
	s.mu.Lock()
	s.after = func(int) { peak = max(peak, int(fsys.open.Load())-s.live) }
	s.mu.Unlock()
	_, err = db.Exec("UPDATE t SET k = CASE WHEN id < 6000 THEN k+100000 ELSE 100001 END, v=randomblob(900)")
	s.mu.Lock()
	s.after = nil
	s.mu.Unlock()
	if err == nil {
		t.Fatal("the update that breaks UNIQUE succeeded")
	}

	if peak < 2 {
		t.Fatalf("at most %d temporary file(s) were open during the update, want the table and the statement journal", peak)
	}

	mustExec(t, db, "COMMIT")
	if got := dump(t, db, "SELECT count(*), sum(k), sum(length(v)) FROM t"); got != "6000|18003000|5406000|\n" {
		t.Fatalf("after the failed update: %q", got)
	}

	integrity(t, db)
}

// TestTempFile checks the file the package keeps SQLite's temporary files in:
// reads past the end, a write that leaves a gap, growth, and truncation in
// both directions.
func TestTempFile(t *testing.T) {
	var f tempFile
	read := func(off int64, n int) (string, error) {
		p := make([]byte, n)
		m, err := f.ReadAt(p, off)
		return string(p[:m]), err
	}
	if got, err := read(0, 4); got != "" || err != io.EOF {
		t.Fatalf("reading an empty file: %q, %v", got, err)
	}

	if n, err := f.WriteAt([]byte("abc"), 5); n != 3 || err != nil {
		t.Fatalf("WriteAt: %d, %v", n, err)
	}

	if got, err := read(0, 8); got != "\x00\x00\x00\x00\x00abc" || err != nil {
		t.Fatalf("after a write that leaves a gap: %q, %v", got, err)
	}

	if got, err := read(6, 8); got != "bc" || err != io.EOF {
		t.Fatalf("reading across the end: %q, %v", got, err)
	}

	if err := f.Truncate(6); err != nil {
		t.Fatal(err)
	}

	// The bytes cut off must not come back when the file grows again.
	if n, err := f.WriteAt([]byte("z"), 9); n != 1 || err != nil {
		t.Fatalf("WriteAt: %d, %v", n, err)
	}

	if got, err := read(0, 10); got != "\x00\x00\x00\x00\x00a\x00\x00\x00z" || err != nil {
		t.Fatalf("after cutting and growing: %q, %v", got, err)
	}

	// Nor when it grows within the memory it already has.
	if err := f.Truncate(6); err != nil {
		t.Fatal(err)
	}

	if n, err := f.WriteAt([]byte("y"), 12); n != 1 || err != nil {
		t.Fatalf("WriteAt: %d, %v", n, err)
	}

	if got, err := read(0, 13); got != "\x00\x00\x00\x00\x00a\x00\x00\x00\x00\x00\x00y" || err != nil {
		t.Fatalf("after cutting and growing again: %q, %v", got, err)
	}

	if err := f.Truncate(15); err != nil {
		t.Fatal(err)
	}

	if size, err := f.Size(); size != 15 || err != nil {
		t.Fatalf("Size: %d, %v", size, err)
	}

	if got, _ := read(12, 3); got != "y\x00\x00" {
		t.Fatalf("after growing by Truncate: %q", got)
	}

	big := bytes.Repeat([]byte{7}, 1<<20)
	if n, err := f.WriteAt(big, 3); n != len(big) || err != nil {
		t.Fatalf("WriteAt: %d, %v", n, err)
	}

	if size, _ := f.Size(); size != 3+1<<20 {
		t.Fatalf("Size after a megabyte: %d", size)
	}

	if _, err := f.WriteAt(nil, -1); err == nil {
		t.Fatal("a write at a negative offset succeeded")
	}

	// Sizes no slice can have are an error, not a panic.
	if _, err := f.WriteAt([]byte("ab"), math.MaxInt64-1); !errors.Is(err, ErrFull) {
		t.Fatalf("a write that ends past the largest offset: %v, want ErrFull", err)
	}

	if strconv.IntSize == 32 {
		if _, err := f.WriteAt([]byte("ab"), math.MaxInt32); !errors.Is(err, ErrFull) {
			t.Fatalf("a write that ends past the largest slice: %v, want ErrFull", err)
		}
	}

	if size, _ := f.Size(); size != 3+1<<20 {
		t.Fatalf("Size after the refused writes: %d", size)
	}

	if err := f.Sync(SyncFull); err != nil {
		t.Fatal(err)
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}

// TestRegisterCurrentTime checks that 'now' has milliseconds on a connection
// through a VFS made by Register: a database that can be written stores it.
func TestRegisterCurrentTime(t *testing.T) {
	checkNow(t, openRW(t, newMemStore(), "test.db", 1))
}
