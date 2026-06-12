# NetBSD/amd64 support — status update

## Summary
`modernc.org/sqlite` builds and its full test suite passes on **NetBSD 10.1 / amd64
/ Go 1.26.3**. The original `lib/sqlite_netbsd_amd64.go` build break (issue #246) is
fixed by a fresh re-transpile, and a NetBSD-specific `modernc.org/libc` bug that caused
a concurrent-WAL `SIGBUS` was root-caused and fixed. The toolchain (`cc/v4`, `ccgo/v4`,
`libc`) is now green on NetBSD as well, clearing the Tier-1 toolchain gate. All changes
are on `master` across the chain (see the cascade table below); the only remaining step
is the maintainer's dependency-ordered re-tagging.

Thanks to **Thomas Klausner (@_wiz_)** for the report and **Leonardo Taccari (@iamleot)**
for the original MRs (sqlite !82, libsqlite3 !1), which this work builds on.

## Tier 2 — DONE (sqlite repo green on NetBSD)
- **Build break fixed.** The committed `lib/sqlite_netbsd_amd64.go` was a stale
  old-generator transpile; replaced with a fresh new-generator transpile (SQLite 3.53.x).
  `GOOS=netbsd GOARCH=amd64 go build ./...`, native build, and `go test -c` all pass.
- **Vendor tool fixed** for a NetBSD-only const-vs-alias collision (spurious
  `const <typename> = 0` macro artifacts); other platforms' output is unchanged.
- **`sqlite-vec` vendored for NetBSD** and `TestVec` passes (the gap was only build tags
  in `vec/patches.go` + `vec_test.go`, now including netbsd).
- **Full suite GREEN on NetBSD**: main package (incl. `-race`), `vfs`, `pcache`, `TestVec`
  — exit 0, no faults.

## Root cause of the concurrent-WAL SIGBUS — a libc bug, now fixed
NetBSD's `mmap(2)` syscall has a `long PAD` argument before `off_t pos`
(`mmap(addr, len, prot, flags, fd, PAD, pos)`). `modernc.org/libc`'s netbsd `Xmmap`
used `unix.Syscall6`, putting the offset in the PAD slot and leaving `pos` as stack
garbage → an unaligned/unbacked mapping → `SIGBUS` in the WAL-index shm. Fix: pass the
offset as the 7th arg via `Syscall9` (matches `x/sys/unix`). Found via `ktrace`/`kdump`
(exactly one of 160 mmaps returned a non-page-aligned address = the fault address).
The fix is ABI-preserving, so dependents need only a **go.mod bump, no re-transpile**.
Released as **`modernc.org/libc v1.73.1`**.

## Tier-1 toolchain gate — MET (cc/v4, ccgo/v4, libc green on NetBSD)
- **libc**: netbsd tests pass; v1.73.1 tagged.
- **cc/v4**: `-short` passes; netbsd/amd64 added to `cc/v4/builder.json`.
- **ccgo/v4**: `make shorttest` now **PASSES deterministically** on NetBSD — all 10
  `TestExec` suites at 0 fails, 0 relocation errors (~49 min on the 2-CPU VM).
  - The intermittent `BUILD FAIL: "missing section for relocation target"` failures were
    a **concurrency race inside the NetBSD Go linker** when run multi-threaded — not a ccgo
    bug. Fix (gated on `runtime.GOOS == "netbsd"`): in `newParallel`, cap the build/link
    worker pool to 1 **and** export `GOMAXPROCS=1` so each `go build` subprocess links
    single-threaded (ccgo's link step inherits `os.Environ`). Both are required.
  - Curated `known_failures_netbsd_amd64_test.go`: deduped a duplicate map key
    (`pr36093.c`) that broke test compilation, and recorded 74 deterministic,
    expected-class feature-gaps the failfast runs had never reached (56 gcc-mirror,
    12 vnmakarov, 6 tcc — SIMD vectors, setjmp/longjmp, bitfield ABI, inline asm, libc
    externals, `__int128`/va-arg-pack). None are NetBSD-specific bugs.
  - Bumped ccgo's `modernc.org/libc` to v1.73.1; added netbsd/amd64 to
    `ccgo/v4/lib/builder.json` (the file the builder resolves for the
    `modernc.org/ccgo/v4/lib` test package — its test + autotag lists).

## Cascade — landed on master
All toolchain + dependent changes are now pushed to `master` (gitlab `cznic/*`),
each at **libc v1.73.1** and carrying its netbsd/amd64 support:

| Repo | master | libc | netbsd |
|------|--------|------|--------|
| `cc/v4` | `f5ed098` | (no libc dep) | builder.json test+autotag |
| `ccgo/v4` | `7ced11c` | v1.73.1 | shorttest green + builder.json |
| `libz` | `fea0036` | v1.73.1 | ✓ |
| `libtcl8.6` | `55054e6d` | v1.73.1 | ✓ |
| `libsqlite3` | `37a0df44` | v1.73.1 | transpile + testfixture |
| `libsqlite_vec` | `c626766` | v1.73.1 | transpile + generator |
| `sqlite` | `0681593` | v1.73.1 | vendored lib + vec |

## Remaining — tag cascade (maintainer)
- Run the builders on the new masters and **tag in dependency order**:
  libz / libtcl8.6 → libsqlite3 → libsqlite_vec → sqlite.
- **`libsqlite_vec`'s netbsd build stays red until then**: its go.mod still resolves the
  released `libsqlite3 v1.13.2` (no netbsd). Once libsqlite3 is re-tagged with netbsd,
  bump libsqlite_vec to that tag and it goes green. `sqlite` is unaffected (it vendors
  lib + vec).
- NetBSD note: with only 2 CPUs the ccgo torture run is serialized for determinism
  (~49 min). The relocation-race is a Go linker concurrency issue on NetBSD, worth an
  upstream Go report.
