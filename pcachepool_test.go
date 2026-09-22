// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcachepool

package sqlite_test

import (
	"errors"
	"testing"

	"modernc.org/sqlite"
	"modernc.org/sqlite/pcache"
)

// Built with -tags pcachepool (make test_pcache), the whole test suite runs
// with the reference page cache registered, so every page SQLite touches in
// every test goes through the pluggable page cache binding. vec_test.go is
// excluded from such a build: importing modernc.org/sqlite/vec initializes
// SQLite before any init here can register a page cache.
func init() {
	sqlite.MustRegisterPageCache(pcache.New())
}

// TestPcachePoolRegistered guards the build tag itself: if the Pool were not
// in effect, the suite would silently test SQLite's own page cache instead.
func TestPcachePoolRegistered(t *testing.T) {
	err := sqlite.RegisterPageCache(pcache.New())
	if !errors.Is(err, sqlite.ErrPageCacheConflict) && !errors.Is(err, sqlite.ErrPageCacheTooLate) {
		t.Fatalf("RegisterPageCache with a second Pool: got %v, want a conflict with the one registered", err)
	}
}
