// Copyright 2024 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build none
// +build none

// Tool for 1.28+ -> 1.29+. Pulls adjusted libsqlite3 code to this repo.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"modernc.org/gc/v3"
)

var (
	// undupPkg is the deduplicator, used only to expand a source checkout that
	// ships deduplicated; see srcDir. The Makefile passes its own pin, so the
	// repo has a single version of record. This default only keeps a bare
	// ./vendor run working.
	undupPkg = flag.String("undup", "modernc.org/undup@v0.0.5", "pkg@version of the deduplicator used to expand a deduplicated source checkout")

	// The source checkouts. Flags so this can be pointed at scratch copies for
	// testing without touching the real ones.
	libsqlite3Dir   = flag.String("libsqlite3", filepath.Join("..", "libsqlite3"), "modernc.org/libsqlite3 checkout to vendor from")
	libsqliteVecDir = flag.String("libsqlite_vec", filepath.Join("..", "libsqlite_vec"), "modernc.org/libsqlite_vec checkout to vendor from")

	// tempDirs is drained by cleanup, which both fail and the end of main run:
	// fail exits the process, so a defer would not be enough.
	tempDirs []string
)

func cleanup() {
	for _, v := range tempDirs {
		os.RemoveAll(v)
	}
	tempDirs = nil
}

func fail(rc int, msg string, args ...any) {
	cleanup()
	fmt.Fprintln(os.Stderr, strings.TrimSpace(fmt.Sprintf(msg, args...)))
	os.Exit(rc)
}

func copyFile(src, dst string) {
	b, err := os.ReadFile(src)
	if err != nil {
		fail(1, "%s\n", err)
	}

	if err := os.WriteFile(dst, b, 0660); err != nil {
		fail(1, "%s\n", err)
	}
}

// srcDir returns a directory holding full, self-contained per-target files of
// the group named base (base_GOOS_GOARCH.go) for the checkout in dir.
//
// The vendoring below reads one full file per target, by name. A checkout can
// also ship deduplicated: modernc.org/undup folds declarations shared across
// targets into base.go + base_g_<hex>.go, leaving each per-target file holding
// only that target's residue. Vendoring those would silently drop most of the
// package. modernc.org/wa2c already ships that way and libsqlite3 is a
// candidate (see modernc.org/builder/NW_GENERALIZATION_HANDOFF.md), so this
// tool detects the form it is given instead of assuming the one that happens to
// be current.
//
// An already-full checkout is returned as is: no copy, no subprocess, exactly
// what this tool has always done. A deduplicated one is copied to a temporary
// directory and expanded THERE — never in place, which would leave the sibling
// checkout dirty and tempt a "restore" that discards whatever else is
// uncommitted in it. go.mod and go.sum travel with the copy because undup
// resolves the package name of every unaliased import by running "go list"
// inside the directory it expands, which needs a module context.
//
// The detection is undup's own rule (see DedupDir): base.go or base_g_*.go
// present means the tree is folded.
//
// One consequence to expect the first time a source checkout does ship folded:
// the vendored files change textually, everywhere, in one commit. undup
// reconstructs declarations in its own order and recomputes each file's imports,
// dropping ccgo's `var _ = math.Pi` / `var _ reflect.Type` / `var _
// unsafe.Pointer` unused-import suppressors along with the imports they exist to
// keep (undup.isSuppressor). The result is the same package — verified by
// vendoring from a deduplicated copy of both checkouts and building lib/ and
// vec/ for linux/{amd64,s390x}, darwin/arm64, windows/{amd64,386}, freebsd/386
// and netbsd/amd64 — but it is not the same bytes. Do not mistake that diff for
// a defect.
func srcDir(dir, base string) string {
	shared, err := filepath.Glob(filepath.Join(dir, base+"_g_*.go"))
	if err != nil {
		fail(1, "%s\n", err)
	}

	universal := filepath.Join(dir, base+".go")
	if _, err := os.Stat(universal); err == nil {
		shared = append(shared, universal)
	}

	if len(shared) == 0 {
		return dir
	}

	group, err := filepath.Glob(filepath.Join(dir, base+"*.go"))
	if err != nil {
		fail(1, "%s\n", err)
	}

	tmp, err := os.MkdirTemp("", "vendor-libs-")
	if err != nil {
		fail(1, "%s\n", err)
	}

	tempDirs = append(tempDirs, tmp)
	fmt.Printf("%s is deduplicated (%d shared files), expanding a copy in %s\n", dir, len(shared), tmp)
	for _, v := range group {
		copyFile(v, filepath.Join(tmp, filepath.Base(v)))
	}

	// go.sum is optional, go.mod is not: without it "go list" has no module to
	// resolve imports against.
	copyFile(filepath.Join(dir, "go.mod"), filepath.Join(tmp, "go.mod"))
	if b, err := os.ReadFile(filepath.Join(dir, "go.sum")); err == nil {
		if err := os.WriteFile(filepath.Join(tmp, "go.sum"), b, 0660); err != nil {
			fail(1, "%s\n", err)
		}
	}

	cmd := exec.Command("go", "run", *undupPkg, "-expand", "-dir", tmp)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fail(1, "go run %s -expand -dir %s: %s\n", *undupPkg, tmp, err)
	}

	// A silent no-op here would vendor residue files as if they were full ones,
	// so verify rather than trust: the shared files must be gone. The likely
	// cause of them surviving is -undup naming a version too old to read the
	// layout the checkout was folded with.
	left, err := filepath.Glob(filepath.Join(tmp, base+"_g_*.go"))
	if err != nil {
		fail(1, "%s\n", err)
	}

	if _, err := os.Stat(filepath.Join(tmp, base+".go")); err == nil {
		left = append(left, filepath.Join(tmp, base+".go"))
	}

	if len(left) != 0 {
		fail(1, "%s: expanding with %s left %d shared file(s) behind, e.g. %s: is that undup version too old for this layout?\n",
			dir, *undupPkg, len(left), left[0])
	}

	return tmp
}

func main() {
	flag.Parse()
	defer cleanup()

	// Both source checkouts may ship in either form; resolve each once, before
	// the per-target loops that read them by name.
	libDir := srcDir(*libsqlite3Dir, "ccgo")
	vecDir := srcDir(*libsqliteVecDir, "ccgo")

	for _, v := range []struct{ goos, goarch string }{
		{"darwin", "amd64"},
		{"darwin", "arm64"},
		{"freebsd", "386"},
		{"freebsd", "amd64"},
		{"freebsd", "arm"},
		{"freebsd", "arm64"},
		{"linux", "386"},
		{"linux", "amd64"},
		{"linux", "arm"},
		{"linux", "arm64"},
		{"linux", "loong64"},
		{"linux", "ppc64le"},
		{"linux", "riscv64"},
		{"linux", "s390x"},
		{"netbsd", "amd64"},
		{"openbsd", "amd64"},
		{"openbsd", "arm64"},
		{"windows", "386"},
		{"windows", "amd64"},
	} {
		base := fmt.Sprintf("ccgo_%s_%s.go", v.goos, v.goarch)
		if v.goos == "windows" && v.goarch == "amd64" {
			base = "ccgo_windows.go"
		}
		ifn := filepath.Join(libDir, base)
		fmt.Printf("%s/%s\t%s\n", v.goos, v.goarch, ifn)
		in, err := os.ReadFile(ifn)
		if err != nil {
			fail(1, "%s\n", err)
		}

		ast, err := gc.ParseFile(ifn, in)
		if err != nil {
			fail(1, "%s\n", err)
		}

		b := bytes.NewBuffer(nil)
		s := ast.SourceFile.PackageClause.Source(true)
		s = strings.Replace(s, "package libsqlite3", "package sqlite3", 1)
		fmt.Fprintln(b, s)
		fmt.Fprint(b, ast.SourceFile.ImportDeclList.Source(true))
		taken := map[string]struct{}{}
		for n := ast.SourceFile.TopLevelDeclList; n != nil; n = n.List {
			switch x := n.TopLevelDecl.(type) {
			case *gc.TypeDeclNode:
				adn := x.TypeSpecList.TypeSpec.(*gc.AliasDeclNode)
				nm := adn.IDENT.Src()
				taken[nm] = struct{}{}
			case *gc.ConstDeclNode:
				// Some targets (e.g. netbsd/amd64) emit spurious
				// `const <typename> = 0` macro-eval artifacts that collide
				// with the un-prefixed type alias generated below. Record
				// const names so the colliding alias is skipped, leaving the
				// const (harmless, unreferenced) as the sole declaration.
				if y, ok := x.ConstSpec.(*gc.ConstSpecNode); ok {
					taken[y.IDENT.Src()] = struct{}{}
				}
			}
		}
	loop:
		for n := ast.SourceFile.TopLevelDeclList; n != nil; n = n.List {
			switch x := n.TopLevelDecl.(type) {
			case *gc.ConstDeclNode:
				switch y := x.ConstSpec.(type) {
				case *gc.ConstSpecNode:
					if y.IDENT.Src() != "SQLITE_TRANSIENT" {
						fmt.Fprintln(b, x.Source(true))
					}
				default:
					panic(fmt.Sprintf("%v: %T %q", x.Position(), y, x.Source(false)))
				}

			case *gc.FunctionDeclNode:
				fmt.Fprintln(b, x.Source(true))
			case *gc.TypeDeclNode:
				fmt.Fprintln(b, x.Source(true))
				adn := x.TypeSpecList.TypeSpec.(*gc.AliasDeclNode)
				nm := adn.IDENT.Src()
				nm2 := nm[1:]
				if _, ok := taken[nm2]; ok {
					break
				}

				if token.IsExported(nm) {
					fmt.Fprintf(b, "\ntype %s = %s\n", nm2, nm)
				}
			case *gc.VarDeclNode:
				fmt.Fprintln(b, x.Source(true))
			default:
				fmt.Printf("%v: TODO %T\n", n.Position(), x)
				break loop
			}
		}

		b.WriteString(`
type Sqlite3_int64 = sqlite3_int64
type Sqlite3_mutex_methods = sqlite3_mutex_methods
type Sqlite3_value = sqlite3_value

type Sqlite3_index_info = sqlite3_index_info
type Sqlite3_module = sqlite3_module
type Sqlite3_vtab = sqlite3_vtab
type Sqlite3_vtab_cursor = sqlite3_vtab_cursor

`)
		base = strings.Replace(base, "ccgo_", "sqlite_", 1)
		if err := os.WriteFile(filepath.Join("lib", base), b.Bytes(), 0660); err != nil {
			fail(1, "%s\n", err)
		}
	}

	{
		// Unlike SQLite, which is public domain, sqlite-vec is MIT-licensed and
		// its notice must travel with the substantial portion of it vendored
		// below. modernc.org/libsqlite_vec's generator extracts the notice from
		// the upstream tarball as LICENSE-SQLITE_VEC; copy it verbatim next to
		// this module's own LICENSE. Read errors are fatal on purpose: shipping
		// the code without the notice is worse than not vendoring at all.
		{
			// From the checkout itself, not vecDir: the license is not part of
			// the deduplicated group and so never travels to the expansion copy.
			const licenseFile = "LICENSE-SQLITE_VEC"
			ifn := filepath.Join(*libsqliteVecDir, licenseFile)
			fmt.Printf("license\t%s\n", ifn)
			license, err := os.ReadFile(ifn)
			if err != nil {
				fail(1, "%s\n", err)
			}

			if err := os.WriteFile(licenseFile, license, 0660); err != nil {
				fail(1, "%s\n", err)
			}
		}

		for _, v := range []struct{ goos, goarch string }{
			{"darwin", "amd64"},
			{"darwin", "arm64"},
			{"freebsd", "386"},
			{"freebsd", "amd64"},
			{"freebsd", "arm"},
			{"freebsd", "arm64"},
			{"linux", "386"},
			{"linux", "amd64"},
			{"linux", "arm"},
			{"linux", "arm64"},
			{"linux", "loong64"},
			{"linux", "ppc64le"},
			{"linux", "riscv64"},
			{"linux", "s390x"},
			{"netbsd", "amd64"},
			{"openbsd", "amd64"},
			{"openbsd", "arm64"},
			{"windows", "386"},
			{"windows", "amd64"},
		} {
			base := fmt.Sprintf("ccgo_%s_%s.go", v.goos, v.goarch)
			if v.goos == "windows" && v.goarch == "amd64" {
				base = "ccgo_windows.go"
			}
			ifn := filepath.Join(vecDir, base)
			fmt.Printf("%s/%s\t%s\n", v.goos, v.goarch, ifn)
			in, err := os.ReadFile(ifn)
			if err != nil {
				fail(1, "%s\n", err)
			}

			ast, err := gc.ParseFile(ifn, in)
			if err != nil {
				fail(1, "%s\n", err)
			}

			b := bytes.NewBuffer(nil)
			s := ast.SourceFile.PackageClause.Source(true)
			s = strings.Replace(s, "package libsqlite_vec", "package vec", 1)
			fmt.Fprintln(b, s)
			s = ast.SourceFile.ImportDeclList.Source(true)
			s = strings.Replace(s, `"modernc.org/libsqlite3"`, `libsqlite3 "modernc.org/sqlite/lib"`, 1)
			fmt.Fprint(b, s)
			taken := map[string]struct{}{}
			for n := ast.SourceFile.TopLevelDeclList; n != nil; n = n.List {
				switch x := n.TopLevelDecl.(type) {
				case *gc.TypeDeclNode:
					adn := x.TypeSpecList.TypeSpec.(*gc.AliasDeclNode)
					nm := adn.IDENT.Src()
					taken[nm] = struct{}{}
				}
			}
		loopvec:
			for n := ast.SourceFile.TopLevelDeclList; n != nil; n = n.List {
				switch x := n.TopLevelDecl.(type) {
				case *gc.ConstDeclNode:
					switch y := x.ConstSpec.(type) {
					case *gc.ConstSpecNode:
						if y.IDENT.Src() != "SQLITE_TRANSIENT" {
							fmt.Fprintln(b, x.Source(true))
						}
					default:
						panic(fmt.Sprintf("%v: %T %q", x.Position(), y, x.Source(false)))
					}

				case *gc.FunctionDeclNode:
					fmt.Fprintln(b, x.Source(true))
				case *gc.TypeDeclNode:
					fmt.Fprintln(b, x.Source(true))
					adn := x.TypeSpecList.TypeSpec.(*gc.AliasDeclNode)
					nm := adn.IDENT.Src()
					nm2 := nm[1:]
					if _, ok := taken[nm2]; ok {
						break
					}

					if token.IsExported(nm) {
						fmt.Fprintf(b, "\ntype %s = %s\n", nm2, nm)
					}
				case *gc.VarDeclNode:
					fmt.Fprintln(b, x.Source(true))
				default:
					fmt.Printf("%v: TODO %T\n", n.Position(), x)
					break loopvec
				}
			}

			base = strings.Replace(base, "ccgo_", "vec_", 1)
			if err := os.WriteFile(filepath.Join("vec", base), b.Bytes(), 0660); err != nil {
				fail(1, "%s\n", err)
			}
		}
	}
}
