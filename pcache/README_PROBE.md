# MR-131 shared-cache concurrency: experiment & findings

**Throwaway branch `tmp-mr-131-probe` — not for merge.** It exists to
answer the question behind MR-131 (`pcache-shared-cache-draft`):

> When the pcache pool is registered and a database is opened with
> `cache=shared`, one `*pCache` is shared across connections. Does the
> current unlocked `pcache` impl actually race?

## TL;DR

**No real race. The impl is functionally safe under `cache=shared`** —
SQLite's per-`BtShared` mutex serialises every pcache2 callback, so the
shared cache is never entered by two goroutines at once.

**But `go test -race` reports a *false positive*** on that workload,
because the Go race detector cannot model libc's pthread-mutex-backed
`BtShared` serialisation as a happens-before edge. So a pcache-pool user
who also uses `cache=shared` *and* runs their suite under `-race` (very
common) will see spurious `DATA RACE` failures.

So MR-131's premise ("the current implementation would race") is
*technically false*, but there is still a real, user-facing problem to
decide on (see "Implications").

## Files

- `shared_cache_integrity_test.go` — `TestSharedCacheTwoConns_Integrity`:
  two `sql.Conn` share one `cache=shared` file, both write concurrently,
  `integrity_check` at the end. Runs against **stock** `pool.go`.
- `shared_cache_probe_test.go` — `TestSharedCacheInflightProbe`: same
  workload, but reads the lock-free in-flight counter added to `pool.go`.
- `pool.go` — instrumented (commit 2) with a **probe**: an atomic
  in-flight counter on every cache method. **Atomics only, no locking**,
  so it measures real overlap without changing whether overlap can occur.

## Reproduce

The race report is on **stock** `pool.go`, so use **commit 1** for it
(the probe's atomics in commit 2 incidentally silence the detector):

```
git checkout HEAD~1            # commit 1: stock pool.go + the test
go test -race -run TestSharedCacheTwoConns_Integrity ./pcache/   # FAILS: DATA RACE
go test       -run TestSharedCacheTwoConns_Integrity -count=20 ./pcache/   # PASSES, clean
```

Then the proof it is a false positive, on **commit 2** (this HEAD):

```
go test -run TestSharedCacheInflightProbe -count=3 ./pcache/
# >>> ProbeMaxInflight=1 ProbeConcurrentEvents=0 <<<
```

## Evidence it is a false positive, not a real race

1. **In-flight probe = 1.** With an atomic-only counter on every cache
   method (no lock), the max number of goroutines simultaneously inside
   any one cache is **1**, across the workload that `-race` flags — under
   both normal and `-race` timing. A data race *requires* concurrent
   access; there is none.
2. **Probe is sound (positive control).** Driving four goroutines into
   one cache's `PageCount` directly reports `ProbeMaxInflight=4`, so the
   probe does detect real overlap when it exists.
3. **20× stock runs without `-race`:** clean — `integrity_check` always
   `ok`, no `concurrent map writes` fatal, no corruption. A real race on
   `cache.pages` / `cache.lru` would have surfaced.
4. **All 23 race blocks are `cache.Fetch` ↔ `cache.Unpin`** on a shared
   page's `pinned`/`lruElem` (one conn's BEGIN/`lockBtree` of page 1 vs
   the other's COMMIT/`releasePageOne`). Both are under `BtShared`'s
   mutex; the detector just can't see it.

Mechanism: `SQLITE_THREADSAFE=1`; `sqlite3BtreeEnter` takes `pBt->mutex`
for sharable btrees; on linux/amd64 that mutex is libc's
`pthread_musl.go` atomic-CAS spinlock — real mutual exclusion, but not
registered as synchronisation by the race detector for these C-memory
mutexes.

## Implications for MR-C

The three RFC questions were premised on a real race. Re-framed:

- **Functionality:** `cache=shared` + pcache pool already works
  correctly. No locking is needed for *correctness*.
- **The real issue is `-race` cleanliness.** A per-cache `sync.Mutex`
  (the RFC's Q2a) does make `-race` clean (it gives the detector the
  happens-before edge it's missing) — but that is paying an always-taken
  lock to *suppress a tooling false positive*, not to fix a bug.
- **Alternative:** keep `cache=shared` out of scope (the package already
  documents it as unsupported) and document the `-race` noise; or
  detect+reject a shared-cache parent at `Create` time.

This is a judgement call about whether to support `cache=shared` and
whether `-race` cleanliness is worth a global lock — worth deciding
before writing any of Q1/Q2/Q3.
