// Copyright 2026 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite3

import (
	"math/bits"
	"runtime/debug"
	"unsafe"

	"modernc.org/libc"
)

func ___umulh(tls *libc.TLS, a, b uint64) uint64 {
	hi, _ := bits.Mul64(a, b)
	return hi
}

func _modernc_seh_try(tls *libc.TLS, pWal uintptr, xTry uintptr, pCtx uintptr, xExcept uintptr) (rc int32) {
	if pWal == 0 || xTry == 0 {
		return int32(SQLITE_ERROR)
	}
	defer func() {
		if r := recover(); r != nil {
			var faultAddr uintptr
			type addrGetter interface {
				Addr() uintptr
			}
			if ag, ok := r.(addrGetter); ok {
				faultAddr = ag.Addr()
			}

			wal := (*TWal)(unsafe.Pointer(pWal))
			inShm := false
			if faultAddr != 0 {
				nPages := int(wal.FnWiData)
				for i := 0; i < nPages; i++ {
					pagePtr := *(*uintptr)(unsafe.Pointer(wal.FapWiData + uintptr(i)*unsafe.Sizeof(uintptr(0))))
					if pagePtr != 0 && faultAddr >= pagePtr && faultAddr < pagePtr+uintptr(32768) {
						inShm = true
						break
					}
				}
			} else {
				inShm = true
			}

			if inShm {
				if xExcept != 0 {
					rc = (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{xExcept})))(tls, pWal)
				} else {
					rc = int32(SQLITE_IOERR_IN_PAGE)
				}
				return
			}

			panic(r)
		}
	}()

	old := debug.SetPanicOnFault(true)
	defer debug.SetPanicOnFault(old)

	rc = (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{xTry})))(tls, pWal, pCtx)
	return rc
}
