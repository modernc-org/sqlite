// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "modernc.org/sqlite"

import (
	"testing"

	"modernc.org/sqlite/internal/vendorstamp"
)

// The builders run this suite on every target and a release is tagged only when
// they are all green, so this is what keeps a release from carrying generated
// code that vendor.json does not describe: a stamp written from a dirty sibling
// checkout, one left behind by an earlier make vendor, or lib/ and vec/ edited
// by hand afterwards.
func TestVendorStamp(t *testing.T) {
	if err := vendorstamp.Check("."); err != nil {
		t.Fatal(err)
	}
}
