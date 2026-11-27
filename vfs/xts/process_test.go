// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

// The tests below run this test binary again as other processes on one
// database. What such a process does is in the environment.
const (
	childRole = "XTS_TEST_ROLE"
	childPath = "XTS_TEST_PATH"
	childKey  = "XTS_TEST_KEY"
)

func child(t *testing.T, role, path string, key byte) *exec.Cmd {
	t.Helper()
	cmd := exec.Command(os.Args[0], "-test.run=^TestChildProcess$", "-test.v")
	cmd.Env = append(os.Environ(), childRole+"="+role, childPath+"="+path, childKey+"="+strconv.Itoa(int(key)))
	cmd.Stderr = os.Stderr
	return cmd
}

// TestChildProcess is the other process. It does nothing unless it is asked.
func TestChildProcess(t *testing.T) {
	role := os.Getenv(childRole)
	if role == "" {
		t.Skip("run by the tests that need another process")
	}

	k, err := strconv.Atoi(os.Getenv(childKey))
	if err != nil {
		t.Fatal(err)
	}

	vfs := register(t, testKey(byte(k)))
	db := open(t, os.Getenv(childPath), vfs, "_pragma=cache_size(10)")
	db.SetMaxOpenConns(1)
	switch role {
	case "die in a transaction":
		tx, err := db.Begin()
		if err != nil {
			t.Fatal(err)
		}

		if _, err := tx.Exec("create table u(x)"); err != nil {
			t.Fatal(err)
		}

		if _, err := tx.Exec("update t set b = 'changed ' || b"); err != nil {
			t.Fatal(err)
		}

		fmt.Println("READY")
		time.Sleep(time.Minute) // killed before that
	case "insert":
		for i := range 200 {
			mustExec(t, db, "insert into t(b) values(?)", fmt.Sprintf("%d from %d", i, os.Getpid()))
		}
	default:
		t.Fatalf("role %q", role)
	}
}

// A writer is killed in the middle of a transaction that has already put
// pages into the database file. The next process with the key rolls it back.
func TestKilledWriter(t *testing.T) {
	const key = 50
	vfs := register(t, testKey(key))
	dir := t.TempDir()
	path := filepath.Join(dir, "a.db")
	db := open(t, path, vfs, "")
	mustExec(t, db, "create table t(a integer primary key, b text)")
	mustExec(t, db, fmt.Sprintf(fill, 2000)+"insert into t(b) select ? || v from s", marker)
	db.Close()

	cmd := child(t, "die in a transaction", path, key)
	out, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	ready := make(chan bool, 1)
	go func() {
		sc := bufio.NewScanner(out)
		for sc.Scan() {
			if sc.Text() == "READY" {
				ready <- true
				return
			}
		}

		ready <- false
	}()

	select {
	case ok := <-ready:
		if !ok {
			cmd.Wait()
			t.Fatal("the writer ended before it was ready")
		}
	case <-time.After(time.Minute):
		cmd.Process.Kill()
		t.Fatal("the writer did not get ready")
	}

	if err := cmd.Process.Kill(); err != nil {
		t.Fatal(err)
	}

	cmd.Wait()
	if _, err := os.Stat(path + "-journal"); err != nil {
		t.Fatalf("the killed writer left no journal: %v", err)
	}

	noMarker(t, dir)
	db = open(t, path, vfs, "")
	if got := queryString(t, db, "select count(*) || ' ' || count(*) filter (where b like 'changed%') || ' ' || (select count(*) from sqlite_schema where name = 'u') from t"); got != "2000 0 0" {
		t.Fatalf("after the recovery: %s", got)
	}

	if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
		t.Fatalf("integrity_check: %s", got)
	}

	if _, err := os.Stat(path + "-journal"); err == nil {
		t.Fatal("the journal is still there")
	}
}

// Processes write one encrypted database at the same time: the locks are
// those of the VFS underneath, and they hold.
func TestProcesses(t *testing.T) {
	const key = 51
	vfs := register(t, testKey(key))
	path := filepath.Join(t.TempDir(), "a.db")
	db := open(t, path, vfs, "")
	db.SetMaxOpenConns(1)
	mustExec(t, db, "create table t(a integer primary key, b text)")
	var cmds []*exec.Cmd
	for range 3 {
		cmd := child(t, "insert", path, key)
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}

		cmds = append(cmds, cmd)
	}

	for i := range 200 {
		mustExec(t, db, "insert into t(b) values(?)", fmt.Sprintf("%d from the test", i))
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			t.Errorf("a writer failed: %v", err)
		}
	}

	if got := queryString(t, db, "select count(*) || ' ' || count(distinct substr(b, instr(b, 'from'))) from t"); got != "800 4" {
		t.Fatalf("rows and writers: %s", got)
	}

	if got := queryString(t, db, "pragma integrity_check"); got != "ok" {
		t.Fatalf("integrity_check: %s", got)
	}
}
