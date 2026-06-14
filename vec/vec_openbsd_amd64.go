// Code generated for openbsd/amd64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build openbsd && amd64

package vec

import (
	"modernc.org/libc"
	libsqlite3 "modernc.org/sqlite/lib"
)

func _int8_vec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var blob, ptr, source uintptr
	var bytes, i, offset, rc, source_len, value_type int32
	var result int64
	var _ /* endptr at bp+32 */ uintptr
	var _ /* res at bp+40 */ Ti8
	var _ /* x at bp+0 */ TArray
	_, _, _, _, _, _, _, _, _, _ = blob, bytes, i, offset, ptr, rc, result, source, source_len, value_type
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	if value_type == int32(m_SQLITE_BLOB) {
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		if bytes == 0 {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		**(**uintptr)(__ccgo_up(vector)) = blob
		**(**Tsize_t)(__ccgo_up(dimensions)) = libc.Uint64FromInt32(bytes)
		**(**Tvector_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(Xvector_cleanup_noop)
		return m_SQLITE_OK
	}
	if value_type == int32(m_SQLITE_TEXT) {
		source = libsqlite3.Xsqlite3_value_text(tls, value)
		source_len = libsqlite3.Xsqlite3_value_bytes(tls, value)
		i = 0
		if source_len == 0 {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		rc = Xarray_init(tls, bp, uint64(1), uint64(libc.Xceil(tls, float64(source_len)/float64(2))))
		if rc != m_SQLITE_OK {
			return rc
		}
		// advance leading whitespace to first '['
		for i < source_len {
			if _vecJsonIsSpaceX[libc.Uint8FromInt8(**(**int8)(__ccgo_up(source + uintptr(i))))] != 0 {
				i = i + 1
				continue
			}
			if int32(**(**int8)(__ccgo_up(source + uintptr(i)))) == int32('[') {
				break
			}
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			Xarray_cleanup(tls, bp)
			return int32(m_SQLITE_ERROR)
		}
		if int32(**(**int8)(__ccgo_up(source + uintptr(i)))) != int32('[') {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			Xarray_cleanup(tls, bp)
			return int32(m_SQLITE_ERROR)
		}
		offset = i + int32(1)
		for offset < source_len {
			ptr = source + uintptr(offset)
			**(**int32)(__ccgo_up(libc.X__errno(tls))) = 0
			result = libc.Xstrtol(tls, ptr, bp+32, int32(10))
			if **(**int32)(__ccgo_up(libc.X__errno(tls))) != 0 && result == 0 || **(**int32)(__ccgo_up(libc.X__errno(tls))) == int32(m_ERANGE) && (result == int64(0x7fffffffffffffff) || result == -libc.Int64FromInt64(0x7fffffffffffffff)-libc.Int64FromInt32(1)) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if **(**uintptr)(__ccgo_up(bp + 32)) == ptr {
				if int32(**(**int8)(__ccgo_up(ptr))) != int32(']') {
					libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
					return int32(m_SQLITE_ERROR)
				}
				goto done
			}
			if result < int64(-libc.Int32FromInt32(0x7f)-libc.Int32FromInt32(1)) || result > int64(m_INT8_MAX) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+341, 0)
				return int32(m_SQLITE_ERROR)
			}
			**(**Ti8)(__ccgo_up(bp + 40)) = int8(result)
			Xarray_append(tls, bp, bp+40)
			offset = int32(int64(offset) + (int64(**(**uintptr)(__ccgo_up(bp + 32))) - int64(ptr)))
			for offset < source_len {
				if _vecJsonIsSpaceX[libc.Uint8FromInt8(**(**int8)(__ccgo_up(source + uintptr(offset))))] != 0 {
					offset = offset + 1
					continue
				}
				if int32(**(**int8)(__ccgo_up(source + uintptr(offset)))) == int32(',') {
					offset = offset + 1
					continue
				}
				if int32(**(**int8)(__ccgo_up(source + uintptr(offset)))) == int32(']') {
					goto done
				}
				break
			}
		}
		goto done
	done:
		;
		if (**(**TArray)(__ccgo_up(bp))).Flength > uint64(0) {
			**(**uintptr)(__ccgo_up(vector)) = (**(**TArray)(__ccgo_up(bp))).Fz
			**(**Tsize_t)(__ccgo_up(dimensions)) = (**(**TArray)(__ccgo_up(bp))).Flength
			**(**Tvector_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(libsqlite3.Xsqlite3_free)
			return m_SQLITE_OK
		}
		libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
		return int32(m_SQLITE_ERROR)
	}
	**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+389, 0)
	return int32(m_SQLITE_ERROR)
}

const m_CHAR_MAX = 0x7f

const m___CET__ = 1

const m___DECIMAL_DIG = 21

const m___FLT_EPSILON = 1.19209290e-07

const m___LDBL_DIG = 18

const m___LDBL_EPSILON = 1.08420217248550443401e-19

const m___LDBL_MANT_DIG = 64

const m___LDBL_MAX = "1.18973149535723176502e+4932"

const m___LDBL_MAX_10_EXP = 4932

const m___LDBL_MAX_EXP = 16384

const m___LDBL_MIN = 3.36210314311209350626e-4932

type t__mbstate_t = struct {
	F__mbstateL [0]t__int64_t
	F__mbstate8 [128]int8
}
