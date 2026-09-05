// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"database/sql"
	"database/sql/driver"
	"testing"
)

// ctxProbe records what the *FunctionContext handed to each user callback
// looked like. The driver hands out a pooled context whose tls and ctx fields
// are populated for the duration of the call; a nil pointer or zero fields
// would mean a trampoline stopped doing so.
type ctxProbe struct {
	scalar, step, inverse, value, final int
	bad                                 []string
}

var theCtxProbe = &ctxProbe{}

func (p *ctxProbe) check(name string, ctx *FunctionContext) {
	if ctx == nil || ctx.tls == nil || ctx.ctx == 0 {
		p.bad = append(p.bad, name)
	}
}

type ctxProbeSum struct{ sum int64 }

func (a *ctxProbeSum) Step(ctx *FunctionContext, args []driver.Value) error {
	theCtxProbe.step++
	theCtxProbe.check("Step", ctx)
	a.sum += args[0].(int64)
	return nil
}

func (a *ctxProbeSum) WindowInverse(ctx *FunctionContext, args []driver.Value) error {
	theCtxProbe.inverse++
	theCtxProbe.check("WindowInverse", ctx)
	a.sum -= args[0].(int64)
	return nil
}

func (a *ctxProbeSum) WindowValue(ctx *FunctionContext) (driver.Value, error) {
	theCtxProbe.value++
	theCtxProbe.check("WindowValue", ctx)
	return a.sum, nil
}

func (a *ctxProbeSum) Final(ctx *FunctionContext) {
	theCtxProbe.final++
	theCtxProbe.check("Final", ctx)
}

func init() {
	MustRegisterDeterministicScalarFunction(
		"ctxprobe_scalar",
		1,
		func(ctx *FunctionContext, args []driver.Value) (driver.Value, error) {
			theCtxProbe.scalar++
			theCtxProbe.check("Scalar", ctx)
			return args[0], nil
		},
	)
	MustRegisterFunction("ctxprobe_sum", &FunctionImpl{
		NArgs:         1,
		Deterministic: true,
		MakeAggregate: func(ctx FunctionContext) (AggregateFunction, error) {
			return &ctxProbeSum{}, nil
		},
	})
}

// TestFunctionContextPerCall checks that every user callback (scalar, Step,
// WindowInverse, WindowValue, Final) receives a non-nil *FunctionContext whose
// tls and ctx fields refer to the current invocation, now that the driver
// hands out a pooled context instead of a fresh zero value per call.
func TestFunctionContextPerCall(t *testing.T) {
	theCtxProbe = &ctxProbe{}
	db, err := sql.Open(driverName, "file::memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE t (a INTEGER)`); err != nil {
		t.Fatal(err)
	}
	for i := int64(1); i <= 5; i++ {
		if _, err := db.Exec(`INSERT INTO t VALUES (?)`, i); err != nil {
			t.Fatal(err)
		}
	}

	var n int64
	if err := db.QueryRow(`SELECT sum(ctxprobe_scalar(a)) FROM t`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 15 {
		t.Fatalf("scalar: got %d, want 15", n)
	}

	if err := db.QueryRow(`SELECT ctxprobe_sum(a) FROM t`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 15 {
		t.Fatalf("aggregate: got %d, want 15", n)
	}

	rows, err := db.Query(`SELECT ctxprobe_sum(a) OVER (ORDER BY a ROWS BETWEEN 1 PRECEDING AND CURRENT ROW) FROM t`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var got []int64
	for rows.Next() {
		if err := rows.Scan(&n); err != nil {
			t.Fatal(err)
		}
		got = append(got, n)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	want := []int64{1, 3, 5, 7, 9}
	if len(got) != len(want) {
		t.Fatalf("window: got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("window: got %v, want %v", got, want)
		}
	}

	p := theCtxProbe
	if p.scalar != 5 || p.step < 5 || p.inverse < 1 || p.value < 1 || p.final < 1 {
		t.Fatalf("callback counts: scalar=%d step=%d inverse=%d value=%d final=%d", p.scalar, p.step, p.inverse, p.value, p.final)
	}
	if len(p.bad) != 0 {
		t.Fatalf("callbacks that received a nil or unpopulated *FunctionContext: %v", p.bad)
	}
}
