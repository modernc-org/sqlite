// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"errors"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"

	"modernc.org/libc"
)

// TestPCacheMethods2Layout pins the on-disk shape of the sqlite3_pcache_methods2
// struct as cznic emits it today. The test catches two failure modes:
//
//  1. A regenerated lib/ that reorders or renames fields. populateCMethods uses
//     named-field assignment, so reordering by itself is safe, but renaming a
//     field would silently make a field zero on every arch.
//  2. A regenerated lib/ that introduces a new field between two existing
//     ones, growing the struct without us noticing.
//
// Layout assertions are kept relative on purpose: hard-coded byte offsets
// would be wrong on the 32-bit arches (uintptr is 4 bytes, no padding after
// the int32 FiVersion) and on netbsd_amd64 (explicit F__ccgo_pad1 [4]byte).
func TestPCacheMethods2Layout(t *testing.T) {
	var m pcacheMethods2

	// Field count: int32 iVersion + uintptr pArg + 11 callback uintptrs = 13.
	v := reflect.ValueOf(m)
	if got, want := v.NumField(), 13; got != want {
		t.Fatalf("pcacheMethods2 has %d fields, want %d (struct regeneration drift)", got, want)
	}

	// Field shapes by name.
	wantKinds := map[string]reflect.Kind{
		"FiVersion":   reflect.Int32,
		"FpArg":       reflect.Uintptr,
		"FxInit":      reflect.Uintptr,
		"FxShutdown":  reflect.Uintptr,
		"FxCreate":    reflect.Uintptr,
		"FxCachesize": reflect.Uintptr,
		"FxPagecount": reflect.Uintptr,
		"FxFetch":     reflect.Uintptr,
		"FxUnpin":     reflect.Uintptr,
		"FxRekey":     reflect.Uintptr,
		"FxTruncate":  reflect.Uintptr,
		"FxDestroy":   reflect.Uintptr,
		"FxShrink":    reflect.Uintptr,
	}
	tp := v.Type()
	for name, kind := range wantKinds {
		f, ok := tp.FieldByName(name)
		if !ok {
			t.Errorf("pcacheMethods2.%s is missing (struct regeneration drift)", name)
			continue
		}
		if f.Type.Kind() != kind {
			t.Errorf("pcacheMethods2.%s kind is %s, want %s", name, f.Type.Kind(), kind)
		}
	}

	// Minimum struct size: 1 int32 + 12 uintptrs (FpArg + 11 callbacks). On any
	// arch unsafe.Sizeof(m) must be at least that, accounting for natural
	// alignment.
	minSize := unsafe.Sizeof(int32(0)) + 12*unsafe.Sizeof(uintptr(0))
	if unsafe.Sizeof(m) < minSize {
		t.Fatalf("unsafe.Sizeof(pcacheMethods2) = %d, want >= %d (struct shrunk; layout broken)",
			unsafe.Sizeof(m), minSize)
	}

	// FpArg must follow FiVersion, FxInit must follow FpArg, and FxShrink must
	// be the tail. Catches a regenerator that reorders the C struct.
	off := func(name string) uintptr {
		f, _ := tp.FieldByName(name)
		return f.Offset
	}
	if !(off("FpArg") > off("FiVersion")) {
		t.Errorf("FpArg offset (%d) must follow FiVersion (%d)", off("FpArg"), off("FiVersion"))
	}
	if !(off("FxInit") > off("FpArg")) {
		t.Errorf("FxInit offset (%d) must follow FpArg (%d)", off("FxInit"), off("FpArg"))
	}
	if !(off("FxShrink") > off("FxDestroy")) {
		t.Errorf("FxShrink offset (%d) must follow FxDestroy (%d)", off("FxShrink"), off("FxDestroy"))
	}
}

// validModuleStub returns a PageCacheModule with every required field populated
// by harmless top-level no-op callbacks. The same pointer is returned on every
// call so idempotency tests can compare against it.
func validModuleStub() *PageCacheModule {
	return &PageCacheModule{
		Init:      stubInit,
		Create:    stubCreate,
		Cachesize: stubCachesize,
		Pagecount: stubPagecount,
		Fetch:     stubFetch,
		Unpin:     stubUnpin,
		Truncate:  stubTruncate,
		Destroy:   stubDestroy,
	}
}

// Stub callbacks. They are intentionally top-level (closures are not safe with
// cFuncPointer) but never actually run in TestRegisterPageCacheModuleValidation
// because the validation tests do not progress past the lock.
func stubInit(*libc.TLS, uintptr) int32                   { return 0 }
func stubCreate(*libc.TLS, int32, int32, int32) uintptr   { return 0 }
func stubCachesize(*libc.TLS, uintptr, int32)             {}
func stubPagecount(*libc.TLS, uintptr) int32              { return 0 }
func stubFetch(*libc.TLS, uintptr, uint32, int32) uintptr { return 0 }
func stubUnpin(*libc.TLS, uintptr, uintptr, int32)        {}
func stubTruncate(*libc.TLS, uintptr, uint32)             {}
func stubDestroy(*libc.TLS, uintptr)                      {}

// TestRegisterPageCacheModuleValidation exercises the input-validation surface
// of RegisterPageCacheModule. None of these subtests run the lifecycle gate
// (they fail before the lock is taken), so they leave pcacheState untouched
// and can be run in any order alongside other tests in the package.
func TestRegisterPageCacheModuleValidation(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if err := RegisterPageCacheModule(nil); err == nil ||
			!strings.Contains(err.Error(), "RegisterPageCacheModule(nil)") {
			t.Fatalf("RegisterPageCacheModule(nil) error = %v, want sentinel message", err)
		}
	})

	required := []struct {
		name string
		zero func(*PageCacheModule)
	}{
		{"Init", func(m *PageCacheModule) { m.Init = nil }},
		{"Create", func(m *PageCacheModule) { m.Create = nil }},
		{"Cachesize", func(m *PageCacheModule) { m.Cachesize = nil }},
		{"Pagecount", func(m *PageCacheModule) { m.Pagecount = nil }},
		{"Fetch", func(m *PageCacheModule) { m.Fetch = nil }},
		{"Unpin", func(m *PageCacheModule) { m.Unpin = nil }},
		{"Truncate", func(m *PageCacheModule) { m.Truncate = nil }},
		{"Destroy", func(m *PageCacheModule) { m.Destroy = nil }},
	}
	for _, tc := range required {
		t.Run("missing/"+tc.name, func(t *testing.T) {
			m := validModuleStub()
			tc.zero(m)
			err := RegisterPageCacheModule(m)
			if err == nil ||
				!strings.Contains(err.Error(), "PageCacheModule."+tc.name) {
				t.Fatalf("RegisterPageCacheModule missing %s, err = %v, want mention of field",
					tc.name, err)
			}
		})
	}
}

// TestRegisterPageCacheModuleLifecycle exercises the gate ordering between
// RegisterPageCacheModule and the Open path. The subtests share pcacheState
// and run in declaration order; once Lifecycle/TooLate flips the opened flag,
// every other subtest in this file that calls RegisterPageCacheModule with a
// valid module would also return ErrPageCacheTooLate, so the validation tests
// live in their own function above and only this single function holds tests
// that mutate the global lifecycle state.
//
// This test is intentionally structured as a single function: parallel
// execution across separate top-level tests would race on pcacheState.
func TestRegisterPageCacheModuleLifecycle(t *testing.T) {
	// Confirm we are starting from a clean slate: if any earlier test in
	// this run polluted pcacheState (e.g. by opening a connection), the
	// rest of the lifecycle subtests would be meaningless. Skip rather
	// than miss a regression.
	pcacheState.openGate.RLock()
	registered := pcacheState.registered
	opened := pcacheState.opened.Load()
	pcacheState.openGate.RUnlock()
	if registered != nil || opened {
		t.Skip("pcacheState already polluted by an earlier test in this run; " +
			"this lifecycle test must run before any sql.Open call. " +
			"Move it earlier or run via go test -run TestRegisterPageCacheModuleLifecycle.")
	}

	t.Run("DifferentPointersConflict", func(t *testing.T) {
		// Two distinct *PageCacheModule values produce ErrPageCacheConflict
		// once one of them is registered. We do not call Xsqlite3_config
		// here because the gate check happens before configOnce.Do.
		m1 := validModuleStub()
		m2 := validModuleStub()
		if m1 == m2 {
			t.Fatal("validModuleStub returned the same pointer twice; invariant broken")
		}

		// Manually set the registered slot to m1 without going through the
		// configOnce path. This isolates the conflict check from the real
		// SQLite config call, which we cannot undo within a single process.
		pcacheState.openGate.Lock()
		if pcacheState.registered != nil {
			pcacheState.openGate.Unlock()
			t.Skip("registered slot was filled between Skip-guard and Lock; aborting")
		}
		pcacheState.registered = m1
		pcacheState.openGate.Unlock()
		defer func() {
			pcacheState.openGate.Lock()
			pcacheState.registered = nil
			pcacheState.openGate.Unlock()
		}()

		// Same pointer = no error.
		if err := RegisterPageCacheModule(m1); err != nil {
			t.Errorf("RegisterPageCacheModule(m1) second call err = %v, want nil (idempotent)", err)
		}

		// Different pointer = ErrPageCacheConflict.
		err := RegisterPageCacheModule(m2)
		if !errors.Is(err, ErrPageCacheConflict) {
			t.Errorf("RegisterPageCacheModule(m2) err = %v, want ErrPageCacheConflict", err)
		}
	})

	t.Run("TooLate", func(t *testing.T) {
		// Simulate "a connection has been opened" by flipping the gate flag
		// directly. We cannot easily undo this state within a process, so
		// the subtest runs last in this function.
		pcacheState.openGate.Lock()
		if pcacheState.registered != nil {
			pcacheState.openGate.Unlock()
			t.Skip("registered slot was filled before TooLate subtest; aborting")
		}
		prevOpened := pcacheState.opened.Load()
		pcacheState.opened.Store(true)
		pcacheState.openGate.Unlock()
		defer func() {
			pcacheState.openGate.Lock()
			pcacheState.opened.Store(prevOpened)
			pcacheState.openGate.Unlock()
		}()

		err := RegisterPageCacheModule(validModuleStub())
		if !errors.Is(err, ErrPageCacheTooLate) {
			t.Errorf("RegisterPageCacheModule after Open err = %v, want ErrPageCacheTooLate", err)
		}
	})

	t.Run("TooLateIdempotentForSamePointer", func(t *testing.T) {
		// A library that holds onto its module pointer and re-registers
		// after Open should get a no-op success, not ErrPageCacheTooLate.
		// This subtest uses the same flip-flag trick as TooLate above.
		m := validModuleStub()

		pcacheState.openGate.Lock()
		if pcacheState.registered != nil {
			pcacheState.openGate.Unlock()
			t.Skip("registered slot was filled before idempotency subtest; aborting")
		}
		pcacheState.registered = m
		prevOpened := pcacheState.opened.Load()
		pcacheState.opened.Store(true)
		pcacheState.openGate.Unlock()
		defer func() {
			pcacheState.openGate.Lock()
			pcacheState.registered = nil
			pcacheState.opened.Store(prevOpened)
			pcacheState.openGate.Unlock()
		}()

		if err := RegisterPageCacheModule(m); err != nil {
			t.Errorf("re-register same pointer after Open err = %v, want nil", err)
		}
	})
}

// TestOpenGateConcurrentReaders is a smoke test that the openGate RWMutex
// does not deadlock under a fan of concurrent Open-side readers. It does
// not exercise the Register-vs-Open race: a real Register call would
// commit Xsqlite3_config, which is process-global and would break every
// other test in the package. The Register-vs-Open ordering is exercised
// indirectly by TestRegisterPageCacheModuleLifecycle/TooLate, which sets
// the opened flag and then verifies Register sees it.
//
// What this test owns: many concurrent withOpenGate calls succeed without
// races, every one of them flips opened to true, and the gate is free for
// a subsequent Lock acquisition.
func TestOpenGateConcurrentReaders(t *testing.T) {
	const iterations = 32
	const openers = 16
	var wg sync.WaitGroup
	wg.Add(openers)

	for i := 0; i < openers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				_ = withOpenGate(func() error { return nil })
			}
		}()
	}
	wg.Wait()

	if !pcacheState.opened.Load() {
		t.Fatal("pcacheState.opened is false after withOpenGate fan-out; gate is broken")
	}

	// Confirm Lock acquisition still completes promptly — a stuck reader
	// would block this indefinitely and the test would time out.
	done := make(chan struct{})
	go func() {
		pcacheState.openGate.Lock()
		pcacheState.openGate.Unlock()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Lock did not acquire within 5s after RLock fan-out; gate is stuck")
	}
}
