// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Per-arch alias for sqlite3_pcache_methods2 on the GOOS/GOARCH pairs
// whose older lib/ generator emits the type as Sqlite3_pcache_methods2
// (no T prefix): freebsd_386, freebsd_arm, netbsd_amd64.
//
// The complementary file pagecache_alias_new.go covers every other
// supported pair, where the regenerated lib/ emits the type as
// Tsqlite3_pcache_methods2. Splitting on the build tag lets pagecache.go
// refer to a single unqualified pcacheMethods2 type uniformly.

//go:build (freebsd && 386) || (freebsd && arm) || (netbsd && amd64)

package sqlite

import sqlite3 "modernc.org/sqlite/lib"

// pcacheMethods2 is the Go view of the C sqlite3_pcache_methods2 struct
// as emitted by the old-style cznic generator. The Sqlite3_pcache_methods2
// alias in those lib files maps to the same underlying struct shape; the
// per-arch Go compiler picks the correct byte layout.
type pcacheMethods2 = sqlite3.Sqlite3_pcache_methods2
