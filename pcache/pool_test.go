// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcache

import (
	"testing"

	"modernc.org/sqlite"
)

// pcacheCreate is a shortcut to exercise the Pool with realistic
// SQLite page sizes (4096 bytes for the data buffer and 200 bytes of
// extra scratch, both within the range SQLite uses in default
// configurations).
func pcacheCreate(t *testing.T, purgeable bool) (*Pool, sqlite.Cache) {
	t.Helper()
	p := New()
	c, err := p.Create(4096, 200, purgeable)
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if c == nil {
		t.Fatal("Create returned nil cache")
	}
	t.Cleanup(c.Destroy)
	return p, c
}

func TestPoolImplementsPageCache(t *testing.T) {
	var _ sqlite.PageCache = New()
}

func TestEmptyCacheState(t *testing.T) {
	_, c := pcacheCreate(t, true)
	if got := c.PageCount(); got != 0 {
		t.Errorf("PageCount on empty cache = %d, want 0", got)
	}
	if pg := c.Fetch(0, sqlite.FetchLookup); pg != nil {
		t.Errorf("Fetch(Lookup) on empty cache = %v, want nil", pg)
	}
}

func TestFetchAllocateAndRetain(t *testing.T) {
	pool, c := pcacheCreate(t, true)
	pg := c.Fetch(7, sqlite.FetchCreateEasy)
	if pg == nil {
		t.Fatal("Fetch(CreateEasy) on empty cache returned nil")
	}
	if pg.Buf() == nil || pg.Extra() == nil {
		t.Fatalf("Buf/Extra returned nil: buf=%v extra=%v", pg.Buf(), pg.Extra())
	}
	if got := c.PageCount(); got != 1 {
		t.Errorf("PageCount after allocate = %d, want 1", got)
	}
	c.Unpin(pg, false)
	pg2 := c.Fetch(7, sqlite.FetchLookup)
	if pg2 != pg {
		t.Errorf("Fetch(Lookup) after Unpin(discard=false) returned different page: %v vs %v", pg2, pg)
	}
	if s := pool.Stats(); s.Allocs != 1 || s.Hits != 1 || s.Misses != 1 {
		t.Errorf("Stats = %+v, want Allocs=1 Hits=1 Misses=1", s)
	}
}

func TestUnpinDiscardFrees(t *testing.T) {
	pool, c := pcacheCreate(t, true)
	pg := c.Fetch(1, sqlite.FetchCreateForce)
	c.Unpin(pg, true)
	if got := c.PageCount(); got != 0 {
		t.Errorf("PageCount after Unpin(discard=true) = %d, want 0", got)
	}
	if s := pool.Stats(); s.Evictions != 1 {
		t.Errorf("Stats.Evictions = %d, want 1", s.Evictions)
	}
	if c.Fetch(1, sqlite.FetchLookup) != nil {
		t.Error("Fetch(Lookup) after discard returned a page")
	}
}

func TestNonPurgeableDropsOnUnpin(t *testing.T) {
	_, c := pcacheCreate(t, false)
	pg := c.Fetch(1, sqlite.FetchCreateForce)
	c.Unpin(pg, false)
	if got := c.PageCount(); got != 0 {
		t.Errorf("non-purgeable cache retained page across Unpin(discard=false): PageCount=%d", got)
	}
}

func TestEasyHonoursCacheSize(t *testing.T) {
	pool, c := pcacheCreate(t, true)
	c.SetSize(2)
	a := c.Fetch(1, sqlite.FetchCreateForce)
	b := c.Fetch(2, sqlite.FetchCreateForce)
	if a == nil || b == nil {
		t.Fatal("setup: failed to allocate two pages within targetSize")
	}
	// At cap with both pinned. Easy must refuse to allocate a third,
	// and the refusal must increment Stats.EasyRefusals so the
	// benchmark can surface I/O-pressure proxies for the strict
	// at-cap behavior (cznic note on !127).
	if pg := c.Fetch(3, sqlite.FetchCreateEasy); pg != nil {
		t.Errorf("FetchCreateEasy at cap returned %v, want nil", pg)
	}
	if got := pool.Stats().EasyRefusals; got != 1 {
		t.Errorf("Stats.EasyRefusals after one Easy refusal = %d, want 1", got)
	}
	// A second Easy refusal at cap accumulates.
	if pg := c.Fetch(4, sqlite.FetchCreateEasy); pg != nil {
		t.Errorf("FetchCreateEasy at cap (second call) returned %v, want nil", pg)
	}
	if got := pool.Stats().EasyRefusals; got != 2 {
		t.Errorf("Stats.EasyRefusals after two Easy refusals = %d, want 2", got)
	}
	// Force at cap with everything pinned: overcommit (matches pcache1).
	if pg := c.Fetch(3, sqlite.FetchCreateForce); pg == nil {
		t.Error("FetchCreateForce returned nil at cap with all pinned")
	}
}

func TestForceEvictsLRU(t *testing.T) {
	pool, c := pcacheCreate(t, true)
	c.SetSize(2)
	a := c.Fetch(1, sqlite.FetchCreateForce)
	b := c.Fetch(2, sqlite.FetchCreateForce)
	c.Unpin(a, false)
	c.Unpin(b, false)
	// Both unpinned, cache at cap. Force(3) must evict the LRU tail
	// (a, since b was unpinned more recently and sits at Front).
	if pg := c.Fetch(3, sqlite.FetchCreateForce); pg == nil {
		t.Fatal("FetchCreateForce returned nil with two unpinned slots available")
	}
	if c.Fetch(1, sqlite.FetchLookup) != nil {
		t.Error("expected key=1 evicted as LRU tail; still present")
	}
	if c.Fetch(2, sqlite.FetchLookup) == nil {
		t.Error("expected key=2 retained; was evicted")
	}
	if s := pool.Stats(); s.Evictions < 1 {
		t.Errorf("Stats.Evictions = %d, want >= 1", s.Evictions)
	}
}

func TestSetSizeShrinks(t *testing.T) {
	_, c := pcacheCreate(t, true)
	pages := make([]sqlite.Page, 5)
	for i := range pages {
		pages[i] = c.Fetch(uint32(i), sqlite.FetchCreateForce)
	}
	for _, pg := range pages {
		c.Unpin(pg, false)
	}
	c.SetSize(2)
	if got := c.PageCount(); got != 2 {
		t.Errorf("PageCount after SetSize(2) = %d, want 2", got)
	}
}

func TestSetSizeWithPinnedOvercommit(t *testing.T) {
	_, c := pcacheCreate(t, true)
	pages := make([]sqlite.Page, 4)
	for i := range pages {
		pages[i] = c.Fetch(uint32(i), sqlite.FetchCreateForce)
	}
	// All four are pinned; SetSize(2) cannot evict them.
	c.SetSize(2)
	if got := c.PageCount(); got != 4 {
		t.Errorf("PageCount after SetSize(2) with all pinned = %d, want 4 (overcommit)", got)
	}
}

func TestRekey(t *testing.T) {
	pool, c := pcacheCreate(t, true)
	pg := c.Fetch(1, sqlite.FetchCreateForce)
	c.Rekey(pg, 1, 5)
	if c.Fetch(1, sqlite.FetchLookup) != nil {
		t.Error("Fetch(Lookup) on old key after Rekey returned a page")
	}
	hit := c.Fetch(5, sqlite.FetchLookup)
	if hit != pg {
		t.Errorf("Fetch(Lookup) on new key returned %v, want %v", hit, pg)
	}
	if s := pool.Stats(); s.Rekeys != 1 {
		t.Errorf("Stats.Rekeys = %d, want 1", s.Rekeys)
	}
}

func TestRekeyEvictsCollider(t *testing.T) {
	_, c := pcacheCreate(t, true)
	moving := c.Fetch(1, sqlite.FetchCreateForce)
	collider := c.Fetch(5, sqlite.FetchCreateForce)
	c.Unpin(collider, false)
	c.Rekey(moving, 1, 5)
	hit := c.Fetch(5, sqlite.FetchLookup)
	if hit != moving {
		t.Errorf("after Rekey collision Fetch(5) returned %v, want %v", hit, moving)
	}
	if got := c.PageCount(); got != 1 {
		t.Errorf("PageCount after Rekey collision = %d, want 1 (collider freed)", got)
	}
}

func TestTruncate(t *testing.T) {
	pool, c := pcacheCreate(t, true)
	for i := 0; i < 6; i++ {
		c.Unpin(c.Fetch(uint32(i), sqlite.FetchCreateForce), false)
	}
	c.Truncate(3)
	if got := c.PageCount(); got != 3 {
		t.Errorf("PageCount after Truncate(3) = %d, want 3", got)
	}
	for i := uint32(0); i < 3; i++ {
		if c.Fetch(i, sqlite.FetchLookup) == nil {
			t.Errorf("key=%d evicted by Truncate(3); want retained", i)
		}
	}
	for i := uint32(3); i < 6; i++ {
		if c.Fetch(i, sqlite.FetchLookup) != nil {
			t.Errorf("key=%d retained after Truncate(3); want evicted", i)
		}
	}
	if s := pool.Stats(); s.Truncates != 1 {
		t.Errorf("Stats.Truncates = %d, want 1", s.Truncates)
	}
}

func TestTruncateEvictsPinned(t *testing.T) {
	_, c := pcacheCreate(t, true)
	pg := c.Fetch(7, sqlite.FetchCreateForce)
	// Page is still pinned; Truncate must evict regardless.
	c.Truncate(0)
	if got := c.PageCount(); got != 0 {
		t.Errorf("PageCount after Truncate(0) = %d, want 0 (pinned eviction)", got)
	}
	_ = pg
}

func TestShrinkDropsUnpinned(t *testing.T) {
	_, c := pcacheCreate(t, true)
	a := c.Fetch(1, sqlite.FetchCreateForce)
	b := c.Fetch(2, sqlite.FetchCreateForce)
	c.Unpin(a, false)
	c.Shrink()
	if got := c.PageCount(); got != 1 {
		t.Errorf("PageCount after Shrink with one pinned = %d, want 1", got)
	}
	_ = b
}

func TestFetchLookupOnDestroyedNoPanic(t *testing.T) {
	_, c := pcacheCreate(t, true)
	c.Destroy()
	// Subsequent calls must be no-ops, not panics. The Cleanup will
	// call Destroy a second time; that path is also exercised.
	if c.Fetch(0, sqlite.FetchLookup) != nil {
		t.Error("Fetch on destroyed cache returned non-nil")
	}
	if c.PageCount() != 0 {
		t.Error("PageCount on destroyed cache non-zero")
	}
	c.SetSize(1)
	c.Shrink()
	c.Truncate(0)
}

func TestDistinctNewValuesUnequal(t *testing.T) {
	a := New()
	b := New()
	if a == b {
		t.Error("Two New() calls returned the same *Pool")
	}
}
