// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	sqlite3 "modernc.org/sqlite/lib"
)

// The registrations below go on the package-level driver and therefore affect
// every connection the test binary opens afterwards. They are named uniquely
// and the connection hook is inert unless the DSN carries connHookMarker, so
// they cannot perturb the rest of the suite.
const connHookMarker = "connector_test_hook_marker"

var connHookCalls int64

func init() {
	MustRegisterDeterministicScalarFunction("connector_test_answer", 0,
		func(*FunctionContext, []driver.Value) (driver.Value, error) {
			return int64(42), nil
		})
	MustRegisterCollationUtf8("connector_test_collation",
		func(left, right string) int { return strings.Compare(left, right) })
	RegisterConnectionHook(func(_ ExecQuerierContext, dsn string) error {
		if strings.Contains(dsn, connHookMarker) {
			atomic.AddInt64(&connHookCalls, 1)
		}
		return nil
	})
}

// TestConnectorOpenDB exercises the sql.OpenDB path end to end.
func TestConnectorOpenDB(t *testing.T) {
	c, err := NewConnector("file::memory:")
	if err != nil {
		t.Fatalf("NewConnector: %v", err)
	}

	db := sql.OpenDB(c)
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE t(i INT); INSERT INTO t VALUES(1), (2)`); err != nil {
		t.Fatalf("Exec: %v", err)
	}

	var n int
	if err := db.QueryRow(`SELECT sum(i) FROM t`).Scan(&n); err != nil {
		t.Fatalf("QueryRow: %v", err)
	}
	if g, e := n, 3; g != e {
		t.Fatalf("got %v, expected %v", g, e)
	}
}

// TestConnectorAppliesGlobalRegistrations is the point of
// https://gitlab.com/cznic/sqlite/-/issues/253: connections handed out by the
// Connector must carry the functions, collations and hooks registered on the
// package-level driver, which a caller-constructed &Driver{} does not.
func TestConnectorAppliesGlobalRegistrations(t *testing.T) {
	before := atomic.LoadInt64(&connHookCalls)

	c, err := NewConnector("file::memory:?_pragma=application_id(1)&x=" + connHookMarker)
	if err != nil {
		t.Fatalf("NewConnector: %v", err)
	}

	db := sql.OpenDB(c)
	defer db.Close()

	var n int
	if err := db.QueryRow(`SELECT connector_test_answer()`).Scan(&n); err != nil {
		t.Fatalf("registered function not available: %v", err)
	}
	if g, e := n, 42; g != e {
		t.Fatalf("got %v, expected %v", g, e)
	}

	var b bool
	if err := db.QueryRow(`SELECT 'a' < 'b' COLLATE connector_test_collation`).Scan(&b); err != nil {
		t.Fatalf("registered collation not available: %v", err)
	}
	if !b {
		t.Fatal("collation returned an unexpected ordering")
	}

	if g := atomic.LoadInt64(&connHookCalls); g <= before {
		t.Fatalf("connection hook not called: %v, was %v", g, before)
	}
}

// TestConnectorDriverIsRegisteredDriver checks that the Connector reports the
// same driver value database/sql hands out for sql.Open("sqlite", ...), which
// is what makes it a drop-in for code reaching the driver through db.Driver().
func TestConnectorDriverIsRegisteredDriver(t *testing.T) {
	viaOpen, err := sql.Open("sqlite", "file::memory:")
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	defer viaOpen.Close()

	c, err := NewConnector("file::memory:")
	if err != nil {
		t.Fatalf("NewConnector: %v", err)
	}

	viaConnector := sql.OpenDB(c)
	defer viaConnector.Close()

	if g, e := viaConnector.Driver(), viaOpen.Driver(); g != e {
		t.Fatalf("got %p, expected %p", g, e)
	}
	if _, ok := viaConnector.Driver().(*Driver); !ok {
		t.Fatalf("got %T, expected *sqlite.Driver", viaConnector.Driver())
	}
}

// countingConnector is the wrapper an instrumentation library writes. Note
// that it needs no sql.Register and therefore no globally unique driver name.
type countingConnector struct {
	driver.Connector

	connects int64
}

func (c *countingConnector) Connect(ctx context.Context) (driver.Conn, error) {
	atomic.AddInt64(&c.connects, 1)
	return c.Connector.Connect(ctx)
}

// TestConnectorWrapping covers the use case the Connector exists for:
// interposing on the physical connections database/sql opens.
func TestConnectorWrapping(t *testing.T) {
	base, err := NewConnector("file:" + filepath.Join(t.TempDir(), "wrap.db"))
	if err != nil {
		t.Fatalf("NewConnector: %v", err)
	}

	wrapped := &countingConnector{Connector: base}
	db := sql.OpenDB(wrapped)
	defer db.Close()
	db.SetMaxOpenConns(2)

	if g := atomic.LoadInt64(&wrapped.connects); g != 0 {
		t.Fatalf("sql.OpenDB connected eagerly: %v", g)
	}

	// Holding two sql.Conn at once forces exactly two physical connections.
	ctx := context.Background()
	c1, err := db.Conn(ctx)
	if err != nil {
		t.Fatalf("Conn: %v", err)
	}
	defer c1.Close()

	c2, err := db.Conn(ctx)
	if err != nil {
		t.Fatalf("Conn: %v", err)
	}
	defer c2.Close()

	if g, e := atomic.LoadInt64(&wrapped.connects), int64(2); g != e {
		t.Fatalf("got %v physical connects, expected %v", g, e)
	}

	// The wrapper must not have cost us the global registrations.
	var n int
	if err := c1.QueryRowContext(ctx, `SELECT connector_test_answer()`).Scan(&n); err != nil {
		t.Fatalf("registered function not available through wrapper: %v", err)
	}
	if g, e := n, 42; g != e {
		t.Fatalf("got %v, expected %v", g, e)
	}
}

// unwrappingConn is the connection wrapper the NewConnector documentation
// asks for: it offers the wrapped connection through Unwrap, so a helper that
// needs the driver's own connection, such as Limit, reaches it. It forwards
// none of the optional interfaces, which this test does not need.
type unwrappingConn struct {
	driver.Conn
}

func (c unwrappingConn) Unwrap() driver.Conn { return c.Conn }

// opaqueConn is a wrapper without Unwrap. Limit cannot see through it.
type opaqueConn struct {
	driver.Conn
}

// nilUnwrapConn is a broken wrapper whose Unwrap returns nothing.
type nilUnwrapConn struct {
	driver.Conn
}

func (nilUnwrapConn) Unwrap() driver.Conn { return nil }

// wrappingConnector wraps every connection it opens with wrap.
type wrappingConnector struct {
	driver.Connector

	wrap func(driver.Conn) driver.Conn
}

func (c *wrappingConnector) Connect(ctx context.Context) (driver.Conn, error) {
	conn, err := c.Connector.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return c.wrap(conn), nil
}

// TestLimitWrappedConn checks that Limit reaches the driver connection through
// a wrapper that has an Unwrap method, and through two of them, and that it
// still reports a wrapper it cannot see through instead of panicking.
func TestLimitWrappedConn(t *testing.T) {
	base, err := NewConnector("file:" + filepath.Join(t.TempDir(), "limit.db"))
	if err != nil {
		t.Fatalf("NewConnector: %v", err)
	}

	for _, tc := range []struct {
		name string
		wrap func(driver.Conn) driver.Conn
		ok   bool
	}{
		{"unwrap", func(c driver.Conn) driver.Conn { return unwrappingConn{c} }, true},
		{"unwrap twice", func(c driver.Conn) driver.Conn { return unwrappingConn{unwrappingConn{c}} }, true},
		{"opaque", func(c driver.Conn) driver.Conn { return opaqueConn{c} }, false},
		{"opaque inside unwrap", func(c driver.Conn) driver.Conn { return unwrappingConn{opaqueConn{c}} }, false},
		{"unwrap to nil", func(c driver.Conn) driver.Conn { return nilUnwrapConn{c} }, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			db := sql.OpenDB(&wrappingConnector{Connector: base, wrap: tc.wrap})
			defer db.Close()

			conn, err := db.Conn(context.Background())
			if err != nil {
				t.Fatalf("Conn: %v", err)
			}

			defer conn.Close()

			n, err := Limit(conn, sqlite3.SQLITE_LIMIT_COLUMN, 100)
			if !tc.ok {
				if err == nil || !strings.Contains(err.Error(), "unexpected driverConn type") {
					t.Fatalf("got err %v, expected the unexpected driverConn type error", err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Limit: %v", err)
			}

			if n <= 0 {
				t.Fatalf("got previous limit %v, expected a positive value", n)
			}

			// The call reached the driver connection: the new limit is in
			// force on it.
			if n, err = Limit(conn, sqlite3.SQLITE_LIMIT_COLUMN, -1); err != nil {
				t.Fatalf("Limit: %v", err)
			}

			if g, e := n, 100; g != e {
				t.Fatalf("got limit %v, expected %v", g, e)
			}
		})
	}
}

// TestConnectorRejectsMalformedQuery covers the eager half of the validation
// contract: what can be rejected without opening a database is rejected at
// construction. That is a query string that does not parse, and conflicting
// vfs parameters.
func TestConnectorRejectsMalformedQuery(t *testing.T) {
	for _, dsn := range []string{
		"file::memory:?_pragma=%zz",
		"file::memory:?%",
		"file::memory:?a=%2",
		"file::memory:?vfs=a&vfs=b",
	} {
		c, err := NewConnector(dsn)
		if err == nil {
			t.Errorf("%q: got a Connector, expected an error", dsn)
			continue
		}
		if c != nil {
			t.Errorf("%q: got a non-nil Connector alongside %v", dsn, err)
		}
	}
}

// TestConnectorDefersValueValidation covers the lazy half: parameter values
// are checked when the connection is opened, not by NewConnector.
func TestConnectorDefersValueValidation(t *testing.T) {
	const dsn = "file::memory:?_txlock=bogus"

	c, err := NewConnector(dsn)
	if err != nil {
		t.Fatalf("NewConnector rejected %q eagerly: %v", dsn, err)
	}

	db := sql.OpenDB(c)
	defer db.Close()

	if err := db.Ping(); err == nil {
		t.Fatal("Ping succeeded, expected the bad _txlock to be reported")
	} else if !strings.Contains(err.Error(), "_txlock") {
		t.Fatalf("got %v, expected it to mention _txlock", err)
	}
}

// TestConnectorContextCanceled checks Connect reports an already-cancelled
// context rather than opening a connection.
func TestConnectorContextCanceled(t *testing.T) {
	c, err := NewConnector("file::memory:")
	if err != nil {
		t.Fatalf("NewConnector: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	conn, err := c.Connect(ctx)
	if err == nil {
		conn.Close()
		t.Fatal("Connect succeeded on a cancelled context")
	}
	if g, e := err, context.Canceled; g != e {
		t.Fatalf("got %v, expected %v", g, e)
	}
}

// TestConnectorDSNSplitMatchesOpen guards the invariant dsnQuery's comment
// relies on: NewConnector must never reject a dsn newConn would have accepted.
func TestConnectorDSNSplitMatchesOpen(t *testing.T) {
	for _, v := range []struct {
		dsn   string
		query string
	}{
		{"", ""},
		{":memory:", ""},
		{"file::memory:", ""},
		{"file::memory:?", ""},
		{"file::memory:?a=b", "a=b"},
		{"file::memory:?a=b&c=d", "a=b&c=d"},
		{"?a=b", ""}, // A '?' in the first position is part of the filename.
		{"x?a=b?c=d", "a=b?c=d"},
	} {
		if g, e := dsnQuery(v.dsn), v.query; g != e {
			t.Errorf("dsnQuery(%q): got %q, expected %q", v.dsn, g, e)
		}
	}

	// Whatever NewConnector rejects, opening must reject too. Every dsn here
	// is memory-backed, so a successful open touches no file.
	for _, dsn := range []string{
		"file::memory:",
		"file::memory:?",
		"file::memory:?_pragma=%zz",
		"file::memory:?%",
		"file::memory:?vfs=a&vfs=b",
		"file::memory:?_error_rc=maybe",
	} {
		_, cErr := NewConnector(dsn)
		if cErr == nil {
			continue
		}

		conn, oErr := d.Open(dsn)
		if oErr == nil {
			conn.Close()
			t.Errorf("%q: NewConnector rejected it (%v) but Open accepted it", dsn, cErr)
		}
	}
}
