// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"testing"
)

func unhex(t *testing.T, s string) []byte {
	t.Helper()
	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}

	return b
}

// The vectors of IEEE Std 1619, Annex B, that golang.org/x/crypto/xts is
// tested with: numbers 1, 2, 3, 4 and 5 for XTS-AES-128 and number 10 for
// XTS-AES-256, in ieee1619_test.go.
func TestIEEE1619(t *testing.T) {
	for i, v := range ieee1619 {
		x, err := newXTS(unhex(t, v.key))
		if err != nil {
			t.Fatal(i, err)
		}

		plain, want := unhex(t, v.plain), unhex(t, v.cipher)
		got := make([]byte, len(plain))
		x.encrypt(got, plain, v.unit, 0)
		if !bytes.Equal(got, want) {
			t.Errorf("vector %d: encrypt\n got %x\nwant %x", i, got, want)
		}

		x.decrypt(got, want, v.unit, 0)
		if !bytes.Equal(got, plain) {
			t.Errorf("vector %d: decrypt\n got %x\nwant %x", i, got, plain)
		}

		// From any block of the unit on, and in place.
		for b := 0; b < len(plain)/blockSize; b++ {
			part := bytes.Clone(plain[b*blockSize:])
			x.encrypt(part, part, v.unit, b)
			if !bytes.Equal(part, want[b*blockSize:]) {
				t.Errorf("vector %d: encrypt from block %d", i, b)
			}

			x.decrypt(part, part, v.unit, b)
			if !bytes.Equal(part, plain[b*blockSize:]) {
				t.Errorf("vector %d: decrypt from block %d", i, b)
			}
		}
	}
}

// A range of a file is encrypted the same whether it goes in one call or in
// pieces, and every 4096 bytes are a data unit of their own.
func TestCryptAt(t *testing.T) {
	x, err := derive(bytes.Repeat([]byte{7}, KeySize), labelDatabase)
	if err != nil {
		t.Fatal(err)
	}

	rnd := rand.New(rand.NewSource(1))
	plain := make([]byte, 3*unitSize+5*blockSize)
	rnd.Read(plain)
	const base = 7 * unitSize
	whole := make([]byte, len(plain))
	x.encryptAt(whole, plain, base)

	for u := 0; u*unitSize < len(plain); u++ {
		end := min((u+1)*unitSize, len(plain))
		want := make([]byte, end-u*unitSize)
		x.encrypt(want, plain[u*unitSize:end], uint64(7+u), 0)
		if !bytes.Equal(whole[u*unitSize:end], want) {
			t.Fatalf("unit %d is not encrypted on its own", u)
		}
	}

	for range 2000 {
		a := rnd.Intn(len(plain)/blockSize) * blockSize
		b := a + rnd.Intn((len(plain)-a)/blockSize+1)*blockSize
		got := make([]byte, b-a)
		x.encryptAt(got, plain[a:b], base+int64(a))
		if !bytes.Equal(got, whole[a:b]) {
			t.Fatalf("encrypt [%d, %d) differs from the whole", a, b)
		}

		x.decryptAt(got, got, base+int64(a))
		if !bytes.Equal(got, plain[a:b]) {
			t.Fatalf("decrypt [%d, %d)", a, b)
		}
	}
}

func TestMulAlpha(t *testing.T) {
	// alpha^128 = alpha^7 + alpha^2 + alpha + 1.
	lo, hi := uint64(1), uint64(0)
	for range 128 {
		lo, hi = mulAlpha(lo, hi)
	}

	if lo != 0x87 || hi != 0 {
		t.Fatalf("alpha^128 = %#x %#x", hi, lo)
	}
}

func TestDerive(t *testing.T) {
	key := bytes.Repeat([]byte{1}, KeySize)
	if _, err := derive(key[:31], labelDatabase); err == nil {
		t.Fatal("a key of 31 bytes was taken")
	}

	block := make([]byte, blockSize)
	seen := map[string]string{}
	for _, label := range []string{labelDatabase, labelJournal, labelJournalLength, labelKeyCheck} {
		x, err := derive(key, label)
		if err != nil {
			t.Fatal(err)
		}

		out := make([]byte, blockSize)
		x.encrypt(out, block, 0, 0)
		if other, ok := seen[string(out)]; ok {
			t.Fatalf("%q and %q give one key", label, other)
		}

		seen[string(out)] = label
	}
}

func BenchmarkEncryptPage(b *testing.B) {
	x, err := derive(make([]byte, KeySize), labelDatabase)
	if err != nil {
		b.Fatal(err)
	}

	page := make([]byte, 4096)
	b.SetBytes(int64(len(page)))
	for i := 0; b.Loop(); i++ {
		x.encryptAt(page, page, int64(i)*4096)
	}
}
