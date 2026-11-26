// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"modernc.org/sqlite/vfs"
)

// The transaction the crash tests interrupt, as in the tests of package vfs:
// with a cache of five pages it writes database pages before it commits, so
// that some of the states in the middle of it hold new pages next to the
// journal that undoes them.
func crashBlob(n, mul int) string {
	return fmt.Sprintf("unhex(substr(replace(hex(zeroblob(%d)), '0000', printf('%%04x', id*%d%%65536)), 1, %d))", n/2, mul, 2*n)
}

var (
	crashSetup = []string{
		"CREATE TABLE %st(id INTEGER PRIMARY KEY, v BLOB)",
		"WITH RECURSIVE n(i) AS (SELECT 1 UNION ALL SELECT i+1 FROM n WHERE i<120) INSERT INTO %st SELECT i, zeroblob(300) FROM n",
	}
	crashChange = []string{
		"UPDATE %st SET v=" + crashBlob(300, 40503) + " WHERE id%2=0",
		"INSERT INTO %st SELECT id+1000, " + crashBlob(200, 30011) + " FROM %st WHERE id<40",
		"DELETE FROM %st WHERE id%5=0",
	}
)

const crashQuery = "SELECT coalesce(group_concat(id || ':' || hex(v), ' '), '') FROM (SELECT id, v FROM t ORDER BY id)"

func execAll(t testing.TB, db *sql.DB, schema string, stmts []string) {
	t.Helper()
	for _, s := range stmts {
		mustExec(t, db, strings.ReplaceAll(s, "%s", schema))
	}
}

// crashRig is a crash store with the VFSes of a test on it: the plain one of
// package vfs and encrypting ones over it, one for each key.
type crashRig struct {
	t     testing.TB
	s     *crashStore
	plain string
	enc   []string
	close func()
}

func newCrashRig(t testing.TB, s *crashStore, keys ...[]byte) *crashRig {
	t.Helper()
	plain, fsys, err := vfs.Register(s)
	if err != nil {
		t.Fatal(err)
	}

	r := &crashRig{t: t, s: s, plain: plain}
	var closers []func() error
	for _, k := range keys {
		name, closer, err := Register(plain, k)
		if err != nil {
			t.Fatal(err)
		}

		r.enc = append(r.enc, name)
		closers = append(closers, closer.Close)
	}

	r.close = func() {
		for _, c := range closers {
			if err := c(); err != nil {
				t.Errorf("closing the encrypting VFS: %v", err)
			}
		}

		if err := fsys.Close(); err != nil {
			t.Errorf("closing the VFS: %v", err)
		}
	}
	return r
}

// db opens name through the VFS called vfsName, with a cache of five pages.
func (r *crashRig) db(name, vfsName string, pragmas ...string) *sql.DB {
	r.t.Helper()
	db, err := sql.Open("sqlite", "file:"+name+"?vfs="+vfsName)
	if err != nil {
		r.t.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	for _, p := range append([]string{"PRAGMA cache_size=5"}, pragmas...) {
		if _, err := db.Exec(p); err != nil {
			db.Close()
			r.t.Fatalf("%s: %v", p, err)
		}
	}

	return db
}

// state opens the database called name of an image and returns what it
// holds, which must be a sound database.
func (r *crashRig) state(name, vfsName, what string) string {
	r.t.Helper()
	db := r.db(name, vfsName)
	defer db.Close()

	if got := queryString(r.t, db, "PRAGMA integrity_check"); got != "ok" {
		r.t.Fatalf("%s: %s: integrity_check: %s", what, name, got)
	}

	return queryString(r.t, db, crashQuery)
}

// images returns what a loss of power at this moment may leave: only what
// was synced, everything written, twice an arbitrary part of what was not
// synced, and unsynced writes that spoil the sectors they touch.
func crashImages(s *crashStore, rng *rand.Rand) [5]map[string][]byte {
	return [5]map[string][]byte{
		s.image(crashSynced, nil), s.image(crashWritten, nil), s.image(crashTorn, rng), s.image(crashTorn, rng), s.image(crashSector, rng),
	}
}

var crashKinds = [5]string{"synced only", "all written", "torn", "torn", "spoiled sectors"}

// TestPowerLoss cuts the power after every operation of a transaction on
// the storage under the encrypting VFS and opens what is left. It must be
// the database before the transaction or the one after it, and the one
// after it once COMMIT has returned, wherever SQLite promises that.
func TestPowerLoss(t *testing.T) {
	key := testKey(40)
	first := newCrashStore()
	rig := newCrashRig(t, first, key)
	db := rig.db("test.db", rig.enc[0])
	execAll(t, db, "", crashSetup)
	old := queryString(t, db, crashQuery)
	db.Close()
	rig.close()
	start := map[string][]byte{}
	for _, name := range first.names() {
		start[name] = first.content(name)
	}

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

			var points []point
			s := reopened(start)
			rng := rand.New(rand.NewSource(1))
			rig := newCrashRig(t, s, key)
			db := rig.db("test.db", rig.enc[0], c.pragmas...)
			s.mu.Lock()
			s.after = func(n int) { points = append(points, point{n, crashImages(s, rng)}) }
			s.mu.Unlock()
			mustExec(t, db, "BEGIN")
			execAll(t, db, "", crashChange)
			mustExec(t, db, "COMMIT")
			s.mu.Lock()
			s.after = nil
			final := crashImages(s, rng)
			s.mu.Unlock()
			new := queryString(t, db, crashQuery)
			db.Close()
			rig.close()
			if new == old || len(points) < 20 {
				t.Fatalf("the transaction changed nothing, or made only %d operations on the storage", len(points))
			}

			recovered := func(img map[string][]byte, what string) string {
				rig := newCrashRig(t, reopened(img), key)
				defer rig.close()

				return rig.state("test.db", rig.enc[0], what)
			}

			var sawOld, sawNew, undone int
			for _, p := range points {
				for i, img := range p.images {
					what := fmt.Sprintf("power lost after operation %d (%s)", p.n, crashKinds[i])
					switch recovered(img, what) {
					case old:
						sawOld++
						if !bytes.Equal(img["test.db"], start["test.db"]) {
							undone++
						}
					case new:
						sawNew++
					default:
						t.Fatalf("%s: the database is neither the one before the transaction nor the one after it\n%s", what, s.trace[max(0, p.n-6):p.n])
					}
				}
			}

			if sawOld == 0 || sawNew == 0 || undone == 0 {
				t.Fatalf("of %d crashes %d came up old, %d new, and %d had new pages undone: the sweep did not cover the transaction", 5*len(points), sawOld, sawNew, undone)
			}

			for i, img := range final {
				got := recovered(img, "power lost after COMMIT ("+crashKinds[i]+")")
				switch {
				case got == new:
				case got == old && !c.durable && i != 1:
					// The deletion of the journal was not synced.
				default:
					t.Fatalf("power lost after COMMIT returned (%s): the committed transaction is not there", crashKinds[i])
				}
			}
		})
	}
}

// TestPowerLossAttached does the same to a transaction over two databases,
// with one key, two keys, and a plain database attached to an encrypted
// main one. Whichever of the two is opened first after the loss of power,
// both are as they were before the transaction or both as it left them.
func TestPowerLossAttached(t *testing.T) {
	keys := [][]byte{testKey(41), testKey(42)}
	// The VFS of the attached database: the first encrypting one, the
	// second, or the plain one.
	pick := func(rig *crashRig, which int) string {
		if which < 0 {
			return rig.plain
		}

		return rig.enc[which]
	}

	// Every operation of the transaction is a point at which the power is
	// cut. With two keys all of them are tried; the other combinations
	// differ from that one in a single thing each and try every third, and
	// all of the last fifteen, where the commit is. The race detector makes
	// the sweep ten times slower: it gets every fourth point with two keys.
	for _, c := range []struct {
		name     string
		attached int
		pragma   string
		every    int
	}{
		{"two keys", 1, "DELETE", 1},
		{"one key", 0, "DELETE", 3},
		{"plain attached", -1, "DELETE", 3},
		{"one key, persist", 0, "PERSIST", 3},
	} {
		t.Run(c.name, func(t *testing.T) {
			if (testing.Short() || raceDetector) && c.every != 1 {
				t.Skip("-short, or the race detector")
			}

			if raceDetector {
				c.every = 4
			}

			type point struct {
				n      int
				images [5]map[string][]byte
			}

			var points []point
			attach := func(rig *crashRig) *sql.DB {
				db := rig.db("test.db", rig.enc[0], "PRAGMA journal_mode="+c.pragma)
				mustExec(t, db, fmt.Sprintf("ATTACH DATABASE 'file:other.db?vfs=%s' AS o", pick(rig, c.attached)))
				mustExec(t, db, "PRAGMA o.cache_size=5")
				mustExec(t, db, "PRAGMA o.journal_mode="+c.pragma)
				return db
			}

			first := newCrashStore()
			rig := newCrashRig(t, first, keys...)
			db := attach(rig)
			execAll(t, db, "", crashSetup)
			execAll(t, db, "o.", crashSetup)
			old := queryString(t, db, crashQuery)
			db.Close()
			rig.close()
			start := map[string][]byte{}
			for _, name := range first.names() {
				start[name] = first.content(name)
			}

			s := reopened(start)
			rig = newCrashRig(t, s, keys...)
			db = attach(rig)
			rng := rand.New(rand.NewSource(2))
			s.mu.Lock()
			s.after = func(n int) { points = append(points, point{n, crashImages(s, rng)}) }
			s.mu.Unlock()
			mustExec(t, db, "BEGIN")
			execAll(t, db, "", crashChange)
			execAll(t, db, "o.", crashChange)
			mustExec(t, db, "COMMIT")
			s.mu.Lock()
			s.after = nil
			final := crashImages(s, rng)
			trace := s.trace
			s.mu.Unlock()
			new := queryString(t, db, crashQuery)
			db.Close()
			rig.close()
			if new == old || len(points) < 40 {
				t.Fatalf("the transaction changed nothing, or made only %d operations on the storage", len(points))
			}

			// both returns the two databases of an image, the one named
			// first opened first.
			both := func(img map[string][]byte, mainFirst bool, what string) (main, other string) {
				rig := newCrashRig(t, reopened(img), keys...)
				defer rig.close()

				if mainFirst {
					main = rig.state("test.db", rig.enc[0], what)
				}

				other = rig.state("other.db", pick(rig, c.attached), what)
				if !mainFirst {
					main = rig.state("test.db", rig.enc[0], what)
				}

				return main, other
			}

			var sawOld, sawNew int
			check := func(img map[string][]byte, what string, committed bool) {
				for _, mainFirst := range []bool{true, false} {
					main, other := both(img, mainFirst, what)
					switch {
					case main == old && other == old && !committed:
						sawOld++
					case main == new && other == new:
						sawNew++
					default:
						name := func(s string) string {
							switch s {
							case old:
								return "old"
							case new:
								return "new"
							}

							return "neither"
						}
						t.Fatalf("%s, main opened first: %v: the main database is %s and the attached one %s", what, mainFirst, name(main), name(other))
					}
				}
			}

			for k, p := range points {
				// The points left out move with the kind of loss, so that
				// each operation is tried in some of them.
				for i, img := range p.images {
					if (k+i)%c.every != 0 && k < len(points)-15 {
						continue // all of the points around the commit are tried
					}

					what := fmt.Sprintf("power lost after operation %d of %d (%s; %s)", p.n, len(points), crashKinds[i], strings.TrimSpace(trace[p.n-1]))
					check(img, what, false)
				}
			}

			if sawOld == 0 || sawNew == 0 {
				t.Fatalf("%d crashes came up old and %d new: the sweep did not cover the transaction", sawOld, sawNew)
			}

			for i, img := range final {
				check(img, "power lost after COMMIT returned ("+crashKinds[i]+")", true)
			}
		})
	}
}
