// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	sqlite3 "modernc.org/sqlite/lib"
)

func registerOver(t testing.TB, base string, key []byte) string {
	t.Helper()
	name, closer, err := Register(base, key)
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

// snapshot is the content of every file of a directory, and restore makes
// the directory hold exactly that again: the files a crash would have left.
func snapshot(t testing.TB, dir string) map[string][]byte {
	ents, err := os.ReadDir(dir)
	if err != nil {
		t.Error(err)
		return nil
	}

	m := map[string][]byte{}
	for _, e := range ents {
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Error(err)
			return nil
		}

		m[e.Name()] = b
	}

	return m
}

func restore(t testing.TB, dir string, snap map[string][]byte) {
	t.Helper()
	ents, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range ents {
		if err := os.Remove(filepath.Join(dir, e.Name())); err != nil {
			t.Fatal(err)
		}
	}

	for name, b := range snap {
		if err := os.WriteFile(filepath.Join(dir, name), b, 0o600); err != nil {
			t.Fatal(err)
		}
	}
}

func isSuper(name string) bool { return strings.Contains(filepath.Base(name), "-mj") }

func fileNames(m map[string][]byte) string {
	var names []string
	for _, n := range []string{"a.db", "a.db-journal", "b.db", "b.db-journal"} {
		if _, ok := m[n]; ok {
			names = append(names, n)
		}
	}

	for n := range m {
		if isSuper(n) {
			names = append(names, "super-journal")
		}
	}

	return strings.Join(names, " ")
}

// A transaction over two databases, stopped where a crash hurts most: just
// before SQLite deletes the super-journal, which is the commit, and just
// after. Whichever database is opened first afterwards, and through
// whichever VFS it lives on, both come back as they were before the
// transaction in the first case and as it left them in the second.
func TestAttachedCommit(t *testing.T) {
	h := newHookVFS(t)
	one, two := registerOver(t, h.name, testKey(20)), registerOver(t, h.name, testKey(21))
	for _, tc := range []struct {
		what    string
		a, b    string // the VFS of the main and of the attached database
		aKeyed  bool
		bKeyed  bool
		persist bool
	}{
		{"one key", one, one, true, true, false},
		{"two keys", one, two, true, true, false},
		{"encrypted main, plain attached", one, h.name, true, false, false},
		{"one key, persistent journals", one, one, true, true, true},
	} {
		t.Run(tc.what, func(t *testing.T) {
			dir := t.TempDir()
			pa, pb := filepath.Join(dir, "a.db"), filepath.Join(dir, "b.db")
			q := "_pragma=cache_size(10)"
			if tc.persist {
				q += "&_pragma=journal_mode(persist)"
			}

			db := open(t, pa, tc.a, q)
			db.SetMaxOpenConns(1)
			mustExec(t, db, fmt.Sprintf("attach database 'file:%s?vfs=%s' as b", pb, tc.b))
			if tc.persist {
				mustExec(t, db, "pragma b.journal_mode=persist")
			}

			for _, s := range []string{"main", "b"} {
				mustExec(t, db, fmt.Sprintf("create table %s.t(a integer primary key, b text)", s))
				mustExec(t, db, fmt.Sprintf("with recursive s(v) as (select 1 union all select v+1 from s where v < 500) insert into %s.t(b) select 'old ' || ? || v from s", s), marker)
			}

			if tc.persist {
				// Longer journals than the transaction below will write.
				mustExec(t, db, "begin; update main.t set b = b || '.'; update b.t set b = b || '.'; delete from main.t where a > 400; delete from b.t where a > 400; commit")
			}

			var before, after map[string][]byte
			h.mu.Lock()
			h.beforeDelete = func(name string) {
				if isSuper(name) && before == nil {
					before = snapshot(t, dir)
				}
			}

			h.afterDelete = func(name string) {
				if isSuper(name) && after == nil {
					after = snapshot(t, dir)
				}
			}

			h.mu.Unlock()
			mustExec(t, db, "begin; update main.t set b = 'new' where a <= 50; update b.t set b = 'new' where a <= 50; commit")
			h.mu.Lock()
			h.beforeDelete, h.afterDelete = nil, nil
			h.mu.Unlock()
			db.Close()

			if got := fileNames(before); got != "a.db a.db-journal b.db b.db-journal super-journal" {
				t.Fatalf("before the commit there are: %s", got)
			}

			if got := fileNames(after); got != "a.db a.db-journal b.db b.db-journal" {
				t.Fatalf("after the commit there are: %s", got)
			}

			for name, content := range before {
				keyed := strings.HasPrefix(name, "a.db") && tc.aKeyed || strings.HasPrefix(name, "b.db") && tc.bKeyed
				if keyed && !isSuper(name) && (bytes.Contains(content, []byte(marker)) || bytes.Contains(content, []byte("SQLite format 3"))) {
					t.Errorf("%s holds plain text", name)
				}
			}

			for _, state := range []struct {
				what string
				snap map[string][]byte
				want string
			}{
				{"before the commit", before, "0"},
				{"after the commit", after, "50"},
			} {
				for _, first := range []string{"a", "b"} {
					what := fmt.Sprintf("%s, %s opened first", state.what, first)
					restore(t, dir, state.snap)
					order := []struct{ path, vfs string }{{pa, tc.a}, {pb, tc.b}}
					if first == "b" {
						order[0], order[1] = order[1], order[0]
					}

					for _, o := range order {
						d := open(t, o.path, o.vfs, "")
						if got := queryString(t, d, "select count(*) filter (where b = 'new') from t"); got != state.want {
							t.Errorf("%s: %s has %s new rows, want %s", what, filepath.Base(o.path), got, state.want)
						}

						if got := queryString(t, d, "pragma integrity_check"); got != "ok" {
							t.Errorf("%s: %s: integrity_check: %s", what, filepath.Base(o.path), got)
						}

						d.Close()
					}

					left := snapshot(t, dir)
					for name := range left {
						if isSuper(name) {
							t.Errorf("%s: the super-journal %s is left", what, name)
						}
					}
				}
			}
		})
	}
}

// What is not supported fails before anything is committed, or behaves as
// it does without the cipher.
func TestAttachedRefused(t *testing.T) {
	h := newHookVFS(t)
	enc := registerOver(t, h.name, testKey(22))
	elsewhere := register(t, testKey(23)) // another base
	for _, tc := range []struct{ what, a, b string }{
		{"plain main, encrypted attached", h.name, enc},
		{"main and attached on different bases", elsewhere, enc},
	} {
		t.Run(tc.what, func(t *testing.T) {
			dir := t.TempDir()
			pa, pb := filepath.Join(dir, "a.db"), filepath.Join(dir, "b.db")
			db := open(t, pa, tc.a, "")
			db.SetMaxOpenConns(1)
			mustExec(t, db, fmt.Sprintf("attach database 'file:%s?vfs=%s' as b", pb, tc.b))
			for _, s := range []string{"main", "b"} {
				mustExec(t, db, fmt.Sprintf("create table %s.t(a integer primary key, b text)", s))
				mustExec(t, db, fmt.Sprintf("insert into %s.t(b) values('old')", s))
			}

			// One database at a time works.
			mustExec(t, db, "update b.t set b = 'old.'")
			_, err := db.Exec("begin; update main.t set b = 'new'; update b.t set b = 'new'; commit")
			if err == nil {
				t.Fatal("the transaction over both was committed")
			}

			db.Exec("rollback")
			if got := queryString(t, db, "select (select b from main.t) || ' ' || (select b from b.t)"); got != "old old." {
				t.Fatalf("after the refusal: %s", got)
			}

			db.Close()
			for _, o := range []struct{ path, vfs string }{{pa, tc.a}, {pb, tc.b}} {
				d := open(t, o.path, o.vfs, "")
				if got := queryString(t, d, "select b from t"); !strings.HasPrefix(got, "old") {
					t.Errorf("%s: %s", filepath.Base(o.path), got)
				}

				if got := queryString(t, d, "pragma integrity_check"); got != "ok" {
					t.Errorf("%s: integrity_check: %s", filepath.Base(o.path), got)
				}

				d.Close()
			}
		})
	}

	// A database in journal_mode OFF or MEMORY keeps no journal on the
	// storage. SQLite writes it no record, nothing is asked of it, and the
	// transaction goes as it does without the cipher: all three databases
	// have the new row.
	for _, mode := range []string{"off", "memory"} {
		t.Run("encrypted attached, journal_mode="+mode, func(t *testing.T) {
			dir := t.TempDir()
			db := open(t, filepath.Join(dir, "a.db"), h.name, "")
			db.SetMaxOpenConns(1)
			mustExec(t, db, fmt.Sprintf("attach database 'file:%s?vfs=%s' as b", filepath.Join(dir, "b.db"), h.name))
			mustExec(t, db, fmt.Sprintf("attach database 'file:%s?vfs=%s' as c", filepath.Join(dir, "c.db"), enc))
			mustExec(t, db, "pragma c.journal_mode="+mode)
			for _, s := range []string{"main", "b", "c"} {
				mustExec(t, db, fmt.Sprintf("create table %s.t(b)", s))
				mustExec(t, db, fmt.Sprintf("insert into %s.t values('old')", s))
			}

			mustExec(t, db, "begin; update main.t set b = 'new'; update b.t set b = 'new'; update c.t set b = 'new'; commit")
			if got := queryString(t, db, "select (select b from main.t) || ' ' || (select b from b.t) || ' ' || (select b from c.t)"); got != "new new new" {
				t.Fatalf("after the commit: %s", got)
			}
		})
	}

	t.Run("main in memory", func(t *testing.T) {
		// SQLite uses no super-journal then, with or without the cipher.
		dir := t.TempDir()
		db := open(t, ":memory:", "", "")
		db.SetMaxOpenConns(1)
		for _, s := range []string{"x", "y"} {
			mustExec(t, db, fmt.Sprintf("attach database 'file:%s?vfs=%s' as %s", filepath.Join(dir, s+".db"), enc, s))
			mustExec(t, db, fmt.Sprintf("create table %s.t(b)", s))
		}

		mustExec(t, db, "begin; insert into x.t values('new'); insert into y.t values('new'); commit")
		if got := queryString(t, db, "select (select b from x.t) || ' ' || (select b from y.t)"); got != "new new" {
			t.Fatalf("after the commit: %s", got)
		}
	})
}

// Before the commit of a transaction over several databases SQLite tells
// every database the name of the super-journal. That is refused when the
// journal the connection has open for the database ends with the record of
// it and the copy for readers without the key is not there, whatever the
// reason, and when the package did not see the super-journal being created.
func TestCommitNeedsCopy(t *testing.T) {
	name, closer, err := Register("", testKey(24))
	if err != nil {
		t.Fatal(err)
	}

	defer closer.Close()

	r := closer.(*registration)
	path := filepath.Join(t.TempDir(), "a.db")
	super := path + "-mj0A1B2C9FF"
	f, rc := rawOpen(t, name, path, rawDB)
	if rc != 0 {
		t.Fatalf("open: %d", rc)
	}

	defer f.close()

	csuper := cstring(f.tls, super)
	defer sqlite3.Xsqlite3_free(f.tls, csuper)

	sync := func() bool {
		rc := f.fileControl(sqlite3.SQLITE_FCNTL_SYNC, csuper)
		return rc == sqlite3.SQLITE_OK || rc == sqlite3.SQLITE_NOTFOUND
	}

	// A connection that has no journal open for the database keeps none on
	// the storage: nothing is asked of it, whoever made the super-journal.
	if !sync() {
		t.Error("refused although the connection has no journal open")
	}

	names.add(super, r.base)
	defer names.remove(super)

	// The journal of the transaction, with the record at its end.
	j := f.journal()
	defer j.close()

	rec := superRecord(super, true)
	if rc := j.write(append(bytes.Clone(journalMagic), make([]byte, 504)...), 0); rc != 0 {
		t.Fatalf("write: %d", rc)
	}

	if !sync() {
		t.Error("refused although the journal has no record yet")
	}

	if rc := j.write(rec, 512); rc != 0 {
		t.Fatalf("write: %d", rc)
	}

	if !sync() {
		t.Error("refused although the record has its copy")
	}

	if rc := f.fileControl(sqlite3.SQLITE_FCNTL_SYNC, 0); rc != sqlite3.SQLITE_OK && rc != sqlite3.SQLITE_NOTFOUND {
		t.Errorf("a commit without a super-journal: %d", rc)
	}

	// The same record written while the name was not known: no copy.
	names.remove(super)
	if rc := j.truncate(0); rc != 0 {
		t.Fatalf("truncate: %d", rc)
	}

	if rc := j.write(append(bytes.Clone(journalMagic), make([]byte, 504)...), 0); rc != 0 {
		t.Fatalf("write: %d", rc)
	}

	if rc := j.write(rec, 512); rc != 0 {
		t.Fatalf("write: %d", rc)
	}

	names.add(super, r.base)
	if sync() {
		t.Error("let through although the record has no copy")
	}

	stored, err := os.ReadFile(path + "-journal")
	if err != nil || plainSuper(stored) != "" {
		t.Fatalf("the stored journal shows %q, %v", plainSuper(stored), err)
	}
}
