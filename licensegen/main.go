// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build none
// +build none

// Command licensegen regenerates LICENSE-3RD-PARTY.md from the module graph
// and from the third-party material vendored in this repository.
//
// Usage:
//
//	make licenses
//
// or, by hand:
//
//	cd licensegen && go build -o ../licgen main.go
//	./licgen [-C dir] [-o file] [-check]
//
// -check writes nothing and exits non-zero if the file on disk is not what the
// current tree would produce. That is the form to wire into CI.
//
// The output is deterministic: it carries no timestamp and changes only when
// the tree's third-party content changes.
//
// The file is named for the LICENSE prefix, not for looks: `go mod vendor`
// selects metadata files by case-sensitive prefix (cmd/go, matchMetadata), so a
// name like 3RD_PARTY_LICENSES.md would never reach a downstream vendor/ tree.
// modernc.org/libc names its own equivalent file the same way.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	defaultOut = "LICENSE-3RD-PARTY.md"
	noticeMark = "\x00notices\x00"
	maxDepth   = 3
)

var (
	oC     = flag.String("C", ".", "directory of the modernc.org/sqlite checkout")
	oOut   = flag.String("o", defaultOut, "license document, relative to -C")
	oMD    = flag.String("sbom-md", "SBOM.md", "SBOM overview document, relative to -C")
	oCDX   = flag.String("sbom-cdx", "sbom.cdx.json", "CycloneDX SBOM, relative to -C")
	oSPDX  = flag.String("sbom-spdx", "sbom.spdx.json", "SPDX SBOM, relative to -C")
	oCheck = flag.Bool("check", false, "do not write; exit non-zero if the file is stale")

	root string

	licNameRE    = regexp.MustCompile(`(?i)^(licen[cs]e|copying)`)
	noticeNameRE = regexp.MustCompile(`(?i)^(licen[cs]e-3rd-party|notice|third[-_ ]?party)`)
	headingRE    = regexp.MustCompile(`^#{1,6}\s`)
	skipDirs     = map[string]bool{".git": true, "testdata": true, "vendor": true, "node_modules": true, "internal": false}
)

// ----------------------------------------------------------------- data model

type sectionID int

const (
	secLinked sectionID = iota
	secTest
	secGraph
	numModSections
)

var sectionTitle = [numModSections]string{
	secLinked: "1. Components linked into your program",
	secTest:   "2. Components used only by the test suite",
	secGraph:  "3. Components in the module graph, not compiled into anything",
}

var sectionIntro = [numModSections]string{
	secLinked: "Everything below ends up in a binary that imports `modernc.org/sqlite`. " +
		"This is the section that matters when you redistribute software built on this driver: " +
		"the notices and license texts it points at are the ones you have to carry with you.",
	secTest: "Compiled into this repository's own tests only. It is not part of the library and " +
		"does not reach a program that imports `modernc.org/sqlite`.",
	secGraph: "Reachable in the module graph -- `go list -m all` reports them, and the Go module " +
		"proxy will fetch them -- but no package of theirs is compiled into the library or into its " +
		"tests. They are listed for completeness, because tools that read `go.sum` or an SBOM " +
		"generated from the module graph will report them. Note in particular that the only " +
		"copyleft-adjacent license anywhere in this project, MPL-2.0, appears here and only here. " +
		"`modernc.org/ccgo` and `modernc.org/cc` are in this section because `modernc.org/libc` " +
		"requires them; they are separately the transpilers that produced the Go code in `lib/` and " +
		"`vec/`, but that happens in the sibling `libsqlite3` repository, not in any build of this one.",
}

// A licFile is one license or notice file found in the tree.
type licFile struct {
	label   string // path as shown to the reader, e.g. "simplelru/LICENSE_list"
	text    string
	notices []string
	grp     *licGroup
}

// A licGroup is a set of licFiles sharing one license text, modulo the
// copyright notices, which are merged and rendered in place.
type licGroup struct {
	id      string
	spdx    string
	masked  string
	notices []string
	members []grpMember
}

// A grpMember is one component covered by a shared license text, together with
// the copyright notice that came with its own copy.
type grpMember struct {
	name    string
	notices []string
}

type component struct {
	name    string
	version string
	url     string
	via     string
	note    string
	sec     sectionID
	kind    string // module, vendored, inherited
	purl    string
	files   []*licFile
	inh     *inheritEntry
}

func (c *component) display() string {
	if c.url == "" {
		return c.name
	}

	return fmt.Sprintf("[%s](%s)", c.name, c.url)
}

// An inherited is a third-party notices document carried by a dependency, the
// transitive tail of the flattening.
type inherited struct {
	id      string
	owner   string // module path
	label   string // file name within the module
	body    string
	entries []*inheritEntry
}

// An inheritEntry is one upstream named by a dependency's notices document.
// These are the transitive tail: musl libc reaches this module this way, and
// nothing in the module graph names it.
type inheritEntry struct {
	name string
	url  string
	body string
	spdx string
	doc  *inherited
}

var urlLineRE = regexp.MustCompile(`(?m)^\s*\*\s+\*\*URL:\*\*\s+(\S+)`)

// parse splits a notices document into one entry per "##" heading and reads the
// license out of each, so the upstreams get real SPDX identifiers instead of a
// pointer at a wall of prose.
func (inh *inherited) parse() {
	var cur *inheritEntry
	var buf []string
	flush := func() {
		if cur == nil {
			return
		}

		cur.body = strings.Trim(strings.Join(buf, "\n"), "\n")
		cur.spdx = spdxOf(cur.body)
		if m := urlLineRE.FindStringSubmatch(cur.body); m != nil {
			cur.url = m[1]
		}
		inh.entries = append(inh.entries, cur)
		buf = nil
	}
	for _, line := range strings.Split(inh.body, "\n") {
		if strings.HasPrefix(line, "## ") {
			flush()
			cur = &inheritEntry{name: strings.TrimSpace(line[3:]), doc: inh}
			continue
		}

		buf = append(buf, line)
	}
	flush()
}

// A repoItem is third-party material living in this repository that is not a Go
// module dependency.
type repoItem struct {
	material string
	location string
	origin   string
	url      string
	file     *licFile
}

var (
	sections  [numModSections][]*component
	inherits  []*inherited
	repoItems []*repoItem
	repoComps []*component
	groups    []*licGroup
	byMasked  = map[string]*licGroup{}
)

// --------------------------------------------------------------------- go(1)

func run(args ...string) string {
	cmd := exec.Command("go", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOWORK=off")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	b, err := cmd.Output()
	if err != nil {
		log.Fatalf("go %s: %v\n%s", strings.Join(args, " "), err, stderr.String())
	}

	return string(b)
}

type modJSON struct {
	Path    string
	Version string
	Dir     string
	Main    bool
}

func modules() (r []modJSON) {
	dec := json.NewDecoder(strings.NewReader(run("list", "-m", "-json", "all")))
	for {
		var m modJSON
		switch err := dec.Decode(&m); err {
		case nil:
			r = append(r, m)
		default:
			return r
		}
	}
}

// download fetches modules into the module cache and reports their
// directories. It deliberately runs in a scratch module: `go mod download` in
// the repository itself would add hashes for modules that no build of this
// module needs to go.sum, and go.sum here is load-bearing.
func download(mods []string) map[string]string {
	dir, err := os.MkdirTemp("", "licensegen")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module licensegen.invalid\n\ngo 1.21\n"), 0o644); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("go", append([]string{"mod", "download", "-json"}, mods...)...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOWORK=off")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	b, err := cmd.Output()
	if err != nil {
		log.Fatalf("go mod download %s: %v\n%s", strings.Join(mods, " "), err, stderr.String())
	}

	r := map[string]string{}
	dec := json.NewDecoder(bytes.NewReader(b))
	for {
		var m modJSON
		if err := dec.Decode(&m); err != nil {
			return r
		}

		r[m.Path+"@"+m.Version] = m.Dir
	}
}

func modSet(args ...string) map[string]bool {
	r := map[string]bool{}
	for _, v := range strings.Split(run(args...), "\n") {
		if v = strings.TrimSpace(v); v != "" {
			r[v] = true
		}
	}
	return r
}

// modDeps maps a module path to the module paths it requires directly. It is
// the module graph as the SBOM's dependency section reports it.
var modDeps = map[string][]string{}

// requirers maps a module path to the module paths requiring it directly.
func requirers() map[string][]string {
	r := map[string]map[string]bool{}
	d := map[string]map[string]bool{}
	for _, line := range strings.Split(run("mod", "graph"), "\n") {
		f := strings.Fields(line)
		if len(f) != 2 {
			continue
		}

		from, to := strip(f[0]), strip(f[1])
		if from == "go" || from == "toolchain" || to == "go" || to == "toolchain" {
			continue
		}

		if r[to] == nil {
			r[to] = map[string]bool{}
		}
		r[to][from] = true
		if d[from] == nil {
			d[from] = map[string]bool{}
		}
		d[from][to] = true
	}
	for from, tos := range d {
		var s []string
		for to := range tos {
			s = append(s, to)
		}
		sort.Strings(s)
		modDeps[from] = s
	}
	m := map[string][]string{}
	for to, froms := range r {
		var s []string
		for from := range froms {
			s = append(s, from)
		}
		sort.Strings(s)
		m[to] = s
	}
	return m
}

func strip(s string) string {
	if i := strings.IndexByte(s, '@'); i >= 0 {
		return s[:i]
	}

	return s
}

// ------------------------------------------------------------- license texts

func normalize(b []byte) string {
	s := strings.ReplaceAll(string(b), "\r\n", "\n")
	lines := strings.Split(s, "\n")
	for i, v := range lines {
		lines[i] = strings.TrimRight(v, " \t")
	}
	return strings.Trim(strings.Join(lines, "\n"), "\n")
}

// mask replaces the first contiguous run of copyright notice lines with a
// marker and returns the notices separately. Later runs, such as the template
// notice in the appendix of Apache-2.0, are left in the body: they are part of
// the text, not an attribution.
func mask(text string) (masked string, notices []string) {
	lines := strings.Split(text, "\n")
	var out []string
	done := false
	for i := 0; i < len(lines); i++ {
		if done || !isNotice(lines[i]) {
			out = append(out, lines[i])
			continue
		}

		for i < len(lines) && isNotice(lines[i]) {
			notices = append(notices, strings.TrimRight(lines[i], " \t"))
			i++
		}
		i--
		out = append(out, noticeMark)
		done = true
	}
	return strings.Join(out, "\n"), notices
}

func isNotice(line string) bool {
	t := strings.TrimSpace(line)
	return strings.HasPrefix(t, "Copyright") || t == "All rights reserved."
}

var spaceRE = regexp.MustCompile(`\s+`)

func spdxOf(text string) string {
	flat := spaceRE.ReplaceAllString(text, " ")
	has := func(s string) bool { return strings.Contains(flat, s) }
	switch {
	case len(strings.Split(strings.TrimSpace(text), "\n")) < 3 && strings.Contains(text, "http"):
		return "attribution reference"
	case has("Apache License") && has("Version 2.0"):
		return "Apache-2.0"
	case has("Mozilla Public License"):
		return "MPL-2.0"
	case has("Permission is hereby granted, free of charge"):
		return "MIT"
	case has("Redistribution and use in source and binary forms"):
		if has("Neither the name") || has("neither the name") {
			return "BSD-3-Clause"
		}

		return "BSD-2-Clause"
	case has("dedicated to the public domain") || has("disclaims copyright"):
		return "public domain"
	}
	return "see text"
}

// newLicFile interns text into a group, merging copyright notices of identical
// texts.
func newLicFile(label, owner, text string) *licFile {
	masked, notices := mask(text)
	g := byMasked[masked]
	if g == nil {
		g = &licGroup{masked: masked, spdx: spdxOf(masked)}
		byMasked[masked] = g
		groups = append(groups, g)
	}
	for _, n := range notices {
		if !contains(g.notices, n) {
			g.notices = append(g.notices, n)
		}
	}
	m := owner
	if label != "LICENSE" && label != "" {
		m = fmt.Sprintf("%s (`%s`)", owner, label)
	}
	seen := false
	for _, v := range g.members {
		if v.name == m {
			seen = true
			break
		}
	}
	if !seen {
		g.members = append(g.members, grpMember{name: m, notices: notices})
	}
	return &licFile{label: label, text: text, notices: notices, grp: g}
}

func contains(s []string, v string) bool {
	for _, w := range s {
		if w == v {
			return true
		}
	}
	return false
}

// render returns the group's text with the merged notices back in place.
func (g *licGroup) render() string {
	if !strings.Contains(g.masked, noticeMark) {
		return g.masked
	}

	return strings.Replace(g.masked, noticeMark, strings.Join(g.notices, "\n"), 1)
}

// ---------------------------------------------------------------- collection

// scan finds the license and notice files of a module directory.
func scan(dir string) (lic []string, not []string) {
	base := filepath.Clean(dir)
	filepath.WalkDir(base, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(base, path)
		if d.IsDir() {
			if path == base {
				return nil
			}

			if skipDirs[d.Name()] || strings.Count(rel, string(filepath.Separator)) >= maxDepth {
				return fs.SkipDir
			}

			return nil
		}

		switch {
		case noticeNameRE.MatchString(d.Name()):
			not = append(not, rel)
		case licNameRE.MatchString(d.Name()):
			lic = append(lic, rel)
		}
		return nil
	})
	sort.Strings(lic)
	sort.Strings(not)
	return lic, not
}

func mustRead(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return normalize(b)
}

func grep(path, expr string) string {
	re := regexp.MustCompile(expr)
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	m := re.FindSubmatch(b)
	if m == nil {
		log.Fatalf("%s: no match for %s", path, expr)
	}

	return string(m[1])
}

func collect() {
	mods := modules()
	var missing []string
	for _, m := range mods {
		if !m.Main && m.Dir == "" {
			missing = append(missing, m.Path+"@"+m.Version)
		}
	}
	if len(missing) != 0 {
		// Modules reachable in the graph but not needed for any build are not
		// in the module cache. Their license still has to be read.
		dirs := download(missing)
		for i, m := range mods {
			if d := dirs[m.Path+"@"+m.Version]; d != "" {
				mods[i].Dir = d
			}
		}
	}

	linked := modSet("list", "-deps", "-f", "{{if .Module}}{{.Module.Path}}{{end}}", ".")
	tested := modSet("list", "-deps", "-test", "-f", "{{if .Module}}{{.Module.Path}}{{end}}", "./...")
	req := requirers()

	for _, m := range mods {
		if m.Main {
			continue
		}

		if m.Dir == "" {
			log.Fatalf("%s@%s: not in the module cache", m.Path, m.Version)
		}

		sec := secGraph
		switch {
		case linked[m.Path]:
			sec = secLinked
		case tested[m.Path]:
			sec = secTest
		}
		c := &component{
			name:    m.Path,
			version: m.Version,
			url:     "https://pkg.go.dev/" + m.Path,
			via:     viaOf(m.Path, req),
			sec:     sec,
			kind:    "module",
			purl:    "pkg:golang/" + m.Path + "@" + m.Version,
		}
		lic, not := scan(m.Dir)
		for _, v := range lic {
			c.files = append(c.files, newLicFile(filepath.ToSlash(v), m.Path, mustRead(filepath.Join(m.Dir, v))))
		}
		sections[sec] = append(sections[sec], c)
		for _, v := range not {
			inh := &inherited{
				owner: m.Path,
				label: filepath.ToSlash(v),
				body:  mustRead(filepath.Join(m.Dir, v)),
			}
			inh.parse()
			inherits = append(inherits, inh)
			// The upstreams reached through that module belong in the same
			// section as the module itself: they are as linked as it is.
			for _, e := range inh.entries {
				sections[sec] = append(sections[sec], &component{
					name: e.name,
					url:  e.url,
					via:  m.Path,
					sec:  sec,
					kind: "inherited",
					inh:  e,
				})
			}
		}
	}

	for i := range sections {
		s := sections[i]
		sort.SliceStable(s, func(a, b int) bool {
			return strings.ToLower(s[a].name) < strings.ToLower(s[b].name)
		})
	}

	collectVendored()
	collectRepo()
}

func viaOf(path string, req map[string][]string) string {
	v := req[path]
	if len(v) == 0 {
		return "`go.mod`"
	}

	var s []string
	for _, w := range v {
		if w == "modernc.org/sqlite" {
			s = append(s, "`go.mod`")
			continue
		}

		s = append(s, w)
	}
	return strings.Join(s, ", ")
}

// collectVendored adds the third-party C that was transpiled into this
// repository. It is not a Go module, so nothing in the module graph names it,
// yet it is the largest single body of third-party code we ship.
func collectVendored() {
	sqliteVer := grep(filepath.Join(root, "lib", "sqlite.go"), `(?m)^const SQLITE_VERSION = "([^"]+)"`)
	vecVer := grep(filepath.Join(root, "vec", "vec.go"), `(?m)^const m_SQLITE_VEC_VERSION = "([^"]+)"`)
	sqliteText := mustRead(filepath.Join(root, "LICENSE-SQLITE"))

	add := func(c *component) {
		c.sec = secLinked
		c.kind = "vendored"
		sections[secLinked] = append(sections[secLinked], c)
	}

	c := &component{
		name:    "SQLite",
		version: sqliteVer,
		url:     "https://sqlite.org/",
		via:     "transpiled into `lib/`",
		purl:    "pkg:generic/sqlite@" + sqliteVer,
		note:    "The SQLite C amalgamation, translated to Go by `modernc.org/ccgo`. Local copy of the upstream dedication: `LICENSE-SQLITE`.",
	}
	c.files = append(c.files, newLicFile("LICENSE-SQLITE", "SQLite "+sqliteVer, sqliteText))
	add(c)

	c = &component{
		name:    "sqlite-vec",
		version: vecVer,
		url:     "https://github.com/asg017/sqlite-vec",
		via:     "transpiled into `vec/`",
		purl:    "pkg:github/asg017/sqlite-vec@" + vecVer,
		note:    "Vector-search extension, linked only when `modernc.org/sqlite/vec` is imported. Local copy: `LICENSE-SQLITE_VEC`.",
	}
	c.files = append(c.files, newLicFile("LICENSE-SQLITE_VEC", "sqlite-vec "+vecVer, mustRead(filepath.Join(root, "LICENSE-SQLITE_VEC"))))
	add(c)

	c = &component{
		name:    "SQLite `test_demovfs.c`",
		version: "--",
		url:     "https://www.sqlite.org/src/doc/trunk/src/test_demovfs.c",
		via:     "`vfs/c/vfs.c`, transpiled into `vfs/`",
		note:    "Starting point of the Go-fs VFS bridge, linked only when `modernc.org/sqlite/vfs` is imported. Carries SQLite's own copyright disclaimer.",
	}
	c.files = append(c.files, newLicFile("", "SQLite `test_demovfs.c`", sqliteText))
	add(c)
}

func collectRepo() {
	sqliteText := mustRead(filepath.Join(root, "LICENSE-SQLITE"))
	suite := newLicFile("", "SQLite test suite (`testdata/`)", sqliteText)
	repoItems = append(repoItems,
		&repoItem{
			material: "SQLite TCL test suite",
			location: "`testdata/tcl/`, `testdata/overlay/`, `testdata/mptest.c`",
			origin:   "SQLite",
			url:      "https://sqlite.org/",
			file:     suite,
		},
		&repoItem{
			material: "Sponsor logos and word marks",
			location: "`sponsors/`",
			origin:   "the respective sponsors",
			url:      "",
		},
	)
	// The test suite is third-party code in the tree, so it is an SBOM
	// component too, scoped as compiled into nothing. Criticizing module-graph
	// SBOMs for omitting what they cannot see is only honest if this one omits
	// nothing it can see.
	repoComps = append(repoComps, &component{
		name:    "SQLite TCL test suite",
		version: "",
		url:     "https://sqlite.org/",
		via:     "`testdata/`",
		note:    "SQLite's own TCL and C test suite, vendored under `testdata/`. Present in a clone and in the module zip; compiled into nothing.",
		sec:     secGraph,
		kind:    "repo",
		files:   []*licFile{suite},
	})
}

// ------------------------------------------------------------------ rendering

func fence(text string) string {
	n := 0
	run := 0
	for _, r := range text {
		if r == '`' {
			run++
			if run > n {
				n = run
			}
			continue
		}

		run = 0
	}
	if n < 3 {
		n = 3
	} else {
		n++
	}
	return strings.Repeat("`", n)
}

func demote(body string) string {
	lines := strings.Split(body, "\n")
	for i, v := range lines {
		if headingRE.MatchString(v) {
			lines[i] = "###" + v
		}
	}
	return strings.Join(lines, "\n")
}

func (c *component) textRef() string {
	if c.inh != nil {
		return fmt.Sprintf("[%s](#%s)", c.inh.doc.id, c.inh.doc.id)
	}
	var s []string
	for _, f := range c.files {
		if v := fmt.Sprintf("[%s](#%s)", f.grp.id, f.grp.id); !contains(s, v) {
			s = append(s, v)
		}
	}
	if len(s) == 0 {
		return "--"
	}

	return strings.Join(s, ", ")
}

func (c *component) licenses() string {
	if c.inh != nil {
		return c.inh.spdx
	}

	var s []string
	for _, f := range c.files {
		if !contains(s, f.grp.spdx) {
			s = append(s, f.grp.spdx)
		}
	}
	if len(s) == 0 {
		return "**none found**"
	}

	return strings.Join(s, ", ")
}

func generate() string {
	// Identifiers are assigned in the order the reader meets them.
	n := 0
	for _, sec := range sections {
		for _, c := range sec {
			for _, f := range c.files {
				if f.grp.id == "" {
					n++
					f.grp.id = fmt.Sprintf("L%d", n)
				}
			}
		}
	}
	for _, it := range repoItems {
		if it.file != nil && it.file.grp.id == "" {
			n++
			it.file.grp.id = fmt.Sprintf("L%d", n)
		}
	}
	for i, inh := range inherits {
		inh.id = fmt.Sprintf("N%d", i+1)
	}
	var used []*licGroup
	for _, g := range groups {
		if g.id != "" {
			used = append(used, g)
		}
	}
	sort.Slice(used, func(i, j int) bool { return idNum(used[i].id) < idNum(used[j].id) })
	groups = used

	w := &strings.Builder{}
	p := func(f string, args ...any) { fmt.Fprintf(w, f+"\n", args...) }

	p("# Third-Party Licenses")
	p("")
	p("`modernc.org/sqlite` is BSD-3-Clause; that license is in [LICENSE](LICENSE) and it is")
	p("not repeated here. This file accounts for everything else -- every other body of code,")
	p("of any origin, that reaches you through this repository.")
	p("")
	p("The list is **transitively flattened**. It is not the direct dependencies of")
	p("`go.mod`: it is the whole module graph, `go list -m all`, plus the third-party C that")
	p("was transpiled into Go and committed here, plus the upstream notices that dependencies")
	p("carry in turn -- musl libc, for one, reaches you through `modernc.org/libc` and appears")
	p("below on its own account. Each component is classified by what it actually does to a")
	p("program that imports this driver, because \"a dependency\" and \"code in your binary\"")
	p("are not the same question:")
	p("")
	for i := range sections {
		p("- [%s](#%s)", sectionTitle[i], anchor(sectionTitle[i]))
	}
	p("- [4. Third-party material in the repository, outside any build](#4-third-party-material-in-the-repository-outside-any-build)")
	p("- [Appendix A: license texts](#appendix-a-license-texts)")
	if len(inherits) != 0 {
		p("- [Appendix B: inherited third-party notices](#appendix-b-inherited-third-party-notices)")
	}
	p("")
	p("Every license text referenced by the tables is reproduced in full in Appendix A. Where")
	p("several components share one text, it appears once, with every copyright notice that")
	p("came with it merged in place -- nothing is summarized, truncated or linked away.")
	p("")
	p("<!--")
	p("Generated by licensegen. Do not edit by hand: run `make licenses`.")
	p("The output is deterministic and carries no timestamp, so a diff here means the tree's")
	p("third-party content really changed.")
	p("-->")

	for i := range sections {
		p("")
		p("## %s", sectionTitle[i])
		p("")
		p("%s", wrap(sectionIntro[i]))
		p("")
		p("| Component | Version | License | Reached via | Text |")
		p("| --- | --- | --- | --- | --- |")
		for _, c := range sections[i] {
			p("| %s | %s | %s | %s | %s |", c.display(), dash(c.version), c.licenses(), dash(c.via), c.textRef())
		}
		var notes []*component
		for _, c := range sections[i] {
			if c.note != "" {
				notes = append(notes, c)
			}
		}
		if len(notes) != 0 {
			p("")
			for _, c := range notes {
				p("- **%s** -- %s", c.name, c.note)
			}
		}
	}

	p("")
	p("## 4. Third-party material in the repository, outside any build")
	p("")
	p("%s", wrap("Present in a clone or a source tarball of this repository, compiled into nothing. "+
		"Listed so that a redistribution of the source tree is covered as well as a redistribution of a binary."))
	p("")
	p("| Material | Location | Origin | License | Text |")
	p("| --- | --- | --- | --- | --- |")
	for _, it := range repoItems {
		lic, ref := "--", "--"
		if it.file != nil {
			lic = it.file.grp.spdx
			ref = fmt.Sprintf("[%s](#%s)", it.file.grp.id, it.file.grp.id)
		}
		origin := it.origin
		if it.url != "" {
			origin = fmt.Sprintf("[%s](%s)", it.origin, it.url)
		}
		p("| %s | %s | %s | %s | %s |", it.material, it.location, origin, lic, ref)
	}
	p("")
	p("%s", wrap("The sponsor logos and word marks in `sponsors/` are the trademarks of their owners. "+
		"They are reproduced to acknowledge sponsorship, are not licensed under this project's license, "+
		"and confer no right to use those marks."))
	p("")
	p("%s", wrap("`vendor_libs/` is a second, self-contained Go module holding the regeneration tool "+
		"invoked by `make vendor`. Its own dependencies are a subset of the modules already listed above."))

	p("")
	p("## Appendix A: license texts")
	p("")
	p("%s", wrap("Verbatim. The only editing is the merging of copyright notices where one text is "+
		"shared by several components; where a component's text is unique, it is byte-for-byte the file "+
		"shipped upstream."))
	for _, g := range groups {
		if g.id == "" {
			continue
		}

		p("")
		p("<a id=\"%s\"></a>", g.id)
		p("")
		p("### %s: %s", g.id, g.spdx)
		p("")
		p("Applies to:")
		p("")
		for _, m := range g.members {
			switch {
			case len(m.notices) == 0:
				p("- %s", m.name)
			default:
				p("- %s -- %s", m.name, strings.Join(m.notices, " "))
			}
		}
		p("")
		f := fence(g.render())
		p("%stext", f)
		p("%s", g.render())
		p("%s", f)
	}

	if len(inherits) != 0 {
		p("")
		p("## Appendix B: inherited third-party notices")
		p("")
		p("%s", wrap("These documents are the transitive tail: notices that a dependency carries for "+
			"code of its own that came from somewhere else. They are reproduced verbatim, with only the "+
			"Markdown heading levels shifted so they nest inside this file."))
		for _, inh := range inherits {
			p("")
			p("<a id=\"%s\"></a>", inh.id)
			p("")
			p("### %s: `%s` from %s", inh.id, inh.label, inh.owner)
			p("")
			var names []string
			for _, e := range inh.entries {
				names = append(names, fmt.Sprintf("%s (%s)", e.name, e.spdx))
			}
			p("Covers: %s.", strings.Join(names, ", "))
			p("")
			p("---")
			p("")
			p("%s", demote(inh.body))
			p("")
			p("---")
		}
	}
	return w.String()
}

func idNum(id string) int {
	n := 0
	fmt.Sscanf(id, "L%d", &n)
	return n
}

func dash(s string) string {
	if s == "" {
		return "--"
	}

	return s
}

func anchor(title string) string {
	s := strings.ToLower(title)
	s = regexp.MustCompile(`[^a-z0-9 -]`).ReplaceAllString(s, "")
	return strings.ReplaceAll(s, " ", "-")
}

// wrap breaks a paragraph at 88 columns so the Markdown source stays readable.
func wrap(s string) string {
	var out []string
	line := ""
	for _, w := range strings.Fields(s) {
		switch {
		case line == "":
			line = w
		case len(line)+1+len(w) > 88:
			out = append(out, line)
			line = w
		default:
			line += " " + w
		}
	}
	if line != "" {
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("licensegen: ")
	flag.Parse()
	abs, err := filepath.Abs(*oC)
	if err != nil {
		log.Fatal(err)
	}

	root = abs
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err != nil {
		log.Fatalf("%s: not a module root", root)
	}

	collect()
	// generate() assigns the appendix identifiers the SBOM's LicenseRef names
	// are built from, so it runs first.
	outs := []struct{ name, content string }{
		{*oOut, generate()},
		{*oMD, generateSBOMmd()},
		{*oCDX, generateCDX()},
		{*oSPDX, generateSPDX()},
	}
	stale := false
	for _, o := range outs {
		path := o.name
		if !filepath.IsAbs(path) {
			path = filepath.Join(root, path)
		}
		if *oCheck {
			old, err := os.ReadFile(path)
			switch {
			case err != nil:
				log.Printf("%v", err)
				stale = true
			case string(old) != o.content:
				log.Printf("%s is stale", o.name)
				stale = true
			}
			continue
		}

		if err := os.WriteFile(path, []byte(o.content), 0o644); err != nil {
			log.Fatal(err)
		}
	}
	if *oCheck {
		if stale {
			log.Fatal("run `make sbom`")
		}

		fmt.Println("license and SBOM documents are up to date")
		return
	}

	fmt.Printf("%d components, %d license texts, %d inherited notice files -> %s, %s, %s, %s\n",
		len(sections[0])+len(sections[1])+len(sections[2])+len(repoItems), len(groups), len(inherits),
		*oOut, *oMD, *oCDX, *oSPDX)
}
