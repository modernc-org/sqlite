// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

// memStore is a store in memory. It records every call that changes it.
type memStore struct {
	b      []byte
	writes []span // the ranges written, in order
	fail   error  // returned by every call when set
}

type span struct{ off, n int64 }

func (m *memStore) readAt(p []byte, off int64) (int, error) {
	if m.fail != nil {
		return 0, m.fail
	}

	if off >= int64(len(m.b)) {
		return 0, nil
	}

	return copy(p, m.b[off:]), nil
}

func (m *memStore) writeAt(p []byte, off int64) error {
	if m.fail != nil {
		return m.fail
	}

	if need := off + int64(len(p)); need > int64(len(m.b)) {
		m.b = append(m.b, make([]byte, need-int64(len(m.b)))...)
	}

	copy(m.b[off:], p)
	m.writes = append(m.writes, span{off, int64(len(p))})
	return nil
}

func (m *memStore) truncate(size int64) error {
	if m.fail != nil {
		return m.fail
	}

	if size <= int64(len(m.b)) {
		m.b = m.b[:size]
		return nil
	}

	m.b = append(m.b, make([]byte, size-int64(len(m.b)))...)
	return nil
}

func (m *memStore) size() (int64, error) { return int64(len(m.b)), m.fail }

// plainModel is the file SQLite believes it has: what it wrote, where, and
// which bytes it never wrote.
type plainModel struct {
	b       []byte
	written []bool
}

func (m *plainModel) grow(size int64) {
	for int64(len(m.b)) < size {
		m.b = append(m.b, 0)
		m.written = append(m.written, false)
	}
}

func (m *plainModel) write(p []byte, off int64) {
	m.grow(off + int64(len(p)))
	copy(m.b[off:], p)
	for i := range p {
		m.written[off+int64(i)] = true
	}
}

func testCipher(t testing.TB, label string) *xts {
	t.Helper()
	c, err := derive(bytes.Repeat([]byte{0x5a}, KeySize), label)
	if err != nil {
		t.Fatal(err)
	}

	return c
}

// checkRead reads [off, off+n) through f and compares it with the model:
// every byte SQLite wrote, the count, and zeros past the end.
func checkRead(t *testing.T, what string, f file, m *plainModel, size, off, n int64) {
	t.Helper()
	p := bytes.Repeat([]byte{0xee}, int(n))
	got, err := f.read(p, off)
	if err != nil {
		t.Fatalf("%s: read: %v", what, err)
	}

	if want := min(max(size-off, 0), n); int64(got) != want {
		t.Fatalf("%s: read [%d, %d) of a file of %d: %d bytes, want %d", what, off, off+n, size, got, want)
	}

	for i := int64(0); i < n; i++ {
		at := off + i
		switch {
		case at >= size:
			if p[i] != 0 {
				t.Fatalf("%s: byte %d, past the end at %d, reads %#x", what, at, size, p[i])
			}
		case at < int64(len(m.b)) && m.written[at]:
			if p[i] != m.b[at] {
				t.Fatalf("%s: byte %d reads %#x, written %#x", what, at, p[i], m.b[at])
			}
		}
	}
}

// The database layout against the plain model: whole pages as SQLite writes
// a database, and writes at any place as it writes temporary files.
func TestDBFileModel(t *testing.T) {
	for _, aligned := range []bool{true, false} {
		t.Run(fmt.Sprintf("aligned=%v", aligned), func(t *testing.T) {
			rnd := rand.New(rand.NewSource(3))
			for round := range 200 {
				st := &memStore{}
				f := &dbFile{st: st, c: testCipher(t, labelDatabase)}
				m := &plainModel{}
				var size int64
				const limit = 6 * unitSize
				for op := range 120 {
					what := fmt.Sprintf("round %d op %d", round, op)
					off, n := rnd.Int63n(limit), 1+rnd.Int63n(2*unitSize)
					if aligned {
						page := int64(512 << rnd.Intn(4))
						off, n = off/page*page, page
					}

					switch k := rnd.Intn(10); {
					case k < 6:
						p := make([]byte, n)
						rnd.Read(p)
						before := len(st.writes)
						if err := f.write(p, off); err != nil {
							t.Fatalf("%s: write: %v", what, err)
						}

						m.write(p, off)
						size = max(size, roundUp(off+n))
						if aligned {
							// Exactly the range SQLite wrote, in one write.
							if w := st.writes[before:]; len(w) != 1 || w[0] != (span{off, n}) {
								t.Fatalf("%s: write [%d, %d) reached the store as %v", what, off, off+n, w)
							}
						}

					case k < 7:
						to := rnd.Int63n(limit)
						if aligned {
							to = to / 512 * 512
						}

						if err := f.truncate(to); err != nil {
							t.Fatalf("%s: truncate: %v", what, err)
						}

						// What is cut off is gone; growing gives bytes
						// SQLite never wrote.
						m.grow(to)
						for i := to; i < int64(len(m.b)); i++ {
							m.b[i], m.written[i] = 0, i < roundUp(to) && to%blockSize != 0
						}

						m.b, m.written = m.b[:max(to, 0)], m.written[:max(to, 0)]
						m.grow(roundUp(to))
						for i := to; i < roundUp(to); i++ {
							m.written[i] = true // the zeros after the new end
						}

						size = roundUp(to)

					default:
						checkRead(t, what, f, m, size, off, n)
					}

					if got, err := f.size(); err != nil || got != size {
						t.Fatalf("%s: size %d %v, want %d", what, got, err, size)
					}

					if int64(len(st.b)) != size {
						t.Fatalf("%s: the store holds %d bytes, the file %d", what, len(st.b), size)
					}
				}

				checkRead(t, "at the end", f, m, size, 0, limit+unitSize)
			}
		})
	}
}

// What the store holds is not what SQLite wrote: no block of the content is
// found in it, and a block written twice at two places is stored twice
// differently.
func TestDBFileStoresCiphertext(t *testing.T) {
	st := &memStore{}
	f := &dbFile{st: st, c: testCipher(t, labelDatabase)}
	page := bytes.Repeat([]byte("sixteen byte txt"), unitSize/blockSize)
	for i := range int64(4) {
		if err := f.write(page, i*unitSize); err != nil {
			t.Fatal(err)
		}
	}

	seen := map[string]bool{}
	for off := 0; off < len(st.b); off += blockSize {
		b := string(st.b[off : off+blockSize])
		if b == "sixteen byte txt" {
			t.Fatalf("the block at %d is stored as written", off)
		}

		if seen[b] {
			t.Fatalf("the block at %d is stored like an earlier one", off)
		}

		seen[b] = true
	}

	// The same content at the same place is stored the same: the cipher is
	// deterministic, which the documentation says.
	before := bytes.Clone(st.b)
	if err := f.write(page, unitSize); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(before, st.b) {
		t.Fatal("a page written again at its place is stored differently")
	}
}

// An error of the store is the error of the call, and nothing is written
// after a failed read.
func TestDBFileErrors(t *testing.T) {
	st := &memStore{}
	f := &dbFile{st: st, c: testCipher(t, labelDatabase)}
	if err := f.write(make([]byte, 100), 0); err != nil {
		t.Fatal(err)
	}

	boom := errors.New("boom")
	st.fail = boom
	n := len(st.writes)
	if _, err := f.read(make([]byte, 10), 3); err != boom {
		t.Fatalf("read: %v", err)
	}

	if err := f.write(make([]byte, 10), 3); err != boom {
		t.Fatalf("write: %v", err)
	}

	if err := f.truncate(5); err != boom {
		t.Fatalf("truncate: %v", err)
	}

	if _, err := f.size(); err != boom {
		t.Fatalf("size: %v", err)
	}

	if len(st.writes) != n {
		t.Fatal("the store was written after a failed read")
	}
}

// A store that ends inside a block, which only something other than the
// layout can make of it: the cut block is not there, for a read and for the
// write that completes it.
func TestDBFileCutBlock(t *testing.T) {
	st := &memStore{}
	f := &dbFile{st: st, c: testCipher(t, labelDatabase)}
	p := bytes.Repeat([]byte{0xab}, 5*blockSize)
	if err := f.write(p, 0); err != nil {
		t.Fatal(err)
	}

	st.b = st.b[:4*blockSize+7]
	got := bytes.Repeat([]byte{0xee}, 6*blockSize)
	n, err := f.read(got, 0)
	if err != nil || n != 4*blockSize {
		t.Fatalf("read: %d bytes, %v; want %d", n, err, 4*blockSize)
	}

	if !bytes.Equal(got[:n], p[:n]) || !bytes.Equal(got[n:], make([]byte, len(got)-n)) {
		t.Fatalf("read: % x", got)
	}

	n, err = f.read(got[:10], 4*blockSize+2)
	if err != nil || n != 0 || !bytes.Equal(got[:10], make([]byte, 10)) {
		t.Fatalf("read inside the cut block: %d bytes, %v, % x", n, err, got[:10])
	}

	if err := f.write([]byte{1, 2, 3}, 4*blockSize+9); err != nil {
		t.Fatal(err)
	}

	want := append(bytes.Clone(p[:4*blockSize]), make([]byte, blockSize)...)
	copy(want[4*blockSize+9:], []byte{1, 2, 3})
	n, err = f.read(got[:5*blockSize], 0)
	if err != nil || n != 5*blockSize || !bytes.Equal(got[:n], want) {
		t.Fatalf("after the write: %d bytes, %v\n got % x\nwant % x", n, err, got[:n], want)
	}
}
