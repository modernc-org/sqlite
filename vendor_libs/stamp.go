// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build none
// +build none

// The provenance stamp, vendor.json, which make vendor writes last.
//
// -preflight runs before make vendor changes anything. It inspects both sibling
// checkouts and refuses if either is dirty, if they declare different
// modernc.org/libc versions, or if libsqlite_vec requires a libsqlite3 other
// than the commit ../libsqlite3 is at. What it saw goes to .vendor-preflight.json.
//
// -stamp runs after everything else in the recipe has succeeded, cross-builds
// included. It inspects the checkouts again, refuses if anything changed while
// make vendor ran, and writes vendor.json with a digest of the files produced.
//
// modernc.org/sqlite/internal/vendorstamp checks the result; the test suite and
// the pipeline run that check, and make vendor runs it right after writing.
// OutputDigest below must stay the same computation as the one there, which that
// run verifies.

package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	stampFile     = "vendor.json"
	preflightFile = ".vendor-preflight.json"
)

// Matches internal/vendorstamp.OutputPatterns.
var outputPatterns = []string{"lib/sqlite*.go", "vec/vec*.go", "LICENSE-SQLITE_VEC"}

type stampSource struct {
	Module   string            `json:"module"`
	Into     string            `json:"into"`
	Commit   string            `json:"commit"`
	Tags     []string          `json:"tags"`
	Dirty    bool              `json:"dirty"`
	Requires map[string]string `json:"requires"`

	// Why the checkout counts as dirty; reported, not recorded.
	reasons []string
}

type stampOutput struct {
	Files  int    `json:"files"`
	SHA256 string `json:"sha256"`
}

type stamp struct {
	Go      string        `json:"go"`
	Undup   string        `json:"undup"`
	Sources []stampSource `json:"sources"`
	Output  *stampOutput  `json:"output,omitempty"`
}

// Matches internal/vendorstamp: all three pseudo-version forms.
var pseudoVersionRE = regexp.MustCompile(`[-.](?:0\.)?[0-9]{14}-([0-9a-f]{12})$`)

func git(dir string, stdin string, args ...string) (string, error) {
	cmd := exec.Command("git", append([]string{"-C", dir, "-c", "core.quotepath=off"}, args...)...)
	cmd.Stdin = strings.NewReader(stdin)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git -C %s %s: %v: %s", dir, strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}

	return string(out), nil
}

// goModFields matches internal/vendorstamp.GoModFields.
func goModFields(line string) []string {
	if i := strings.Index(line, "//"); i >= 0 {
		line = line[:i]
	}
	f := strings.Fields(strings.NewReplacer("(", " ( ", ")", " ) ").Replace(line))
	for i, v := range f {
		if u, err := strconv.Unquote(v); err == nil {
			f[i] = u
		}
	}
	return f
}

// goMod returns the module path of the go.mod file at name and the versions it
// requires of the paths given. A path required twice is an error.
func goMod(name string, paths ...string) (module string, requires map[string]string, err error) {
	f, err := os.Open(name)
	if err != nil {
		return "", nil, err
	}

	defer f.Close()

	requires = map[string]string{}
	block := false
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fields := goModFields(sc.Text())
		switch {
		case len(fields) == 0:
			continue
		case fields[0] == "module" && len(fields) == 2:
			module = fields[1]
			continue
		case fields[0] == "require" && len(fields) == 2 && fields[1] == "(":
			block = true
			continue
		case block && fields[0] == ")":
			block = false
			continue
		case fields[0] == "require" && len(fields) == 3:
			fields = fields[1:]
		case !block:
			continue
		}

		for _, p := range paths {
			if len(fields) == 2 && fields[0] == p {
				if _, ok := requires[p]; ok {
					return "", nil, fmt.Errorf("%s requires %s twice", name, p)
				}

				requires[p] = fields[1]
			}
		}
	}
	return module, requires, sc.Err()
}

// inspect describes the checkout in dir, which must be a git checkout of
// module. inputs are the file patterns at its top level that vendoring reads.
//
// The checkout is dirty if git status reports anything, untracked files and
// submodules included whatever the local configuration says, or if the files
// vendoring reads differ in any way from the committed ones: an input file that
// is not committed (an ignored file, say), a committed one missing from disk (a
// sparse checkout), or one whose content differs from its committed blob (which
// git status misses for files marked assume-unchanged or skip-worktree).
//
// The comparison is of raw bytes, git's filters off, because vendoring reads
// raw bytes: a checkout converting line endings (core.autocrlf, eol=crlf)
// counts as dirty even when git status is empty. Check the siblings out
// without conversion.
//
// A change made to a checkout while make vendor runs is caught if it is still
// there when -stamp inspects the checkouts again; one made and undone in the
// meantime is not. Leave the checkouts alone while make vendor runs.
func inspect(dir, module, into string, inputs []string) (stampSource, error) {
	s := stampSource{Module: module, Into: into, Tags: []string{}}

	m, requires, err := goMod(filepath.Join(dir, "go.mod"), "modernc.org/libc", "modernc.org/ccgo/v4", "modernc.org/libsqlite3")
	if err != nil {
		return s, err
	}

	if m != module {
		return s, fmt.Errorf("%s is a checkout of %q, want %s", dir, m, module)
	}

	s.Requires = requires

	out, err := git(dir, "", "rev-parse", "--verify", "HEAD^{commit}")
	if err != nil {
		return s, fmt.Errorf("%s must be a git checkout: %v", dir, err)
	}

	s.Commit = strings.TrimSpace(out)

	if out, err = git(dir, "", "tag", "--points-at", "HEAD"); err != nil {
		return s, err
	}

	s.Tags = append(s.Tags, strings.Fields(out)...)
	sort.Strings(s.Tags)

	if out, err = git(dir, "", "status", "--porcelain=v1", "--untracked-files=all", "--ignore-submodules=none"); err != nil {
		return s, err
	}

	if status := strings.TrimSpace(out); status != "" {
		lines := strings.Split(status, "\n")
		s.reasons = append(s.reasons, fmt.Sprintf("git status reports %d change(s), e.g. %q", len(lines), strings.TrimSpace(lines[0])))
	}

	// NUL-terminated, so that no file name is quoted or split.
	if out, err = git(dir, "", "ls-tree", "-z", "HEAD"); err != nil {
		return s, err
	}

	committed := map[string]string{} // input file name -> blob hash
	for _, line := range strings.Split(strings.TrimRight(out, "\x00"), "\x00") {
		// <mode> SP <type> SP <object> TAB <file>
		meta, name, ok := strings.Cut(line, "\t")
		f := strings.Fields(meta)
		if !ok || len(f) != 3 {
			continue
		}

		for _, p := range inputs {
			if matched, _ := filepath.Match(p, name); matched {
				committed[name] = f[2]
			}
		}
	}

	var onDisk []string
	for _, p := range inputs {
		matches, err := filepath.Glob(filepath.Join(dir, p))
		if err != nil {
			return s, err
		}

		for _, v := range matches {
			onDisk = append(onDisk, filepath.Base(v))
		}
	}
	sort.Strings(onDisk)

	seen := map[string]bool{}
	for _, v := range onDisk {
		seen[v] = true
		if _, ok := committed[v]; !ok {
			s.reasons = append(s.reasons, fmt.Sprintf("%s is read by vendoring but not committed", v))
		}
	}

	for v := range committed {
		if !seen[v] {
			s.reasons = append(s.reasons, fmt.Sprintf("%s is committed but missing from disk", v))
		}
	}

	var paths []string
	for _, v := range onDisk {
		if _, ok := committed[v]; ok {
			paths = append(paths, v)
		}
	}

	if len(paths) != 0 {
		// --no-filters: the committed blobs are the bytes vendoring must read.
		if out, err = git(dir, strings.Join(paths, "\n")+"\n", "hash-object", "--no-filters", "--stdin-paths"); err != nil {
			return s, err
		}

		hashes := strings.Fields(out)
		if len(hashes) != len(paths) {
			return s, fmt.Errorf("git hash-object returned %d hashes for %d files", len(hashes), len(paths))
		}

		for i, v := range paths {
			if hashes[i] != committed[v] {
				s.reasons = append(s.reasons, fmt.Sprintf("%s differs from its committed content", v))
			}
		}
	}

	sort.Strings(s.reasons)
	s.Dirty = len(s.reasons) != 0
	return s, nil
}

func goVersion() (string, error) {
	out, err := exec.Command("go", "env", "GOVERSION").Output()
	if err != nil {
		return "", fmt.Errorf("go env GOVERSION: %v", err)
	}

	return strings.TrimSpace(string(out)), nil
}

// inspectAll inspects both checkouts and applies the refusals that do not
// depend on the vendored output.
func inspectAll(allowDirty bool) (*stamp, error) {
	gv, err := goVersion()
	if err != nil {
		return nil, err
	}

	lib, err := inspect(*libsqlite3Dir, "modernc.org/libsqlite3", "lib", []string{"ccgo*.go", "go.mod", "go.sum"})
	if err != nil {
		return nil, err
	}

	vec, err := inspect(*libsqliteVecDir, "modernc.org/libsqlite_vec", "vec", []string{"ccgo*.go", "go.mod", "go.sum", "LICENSE-SQLITE_VEC"})
	if err != nil {
		return nil, err
	}

	var problems []string
	for _, src := range []stampSource{lib, vec} {
		if src.Dirty && !allowDirty {
			for _, r := range src.reasons {
				problems = append(problems, fmt.Sprintf("%s is dirty: %s", src.Module, r))
			}
		}

		if src.Requires["modernc.org/libc"] == "" {
			problems = append(problems, fmt.Sprintf("%s: go.mod requires no modernc.org/libc", src.Module))
		}
	}

	if a, b := lib.Requires["modernc.org/libc"], vec.Requires["modernc.org/libc"]; a != b {
		problems = append(problems, fmt.Sprintf("the checkouts require different modernc.org/libc versions, %s and %s; lib/ and vec/ must come from the same libc", a, b))
	}

	// vec/ calls into lib/: the libsqlite3 that libsqlite_vec requires must be
	// the commit ../libsqlite3 is at, compared by commit, not by tag name.
	switch want := vec.Requires["modernc.org/libsqlite3"]; {
	case want == "":
		problems = append(problems, fmt.Sprintf("%s: go.mod requires no modernc.org/libsqlite3", vec.Module))
	default:
		if m := pseudoVersionRE.FindStringSubmatch(want); m != nil {
			if !strings.HasPrefix(lib.Commit, m[1]) {
				problems = append(problems, fmt.Sprintf("%s requires modernc.org/libsqlite3 %s, but %s is at %s", vec.Module, want, *libsqlite3Dir, lib.Commit))
			}
			break
		}

		out, err := git(*libsqlite3Dir, "", "rev-parse", "--verify", "--quiet", "refs/tags/"+want+"^{commit}")
		if err != nil {
			problems = append(problems, fmt.Sprintf("%s requires modernc.org/libsqlite3 %s, which is not a tag in %s; fetch its tags", vec.Module, want, *libsqlite3Dir))
			break
		}

		if got := strings.TrimSpace(out); got != lib.Commit {
			problems = append(problems, fmt.Sprintf("%s requires modernc.org/libsqlite3 %s = %s, but %s is at %s", vec.Module, want, got, *libsqlite3Dir, lib.Commit))
		}
	}

	if len(problems) != 0 {
		return nil, fmt.Errorf("refusing to vendor:\n\t%s", strings.Join(problems, "\n\t"))
	}

	return &stamp{Go: gv, Undup: *undupPkg, Sources: []stampSource{lib, vec}}, nil
}

// outputDigest must compute what internal/vendorstamp.OutputDigest computes.
func outputDigest() (int, string, error) {
	var files []string
	for _, p := range outputPatterns {
		m, err := filepath.Glob(filepath.FromSlash(p))
		if err != nil {
			return 0, "", err
		}

		for _, v := range m {
			files = append(files, filepath.ToSlash(v))
		}
	}
	sort.Strings(files)

	manifest := sha256.New()
	for _, v := range files {
		b, err := os.ReadFile(filepath.FromSlash(v))
		if err != nil {
			return 0, "", err
		}

		sum := sha256.Sum256(bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n")))
		fmt.Fprintf(manifest, "%s  %s\n", hex.EncodeToString(sum[:]), v)
	}
	return len(files), hex.EncodeToString(manifest.Sum(nil)), nil
}

func writeJSON(name string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	b = append(b, '\n')
	tmp := name + ".tmp"
	if err := os.WriteFile(tmp, b, 0664); err != nil {
		return err
	}

	return os.Rename(tmp, name)
}

func preflight(allowDirty bool) {
	s, err := inspectAll(allowDirty)
	if err != nil {
		fail(1, "%v\n", err)
	}

	if err := writeJSON(preflightFile, s); err != nil {
		fail(1, "%v\n", err)
	}

	for _, src := range s.Sources {
		fmt.Printf("%s\t%s %s dirty=%v\n", src.Module, src.Commit, strings.Join(src.Tags, ","), src.Dirty)
	}
}

func writeStamp(allowDirty bool) {
	b, err := os.ReadFile(preflightFile)
	if err != nil {
		fail(1, "%v: run the whole make vendor recipe, which starts with ./vendor -preflight\n", err)
	}

	var before stamp
	if err := json.Unmarshal(b, &before); err != nil {
		fail(1, "%s: %v\n", preflightFile, err)
	}

	now, err := inspectAll(allowDirty)
	if err != nil {
		fail(1, "%v\n", err)
	}

	// The checkouts, the toolchain and the pin must be what they were when
	// vendoring started; otherwise the output came from something else.
	now.Output, before.Output = nil, nil
	for i := range now.Sources {
		now.Sources[i].reasons = nil
	}
	if !reflect.DeepEqual(&before, now) {
		fail(1, "a sibling checkout, the Go toolchain or the undup pin changed while make vendor ran; run it again\n\tbefore: %s\n\tnow:    %s\n", mustJSON(&before), mustJSON(now))
	}

	n, sum, err := outputDigest()
	if err != nil {
		fail(1, "%v\n", err)
	}

	now.Output = &stampOutput{Files: n, SHA256: sum}
	if err := writeJSON(stampFile, now); err != nil {
		fail(1, "%v\n", err)
	}

	os.Remove(preflightFile)
	fmt.Printf("%s\t%d files, sha256 %s\n", stampFile, n, sum)
}

func mustJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}

	return string(b)
}
