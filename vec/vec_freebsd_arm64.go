// Code generated for freebsd/arm64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build freebsd && arm64

package vec

import (
	"unsafe"

	"modernc.org/libc"
	libsqlite3 "modernc.org/sqlite/lib"
)

func Xvec0Filter_knn_chunks_iter(tls *libc.TLS, p uintptr, stmtChunks uintptr, vector_column uintptr, vectorColumnIdx int32, arrayRowidsIn uintptr, aMetadataIn uintptr, idxStr uintptr, argc int32, argv uintptr, queryVector uintptr, k Ti64, out_topk_rowids uintptr, out_topk_distances uintptr, out_used uintptr) (r int32) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var b, bTaken, baseVectors, base_i, base_i1, base_i2, bmMetadata, bmRowids, chunkRowids, chunkValidity, chunk_distances, chunk_topk_idxs, in, tmp_topk_distances, tmp_topk_rowids, topk_distances, topk_rowids, v1 uintptr
	var baseVectorsSize, chunk_id, currentBaseVectorsSize, expectedBaseVectorsSize, k_used, rowidsSize, validitySize Ti64
	var hasDistanceConstraints, hasMetadataFilters, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, idx, idx1, idx2, idxStrLength, metadata_idx, numValueEntries, operator, rc, v4 int32
	var kind, kind1, kind2 uint8
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
	baseVectorsSize = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(vector_column))))
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
		libc.Xmemset(tls, chunk_distances, 0, libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(4))
		libc.Xmemset(tls, chunk_topk_idxs, 0, uint64(libc.Uint64FromInt64(k)*uint64(4)))
		Xbitmap_clear(tls, b, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		chunk_id = libsqlite3.Xsqlite3_column_int64(tls, stmtChunks, 0)
		chunkValidity = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(1))
		validitySize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(1)))
		if validitySize != int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT)) {
			// IMP: V05271_22109
			Xvtab_set_error(tls, p, __ccgo_ts+9715, libc.VaList(bp+168, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m___CHAR_BIT), validitySize))
			rc = int32(m_SQLITE_ERROR)
			goto cleanup
		}
		chunkRowids = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(2))
		rowidsSize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(2)))
		if libc.Uint64FromInt64(rowidsSize) != uint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8)) {
			// IMP: V02796_19635
			Xvtab_set_error(tls, p, __ccgo_ts+9777, 0)
			Xvtab_set_error(tls, p, __ccgo_ts+9803, libc.VaList(bp+168, libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8), rowidsSize))
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
		expectedBaseVectorsSize = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(vector_column))))
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
				kind1 = **(**uint8)(__ccgo_up(idxStr + uintptr(idx1+0)))
				if libc.Int32FromUint8(kind1) != int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) {
					goto _5
				}
				metadata_idx = libc.Int32FromUint8(**(**uint8)(__ccgo_up(idxStr + uintptr(idx1+int32(1))))) - int32('A')
				operator = libc.Int32FromUint8(**(**uint8)(__ccgo_up(idxStr + uintptr(idx1+int32(2)))))
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
				base_i = baseVectors + uintptr(libc.Uint64FromInt32(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)*4
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
				base_i1 = baseVectors + uintptr(libc.Uint64FromInt32(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)
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
				base_i2 = baseVectors + uintptr(libc.Uint64FromInt32(i3)*((*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions/libc.Uint64FromInt32(m___CHAR_BIT)))
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
				kind2 = **(**uint8)(__ccgo_up(idxStr + uintptr(idx2+0)))
				// TODO casts f64 to f32, is that a problem?
				target = float32(libsqlite3.Xsqlite3_value_double(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i4)*8))))
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

func Xvec0Update_Delete_ClearValidity(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var brc, rc, validityOffset int32
	var mask uint8
	var _ /* blobChunksValidity at bp+0 */ uintptr
	var _ /* bx at bp+8 */ uint8
	var _ /* result at bp+9 */ uint8
	_, _, _, _ = brc, mask, rc, validityOffset
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	validityOffset = libc.Int32FromUint64(chunk_offset / uint64(m___CHAR_BIT))
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
	if !(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 8)))>>(chunk_offset%libc.Uint64FromInt32(m___CHAR_BIT)) != 0) {
		// IMP: V21193_05263
		rc = int32(m_SQLITE_ERROR)
		Xvtab_set_error(tls, p, __ccgo_ts+13831, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		goto cleanup
	}
	mask = libc.Uint8FromInt32(^(libc.Int32FromInt32(1) << (chunk_offset % libc.Uint64FromInt32(m___CHAR_BIT))))
	**(**uint8)(__ccgo_up(bp + 9)) = libc.Uint8FromInt32(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 8))) & libc.Int32FromUint8(mask))
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

func _distance_cosine_float(tls *libc.TLS, pVect1v uintptr, pVect2v uintptr, qty_ptr uintptr) (r Tf32) {
	var aMag, bMag, dot Tf32
	var i, qty Tsize_t
	var pVect1, pVect2 uintptr
	_, _, _, _, _, _, _ = aMag, bMag, dot, i, pVect1, pVect2, qty
	pVect1 = pVect1v
	pVect2 = pVect2v
	qty = **(**Tsize_t)(__ccgo_up(qty_ptr))
	dot = libc.Float32FromInt32(0)
	aMag = libc.Float32FromInt32(0)
	bMag = libc.Float32FromInt32(0)
	i = uint64(0)
	for {
		if !(i < qty) {
			break
		}
		dot = dot + **(**Tf32)(__ccgo_up(pVect1))***(**Tf32)(__ccgo_up(pVect2))
		aMag = aMag + **(**Tf32)(__ccgo_up(pVect1))***(**Tf32)(__ccgo_up(pVect1))
		bMag = bMag + **(**Tf32)(__ccgo_up(pVect2))***(**Tf32)(__ccgo_up(pVect2))
		pVect1 += 4
		pVect2 += 4
		goto _1
	_1:
		;
		i = i + 1
	}
	return float32(libc.Float64FromInt32(1) - float64(dot)/(libc.Xsqrt(tls, float64(aMag))*libc.Xsqrt(tls, float64(bMag))))
}

func _distance_cosine_int8(tls *libc.TLS, pA uintptr, pB uintptr, pD uintptr) (r Tf32) {
	var a, b uintptr
	var aMag, bMag, dot Tf32
	var d, i Tsize_t
	_, _, _, _, _, _, _ = a, aMag, b, bMag, d, dot, i
	a = pA
	b = pB
	d = **(**Tsize_t)(__ccgo_up(pD))
	dot = libc.Float32FromInt32(0)
	aMag = libc.Float32FromInt32(0)
	bMag = libc.Float32FromInt32(0)
	i = uint64(0)
	for {
		if !(i < d) {
			break
		}
		dot = dot + float32(int32(**(**Ti8)(__ccgo_up(a)))*int32(**(**Ti8)(__ccgo_up(b))))
		aMag = aMag + float32(int32(**(**Ti8)(__ccgo_up(a)))*int32(**(**Ti8)(__ccgo_up(a))))
		bMag = bMag + float32(int32(**(**Ti8)(__ccgo_up(b)))*int32(**(**Ti8)(__ccgo_up(b))))
		a = a + 1
		b = b + 1
		goto _1
	_1:
		;
		i = i + 1
	}
	return float32(libc.Float64FromInt32(1) - float64(dot)/(libc.Xsqrt(tls, float64(aMag))*libc.Xsqrt(tls, float64(bMag))))
}

func _fvec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var blob, buf, ptr, source uintptr
	var bytes, i, offset, rc, source_len, value_type int32
	var result float64
	var _ /* endptr at bp+32 */ uintptr
	var _ /* res at bp+40 */ Tf32
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
		if libc.Uint64FromInt32(bytes)%uint64(4) != uint64(0) {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+86, libc.VaList(bp+56, uint64(4), bytes))
			return int32(m_SQLITE_ERROR)
		}
		buf = libsqlite3.Xsqlite3_malloc(tls, bytes)
		if !(buf != 0) {
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+156, 0)
			return int32(m_SQLITE_NOMEM)
		}
		libc.Xmemcpy(tls, buf, blob, libc.Uint64FromInt32(bytes))
		**(**uintptr)(__ccgo_up(vector)) = buf
		**(**Tsize_t)(__ccgo_up(dimensions)) = libc.Uint64FromInt32(bytes) / uint64(4)
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
			result = libc.Xstrtod(tls, ptr, bp+32)
			if **(**int32)(__ccgo_up(libc.X__error(tls))) != 0 && result == libc.Float64FromInt32(0) || **(**int32)(__ccgo_up(libc.X__error(tls))) == int32(m_ERANGE) && (result == libc.X__builtin_huge_val(tls) || result == -libc.X__builtin_huge_val(tls)) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if **(**uintptr)(__ccgo_up(bp + 32)) == ptr {
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(ptr))) != int32(']') {
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
			result = libc.Xstrtol(tls, ptr, bp+32, int32(10))
			if **(**int32)(__ccgo_up(libc.X__error(tls))) != 0 && result == 0 || **(**int32)(__ccgo_up(libc.X__error(tls))) == int32(m_ERANGE) && (result == int64(0x7fffffffffffffff) || result == -libc.Int64FromInt64(0x7fffffffffffffff)-libc.Int64FromInt32(1)) {
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				return int32(m_SQLITE_ERROR)
			}
			if **(**uintptr)(__ccgo_up(bp + 32)) == ptr {
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(ptr))) != int32(']') {
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

func _l2_sqr_float(tls *libc.TLS, pVect1v uintptr, pVect2v uintptr, qty_ptr uintptr) (r Tf32) {
	var i, qty Tsize_t
	var pVect1, pVect2 uintptr
	var res, t Tf32
	_, _, _, _, _, _ = i, pVect1, pVect2, qty, res, t
	pVect1 = pVect1v
	pVect2 = pVect2v
	qty = **(**Tsize_t)(__ccgo_up(qty_ptr))
	res = libc.Float32FromInt32(0)
	i = uint64(0)
	for {
		if !(i < qty) {
			break
		}
		t = **(**Tf32)(__ccgo_up(pVect1)) - **(**Tf32)(__ccgo_up(pVect2))
		pVect1 += 4
		pVect2 += 4
		res = res + t*t
		goto _1
	_1:
		;
		i = i + 1
	}
	return float32(libc.Xsqrt(tls, float64(res)))
}

func _l2_sqr_int8(tls *libc.TLS, pA uintptr, pB uintptr, pD uintptr) (r Tf32) {
	var a, b uintptr
	var d, i Tsize_t
	var res, t Tf32
	_, _, _, _, _, _ = a, b, d, i, res, t
	a = pA
	b = pB
	d = **(**Tsize_t)(__ccgo_up(pD))
	res = libc.Float32FromInt32(0)
	i = uint64(0)
	for {
		if !(i < d) {
			break
		}
		t = float32(int32(**(**Ti8)(__ccgo_up(a))) - int32(**(**Ti8)(__ccgo_up(b))))
		a = a + 1
		b = b + 1
		res = res + t*t
		goto _1
	_1:
		;
		i = i + 1
	}
	return float32(libc.Xsqrt(tls, float64(res)))
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
	outSize = libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8)) * uint64(4))
	out = libsqlite3.Xsqlite3_malloc(tls, outSize)
	if !(out != 0) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
		libsqlite3.Xsqlite3_result_error_code(tls, context, int32(m_SQLITE_NOMEM))
		return
	}
	libc.Xmemset(tls, out, 0, libc.Uint64FromInt32(outSize))
	v = **(**uintptr)(__ccgo_up(bp))
	norm = libc.Float32FromInt32(0)
	i = uint64(0)
	for {
		if !(i < **(**Tsize_t)(__ccgo_up(bp + 8))) {
			break
		}
		norm = norm + **(**Tf32)(__ccgo_up(v + uintptr(i)*4))***(**Tf32)(__ccgo_up(v + uintptr(i)*4))
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
	libsqlite3.Xsqlite3_result_blob(tls, context, out, libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8))*uint64(4)), __ccgo_fp(libsqlite3.Xsqlite3_free))
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

const m_LDBL_DECIMAL_DIG = 36

const m_LDBL_DIG = 33

const m_LDBL_EPSILON = 1.925929944387235853055977942584927319e-34

const m_LDBL_MANT_DIG = 113

const m_LDBL_MAX = "1.189731495357231765085759326628007016E+4932"

const m_LDBL_MIN = 3.362103143112093506262677817321752603e-4932

const m_LDBL_TRUE_MIN = 6.475175119438025110924438958227646552e-4966

type t__mbstate_t = struct {
	F_mbstateL  [0]t__int64_t
	F__mbstate8 [128]uint8
}
