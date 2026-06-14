// Code generated for netbsd/amd64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build netbsd && amd64

package vec

import (
	"modernc.org/libc"
	libsqlite3 "modernc.org/sqlite/lib"
)

type TFILE = struct {
	F_p         uintptr
	F_r         int32
	F_w         int32
	F_flags     uint16
	F_file      int16
	F_bf        t__sbuf
	F_lbfsize   int32
	F_cookie    uintptr
	F_close     uintptr
	F_read      uintptr
	F_seek      uintptr
	F_write     uintptr
	F_ext       t__sbuf
	F_up        uintptr
	F_ur        int32
	F_ubuf      [3]uint8
	F_nbuf      [1]uint8
	F_flush     uintptr
	F_lb_unused [8]int8
	F_blksize   int32
	F_offset    t__off_t
}

type Taccmode_t = uint32

type Tclock_t = uint32

type Tdev_t = uint64

type Tfd_set = struct {
	Ffds_bits [8]t__fd_mask
}

type Tfpos_t = struct {
	F_pos         t__off_t
	F_mbstate_in  t__mbstate_t
	F_mbstate_out t__mbstate_t
}

type Tin_addr_t = uint32

type Tin_port_t = uint16

type Tkauth_cred_t = uintptr

type Tlonglong_t = int64

type Tlwpid_t = int32

type Tmqd_t = int32

type Tpri_t = int32

type Tpsetid_t = int32

type Tpthread_attr_t = struct {
	Fpta_magic   uint32
	Fpta_flags   int32
	Fpta_private uintptr
}

type Tpthread_barrier_t = struct {
	Fptb_magic      uint32
	Fptb_lock       Tpthread_spin_t
	Fptb_waiters    Tpthread_queue_t
	Fptb_initcount  uint32
	Fptb_curcount   uint32
	Fptb_generation uint32
	Fptb_private    uintptr
}

type Tpthread_barrierattr_t = struct {
	Fptba_magic   uint32
	Fptba_private uintptr
}

type Tpthread_cond_t = struct {
	Fptc_magic   uint32
	Fptc_lock    t__pthread_spin_t
	Fptc_waiters uintptr
	Fptc_spare   uintptr
	Fptc_mutex   uintptr
	Fptc_private uintptr
}

type Tpthread_condattr_t = struct {
	Fptca_magic   uint32
	Fptca_private uintptr
}

type Tpthread_key_t = int32

type Tpthread_mutex_t = struct {
	Fptm_magic      uint32
	Fptm_errorcheck t__pthread_spin_t
	Fptm_pad1       [3]Tuint8_t
	F__ccgo3_8      struct {
		Fptm_unused  [0]t__pthread_spin_t
		Fptm_ceiling uint8
	}
	Fptm_pad2     [3]Tuint8_t
	Fptm_owner    Tpthread_t
	Fptm_waiters  uintptr
	Fptm_recursed uint32
	Fptm_spare2   uintptr
}

type Tpthread_mutexattr_t = struct {
	Fptma_magic   uint32
	Fptma_private uintptr
}

type Tpthread_once_t = struct {
	Fpto_mutex Tpthread_mutex_t
	Fpto_done  int32
}

type Tpthread_queue_struct_t = struct {
	Fptqh_first uintptr
	Fptqh_last  uintptr
}

type Tpthread_queue_t = struct {
	Fptqh_first uintptr
	Fptqh_last  uintptr
}

type Tpthread_rwlock_t = struct {
	Fptr_magic     uint32
	Fptr_interlock t__pthread_spin_t
	Fptr_rblocked  Tpthread_queue_t
	Fptr_wblocked  Tpthread_queue_t
	Fptr_nreaders  uint32
	Fptr_owner     Tpthread_t
	Fptr_private   uintptr
}

type Tpthread_rwlockattr_t = struct {
	Fptra_magic   uint32
	Fptra_private uintptr
}

type Tpthread_spin_t = uint8

type Tpthread_spinlock_t = struct {
	Fpts_magic uint32
	Fpts_spin  t__pthread_spin_t
	Fpts_flags int32
}

type Tpthread_t = uintptr

type Tqaddr_t = uintptr

type Tsuseconds_t = int32

type Tswblk_t = int32

type Tu_longlong_t = uint64

const _fdlibm_ieee = -1

const _fdlibm_posix = 2

const _fdlibm_svid = 0

const _fdlibm_xopen = 1

type _fdversion = int32

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
			if result < int64(-libc.Int32FromInt32(m___INT8_MAX__)-libc.Int32FromInt32(1)) || result > int64(m___INT8_MAX__) {
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

type accmode_t = Taccmode_t

type fd_set = Tfd_set

type fts5_api = Tfts5_api

//TODO "sys_errlist" // errno.h:61:19:

type in_addr_t = Tin_addr_t

type in_port_t = Tin_port_t

type kauth_cred_t = Tkauth_cred_t

type longlong_t = Tlonglong_t

type lwpid_t = Tlwpid_t

const m_BIG_ENDIAN = 4321

const m_CHILD_MAX = 160

const m_EBADMSG = 88

const m_ECANCELED = 87

const m_EILSEQ = 85

const m_ELAST = 98

const m_EMULTIHOP = 94

const m_ENODATA = 89

const m_ENOLINK = 95

const m_ENOSR = 90

const m_ENOSTR = 91

const m_ENOTRECOVERABLE = 98

const m_ENOTSUP = 86

const m_EOWNERDEAD = 97

const m_EPROTO = 96

const m_ETIME = 92

const m_FD_SETSIZE = 256

const m_FPARSELN_UNESCALL = 0x0f

const m_FPARSELN_UNESCCOMM = 0x04

const m_FPARSELN_UNESCCONT = 0x02

const m_FPARSELN_UNESCESC = 0x01

const m_FPARSELN_UNESCREST = 0x08

const m_FP_ILOGB0 = "INT_MIN"

const m_FP_INFINITE = 0x00

const m_FP_NAN = 0x01

const m_FP_NORMAL = 0x02

const m_FP_SUBNORMAL = 0x03

const m_FP_ZERO = 0x04

const m_GID_MAX = 2147483647

const m_HN_AUTOSCALE = 0x20

const m_HN_B = 0x04

const m_HN_DECIMAL = 0x01

const m_HN_DIVISOR_1000 = 0x08

const m_HN_GETSCALE = 0x10

const m_HN_NOSPACE = 0x02

const m_INT16_MAX = "__INT16_MAX__"

const m_INT32_MAX = "__INT32_MAX__"

const m_INT64_MAX = "__INT64_MAX__"

const m_INT8_MAX = "__INT8_MAX__"

const m_INT_FAST16_MAX = "__INT_FAST16_MAX__"

const m_INT_FAST32_MAX = "__INT_FAST32_MAX__"

const m_INT_FAST64_MAX = "__INT_FAST64_MAX__"

const m_INT_FAST8_MAX = "__INT_FAST8_MAX__"

const m_INT_LEAST16_MAX = "__INT_LEAST16_MAX__"

const m_INT_LEAST32_MAX = "__INT_LEAST32_MAX__"

const m_INT_LEAST64_MAX = "__INT_LEAST64_MAX__"

const m_INT_LEAST8_MAX = "__INT_LEAST8_MAX__"

const m_LITTLE_ENDIAN = 1234

const m_LOGIN_NAME_MAX = 17

const m_L_cuserid = 9

const m_MB_LEN_MAX = 32

const m_NAME_MAX = 511

const m_NBBY = 8

const m_NFDBITS = "__NFDBITS"

const m_OPEN_MAX = 128

const m_PDP_ENDIAN = 3412

const m_PRIXMAX = "lX"

const m_PRIdMAX = "ld"

const m_PRIiMAX = "li"

const m_PRIoMAX = "lo"

const m_PRIuMAX = "lu"

const m_PRIxMAX = "lx"

const m_PTHREAD_DESTRUCTOR_ITERATIONS = "_POSIX_THREAD_DESTRUCTOR_ITERATIONS"

const m_PTHREAD_KEYS_MAX = 256

const m_PTHREAD_STACK_MIN = 4096

const m_PTHREAD_THREADS_MAX = "_POSIX_THREAD_THREADS_MAX"

const m_RANDOM_MAX = 0x7fffffff

const m_SCNdMAX = "ld"

const m_SCNiMAX = "li"

const m_SCNoMAX = "lo"

const m_SCNuMAX = "lu"

const m_SCNxMAX = "lx"

const m_SIG_ATOMIC_MAX = "__SIG_ATOMIC_MAX__"

const m_SSIZE_MIN = "LONG_MIN"

const m_UID_MAX = 2147483647

const m_UINT16_MAX = "__UINT16_MAX__"

const m_UINT32_MAX = "__UINT32_MAX__"

const m_UINT64_MAX = "__UINT64_MAX__"

const m_UINT8_MAX = "__UINT8_MAX__"

const m_UINT_FAST16_MAX = "__UINT_FAST16_MAX__"

const m_UINT_FAST32_MAX = "__UINT_FAST32_MAX__"

const m_UINT_FAST64_MAX = "__UINT_FAST64_MAX__"

const m_UINT_FAST8_MAX = "__UINT_FAST8_MAX__"

const m_UINT_LEAST16_MAX = "__UINT_LEAST16_MAX__"

const m_UINT_LEAST32_MAX = "__UINT_LEAST32_MAX__"

const m_UINT_LEAST64_MAX = "__UINT_LEAST64_MAX__"

const m_UINT_LEAST8_MAX = "__UINT_LEAST8_MAX__"

const m_WCHAR_MAX = 0x7fffffff

const m_WINT_MAX = 0x7fffffff

const m__BSD_MBSTATE_T_ = "__mbstate_t"

const m__BSD_WCTRANS_T_ = "__wctrans_t"

const m__BSD_WCTYPE_T_ = "__wctype_t"

const m__BSD_WINT_T_ = "__WINT_TYPE__"

const m__FLOAT_IEEE754 = 1

const m__FP_HIMD = 0xff

const m__FP_LOMD = 0x80

const m__GETGR_R_SIZE_MAX = 1024

const m__GETPW_R_SIZE_MAX = 1024

const m__IEEE_ = "fdlibm_ieee"

const m__LIB_VERSION = "_fdlib_version"

const m__NETBSD_SOURCE = 1

const m__POSIX_ = "fdlibm_posix"

const m__POSIX_REALTIME_SIGNALS = 200112

const m__PT_BARRIERATTR_DEAD = 0xDEAD0808

const m__PT_BARRIERATTR_MAGIC = 0x88880808

const m__PT_BARRIER_DEAD = 0xDEAD0008

const m__PT_BARRIER_MAGIC = 0x88880008

const m__PT_CONDATTR_DEAD = 0xDEAD0006

const m__PT_CONDATTR_MAGIC = 0x66660006

const m__PT_COND_DEAD = 0xDEAD0005

const m__PT_COND_MAGIC = 0x55550005

const m__PT_MUTEXATTR_DEAD = 0xDEAD0004

const m__PT_MUTEXATTR_MAGIC = 0x44440004

const m__PT_MUTEX_DEAD = 0xDEAD0003

const m__PT_MUTEX_MAGIC = 0x33330003

const m__PT_RWLOCKATTR_DEAD = 0xDEAD0909

const m__PT_RWLOCKATTR_MAGIC = 0x99990909

const m__PT_RWLOCK_DEAD = 0xDEAD0009

const m__PT_RWLOCK_MAGIC = 0x99990009

const m__PT_SPINLOCK_DEAD = 0xDEAD0007

const m__PT_SPINLOCK_MAGIC = 0x77770007

const m__PT_SPINLOCK_PSHARED = 0x00000001

const m__SVID_ = "fdlibm_svid"

const m__XOPEN_ = "fdlibm_xopen"

const m__XOPEN_NAME_MAX = 256

const m___BEGIN_DECLS = "__BEGIN_PUBLIC_DECLS"

const m___BYTE_SWAP_U16_VARIABLE = "__byte_swap_u16_variable"

const m___BYTE_SWAP_U32_VARIABLE = "__byte_swap_u32_variable"

const m___BYTE_SWAP_U64_VARIABLE = "__byte_swap_u64_variable"

const m___END_DECLS = "__END_PUBLIC_DECLS"

const m___GNUC_MINOR__ = 5

const m___GNUC__ = 10

const m___GXX_ABI_VERSION = 1014

const m___INT_FAST8_MAX__ = 0x7fffffff

const m___INT_FAST8_TYPE__ = "int"

const m___INT_FAST8_WIDTH__ = 32

const m___NFDBITS = 32

const m___NFDSHIFT = 5

const m___NetBSD__ = 1

const m___PRIuBIT = "PRIuMAX"

const m___PRIuBITS = "__PRIuBIT"

const m___PRIxBIT = "PRIxMAX"

const m___PRIxBITS = "__PRIxBIT"

const m___SIMPLELOCK_LOCKED = 1

const m___SIMPLELOCK_UNLOCKED = 0

const m___UINT_FAST8_MAX__ = 0xffffffff

const m___VERSION__ = "10.5.0"

const m___WINT_MAX__ = 0x7fffffff

const m___assert_function__ = "__func__"

const m___debugused = "__unused"

const m___diagused = "__unused"

const m___pthread_volatile = "volatile"

const m___syslog_attribute__ = 1

const m___tune_nocona__ = 1

const m___weakref_visible = "static"

const m_accmode_t = "__accmode_t"

const m_caddr_t = "__caddr_t"

const m_devmajor_t = "__devmajor_t"

const m_devminor_t = "__devminor_t"

const m_fd_mask = "__fd_mask"

const m_fsblkcnt_t = "__fsblkcnt_t"

const m_fsfilcnt_t = "__fsfilcnt_t"

const m_gid_t = "__gid_t"

const m_in_addr_t = "__in_addr_t"

const m_in_port_t = "__in_port_t"

const m_mode_t = "__mode_t"

const m_off_t = "__off_t"

const m_pid_t = "__pid_t"

const m_uid_t = "__uid_t"

const m_va_arg = "__builtin_va_arg"

type mqd_t = Tmqd_t

type pri_t = Tpri_t

type psetid_t = Tpsetid_t

type pthread_barrier_t = Tpthread_barrier_t

type pthread_barrierattr_t = Tpthread_barrierattr_t

type pthread_cond_t = Tpthread_cond_t

type pthread_condattr_t = Tpthread_condattr_t

type pthread_key_t = Tpthread_key_t

type pthread_mutex_t = Tpthread_mutex_t

type pthread_mutexattr_t = Tpthread_mutexattr_t

type pthread_once_t = Tpthread_once_t

type pthread_queue_struct_t = Tpthread_queue_struct_t

type pthread_queue_t = Tpthread_queue_t

type pthread_rwlock_t = Tpthread_rwlock_t

type pthread_rwlockattr_t = Tpthread_rwlockattr_t

type pthread_spin_t = Tpthread_spin_t

type pthread_spinlock_t = Tpthread_spinlock_t

type pthread_t = Tpthread_t

type qaddr_t = Tqaddr_t

type swblk_t = Tswblk_t

type t__accmode_t = uint32

type t__caddr_t = uintptr

type t__cpu_simple_lock_nv_t = uint8

type t__cpu_simple_lock_t = uint8

type t__devmajor_t = int32

type t__devminor_t = int32

type t__double_u = struct {
	F__val   [0]float64
	F__dummy [8]uint8
}

type t__fd_mask = uint32

type t__float_u = struct {
	F__val   [0]float32
	F__dummy [4]uint8
}

type t__long_double_u = struct {
	F__val   [0]float64
	F__dummy [8]uint8
}

type t__mbstate_t = struct {
	F__mbstate8  [0][128]int8
	F__mbstateL  t__int64_t
	F__ccgo_pad2 [120]byte
}

type t__pthread_attr_st = struct {
	Fpta_magic   uint32
	Fpta_flags   int32
	Fpta_private uintptr
}

type t__pthread_barrier_st = struct {
	Fptb_magic      uint32
	Fptb_lock       Tpthread_spin_t
	Fptb_waiters    Tpthread_queue_t
	Fptb_initcount  uint32
	Fptb_curcount   uint32
	Fptb_generation uint32
	Fptb_private    uintptr
}

type t__pthread_barrierattr_st = struct {
	Fptba_magic   uint32
	Fptba_private uintptr
}

type t__pthread_cond_st = struct {
	Fptc_magic   uint32
	Fptc_lock    t__pthread_spin_t
	Fptc_waiters uintptr
	Fptc_spare   uintptr
	Fptc_mutex   uintptr
	Fptc_private uintptr
}

type t__pthread_condattr_st = struct {
	Fptca_magic   uint32
	Fptca_private uintptr
}

type t__pthread_mutex_st = struct {
	Fptm_magic      uint32
	Fptm_errorcheck t__pthread_spin_t
	Fptm_pad1       [3]Tuint8_t
	F__ccgo3_8      struct {
		Fptm_unused  [0]t__pthread_spin_t
		Fptm_ceiling uint8
	}
	Fptm_pad2     [3]Tuint8_t
	Fptm_owner    Tpthread_t
	Fptm_waiters  uintptr
	Fptm_recursed uint32
	Fptm_spare2   uintptr
}

type t__pthread_mutexattr_st = struct {
	Fptma_magic   uint32
	Fptma_private uintptr
}

type t__pthread_once_st = Tpthread_once_t

type t__pthread_rwlock_st = struct {
	Fptr_magic     uint32
	Fptr_interlock t__pthread_spin_t
	Fptr_rblocked  Tpthread_queue_t
	Fptr_wblocked  Tpthread_queue_t
	Fptr_nreaders  uint32
	Fptr_owner     Tpthread_t
	Fptr_private   uintptr
}

type t__pthread_rwlockattr_st = struct {
	Fptra_magic   uint32
	Fptra_private uintptr
}

type t__pthread_spin_t = uint8

type t__pthread_spinlock_st = Tpthread_spinlock_t

type t__sFILE = TFILE

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

type t__sfpos = Tfpos_t

type u_longlong_t = Tu_longlong_t
