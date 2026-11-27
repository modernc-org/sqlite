// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// The eight bytes SQLite begins a journal header and ends a super-journal
// record with, aJournalMagic in pager.c.
var journalMagic = []byte{0xd9, 0xd5, 0x05, 0xf9, 0x20, 0xa1, 0x63, 0xd7}

// What a stored journal begins with, and what its length block begins with.
// The first byte is zero for a reason, see journalFile.
var (
	prefixMagic = []byte("\x00sqlite vfs/xts\x01")
	lengthMagic = []byte("xtsJLEN1")
)

const (
	minShift    = 4096  // the smallest prefix
	maxShift    = 65536 // the largest sector SQLite works with
	prefixCheck = 32    // where the key check block is in the prefix
	prefixRead  = prefixCheck + blockSize
)

var errGrow = errors.New("xts: a journal cannot be extended by a truncation")

// journalFile is the layout of a rollback journal. The stored file is
//
//	prefix | blocks | length block | copy of a super-journal record
//
// The prefix is as long as a sector, so that what SQLite aligns to sectors
// in the journal stays aligned in the store. Its first byte is zero: SQLite
// takes a journal for hot when its first byte is not, and a program that
// opens the database without the cipher would take a stored journal for a
// hot one, find no header in it, call that the end of the recovery and
// delete it. The prefix also holds a block that only the key gives, by which
// the key is known to be the right one when page 1 of the database is torn.
//
// The blocks are the content, the block at offset x of the content at offset
// x after the prefix, with the tweak of x. Unlike a database a journal is
// written at any offset and SQLite reads its length to the byte: it finds
// the name of a super-journal at the end of a journal by the size of the
// file. So the length is kept in a block of its own after the last block of
// the content, written in the same call as the data that changes the
// length. What lies between the end of the content and the end of its last
// block, and between an old end and a write beyond it, is not content:
// SQLite does not read there.
//
// The copy is there for readers that do not have the key. When SQLite
// recovers one database of a transaction over several, it reads the end of
// the journals of the others, through the VFS of the one it recovers, to
// see whether they still name the super-journal, and deletes the
// super-journal if none does; a journal whose super-journal is gone is not
// rolled back. So a journal that ends with the record of a super-journal
// repeats that record in plain text at the end of the stored file, where
// SQLite reads it. The bytes of the copy are made from a name the package
// was given when SQLite created the super-journal, never taken from the
// journal: content of the database that looks like a record does not reach
// the store in plain text. And the owner of the journal is shown a record
// only while its copy is there: without one it is given the stored size, at
// which SQLite finds no record and rolls the journal back.
type journalFile struct {
	st      store
	c, cl   *xts
	check   []byte // the key check block
	shift   int64  // the prefix of a journal this file creates
	maxName int64  // the longest name of the base VFS

	// known reports whether name is a super-journal SQLite created through
	// the package, on the storage this journal is on.
	known func(name []byte) bool

	buf []byte
}

// jstate is what the store holds.
type jstate struct {
	r     int64 // the size of the store
	s     int64 // the prefix; 0 if there is no journal the key can read
	l     int64 // the length of the content; r-s without a length block
	exact bool  // a length block says l
	copyN int64 // the length of the copy
}

func (f *journalFile) scratch(n int64) []byte {
	if int64(cap(f.buf)) < n {
		f.buf = make([]byte, n)
	}

	b := f.buf[:n]
	clear(b)
	return b
}

// full reads all of p from the store at off, or reports that it could not.
func (f *journalFile) full(p []byte, off int64) (bool, error) {
	n, err := f.st.readAt(p, off)
	return n == len(p), err
}

// prefix returns the length of the prefix of the stored journal, or 0 if
// the store does not begin with a prefix of this key.
func (f *journalFile) prefix(r int64) (int64, error) {
	if r < minShift {
		return 0, nil
	}

	var p [prefixRead]byte
	ok, err := f.full(p[:], 0)
	if err != nil || !ok {
		return 0, err
	}

	s := int64(binary.LittleEndian.Uint32(p[len(prefixMagic):]))
	if !bytes.Equal(p[:len(prefixMagic)], prefixMagic) ||
		s < minShift || s > maxShift || s&(s-1) != 0 || r < s ||
		!bytes.Equal(p[prefixCheck:], f.check) {
		return 0, nil
	}

	return s, nil
}

// lengthAt reads the length block at offset x of the content.
func (f *journalFile) lengthAt(s, x int64) (int64, bool, error) {
	var b [blockSize]byte
	ok, err := f.full(b[:], s+x)
	if err != nil || !ok {
		return 0, false, err
	}

	f.cl.decryptAt(b[:], b[:], x)
	l := int64(binary.LittleEndian.Uint64(b[len(lengthMagic):]))
	if !bytes.Equal(b[:len(lengthMagic)], lengthMagic) || l <= 0 || roundUp(l) != x {
		return 0, false, nil
	}

	return l, true, nil
}

// plain reads n bytes of the content at x, all of which the store must have.
func (f *journalFile) plain(s, x, n int64) ([]byte, error) {
	lo, hi := roundDown(x), roundUp(x+n)
	b := make([]byte, hi-lo)
	ok, err := f.full(b, s+lo)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("xts: the journal is shorter than its length block says")
	}

	f.c.decryptAt(b, b, lo)
	return b[x-lo : x-lo+n], nil
}

// recordForm reports whether tail, the end of a content, has the form of a
// super-journal record, and returns the length of the name in it.
func (f *journalFile) recordForm(tail []byte) (int64, bool) {
	if len(tail) < 2*len(journalMagic) || !bytes.HasSuffix(tail, journalMagic) {
		return 0, false
	}

	// SQLite reads back a name of up to one byte more than the longest
	// name of the VFS (pager_playback, readSuperJournal).
	n := int64(binary.BigEndian.Uint32(tail[len(tail)-16:]))
	return n, n >= 1 && n <= f.maxName+1
}

// state reads what the store holds.
func (f *journalFile) state() (jstate, error) {
	r, err := f.st.size()
	if err != nil {
		return jstate{}, err
	}

	s, err := f.prefix(r)
	if err != nil || s == 0 {
		return jstate{r: r}, err
	}

	// Without a length block the journal is as long as the whole blocks the
	// store holds after the prefix: what is read is whole blocks, and every
	// byte below the size SQLite is told must be there to read.
	st := jstate{r: r, s: s, l: roundDown(r - s)}
	if r == s {
		st.exact = true
		return st, nil
	}

	// A copy at the end says where the length block is.
	var tail [16]byte
	if r-s >= 3*blockSize {
		if _, err := f.full(tail[:], r-16); err != nil {
			return jstate{}, err
		}

		if n, ok := f.recordForm(tail[:]); ok {
			c := n + 16
			if x := r - c - blockSize - s; x >= blockSize && x%blockSize == 0 {
				l, ok, err := f.lengthAt(s, x)
				if err != nil {
					return jstate{}, err
				}

				if ok && l >= c {
					stored := make([]byte, c)
					if _, err := f.full(stored, r-c); err != nil {
						return jstate{}, err
					}

					end, err := f.plain(s, l-c, c)
					if err != nil {
						return jstate{}, err
					}

					if bytes.Equal(stored, end) {
						st.l, st.exact, st.copyN = l, true, c
						return st, nil
					}
				}
			}
		}
	}

	// Without a copy the length block is the last block.
	if x := r - blockSize - s; x >= blockSize && x%blockSize == 0 {
		l, ok, err := f.lengthAt(s, x)
		if err != nil {
			return jstate{}, err
		}

		if ok {
			st.l, st.exact = l, true
		}
	}

	return st, nil
}

func (f *journalFile) size() (int64, error) {
	st, err := f.state()
	if err != nil || st.s == 0 {
		return 0, err
	}

	return f.shown(st)
}

// shown returns the length of the journal as its owner is told it. A record
// at the end without its copy is not shown: the journal is then said to end
// where SQLite finds no record. With a length block that is behind the
// block, in what is no content. Without one it is a block or more before
// the end, which is still behind everything a sync of the journal made safe
// before the record was written.
func (f *journalFile) shown(st jstate) (int64, error) {
	if st.copyN != 0 {
		return st.l, nil
	}

	for end := st.l; ; end -= blockSize {
		if end < 16 {
			return end, nil
		}

		tail, err := f.plain(st.s, end-16, 16)
		if err != nil {
			return 0, err
		}

		if n, ok := f.recordForm(tail); !ok || end < n+16 {
			return end, nil
		}

		if st.exact {
			return roundDown(st.r - st.s), nil
		}
	}
}

// mayCommit reports whether a transaction over several databases, whose
// super-journal is called name, may commit with this journal as it is: the
// journal of the transaction, open and written by the connection that asks.
// It may not if the journal has no length block, which the writes of one
// transaction never leave, or if it ends with the record of the
// super-journal and the copy of that record is not there.
func (f *journalFile) mayCommit(name string) (bool, error) {
	st, err := f.state()
	if err != nil || st.s == 0 {
		return err == nil, err
	}

	n := int64(len(name))
	if !st.exact || st.l < n+16 {
		return st.exact, nil
	}

	tail, err := f.plain(st.s, st.l-n-16, n+16)
	if err != nil {
		return false, err
	}

	if string(tail[:n]) != name || int64(binary.BigEndian.Uint32(tail[n:])) != n || !bytes.HasSuffix(tail, journalMagic) {
		return true, nil
	}

	return st.copyN == n+16, nil
}

func (f *journalFile) read(p []byte, off int64) (int, error) {
	r, err := f.st.size()
	if err != nil {
		return 0, err
	}

	s, err := f.prefix(r)
	if err != nil || s == 0 {
		clear(p)
		return 0, err
	}

	// Everything after the prefix can be read as content: SQLite reads no
	// further than the size it was told, and when that is the stored size
	// it must find there what is no journal header and no record.
	lo, hi := roundDown(off), roundUp(off+int64(len(p)))
	buf := f.scratch(hi - lo)
	n, err := f.st.readAt(buf, s+lo)
	if err != nil {
		return 0, err
	}

	n = int(roundDown(int64(n)))
	f.c.decryptAt(buf[:n], buf[:n], lo)
	clear(buf[n:])
	copy(p, buf[off-lo:])
	return int(min(max(lo+int64(n)-off, 0), int64(len(p)))), nil
}

// copyFor returns the record to repeat in plain text after a content that
// ends with tail, of which the caller has at least the last 16 bytes; more
// returns the last n bytes of the content. The record is nil if the content
// does not end with the record of a super-journal the package knows.
func (f *journalFile) copyFor(tail []byte, more func(n int64) ([]byte, error)) ([]byte, error) {
	n, ok := f.recordForm(tail)
	if !ok {
		return nil, nil
	}

	rec, err := more(n + 16)
	if err != nil || rec == nil {
		return nil, err
	}

	name := rec[:n]
	if bytes.IndexByte(name, 0) >= 0 || !f.known(name) {
		return nil, nil
	}

	// SQLite's checksum of the name is the sum of its bytes as C chars,
	// which are signed on some targets and unsigned on others.
	var signed, unsigned uint32
	for _, b := range name {
		signed += uint32(int32(int8(b)))
		unsigned += uint32(b)
	}

	for _, sum := range []uint32{signed, unsigned} {
		want := make([]byte, 0, n+16)
		want = append(want, name...)
		want = binary.BigEndian.AppendUint32(want, uint32(n))
		want = binary.BigEndian.AppendUint32(want, sum)
		want = append(want, journalMagic...)
		if bytes.Equal(rec, want) {
			return want, nil
		}
	}

	return nil, nil
}

// keyCheck returns the block of the prefix that only the key gives.
func keyCheck(c *xts) []byte {
	b := []byte("xts key check 1\x00")
	c.encrypt(b, b, 0, 0)
	return b
}

func (f *journalFile) lengthBlock(l int64) []byte {
	b := make([]byte, blockSize)
	copy(b, lengthMagic)
	binary.LittleEndian.PutUint64(b[len(lengthMagic):], uint64(l))
	f.cl.encryptAt(b, b, roundUp(l))
	return b
}

func (f *journalFile) write(p []byte, off int64) error {
	if len(p) == 0 {
		return nil
	}

	st, err := f.state()
	if err != nil {
		return err
	}

	s, l, fresh := st.s, st.l, st.s == 0
	if !fresh && off == 0 && (s%f.shift != 0 || !st.exact) {
		// SQLite begins a journal anew at offset 0, and what an earlier
		// transaction left behind it is then dead. Two kinds of journal are
		// not written into but begun anew with it: one that was begun on a
		// storage with smaller sectors, whose prefix would put what SQLite
		// aligns to the sectors of this one between them, and one without a
		// length block, the rest of an interrupted write, in which nothing
		// says where a record at its end belongs.
		fresh = true
	}

	if fresh {
		s, l = f.shift, 0
	}

	end := off + int64(len(p))
	lo, hi := roundDown(off), roundUp(end)
	if end < l {
		// Inside the journal: its length does not change.
		buf := f.scratch(hi - lo)
		for _, at := range []int64{lo, hi - blockSize} {
			if ok, err := f.full(buf[at-lo:at-lo+blockSize], s+at); err != nil {
				return err
			} else if ok {
				f.c.decryptAt(buf[at-lo:at-lo+blockSize], buf[at-lo:at-lo+blockSize], at)
			} else {
				clear(buf[at-lo : at-lo+blockSize])
			}
		}

		copy(buf[off-lo:], p)
		f.c.encryptAt(buf, buf, lo)
		return f.st.writeAt(buf, s+lo)
	}

	// The write ends where the journal ends, or beyond: the end is written
	// anew, also when it does not move, because what the journal ends with
	// decides whether a copy follows. What goes to the store, in one call, is
	// the blocks the write lies in, the padding, the length block and the
	// copy, and nothing before the block the write begins in: SQLite begins
	// what it writes after a sync at the next sector, so that a write the
	// power cuts short cannot damage what the sync made safe, and the store
	// must be written no further back than SQLite writes. The bytes between
	// the old end and the write, which SQLite never reads, are whatever the
	// store holds there; the old length block is among them.
	head := int64(0)
	if fresh {
		// A new journal begins with its prefix, and the store has no gap.
		head, lo = s, 0
	}

	buf := f.scratch(head + hi - lo + blockSize)
	body := buf[head : head+hi-lo]
	if !fresh && off != lo {
		if ok, err := f.full(body[:blockSize], s+lo); err != nil {
			return err
		} else if ok {
			f.c.decryptAt(body[:blockSize], body[:blockSize], lo)
		} else {
			clear(body[:blockSize])
		}
	}

	copy(body[off-lo:], p)
	more := func(n int64) ([]byte, error) {
		if n > end {
			return nil, nil
		}

		if end-n >= lo {
			return body[end-n-lo : end-lo], nil
		}

		before, err := f.plain(s, end-n, lo-(end-n))
		if err != nil {
			return nil, err
		}

		return append(before, body[:end-lo]...), nil
	}

	var rec []byte
	if k := min(end-lo, int64(len(journalMagic))); bytes.Equal(body[end-lo-k:end-lo], journalMagic[int64(len(journalMagic))-k:]) {
		tail, err := more(16)
		if err != nil {
			return err
		}

		if tail != nil {
			if rec, err = f.copyFor(tail, more); err != nil {
				return err
			}
		}
	}

	f.c.encryptAt(body, body, lo)
	copy(buf[head+hi-lo:], f.lengthBlock(end))
	out := append(buf, rec...)
	at := s + lo
	if fresh {
		copy(out, prefixMagic)
		binary.LittleEndian.PutUint32(out[len(prefixMagic):], uint32(s))
		copy(out[prefixCheck:], f.check)
		at = 0
		if st.r != 0 {
			// What is there is no journal of this key.
			if err := f.st.truncate(0); err != nil {
				return err
			}

			st.r = 0
		}
	}

	if err := f.st.writeAt(out, at); err != nil {
		return err
	}

	if stored := at + int64(len(out)); st.r > stored {
		return f.st.truncate(stored)
	}

	return nil
}

func (f *journalFile) truncate(size int64) error {
	st, err := f.state()
	if err != nil {
		return err
	}

	switch {
	case size == 0:
		if st.r == 0 {
			return nil
		}

		return f.st.truncate(0)
	case st.s == 0 || size > st.l:
		return errGrow
	}

	// The new end goes to the store first, over what is being cut off: the
	// length block and the copy, after the block the journal now ends in.
	// Until the store is cut the old length block is its last block and the
	// journal is as long as it was.
	hi := roundUp(size)
	out := f.lengthBlock(size)
	more := func(n int64) ([]byte, error) {
		if n > size || st.r < st.s+hi {
			return nil, nil
		}

		return f.plain(st.s, size-n, n)
	}

	if tail, err := more(16); err != nil {
		return err
	} else if tail != nil {
		rec, err := f.copyFor(tail, more)
		if err != nil {
			return err
		}

		out = append(out, rec...)
	}

	if st.r < st.s+hi {
		// A journal without a length block that ends inside a block: the
		// block is completed with zeros.
		if err := f.st.truncate(st.s + hi); err != nil {
			return err
		}
	}

	if err := f.st.writeAt(out, st.s+hi); err != nil {
		return err
	}

	return f.st.truncate(st.s + hi + int64(len(out)))
}
