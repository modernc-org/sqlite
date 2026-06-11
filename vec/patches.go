//go:build linux || darwin || freebsd || netbsd || openbsd || windows

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
