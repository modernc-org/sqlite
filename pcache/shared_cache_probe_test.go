// EXPERIMENT ONLY — not for merge. Distinguishes a real data race from a
// race-detector false positive on the write-heavy shared-cache workload.

package pcache_test

import (
	"context"
	"database/sql"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"

	"modernc.org/sqlite/pcache"
)

// TestSharedCacheInflightProbe runs the same write-heavy two-connection
// workload as TestSharedCacheTwoConns_Integrity but measures, via an
// atomic in-flight counter inside the cache (NO lock; independent of the
// race detector), whether two goroutines ever execute the shared cache's
// callbacks at the same instant:
//
//	ProbeMaxInflight == 1 -> callbacks serialised; a -race report would be
//	                         a false positive.
//	ProbeMaxInflight  > 1 -> genuine concurrent entry; the race is real.
//
// Run WITHOUT -race (a real map race surfaces here as a "concurrent map
// writes" fatal, which is itself proof):
//
//	go test -run TestSharedCacheInflightProbe -v ./pcache/
func TestSharedCacheInflightProbe(t *testing.T) {
	pcache.ResetProbe()

	dsn := "file:" + filepath.Join(t.TempDir(), "shared.db") + "?cache=shared"
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)

	ctx := context.Background()
	c1, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer c1.Close()
	c2, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close()

	if _, err := c1.ExecContext(ctx,
		`CREATE TABLE t(k INTEGER PRIMARY KEY, v BLOB)`); err != nil {
		t.Fatal(err)
	}

	const iters = 1500
	writer := func(conn *sql.Conn, base int) {
		for i := 0; i < iters; i++ {
			for {
				_, err := conn.ExecContext(ctx,
					"INSERT OR REPLACE INTO t(k, v) VALUES(?, randomblob(80))", base+i)
				if err == nil {
					break
				}
				if isBusyOrLocked(err) {
					continue
				}
				t.Errorf("write: %v", err)
				return
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); writer(c1, 0) }()
	go func() { defer wg.Done(); writer(c2, 1_000_000) }()
	wg.Wait()

	maxInflight := atomic.LoadInt32(&pcache.ProbeMaxInflight)
	events := atomic.LoadInt64(&pcache.ProbeConcurrentEvents)
	t.Logf(">>> ProbeMaxInflight=%d ProbeConcurrentEvents=%d <<<", maxInflight, events)
	if maxInflight > 1 {
		t.Logf("=> REAL race: two goroutines genuinely entered the shared cache at once")
	} else {
		t.Logf("=> callbacks serialised; a -race report on this workload would be a false positive")
	}
}
