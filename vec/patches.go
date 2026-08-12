//go:build linux || darwin || freebsd || netbsd || openbsd || windows

// Package vec bundles sqlite-vec, the vector search extension for SQLite, in
// the same CGo-free form as the rest of this module. The bundled version is
// v0.1.9; see https://github.com/asg017/sqlite-vec for the extension's own
// documentation.
//
// The package exports no API. Import it for its side effect, which installs
// the extension with sqlite3_auto_extension so that every connection opened
// afterwards has it loaded:
//
//	import (
//		"database/sql"
//
//		_ "modernc.org/sqlite"
//		_ "modernc.org/sqlite/vec"
//	)
//
// The vec0 virtual table module and the vec_* SQL functions are then
// available:
//
//	CREATE VIRTUAL TABLE vec_examples USING vec0(embedding float[8])
//
// # License
//
// sqlite-vec is Copyright (c) 2024 Alex Garcia and is dual-licensed Apache-2.0
// OR MIT; it is used here under the MIT license. That is a different license
// from the rest of this module -- SQLite itself is public domain and this
// module's own code is BSD-3-Clause -- and the MIT notice must accompany
// redistribution of the sources in this package. It travels with them as
// LICENSE-SQLITE_VEC in the module root.
package vec

import (
	"modernc.org/libc"
	sqlite3 "modernc.org/sqlite/lib"
)

func init() {
	tls := libc.NewTLS()
	defer tls.Close()
	sqlite3.Xsqlite3_auto_extension(tls, __ccgo_fp(Xsqlite3_vec_init))
}
