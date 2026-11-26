// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hkdf"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// blockSize is the block of AES. Nothing shorter is ever encrypted.
const blockSize = aes.BlockSize

// unitSize is the data unit of XTS in a file: the tweak of a block is made
// of the number of the 4096-byte unit it lies in and its index in that unit.
const unitSize = 4096

// xts is AES in XTS mode, IEEE Std 1619 and NIST SP 800-38E, without
// ciphertext stealing: the lengths are multiples of the block. It differs
// from golang.org/x/crypto/xts in one thing the files need, that a call can
// begin at any block of a data unit, not only at its first.
type xts struct {
	data, tweak cipher.Block
}

// newXTS takes the two AES keys of XTS, one after the other.
func newXTS(key []byte) (*xts, error) {
	half := len(key) / 2
	data, err := aes.NewCipher(key[:half])
	if err != nil {
		return nil, err
	}

	tweak, err := aes.NewCipher(key[half:])
	if err != nil {
		return nil, err
	}

	return &xts{data, tweak}, nil
}

// start returns the tweak of block number block of data unit unit, as the two
// halves of a little-endian 128-bit number: the unit number encrypted under
// the tweak key, times alpha to the power of block.
func (x *xts) start(unit uint64, block int) (lo, hi uint64) {
	var t [blockSize]byte
	binary.LittleEndian.PutUint64(t[:], unit)
	x.tweak.Encrypt(t[:], t[:])
	lo, hi = binary.LittleEndian.Uint64(t[:]), binary.LittleEndian.Uint64(t[8:])
	for ; block > 0; block-- {
		lo, hi = mulAlpha(lo, hi)
	}
	return lo, hi
}

// mulAlpha multiplies a tweak by alpha in GF(2^128) modulo
// x^128 + x^7 + x^2 + x + 1, in time that does not depend on the tweak.
func mulAlpha(lo, hi uint64) (uint64, uint64) {
	carry := hi >> 63
	return lo<<1 ^ (-carry & 0x87), hi<<1 | lo>>63
}

// encrypt encrypts src, whole blocks of data unit unit that begin at its
// block number block, into dst. The two may be the same slice.
func (x *xts) encrypt(dst, src []byte, unit uint64, block int) {
	x.crypt(dst, src, unit, block, x.data.Encrypt)
}

// decrypt is the inverse of encrypt.
func (x *xts) decrypt(dst, src []byte, unit uint64, block int) {
	x.crypt(dst, src, unit, block, x.data.Decrypt)
}

func (x *xts) crypt(dst, src []byte, unit uint64, block int, f func(dst, src []byte)) {
	if len(src)%blockSize != 0 || len(dst) < len(src) {
		panic("xts: not whole blocks")
	}

	lo, hi := x.start(unit, block)
	var b [blockSize]byte
	for i := 0; i < len(src); i += blockSize {
		binary.LittleEndian.PutUint64(b[:], binary.LittleEndian.Uint64(src[i:])^lo)
		binary.LittleEndian.PutUint64(b[8:], binary.LittleEndian.Uint64(src[i+8:])^hi)
		f(b[:], b[:])
		binary.LittleEndian.PutUint64(dst[i:], binary.LittleEndian.Uint64(b[:])^lo)
		binary.LittleEndian.PutUint64(dst[i+8:], binary.LittleEndian.Uint64(b[8:])^hi)
		lo, hi = mulAlpha(lo, hi)
	}
}

// encryptAt encrypts src, whole blocks that lie at offset off of a file, into
// dst: every 4096 bytes of the file are one data unit. The two slices may be
// the same.
func (x *xts) encryptAt(dst, src []byte, off int64) { x.cryptAt(dst, src, off, x.encrypt) }

// decryptAt is the inverse of encryptAt.
func (x *xts) decryptAt(dst, src []byte, off int64) { x.cryptAt(dst, src, off, x.decrypt) }

func (x *xts) cryptAt(dst, src []byte, off int64, f func(dst, src []byte, unit uint64, block int)) {
	if off%blockSize != 0 {
		panic("xts: not at a block")
	}

	for len(src) > 0 {
		in := int(off % unitSize)
		n := min(len(src), unitSize-in)
		f(dst[:n], src[:n], uint64(off/unitSize), in/blockSize)
		dst, src, off = dst[n:], src[n:], off+int64(n)
	}
}

// KeySize is the length of the key [Register] takes.
const KeySize = 32

// The purposes a key is derived for. The labels are part of format 1: a
// change of one makes every file written before it unreadable.
const (
	labelDatabase      = "modernc.org/sqlite/vfs/xts format 1: database"
	labelJournal       = "modernc.org/sqlite/vfs/xts format 1: journal"
	labelJournalLength = "modernc.org/sqlite/vfs/xts format 1: journal length"
	labelKeyCheck      = "modernc.org/sqlite/vfs/xts format 1: key check"
)

// derive makes the cipher of one purpose from the key of the registration:
// 64 bytes of HKDF-SHA256 output, the two AES-256 keys of XTS.
func derive(key []byte, label string) (*xts, error) {
	if len(key) != KeySize {
		return nil, fmt.Errorf("xts: the key must be %d bytes, not %d", KeySize, len(key))
	}

	k, err := hkdf.Key(sha256.New, key, nil, label, 2*KeySize)
	if err != nil {
		return nil, err
	}

	defer clear(k)

	return newXTS(k)
}
