// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"slices"
	"strings"
	"sync/atomic"
	"testing"

	"modernc.org/sqlite/vtab"
)

// The TEXT values the tests below pass around: with a NUL in the middle, at
// the start and at the end, alone, and without one.
var textNULValues = []string{"alice\x00x", "\x00alice", "alice\x00", "\x00", "a\x00b\x00c", "alice", ""}

// textNULRun numbers the runs of the tests below. Functions and modules are
// registered for the life of the process, so each run needs names of its own
// to work under -count.
var textNULRun atomic.Int32

// textNULKeep returns what a callback may keep of a TEXT argument. Without
// VolatileArgs the argument is the callback's own and is kept as it came, so
// that the tests also notice one that is still SQLite's memory; with
// VolatileArgs it has to be copied.
func textNULKeep(v any, volatile bool) string {
	if volatile {
		return strings.Clone(v.(string))
	}

	return v.(string)
}

// textNULRecorder is an aggregate that records the TEXT argument of every
// Step and WindowInverse.
type textNULRecorder struct {
	step, inverse *[]string
	volatile      bool
}

func (r *textNULRecorder) Step(ctx *FunctionContext, args []driver.Value) error {
	*r.step = append(*r.step, textNULKeep(args[0], r.volatile))
	return nil
}

func (r *textNULRecorder) WindowInverse(ctx *FunctionContext, args []driver.Value) error {
	*r.inverse = append(*r.inverse, textNULKeep(args[0], r.volatile))
	return nil
}

func (r *textNULRecorder) WindowValue(ctx *FunctionContext) (driver.Value, error) { return nil, nil }

func (r *textNULRecorder) Final(ctx *FunctionContext) {}

// TestTextNULFunctionArgs checks that a function written in Go gets a TEXT
// argument whole when it holds a NUL: a scalar function, the Step of an
// aggregate and the WindowInverse of a window function, each with and without
// VolatileArgs. Without it the argument used to end at the first NUL.
func TestTextNULFunctionArgs(t *testing.T) {
	for _, volatile := range []bool{false, true} {
		name := "default"
		if volatile {
			name = "volatile"
		}
		t.Run(name, func(t *testing.T) {
			name := fmt.Sprintf("%s_%d", name, textNULRun.Add(1))
			var scalar, step, inverse []string
			MustRegisterFunction("textnul_scalar_"+name, &FunctionImpl{
				NArgs:        1,
				VolatileArgs: volatile,
				Scalar: func(ctx *FunctionContext, args []driver.Value) (driver.Value, error) {
					scalar = append(scalar, textNULKeep(args[0], volatile))
					return int64(len(args[0].(string))), nil
				},
			})
			MustRegisterFunction("textnul_agg_"+name, &FunctionImpl{
				NArgs:        1,
				VolatileArgs: volatile,
				MakeAggregate: func(ctx FunctionContext) (AggregateFunction, error) {
					return &textNULRecorder{&step, &inverse, volatile}, nil
				},
			})

			db, err := sql.Open(driverName, "file::memory:")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			db.SetMaxOpenConns(1)
			if _, err := db.Exec("CREATE TABLE t(i INTEGER PRIMARY KEY, s TEXT)"); err != nil {
				t.Fatal(err)
			}

			for _, s := range textNULValues {
				if _, err := db.Exec("INSERT INTO t(s) VALUES(?)", s); err != nil {
					t.Fatal(err)
				}

				// As a bound parameter the value does not come from a row.
				var n int
				if err := db.QueryRow("SELECT textnul_scalar_"+name+"(?)", s).Scan(&n); err != nil || n != len(s) {
					t.Fatalf("the scalar function returned a length of %d for %q, %v", n, s, err)
				}
			}

			for _, q := range []string{
				"SELECT textnul_scalar_" + name + "(s) FROM t ORDER BY i",
				"SELECT textnul_agg_" + name + "(s) FROM t",
				// A window of one row that moves: every row is stepped in
				// and taken out again but the last.
				"SELECT textnul_agg_" + name + "(s) OVER (ORDER BY i ROWS BETWEEN CURRENT ROW AND CURRENT ROW) FROM t",
			} {
				rows, err := db.Query(q)
				if err != nil {
					t.Fatalf("%s: %v", q, err)
				}

				for rows.Next() {
				}
				if err := rows.Err(); err != nil {
					t.Fatalf("%s: %v", q, err)
				}

				rows.Close()
			}

			twice := append(slices.Clone(textNULValues), textNULValues...)
			if !slices.Equal(scalar, twice) {
				t.Errorf("the scalar function got\n%q\nwant\n%q", scalar, twice)
			}

			if !slices.Equal(step, twice) {
				t.Errorf("Step got\n%q\nwant\n%q", step, twice)
			}

			if want := textNULValues[:len(textNULValues)-1]; !slices.Equal(inverse, want) {
				t.Errorf("WindowInverse got\n%q\nwant\n%q", inverse, want)
			}
		})
	}
}

// TestTextNULVirtualTable checks the same for the arguments of a virtual
// table: the value Cursor.Filter gets for a constraint, and the columns
// Updater.Insert and Updater.Update get. It matters most in Filter: a module
// that takes a constraint for itself, as these do with Omit in BestIndex, is
// not checked by SQLite, so with a value cut at its first NUL it would
// answer with the rows of the shorter one.
func TestTextNULVirtualTable(t *testing.T) {
	for _, volatile := range []bool{false, true} {
		name := "default"
		if volatile {
			name = "volatile"
		}
		t.Run(name, func(t *testing.T) {
			name := fmt.Sprintf("%s_%d", name, textNULRun.Add(1))
			var filtered, inserted, updated []string
			record := func(to *[]string) func([]vtab.Value) {
				return func(vals []vtab.Value) { *to = append(*to, textNULKeep(vals[0], volatile)) }
			}
			if err := vtab.RegisterModule(nil, "textnul_filter_"+name, &volNoopModule{volatile: volatile, recordFilter: record(&filtered)}); err != nil {
				t.Fatal(err)
			}

			stored := &volStoredModule{volatile: volatile, recordInsert: record(&inserted)}
			stored.recordUpdate = func(oldRowid int64, cols []vtab.Value) { record(&updated)(cols) }
			if err := vtab.RegisterModule(nil, "textnul_stored_"+name, stored); err != nil {
				t.Fatal(err)
			}

			db, err := sql.Open(driverName, "file::memory:")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			db.SetMaxOpenConns(1)
			for _, q := range []string{
				"CREATE VIRTUAL TABLE f USING textnul_filter_" + name + "()",
				"CREATE VIRTUAL TABLE u USING textnul_stored_" + name + "()",
			} {
				if _, err := db.Exec(q); err != nil {
					t.Fatalf("%s: %v", q, err)
				}
			}

			for i, s := range textNULValues {
				rows, err := db.Query("SELECT * FROM f WHERE s = ?", s)
				if err != nil {
					t.Fatal(err)
				}

				for rows.Next() {
				}
				if err := rows.Err(); err != nil {
					t.Fatal(err)
				}

				rows.Close()
				if _, err := db.Exec("INSERT INTO u(s, b) VALUES(?, x'00')", s); err != nil {
					t.Fatalf("INSERT of %q: %v", s, err)
				}

				if _, err := db.Exec("UPDATE u SET s = ? WHERE rowid = ?", s+"\x00z", i+1); err != nil {
					t.Fatalf("UPDATE to %q: %v", s+"\x00z", err)
				}
			}

			if !slices.Equal(filtered, textNULValues) {
				t.Errorf("Filter got\n%q\nwant\n%q", filtered, textNULValues)
			}

			if !slices.Equal(inserted, textNULValues) {
				t.Errorf("Insert got\n%q\nwant\n%q", inserted, textNULValues)
			}

			var want []string
			for _, s := range textNULValues {
				want = append(want, s+"\x00z")
			}
			if !slices.Equal(updated, want) {
				t.Errorf("Update got\n%q\nwant\n%q", updated, want)
			}
		})
	}
}
