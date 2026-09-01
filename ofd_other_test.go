// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux

package sqlite // import "modernc.org/sqlite"

import "testing"

// TestOFDLockingUnavailable: everywhere but Linux the OFD locking switch
// reports itself unavailable and changes nothing.
func TestOFDLockingUnavailable(t *testing.T) {
	if OFDLockingEnabled() {
		t.Fatal("OFDLockingEnabled() = true")
	}
	for _, on := range []bool{true, false} {
		if _, err := OFDLocking(on); err != ErrOFDLockingUnavailable {
			t.Fatalf("OFDLocking(%v): err = %v, want ErrOFDLockingUnavailable", on, err)
		}
	}
}
