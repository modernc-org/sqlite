// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xts keeps a database encrypted on its storage: a VFS that goes on
// top of another one, SQLite's own by default, and encrypts what SQLite
// writes through it with AES-256 in XTS mode.
//
//	name, closer, err := xts.Register("", key)
//	if err != nil {
//		return err
//	}
//
//	defer closer.Close()
//
//	db, err := sql.Open("sqlite", "file:app.db?vfs="+name)
//
// The key is 32 bytes that nobody else knows. It is given to [Register] and
// to nothing else: not to a DSN, which ends up in logs and in configuration.
// The package derives no key from a password.
//
// Locking, also between processes, and the code for each platform are those
// of the VFS underneath. An encrypted database is used like any other, with
// the differences this text lists.
//
// # What it protects
//
// Without the key nobody reads the content of the database, of its rollback
// journal or of the temporary files SQLite makes while it works on it.
//
// It is the protection a disk encryption gives, and no more:
//
//   - A file that was changed is not detected. Whoever can write the file
//     cannot make a change of his choosing, but he can spoil 16 bytes at a
//     time, and he can put back an older copy of the file or of any part of
//     it. Nothing is authenticated.
//   - The cipher is deterministic. Two copies of a file show which 16-byte
//     blocks changed between them, and two files under one key show which
//     blocks at equal offsets are equal. Use a key for one database only.
//   - The sizes of the files, their names and the times they change are
//     there for anyone to see.
//   - Two things are stored in plain text, both names of files and no
//     content: a super-journal, the file SQLite keeps for the time of a
//     transaction over several databases, which holds the names of their
//     journals, and the name of that file at the end of each such journal.
//   - The key is in the memory of the process for as long as the VFS is
//     registered.
//
// # Use it for the main database
//
// Open the encrypted database as the main database of a connection, not as
// one attached to a plain one. The temporary files of a connection, where
// sorts and temporary tables too large for memory go, are written through
// the VFS of its main database. With a plain main database they are plain,
// whatever is attached, unless "PRAGMA temp_store=MEMORY" keeps them off the
// storage. SQLite also creates the super-journal through that VFS, and a
// transaction that changes an encrypted database together with another
// database commits only if the package saw that file being created: with an
// encrypted main database, to which plain databases and databases under
// other keys on the same underlying VFS may be attached. Otherwise the
// transaction fails with an I/O error and is rolled back. So does one whose
// super-journal has a name longer than SQLite itself reads back from a
// journal, the longest name of the VFS and a byte, which a path of the main
// database within 12 bytes of that limit gives. Two cases are left as they
// are without the cipher. A database in journal_mode OFF or MEMORY keeps no
// journal on the storage, is asked nothing and commits with the others as
// SQLite has it. And when the main database is in memory or temporary,
// SQLite uses no super-journal and commits the databases one after another.
//
// # What it does not do
//
// There is no write-ahead log. "PRAGMA journal_mode=WAL" fails with an
// error, also under "PRAGMA locking_mode=EXCLUSIVE", and a database that
// says in its header that it is in WAL mode is not written: such a header
// comes with a copy that SQLite's backup interface makes of a WAL-mode
// database, and that copy fails. If its destination has a journal, the
// destination is as it was; with "PRAGMA journal_mode=OFF" it is damaged, as
// after any failed write in that mode. Copy a WAL-mode database with VACUUM
// INTO instead, or switch it to a rollback mode first.
//
// There is no shared memory and no memory mapping, and the query parameter
// nolock=1 is refused: SQLite would then lock nothing and still act on
// journals, with no moment left at which the key could be checked first.
//
// The files are not those of SQLCipher or of the SQLite Encryption
// Extension, and not those of the vfs/xts package of
// github.com/ncruces/go-sqlite3 either, although the cipher is the same.
//
// # The wrong key, and programs without the key
//
// With another key, and through a VFS that does not encrypt, every statement
// that reads or writes the database fails with SQLITE_NOTADB, and no file is
// changed: the key is checked at the beginning of every transaction, before
// SQLite looks at a journal, and a stored journal begins in a way that keeps
// SQLite from taking it for the journal of a writer that died. Two cases are
// not covered. SQLite deletes a journal it finds next to an empty database
// file, with the cipher as without it. And a program that has had the file
// open, without the cipher, since it was empty can delete a journal by
// changing its journal mode.
//
// # Making and unmaking an encrypted database
//
// "VACUUM INTO 'file:new.db?vfs=NAME'", run on the plain database, writes an
// encrypted copy of it. The other way round the plain VFS has to be named,
// "unix" or "win32": a copy made from a connection that is on the encrypting
// VFS goes through it as well unless the URI says otherwise. The backup
// interface copies into an encrypted database too, from a source that is
// not in WAL mode.
//
// # What it costs
//
// AES itself, about 6 microseconds for a page of 4096 bytes on a recent
// processor. Beyond that the VFS tells SQLite less than SQLite's own does:
// that a write may damage the whole sector it touches. SQLite then journals
// by sector and begins a journal with a header of a sector, 4096 bytes,
// where it otherwise writes 512; with the prefix and the length block a
// journal is some 7700 bytes longer than it would be, and more with every
// further header a long transaction puts into it. And every
// write to a journal first reads where the stored journal ends, its size,
// its prefix and its length block, about four small reads for a write, and
// the blocks at its two ends if it does not begin and end on 16 bytes.
// Writes to the database file cost no read.
//
// # Format 1
//
// What [Register] writes is fixed: a later format will be asked for in
// another way, and files written now stay readable.
//
// The cipher is XTS-AES-256 as in IEEE Std 1619, without ciphertext
// stealing: only whole 16-byte blocks are encrypted. A data unit is 4096
// bytes. The tweak of a data unit is its number as a 64-bit little-endian
// integer in the first 8 bytes of a block whose other bytes are zero,
// encrypted with the tweak key; block j of the unit uses that value times
// alpha to the power j. A range of a file that begins and ends on a block is
// therefore encrypted on its own, whatever is around it.
//
// Each XTS key is 64 bytes of HKDF-SHA256 (RFC 5869) with the key of the
// registration as the input keying material, no salt, and one of these
// strings as the info: the first 32 bytes are the AES-256 key of the data,
// the last 32 that of the tweak.
//
//	modernc.org/sqlite/vfs/xts format 1: database
//	modernc.org/sqlite/vfs/xts format 1: journal
//	modernc.org/sqlite/vfs/xts format 1: journal length
//	modernc.org/sqlite/vfs/xts format 1: key check
//
// A database file: the block at offset x of the content is the block at
// offset x of the stored file, encrypted under the database key with the
// tweak of its place, x/4096 and (x mod 4096)/16. Nothing is added and
// nothing is moved; the file is as long as SQLite made it, which is whole
// pages and so whole blocks.
//
// A temporary file: the same, under 64 random bytes as the XTS key, made for
// that file alone. If SQLite ends it inside a block, the block is completed
// with zeros.
//
// A rollback journal of n bytes is stored as a prefix, the blocks of its
// content, a length block and, in one case, a copy.
//
//   - The prefix is as long as the sector the VFS reports, at least 4096
//     bytes, a power of two. It holds a zero byte, the 14 bytes "sqlite
//     vfs/xts", the byte 1, its own length as a 32-bit little-endian integer,
//     12 zero bytes, and the key check block: 16 bytes, the 15 of "xts key
//     check 1" and a zero byte, encrypted under the key check key as block 0
//     of data unit 0. The rest is zeros.
//   - The block at offset x of the content follows the prefix at offset x,
//     encrypted under the journal key with the tweak of x. If the content
//     does not end on a block, the rest of the last block is not content.
//   - The length block follows the last block of the content. It is the 8
//     bytes "xtsJLEN1" and n as a 64-bit little-endian integer, encrypted
//     under the journal length key with the tweak of the offset in the
//     content at which it lies, n rounded up to 16.
//   - If the content ends with the record SQLite writes for a super-journal,
//     and the package was given that name when SQLite created the
//     super-journal, the record follows the length block in plain text, as
//     SQLite reads it from the end of a file: the name, its length and its
//     checksum as 32-bit big-endian integers, and the 8 bytes d9 d5 05 f9 20
//     a1 63 d7.
//
// A journal without content is a stored file without content: no prefix and
// no length block. The prefix is written with the first content.
//
// What a reader of a stored journal does, which a later version has to go
// on doing for files written now:
//
//   - A prefix is accepted if its length is a power of two from 4096 to
//     65536 and not more than the file, and its key check block is the one
//     of the key. A file without such a prefix holds no journal.
//   - If the file ends with the 8 bytes of a record, and the length field 16
//     bytes before the end says a name of n bytes, and the block n+32 bytes
//     before the end is a valid length block, and the n+16 bytes at the end
//     of the content it tells are the n+16 bytes at the end of the file,
//     then that is the length block and the end of the file is the copy.
//     Otherwise the last block of the file is the length block, if it is a
//     valid one: its length is not zero and, rounded up to 16, is the offset
//     in the content at which the block lies.
//   - Without a valid length block the content is the whole blocks that
//     follow the prefix: the rest of an interrupted write.
//   - A record at the end of the content counts only with its copy. A
//     journal that ends with a record and has no copy of it is read with an
//     end at which SQLite finds no record: behind the length block if there
//     is one, else one or more blocks before the end of the content.
//   - What lies between an earlier end of the content and a later write
//     beyond it is arbitrary; earlier length blocks are among it. SQLite
//     does not read there.
//
// A file SQLite opens as a super-journal is stored as it is.
package xts // import "modernc.org/sqlite/vfs/xts"
