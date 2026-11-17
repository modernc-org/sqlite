// Copyright 2022 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || arm
// +build 386 arm

package vfs

import (
	"unsafe"

	"modernc.org/libc"
)

func init() {
	*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 4)) = vfsClose
	*(*func(*libc.TLS, uintptr, uintptr, int32, sqlite_int64) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 8)) = vfsRead
	*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 24)) = vfsFileSize
	*(*func(*libc.TLS, uintptr, int32) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 28)) = vfsLock
	*(*func(*libc.TLS, uintptr, int32) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 32)) = vfsUnlock
	*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 36)) = vfsCheckReservedLock
	*(*func(*libc.TLS, uintptr, int32, uintptr) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 40)) = vfsFileControl
	*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 44)) = vfsSectorSize
	*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&vfsio)) + 48)) = vfsDeviceCharacteristics
}
