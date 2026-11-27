// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"context"
	"database/sql/driver"
	"errors"
	"path/filepath"
	"testing"
)

type closeTestConn interface {
	driver.Conn
	driver.ExecerContext
	driver.QueryerContext
	driver.ConnPrepareContext
}

func closeTestConnect(t *testing.T, c driver.Connector) closeTestConn {
	t.Helper()
	conn, err := c.Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	return conn.(closeTestConn)
}

func closeTestExec(t *testing.T, c closeTestConn, sql string) {
	t.Helper()
	if _, err := c.ExecContext(context.Background(), sql, nil); err != nil {
		t.Fatalf("%s: %v", sql, err)
	}
}

// TestCloseFinalizesStatements: a statement that is never closed, used
// through the driver interfaces directly, no longer keeps its database alive
// after the connection is closed (https://github.com/modernc-org/sqlite/issues/16).
// Before, sqlite3_close_v2 left the connection a zombie that held its
// transaction, and with it the lock that keeps another connection from
// writing, until the process ended; closing the statement afterwards
// panicked.
func TestCloseFinalizesStatements(t *testing.T) {
	ctx := context.Background()
	con, err := NewConnector("file:" + filepath.Join(t.TempDir(), "test.db") + "?_pragma=busy_timeout(0)")
	if err != nil {
		t.Fatal(err)
	}

	c1 := closeTestConnect(t, con)
	closeTestExec(t, c1, "CREATE TABLE t(i)")
	closeTestExec(t, c1, "INSERT INTO t VALUES(1), (2)")
	// FTS5 prepares statements of its own on the handle and finalizes them
	// when it is disconnected: Close must leave those alone.
	closeTestExec(t, c1, "CREATE VIRTUAL TABLE f USING fts5(x)")
	closeTestExec(t, c1, "INSERT INTO f VALUES('alpha beta'), ('beta gamma')")
	closeTestExec(t, c1, "BEGIN IMMEDIATE")
	closeTestExec(t, c1, "INSERT INTO t VALUES(3)")
	closeTestExec(t, c1, "INSERT INTO f VALUES('gamma delta')")

	// A prepared statement run once and never closed, and the rows of a
	// query left in the middle.
	stmt, err := c1.PrepareContext(ctx, "SELECT count(*) FROM t")
	if err != nil {
		t.Fatal(err)
	}

	rows, err := stmt.(driver.StmtQueryContext).QueryContext(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	dest := make([]driver.Value, 1)
	if err := rows.Next(dest); err != nil {
		t.Fatal(err)
	}

	if err := rows.Close(); err != nil {
		t.Fatal(err)
	}

	open, err := c1.QueryContext(ctx, "SELECT i FROM t", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := open.Next(dest); err != nil {
		t.Fatal(err)
	}

	// And a query of the virtual table left in the middle, its cursor open.
	match, err := c1.QueryContext(ctx, "SELECT x FROM f WHERE f MATCH 'beta'", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := match.Next(dest); err != nil {
		t.Fatal(err)
	}

	if err := c1.Close(); err != nil {
		t.Fatal(err)
	}

	// The transaction ended with the connection: another one writes, and
	// the row it did not commit is gone.
	c2 := closeTestConnect(t, con)
	defer c2.Close()
	closeTestExec(t, c2, "BEGIN IMMEDIATE")
	closeTestExec(t, c2, "INSERT INTO t VALUES(4)")
	closeTestExec(t, c2, "COMMIT")
	got, err := c2.QueryContext(ctx, "SELECT count(*), sum(i) FROM t", nil)
	if err != nil {
		t.Fatal(err)
	}

	dest = make([]driver.Value, 2)
	if err := got.Next(dest); err != nil {
		t.Fatal(err)
	}

	if dest[0] != int64(3) || dest[1] != int64(7) {
		t.Errorf("count, sum = %v, %v, want 3, 7", dest[0], dest[1])
	}

	// Rows closed twice on a live connection: the second Close does nothing.
	if err := got.Close(); err != nil {
		t.Fatal(err)
	}

	if err := got.Close(); err != nil {
		t.Fatalf("second Close: %v", err)
	}

	// What was left on the closed connection reports it and closes without
	// touching the statements it had.
	if _, err := stmt.(driver.StmtQueryContext).QueryContext(ctx, nil); !errors.Is(err, driver.ErrBadConn) {
		t.Errorf("Query after conn.Close: %v, want %v", err, driver.ErrBadConn)
	}

	if _, err := stmt.(driver.StmtExecContext).ExecContext(ctx, nil); !errors.Is(err, driver.ErrBadConn) {
		t.Errorf("Exec after conn.Close: %v, want %v", err, driver.ErrBadConn)
	}

	if err := open.Next(dest[:1]); !errors.Is(err, driver.ErrBadConn) {
		t.Errorf("Next after conn.Close: %v, want %v", err, driver.ErrBadConn)
	}

	ct := open.(driver.RowsColumnTypeLength)
	if _, ok := ct.ColumnTypeLength(0); ok {
		t.Error("ColumnTypeLength after conn.Close: ok")
	}

	_ = open.(driver.RowsColumnTypeScanType).ColumnTypeScanType(0)
	for i := 0; i < 2; i++ {
		if err := open.Close(); err != nil {
			t.Errorf("rows.Close %d after conn.Close: %v", i, err)
		}

		if err := match.Close(); err != nil {
			t.Errorf("rows.Close %d of the FTS5 query after conn.Close: %v", i, err)
		}

		if err := stmt.Close(); err != nil {
			t.Errorf("stmt.Close %d after conn.Close: %v", i, err)
		}
	}
}
