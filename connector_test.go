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
