// Command connector demonstrates sqlite.NewConnector: opening a database
// through database/sql's sql.OpenDB instead of sql.Open, and interposing on
// the physical connections database/sql opens.
//
// The interposing is the reason the API exists. sql.Open reaches the driver
// only by its registered name, so a library wanting to wrap connections has to
// call sql.Register with a name of its own -- which is process-global, panics
// on a name it has already seen, and cannot be undone, so the name must be
// made unique per configuration. sql.OpenDB takes the connector directly and
// registers nothing.
//
// Usage:
//
//	go run ./examples/connector
package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"

	"modernc.org/sqlite"
)

// tracingConnector is the wrapper an instrumentation library writes: it
// embeds the Connector sqlite.NewConnector returned and overrides Connect.
//
// Wrapping the returned driver.Conn as well is possible but wants care: the
// driver's connection implements optional interfaces such as
// driver.QueryerContext and driver.SessionResetter, and a wrapper type that
// does not forward them silently costs database/sql the corresponding fast
// paths. This example wraps only the Connector, which has no such hazard.
type tracingConnector struct {
	driver.Connector

	connects atomic.Int64
}

func (c *tracingConnector) Connect(ctx context.Context) (driver.Conn, error) {
	n := c.connects.Add(1)
	fmt.Printf("  [trace] physical connection #%d\n", n)

	return c.Connector.Connect(ctx)
}

func main() {
	if err := main1(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main1() error {
	// A function on the package-level driver, as an application registers it.
	// Connections handed out by the Connector carry it; connections from a
	// caller-constructed &sqlite.Driver{} would not, because that type's
	// fields are unexported and hold none of the registrations.
	if err := sqlite.RegisterDeterministicScalarFunction("meaning", 0,
		func(*sqlite.FunctionContext, []driver.Value) (driver.Value, error) {
			return int64(42), nil
		},
	); err != nil {
		return err
	}

	dir, err := os.MkdirTemp("", "connector-example-")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	// The dsn syntax is the one documented on sqlite.Driver.Open.
	dsn := "file:" + filepath.Join(dir, "db") + "?_pragma=foreign_keys(1)"

	base, err := sqlite.NewConnector(dsn)
	if err != nil {
		return err
	}

	// No sql.Register, and so no driver name to invent.
	tracer := &tracingConnector{Connector: base}
	db := sql.OpenDB(tracer)

	defer db.Close()

	db.SetMaxOpenConns(2)

	fmt.Printf("sql.OpenDB returned; physical connections so far: %d\n", tracer.connects.Load())

	if _, err := db.Exec(`CREATE TABLE t(i INT)`); err != nil {
		return err
	}

	// Holding two sql.Conn at once forces the pool to open a second
	// connection, so the trace below shows more than one.
	fmt.Println("taking two pooled connections:")
	ctx := context.Background()
	c1, err := db.Conn(ctx)
	if err != nil {
		return err
	}

	defer c1.Close()

	c2, err := db.Conn(ctx)
	if err != nil {
		return err
	}

	defer c2.Close()

	// The registered function is available on connections reached through the
	// Connector, and the PRAGMA in the dsn took effect on each of them.
	for i, c := range []*sql.Conn{c1, c2} {
		var meaning, foreignKeys int
		if err := c.QueryRowContext(ctx, `SELECT meaning()`).Scan(&meaning); err != nil {
			return err
		}

		if err := c.QueryRowContext(ctx, `PRAGMA foreign_keys`).Scan(&foreignKeys); err != nil {
			return err
		}

		fmt.Printf("conn %d: meaning()=%d foreign_keys=%d\n", i+1, meaning, foreignKeys)
	}

	// db.Driver reports the driver registered as "sqlite", the same value
	// sql.Open("sqlite", dsn) would have produced.
	fmt.Printf("db.Driver() is %T\n", db.Driver())
	fmt.Printf("total physical connections: %d\n", tracer.connects.Load())

	return nil
}
