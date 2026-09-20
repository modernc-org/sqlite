// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build none
// +build none

// SBOM generation: CycloneDX 1.6 and SPDX 2.3, from the same inventory the
// license document is built from.
//
// The point of generating these here rather than with a stock Go SBOM tool is
// that the largest component this module ships -- SQLite itself -- is C that
// was transpiled into Go and committed. No go.mod names it, so no tool reading
// the module graph can see it. Same for sqlite-vec, for the VFS bridge, and
// for the upstreams that modernc.org/libc carries notices for.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	rootName = "modernc.org/sqlite"
	rootRepo = "https://gitlab.com/cznic/sqlite"
	rootRef  = "pkg:golang/modernc.org/sqlite"
	rootSPDX = "SPDXRef-Package-modernc.org-sqlite"

	// SPDX requires a creation timestamp and CycloneDX offers one. A real
	// timestamp would change on every run and make the committed files differ
	// from a fresh generation for no reason, which is exactly what -check is
	// there to catch. The authoritative date of this document is the date of
	// the commit that last changed it.
	fixedTime = "1970-01-01T00:00:00Z"
)

var (
	spdxStandard = map[string]bool{
		"Apache-2.0":   true,
		"BSD-2-Clause": true,
		"BSD-3-Clause": true,
		"MIT":          true,
		"MPL-2.0":      true,
	}
	notIDRE = regexp.MustCompile(`[^A-Za-z0-9.-]+`)
)

// --------------------------------------------------------------- inventory

func allComponents() (r []*component) {
	for i := range sections {
		r = append(r, sections[i]...)
	}
	return append(r, repoComps...)
}

func sanitize(s string) string {
	return strings.Trim(notIDRE.ReplaceAllString(s, "-"), "-")
}

func (c *component) ref() string {
	if c.purl != "" {
		return c.purl
	}

	return c.kind + ":" + sanitize(c.name)
}

func (c *component) spdxID() string {
	return "SPDXRef-Package-" + sanitize(c.name+"-"+c.version)
}

// escapeModPath applies the module proxy's case encoding, so the download
// location is one a consumer can actually fetch.
func escapeModPath(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			b.WriteByte('!')
			b.WriteRune(r + ('a' - 'A'))
			continue
		}

		b.WriteRune(r)
	}
	return b.String()
}

func (c *component) downloadLocation() string {
	switch {
	case c.kind == "module":
		return fmt.Sprintf("https://proxy.golang.org/%s/@v/%s.zip", escapeModPath(c.name), c.version)
	case c.url != "":
		return c.url
	}
	return "NOASSERTION"
}

// licenseID maps a license group to an SPDX identifier, minting a LicenseRef
// for anything that is not a standard one. The LicenseRef names the appendix
// entry of LICENSE-3RD-PARTY.md that holds the text, and the text travels with
// the SBOM too, so neither document depends on the other.
func licenseID(g *licGroup) string {
	if spdxStandard[g.spdx] {
		return g.spdx
	}

	if strings.Contains(g.masked, "SQLite Is Public Domain") {
		return "LicenseRef-SQLite-PublicDomain"
	}

	return "LicenseRef-" + g.id
}

func (c *component) licenseGroups() (r []*licGroup) {
	seen := map[*licGroup]bool{}
	for _, f := range c.files {
		if !seen[f.grp] {
			seen[f.grp] = true
			r = append(r, f.grp)
		}
	}
	return r
}

// licenseExpr is the component's SPDX license expression. Several license files
// in one component means all of them apply, hence AND.
func (c *component) licenseExpr() string {
	if c.inh != nil {
		if spdxStandard[c.inh.spdx] {
			return c.inh.spdx
		}

		return "NOASSERTION"
	}

	var s []string
	for _, g := range c.licenseGroups() {
		if id := licenseID(g); !contains(s, id) {
			s = append(s, id)
		}
	}
	switch len(s) {
	case 0:
		return "NOASSERTION"
	case 1:
		return s[0]
	}
	return strings.Join(s, " AND ")
}

func (c *component) copyright() string {
	var s []string
	for _, f := range c.files {
		for _, n := range f.notices {
			if n = strings.TrimSpace(n); n != "" && !contains(s, n) {
				s = append(s, n)
			}
		}
	}
	if len(s) == 0 {
		return "NOASSERTION"
	}

	return strings.Join(s, "\n")
}

// scope tells a consumer what the component does to them. It is the single most
// useful field in this document and the one a module graph alone cannot fill in.
func (c *component) scope() (cdx, spdxRel string) {
	if c.kind == "repo" {
		return "excluded", "CONTAINS"
	}

	switch c.sec {
	case secLinked:
		if c.kind == "vendored" {
			return "required", "CONTAINS"
		}

		if c.kind == "inherited" {
			return "required", "CONTAINS"
		}

		return "required", "DEPENDS_ON"
	case secTest:
		return "optional", "TEST_DEPENDENCY_OF"
	}
	return "excluded", "OPTIONAL_DEPENDENCY_OF"
}

// extractedLicenses returns the non-standard licenses in use, for the documents
// that must carry their texts.
func extractedLicenses() (r []*licGroup) {
	seen := map[string]bool{}
	for _, c := range allComponents() {
		for _, g := range c.licenseGroups() {
			id := licenseID(g)
			if !spdxStandard[g.spdx] && !seen[id] {
				seen[id] = true
				r = append(r, g)
			}
		}
	}
	sort.Slice(r, func(i, j int) bool { return licenseID(r[i]) < licenseID(r[j]) })
	return r
}

// ---------------------------------------------------------------- CycloneDX

type cdxText struct {
	ContentType string `json:"contentType,omitempty"`
	Content     string `json:"content"`
}

type cdxLicense struct {
	ID   string   `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
	Text *cdxText `json:"text,omitempty"`
}

type cdxLicenseChoice struct {
	License *cdxLicense `json:"license,omitempty"`
}

type cdxExtRef struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type cdxProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type cdxPedigree struct {
	Notes string `json:"notes"`
}

type cdxComponent struct {
	Type         string             `json:"type"`
	BOMRef       string             `json:"bom-ref"`
	Name         string             `json:"name"`
	Version      string             `json:"version,omitempty"`
	Description  string             `json:"description,omitempty"`
	Scope        string             `json:"scope,omitempty"`
	PURL         string             `json:"purl,omitempty"`
	Licenses     []cdxLicenseChoice `json:"licenses,omitempty"`
	ExternalRefs []cdxExtRef        `json:"externalReferences,omitempty"`
	Pedigree     *cdxPedigree       `json:"pedigree,omitempty"`
	Properties   []cdxProperty      `json:"properties,omitempty"`
}

type cdxTools struct {
	Components []cdxComponent `json:"components"`
}

type cdxMetadata struct {
	Component  cdxComponent  `json:"component"`
	Tools      cdxTools      `json:"tools"`
	Properties []cdxProperty `json:"properties,omitempty"`
}

type cdxDependency struct {
	Ref       string   `json:"ref"`
	DependsOn []string `json:"dependsOn,omitempty"`
}

type cdxBOM struct {
	Format       string          `json:"bomFormat"`
	SpecVersion  string          `json:"specVersion"`
	Version      int             `json:"version"`
	Metadata     cdxMetadata     `json:"metadata"`
	Components   []cdxComponent  `json:"components"`
	Dependencies []cdxDependency `json:"dependencies"`
}

func cdxLicensesOf(c *component) (r []cdxLicenseChoice) {
	if c.inh != nil {
		if spdxStandard[c.inh.spdx] {
			return []cdxLicenseChoice{{License: &cdxLicense{ID: c.inh.spdx}}}
		}

		return nil
	}

	for _, g := range c.licenseGroups() {
		id := licenseID(g)
		if spdxStandard[g.spdx] {
			r = append(r, cdxLicenseChoice{License: &cdxLicense{ID: id}})
			continue
		}

		r = append(r, cdxLicenseChoice{License: &cdxLicense{
			Name: id,
			Text: &cdxText{ContentType: "text/plain", Content: g.render()},
		}})
	}
	return r
}

func generateCDX() string {
	comps := allComponents()
	sort.Slice(comps, func(i, j int) bool { return comps[i].ref() < comps[j].ref() })

	bom := cdxBOM{
		Format:      "CycloneDX",
		SpecVersion: "1.6",
		Version:     1,
		Metadata: cdxMetadata{
			Component: cdxComponent{
				Type:        "library",
				BOMRef:      rootRef,
				Name:        rootName,
				Description: "Pure-Go SQLite driver, no cgo. Describes the source tree; the release version is applied by the maintainer when a tag is pushed.",
				PURL:        rootRef,
				ExternalRefs: []cdxExtRef{
					{Type: "vcs", URL: rootRepo},
					{Type: "distribution", URL: "https://pkg.go.dev/" + rootName},
				},
			},
			Tools: cdxTools{Components: []cdxComponent{{
				Type:   "application",
				BOMRef: "licensegen",
				Name:   "licensegen",
			}}},
			Properties: []cdxProperty{
				{Name: "modernc:generator", Value: "licensegen, in this repository; run make sbom"},
				{Name: "modernc:reproducible", Value: "no timestamp and no serial number, so regeneration of an unchanged tree is byte-identical"},
				{Name: "modernc:notes", Value: "See SBOM.md. Components whose C was transpiled into Go are marked with a pedigree note; no Go module graph names them."},
			},
		},
	}

	for _, c := range comps {
		scope, _ := c.scope()
		cc := cdxComponent{
			Type:        "library",
			BOMRef:      c.ref(),
			Name:        c.name,
			Version:     c.version,
			Description: strings.TrimSpace(stripMarkdown(c.note)),
			Scope:       scope,
			PURL:        c.purl,
			Licenses:    cdxLicensesOf(c),
		}
		if c.url != "" {
			cc.ExternalRefs = append(cc.ExternalRefs, cdxExtRef{Type: "website", URL: c.url})
		}
		if c.kind == "module" {
			cc.ExternalRefs = append(cc.ExternalRefs, cdxExtRef{Type: "distribution", URL: c.downloadLocation()})
		}
		switch c.kind {
		case "repo":
			cc.Pedigree = &cdxPedigree{Notes: "Third-party material committed to this repository. It is compiled into nothing; it ships because it is in the tree."}
		case "vendored":
			cc.Pedigree = &cdxPedigree{Notes: "C source transpiled to Go by modernc.org/ccgo and committed to this repository. It is not a Go module dependency and appears in no module graph."}
		case "inherited":
			cc.Pedigree = &cdxPedigree{Notes: "Reached through " + c.via + ", which carries the upstream notice. The upstream revision is not recorded there, hence no version."}
		}
		cc.Properties = append(cc.Properties, cdxProperty{Name: "modernc:reachedVia", Value: stripMarkdown(c.via)})
		bom.Components = append(bom.Components, cc)
	}

	// Dependencies: the real module graph, plus containment for what was
	// vendored or inherited, so nothing in the inventory is left unattached.
	byPath := map[string]*component{}
	for _, c := range comps {
		if c.kind == "module" {
			byPath[c.name] = c
		}
	}
	root := cdxDependency{Ref: rootRef}
	var rest []cdxDependency
	for _, c := range comps {
		switch c.kind {
		case "module":
			if strings.Contains(c.via, "go.mod") {
				root.DependsOn = append(root.DependsOn, c.ref())
			}
			var on []string
			for _, dep := range modDeps[c.name] {
				if d := byPath[dep]; d != nil {
					on = append(on, d.ref())
				}
			}
			sort.Strings(on)
			rest = append(rest, cdxDependency{Ref: c.ref(), DependsOn: on})
		case "vendored", "repo":
			root.DependsOn = append(root.DependsOn, c.ref())
		case "inherited":
			if d := byPath[c.via]; d != nil {
				for i := range rest {
					if rest[i].Ref == d.ref() {
						rest[i].DependsOn = append(rest[i].DependsOn, c.ref())
						sort.Strings(rest[i].DependsOn)
					}
				}
			}
		}
	}
	sort.Strings(root.DependsOn)
	bom.Dependencies = append([]cdxDependency{root}, rest...)

	b, err := json.MarshalIndent(bom, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b) + "\n"
}

// --------------------------------------------------------------------- SPDX

type spdxCreationInfo struct {
	Created  string   `json:"created"`
	Creators []string `json:"creators"`
	Comment  string   `json:"comment,omitempty"`
}

type spdxExternalRef struct {
	Category string `json:"referenceCategory"`
	Type     string `json:"referenceType"`
	Locator  string `json:"referenceLocator"`
}

type spdxPackage struct {
	SPDXID           string            `json:"SPDXID"`
	Name             string            `json:"name"`
	VersionInfo      string            `json:"versionInfo,omitempty"`
	DownloadLocation string            `json:"downloadLocation"`
	FilesAnalyzed    bool              `json:"filesAnalyzed"`
	Homepage         string            `json:"homepage,omitempty"`
	LicenseConcluded string            `json:"licenseConcluded"`
	LicenseDeclared  string            `json:"licenseDeclared"`
	CopyrightText    string            `json:"copyrightText"`
	Comment          string            `json:"comment,omitempty"`
	ExternalRefs     []spdxExternalRef `json:"externalRefs,omitempty"`
}

type spdxExtractedLicense struct {
	LicenseID     string `json:"licenseId"`
	Name          string `json:"name"`
	ExtractedText string `json:"extractedText"`
}

type spdxRelationship struct {
	SPDXElementID      string `json:"spdxElementId"`
	RelationshipType   string `json:"relationshipType"`
	RelatedSPDXElement string `json:"relatedSpdxElement"`
}

type spdxDoc struct {
	SPDXVersion       string                 `json:"spdxVersion"`
	DataLicense       string                 `json:"dataLicense"`
	SPDXID            string                 `json:"SPDXID"`
	Name              string                 `json:"name"`
	DocumentNamespace string                 `json:"documentNamespace"`
	CreationInfo      spdxCreationInfo       `json:"creationInfo"`
	DocumentDescribes []string               `json:"documentDescribes"`
	Packages          []spdxPackage          `json:"packages"`
	ExtractedLicenses []spdxExtractedLicense `json:"hasExtractedLicensingInfos,omitempty"`
	Relationships     []spdxRelationship     `json:"relationships"`
}

func generateSPDX() string {
	comps := allComponents()
	sort.Slice(comps, func(i, j int) bool { return comps[i].spdxID() < comps[j].spdxID() })

	doc := spdxDoc{
		SPDXVersion: "SPDX-2.3",
		DataLicense: "CC0-1.0",
		SPDXID:      "SPDXRef-DOCUMENT",
		Name:        rootName,
		CreationInfo: spdxCreationInfo{
			Created:  fixedTime,
			Creators: []string{"Tool: licensegen", "Organization: The Sqlite Authors"},
			Comment: "Generated by licensegen in the modernc.org/sqlite repository; run make sbom. " +
				"The creation timestamp is fixed so that regenerating an unchanged tree is byte-identical " +
				"and a staleness check is meaningful; the authoritative date is that of the commit which " +
				"last changed this file. The document namespace is derived from the document's own content. " +
				"See SBOM.md.",
		},
		DocumentDescribes: []string{rootSPDX},
		Packages: []spdxPackage{{
			SPDXID:           rootSPDX,
			Name:             rootName,
			VersionInfo:      "NOASSERTION",
			DownloadLocation: rootRepo,
			FilesAnalyzed:    false,
			Homepage:         "https://pkg.go.dev/" + rootName,
			LicenseConcluded: "BSD-3-Clause",
			LicenseDeclared:  "BSD-3-Clause",
			CopyrightText:    "Copyright (c) 2017 The Sqlite Authors. All rights reserved.",
			Comment:          "Describes the source tree. The release version is applied by the maintainer when a tag is pushed, so no version is asserted here.",
			ExternalRefs: []spdxExternalRef{{
				Category: "PACKAGE-MANAGER",
				Type:     "purl",
				Locator:  rootRef,
			}},
		}},
		Relationships: []spdxRelationship{{
			SPDXElementID:      "SPDXRef-DOCUMENT",
			RelationshipType:   "DESCRIBES",
			RelatedSPDXElement: rootSPDX,
		}},
	}

	for _, g := range extractedLicenses() {
		name := g.spdx
		if name == "" {
			name = "Unknown"
		}
		doc.ExtractedLicenses = append(doc.ExtractedLicenses, spdxExtractedLicense{
			LicenseID:     licenseID(g),
			Name:          name + " (" + g.id + " in LICENSE-3RD-PARTY.md)",
			ExtractedText: g.render(),
		})
	}

	for _, c := range comps {
		_, rel := c.scope()
		version := c.version
		if version == "" || version == "--" {
			version = "NOASSERTION"
		}
		p := spdxPackage{
			SPDXID:           c.spdxID(),
			Name:             c.name,
			VersionInfo:      version,
			DownloadLocation: c.downloadLocation(),
			FilesAnalyzed:    false,
			Homepage:         c.url,
			LicenseConcluded: c.licenseExpr(),
			LicenseDeclared:  c.licenseExpr(),
			CopyrightText:    c.copyright(),
			Comment:          spdxComment(c),
		}
		if c.purl != "" {
			p.ExternalRefs = append(p.ExternalRefs, spdxExternalRef{
				Category: "PACKAGE-MANAGER",
				Type:     "purl",
				Locator:  c.purl,
			})
		}
		doc.Packages = append(doc.Packages, p)

		owner := rootSPDX
		if c.kind == "inherited" {
			for _, o := range comps {
				if o.kind == "module" && o.name == c.via {
					owner = o.spdxID()
					break
				}
			}
		}
		switch rel {
		case "DEPENDS_ON", "CONTAINS":
			doc.Relationships = append(doc.Relationships, spdxRelationship{
				SPDXElementID:      owner,
				RelationshipType:   rel,
				RelatedSPDXElement: c.spdxID(),
			})
		default:
			doc.Relationships = append(doc.Relationships, spdxRelationship{
				SPDXElementID:      c.spdxID(),
				RelationshipType:   rel,
				RelatedSPDXElement: rootSPDX,
			})
		}
	}

	// The namespace must be unique per document. Deriving it from the content
	// keeps that true without a UUID, which would change on every run.
	doc.DocumentNamespace = ""
	b, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		panic(err)
	}

	sum := sha256.Sum256(b)
	doc.DocumentNamespace = rootRepo + "/-/spdx/" + hex.EncodeToString(sum[:])
	if b, err = json.MarshalIndent(doc, "", "  "); err != nil {
		panic(err)
	}

	return string(b) + "\n"
}

func spdxComment(c *component) string {
	switch c.kind {
	case "repo":
		return stripMarkdown(c.note)
	case "vendored":
		return "C source transpiled to Go by modernc.org/ccgo and committed to this repository; it appears in no Go module graph. " + stripMarkdown(c.note)
	case "inherited":
		return "Reached through " + c.via + ", which carries the upstream notice. That notice does not record the upstream revision, so no version is asserted."
	}
	return "Reached via " + stripMarkdown(c.via) + "."
}

func stripMarkdown(s string) string {
	return strings.ReplaceAll(s, "`", "")
}

// ------------------------------------------------------------------ SBOM.md

func generateSBOMmd() string {
	w := &strings.Builder{}
	p := func(f string, args ...any) { fmt.Fprintf(w, f+"\n", args...) }

	var linked, test, graph int
	for _, c := range sections[secLinked] {
		_ = c
		linked++
	}
	test = len(sections[secTest])
	graph = len(sections[secGraph])

	p("# Software Bill of Materials")
	p("")
	p("%s", wrap("This repository ships a machine-readable SBOM in both common formats, generated from "+
		"the same inventory as [LICENSE-3RD-PARTY.md](LICENSE-3RD-PARTY.md):"))
	p("")
	p("| File | Format | For |")
	p("| --- | --- | --- |")
	p("| [`sbom.cdx.json`](sbom.cdx.json) | CycloneDX 1.6 | Scanners and policy engines. Carries `pedigree` notes and per-component `scope`. |")
	p("| [`sbom.spdx.json`](sbom.spdx.json) | SPDX 2.3 | License and provenance review. Carries precise relationships and the full text of every non-standard license. |")
	p("| [`LICENSE-3RD-PARTY.md`](LICENSE-3RD-PARTY.md) | Markdown | People. Full license texts, transitively flattened. |")
	p("")
	p("This page is generated too. Regenerate all four documents with `make sbom`; `licgen -check`")
	p("writes nothing and fails if any of them is stale, which is the form to run in CI.")
	p("")
	p("## Why this is not what a stock Go SBOM tool produces")
	p("")
	p("%s", wrap("Run `cyclonedx-gomod` or `syft` against this module and the result will be a clean, "+
		"complete-looking list of Go modules that is missing the single largest thing the module ships. "+
		"`lib/` is SQLite itself: the C amalgamation, transpiled to Go by `modernc.org/ccgo` and committed "+
		"here. It is not a dependency, it has no `go.mod` entry, and nothing in the module graph names it. "+
		"The same is true of `sqlite-vec` in `vec/` and of the VFS bridge in `vfs/`."))
	p("")
	p("%s", wrap("A consumer asking \"does this project ship SQLite 3.53.4, and am I exposed to a CVE "+
		"against it\" gets the wrong answer from a module-graph SBOM: not a cautious answer, a wrong one. "+
		"These documents name it, version it, and mark it with a pedigree note saying where it came from."))
	p("")
	p("%s", wrap("The other direction matters too. `modernc.org/libc` carries notices for upstreams "+
		"vendored into it -- musl libc among them -- which no module graph reports either. Those are listed "+
		"as components in their own right, related to `modernc.org/libc` by containment."))
	p("")
	p("## What the scopes mean")
	p("")
	p("%s", wrap("Every component is classified by what it does to a program that imports this driver. "+
		"This is the field worth reading first, and the one a module list cannot fill in:"))
	p("")
	p("| Components | CycloneDX `scope` | SPDX relationship | Meaning |")
	p("| --- | --- | --- | --- |")
	p("| %d | `required` | `DEPENDS_ON` / `CONTAINS` | Linked into your binary. Their licenses are the ones you carry when you redistribute. |", linked)
	p("| %d | `optional` | `TEST_DEPENDENCY_OF` | Compiled into this repository's own tests only. Never reaches your program. |", test)
	p("| %d | `excluded` | `OPTIONAL_DEPENDENCY_OF` | Present in `go list -m all`, compiled into nothing here. Reported because `go.sum` and module-graph scanners will surface them. |", graph)
	p("| %d | `excluded` | `CONTAINS` | Third-party material committed to this repository and shipped in the module zip, compiled into nothing. |", len(repoComps))
	p("")
	p("%s", wrap("The distinction is not academic. The only copyleft-adjacent license anywhere in this "+
		"project, MPL-2.0, sits in the third group: `github.com/hashicorp/golang-lru/v2`, reached through "+
		"`modernc.org/libc`'s `go.mod` and in no import closure. A scanner that reads the module graph will "+
		"flag it; these documents say plainly that nothing of it is compiled in."))
	p("")
	p("## Inventory")
	p("")
	p("%s", wrap("Full license texts are in [LICENSE-3RD-PARTY.md](LICENSE-3RD-PARTY.md); a `LicenseRef-` "+
		"identifier below names the appendix entry there that holds the text, and the text is embedded in "+
		"both JSON documents as well."))
	for i := range sections {
		p("")
		p("### %s", sectionTitle[i])
		p("")
		p("| Component | Version | License | Package URL |")
		p("| --- | --- | --- | --- |")
		cs := append([]*component(nil), sections[i]...)
		sort.SliceStable(cs, func(a, b int) bool {
			return strings.ToLower(cs[a].name) < strings.ToLower(cs[b].name)
		})
		for _, c := range cs {
			p("| %s | %s | `%s` | %s |", c.name, dash(c.version), c.licenseExpr(), code(c.purl))
		}
	}
	p("")
	p("### 4. Third-party material in the repository, outside any build")
	p("")
	p("| Component | Version | License | Package URL |")
	p("| --- | --- | --- | --- |")
	for _, c := range repoComps {
		p("| %s | %s | `%s` | %s |", c.name, dash(c.version), c.licenseExpr(), code(c.purl))
	}
	p("")
	p("## What these documents do not claim")
	p("")
	p("- %s", wrap1("No file-level hashes or a package verification code: `filesAnalyzed` is `false` "+
		"throughout. The documents describe components, not a file manifest."))
	p("- %s", wrap1("No build attestation or signature. These say what is in the tree, not who built "+
		"a given binary from it."))
	p("- %s", wrap1("The upstreams reached through a dependency's notices carry no version. The notice "+
		"names the project, not the revision that was vendored, and guessing one would be worse than saying so."))
	p("- %s", wrap1("The subject's own version is `NOASSERTION`. The document describes a source tree; "+
		"the release tag is applied later, by hand, once the builders are green."))
	p("- %s", wrap1("Timestamps are fixed at the Unix epoch and the SPDX document namespace is derived "+
		"from the document's content. Both are deliberate: a regenerated document for an unchanged tree is "+
		"byte-identical, which is what makes `licgen -check` meaningful. The real date of the document is "+
		"the date of the commit that last changed it."))
	p("")
	p("<!--")
	p("Generated by licensegen. Do not edit by hand: run `make sbom`.")
	p("-->")
	return w.String()
}

func code(s string) string {
	if s == "" {
		return "--"
	}

	return "`" + s + "`"
}

// wrap1 wraps a bullet's continuation lines under its text.
func wrap1(s string) string {
	return strings.ReplaceAll(wrap(s), "\n", "\n  ")
}
