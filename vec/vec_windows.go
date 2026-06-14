// Code generated for windows/amd64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors --cpp /usr/bin/x86_64-w64-mingw32-gcc --goarch amd64 --goos windows -build-lines \/\/go:build windows && (amd64 || arm64)\n -map ar=x86_64-w64-mingw32-ar,gcc=x86_64-w64-mingw32-gcc -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build windows && (amd64 || arm64)

package vec

import (
	"unsafe"

	"modernc.org/libc"
	libsqlite3 "modernc.org/sqlite/lib"
)

type SETJMP_FLOAT128 = TSETJMP_FLOAT128

type TSETJMP_FLOAT128 = struct {
	FPart [2]uint64
}

type T_CRT_DOUBLE = struct {
	Fx float64
}

type T_JBTYPE = struct {
	FPart [2]uint64
}

type T_JUMP_BUFFER = struct {
	FFrame uint64
	FRbx   uint64
	FRsp   uint64
	FRbp   uint64
	FRsi   uint64
	FRdi   uint64
	FR12   uint64
	FR13   uint64
	FR14   uint64
	FR15   uint64
	FRip   uint64
	FMxCsr uint32
	FFpCsr uint16
	FSpare uint16
	FXmm6  TSETJMP_FLOAT128
	FXmm7  TSETJMP_FLOAT128
	FXmm8  TSETJMP_FLOAT128
	FXmm9  TSETJMP_FLOAT128
	FXmm10 TSETJMP_FLOAT128
	FXmm11 TSETJMP_FLOAT128
	FXmm12 TSETJMP_FLOAT128
	FXmm13 TSETJMP_FLOAT128
	FXmm14 TSETJMP_FLOAT128
	FXmm15 TSETJMP_FLOAT128
}

type T_LONGDOUBLE = struct {
	Fx float64
}

type T_SETJMP_FLOAT128 = TSETJMP_FLOAT128

type T_complex = struct {
	Fx float64
	Fy float64
}

type T_exception = struct {
	Ftype1  int32
	Fname   uintptr
	Farg1   float64
	Farg2   float64
	Fretval float64
}

type Tjmp_buf = [16]T_JBTYPE

type Tmax_align_t = struct {
	F__max_align_ll int64
	F__max_align_ld float64
}

type Toff_t = int32

// C documentation
//
//	/**
//	 * @brief Initial an array with the given element size and capacity.
//	 *
//	 * @param array
//	 * @param element_size
//	 * @param init_capacity
//	 * @return SQLITE_OK on success, error code on failure. Only error is
//	 * SQLITE_NOMEM
//	 */
func Xarray_init(tls *libc.TLS, array uintptr, element_size Tsize_t, init_capacity Tsize_t) (r int32) {
	var sz int32
	var z uintptr
	_, _ = sz, z
	sz = int32(element_size * init_capacity)
	z = libsqlite3.Xsqlite3_malloc(tls, sz)
	if !(z != 0) {
		return int32(m_SQLITE_NOMEM)
	}
	libc.Xmemset(tls, z, 0, uint64(sz))
	(*TArray)(unsafe.Pointer(array)).Felement_size = element_size
	(*TArray)(unsafe.Pointer(array)).Flength = uint64(0)
	(*TArray)(unsafe.Pointer(array)).Fcapacity = init_capacity
	(*TArray)(unsafe.Pointer(array)).Fz = z
	return m_SQLITE_OK
}

func Xbitmap_clear(tls *libc.TLS, bitmap uintptr, n Ti32) {
	libc.Xmemset(tls, bitmap, 0, uint64(n/int32(m___CHAR_BIT__)))
}

func Xbitmap_copy(tls *libc.TLS, base uintptr, from uintptr, n Ti32) {
	libc.Xmemcpy(tls, base, from, uint64(n/int32(m___CHAR_BIT__)))
}

func Xbitmap_fill(tls *libc.TLS, bitmap uintptr, n Ti32) {
	libc.Xmemset(tls, bitmap, int32(0xFF), uint64(n/int32(m___CHAR_BIT__)))
}

func Xbitmap_new(tls *libc.TLS, n Ti32) (r uintptr) {
	var p uintptr
	_ = p
	p = libsqlite3.Xsqlite3_malloc(tls, int32(uint64(n)*uint64(1)/uint64(m___CHAR_BIT__)))
	if p != 0 {
		libc.Xmemset(tls, p, 0, uint64(n)*uint64(1)/uint64(m___CHAR_BIT__))
	}
	return p
}

func Xbitmap_new_from(tls *libc.TLS, n Ti32, from uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = libsqlite3.Xsqlite3_malloc(tls, int32(uint64(n)*uint64(1)/uint64(m___CHAR_BIT__)))
	if p != 0 {
		libc.Xmemcpy(tls, p, from, uint64(n/int32(m___CHAR_BIT__)))
	}
	return p
}

func Xparse_npy_buffer(tls *libc.TLS, pVTab uintptr, buffer uintptr, bufferLength int32, data uintptr, numElements uintptr, numDimensions uintptr, element_type uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var dataSize, expectedDataSize, totalHeaderLength Ti32
	var header uintptr
	var major, minor Tu8
	var rc int32
	var _ /* fortran_order at bp+4 */ int32
	var _ /* headerLength at bp+0 */ Tuint16_t
	_, _, _, _, _, _, _ = dataSize, expectedDataSize, header, major, minor, rc, totalHeaderLength
	if bufferLength < int32(10) {
		// IMP: V03312_20150
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3171, 0)
		return int32(m_SQLITE_ERROR)
	}
	if libc.Xmemcmp(tls, uintptr(unsafe.Pointer(&_NPY_MAGIC)), buffer, uint64(6)) != 0 {
		// V11954_28792
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3193, 0)
		return int32(m_SQLITE_ERROR)
	}
	major = **(**uint8)(__ccgo_up(buffer + 6))
	minor = **(**uint8)(__ccgo_up(buffer + 7))
	**(**Tuint16_t)(__ccgo_up(bp)) = uint16(0)
	libc.Xmemcpy(tls, bp, buffer+8, uint64(2))
	totalHeaderLength = int32(libc.Uint64FromInt64(6) + libc.Uint64FromInt64(1) + libc.Uint64FromInt64(1) + libc.Uint64FromInt64(2) + uint64(**(**Tuint16_t)(__ccgo_up(bp))))
	dataSize = bufferLength - totalHeaderLength
	if dataSize < 0 {
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3241, 0)
		return int32(m_SQLITE_ERROR)
	}
	header = buffer + 10
	rc = Xparse_npy_header(tls, pVTab, header, uint64(**(**Tuint16_t)(__ccgo_up(bp))), element_type, bp+4, numElements, numDimensions)
	if rc != m_SQLITE_OK {
		return rc
	}
	expectedDataSize = int32(**(**Tsize_t)(__ccgo_up(numElements)) * Xvector_byte_size(tls, **(**_VectorElementType)(__ccgo_up(element_type)), **(**Tsize_t)(__ccgo_up(numDimensions))))
	if expectedDataSize != dataSize {
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3278, libc.VaList(bp+16, expectedDataSize, dataSize))
		return int32(m_SQLITE_ERROR)
	}
	**(**uintptr)(__ccgo_up(data)) = buffer + uintptr(totalHeaderLength)
	return m_SQLITE_OK
}

func Xparse_npy_file(tls *libc.TLS, pVTab uintptr, file uintptr, pCur uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var dataSize, expectedDataSize Ti32
	var fileSize, n, rc int32
	var headerX uintptr
	var major, minor Tu8
	var totalHeaderLength Tsize_t
	var _ /* element_type at bp+16 */ _VectorElementType
	var _ /* fortran_order at bp+12 */ int32
	var _ /* header at bp+0 */ [10]uint8
	var _ /* headerLength at bp+10 */ Tuint16_t
	var _ /* numDimensions at bp+32 */ Tsize_t
	var _ /* numElements at bp+24 */ Tsize_t
	_, _, _, _, _, _, _, _, _ = dataSize, expectedDataSize, fileSize, headerX, major, minor, n, rc, totalHeaderLength
	libc.Xfseek(tls, file, 0, int32(m_SEEK_END))
	fileSize = libc.Xftell(tls, file)
	libc.Xfseek(tls, file, 0, m_SEEK_SET)
	n = int32(libc.Xfread(tls, bp, uint64(1), uint64(10), file))
	if n != int32(10) {
		Xvtab_set_error(tls, pVTab, __ccgo_ts+2988, 0)
		return int32(m_SQLITE_ERROR)
	}
	if libc.Xmemcmp(tls, uintptr(unsafe.Pointer(&_NPY_MAGIC)), bp, uint64(6)) != 0 {
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3015, 0)
		return int32(m_SQLITE_ERROR)
	}
	major = (**(**[10]uint8)(__ccgo_up(bp)))[int32(6)]
	minor = (**(**[10]uint8)(__ccgo_up(bp)))[int32(7)]
	**(**Tuint16_t)(__ccgo_up(bp + 10)) = uint16(0)
	libc.Xmemcpy(tls, bp+10, bp+8, uint64(2))
	totalHeaderLength = libc.Uint64FromInt64(6) + libc.Uint64FromInt64(1) + libc.Uint64FromInt64(1) + libc.Uint64FromInt64(2) + uint64(**(**Tuint16_t)(__ccgo_up(bp + 10)))
	dataSize = int32(uint64(fileSize) - totalHeaderLength)
	if dataSize < 0 {
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3068, 0)
		return int32(m_SQLITE_ERROR)
	}
	headerX = libsqlite3.Xsqlite3_malloc(tls, int32(**(**Tuint16_t)(__ccgo_up(bp + 10))))
	if **(**Tuint16_t)(__ccgo_up(bp + 10)) != 0 && !(headerX != 0) {
		return int32(m_SQLITE_NOMEM)
	}
	n = int32(libc.Xfread(tls, headerX, uint64(1), uint64(**(**Tuint16_t)(__ccgo_up(bp + 10))), file))
	if n != int32(**(**Tuint16_t)(__ccgo_up(bp + 10))) {
		libsqlite3.Xsqlite3_free(tls, headerX)
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3068, 0)
		return int32(m_SQLITE_ERROR)
	}
	rc = Xparse_npy_header(tls, pVTab, headerX, uint64(**(**Tuint16_t)(__ccgo_up(bp + 10))), bp+16, bp+12, bp+24, bp+32)
	libsqlite3.Xsqlite3_free(tls, headerX)
	if rc != m_SQLITE_OK {
		// parse_npy_header already attackes an error emssage
		return rc
	}
	expectedDataSize = int32(**(**Tsize_t)(__ccgo_up(bp + 24)) * Xvector_byte_size(tls, **(**_VectorElementType)(__ccgo_up(bp + 16)), **(**Tsize_t)(__ccgo_up(bp + 32))))
	if expectedDataSize != dataSize {
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3110, libc.VaList(bp+48, expectedDataSize, dataSize))
		return int32(m_SQLITE_ERROR)
	}
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FmaxChunks = uint64(1024)
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBufferSize = Xvector_byte_size(tls, **(**_VectorElementType)(__ccgo_up(bp + 16)), **(**Tsize_t)(__ccgo_up(bp + 32))) * (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FmaxChunks
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer = libsqlite3.Xsqlite3_malloc(tls, int32((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBufferSize))
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBufferSize != 0 && !((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer != 0) {
		return int32(m_SQLITE_NOMEM)
	}
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkSize = libc.Xfread(tls, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer, Xvector_byte_size(tls, **(**_VectorElementType)(__ccgo_up(bp + 16)), **(**Tsize_t)(__ccgo_up(bp + 32))), (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FmaxChunks, file)
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkIndex = uint64(0)
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType = **(**_VectorElementType)(__ccgo_up(bp + 16))
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnElements = **(**Tsize_t)(__ccgo_up(bp + 24))
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions = **(**Tsize_t)(__ccgo_up(bp + 32))
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Finput_type = int32(_VEC_NPY_EACH_INPUT_FILE)
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Feof = libc.BoolInt32((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkSize == uint64(0))
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile = file
	return m_SQLITE_OK
}

func Xparse_npy_header(tls *libc.TLS, pVTab uintptr, header uintptr, headerLength Tsize_t, out_element_type uintptr, fortran_order uintptr, numElements uintptr, numDimensions uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var first Tsize_t
	var key uintptr
	var rc, v1 int32
	var _ /* scanner at bp+0 */ TNpyScanner
	var _ /* token at bp+24 */ TNpyToken
	_, _, _, _ = first, key, rc, v1
	Xnpy_scanner_init(tls, bp, header, int32(headerLength))
	if Xnpy_scanner_next(tls, bp, bp+24) != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_LBRACE) {
		Xvtab_set_error(tls, pVTab, __ccgo_ts+1922, 0)
		return int32(m_SQLITE_ERROR)
	}
	for int32(1) != 0 {
		rc = Xnpy_scanner_next(tls, bp, bp+24)
		if rc != int32(m_VEC0_TOKEN_RESULT_SOME) {
			Xvtab_set_error(tls, pVTab, __ccgo_ts+1985, 0)
			return int32(m_SQLITE_ERROR)
		}
		if (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type == int32(_NPY_TOKEN_TYPE_RBRACE) {
			break
		}
		if (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_STRING) {
			Xvtab_set_error(tls, pVTab, __ccgo_ts+2041, 0)
			return int32(m_SQLITE_ERROR)
		}
		key = (**(**TNpyToken)(__ccgo_up(bp + 24))).Fstart
		rc = Xnpy_scanner_next(tls, bp, bp+24)
		if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_COLON) {
			Xvtab_set_error(tls, pVTab, __ccgo_ts+2109, 0)
			return int32(m_SQLITE_ERROR)
		}
		if libc.Xstrncmp(tls, key, __ccgo_ts+2177, libc.Xstrlen(tls, __ccgo_ts+2177)) == 0 {
			rc = Xnpy_scanner_next(tls, bp, bp+24)
			if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_STRING) {
				Xvtab_set_error(tls, pVTab, __ccgo_ts+2185, 0)
				return int32(m_SQLITE_ERROR)
			}
			if libc.Xstrncmp(tls, (**(**TNpyToken)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+2254, libc.Xstrlen(tls, __ccgo_ts+2254)) != 0 {
				Xvtab_set_error(tls, pVTab, __ccgo_ts+2260, 0)
				return int32(m_SQLITE_ERROR)
			}
			**(**_VectorElementType)(__ccgo_up(out_element_type)) = int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32)
		} else {
			if libc.Xstrncmp(tls, key, __ccgo_ts+2349, libc.Xstrlen(tls, __ccgo_ts+2349)) == 0 {
				rc = Xnpy_scanner_next(tls, bp, bp+24)
				if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_FALSE) {
					Xvtab_set_error(tls, pVTab, __ccgo_ts+2365, 0)
					return int32(m_SQLITE_ERROR)
				}
				**(**int32)(__ccgo_up(fortran_order)) = 0
			} else {
				if libc.Xstrncmp(tls, key, __ccgo_ts+2462, libc.Xstrlen(tls, __ccgo_ts+2462)) == 0 {
					rc = Xnpy_scanner_next(tls, bp, bp+24)
					if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_LPAREN) {
						Xvtab_set_error(tls, pVTab, __ccgo_ts+2470, 0)
						return int32(m_SQLITE_ERROR)
					}
					rc = Xnpy_scanner_next(tls, bp, bp+24)
					if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_NUMBER) {
						Xvtab_set_error(tls, pVTab, __ccgo_ts+2543, 0)
						return int32(m_SQLITE_ERROR)
					}
					first = uint64(libc.Xstrtol(tls, (**(**TNpyToken)(__ccgo_up(bp + 24))).Fstart, libc.UintptrFromInt32(0), int32(10)))
					rc = Xnpy_scanner_next(tls, bp, bp+24)
					if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_COMMA) {
						Xvtab_set_error(tls, pVTab, __ccgo_ts+2612, 0)
						return int32(m_SQLITE_ERROR)
					}
					rc = Xnpy_scanner_next(tls, bp, bp+24)
					if rc != int32(m_VEC0_TOKEN_RESULT_SOME) {
						Xvtab_set_error(tls, pVTab, __ccgo_ts+2678, 0)
						return int32(m_SQLITE_ERROR)
					}
					if (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type == int32(_NPY_TOKEN_TYPE_NUMBER) {
						**(**Tsize_t)(__ccgo_up(numElements)) = first
						**(**Tsize_t)(__ccgo_up(numDimensions)) = uint64(libc.Xstrtol(tls, (**(**TNpyToken)(__ccgo_up(bp + 24))).Fstart, libc.UintptrFromInt32(0), int32(10)))
						rc = Xnpy_scanner_next(tls, bp, bp+24)
						if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_RPAREN) {
							Xvtab_set_error(tls, pVTab, __ccgo_ts+2747, 0)
							return int32(m_SQLITE_ERROR)
						}
					} else {
						if (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type == int32(_NPY_TOKEN_TYPE_RPAREN) {
							// '(0,)' means an empty array!
							if first != 0 {
								v1 = int32(1)
							} else {
								v1 = 0
							}
							**(**Tsize_t)(__ccgo_up(numElements)) = uint64(v1)
							**(**Tsize_t)(__ccgo_up(numDimensions)) = first
						} else {
							Xvtab_set_error(tls, pVTab, __ccgo_ts+2819, 0)
							return int32(m_SQLITE_ERROR)
						}
					}
				} else {
					Xvtab_set_error(tls, pVTab, __ccgo_ts+2874, 0)
					return int32(m_SQLITE_ERROR)
				}
			}
		}
		rc = Xnpy_scanner_next(tls, bp, bp+24)
		if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_COMMA) {
			Xvtab_set_error(tls, pVTab, __ccgo_ts+2929, 0)
			return int32(m_SQLITE_ERROR)
		}
	}
	return m_SQLITE_OK
}

func Xsqlite3_vec_init(tls *libc.TLS, db uintptr, pzErrMsg uintptr, pApi uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, i1 uint32
	var rc int32
	_, _, _ = i, i1, rc
	rc = m_SQLITE_OK
	rc = libsqlite3.Xsqlite3_create_function_v2(tls, db, __ccgo_ts+16107, 0, libc.Int32FromInt32(m_SQLITE_UTF8)|libc.Int32FromInt32(m_SQLITE_INNOCUOUS)|libc.Int32FromInt32(m_SQLITE_DETERMINISTIC), __ccgo_ts+6911, __ccgo_fp(__static_text_func), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
	if rc != m_SQLITE_OK {
		return rc
	}
	rc = libsqlite3.Xsqlite3_create_function_v2(tls, db, __ccgo_ts+16119, 0, libc.Int32FromInt32(m_SQLITE_UTF8)|libc.Int32FromInt32(m_SQLITE_INNOCUOUS)|libc.Int32FromInt32(m_SQLITE_DETERMINISTIC), __ccgo_ts+16129, __ccgo_fp(__static_text_func), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
	if rc != m_SQLITE_OK {
		return rc
	}
	i = uint32(0)
	for {
		if !(uint64(i) < libc.Uint64FromInt64(384)/libc.Uint64FromInt64(24) && rc == m_SQLITE_OK) {
			break
		}
		rc = libsqlite3.Xsqlite3_create_function_v2(tls, db, _aFunc[i].FzFName, _aFunc[i].FnArg, _aFunc[i].Fflags, libc.UintptrFromInt32(0), _aFunc[i].FxFunc, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
		if rc != m_SQLITE_OK {
			**(**uintptr)(__ccgo_up(pzErrMsg)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+16423, libc.VaList(bp+8, _aFunc[i].FzFName, libsqlite3.Xsqlite3_errmsg(tls, db)))
			return rc
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	i1 = uint32(0)
	for {
		if !(uint64(i1) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(32) && rc == m_SQLITE_OK) {
			break
		}
		rc = libsqlite3.Xsqlite3_create_module_v2(tls, db, _aMod[i1].Fname, _aMod[i1].Fmodule, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
		if rc != m_SQLITE_OK {
			**(**uintptr)(__ccgo_up(pzErrMsg)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+16454, libc.VaList(bp+8, _aMod[i1].Fname, libsqlite3.Xsqlite3_errmsg(tls, db)))
			return rc
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	return m_SQLITE_OK
}

func Xvec0Filter_knn_chunks_iter(tls *libc.TLS, p uintptr, stmtChunks uintptr, vector_column uintptr, vectorColumnIdx int32, arrayRowidsIn uintptr, aMetadataIn uintptr, idxStr uintptr, argc int32, argv uintptr, queryVector uintptr, k Ti64, out_topk_rowids uintptr, out_topk_distances uintptr, out_used uintptr) (r int32) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var b, bTaken, baseVectors, base_i, base_i1, base_i2, bmMetadata, bmRowids, chunkRowids, chunkValidity, chunk_distances, chunk_topk_idxs, in, tmp_topk_distances, tmp_topk_rowids, topk_distances, topk_rowids, v1 uintptr
	var baseVectorsSize, chunk_id, currentBaseVectorsSize, expectedBaseVectorsSize, k_used, rowidsSize, validitySize Ti64
	var hasDistanceConstraints, hasMetadataFilters, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, idx, idx1, idx2, idxStrLength, metadata_idx, numValueEntries, operator, rc, v4 int32
	var kind, kind1, kind2 int8
	var op Tvec0_distance_constraint_operator
	var result, target Tf32
	var v12, v13, v14 int64
	var _ /* blobVectors at bp+0 */ uintptr
	var _ /* metadataBlobs at bp+8 */ [16]uintptr
	var _ /* rowid at bp+136 */ Ti64
	var _ /* used at bp+152 */ Ti64
	var _ /* used1 at bp+144 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b, bTaken, baseVectors, baseVectorsSize, base_i, base_i1, base_i2, bmMetadata, bmRowids, chunkRowids, chunkValidity, chunk_distances, chunk_id, chunk_topk_idxs, currentBaseVectorsSize, expectedBaseVectorsSize, hasDistanceConstraints, hasMetadataFilters, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, idx, idx1, idx2, idxStrLength, in, k_used, kind, kind1, kind2, metadata_idx, numValueEntries, op, operator, rc, result, rowidsSize, target, tmp_topk_distances, tmp_topk_rowids, topk_distances, topk_rowids, validitySize, v1, v12, v13, v14, v4
	// for each chunk, get top min(k, chunk_size) rowid + distances to query vec.
	// then reconcile all topk_chunks for a true top k.
	// output only rowids + distances for now
	rc = m_SQLITE_OK
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
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
	topk_rowids = libsqlite3.Xsqlite3_malloc(tls, int32(uint64(k)*uint64(8)))
	if !(topk_rowids != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, topk_rowids, 0, uint64(k)*uint64(8))
	topk_distances = libsqlite3.Xsqlite3_malloc(tls, int32(uint64(k)*uint64(4)))
	if !(topk_distances != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, topk_distances, 0, uint64(k)*uint64(4))
	tmp_topk_rowids = libsqlite3.Xsqlite3_malloc(tls, int32(uint64(k)*uint64(8)))
	if !(tmp_topk_rowids != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, tmp_topk_rowids, 0, uint64(k)*uint64(8))
	tmp_topk_distances = libsqlite3.Xsqlite3_malloc(tls, int32(uint64(k)*uint64(4)))
	if !(tmp_topk_distances != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, tmp_topk_distances, 0, uint64(k)*uint64(4))
	k_used = 0
	baseVectorsSize = int64(uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(vector_column))))
	baseVectors = libsqlite3.Xsqlite3_malloc(tls, int32(baseVectorsSize))
	if !(baseVectors != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	chunk_distances = libsqlite3.Xsqlite3_malloc(tls, int32(uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(4)))
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
	chunk_topk_idxs = libsqlite3.Xsqlite3_malloc(tls, int32(uint64(k)*uint64(4)))
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
	idxStrLength = int32(libc.Xstrlen(tls, idxStr))
	numValueEntries = (idxStrLength - int32(1)) / int32(4)
	hasMetadataFilters = 0
	hasDistanceConstraints = 0
	i = 0
	for {
		if !(i < argc) {
			break
		}
		idx = int32(1) + i*int32(4)
		kind = **(**int8)(__ccgo_up(idxStr + uintptr(idx+0)))
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
		libc.Xmemset(tls, chunk_distances, 0, uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(4))
		libc.Xmemset(tls, chunk_topk_idxs, 0, uint64(k)*uint64(4))
		Xbitmap_clear(tls, b, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		chunk_id = libsqlite3.Xsqlite3_column_int64(tls, stmtChunks, 0)
		chunkValidity = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(1))
		validitySize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(1)))
		if validitySize != int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT__)) {
			// IMP: V05271_22109
			Xvtab_set_error(tls, p, __ccgo_ts+9715, libc.VaList(bp+168, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT__), validitySize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		chunkRowids = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(2))
		rowidsSize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(2)))
		if uint64(rowidsSize) != uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8) {
			// IMP: V02796_19635
			Xvtab_set_error(tls, p, __ccgo_ts+9777, 0)
			Xvtab_set_error(tls, p, __ccgo_ts+9803, libc.VaList(bp+168, uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8), rowidsSize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		// open the vector chunk blob for the current chunk
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(vectorColumnIdx)*8)), __ccgo_ts+3712, chunk_id, 0, bp)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+9863, libc.VaList(bp+168, chunk_id))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		currentBaseVectorsSize = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp))))
		expectedBaseVectorsSize = int64(uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(vector_column))))
		if currentBaseVectorsSize != expectedBaseVectorsSize {
			// IMP: V16465_00535
			Xvtab_set_error(tls, p, __ccgo_ts+9906, libc.VaList(bp+168, expectedBaseVectorsSize, currentBaseVectorsSize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), baseVectors, int32(currentBaseVectorsSize), 0)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+9966, libc.VaList(bp+168, chunk_id))
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
				**(**Ti64)(__ccgo_up(bp + 136)) = **(**Ti64)(__ccgo_up(chunkRowids + uintptr(i1)*8))
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
				kind1 = **(**int8)(__ccgo_up(idxStr + uintptr(idx1+0)))
				if int32(kind1) != int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) {
					goto _5
				}
				metadata_idx = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx1+int32(1))))) - int32('A')
				operator = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx1+int32(2)))))
				if !((**(**[16]uintptr)(__ccgo_up(bp + 8)))[metadata_idx] != 0) {
					rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 480 + uintptr(metadata_idx)*8)), __ccgo_ts+4053, chunk_id, 0, bp+8+uintptr(metadata_idx)*8)
					Xvtab_set_error(tls, p, __ccgo_ts+9999, 0)
					if rc != m_SQLITE_OK {
						goto cleanup
					}
				}
				Xbitmap_clear(tls, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
				rc = Xvec0_set_metadata_filter_bitmap(tls, p, metadata_idx, operator, **(**uintptr)(__ccgo_up(argv + uintptr(i2)*8)), (**(**[16]uintptr)(__ccgo_up(bp + 8)))[metadata_idx], chunk_id, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size, aMetadataIn, i2)
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
				base_i = baseVectors + uintptr(uint64(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)*4
				switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdistance_metric {
				case int32(_VEC0_DISTANCE_METRIC_L2):
					result = _distance_l2_sqr_float(tls, base_i, queryVector, vector_column+16)
				case int32(_VEC0_DISTANCE_METRIC_L1):
					result = float32(_distance_l1_f32(tls, base_i, queryVector, vector_column+16))
				case int32(_VEC0_DISTANCE_METRIC_COSINE):
					result = _distance_cosine_float(tls, base_i, queryVector, vector_column+16)
					break
				}
			case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
				base_i1 = baseVectors + uintptr(uint64(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)
				switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdistance_metric {
				case int32(_VEC0_DISTANCE_METRIC_L2):
					result = _distance_l2_sqr_int8(tls, base_i1, queryVector, vector_column+16)
				case int32(_VEC0_DISTANCE_METRIC_L1):
					result = float32(_distance_l1_int8(tls, base_i1, queryVector, vector_column+16))
				case int32(_VEC0_DISTANCE_METRIC_COSINE):
					result = _distance_cosine_int8(tls, base_i1, queryVector, vector_column+16)
					break
				}
			case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
				base_i2 = baseVectors + uintptr(uint64(i3)*((*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions/libc.Uint64FromInt32(m___CHAR_BIT__)))
				result = _distance_hamming(tls, base_i2, queryVector, vector_column+16)
				break
			}
			**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i3)*4)) = result
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
				kind2 = **(**int8)(__ccgo_up(idxStr + uintptr(idx2+0)))
				// TODO casts f64 to f32, is that a problem?
				target = float32(libsqlite3.Xsqlite3_value_double(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i4)*8))))
				if int32(kind2) != int32(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT) {
					goto _7
				}
				op = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx2+int32(1)))))
				switch op {
				case int32(_VEC0_DISTANCE_CONSTRAINT_GE):
					i5 = 0
					for {
						if !(i5 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						if Xbitmap_get(tls, b, i5) != 0 && !(**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i5)*4)) >= target) {
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
						if Xbitmap_get(tls, b, i6) != 0 && !(**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i6)*4)) > target) {
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
						if Xbitmap_get(tls, b, i7) != 0 && !(**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i7)*4)) <= target) {
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
						if Xbitmap_get(tls, b, i8) != 0 && !(**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i8)*4)) < target) {
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
		Xmin_idx(tls, chunk_distances, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size, b, chunk_topk_idxs, int32(v12), bTaken, bp+144)
		if k <= int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
			v13 = k
		} else {
			v13 = int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		}
		if v13 <= int64(**(**int32)(__ccgo_up(bp + 144))) {
			if k <= int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
				v14 = k
			} else {
				v14 = int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
			}
			v12 = v14
		} else {
			v12 = int64(**(**int32)(__ccgo_up(bp + 144)))
		}
		Xmerge_sorted_lists(tls, topk_distances, topk_rowids, k_used, chunk_distances, chunkRowids, chunk_topk_idxs, v12, tmp_topk_distances, tmp_topk_rowids, k, bp+152)
		i9 = 0
		for {
			if !(int64(i9) < **(**Ti64)(__ccgo_up(bp + 152))) {
				break
			}
			**(**Ti64)(__ccgo_up(topk_rowids + uintptr(i9)*8)) = **(**Ti64)(__ccgo_up(tmp_topk_rowids + uintptr(i9)*8))
			**(**Tf32)(__ccgo_up(topk_distances + uintptr(i9)*4)) = **(**Tf32)(__ccgo_up(tmp_topk_distances + uintptr(i9)*4))
			goto _16
		_16:
			;
			i9 = i9 + 1
		}
		k_used = **(**Ti64)(__ccgo_up(bp + 152))
		// blobVectors is always opened with read-only permissions, so this never
		// fails.
		libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
		**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	}
	**(**uintptr)(__ccgo_up(out_topk_rowids)) = topk_rowids
	**(**uintptr)(__ccgo_up(out_topk_distances)) = topk_distances
	**(**Ti64)(__ccgo_up(out_used)) = k_used
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
		libsqlite3.Xsqlite3_blob_close(tls, (**(**[16]uintptr)(__ccgo_up(bp + 8)))[i10])
		goto _17
	_17:
		;
		i10 = i10 + 1
	}
	// blobVectors is always opened with read-only permissions, so this never
	// fails.
	libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	return rc
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
	kind = (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 480 + uintptr(metadata_idx)*8)), __ccgo_ts+4053, chunk_id, int32(1), bp)
	if rc != m_SQLITE_OK {
		return rc
	}
	switch kind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), int32(chunk_offset/libc.Uint64FromInt32(m___CHAR_BIT__)))
		if rc != m_SQLITE_OK {
			goto done
		}
		**(**Tu8)(__ccgo_up(bp + 8)) = uint8(int32(**(**Tu8)(__ccgo_up(bp + 8))) & ^(libc.Int32FromInt32(1) << (chunk_offset % libc.Uint64FromInt32(m___CHAR_BIT__))))
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), int32(chunk_offset/uint64(m___CHAR_BIT__)))
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		**(**Ti64)(__ccgo_up(bp + 16)) = 0
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+16, int32(8), int32(chunk_offset*uint64(8)))
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		**(**float64)(__ccgo_up(bp + 24)) = libc.Float64FromInt32(0)
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+24, int32(8), int32(chunk_offset*uint64(8)))
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+32, int32(4), int32(chunk_offset*uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if rc != m_SQLITE_OK {
			goto done
		}
		libc.Xmemset(tls, bp+36, 0, uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+36, int32(16), int32(chunk_offset*uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if rc != m_SQLITE_OK {
			goto done
		}
		if **(**int32)(__ccgo_up(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12987, libc.VaList(bp+72, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_idx))
			if !(zSql != 0) {
				rc = int32(m_SQLITE_NOMEM)
				goto done
			}
			rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+56, libc.UintptrFromInt32(0))
			if rc != m_SQLITE_OK {
				goto done
			}
			libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 56)), int32(1), rowid)
			rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 56)))
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 56)))
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
	rc2 = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	if rc == m_SQLITE_OK {
		return rc2
	}
	return rc
}

func Xvec0Update_Delete_ClearValidity(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var brc, rc, validityOffset int32
	var mask uint8
	var _ /* blobChunksValidity at bp+0 */ uintptr
	var _ /* bx at bp+8 */ uint8
	var _ /* result at bp+9 */ int8
	_, _, _, _ = brc, mask, rc, validityOffset
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	validityOffset = int32(chunk_offset / uint64(m___CHAR_BIT__))
	// 2. ensure chunks.validity bit is 1, then set to 0
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+11207, chunk_id, int32(1), bp)
	if rc != m_SQLITE_OK {
		// IMP: V26002_10073
		Xvtab_set_error(tls, p, __ccgo_ts+13737, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id))
		return int32(m_SQLITE_ERROR)
	}
	// will skip the sqlite3_blob_bytes(blobChunksValidity) check for now,
	// the read below would catch it
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), validityOffset)
	if rc != m_SQLITE_OK {
		// IMP: V21193_05263
		Xvtab_set_error(tls, p, __ccgo_ts+13781, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		goto cleanup
	}
	if !(int32(**(**uint8)(__ccgo_up(bp + 8)))>>(chunk_offset%libc.Uint64FromInt32(m___CHAR_BIT__)) != 0) {
		// IMP: V21193_05263
		rc = int32(m_SQLITE_ERROR)
		Xvtab_set_error(tls, p, __ccgo_ts+13831, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		goto cleanup
	}
	mask = uint8(^(libc.Int32FromInt32(1) << (chunk_offset % libc.Uint64FromInt32(m___CHAR_BIT__))))
	**(**int8)(__ccgo_up(bp + 9)) = int8(int32(**(**uint8)(__ccgo_up(bp + 8))) & int32(mask))
	rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+9, int32(1), validityOffset)
	if rc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+13897, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		goto cleanup
	}
	goto cleanup
cleanup:
	;
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	if rc != m_SQLITE_OK {
		return rc
	}
	if brc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+13951, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		return brc
	}
	return m_SQLITE_OK
}

func Xvec0Update_Delete_ClearVectors(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var brc, i, rc int32
	var n Tsize_t
	var zeroBuf uintptr
	var _ /* blobVectors at bp+0 */ uintptr
	_, _, _, _, _ = brc, i, n, rc, zeroBuf
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
		n = Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32)))
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), __ccgo_ts+3712, chunk_id, int32(1), bp)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+14213, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id, i))
			return int32(m_SQLITE_ERROR)
		}
		zeroBuf = libsqlite3.Xsqlite3_malloc(tls, int32(n))
		if !(zeroBuf != 0) {
			libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
			return int32(m_SQLITE_NOMEM)
		}
		libc.Xmemset(tls, zeroBuf, 0, n)
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), zeroBuf, int32(n), int32(chunk_offset*n))
		libsqlite3.Xsqlite3_free(tls, zeroBuf)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+14265, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id, chunk_offset, i))
		}
		brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
		if rc != m_SQLITE_OK {
			return rc
		}
		if brc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+14329, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id, i))
			return brc
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return m_SQLITE_OK
}

func Xvec0Update_Delete_DeleteChunkIfEmpty(tls *libc.TLS, p uintptr, chunk_id Ti64, deleted uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var allZero, brc, i, i1, i2, rc, validitySize int32
	var validityBuf, zSql uintptr
	var _ /* blobValidity at bp+0 */ uintptr
	var _ /* stmt at bp+8 */ uintptr
	_, _, _, _, _, _, _, _, _ = allZero, brc, i, i1, i2, rc, validityBuf, validitySize, zSql
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	**(**int32)(__ccgo_up(deleted)) = 0
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+11207, chunk_id, 0, bp)
	if rc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+14414, libc.VaList(bp+24, chunk_id))
		return int32(m_SQLITE_ERROR)
	}
	validitySize = libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp)))
	validityBuf = libsqlite3.Xsqlite3_malloc(tls, validitySize)
	if !(validityBuf != 0) {
		libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
		return int32(m_SQLITE_NOMEM)
	}
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), validityBuf, validitySize, 0)
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_free(tls, validityBuf)
		return rc
	}
	if brc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_free(tls, validityBuf)
		return brc
	}
	allZero = int32(1)
	i = 0
	for {
		if !(i < validitySize) {
			break
		}
		if int32(**(**uint8)(__ccgo_up(validityBuf + uintptr(i)))) != 0 {
			allZero = 0
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	libsqlite3.Xsqlite3_free(tls, validityBuf)
	if !(allZero != 0) {
		return m_SQLITE_OK
	}
	// Delete from _chunks
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14458, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	if !(zSql != 0) {
		return int32(m_SQLITE_NOMEM)
	}
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+8, libc.UintptrFromInt32(0))
	libsqlite3.Xsqlite3_free(tls, zSql)
	if rc != m_SQLITE_OK {
		return rc
	}
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 8)), int32(1), chunk_id)
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	if rc != int32(m_SQLITE_DONE) {
		return int32(m_SQLITE_ERROR)
	}
	// Delete from each _vector_chunksNN
	i1 = 0
	for {
		if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14503, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, i1))
		if !(zSql != 0) {
			return int32(m_SQLITE_NOMEM)
		}
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+8, libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_free(tls, zSql)
		if rc != m_SQLITE_OK {
			return rc
		}
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 8)), int32(1), chunk_id)
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		if rc != int32(m_SQLITE_DONE) {
			return int32(m_SQLITE_ERROR)
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// Delete from each _metadatachunksNN
	i2 = 0
	for {
		if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumMetadataColumns) {
			break
		}
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14559, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, i2))
		if !(zSql != 0) {
			return int32(m_SQLITE_NOMEM)
		}
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+8, libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_free(tls, zSql)
		if rc != m_SQLITE_OK {
			return rc
		}
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 8)), int32(1), chunk_id)
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		if rc != int32(m_SQLITE_DONE) {
			return int32(m_SQLITE_ERROR)
		}
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	// Invalidate cached stmtLatestChunk so it gets re-prepared on next insert
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk != 0 {
		libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk)
		(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk = libc.UintptrFromInt32(0)
	}
	**(**int32)(__ccgo_up(deleted)) = int32(1)
	return m_SQLITE_OK
}

// C documentation
//
//	/**
//	 * @brief Handles INSERT INTO operations on a vec0 table.
//	 *
//	 * @return int SQLITE_OK on success, otherwise error code on failure
//	 */
func Xvec0Update_Insert(tls *libc.TLS, pVTab uintptr, argc int32, argv uintptr, pRowid uintptr) (r int32) {
	bp := tls.Alloc(400)
	defer tls.Free(400)
	var auxiliary_key_idx, brc, i, i1, i2, i3, i4, i5, i6, metadata_idx, new_value_type, numReadVectors, partition_key_idx, rc, v_type, vector_column_idx int32
	var p, s, v, v1, valueVector, zSql uintptr
	var _ /* blobChunksValidity at bp+312 */ uintptr
	var _ /* bufferChunksValidity at bp+320 */ uintptr
	var _ /* chunk_offset at bp+304 */ Ti64
	var _ /* chunk_rowid at bp+296 */ Ti64
	var _ /* cleanups at bp+136 */ [16]Tvector_cleanup
	var _ /* dimensions at bp+328 */ Tsize_t
	var _ /* elementType at bp+344 */ _VectorElementType
	var _ /* partitionKeyValues at bp+264 */ [4]uintptr
	var _ /* pzError at bp+336 */ uintptr
	var _ /* rowid at bp+0 */ Ti64
	var _ /* stmt at bp+352 */ uintptr
	var _ /* vectorDatas at bp+8 */ [16]uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = auxiliary_key_idx, brc, i, i1, i2, i3, i4, i5, i6, metadata_idx, new_value_type, numReadVectors, p, partition_key_idx, rc, s, v, v1, v_type, valueVector, vector_column_idx, zSql
	_ = argc
	p = pVTab
	// a write-able blob of the validity column for the given chunk. Used to mark
	// validity bit
	**(**uintptr)(__ccgo_up(bp + 312)) = libc.UintptrFromInt32(0)
	// buffer for the valididty column for the given chunk. Maybe not needed here?
	**(**uintptr)(__ccgo_up(bp + 320)) = libc.UintptrFromInt32(0)
	numReadVectors = 0
	// Read all provided partition key values into partitionKeyValues
	i = 0
	for {
		if !(i < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION) {
			goto _1
		}
		partition_key_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i))))
		(**(**[4]uintptr)(__ccgo_up(bp + 264)))[partition_key_idx] = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i)*8))
		new_value_type = libsqlite3.Xsqlite3_value_type(tls, (**(**[4]uintptr)(__ccgo_up(bp + 264)))[partition_key_idx])
		if new_value_type != int32(m_SQLITE_NULL) && new_value_type != (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(partition_key_idx)*24))).Ftype1 {
			// IMP: V11454_28292
			Xvtab_set_error(tls, pVTab, __ccgo_ts+13042, libc.VaList(bp+368, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(partition_key_idx)*24))).Fname_length, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(partition_key_idx)*24))).Fname, Xtype_name(tls, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(partition_key_idx)*24))).Ftype1), Xtype_name(tls, new_value_type)))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// read all the inserted vectors  into vectorDatas, validate their lengths.
	i1 = 0
	for {
		if !(i1 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i1)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR) {
			goto _2
		}
		vector_column_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i1))))
		valueVector = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i1)*8))
		rc = Xvector_from_value(tls, valueVector, bp+8+uintptr(vector_column_idx)*8, bp+328, bp+344, bp+136+uintptr(vector_column_idx)*8, bp+336)
		if rc != m_SQLITE_OK {
			// IMP: V06519_23358
			Xvtab_set_error(tls, pVTab, __ccgo_ts+13134, libc.VaList(bp+368, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fname, **(**uintptr)(__ccgo_up(bp + 336))))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		numReadVectors = numReadVectors + 1
		if **(**_VectorElementType)(__ccgo_up(bp + 344)) != (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Felement_type {
			// IMP: V08221_25059
			Xvtab_set_error(tls, pVTab, __ccgo_ts+13187, libc.VaList(bp+368, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i1)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i1)*32))).Fname, Xvector_subtype_name(tls, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i1)*32))).Felement_type), Xvector_subtype_name(tls, **(**_VectorElementType)(__ccgo_up(bp + 344)))))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		if **(**Tsize_t)(__ccgo_up(bp + 328)) != (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fdimensions {
			// IMP: V01145_17984
			Xvtab_set_error(tls, pVTab, __ccgo_ts+13285, libc.VaList(bp+368, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fname, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fdimensions, **(**Tsize_t)(__ccgo_up(bp + 328))))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// Cannot insert a value in the hidden "distance" column
	if libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv + uintptr(int32(2)+Xvec0_column_distance_idx(tls, p))*8))) != int32(m_SQLITE_NULL) {
		// IMP: V24228_08298
		Xvtab_set_error(tls, pVTab, __ccgo_ts+13387, 0)
		rc = int32(m_SQLITE_ERROR)
		goto cleanup
	}
	// Cannot insert a value in the hidden "k" column
	if libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv + uintptr(int32(2)+Xvec0_column_k_idx(tls, p))*8))) != int32(m_SQLITE_NULL) {
		// IMP: V11875_28713
		Xvtab_set_error(tls, pVTab, __ccgo_ts+13442, 0)
		rc = int32(m_SQLITE_ERROR)
		goto cleanup
	}
	// Step #1: Insert/get a rowid for this row, from the _rowids table.
	rc = Xvec0Update_InsertRowidStep(tls, p, **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_ID))*8)), bp)
	if rc != m_SQLITE_OK {
		goto cleanup
	}
	// Step #2: Find the next "available" position in the _chunks table for this
	// row.
	rc = Xvec0Update_InsertNextAvailableStep(tls, p, bp+264, bp+296, bp+304, bp+312, bp+320)
	if rc != m_SQLITE_OK {
		goto cleanup
	}
	// Step #3: With the next available chunk position, write out all the vectors
	//          to their specified location.
	rc = Xvec0Update_InsertWriteFinalStep(tls, p, **(**Ti64)(__ccgo_up(bp + 296)), **(**Ti64)(__ccgo_up(bp + 304)), **(**Ti64)(__ccgo_up(bp)), bp+8, **(**uintptr)(__ccgo_up(bp + 312)), **(**uintptr)(__ccgo_up(bp + 320)))
	if rc != m_SQLITE_OK {
		goto cleanup
	}
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns > 0 {
		s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+13490, libc.VaList(bp+368, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		i2 = 0
		for {
			if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns) {
				break
			}
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+8098, libc.VaList(bp+368, i2))
			goto _3
		_3:
			;
			i2 = i2 + 1
		}
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+13529)
		i3 = 0
		for {
			if !(i3 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns) {
				break
			}
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5252)
			goto _4
		_4:
			;
			i3 = i3 + 1
		}
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5256)
		zSql = libsqlite3.Xsqlite3_str_finish(tls, s)
		// TODO double check error handling ehre
		if !(zSql != 0) {
			rc = int32(m_SQLITE_NOMEM)
			goto cleanup
		}
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+352, libc.UintptrFromInt32(0))
		if rc != m_SQLITE_OK {
			goto cleanup
		}
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 352)), int32(1), **(**Ti64)(__ccgo_up(bp)))
		i4 = 0
		for {
			if !(i4 < Xvec0_num_defined_user_columns(tls, p)) {
				break
			}
			if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i4)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY) {
				goto _5
			}
			auxiliary_key_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i4))))
			v = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i4)*8))
			v_type = libsqlite3.Xsqlite3_value_type(tls, v)
			if v_type != int32(m_SQLITE_NULL) && v_type != (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(auxiliary_key_idx)*24))).Ftype1 {
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 352)))
				rc = int32(m_SQLITE_CONSTRAINT)
				Xvtab_set_error(tls, pVTab, __ccgo_ts+13542, libc.VaList(bp+368, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(auxiliary_key_idx)*24))).Fname_length, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(auxiliary_key_idx)*24))).Fname, Xtype_name(tls, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(auxiliary_key_idx)*24))).Ftype1), Xtype_name(tls, v_type)))
				goto cleanup
			}
			// first 1 is for 1-based indexing on sqlite3_bind_*, second 1 is to account for initial rowid parameter
			libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(bp + 352)), libc.Int32FromInt32(1)+libc.Int32FromInt32(1)+auxiliary_key_idx, v)
			goto _5
		_5:
			;
			i4 = i4 + 1
		}
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 352)))
		if rc != int32(m_SQLITE_DONE) {
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 352)))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 352)))
	}
	i5 = 0
	for {
		if !(i5 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i5)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA) {
			goto _6
		}
		metadata_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i5))))
		v1 = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i5)*8))
		rc = Xvec0_write_metadata_value(tls, p, metadata_idx, **(**Ti64)(__ccgo_up(bp)), **(**Ti64)(__ccgo_up(bp + 296)), **(**Ti64)(__ccgo_up(bp + 304)), v1, 0)
		if rc != m_SQLITE_OK {
			goto cleanup
		}
		goto _6
	_6:
		;
		i5 = i5 + 1
	}
	**(**Tsqlite_int64)(__ccgo_up(pRowid)) = **(**Ti64)(__ccgo_up(bp))
	rc = m_SQLITE_OK
	goto cleanup
cleanup:
	;
	i6 = 0
	for {
		if !(i6 < numReadVectors) {
			break
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(**(**[16]Tvector_cleanup)(__ccgo_up(bp + 136)))[i6]})))(tls, (**(**[16]uintptr)(__ccgo_up(bp + 8)))[i6])
		goto _7
	_7:
		;
		i6 = i6 + 1
	}
	libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 320)))
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 312)))
	if rc == m_SQLITE_OK && brc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+13634, 0)
		return brc
	}
	return rc
}

// C documentation
//
//	/**
//	 * @brief
//	 *
//	 * @param p vec0 virtual table
//	 * @param chunk_rowid: which chunk to write to
//	 * @param chunk_offset: the offset inside the chunk to write the vector to.
//	 * @param rowid: the rowid of the inserting row
//	 * @param vectorDatas: array of the vector data to insert
//	 * @param blobValidity: writeable validity blob of the row's assigned chunk.
//	 * @param validity: snapshot buffer of the valdity column from the row's
//	 * assigned chunk.
//	 * @return int SQLITE_OK on success, error code on failure
//	 */
func Xvec0Update_InsertWriteFinalStep(tls *libc.TLS, p uintptr, chunk_rowid Ti64, chunk_offset Ti64, _rowid Ti64, vectorDatas uintptr, blobChunksValidity uintptr, bufferChunksValidity uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*Ti64)(unsafe.Pointer(bp)) = _rowid
	var actual, actual1, expected, expected1 Ti64
	var brc, i, rc int32
	var _ /* blobChunksRowids at bp+8 */ uintptr
	var _ /* blobVectors at bp+24 */ uintptr
	var _ /* bx at bp+16 */ uint8
	_, _, _, _, _, _, _ = actual, actual1, brc, expected, expected1, i, rc
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	// mark the validity bit for this row in the chunk's validity bitmap
	// Get the byte offset of the bitmap
	**(**uint8)(__ccgo_up(bp + 16)) = **(**uint8)(__ccgo_up(bufferChunksValidity + uintptr(chunk_offset/int64(m___CHAR_BIT__))))
	// set the bit at the chunk_offset position inside that byte
	**(**uint8)(__ccgo_up(bp + 16)) = uint8(int32(**(**uint8)(__ccgo_up(bp + 16))) | int32(1)<<(chunk_offset%int64(m___CHAR_BIT__)))
	// write that 1 byte
	rc = libsqlite3.Xsqlite3_blob_write(tls, blobChunksValidity, bp+16, int32(1), int32(chunk_offset/int64(m___CHAR_BIT__)))
	if rc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+11995, 0)
		return rc
	}
	// Go insert the vector data into the vector chunk shadow tables
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), __ccgo_ts+3712, chunk_rowid, int32(1), bp+24)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+12051, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_rowid))
			goto cleanup
		}
		expected = int64(uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))))
		actual = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp + 24))))
		if actual != expected {
			// IMP: V16386_00456
			Xvtab_set_error(tls, p, __ccgo_ts+12091, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_rowid, expected, actual))
			rc = int32(m_SQLITE_ERROR)
			// already error, can ignore result code
			libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 24)))
			goto cleanup
		}
		rc = _vec0_write_vector_to_vector_blob(tls, **(**uintptr)(__ccgo_up(bp + 24)), chunk_offset, **(**uintptr)(__ccgo_up(vectorDatas + uintptr(i)*8)), (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fdimensions, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Felement_type)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+12186, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_rowid))
			rc = int32(m_SQLITE_ERROR)
			// already error, can ignore result code
			libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 24)))
			goto cleanup
		}
		rc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+12255, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_rowid))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// write the new rowid to the rowids column of the _chunks table
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+9690, chunk_rowid, int32(1), bp+8)
	if rc != m_SQLITE_OK {
		// IMP: V09221_26060
		Xvtab_set_error(tls, p, __ccgo_ts+12324, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_rowid))
		goto cleanup
	}
	expected1 = int64(uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * uint64(8))
	actual1 = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp + 8))))
	if expected1 != actual1 {
		// IMP: V12779_29618
		Xvtab_set_error(tls, p, __ccgo_ts+12392, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_rowid, expected1, actual1))
		rc = int32(m_SQLITE_ERROR)
		goto cleanup
	}
	rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp + 8)), bp, int32(8), int32(uint64(chunk_offset)*uint64(8)))
	if rc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+12487, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_rowid))
		rc = int32(m_SQLITE_ERROR)
		goto cleanup
	}
	// Now with all the vectors inserted, go back and update the _rowids table
	// with the new chunk_rowid/chunk_offset values
	rc = Xvec0_rowids_update_position(tls, p, **(**Ti64)(__ccgo_up(bp)), chunk_rowid, chunk_offset)
	goto cleanup
cleanup:
	;
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	if rc == m_SQLITE_OK && brc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+12556, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_rowid))
		return brc
	}
	return rc
}

func Xvec0Update_Update(tls *libc.TLS, pVTab uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var a, b, p, value, value1, value2, valueVector uintptr
	var auxiliary_column_idx, i, i1, i2, i3, metadata_column_idx, rc, vector_idx int32
	var _ /* chunk_id at bp+0 */ Ti64
	var _ /* chunk_offset at bp+8 */ Ti64
	var _ /* rowid at bp+16 */ Ti64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a, auxiliary_column_idx, b, i, i1, i2, i3, metadata_column_idx, p, rc, value, value1, value2, valueVector, vector_idx
	_ = argc
	p = pVTab
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FpkIsText != 0 {
		a = libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv)))
		b = libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv + 1*8)))
		// IMP: V08886_25725
		if libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv))) != libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv + 1*8))) || libc.Xstrncmp(tls, a, b, uint64(libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv))))) != 0 {
			Xvtab_set_error(tls, pVTab, __ccgo_ts+15180, 0)
			return int32(m_SQLITE_ERROR)
		}
		rc = Xvec0_rowid_from_id(tls, p, **(**uintptr)(__ccgo_up(argv)), bp+16)
		if rc != m_SQLITE_OK {
			return rc
		}
	} else {
		**(**Ti64)(__ccgo_up(bp + 16)) = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv)))
	}
	// 1) get chunk_id and chunk_offset from _rowids
	rc = Xvec0_get_chunk_position(tls, p, **(**Ti64)(__ccgo_up(bp + 16)), libc.UintptrFromInt32(0), bp, bp+8)
	if rc != m_SQLITE_OK {
		return rc
	}
	// 2) update any partition key values
	i = 0
	for {
		if !(i < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION) {
			goto _1
		}
		value = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i)*8))
		if libsqlite3.Xsqlite3_value_nochange(tls, value) != 0 {
			goto _1
		}
		Xvtab_set_error(tls, pVTab, __ccgo_ts+15232, 0)
		return int32(m_SQLITE_ERROR)
		goto _1
	_1:
		;
		i = i + 1
	}
	// 3) handle auxiliary column updates
	i1 = 0
	for {
		if !(i1 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i1)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY) {
			goto _2
		}
		auxiliary_column_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i1))))
		value1 = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i1)*8))
		if libsqlite3.Xsqlite3_value_nochange(tls, value1) != 0 {
			goto _2
		}
		rc = Xvec0Update_UpdateAuxColumn(tls, p, auxiliary_column_idx, value1, **(**Ti64)(__ccgo_up(bp + 16)))
		if rc != m_SQLITE_OK {
			return int32(m_SQLITE_ERROR)
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// 4) handle metadata column updates
	i2 = 0
	for {
		if !(i2 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i2)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA) {
			goto _3
		}
		metadata_column_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i2))))
		value2 = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i2)*8))
		if libsqlite3.Xsqlite3_value_nochange(tls, value2) != 0 {
			goto _3
		}
		rc = Xvec0_write_metadata_value(tls, p, metadata_column_idx, **(**Ti64)(__ccgo_up(bp + 16)), **(**Ti64)(__ccgo_up(bp)), **(**Ti64)(__ccgo_up(bp + 8)), value2, int32(1))
		if rc != m_SQLITE_OK {
			return rc
		}
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	// 5) iterate over all new vectors, update the vectors
	i3 = 0
	for {
		if !(i3 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i3)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR) {
			goto _4
		}
		vector_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i3))))
		valueVector = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i3)*8))
		// in vec0Column, we check sqlite3_vtab_nochange() on vector columns.
		// If the vector column isn't being changed, we return NULL;
		// That's not great, that means vector columns can never be NULLABLE
		// (bc we cant distinguish if an updated vector is truly NULL or nochange).
		// Also it means that if someone tries to run `UPDATE v SET X = NULL`,
		// we can't effectively detect and raise an error.
		// A better solution would be to use a custom result_type for "empty",
		// but subtypes don't appear to survive xColumn -> xUpdate, it's always 0.
		// So for now, we'll just use NULL and warn people to not SET X = NULL
		// in the docs.
		if libsqlite3.Xsqlite3_value_type(tls, valueVector) == int32(m_SQLITE_NULL) {
			goto _4
		}
		rc = Xvec0Update_UpdateVectorColumn(tls, p, **(**Ti64)(__ccgo_up(bp)), **(**Ti64)(__ccgo_up(bp + 8)), vector_idx, valueVector)
		if rc != m_SQLITE_OK {
			return int32(m_SQLITE_ERROR)
		}
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
	return m_SQLITE_OK
}

// C documentation
//
//	/**
//	 * @brief Crete at "iterator" (sqlite3_stmt) of chunks with the given constraints
//	 *
//	 * Any VEC0_IDXSTR_KIND_KNN_PARTITON_CONSTRAINT values in idxStr/argv will be applied
//	 * as WHERE constraints in the underlying stmt SQL, and any consumer of the stmt
//	 * can freely step through the stmt with all constraints satisfied.
//	 *
//	 * @param p - vec0_vtab
//	 * @param idxStr - the xBestIndex/xFilter idxstr containing VEC0_IDXSTR values
//	 * @param argc - number of argv values from xFilter
//	 * @param argv - array of sqlite3_value from xFilter
//	 * @param outStmt - output sqlite3_stmt of chunks with all filters applied
//	 * @return int SQLITE_OK on success, error code otherwise
//	 */
func Xvec0_chunks_iter(tls *libc.TLS, p uintptr, idxStr uintptr, argc int32, argv uintptr, outStmt uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var appendedWhere, i, i1, idx, idx1, idxStrLength, n, numValueEntries, operator, partition_idx, rc, v3 int32
	var kind, kind1 int8
	var s, zSql, zSql1 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = appendedWhere, i, i1, idx, idx1, idxStrLength, kind, kind1, n, numValueEntries, operator, partition_idx, rc, s, zSql, zSql1, v3
	// always null terminated, enforced by SQLite
	idxStrLength = int32(libc.Xstrlen(tls, idxStr))
	// "1" refers to the initial vec0_query_plan char, 4 is the number of chars per "element"
	numValueEntries = (idxStrLength - int32(1)) / int32(4)
	s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
	libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9522, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	appendedWhere = 0
	i = 0
	for {
		if !(i < numValueEntries) {
			break
		}
		idx = int32(1) + i*int32(4)
		kind = **(**int8)(__ccgo_up(idxStr + uintptr(idx+0)))
		if int32(kind) != int32(_VEC0_IDXSTR_KIND_KNN_PARTITON_CONSTRAINT) {
			goto _1
		}
		partition_idx = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx+int32(1))))) - int32('A')
		operator = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx+int32(2)))))
		// idxStr[idx + 3] is just null, a '_' placeholder
		if !(appendedWhere != 0) {
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+9579)
			appendedWhere = int32(1)
		} else {
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+4165)
		}
		switch operator {
		case int32(_VEC0_PARTITION_OPERATOR_EQ):
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+4171, libc.VaList(bp+8, partition_idx))
		case int32(_VEC0_PARTITION_OPERATOR_GT):
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9587, libc.VaList(bp+8, partition_idx))
		case int32(_VEC0_PARTITION_OPERATOR_LE):
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9607, libc.VaList(bp+8, partition_idx))
		case int32(_VEC0_PARTITION_OPERATOR_LT):
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9628, libc.VaList(bp+8, partition_idx))
		case int32(_VEC0_PARTITION_OPERATOR_GE):
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9648, libc.VaList(bp+8, partition_idx))
		case int32(_VEC0_PARTITION_OPERATOR_NE):
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9669, libc.VaList(bp+8, partition_idx))
		default:
			zSql = libsqlite3.Xsqlite3_str_finish(tls, s)
			libsqlite3.Xsqlite3_free(tls, zSql)
			return int32(m_SQLITE_ERROR)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	zSql1 = libsqlite3.Xsqlite3_str_finish(tls, s)
	if !(zSql1 != 0) {
		return int32(m_SQLITE_NOMEM)
	}
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql1, -int32(1), outStmt, libc.UintptrFromInt32(0))
	libsqlite3.Xsqlite3_free(tls, zSql1)
	if rc != m_SQLITE_OK {
		return rc
	}
	n = int32(1)
	i1 = 0
	for {
		if !(i1 < numValueEntries) {
			break
		}
		idx1 = int32(1) + i1*int32(4)
		kind1 = **(**int8)(__ccgo_up(idxStr + uintptr(idx1+0)))
		if int32(kind1) != int32(_VEC0_IDXSTR_KIND_KNN_PARTITON_CONSTRAINT) {
			goto _2
		}
		v3 = n
		n = n + 1
		libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(outStmt)), v3, **(**uintptr)(__ccgo_up(argv + uintptr(i1)*8)))
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	return rc
}

// C documentation
//
//	/**
//	 * Returns the auxiliary column index of the given user column index.
//	 * ONLY call if validated with vec0_column_idx_to_partition_idx before
//	 */
func Xvec0_column_idx_to_auxiliary_idx(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	_ = pVtab
	return int32(**(**Tuint8_t)(__ccgo_up(pVtab + 296 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START)))))
}

// C documentation
//
//	/**
//	 * Returns the metadata column index of the given user column index.
//	 * ONLY call if validated with vec0_column_idx_is_metadata before
//	 */
func Xvec0_column_idx_to_metadata_idx(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	_ = pVtab
	return int32(**(**Tuint8_t)(__ccgo_up(pVtab + 296 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START)))))
}

// C documentation
//
//	/**
//	 * Returns the partition column index of the given user column index.
//	 * ONLY call if validated with vec0_column_idx_is_vector before
//	 */
func Xvec0_column_idx_to_partition_idx(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	_ = pVtab
	return int32(**(**Tuint8_t)(__ccgo_up(pVtab + 296 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START)))))
}

// C documentation
//
//	/**
//	 * Returns the vector index of the given user column index.
//	 * ONLY call if validated with vec0_column_idx_is_vector before
//	 */
func Xvec0_column_idx_to_vector_idx(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	_ = pVtab
	return int32(**(**Tuint8_t)(__ccgo_up(pVtab + 296 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START)))))
}

// C documentation
//
//	/**
//	 * @brief
//	 *
//	 * @param pVtab: virtual table to query
//	 * @param rowid: row to lookup
//	 * @param vector_column_idx: which vector column to query
//	 * @param outVector: Output pointer to the vector buffer.
//	 *                    Must be sqlite3_free()'ed.
//	 * @param outVectorSize: Pointer to a int where the size of outVector
//	 *                       will be stored.
//	 * @return int SQLITE_OK on success.
//	 */
func Xvec0_get_vector_data(tls *libc.TLS, pVtab uintptr, rowid Ti64, vector_column_idx int32, outVector uintptr, outVectorSize uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var blobOffset, brc, rc int32
	var buf, p uintptr
	var size Tsize_t
	var _ /* chunk_id at bp+0 */ Ti64
	var _ /* chunk_offset at bp+8 */ Ti64
	var _ /* vectorBlob at bp+16 */ uintptr
	_, _, _, _, _, _ = blobOffset, brc, buf, p, rc, size
	p = pVtab
	buf = libc.UintptrFromInt32(0)
	**(**uintptr)(__ccgo_up(bp + 16)) = libc.UintptrFromInt32(0)
	rc = Xvec0_get_chunk_position(tls, pVtab, rowid, libc.UintptrFromInt32(0), bp, bp+8)
	if rc == int32(m_SQLITE_EMPTY) {
		Xvtab_set_error(tls, pVtab, __ccgo_ts+3675, libc.VaList(bp+32, rowid))
		goto cleanup
	}
	if rc != m_SQLITE_OK {
		goto cleanup
	}
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(vector_column_idx)*8)), __ccgo_ts+3712, **(**Ti64)(__ccgo_up(bp)), 0, bp+16)
	if rc != m_SQLITE_OK {
		Xvtab_set_error(tls, pVtab, __ccgo_ts+3720, libc.VaList(bp+32, rowid))
		rc = int32(m_SQLITE_ERROR)
		goto cleanup
	}
	size = Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_column_idx)*32)))
	blobOffset = int32(uint64(**(**Ti64)(__ccgo_up(bp + 8))) * size)
	buf = libsqlite3.Xsqlite3_malloc(tls, int32(size))
	if !(buf != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), buf, int32(size), blobOffset)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_free(tls, buf)
		buf = libc.UintptrFromInt32(0)
		Xvtab_set_error(tls, pVtab, __ccgo_ts+3778, libc.VaList(bp+32, rowid))
		rc = int32(m_SQLITE_ERROR)
		goto cleanup
	}
	**(**uintptr)(__ccgo_up(outVector)) = buf
	if outVectorSize != 0 {
		**(**int32)(__ccgo_up(outVectorSize)) = int32(size)
	}
	rc = m_SQLITE_OK
	goto cleanup
cleanup:
	;
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 16)))
	if rc == m_SQLITE_OK && brc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+3841, 0)
		return brc
	}
	return rc
}

func Xvec0_metadata_chunk_size(tls *libc.TLS, kind Tvec0_metadata_column_kind, chunk_size int32) (r int32) {
	switch kind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		return chunk_size / int32(8)
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		return int32(uint64(chunk_size) * uint64(8))
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		return int32(uint64(chunk_size) * uint64(8))
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		return chunk_size * int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)
	}
	return 0
}

func Xvec0_metadata_filter_text(tls *libc.TLS, p uintptr, value uintptr, buffer uintptr, size int32, op Tvec0_metadata_operator, b uintptr, metadata_idx int32, chunk_rowid int32, aMetadataIn uintptr, argv_idx int32) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var aTarget, entry, metadataIn, metadataIn1, rowids, sPrefix, sPrefix1, sTarget, view, view1 uintptr
	var cmpPrefix, cmpPrefix1, cmpPrefix2, cmpPrefix3, cmpPrefix4, cmpPrefix5, cmpPrefix6, i, i1, i2, i3, i4, i5, i7, nPrefix, nPrefix1, nTarget, rc, v10, v12, v14 int32
	var i6, metadataInIdx, target_idx Tsize_t
	var _ /* nFull at bp+24 */ int32
	var _ /* nFull at bp+40 */ int32
	var _ /* rowidsBlob at bp+8 */ uintptr
	var _ /* sFull at bp+16 */ uintptr
	var _ /* sFull at bp+32 */ uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aTarget, cmpPrefix, cmpPrefix1, cmpPrefix2, cmpPrefix3, cmpPrefix4, cmpPrefix5, cmpPrefix6, entry, i, i1, i2, i3, i4, i5, i6, i7, metadataIn, metadataIn1, metadataInIdx, nPrefix, nPrefix1, nTarget, rc, rowids, sPrefix, sPrefix1, sTarget, target_idx, view, view1, v10, v12, v14
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	rowids = libc.UintptrFromInt32(0)
	sTarget = libsqlite3.Xsqlite3_value_text(tls, value)
	nTarget = libsqlite3.Xsqlite3_value_bytes(tls, value)
	// TODO(perf): only text metadata news the rowids BLOB. Make it so that
	// rowids BLOB is re-used when multiple fitlers on text columns,
	// ex "name BETWEEN 'a' and 'b'""
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+9690, int64(chunk_rowid), 0, bp+8)
	if rc != m_SQLITE_OK {
		return rc
	}
	rowids = libsqlite3.Xsqlite3_malloc(tls, libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp + 8))))
	if !(rowids != 0) {
		libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		return int32(m_SQLITE_NOMEM)
	}
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 8)), rowids, libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp + 8))), 0)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		return rc
	}
	libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	switch op {
	case int32(_VEC0_METADATA_OPERATOR_EQ):
		goto _1
	case int32(_VEC0_METADATA_OPERATOR_NE):
		goto _2
	case int32(_VEC0_METADATA_OPERATOR_GT):
		goto _3
	case int32(_VEC0_METADATA_OPERATOR_GE):
		goto _4
	case int32(_VEC0_METADATA_OPERATOR_LE):
		goto _5
	case int32(_VEC0_METADATA_OPERATOR_LT):
		goto _6
	case int32(_VEC0_METADATA_OPERATOR_IN):
		goto _7
	}
	goto _8
_1:
	;
	i = 0
	for {
		if !(i < size) {
			break
		}
		view = buffer + uintptr(i*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		nPrefix = **(**int32)(__ccgo_up(view))
		sPrefix = view + 4
		// for EQ the text lengths must match
		if nPrefix != nTarget {
			Xbitmap_set(tls, b, i, 0)
			goto _9
		}
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			v10 = nPrefix
		} else {
			v10 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
		}
		cmpPrefix = libc.Xstrncmp(tls, sPrefix, sTarget, uint64(v10))
		// for short strings, use the prefix comparison direclty
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			Xbitmap_set(tls, b, i, libc.BoolInt32(cmpPrefix == 0))
			goto _9
		}
		// for EQ on longs strings, the prefix must match
		if cmpPrefix != 0 {
			Xbitmap_set(tls, b, i, 0)
			goto _9
		}
		// consult the full string
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i)*8)), bp+24, bp+16)
		if rc != m_SQLITE_OK {
			goto done
		}
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		Xbitmap_set(tls, b, i, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, uint64(**(**int32)(__ccgo_up(bp + 24)))) == 0))
		goto _9
	_9:
		;
		i = i + 1
	}
	goto _8
_2:
	;
	i1 = 0
	for {
		if !(i1 < size) {
			break
		}
		view = buffer + uintptr(i1*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		nPrefix = **(**int32)(__ccgo_up(view))
		sPrefix = view + 4
		// for NE if text lengths dont match, it never will
		if nPrefix != nTarget {
			Xbitmap_set(tls, b, i1, int32(1))
			goto _11
		}
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			v10 = nPrefix
		} else {
			v10 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
		}
		cmpPrefix1 = libc.Xstrncmp(tls, sPrefix, sTarget, uint64(v10))
		// for short strings, use the prefix comparison direclty
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			Xbitmap_set(tls, b, i1, libc.BoolInt32(cmpPrefix1 != 0))
			goto _11
		}
		// for NE on longs strings, if prefixes dont match, then long string wont
		if cmpPrefix1 != 0 {
			Xbitmap_set(tls, b, i1, int32(1))
			goto _11
		}
		// consult the full string
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i1)*8)), bp+24, bp+16)
		if rc != m_SQLITE_OK {
			goto done
		}
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		Xbitmap_set(tls, b, i1, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, uint64(**(**int32)(__ccgo_up(bp + 24)))) != 0))
		goto _11
	_11:
		;
		i1 = i1 + 1
	}
	goto _8
_3:
	;
	i2 = 0
	for {
		if !(i2 < size) {
			break
		}
		view = buffer + uintptr(i2*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		nPrefix = **(**int32)(__ccgo_up(view))
		sPrefix = view + 4
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			v12 = nPrefix
		} else {
			v12 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
		}
		if v12 <= nTarget {
			if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				v14 = nPrefix
			} else {
				v14 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
			}
			v10 = v14
		} else {
			v10 = nTarget
		}
		cmpPrefix2 = libc.Xstrncmp(tls, sPrefix, sTarget, uint64(v10))
		if nPrefix < int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// if prefix match, check which is longer
			if cmpPrefix2 == 0 {
				Xbitmap_set(tls, b, i2, libc.BoolInt32(nPrefix > nTarget))
			} else {
				Xbitmap_set(tls, b, i2, libc.BoolInt32(cmpPrefix2 > 0))
			}
			goto _13
		}
		// TODO(perf): may not need to compare full text in some cases
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i2)*8)), bp+24, bp+16)
		if rc != m_SQLITE_OK {
			goto done
		}
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		Xbitmap_set(tls, b, i2, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, uint64(**(**int32)(__ccgo_up(bp + 24)))) > 0))
		goto _13
	_13:
		;
		i2 = i2 + 1
	}
	goto _8
_4:
	;
	i3 = 0
	for {
		if !(i3 < size) {
			break
		}
		view = buffer + uintptr(i3*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		nPrefix = **(**int32)(__ccgo_up(view))
		sPrefix = view + 4
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			v12 = nPrefix
		} else {
			v12 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
		}
		if v12 <= nTarget {
			if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				v14 = nPrefix
			} else {
				v14 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
			}
			v10 = v14
		} else {
			v10 = nTarget
		}
		cmpPrefix3 = libc.Xstrncmp(tls, sPrefix, sTarget, uint64(v10))
		if nPrefix < int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// if prefix match, check which is longer
			if cmpPrefix3 == 0 {
				Xbitmap_set(tls, b, i3, libc.BoolInt32(nPrefix >= nTarget))
			} else {
				Xbitmap_set(tls, b, i3, libc.BoolInt32(cmpPrefix3 >= 0))
			}
			goto _17
		}
		// TODO(perf): may not need to compare full text in some cases
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i3)*8)), bp+24, bp+16)
		if rc != m_SQLITE_OK {
			goto done
		}
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		Xbitmap_set(tls, b, i3, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, uint64(**(**int32)(__ccgo_up(bp + 24)))) >= 0))
		goto _17
	_17:
		;
		i3 = i3 + 1
	}
	goto _8
_5:
	;
	i4 = 0
	for {
		if !(i4 < size) {
			break
		}
		view = buffer + uintptr(i4*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		nPrefix = **(**int32)(__ccgo_up(view))
		sPrefix = view + 4
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			v12 = nPrefix
		} else {
			v12 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
		}
		if v12 <= nTarget {
			if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				v14 = nPrefix
			} else {
				v14 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
			}
			v10 = v14
		} else {
			v10 = nTarget
		}
		cmpPrefix4 = libc.Xstrncmp(tls, sPrefix, sTarget, uint64(v10))
		if nPrefix < int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// if prefix match, check which is longer
			if cmpPrefix4 == 0 {
				Xbitmap_set(tls, b, i4, libc.BoolInt32(nPrefix <= nTarget))
			} else {
				Xbitmap_set(tls, b, i4, libc.BoolInt32(cmpPrefix4 <= 0))
			}
			goto _21
		}
		// TODO(perf): may not need to compare full text in some cases
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i4)*8)), bp+24, bp+16)
		if rc != m_SQLITE_OK {
			goto done
		}
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		Xbitmap_set(tls, b, i4, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, uint64(**(**int32)(__ccgo_up(bp + 24)))) <= 0))
		goto _21
	_21:
		;
		i4 = i4 + 1
	}
	goto _8
_6:
	;
	i5 = 0
	for {
		if !(i5 < size) {
			break
		}
		view = buffer + uintptr(i5*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		nPrefix = **(**int32)(__ccgo_up(view))
		sPrefix = view + 4
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			v12 = nPrefix
		} else {
			v12 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
		}
		if v12 <= nTarget {
			if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				v14 = nPrefix
			} else {
				v14 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
			}
			v10 = v14
		} else {
			v10 = nTarget
		}
		cmpPrefix5 = libc.Xstrncmp(tls, sPrefix, sTarget, uint64(v10))
		if nPrefix < int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// if prefix match, check which is longer
			if cmpPrefix5 == 0 {
				Xbitmap_set(tls, b, i5, libc.BoolInt32(nPrefix < nTarget))
			} else {
				Xbitmap_set(tls, b, i5, libc.BoolInt32(cmpPrefix5 < 0))
			}
			goto _25
		}
		// TODO(perf): may not need to compare full text in some cases
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i5)*8)), bp+24, bp+16)
		if rc != m_SQLITE_OK {
			goto done
		}
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		Xbitmap_set(tls, b, i5, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, uint64(**(**int32)(__ccgo_up(bp + 24)))) < 0))
		goto _25
	_25:
		;
		i5 = i5 + 1
	}
	goto _8
_7:
	;
	metadataInIdx = uint64(-libc.Int32FromInt32(1))
	i6 = uint64(0)
	for {
		if !(i6 < (*TArray)(unsafe.Pointer(aMetadataIn)).Flength) {
			break
		}
		metadataIn = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(i6)*40
		if (*TVec0MetadataIn)(unsafe.Pointer(metadataIn)).Fargv_idx == argv_idx {
			metadataInIdx = i6
			break
		}
		goto _29
	_29:
		;
		i6 = i6 + 1
	}
	if metadataInIdx < uint64(0) {
		rc = int32(m_SQLITE_ERROR)
		goto done
	}
	metadataIn1 = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(metadataInIdx)*40
	aTarget = metadataIn1 + 8
	i7 = 0
	for {
		if !(i7 < size) {
			break
		}
		view1 = buffer + uintptr(i7*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		nPrefix1 = **(**int32)(__ccgo_up(view1))
		sPrefix1 = view1 + 4
		target_idx = uint64(0)
		for {
			if !(target_idx < (*TArray)(unsafe.Pointer(aTarget)).Flength) {
				break
			}
			entry = (*TArray)(unsafe.Pointer(aTarget)).Fz + uintptr(target_idx)*16
			if (*TVec0MetadataInTextEntry)(unsafe.Pointer(entry)).Fn != nPrefix1 {
				goto _31
			}
			if nPrefix1 <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				v10 = nPrefix1
			} else {
				v10 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
			}
			cmpPrefix6 = libc.Xstrncmp(tls, sPrefix1, (*TVec0MetadataInTextEntry)(unsafe.Pointer(entry)).FzString, uint64(v10))
			if nPrefix1 <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				if cmpPrefix6 == 0 {
					Xbitmap_set(tls, b, i7, int32(1))
					break
				}
				goto _31
			}
			if cmpPrefix6 != 0 {
				goto _31
			}
			rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i7)*8)), bp+40, bp+32)
			if rc != m_SQLITE_OK {
				goto done
			}
			if nPrefix1 != **(**int32)(__ccgo_up(bp + 40)) {
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
			if libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 32)), (*TVec0MetadataInTextEntry)(unsafe.Pointer(entry)).FzString, uint64(**(**int32)(__ccgo_up(bp + 40)))) == 0 {
				Xbitmap_set(tls, b, i7, int32(1))
				break
			}
			goto _31
		_31:
			;
			target_idx = target_idx + 1
		}
		goto _30
	_30:
		;
		i7 = i7 + 1
	}
	goto _8
_8:
	;
	rc = m_SQLITE_OK
	goto done
done:
	;
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	libsqlite3.Xsqlite3_free(tls, rowids)
	return rc
}

// C documentation
//
//	/**
//	 * @brief Adds a new chunk for the vec0 table, and the corresponding vector
//	 * chunks.
//	 *
//	 * Inserts a new row into the _chunks table, with blank data, and uses that new
//	 * rowid to insert new blank rows into _vector_chunksXX tables.
//	 *
//	 * @param p: vec0 table to add new chunk
//	 * @param paritionKeyValues: Array of partition key valeus for the new chunk, if available
//	 * @param chunk_rowid: Output pointer, if not NULL, then will be filled with the
//	 * new chunk rowid.
//	 * @return int SQLITE_OK on success, error code otherwise.
//	 */
func Xvec0_new_chunk(tls *libc.TLS, p uintptr, partitionKeyValues uintptr, chunk_rowid uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var failed, i, i1, i2, i3, i4, metadata_column_idx, rc, vector_column_idx int32
	var rowid, vectorsSize Ti64
	var s, zSql uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = failed, i, i1, i2, i3, i4, metadata_column_idx, rc, rowid, s, vector_column_idx, vectorsSize, zSql
	// Step 1: Insert a new row in _chunks, capture that new rowid
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns > 0 {
		s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+5165, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5194)
		i = 0
		for {
			if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
				break
			}
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+5218, libc.VaList(bp+16, i))
			goto _1
		_1:
			;
			i = i + 1
		}
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5234)
		i1 = 0
		for {
			if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
				break
			}
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5252)
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5256)
		zSql = libsqlite3.Xsqlite3_str_finish(tls, s)
	} else {
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5258, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	}
	if !(zSql != 0) {
		return int32(m_SQLITE_NOMEM)
	}
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
	libsqlite3.Xsqlite3_free(tls, zSql)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		return rc
	}
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size))                      // size
	libsqlite3.Xsqlite3_bind_zeroblob(tls, **(**uintptr)(__ccgo_up(bp)), int32(2), (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT__))    // validity bitmap
	libsqlite3.Xsqlite3_bind_zeroblob(tls, **(**uintptr)(__ccgo_up(bp)), int32(3), int32(uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8))) // rowids
	i2 = 0
	for {
		if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
			break
		}
		libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(bp)), int32(4)+i2, **(**uintptr)(__ccgo_up(partitionKeyValues + uintptr(i2)*8)))
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
	failed = libc.BoolInt32(rc != int32(m_SQLITE_DONE))
	rowid = libsqlite3.Xsqlite3_last_insert_rowid(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb)
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	if failed != 0 {
		return int32(m_SQLITE_ERROR)
	}
	// Step 2: Create new vector chunks for each vector column, with
	//          that new chunk_rowid.
	//
	// SHADOW_TABLE_ROWID_QUIRK: The _vector_chunksNN and _metadatachunksNN
	// shadow tables declare "rowid PRIMARY KEY" without the INTEGER type, so
	// the user-defined "rowid" column is NOT an alias for the internal SQLite
	// rowid (_rowid_). When only appending rows these two happen to stay in
	// sync, but after a chunk is deleted (vec0Update_Delete_DeleteChunkIfEmpty)
	// and a new one is created, the auto-assigned _rowid_ can diverge from the
	// user "rowid" value. Since sqlite3_blob_open() addresses rows by internal
	// _rowid_, we must explicitly set BOTH _rowid_ and "rowid" to the same
	// value so that later blob operations can find the row.
	//
	// The correct long-term fix is changing the schema to
	//   "rowid INTEGER PRIMARY KEY"
	// which makes it a true alias, but that would break existing databases.
	i3 = 0
	for {
		if !(i3 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i3)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR) {
			goto _4
		}
		vector_column_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i3))))
		vectorsSize = int64(uint64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))))
		// See SHADOW_TABLE_ROWID_QUIRK above for why _rowid_ and rowid are both set.
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5329, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, vector_column_idx))
		if !(zSql != 0) {
			return int32(m_SQLITE_NOMEM)
		}
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_free(tls, zSql)
		if rc != m_SQLITE_OK {
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
			return rc
		}
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), rowid) // _rowid_ (internal SQLite rowid)
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(2), rowid) // rowid   (user-defined column)
		libsqlite3.Xsqlite3_bind_zeroblob64(tls, **(**uintptr)(__ccgo_up(bp)), int32(3), uint64(vectorsSize))
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		if rc != int32(m_SQLITE_DONE) {
			return rc
		}
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
	// Step 3: Create new metadata chunks for each metadata column
	i4 = 0
	for {
		if !(i4 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i4)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA) {
			goto _5
		}
		metadata_column_idx = int32(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i4))))
		// See SHADOW_TABLE_ROWID_QUIRK above for why _rowid_ and rowid are both set.
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5410, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
		if !(zSql != 0) {
			return int32(m_SQLITE_NOMEM)
		}
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_free(tls, zSql)
		if rc != m_SQLITE_OK {
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
			return rc
		}
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), rowid) // _rowid_ (internal SQLite rowid)
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(2), rowid) // rowid   (user-defined column)
		libsqlite3.Xsqlite3_bind_zeroblob64(tls, **(**uintptr)(__ccgo_up(bp)), int32(3), uint64(Xvec0_metadata_chunk_size(tls, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_column_idx)*24))).Fkind, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)))
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		if rc != int32(m_SQLITE_DONE) {
			return rc
		}
		goto _5
	_5:
		;
		i4 = i4 + 1
	}
	if chunk_rowid != 0 {
		**(**Ti64)(__ccgo_up(chunk_rowid)) = rowid
	}
	return m_SQLITE_OK
}

// C documentation
//
//	/**
//	 * @brief Parse an vec0 vtab argv[i] column definition and see if
//	 * it's a vector column defintion, ex `contents_embedding float[768]`.
//	 *
//	 * @param source vec0 argv[i] item
//	 * @param source_length length of source in bytes
//	 * @param outColumn Output the parse vector column to this struct, if success
//	 * @return int SQLITE_OK on success, SQLITE_EMPTY is it's not a vector column
//	 * definition, SQLITE_ERROR on error.
//	 */
func Xvec0_parse_vector_column(tls *libc.TLS, source uintptr, source_length int32, outColumn uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var dimensions, keyLength, nameLength, rc, valueLength int32
	var distanceMetric _Vec0DistanceMetrics
	var elementType _VectorElementType
	var key, name, value uintptr
	var _ /* scanner at bp+0 */ TVec0Scanner
	var _ /* token at bp+24 */ TVec0Token
	_, _, _, _, _, _, _, _, _, _ = dimensions, distanceMetric, elementType, key, keyLength, name, nameLength, rc, value, valueLength
	distanceMetric = int32(_VEC0_DISTANCE_METRIC_L2)
	Xvec0_scanner_init(tls, bp, source, source_length)
	// starts with an identifier
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		return int32(m_SQLITE_EMPTY)
	}
	name = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	nameLength = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// vector column type comes next: float, int, or bit
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		return int32(m_SQLITE_EMPTY)
	}
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1771, int32(5)) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1838, int32(3)) == 0 {
		elementType = int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32)
	} else {
		if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+8, int32(4)) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1842, int32(2)) == 0 {
			elementType = int32(_SQLITE_VEC_ELEMENT_TYPE_INT8)
		} else {
			if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+13, int32(3)) == 0 {
				elementType = int32(_SQLITE_VEC_ELEMENT_TYPE_BIT)
			} else {
				return int32(m_SQLITE_EMPTY)
			}
		}
	}
	// left '[' bracket
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_LBRACKET) {
		return int32(m_SQLITE_EMPTY)
	}
	// digit, for vector dimension length
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_DIGIT) {
		return int32(m_SQLITE_ERROR)
	}
	dimensions = libc.Xatoi(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart)
	if dimensions <= 0 {
		return int32(m_SQLITE_ERROR)
	}
	// // right ']' bracket
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_RBRACKET) {
		return int32(m_SQLITE_ERROR)
	}
	// any other tokens left should be column-level options , ex `key=value`
	// ex `distance_metric=L2 distance_metric=cosine` should error
	for int32(1) != 0 {
		// should be EOF or identifier (option key)
		rc = Xvec0_scanner_next(tls, bp, bp+24)
		if rc == int32(m_VEC0_TOKEN_RESULT_EOF) {
			break
		}
		if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
			return int32(m_SQLITE_ERROR)
		}
		key = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
		keyLength = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
		if libsqlite3.Xsqlite3_strnicmp(tls, key, __ccgo_ts+1845, keyLength) == 0 {
			if elementType == int32(_SQLITE_VEC_ELEMENT_TYPE_BIT) {
				return int32(m_SQLITE_ERROR)
			}
			// ensure equal sign after distance_metric
			rc = Xvec0_scanner_next(tls, bp, bp+24)
			if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_EQ) {
				return int32(m_SQLITE_ERROR)
			}
			// distance_metric value, an identifier (L2, cosine, etc)
			rc = Xvec0_scanner_next(tls, bp, bp+24)
			if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
				return int32(m_SQLITE_ERROR)
			}
			value = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
			valueLength = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
			if libsqlite3.Xsqlite3_strnicmp(tls, value, __ccgo_ts+1861, valueLength) == 0 {
				distanceMetric = int32(_VEC0_DISTANCE_METRIC_L2)
			} else {
				if libsqlite3.Xsqlite3_strnicmp(tls, value, __ccgo_ts+1864, valueLength) == 0 {
					distanceMetric = int32(_VEC0_DISTANCE_METRIC_L1)
				} else {
					if libsqlite3.Xsqlite3_strnicmp(tls, value, __ccgo_ts+1867, valueLength) == 0 {
						distanceMetric = int32(_VEC0_DISTANCE_METRIC_COSINE)
					} else {
						return int32(m_SQLITE_ERROR)
					}
				}
			}
		} else {
			return int32(m_SQLITE_ERROR)
		}
	}
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+56, nameLength, name))
	if !((*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fname != 0) {
		return int32(m_SQLITE_ERROR)
	}
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fname_length = nameLength
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fdistance_metric = distanceMetric
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Felement_type = elementType
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fdimensions = uint64(dimensions)
	return m_SQLITE_OK
}

// C documentation
//
//	/**
//	 * @brief Result the given metadata value for the given row and metadata column index.
//	 * Will traverse the metadatachunksNN table with BLOB I/0 for the given rowid.
//	 *
//	 * @param p
//	 * @param rowid
//	 * @param metadata_idx
//	 * @param context
//	 * @return int
//	 */
func Xvec0_result_metadata_value_for_rowid(tls *libc.TLS, p uintptr, rowid Ti64, metadata_idx int32, context uintptr) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var length, rc, value int32
	var zSql uintptr
	var _ /* blobValue at bp+16 */ uintptr
	var _ /* block at bp+24 */ Tu8
	var _ /* chunk_id at bp+0 */ Ti64
	var _ /* chunk_offset at bp+8 */ Ti64
	var _ /* stmt at bp+64 */ uintptr
	var _ /* value at bp+32 */ Ti64
	var _ /* value at bp+40 */ float64
	var _ /* view at bp+48 */ [16]Tu8
	_, _, _, _ = length, rc, value, zSql
	rc = Xvec0_get_chunk_position(tls, p, rowid, libc.UintptrFromInt32(0), bp, bp+8)
	if rc != m_SQLITE_OK {
		return rc
	}
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 480 + uintptr(metadata_idx)*8)), __ccgo_ts+4053, **(**Ti64)(__ccgo_up(bp)), 0, bp+16)
	if rc != m_SQLITE_OK {
		return rc
	}
	switch (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), bp+24, int32(1), int32(**(**Ti64)(__ccgo_up(bp + 8))/int64(m___CHAR_BIT__)))
		if rc != m_SQLITE_OK {
			goto done
		}
		value = int32(**(**Tu8)(__ccgo_up(bp + 24))) >> (**(**Ti64)(__ccgo_up(bp + 8)) % int64(m___CHAR_BIT__)) & int32(1)
		libsqlite3.Xsqlite3_result_int(tls, context, value)
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), bp+32, int32(8), int32(uint64(**(**Ti64)(__ccgo_up(bp + 8)))*uint64(8)))
		if rc != m_SQLITE_OK {
			goto done
		}
		libsqlite3.Xsqlite3_result_int64(tls, context, **(**Ti64)(__ccgo_up(bp + 32)))
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), bp+40, int32(8), int32(uint64(**(**Ti64)(__ccgo_up(bp + 8)))*uint64(8)))
		if rc != m_SQLITE_OK {
			goto done
		}
		libsqlite3.Xsqlite3_result_double(tls, context, **(**float64)(__ccgo_up(bp + 40)))
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), bp+48, int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH), int32(**(**Ti64)(__ccgo_up(bp + 8))*int64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if rc != m_SQLITE_OK {
			goto done
		}
		length = **(**int32)(__ccgo_up(bp + 48))
		if length <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			libsqlite3.Xsqlite3_result_text(tls, context, bp+48+libc.UintptrFromInt32(4), length, uintptr(-libc.Int32FromInt32(1)))
		} else {
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+4058, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_idx))
			if !(zSql != 0) {
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
			rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+64, libc.UintptrFromInt32(0))
			libsqlite3.Xsqlite3_free(tls, zSql)
			if rc != m_SQLITE_OK {
				goto done
			}
			libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 64)), int32(1), rowid)
			rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 64)))
			if rc != int32(m_SQLITE_ROW) {
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 64)))
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
			libsqlite3.Xsqlite3_result_value(tls, context, libsqlite3.Xsqlite3_column_value(tls, **(**uintptr)(__ccgo_up(bp + 64)), 0))
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 64)))
			rc = m_SQLITE_OK
		}
		break
	}
	goto done
done:
	;
	// blobValue is read-only, will not fail on close
	libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 16)))
	return rc
}

// C documentation
//
//	/**
//	 * @brief Fill in bitmap of chunk values, whether or not the values match a metadata constraint
//	 *
//	 * @param p vec0_vtab
//	 * @param metadata_idx index of the metatadata column to perfrom constraints on
//	 * @param value sqlite3_value of the constraints value
//	 * @param blob sqlite3_blob that is already opened on the metdata column's shadow chunk table
//	 * @param chunk_rowid rowid of the chunk to calculate on
//	 * @param b pre-allocated and zero'd out bitmap to write results to
//	 * @param size size of the chunk
//	 * @return int SQLITE_OK on success, error code otherwise
//	 */
func Xvec0_set_metadata_filter_bitmap(tls *libc.TLS, p uintptr, metadata_idx int32, op Tvec0_metadata_operator, value uintptr, blob uintptr, chunk_rowid Ti64, b uintptr, size int32, aMetadataIn uintptr, argv_idx int32) (r int32) {
	var aTarget, array, array1, buffer, metadataIn, metadataIn1 uintptr
	var blobSize, i, i1, i10, i11, i12, i13, i14, i15, i2, i3, i4, i5, i6, i7, i9, metadataInIdx, rc, szMatch, target int32
	var i8, target_idx Tsize_t
	var kind Tvec0_metadata_column_kind
	var target1 Ti64
	var target2 float64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aTarget, array, array1, blobSize, buffer, i, i1, i10, i11, i12, i13, i14, i15, i2, i3, i4, i5, i6, i7, i8, i9, kind, metadataIn, metadataIn1, metadataInIdx, rc, szMatch, target, target1, target2, target_idx
	rc = libsqlite3.Xsqlite3_blob_reopen(tls, blob, chunk_rowid)
	if rc != m_SQLITE_OK {
		return rc
	}
	kind = (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind
	szMatch = 0
	blobSize = libsqlite3.Xsqlite3_blob_bytes(tls, blob)
	switch kind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		szMatch = libc.BoolInt32(blobSize == size/int32(m___CHAR_BIT__))
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		szMatch = libc.BoolInt32(uint64(blobSize) == uint64(size)*uint64(8))
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		szMatch = libc.BoolInt32(uint64(blobSize) == uint64(size)*uint64(8))
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		szMatch = libc.BoolInt32(blobSize == size*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		break
	}
	if !(szMatch != 0) {
		return int32(m_SQLITE_ERROR)
	}
	buffer = libsqlite3.Xsqlite3_malloc(tls, blobSize)
	if !(buffer != 0) {
		return int32(m_SQLITE_NOMEM)
	}
	rc = libsqlite3.Xsqlite3_blob_read(tls, blob, buffer, blobSize, 0)
	if rc != m_SQLITE_OK {
		goto done
	}
	switch kind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		target = libsqlite3.Xsqlite3_value_int(tls, value)
		if target != 0 && op == int32(_VEC0_METADATA_OPERATOR_EQ) || !(target != 0) && op == int32(_VEC0_METADATA_OPERATOR_NE) {
			i = 0
			for {
				if !(i < size) {
					break
				}
				Xbitmap_set(tls, b, i, Xbitmap_get(tls, buffer, i))
				goto _1
			_1:
				;
				i = i + 1
			}
		} else {
			i1 = 0
			for {
				if !(i1 < size) {
					break
				}
				Xbitmap_set(tls, b, i1, libc.BoolInt32(!(Xbitmap_get(tls, buffer, i1) != 0)))
				goto _2
			_2:
				;
				i1 = i1 + 1
			}
		}
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		array = buffer
		target1 = libsqlite3.Xsqlite3_value_int64(tls, value)
		switch op {
		case int32(_VEC0_METADATA_OPERATOR_EQ):
			i2 = 0
			for {
				if !(i2 < size) {
					break
				}
				Xbitmap_set(tls, b, i2, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i2)*8)) == target1))
				goto _3
			_3:
				;
				i2 = i2 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_GT):
			i3 = 0
			for {
				if !(i3 < size) {
					break
				}
				Xbitmap_set(tls, b, i3, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i3)*8)) > target1))
				goto _4
			_4:
				;
				i3 = i3 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_LE):
			i4 = 0
			for {
				if !(i4 < size) {
					break
				}
				Xbitmap_set(tls, b, i4, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i4)*8)) <= target1))
				goto _5
			_5:
				;
				i4 = i4 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_LT):
			i5 = 0
			for {
				if !(i5 < size) {
					break
				}
				Xbitmap_set(tls, b, i5, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i5)*8)) < target1))
				goto _6
			_6:
				;
				i5 = i5 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_GE):
			i6 = 0
			for {
				if !(i6 < size) {
					break
				}
				Xbitmap_set(tls, b, i6, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i6)*8)) >= target1))
				goto _7
			_7:
				;
				i6 = i6 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_NE):
			i7 = 0
			for {
				if !(i7 < size) {
					break
				}
				Xbitmap_set(tls, b, i7, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i7)*8)) != target1))
				goto _8
			_8:
				;
				i7 = i7 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_IN):
			metadataInIdx = -int32(1)
			i8 = uint64(0)
			for {
				if !(i8 < (*TArray)(unsafe.Pointer(aMetadataIn)).Flength) {
					break
				}
				metadataIn = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(i8)*40
				if (*TVec0MetadataIn)(unsafe.Pointer(metadataIn)).Fargv_idx == argv_idx {
					metadataInIdx = int32(i8)
					break
				}
				goto _9
			_9:
				;
				i8 = i8 + 1
			}
			if metadataInIdx < 0 {
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
			metadataIn1 = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(metadataInIdx)*40
			aTarget = metadataIn1 + 8
			i9 = 0
			for {
				if !(i9 < size) {
					break
				}
				target_idx = uint64(0)
				for {
					if !(target_idx < (*TArray)(unsafe.Pointer(aTarget)).Flength) {
						break
					}
					if **(**Ti64)(__ccgo_up((*TArray)(unsafe.Pointer(aTarget)).Fz + uintptr(target_idx)*8)) == **(**Ti64)(__ccgo_up(array + uintptr(i9)*8)) {
						Xbitmap_set(tls, b, i9, int32(1))
						break
					}
					goto _11
				_11:
					;
					target_idx = target_idx + 1
				}
				goto _10
			_10:
				;
				i9 = i9 + 1
			}
			break
		}
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		array1 = buffer
		target2 = libsqlite3.Xsqlite3_value_double(tls, value)
		switch op {
		case int32(_VEC0_METADATA_OPERATOR_EQ):
			i10 = 0
			for {
				if !(i10 < size) {
					break
				}
				Xbitmap_set(tls, b, i10, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i10)*8)) == target2))
				goto _12
			_12:
				;
				i10 = i10 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_GT):
			i11 = 0
			for {
				if !(i11 < size) {
					break
				}
				Xbitmap_set(tls, b, i11, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i11)*8)) > target2))
				goto _13
			_13:
				;
				i11 = i11 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_LE):
			i12 = 0
			for {
				if !(i12 < size) {
					break
				}
				Xbitmap_set(tls, b, i12, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i12)*8)) <= target2))
				goto _14
			_14:
				;
				i12 = i12 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_LT):
			i13 = 0
			for {
				if !(i13 < size) {
					break
				}
				Xbitmap_set(tls, b, i13, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i13)*8)) < target2))
				goto _15
			_15:
				;
				i13 = i13 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_GE):
			i14 = 0
			for {
				if !(i14 < size) {
					break
				}
				Xbitmap_set(tls, b, i14, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i14)*8)) >= target2))
				goto _16
			_16:
				;
				i14 = i14 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_NE):
			i15 = 0
			for {
				if !(i15 < size) {
					break
				}
				Xbitmap_set(tls, b, i15, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i15)*8)) != target2))
				goto _17
			_17:
				;
				i15 = i15 + 1
			}
		case int32(_VEC0_METADATA_OPERATOR_IN):
			// should never be reached
			break
		}
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		rc = Xvec0_metadata_filter_text(tls, p, value, buffer, size, op, b, metadata_idx, int32(chunk_rowid), aMetadataIn, argv_idx)
		if rc != m_SQLITE_OK {
			goto done
		}
		break
	}
	goto done
done:
	;
	libsqlite3.Xsqlite3_free(tls, buffer)
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
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 480 + uintptr(metadata_column_idx)*8)), __ccgo_ts+4053, chunk_id, int32(1), bp)
	if rc != m_SQLITE_OK {
		goto done
	}
	switch kind {
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		value = libsqlite3.Xsqlite3_value_int(tls, v)
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), int32(chunk_offset/libc.Int64FromInt32(m___CHAR_BIT__)))
		if rc != m_SQLITE_OK {
			goto done
		}
		if value != 0 {
			**(**Tu8)(__ccgo_up(bp + 8)) = uint8(int32(**(**Tu8)(__ccgo_up(bp + 8))) | libc.Int32FromInt32(1)<<(chunk_offset%libc.Int64FromInt32(m___CHAR_BIT__)))
		} else {
			**(**Tu8)(__ccgo_up(bp + 8)) = uint8(int32(**(**Tu8)(__ccgo_up(bp + 8))) & ^(libc.Int32FromInt32(1) << (chunk_offset % libc.Int64FromInt32(m___CHAR_BIT__))))
		}
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), int32(chunk_offset/int64(m___CHAR_BIT__)))
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		**(**Ti64)(__ccgo_up(bp + 16)) = libsqlite3.Xsqlite3_value_int64(tls, v)
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+16, int32(8), int32(uint64(chunk_offset)*uint64(8)))
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		**(**float64)(__ccgo_up(bp + 24)) = libsqlite3.Xsqlite3_value_double(tls, v)
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+24, int32(8), int32(uint64(chunk_offset)*uint64(8)))
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+32, int32(4), int32(chunk_offset*int64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if rc != m_SQLITE_OK {
			goto done
		}
		s = libsqlite3.Xsqlite3_value_text(tls, v)
		**(**int32)(__ccgo_up(bp + 36)) = libsqlite3.Xsqlite3_value_bytes(tls, v)
		libc.Xmemset(tls, bp+40, 0, uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		libc.Xmemcpy(tls, bp+40, bp+36, uint64(4))
		if **(**int32)(__ccgo_up(bp + 36)) <= libc.Int32FromInt32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)-libc.Int32FromInt32(4) {
			v1 = **(**int32)(__ccgo_up(bp + 36))
		} else {
			v1 = libc.Int32FromInt32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH) - libc.Int32FromInt32(4)
		}
		libc.Xmemcpy(tls, bp+40+uintptr(4), s, uint64(v1))
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+40, int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH), int32(chunk_offset*int64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		if **(**int32)(__ccgo_up(bp + 36)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			if isupdate != 0 && **(**int32)(__ccgo_up(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
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
			libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 56)), int32(1), rowid)
			libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 56)), int32(2), s, **(**int32)(__ccgo_up(bp + 36)), libc.UintptrFromInt32(0))
			rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 56)))
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 56)))
			if rc != int32(m_SQLITE_DONE) {
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
		} else {
			if **(**int32)(__ccgo_up(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12987, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
				if !(zSql1 != 0) {
					rc = int32(m_SQLITE_NOMEM)
					goto done
				}
				rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql1, -int32(1), bp+64, libc.UintptrFromInt32(0))
				if rc != m_SQLITE_OK {
					goto done
				}
				libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 64)), int32(1), rowid)
				rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 64)))
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 64)))
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
	rc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	if rc != m_SQLITE_OK {
		goto done
	}
	goto done
done:
	;
	return rc
	return r
}

type _JBTYPE = T_JBTYPE

type _SETJMP_FLOAT128 = T_SETJMP_FLOAT128

func _bitvec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	var blob uintptr
	var bytes, value_type int32
	_, _, _ = blob, bytes, value_type
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	if value_type == int32(m_SQLITE_BLOB) {
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		if bytes == 0 {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		**(**uintptr)(__ccgo_up(vector)) = blob
		**(**Tsize_t)(__ccgo_up(dimensions)) = uint64(bytes * int32(m___CHAR_BIT__))
		**(**Tvector_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(Xvector_cleanup_noop)
		return m_SQLITE_OK
	}
	**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+313, 0)
	return int32(m_SQLITE_ERROR)
}

func _distance_hamming_u64(tls *libc.TLS, a uintptr, b uintptr, n Tsize_t) (r Tf32) {
	var i uint32
	var same int32
	var v2 uint64
	_, _, _ = i, same, v2
	same = 0
	i = uint32(0)
	for {
		if !(uint64(i) < n) {
			break
		}
		v2 = uint64(libc.X__builtin_popcountll(tls, **(**Tu64)(__ccgo_up(a + uintptr(i)*8))^**(**Tu64)(__ccgo_up(b + uintptr(i)*8))))
		goto _3
	_3:
		same = int32(uint64(same) + v2)
		goto _1
	_1:
		;
		i = i + 1
	}
	return float32(same)
}

func _distance_hamming_u8(tls *libc.TLS, a uintptr, b uintptr, n Tsize_t) (r Tf32) {
	var i uint32
	var same int32
	_, _ = i, same
	same = 0
	i = uint32(0)
	for {
		if !(uint64(i) < n) {
			break
		}
		same = same + int32(_hamdist_table[int32(**(**Tu8)(__ccgo_up(a + uintptr(i))))^int32(**(**Tu8)(__ccgo_up(b + uintptr(i))))])
		goto _1
	_1:
		;
		i = i + 1
	}
	return float32(same)
}

func _fvec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var blob, buf, ptr, source uintptr
	var bytes, i, offset, rc, source_len, value_type int32
	var result, v1 float64
	var _ /* endptr at bp+32 */ uintptr
	var _ /* res at bp+40 */ Tf32
	var _ /* x at bp+0 */ TArray
	_, _, _, _, _, _, _, _, _, _, _, _ = blob, buf, bytes, i, offset, ptr, rc, result, source, source_len, value_type, v1
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	if value_type == int32(m_SQLITE_BLOB) {
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		if bytes == 0 {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		if uint64(bytes)%uint64(4) != uint64(0) {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+86, libc.VaList(bp+56, uint64(4), bytes))
			return int32(m_SQLITE_ERROR)
		}
		buf = libsqlite3.Xsqlite3_malloc(tls, bytes)
		if !(buf != 0) {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+156, 0)
			return int32(m_SQLITE_NOMEM)
		}
		libc.Xmemcpy(tls, buf, blob, uint64(bytes))
		**(**uintptr)(__ccgo_up(vector)) = buf
		**(**Tsize_t)(__ccgo_up(dimensions)) = uint64(bytes) / uint64(4)
		**(**Tfvec_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(libsqlite3.Xsqlite3_free)
		return m_SQLITE_OK
	}
	if value_type == int32(m_SQLITE_TEXT) {
		source = libsqlite3.Xsqlite3_value_text(tls, value)
		source_len = libsqlite3.Xsqlite3_value_bytes(tls, value)
		if source_len == 0 {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		i = 0
		rc = Xarray_init(tls, bp, uint64(4), uint64(libc.Xceil(tls, float64(source_len)/float64(2))))
		if rc != m_SQLITE_OK {
			return rc
		}
		// advance leading whitespace to first '['
		for i < source_len {
			if _vecJsonIsSpaceX[uint8(**(**int8)(__ccgo_up(source + uintptr(i))))] != 0 {
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
			**(**int32)(__ccgo_up(libc.X_errno(tls))) = 0
			v1 = libc.X__mingw_strtod(tls, ptr, bp+32)
			goto _2
		_2:
			result = v1
			if **(**int32)(__ccgo_up(libc.X_errno(tls))) != 0 && result == libc.Float64FromInt32(0) || **(**int32)(__ccgo_up(libc.X_errno(tls))) == int32(m_ERANGE) && (result == libc.X__builtin_huge_val(tls) || result == -libc.X__builtin_huge_val(tls)) {
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
			**(**Tf32)(__ccgo_up(bp + 40)) = float32(result)
			Xarray_append(tls, bp, bp+40)
			offset = int32(int64(offset) + (int64(**(**uintptr)(__ccgo_up(bp + 32))) - int64(ptr)))
			for offset < source_len {
				if _vecJsonIsSpaceX[uint8(**(**int8)(__ccgo_up(source + uintptr(offset))))] != 0 {
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
			**(**Tfvec_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(libsqlite3.Xsqlite3_free)
			return m_SQLITE_OK
		}
		libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
		return int32(m_SQLITE_ERROR)
	}
	**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+245, libc.VaList(bp+56, Xtype_name(tls, value_type)))
	return int32(m_SQLITE_ERROR)
}

func _int8_vec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var blob, ptr, source uintptr
	var bytes, i, offset, rc, result, source_len, value_type int32
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
		**(**Tsize_t)(__ccgo_up(dimensions)) = uint64(bytes)
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
			if _vecJsonIsSpaceX[uint8(**(**int8)(__ccgo_up(source + uintptr(i))))] != 0 {
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
			**(**int32)(__ccgo_up(libc.X_errno(tls))) = 0
			result = libc.Xstrtol(tls, ptr, bp+32, int32(10))
			if **(**int32)(__ccgo_up(libc.X_errno(tls))) != 0 && result == 0 || **(**int32)(__ccgo_up(libc.X_errno(tls))) == int32(m_ERANGE) && (result == int32(0x7fffffff) || result == -libc.Int32FromInt32(0x7fffffff)-libc.Int32FromInt32(1)) {
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
			if result < int32(-libc.Int32FromInt32(128)) || result > int32(m_INT8_MAX) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+341, 0)
				return int32(m_SQLITE_ERROR)
			}
			**(**Ti8)(__ccgo_up(bp + 40)) = int8(result)
			Xarray_append(tls, bp, bp+40)
			offset = int32(int64(offset) + (int64(**(**uintptr)(__ccgo_up(bp + 32))) - int64(ptr)))
			for offset < source_len {
				if _vecJsonIsSpaceX[uint8(**(**int8)(__ccgo_up(source + uintptr(offset))))] != 0 {
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

func _vec0BestIndex(tls *libc.TLS, pVTab uintptr, pIdxInfo uintptr) (r int32) {
	var argvIndex, hasAuxConstraint, i, i1, i2, i3, iColumn, iColumn1, iColumn2, iColumn3, iKTerm, iLimitTerm, iMatchTerm, iMatchVectorTerm, iRowidInTerm, iRowidTerm, metadata_idx, op, op1, op2, op3, partition_idx, rc, vtabIn1, v2 int32
	var idxStr, p uintptr
	var value, value1, value2 int8
	var vtabIn Tu8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = argvIndex, hasAuxConstraint, i, i1, i2, i3, iColumn, iColumn1, iColumn2, iColumn3, iKTerm, iLimitTerm, iMatchTerm, iMatchVectorTerm, iRowidInTerm, iRowidTerm, idxStr, metadata_idx, op, op1, op2, op3, p, partition_idx, rc, value, value1, value2, vtabIn, vtabIn1, v2
	p = pVTab
	/**
	 * Possible query plans are:
	 * 1. KNN when:
	 *    a) An `MATCH` op on vector column
	 *    b) ORDER BY on distance column
	 *    c) LIMIT
	 *    d) rowid in (...) OPTIONAL
	 * 2. Point when:
	 *    a) An `EQ` op on rowid column
	 * 3. else: fullscan
	 *
	 */
	iMatchTerm = -int32(1)
	iMatchVectorTerm = -int32(1)
	iLimitTerm = -int32(1)
	iRowidTerm = -int32(1)
	iKTerm = -int32(1)
	iRowidInTerm = -int32(1)
	hasAuxConstraint = 0
	i = 0
	for {
		if !(i < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
			break
		}
		vtabIn = uint8(0)
		if libsqlite3.Xsqlite3_libversion_number(tls) >= int32(3038000) {
			vtabIn = uint8(libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, i, -int32(1)))
		}
		if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).Fusable != 0) {
			goto _1
		}
		iColumn = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).FiColumn
		op = int32((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).Fop)
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) {
			iLimitTerm = i
		}
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_MATCH) && Xvec0_column_idx_is_vector(tls, p, iColumn) != 0 {
			if iMatchTerm > -int32(1) {
				Xvtab_set_error(tls, pVTab, __ccgo_ts+8434, 0)
				return int32(m_SQLITE_ERROR)
			}
			iMatchTerm = i
			iMatchVectorTerm = Xvec0_column_idx_to_vector_idx(tls, p, iColumn)
		}
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_EQ) && iColumn == m_VEC0_COLUMN_ID {
			if vtabIn != 0 {
				if iRowidInTerm != -int32(1) {
					Xvtab_set_error(tls, pVTab, __ccgo_ts+8490, 0)
					return int32(m_SQLITE_ERROR)
				}
				iRowidInTerm = i
			} else {
				iRowidTerm = i
			}
		}
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_EQ) && iColumn == Xvec0_column_k_idx(tls, p) {
			iKTerm = i
		}
		if op != int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) && op != int32(m_SQLITE_INDEX_CONSTRAINT_OFFSET) && Xvec0_column_idx_is_auxiliary(tls, p, iColumn) != 0 {
			hasAuxConstraint = int32(1)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	idxStr = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
	if iMatchTerm >= 0 {
		if iLimitTerm < 0 && iKTerm < 0 {
			Xvtab_set_error(tls, pVTab, __ccgo_ts+8556, 0)
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		if iLimitTerm >= 0 && iKTerm >= 0 {
			Xvtab_set_error(tls, pVTab, __ccgo_ts+8619, 0)
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		if (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnOrderBy != 0 {
			if (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnOrderBy > int32(1) {
				Xvtab_set_error(tls, pVTab, __ccgo_ts+8666, 0)
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
			if (**(**Tsqlite3_index_orderby)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaOrderBy))).FiColumn != Xvec0_column_distance_idx(tls, p) {
				Xvtab_set_error(tls, pVTab, __ccgo_ts+8738, 0)
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
			if (**(**Tsqlite3_index_orderby)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaOrderBy))).Fdesc != 0 {
				Xvtab_set_error(tls, pVTab, __ccgo_ts+8832, 0)
				rc = int32(m_SQLITE_ERROR)
				goto done
			}
		}
		if hasAuxConstraint != 0 {
			// IMP: V25623_09693
			Xvtab_set_error(tls, pVTab, __ccgo_ts+8916, 0)
			rc = int32(m_SQLITE_ERROR)
			goto done
		}
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_QUERY_PLAN_KNN))
		argvIndex = int32(1)
		v2 = argvIndex
		argvIndex = argvIndex + 1
		(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iMatchTerm)*8))).FargvIndex = v2
		(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iMatchTerm)*8))).Fomit = uint8(1)
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_MATCH))
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(3), int8('_'))
		if iLimitTerm >= 0 {
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iLimitTerm)*8))).FargvIndex = v2
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iLimitTerm)*8))).Fomit = uint8(1)
		} else {
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iKTerm)*8))).FargvIndex = v2
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iKTerm)*8))).Fomit = uint8(1)
		}
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_K))
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(3), int8('_'))
		if iRowidInTerm >= 0 {
			// already validated as  >= SQLite 3.38 bc iRowidInTerm is only >= 0 when
			// vtabIn == 1
			libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, iRowidInTerm, int32(1))
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iRowidInTerm)*8))).FargvIndex = v2
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iRowidInTerm)*8))).Fomit = uint8(1)
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_ROWID_IN))
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(3), int8('_'))
		}
		// find any PARTITION KEY column constraints
		i1 = 0
		for {
			if !(i1 < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
				break
			}
			if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i1)*12))).Fusable != 0) {
				goto _6
			}
			iColumn1 = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i1)*12))).FiColumn
			op1 = int32((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i1)*12))).Fop)
			if op1 == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) || op1 == int32(m_SQLITE_INDEX_CONSTRAINT_OFFSET) {
				goto _6
			}
			if !(Xvec0_column_idx_is_partition(tls, p, iColumn1) != 0) {
				goto _6
			}
			partition_idx = Xvec0_column_idx_to_partition_idx(tls, p, iColumn1)
			value = 0
			switch op1 {
			case int32(m_SQLITE_INDEX_CONSTRAINT_EQ):
				value = int8(_VEC0_PARTITION_OPERATOR_EQ)
			case int32(m_SQLITE_INDEX_CONSTRAINT_GT):
				value = int8(_VEC0_PARTITION_OPERATOR_GT)
			case int32(m_SQLITE_INDEX_CONSTRAINT_LE):
				value = int8(_VEC0_PARTITION_OPERATOR_LE)
			case int32(m_SQLITE_INDEX_CONSTRAINT_LT):
				value = int8(_VEC0_PARTITION_OPERATOR_LT)
			case int32(m_SQLITE_INDEX_CONSTRAINT_GE):
				value = int8(_VEC0_PARTITION_OPERATOR_GE)
			case int32(m_SQLITE_INDEX_CONSTRAINT_NE):
				value = int8(_VEC0_PARTITION_OPERATOR_NE)
				break
			}
			if value != 0 {
				v2 = argvIndex
				argvIndex = argvIndex + 1
				(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i1)*8))).FargvIndex = v2
				(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i1)*8))).Fomit = uint8(1)
				libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_PARTITON_CONSTRAINT))
				libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(int32('A')+partition_idx))
				libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), value)
				libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8('_'))
			}
			goto _6
		_6:
			;
			i1 = i1 + 1
		}
		// find any metadata column constraints
		i2 = 0
		for {
			if !(i2 < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
				break
			}
			if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i2)*12))).Fusable != 0) {
				goto _8
			}
			iColumn2 = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i2)*12))).FiColumn
			op2 = int32((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i2)*12))).Fop)
			if op2 == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) || op2 == int32(m_SQLITE_INDEX_CONSTRAINT_OFFSET) {
				goto _8
			}
			if !(Xvec0_column_idx_is_metadata(tls, p, iColumn2) != 0) {
				goto _8
			}
			metadata_idx = Xvec0_column_idx_to_metadata_idx(tls, p, iColumn2)
			value1 = 0
			switch op2 {
			case int32(m_SQLITE_INDEX_CONSTRAINT_EQ):
				vtabIn1 = 0
				if libsqlite3.Xsqlite3_libversion_number(tls) >= int32(3038000) {
					vtabIn1 = libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, i2, -int32(1))
				}
				if vtabIn1 != 0 {
					switch (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind {
					case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
						fallthrough
					case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
						// IMP: V15248_32086
						rc = int32(m_SQLITE_ERROR)
						Xvtab_set_error(tls, pVTab, __ccgo_ts+9000, 0)
						goto done
					case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
						fallthrough
					case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
						break
					}
					value1 = int8(_VEC0_METADATA_OPERATOR_IN)
					libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, i2, int32(1))
				} else {
					value1 = int8(_VEC0_PARTITION_OPERATOR_EQ)
				}
			case int32(m_SQLITE_INDEX_CONSTRAINT_GT):
				value1 = int8(_VEC0_METADATA_OPERATOR_GT)
			case int32(m_SQLITE_INDEX_CONSTRAINT_LE):
				value1 = int8(_VEC0_METADATA_OPERATOR_LE)
			case int32(m_SQLITE_INDEX_CONSTRAINT_LT):
				value1 = int8(_VEC0_METADATA_OPERATOR_LT)
			case int32(m_SQLITE_INDEX_CONSTRAINT_GE):
				value1 = int8(_VEC0_METADATA_OPERATOR_GE)
			case int32(m_SQLITE_INDEX_CONSTRAINT_NE):
				value1 = int8(_VEC0_METADATA_OPERATOR_NE)
			default:
				// IMP: V16511_00582
				rc = int32(m_SQLITE_ERROR)
				Xvtab_set_error(tls, pVTab, __ccgo_ts+9070, 0)
				goto done
			}
			if (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind == int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN) {
				if !(int32(value1) == int32(_VEC0_METADATA_OPERATOR_EQ) || int32(value1) == int32(_VEC0_METADATA_OPERATOR_NE)) {
					// IMP: V10145_26984
					rc = int32(m_SQLITE_ERROR)
					Xvtab_set_error(tls, pVTab, __ccgo_ts+9264, 0)
					goto done
				}
			}
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i2)*8))).FargvIndex = v2
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i2)*8))).Fomit = uint8(1)
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT))
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(int32('A')+metadata_idx))
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), value1)
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8('_'))
			goto _8
		_8:
			;
			i2 = i2 + 1
		}
		// find any distance column constraints
		i3 = 0
		for {
			if !(i3 < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
				break
			}
			if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i3)*12))).Fusable != 0) {
				goto _10
			}
			iColumn3 = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i3)*12))).FiColumn
			op3 = int32((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i3)*12))).Fop)
			if op3 == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) || op3 == int32(m_SQLITE_INDEX_CONSTRAINT_OFFSET) {
				goto _10
			}
			if Xvec0_column_distance_idx(tls, p) != iColumn3 {
				goto _10
			}
			value2 = 0
			switch op3 {
			case int32(m_SQLITE_INDEX_CONSTRAINT_GT):
				value2 = int8(_VEC0_DISTANCE_CONSTRAINT_GT)
			case int32(m_SQLITE_INDEX_CONSTRAINT_GE):
				value2 = int8(_VEC0_DISTANCE_CONSTRAINT_GE)
			case int32(m_SQLITE_INDEX_CONSTRAINT_LT):
				value2 = int8(_VEC0_DISTANCE_CONSTRAINT_LT)
			case int32(m_SQLITE_INDEX_CONSTRAINT_LE):
				value2 = int8(_VEC0_DISTANCE_CONSTRAINT_LE)
			default:
				// IMP TODO
				rc = int32(m_SQLITE_ERROR)
				Xvtab_set_error(tls, pVTab, __ccgo_ts+9350, 0)
				goto done
			}
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i3)*8))).FargvIndex = v2
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i3)*8))).Fomit = uint8(1)
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT))
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), value2)
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8('_'))
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8('_'))
			goto _10
		_10:
			;
			i3 = i3 + 1
		}
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxNum = iMatchVectorTerm
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = float64(30)
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(10)
	} else {
		if iRowidTerm >= 0 {
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_QUERY_PLAN_POINT))
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iRowidTerm)*8))).FargvIndex = int32(1)
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iRowidTerm)*8))).Fomit = uint8(1)
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_POINT_ID))
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(3), int8('_'))
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxNum = int32((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FcolUsed)
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = float64(10)
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(1)
		} else {
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_QUERY_PLAN_FULLSCAN))
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = float64(3e+06)
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(100000)
		}
	}
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxStr = libsqlite3.Xsqlite3_str_finish(tls, idxStr)
	idxStr = libc.UintptrFromInt32(0)
	if !((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxStr != 0) {
		rc = m_SQLITE_OK
		goto done
	}
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FneedToFreeIdxStr = int32(1)
	rc = m_SQLITE_OK
	goto done
done:
	;
	if idxStr != 0 {
		libsqlite3.Xsqlite3_str_finish(tls, idxStr)
	}
	return rc
}

func _vec0Column_fullscan(tls *libc.TLS, pVtab uintptr, pCur uintptr, context uintptr, i int32) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rc3, vector_idx int32
	var rowid Ti64
	var zErr uintptr
	var _ /* sz at bp+8 */ int32
	var _ /* v at bp+0 */ uintptr
	var _ /* v at bp+16 */ uintptr
	var _ /* v at bp+24 */ uintptr
	_, _, _, _, _, _, _, _, _, _ = auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rc3, rowid, vector_idx, zErr
	if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data != 0) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+10841, -int32(1))
		return int32(m_SQLITE_ERROR)
	}
	rowid = libsqlite3.Xsqlite3_column_int64(tls, (*Tvec0_query_fullscan_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data)).Frowids_stmt, 0)
	if i == m_VEC0_COLUMN_ID {
		return Xvec0_result_id(tls, pVtab, context, rowid)
	} else {
		if Xvec0_column_idx_is_vector(tls, pVtab, i) != 0 {
			vector_idx = Xvec0_column_idx_to_vector_idx(tls, pVtab, i)
			rc = Xvec0_get_vector_data(tls, pVtab, rowid, vector_idx, bp, bp+8)
			if rc != m_SQLITE_OK {
				return rc
			}
			libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), **(**int32)(__ccgo_up(bp + 8)), __ccgo_fp(libsqlite3.Xsqlite3_free))
			libsqlite3.Xsqlite3_result_subtype(tls, context, uint32((**(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_idx)*32))).Felement_type))
		} else {
			if i == Xvec0_column_distance_idx(tls, pVtab) {
				libsqlite3.Xsqlite3_result_null(tls, context)
			} else {
				if Xvec0_column_idx_is_partition(tls, pVtab, i) != 0 {
					partition_idx = Xvec0_column_idx_to_partition_idx(tls, pVtab, i)
					rc1 = Xvec0_get_partition_value_for_rowid(tls, pVtab, rowid, partition_idx, bp+16)
					if rc1 == m_SQLITE_OK {
						libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 16)))
						libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 16)))
					} else {
						libsqlite3.Xsqlite3_result_error_code(tls, context, rc1)
					}
				} else {
					if Xvec0_column_idx_is_auxiliary(tls, pVtab, i) != 0 {
						auxiliary_idx = Xvec0_column_idx_to_auxiliary_idx(tls, pVtab, i)
						rc2 = Xvec0_get_auxiliary_value_for_rowid(tls, pVtab, rowid, auxiliary_idx, bp+24)
						if rc2 == m_SQLITE_OK {
							libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 24)))
							libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
						} else {
							libsqlite3.Xsqlite3_result_error_code(tls, context, rc2)
						}
					} else {
						if Xvec0_column_idx_is_metadata(tls, pVtab, i) != 0 {
							if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
								return m_SQLITE_OK
							}
							metadata_idx = Xvec0_column_idx_to_metadata_idx(tls, pVtab, i)
							rc3 = Xvec0_result_metadata_value_for_rowid(tls, pVtab, rowid, metadata_idx, context)
							if rc3 != m_SQLITE_OK {
								// IMP: V15466_32305
								zErr = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+10891, libc.VaList(bp+40, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname_length, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname, rowid))
								if zErr != 0 {
									libsqlite3.Xsqlite3_result_error(tls, context, zErr, -int32(1))
									libsqlite3.Xsqlite3_free(tls, zErr)
								} else {
									libsqlite3.Xsqlite3_result_error_nomem(tls, context)
								}
							}
						}
					}
				}
			}
		}
	}
	return m_SQLITE_OK
}

func _vec0Column_knn(tls *libc.TLS, pVtab uintptr, pCur uintptr, context uintptr, i int32) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rc3, vector_idx int32
	var rowid, rowid1, rowid2, rowid3 Ti64
	var zErr uintptr
	var _ /* out at bp+0 */ uintptr
	var _ /* sz at bp+8 */ int32
	var _ /* v at bp+16 */ uintptr
	var _ /* v at bp+24 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rc3, rowid, rowid1, rowid2, rowid3, vector_idx, zErr
	if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data != 0) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+11001, -int32(1))
		return int32(m_SQLITE_ERROR)
	}
	if i == m_VEC0_COLUMN_ID {
		rowid = **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8))
		return Xvec0_result_id(tls, pVtab, context, rowid)
	} else {
		if i == Xvec0_column_distance_idx(tls, pVtab) {
			libsqlite3.Xsqlite3_result_double(tls, context, float64(**(**Tf32)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fdistances + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*4))))
			return m_SQLITE_OK
		} else {
			if Xvec0_column_idx_is_vector(tls, pVtab, i) != 0 {
				vector_idx = Xvec0_column_idx_to_vector_idx(tls, pVtab, i)
				rc = Xvec0_get_vector_data(tls, pVtab, **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8)), vector_idx, bp, bp+8)
				if rc != m_SQLITE_OK {
					return rc
				}
				libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), **(**int32)(__ccgo_up(bp + 8)), __ccgo_fp(libsqlite3.Xsqlite3_free))
				libsqlite3.Xsqlite3_result_subtype(tls, context, uint32((**(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_idx)*32))).Felement_type))
				return m_SQLITE_OK
			} else {
				if Xvec0_column_idx_is_partition(tls, pVtab, i) != 0 {
					partition_idx = Xvec0_column_idx_to_partition_idx(tls, pVtab, i)
					rowid1 = **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8))
					rc1 = Xvec0_get_partition_value_for_rowid(tls, pVtab, rowid1, partition_idx, bp+16)
					if rc1 == m_SQLITE_OK {
						libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 16)))
						libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 16)))
					} else {
						libsqlite3.Xsqlite3_result_error_code(tls, context, rc1)
					}
				} else {
					if Xvec0_column_idx_is_auxiliary(tls, pVtab, i) != 0 {
						auxiliary_idx = Xvec0_column_idx_to_auxiliary_idx(tls, pVtab, i)
						rowid2 = **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8))
						rc2 = Xvec0_get_auxiliary_value_for_rowid(tls, pVtab, rowid2, auxiliary_idx, bp+24)
						if rc2 == m_SQLITE_OK {
							libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 24)))
							libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
						} else {
							libsqlite3.Xsqlite3_result_error_code(tls, context, rc2)
						}
					} else {
						if Xvec0_column_idx_is_metadata(tls, pVtab, i) != 0 {
							metadata_idx = Xvec0_column_idx_to_metadata_idx(tls, pVtab, i)
							rowid3 = **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8))
							rc3 = Xvec0_result_metadata_value_for_rowid(tls, pVtab, rowid3, metadata_idx, context)
							if rc3 != m_SQLITE_OK {
								zErr = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+10891, libc.VaList(bp+40, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname_length, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname, rowid3))
								if zErr != 0 {
									libsqlite3.Xsqlite3_result_error(tls, context, zErr, -int32(1))
									libsqlite3.Xsqlite3_free(tls, zErr)
								} else {
									libsqlite3.Xsqlite3_result_error_nomem(tls, context)
								}
							}
						}
					}
				}
			}
		}
	}
	return m_SQLITE_OK
}

func _vec0Column_point(tls *libc.TLS, pVtab uintptr, pCur uintptr, context uintptr, i int32) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, vector_idx int32
	var rowid, rowid1, rowid2 Ti64
	var zErr uintptr
	var _ /* v at bp+0 */ uintptr
	var _ /* v at bp+8 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _ = auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rowid, rowid1, rowid2, vector_idx, zErr
	if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data != 0) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+10954, -int32(1))
		return int32(m_SQLITE_ERROR)
	}
	if i == m_VEC0_COLUMN_ID {
		return Xvec0_result_id(tls, pVtab, context, (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid)
	} else {
		if i == Xvec0_column_distance_idx(tls, pVtab) {
			libsqlite3.Xsqlite3_result_null(tls, context)
			return m_SQLITE_OK
		} else {
			if Xvec0_column_idx_is_vector(tls, pVtab, i) != 0 {
				if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
					libsqlite3.Xsqlite3_result_null(tls, context)
					return m_SQLITE_OK
				}
				vector_idx = Xvec0_column_idx_to_vector_idx(tls, pVtab, i)
				libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data + 8 + uintptr(vector_idx)*8)), int32(Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_idx)*32)))), uintptr(-libc.Int32FromInt32(1)))
				libsqlite3.Xsqlite3_result_subtype(tls, context, uint32((**(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_idx)*32))).Felement_type))
				return m_SQLITE_OK
			} else {
				if Xvec0_column_idx_is_partition(tls, pVtab, i) != 0 {
					if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
						return m_SQLITE_OK
					}
					partition_idx = Xvec0_column_idx_to_partition_idx(tls, pVtab, i)
					rowid = (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid
					rc = Xvec0_get_partition_value_for_rowid(tls, pVtab, rowid, partition_idx, bp)
					if rc == m_SQLITE_OK {
						libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp)))
						libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp)))
					} else {
						libsqlite3.Xsqlite3_result_error_code(tls, context, rc)
					}
				} else {
					if Xvec0_column_idx_is_auxiliary(tls, pVtab, i) != 0 {
						if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
							return m_SQLITE_OK
						}
						rowid1 = (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid
						auxiliary_idx = Xvec0_column_idx_to_auxiliary_idx(tls, pVtab, i)
						rc1 = Xvec0_get_auxiliary_value_for_rowid(tls, pVtab, rowid1, auxiliary_idx, bp+8)
						if rc1 == m_SQLITE_OK {
							libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 8)))
							libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 8)))
						} else {
							libsqlite3.Xsqlite3_result_error_code(tls, context, rc1)
						}
					} else {
						if Xvec0_column_idx_is_metadata(tls, pVtab, i) != 0 {
							if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
								return m_SQLITE_OK
							}
							rowid2 = (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid
							metadata_idx = Xvec0_column_idx_to_metadata_idx(tls, pVtab, i)
							rc2 = Xvec0_result_metadata_value_for_rowid(tls, pVtab, rowid2, metadata_idx, context)
							if rc2 != m_SQLITE_OK {
								zErr = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+10891, libc.VaList(bp+24, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname_length, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname, rowid2))
								if zErr != 0 {
									libsqlite3.Xsqlite3_result_error(tls, context, zErr, -int32(1))
									libsqlite3.Xsqlite3_free(tls, zErr)
								} else {
									libsqlite3.Xsqlite3_result_error_nomem(tls, context)
								}
							}
						}
					}
				}
			}
		}
	}
	return m_SQLITE_OK
}

func _vec0_init(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr, isCreate uint8) (r int32) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var auxiliary_idx, chunk_size, i, i1, i2, i3, i4, i5, i6, i7, metadata_idx, numAuxiliaryColumns, numMetadataColumns, numPartitionColumns, numVectorColumns, partition_idx, pkColumnNameLength, pkColumnType, rc, rc1, user_column_idx, vector_idx int32
	var createStr, pNew, pkColumnName, s, s1, schemaName, tableName, zCreateInfo, zCreateShadowChunks, zCreateShadowRowids, zSeedInfo, zSql, zSql1, zSql2, zSql3, zSql4 uintptr
	var _ /* auxColumn at bp+56 */ TVec0AuxiliaryColumnDefinition
	var _ /* cName at bp+104 */ uintptr
	var _ /* cNameLength at bp+112 */ int32
	var _ /* cType at bp+116 */ int32
	var _ /* key at bp+128 */ uintptr
	var _ /* keyLength at bp+144 */ int32
	var _ /* kind at bp+120 */ Tvec0_metadata_column_kind
	var _ /* metadataColumn at bp+80 */ TVec0MetadataColumnDefinition
	var _ /* partitionColumn at bp+32 */ TVec0PartitionColumnDefinition
	var _ /* stmt at bp+152 */ uintptr
	var _ /* stmt at bp+160 */ uintptr
	var _ /* value at bp+136 */ uintptr
	var _ /* valueLength at bp+148 */ int32
	var _ /* vecColumn at bp+0 */ TVectorColumnDefinition
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = auxiliary_idx, chunk_size, createStr, i, i1, i2, i3, i4, i5, i6, i7, metadata_idx, numAuxiliaryColumns, numMetadataColumns, numPartitionColumns, numVectorColumns, pNew, partition_idx, pkColumnName, pkColumnNameLength, pkColumnType, rc, rc1, s, s1, schemaName, tableName, user_column_idx, vector_idx, zCreateInfo, zCreateShadowChunks, zCreateShadowRowids, zSeedInfo, zSql, zSql1, zSql2, zSql3, zSql4
	_ = pAux
	pNew = libsqlite3.Xsqlite3_malloc(tls, int32(2032))
	if pNew == uintptr(0) {
		return int32(m_SQLITE_NOMEM)
	}
	libc.Xmemset(tls, pNew, 0, uint64(2032))
	// Declared chunk_size=N for entire table.
	// -1 to use the defualt, otherwise will get re-assigned on `chunk_size=N`
	// option
	chunk_size = -int32(1)
	numVectorColumns = 0
	numPartitionColumns = 0
	numAuxiliaryColumns = 0
	numMetadataColumns = 0
	user_column_idx = 0
	// track if a "primary key" column is defined
	pkColumnName = libc.UintptrFromInt32(0)
	pkColumnType = int32(m_SQLITE_INTEGER)
	i = int32(3)
	for {
		if !(i < argc) {
			break
		}
		**(**uintptr)(__ccgo_up(bp + 104)) = libc.UintptrFromInt32(0)
		// Scenario #1: Constructor argument is a vector column definition, ie `foo float[1024]`
		rc = Xvec0_parse_vector_column(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), int32(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp)
		if rc == int32(m_SQLITE_ERROR) {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5489, libc.VaList(bp+176, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
			goto error
		}
		if rc == m_SQLITE_OK {
			if numVectorColumns >= int32(m_VEC0_MAX_VECTOR_COLUMNS) {
				libsqlite3.Xsqlite3_free(tls, (**(**TVectorColumnDefinition)(__ccgo_up(bp))).Fname)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5548, libc.VaList(bp+176, int32(m_VEC0_MAX_VECTOR_COLUMNS)))
				goto error
			}
			if (**(**TVectorColumnDefinition)(__ccgo_up(bp))).Fdimensions > uint64(m_SQLITE_VEC_VEC0_MAX_DIMENSIONS) {
				libsqlite3.Xsqlite3_free(tls, (**(**TVectorColumnDefinition)(__ccgo_up(bp))).Fname)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5617, libc.VaList(bp+176, int64((**(**TVectorColumnDefinition)(__ccgo_up(bp))).Fdimensions), int32(m_SQLITE_VEC_VEC0_MAX_DIMENSIONS)))
				goto error
			}
			**(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(user_column_idx)*4)) = int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR)
			**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(user_column_idx))) = uint8(numVectorColumns)
			libc.Xmemcpy(tls, pNew+608+uintptr(numVectorColumns)*32, bp, uint64(32))
			numVectorColumns = numVectorColumns + 1
			(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumVectorColumns = numVectorColumns
			user_column_idx = user_column_idx + 1
			goto _1
		}
		// Scenario #2: Constructor argument is a partition key column definition, ie `user_id text partition key`
		rc = Xvec0_parse_partition_key_definition(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), int32(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+104, bp+112, bp+116)
		if rc == m_SQLITE_OK {
			if numPartitionColumns >= int32(m_VEC0_MAX_PARTITION_COLUMNS) {
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5707, libc.VaList(bp+176, int32(m_VEC0_MAX_PARTITION_COLUMNS)))
				goto error
			}
			(**(**TVec0PartitionColumnDefinition)(__ccgo_up(bp + 32))).Ftype1 = **(**int32)(__ccgo_up(bp + 116))
			(**(**TVec0PartitionColumnDefinition)(__ccgo_up(bp + 32))).Fname_length = **(**int32)(__ccgo_up(bp + 112))
			(**(**TVec0PartitionColumnDefinition)(__ccgo_up(bp + 32))).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+176, **(**int32)(__ccgo_up(bp + 112)), **(**uintptr)(__ccgo_up(bp + 104))))
			if !((**(**TVec0PartitionColumnDefinition)(__ccgo_up(bp + 32))).Fname != 0) {
				rc = int32(m_SQLITE_NOMEM)
				goto error
			}
			**(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(user_column_idx)*4)) = int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION)
			**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(user_column_idx))) = uint8(numPartitionColumns)
			libc.Xmemcpy(tls, pNew+1120+uintptr(numPartitionColumns)*24, bp+32, uint64(24))
			numPartitionColumns = numPartitionColumns + 1
			(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumPartitionColumns = numPartitionColumns
			user_column_idx = user_column_idx + 1
			goto _1
		}
		// Scenario #3: Constructor argument is a primary key column definition, ie `article_id text primary key`
		rc = Xvec0_parse_primary_key_definition(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), int32(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+104, bp+112, bp+116)
		if rc == m_SQLITE_OK {
			if pkColumnName != 0 {
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5780, libc.VaList(bp+176, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
				goto error
			}
			pkColumnName = **(**uintptr)(__ccgo_up(bp + 104))
			pkColumnNameLength = **(**int32)(__ccgo_up(bp + 112))
			pkColumnType = **(**int32)(__ccgo_up(bp + 116))
			goto _1
		}
		// Scenario #4: Constructor argument is a auxiliary column definition, ie `+contents text`
		rc = Xvec0_parse_auxiliary_column_definition(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), int32(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+104, bp+112, bp+116)
		if rc == m_SQLITE_OK {
			if numAuxiliaryColumns >= int32(m_VEC0_MAX_AUXILIARY_COLUMNS) {
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5901, libc.VaList(bp+176, int32(m_VEC0_MAX_AUXILIARY_COLUMNS)))
				goto error
			}
			(**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(bp + 56))).Ftype1 = **(**int32)(__ccgo_up(bp + 116))
			(**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(bp + 56))).Fname_length = **(**int32)(__ccgo_up(bp + 112))
			(**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(bp + 56))).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+176, **(**int32)(__ccgo_up(bp + 112)), **(**uintptr)(__ccgo_up(bp + 104))))
			if !((**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(bp + 56))).Fname != 0) {
				rc = int32(m_SQLITE_NOMEM)
				goto error
			}
			**(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(user_column_idx)*4)) = int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY)
			**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(user_column_idx))) = uint8(numAuxiliaryColumns)
			libc.Xmemcpy(tls, pNew+1216+uintptr(numAuxiliaryColumns)*24, bp+56, uint64(24))
			numAuxiliaryColumns = numAuxiliaryColumns + 1
			(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumAuxiliaryColumns = numAuxiliaryColumns
			user_column_idx = user_column_idx + 1
			goto _1
		}
		rc = Xvec0_parse_metadata_column_definition(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), int32(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+104, bp+112, bp+120)
		if rc == m_SQLITE_OK {
			if numMetadataColumns >= int32(m_VEC0_MAX_METADATA_COLUMNS) {
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5970, libc.VaList(bp+176, int32(m_VEC0_MAX_METADATA_COLUMNS)))
				goto error
			}
			(**(**TVec0MetadataColumnDefinition)(__ccgo_up(bp + 80))).Fkind = **(**Tvec0_metadata_column_kind)(__ccgo_up(bp + 120))
			(**(**TVec0MetadataColumnDefinition)(__ccgo_up(bp + 80))).Fname_length = **(**int32)(__ccgo_up(bp + 112))
			(**(**TVec0MetadataColumnDefinition)(__ccgo_up(bp + 80))).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+176, **(**int32)(__ccgo_up(bp + 112)), **(**uintptr)(__ccgo_up(bp + 104))))
			if !((**(**TVec0MetadataColumnDefinition)(__ccgo_up(bp + 80))).Fname != 0) {
				rc = int32(m_SQLITE_NOMEM)
				goto error
			}
			**(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(user_column_idx)*4)) = int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA)
			**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(user_column_idx))) = uint8(numMetadataColumns)
			libc.Xmemcpy(tls, pNew+1600+uintptr(numMetadataColumns)*24, bp+80, uint64(24))
			numMetadataColumns = numMetadataColumns + 1
			(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumMetadataColumns = numMetadataColumns
			user_column_idx = user_column_idx + 1
			goto _1
		}
		rc = Xvec0_parse_table_option(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), int32(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+128, bp+144, bp+136, bp+148)
		if rc == int32(m_SQLITE_ERROR) {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6038, libc.VaList(bp+176, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
			goto error
		}
		if rc == m_SQLITE_OK {
			if libsqlite3.Xsqlite3_strnicmp(tls, **(**uintptr)(__ccgo_up(bp + 128)), __ccgo_ts+6096, **(**int32)(__ccgo_up(bp + 144))) == 0 {
				chunk_size = libc.Xatoi(tls, **(**uintptr)(__ccgo_up(bp + 136)))
				if chunk_size <= 0 {
					// IMP: V01931_18769
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6107, 0)
					goto error
				}
				if chunk_size%int32(8) != 0 {
					// IMP: V14110_30948
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6178, 0)
					goto error
				}
				if chunk_size > int32(m_SQLITE_VEC_CHUNK_SIZE_MAX) {
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6236, 0)
					goto error
				}
			} else {
				// IMP: V27642_11712
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6281, libc.VaList(bp+176, **(**int32)(__ccgo_up(bp + 144)), **(**uintptr)(__ccgo_up(bp + 128))))
				goto error
			}
			goto _1
		}
		// Scenario #5: Unknown constructor argument
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6332, libc.VaList(bp+176, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
		goto error
		goto _1
	_1:
		;
		i = i + 1
	}
	if chunk_size < 0 {
		chunk_size = int32(1024)
	}
	if numVectorColumns <= 0 {
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6377, 0)
		goto error
	}
	createStr = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
	libsqlite3.Xsqlite3_str_appendall(tls, createStr, __ccgo_ts+6440)
	if pkColumnName != 0 {
		libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6456, libc.VaList(bp+176, pkColumnNameLength, pkColumnName))
	} else {
		libsqlite3.Xsqlite3_str_appendall(tls, createStr, __ccgo_ts+6477)
	}
	i1 = 0
	for {
		if !(i1 < numVectorColumns+numPartitionColumns+numAuxiliaryColumns+numMetadataColumns) {
			break
		}
		switch **(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(i1)*4)) {
		case int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR):
			vector_idx = int32(**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(i1))))
			libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6485, libc.VaList(bp+176, (**(**TVectorColumnDefinition)(__ccgo_up(pNew + 608 + uintptr(vector_idx)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(pNew + 608 + uintptr(vector_idx)*32))).Fname))
		case int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION):
			partition_idx = int32(**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(i1))))
			libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6485, libc.VaList(bp+176, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(pNew + 1120 + uintptr(partition_idx)*24))).Fname_length, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(pNew + 1120 + uintptr(partition_idx)*24))).Fname))
		case int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY):
			auxiliary_idx = int32(**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(i1))))
			libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6485, libc.VaList(bp+176, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(pNew + 1216 + uintptr(auxiliary_idx)*24))).Fname_length, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(pNew + 1216 + uintptr(auxiliary_idx)*24))).Fname))
		case int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA):
			metadata_idx = int32(**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(i1))))
			libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6485, libc.VaList(bp+176, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pNew + 1600 + uintptr(metadata_idx)*24))).Fname_length, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pNew + 1600 + uintptr(metadata_idx)*24))).Fname))
			break
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	libsqlite3.Xsqlite3_str_appendall(tls, createStr, __ccgo_ts+6494)
	if pkColumnName != 0 {
		libsqlite3.Xsqlite3_str_appendall(tls, createStr, __ccgo_ts+6523)
	}
	zSql = libsqlite3.Xsqlite3_str_finish(tls, createStr)
	if !(zSql != 0) {
		goto error
	}
	rc = libsqlite3.Xsqlite3_declare_vtab(tls, db, zSql)
	libsqlite3.Xsqlite3_free(tls, zSql)
	if rc != m_SQLITE_OK {
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6538, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
		goto error
	}
	schemaName = **(**uintptr)(__ccgo_up(argv + 1*8))
	tableName = **(**uintptr)(__ccgo_up(argv + 2*8))
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).Fdb = db
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FpkIsText = libc.BoolInt32(pkColumnType == int32(m_SQLITE_TEXT))
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6600, libc.VaList(bp+176, schemaName))
	if !((*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName != 0) {
		goto error
	}
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6600, libc.VaList(bp+176, tableName))
	if !((*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName != 0) {
		goto error
	}
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FshadowRowidsName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6603, libc.VaList(bp+176, tableName))
	if !((*Tvec0_vtab)(unsafe.Pointer(pNew)).FshadowRowidsName != 0) {
		goto error
	}
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FshadowChunksName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6613, libc.VaList(bp+176, tableName))
	if !((*Tvec0_vtab)(unsafe.Pointer(pNew)).FshadowChunksName != 0) {
		goto error
	}
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumVectorColumns = numVectorColumns
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumPartitionColumns = numPartitionColumns
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumAuxiliaryColumns = numAuxiliaryColumns
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumMetadataColumns = numMetadataColumns
	i2 = 0
	for {
		if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumVectorColumns) {
			break
		}
		**(**uintptr)(__ccgo_up(pNew + 352 + uintptr(i2)*8)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6623, libc.VaList(bp+176, tableName, i2))
		if !(**(**uintptr)(__ccgo_up(pNew + 352 + uintptr(i2)*8)) != 0) {
			goto error
		}
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	i3 = 0
	for {
		if !(i3 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumMetadataColumns) {
			break
		}
		**(**uintptr)(__ccgo_up(pNew + 480 + uintptr(i3)*8)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6644, libc.VaList(bp+176, tableName, i3))
		if !(**(**uintptr)(__ccgo_up(pNew + 480 + uintptr(i3)*8)) != 0) {
			goto error
		}
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).Fchunk_size = chunk_size
	// if xCreate, then create the necessary shadow tables
	if isCreate != 0 {
		zCreateInfo = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6666, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		if !(zCreateInfo != 0) {
			goto error
		}
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zCreateInfo, -int32(1), bp+152, libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_free(tls, zCreateInfo)
		if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
			// TODO(IMP)
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6728, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			goto error
		}
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
		zSeedInfo = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6770, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		if !(zSeedInfo != 0) {
			goto error
		}
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSeedInfo, -int32(1), bp+152, libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_free(tls, zSeedInfo)
		if rc1 != m_SQLITE_OK {
			// TODO(IMP)
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6856, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			goto error
		}
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(1), __ccgo_ts+6896, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(2), __ccgo_ts+6911, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(3), __ccgo_ts+6918, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(4), m_SQLITE_VEC_VERSION_MAJOR)
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(5), __ccgo_ts+6939, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(6), int32(m_SQLITE_VEC_VERSION_MINOR))
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(7), __ccgo_ts+6960, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(8), int32(m_SQLITE_VEC_VERSION_PATCH))
		if libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
			// TODO(IMP)
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6856, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			goto error
		}
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
		// create the _chunks shadow table
		zCreateShadowChunks = libc.UintptrFromInt32(0)
		if (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumPartitionColumns != 0 {
			s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+6981, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+7012)
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+7078)
			i4 = 0
			for {
				if !(i4 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumPartitionColumns) {
					break
				}
				libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+7099, libc.VaList(bp+176, i4))
				goto _5
			_5:
				;
				i4 = i4 + 1
			}
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+7114)
			zCreateShadowChunks = libsqlite3.Xsqlite3_str_finish(tls, s)
		} else {
			zCreateShadowChunks = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7161, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		}
		if !(zCreateShadowChunks != 0) {
			goto error
		}
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zCreateShadowChunks, -int32(1), bp+152, uintptr(0))
		libsqlite3.Xsqlite3_free(tls, zCreateShadowChunks)
		if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
			// IMP: V17740_01811
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7302, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			goto error
		}
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
		if (*Tvec0_vtab)(unsafe.Pointer(pNew)).FpkIsText != 0 {
			// adds a "text unique not null" constraint to the id column
			zCreateShadowRowids = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7346, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		} else {
			zCreateShadowRowids = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7480, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		}
		if !(zCreateShadowRowids != 0) {
			goto error
		}
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zCreateShadowRowids, -int32(1), bp+152, uintptr(0))
		libsqlite3.Xsqlite3_free(tls, zCreateShadowRowids)
		if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
			// IMP: V11631_28470
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7593, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			goto error
		}
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
		i5 = 0
		for {
			if !(i5 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumVectorColumns) {
				break
			}
			zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7637, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName, i5))
			if !(zSql1 != 0) {
				goto error
			}
			rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSql1, -int32(1), bp+152, uintptr(0))
			libsqlite3.Xsqlite3_free(tls, zSql1)
			if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
				// IMP: V25919_09989
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7720, libc.VaList(bp+176, i5, libsqlite3.Xsqlite3_errmsg(tls, db)))
				goto error
			}
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			goto _6
		_6:
			;
			i5 = i5 + 1
		}
		// See SHADOW_TABLE_ROWID_QUIRK in vec0_new_chunk() — same "rowid PRIMARY KEY"
		// without INTEGER type issue applies here.
		i6 = 0
		for {
			if !(i6 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumMetadataColumns) {
				break
			}
			zSql2 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7775, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName, i6))
			if !(zSql2 != 0) {
				goto error
			}
			rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSql2, -int32(1), bp+152, uintptr(0))
			libsqlite3.Xsqlite3_free(tls, zSql2)
			if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7857, libc.VaList(bp+176, i6, libsqlite3.Xsqlite3_errmsg(tls, db)))
				goto error
			}
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			if (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pNew + 1600 + uintptr(i6)*24))).Fkind == int32(_VEC0_METADATA_COLUMN_KIND_TEXT) {
				zSql3 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7912, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName, i6))
				if !(zSql3 != 0) {
					goto error
				}
				rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSql3, -int32(1), bp+152, uintptr(0))
				libsqlite3.Xsqlite3_free(tls, zSql3)
				if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
					libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7983, libc.VaList(bp+176, i6, libsqlite3.Xsqlite3_errmsg(tls, db)))
					goto error
				}
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			}
			goto _7
		_7:
			;
			i6 = i6 + 1
		}
		if (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumAuxiliaryColumns > 0 {
			s1 = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
			libsqlite3.Xsqlite3_str_appendf(tls, s1, __ccgo_ts+8037, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
			i7 = 0
			for {
				if !(i7 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumAuxiliaryColumns) {
					break
				}
				libsqlite3.Xsqlite3_str_appendf(tls, s1, __ccgo_ts+8098, libc.VaList(bp+176, i7))
				goto _8
			_8:
				;
				i7 = i7 + 1
			}
			libsqlite3.Xsqlite3_str_appendall(tls, s1, __ccgo_ts+5256)
			zSql4 = libsqlite3.Xsqlite3_str_finish(tls, s1)
			if !(zSql4 != 0) {
				goto error
			}
			rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSql4, -int32(1), bp+160, libc.UintptrFromInt32(0))
			if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 160))) != int32(m_SQLITE_DONE) {
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 160)))
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8110, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
				goto error
			}
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 160)))
		}
	}
	**(**uintptr)(__ccgo_up(ppVtab)) = pNew
	return m_SQLITE_OK
	goto error
error:
	;
	Xvec0_free(tls, pNew)
	libsqlite3.Xsqlite3_free(tls, pNew)
	return int32(m_SQLITE_ERROR)
}

// C documentation
//
//	/**
//	 * @brief Write the vector data into the provided vector blob at the given
//	 * offset
//	 *
//	 * @param blobVectors SQLite BLOB to write to
//	 * @param chunk_offset the "offset" (ie validity bitmap position) to write the
//	 * vector to
//	 * @param bVector pointer to the vector containing data
//	 * @param dimensions how many dimensions the vector has
//	 * @param element_type the vector type
//	 * @return result of sqlite3_blob_write, SQLITE_OK on success, otherwise failure
//	 */
func _vec0_write_vector_to_vector_blob(tls *libc.TLS, blobVectors uintptr, chunk_offset Ti64, bVector uintptr, dimensions Tsize_t, element_type _VectorElementType) (r int32) {
	var n, offset int32
	_, _ = n, offset
	switch element_type {
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		n = int32(dimensions * uint64(4))
		offset = int32(uint64(chunk_offset) * dimensions * uint64(4))
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		n = int32(dimensions * uint64(1))
		offset = int32(uint64(chunk_offset) * dimensions * uint64(1))
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		n = int32(dimensions / uint64(m___CHAR_BIT__))
		offset = int32(uint64(chunk_offset) * dimensions / uint64(m___CHAR_BIT__))
		break
	}
	return libsqlite3.Xsqlite3_blob_write(tls, blobVectors, bVector, n, offset)
}

func _vec_add(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i, i1, outSize, outSize1 Tsize_t
	var out, out1 uintptr
	var rc int32
	var _ /* a at bp+0 */ uintptr
	var _ /* aCleanup at bp+24 */ Tvector_cleanup
	var _ /* b at bp+8 */ uintptr
	var _ /* bCleanup at bp+32 */ Tvector_cleanup
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+48 */ _VectorElementType
	var _ /* error at bp+40 */ uintptr
	_, _, _, _, _, _, _ = i, i1, out, out1, outSize, outSize1, rc
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	rc = Xensure_vector_match(tls, **(**uintptr)(__ccgo_up(argv)), **(**uintptr)(__ccgo_up(argv + 1*8)), bp, bp+8, bp+48, bp+16, bp+24, bp+32, bp+40)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 40)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 40)))
		return
	}
	switch **(**_VectorElementType)(__ccgo_up(bp + 48)) {
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1183, -int32(1))
		goto finish
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		outSize = **(**Tsize_t)(__ccgo_up(bp + 16)) * uint64(4)
		out = libsqlite3.Xsqlite3_malloc(tls, int32(outSize))
		if !(out != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			goto finish
		}
		libc.Xmemset(tls, out, 0, outSize)
		i = uint64(0)
		for {
			if !(i < **(**Tsize_t)(__ccgo_up(bp + 16))) {
				break
			}
			**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4)) + **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i)*4))
			goto _1
		_1:
			;
			i = i + 1
		}
		libsqlite3.Xsqlite3_result_blob(tls, context, out, int32(outSize), __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
		goto finish
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		outSize1 = **(**Tsize_t)(__ccgo_up(bp + 16)) * uint64(1)
		out1 = libsqlite3.Xsqlite3_malloc(tls, int32(outSize1))
		if !(out1 != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			goto finish
		}
		libc.Xmemset(tls, out1, 0, outSize1)
		i1 = uint64(0)
		for {
			if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 16))) {
				break
			}
			**(**Ti8)(__ccgo_up(out1 + uintptr(i1))) = int8(int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i1)))) + int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i1)))))
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		libsqlite3.Xsqlite3_result_blob(tls, context, out1, int32(outSize1), __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
		goto finish
	}
	goto finish
finish:
	;
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp)))
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	return
}

func _vec_bit(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var rc int32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* errmsg at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	rc = _bitvec_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		return
	}
	libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), int32(**(**Tsize_t)(__ccgo_up(bp + 8))/uint64(m___CHAR_BIT__)), uintptr(-libc.Int32FromInt32(1)))
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_BIT))
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

func _vec_f32(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var rc int32
	var _ /* cleanup at bp+16 */ Tfvec_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* errmsg at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	rc = _fvec_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		return
	}
	libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), int32(**(**Tsize_t)(__ccgo_up(bp + 8))*uint64(4)), **(**Tfvec_cleanup)(__ccgo_up(bp + 16)))
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
}

func _vec_int8(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var rc int32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* errmsg at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	rc = _int8_vec_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		return
	}
	libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), int32(**(**Tsize_t)(__ccgo_up(bp + 8))), uintptr(-libc.Int32FromInt32(1)))
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

func _vec_length(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var rc int32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* errmsg at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		return
	}
	libsqlite3.Xsqlite3_result_int64(tls, context, int64(**(**Tsize_t)(__ccgo_up(bp + 8))))
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

func _vec_normalize(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i, i1 Tsize_t
	var norm Tf32
	var out, v uintptr
	var outSize, rc int32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* err at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_, _, _, _, _, _, _ = i, i1, norm, out, outSize, rc, v
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		return
	}
	if **(**_VectorElementType)(__ccgo_up(bp + 32)) != int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1688, -int32(1))
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
		return
	}
	outSize = int32(**(**Tsize_t)(__ccgo_up(bp + 8)) * uint64(4))
	out = libsqlite3.Xsqlite3_malloc(tls, outSize)
	if !(out != 0) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
		libsqlite3.Xsqlite3_result_error_code(tls, context, int32(m_SQLITE_NOMEM))
		return
	}
	libc.Xmemset(tls, out, 0, uint64(outSize))
	v = **(**uintptr)(__ccgo_up(bp))
	norm = libc.Float32FromInt32(0)
	i = uint64(0)
	for {
		if !(i < **(**Tsize_t)(__ccgo_up(bp + 8))) {
			break
		}
		norm = norm + Tf32(**(**Tf32)(__ccgo_up(v + uintptr(i)*4))***(**Tf32)(__ccgo_up(v + uintptr(i)*4)))
		goto _1
	_1:
		;
		i = i + 1
	}
	norm = float32(libc.Xsqrt(tls, float64(norm)))
	i1 = uint64(0)
	for {
		if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 8))) {
			break
		}
		**(**Tf32)(__ccgo_up(out + uintptr(i1)*4)) = **(**Tf32)(__ccgo_up(v + uintptr(i1)*4)) / norm
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	libsqlite3.Xsqlite3_result_blob(tls, context, out, int32(**(**Tsize_t)(__ccgo_up(bp + 8))*uint64(4)), __ccgo_fp(libsqlite3.Xsqlite3_free))
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

func _vec_npy_eachColumnBuffer(tls *libc.TLS, pCur uintptr, context uintptr, i int32) (r int32) {
	switch i {
	case m_VEC_NPY_EACH_COLUMN_VECTOR:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType))
		switch (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType {
		case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
			libsqlite3.Xsqlite3_result_blob(tls, context, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Fvector+uintptr(uint64((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FiRowid)*(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions*uint64(4)), int32((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions*uint64(4)), uintptr(-libc.Int32FromInt32(1)))
		case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
			fallthrough
		case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
			// https://github.com/asg017/sqlite-vec/issues/42
			libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+3426, -int32(1))
			break
		}
		break
	}
	return m_SQLITE_OK
}

func _vec_npy_eachColumnFile(tls *libc.TLS, pCur uintptr, context uintptr, i int32) (r int32) {
	switch i {
	case m_VEC_NPY_EACH_COLUMN_VECTOR:
		switch (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType {
		case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
			libsqlite3.Xsqlite3_result_blob(tls, context, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer+uintptr((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkIndex*(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions*uint64(4)), int32((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions*uint64(4)), uintptr(-libc.Int32FromInt32(1)))
		case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
			fallthrough
		case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
			// https://github.com/asg017/sqlite-vec/issues/42
			libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+3426, -int32(1))
			break
		}
		break
	}
	return m_SQLITE_OK
}

func _vec_npy_eachEof(tls *libc.TLS, cur uintptr) (r int32) {
	var pCur uintptr
	_ = pCur
	pCur = cur
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Finput_type == int32(_VEC_NPY_EACH_INPUT_BUFFER) {
		return libc.BoolInt32(!((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnElements != 0) || uint64((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FiRowid) >= (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnElements)
	}
	return (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Feof
}

func _vec_npy_file(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	var f, path uintptr
	var pathLength Tsize_t
	_, _, _ = f, path, pathLength
	path = libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv)))
	pathLength = uint64(libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv))))
	f = libsqlite3.Xsqlite3_malloc(tls, int32(16))
	if !(f != 0) {
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
		return
	}
	libc.Xmemset(tls, f, 0, uint64(16))
	(*TVecNpyFile)(unsafe.Pointer(f)).Fpath = path
	(*TVecNpyFile)(unsafe.Pointer(f)).FpathLength = pathLength
	libsqlite3.Xsqlite3_result_pointer(tls, context, f, __ccgo_ts+674, __ccgo_fp(libsqlite3.Xsqlite3_free))
}

func _vec_quantize_binary(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i, i1 Tsize_t
	var out, v2 uintptr
	var rc, res, res1, sz int32
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* pzError at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	var _ /* vectorCleanup at bp+16 */ Tvector_cleanup
	_, _, _, _, _, _, _, _ = i, i1, out, rc, res, res1, sz, v2
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		return
	}
	if **(**Tsize_t)(__ccgo_up(bp + 8)) <= uint64(0) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+974, -int32(1))
		goto cleanup
		return
	}
	if **(**Tsize_t)(__ccgo_up(bp + 8))%uint64(m___CHAR_BIT__) != uint64(0) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1013, -int32(1))
		goto cleanup
		return
	}
	sz = int32(**(**Tsize_t)(__ccgo_up(bp + 8)) / uint64(m___CHAR_BIT__))
	out = libsqlite3.Xsqlite3_malloc(tls, sz)
	if !(out != 0) {
		libsqlite3.Xsqlite3_result_error_code(tls, context, int32(m_SQLITE_NOMEM))
		goto cleanup
		return
	}
	libc.Xmemset(tls, out, 0, uint64(sz))
	switch **(**_VectorElementType)(__ccgo_up(bp + 32)) {
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		i = uint64(0)
		for {
			if !(i < **(**Tsize_t)(__ccgo_up(bp + 8))) {
				break
			}
			res = libc.BoolInt32(float64(**(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4))) > float64(0))
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
			if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 8))) {
				break
			}
			res1 = libc.BoolInt32(int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i1)))) > 0)
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
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

func _vec_quantize_int8(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i Tsize_t
	var out uintptr
	var rc, sz int32
	var step Tf32
	var val float64
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* err at bp+24 */ uintptr
	var _ /* srcCleanup at bp+16 */ Tfvec_cleanup
	var _ /* srcVector at bp+0 */ uintptr
	_, _, _, _, _, _ = i, out, rc, step, sz, val
	out = libc.UintptrFromInt32(0)
	rc = _fvec_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		return
	}
	sz = int32(**(**Tsize_t)(__ccgo_up(bp + 8)) * uint64(1))
	out = libsqlite3.Xsqlite3_malloc(tls, sz)
	if !(out != 0) {
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
		goto cleanup
	}
	libc.Xmemset(tls, out, 0, uint64(sz))
	if libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv + 1*8))) != int32(m_SQLITE_TEXT) || uint64(libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv + 1*8)))) != libc.Xstrlen(tls, __ccgo_ts+1126) || libsqlite3.Xsqlite3_stricmp(tls, libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv + 1*8))), __ccgo_ts+1126) != 0 {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1131, -int32(1))
		libsqlite3.Xsqlite3_free(tls, out)
		goto cleanup
	}
	step = float32((float64(1) - -libc.Float64FromFloat64(1)) / libc.Float64FromInt32(255))
	i = uint64(0)
	for {
		if !(i < **(**Tsize_t)(__ccgo_up(bp + 8))) {
			break
		}
		val = (float64(**(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4))) - -libc.Float64FromFloat64(1))/float64(step) - libc.Float64FromInt32(128)
		if !(val <= libc.Float64FromFloat64(127)) {
			val = float64(127)
		} /* also clamps NaN */
		if !(val >= -libc.Float64FromFloat64(128)) {
			val = -libc.Float64FromFloat64(128)
		}
		**(**Ti8)(__ccgo_up(out + uintptr(i))) = int8(val)
		goto _1
	_1:
		;
		i = i + 1
	}
	libsqlite3.Xsqlite3_result_blob(tls, context, out, int32(**(**Tsize_t)(__ccgo_up(bp + 8))*uint64(1)), __ccgo_fp(libsqlite3.Xsqlite3_free))
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
	goto cleanup
cleanup:
	;
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

func _vec_slice(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var end, outSize, outSize1, outSize2, rc, start int32
	var i, i1, i2, n Tsize_t
	var out, out1, out2 uintptr
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* err at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = end, i, i1, i2, n, out, out1, out2, outSize, outSize1, outSize2, rc, start
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		return
	}
	start = libsqlite3.Xsqlite3_value_int(tls, **(**uintptr)(__ccgo_up(argv + 1*8)))
	end = libsqlite3.Xsqlite3_value_int(tls, **(**uintptr)(__ccgo_up(argv + 2*8)))
	if start < 0 {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1260, -int32(1))
		goto done
	}
	if end < 0 {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1306, -int32(1))
		goto done
	}
	if uint64(start) > **(**Tsize_t)(__ccgo_up(bp + 8)) {
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1350, -int32(1))
		goto done
	}
	if uint64(end) > **(**Tsize_t)(__ccgo_up(bp + 8)) {
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
	n = uint64(end - start)
	switch **(**_VectorElementType)(__ccgo_up(bp + 32)) {
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		outSize = int32(n * uint64(4))
		out = libsqlite3.Xsqlite3_malloc(tls, outSize)
		if !(out != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			goto done
		}
		libc.Xmemset(tls, out, 0, uint64(outSize))
		i = uint64(0)
		for {
			if !(i < n) {
				break
			}
			**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(uint64(start)+i)*4))
			goto _1
		_1:
			;
			i = i + 1
		}
		libsqlite3.Xsqlite3_result_blob(tls, context, out, outSize, __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
		goto done
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		outSize1 = int32(n * uint64(1))
		out1 = libsqlite3.Xsqlite3_malloc(tls, outSize1)
		if !(out1 != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			return
		}
		libc.Xmemset(tls, out1, 0, uint64(outSize1))
		i1 = uint64(0)
		for {
			if !(i1 < n) {
				break
			}
			**(**Ti8)(__ccgo_up(out1 + uintptr(i1))) = **(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(uint64(start)+i1)))
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
		outSize2 = int32(n / uint64(m___CHAR_BIT__))
		out2 = libsqlite3.Xsqlite3_malloc(tls, outSize2)
		if !(out2 != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			return
		}
		libc.Xmemset(tls, out2, 0, uint64(outSize2))
		i2 = uint64(0)
		for {
			if !(i2 < n/uint64(m___CHAR_BIT__)) {
				break
			}
			**(**Tu8)(__ccgo_up(out2 + uintptr(i2))) = **(**Tu8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(uint64(start/libc.Int32FromInt32(m___CHAR_BIT__))+i2)))
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
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

func _vec_static_blob_entriesColumn(tls *libc.TLS, cur uintptr, context uintptr, i int32) (r int32) {
	var p, pCur uintptr
	var rowid Ti32
	_, _, _ = p, pCur, rowid
	pCur = cur
	p = (*Tsqlite3_vtab_cursor)(unsafe.Pointer(cur)).FpVtab
	switch (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	case int32(_VEC_SBE__QUERYPLAN_FULLSCAN):
		switch i {
		case m_VEC_STATIC_BLOB_ENTRIES_VECTOR:
			libsqlite3.Xsqlite3_result_blob(tls, context, (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fp+uintptr(uint64((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid)*(*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions*libc.Uint64FromInt64(4)), int32((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions*uint64(4)), uintptr(-libc.Int32FromInt32(1)))
			libsqlite3.Xsqlite3_result_subtype(tls, context, uint32((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Felement_type))
			break
		}
		return m_SQLITE_OK
	case int32(_VEC_SBE__QUERYPLAN_KNN):
		switch i {
		case m_VEC_STATIC_BLOB_ENTRIES_VECTOR:
			rowid = **(**Ti32)(__ccgo_up((*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*4))
			libsqlite3.Xsqlite3_result_blob(tls, context, (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fp+uintptr(uint64(rowid)*(*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions*libc.Uint64FromInt64(4)), int32((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions*uint64(4)), uintptr(-libc.Int32FromInt32(1)))
			libsqlite3.Xsqlite3_result_subtype(tls, context, uint32((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Felement_type))
			break
		}
		return m_SQLITE_OK
	}
	return int32(m_SQLITE_ERROR)
}

func _vec_static_blob_entriesEof(tls *libc.TLS, cur uintptr) (r int32) {
	var p, pCur uintptr
	_, _ = p, pCur
	pCur = cur
	p = (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fbase.FpVtab
	switch (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	case int32(_VEC_SBE__QUERYPLAN_FULLSCAN):
		return libc.BoolInt32(uint64((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid) >= (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors)
	case int32(_VEC_SBE__QUERYPLAN_KNN):
		return libc.BoolInt32((*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx >= (*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fk)
	}
	return int32(m_SQLITE_ERROR)
}

func _vec_static_blob_entriesFilter(tls *libc.TLS, pVtabCursor uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var bsize, i, i1 Tsize_t
	var candidates, distances, knn_data, p, pCur, taken, topk_rowids, v uintptr
	var k Ti64
	var rc int32
	var v1 int64
	var _ /* cleanup at bp+24 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+16 */ _VectorElementType
	var _ /* err at bp+32 */ uintptr
	var _ /* k_used at bp+40 */ Ti32
	var _ /* queryVector at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = bsize, candidates, distances, i, i1, k, knn_data, p, pCur, rc, taken, topk_rowids, v, v1
	_ = idxStr
	pCur = pVtabCursor
	p = (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fbase.FpVtab
	if idxNum == int32(_VEC_SBE__QUERYPLAN_KNN) {
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC_SBE__QUERYPLAN_KNN)
		knn_data = libsqlite3.Xsqlite3_malloc(tls, int32(40))
		if !(knn_data != 0) {
			return int32(m_SQLITE_NOMEM)
		}
		libc.Xmemset(tls, knn_data, 0, uint64(40))
		rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24, bp+32)
		if rc != m_SQLITE_OK {
			return int32(m_SQLITE_ERROR)
		}
		if **(**_VectorElementType)(__ccgo_up(bp + 16)) != (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Felement_type {
			return int32(m_SQLITE_ERROR)
		}
		if **(**Tsize_t)(__ccgo_up(bp + 8)) != (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions {
			return int32(m_SQLITE_ERROR)
		}
		if libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + 1*8))) <= int64((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors) {
			v1 = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + 1*8)))
		} else {
			v1 = int64((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors)
		}
		k = v1
		if k < 0 {
			// HANDLE https://github.com/asg017/sqlite-vec/issues/55
			return int32(m_SQLITE_ERROR)
		}
		if k == 0 {
			(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fk = 0
			(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data = knn_data
			return m_SQLITE_OK
		}
		bsize = ((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors + uint64(7)) & uint64(^libc.Int32FromInt32(7))
		topk_rowids = libsqlite3.Xsqlite3_malloc(tls, int32(uint64(k)*uint64(4)))
		if !(topk_rowids != 0) {
			// HANDLE https://github.com/asg017/sqlite-vec/issues/55
			return int32(m_SQLITE_ERROR)
		}
		distances = libsqlite3.Xsqlite3_malloc(tls, int32(bsize*uint64(4)))
		if !(distances != 0) {
			// HANDLE https://github.com/asg017/sqlite-vec/issues/55
			return int32(m_SQLITE_ERROR)
		}
		i = uint64(0)
		for {
			if !(i < (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors) {
				break
			}
			// https://github.com/asg017/sqlite-vec/issues/52
			v = (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fp + uintptr(i*(*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions)*4
			**(**Tf32)(__ccgo_up(distances + uintptr(i)*4)) = _distance_l2_sqr_float(tls, v, **(**uintptr)(__ccgo_up(bp)), (*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob+16)
			goto _2
		_2:
			;
			i = i + 1
		}
		candidates = Xbitmap_new(tls, int32(bsize))
		taken = Xbitmap_new(tls, int32(bsize))
		Xbitmap_fill(tls, candidates, int32(bsize))
		i1 = bsize
		for {
			if !(i1 >= (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors) {
				break
			}
			Xbitmap_set(tls, candidates, int32(i1), 0)
			goto _3
		_3:
			;
			i1 = i1 - 1
		}
		**(**Ti32)(__ccgo_up(bp + 40)) = 0
		Xmin_idx(tls, distances, int32(bsize), candidates, topk_rowids, int32(k), taken, bp+40)
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fcurrent_idx = 0
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances = distances
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fk = k
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Frowids = topk_rowids
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data = knn_data
	} else {
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC_SBE__QUERYPLAN_FULLSCAN)
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid = 0
	}
	return m_SQLITE_OK
}

func _vec_static_blob_from_raw(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	var p uintptr
	_ = p
	p = libsqlite3.Xsqlite3_malloc(tls, int32(32))
	if !(p != 0) {
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
		return
	}
	libc.Xmemset(tls, p, 0, uint64(32))
	(*Tstatic_blob_definition)(unsafe.Pointer(p)).Fp = uintptr(libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv))))
	(*Tstatic_blob_definition)(unsafe.Pointer(p)).Felement_type = int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32)
	(*Tstatic_blob_definition)(unsafe.Pointer(p)).Fdimensions = uint64(libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + 2*8))))
	(*Tstatic_blob_definition)(unsafe.Pointer(p)).Fnvectors = uint64(libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + 3*8))))
	libsqlite3.Xsqlite3_result_pointer(tls, context, p, _POINTER_NAME_STATIC_BLOB_DEF, __ccgo_fp(libsqlite3.Xsqlite3_free))
}

func _vec_static_blobsColumn(tls *libc.TLS, cur uintptr, context uintptr, i int32) (r int32) {
	var p, pCur uintptr
	_, _ = p, pCur
	pCur = cur
	p = (*Tsqlite3_vtab_cursor)(unsafe.Pointer(cur)).FpVtab
	switch i {
	case m_VEC_STATIC_BLOBS_NAME:
		libsqlite3.Xsqlite3_result_text(tls, context, (**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr((*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid)*40))).Fname, -int32(1), uintptr(-libc.Int32FromInt32(1)))
	case int32(m_VEC_STATIC_BLOBS_DATA):
		libsqlite3.Xsqlite3_result_null(tls, context)
	case int32(m_VEC_STATIC_BLOBS_DIMENSIONS):
		libsqlite3.Xsqlite3_result_int64(tls, context, int64((**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr((*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid)*40))).Fdimensions))
	case int32(m_VEC_STATIC_BLOBS_COUNT):
		libsqlite3.Xsqlite3_result_int64(tls, context, int64((**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr((*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid)*40))).Fnvectors))
		break
	}
	return m_SQLITE_OK
}

func _vec_sub(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i, i1, outSize, outSize1 Tsize_t
	var out, out1 uintptr
	var rc int32
	var _ /* a at bp+0 */ uintptr
	var _ /* aCleanup at bp+24 */ Tvector_cleanup
	var _ /* b at bp+8 */ uintptr
	var _ /* bCleanup at bp+32 */ Tvector_cleanup
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+48 */ _VectorElementType
	var _ /* error at bp+40 */ uintptr
	_, _, _, _, _, _, _ = i, i1, out, out1, outSize, outSize1, rc
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	rc = Xensure_vector_match(tls, **(**uintptr)(__ccgo_up(argv)), **(**uintptr)(__ccgo_up(argv + 1*8)), bp, bp+8, bp+48, bp+16, bp+24, bp+32, bp+40)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 40)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 40)))
		return
	}
	switch **(**_VectorElementType)(__ccgo_up(bp + 48)) {
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1219, -int32(1))
		goto finish
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		outSize = **(**Tsize_t)(__ccgo_up(bp + 16)) * uint64(4)
		out = libsqlite3.Xsqlite3_malloc(tls, int32(outSize))
		if !(out != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			goto finish
		}
		libc.Xmemset(tls, out, 0, outSize)
		i = uint64(0)
		for {
			if !(i < **(**Tsize_t)(__ccgo_up(bp + 16))) {
				break
			}
			**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4)) - **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i)*4))
			goto _1
		_1:
			;
			i = i + 1
		}
		libsqlite3.Xsqlite3_result_blob(tls, context, out, int32(outSize), __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
		goto finish
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		outSize1 = **(**Tsize_t)(__ccgo_up(bp + 16)) * uint64(1)
		out1 = libsqlite3.Xsqlite3_malloc(tls, int32(outSize1))
		if !(out1 != 0) {
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			goto finish
		}
		libc.Xmemset(tls, out1, 0, outSize1)
		i1 = uint64(0)
		for {
			if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 16))) {
				break
			}
			**(**Ti8)(__ccgo_up(out1 + uintptr(i1))) = int8(int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i1)))) - int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i1)))))
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		libsqlite3.Xsqlite3_result_blob(tls, context, out1, int32(outSize1), __ccgo_fp(libsqlite3.Xsqlite3_free))
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
		goto finish
	}
	goto finish
finish:
	;
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp)))
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	return
}

func _vec_to_json(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var b Tu8
	var i, len1, rc, v2 int32
	var i1 Tsize_t
	var s, str uintptr
	var value Tf32
	var _ /* cleanup at bp+24 */ Tvector_cleanup
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+40 */ _VectorElementType
	var _ /* err at bp+32 */ uintptr
	var _ /* hlp at bp+0 */ t__mingw_flt_type_t
	var _ /* vector at bp+8 */ uintptr
	_, _, _, _, _, _, _, _, _ = b, i, i1, len1, rc, s, str, value, v2
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp+8, bp+16, bp+40, bp+24, bp+32)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 32)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 32)))
		return
	}
	str = libsqlite3.Xsqlite3_str_new(tls, libsqlite3.Xsqlite3_context_db_handle(tls, context))
	libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1671)
	i1 = uint64(0)
	for {
		if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 16))) {
			break
		}
		if i1 != uint64(0) {
			libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1673)
		}
		if **(**_VectorElementType)(__ccgo_up(bp + 40)) == int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32) {
			value = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i1)*4))
			*(*float32)(unsafe.Pointer(bp)) = value
			i = int32(*(*uint32)(unsafe.Pointer(bp)) & uint32(0x7fffffff))
			i = int32(0x7f800000) - i
			v2 = int32(uint32(i) >> libc.Int32FromInt32(31))
			goto _3
		_3:
			if v2 != 0 {
				libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1675)
			} else {
				libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1680, libc.VaList(bp+56, float64(value)))
			}
		} else {
			if **(**_VectorElementType)(__ccgo_up(bp + 40)) == int32(_SQLITE_VEC_ELEMENT_TYPE_INT8) {
				libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1683, libc.VaList(bp+56, int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i1))))))
			} else {
				if **(**_VectorElementType)(__ccgo_up(bp + 40)) == int32(_SQLITE_VEC_ELEMENT_TYPE_BIT) {
					b = uint8(int32(**(**Tu8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i1/uint64(8))))) >> (i1 % uint64(m___CHAR_BIT__)) & int32(1))
					libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1683, libc.VaList(bp+56, int32(b)))
				}
			}
		}
		goto _1
	_1:
		;
		i1 = i1 + 1
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
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
}

const m_FP_NAN = 256

const m_FP_NORMAL = 1024

const m_FP_ZERO = 16384

const m_PRIXPTR = "PRIX64"

const m_PRIdPTR = "PRId64"

const m_PRIiPTR = "PRIi64"

const m_PRIoPTR = "PRIo64"

const m_PRIuPTR = "PRIu64"

const m_PRIxPTR = "PRIx64"

const m_SCNdPTR = "PRId64"

const m_SCNiPTR = "PRIi64"

const m_SCNoPTR = "PRIo64"

const m_SCNuPTR = "PRIu64"

const m_SCNxPTR = "PRIx64"

const m_SSIZE_MAX = "_I64_MAX"

const m_UNALIGNED = "__unaligned"

const m_WIN64 = 1

const m__ALLOCA_S_MARKER_SIZE = 16

const m__HEAP_MAXREQ = 0xFFFFFFFFFFFFFFE0

const m__M_AMD64 = 100

const m__M_X64 = 100

const m__WIN64 = 1

const m___MACHINEW64 = "__MACHINE"

const m___MACHINEX64 = "__MACHINE"

const m___MINGW64__ = 1

const m___MINGW_USE_UNDERSCORE_PREFIX = 0

const m___SEH__ = 1

const m___SIZE_MAX__ = "0xffffffffffffffffU"

const m___UINTPTR_MAX__ = "0xffffffffffffffffU"

const m___WIN64 = 1

const m___WIN64__ = 1

const m___code_model_medium__ = 1

type t__mingw_dbl_type_t = struct {
	Fval [0]uint64
	Flh  [0]struct {
		Flow  uint32
		Fhigh uint32
	}
	Fx float64
}

type t__mingw_ldbl_type_t = struct {
	Flh [0]struct {
		Flow      uint32
		Fhigh     uint32
		F__ccgo8  uint32
		F__ccgo12 uint32
	}
	Fx           float64
	F__ccgo_pad2 [8]byte
}

type t__uintr_frame = struct {
	Frip    uint64
	Frflags uint64
	Frsp    uint64
}
