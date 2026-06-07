// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libc/sys/types"
	sqlite3 "modernc.org/sqlite/lib"
)

// ErrPageCacheTooLate is returned by RegisterPageCacheModule when a SQLite
// connection has already been opened in this process. SQLITE_CONFIG_PCACHE2
// must be installed before sqlite3_initialize, which is called implicitly by
// the first sqlite3_open_v2. After that point SQLite returns SQLITE_MISUSE
// and the engine cannot switch its page cache backend.
var ErrPageCacheTooLate = errors.New(
	"sqlite: RegisterPageCacheModule called after first Open; " +
		"SQLITE_CONFIG_PCACHE2 must be installed before any connection is opened")

// ErrPageCacheConflict is returned when a different PageCacheModule has
// already been registered in this process. The same *PageCacheModule pointer
// may be re-registered without error, which lets multiple library imports
// share a singleton without coordination.
var ErrPageCacheConflict = errors.New(
	"sqlite: a different page cache module is already registered")

// PageCacheModule is the Go-facing view of sqlite3_pcache_methods2. It maps
// 1:1 onto the SQLite contract documented at
// https://sqlite.org/c3ref/pcache_methods2.html.
//
// All callbacks MUST be top-level function declarations. Closures and method
// values are not supported because the binding stores function-descriptor
// addresses obtained via cFuncPointer; a captured closure would put its
// state on the heap and the extracted pointer would not survive.
//
// The tls parameter is the per-thread libc context owned by SQLite. Forward
// it verbatim to any sqlite3.Xsqlite3_* call made from inside a callback.
//
// Per SQLite's contract Init and Shutdown are serialized by the engine; all
// other methods MAY be invoked concurrently from multiple goroutines and
// MUST be threadsafe (this build assumes SQLITE_THREADSAFE=1).
//
// Required fields: Init, Create, Cachesize, Pagecount, Fetch, Unpin,
// Truncate, Destroy. Optional fields (SQLite tolerates a NULL slot):
// Shutdown, Rekey, Shrink. Arg is passed verbatim as the first non-tls
// argument to Init and Shutdown; it is ignored by every other callback.
//
// Callbacks must not call RegisterPageCacheModule directly or
// transitively. Init in particular runs under the openGate read lock
// held by the Open path, so a re-entrant RegisterPageCacheModule would
// deadlock on the same gate's write lock.
type PageCacheModule struct {
	Init      func(tls *libc.TLS, arg uintptr) int32
	Shutdown  func(tls *libc.TLS, arg uintptr)
	Create    func(tls *libc.TLS, szPage, szExtra, bPurgeable int32) uintptr
	Cachesize func(tls *libc.TLS, pCache uintptr, nCachesize int32)
	Pagecount func(tls *libc.TLS, pCache uintptr) int32
	Fetch     func(tls *libc.TLS, pCache uintptr, key uint32, createFlag int32) uintptr
	Unpin     func(tls *libc.TLS, pCache, pPage uintptr, discard int32)
	Rekey     func(tls *libc.TLS, pCache, pPage uintptr, oldKey, newKey uint32)
	Truncate  func(tls *libc.TLS, pCache uintptr, iLimit uint32)
	Destroy   func(tls *libc.TLS, pCache uintptr)
	Shrink    func(tls *libc.TLS, pCache uintptr)
	Arg       uintptr
}

// Page describes one logical cache entry. It is provided as a stable
// interface surface for impl authors writing custom modules in the
// follow-up MR; the binding does not use Page directly in this MR.
//
// Buf must point at szPage bytes of writable memory whose address does not
// move between the Fetch that first returned the entry and the eviction
// signal (Unpin with discard=true, a Truncate sweep covering the key, or
// Destroy on the owning cache). Extra must point at szExtra bytes of
// memory with the same stability guarantee. Go-heap memory that the GC may
// relocate is forbidden; use libc.Xmalloc, runtime.Pinner-pinned slices,
// or an off-heap allocator.
type Page interface {
	Buf() unsafe.Pointer
	Extra() unsafe.Pointer
}

// PageEq is an optional fast-path for cache implementations that wrap
// entries in throwaway interface values per Fetch. When implemented, the
// binding uses Same to test entry identity without falling back to the
// pointer comparison of Buf and Extra.
//
// Same MUST satisfy reflexivity (a.Same(a) is true) and symmetry
// (a.Same(b) implies b.Same(a)). Transitivity is not required because the
// binding only compares pairwise. Correctness does not depend on PageEq;
// it is purely an allocation-shape optimization.
type PageEq interface {
	Page
	Same(other Page) bool
}

// pcacheState holds the package-global state shared between
// RegisterPageCacheModule and the Open path.
//
// Locking discipline:
//
//   - openGate.RLock is held for the body of withOpenGate (called from the
//     Driver.Open path). Many opens may proceed concurrently; what is
//     forbidden is registering a page cache while any open is in flight.
//   - openGate.Lock is held for the body of RegisterPageCacheModule. The
//     write lock drains all in-flight opens and blocks all subsequent
//     opens until Xsqlite3_config completes.
//   - opened is set with an unconditional Store on every Open. The Store
//     is intentional and cheap; using CompareAndSwap to only mutate on
//     the first Open buys nothing because the read-side under the write
//     lock is uncontended. The atomic.Bool lets the hot Open path read
//     under RLock without paying for a full mutex acquisition.
//   - configOnce guarantees Xsqlite3_config runs at most once per process.
//     A non-OK return code or an OOM during the methods-table allocation
//     leaves configErr set and unregisters the module; a subsequent
//     Register with the same pointer will retry through configOnce.Do
//     only if the once has not fired, which after a successful initial
//     call it has. Reload is therefore not supported in this MR.
//   - registered holds the canonical *PageCacheModule pointer for
//     idempotency comparison.
//   - cMethods is the libc.Xcalloc-owned C struct SQLite reads at
//     sqlite3_initialize time. Allocated once and lives until process
//     exit; allocating via libc avoids tripping Go's checkptr when the
//     transpiled C code reads the struct (the same reasoning that
//     vtab.go:130-141 uses for sqlite3_module).
var pcacheState struct {
	openGate   sync.RWMutex
	opened     atomic.Bool
	configOnce sync.Once
	configErr  error
	registered *PageCacheModule
	cMethods   uintptr
}

// markConnectionOpened is called from the Open path under
// pcacheState.openGate.RLock before sqlite3_open_v2. The RLock-side store
// happens-before RUnlock; RegisterPageCacheModule's Lock acquisition
// waits for all readers to drain, so its subsequent Load observes every
// prior store.
func markConnectionOpened() {
	pcacheState.opened.Store(true)
}

// withOpenGate runs fn while holding the openGate read lock. The Open
// path wraps its entire body in this so a concurrent
// RegisterPageCacheModule cannot squeeze in between the opened-flag
// store and sqlite3_open_v2.
func withOpenGate(fn func() error) error {
	pcacheState.openGate.RLock()
	defer pcacheState.openGate.RUnlock()
	markConnectionOpened()
	return fn()
}

// RegisterPageCacheModule installs m as the process-global SQLite page
// cache via SQLITE_CONFIG_PCACHE2. It MUST be called before the first
// sql.Open or driver.Open in the program.
//
// Concurrency contract:
//
//   - Safe to call concurrently with itself and with other Register*
//     entry points.
//   - Blocks until any sql.Open calls currently in progress complete.
//     Trade-off: a Register call may block for the duration of an
//     in-flight Open. WAL recovery or cold-file-lock contention can make
//     that wait visible.
//   - Once any connection has been opened, returns ErrPageCacheTooLate
//     without mutating the global module slot.
//   - Calling twice with the same *PageCacheModule pointer is a no-op
//     success. Calling twice with a different pointer returns
//     ErrPageCacheConflict.
//
// Required fields on m: Init, Create, Cachesize, Pagecount, Fetch, Unpin,
// Truncate, Destroy. Passing nil for any required field returns an error
// naming the missing field.
func RegisterPageCacheModule(m *PageCacheModule) error {
	if m == nil {
		return errors.New("sqlite: RegisterPageCacheModule(nil)")
	}
	if err := validatePageCacheModule(m); err != nil {
		return err
	}

	pcacheState.openGate.Lock()
	defer pcacheState.openGate.Unlock()

	// Idempotency / conflict / too-late checks before we touch the once.
	// pcacheState.registered is non-nil only after a SUCCESSFUL install,
	// so observing it here means a prior Register completed without
	// returning configErr.
	if pcacheState.registered != nil {
		if pcacheState.registered == m {
			return nil
		}
		if pcacheState.opened.Load() {
			return ErrPageCacheTooLate
		}
		return ErrPageCacheConflict
	}
	if pcacheState.opened.Load() {
		return ErrPageCacheTooLate
	}

	// First-time install. The once captures m by closure but does not
	// commit pcacheState.registered = m until Xsqlite3_config succeeds.
	// If the body panics or returns early on OOM / SQLITE error,
	// configErr is set, registered stays nil, and every subsequent
	// Register call returns the sticky configErr because the once is
	// already fired. A second install attempt within the process is
	// therefore not possible; this matches the "reload not supported"
	// invariant documented on pcacheState.
	pcacheState.configOnce.Do(func() {
		defer func() {
			if r := recover(); r != nil {
				pcacheState.configErr = fmt.Errorf(
					"sqlite: panic during PCACHE2 install: %v", r)
				pcacheState.registered = nil
				pcacheState.cMethods = 0
			}
		}()

		tls := libc.NewTLS()
		defer tls.Close()

		methodsPtr := libc.Xcalloc(tls, 1, types.Size_t(unsafe.Sizeof(pcacheMethods2{})))
		if methodsPtr == 0 {
			pcacheState.configErr = errors.New("sqlite: out of memory allocating pcache_methods2")
			return
		}
		populateCMethods(methodsPtr, m)

		varArgs := libc.Xmalloc(tls, types.Size_t(unsafe.Sizeof(uintptr(0))))
		if varArgs == 0 {
			libc.Xfree(tls, methodsPtr)
			pcacheState.configErr = errors.New("sqlite: out of memory allocating va_list")
			return
		}
		defer libc.Xfree(tls, varArgs)

		rc := sqlite3.Xsqlite3_config(tls,
			int32(sqlite3.SQLITE_CONFIG_PCACHE2),
			libc.VaList(varArgs, methodsPtr))
		if rc != sqlite3.SQLITE_OK {
			libc.Xfree(tls, methodsPtr)
			pcacheState.configErr = fmt.Errorf(
				"sqlite: Xsqlite3_config(SQLITE_CONFIG_PCACHE2) returned %d", rc)
			return
		}

		// Commit only after every fallible step succeeded.
		pcacheState.cMethods = methodsPtr
		pcacheState.registered = m
	})

	return pcacheState.configErr
}

// MustRegisterPageCacheModule is like RegisterPageCacheModule but panics
// on any error. Intended for init() use where a missing page cache is
// fatal. Mirrors the precedent set by MustRegisterDeterministicScalarFunction.
func MustRegisterPageCacheModule(m *PageCacheModule) {
	if err := RegisterPageCacheModule(m); err != nil {
		panic(err)
	}
}

// validatePageCacheModule reports the first missing required field on m.
// SQLite tolerates NULL for Shutdown, Rekey, and Shrink; every other
// callback must be present.
func validatePageCacheModule(m *PageCacheModule) error {
	switch {
	case m.Init == nil:
		return errors.New("sqlite: PageCacheModule.Init is required")
	case m.Create == nil:
		return errors.New("sqlite: PageCacheModule.Create is required")
	case m.Cachesize == nil:
		return errors.New("sqlite: PageCacheModule.Cachesize is required")
	case m.Pagecount == nil:
		return errors.New("sqlite: PageCacheModule.Pagecount is required")
	case m.Fetch == nil:
		return errors.New("sqlite: PageCacheModule.Fetch is required")
	case m.Unpin == nil:
		return errors.New("sqlite: PageCacheModule.Unpin is required")
	case m.Truncate == nil:
		return errors.New("sqlite: PageCacheModule.Truncate is required")
	case m.Destroy == nil:
		return errors.New("sqlite: PageCacheModule.Destroy is required")
	}
	return nil
}

// populateCMethods writes function-descriptor uintptrs into the libc-owned
// SQLite C methods table using named-field assignment. Named fields are
// portable across all 19 supported GOOS/GOARCH pairs; the Go compiler
// emits the correct per-arch offsets. Hardcoded byte offsets would
// corrupt the struct on every 32-bit arch (no padding after the int32
// FiVersion when uintptr is 4 bytes) and on netbsd_amd64 (explicit
// F__ccgo_pad1 [4]byte in the regenerated layout).
func populateCMethods(ptr uintptr, m *PageCacheModule) {
	dst := (*pcacheMethods2)(unsafe.Pointer(ptr))
	dst.FiVersion = 1
	dst.FpArg = m.Arg
	dst.FxInit = cFuncPointer(m.Init)
	if m.Shutdown != nil {
		dst.FxShutdown = cFuncPointer(m.Shutdown)
	}
	dst.FxCreate = cFuncPointer(m.Create)
	dst.FxCachesize = cFuncPointer(m.Cachesize)
	dst.FxPagecount = cFuncPointer(m.Pagecount)
	dst.FxFetch = cFuncPointer(m.Fetch)
	dst.FxUnpin = cFuncPointer(m.Unpin)
	if m.Rekey != nil {
		dst.FxRekey = cFuncPointer(m.Rekey)
	}
	dst.FxTruncate = cFuncPointer(m.Truncate)
	dst.FxDestroy = cFuncPointer(m.Destroy)
	if m.Shrink != nil {
		dst.FxShrink = cFuncPointer(m.Shrink)
	}
}
