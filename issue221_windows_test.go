// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package sqlite_test

import (
	"testing"
)

func TestIssue221_Windows(t *testing.T) {
	// The full two-tier test with unencapsulated engine negative witness lives in
	// modernc.org/libsqlite3 (issue221_windows_test.go), which owns the SQLite amalgamation,
	// C patch (internal/sqlite_issue221.patch), and transpiler configuration (generator.go).
	// The vendored code in modernc.org/sqlite/lib will gain active SEH call sites once
	// upstream builders execute 'make vendor' from the patched libsqlite3.
	t.Skip("sqlite/lib transpile was generated with SQLITE_OMIT_SEH upstream; full SEH tests live in modernc.org/libsqlite3 pending upstream builder regeneration (make vendor)")
}
