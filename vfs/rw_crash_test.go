// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vfs

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// The transaction the crash tests interrupt. With a cache of five pages it
// writes database pages before it commits, so that some of the storage
// states in the middle of it hold new pages next to the journal that undoes
// them. What it writes depends on the row and on nothing else, so every run
// of it ends in the same database.
var (
	crashSetup = []string{
		"CREATE TABLE t(id INTEGER PRIMARY KEY, v BLOB)",
		"WITH RECURSIVE n(i) AS (SELECT 1 UNION ALL SELECT i+1 FROM n WHERE i<120) INSERT INTO t SELECT i, zeroblob(300) FROM n",
	}
	crashTxn = []string{
		"BEGIN",
		"UPDATE t SET v=" + crashBlob(300, 40503) + " WHERE id%2=0",
		"INSERT INTO t SELECT id+1000, " + crashBlob(200, 30011) + " FROM t WHERE id<40",
		"DELETE FROM t WHERE id%5=0",
		"COMMIT",
	}
)

// crashBlob returns an SQL expression for a blob of n bytes, n even, that
// differs from row to row: four hexadecimal digits made of the id, repeated.
func crashBlob(n, mul int) string {
	return fmt.Sprintf("CAST(substr(replace(hex(zeroblob(%d)), '00', printf('%%04x', id*%d%%65536)), 1, %d) AS BLOB)", n/2, mul, n)
}

// Every row with every byte of it.
const crashQuery = "SELECT id, hex(v) FROM t ORDER BY id"

// allFiles returns a copy of every file of s as the process sees it.
func allFiles(s *memStore) map[string][]byte {
	img := map[string][]byte{}
	for _, name := range s.names() {
		img[name] = s.content(name)
	}
	return img
}

// withDB opens test.db on s through a registration of its own, with a cache
// of five pages, and closes both again.
func withDB(t *testing.T, s *memStore, pragmas []string, f func(db *sql.DB)) {
	t.Helper()
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
			t.Errorf("closing the database: %v", err)
		}
		if err := fsys.Close(); err != nil {
			t.Errorf("closing the VFS: %v", err)
		}
	}()
	mustExec(t, db, append([]string{"PRAGMA cache_size=5"}, pragmas...)...)
	f(db)
}

// crashStart returns the files of a database holding crashSetup, and what
// crashQuery says of it.
func crashStart(t *testing.T) (image map[string][]byte, old string) {
	t.Helper()
	s := newMemStore()
	withDB(t, s, nil, func(db *sql.DB) {
		mustExec(t, db, crashSetup...)
		old = dump(t, db, crashQuery)
	})
	return allFiles(s), old
}

// recovered opens the files a crash left and returns what crashQuery says of
// the database SQLite makes of them, which must be a sound one.
func recovered(t *testing.T, image map[string][]byte, what string) (state string) {
	t.Helper()
	withDB(t, reopened(image), nil, func(db *sql.DB) {
		if got := dump(t, db, "PRAGMA integrity_check"); got != "ok|\n" {
			t.Fatalf("%s: integrity_check: %s", what, got)
		}
		state = dump(t, db, crashQuery)
	})
	return state
}

// TestRegisterPowerLoss cuts the power after every operation of a
// transaction on the storage and opens what is left, in four versions: only
// what was synced, everything that was written, of what was not synced an
// arbitrary set of 512-byte pieces, and of the writes that were not synced
// some done, some not, and some leaving garbage in the 4096-byte sectors they
// touch, which is the damage SQLite is told to expect. Each must come up as
// the database before the transaction or after it. Once COMMIT has returned
// it must be the one after it wherever SQLite promises that: with a truncated
// or persistent journal, and with a deleted one under synchronous=EXTRA,
// which is what syncs the directory after the deletion. Under
// synchronous=FULL the deletion may be lost with the power, and the journal
// that comes back undoes the commit.
func TestRegisterPowerLoss(t *testing.T) {
	start, old := crashStart(t)
	for _, c := range []struct {
		name    string
		pragmas []string
		durable bool
	}{
		{"delete, FULL", []string{"PRAGMA journal_mode=DELETE", "PRAGMA synchronous=FULL"}, false},
		{"delete, EXTRA", []string{"PRAGMA journal_mode=DELETE", "PRAGMA synchronous=EXTRA"}, true},
		{"truncate", []string{"PRAGMA journal_mode=TRUNCATE", "PRAGMA synchronous=FULL"}, true},
		{"persist", []string{"PRAGMA journal_mode=PERSIST", "PRAGMA synchronous=FULL"}, true},
	} {
		t.Run(c.name, func(t *testing.T) {
			type point struct {
				n      int
				images [5]map[string][]byte
			}
			var (
				points []point
				final  point
				new    string
			)
			s := reopened(start)
			rng := rand.New(rand.NewSource(1))
			snap := func(n int) point {
				return point{n, [5]map[string][]byte{
					s.image(crashSynced, nil), s.image(crashWritten, nil), s.image(crashTorn, rng), s.image(crashTorn, rng), s.image(crashSector, rng),
				}}
			}
			withDB(t, s, c.pragmas, func(db *sql.DB) {
				s.mu.Lock()
				s.after = func(n int) { points = append(points, snap(n)) }
				s.mu.Unlock()
				mustExec(t, db, crashTxn...)
				s.mu.Lock()
				s.after = nil
				final = snap(s.ops)
				s.mu.Unlock()
				new = dump(t, db, crashQuery)
			})
			if new == old {
				t.Fatal("the transaction changed nothing")
			}

			if len(points) < 20 {
				t.Fatalf("the transaction made %d operations on the storage, too few to sweep", len(points))
			}

			kinds := [5]string{"synced only", "all written", "torn", "torn", "spoiled sectors"}
			var sawOld, sawNew, undone int
			for _, p := range points {
				for i, img := range p.images {
					switch recovered(t, img, c.name) {
					case old:
						sawOld++
						if !bytes.Equal(img["test.db"], start["test.db"]) {
							undone++ // new pages were in the file and the journal took them out
						}
					case new:
						sawNew++
					default:
						t.Fatalf("power lost after operation %d (%s): the database is neither the one before the transaction nor the one after it\n%s",
							p.n, kinds[i], s.trace[max(0, p.n-6):p.n])
					}
				}
			}
			if sawOld == 0 || sawNew == 0 || undone == 0 {
				t.Fatalf("of %d crashes %d came up old, %d new, and %d had new pages undone: the sweep did not cover the transaction", 5*len(points), sawOld, sawNew, undone)
			}

			for i, img := range final.images {
				got := recovered(t, img, c.name)
				switch {
				case got == new:
				case got == old && !c.durable && i != 1:
					// The deletion of the journal was not synced.
				default:
					t.Fatalf("power lost after COMMIT returned (%s): the committed transaction is not there", kinds[i])
				}
			}
		})
	}
}

// TestRegisterSectorSize is about the sector size the VFS reports. With pages
// of 1024 bytes a transaction that changes three rows writes a few pages, and
// a write the power cuts short may spoil the rest of its 4096-byte sector:
// pages the transaction never touched. SQLite puts those in the journal as
// well because the VFS says a sector has 4096 bytes. Told 512, which is what
// it makes of 0, it would not, and they would stay spoiled.
func TestRegisterSectorSize(t *testing.T) {
	var old, new string
	s := newMemStore()
	withDB(t, s, []string{"PRAGMA page_size=1024"}, func(db *sql.DB) {
		mustExec(t, db, crashSetup...)
		if got := dump(t, db, "PRAGMA page_size"); got != "1024|\n" {
			t.Fatalf("page_size is %s", got)
		}
		old = dump(t, db, crashQuery)
	})
	start := allFiles(s)

	var images []map[string][]byte
	s = reopened(start)
	rng := rand.New(rand.NewSource(2))
	withDB(t, s, []string{"PRAGMA journal_mode=TRUNCATE"}, func(db *sql.DB) {
		s.mu.Lock()
		s.after = func(int) {
			for i := 0; i < 8; i++ {
				images = append(images, s.image(crashSector, rng))
			}
		}
		s.mu.Unlock()
		mustExec(t, db, "BEGIN", "UPDATE t SET v="+crashBlob(300, 40503)+" WHERE id IN (7, 60, 113)", "COMMIT")
		s.mu.Lock()
		s.after = nil
		s.mu.Unlock()
		new = dump(t, db, crashQuery)
	})

	var sawOld, sawNew, spoiled int
	for _, img := range images {
		if bytes.Contains(img["test.db"], bytes.Repeat([]byte{0xee}, 2048)) {
			spoiled++ // more than the one page of a write
		}
		switch recovered(t, img, "spoiled sectors") {
		case old:
			sawOld++
		case new:
			sawNew++
		default:
			t.Fatal("the database is neither the one before the transaction nor the one after it")
		}
	}
	if spoiled < 10 || sawOld == 0 {
		t.Fatalf("of %d crashes %d had a sector spoiled in the database and %d came up old: the sweep did not cover the case", len(images), spoiled, sawOld)
	}
	t.Logf("%d crashes: %d with a spoiled sector in the database, %d came up old, %d new", len(images), spoiled, sawOld, sawNew)
}

// TestRegisterFaults fails every operation of the transaction on the storage
// in turn. The statement that meets the failure must return an I/O error, or
// succeed if SQLite does not need the operation, and whichever it was, the
// files then hold a sound database from before the transaction or after it.
func TestRegisterFaults(t *testing.T) {
	start, old := crashStart(t)
	pragmas := []string{"PRAGMA journal_mode=DELETE"}

	// A run without faults counts the operations and gives the end state.
	var (
		total int
		new   string
	)
	s := reopened(start)
	withDB(t, s, pragmas, func(db *sql.DB) {
		before := s.ops
		mustExec(t, db, crashTxn...)
		total = s.ops - before
		new = dump(t, db, crashQuery)
	})
	if total < 20 {
		t.Fatalf("the transaction made %d operations on the storage, too few to sweep", total)
	}

	boom := errors.New("injected")
	var failed, sawOld, sawNew int
	for k := 1; k <= total; k++ {
		s := reopened(start)
		var (
			injected string
			txnErr   error
		)
		withDB(t, s, pragmas, func(db *sql.DB) {
			s.mu.Lock()
			at := s.ops + k
			s.fail = func(n int, op, name string) error {
				if n == at {
					injected = op + " " + kindOf(name)
					return boom
				}
				return nil
			}
			s.mu.Unlock()
			for _, stmt := range crashTxn {
				if _, txnErr = db.Exec(stmt); txnErr != nil {
					break
				}
			}
			s.mu.Lock()
			s.fail = nil
			s.mu.Unlock()
			if txnErr != nil {
				db.Exec("ROLLBACK")
			}
		})
		if injected == "" {
			t.Fatalf("operation %d of %d was not reached", k, total)
		}

		if txnErr != nil {
			failed++
			var se *sqlite.Error
			if !errors.As(txnErr, &se) || (se.Code()&0xff != sqlite3.SQLITE_IOERR && se.Code()&0xff != sqlite3.SQLITE_CANTOPEN) {
				t.Fatalf("operation %d (%s) failed: the statement returned %v, want an I/O error", k, injected, txnErr)
			}
		}

		switch recovered(t, allFiles(s), injected) {
		case old:
			sawOld++
		case new:
			sawNew++
			if txnErr != nil && injected != "delete journal" {
				t.Fatalf("operation %d (%s) failed and the statement returned %v, but the transaction is in the database", k, injected, txnErr)
			}
		default:
			t.Fatalf("operation %d (%s) failed: the database is neither the one before the transaction nor the one after it", k, injected)
		}
	}
	if failed < total/2 || sawOld == 0 {
		t.Fatalf("of %d faults %d failed a statement and %d left the old database: the sweep did not cover the transaction", total, failed, sawOld)
	}
	t.Logf("%d faults: %d failed a statement, %d left the old database, %d the new one", total, failed, sawOld, sawNew)
}
