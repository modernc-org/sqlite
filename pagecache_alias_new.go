// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Per-arch alias for sqlite3_pcache_methods2 on the GOOS/GOARCH pairs
// whose regenerated lib/ emits the type as Tsqlite3_pcache_methods2.
//
// The complementary file pagecache_alias_old.go covers freebsd_386,
// freebsd_arm, and netbsd_amd64, whose older generator emits the type
// as Sqlite3_pcache_methods2 (no T prefix). The split exists so the
// pagecache.go body can refer to a single unqualified pcacheMethods2
// type regardless of which arch the binary is built for.

//go:build !((freebsd && 386) || (freebsd && arm) || (netbsd && amd64))

package sqlite

import sqlite3 "modernc.org/sqlite/lib"

// pcacheMethods2 is the Go view of the C sqlite3_pcache_methods2 struct
// as emitted by the new-style cznic generator. Field names (FiVersion,
// FpArg, FxInit, ...) are stable across all arches in this group; the
// per-arch Go compiler picks the correct byte layout.
type pcacheMethods2 = sqlite3.Tsqlite3_pcache_methods2
