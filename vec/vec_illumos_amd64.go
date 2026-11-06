// Code generated for illumos/amd64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build illumos && amd64

package vec

import (
	"unsafe"

	"modernc.org/libc"
	libsqlite3 "modernc.org/sqlite/lib"
)

type TFILE = struct {
	F__pad [16]int64
}

type Tdecimal_mode = struct {
	Frd      _fp_direction_type
	Fdf      _decimal_form
	Fndigits int32
}

type Tdecimal_record = struct {
	Ffpclass  _fp_class_type
	Fsign     int32
	Fexponent int32
	Fds       Tdecimal_string
	Fmore     int32
	Fndigits  int32
}

type Tdecimal_string = [512]int8

type Textended = [3]uint32

type Tfp_exception_field_type = uint32

type Tfpos64_t = int64

type Tquadruple = float64

type Tsigfpe_code_type = int32

type Tsigfpe_handler_type = uintptr

type Tsingle = float32

func Xbitmap_clear(tls *libc.TLS, bitmap uintptr, n Ti32) {
	libc.Xmemset(tls, bitmap, 0, libc.Uint64FromInt32(n/int32(m___CHAR_BIT__)))
}

func Xbitmap_copy(tls *libc.TLS, base uintptr, from uintptr, n Ti32) {
	libc.Xmemcpy(tls, base, from, libc.Uint64FromInt32(n/int32(m___CHAR_BIT__)))
}

func Xbitmap_fill(tls *libc.TLS, bitmap uintptr, n Ti32) {
	libc.Xmemset(tls, bitmap, int32(0xFF), libc.Uint64FromInt32(n/int32(m___CHAR_BIT__)))
}

func Xbitmap_new(tls *libc.TLS, n Ti32) (r uintptr) {
	var p uintptr
	_ = p
	p = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt32(n)*uint64(1)/uint64(m___CHAR_BIT__)))
	if p != 0 {
		libc.Xmemset(tls, p, 0, libc.Uint64FromInt32(n)*uint64(1)/uint64(m___CHAR_BIT__))
	}
	return p
}

func Xbitmap_new_from(tls *libc.TLS, n Ti32, from uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt32(n)*uint64(1)/uint64(m___CHAR_BIT__)))
	if p != 0 {
		libc.Xmemcpy(tls, p, from, libc.Uint64FromInt32(n/int32(m___CHAR_BIT__)))
	}
	return p
}

func Xvec0Update_Delete_ClearMetadata(tls *libc.TLS, p uintptr, metadata_idx int32, rowid Ti64, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var kind Tvec0_metadata_column_kind
	var rc, rc2 int32
	var zSql uintptr
	var _ /* blobValue at bp+0 */ uintptr
	var _ /* block at bp+8 */ Tu8
	var _ /* n at bp+32 */ int32
	var _ /* stmt at bp+56 */ uintptr
	var _ /* v at bp+16 */ Ti64
	var _ /* v at bp+24 */ float64
	var _ /* view at bp+36 */ [16]Tu8
	_, _, _, _ = kind, rc, rc2, zSql
	kind = (*(*TVec0MetadataColumnDefinition)(__ccgo_ptr(p + 1600 + uintptr(metadata_idx)*24))).Fkind
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, *(*uintptr)(__ccgo_ptr(p + 480 + uintptr(metadata_idx)*8)), __ccgo_ts+4053, chunk_id, int32(1), bp)
	if rc != m_SQLITE_OK {
		return rc
	}
	switch kind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		rc = libsqlite3.Xsqlite3_blob_read(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+8, int32(1), libc.Int32FromUint64(chunk_offset/libc.Uint64FromInt32(m___CHAR_BIT__)))
		if rc != m_SQLITE_OK {
			goto done
		}
		*(*Tu8)(__ccgo_ptr(bp + 8)) = libc.Uint8FromInt32(int32(*(*Tu8)(__ccgo_ptr(bp + 8))) & ^(libc.Int32FromInt32(1) << (chunk_offset % libc.Uint64FromInt32(m___CHAR_BIT__))))
		rc = libsqlite3.Xsqlite3_blob_write(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+8, int32(1), libc.Int32FromUint64(chunk_offset/uint64(m___CHAR_BIT__)))
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		*(*Ti64)(__ccgo_ptr(bp + 16)) = 0
		rc = libsqlite3.Xsqlite3_blob_write(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+16, int32(8), libc.Int32FromUint64(chunk_offset*uint64(8)))
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		*(*float64)(__ccgo_ptr(bp + 24)) = libc.Float64FromInt32(0)
		rc = libsqlite3.Xsqlite3_blob_write(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+24, int32(8), libc.Int32FromUint64(chunk_offset*uint64(8)))
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		rc = libsqlite3.Xsqlite3_blob_read(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+32, int32(4), libc.Int32FromUint64(chunk_offset*uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if rc != m_SQLITE_OK {
			goto done
		}
		libc.Xmemset(tls, bp+36, 0, uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		rc = libsqlite3.Xsqlite3_blob_write(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+36, int32(16), libc.Int32FromUint64(chunk_offset*uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if rc != m_SQLITE_OK {
			goto done
		}
		if *(*int32)(__ccgo_ptr(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12987, libc.VaList(bp+72, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_idx))
			if !(zSql != 0) {
				rc = int32(m_SQLITE_NOMEM)
				goto done
			}
			rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+56, libc.UintptrFromInt32(0))
			if rc != m_SQLITE_OK {
				goto done
			}
			libsqlite3.Xsqlite3_bind_int64(tls, *(*uintptr)(__ccgo_ptr(bp + 56)), int32(1), rowid)
			rc = libsqlite3.Xsqlite3_step(tls, *(*uintptr)(__ccgo_ptr(bp + 56)))
			libsqlite3.Xsqlite3_finalize(tls, *(*uintptr)(__ccgo_ptr(bp + 56)))
			if rc != int32(m_SQLITE_DONE) {
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
			// Fix for https://github.com/asg017/sqlite-vec/issues/274
			// sqlite3_step returns SQLITE_DONE (101) on DML success, but the
			// `done:` epilogue treats anything other than SQLITE_OK as an error.
			// Without this, SQLITE_DONE propagates up to vec0Update_Delete,
			// which aborts the DELETE scan and silently drops remaining rows.
			rc = m_SQLITE_OK
		}
		break
	}
	goto done
done:
	;
	rc2 = libsqlite3.Xsqlite3_blob_close(tls, *(*uintptr)(__ccgo_ptr(bp)))
	if rc == m_SQLITE_OK {
		return rc2
	}
	return rc
}

func Xvec0_write_metadata_value(tls *libc.TLS, p uintptr, metadata_column_idx int32, rowid Ti64, chunk_id Ti64, chunk_offset Ti64, v uintptr, isupdate int32) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var kind Tvec0_metadata_column_kind
	var metadata_column, s, zSql, zSql1 uintptr
	var rc, value, v1 int32
	var _ /* blobValue at bp+0 */ uintptr
	var _ /* block at bp+8 */ Tu8
	var _ /* n at bp+36 */ int32
	var _ /* prev_n at bp+32 */ int32
	var _ /* stmt at bp+56 */ uintptr
	var _ /* stmt at bp+64 */ uintptr
	var _ /* value at bp+16 */ Ti64
	var _ /* value at bp+24 */ float64
	var _ /* view at bp+40 */ [16]Tu8
	_, _, _, _, _, _, _, _ = kind, metadata_column, rc, s, value, zSql, zSql1, v1
	metadata_column = p + 1600 + uintptr(metadata_column_idx)*24
	kind = (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fkind
	// verify input value matches column type
	switch kind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		if libsqlite3.Xsqlite3_value_type(tls, v) != int32(m_SQLITE_INTEGER) || libsqlite3.Xsqlite3_value_int(tls, v) != 0 && libsqlite3.Xsqlite3_value_int(tls, v) != int32(1) {
			rc = int32(m_SQLITE_ERROR)
			Xvtab_set_error(tls, p, __ccgo_ts+12625, libc.VaList(bp+80, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname_length, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname))
			goto done
		}
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		if libsqlite3.Xsqlite3_value_type(tls, v) != int32(m_SQLITE_INTEGER) {
			rc = int32(m_SQLITE_ERROR)
			Xvtab_set_error(tls, p, __ccgo_ts+12674, libc.VaList(bp+80, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname_length, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname, Xtype_name(tls, libsqlite3.Xsqlite3_value_type(tls, v))))
			goto done
		}
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		if libsqlite3.Xsqlite3_value_type(tls, v) != int32(m_SQLITE_FLOAT) {
			rc = int32(m_SQLITE_ERROR)
			Xvtab_set_error(tls, p, __ccgo_ts+12737, libc.VaList(bp+80, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname_length, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname, Xtype_name(tls, libsqlite3.Xsqlite3_value_type(tls, v))))
			goto done
		}
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		if libsqlite3.Xsqlite3_value_type(tls, v) != int32(m_SQLITE_TEXT) {
			rc = int32(m_SQLITE_ERROR)
			Xvtab_set_error(tls, p, __ccgo_ts+12796, libc.VaList(bp+80, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname_length, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname, Xtype_name(tls, libsqlite3.Xsqlite3_value_type(tls, v))))
			goto done
		}
		break
	}
	*(*uintptr)(__ccgo_ptr(bp)) = libc.UintptrFromInt32(0)
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, *(*uintptr)(__ccgo_ptr(p + 480 + uintptr(metadata_column_idx)*8)), __ccgo_ts+4053, chunk_id, int32(1), bp)
	if rc != m_SQLITE_OK {
		goto done
	}
	switch kind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		value = libsqlite3.Xsqlite3_value_int(tls, v)
		rc = libsqlite3.Xsqlite3_blob_read(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+8, int32(1), int32(chunk_offset/libc.Int64FromInt32(m___CHAR_BIT__)))
		if rc != m_SQLITE_OK {
			goto done
		}
		if value != 0 {
			*(*Tu8)(__ccgo_ptr(bp + 8)) = libc.Uint8FromInt32(int32(*(*Tu8)(__ccgo_ptr(bp + 8))) | libc.Int32FromInt32(1)<<(chunk_offset%libc.Int64FromInt32(m___CHAR_BIT__)))
		} else {
			*(*Tu8)(__ccgo_ptr(bp + 8)) = libc.Uint8FromInt32(int32(*(*Tu8)(__ccgo_ptr(bp + 8))) & ^(libc.Int32FromInt32(1) << (chunk_offset % libc.Int64FromInt32(m___CHAR_BIT__))))
		}
		rc = libsqlite3.Xsqlite3_blob_write(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+8, int32(1), int32(chunk_offset/int64(m___CHAR_BIT__)))
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		*(*Ti64)(__ccgo_ptr(bp + 16)) = libsqlite3.Xsqlite3_value_int64(tls, v)
		rc = libsqlite3.Xsqlite3_blob_write(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+16, int32(8), libc.Int32FromUint64(libc.Uint64FromInt64(chunk_offset)*uint64(8)))
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		*(*float64)(__ccgo_ptr(bp + 24)) = libsqlite3.Xsqlite3_value_double(tls, v)
		rc = libsqlite3.Xsqlite3_blob_write(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+24, int32(8), libc.Int32FromUint64(libc.Uint64FromInt64(chunk_offset)*uint64(8)))
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		rc = libsqlite3.Xsqlite3_blob_read(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+32, int32(4), int32(chunk_offset*int64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if rc != m_SQLITE_OK {
			goto done
		}
		s = libsqlite3.Xsqlite3_value_text(tls, v)
		*(*int32)(__ccgo_ptr(bp + 36)) = libsqlite3.Xsqlite3_value_bytes(tls, v)
		libc.Xmemset(tls, bp+40, 0, uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		libc.Xmemcpy(tls, bp+40, bp+36, uint64(4))
		if *(*int32)(__ccgo_ptr(bp + 36)) <= libc.Int32FromInt32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)-libc.Int32FromInt32(4) {
			v1 = *(*int32)(__ccgo_ptr(bp + 36))
		} else {
			v1 = libc.Int32FromInt32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH) - libc.Int32FromInt32(4)
		}
		libc.Xmemcpy(tls, bp+40+uintptr(4), s, libc.Uint64FromInt32(v1))
		rc = libsqlite3.Xsqlite3_blob_write(tls, *(*uintptr)(__ccgo_ptr(bp)), bp+40, int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH), int32(chunk_offset*int64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if *(*int32)(__ccgo_ptr(bp + 36)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			if isupdate != 0 && *(*int32)(__ccgo_ptr(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12853, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
			} else {
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12918, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
			}
			if !(zSql != 0) {
				rc = int32(m_SQLITE_NOMEM)
				goto done
			}
			rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+56, libc.UintptrFromInt32(0))
			if rc != m_SQLITE_OK {
				goto done
			}
			libsqlite3.Xsqlite3_bind_int64(tls, *(*uintptr)(__ccgo_ptr(bp + 56)), int32(1), rowid)
			libsqlite3.Xsqlite3_bind_text(tls, *(*uintptr)(__ccgo_ptr(bp + 56)), int32(2), s, *(*int32)(__ccgo_ptr(bp + 36)), libc.UintptrFromInt32(0))
			rc = libsqlite3.Xsqlite3_step(tls, *(*uintptr)(__ccgo_ptr(bp + 56)))
			libsqlite3.Xsqlite3_finalize(tls, *(*uintptr)(__ccgo_ptr(bp + 56)))
			if rc != int32(m_SQLITE_DONE) {
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
		} else {
			if *(*int32)(__ccgo_ptr(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12987, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
				if !(zSql1 != 0) {
					rc = int32(m_SQLITE_NOMEM)
					goto done
				}
				rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql1, -int32(1), bp+64, libc.UintptrFromInt32(0))
				if rc != m_SQLITE_OK {
					goto done
				}
				libsqlite3.Xsqlite3_bind_int64(tls, *(*uintptr)(__ccgo_ptr(bp + 64)), int32(1), rowid)
				rc = libsqlite3.Xsqlite3_step(tls, *(*uintptr)(__ccgo_ptr(bp + 64)))
				libsqlite3.Xsqlite3_finalize(tls, *(*uintptr)(__ccgo_ptr(bp + 64)))
				if rc != int32(m_SQLITE_DONE) {
					rc = int32(m_SQLITE_ERROR)
					goto done
				}
			}
		}
		break
	}
	if rc != m_SQLITE_OK {
	}
	rc = libsqlite3.Xsqlite3_blob_close(tls, *(*uintptr)(__ccgo_ptr(bp)))
	if rc != m_SQLITE_OK {
		goto done
	}
	goto done
done:
	;
	return rc
	return r
}

const _ansi_1 = 1

const _c_issue_4 = 0

type _decimal_form = int32

type _decimal_string_form = int32

const _fixed_dotfrac_form = 4

const _fixed_form = 0

const _fixed_int_form = 2

const _fixed_intdot_form = 3

const _fixed_intdotfrac_form = 5

const _floating_dotfrac_form = 8

const _floating_form = 1

const _floating_int_form = 6

const _floating_intdot_form = 7

const _floating_intdotfrac_form = 9

type _fp_class_type = int32

const _fp_denormalized = 1

type _fp_direction_type = int32

const _fp_division = 2

const _fp_double = 2

type _fp_exception_type = int32

const _fp_extended = 3

const _fp_inexact = 5

const _fp_infinity = 3

const _fp_invalid = 0

const _fp_nearest = 0

const _fp_negative = 1

const _fp_normal = 2

const _fp_overflow = 3

const _fp_positive = 2

const _fp_precision_3 = 1

type _fp_precision_type = int32

const _fp_quiet = 4

const _fp_signaling = 5

const _fp_single = 0

const _fp_subnormal = 1

const _fp_tozero = 3

const _fp_trap_denormalized = 1

const _fp_trap_division = 2

type _fp_trap_enable_type = int32

const _fp_trap_inexact = 5

const _fp_trap_invalid = 0

const _fp_trap_overflow = 3

const _fp_trap_underflow = 4

const _fp_underflow = 4

const _fp_zero = 0

func _fvec_from_value(tls *libc.TLS, value uintptr, vector unsafe.Pointer, dimensions unsafe.Pointer, __ccgo_fp_cleanup unsafe.Pointer, pzErr unsafe.Pointer) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var blob, buf, ptr, source uintptr
	var bytes, i, offset, rc, source_len, value_type int32
	var result float64
	var x TArray
	var _ /* endptr at bp+0 */ uintptr
	var _ /* res at bp+8 */ Tf32
	_, _, _, _, _, _, _, _, _, _, _, _ = blob, buf, bytes, i, offset, ptr, rc, result, source, source_len, value_type, x
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	if value_type == int32(m_SQLITE_BLOB) {
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		if bytes == 0 {
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		if libc.Uint64FromInt32(bytes)%uint64(4) != uint64(0) {
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+86, libc.VaList(bp+24, uint64(4), bytes))
			return int32(m_SQLITE_ERROR)
		}
		buf = libsqlite3.Xsqlite3_malloc(tls, bytes)
		if !(buf != 0) {
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+156, 0)
			return int32(m_SQLITE_NOMEM)
		}
		libc.Xmemcpy(tls, buf, blob, libc.Uint64FromInt32(bytes))
		*(*uintptr)(vector) = buf
		*(*Tsize_t)(dimensions) = libc.Uint64FromInt32(bytes) / uint64(4)
		*(*uintptr)(__ccgo_fp_cleanup) = __ccgo_fp(libsqlite3.Xsqlite3_free)
		return m_SQLITE_OK
	}
	if value_type == int32(m_SQLITE_TEXT) {
		source = libsqlite3.Xsqlite3_value_text(tls, value)
		source_len = libsqlite3.Xsqlite3_value_bytes(tls, value)
		if source_len == 0 {
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		i = 0
		rc = _array_init__ccgo_gp(tls, unsafe.Pointer(&x), uint64(4), uint64(libc.Xceil(tls, float64(source_len)/float64(2))))
		if rc != m_SQLITE_OK {
			return rc
		}
		// advance leading whitespace to first '['
		for i < source_len {
			if _vecJsonIsSpaceX[libc.Uint8FromInt8(*(*int8)(__ccgo_ptr(source + uintptr(i))))] != 0 {
				i = i + 1
				continue
			}
			if int32(*(*int8)(__ccgo_ptr(source + uintptr(i)))) == int32('[') {
				break
			}
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			_array_cleanup__ccgo_gp(tls, unsafe.Pointer(&x))
			return int32(m_SQLITE_ERROR)
		}
		if int32(*(*int8)(__ccgo_ptr(source + uintptr(i)))) != int32('[') {
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			_array_cleanup__ccgo_gp(tls, unsafe.Pointer(&x))
			return int32(m_SQLITE_ERROR)
		}
		offset = i + int32(1)
		for offset < source_len {
			ptr = source + uintptr(offset)
			*(*int32)(__ccgo_ptr(libc.X___errno(tls))) = 0
			result = libc.Xstrtod(tls, ptr, bp)
			if *(*int32)(__ccgo_ptr(libc.X___errno(tls))) != 0 && result == libc.Float64FromInt32(0) || *(*int32)(__ccgo_ptr(libc.X___errno(tls))) == int32(m_ERANGE) && (result == libc.X__builtin_huge_val(tls) || result == -libc.X__builtin_huge_val(tls)) {
				libsqlite3.Xsqlite3_free(tls, x.Fz)
				*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if *(*uintptr)(__ccgo_ptr(bp)) == ptr {
				if int32(*(*int8)(__ccgo_ptr(ptr))) != int32(']') {
					libsqlite3.Xsqlite3_free(tls, x.Fz)
					*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
					return int32(m_SQLITE_ERROR)
				}
				goto done
			}
			*(*Tf32)(__ccgo_ptr(bp + 8)) = float32(result)
			_array_append__ccgo_gp(tls, unsafe.Pointer(&x), bp+8)
			offset = int32(int64(offset) + (int64(*(*uintptr)(__ccgo_ptr(bp))) - int64(ptr)))
			for offset < source_len {
				if _vecJsonIsSpaceX[libc.Uint8FromInt8(*(*int8)(__ccgo_ptr(source + uintptr(offset))))] != 0 {
					offset = offset + 1
					continue
				}
				if int32(*(*int8)(__ccgo_ptr(source + uintptr(offset)))) == int32(',') {
					offset = offset + 1
					continue
				}
				if int32(*(*int8)(__ccgo_ptr(source + uintptr(offset)))) == int32(']') {
					goto done
				}
				break
			}
		}
		goto done
	done:
		;
		if x.Flength > uint64(0) {
			*(*uintptr)(vector) = x.Fz
			*(*Tsize_t)(dimensions) = x.Flength
			*(*uintptr)(__ccgo_fp_cleanup) = __ccgo_fp(libsqlite3.Xsqlite3_free)
			return m_SQLITE_OK
		}
		libsqlite3.Xsqlite3_free(tls, x.Fz)
		*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
		return int32(m_SQLITE_ERROR)
	}
	*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+245, libc.VaList(bp+24, Xtype_name(tls, value_type)))
	return int32(m_SQLITE_ERROR)
}

const _inf_form = 10

const _infinity_form = 11

func _int8_vec_from_value(tls *libc.TLS, value uintptr, vector unsafe.Pointer, dimensions unsafe.Pointer, __ccgo_fp_cleanup unsafe.Pointer, pzErr unsafe.Pointer) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var blob, ptr, source uintptr
	var bytes, i, offset, rc, source_len, value_type int32
	var result int64
	var x TArray
	var _ /* endptr at bp+0 */ uintptr
	var _ /* res at bp+8 */ Ti8
	_, _, _, _, _, _, _, _, _, _, _ = blob, bytes, i, offset, ptr, rc, result, source, source_len, value_type, x
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	if value_type == int32(m_SQLITE_BLOB) {
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		if bytes == 0 {
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		*(*uintptr)(vector) = blob
		*(*Tsize_t)(dimensions) = libc.Uint64FromInt32(bytes)
		*(*uintptr)(__ccgo_fp_cleanup) = __ccgo_fp(Xvector_cleanup_noop)
		return m_SQLITE_OK
	}
	if value_type == int32(m_SQLITE_TEXT) {
		source = libsqlite3.Xsqlite3_value_text(tls, value)
		source_len = libsqlite3.Xsqlite3_value_bytes(tls, value)
		i = 0
		if source_len == 0 {
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		rc = _array_init__ccgo_gp(tls, unsafe.Pointer(&x), uint64(1), uint64(libc.Xceil(tls, float64(source_len)/float64(2))))
		if rc != m_SQLITE_OK {
			return rc
		}
		// advance leading whitespace to first '['
		for i < source_len {
			if _vecJsonIsSpaceX[libc.Uint8FromInt8(*(*int8)(__ccgo_ptr(source + uintptr(i))))] != 0 {
				i = i + 1
				continue
			}
			if int32(*(*int8)(__ccgo_ptr(source + uintptr(i)))) == int32('[') {
				break
			}
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			_array_cleanup__ccgo_gp(tls, unsafe.Pointer(&x))
			return int32(m_SQLITE_ERROR)
		}
		if int32(*(*int8)(__ccgo_ptr(source + uintptr(i)))) != int32('[') {
			*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			_array_cleanup__ccgo_gp(tls, unsafe.Pointer(&x))
			return int32(m_SQLITE_ERROR)
		}
		offset = i + int32(1)
		for offset < source_len {
			ptr = source + uintptr(offset)
			*(*int32)(__ccgo_ptr(libc.X___errno(tls))) = 0
			result = libc.Xstrtol(tls, ptr, bp, int32(10))
			if *(*int32)(__ccgo_ptr(libc.X___errno(tls))) != 0 && result == 0 || *(*int32)(__ccgo_ptr(libc.X___errno(tls))) == int32(m_ERANGE) && (result == int64(0x7fffffffffffffff) || result == -libc.Int64FromInt64(0x7fffffffffffffff)-libc.Int64FromInt64(1)) {
				libsqlite3.Xsqlite3_free(tls, x.Fz)
				*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if *(*uintptr)(__ccgo_ptr(bp)) == ptr {
				if int32(*(*int8)(__ccgo_ptr(ptr))) != int32(']') {
					libsqlite3.Xsqlite3_free(tls, x.Fz)
					*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
					return int32(m_SQLITE_ERROR)
				}
				goto done
			}
			if result < int64(-libc.Int32FromInt32(128)) || result > int64(libc.Int32FromInt32(m_INT8_MAX)) {
				libsqlite3.Xsqlite3_free(tls, x.Fz)
				*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+341, 0)
				return int32(m_SQLITE_ERROR)
			}
			*(*Ti8)(__ccgo_ptr(bp + 8)) = int8(result)
			_array_append__ccgo_gp(tls, unsafe.Pointer(&x), bp+8)
			offset = int32(int64(offset) + (int64(*(*uintptr)(__ccgo_ptr(bp))) - int64(ptr)))
			for offset < source_len {
				if _vecJsonIsSpaceX[libc.Uint8FromInt8(*(*int8)(__ccgo_ptr(source + uintptr(offset))))] != 0 {
					offset = offset + 1
					continue
				}
				if int32(*(*int8)(__ccgo_ptr(source + uintptr(offset)))) == int32(',') {
					offset = offset + 1
					continue
				}
				if int32(*(*int8)(__ccgo_ptr(source + uintptr(offset)))) == int32(']') {
					goto done
				}
				break
			}
		}
		goto done
	done:
		;
		if x.Flength > uint64(0) {
			*(*uintptr)(vector) = x.Fz
			*(*Tsize_t)(dimensions) = x.Flength
			*(*uintptr)(__ccgo_fp_cleanup) = __ccgo_fp(libsqlite3.Xsqlite3_free)
			return m_SQLITE_OK
		}
		libsqlite3.Xsqlite3_free(tls, x.Fz)
		*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
		return int32(m_SQLITE_ERROR)
	}
	*(*uintptr)(pzErr) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+389, 0)
	return int32(m_SQLITE_ERROR)
}

const _invalid_form = 0

const _libm_ieee = -1

const _nan_form = 12

const _nanstring_form = 13

const _strict_ansi = 2

func _vec0Filter_knn_chunks_iter__ccgo_gp(tls *libc.TLS, p uintptr, stmtChunks uintptr, vector_column uintptr, vectorColumnIdx int32, arrayRowidsIn uintptr, aMetadataIn uintptr, idxStr uintptr, argc int32, argv uintptr, queryVector uintptr, k Ti64, out_topk_rowids unsafe.Pointer, out_topk_distances unsafe.Pointer, out_used unsafe.Pointer) (r int32) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var b, bTaken, baseVectors, base_i, base_i1, base_i2, bmMetadata, bmRowids, chunkRowids, chunkValidity, chunk_distances, chunk_topk_idxs, in, tmp_topk_distances, tmp_topk_rowids, topk_distances, topk_rowids, v1 uintptr
	var baseVectorsSize, chunk_id, currentBaseVectorsSize, expectedBaseVectorsSize, k_used, rowidsSize, used, validitySize Ti64
	var hasDistanceConstraints, hasMetadataFilters, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, idx, idx1, idx2, idxStrLength, metadata_idx, numValueEntries, operator, rc, used1, v4 int32
	var kind, kind1, kind2 int8
	var op Tvec0_distance_constraint_operator
	var result, target Tf32
	var v12, v13, v14 int64
	var _ /* blobVectors at bp+0 */ uintptr
	var _ /* metadataBlobs at bp+8 */ [16]uintptr
	var _ /* rowid at bp+136 */ Ti64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b, bTaken, baseVectors, baseVectorsSize, base_i, base_i1, base_i2, bmMetadata, bmRowids, chunkRowids, chunkValidity, chunk_distances, chunk_id, chunk_topk_idxs, currentBaseVectorsSize, expectedBaseVectorsSize, hasDistanceConstraints, hasMetadataFilters, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, idx, idx1, idx2, idxStrLength, in, k_used, kind, kind1, kind2, metadata_idx, numValueEntries, op, operator, rc, result, rowidsSize, target, tmp_topk_distances, tmp_topk_rowids, topk_distances, topk_rowids, used, used1, validitySize, v1, v12, v13, v14, v4
	// for each chunk, get top min(k, chunk_size) rowid + distances to query vec.
	// then reconcile all topk_chunks for a true top k.
	// output only rowids + distances for now
	rc = m_SQLITE_OK
	*(*uintptr)(__ccgo_ptr(bp)) = libc.UintptrFromInt32(0)
	baseVectors = libc.UintptrFromInt32(0) // memory: chunk_size * dimensions * element_size
	// OWNED BY CALLER ON SUCCESS
	topk_rowids = libc.UintptrFromInt32(0) // memory: k * 4
	// OWNED BY CALLER ON SUCCESS
	topk_distances = libc.UintptrFromInt32(0)     // memory: k * 4
	tmp_topk_rowids = libc.UintptrFromInt32(0)    // memory: k * 4
	tmp_topk_distances = libc.UintptrFromInt32(0) // memory: k * 4
	chunk_distances = libc.UintptrFromInt32(0)    // memory: chunk_size * 4
	b = libc.UintptrFromInt32(0)                  // memory: chunk_size / 8
	bTaken = libc.UintptrFromInt32(0)             // memory: chunk_size / 8
	chunk_topk_idxs = libc.UintptrFromInt32(0)    // memory: k * 4
	bmRowids = libc.UintptrFromInt32(0)           // memory: chunk_size / 8
	bmMetadata = libc.UintptrFromInt32(0)         // memory: chunk_size / 8
	//                        // total: a lot???
	// 6 * (k * 4) + (k * 2) + (chunk_size / 8) + (chunk_size * dimensions * 4)
	topk_rowids = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(8)))
	if !(topk_rowids != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, topk_rowids, 0, uint64(libc.Uint64FromInt64(k)*uint64(8)))
	topk_distances = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(4)))
	if !(topk_distances != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, topk_distances, 0, uint64(libc.Uint64FromInt64(k)*uint64(4)))
	tmp_topk_rowids = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(8)))
	if !(tmp_topk_rowids != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, tmp_topk_rowids, 0, uint64(libc.Uint64FromInt64(k)*uint64(8)))
	tmp_topk_distances = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(4)))
	if !(tmp_topk_distances != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, tmp_topk_distances, 0, uint64(libc.Uint64FromInt64(k)*uint64(4)))
	k_used = 0
	baseVectorsSize = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, *(*TVectorColumnDefinition)(__ccgo_ptr(vector_column))))
	baseVectors = libsqlite3.Xsqlite3_malloc(tls, int32(baseVectorsSize))
	if !(baseVectors != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	chunk_distances = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(4)))
	if !(chunk_distances != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	b = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	if !(b != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	bTaken = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	if !(bTaken != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	chunk_topk_idxs = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(4)))
	if !(chunk_topk_idxs != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	if arrayRowidsIn != 0 {
		v1 = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	bmRowids = v1
	if arrayRowidsIn != 0 && !(bmRowids != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, bp+8, 0, libc.Uint64FromInt64(8)*libc.Uint64FromInt32(m_VEC0_MAX_METADATA_COLUMNS))
	bmMetadata = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	if !(bmMetadata != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	idxStrLength = libc.Int32FromUint64(libc.Xstrlen(tls, idxStr))
	numValueEntries = (idxStrLength - int32(1)) / int32(4)
	hasMetadataFilters = 0
	hasDistanceConstraints = 0
	i = 0
	for {
		if !(i < argc) {
			break
		}
		idx = int32(1) + i*int32(4)
		kind = *(*int8)(__ccgo_ptr(idxStr + uintptr(idx+0)))
		if int32(kind) == int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) {
			hasMetadataFilters = int32(1)
		} else {
			if int32(kind) == int32(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT) {
				hasDistanceConstraints = int32(1)
			}
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	for int32(m_true) != 0 {
		rc = libsqlite3.Xsqlite3_step(tls, stmtChunks)
		if rc == int32(m_SQLITE_DONE) {
			break
		}
		if rc != int32(m_SQLITE_ROW) {
			Xvtab_set_error(tls, p, __ccgo_ts+9697, 0)
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		libc.Xmemset(tls, chunk_distances, 0, libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(4))
		libc.Xmemset(tls, chunk_topk_idxs, 0, uint64(libc.Uint64FromInt64(k)*uint64(4)))
		Xbitmap_clear(tls, b, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		chunk_id = libsqlite3.Xsqlite3_column_int64(tls, stmtChunks, 0)
		chunkValidity = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(1))
		validitySize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(1)))
		if validitySize != int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT__)) {
			// IMP: V05271_22109
			Xvtab_set_error(tls, p, __ccgo_ts+9715, libc.VaList(bp+152, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT__), validitySize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		chunkRowids = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(2))
		rowidsSize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(2)))
		if libc.Uint64FromInt64(rowidsSize) != uint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8)) {
			// IMP: V02796_19635
			Xvtab_set_error(tls, p, __ccgo_ts+9777, 0)
			Xvtab_set_error(tls, p, __ccgo_ts+9803, libc.VaList(bp+152, libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8), rowidsSize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		// open the vector chunk blob for the current chunk
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, *(*uintptr)(__ccgo_ptr(p + 352 + uintptr(vectorColumnIdx)*8)), __ccgo_ts+3712, chunk_id, 0, bp)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+9863, libc.VaList(bp+152, chunk_id))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		currentBaseVectorsSize = int64(libsqlite3.Xsqlite3_blob_bytes(tls, *(*uintptr)(__ccgo_ptr(bp))))
		expectedBaseVectorsSize = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, *(*TVectorColumnDefinition)(__ccgo_ptr(vector_column))))
		if currentBaseVectorsSize != expectedBaseVectorsSize {
			// IMP: V16465_00535
			Xvtab_set_error(tls, p, __ccgo_ts+9906, libc.VaList(bp+152, expectedBaseVectorsSize, currentBaseVectorsSize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		rc = libsqlite3.Xsqlite3_blob_read(tls, *(*uintptr)(__ccgo_ptr(bp)), baseVectors, int32(currentBaseVectorsSize), 0)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+9966, libc.VaList(bp+152, chunk_id))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		Xbitmap_copy(tls, b, chunkValidity, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		if arrayRowidsIn != 0 {
			Xbitmap_clear(tls, bmRowids, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
			i1 = 0
			for {
				if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
					break
				}
				if !(Xbitmap_get(tls, chunkValidity, i1) != 0) {
					goto _3
				}
				*(*Ti64)(__ccgo_ptr(bp + 136)) = *(*Ti64)(__ccgo_ptr(chunkRowids + uintptr(i1)*8))
				in = libc.Xbsearch(tls, bp+136, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Fz, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Flength, uint64(8), __ccgo_fp(X_cmp))
				if in != 0 {
					v4 = int32(1)
				} else {
					v4 = 0
				}
				Xbitmap_set(tls, bmRowids, i1, v4)
				goto _3
			_3:
				;
				i1 = i1 + 1
			}
			Xbitmap_and_inplace(tls, b, bmRowids, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		}
		if hasMetadataFilters != 0 {
			i2 = 0
			for {
				if !(i2 < argc) {
					break
				}
				idx1 = int32(1) + i2*int32(4)
				kind1 = *(*int8)(__ccgo_ptr(idxStr + uintptr(idx1+0)))
				if int32(kind1) != int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) {
					goto _5
				}
				metadata_idx = int32(*(*int8)(__ccgo_ptr(idxStr + uintptr(idx1+int32(1))))) - int32('A')
				operator = int32(*(*int8)(__ccgo_ptr(idxStr + uintptr(idx1+int32(2)))))
				if !((*(*[16]uintptr)(__ccgo_ptr(bp + 8)))[metadata_idx] != 0) {
					rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, *(*uintptr)(__ccgo_ptr(p + 480 + uintptr(metadata_idx)*8)), __ccgo_ts+4053, chunk_id, 0, bp+8+uintptr(metadata_idx)*8)
					Xvtab_set_error(tls, p, __ccgo_ts+9999, 0)
					if rc != m_SQLITE_OK {
						goto cleanup
					}
				}
				Xbitmap_clear(tls, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
				rc = Xvec0_set_metadata_filter_bitmap(tls, p, metadata_idx, operator, *(*uintptr)(__ccgo_ptr(argv + uintptr(i2)*8)), (*(*[16]uintptr)(__ccgo_ptr(bp + 8)))[metadata_idx], chunk_id, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size, aMetadataIn, i2)
				if rc != m_SQLITE_OK {
					Xvtab_set_error(tls, p, __ccgo_ts+10028, 0)
					if rc != m_SQLITE_OK {
						goto cleanup
					}
				}
				Xbitmap_and_inplace(tls, b, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
				goto _5
			_5:
				;
				i2 = i2 + 1
			}
		}
		i3 = 0
		for {
			if !(i3 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
				break
			}
			if !(Xbitmap_get(tls, b, i3) != 0) {
				goto _6
			}
			switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Felement_type {
			case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
				base_i = baseVectors + uintptr(libc.Uint64FromInt32(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)*4
				switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdistance_metric {
				case int32(_VEC0_DISTANCE_METRIC_L2):
					result = _distance_l2_sqr_float(tls, base_i, queryVector, unsafe.Pointer(vector_column+16))
				case int32(_VEC0_DISTANCE_METRIC_L1):
					result = float32(_distance_l1_f32(tls, base_i, queryVector, unsafe.Pointer(vector_column+16)))
				case int32(_VEC0_DISTANCE_METRIC_COSINE):
					result = _distance_cosine_float(tls, base_i, queryVector, unsafe.Pointer(vector_column+16))
					break
				}
			case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
				base_i1 = baseVectors + uintptr(libc.Uint64FromInt32(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)
				switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdistance_metric {
				case int32(_VEC0_DISTANCE_METRIC_L2):
					result = _distance_l2_sqr_int8(tls, base_i1, queryVector, unsafe.Pointer(vector_column+16))
				case int32(_VEC0_DISTANCE_METRIC_L1):
					result = float32(_distance_l1_int8(tls, base_i1, queryVector, unsafe.Pointer(vector_column+16)))
				case int32(_VEC0_DISTANCE_METRIC_COSINE):
					result = _distance_cosine_int8(tls, base_i1, queryVector, unsafe.Pointer(vector_column+16))
					break
				}
			case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
				base_i2 = baseVectors + uintptr(libc.Uint64FromInt32(i3)*((*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions/libc.Uint64FromInt32(m___CHAR_BIT__)))
				result = _distance_hamming(tls, base_i2, queryVector, unsafe.Pointer(vector_column+16))
				break
			}
			*(*Tf32)(__ccgo_ptr(chunk_distances + uintptr(i3)*4)) = result
			goto _6
		_6:
			;
			i3 = i3 + 1
		}
		if hasDistanceConstraints != 0 {
			i4 = 0
			for {
				if !(i4 < argc) {
					break
				}
				idx2 = int32(1) + i4*int32(4)
				kind2 = *(*int8)(__ccgo_ptr(idxStr + uintptr(idx2+0)))
				// TODO casts f64 to f32, is that a problem?
				target = float32(libsqlite3.Xsqlite3_value_double(tls, *(*uintptr)(__ccgo_ptr(argv + uintptr(i4)*8))))
				if int32(kind2) != int32(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT) {
					goto _7
				}
				op = int32(*(*int8)(__ccgo_ptr(idxStr + uintptr(idx2+int32(1)))))
				switch op {
				case int32(_VEC0_DISTANCE_CONSTRAINT_GE):
					i5 = 0
					for {
						if !(i5 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						if Xbitmap_get(tls, b, i5) != 0 && !(*(*Tf32)(__ccgo_ptr(chunk_distances + uintptr(i5)*4)) >= target) {
							Xbitmap_set(tls, b, i5, 0)
						}
						goto _8
					_8:
						;
						i5 = i5 + 1
					}
				case int32(_VEC0_DISTANCE_CONSTRAINT_GT):
					i6 = 0
					for {
						if !(i6 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						if Xbitmap_get(tls, b, i6) != 0 && !(*(*Tf32)(__ccgo_ptr(chunk_distances + uintptr(i6)*4)) > target) {
							Xbitmap_set(tls, b, i6, 0)
						}
						goto _9
					_9:
						;
						i6 = i6 + 1
					}
				case int32(_VEC0_DISTANCE_CONSTRAINT_LE):
					i7 = 0
					for {
						if !(i7 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						if Xbitmap_get(tls, b, i7) != 0 && !(*(*Tf32)(__ccgo_ptr(chunk_distances + uintptr(i7)*4)) <= target) {
							Xbitmap_set(tls, b, i7, 0)
						}
						goto _10
					_10:
						;
						i7 = i7 + 1
					}
				case int32(_VEC0_DISTANCE_CONSTRAINT_LT):
					i8 = 0
					for {
						if !(i8 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						if Xbitmap_get(tls, b, i8) != 0 && !(*(*Tf32)(__ccgo_ptr(chunk_distances + uintptr(i8)*4)) < target) {
							Xbitmap_set(tls, b, i8, 0)
						}
						goto _11
					_11:
						;
						i8 = i8 + 1
					}
					break
				}
				goto _7
			_7:
				;
				i4 = i4 + 1
			}
		}
		if k <= int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
			v12 = k
		} else {
			v12 = int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		}
		_min_idx__ccgo_gp(tls, chunk_distances, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size, b, chunk_topk_idxs, int32(v12), bTaken, unsafe.Pointer(&used1))
		if k <= int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
			v13 = k
		} else {
			v13 = int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		}
		if v13 <= int64(used1) {
			if k <= int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
				v14 = k
			} else {
				v14 = int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
			}
			v12 = v14
		} else {
			v12 = int64(used1)
		}
		_merge_sorted_lists__ccgo_gp(tls, topk_distances, topk_rowids, k_used, chunk_distances, chunkRowids, chunk_topk_idxs, v12, tmp_topk_distances, tmp_topk_rowids, k, unsafe.Pointer(&used))
		i9 = 0
		for {
			if !(int64(i9) < used) {
				break
			}
			*(*Ti64)(__ccgo_ptr(topk_rowids + uintptr(i9)*8)) = *(*Ti64)(__ccgo_ptr(tmp_topk_rowids + uintptr(i9)*8))
			*(*Tf32)(__ccgo_ptr(topk_distances + uintptr(i9)*4)) = *(*Tf32)(__ccgo_ptr(tmp_topk_distances + uintptr(i9)*4))
			goto _16
		_16:
			;
			i9 = i9 + 1
		}
		k_used = used
		// blobVectors is always opened with read-only permissions, so this never
		// fails.
		libsqlite3.Xsqlite3_blob_close(tls, *(*uintptr)(__ccgo_ptr(bp)))
		*(*uintptr)(__ccgo_ptr(bp)) = libc.UintptrFromInt32(0)
	}
	*(*uintptr)(out_topk_rowids) = topk_rowids
	*(*uintptr)(out_topk_distances) = topk_distances
	*(*Ti64)(out_used) = k_used
	rc = m_SQLITE_OK
	goto cleanup
cleanup:
	;
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_free(tls, topk_rowids)
		libsqlite3.Xsqlite3_free(tls, topk_distances)
	}
	libsqlite3.Xsqlite3_free(tls, chunk_topk_idxs)
	libsqlite3.Xsqlite3_free(tls, tmp_topk_rowids)
	libsqlite3.Xsqlite3_free(tls, tmp_topk_distances)
	libsqlite3.Xsqlite3_free(tls, b)
	libsqlite3.Xsqlite3_free(tls, bTaken)
	libsqlite3.Xsqlite3_free(tls, bmRowids)
	libsqlite3.Xsqlite3_free(tls, baseVectors)
	libsqlite3.Xsqlite3_free(tls, chunk_distances)
	libsqlite3.Xsqlite3_free(tls, bmMetadata)
	i10 = 0
	for {
		if !(i10 < int32(m_VEC0_MAX_METADATA_COLUMNS)) {
			break
		}
		libsqlite3.Xsqlite3_blob_close(tls, (*(*[16]uintptr)(__ccgo_ptr(bp + 8)))[i10])
		goto _17
	_17:
		;
		i10 = i10 + 1
	}
	// blobVectors is always opened with read-only permissions, so this never
	// fails.
	libsqlite3.Xsqlite3_blob_close(tls, *(*uintptr)(__ccgo_ptr(bp)))
	return rc
}

func _vec_quantize_binary(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	var dimensions, i, i1 Tsize_t
	var elementType _VectorElementType
	var out, pzError, vector, v2 uintptr
	var rc, res, res1, sz int32
	var vectorCleanup Tvector_cleanup
	_, _, _, _, _, _, _, _, _, _, _, _, _ = dimensions, elementType, i, i1, out, pzError, rc, res, res1, sz, vector, vectorCleanup, v2
	rc = _vector_from_value__ccgo_gp(tls, *(*uintptr)(__ccgo_ptr(argv)), unsafe.Pointer(&vector), unsafe.Pointer(&dimensions), unsafe.Pointer(&elementType), unsafe.Pointer(&vectorCleanup), unsafe.Pointer(&pzError))
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, pzError, -int32(1))
		libsqlite3.Xsqlite3_free(tls, pzError)
		return
	}
	if dimensions <= uint64(0) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+974, -int32(1))
		goto cleanup
		return
	}
	if dimensions%uint64(m___CHAR_BIT__) != uint64(0) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1013, -int32(1))
		goto cleanup
		return
	}
	sz = libc.Int32FromUint64(dimensions / uint64(m___CHAR_BIT__))
	out = libsqlite3.Xsqlite3_malloc(tls, sz)
	if !(out != 0) {
		libsqlite3.Xsqlite3_result_error_code(tls, context, int32(m_SQLITE_NOMEM))
		goto cleanup
		return
	}
	libc.Xmemset(tls, out, 0, libc.Uint64FromInt32(sz))
	switch elementType {
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		i = uint64(0)
		for {
			if !(i < dimensions) {
				break
			}
			res = libc.BoolInt32(float64(*(*Tf32)(__ccgo_ptr(vector + uintptr(i)*4))) > float64(0))
			v2 = out + uintptr(i/uint64(8))
			*(*Tu8)(unsafe.Pointer(v2)) = Tu8(int32(*(*Tu8)(unsafe.Pointer(v2))) | res<<(i%libc.Uint64FromInt32(8)))
			goto _1
		_1:
			;
			i = i + 1
		}
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		i1 = uint64(0)
		for {
			if !(i1 < dimensions) {
				break
			}
			res1 = libc.BoolInt32(int32(*(*Ti8)(__ccgo_ptr(vector + uintptr(i1)))) > 0)
			v2 = out + uintptr(i1/uint64(8))
			*(*Tu8)(unsafe.Pointer(v2)) = Tu8(int32(*(*Tu8)(unsafe.Pointer(v2))) | res1<<(i1%libc.Uint64FromInt32(8)))
			goto _3
		_3:
			;
			i1 = i1 + 1
		}
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1079, -int32(1))
		libsqlite3.Xsqlite3_free(tls, out)
		return
	}
	libsqlite3.Xsqlite3_result_blob(tls, context, out, sz, __ccgo_fp(libsqlite3.Xsqlite3_free))
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_BIT))
	goto cleanup
cleanup:
	;
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{vectorCleanup})))(tls, vector)
}

func _vec_slice(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	var cleanup Tvector_cleanup
	var dimensions, i, i1, i2, n Tsize_t
	var elementType _VectorElementType
	var end, outSize, outSize1, outSize2, rc, start int32
	var err, out, out1, out2, vector uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cleanup, dimensions, elementType, end, err, i, i1, i2, n, out, out1, out2, outSize, outSize1, outSize2, rc, start, vector
	rc = _vector_from_value__ccgo_gp(tls, *(*uintptr)(__ccgo_ptr(argv)), unsafe.Pointer(&vector), unsafe.Pointer(&dimensions), unsafe.Pointer(&elementType), unsafe.Pointer(&cleanup), unsafe.Pointer(&err))
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, err, -int32(1))
		libsqlite3.Xsqlite3_free(tls, err)
		return
	}
	start = libsqlite3.Xsqlite3_value_int(tls, *(*uintptr)(__ccgo_ptr(argv + 1*8)))
	end = libsqlite3.Xsqlite3_value_int(tls, *(*uintptr)(__ccgo_ptr(argv + 2*8)))
	if start < 0 {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1260, -int32(1))
		goto done
	}
	if end < 0 {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1306, -int32(1))
		goto done
	}
	if libc.Uint64FromInt32(start) > dimensions {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1350, -int32(1))
		goto done
	}
	if libc.Uint64FromInt32(end) > dimensions {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1411, -int32(1))
		goto done
	}
	if start > end {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1470, -int32(1))
		goto done
	}
	if start == end {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1518, -int32(1))
		goto done
	}
	n = libc.Uint64FromInt32(end - start)
	switch elementType {
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		outSize = libc.Int32FromUint64(n * uint64(4))
		out = libsqlite3.Xsqlite3_malloc(tls, outSize)
		if !(out != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			goto done
		}
		libc.Xmemset(tls, out, 0, libc.Uint64FromInt32(outSize))
		i = uint64(0)
		for {
			if !(i < n) {
				break
			}
			*(*Tf32)(__ccgo_ptr(out + uintptr(i)*4)) = *(*Tf32)(__ccgo_ptr(vector + uintptr(libc.Uint64FromInt32(start)+i)*4))
			goto _1
		_1:
			;
			i = i + 1
		}
		libsqlite3.Xsqlite3_result_blob(tls, context, out, outSize, __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
		goto done
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		outSize1 = libc.Int32FromUint64(n * uint64(1))
		out1 = libsqlite3.Xsqlite3_malloc(tls, outSize1)
		if !(out1 != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			return
		}
		libc.Xmemset(tls, out1, 0, libc.Uint64FromInt32(outSize1))
		i1 = uint64(0)
		for {
			if !(i1 < n) {
				break
			}
			*(*Ti8)(__ccgo_ptr(out1 + uintptr(i1))) = *(*Ti8)(__ccgo_ptr(vector + uintptr(libc.Uint64FromInt32(start)+i1)))
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		libsqlite3.Xsqlite3_result_blob(tls, context, out1, outSize1, __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
		goto done
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		if start%int32(m___CHAR_BIT__) != 0 {
			libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1601, -int32(1))
			goto done
		}
		if end%int32(m___CHAR_BIT__) != 0 {
			libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1637, -int32(1))
			goto done
		}
		outSize2 = libc.Int32FromUint64(n / uint64(m___CHAR_BIT__))
		out2 = libsqlite3.Xsqlite3_malloc(tls, outSize2)
		if !(out2 != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			return
		}
		libc.Xmemset(tls, out2, 0, libc.Uint64FromInt32(outSize2))
		i2 = uint64(0)
		for {
			if !(i2 < n/uint64(m___CHAR_BIT__)) {
				break
			}
			*(*Tu8)(__ccgo_ptr(out2 + uintptr(i2))) = *(*Tu8)(__ccgo_ptr(vector + uintptr(libc.Uint64FromInt32(start/libc.Int32FromInt32(m___CHAR_BIT__))+i2)))
			goto _3
		_3:
			;
			i2 = i2 + 1
		}
		libsqlite3.Xsqlite3_result_blob(tls, context, out2, outSize2, __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_BIT))
		goto done
	}
	goto done
done:
	;
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{cleanup})))(tls, vector)
}

func _vec_to_json(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var b Tu8
	var cleanup Tvector_cleanup
	var dimensions, i Tsize_t
	var elementType _VectorElementType
	var err, s, str, vector uintptr
	var len1, rc int32
	var value Tf32
	_, _, _, _, _, _, _, _, _, _, _, _ = b, cleanup, dimensions, elementType, err, i, len1, rc, s, str, value, vector
	rc = _vector_from_value__ccgo_gp(tls, *(*uintptr)(__ccgo_ptr(argv)), unsafe.Pointer(&vector), unsafe.Pointer(&dimensions), unsafe.Pointer(&elementType), unsafe.Pointer(&cleanup), unsafe.Pointer(&err))
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, err, -int32(1))
		libsqlite3.Xsqlite3_free(tls, err)
		return
	}
	str = libsqlite3.Xsqlite3_str_new(tls, libsqlite3.Xsqlite3_context_db_handle(tls, context))
	libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1671)
	i = uint64(0)
	for {
		if !(i < dimensions) {
			break
		}
		if i != uint64(0) {
			libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1673)
		}
		if elementType == int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32) {
			value = *(*Tf32)(__ccgo_ptr(vector + uintptr(i)*4))
			if libc.X__builtin_isnan(tls, float64(value)) != 0 {
				libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1675)
			} else {
				libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1680, libc.VaList(bp+8, float64(value)))
			}
		} else {
			if elementType == int32(_SQLITE_VEC_ELEMENT_TYPE_INT8) {
				libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1683, libc.VaList(bp+8, int32(*(*Ti8)(__ccgo_ptr(vector + uintptr(i))))))
			} else {
				if elementType == int32(_SQLITE_VEC_ELEMENT_TYPE_BIT) {
					b = libc.Uint8FromInt32(libc.Int32FromUint8(*(*Tu8)(__ccgo_ptr(vector + uintptr(i/uint64(8))))) >> (i % uint64(m___CHAR_BIT__)) & int32(1))
					libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1683, libc.VaList(bp+8, libc.Int32FromUint8(b)))
				}
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1686)
	len1 = libsqlite3.Xsqlite3_str_length(tls, str)
	s = libsqlite3.Xsqlite3_str_finish(tls, str)
	if s != 0 {
		libsqlite3.Xsqlite3_result_text(tls, context, s, len1, __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(m_JSON_SUBTYPE))
	} else {
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
	}
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{cleanup})))(tls, vector)
}

type _version = int32

const _whitespace_form = 1

type decimal_mode = Tdecimal_mode

type decimal_record = Tdecimal_record

type decimal_string = Tdecimal_string

type extended = Textended

type fp_exception_field_type = Tfp_exception_field_type

type fpos64_t = Tfpos64_t

/*
** 2001-09-15
**
** The author disclaims copyright to this source code.  In place of
** a legal notice, here is a blessing:
**
**    May you do good and not evil.
**    May you find forgiveness for yourself and forgive others.
**    May you share freely, never taking more than you give.
**
*************************************************************************
** This header file defines the interface that the SQLite library
** presents to client programs.  If a C-function, structure, datatype,
** or constant definition does not appear in this file, then it is
** not a published API of SQLite, is subject to change without
** notice, and should not be referenced by programs that use SQLite.
**
** Some of the definitions that are in this file are marked as
** "experimental".  Experimental interfaces are normally new
** features recently added to SQLite.  We do not anticipate changes
** to experimental interfaces but reserve the right to make minor changes
** if experience from use "in the wild" suggest such changes are prudent.
**
** The official C-language API documentation for SQLite is derived
** from comments in this file.  This file is the authoritative source
** on how SQLite interfaces are supposed to operate.
**
** The name of this file under configuration management is "sqlite.h.in".
** The makefile makes some minor changes to this file (such as inserting
** the version number) and changes its name to "sqlite3.h" as
** part of the build process.
 */

const m_ARG_MAX = "_ARG_MAX64"

const m_BC_BASE_MAX = "_POSIX2_BC_BASE_MAX"

const m_BC_DIM_MAX = "_POSIX2_BC_DIM_MAX"

const m_BC_SCALE_MAX = "_POSIX2_BC_SCALE_MAX"

const m_BC_STRING_MAX = "_POSIX2_BC_STRING_MAX"

const m_CHARCLASS_NAME_MAX = "_POSIX2_CHARCLASS_NAME_MAX"

const m_CHILD_MAX = 25

const m_DECIMAL_STRING_LENGTH = 512

const m_EADDRINUSE = 125

const m_EADDRNOTAVAIL = 126

const m_EAFNOSUPPORT = 124

const m_EALREADY = 149

const m_EBADE = 50

const m_EBADFD = 81

const m_EBADMSG = 77

const m_EBADR = 51

const m_EBADRQC = 54

const m_EBADSLT = 55

const m_EBFONT = 57

const m_ECANCELED = 47

const m_ECHRNG = 37

const m_ECONNABORTED = 130

const m_ECONNREFUSED = 146

const m_ECONNRESET = 131

const m_EDEADLK = 45

const m_EDEADLOCK = 56

const m_EDESTADDRREQ = 96

const m_EDQUOT = 49

const m_EHOSTDOWN = 147

const m_EHOSTUNREACH = 148

const m_EIDRM = 36

const m_EILSEQ = 88

const m_EINPROGRESS = 150

const m_EISCONN = 133

const m_EL2HLT = 44

const m_EL2NSYNC = 38

const m_EL3HLT = 39

const m_EL3RST = 40

const m_ELIBACC = 83

const m_ELIBBAD = 84

const m_ELIBEXEC = 87

const m_ELIBMAX = 86

const m_ELIBSCN = 85

const m_ELNRNG = 41

const m_ELOCKUNMAPPED = 72

const m_ELOOP = 90

const m_EMSGSIZE = 97

const m_EMULTIHOP = 74

const m_ENAMETOOLONG = 78

const m_ENETDOWN = 127

const m_ENETRESET = 129

const m_ENETUNREACH = 128

const m_ENOANO = 53

const m_ENOBUFS = 132

const m_ENOCSI = 43

const m_ENOLCK = 46

const m_ENOMSG = 35

const m_ENOPROTOOPT = 99

const m_ENOSYS = 89

const m_ENOTACTIVE = 73

const m_ENOTCONN = 134

const m_ENOTEMPTY = 93

const m_ENOTRECOVERABLE = 59

const m_ENOTSOCK = 95

const m_ENOTSUP = 48

const m_ENOTUNIQ = 80

const m_EOPNOTSUPP = 122

const m_EOVERFLOW = 79

const m_EOWNERDEAD = 58

const m_EPFNOSUPPORT = 123

const m_EPROTONOSUPPORT = 120

const m_EPROTOTYPE = 98

const m_EREMCHG = 82

const m_ERESTART = 91

const m_ESHUTDOWN = 143

const m_ESOCKTNOSUPPORT = 121

const m_ESTALE = 151

const m_ESTRPIPE = 92

const m_ETIMEDOUT = 145

const m_ETOOMANYREFS = 144

const m_EUNATCH = 42

const m_EUSERS = 94

const m_EXFULL = 52

const m_EXPR_NEST_MAX = "_POSIX2_EXPR_NEST_MAX"

const m_FCHR_MAX = 1048576

const m_FOPEN_MAX = "_NFILE"

const m_LINE_MAX = "_POSIX2_LINE_MAX"

const m_LOGNAME_MAX = 32

const m_LOGNAME_MAX_TRAD = 8

const m_L_ctermid = 9

const m_L_tmpnam = 25

const m_MAX_CANON = 256

const m_MAX_INPUT = 512

const m_N_IEEE_EXCEPTION = 6

const m_OPEN_MAX = 256

const m_PASS_MAX = "_PASS_MAX"

const m_PID_MAX = 999999

const m_PIPE_BUF = 5120

const m_PIPE_MAX = 5120

const m_PRIXFAST16 = "PRIX32"

const m_PRIdFAST16 = "PRId32"

const m_PRIiFAST16 = "PRIi32"

const m_PRIoFAST16 = "PRIo32"

const m_PRIuFAST16 = "PRIu32"

const m_PRIxFAST16 = "PRIx32"

const m_PTRDIFF_MAX = 9223372036854775807

const m_RAND_MAX = 32767

const m_RE_DUP_MAX = "_POSIX2_RE_DUP_MAX"

const m_SCNX16 = "hX"

const m_SCNX32 = "X"

const m_SCNX64 = "PRIX64"

const m_SCNX8 = "hhX"

const m_SCNXFAST16 = "SCNX32"

const m_SCNXFAST32 = "SCNX32"

const m_SCNXFAST64 = "PRIX64"

const m_SCNXFAST8 = "SCNX8"

const m_SCNXLEAST16 = "SCNX16"

const m_SCNXLEAST32 = "SCNX32"

const m_SCNXLEAST64 = "PRIX64"

const m_SCNXLEAST8 = "SCNX8"

const m_SCNXMAX = "PRIXMAX"

const m_SCNXPTR = "PRIXPTR"

const m_SCNdFAST16 = "SCNd32"

const m_SCNdMAX = "PRIdMAX"

const m_SCNdPTR = "PRIdPTR"

const m_SCNiFAST16 = "SCNi32"

const m_SCNiMAX = "PRIiMAX"

const m_SCNiPTR = "PRIiPTR"

const m_SCNoFAST16 = "SCNo32"

const m_SCNoMAX = "PRIoMAX"

const m_SCNoPTR = "PRIoPTR"

const m_SCNuFAST16 = "SCNu32"

const m_SCNuMAX = "PRIuMAX"

const m_SCNuPTR = "PRIuPTR"

const m_SCNxFAST16 = "SCNx32"

const m_SCNxMAX = "PRIxMAX"

const m_SCNxPTR = "PRIxPTR"

const m_SIG_ATOMIC_MAX = 2147483647

const m_SIZE_MAX = 18446744073709551615

const m_STD_BLK = 1024

const m_SYMLINK_MAX = 1024

const m_SYSPID_MAX = 1

const m_SYS_NMLN = 257

const m_TMP_MAX = 17576

const m_TTYNAME_MAX = 128

const m_UINT64_MAX = 18446744073709551615

const m_USI_MAX = 4294967295

const m_WCHAR_MAX = 2147483647

const m_WINT_MAX = 2147483647

const m__ALIGNMENT_REQUIRED = 1

const m__ARG_MAX32 = 1048320

const m__ARG_MAX64 = 2096640

const m__BOOL_ALIGNMENT = 1

const m__CACHE_LINE_SHIFT = 6

const m__CHAR_ALIGNMENT = 1

const m__DOUBLE_ALIGNMENT = 8

const m__DOUBLE_COMPLEX_ALIGNMENT = 8

const m__DTRACE_VERSION = 1

const m__FLOAT_ALIGNMENT = 4

const m__FLOAT_COMPLEX_ALIGNMENT = 4

const m__INT_ALIGNMENT = 4

const m__IOEOF = 0020

const m__IOERR = 0040

const m__IOFBF = 0000

const m__IOLBF = 0100

const m__IOMYBUF = 0010

const m__IONBF = 0004

const m__IOREAD = 0001

const m__IORW = 0200

const m__IOWRT = 0002

const m__LONG_ALIGNMENT = 8

const m__LONG_DOUBLE_ALIGNMENT = 16

const m__LONG_DOUBLE_COMPLEX_ALIGNMENT = 16

const m__LONG_LONG_ALIGNMENT = 8

const m__LONG_LONG_ALIGNMENT_32 = 4

const m__MAX_ALIGNMENT = 16

const m__MODF16 = "h"

const m__MODF8 = "hh"

const m__NFILE = 20

const m__NORETURN_KYWD = "_Noreturn"

const m__PASS_MAX = 256

const m__PASS_MAX_XPG = 8

const m__POINTER_ALIGNMENT = 8

const m__POSIX_CHILD_MAX = 6

const m__POSIX_NGROUPS_MAX = 0

const m__POSIX_OPEN_MAX = 16

const m__POSIX_PATH_MAX = 255

const m__POSIX_TZNAME_MAX = 3

const m__PRIX = "X"

const m__PRIXID = "X"

const m__PRIXWC = "_PRIXID"

const m__PRId = "d"

const m__PRIdID = "d"

const m__PRIdWC = "_PRIdID"

const m__PRIi = "i"

const m__PRIiID = "i"

const m__PRIiWC = "_PRIiID"

const m__PRIo = "o"

const m__PRIoID = "o"

const m__PRIoWC = "_PRIoID"

const m__PRIu = "u"

const m__PRIuID = "u"

const m__PRIuWC = "_PRIuID"

const m__PRIx = "x"

const m__PRIxID = "x"

const m__PRIxWC = "_PRIxID"

const m__REENTRANT = 1

const m__RESTRICT_KYWD = "restrict"

const m__SBFSIZ = 8

const m__SCNXID = "_PRIXID"

const m__SCNXWC = "_PRIXID"

const m__SCNdID = "_PRIdID"

const m__SCNdWC = "_PRIdID"

const m__SCNiID = "_PRIiID"

const m__SCNiWC = "_PRIiID"

const m__SCNoID = "_PRIoID"

const m__SCNoWC = "_PRIoID"

const m__SCNuID = "_PRIuID"

const m__SCNuWC = "_PRIuID"

const m__SCNxID = "_PRIxID"

const m__SCNxWC = "_PRIxID"

const m__SHORT_ALIGNMENT = 2

const m__XOPEN_VERSION = 3

const m___BFLT16_DECIMAL_DIG__ = 4

const m___BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"

const m___BFLT16_DIG__ = 2

const m___BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"

const m___BFLT16_HAS_DENORM__ = 1

const m___BFLT16_HAS_INFINITY__ = 1

const m___BFLT16_HAS_QUIET_NAN__ = 1

const m___BFLT16_IS_IEC_60559__ = 0

const m___BFLT16_MANT_DIG__ = 8

const m___BFLT16_MAX_10_EXP__ = 38

const m___BFLT16_MAX_EXP__ = 128

const m___BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"

const m___BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"

const m___BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"

const m___EXT1_VISIBLE = 0

const m___FILE_TAG = "__FILE"

const m___GNUC_MINOR__ = 3

const m___INT8_TYPE__ = "char"

const m___INT_FAST8_TYPE__ = "char"

const m___INT_LEAST8_TYPE__ = "char"

const m___SVR4 = 1

const m___VERSION__ = "14.3.0"

const m___illumos__ = 1

const m___sun = 1

const m___sun__ = 1

const m___svr4__ = 1

const m___va_alist_type = "long"

const m___va_ptr_base = "void"

const m_sun = 1

type quadruple = Tquadruple

type sigfpe_code_type = Tsigfpe_code_type

type sigfpe_handler_type = Tsigfpe_handler_type

type single = Tsingle

type t__FILE = struct {
	F__pad [16]int64
}
