// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"testing"
)

const testShift = 4096

func testJournal(t testing.TB, st store, known ...string) *journalFile {
	t.Helper()
	return &journalFile{
		st:      st,
		c:       testCipher(t, labelJournal),
		cl:      testCipher(t, labelJournalLength),
		check:   keyCheck(testCipher(t, labelKeyCheck)),
		shift:   testShift,
		maxName: 512,
		known: func(name []byte) bool {
			for _, k := range known {
				if string(name) == k {
					return true
				}
			}

			return false
		},
	}
}

// superRecord is what SQLite's writeSuperJournal appends to a journal: a page
// number no page has, the name, its length, its checksum, the magic. The
// checksum is the sum of the name's bytes as signed or as unsigned chars.
func superRecord(name string, signed bool) []byte {
	var sum uint32
	for _, b := range []byte(name) {
		if signed {
			sum += uint32(int32(int8(b)))
		} else {
			sum += uint32(b)
		}
	}

	rec := binary.BigEndian.AppendUint32(nil, 0x40000000/4096+1)
	rec = append(rec, name...)
	rec = binary.BigEndian.AppendUint32(rec, uint32(len(name)))
	rec = binary.BigEndian.AppendUint32(rec, sum)
	return append(rec, journalMagic...)
}

// plainSuper reads the name of a super-journal from the end of a file the
// way SQLite's readSuperJournal does, here from the stored bytes: what a
// reader without the key sees.
func plainSuper(b []byte) string {
	sz := len(b)
	if sz < 16 || !bytes.Equal(b[sz-8:], journalMagic) {
		return ""
	}

	n := int(binary.BigEndian.Uint32(b[sz-16:]))
	if n == 0 || n > sz-16 {
		return ""
	}

	name := b[sz-16-n : sz-16]
	sum := binary.BigEndian.Uint32(b[sz-12:])
	for _, c := range name {
		sum -= uint32(int32(int8(c)))
	}

	if sum != 0 {
		// The other kind of char.
		sum = binary.BigEndian.Uint32(b[sz-12:])
		for _, c := range name {
			sum -= uint32(c)
		}
	}

	if sum != 0 {
		return ""
	}

	return string(name)
}

// ownerSuper does the same through the journal, as its owner does.
func ownerSuper(t *testing.T, f *journalFile) string {
	t.Helper()
	sz, err := f.size()
	if err != nil {
		t.Fatal(err)
	}

	b := make([]byte, sz)
	if _, err := f.read(b, 0); err != nil {
		t.Fatal(err)
	}

	return plainSuper(b)
}

func mustWrite(t *testing.T, f file, p []byte, off int64) {
	t.Helper()
	if err := f.write(p, off); err != nil {
		t.Fatalf("write %d bytes at %d: %v", len(p), off, err)
	}
}

func mustSize(t *testing.T, f file, want int64) {
	t.Helper()
	if got, err := f.size(); err != nil || got != want {
		t.Fatalf("size %d, %v; want %d", got, err, want)
	}
}

// The journal layout against a plain file: writes at any offset, with gaps,
// truncations, reads and sizes to the byte, and the form of the store after
// every step.
func TestJournalModel(t *testing.T) {
	rnd := rand.New(rand.NewSource(5))
	for round := range 300 {
		st := &memStore{}
		f := testJournal(t, st)
		var m []byte
		var written []bool // the bytes SQLite wrote; it reads no others
		const limit = 5 * unitSize
		for op := range 150 {
			what := fmt.Sprintf("round %d op %d", round, op)
			off, n := rnd.Int63n(limit), 1+rnd.Int63n(unitSize+100)
			switch rnd.Intn(4) {
			case 0: // as SQLite appends: at the end, small or a page
				off, n = int64(len(m)), []int64{4, 4, 8, 12, 28, 512, 4096}[rnd.Intn(7)]
			case 1: // a header at the next sector
				off, n = (int64(len(m))+511)/512*512, 512
			}

			switch k := rnd.Intn(12); {
			case k < 7:
				p := make([]byte, n)
				rnd.Read(p)
				before, sizeBefore := len(st.writes), int64(len(st.b))
				mustWrite(t, f, p, off)
				w := st.writes[before:]
				if len(w) != 1 {
					t.Fatalf("%s: a write reached the store as %d writes", what, len(w))
				}

				// No further back than the block SQLite writes in: what is
				// before it may have been synced.
				if at := testShift + roundDown(off); sizeBefore != 0 && w[0].off != at {
					t.Fatalf("%s: a write at %d reached the store at %d, want %d", what, off, w[0].off, at)
				}

				for int64(len(m)) < off+n {
					m, written = append(m, 0), append(written, false)
				}

				copy(m[off:], p)
				for i := off; i < off+n; i++ {
					written[i] = true
				}

			case k < 9:
				to := int64(0)
				if len(m) > 0 && rnd.Intn(5) > 0 {
					to = rnd.Int63n(int64(len(m)) + 1)
				}

				if err := f.truncate(to); err != nil {
					t.Fatalf("%s: truncate to %d of %d: %v", what, to, len(m), err)
				}

				m, written = m[:to], written[:to]

			default:
				p := bytes.Repeat([]byte{0xee}, int(n))
				got, err := f.read(p, off)
				if err != nil {
					t.Fatalf("%s: read: %v", what, err)
				}

				in := min(max(int64(len(m))-off, 0), n)
				if int64(got) < in {
					t.Fatalf("%s: read [%d, %d) of %d: %d bytes", what, off, off+n, len(m), got)
				}

				for i := int64(0); i < in; i++ {
					if written[off+i] && p[i] != m[off+i] {
						t.Fatalf("%s: byte %d reads %#x, written %#x", what, off+i, p[i], m[off+i])
					}
				}
			}

			mustSize(t, f, int64(len(m)))
			switch {
			case len(m) == 0 && len(st.b) != 0:
				t.Fatalf("%s: an empty journal is stored as %d bytes", what, len(st.b))
			case len(m) != 0:
				if want := testShift + roundUp(int64(len(m))) + blockSize; int64(len(st.b)) != want {
					t.Fatalf("%s: a journal of %d bytes is stored as %d, want %d", what, len(m), len(st.b), want)
				}

				if st.b[0] != 0 || !bytes.HasPrefix(st.b, prefixMagic) {
					t.Fatalf("%s: the store begins % x", what, st.b[:16])
				}
			}
		}

		// Nothing SQLite wrote is in the store as it was written.
		for off := 0; off+blockSize <= len(m); off += blockSize {
			if b := m[off : off+blockSize]; !slices.Contains(written[off:off+blockSize], false) && bytes.Contains(st.b, b) {
				t.Fatalf("round %d: the block at %d is in the store as written", round, off)
			}
		}
	}
}

// journalWith writes what SQLite writes into a journal: a header of a
// sector, some records, and, in the five writes of writeSuperJournal, rec.
func journalWith(t *testing.T, f *journalFile, pages int, rec []byte) int64 {
	t.Helper()
	rnd := rand.New(rand.NewSource(int64(pages)))
	hdr := make([]byte, 512)
	copy(hdr, journalMagic)
	rnd.Read(hdr[8:28])
	mustWrite(t, f, hdr, 0)
	off := int64(512)
	for range pages {
		page := make([]byte, 1024)
		rnd.Read(page)
		mustWrite(t, f, []byte{0, 0, 0, 7}, off)
		mustWrite(t, f, page, off+4)
		mustWrite(t, f, []byte{1, 2, 3, 4}, off+4+1024)
		off += 4 + 1024 + 4
	}

	if rec != nil {
		n := int64(len(rec)) - 20
		mustWrite(t, f, rec[:4], off)
		mustWrite(t, f, rec[4:4+n], off+4)
		mustWrite(t, f, rec[4+n:8+n], off+4+n)
		mustWrite(t, f, rec[8+n:12+n], off+8+n)
		mustWrite(t, f, rec[12+n:], off+12+n)
		off += int64(len(rec))
	}

	return off
}

// A journal that ends with the record of a super-journal the package knows
// shows it to a reader without the key; one that ends with anything else
// that looks like a record shows it to nobody.
func TestJournalSuperRecord(t *testing.T) {
	const name = "/data/main.db-mj0A1B2C9FF"
	const utf = "/data/\xd0\xb1\xd0\xb0\xd0\xb7\xd0\xb0.db-mj0A1B2C9FF"
	for _, tc := range []struct {
		what  string
		rec   []byte
		known string
		shown string
	}{
		{"known name", superRecord(name, true), name, name},
		{"unknown name", superRecord(name, true), "", ""},
		{"known name, signed chars", superRecord(utf, true), utf, utf},
		{"known name, unsigned chars", superRecord(utf, false), utf, utf},
		{"another name is known", superRecord(name, true), name + "0", ""},
		{"known name, bytes of a row in the checksum", func() []byte {
			r := superRecord(name, true)
			copy(r[len(r)-12:], "SECR")
			return r
		}(), name, ""},
		{"a longer name that holds a known one and a NUL", func() []byte {
			r := superRecord("private-token\x00"+name, true)
			return r
		}(), name, ""},
	} {
		t.Run(tc.what, func(t *testing.T) {
			st := &memStore{}
			f := testJournal(t, st, tc.known)
			end := journalWith(t, f, 3, tc.rec)
			if got := plainSuper(st.b); got != tc.shown {
				t.Fatalf("a reader without the key sees %q, want %q", got, tc.shown)
			}

			if got := ownerSuper(t, f); got != tc.shown {
				t.Fatalf("the owner sees %q, want %q", got, tc.shown)
			}

			if tc.shown != "" {
				mustSize(t, f, end)
				n := len(tc.shown) + 16
				if !bytes.Equal(st.b[len(st.b)-n:], tc.rec[4:]) {
					t.Fatal("the copy is not the record")
				}
			} else {
				// No byte of the record is in the store.
				for i := 4; i+8 <= len(tc.rec)-8; i++ {
					if bytes.Contains(st.b, tc.rec[i:i+8]) {
						t.Fatalf("bytes %d.. of the record are in the store in plain text", i)
					}
				}
			}

			// More of the journal after the record: the copy goes, and the
			// length is exact again.
			mustWrite(t, f, []byte{9, 9, 9, 9}, end)
			mustSize(t, f, end+4)
			if got := plainSuper(st.b); got != "" {
				t.Fatalf("after an append a reader without the key sees %q", got)
			}

			if want := testShift + roundUp(end+4) + blockSize; int64(len(st.b)) != want {
				t.Fatalf("after an append the store holds %d bytes, want %d", len(st.b), want)
			}
		})
	}
}

// A page whose image ends like a record, as SQLite writes it: the number,
// the page, and only then the checksum of the page. Between the last two
// writes the journal ends with the false record.
func TestJournalPageLikeRecord(t *testing.T) {
	const name = "/data/main.db-mj0A1B2C9FF"
	st := &memStore{}
	f := testJournal(t, st) // the name is not one SQLite created
	end := journalWith(t, f, 2, nil)
	page := bytes.Repeat([]byte("secret of a row "), 64)
	rec := superRecord(name, true)[4:]
	copy(page[len(page)-len(rec):], rec)
	mustWrite(t, f, []byte{0, 0, 0, 9}, end)
	mustWrite(t, f, page, end+4)
	if bytes.Contains(st.b, []byte("secret of a row ")) || bytes.Contains(st.b, []byte(name)) {
		t.Fatal("the page is in the store in plain text")
	}

	if got := ownerSuper(t, f); got != "" {
		t.Fatalf("the owner sees the super-journal %q in a page", got)
	}

	mustWrite(t, f, []byte{1, 2, 3, 4}, end+4+int64(len(page)))
	mustSize(t, f, end+8+int64(len(page)))
}

// A long journal of an earlier transaction is used again by a shorter one
// that ends with a record: SQLite writes the record below the old end and
// then truncates. The record and its copy appear together, at the cut.
func TestJournalTruncateToRecord(t *testing.T) {
	const name = "/data/main.db-mj0A1B2C9FF"
	for _, known := range []bool{true, false} {
		st := &memStore{}
		f := testJournal(t, st, map[bool]string{true: name}[known])
		old := journalWith(t, f, 20, nil)
		end := journalWith(t, f, 3, superRecord(name, true))
		mustSize(t, f, old)
		if got := plainSuper(st.b); got != "" {
			t.Fatalf("before the truncation a reader without the key sees %q", got)
		}

		if got := ownerSuper(t, f); got != "" {
			t.Fatalf("before the truncation the owner sees %q", got)
		}

		// The new end is written before the store is cut: at every moment
		// the store ends with a length block the owner can read.
		before := len(st.writes)
		probe := &orderStore{memStore: st}
		f.st = probe
		if err := f.truncate(end); err != nil {
			t.Fatal(err)
		}

		f.st = st
		if len(st.writes) != before+1 || probe.order != "write truncate " {
			t.Fatalf("the truncation reached the store as %q", probe.order)
		}

		want := map[bool]string{true: name}[known]
		if got := plainSuper(st.b); got != want {
			t.Fatalf("known=%v: a reader without the key sees %q", known, got)
		}

		if got := ownerSuper(t, f); got != want {
			t.Fatalf("known=%v: the owner sees %q", known, got)
		}

		if known {
			mustSize(t, f, end)
		}
	}
}

// orderStore records the order of the calls that change a store.
type orderStore struct {
	*memStore
	order string
}

func (o *orderStore) writeAt(p []byte, off int64) error {
	o.order += "write "
	return o.memStore.writeAt(p, off)
}

func (o *orderStore) truncate(size int64) error {
	o.order += "truncate "
	return o.memStore.truncate(size)
}

// What a crash can leave of the end of a stored journal, and a journal of
// another key.
func TestJournalDamagedEnd(t *testing.T) {
	st := &memStore{}
	f := testJournal(t, st)
	end := journalWith(t, f, 4, nil)
	whole := bytes.Clone(st.b)

	// The length block is cut: the journal is as long as the store.
	for _, cut := range []int{1, 7, 16, 17, 40} {
		st.b = bytes.Clone(whole[:len(whole)-cut])
		size := roundDown(int64(len(st.b)) - testShift)
		mustSize(t, f, size)
		// Every byte below the size is there to read: SQLite takes a short
		// read inside the journal for an error and gives the recovery up.
		if n, err := f.read(make([]byte, size), 0); err != nil || int64(n) != size {
			t.Fatalf("cut %d: %d of %d bytes read, %v", cut, n, size, err)
		}

		got := make([]byte, 512)
		if n, err := f.read(got, 0); err != nil || n != 512 || !bytes.HasPrefix(got, journalMagic) {
			t.Fatalf("cut %d: the header reads %d bytes, %v", cut, n, err)
		}

		// SQLite truncates it, or writes on: the length is exact again.
		if err := f.truncate(end - 100); err != nil {
			t.Fatalf("cut %d: truncate: %v", cut, err)
		}

		mustSize(t, f, end-100)
		mustWrite(t, f, make([]byte, 300), end-100)
		mustSize(t, f, end+200)
	}

	// Another key: no journal at all, and the next write begins a new one.
	st.b = bytes.Clone(whole)
	other := testJournal(t, st)
	c, err := derive(bytes.Repeat([]byte{1}, KeySize), labelKeyCheck)
	if err != nil {
		t.Fatal(err)
	}

	other.check = keyCheck(c)
	mustSize(t, other, 0)
	if n, err := other.read(make([]byte, 100), 0); err != nil || n != 0 {
		t.Fatalf("a journal of another key reads %d bytes, %v", n, err)
	}

	mustWrite(t, other, []byte("new journal"), 0)
	mustSize(t, other, 11)
	if want := int64(testShift + 16 + 16); int64(len(st.b)) != want {
		t.Fatalf("the new journal is stored as %d bytes, want %d", len(st.b), want)
	}

	// A truncation cannot make a journal longer.
	if err := other.truncate(12); err != errGrow {
		t.Fatalf("truncate beyond the end: %v", err)
	}

	if err := other.truncate(0); err != nil || len(st.b) != 0 {
		t.Fatalf("truncate to zero: %v, %d bytes left", err, len(st.b))
	}
}

// SQLite begins what it writes after a sync of the journal at the next
// sector, so that a write the power cuts short cannot damage what the sync
// made safe. The store is written no further back than that: not the last
// block of what was there, not the old length block.
func TestJournalWritesNoFurtherBack(t *testing.T) {
	st := &memStore{}
	f := testJournal(t, st)
	end := journalWith(t, f, 3, nil) // ends inside a sector
	synced := bytes.Clone(st.b[:testShift+roundDown(end)])
	next := (end + 4095) / 4096 * 4096
	before := len(st.writes)
	mustWrite(t, f, bytes.Repeat([]byte{7}, 512), next)
	if w := st.writes[before:]; len(w) != 1 || w[0].off != testShift+next {
		t.Fatalf("the write at %d reached the store as %v", next, w)
	}

	if !bytes.Equal(st.b[:len(synced)], synced) {
		t.Fatal("what was there before the write was changed")
	}

	mustSize(t, f, next+512)
	got := make([]byte, 512)
	if n, err := f.read(got, next); err != nil || n != 512 || !bytes.Equal(got, bytes.Repeat([]byte{7}, 512)) {
		t.Fatalf("read back: %d bytes, %v", n, err)
	}

	hdr := make([]byte, 8)
	if n, err := f.read(hdr, 0); err != nil || n != 8 || !bytes.Equal(hdr, journalMagic) {
		t.Fatalf("the header reads % x", hdr)
	}

	// A truncation writes after the block the journal ends in, and only
	// the length block.
	before = len(st.writes)
	if err := f.truncate(end - 3); err != nil {
		t.Fatal(err)
	}

	if w := st.writes[before:]; len(w) != 1 || w[0] != (span{testShift + roundUp(end-3), blockSize}) {
		t.Fatalf("the truncation reached the store as %v", w)
	}

	mustSize(t, f, end-3)
}

// A journal kept from one transaction to the next, on a storage whose
// sectors are larger than its prefix: the next transaction, which SQLite
// begins at offset 0, gets a prefix as long as a sector there.
func TestJournalPrefixAndSector(t *testing.T) {
	st := &memStore{}
	f := testJournal(t, st)
	journalWith(t, f, 3, nil)
	f.shift = 8192
	mustWrite(t, f, []byte{1, 2, 3, 4}, 600) // not a beginning: the journal stays
	if got := int64(binary.LittleEndian.Uint32(st.b[len(prefixMagic):])); got != testShift {
		t.Fatalf("a write inside the journal changed the prefix to %d", got)
	}

	hdr := make([]byte, 512)
	copy(hdr, journalMagic)
	mustWrite(t, f, hdr, 0)
	if got := int64(binary.LittleEndian.Uint32(st.b[len(prefixMagic):])); got != 8192 {
		t.Fatalf("the prefix of the new journal is %d bytes", got)
	}

	mustSize(t, f, 512)
	if want := int64(8192 + 512 + blockSize); int64(len(st.b)) != want {
		t.Fatalf("the new journal is stored as %d bytes, want %d", len(st.b), want)
	}
}

// A journal kept from an earlier transaction is written again from its
// beginning, and the new content ends exactly where the old one did, or the
// journal has lost its length block and is taken to end where the store
// does. A record at that end gets its copy all the same, and without a copy
// it is shown to nobody.
func TestJournalRecordAtTheOldEnd(t *testing.T) {
	const name = "/data/main.db-mj0A1B2C9FF"
	rec := superRecord(name, true)
	for _, lost := range []bool{false, true} {
		for _, known := range []bool{true, false} {
			what := fmt.Sprintf("length block lost=%v, known=%v", lost, known)
			st := &memStore{}
			f := testJournal(t, st, map[bool]string{true: name}[known])
			// An earlier journal of the length the new one will have: 16
			// more bytes than the record, so that it ends on a block.
			pad := (blockSize - (512+3*1032+int64(len(rec)))%blockSize) % blockSize
			end := journalWith(t, f, 3, nil)
			mustWrite(t, f, make([]byte, int64(len(rec))+pad), end)
			end += int64(len(rec)) + pad
			mustSize(t, f, end)
			if lost {
				st.b = st.b[:len(st.b)-blockSize]
				mustSize(t, f, end)
			}

			// The same transaction again, with the record at its end.
			if got := journalWith(t, f, 3, nil); got+pad+int64(len(rec)) != end {
				t.Fatalf("%s: the journal would end at %d, not %d", what, got+pad+int64(len(rec)), end)
			}

			at := end - int64(len(rec))
			n := int64(len(rec)) - 20
			mustWrite(t, f, rec[:4], at)
			mustWrite(t, f, rec[4:4+n], at+4)
			mustWrite(t, f, rec[4+n:8+n], at+4+n)
			mustWrite(t, f, rec[8+n:12+n], at+8+n)
			mustWrite(t, f, rec[12+n:], at+12+n)

			want := map[bool]string{true: name}[known]
			if got := plainSuper(st.b); got != want {
				t.Errorf("%s: a reader without the key sees %q", what, got)
			}

			if got := ownerSuper(t, f); got != want {
				t.Errorf("%s: the owner sees %q", what, got)
			}

			if ok, err := f.mayCommit(name); err != nil || ok != known {
				t.Errorf("%s: mayCommit: %v, %v", what, ok, err)
			}

			// SQLite truncates a journal it is told is longer than it
			// wrote; that changes nothing of what the two readers see.
			if err := f.truncate(end); err != nil {
				t.Fatalf("%s: truncate: %v", what, err)
			}

			if got := plainSuper(st.b); got != want {
				t.Errorf("%s: after the truncation a reader without the key sees %q", what, got)
			}

			if got := ownerSuper(t, f); got != want {
				t.Errorf("%s: after the truncation the owner sees %q", what, got)
			}
		}
	}

	// The record is in the journal and its copy and the length block are
	// not, which a loss of power can leave: nobody is shown it.
	st := &memStore{}
	f := testJournal(t, st, name)
	pad := (blockSize - (512+3*1032+int64(len(rec)))%blockSize) % blockSize
	end := journalWith(t, f, 3, nil)
	mustWrite(t, f, make([]byte, pad), end)
	mustWrite(t, f, rec, end+pad)
	if got := ownerSuper(t, f); got != name {
		t.Fatalf("the owner sees %q", got)
	}

	st.b = st.b[:testShift+end+pad+int64(len(rec))]
	if got := plainSuper(st.b); got != "" {
		t.Fatalf("a reader without the key sees %q in the blocks", got)
	}

	if got := ownerSuper(t, f); got != "" {
		t.Fatalf("without the copy the owner sees %q", got)
	}

	if ok, err := f.mayCommit(name); err != nil || ok {
		t.Fatalf("mayCommit: %v, %v", ok, err)
	}
}

// A journal without a length block whose content ends with something of the
// form of a record, and with a real record a block before that: hiding the
// one must not show the other, which has no copy either. And a transaction
// that begins in such a journal begins it anew.
func TestJournalRecordBeforeTheEnd(t *testing.T) {
	const name = "/data/main.db-mj0A1B2C9FF"
	rec := superRecord(name, true)
	st := &memStore{}
	f := testJournal(t, st) // no name is known: no copy is made
	pad := (blockSize - (512+3*1032+int64(len(rec)))%blockSize) % blockSize
	end := journalWith(t, f, 3, nil)
	mustWrite(t, f, make([]byte, pad), end)
	mustWrite(t, f, rec, end+pad)
	end += pad + int64(len(rec))
	mustWrite(t, f, rec[len(rec)-16:], end) // the last block of another record
	st.b = st.b[:testShift+end+16]
	if got := ownerSuper(t, f); got != "" {
		t.Fatalf("the owner sees %q", got)
	}

	if size, err := f.size(); err != nil || size >= end {
		t.Fatalf("size %d, %v: not before the record at %d", size, err, end)
	}

	if ok, err := f.mayCommit(name); err != nil || ok {
		t.Fatalf("mayCommit of a journal without a length block: %v, %v", ok, err)
	}

	hdr := make([]byte, 512)
	copy(hdr, journalMagic)
	mustWrite(t, f, hdr, 0)
	mustSize(t, f, 512)
	if want := int64(testShift + 512 + blockSize); int64(len(st.b)) != want {
		t.Fatalf("the journal begun anew is stored as %d bytes, want %d", len(st.b), want)
	}
}

// SQLite reads back a super-journal name of up to one byte more than the
// longest name of the VFS. Such a name gets its copy; a longer one, which
// SQLite itself would not find again after a crash, gets none, and the
// transaction does not commit.
func TestJournalLongName(t *testing.T) {
	for _, tc := range []struct {
		n    int
		copy bool
	}{{512, true}, {513, true}, {514, false}, {600, false}} {
		name := strings.Repeat("d", tc.n-len("/main.db-mj0A1B2C9FF")) + "/main.db-mj0A1B2C9FF"
		st := &memStore{}
		f := testJournal(t, st, name)
		journalWith(t, f, 2, superRecord(name, true))
		if got := plainSuper(st.b) == name; got != tc.copy {
			t.Errorf("a name of %d bytes: copy %v", tc.n, got)
		}

		if ok, err := f.mayCommit(name); err != nil || ok != tc.copy {
			t.Errorf("a name of %d bytes: mayCommit %v, %v", tc.n, ok, err)
		}
	}
}
