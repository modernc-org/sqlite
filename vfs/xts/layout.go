// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

// store is the file under the cipher: the file of the base VFS, or a slice
// of bytes in the tests.
type store interface {
	// readAt reads len(p) bytes at off and returns how many of them the
	// file has: fewer than len(p) only where the file ends.
	readAt(p []byte, off int64) (int, error)

	// writeAt writes p at off. The bytes between the end of the file and
	// off, if any, are whatever the file gives for bytes never written.
	writeAt(p []byte, off int64) error

	truncate(size int64) error
	size() (int64, error)
}

// file is what SQLite sees of a store: the plain content.
type file interface {
	// read fills p from off and returns how many of its bytes lie in the
	// file; the rest of p is zeros.
	read(p []byte, off int64) (int, error)

	write(p []byte, off int64) error
	truncate(size int64) error
	size() (int64, error)
}

func roundDown(n int64) int64 { return n &^ (blockSize - 1) }
func roundUp(n int64) int64   { return (n + blockSize - 1) &^ (blockSize - 1) }

// dbFile is the layout of a database file and of every temporary file: the
// block at offset x of the content is the block at offset x of the store,
// encrypted with the tweak of its place. Nothing is added and nothing is
// moved, so the locks SQLite takes on byte ranges of the file, and the page
// it never touches because of them, are where they are without the cipher.
//
// SQLite writes a database in whole pages, and a page is whole blocks: the
// store then gets exactly the range SQLite wrote and nothing is read. Only
// temporary files are written at other places; there the blocks at the two
// ends of a write are read, changed and written back, and the file ends on a
// block, with zeros after the last byte SQLite wrote.
//
// Bytes that lie inside the file and were never written read as the
// decryption of what the store holds there, not as zeros. SQLite does not
// look at the content of what it has not written.
type dbFile struct {
	st  store
	c   *xts
	buf []byte
}

func (f *dbFile) scratch(n int64) []byte {
	if int64(cap(f.buf)) < n {
		f.buf = make([]byte, n)
	}

	return f.buf[:n]
}

// block reads the block at off into dst, decrypted. A block that is not, or
// not wholly, in the store reads as zeros.
func (f *dbFile) block(dst []byte, off int64) error {
	n, err := f.st.readAt(dst[:blockSize], off)
	if err != nil {
		return err
	}

	if n < blockSize {
		clear(dst[:blockSize])
		return nil
	}

	f.c.decryptAt(dst[:blockSize], dst[:blockSize], off)
	return nil
}

func (f *dbFile) read(p []byte, off int64) (int, error) {
	lo, hi := roundDown(off), roundUp(off+int64(len(p)))
	buf := p
	if lo != off || hi != off+int64(len(p)) {
		buf = f.scratch(hi - lo)
	}

	n, err := f.st.readAt(buf, lo)
	if err != nil {
		return 0, err
	}

	// A block the end of the store cuts through is not there.
	n = int(roundDown(int64(n)))
	f.c.decryptAt(buf[:n], buf[:n], lo)
	clear(buf[n:])
	got := min(max(lo+int64(n)-off, 0), int64(len(p)))
	if lo != off || hi != off+int64(len(p)) {
		copy(p, buf[off-lo:])
	}

	return int(got), nil
}

func (f *dbFile) write(p []byte, off int64) error {
	if len(p) == 0 {
		return nil
	}

	lo, hi := roundDown(off), roundUp(off+int64(len(p)))
	buf := f.scratch(hi - lo)
	if lo != off {
		if err := f.block(buf, lo); err != nil {
			return err
		}
	}

	if end := off + int64(len(p)); hi != end && (hi-blockSize != lo || lo == off) {
		if err := f.block(buf[hi-blockSize-lo:], hi-blockSize); err != nil {
			return err
		}
	}

	copy(buf[off-lo:], p)
	f.c.encryptAt(buf, buf, lo)
	return f.st.writeAt(buf, lo)
}

func (f *dbFile) truncate(size int64) error {
	hi := roundUp(size)
	if hi == size {
		return f.st.truncate(size)
	}

	// The file is to end inside a block: the block stays, with zeros from
	// the new end on.
	buf := f.scratch(blockSize)
	if err := f.block(buf, hi-blockSize); err != nil {
		return err
	}

	clear(buf[size-(hi-blockSize):])
	f.c.encryptAt(buf, buf, hi-blockSize)
	if err := f.st.truncate(hi - blockSize); err != nil {
		return err
	}

	return f.st.writeAt(buf, hi-blockSize)
}

func (f *dbFile) size() (int64, error) { return f.st.size() }
