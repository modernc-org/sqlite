// Code generated for linux/386 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build linux && 386

package vec

import (
	"unsafe"

	"modernc.org/libc"
	libsqlite3 "modernc.org/sqlite/lib"
)

func Xnpy_token_next(tls *libc.TLS, start uintptr, end uintptr, out uintptr) (r int32) {
	var curr uint8
	var ptr, start1, start2, v1 uintptr
	_, _, _, _, _ = curr, ptr, start1, start2, v1
	ptr = start
	for ptr < end {
		curr = **(**uint8)(__ccgo_up(ptr))
		if Xis_whitespace(tls, libc.Int8FromUint8(curr)) != 0 {
			ptr = ptr + 1
			continue
		} else {
			if libc.Int32FromUint8(curr) == int32('(') {
				v1 = ptr
				ptr = ptr + 1
				(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
				(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
				(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_LPAREN)
				return int32(m_VEC0_TOKEN_RESULT_SOME)
			} else {
				if libc.Int32FromUint8(curr) == int32(')') {
					v1 = ptr
					ptr = ptr + 1
					(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
					(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
					(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_RPAREN)
					return int32(m_VEC0_TOKEN_RESULT_SOME)
				} else {
					if libc.Int32FromUint8(curr) == int32('{') {
						v1 = ptr
						ptr = ptr + 1
						(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
						(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
						(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_LBRACE)
						return int32(m_VEC0_TOKEN_RESULT_SOME)
					} else {
						if libc.Int32FromUint8(curr) == int32('}') {
							v1 = ptr
							ptr = ptr + 1
							(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
							(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
							(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_RBRACE)
							return int32(m_VEC0_TOKEN_RESULT_SOME)
						} else {
							if libc.Int32FromUint8(curr) == int32(':') {
								v1 = ptr
								ptr = ptr + 1
								(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
								(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
								(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_COLON)
								return int32(m_VEC0_TOKEN_RESULT_SOME)
							} else {
								if libc.Int32FromUint8(curr) == int32(',') {
									v1 = ptr
									ptr = ptr + 1
									(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
									(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
									(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_COMMA)
									return int32(m_VEC0_TOKEN_RESULT_SOME)
								} else {
									if libc.Int32FromUint8(curr) == int32('\'') {
										start1 = ptr
										ptr = ptr + 1
										for ptr < end {
											if libc.Int32FromUint8(**(**uint8)(__ccgo_up(ptr))) == int32('\'') {
												break
											}
											ptr = ptr + 1
										}
										if ptr >= end || libc.Int32FromUint8(**(**uint8)(__ccgo_up(ptr))) != int32('\'') {
											return int32(m_VEC0_TOKEN_RESULT_ERROR)
										}
										(*TNpyToken)(unsafe.Pointer(out)).Fstart = start1
										ptr = ptr + 1
										v1 = ptr
										(*TNpyToken)(unsafe.Pointer(out)).Fend = v1
										(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_STRING)
										return int32(m_VEC0_TOKEN_RESULT_SOME)
									} else {
										if libc.Int32FromUint8(curr) == int32('F') && libc.Xstrncmp(tls, ptr, __ccgo_ts+1916, libc.Xstrlen(tls, __ccgo_ts+1916)) == 0 {
											(*TNpyToken)(unsafe.Pointer(out)).Fstart = ptr
											(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr + uintptr(libc.Int32FromUint32(libc.Xstrlen(tls, __ccgo_ts+1916)))
											ptr = (*TNpyToken)(unsafe.Pointer(out)).Fend
											(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_FALSE)
											return int32(m_VEC0_TOKEN_RESULT_SOME)
										} else {
											if Xis_digit(tls, libc.Int8FromUint8(curr)) != 0 {
												start2 = ptr
												for ptr < end && Xis_digit(tls, libc.Int8FromUint8(**(**uint8)(__ccgo_up(ptr)))) != 0 {
													ptr = ptr + 1
												}
												(*TNpyToken)(unsafe.Pointer(out)).Fstart = start2
												(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
												(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_NUMBER)
												return int32(m_VEC0_TOKEN_RESULT_SOME)
											} else {
												return int32(m_VEC0_TOKEN_RESULT_ERROR)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return int32(m_VEC0_TOKEN_RESULT_ERROR)
}

func Xvec0Filter_knn_chunks_iter(tls *libc.TLS, p uintptr, stmtChunks uintptr, vector_column uintptr, vectorColumnIdx int32, arrayRowidsIn uintptr, aMetadataIn uintptr, idxStr uintptr, argc int32, argv uintptr, queryVector uintptr, k Ti64, out_topk_rowids uintptr, out_topk_distances uintptr, out_used uintptr) (r int32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var b, bTaken, baseVectors, base_i, base_i1, base_i2, bmMetadata, bmRowids, chunkRowids, chunkValidity, chunk_distances, chunk_topk_idxs, in, tmp_topk_distances, tmp_topk_rowids, topk_distances, topk_rowids, v1 uintptr
	var baseVectorsSize, chunk_id, currentBaseVectorsSize, expectedBaseVectorsSize, k_used, rowidsSize, validitySize Ti64
	var hasDistanceConstraints, hasMetadataFilters, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, idx, idx1, idx2, idxStrLength, metadata_idx, numValueEntries, operator, rc, v4 int32
	var kind, kind1, kind2 int8
	var op Tvec0_distance_constraint_operator
	var result, target Tf32
	var v12, v13, v14 int64
	var _ /* blobVectors at bp+0 */ uintptr
	var _ /* metadataBlobs at bp+4 */ [16]uintptr
	var _ /* rowid at bp+72 */ Ti64
	var _ /* used at bp+88 */ Ti64
	var _ /* used1 at bp+80 */ int32
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
	topk_rowids = libsqlite3.Xsqlite3_malloc(tls, int32(k*int64(8)))
	if !(topk_rowids != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, topk_rowids, 0, libc.Uint32FromInt64(k*int64(8)))
	topk_distances = libsqlite3.Xsqlite3_malloc(tls, int32(k*int64(4)))
	if !(topk_distances != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, topk_distances, 0, libc.Uint32FromInt64(k*int64(4)))
	tmp_topk_rowids = libsqlite3.Xsqlite3_malloc(tls, int32(k*int64(8)))
	if !(tmp_topk_rowids != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, tmp_topk_rowids, 0, libc.Uint32FromInt64(k*int64(8)))
	tmp_topk_distances = libsqlite3.Xsqlite3_malloc(tls, int32(k*int64(4)))
	if !(tmp_topk_distances != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	libc.Xmemset(tls, tmp_topk_distances, 0, libc.Uint32FromInt64(k*int64(4)))
	k_used = 0
	baseVectorsSize = libc.Int64FromUint32(libc.Uint32FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(vector_column))))
	baseVectors = libsqlite3.Xsqlite3_malloc(tls, int32(baseVectorsSize))
	if !(baseVectors != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	chunk_distances = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint32(libc.Uint32FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint32(4)))
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
	chunk_topk_idxs = libsqlite3.Xsqlite3_malloc(tls, int32(k*int64(4)))
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
	libc.Xmemset(tls, bp+4, 0, libc.Uint32FromInt64(4)*libc.Uint32FromInt32(m_VEC0_MAX_METADATA_COLUMNS))
	bmMetadata = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	if !(bmMetadata != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto cleanup
	}
	idxStrLength = libc.Int32FromUint32(libc.Xstrlen(tls, idxStr))
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
		libc.Xmemset(tls, chunk_distances, 0, libc.Uint32FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint32(4))
		libc.Xmemset(tls, chunk_topk_idxs, 0, libc.Uint32FromInt64(k*int64(4)))
		Xbitmap_clear(tls, b, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		chunk_id = libsqlite3.Xsqlite3_column_int64(tls, stmtChunks, 0)
		chunkValidity = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(1))
		validitySize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(1)))
		if validitySize != int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m_CHAR_BIT)) {
			// IMP: V05271_22109
			Xvtab_set_error(tls, p, __ccgo_ts+9715, libc.VaList(bp+104, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m_CHAR_BIT), validitySize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		chunkRowids = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(2))
		rowidsSize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(2)))
		if rowidsSize != libc.Int64FromUint32(libc.Uint32FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint32(8)) {
			// IMP: V02796_19635
			Xvtab_set_error(tls, p, __ccgo_ts+9777, 0)
			Xvtab_set_error(tls, p, __ccgo_ts+9803, libc.VaList(bp+104, libc.Uint32FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint32(8), rowidsSize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		// open the vector chunk blob for the current chunk
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 312 + uintptr(vectorColumnIdx)*4)), __ccgo_ts+3712, chunk_id, 0, bp)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+9863, libc.VaList(bp+104, chunk_id))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		currentBaseVectorsSize = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp))))
		expectedBaseVectorsSize = libc.Int64FromUint32(libc.Uint32FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(vector_column))))
		if currentBaseVectorsSize != expectedBaseVectorsSize {
			// IMP: V16465_00535
			Xvtab_set_error(tls, p, __ccgo_ts+9906, libc.VaList(bp+104, expectedBaseVectorsSize, currentBaseVectorsSize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), baseVectors, int32(currentBaseVectorsSize), 0)
		if rc != m_SQLITE_OK {
			Xvtab_set_error(tls, p, __ccgo_ts+9966, libc.VaList(bp+104, chunk_id))
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
				**(**Ti64)(__ccgo_up(bp + 72)) = **(**Ti64)(__ccgo_up(chunkRowids + uintptr(i1)*8))
				in = libc.Xbsearch(tls, bp+72, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Fz, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Flength, uint32(8), __ccgo_fp(X_cmp))
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
				if !((**(**[16]uintptr)(__ccgo_up(bp + 4)))[metadata_idx] != 0) {
					rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 376 + uintptr(metadata_idx)*4)), __ccgo_ts+4053, chunk_id, 0, bp+4+uintptr(metadata_idx)*4)
					Xvtab_set_error(tls, p, __ccgo_ts+9999, 0)
					if rc != m_SQLITE_OK {
						goto cleanup
					}
				}
				Xbitmap_clear(tls, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
				rc = Xvec0_set_metadata_filter_bitmap(tls, p, metadata_idx, operator, **(**uintptr)(__ccgo_up(argv + uintptr(i2)*4)), (**(**[16]uintptr)(__ccgo_up(bp + 4)))[metadata_idx], chunk_id, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size, aMetadataIn, i2)
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
				base_i = baseVectors + uintptr(libc.Uint32FromInt32(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)*4
				switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdistance_metric {
				case int32(_VEC0_DISTANCE_METRIC_L2):
					result = _distance_l2_sqr_float(tls, base_i, queryVector, vector_column+8)
				case int32(_VEC0_DISTANCE_METRIC_L1):
					result = float32(_distance_l1_f32(tls, base_i, queryVector, vector_column+8))
				case int32(_VEC0_DISTANCE_METRIC_COSINE):
					result = _distance_cosine_float(tls, base_i, queryVector, vector_column+8)
					break
				}
			case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
				base_i1 = baseVectors + uintptr(libc.Uint32FromInt32(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)
				switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdistance_metric {
				case int32(_VEC0_DISTANCE_METRIC_L2):
					result = _distance_l2_sqr_int8(tls, base_i1, queryVector, vector_column+8)
				case int32(_VEC0_DISTANCE_METRIC_L1):
					result = float32(_distance_l1_int8(tls, base_i1, queryVector, vector_column+8))
				case int32(_VEC0_DISTANCE_METRIC_COSINE):
					result = _distance_cosine_int8(tls, base_i1, queryVector, vector_column+8)
					break
				}
			case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
				base_i2 = baseVectors + uintptr(libc.Uint32FromInt32(i3)*((*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions/libc.Uint32FromInt32(m_CHAR_BIT)))
				result = _distance_hamming(tls, base_i2, queryVector, vector_column+8)
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
				target = float32(libsqlite3.Xsqlite3_value_double(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i4)*4))))
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
		Xmin_idx(tls, chunk_distances, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size, b, chunk_topk_idxs, int32(v12), bTaken, bp+80)
		if k <= int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
			v13 = k
		} else {
			v13 = int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		}
		if v13 <= int64(**(**int32)(__ccgo_up(bp + 80))) {
			if k <= int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
				v14 = k
			} else {
				v14 = int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
			}
			v12 = v14
		} else {
			v12 = int64(**(**int32)(__ccgo_up(bp + 80)))
		}
		Xmerge_sorted_lists(tls, topk_distances, topk_rowids, k_used, chunk_distances, chunkRowids, chunk_topk_idxs, v12, tmp_topk_distances, tmp_topk_rowids, k, bp+88)
		i9 = 0
		for {
			if !(int64(i9) < **(**Ti64)(__ccgo_up(bp + 88))) {
				break
			}
			**(**Ti64)(__ccgo_up(topk_rowids + uintptr(i9)*8)) = **(**Ti64)(__ccgo_up(tmp_topk_rowids + uintptr(i9)*8))
			**(**Tf32)(__ccgo_up(topk_distances + uintptr(i9)*4)) = **(**Tf32)(__ccgo_up(tmp_topk_distances + uintptr(i9)*4))
			goto _16
		_16:
			;
			i9 = i9 + 1
		}
		k_used = **(**Ti64)(__ccgo_up(bp + 88))
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
		libsqlite3.Xsqlite3_blob_close(tls, (**(**[16]uintptr)(__ccgo_up(bp + 4)))[i10])
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

func Xvec0Filter_point(tls *libc.TLS, pCur uintptr, p uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, rc int32
	var point_data uintptr
	var _ /* rowid at bp+0 */ Ti64
	_, _, _ = i, point_data, rc
	point_data = libc.UintptrFromInt32(0)
	point_data = libsqlite3.Xsqlite3_malloc(tls, int32(76))
	if !(point_data != 0) {
		rc = int32(m_SQLITE_NOMEM)
		goto error
	}
	libc.Xmemset(tls, point_data, 0, uint32(76))
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FpkIsText != 0 {
		rc = Xvec0_rowid_from_id(tls, p, **(**uintptr)(__ccgo_up(argv)), bp)
		if rc == int32(m_SQLITE_EMPTY) {
			goto eof
		}
		if rc != m_SQLITE_OK {
			goto error
		}
	} else {
		**(**Ti64)(__ccgo_up(bp)) = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv)))
	}
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		rc = Xvec0_get_vector_data(tls, p, **(**Ti64)(__ccgo_up(bp)), i, point_data+8+uintptr(i)*4, libc.UintptrFromInt32(0))
		if rc == int32(m_SQLITE_EMPTY) {
			goto eof
		}
		if rc != m_SQLITE_OK {
			goto error
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	(*Tvec0_query_point_data)(unsafe.Pointer(point_data)).Frowid = **(**Ti64)(__ccgo_up(bp))
	(*Tvec0_query_point_data)(unsafe.Pointer(point_data)).Fdone = 0
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data = point_data
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC0_QUERY_PLAN_POINT)
	return m_SQLITE_OK
	goto eof
eof:
	;
	(*Tvec0_query_point_data)(unsafe.Pointer(point_data)).Frowid = **(**Ti64)(__ccgo_up(bp))
	(*Tvec0_query_point_data)(unsafe.Pointer(point_data)).Fdone = int32(1)
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data = point_data
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC0_QUERY_PLAN_POINT)
	return m_SQLITE_OK
	goto error
error:
	;
	Xvec0_query_point_data_clear(tls, point_data)
	libsqlite3.Xsqlite3_free(tls, point_data)
	return rc
}

func Xvec0Update_Delete_ClearValidity(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var brc, rc, validityOffset int32
	var mask uint8
	var _ /* blobChunksValidity at bp+0 */ uintptr
	var _ /* bx at bp+4 */ uint8
	var _ /* result at bp+5 */ int8
	_, _, _, _ = brc, mask, rc, validityOffset
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	validityOffset = libc.Int32FromUint64(chunk_offset / uint64(m_CHAR_BIT))
	// 2. ensure chunks.validity bit is 1, then set to 0
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+11207, chunk_id, int32(1), bp)
	if rc != m_SQLITE_OK {
		// IMP: V26002_10073
		Xvtab_set_error(tls, p, __ccgo_ts+13737, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id))
		return int32(m_SQLITE_ERROR)
	}
	// will skip the sqlite3_blob_bytes(blobChunksValidity) check for now,
	// the read below would catch it
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+4, int32(1), validityOffset)
	if rc != m_SQLITE_OK {
		// IMP: V21193_05263
		Xvtab_set_error(tls, p, __ccgo_ts+13781, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		goto cleanup
	}
	if !(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 4)))>>(chunk_offset%libc.Uint64FromInt32(m_CHAR_BIT)) != 0) {
		// IMP: V21193_05263
		rc = int32(m_SQLITE_ERROR)
		Xvtab_set_error(tls, p, __ccgo_ts+13831, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		goto cleanup
	}
	mask = libc.Uint8FromInt32(^(libc.Int32FromInt32(1) << (chunk_offset % libc.Uint64FromInt32(m_CHAR_BIT))))
	**(**int8)(__ccgo_up(bp + 5)) = int8(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 4))) & libc.Int32FromUint8(mask))
	rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+5, int32(1), validityOffset)
	if rc != m_SQLITE_OK {
		Xvtab_set_error(tls, p, __ccgo_ts+13897, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
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
		Xvtab_set_error(tls, p, __ccgo_ts+13951, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		return brc
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
	idxStrLength = libc.Int32FromUint32(libc.Xstrlen(tls, idxStr))
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
		libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(outStmt)), v3, **(**uintptr)(__ccgo_up(argv + uintptr(i1)*4)))
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	return rc
}

func _fvec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var blob, buf, ptr, source uintptr
	var bytes, i, offset, rc, source_len, value_type int32
	var result float64
	var _ /* endptr at bp+16 */ uintptr
	var _ /* res at bp+20 */ Tf32
	var _ /* x at bp+0 */ TArray
	_, _, _, _, _, _, _, _, _, _, _ = blob, buf, bytes, i, offset, ptr, rc, result, source, source_len, value_type
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	if value_type == int32(m_SQLITE_BLOB) {
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		if bytes == 0 {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			return int32(m_SQLITE_ERROR)
		}
		if libc.Uint32FromInt32(bytes)%uint32(4) != uint32(0) {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+86, libc.VaList(bp+32, uint32(4), bytes))
			return int32(m_SQLITE_ERROR)
		}
		buf = libsqlite3.Xsqlite3_malloc(tls, bytes)
		if !(buf != 0) {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+156, 0)
			return int32(m_SQLITE_NOMEM)
		}
		libc.Xmemcpy(tls, buf, blob, libc.Uint32FromInt32(bytes))
		**(**uintptr)(__ccgo_up(vector)) = buf
		**(**Tsize_t)(__ccgo_up(dimensions)) = libc.Uint32FromInt32(bytes) / uint32(4)
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
		rc = Xarray_init(tls, bp, uint32(4), uint32(libc.Xceil(tls, float64(source_len)/float64(2))))
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
			**(**int32)(__ccgo_up(libc.X__errno_location(tls))) = 0
			result = libc.Xstrtod(tls, ptr, bp+16)
			if **(**int32)(__ccgo_up(libc.X__errno_location(tls))) != 0 && result == libc.Float64FromInt32(0) || **(**int32)(__ccgo_up(libc.X__errno_location(tls))) == int32(m_ERANGE) && (result == float64(libc.X__builtin_inff(tls)) || result == -float64(libc.X__builtin_inff(tls))) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if **(**uintptr)(__ccgo_up(bp + 16)) == ptr {
				if int32(**(**int8)(__ccgo_up(ptr))) != int32(']') {
					libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
					return int32(m_SQLITE_ERROR)
				}
				goto done
			}
			**(**Tf32)(__ccgo_up(bp + 20)) = float32(result)
			Xarray_append(tls, bp, bp+20)
			offset = offset + (int32(**(**uintptr)(__ccgo_up(bp + 16))) - int32(ptr))
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
		if (**(**TArray)(__ccgo_up(bp))).Flength > uint32(0) {
			**(**uintptr)(__ccgo_up(vector)) = (**(**TArray)(__ccgo_up(bp))).Fz
			**(**Tsize_t)(__ccgo_up(dimensions)) = (**(**TArray)(__ccgo_up(bp))).Flength
			**(**Tfvec_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(libsqlite3.Xsqlite3_free)
			return m_SQLITE_OK
		}
		libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
		return int32(m_SQLITE_ERROR)
	}
	**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+245, libc.VaList(bp+32, Xtype_name(tls, value_type)))
	return int32(m_SQLITE_ERROR)
}

func _int8_vec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var blob, ptr, source uintptr
	var bytes, i, offset, rc, result, source_len, value_type int32
	var _ /* endptr at bp+16 */ uintptr
	var _ /* res at bp+20 */ Ti8
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
		**(**Tsize_t)(__ccgo_up(dimensions)) = libc.Uint32FromInt32(bytes)
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
		rc = Xarray_init(tls, bp, uint32(1), uint32(libc.Xceil(tls, float64(source_len)/float64(2))))
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
			**(**int32)(__ccgo_up(libc.X__errno_location(tls))) = 0
			result = libc.Xstrtol(tls, ptr, bp+16, int32(10))
			if **(**int32)(__ccgo_up(libc.X__errno_location(tls))) != 0 && result == 0 || **(**int32)(__ccgo_up(libc.X__errno_location(tls))) == int32(m_ERANGE) && (result == int32(0x7fffffff) || result == -libc.Int32FromInt32(0x7fffffff)-libc.Int32FromInt32(1)) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if **(**uintptr)(__ccgo_up(bp + 16)) == ptr {
				if int32(**(**int8)(__ccgo_up(ptr))) != int32(']') {
					libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
					return int32(m_SQLITE_ERROR)
				}
				goto done
			}
			if result < int32(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7f)) || result > int32(libc.Int32FromInt32(m_INT8_MAX)) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+341, 0)
				return int32(m_SQLITE_ERROR)
			}
			**(**Ti8)(__ccgo_up(bp + 20)) = int8(result)
			Xarray_append(tls, bp, bp+20)
			offset = offset + (int32(**(**uintptr)(__ccgo_up(bp + 16))) - int32(ptr))
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
		if (**(**TArray)(__ccgo_up(bp))).Flength > uint32(0) {
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
			vtabIn = libc.Uint8FromInt32(libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, i, -int32(1)))
		}
		if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).Fusable != 0) {
			goto _1
		}
		iColumn = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).FiColumn
		op = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).Fop)
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
			op1 = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i1)*12))).Fop)
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
			op2 = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i2)*12))).Fop)
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
					switch (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1000 + uintptr(metadata_idx)*12))).Fkind {
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
			if (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1000 + uintptr(metadata_idx)*12))).Fkind == int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN) {
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
			op3 = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i3)*12))).Fop)
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
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxNum = libc.Int32FromUint64((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FcolUsed)
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

func _vec0Filter(tls *libc.TLS, pVtabCursor uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var idxStrLength, numValueEntries int32
	var p, pCur uintptr
	var query_plan int8
	_, _, _, _, _ = idxStrLength, numValueEntries, p, pCur, query_plan
	p = (*Tsqlite3_vtab_cursor)(unsafe.Pointer(pVtabCursor)).FpVtab
	pCur = pVtabCursor
	Xvec0_cursor_clear(tls, pCur)
	idxStrLength = libc.Int32FromUint32(libc.Xstrlen(tls, idxStr))
	if idxStrLength <= 0 {
		return int32(m_SQLITE_ERROR)
	}
	if (idxStrLength-int32(1))%int32(4) != 0 {
		return int32(m_SQLITE_ERROR)
	}
	numValueEntries = (idxStrLength - int32(1)) / int32(4)
	if numValueEntries != argc {
		return int32(m_SQLITE_ERROR)
	}
	query_plan = **(**int8)(__ccgo_up(idxStr))
	switch int32(query_plan) {
	case int32(_VEC0_QUERY_PLAN_FULLSCAN):
		return Xvec0Filter_fullscan(tls, p, pCur)
	case int32(_VEC0_QUERY_PLAN_KNN):
		return Xvec0Filter_knn(tls, pCur, p, idxNum, idxStr, argc, argv)
	case int32(_VEC0_QUERY_PLAN_POINT):
		return Xvec0Filter_point(tls, pCur, p, argc, argv)
	default:
		Xvtab_set_error(tls, (*Tsqlite3_vtab_cursor)(unsafe.Pointer(pVtabCursor)).FpVtab, __ccgo_ts+10745, libc.VaList(bp+8, idxStr))
		return int32(m_SQLITE_ERROR)
	}
	return r
}

func _vec_eachFilter(tls *libc.TLS, pVtabCursor uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var pCur uintptr
	var rc int32
	var _ /* pzErrMsg at bp+0 */ uintptr
	_, _ = pCur, rc
	_ = idxNum
	_ = idxStr
	pCur = pVtabCursor
	if (*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fcleanup})))(tls, (*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector)
		(*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector = libc.UintptrFromInt32(0)
	}
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), pCur+16, pCur+20, pCur+12, pCur+24, bp)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp)))
		return int32(m_SQLITE_ERROR)
	}
	(*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid = 0
	return m_SQLITE_OK
}

func _vec_eachOpen(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	var pCur uintptr
	_ = pCur
	_ = p
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(28))
	if pCur == uintptr(0) {
		return int32(m_SQLITE_NOMEM)
	}
	libc.Xmemset(tls, pCur, 0, uint32(28))
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	return m_SQLITE_OK
}

func _vec_npy_eachOpen(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	var pCur uintptr
	_ = pCur
	_ = p
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(60))
	if pCur == uintptr(0) {
		return int32(m_SQLITE_NOMEM)
	}
	libc.Xmemset(tls, pCur, 0, uint32(60))
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	return m_SQLITE_OK
}

func _vec_static_blob_entriesOpen(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	var pCur uintptr
	_ = pCur
	_ = p
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(20))
	if pCur == uintptr(0) {
		return int32(m_SQLITE_NOMEM)
	}
	libc.Xmemset(tls, pCur, 0, uint32(20))
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	return m_SQLITE_OK
}

func _vec_static_blobsOpen(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	var pCur uintptr
	_ = pCur
	_ = p
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(12))
	if pCur == uintptr(0) {
		return int32(m_SQLITE_NOMEM)
	}
	libc.Xmemset(tls, pCur, 0, uint32(12))
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	return m_SQLITE_OK
}
