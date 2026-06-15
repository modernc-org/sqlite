// Code generated for freebsd/arm by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build freebsd && arm

package vec

import (
	"unsafe"

	"modernc.org/libc"
	libsqlite3 "modernc.org/sqlite/lib"
)

type TFILE = struct {
	F__ccgo_align [0]uint32
	F_p           uintptr
	F_r           int32
	F_w           int32
	F_flags       int16
	F_file        int16
	F_bf          t__sbuf
	F_lbfsize     int32
	F_cookie      uintptr
	F_close       uintptr
	F_read        uintptr
	F_seek        uintptr
	F_write       uintptr
	F_ub          t__sbuf
	F_up          uintptr
	F_ur          int32
	F_ubuf        [3]uint8
	F_nbuf        [1]uint8
	F_lb          t__sbuf
	F_blksize     int32
	F_offset      Tfpos_t
	F_fl_mutex    uintptr
	F_fl_owner    uintptr
	F_fl_count    int32
	F_orientation int32
	F_mbstate     t__mbstate_t
	F_flags2      int32
	F__ccgo_pad26 [4]byte
}

func Xvec0Filter_knn_chunks_iter(tls *libc.TLS, p uintptr, stmtChunks uintptr, vector_column uintptr, vectorColumnIdx int32, arrayRowidsIn uintptr, aMetadataIn uintptr, idxStr uintptr, argc int32, argv uintptr, queryVector uintptr, k Ti64, out_topk_rowids uintptr, out_topk_distances uintptr, out_used uintptr) (r int32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var b, bTaken, baseVectors, base_i, base_i1, base_i2, bmMetadata, bmRowids, chunkRowids, chunkValidity, chunk_distances, chunk_topk_idxs, in, tmp_topk_distances, tmp_topk_rowids, topk_distances, topk_rowids, v1 uintptr
	var baseVectorsSize, chunk_id, currentBaseVectorsSize, expectedBaseVectorsSize, k_used, rowidsSize, validitySize Ti64
	var hasDistanceConstraints, hasMetadataFilters, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, idx, idx1, idx2, idxStrLength, metadata_idx, numValueEntries, operator, rc, v4 int32
	var kind, kind1, kind2 uint8
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
		kind = **(**uint8)(__ccgo_up(idxStr + uintptr(idx+0)))
		if libc.Int32FromUint8(kind) == int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) {
			hasMetadataFilters = int32(1)
		} else {
			if libc.Int32FromUint8(kind) == int32(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT) {
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
		if validitySize != int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT)) {
			// IMP: V05271_22109
			Xvtab_set_error(tls, p, __ccgo_ts+9715, libc.VaList(bp+104, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT), validitySize))
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
				kind1 = **(**uint8)(__ccgo_up(idxStr + uintptr(idx1+0)))
				if libc.Int32FromUint8(kind1) != int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) {
					goto _5
				}
				metadata_idx = libc.Int32FromUint8(**(**uint8)(__ccgo_up(idxStr + uintptr(idx1+int32(1))))) - int32('A')
				operator = libc.Int32FromUint8(**(**uint8)(__ccgo_up(idxStr + uintptr(idx1+int32(2)))))
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
				base_i2 = baseVectors + uintptr(libc.Uint32FromInt32(i3)*((*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions/libc.Uint32FromInt32(m___CHAR_BIT)))
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
				kind2 = **(**uint8)(__ccgo_up(idxStr + uintptr(idx2+0)))
				// TODO casts f64 to f32, is that a problem?
				target = float32(libsqlite3.Xsqlite3_value_double(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i4)*4))))
				if libc.Int32FromUint8(kind2) != int32(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT) {
					goto _7
				}
				op = libc.Int32FromUint8(**(**uint8)(__ccgo_up(idxStr + uintptr(idx2+int32(1)))))
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

func Xvec0Update_Delete_ClearValidity(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var brc, rc, validityOffset int32
	var mask uint8
	var _ /* blobChunksValidity at bp+0 */ uintptr
	var _ /* bx at bp+4 */ uint8
	var _ /* result at bp+5 */ uint8
	_, _, _, _ = brc, mask, rc, validityOffset
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	validityOffset = libc.Int32FromUint64(chunk_offset / uint64(m___CHAR_BIT))
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
	if !(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 4)))>>(chunk_offset%libc.Uint64FromInt32(m___CHAR_BIT)) != 0) {
		// IMP: V21193_05263
		rc = int32(m_SQLITE_ERROR)
		Xvtab_set_error(tls, p, __ccgo_ts+13831, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		goto cleanup
	}
	mask = libc.Uint8FromInt32(^(libc.Int32FromInt32(1) << (chunk_offset % libc.Uint64FromInt32(m___CHAR_BIT))))
	**(**uint8)(__ccgo_up(bp + 5)) = libc.Uint8FromInt32(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 4))) & libc.Int32FromUint8(mask))
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
			if _vecJsonIsSpaceX[uint8(**(**uint8)(__ccgo_up(source + uintptr(i))))] != 0 {
				i = i + 1
				continue
			}
			if libc.Int32FromUint8(**(**uint8)(__ccgo_up(source + uintptr(i)))) == int32('[') {
				break
			}
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			Xarray_cleanup(tls, bp)
			return int32(m_SQLITE_ERROR)
		}
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(source + uintptr(i)))) != int32('[') {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			Xarray_cleanup(tls, bp)
			return int32(m_SQLITE_ERROR)
		}
		offset = i + int32(1)
		for offset < source_len {
			ptr = source + uintptr(offset)
			**(**int32)(__ccgo_up(libc.X__error(tls))) = 0
			result = libc.Xstrtod(tls, ptr, bp+16)
			if **(**int32)(__ccgo_up(libc.X__error(tls))) != 0 && result == libc.Float64FromInt32(0) || **(**int32)(__ccgo_up(libc.X__error(tls))) == int32(m_ERANGE) && (result == libc.X__builtin_huge_val(tls) || result == -libc.X__builtin_huge_val(tls)) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if **(**uintptr)(__ccgo_up(bp + 16)) == ptr {
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(ptr))) != int32(']') {
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
				if _vecJsonIsSpaceX[uint8(**(**uint8)(__ccgo_up(source + uintptr(offset))))] != 0 {
					offset = offset + 1
					continue
				}
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(source + uintptr(offset)))) == int32(',') {
					offset = offset + 1
					continue
				}
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(source + uintptr(offset)))) == int32(']') {
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
			if _vecJsonIsSpaceX[uint8(**(**uint8)(__ccgo_up(source + uintptr(i))))] != 0 {
				i = i + 1
				continue
			}
			if libc.Int32FromUint8(**(**uint8)(__ccgo_up(source + uintptr(i)))) == int32('[') {
				break
			}
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			Xarray_cleanup(tls, bp)
			return int32(m_SQLITE_ERROR)
		}
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(source + uintptr(i)))) != int32('[') {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			Xarray_cleanup(tls, bp)
			return int32(m_SQLITE_ERROR)
		}
		offset = i + int32(1)
		for offset < source_len {
			ptr = source + uintptr(offset)
			**(**int32)(__ccgo_up(libc.X__error(tls))) = 0
			result = libc.Xstrtol(tls, ptr, bp+16, int32(10))
			if **(**int32)(__ccgo_up(libc.X__error(tls))) != 0 && result == 0 || **(**int32)(__ccgo_up(libc.X__error(tls))) == int32(m_ERANGE) && (result == int32(0x7fffffff) || result == -libc.Int32FromInt32(0x7fffffff)-libc.Int32FromInt32(1)) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if **(**uintptr)(__ccgo_up(bp + 16)) == ptr {
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(ptr))) != int32(']') {
					libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
					return int32(m_SQLITE_ERROR)
				}
				goto done
			}
			if result < int32(-libc.Int32FromInt32(0x7f)-libc.Int32FromInt32(1)) || result > int32(m_INT8_MAX) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+341, 0)
				return int32(m_SQLITE_ERROR)
			}
			**(**Ti8)(__ccgo_up(bp + 20)) = int8(result)
			Xarray_append(tls, bp, bp+20)
			offset = offset + (int32(**(**uintptr)(__ccgo_up(bp + 16))) - int32(ptr))
			for offset < source_len {
				if _vecJsonIsSpaceX[uint8(**(**uint8)(__ccgo_up(source + uintptr(offset))))] != 0 {
					offset = offset + 1
					continue
				}
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(source + uintptr(offset)))) == int32(',') {
					offset = offset + 1
					continue
				}
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(source + uintptr(offset)))) == int32(']') {
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

func _vec_to_json(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var b Tu8
	var i Tsize_t
	var len1, rc int32
	var s, str uintptr
	var value Tf32
	var _ /* cleanup at bp+8 */ Tvector_cleanup
	var _ /* dimensions at bp+4 */ Tsize_t
	var _ /* elementType at bp+16 */ _VectorElementType
	var _ /* err at bp+12 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_, _, _, _, _, _, _ = b, i, len1, rc, s, str, value
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+4, bp+16, bp+8, bp+12)
	if rc != m_SQLITE_OK {
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 12)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 12)))
		return
	}
	str = libsqlite3.Xsqlite3_str_new(tls, libsqlite3.Xsqlite3_context_db_handle(tls, context))
	libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1671)
	i = uint32(0)
	for {
		if !(i < **(**Tsize_t)(__ccgo_up(bp + 4))) {
			break
		}
		if i != uint32(0) {
			libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1673)
		}
		if **(**_VectorElementType)(__ccgo_up(bp + 16)) == int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32) {
			value = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4))
			if ___inline_isnanf(tls, value) != 0 {
				libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1675)
			} else {
				libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1680, libc.VaList(bp+32, float64(value)))
			}
		} else {
			if **(**_VectorElementType)(__ccgo_up(bp + 16)) == int32(_SQLITE_VEC_ELEMENT_TYPE_INT8) {
				libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1683, libc.VaList(bp+32, int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i))))))
			} else {
				if **(**_VectorElementType)(__ccgo_up(bp + 16)) == int32(_SQLITE_VEC_ELEMENT_TYPE_BIT) {
					b = libc.Uint8FromInt32(libc.Int32FromUint8(**(**Tu8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i/uint32(8))))) >> (i % uint32(m___CHAR_BIT)) & int32(1))
					libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1683, libc.VaList(bp+32, libc.Int32FromUint8(b)))
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
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 8)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

const m_LDBL_DECIMAL_DIG = "DBL_DECIMAL_DIG"

const m_LDBL_DIG = "DBL_DIG"

const m_LDBL_HAS_SUBNORM = "DBL_HAS_SUBNORM"

const m_LDBL_MANT_DIG = "DBL_MANT_DIG"

const m_LDBL_MAX_10_EXP = "DBL_MAX_10_EXP"

const m_LDBL_MAX_EXP = "DBL_MAX_EXP"

const m_LDBL_MIN_10_EXP = "DBL_MIN_10_EXP"

const m_LDBL_MIN_EXP = "DBL_MIN_EXP"

const m___ARM_ARCH = 7

const m___ARM_ARCH_7A__ = 1

const m___ARM_ARCH_ISA_THUMB = 2

const m___ARM_FEATURE_COPROC = 0xf

const m___ARM_FEATURE_LDREX = 0xf

const m___ARM_FP = 0xc

const m___ARM_NEON_FP = 0x4

const m___ARM_PCS = 1

const m___ARM_VFPV2__ = 1

const m___ARM_VFPV3__ = 1

const m___arm = 1

type t__max_align_t = struct {
	F__ccgo_align [0]uint32
	F__max_align1 int64
	F__max_align2 float64
}

type t__mbstate_t = struct {
	F__ccgo_align [0]uint32
	F_mbstateL    [0]t__int64_t
	F__mbstate8   [128]uint8
}

type t__sFILE = struct {
	F__ccgo_align [0]uint32
	F_p           uintptr
	F_r           int32
	F_w           int32
	F_flags       int16
	F_file        int16
	F_bf          t__sbuf
	F_lbfsize     int32
	F_cookie      uintptr
	F_close       uintptr
	F_read        uintptr
	F_seek        uintptr
	F_write       uintptr
	F_ub          t__sbuf
	F_up          uintptr
	F_ur          int32
	F_ubuf        [3]uint8
	F_nbuf        [1]uint8
	F_lb          t__sbuf
	F_blksize     int32
	F_offset      Tfpos_t
	F_fl_mutex    uintptr
	F_fl_owner    uintptr
	F_fl_count    int32
	F_orientation int32
	F_mbstate     t__mbstate_t
	F_flags2      int32
	F__ccgo_pad26 [4]byte
}

type t__vm_paddr_t = uint32
