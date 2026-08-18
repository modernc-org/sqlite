// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"modernc.org/sqlite/vtab"
)

// reachModule reports that it was reached by failing Create/Connect with a
// recognisable message, which is enough to tell which implementation a
// CREATE VIRTUAL TABLE resolved to without building a working table.
type reachModule struct{ tag string }

func (m reachModule) Create(vtab.Context, []string) (vtab.Table, error) {
	return nil, errors.New("reached:" + m.tag)
}

func (m reachModule) Connect(vtab.Context, []string) (vtab.Table, error) {
	return nil, errors.New("reached:" + m.tag)
}

var driverRegisterSeq atomic.Int64

// uniqueDriverName keeps sql.Register, which panics on a duplicate name, safe
// across tests and repeated -count runs.
func uniqueDriverName(t *testing.T) string {
	t.Helper()
	return fmt.Sprintf("sqlite-test-%s-%d", t.Name(), driverRegisterSeq.Add(1))
}

// openOn registers d under a fresh name and opens an in-memory database on it.
func openOn(t *testing.T, d *Driver) *sql.DB {
	t.Helper()
	name := uniqueDriverName(t)
	sql.Register(name, d)
	db, err := sql.Open(name, "file::memory:")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

// TestDriverRegisterOwnFunctions verifies that a caller-constructed Driver can
// be filled in with its own functions and collations, and that a zero-value
// Driver is usable as-is.
func TestDriverRegisterOwnFunctions(t *testing.T) {
	d := &Driver{} // deliberately zero: the maps must be created on demand

	if err := d.RegisterDeterministicScalarFunction("own_det", 0,
		func(*FunctionContext, []driver.Value) (driver.Value, error) { return int64(42), nil }); err != nil {
		t.Fatal(err)
	}
	if err := d.RegisterScalarFunction("own_scalar", 1,
		func(_ *FunctionContext, args []driver.Value) (driver.Value, error) { return args[0], nil }); err != nil {
		t.Fatal(err)
	}
	if err := d.RegisterFunction("own_impl", &FunctionImpl{
		NArgs:         0,
		Deterministic: true,
		Scalar:        func(*FunctionContext, []driver.Value) (driver.Value, error) { return int64(7), nil },
	}); err != nil {
		t.Fatal(err)
	}
	if err := d.RegisterCollationUtf8("own_coll", strings.Compare); err != nil {
		t.Fatal(err)
	}

	db := openOn(t, d)

	var n int64
	if err := db.QueryRow(`SELECT own_det()`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 42 {
		t.Fatalf("own_det() = %d, want 42", n)
	}
	if err := db.QueryRow(`SELECT own_scalar(5)`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatalf("own_scalar(5) = %d, want 5", n)
	}
	if err := db.QueryRow(`SELECT own_impl()`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 7 {
		t.Fatalf("own_impl() = %d, want 7", n)
	}
	if _, err := db.Exec(`CREATE TABLE t(x TEXT COLLATE own_coll)`); err != nil {
		t.Fatal(err)
	}
}

// TestDriverRegisterDuplicate verifies the per-Driver registrations reject a
// duplicate name the same way the package-level ones do.
func TestDriverRegisterDuplicate(t *testing.T) {
	d := &Driver{}
	fn := func(*FunctionContext, []driver.Value) (driver.Value, error) { return nil, nil }

	if err := d.RegisterScalarFunction("dup_fn", 0, fn); err != nil {
		t.Fatal(err)
	}
	if err := d.RegisterScalarFunction("dup_fn", 0, fn); err == nil {
		t.Fatal("second RegisterScalarFunction with the same name: got nil error")
	}
	if err := d.RegisterCollationUtf8("dup_coll", strings.Compare); err != nil {
		t.Fatal(err)
	}
	if err := d.RegisterCollationUtf8("dup_coll", strings.Compare); err == nil {
		t.Fatal("second RegisterCollationUtf8 with the same name: got nil error")
	}
	if err := d.RegisterModule("dup_mod", reachModule{"a"}); err != nil {
		t.Fatal(err)
	}
	if err := d.RegisterModule("dup_mod", reachModule{"b"}); err == nil {
		t.Fatal("second RegisterModule with the same name: got nil error")
	}
}

// TestDriverRegisterIsolation verifies that what one constructed Driver
// registers does not leak into another, nor into the package-level driver.
func TestDriverRegisterIsolation(t *testing.T) {
	a := &Driver{}
	b := &Driver{}
	if err := a.RegisterDeterministicScalarFunction("only_on_a", 0,
		func(*FunctionContext, []driver.Value) (driver.Value, error) { return int64(1), nil }); err != nil {
		t.Fatal(err)
	}
	if err := a.RegisterModule("mod_only_on_a", reachModule{"a"}); err != nil {
		t.Fatal(err)
	}

	dbA := openOn(t, a)
	dbB := openOn(t, b)

	var n int64
	if err := dbA.QueryRow(`SELECT only_on_a()`).Scan(&n); err != nil {
		t.Fatalf("function on its own driver: %v", err)
	}
	if err := dbB.QueryRow(`SELECT only_on_a()`).Scan(&n); err == nil {
		t.Fatal("function registered on driver a resolved on driver b")
	}

	if _, err := dbA.Exec(`CREATE VIRTUAL TABLE va USING mod_only_on_a()`); err == nil ||
		!strings.Contains(err.Error(), "reached:a") {
		t.Fatalf("module on its own driver: %v, want it to be reached", err)
	}
	if _, err := dbB.Exec(`CREATE VIRTUAL TABLE vb USING mod_only_on_a()`); err == nil ||
		strings.Contains(err.Error(), "reached:") {
		t.Fatalf("module registered on driver a resolved on driver b: %v", err)
	}

	// The package-level driver must not have picked either of them up.
	if _, ok := defaultDriver().udfs["only_on_a"]; ok {
		t.Fatal("per-driver function leaked into the package-level driver")
	}
	if _, ok := defaultDriver().modules["mod_only_on_a"]; ok {
		t.Fatal("per-driver module leaked into the package-level driver")
	}
}

// TestDriverRegisterSameModuleName verifies isolation where it can actually
// fail: two Drivers registering the same module name with different
// implementations. The module ID handed to sqlite3_create_module_v2 as pAux is
// what every trampoline dispatches on, so an ID allocated per name rather than
// per (Driver, name) would make the drivers share one dispatch entry, with the
// last registration winning process-wide -- including on connections that were
// already open. The connection pinned before b's first open is the case that
// catches the retroactive overwrite; TestDriverRegisterIsolation cannot see
// any of this, because its two drivers use two different module names.
func TestDriverRegisterSameModuleName(t *testing.T) {
	a := &Driver{}
	b := &Driver{}
	const modName = "same_name_mod"
	if err := a.RegisterModule(modName, reachModule{"a"}); err != nil {
		t.Fatal(err)
	}
	if err := b.RegisterModule(modName, reachModule{"b"}); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	dbA := openOn(t, a)
	connA, err := dbA.Conn(ctx) // pin one of a's connections before b opens any
	if err != nil {
		t.Fatal(err)
	}
	defer connA.Close()

	reachedBy := func(err error, tag string) bool {
		return err != nil && strings.Contains(err.Error(), "reached:"+tag)
	}

	if _, err := connA.ExecContext(ctx, `CREATE VIRTUAL TABLE v1 USING same_name_mod()`); !reachedBy(err, "a") {
		t.Fatalf("driver a before b opened: %v, want reached:a", err)
	}

	dbB := openOn(t, b)
	if _, err := dbB.Exec(`CREATE VIRTUAL TABLE v2 USING same_name_mod()`); !reachedBy(err, "b") {
		t.Fatalf("driver b: %v, want reached:b", err)
	}

	// The pinned connection must still dispatch to a's implementation; with a
	// per-name ID it answers with b's.
	if _, err := connA.ExecContext(ctx, `CREATE VIRTUAL TABLE v3 USING same_name_mod()`); !reachedBy(err, "a") {
		t.Fatalf("driver a's pre-existing connection after b opened: %v, want reached:a", err)
	}

	// And so must a connection a opens from here on.
	if _, err := dbA.Exec(`CREATE VIRTUAL TABLE v4 USING same_name_mod()`); !reachedBy(err, "a") {
		t.Fatalf("driver a after b opened: %v, want reached:a", err)
	}
}

// TestDriverRegisterConcurrent verifies the documented claim that a Driver's
// registration methods are safe to call concurrently with each other,
// including the very first registrations, which create the maps.
func TestDriverRegisterConcurrent(t *testing.T) {
	d := &Driver{}
	const n = 8
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			fn := func(*FunctionContext, []driver.Value) (driver.Value, error) { return int64(i), nil }
			if err := d.RegisterScalarFunction(fmt.Sprintf("conc_fn_%d", i), 0, fn); err != nil {
				t.Error(err)
			}
			if err := d.RegisterCollationUtf8(fmt.Sprintf("conc_coll_%d", i), strings.Compare); err != nil {
				t.Error(err)
			}
			if err := d.RegisterModule(fmt.Sprintf("conc_mod_%d", i), reachModule{"c"}); err != nil {
				t.Error(err)
			}
			d.RegisterConnectionHook(func(ExecQuerierContext, string) error { return nil })
		}(i)
	}
	wg.Wait()

	if got := len(d.udfs); got != n {
		t.Fatalf("got %d functions, want %d", got, n)
	}
	if got := len(d.collations); got != n {
		t.Fatalf("got %d collations, want %d", got, n)
	}
	if got := len(d.modules); got != n {
		t.Fatalf("got %d modules, want %d", got, n)
	}
	if got := len(d.connectionHooks); got != n {
		t.Fatalf("got %d connection hooks, want %d", got, n)
	}
}

// TestDriverGlobalModulesStillApply pins the behavior that must not change:
// modules registered through the package-level path keep reaching connections
// opened by a caller-constructed Driver, as they have since module support was
// added. See https://gitlab.com/cznic/sqlite/-/issues/254.
func TestDriverGlobalModulesStillApply(t *testing.T) {
	const name = "global_reach_mod"
	if err := vtab.RegisterModule(nil, name, reachModule{"global"}); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		g := defaultDriver()
		g.mu.Lock()
		defer g.mu.Unlock()

		delete(g.modules, name)
	})

	db := openOn(t, &Driver{})
	if _, err := db.Exec(`CREATE VIRTUAL TABLE vg USING ` + name + `()`); err == nil ||
		!strings.Contains(err.Error(), "reached:global") {
		t.Fatalf("globally registered module on a constructed Driver: %v, want it to be reached", err)
	}
}

// TestVtabRegisterModuleHonoursDB verifies that the db argument of
// vtab.RegisterModule now selects the driver the module lands on, and that a
// nil db still targets the package-level driver.
func TestVtabRegisterModuleHonoursDB(t *testing.T) {
	own := &Driver{}
	other := &Driver{}
	dbOwn := openOn(t, own)
	dbOther := openOn(t, other)

	if err := vtab.RegisterModule(dbOwn, "via_db_mod", reachModule{"viadb"}); err != nil {
		t.Fatal(err)
	}
	if _, ok := own.modules["via_db_mod"]; !ok {
		t.Fatal("vtab.RegisterModule with a non-nil db did not register on that db's driver")
	}
	if _, ok := defaultDriver().modules["via_db_mod"]; ok {
		t.Fatal("vtab.RegisterModule with a non-nil db also registered on the package-level driver")
	}

	// Reaching it requires a connection opened after registration.
	if _, err := dbOwn.Exec(`CREATE VIRTUAL TABLE v1 USING via_db_mod()`); err == nil ||
		!strings.Contains(err.Error(), "reached:viadb") {
		t.Fatalf("module via db argument: %v, want it to be reached", err)
	}
	if _, err := dbOther.Exec(`CREATE VIRTUAL TABLE v2 USING via_db_mod()`); err == nil ||
		strings.Contains(err.Error(), "reached:") {
		t.Fatalf("module leaked to an unrelated driver: %v", err)
	}
}

// TestVtabRegisterModuleValidates keeps the argument checks in front of the
// driver lookup, so a bad call fails the same way whether db is nil or not.
func TestVtabRegisterModuleValidates(t *testing.T) {
	db := openOn(t, &Driver{})
	for _, tc := range []struct {
		name string
		db   *sql.DB
		mod  vtab.Module
		want string
	}{
		{"", nil, reachModule{"x"}, "non-empty"},
		{"", db, reachModule{"x"}, "non-empty"},
		{"nil_mod_nil_db", nil, nil, "nil"},
		{"nil_mod_with_db", db, nil, "nil"},
	} {
		err := vtab.RegisterModule(tc.db, tc.name, tc.mod)
		if err == nil || !strings.Contains(err.Error(), tc.want) {
			t.Fatalf("RegisterModule(%v, %q, %v) = %v, want an error containing %q",
				tc.db != nil, tc.name, tc.mod, err, tc.want)
		}
	}
}
