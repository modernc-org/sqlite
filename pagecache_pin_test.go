// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"strings"
	"testing"
	"unsafe"

	"modernc.org/libc"
)

// pinTestPage points into the Go heap, which the Page contract forbids. That
// is harmless here: no SQLite is involved and the binding never dereferences
// Buf or Extra, it only records them.
type pinTestPage struct{ id int }

func (p *pinTestPage) Buf() unsafe.Pointer   { return unsafe.Pointer(p) }
func (p *pinTestPage) Extra() unsafe.Pointer { return unsafe.Pointer(p) }

// pinTestCache returns whatever next holds from Fetch, so a test can make it
// keep, replace or drop a page at will, including in breach of the contract.
type pinTestCache struct{ next Page }

func (c *pinTestCache) SetSize(int)                  {}
func (c *pinTestCache) PageCount() int               { return 0 }
func (c *pinTestCache) Fetch(uint32, FetchMode) Page { return c.next }
func (c *pinTestCache) Unpin(Page, bool)             {}
func (c *pinTestCache) Rekey(Page, uint32, uint32)   {}
func (c *pinTestCache) Truncate(uint32)              {}
func (c *pinTestCache) Destroy()                     {}
func (c *pinTestCache) Shrink()                      {}

// TestPcacheBindingPinnedContract drives the trampolines the way SQLite does
// and checks that the binding frees a stub only once SQLite has unpinned it:
// a Cache that drops or replaces a pinned page makes the binding panic rather
// than leave SQLite holding freed memory.
func TestPcacheBindingPinnedContract(t *testing.T) {
	const key = 7
	p1, p2 := &pinTestPage{1}, &pinTestPage{2}

	for _, tc := range []struct {
		name      string
		unpin     bool // Unpin(discard=false) between the two Fetches
		second    Page
		wantPanic string
	}{
		{"retained while pinned", false, p1, ""},
		{"replaced while pinned", false, p2, "returned a different Page for page 7"},
		{"dropped while pinned", false, nil, "returned nil for page 7"},
		{"replaced after unpin", true, p2, ""},
		{"dropped after unpin", true, nil, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tls := libc.NewTLS()
			defer tls.Close()

			c := &pinTestCache{next: p1}
			h := registerPcacheBinding(&pcacheBinding{
				cache:  c,
				byKey:  map[uint32]uintptr{},
				byStub: map[uintptr]pcacheEntry{},
			})
			defer pcacheTrampolineDestroy(tls, h)

			stub := pcacheTrampolineFetch(tls, h, key, 1)
			if stub == 0 {
				t.Fatal("first Fetch returned no stub")
			}

			if tc.unpin {
				pcacheTrampolineUnpin(tls, h, stub, 0)
			}

			c.next = tc.second
			var got uintptr
			panicked := func() (r any) {
				defer func() { r = recover() }()
				got = pcacheTrampolineFetch(tls, h, key, 1)
				return nil
			}()

			switch {
			case tc.wantPanic == "" && panicked != nil:
				t.Fatalf("unexpected panic: %v", panicked)
			case tc.wantPanic != "" && panicked == nil:
				t.Fatalf("no panic; want one mentioning %q", tc.wantPanic)
			case tc.wantPanic != "":
				if s, _ := panicked.(string); !strings.Contains(s, tc.wantPanic) {
					t.Fatalf("panic %q does not mention %q", panicked, tc.wantPanic)
				}
				return
			}

			if tc.second == p1 && got != stub {
				t.Fatalf("retained page got stub %#x, want the same %#x", got, stub)
			}

			if tc.second == nil && got != 0 {
				t.Fatalf("dropped page got stub %#x, want none", got)
			}
		})
	}
}
