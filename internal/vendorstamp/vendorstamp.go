// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vendorstamp checks vendor.json, the record make vendor writes of
// where the generated code in lib/ and vec/ came from.
//
// make vendor copies transpiled Go from sibling checkouts of
// modernc.org/libsqlite3 and modernc.org/libsqlite_vec. Neither is a module
// dependency, so without this record the revisions that produced a release
// could not be recovered from the repository. vendor.json names them, the Go
// toolchain the vendoring ran under, and a digest of the files it produced, so
// that a later edit of those files, or a stamp left over from an earlier run,
// is caught.
//
// Check is what the test suite and the pipeline run. The file is written by
// vendor_libs/main.go, which is a separate, build-tagged tool; it computes
// OutputDigest the same way, and a disagreement between the two fails Check.
package vendorstamp // import "modernc.org/sqlite/internal/vendorstamp"

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// FileName is the stamp, at the repository root.
const FileName = "vendor.json"

// OutputPatterns are the files make vendor produces, as globs relative to the
// repository root. The hand-written files in lib/ and vec/ match none of them.
var OutputPatterns = []string{"lib/sqlite*.go", "vec/vec*.go", "LICENSE-SQLITE_VEC"}

// Stamp is the content of vendor.json.
type Stamp struct {
	// Go is the toolchain make vendor ran under, as go env GOVERSION reports it.
	// It matters: gofmt output differs between Go releases.
	Go string `json:"go"`

	// Undup is the pinned deduplicator, as pkg@version.
	Undup string `json:"undup"`

	Sources []Source `json:"sources"`

	Output Output `json:"output"`
}

// Source is one sibling checkout make vendor read from.
type Source struct {
	Module string `json:"module"`
	// Into is the directory of this repository the source was vendored into.
	Into   string   `json:"into"`
	Commit string   `json:"commit"`
	Tags   []string `json:"tags"`
	// Dirty is a pointer so that an absent value is an error, not false.
	Dirty *bool `json:"dirty"`
	// Requires holds what the checkout's go.mod declares, which is not proof
	// of the binaries that generated its files: a go.work workspace can
	// override them. The first line of each generated file records the
	// generator invocation itself.
	Requires map[string]string `json:"requires"`
}

// Output identifies the files make vendor produced.
type Output struct {
	Files  int    `json:"files"`
	SHA256 string `json:"sha256"`
}

var (
	commitRE = regexp.MustCompile(`^[0-9a-f]{40}$`)
	digestRE = regexp.MustCompile(`^[0-9a-f]{64}$`)
	// A pseudo-version ends in a UTC timestamp and the first twelve hex digits
	// of its commit, in all three forms: vX.0.0-T-H, vX.Y.Z-pre.0.T-H and
	// vX.Y.(Z+1)-0.T-H.
	pseudoRE = regexp.MustCompile(`[-.](?:0\.)?[0-9]{14}-([0-9a-f]{12})$`)
)

var expected = []struct{ module, into string }{
	{"modernc.org/libsqlite3", "lib"},
	{"modernc.org/libsqlite_vec", "vec"},
}

// OutputFiles lists the files matching OutputPatterns under root, sorted,
// with slash-separated paths relative to root.
func OutputFiles(root string) ([]string, error) {
	var r []string
	for _, pattern := range OutputPatterns {
		m, err := filepath.Glob(filepath.Join(root, filepath.FromSlash(pattern)))
		if err != nil {
			return nil, err
		}

		for _, v := range m {
			rel, err := filepath.Rel(root, v)
			if err != nil {
				return nil, err
			}

			r = append(r, filepath.ToSlash(rel))
		}
	}
	sort.Strings(r)
	return r, nil
}

// OutputDigest returns the number of output files under root and a digest of
// them: the SHA-256 of lines "<sha256 of file>  <path>\n", in path order, the
// format sha256sum prints. Carriage returns before a line feed are dropped
// before hashing a file, so that a checkout converting line endings does not
// change the digest; they carry no meaning in Go source or in the notice.
func OutputDigest(root string) (int, string, error) {
	files, err := OutputFiles(root)
	if err != nil {
		return 0, "", err
	}

	manifest := sha256.New()
	for _, v := range files {
		b, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(v)))
		if err != nil {
			return 0, "", err
		}

		sum := sha256.Sum256(bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n")))
		fmt.Fprintf(manifest, "%s  %s\n", hex.EncodeToString(sum[:]), v)
	}
	return len(files), hex.EncodeToString(manifest.Sum(nil)), nil
}

// Read decodes the stamp at root.
func Read(root string) (*Stamp, error) {
	b, err := os.ReadFile(filepath.Join(root, FileName))
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	var s Stamp
	if err := dec.Decode(&s); err != nil {
		return nil, fmt.Errorf("%s: %v", FileName, err)
	}

	if _, err := dec.Token(); err != io.EOF {
		return nil, fmt.Errorf("%s: data after the JSON object", FileName)
	}

	return &s, nil
}

// Validate checks the stamp on its own and against goModLibc, the version of
// modernc.org/libc this repository's go.mod requires.
func (s *Stamp) Validate(goModLibc string) error {
	if !strings.HasPrefix(s.Go, "go1.") {
		return fmt.Errorf("%s: go %q is not a Go toolchain version", FileName, s.Go)
	}

	if !strings.Contains(s.Undup, "@v") {
		return fmt.Errorf("%s: undup %q is not pkg@version", FileName, s.Undup)
	}

	if len(s.Sources) != len(expected) {
		return fmt.Errorf("%s: %d sources, want %d", FileName, len(s.Sources), len(expected))
	}

	for i, want := range expected {
		src := s.Sources[i]
		if src.Module != want.module || src.Into != want.into {
			return fmt.Errorf("%s: source %d is %s into %s, want %s into %s", FileName, i, src.Module, src.Into, want.module, want.into)
		}

		if !commitRE.MatchString(src.Commit) {
			return fmt.Errorf("%s: %s: commit %q is not a full hash", FileName, src.Module, src.Commit)
		}

		if src.Dirty == nil {
			return fmt.Errorf("%s: %s: dirty is missing", FileName, src.Module)
		}

		if *src.Dirty {
			return fmt.Errorf("%s: %s was vendored from a checkout with uncommitted changes; clean it and run make vendor again", FileName, src.Module)
		}

		if src.Requires["modernc.org/libc"] == "" {
			return fmt.Errorf("%s: %s: no modernc.org/libc requirement recorded", FileName, src.Module)
		}
	}

	lib, vec := s.Sources[0], s.Sources[1]
	if a, b := lib.Requires["modernc.org/libc"], vec.Requires["modernc.org/libc"]; a != b {
		return fmt.Errorf("%s: lib/ was vendored against modernc.org/libc %s and vec/ against %s; they must be the same", FileName, a, b)
	}

	if a := lib.Requires["modernc.org/libc"]; a != goModLibc {
		return fmt.Errorf("%s: lib/ and vec/ were vendored against modernc.org/libc %s, but go.mod requires %s", FileName, a, goModLibc)
	}

	// vec/ calls into lib/, so libsqlite_vec must have been built against the
	// libsqlite3 revision lib/ came from.
	want := vec.Requires["modernc.org/libsqlite3"]
	switch m := pseudoRE.FindStringSubmatch(want); {
	case want == "":
		return fmt.Errorf("%s: %s: no modernc.org/libsqlite3 requirement recorded", FileName, vec.Module)
	case m != nil:
		if !strings.HasPrefix(lib.Commit, m[1]) {
			return fmt.Errorf("%s: vec/ requires modernc.org/libsqlite3 %s, but lib/ came from commit %s", FileName, want, lib.Commit)
		}
	default:
		found := false
		for _, v := range lib.Tags {
			found = found || v == want
		}
		if !found {
			return fmt.Errorf("%s: vec/ requires modernc.org/libsqlite3 %s, but lib/ came from %s, tagged %q", FileName, want, lib.Commit, lib.Tags)
		}
	}

	if !digestRE.MatchString(s.Output.SHA256) {
		return fmt.Errorf("%s: output sha256 %q is not a SHA-256", FileName, s.Output.SHA256)
	}

	return nil
}

// GoModFields splits one line of a go.mod file into its tokens: the comment
// is dropped, parentheses are tokens of their own and quoted strings are
// unquoted, as the go.mod grammar allows. Module paths contain neither
// spaces nor "//".
func GoModFields(line string) []string {
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

// GoModRequire returns the version of path that the go.mod file at name
// requires, or "" if it requires none. A path required twice is an error
// rather than a guess at which requirement wins.
func GoModRequire(name, path string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}

	defer f.Close()

	version, block := "", false
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fields := GoModFields(sc.Text())
		switch {
		case len(fields) == 0:
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

		if len(fields) == 2 && fields[0] == path {
			if version != "" {
				return "", fmt.Errorf("%s requires %s twice", name, path)
			}

			version = fields[1]
		}
	}
	return version, sc.Err()
}

// Check reads the stamp at root and verifies it: its fields, its libc against
// root's go.mod, and its output digest against the files on disk.
func Check(root string) error {
	s, err := Read(root)
	if err != nil {
		return err
	}

	libc, err := GoModRequire(filepath.Join(root, "go.mod"), "modernc.org/libc")
	if err != nil {
		return err
	}

	if err := s.Validate(libc); err != nil {
		return err
	}

	n, sum, err := OutputDigest(root)
	if err != nil {
		return err
	}

	if n != s.Output.Files || sum != s.Output.SHA256 {
		return fmt.Errorf("%s: lib/ and vec/ differ from what make vendor produced (%d files, sha256 %s; on disk %d files, sha256 %s); generated files must not be edited by hand, run make vendor instead", FileName, s.Output.Files, s.Output.SHA256, n, sum)
	}

	return nil
}
