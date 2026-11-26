// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xts

// The storage of the crash tests: the one the tests of package vfs use for
// the same purpose, copied here because a test file cannot be imported.

import (
	"fmt"
	"io"
	"io/fs"
	"maps"
	"math/rand"
	"slices"
	"sort"
	"strings"
	"sync"

	"modernc.org/sqlite/vfs"
)

// crashStore is the storage behind the VFS the tests register. Next to what a
// running process sees of each file it keeps what would survive a loss of
// power, which is what the last Sync of the file made durable, so that a test
// can reopen the storage as it would be after a crash.
type crashStore struct {
	mu      sync.Mutex
	files   map[string]*crashData // the files as the process sees them
	durable map[string][]byte     // the files as the medium holds them

	live  int      // files open through Open and not closed yet
	ops   int      // writes, truncates, syncs, creations and deletions so far
	trace []string // one line for each of them

	// fail, if set, is asked before each counted operation and its error is
	// returned in place of doing it.
	fail func(n int, op, name string) error
	// after, if set, runs after each counted operation, with mu held.
	after func(n int)

	canon    func(name string) string              // FullPathname
	outFlags func(flags vfs.OpenFlag) vfs.OpenFlag // what Open reports
	opened   []string                              // every name passed to Open
	readOnly bool                                  // Open refuses to write
}

type crashData struct {
	b       []byte
	base    []byte       // b at the last Sync of this file, nil if never synced
	pending []crashWrite // what was done to the file since
}

// crashWrite is a write, or with data == nil a truncation to off.
type crashWrite struct {
	off  int64
	data []byte
}

func newCrashStore() *crashStore {
	return &crashStore{files: map[string]*crashData{}, durable: map[string][]byte{}}
}

// reopened returns a new storage holding the given files, all of them durable.
func reopened(image map[string][]byte) *crashStore {
	s := newCrashStore()
	for name, b := range image {
		b = append([]byte(nil), b...)
		s.files[name] = &crashData{b: b, base: append([]byte(nil), b...)}
		s.durable[name] = append([]byte(nil), b...)
	}
	return s
}

// count records a counted operation and returns the error to inject, if any.
// The caller holds mu.
func (s *crashStore) count(op, name string, format string, a ...any) error {
	s.ops++
	s.trace = append(s.trace, fmt.Sprintf("%s %s", op, kindOf(name))+fmt.Sprintf(format, a...))
	if s.fail != nil {
		if err := s.fail(s.ops, op, name); err != nil {
			return err
		}
	}
	return nil
}

func (s *crashStore) done() {
	if s.after != nil {
		s.after(s.ops)
	}
}

// kindOf names a file by its role, so that traces do not depend on names.
func kindOf(name string) string {
	switch {
	case strings.HasSuffix(name, "-journal"):
		return "journal"
	case strings.HasSuffix(name, "-wal"):
		return "wal"
	case strings.Contains(name, "-mj"):
		return "super"
	}
	return "db"
}

func (s *crashStore) Open(name string, flags vfs.OpenFlag) (vfs.File, vfs.OpenFlag, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.opened = append(s.opened, name)
	d := s.files[name]
	switch {
	case d == nil && flags&vfs.OpenCreate == 0:
		return nil, 0, fs.ErrNotExist
	case d != nil && flags&vfs.OpenCreate != 0 && flags&vfs.OpenExclusive != 0:
		return nil, 0, fs.ErrExist
	case d == nil && s.readOnly:
		return nil, 0, fs.ErrPermission
	case d == nil:
		if err := s.count("create", name, ""); err != nil {
			return nil, 0, err
		}

		d = &crashData{}
		s.files[name] = d
		s.done()
	}

	out := flags
	if s.readOnly {
		out = out&^vfs.OpenReadWrite | vfs.OpenReadOnly
	}
	if s.outFlags != nil {
		out = s.outFlags(out)
	}
	s.live++
	return &crashFile{s: s, name: name, d: d}, out, nil
}

func (s *crashStore) Delete(name string, syncDir bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.count("delete", name, " syncDir=%v", syncDir); err != nil {
		return err
	}

	if s.files[name] == nil {
		return fs.ErrNotExist
	}

	delete(s.files, name)
	if syncDir {
		delete(s.durable, name)
	}
	s.done()
	return nil
}

func (s *crashStore) Access(name string, flags vfs.AccessFlag) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.files[name] != nil && !(flags == vfs.AccessReadWrite && s.readOnly), nil
}

func (s *crashStore) FullPathname(name string) (string, error) {
	if s.canon != nil {
		return s.canon(name), nil
	}

	return name, nil
}

// content returns a copy of the named file as the process sees it.
func (s *crashStore) content(name string) []byte {
	s.mu.Lock()
	defer s.mu.Unlock()

	if d := s.files[name]; d != nil {
		return append([]byte(nil), d.b...)
	}

	return nil
}

func (s *crashStore) names() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var r []string
	for name := range s.files {
		r = append(r, name)
	}
	sort.Strings(r)
	return r
}

// What a loss of power may leave of the storage.
const (
	crashSynced  = iota // only what was synced
	crashWritten        // everything written, synced or not
	crashTorn           // of what was not synced, some sectors
	crashSector         // of what was not synced, some writes, others spoiling their sectors
)

// image returns the files a loss of power at this moment may leave. The caller
// holds mu. The files are gone through by name, so that one seed gives one
// image.
func (s *crashStore) image(how int, rng *rand.Rand) map[string][]byte {
	const sector = 512
	img := map[string][]byte{}
	files, durable := slices.Sorted(maps.Keys(s.files)), slices.Sorted(maps.Keys(s.durable))
	switch how {
	case crashSynced:
		for name, b := range s.durable {
			img[name] = append([]byte(nil), b...)
		}
	case crashWritten:
		for name, d := range s.files {
			img[name] = append([]byte(nil), d.b...)
		}
	case crashTorn:
		for _, name := range durable {
			b := s.durable[name]
			// A file deleted without a directory sync, or not yet synced
			// after it was created again, may be the old one still.
			if s.files[name] == nil || s.files[name].base == nil {
				if rng.Intn(2) == 0 {
					img[name] = append([]byte(nil), b...)
				}
			}
		}
		for _, name := range files {
			d := s.files[name]
			if d.base == nil {
				if _, ok := img[name]; ok || rng.Intn(2) == 0 {
					continue // the old file, or no file at all
				}
			}

			b := append([]byte(nil), d.base...)
			for _, w := range d.pending {
				if w.data == nil {
					if rng.Intn(2) == 0 {
						b = resize(b, w.off)
					}
					continue
				}

				for i := int64(0); i < int64(len(w.data)); i += sector {
					if rng.Intn(2) == 0 {
						continue
					}

					part := w.data[i:min(i+sector, int64(len(w.data)))]
					if end := w.off + i + int64(len(part)); end > int64(len(b)) {
						b = resize(b, end)
					}
					copy(b[w.off+i:], part)
				}
			}
			img[name] = b
		}
	case crashSector:
		// SQLite is told a sector of 4096 bytes: a write the power cut short
		// may leave anything in the sectors it touches.
		const block = 4096
		for _, name := range files {
			d := s.files[name]
			if d.base == nil {
				if b, ok := s.durable[name]; ok {
					img[name] = append([]byte(nil), b...) // the old file
				}
				continue
			}

			b := append([]byte(nil), d.base...)
			for _, w := range d.pending {
				switch what := rng.Intn(3); {
				case w.data == nil && what == 0:
					b = resize(b, w.off)
				case w.data == nil || what == 1:
					// not done
				case what == 0:
					if end := w.off + int64(len(w.data)); end > int64(len(b)) {
						b = resize(b, end)
					}
					copy(b[w.off:], w.data)
				default:
					// The file is as long as the write would have left
					// it: the damage is to what the sectors hold.
					end := w.off + int64(len(w.data))
					if end > int64(len(b)) {
						b = resize(b, end)
					}
					from, to := w.off/block*block, min((end+block-1)/block*block, int64(len(b)))
					for i := from; i < to; i++ {
						b[i] = 0xee
					}
				}
			}
			img[name] = b
		}
		for _, name := range durable {
			if s.files[name] == nil && rng.Intn(2) == 0 {
				img[name] = append([]byte(nil), s.durable[name]...)
			}
		}
	}
	return img
}

func resize(b []byte, size int64) []byte {
	if size <= int64(len(b)) {
		return b[:size]
	}

	return append(b, make([]byte, size-int64(len(b)))...)
}

type crashFile struct {
	s    *crashStore
	name string
	d    *crashData
}

func (f *crashFile) ReadAt(p []byte, off int64) (int, error) {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	if off >= int64(len(f.d.b)) {
		return 0, io.EOF
	}

	n := copy(p, f.d.b[off:])
	if n < len(p) {
		return n, io.EOF
	}

	return n, nil
}

func (f *crashFile) WriteAt(p []byte, off int64) (int, error) {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	if err := f.s.count("write", f.name, " %d+%d", off, len(p)); err != nil {
		return 0, err
	}

	if f.s.readOnly {
		return 0, fs.ErrPermission
	}

	if end := off + int64(len(p)); end > int64(len(f.d.b)) {
		f.d.b = resize(f.d.b, end)
	}
	copy(f.d.b[off:], p)
	f.d.pending = append(f.d.pending, crashWrite{off, append([]byte{}, p...)})
	f.s.done()
	return len(p), nil
}

func (f *crashFile) Truncate(size int64) error {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	if err := f.s.count("truncate", f.name, " %d", size); err != nil {
		return err
	}

	f.d.b = resize(f.d.b, size)
	f.d.pending = append(f.d.pending, crashWrite{off: size})
	f.s.done()
	return nil
}

func (f *crashFile) Sync(flags vfs.SyncFlag) error {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	if err := f.s.count("sync", f.name, " %#x", int32(flags)); err != nil {
		return err
	}

	// A file deleted while it is open is synced to nowhere.
	if f.s.files[f.name] == f.d {
		f.d.base = append([]byte{}, f.d.b...)
		f.d.pending = nil
		f.s.durable[f.name] = append([]byte(nil), f.d.b...)
	}
	f.s.done()
	return nil
}

func (f *crashFile) Size() (int64, error) {
	f.s.mu.Lock()
	defer f.s.mu.Unlock()

	return int64(len(f.d.b)), nil
}

func (f *crashFile) Close() error {
	f.s.mu.Lock()
	f.s.live--
	f.s.mu.Unlock()
	return nil
}
