// Code generated for netbsd/amd64 by 'generator -absolute-paths -keep-object-files -positions --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build netbsd && amd64

package vec

import (
	"math"
	"reflect"
	"unsafe"

	"modernc.org/libc"
	libsqlite3 "modernc.org/sqlite/lib"
)

var _ = math.Pi
var _ reflect.Type
var _ unsafe.Pointer

// /usr/include/sys/syslimits.h:63:9:
const m_BC_BASE_MAX = "INT_MAX"

// /usr/include/sys/syslimits.h:64:9:
const m_BC_DIM_MAX = 65535

// /usr/include/sys/syslimits.h:65:9:
const m_BC_SCALE_MAX = "INT_MAX"

// /usr/include/sys/syslimits.h:66:9:
const m_BC_STRING_MAX = "INT_MAX"

// /usr/include/sys/endian.h:103:9:
const m_BIG_ENDIAN = 4321

// /usr/include/stdio.h:183:9:
const m_BUFSIZ = 1024

// /usr/include/sys/endian.h:105:9:
const m_BYTE_ORDER = "_BYTE_ORDER"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11343:9:
const m_CARRAY_BLOB = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11341:9:
const m_CARRAY_DOUBLE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11339:9:
const m_CARRAY_INT32 = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11340:9:
const m_CARRAY_INT64 = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11342:9:
const m_CARRAY_TEXT = 3

// /usr/include/limits.h:113:9:
const m_CHARCLASS_NAME_MAX = 14

// /usr/include/machine/limits.h:41:9:
const m_CHAR_BIT = 8

// /usr/include/limits.h:150:10:
const m_CHAR_MAX = "SCHAR_MAX"

// /usr/include/limits.h:149:10:
const m_CHAR_MIN = "SCHAR_MIN"

// /usr/include/sys/syslimits.h:47:9:
const m_CHILD_MAX = 160

// /usr/include/sys/syslimits.h:67:9:
const m_COLL_WEIGHTS_MAX = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:93:9:
const m_COMPILER_SUPPORTS_VTAB_IN = 1

// /usr/include/machine/limits.h:85:9:
const m_DBL_DIG = "__DBL_DIG__"

// /usr/include/sys/float_ieee754.h:87:9:
const m_DBL_EPSILON = "__DBL_EPSILON__"

// /usr/include/sys/float_ieee754.h:86:9:
const m_DBL_MANT_DIG = "__DBL_MANT_DIG__"

// /usr/include/machine/limits.h:86:9:
const m_DBL_MAX = "__DBL_MAX__"

// /usr/include/sys/float_ieee754.h:94:9:
const m_DBL_MAX_10_EXP = "__DBL_MAX_10_EXP__"

// /usr/include/sys/float_ieee754.h:92:9:
const m_DBL_MAX_EXP = "__DBL_MAX_EXP__"

// /usr/include/machine/limits.h:87:9:
const m_DBL_MIN = "__DBL_MIN__"

// /usr/include/sys/float_ieee754.h:91:9:
const m_DBL_MIN_10_EXP = "__DBL_MIN_10_EXP__"

// /usr/include/sys/float_ieee754.h:89:9:
const m_DBL_MIN_EXP = "__DBL_MIN_EXP__"

// /usr/include/x86/float.h:26:9:
const m_DECIMAL_DIG = 21

// /usr/include/math.h:225:9:
const m_DOMAIN = 1

// /usr/include/sys/errno.h:48:9:
const m_E2BIG = 7

// /usr/include/sys/errno.h:55:9:
const m_EACCES = 13

// /usr/include/sys/errno.h:97:9:
const m_EADDRINUSE = 48

// /usr/include/sys/errno.h:98:9:
const m_EADDRNOTAVAIL = 49

// /usr/include/sys/errno.h:96:9:
const m_EAFNOSUPPORT = 47

// /usr/include/sys/errno.h:81:9:
const m_EAGAIN = 35

// /usr/include/sys/errno.h:84:9:
const m_EALREADY = 37

// /usr/include/sys/errno.h:140:9:
const m_EAUTH = 80

// /usr/include/sys/errno.h:50:9:
const m_EBADF = 9

// /usr/include/sys/errno.h:159:9:
const m_EBADMSG = 88

// /usr/include/sys/errno.h:130:9:
const m_EBADRPC = 72

// /usr/include/sys/errno.h:58:9:
const m_EBUSY = 16

// /usr/include/sys/errno.h:156:9:
const m_ECANCELED = 87

// /usr/include/sys/errno.h:51:9:
const m_ECHILD = 10

// /usr/include/sys/errno.h:104:9:
const m_ECONNABORTED = 53

// /usr/include/sys/errno.h:112:9:
const m_ECONNREFUSED = 61

// /usr/include/sys/errno.h:105:9:
const m_ECONNRESET = 54

// /usr/include/sys/errno.h:52:9:
const m_EDEADLK = 11

// /usr/include/sys/errno.h:88:9:
const m_EDESTADDRREQ = 39

// /usr/include/sys/errno.h:77:9:
const m_EDOM = 33

// /usr/include/sys/errno.h:125:9:
const m_EDQUOT = 69

// /usr/include/sys/errno.h:59:9:
const m_EEXIST = 17

// /usr/include/sys/errno.h:56:9:
const m_EFAULT = 14

// /usr/include/sys/errno.h:69:9:
const m_EFBIG = 27

// /usr/include/sys/errno.h:139:9:
const m_EFTYPE = 79

// /usr/include/sys/errno.h:118:9:
const m_EHOSTDOWN = 64

// /usr/include/sys/errno.h:119:9:
const m_EHOSTUNREACH = 65

// /usr/include/sys/errno.h:144:9:
const m_EIDRM = 82

// /usr/include/sys/errno.h:149:9:
const m_EILSEQ = 85

// /usr/include/sys/errno.h:83:9:
const m_EINPROGRESS = 36

// /usr/include/sys/errno.h:45:9:
const m_EINTR = 4

// /usr/include/sys/errno.h:64:9:
const m_EINVAL = 22

// /usr/include/sys/errno.h:46:9:
const m_EIO = 5

// /usr/include/sys/errno.h:107:9:
const m_EISCONN = 56

// /usr/include/sys/errno.h:63:9:
const m_EISDIR = 21

// /usr/include/sys/errno.h:179:9:
const m_ELAST = 98

// /usr/include/sys/errno.h:114:9:
const m_ELOOP = 62

// /usr/include/sys/errno.h:66:9:
const m_EMFILE = 24

// /usr/include/sys/errno.h:73:9:
const m_EMLINK = 31

// /usr/include/sys/errno.h:89:9:
const m_EMSGSIZE = 40

// /usr/include/sys/errno.h:171:9:
const m_EMULTIHOP = 94

// /usr/include/sys/errno.h:115:9:
const m_ENAMETOOLONG = 63

// /usr/include/sys/errno.h:141:9:
const m_ENEEDAUTH = 81

// /usr/include/sys/errno.h:101:9:
const m_ENETDOWN = 50

// /usr/include/sys/errno.h:103:9:
const m_ENETRESET = 52

// /usr/include/sys/errno.h:102:9:
const m_ENETUNREACH = 51

// /usr/include/sys/errno.h:65:9:
const m_ENFILE = 23

// /usr/include/sys/errno.h:168:9:
const m_ENOATTR = 93

// /usr/include/sys/errno.h:106:9:
const m_ENOBUFS = 55

// /usr/include/sys/errno.h:162:9:
const m_ENODATA = 89

// /usr/include/sys/errno.h:61:9:
const m_ENODEV = 19

// /usr/include/sys/errno.h:43:9:
const m_ENOENT = 2

// /usr/include/sys/errno.h:49:9:
const m_ENOEXEC = 8

// /usr/include/sys/errno.h:136:9:
const m_ENOLCK = 77

// /usr/include/sys/errno.h:172:9:
const m_ENOLINK = 95

// /usr/include/sys/errno.h:54:9:
const m_ENOMEM = 12

// /usr/include/sys/errno.h:145:9:
const m_ENOMSG = 83

// /usr/include/sys/errno.h:91:9:
const m_ENOPROTOOPT = 42

// /usr/include/sys/errno.h:70:9:
const m_ENOSPC = 28

// /usr/include/sys/errno.h:163:9:
const m_ENOSR = 90

// /usr/include/sys/errno.h:164:9:
const m_ENOSTR = 91

// /usr/include/sys/errno.h:137:9:
const m_ENOSYS = 78

// /usr/include/sys/errno.h:57:9:
const m_ENOTBLK = 15

// /usr/include/sys/errno.h:108:9:
const m_ENOTCONN = 57

// /usr/include/sys/errno.h:62:9:
const m_ENOTDIR = 20

// /usr/include/sys/errno.h:120:9:
const m_ENOTEMPTY = 66

// /usr/include/sys/errno.h:177:9:
const m_ENOTRECOVERABLE = 98

// /usr/include/sys/errno.h:87:9:
const m_ENOTSOCK = 38

// /usr/include/sys/errno.h:153:9:
const m_ENOTSUP = 86

// /usr/include/sys/errno.h:67:9:
const m_ENOTTY = 25

// /usr/include/sys/errno.h:47:9:
const m_ENXIO = 6

// /usr/include/sys/errno.h:94:9:
const m_EOPNOTSUPP = 45

// /usr/include/sys/errno.h:146:9:
const m_EOVERFLOW = 84

// /usr/include/sys/errno.h:176:9:
const m_EOWNERDEAD = 97

// /usr/include/sys/errno.h:42:9:
const m_EPERM = 1

// /usr/include/sys/errno.h:95:9:
const m_EPFNOSUPPORT = 46

// /usr/include/sys/errno.h:74:9:
const m_EPIPE = 32

// /usr/include/sys/errno.h:123:9:
const m_EPROCLIM = 67

// /usr/include/sys/errno.h:134:9:
const m_EPROCUNAVAIL = 76

// /usr/include/sys/errno.h:133:9:
const m_EPROGMISMATCH = 75

// /usr/include/sys/errno.h:132:9:
const m_EPROGUNAVAIL = 74

// /usr/include/sys/errno.h:173:9:
const m_EPROTO = 96

// /usr/include/sys/errno.h:92:9:
const m_EPROTONOSUPPORT = 43

// /usr/include/sys/errno.h:90:9:
const m_EPROTOTYPE = 41

// /usr/include/sys/errno.h:78:9:
const m_ERANGE = 34

// /usr/include/sys/errno.h:129:9:
const m_EREMOTE = 71

// /usr/include/sys/errno.h:72:9:
const m_EROFS = 30

// /usr/include/sys/errno.h:131:9:
const m_ERPCMISMATCH = 73

// /usr/include/sys/errno.h:109:9:
const m_ESHUTDOWN = 58

// /usr/include/sys/errno.h:93:9:
const m_ESOCKTNOSUPPORT = 44

// /usr/include/sys/errno.h:71:9:
const m_ESPIPE = 29

// /usr/include/sys/errno.h:44:9:
const m_ESRCH = 3

// /usr/include/sys/errno.h:128:9:
const m_ESTALE = 70

// /usr/include/sys/errno.h:165:9:
const m_ETIME = 92

// /usr/include/sys/errno.h:111:9:
const m_ETIMEDOUT = 60

// /usr/include/sys/errno.h:110:9:
const m_ETOOMANYREFS = 59

// /usr/include/sys/errno.h:68:9:
const m_ETXTBSY = 26

// /usr/include/sys/errno.h:124:9:
const m_EUSERS = 68

// /usr/include/sys/errno.h:82:9:
const m_EWOULDBLOCK = "EAGAIN"

// /usr/include/sys/errno.h:60:9:
const m_EXDEV = 18

// /usr/include/stdlib.h:86:9:
const m_EXIT_FAILURE = 1

// /usr/include/stdlib.h:87:9:
const m_EXIT_SUCCESS = 0

// /usr/include/sys/syslimits.h:68:9:
const m_EXPR_NEST_MAX = 32

// /usr/include/sys/fd_set.h:59:9:
const m_FD_SETSIZE = 256

// /usr/include/stdio.h:193:9:
const m_FILENAME_MAX = 1024

// /usr/include/machine/limits.h:89:9:
const m_FLT_DIG = "__FLT_DIG__"

// /usr/include/sys/float_ieee754.h:77:9:
const m_FLT_EPSILON = "__FLT_EPSILON__"

// /usr/include/sys/float_ieee754.h:65:9:
const m_FLT_EVAL_METHOD = "__FLT_EVAL_METHOD__"

// /usr/include/sys/float_ieee754.h:76:9:
const m_FLT_MANT_DIG = "__FLT_MANT_DIG__"

// /usr/include/machine/limits.h:90:9:
const m_FLT_MAX = "__FLT_MAX__"

// /usr/include/sys/float_ieee754.h:84:9:
const m_FLT_MAX_10_EXP = "__FLT_MAX_10_EXP__"

// /usr/include/sys/float_ieee754.h:82:9:
const m_FLT_MAX_EXP = "__FLT_MAX_EXP__"

// /usr/include/machine/limits.h:91:9:
const m_FLT_MIN = "__FLT_MIN__"

// /usr/include/sys/float_ieee754.h:81:9:
const m_FLT_MIN_10_EXP = "__FLT_MIN_10_EXP__"

// /usr/include/sys/float_ieee754.h:79:9:
const m_FLT_MIN_EXP = "__FLT_MIN_EXP__"

// /usr/include/sys/float_ieee754.h:74:9:
const m_FLT_RADIX = "__FLT_RADIX__"

// /usr/include/stdio.h:192:9:
const m_FOPEN_MAX = 20

// /usr/include/stdio.h:409:9:
const m_FPARSELN_UNESCALL = 0x0f

// /usr/include/stdio.h:407:9:
const m_FPARSELN_UNESCCOMM = 0x04

// /usr/include/stdio.h:406:9:
const m_FPARSELN_UNESCCONT = 0x02

// /usr/include/stdio.h:405:9:
const m_FPARSELN_UNESCESC = 0x01

// /usr/include/stdio.h:408:9:
const m_FPARSELN_UNESCREST = 0x08

// /usr/include/math.h:150:9:
const m_FP_ILOGB0 = "INT_MIN"

// /usr/include/math.h:151:9:
const m_FP_ILOGBNAN = "INT_MAX"

// /usr/include/math.h:127:9:
const m_FP_INFINITE = 0x00

// /usr/include/math.h:128:9:
const m_FP_NAN = 0x01

// /usr/include/math.h:129:9:
const m_FP_NORMAL = 0x02

// /usr/include/math.h:130:9:
const m_FP_SUBNORMAL = 0x03

// /usr/include/math.h:131:9:
const m_FP_ZERO = 0x04

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14273:9:
const m_FTS5_TOKENIZE_AUX = 0x0008

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14272:9:
const m_FTS5_TOKENIZE_DOCUMENT = 0x0004

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14271:9:
const m_FTS5_TOKENIZE_PREFIX = 0x0002

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14270:9:
const m_FTS5_TOKENIZE_QUERY = 0x0001

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14277:9:
const m_FTS5_TOKEN_COLOCATED = 0x0001

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11478:9:
const m_FULLY_WITHIN = 2

// /usr/include/limits.h:136:9:
const m_GETENTROPY_MAX = 256

// /usr/include/sys/syslimits.h:49:9:
const m_GID_MAX = 2147483647

// /usr/include/stdlib.h:306:9:
const m_HN_AUTOSCALE = 0x20

// /usr/include/stdlib.h:302:9:
const m_HN_B = 0x04

// /usr/include/stdlib.h:300:9:
const m_HN_DECIMAL = 0x01

// /usr/include/stdlib.h:303:9:
const m_HN_DIVISOR_1000 = 0x08

// /usr/include/stdlib.h:305:9:
const m_HN_GETSCALE = 0x10

// /usr/include/stdlib.h:301:9:
const m_HN_NOSPACE = 0x02

// /usr/include/math.h:216:9:
const m_HUGE = "MAXFLOAT"

// /usr/include/sys/common_int_limits.h:53:9:
const m_INT16_MAX = "__INT16_MAX__"

// /usr/include/sys/common_int_limits.h:54:9:
const m_INT32_MAX = "__INT32_MAX__"

// /usr/include/sys/common_int_limits.h:55:9:
const m_INT64_MAX = "__INT64_MAX__"

// /usr/include/sys/common_int_limits.h:52:9:
const m_INT8_MAX = "__INT8_MAX__"

// /usr/include/sys/common_int_limits.h:111:9:
const m_INTMAX_MAX = "__INTMAX_MAX__"

// /usr/include/sys/common_int_limits.h:105:9:
const m_INTPTR_MAX = "__INTPTR_MAX__"

// /usr/include/sys/common_int_limits.h:93:9:
const m_INT_FAST16_MAX = "__INT_FAST16_MAX__"

// /usr/include/sys/common_int_limits.h:94:9:
const m_INT_FAST32_MAX = "__INT_FAST32_MAX__"

// /usr/include/sys/common_int_limits.h:95:9:
const m_INT_FAST64_MAX = "__INT_FAST64_MAX__"

// /usr/include/sys/common_int_limits.h:92:9:
const m_INT_FAST8_MAX = "__INT_FAST8_MAX__"

// /usr/include/sys/common_int_limits.h:73:9:
const m_INT_LEAST16_MAX = "__INT_LEAST16_MAX__"

// /usr/include/sys/common_int_limits.h:74:9:
const m_INT_LEAST32_MAX = "__INT_LEAST32_MAX__"

// /usr/include/sys/common_int_limits.h:75:9:
const m_INT_LEAST64_MAX = "__INT_LEAST64_MAX__"

// /usr/include/sys/common_int_limits.h:72:9:
const m_INT_LEAST8_MAX = "__INT_LEAST8_MAX__"

// /usr/include/machine/limits.h:52:9:
const m_INT_MAX = 0x7fffffff

// /usr/include/sys/syslimits.h:84:9:
const m_IOV_MAX = 1024

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:601:9:
const m_JSON_SUBTYPE = 74

// /usr/include/x86/float.h:10:9:
const m_LDBL_DIG = 18

// /usr/include/x86/float.h:9:9:
const m_LDBL_EPSILON = 1.0842021724855044340e-19

// /usr/include/x86/float.h:8:9:
const m_LDBL_MANT_DIG = 64

// /usr/include/x86/float.h:15:9:
const m_LDBL_MAX = "1.1897314953572317650E+4932"

// /usr/include/x86/float.h:16:9:
const m_LDBL_MAX_10_EXP = 4932

// /usr/include/x86/float.h:14:9:
const m_LDBL_MAX_EXP = 16384

// /usr/include/x86/float.h:12:9:
const m_LDBL_MIN = 3.3621031431120935063e-4932

// /usr/include/sys/syslimits.h:69:9:
const m_LINE_MAX = 2048

// /usr/include/sys/syslimits.h:50:9:
const m_LINK_MAX = 32767

// /usr/include/sys/endian.h:102:9:
const m_LITTLE_ENDIAN = 1234

// /usr/include/machine/limits.h:62:9:
const m_LLONG_MAX = 0x7fffffffffffffff

// /usr/include/sys/syslimits.h:77:9:
const m_LOGIN_NAME_MAX = 17

// /usr/include/machine/limits.h:82:9:
const m_LONG_BIT = 64

// /usr/include/machine/limits.h:56:9:
const m_LONG_MAX = 9223372036854775807

// /usr/include/stdio.h:293:9:
const m_L_ctermid = 1024

// /usr/include/stdio.h:294:9:
const m_L_cuserid = 9

// /usr/include/stdio.h:199:9:
const m_L_tmpnam = 1024

// /usr/include/math.h:155:9:
const m_MATH_ERREXCEPT = 2

// /usr/include/math.h:154:9:
const m_MATH_ERRNO = 1

// /usr/include/sys/syslimits.h:51:9:
const m_MAX_CANON = 255

// /usr/include/sys/syslimits.h:52:9:
const m_MAX_INPUT = 255

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9424:9:
const m_MAX_STATIC_BLOBS = 16

// /usr/include/stdlib.h:92:9:
const m_MB_CUR_MAX = "__mb_cur_max"

// /usr/include/limits.h:141:9:
const m_MB_LEN_MAX = 32

// /usr/include/math.h:176:9:
const m_M_1_PI = 0.31830988618379067154

// /usr/include/math.h:177:9:
const m_M_2_PI = 0.63661977236758134308

// /usr/include/math.h:178:9:
const m_M_2_SQRTPI = 1.12837916709551257390

// /usr/include/math.h:168:9:
const m_M_E = 2.7182818284590452354

// /usr/include/math.h:172:9:
const m_M_LN10 = 2.30258509299404568402

// /usr/include/math.h:171:9:
const m_M_LN2 = 0.69314718055994530942

// /usr/include/math.h:170:9:
const m_M_LOG10E = 0.43429448190325182765

// /usr/include/math.h:169:9:
const m_M_LOG2E = 1.4426950408889634074

// /usr/include/math.h:173:9:
const m_M_PI = 3.14159265358979323846

// /usr/include/math.h:174:9:
const m_M_PI_2 = 1.57079632679489661923

// /usr/include/math.h:175:9:
const m_M_PI_4 = 0.78539816339744830962

// /usr/include/math.h:180:9:
const m_M_SQRT1_2 = 0.70710678118654752440

// /usr/include/math.h:179:9:
const m_M_SQRT2 = 1.41421356237309504880

// /usr/include/sys/syslimits.h:53:9:
const m_NAME_MAX = 511

// /usr/include/sys/types.h:331:9:
const m_NBBY = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<command-line>:1:9:
const m_NDEBUG = 1

// /usr/include/sys/fd_set.h:93:9:
const m_NFDBITS = "__NFDBITS"

// /usr/include/sys/syslimits.h:55:9:
const m_NGROUPS_MAX = 16

// /usr/include/limits.h:114:9:
const m_NL_ARGMAX = 9

// /usr/include/limits.h:115:9:
const m_NL_LANGMAX = 14

// /usr/include/limits.h:116:9:
const m_NL_MSGMAX = 32767

// /usr/include/limits.h:117:9:
const m_NL_NMAX = 1

// /usr/include/limits.h:118:9:
const m_NL_SETMAX = 255

// /usr/include/limits.h:119:9:
const m_NL_TEXTMAX = 2048

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11476:9:
const m_NOT_WITHIN = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2782:9:
const m_NPY_PARSE_ERROR = "Error parsing numpy array: numpy header did not start with '{'"

// /usr/include/sys/syslimits.h:85:9:
const m_NZERO = 20

// /usr/include/sys/syslimits.h:58:9:
const m_OPEN_MAX = 128

// /usr/include/math.h:227:9:
const m_OVERFLOW = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11477:9:
const m_PARTLY_WITHIN = 1

// /usr/include/limits.h:111:9:
const m_PASS_MAX = 128

// /usr/include/sys/syslimits.h:60:9:
const m_PATH_MAX = 1024

// /usr/include/sys/endian.h:104:9:
const m_PDP_ENDIAN = 3412

// /usr/include/sys/syslimits.h:61:9:
const m_PIPE_BUF = 512

// /usr/include/math.h:230:9:
const m_PLOSS = 6

// /usr/include/machine/int_fmtio.h:125:9:
const m_PRIX16 = "X"

// /usr/include/machine/int_fmtio.h:126:9:
const m_PRIX32 = "X"

// /usr/include/machine/int_fmtio.h:127:9:
const m_PRIX64 = "lX"

// /usr/include/machine/int_fmtio.h:124:9:
const m_PRIX8 = "X"

// /usr/include/machine/int_fmtio.h:133:9:
const m_PRIXFAST16 = "X"

// /usr/include/machine/int_fmtio.h:134:9:
const m_PRIXFAST32 = "X"

// /usr/include/machine/int_fmtio.h:135:9:
const m_PRIXFAST64 = "lX"

// /usr/include/machine/int_fmtio.h:132:9:
const m_PRIXFAST8 = "X"

// /usr/include/machine/int_fmtio.h:129:9:
const m_PRIXLEAST16 = "X"

// /usr/include/machine/int_fmtio.h:130:9:
const m_PRIXLEAST32 = "X"

// /usr/include/machine/int_fmtio.h:131:9:
const m_PRIXLEAST64 = "lX"

// /usr/include/machine/int_fmtio.h:128:9:
const m_PRIXLEAST8 = "X"

// /usr/include/machine/int_fmtio.h:136:9:
const m_PRIXMAX = "lX"

// /usr/include/machine/int_fmtio.h:137:9:
const m_PRIXPTR = "lX"

// /usr/include/machine/int_fmtio.h:48:9:
const m_PRId16 = "d"

// /usr/include/machine/int_fmtio.h:49:9:
const m_PRId32 = "d"

// /usr/include/machine/int_fmtio.h:50:9:
const m_PRId64 = "ld"

// /usr/include/machine/int_fmtio.h:47:9:
const m_PRId8 = "d"

// /usr/include/machine/int_fmtio.h:56:9:
const m_PRIdFAST16 = "d"

// /usr/include/machine/int_fmtio.h:57:9:
const m_PRIdFAST32 = "d"

// /usr/include/machine/int_fmtio.h:58:9:
const m_PRIdFAST64 = "ld"

// /usr/include/machine/int_fmtio.h:55:9:
const m_PRIdFAST8 = "d"

// /usr/include/machine/int_fmtio.h:52:9:
const m_PRIdLEAST16 = "d"

// /usr/include/machine/int_fmtio.h:53:9:
const m_PRIdLEAST32 = "d"

// /usr/include/machine/int_fmtio.h:54:9:
const m_PRIdLEAST64 = "ld"

// /usr/include/machine/int_fmtio.h:51:9:
const m_PRIdLEAST8 = "d"

// /usr/include/machine/int_fmtio.h:59:9:
const m_PRIdMAX = "ld"

// /usr/include/machine/int_fmtio.h:60:9:
const m_PRIdPTR = "ld"

// /usr/include/machine/int_fmtio.h:63:9:
const m_PRIi16 = "i"

// /usr/include/machine/int_fmtio.h:64:9:
const m_PRIi32 = "i"

// /usr/include/machine/int_fmtio.h:65:9:
const m_PRIi64 = "li"

// /usr/include/machine/int_fmtio.h:62:9:
const m_PRIi8 = "i"

// /usr/include/machine/int_fmtio.h:71:9:
const m_PRIiFAST16 = "i"

// /usr/include/machine/int_fmtio.h:72:9:
const m_PRIiFAST32 = "i"

// /usr/include/machine/int_fmtio.h:73:9:
const m_PRIiFAST64 = "li"

// /usr/include/machine/int_fmtio.h:70:9:
const m_PRIiFAST8 = "i"

// /usr/include/machine/int_fmtio.h:67:9:
const m_PRIiLEAST16 = "i"

// /usr/include/machine/int_fmtio.h:68:9:
const m_PRIiLEAST32 = "i"

// /usr/include/machine/int_fmtio.h:69:9:
const m_PRIiLEAST64 = "li"

// /usr/include/machine/int_fmtio.h:66:9:
const m_PRIiLEAST8 = "i"

// /usr/include/machine/int_fmtio.h:74:9:
const m_PRIiMAX = "li"

// /usr/include/machine/int_fmtio.h:75:9:
const m_PRIiPTR = "li"

// /usr/include/machine/int_fmtio.h:80:9:
const m_PRIo16 = "o"

// /usr/include/machine/int_fmtio.h:81:9:
const m_PRIo32 = "o"

// /usr/include/machine/int_fmtio.h:82:9:
const m_PRIo64 = "lo"

// /usr/include/machine/int_fmtio.h:79:9:
const m_PRIo8 = "o"

// /usr/include/machine/int_fmtio.h:88:9:
const m_PRIoFAST16 = "o"

// /usr/include/machine/int_fmtio.h:89:9:
const m_PRIoFAST32 = "o"

// /usr/include/machine/int_fmtio.h:90:9:
const m_PRIoFAST64 = "lo"

// /usr/include/machine/int_fmtio.h:87:9:
const m_PRIoFAST8 = "o"

// /usr/include/machine/int_fmtio.h:84:9:
const m_PRIoLEAST16 = "o"

// /usr/include/machine/int_fmtio.h:85:9:
const m_PRIoLEAST32 = "o"

// /usr/include/machine/int_fmtio.h:86:9:
const m_PRIoLEAST64 = "lo"

// /usr/include/machine/int_fmtio.h:83:9:
const m_PRIoLEAST8 = "o"

// /usr/include/machine/int_fmtio.h:91:9:
const m_PRIoMAX = "lo"

// /usr/include/machine/int_fmtio.h:92:9:
const m_PRIoPTR = "lo"

// /usr/include/machine/int_fmtio.h:95:9:
const m_PRIu16 = "u"

// /usr/include/machine/int_fmtio.h:96:9:
const m_PRIu32 = "u"

// /usr/include/machine/int_fmtio.h:97:9:
const m_PRIu64 = "lu"

// /usr/include/machine/int_fmtio.h:94:9:
const m_PRIu8 = "u"

// /usr/include/machine/int_fmtio.h:103:9:
const m_PRIuFAST16 = "u"

// /usr/include/machine/int_fmtio.h:104:9:
const m_PRIuFAST32 = "u"

// /usr/include/machine/int_fmtio.h:105:9:
const m_PRIuFAST64 = "lu"

// /usr/include/machine/int_fmtio.h:102:9:
const m_PRIuFAST8 = "u"

// /usr/include/machine/int_fmtio.h:99:9:
const m_PRIuLEAST16 = "u"

// /usr/include/machine/int_fmtio.h:100:9:
const m_PRIuLEAST32 = "u"

// /usr/include/machine/int_fmtio.h:101:9:
const m_PRIuLEAST64 = "lu"

// /usr/include/machine/int_fmtio.h:98:9:
const m_PRIuLEAST8 = "u"

// /usr/include/machine/int_fmtio.h:106:9:
const m_PRIuMAX = "lu"

// /usr/include/machine/int_fmtio.h:107:9:
const m_PRIuPTR = "lu"

// /usr/include/machine/int_fmtio.h:110:9:
const m_PRIx16 = "x"

// /usr/include/machine/int_fmtio.h:111:9:
const m_PRIx32 = "x"

// /usr/include/machine/int_fmtio.h:112:9:
const m_PRIx64 = "lx"

// /usr/include/machine/int_fmtio.h:109:9:
const m_PRIx8 = "x"

// /usr/include/machine/int_fmtio.h:118:9:
const m_PRIxFAST16 = "x"

// /usr/include/machine/int_fmtio.h:119:9:
const m_PRIxFAST32 = "x"

// /usr/include/machine/int_fmtio.h:120:9:
const m_PRIxFAST64 = "lx"

// /usr/include/machine/int_fmtio.h:117:9:
const m_PRIxFAST8 = "x"

// /usr/include/machine/int_fmtio.h:114:9:
const m_PRIxLEAST16 = "x"

// /usr/include/machine/int_fmtio.h:115:9:
const m_PRIxLEAST32 = "x"

// /usr/include/machine/int_fmtio.h:116:9:
const m_PRIxLEAST64 = "lx"

// /usr/include/machine/int_fmtio.h:113:9:
const m_PRIxLEAST8 = "x"

// /usr/include/machine/int_fmtio.h:121:9:
const m_PRIxMAX = "lx"

// /usr/include/machine/int_fmtio.h:122:9:
const m_PRIxPTR = "lx"

// /usr/include/limits.h:78:9:
const m_PTHREAD_DESTRUCTOR_ITERATIONS = "_POSIX_THREAD_DESTRUCTOR_ITERATIONS"

// /usr/include/limits.h:79:9:
const m_PTHREAD_KEYS_MAX = 256

// /usr/include/limits.h:80:9:
const m_PTHREAD_STACK_MIN = 4096

// /usr/include/limits.h:81:9:
const m_PTHREAD_THREADS_MAX = "_POSIX_THREAD_THREADS_MAX"

// /usr/include/sys/common_int_limits.h:121:9:
const m_PTRDIFF_MAX = "__PTRDIFF_MAX__"

// /usr/include/stdio.h:197:9:
const m_P_tmpdir = "/tmp/"

// /usr/include/machine/limits.h:75:9:
const m_QUAD_MAX = 0x7fffffffffffffff

// /usr/include/stdlib.h:186:9:
const m_RANDOM_MAX = 0x7fffffff

// /usr/include/stdlib.h:89:9:
const m_RAND_MAX = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3479:9:
const m_REPORT_URL = "https://github.com/asg017/sqlite-vec/issues/new"

// /usr/include/sys/syslimits.h:70:9:
const m_RE_DUP_MAX = 255

// /usr/include/machine/limits.h:44:9:
const m_SCHAR_MAX = 0x7f

// /usr/include/machine/int_fmtio.h:142:9:
const m_SCNd16 = "hd"

// /usr/include/machine/int_fmtio.h:143:9:
const m_SCNd32 = "d"

// /usr/include/machine/int_fmtio.h:144:9:
const m_SCNd64 = "ld"

// /usr/include/machine/int_fmtio.h:141:9:
const m_SCNd8 = "hhd"

// /usr/include/machine/int_fmtio.h:150:9:
const m_SCNdFAST16 = "d"

// /usr/include/machine/int_fmtio.h:151:9:
const m_SCNdFAST32 = "d"

// /usr/include/machine/int_fmtio.h:152:9:
const m_SCNdFAST64 = "ld"

// /usr/include/machine/int_fmtio.h:149:9:
const m_SCNdFAST8 = "d"

// /usr/include/machine/int_fmtio.h:146:9:
const m_SCNdLEAST16 = "hd"

// /usr/include/machine/int_fmtio.h:147:9:
const m_SCNdLEAST32 = "d"

// /usr/include/machine/int_fmtio.h:148:9:
const m_SCNdLEAST64 = "ld"

// /usr/include/machine/int_fmtio.h:145:9:
const m_SCNdLEAST8 = "hhd"

// /usr/include/machine/int_fmtio.h:153:9:
const m_SCNdMAX = "ld"

// /usr/include/machine/int_fmtio.h:154:9:
const m_SCNdPTR = "ld"

// /usr/include/machine/int_fmtio.h:157:9:
const m_SCNi16 = "hi"

// /usr/include/machine/int_fmtio.h:158:9:
const m_SCNi32 = "i"

// /usr/include/machine/int_fmtio.h:159:9:
const m_SCNi64 = "li"

// /usr/include/machine/int_fmtio.h:156:9:
const m_SCNi8 = "hhi"

// /usr/include/machine/int_fmtio.h:165:9:
const m_SCNiFAST16 = "i"

// /usr/include/machine/int_fmtio.h:166:9:
const m_SCNiFAST32 = "i"

// /usr/include/machine/int_fmtio.h:167:9:
const m_SCNiFAST64 = "li"

// /usr/include/machine/int_fmtio.h:164:9:
const m_SCNiFAST8 = "i"

// /usr/include/machine/int_fmtio.h:161:9:
const m_SCNiLEAST16 = "hi"

// /usr/include/machine/int_fmtio.h:162:9:
const m_SCNiLEAST32 = "i"

// /usr/include/machine/int_fmtio.h:163:9:
const m_SCNiLEAST64 = "li"

// /usr/include/machine/int_fmtio.h:160:9:
const m_SCNiLEAST8 = "hhi"

// /usr/include/machine/int_fmtio.h:168:9:
const m_SCNiMAX = "li"

// /usr/include/machine/int_fmtio.h:169:9:
const m_SCNiPTR = "li"

// /usr/include/machine/int_fmtio.h:174:9:
const m_SCNo16 = "ho"

// /usr/include/machine/int_fmtio.h:175:9:
const m_SCNo32 = "o"

// /usr/include/machine/int_fmtio.h:176:9:
const m_SCNo64 = "lo"

// /usr/include/machine/int_fmtio.h:173:9:
const m_SCNo8 = "hho"

// /usr/include/machine/int_fmtio.h:182:9:
const m_SCNoFAST16 = "o"

// /usr/include/machine/int_fmtio.h:183:9:
const m_SCNoFAST32 = "o"

// /usr/include/machine/int_fmtio.h:184:9:
const m_SCNoFAST64 = "lo"

// /usr/include/machine/int_fmtio.h:181:9:
const m_SCNoFAST8 = "o"

// /usr/include/machine/int_fmtio.h:178:9:
const m_SCNoLEAST16 = "ho"

// /usr/include/machine/int_fmtio.h:179:9:
const m_SCNoLEAST32 = "o"

// /usr/include/machine/int_fmtio.h:180:9:
const m_SCNoLEAST64 = "lo"

// /usr/include/machine/int_fmtio.h:177:9:
const m_SCNoLEAST8 = "hho"

// /usr/include/machine/int_fmtio.h:185:9:
const m_SCNoMAX = "lo"

// /usr/include/machine/int_fmtio.h:186:9:
const m_SCNoPTR = "lo"

// /usr/include/machine/int_fmtio.h:189:9:
const m_SCNu16 = "hu"

// /usr/include/machine/int_fmtio.h:190:9:
const m_SCNu32 = "u"

// /usr/include/machine/int_fmtio.h:191:9:
const m_SCNu64 = "lu"

// /usr/include/machine/int_fmtio.h:188:9:
const m_SCNu8 = "hhu"

// /usr/include/machine/int_fmtio.h:197:9:
const m_SCNuFAST16 = "u"

// /usr/include/machine/int_fmtio.h:198:9:
const m_SCNuFAST32 = "u"

// /usr/include/machine/int_fmtio.h:199:9:
const m_SCNuFAST64 = "lu"

// /usr/include/machine/int_fmtio.h:196:9:
const m_SCNuFAST8 = "u"

// /usr/include/machine/int_fmtio.h:193:9:
const m_SCNuLEAST16 = "hu"

// /usr/include/machine/int_fmtio.h:194:9:
const m_SCNuLEAST32 = "u"

// /usr/include/machine/int_fmtio.h:195:9:
const m_SCNuLEAST64 = "lu"

// /usr/include/machine/int_fmtio.h:192:9:
const m_SCNuLEAST8 = "hhu"

// /usr/include/machine/int_fmtio.h:200:9:
const m_SCNuMAX = "lu"

// /usr/include/machine/int_fmtio.h:201:9:
const m_SCNuPTR = "lu"

// /usr/include/machine/int_fmtio.h:204:9:
const m_SCNx16 = "hx"

// /usr/include/machine/int_fmtio.h:205:9:
const m_SCNx32 = "x"

// /usr/include/machine/int_fmtio.h:206:9:
const m_SCNx64 = "lx"

// /usr/include/machine/int_fmtio.h:203:9:
const m_SCNx8 = "hhx"

// /usr/include/machine/int_fmtio.h:212:9:
const m_SCNxFAST16 = "x"

// /usr/include/machine/int_fmtio.h:213:9:
const m_SCNxFAST32 = "x"

// /usr/include/machine/int_fmtio.h:214:9:
const m_SCNxFAST64 = "lx"

// /usr/include/machine/int_fmtio.h:211:9:
const m_SCNxFAST8 = "x"

// /usr/include/machine/int_fmtio.h:208:9:
const m_SCNxLEAST16 = "hx"

// /usr/include/machine/int_fmtio.h:209:9:
const m_SCNxLEAST32 = "x"

// /usr/include/machine/int_fmtio.h:210:9:
const m_SCNxLEAST64 = "lx"

// /usr/include/machine/int_fmtio.h:207:9:
const m_SCNxLEAST8 = "hhx"

// /usr/include/machine/int_fmtio.h:215:9:
const m_SCNxMAX = "lx"

// /usr/include/machine/int_fmtio.h:216:9:
const m_SCNxPTR = "lx"

// /usr/include/stdio.h:210:9:
const m_SEEK_CUR = 1

// /usr/include/stdio.h:213:9:
const m_SEEK_END = 2

// /usr/include/stdio.h:207:9:
const m_SEEK_SET = 0

// /usr/include/machine/limits.h:48:9:
const m_SHRT_MAX = 0x7fff

// /usr/include/sys/common_int_limits.h:125:9:
const m_SIG_ATOMIC_MAX = "__SIG_ATOMIC_MAX__"

// /usr/include/math.h:226:9:
const m_SING = 2

// /usr/include/sys/common_int_limits.h:128:9:
const m_SIZE_MAX = "__SIZE_MAX__"

// /usr/include/machine/limits.h:72:9:
const m_SIZE_T_MAX = "ULONG_MAX"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5363:9:
const m_SQLITE3_TEXT = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:454:9:
const m_SQLITE_ABORT = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1572:9:
const m_SQLITE_ACCESS_EXISTS = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1574:9:
const m_SQLITE_ACCESS_READ = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1573:9:
const m_SQLITE_ACCESS_READWRITE = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3529:9:
const m_SQLITE_ALTER_TABLE = 26

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3531:9:
const m_SQLITE_ANALYZE = 28

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5880:9:
const m_SQLITE_ANY = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3527:9:
const m_SQLITE_ATTACH = 24

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:473:9:
const m_SQLITE_AUTH = 23

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5356:9:
const m_SQLITE_BLOB = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:455:9:
const m_SQLITE_BUSY = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:464:9:
const m_SQLITE_CANTOPEN = 14

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11333:9:
const m_SQLITE_CARRAY_BLOB = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11331:9:
const m_SQLITE_CARRAY_DOUBLE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11329:9:
const m_SQLITE_CARRAY_INT32 = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11330:9:
const m_SQLITE_CARRAY_INT64 = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11332:9:
const m_SQLITE_CARRAY_TEXT = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10178:9:
const m_SQLITE_CHECKPOINT_FULL = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10177:9:
const m_SQLITE_CHECKPOINT_PASSIVE = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10179:9:
const m_SQLITE_CHECKPOINT_RESTART = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10180:9:
const m_SQLITE_CHECKPOINT_TRUNCATE = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2233:9:
const m_SQLITE_CONFIG_COVERING_INDEX_SCAN = 20

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2218:9:
const m_SQLITE_CONFIG_GETMALLOC = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2224:9:
const m_SQLITE_CONFIG_GETMUTEX = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2228:9:
const m_SQLITE_CONFIG_GETPCACHE = 15

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2232:9:
const m_SQLITE_CONFIG_GETPCACHE2 = 19

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2221:9:
const m_SQLITE_CONFIG_HEAP = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2229:9:
const m_SQLITE_CONFIG_LOG = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2226:9:
const m_SQLITE_CONFIG_LOOKASIDE = 13

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2217:9:
const m_SQLITE_CONFIG_MALLOC = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2242:9:
const m_SQLITE_CONFIG_MEMDB_MAXSIZE = 29

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2222:9:
const m_SQLITE_CONFIG_MEMSTATUS = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2235:9:
const m_SQLITE_CONFIG_MMAP_SIZE = 22

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2215:9:
const m_SQLITE_CONFIG_MULTITHREAD = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2223:9:
const m_SQLITE_CONFIG_MUTEX = 10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2220:9:
const m_SQLITE_CONFIG_PAGECACHE = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2227:9:
const m_SQLITE_CONFIG_PCACHE = 14

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2231:9:
const m_SQLITE_CONFIG_PCACHE2 = 18

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2237:9:
const m_SQLITE_CONFIG_PCACHE_HDRSZ = 24

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2238:9:
const m_SQLITE_CONFIG_PMASZ = 25

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2243:9:
const m_SQLITE_CONFIG_ROWID_IN_VIEW = 30

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2219:9:
const m_SQLITE_CONFIG_SCRATCH = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2216:9:
const m_SQLITE_CONFIG_SERIALIZED = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2214:9:
const m_SQLITE_CONFIG_SINGLETHREAD = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2240:9:
const m_SQLITE_CONFIG_SMALL_MALLOC = 27

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2241:9:
const m_SQLITE_CONFIG_SORTERREF_SIZE = 28

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2234:9:
const m_SQLITE_CONFIG_SQLLOG = 21

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2239:9:
const m_SQLITE_CONFIG_STMTJRNL_SPILL = 26

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2230:9:
const m_SQLITE_CONFIG_URI = 17

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2236:9:
const m_SQLITE_CONFIG_WIN32_HEAPSIZE = 23

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:469:9:
const m_SQLITE_CONSTRAINT = 19

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3536:9:
const m_SQLITE_COPY = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<command-line>:3:9:
const m_SQLITE_CORE = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:461:9:
const m_SQLITE_CORRUPT = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3504:9:
const m_SQLITE_CREATE_INDEX = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3505:9:
const m_SQLITE_CREATE_TABLE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3506:9:
const m_SQLITE_CREATE_TEMP_INDEX = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3507:9:
const m_SQLITE_CREATE_TEMP_TABLE = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3508:9:
const m_SQLITE_CREATE_TEMP_TRIGGER = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3509:9:
const m_SQLITE_CREATE_TEMP_VIEW = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3510:9:
const m_SQLITE_CREATE_TRIGGER = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3511:9:
const m_SQLITE_CREATE_VIEW = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3532:9:
const m_SQLITE_CREATE_VTABLE = 29

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2709:9:
const m_SQLITE_DBCONFIG_DEFENSIVE = 1010

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2713:9:
const m_SQLITE_DBCONFIG_DQS_DDL = 1014

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2712:9:
const m_SQLITE_DBCONFIG_DQS_DML = 1013

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2719:9:
const m_SQLITE_DBCONFIG_ENABLE_ATTACH_CREATE = 1020

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2720:9:
const m_SQLITE_DBCONFIG_ENABLE_ATTACH_WRITE = 1021

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2721:9:
const m_SQLITE_DBCONFIG_ENABLE_COMMENTS = 1022

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2701:9:
const m_SQLITE_DBCONFIG_ENABLE_FKEY = 1002

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2703:9:
const m_SQLITE_DBCONFIG_ENABLE_FTS3_TOKENIZER = 1004

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2704:9:
const m_SQLITE_DBCONFIG_ENABLE_LOAD_EXTENSION = 1005

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2706:9:
const m_SQLITE_DBCONFIG_ENABLE_QPSG = 1007

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2702:9:
const m_SQLITE_DBCONFIG_ENABLE_TRIGGER = 1003

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2714:9:
const m_SQLITE_DBCONFIG_ENABLE_VIEW = 1015

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2722:9:
const m_SQLITE_DBCONFIG_FP_DIGITS = 1023

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2711:9:
const m_SQLITE_DBCONFIG_LEGACY_ALTER_TABLE = 1012

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2715:9:
const m_SQLITE_DBCONFIG_LEGACY_FILE_FORMAT = 1016

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2700:9:
const m_SQLITE_DBCONFIG_LOOKASIDE = 1001

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2699:9:
const m_SQLITE_DBCONFIG_MAINDBNAME = 1000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2723:9:
const m_SQLITE_DBCONFIG_MAX = 1023

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2705:9:
const m_SQLITE_DBCONFIG_NO_CKPT_ON_CLOSE = 1006

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2708:9:
const m_SQLITE_DBCONFIG_RESET_DATABASE = 1009

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2718:9:
const m_SQLITE_DBCONFIG_REVERSE_SCANORDER = 1019

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2717:9:
const m_SQLITE_DBCONFIG_STMT_SCANSTATUS = 1018

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2707:9:
const m_SQLITE_DBCONFIG_TRIGGER_EQP = 1008

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2716:9:
const m_SQLITE_DBCONFIG_TRUSTED_SCHEMA = 1017

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:2710:9:
const m_SQLITE_DBCONFIG_WRITABLE_SCHEMA = 1011

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9201:9:
const m_SQLITE_DBSTATUS_CACHE_HIT = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9202:9:
const m_SQLITE_DBSTATUS_CACHE_MISS = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9206:9:
const m_SQLITE_DBSTATUS_CACHE_SPILL = 12

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9195:9:
const m_SQLITE_DBSTATUS_CACHE_USED = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9205:9:
const m_SQLITE_DBSTATUS_CACHE_USED_SHARED = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9203:9:
const m_SQLITE_DBSTATUS_CACHE_WRITE = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9204:9:
const m_SQLITE_DBSTATUS_DEFERRED_FKS = 10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9198:9:
const m_SQLITE_DBSTATUS_LOOKASIDE_HIT = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9200:9:
const m_SQLITE_DBSTATUS_LOOKASIDE_MISS_FULL = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9199:9:
const m_SQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9194:9:
const m_SQLITE_DBSTATUS_LOOKASIDE_USED = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9208:9:
const m_SQLITE_DBSTATUS_MAX = 13

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9196:9:
const m_SQLITE_DBSTATUS_SCHEMA_USED = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9197:9:
const m_SQLITE_DBSTATUS_STMT_USED = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9207:9:
const m_SQLITE_DBSTATUS_TEMPBUF_SPILL = 13

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3512:9:
const m_SQLITE_DELETE = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3481:9:
const m_SQLITE_DENY = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11270:9:
const m_SQLITE_DESERIALIZE_FREEONCLOSE = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11272:9:
const m_SQLITE_DESERIALIZE_READONLY = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11271:9:
const m_SQLITE_DESERIALIZE_RESIZEABLE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3528:9:
const m_SQLITE_DETACH = 25

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5983:9:
const m_SQLITE_DETERMINISTIC = 2048

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5984:9:
const m_SQLITE_DIRECTONLY = 0x000080000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:480:9:
const m_SQLITE_DONE = 101

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3513:9:
const m_SQLITE_DROP_INDEX = 10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3514:9:
const m_SQLITE_DROP_TABLE = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3515:9:
const m_SQLITE_DROP_TEMP_INDEX = 12

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3516:9:
const m_SQLITE_DROP_TEMP_TABLE = 13

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3517:9:
const m_SQLITE_DROP_TEMP_TRIGGER = 14

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3518:9:
const m_SQLITE_DROP_TEMP_VIEW = 15

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3519:9:
const m_SQLITE_DROP_TRIGGER = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3520:9:
const m_SQLITE_DROP_VIEW = 17

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3533:9:
const m_SQLITE_DROP_VTABLE = 30

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:466:9:
const m_SQLITE_EMPTY = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:451:9:
const m_SQLITE_ERROR = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:72:10:
const m_SQLITE_EXTERN = "extern"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10623:9:
const m_SQLITE_FAIL = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1272:9:
const m_SQLITE_FCNTL_BEGIN_ATOMIC_WRITE = 31

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1285:9:
const m_SQLITE_FCNTL_BLOCK_ON_CONNECT = 44

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1257:9:
const m_SQLITE_FCNTL_BUSYHANDLER = 15

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1248:9:
const m_SQLITE_FCNTL_CHUNK_SIZE = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1278:9:
const m_SQLITE_FCNTL_CKPT_DONE = 37

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1280:9:
const m_SQLITE_FCNTL_CKPT_START = 39

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1282:9:
const m_SQLITE_FCNTL_CKSM_FILE = 41

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1273:9:
const m_SQLITE_FCNTL_COMMIT_ATOMIC_WRITE = 32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1263:9:
const m_SQLITE_FCNTL_COMMIT_PHASETWO = 22

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1276:9:
const m_SQLITE_FCNTL_DATA_VERSION = 35

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1281:9:
const m_SQLITE_FCNTL_EXTERNAL_READER = 40

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1286:9:
const m_SQLITE_FCNTL_FILESTAT = 45

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1249:9:
const m_SQLITE_FCNTL_FILE_POINTER = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1244:9:
const m_SQLITE_FCNTL_GET_LOCKPROXYFILE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1261:9:
const m_SQLITE_FCNTL_HAS_MOVED = 20

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1269:9:
const m_SQLITE_FCNTL_JOURNAL_POINTER = 28

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1246:9:
const m_SQLITE_FCNTL_LAST_ERRNO = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1243:9:
const m_SQLITE_FCNTL_LOCKSTATE = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1275:9:
const m_SQLITE_FCNTL_LOCK_TIMEOUT = 34

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1259:9:
const m_SQLITE_FCNTL_MMAP_SIZE = 18

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1284:9:
const m_SQLITE_FCNTL_NULL_IO = 43

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1253:9:
const m_SQLITE_FCNTL_OVERWRITE = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1271:9:
const m_SQLITE_FCNTL_PDB = 30

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1252:9:
const m_SQLITE_FCNTL_PERSIST_WAL = 10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1255:9:
const m_SQLITE_FCNTL_POWERSAFE_OVERWRITE = 13

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1256:9:
const m_SQLITE_FCNTL_PRAGMA = 14

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1267:9:
const m_SQLITE_FCNTL_RBU = 26

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1279:9:
const m_SQLITE_FCNTL_RESERVE_BYTES = 38

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1283:9:
const m_SQLITE_FCNTL_RESET_CACHE = 42

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1274:9:
const m_SQLITE_FCNTL_ROLLBACK_ATOMIC_WRITE = 33

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1245:9:
const m_SQLITE_FCNTL_SET_LOCKPROXYFILE = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1247:9:
const m_SQLITE_FCNTL_SIZE_HINT = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1277:9:
const m_SQLITE_FCNTL_SIZE_LIMIT = 36

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1262:9:
const m_SQLITE_FCNTL_SYNC = 21

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1250:9:
const m_SQLITE_FCNTL_SYNC_OMITTED = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1258:9:
const m_SQLITE_FCNTL_TEMPFILENAME = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1260:9:
const m_SQLITE_FCNTL_TRACE = 19

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1254:9:
const m_SQLITE_FCNTL_VFSNAME = 12

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1268:9:
const m_SQLITE_FCNTL_VFS_POINTER = 27

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1265:9:
const m_SQLITE_FCNTL_WAL_BLOCK = 24

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1251:9:
const m_SQLITE_FCNTL_WIN32_AV_RETRY = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1270:9:
const m_SQLITE_FCNTL_WIN32_GET_HANDLE = 29

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1264:9:
const m_SQLITE_FCNTL_WIN32_SET_HANDLE = 23

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1266:9:
const m_SQLITE_FCNTL_ZIPVFS = 25

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5355:9:
const m_SQLITE_FLOAT = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:474:9:
const m_SQLITE_FORMAT = 24

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:463:9:
const m_SQLITE_FULL = 13

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3534:9:
const m_SQLITE_FUNCTION = 31

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1289:9:
const m_SQLITE_GET_LOCKPROXYFILE = "SQLITE_FCNTL_GET_LOCKPROXYFILE"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3482:9:
const m_SQLITE_IGNORE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7911:9:
const m_SQLITE_INDEX_CONSTRAINT_EQ = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7927:9:
const m_SQLITE_INDEX_CONSTRAINT_FUNCTION = 150

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7915:9:
const m_SQLITE_INDEX_CONSTRAINT_GE = 32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7918:9:
const m_SQLITE_INDEX_CONSTRAINT_GLOB = 66

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7912:9:
const m_SQLITE_INDEX_CONSTRAINT_GT = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7924:9:
const m_SQLITE_INDEX_CONSTRAINT_IS = 72

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7921:9:
const m_SQLITE_INDEX_CONSTRAINT_ISNOT = 69

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7922:9:
const m_SQLITE_INDEX_CONSTRAINT_ISNOTNULL = 70

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7923:9:
const m_SQLITE_INDEX_CONSTRAINT_ISNULL = 71

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7913:9:
const m_SQLITE_INDEX_CONSTRAINT_LE = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7917:9:
const m_SQLITE_INDEX_CONSTRAINT_LIKE = 65

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7925:9:
const m_SQLITE_INDEX_CONSTRAINT_LIMIT = 73

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7914:9:
const m_SQLITE_INDEX_CONSTRAINT_LT = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7916:9:
const m_SQLITE_INDEX_CONSTRAINT_MATCH = 64

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7920:9:
const m_SQLITE_INDEX_CONSTRAINT_NE = 68

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7926:9:
const m_SQLITE_INDEX_CONSTRAINT_OFFSET = 74

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7919:9:
const m_SQLITE_INDEX_CONSTRAINT_REGEXP = 67

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7870:9:
const m_SQLITE_INDEX_SCAN_HEX = 0x00000002

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7869:9:
const m_SQLITE_INDEX_SCAN_UNIQUE = 0x00000001

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5986:9:
const m_SQLITE_INNOCUOUS = 2097152

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3521:9:
const m_SQLITE_INSERT = 18

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5354:9:
const m_SQLITE_INTEGER = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:452:9:
const m_SQLITE_INTERNAL = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:459:9:
const m_SQLITE_INTERRUPT = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:671:9:
const m_SQLITE_IOCAP_ATOMIC = 0x00000001

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:677:9:
const m_SQLITE_IOCAP_ATOMIC16K = 0x00000040

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:673:9:
const m_SQLITE_IOCAP_ATOMIC1K = 0x00000004

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:674:9:
const m_SQLITE_IOCAP_ATOMIC2K = 0x00000008

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:678:9:
const m_SQLITE_IOCAP_ATOMIC32K = 0x00000080

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:675:9:
const m_SQLITE_IOCAP_ATOMIC4K = 0x00000010

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:672:9:
const m_SQLITE_IOCAP_ATOMIC512 = 0x00000002

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:679:9:
const m_SQLITE_IOCAP_ATOMIC64K = 0x00000100

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:676:9:
const m_SQLITE_IOCAP_ATOMIC8K = 0x00000020

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:685:9:
const m_SQLITE_IOCAP_BATCH_ATOMIC = 0x00004000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:684:9:
const m_SQLITE_IOCAP_IMMUTABLE = 0x00002000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:683:9:
const m_SQLITE_IOCAP_POWERSAFE_OVERWRITE = 0x00001000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:680:9:
const m_SQLITE_IOCAP_SAFE_APPEND = 0x00000200

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:681:9:
const m_SQLITE_IOCAP_SEQUENTIAL = 0x00000400

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:686:9:
const m_SQLITE_IOCAP_SUBPAGE_READ = 0x00008000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:682:9:
const m_SQLITE_IOCAP_UNDELETABLE_WHEN_OPEN = 0x00000800

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:460:9:
const m_SQLITE_IOERR = 10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1291:9:
const m_SQLITE_LAST_ERRNO = "SQLITE_FCNTL_LAST_ERRNO"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4414:9:
const m_SQLITE_LIMIT_ATTACHED = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4409:9:
const m_SQLITE_LIMIT_COLUMN = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4411:9:
const m_SQLITE_LIMIT_COMPOUND_SELECT = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4410:9:
const m_SQLITE_LIMIT_EXPR_DEPTH = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4413:9:
const m_SQLITE_LIMIT_FUNCTION_ARG = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4407:9:
const m_SQLITE_LIMIT_LENGTH = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4415:9:
const m_SQLITE_LIMIT_LIKE_PATTERN_LENGTH = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4419:9:
const m_SQLITE_LIMIT_PARSER_DEPTH = 12

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4408:9:
const m_SQLITE_LIMIT_SQL_LENGTH = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4417:9:
const m_SQLITE_LIMIT_TRIGGER_DEPTH = 10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4416:9:
const m_SQLITE_LIMIT_VARIABLE_NUMBER = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4412:9:
const m_SQLITE_LIMIT_VDBE_OP = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4418:9:
const m_SQLITE_LIMIT_WORKER_THREADS = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:456:9:
const m_SQLITE_LOCKED = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:703:9:
const m_SQLITE_LOCK_EXCLUSIVE = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:699:9:
const m_SQLITE_LOCK_NONE = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:702:9:
const m_SQLITE_LOCK_PENDING = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:701:9:
const m_SQLITE_LOCK_RESERVED = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:700:9:
const m_SQLITE_LOCK_SHARED = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:470:9:
const m_SQLITE_MISMATCH = 20

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:471:9:
const m_SQLITE_MISUSE = 21

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8580:9:
const m_SQLITE_MUTEX_FAST = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8581:9:
const m_SQLITE_MUTEX_RECURSIVE = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8590:9:
const m_SQLITE_MUTEX_STATIC_APP1 = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8591:9:
const m_SQLITE_MUTEX_STATIC_APP2 = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8592:9:
const m_SQLITE_MUTEX_STATIC_APP3 = 10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8587:9:
const m_SQLITE_MUTEX_STATIC_LRU = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8588:9:
const m_SQLITE_MUTEX_STATIC_LRU2 = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8582:9:
const m_SQLITE_MUTEX_STATIC_MAIN = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8598:9:
const m_SQLITE_MUTEX_STATIC_MASTER = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8583:9:
const m_SQLITE_MUTEX_STATIC_MEM = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8584:9:
const m_SQLITE_MUTEX_STATIC_MEM2 = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8585:9:
const m_SQLITE_MUTEX_STATIC_OPEN = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8589:9:
const m_SQLITE_MUTEX_STATIC_PMEM = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8586:9:
const m_SQLITE_MUTEX_STATIC_PRNG = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8593:9:
const m_SQLITE_MUTEX_STATIC_VFS1 = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8594:9:
const m_SQLITE_MUTEX_STATIC_VFS2 = 12

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8595:9:
const m_SQLITE_MUTEX_STATIC_VFS3 = 13

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:472:9:
const m_SQLITE_NOLFS = 22

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:457:9:
const m_SQLITE_NOMEM = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:476:9:
const m_SQLITE_NOTADB = 26

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:462:9:
const m_SQLITE_NOTFOUND = 12

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:477:9:
const m_SQLITE_NOTICE = 27

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5357:9:
const m_SQLITE_NULL = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:449:9:
const m_SQLITE_OK = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:608:9:
const m_SQLITE_OPEN_AUTOPROXY = 0x00000020

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:605:9:
const m_SQLITE_OPEN_CREATE = 0x00000004

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:606:9:
const m_SQLITE_OPEN_DELETEONCLOSE = 0x00000008

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:607:9:
const m_SQLITE_OPEN_EXCLUSIVE = 0x00000010

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:624:9:
const m_SQLITE_OPEN_EXRESCODE = 0x02000000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:619:9:
const m_SQLITE_OPEN_FULLMUTEX = 0x00010000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:611:9:
const m_SQLITE_OPEN_MAIN_DB = 0x00000100

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:614:9:
const m_SQLITE_OPEN_MAIN_JOURNAL = 0x00000800

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:628:9:
const m_SQLITE_OPEN_MASTER_JOURNAL = 0x00004000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:610:9:
const m_SQLITE_OPEN_MEMORY = 0x00000080

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:623:9:
const m_SQLITE_OPEN_NOFOLLOW = 0x01000000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:618:9:
const m_SQLITE_OPEN_NOMUTEX = 0x00008000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:621:9:
const m_SQLITE_OPEN_PRIVATECACHE = 0x00040000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:603:9:
const m_SQLITE_OPEN_READONLY = 0x00000001

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:604:9:
const m_SQLITE_OPEN_READWRITE = 0x00000002

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:620:9:
const m_SQLITE_OPEN_SHAREDCACHE = 0x00020000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:616:9:
const m_SQLITE_OPEN_SUBJOURNAL = 0x00002000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:617:9:
const m_SQLITE_OPEN_SUPER_JOURNAL = 0x00004000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:612:9:
const m_SQLITE_OPEN_TEMP_DB = 0x00000200

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:615:9:
const m_SQLITE_OPEN_TEMP_JOURNAL = 0x00001000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:613:9:
const m_SQLITE_OPEN_TRANSIENT_DB = 0x00000400

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:609:9:
const m_SQLITE_OPEN_URI = 0x00000040

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:622:9:
const m_SQLITE_OPEN_WAL = 0x00080000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:453:9:
const m_SQLITE_PERM = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3522:9:
const m_SQLITE_PRAGMA = 19

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4485:9:
const m_SQLITE_PREPARE_DONT_LOG = 0x10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4486:9:
const m_SQLITE_PREPARE_FROM_DDL = 0x20

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4483:9:
const m_SQLITE_PREPARE_NORMALIZE = 0x02

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4484:9:
const m_SQLITE_PREPARE_NO_VTAB = 0x04

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:4482:9:
const m_SQLITE_PREPARE_PERSISTENT = 0x01

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:465:9:
const m_SQLITE_PROTOCOL = 15

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:475:9:
const m_SQLITE_RANGE = 25

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3523:9:
const m_SQLITE_READ = 20

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:458:9:
const m_SQLITE_READONLY = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3537:9:
const m_SQLITE_RECURSIVE = 33

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3530:9:
const m_SQLITE_REINDEX = 27

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10625:9:
const m_SQLITE_REPLACE = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5987:9:
const m_SQLITE_RESULT_SUBTYPE = 16777216

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10621:9:
const m_SQLITE_ROLLBACK = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:479:9:
const m_SQLITE_ROW = 100

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3535:9:
const m_SQLITE_SAVEPOINT = 32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10755:9:
const m_SQLITE_SCANSTAT_COMPLEX = 0x0001

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10692:9:
const m_SQLITE_SCANSTAT_EST = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10694:9:
const m_SQLITE_SCANSTAT_EXPLAIN = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10693:9:
const m_SQLITE_SCANSTAT_NAME = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10697:9:
const m_SQLITE_SCANSTAT_NCYCLE = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10690:9:
const m_SQLITE_SCANSTAT_NLOOP = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10691:9:
const m_SQLITE_SCANSTAT_NVISIT = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10696:9:
const m_SQLITE_SCANSTAT_PARENTID = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10695:9:
const m_SQLITE_SCANSTAT_SELECTID = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:467:9:
const m_SQLITE_SCHEMA = 17

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:152:9:
const m_SQLITE_SCM_BRANCH = "branch-3.53"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:154:9:
const m_SQLITE_SCM_DATETIME = "2026-06-03T19:12:13.350Z"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:153:9:
const m_SQLITE_SCM_TAGS = "release version-3.53.2"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3524:9:
const m_SQLITE_SELECT = 21

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5988:9:
const m_SQLITE_SELFORDER1 = 0x002000000

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11194:9:
const m_SQLITE_SERIALIZE_NOCOPY = 0x001

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3113:9:
const m_SQLITE_SETLK_BLOCK_ON_CONNECT = 0x01

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1290:9:
const m_SQLITE_SET_LOCKPROXYFILE = "SQLITE_FCNTL_SET_LOCKPROXYFILE"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1601:9:
const m_SQLITE_SHM_EXCLUSIVE = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1599:9:
const m_SQLITE_SHM_LOCK = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1611:9:
const m_SQLITE_SHM_NLOCK = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1600:9:
const m_SQLITE_SHM_SHARED = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1598:9:
const m_SQLITE_SHM_UNLOCK = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:151:9:
const m_SQLITE_SOURCE_ID = "2026-06-03 19:12:13 d6e03d8c777cfa2d35e3b60d8ec3e0187f3e9f99d8e2ee9cac695fd6fcdf1a24"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9029:9:
const m_SQLITE_STATUS_MALLOC_COUNT = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9025:9:
const m_SQLITE_STATUS_MALLOC_SIZE = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9020:9:
const m_SQLITE_STATUS_MEMORY_USED = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9022:9:
const m_SQLITE_STATUS_PAGECACHE_OVERFLOW = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9027:9:
const m_SQLITE_STATUS_PAGECACHE_SIZE = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9021:9:
const m_SQLITE_STATUS_PAGECACHE_USED = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9026:9:
const m_SQLITE_STATUS_PARSER_STACK = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9024:9:
const m_SQLITE_STATUS_SCRATCH_OVERFLOW = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9028:9:
const m_SQLITE_STATUS_SCRATCH_SIZE = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9023:9:
const m_SQLITE_STATUS_SCRATCH_USED = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:84:10:
const m_SQLITE_STDCALL = "SQLITE_APICALL"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9304:9:
const m_SQLITE_STMTSTATUS_AUTOINDEX = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9309:9:
const m_SQLITE_STMTSTATUS_FILTER_HIT = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9308:9:
const m_SQLITE_STMTSTATUS_FILTER_MISS = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9302:9:
const m_SQLITE_STMTSTATUS_FULLSCAN_STEP = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9310:9:
const m_SQLITE_STMTSTATUS_MEMUSED = 99

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9306:9:
const m_SQLITE_STMTSTATUS_REPREPARE = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9307:9:
const m_SQLITE_STMTSTATUS_RUN = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9303:9:
const m_SQLITE_STMTSTATUS_SORT = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9305:9:
const m_SQLITE_STMTSTATUS_VM_STEP = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5985:9:
const m_SQLITE_SUBTYPE = 1048576

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:733:9:
const m_SQLITE_SYNC_DATAONLY = 0x00010

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:732:9:
const m_SQLITE_SYNC_FULL = 0x00003

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:731:9:
const m_SQLITE_SYNC_NORMAL = 0x00002

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8696:9:
const m_SQLITE_TESTCTRL_ALWAYS = 13

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8695:9:
const m_SQLITE_TESTCTRL_ASSERT = 12

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8722:9:
const m_SQLITE_TESTCTRL_ATOF = 34

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8693:9:
const m_SQLITE_TESTCTRL_BENIGN_MALLOC_HOOKS = 10

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8691:9:
const m_SQLITE_TESTCTRL_BITVEC_TEST = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8709:9:
const m_SQLITE_TESTCTRL_BYTEORDER = 22

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8705:9:
const m_SQLITE_TESTCTRL_EXPLAIN_STMT = 19

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8716:9:
const m_SQLITE_TESTCTRL_EXTRA_SCHEMA_CHECKS = 29

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8692:9:
const m_SQLITE_TESTCTRL_FAULT_INSTALL = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8686:9:
const m_SQLITE_TESTCTRL_FIRST = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8690:9:
const m_SQLITE_TESTCTRL_FK_NO_ACTION = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8701:9:
const m_SQLITE_TESTCTRL_GETOPT = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8712:9:
const m_SQLITE_TESTCTRL_IMPOSTER = 25

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8703:9:
const m_SQLITE_TESTCTRL_INTERNAL_FUNCTIONS = 17

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8710:9:
const m_SQLITE_TESTCTRL_ISINIT = 23

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8700:9:
const m_SQLITE_TESTCTRL_ISKEYWORD = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8698:9:
const m_SQLITE_TESTCTRL_JSON_SELFCHECK = 14

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8723:9:
const m_SQLITE_TESTCTRL_LAST = 34

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8704:9:
const m_SQLITE_TESTCTRL_LOCALTIME_FAULT = 18

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8720:9:
const m_SQLITE_TESTCTRL_LOGEST = 33

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8707:9:
const m_SQLITE_TESTCTRL_NEVER_CORRUPT = 20

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8706:9:
const m_SQLITE_TESTCTRL_ONCE_RESET_THRESHOLD = 19

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8699:9:
const m_SQLITE_TESTCTRL_OPTIMIZATIONS = 15

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8713:9:
const m_SQLITE_TESTCTRL_PARSER_COVERAGE = 26

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8694:9:
const m_SQLITE_TESTCTRL_PENDING_BYTE = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8689:9:
const m_SQLITE_TESTCTRL_PRNG_RESET = 7

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8688:9:
const m_SQLITE_TESTCTRL_PRNG_RESTORE = 6

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8687:9:
const m_SQLITE_TESTCTRL_PRNG_SAVE = 5

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8715:9:
const m_SQLITE_TESTCTRL_PRNG_SEED = 28

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8697:9:
const m_SQLITE_TESTCTRL_RESERVE = 14

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8714:9:
const m_SQLITE_TESTCTRL_RESULT_INTREAL = 27

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8702:9:
const m_SQLITE_TESTCTRL_SCRATCHMALLOC = 17

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8717:9:
const m_SQLITE_TESTCTRL_SEEK_COUNT = 30

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8711:9:
const m_SQLITE_TESTCTRL_SORTER_MMAP = 24

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8718:9:
const m_SQLITE_TESTCTRL_TRACEFLAGS = 31

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8719:9:
const m_SQLITE_TESTCTRL_TUNE = 32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8721:9:
const m_SQLITE_TESTCTRL_USELONGDOUBLE = 34

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8708:9:
const m_SQLITE_TESTCTRL_VDBE_COVERAGE = 21

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5361:10:
const m_SQLITE_TEXT = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11360:11:
const m_SQLITE_THREADSAFE = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:468:9:
const m_SQLITE_TOOBIG = 18

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3631:9:
const m_SQLITE_TRACE_CLOSE = 0x08

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3629:9:
const m_SQLITE_TRACE_PROFILE = 0x02

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3630:9:
const m_SQLITE_TRACE_ROW = 0x04

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3628:9:
const m_SQLITE_TRACE_STMT = 0x01

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3525:9:
const m_SQLITE_TRANSACTION = 22

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7089:9:
const m_SQLITE_TXN_NONE = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7090:9:
const m_SQLITE_TXN_READ = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7091:9:
const m_SQLITE_TXN_WRITE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:3526:9:
const m_SQLITE_UPDATE = 23

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5879:9:
const m_SQLITE_UTF16 = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5878:9:
const m_SQLITE_UTF16BE = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5877:9:
const m_SQLITE_UTF16LE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5881:9:
const m_SQLITE_UTF16_ALIGNED = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5876:9:
const m_SQLITE_UTF8 = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:5882:9:
const m_SQLITE_UTF8_ZT = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4921:9:
const m_SQLITE_VEC_CHUNK_SIZE_MAX = 4096

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.h:22:9:
const m_SQLITE_VEC_DATE = ""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10047:9:
const m_SQLITE_VEC_DEBUG_BUILD_AVX = ""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10052:9:
const m_SQLITE_VEC_DEBUG_BUILD_NEON = ""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1072:9:
const m_SQLITE_VEC_NPY_FILE_NAME = "vec0-npy-file"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.h:23:9:
const m_SQLITE_VEC_SOURCE = ""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<command-line>:4:9:
const m_SQLITE_VEC_STATIC = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7111:9:
const m_SQLITE_VEC_VEC0_K_MAX = 4096

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3488:9:
const m_SQLITE_VEC_VEC0_MAX_DIMENSIONS = 8192

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.h:20:9:
const m_SQLITE_VEC_VERSION = "v0.1.9"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.h:26:9:
const m_SQLITE_VEC_VERSION_MAJOR = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.h:27:9:
const m_SQLITE_VEC_VERSION_MINOR = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.h:28:9:
const m_SQLITE_VEC_VERSION_PATCH = 9

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:149:9:
const m_SQLITE_VERSION = "3.53.2"

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:150:9:
const m_SQLITE_VERSION_NUMBER = 3053002

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10273:9:
const m_SQLITE_VTAB_CONSTRAINT_SUPPORT = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10275:9:
const m_SQLITE_VTAB_DIRECTONLY = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10274:9:
const m_SQLITE_VTAB_INNOCUOUS = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10276:9:
const m_SQLITE_VTAB_USES_ALL_SCHEMAS = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:478:9:
const m_SQLITE_WARNING = 28

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:6938:9:
const m_SQLITE_WIN32_DATA_DIRECTORY_TYPE = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:6939:9:
const m_SQLITE_WIN32_TEMP_DIRECTORY_TYPE = 2

// /usr/include/machine/limits.h:68:9:
const m_SSIZE_MAX = "LONG_MAX"

// /usr/include/machine/limits.h:71:9:
const m_SSIZE_MIN = "LONG_MIN"

// /usr/include/math.h:229:9:
const m_TLOSS = 5

// /usr/include/limits.h:127:9:
const m_TMP_MAX = 308915776

// /usr/include/machine/limits.h:43:9:
const m_UCHAR_MAX = 0xff

// /usr/include/sys/syslimits.h:56:9:
const m_UID_MAX = 2147483647

// /usr/include/sys/common_int_limits.h:59:9:
const m_UINT16_MAX = "__UINT16_MAX__"

// /usr/include/sys/common_int_limits.h:60:9:
const m_UINT32_MAX = "__UINT32_MAX__"

// /usr/include/sys/common_int_limits.h:61:9:
const m_UINT64_MAX = "__UINT64_MAX__"

// /usr/include/sys/common_int_limits.h:58:9:
const m_UINT8_MAX = "__UINT8_MAX__"

// /usr/include/sys/common_int_limits.h:112:9:
const m_UINTMAX_MAX = "__UINTMAX_MAX__"

// /usr/include/sys/common_int_limits.h:106:9:
const m_UINTPTR_MAX = "__UINTPTR_MAX__"

// /usr/include/sys/common_int_limits.h:99:9:
const m_UINT_FAST16_MAX = "__UINT_FAST16_MAX__"

// /usr/include/sys/common_int_limits.h:100:9:
const m_UINT_FAST32_MAX = "__UINT_FAST32_MAX__"

// /usr/include/sys/common_int_limits.h:101:9:
const m_UINT_FAST64_MAX = "__UINT_FAST64_MAX__"

// /usr/include/sys/common_int_limits.h:98:9:
const m_UINT_FAST8_MAX = "__UINT_FAST8_MAX__"

// /usr/include/sys/common_int_limits.h:79:9:
const m_UINT_LEAST16_MAX = "__UINT_LEAST16_MAX__"

// /usr/include/sys/common_int_limits.h:80:9:
const m_UINT_LEAST32_MAX = "__UINT_LEAST32_MAX__"

// /usr/include/sys/common_int_limits.h:81:9:
const m_UINT_LEAST64_MAX = "__UINT_LEAST64_MAX__"

// /usr/include/sys/common_int_limits.h:78:9:
const m_UINT_LEAST8_MAX = "__UINT_LEAST8_MAX__"

// /usr/include/machine/limits.h:51:9:
const m_UINT_MAX = 0xffffffff

// /usr/include/machine/limits.h:61:9:
const m_ULLONG_MAX = "0xffffffffffffffffU"

// /usr/include/machine/limits.h:55:9:
const m_ULONG_MAX = 0xffffffffffffffff

// /usr/include/math.h:228:9:
const m_UNDERFLOW = 4

// /usr/include/machine/limits.h:74:9:
const m_UQUAD_MAX = "0xffffffffffffffffU"

// /usr/include/machine/limits.h:47:9:
const m_USHRT_MAX = 0xffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3418:9:
const m_VEC0_COLUMN_ID = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3420:9:
const m_VEC0_COLUMN_OFFSET_DISTANCE = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3421:9:
const m_VEC0_COLUMN_OFFSET_K = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3419:9:
const m_VEC0_COLUMN_USERN_START = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3485:9:
const m_VEC0_MAX_AUXILIARY_COLUMNS = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3486:9:
const m_VEC0_MAX_METADATA_COLUMNS = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3484:9:
const m_VEC0_MAX_PARTITION_COLUMNS = 4

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3483:9:
const m_VEC0_MAX_VECTOR_COLUMNS = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3489:9:
const m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3490:9:
const m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH = 12

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3473:9:
const m_VEC0_SHADOW_AUXILIARY_NAME = "\\\"%w\\\".\\\"%w_auxiliary\\\""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3425:9:
const m_VEC0_SHADOW_CHUNKS_NAME = "\\\"%w\\\".\\\"%w_chunks\\\""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3423:9:
const m_VEC0_SHADOW_INFO_NAME = "\\\"%w\\\".\\\"%w_info\\\""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3475:9:
const m_VEC0_SHADOW_METADATA_N_NAME = "\\\"%w\\\".\\\"%w_metadatachunks%02d\\\""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3476:9:
const m_VEC0_SHADOW_METADATA_TEXT_DATA_NAME = "\\\"%w\\\".\\\"%w_metadatatext%02d\\\""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3435:9:
const m_VEC0_SHADOW_ROWIDS_NAME = "\\\"%w\\\".\\\"%w_rowids\\\""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3457:9:
const m_VEC0_SHADOW_VECTOR_N_NAME = "\\\"%w\\\".\\\"%w_vector_chunks%02d\\\""

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1844:9:
const m_VEC0_TOKEN_RESULT_EOF = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1846:9:
const m_VEC0_TOKEN_RESULT_ERROR = 3

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1845:9:
const m_VEC0_TOKEN_RESULT_SOME = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4720:9:
const m_VEC_CONSTRUCTOR_ERROR = "vec0 constructor error: could not parse vector column '%s'"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2497:9:
const m_VEC_EACH_COLUMN_VALUE = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2498:9:
const m_VEC_EACH_COLUMN_VECTOR = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3478:9:
const m_VEC_INTERAL_ERROR = "Internal sqlite-vec error: could not initialize 'rowids get chunk position' statement"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3133:9:
const m_VEC_NPY_EACH_COLUMN_INPUT = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3132:9:
const m_VEC_NPY_EACH_COLUMN_VECTOR = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9463:9:
const m_VEC_STATIC_BLOBS_COUNT = 3

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9461:9:
const m_VEC_STATIC_BLOBS_DATA = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9462:9:
const m_VEC_STATIC_BLOBS_DIMENSIONS = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9460:9:
const m_VEC_STATIC_BLOBS_NAME = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9705:9:
const m_VEC_STATIC_BLOB_ENTRIES_DISTANCE = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9706:9:
const m_VEC_STATIC_BLOB_ENTRIES_K = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9704:9:
const m_VEC_STATIC_BLOB_ENTRIES_VECTOR = 0

// /usr/include/machine/wchar_limits.h:41:9:
const m_WCHAR_MAX = 0x7fffffff

// /usr/include/machine/wchar_limits.h:45:9:
const m_WINT_MAX = 0x7fffffff

// /usr/include/machine/limits.h:83:9:
const m_WORD_BIT = 32

// /usr/include/math.h:223:9:
const m_X_TLOSS = 1.41484755040568800000e+16

// /usr/include/sys/endian.h:44:9:
const m__BIG_ENDIAN = 4321

// /usr/include/sys/ansi.h:68:9:
const m__BSD_MBSTATE_T_ = "__mbstate_t"

// /usr/include/sys/ansi.h:66:9:
const m__BSD_WCTRANS_T_ = "__wctrans_t"

// /usr/include/sys/ansi.h:67:9:
const m__BSD_WCTYPE_T_ = "__wctype_t"

// /usr/include/sys/common_ansi.h:74:9:
const m__BSD_WINT_T_ = "__WINT_TYPE__"

// /usr/include/machine/endian_machdep.h:3:9:
const m__BYTE_ORDER = "_LITTLE_ENDIAN"

// /usr/include/sys/float_ieee754.h:48:9:
const m__FLOAT_IEEE754 = 1

// /usr/include/math.h:134:9:
const m__FP_HIMD = 0xff

// /usr/include/math.h:133:9:
const m__FP_LOMD = 0x80

// /usr/include/limits.h:122:9:
const m__GETGR_R_SIZE_MAX = 1024

// /usr/include/limits.h:123:9:
const m__GETPW_R_SIZE_MAX = 1024

// /usr/include/math.h:201:9:
const m__IEEE_ = "fdlibm_ieee"

// /usr/include/stdio.h:179:9:
const m__IOFBF = 0

// /usr/include/stdio.h:180:9:
const m__IOLBF = 1

// /usr/include/stdio.h:181:9:
const m__IONBF = 2

// /usr/include/math.h:190:9:
const m__LIB_VERSION = "_fdlib_version"

// /usr/include/sys/endian.h:43:9:
const m__LITTLE_ENDIAN = 1234

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:261:9:
const m__LP64 = 1

// /usr/include/sys/featuretest.h:70:9:
const m__NETBSD_SOURCE = 1

// /usr/include/sys/endian.h:45:9:
const m__PDP_ENDIAN = 3412

// /usr/include/limits.h:91:9:
const m__POSIX2_BC_BASE_MAX = 99

// /usr/include/limits.h:92:9:
const m__POSIX2_BC_DIM_MAX = 2048

// /usr/include/limits.h:93:9:
const m__POSIX2_BC_SCALE_MAX = 99

// /usr/include/limits.h:94:9:
const m__POSIX2_BC_STRING_MAX = 1000

// /usr/include/limits.h:95:9:
const m__POSIX2_CHARCLASS_NAME_MAX = 14

// /usr/include/limits.h:96:9:
const m__POSIX2_COLL_WEIGHTS_MAX = 2

// /usr/include/limits.h:97:9:
const m__POSIX2_EXPR_NEST_MAX = 32

// /usr/include/limits.h:98:9:
const m__POSIX2_LINE_MAX = 2048

// /usr/include/limits.h:99:9:
const m__POSIX2_RE_DUP_MAX = 255

// /usr/include/math.h:204:9:
const m__POSIX_ = "fdlibm_posix"

// /usr/include/limits.h:41:9:
const m__POSIX_AIO_LISTIO_MAX = 2

// /usr/include/limits.h:42:9:
const m__POSIX_AIO_MAX = 1

// /usr/include/limits.h:43:9:
const m__POSIX_ARG_MAX = 4096

// /usr/include/limits.h:44:9:
const m__POSIX_CHILD_MAX = 25

// /usr/include/limits.h:87:9:
const m__POSIX_DELAYTIMER_MAX = 32

// /usr/include/limits.h:45:9:
const m__POSIX_HOST_NAME_MAX = 255

// /usr/include/limits.h:46:9:
const m__POSIX_LINK_MAX = 8

// /usr/include/limits.h:47:9:
const m__POSIX_LOGIN_NAME_MAX = 9

// /usr/include/limits.h:48:9:
const m__POSIX_MAX_CANON = 255

// /usr/include/limits.h:49:9:
const m__POSIX_MAX_INPUT = 255

// /usr/include/limits.h:50:9:
const m__POSIX_MQ_OPEN_MAX = 8

// /usr/include/limits.h:51:9:
const m__POSIX_MQ_PRIO_MAX = 32

// /usr/include/limits.h:52:9:
const m__POSIX_NAME_MAX = 14

// /usr/include/limits.h:53:9:
const m__POSIX_NGROUPS_MAX = 8

// /usr/include/limits.h:54:9:
const m__POSIX_OPEN_MAX = 20

// /usr/include/limits.h:55:9:
const m__POSIX_PATH_MAX = 256

// /usr/include/limits.h:56:9:
const m__POSIX_PIPE_BUF = 512

// /usr/include/limits.h:86:9:
const m__POSIX_REALTIME_SIGNALS = 200112

// /usr/include/limits.h:57:9:
const m__POSIX_RE_DUP_MAX = 255

// /usr/include/limits.h:84:9:
const m__POSIX_SEM_NSEMS_MAX = 256

// /usr/include/limits.h:85:9:
const m__POSIX_SIGQUEUE_MAX = 32

// /usr/include/limits.h:58:9:
const m__POSIX_SSIZE_MAX = 32767

// /usr/include/limits.h:59:9:
const m__POSIX_STREAM_MAX = 8

// /usr/include/limits.h:60:9:
const m__POSIX_SYMLINK_MAX = 255

// /usr/include/limits.h:61:9:
const m__POSIX_SYMLOOP_MAX = 8

// /usr/include/limits.h:68:9:
const m__POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4

// /usr/include/limits.h:69:9:
const m__POSIX_THREAD_KEYS_MAX = 128

// /usr/include/limits.h:70:9:
const m__POSIX_THREAD_THREADS_MAX = 64

// /usr/include/limits.h:83:9:
const m__POSIX_TIMER_MAX = 32

// /usr/include/limits.h:88:9:
const m__POSIX_TTY_NAME_MAX = 9

// /usr/include/limits.h:89:9:
const m__POSIX_TZNAME_MAX = 6

// /usr/include/pthread_types.h:280:9:
const m__PT_BARRIERATTR_DEAD = 0xDEAD0808

// /usr/include/pthread_types.h:279:9:
const m__PT_BARRIERATTR_MAGIC = 0x88880808

// /usr/include/pthread_types.h:272:9:
const m__PT_BARRIER_DEAD = 0xDEAD0008

// /usr/include/pthread_types.h:271:9:
const m__PT_BARRIER_MAGIC = 0x88880008

// /usr/include/pthread_types.h:199:9:
const m__PT_CONDATTR_DEAD = 0xDEAD0006

// /usr/include/pthread_types.h:198:9:
const m__PT_CONDATTR_MAGIC = 0x66660006

// /usr/include/pthread_types.h:183:9:
const m__PT_COND_DEAD = 0xDEAD0005

// /usr/include/pthread_types.h:182:9:
const m__PT_COND_MAGIC = 0x55550005

// /usr/include/pthread_types.h:167:9:
const m__PT_MUTEXATTR_DEAD = 0xDEAD0004

// /usr/include/pthread_types.h:166:9:
const m__PT_MUTEXATTR_MAGIC = 0x44440004

// /usr/include/pthread_types.h:139:9:
const m__PT_MUTEX_DEAD = 0xDEAD0003

// /usr/include/pthread_types.h:138:9:
const m__PT_MUTEX_MAGIC = 0x33330003

// /usr/include/pthread_types.h:255:9:
const m__PT_RWLOCKATTR_DEAD = 0xDEAD0909

// /usr/include/pthread_types.h:254:9:
const m__PT_RWLOCKATTR_MAGIC = 0x99990909

// /usr/include/pthread_types.h:238:9:
const m__PT_RWLOCK_DEAD = 0xDEAD0009

// /usr/include/pthread_types.h:237:9:
const m__PT_RWLOCK_MAGIC = 0x99990009

// /usr/include/pthread_types.h:215:9:
const m__PT_SPINLOCK_DEAD = 0xDEAD0007

// /usr/include/pthread_types.h:214:9:
const m__PT_SPINLOCK_MAGIC = 0x77770007

// /usr/include/pthread_types.h:216:9:
const m__PT_SPINLOCK_PSHARED = 0x00000001

// /usr/include/sys/endian.h:86:9:
const m__QUAD_HIGHWORD = 1

// /usr/include/sys/endian.h:87:9:
const m__QUAD_LOWWORD = 0

// /usr/include/math.h:202:9:
const m__SVID_ = "fdlibm_svid"

// /usr/include/math.h:203:9:
const m__XOPEN_ = "fdlibm_xopen"

// /usr/include/limits.h:107:9:
const m__XOPEN_IOV_MAX = 16

// /usr/include/limits.h:108:9:
const m__XOPEN_NAME_MAX = 256

// /usr/include/limits.h:109:9:
const m__XOPEN_PATH_MAX = 1024

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:3:9:
const m___ATOMIC_ACQUIRE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:324:9:
const m___ATOMIC_ACQ_REL = 4

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:301:9:
const m___ATOMIC_CONSUME = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:113:9:
const m___ATOMIC_HLE_ACQUIRE = 65536

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:110:9:
const m___ATOMIC_HLE_RELEASE = 131072

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:258:9:
const m___ATOMIC_RELAXED = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:325:9:
const m___ATOMIC_RELEASE = 3

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:311:9:
const m___ATOMIC_SEQ_CST = 5

// /usr/include/sys/cdefs.h:431:9:
const m___BEGIN_DECLS = "__BEGIN_PUBLIC_DECLS"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:90:9:
const m___BIGGEST_ALIGNMENT__ = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:283:9:
const m___BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"

// /usr/include/machine/byte_swap.h:63:9:
const m___BYTE_SWAP_U16_VARIABLE = "__byte_swap_u16_variable"

// /usr/include/machine/byte_swap.h:54:9:
const m___BYTE_SWAP_U32_VARIABLE = "__byte_swap_u32_variable"

// /usr/include/machine/byte_swap.h:45:9:
const m___BYTE_SWAP_U64_VARIABLE = "__byte_swap_u64_variable"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<command-line>:5:9:
const m___CCGO__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:10:9:
const m___CHAR_BIT__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:277:9:
const m___DBL_DECIMAL_DIG__ = 17

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:66:9:
const m___DBL_DIG__ = 15

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:139:9:
const m___DBL_HAS_DENORM__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:94:9:
const m___DBL_HAS_INFINITY__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:132:9:
const m___DBL_HAS_QUIET_NAN__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:238:9:
const m___DBL_MANT_DIG__ = 53

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:306:9:
const m___DBL_MAX_10_EXP__ = 308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:106:9:
const m___DBL_MAX_EXP__ = 1024

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:80:9:
const m___DECIMAL_DIG__ = 17

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:144:9:
const m___DEC_EVAL_METHOD__ = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:167:9:
const m___ELF__ = 1

// /usr/include/sys/cdefs.h:432:9:
const m___END_DECLS = "__END_PUBLIC_DECLS"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:32:9:
const m___FINITE_MATH_ONLY__ = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:153:9:
const m___FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:316:9:
const m___FLT128_DECIMAL_DIG__ = 36

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:46:9:
const m___FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:158:9:
const m___FLT128_DIG__ = 33

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:128:9:
const m___FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:155:9:
const m___FLT128_HAS_DENORM__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:105:9:
const m___FLT128_HAS_INFINITY__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:270:9:
const m___FLT128_HAS_QUIET_NAN__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:56:9:
const m___FLT128_MANT_DIG__ = 113

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:4:9:
const m___FLT128_MAX_10_EXP__ = 4932

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:151:9:
const m___FLT128_MAX_EXP__ = 16384

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:210:9:
const m___FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:260:9:
const m___FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:268:9:
const m___FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:25:9:
const m___FLT32X_DECIMAL_DIG__ = 17

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:305:9:
const m___FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:214:9:
const m___FLT32X_DIG__ = 15

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:133:9:
const m___FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:100:9:
const m___FLT32X_HAS_DENORM__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:74:9:
const m___FLT32X_HAS_INFINITY__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:300:9:
const m___FLT32X_HAS_QUIET_NAN__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:61:9:
const m___FLT32X_MANT_DIG__ = 53

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:254:9:
const m___FLT32X_MAX_10_EXP__ = 308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:33:9:
const m___FLT32X_MAX_EXP__ = 1024

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:281:9:
const m___FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:176:9:
const m___FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:290:9:
const m___FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:157:9:
const m___FLT32_DECIMAL_DIG__ = 9

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:202:9:
const m___FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:67:9:
const m___FLT32_DIG__ = 6

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:276:9:
const m___FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:34:9:
const m___FLT32_HAS_DENORM__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:231:9:
const m___FLT32_HAS_INFINITY__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:318:9:
const m___FLT32_HAS_QUIET_NAN__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:152:9:
const m___FLT32_MANT_DIG__ = 24

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:36:9:
const m___FLT32_MAX_10_EXP__ = 38

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:119:9:
const m___FLT32_MAX_EXP__ = 128

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:108:9:
const m___FLT32_MAX__ = 3.40282346638528859811704183484516925e+38

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:141:9:
const m___FLT32_MIN__ = 1.17549435082228750796873653722224568e-38

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:245:9:
const m___FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:136:9:
const m___FLT64X_DECIMAL_DIG__ = 36

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:73:9:
const m___FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:165:9:
const m___FLT64X_DIG__ = 33

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:64:9:
const m___FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:190:9:
const m___FLT64X_HAS_DENORM__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:98:9:
const m___FLT64X_HAS_INFINITY__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:43:9:
const m___FLT64X_HAS_QUIET_NAN__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:85:9:
const m___FLT64X_MANT_DIG__ = 113

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:42:9:
const m___FLT64X_MAX_10_EXP__ = 4932

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:232:9:
const m___FLT64X_MAX_EXP__ = 16384

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:218:9:
const m___FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:273:9:
const m___FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:129:9:
const m___FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:27:9:
const m___FLT64_DECIMAL_DIG__ = 17

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:126:9:
const m___FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:184:9:
const m___FLT64_DIG__ = 15

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:81:9:
const m___FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:275:9:
const m___FLT64_HAS_DENORM__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:217:9:
const m___FLT64_HAS_INFINITY__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:212:9:
const m___FLT64_HAS_QUIET_NAN__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:84:9:
const m___FLT64_MANT_DIG__ = 53

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:91:9:
const m___FLT64_MAX_10_EXP__ = 308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:263:9:
const m___FLT64_MAX_EXP__ = 1024

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:248:9:
const m___FLT64_MAX__ = 1.79769313486231570814527423731704357e+308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:297:9:
const m___FLT64_MIN__ = 2.22507385850720138309023271733240406e-308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:269:9:
const m___FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:319:9:
const m___FLT_DECIMAL_DIG__ = 9

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:286:9:
const m___FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:143:9:
const m___FLT_DIG__ = 6

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:68:9:
const m___FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:48:9:
const m___FLT_EVAL_METHOD_TS_18661_3__ = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:26:9:
const m___FLT_EVAL_METHOD__ = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:88:9:
const m___FLT_HAS_DENORM__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:191:9:
const m___FLT_HAS_INFINITY__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:187:9:
const m___FLT_HAS_QUIET_NAN__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:145:9:
const m___FLT_MANT_DIG__ = 24

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:188:9:
const m___FLT_MAX_10_EXP__ = 38

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:236:9:
const m___FLT_MAX_EXP__ = 128

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:292:9:
const m___FLT_MAX__ = 3.40282346638528859811704183484516925e+38

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:5:9:
const m___FLT_MIN__ = 1.17549435082228750796873653722224568e-38

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:230:9:
const m___FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:170:9:
const m___FLT_RADIX__ = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:34:9:
const m___FUNCTION__ = "__func__"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:280:9:
const m___FXSR__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:168:9:
const m___GCC_ASM_FLAG_OUTPUTS__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:45:9:
const m___GCC_ATOMIC_BOOL_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:241:9:
const m___GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:28:9:
const m___GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:23:9:
const m___GCC_ATOMIC_CHAR_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:150:9:
const m___GCC_ATOMIC_INT_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:317:9:
const m___GCC_ATOMIC_LLONG_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:221:9:
const m___GCC_ATOMIC_LONG_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:59:9:
const m___GCC_ATOMIC_POINTER_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:322:9:
const m___GCC_ATOMIC_SHORT_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:229:9:
const m___GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:178:9:
const m___GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:122:9:
const m___GCC_HAVE_DWARF2_CFI_ASM = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:18:9:
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:19:9:
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:20:9:
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:22:9:
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:24:9:
const m___GCC_IEC_559 = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:6:9:
const m___GCC_IEC_559_COMPLEX = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:302:9:
const m___GNUC_MINOR__ = 5

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:267:9:
const m___GNUC_PATCHLEVEL__ = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:274:9:
const m___GNUC_STDC_INLINE__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:86:9:
const m___GNUC__ = 10

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:120:9:
const m___GXX_ABI_VERSION = 1014

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:96:9:
const m___HAVE_SPECULATION_SAFE_VALUE = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:123:9:
const m___INT16_MAX__ = 0x7fff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:75:9:
const m___INT32_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:294:9:
const m___INT32_TYPE__ = "int"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:200:9:
const m___INT64_MAX__ = 0x7fffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:287:9:
const m___INT8_MAX__ = 127

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:271:9:
const m___INTMAX_MAX__ = 0x7fffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:284:9:
const m___INTMAX_WIDTH__ = 64

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:211:9:
const m___INTPTR_MAX__ = 0x7fffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:97:9:
const m___INTPTR_WIDTH__ = 64

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:183:9:
const m___INT_FAST16_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:101:9:
const m___INT_FAST16_TYPE__ = "int"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:303:9:
const m___INT_FAST16_WIDTH__ = 32

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:93:9:
const m___INT_FAST32_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:161:9:
const m___INT_FAST32_TYPE__ = "int"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:193:9:
const m___INT_FAST32_WIDTH__ = 32

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:228:9:
const m___INT_FAST64_MAX__ = 0x7fffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:227:9:
const m___INT_FAST64_WIDTH__ = 64

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:209:9:
const m___INT_FAST8_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:272:9:
const m___INT_FAST8_TYPE__ = "int"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:279:9:
const m___INT_FAST8_WIDTH__ = 32

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:199:9:
const m___INT_LEAST16_MAX__ = 0x7fff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:54:9:
const m___INT_LEAST16_WIDTH__ = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:149:9:
const m___INT_LEAST32_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:264:9:
const m___INT_LEAST32_TYPE__ = "int"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:298:9:
const m___INT_LEAST32_WIDTH__ = 32

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:240:9:
const m___INT_LEAST64_MAX__ = 0x7fffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:252:9:
const m___INT_LEAST64_WIDTH__ = 64

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:251:9:
const m___INT_LEAST8_MAX__ = 0x7f

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:38:9:
const m___INT_LEAST8_WIDTH__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:234:9:
const m___INT_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:77:9:
const m___INT_WIDTH__ = 32

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:146:9:
const m___LDBL_DECIMAL_DIG__ = 17

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:307:9:
const m___LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:180:9:
const m___LDBL_DIG__ = 15

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:172:9:
const m___LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:104:9:
const m___LDBL_HAS_DENORM__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:140:9:
const m___LDBL_HAS_INFINITY__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:83:9:
const m___LDBL_HAS_QUIET_NAN__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:216:9:
const m___LDBL_MANT_DIG__ = 53

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:257:9:
const m___LDBL_MAX_10_EXP__ = 308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:253:9:
const m___LDBL_MAX_EXP__ = 1024

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:41:9:
const m___LDBL_MAX__ = 1.79769313486231570814527423731704357e+308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:71:9:
const m___LDBL_MIN__ = 2.22507385850720138309023271733240406e-308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:321:9:
const m___LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:63:9:
const m___LONG_DOUBLE_64__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:114:9:
const m___LONG_LONG_MAX__ = 0x7fffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:118:9:
const m___LONG_LONG_WIDTH__ = 64

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:189:9:
const m___LONG_MAX__ = 0x7fffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:288:9:
const m___LONG_WIDTH__ = 64

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:131:9:
const m___LP64__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:102:9:
const m___MMX_WITH_SSE__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:87:9:
const m___MMX__ = 1

// /usr/include/sys/fd_set.h:49:9:
const m___NFDBITS = 32

// /usr/include/sys/fd_set.h:50:9:
const m___NFDSHIFT = 5

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:70:9:
const m___NetBSD__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:49:9:
const m___OPTIMIZE__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:237:9:
const m___ORDER_BIG_ENDIAN__ = 4321

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:15:9:
const m___ORDER_LITTLE_ENDIAN__ = 1234

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:160:9:
const m___ORDER_PDP_ENDIAN__ = 3412

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:196:9:
const m___PRAGMA_REDEFINE_EXTNAME = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:35:9:
const m___PRETTY_FUNCTION__ = "__func__"

// /usr/include/sys/cdefs.h:670:9:
const m___PRIuBIT = "PRIuMAX"

// /usr/include/sys/cdefs.h:671:9:
const m___PRIuBITS = "__PRIuBIT"

// /usr/include/sys/cdefs.h:673:9:
const m___PRIxBIT = "PRIxMAX"

// /usr/include/sys/cdefs.h:674:9:
const m___PRIxBITS = "__PRIxBIT"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:111:9:
const m___PTRDIFF_MAX__ = 0x7fffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:215:9:
const m___PTRDIFF_WIDTH__ = 64

// /usr/include/stdio.h:168:9:
const m___SALC = 0x4000

// /usr/include/stdio.h:162:9:
const m___SAPP = 0x0100

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:55:9:
const m___SCHAR_MAX__ = 0x7f

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:12:9:
const m___SCHAR_WIDTH__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:198:9:
const m___SEG_FS = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:201:9:
const m___SEG_GS = 1

// /usr/include/stdio.h:159:9:
const m___SEOF = 0x0020

// /usr/include/stdio.h:160:9:
const m___SERR = 0x0040

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:40:9:
const m___SHRT_MAX__ = 0x7fff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:69:9:
const m___SHRT_WIDTH__ = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:177:9:
const m___SIG_ATOMIC_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:30:9:
const m___SIG_ATOMIC_TYPE__ = "int"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:203:9:
const m___SIG_ATOMIC_WIDTH__ = 32

// /usr/include/machine/types.h:74:9:
const m___SIMPLELOCK_LOCKED = 1

// /usr/include/machine/types.h:75:9:
const m___SIMPLELOCK_UNLOCKED = 0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:295:9:
const m___SIZEOF_DOUBLE__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:239:9:
const m___SIZEOF_FLOAT128__ = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:8:9:
const m___SIZEOF_FLOAT80__ = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:95:9:
const m___SIZEOF_FLOAT__ = 4

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:256:9:
const m___SIZEOF_INT128__ = 16

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:208:9:
const m___SIZEOF_INT__ = 4

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:89:9:
const m___SIZEOF_LONG_DOUBLE__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:315:9:
const m___SIZEOF_LONG_LONG__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:78:9:
const m___SIZEOF_LONG__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:130:9:
const m___SIZEOF_POINTER__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:179:9:
const m___SIZEOF_PTRDIFF_T__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:244:9:
const m___SIZEOF_SHORT__ = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:115:9:
const m___SIZEOF_SIZE_T__ = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:265:9:
const m___SIZEOF_WCHAR_T__ = 4

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:117:9:
const m___SIZEOF_WINT_T__ = 4

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:16:9:
const m___SIZE_MAX__ = 0xffffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:197:9:
const m___SIZE_WIDTH__ = 64

// /usr/include/stdio.h:153:9:
const m___SLBF = 0x0001

// /usr/include/stdio.h:161:9:
const m___SMBF = 0x0080

// /usr/include/stdio.h:167:9:
const m___SMOD = 0x2000

// /usr/include/stdio.h:154:9:
const m___SNBF = 0x0002

// /usr/include/stdio.h:165:9:
const m___SNPT = 0x0800

// /usr/include/stdio.h:166:9:
const m___SOFF = 0x1000

// /usr/include/stdio.h:164:9:
const m___SOPT = 0x0400

// /usr/include/stdio.h:155:9:
const m___SRD = 0x0004

// /usr/include/stdio.h:158:9:
const m___SRW = 0x0010

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:109:9:
const m___SSE2_MATH__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:293:9:
const m___SSE2__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:174:9:
const m___SSE_MATH__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:246:9:
const m___SSE__ = 1

// /usr/include/stdio.h:163:9:
const m___SSTR = 0x0200

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:65:9:
const m___STDC_HOSTED__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:72:9:
const m___STDC_UTF_16__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:278:9:
const m___STDC_UTF_32__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:207:9:
const m___STDC_VERSION__ = 201710

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:309:9:
const m___STDC__ = 1

// /usr/include/stdio.h:156:9:
const m___SWR = 0x0008

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:138:9:
const m___UINT16_MAX__ = 0xffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:50:9:
const m___UINT32_MAX__ = 0xffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:164:9:
const m___UINT64_MAX__ = 0xffffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:11:9:
const m___UINT8_MAX__ = 0xff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:304:9:
const m___UINTMAX_MAX__ = 0xffffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:226:9:
const m___UINTPTR_MAX__ = 0xffffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:320:9:
const m___UINT_FAST16_MAX__ = 0xffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:185:9:
const m___UINT_FAST32_MAX__ = 0xffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:29:9:
const m___UINT_FAST64_MAX__ = 0xffffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:35:9:
const m___UINT_FAST8_MAX__ = 0xffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:2:9:
const m___UINT_LEAST16_MAX__ = 0xffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:99:9:
const m___UINT_LEAST32_MAX__ = 0xffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:39:9:
const m___UINT_LEAST64_MAX__ = 0xffffffffffffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:44:9:
const m___UINT_LEAST8_MAX__ = 0xff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:147:9:
const m___VERSION__ = "10.5.0"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:17:9:
const m___WCHAR_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:225:9:
const m___WCHAR_TYPE__ = "int"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:107:9:
const m___WCHAR_WIDTH__ = 32

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:13:9:
const m___WINT_MAX__ = 0x7fffffff

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:242:9:
const m___WINT_TYPE__ = "int"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:250:9:
const m___WINT_WIDTH__ = 32

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:112:9:
const m___amd64 = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:249:9:
const m___amd64__ = 1

// /usr/include/assert.h:91:9:
const m___assert_function__ = "__func__"

// /usr/include/stdbool.h:44:9:
const m___bool_true_false_are_defined = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:220:9:
const m___code_model_small__ = 1

// /usr/include/sys/cdefs.h:114:9:
const m___const = "const"

// /usr/include/sys/cdefs.h:331:9:
const m___debugused = "__unused"

// /usr/include/sys/cdefs.h:321:9:
const m___diagused = "__unused"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:175:9:
const m___k8 = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:222:9:
const m___k8__ = 1

// /usr/include/pthread_types.h:49:9:
const m___pthread_volatile = "volatile"

// /usr/include/sys/cdefs.h:473:9:
const m___restrict = "restrict"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:3:9:
const m___restrict_arr = "restrict"

// /usr/include/sys/cdefs.h:115:9:
const m___signed = "signed"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:255:9:
const m___syslog_attribute__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:154:9:
const m___tune_nocona__ = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:76:9:
const m___unix__ = 1

// /usr/include/sys/cdefs.h:116:9:
const m___volatile = "volatile"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<command-line>:2:9:
const m___wasi__ = 1

// /usr/include/sys/cdefs_elf.h:80:9:
const m___weakref_visible = "static"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:124:9:
const m___x86_64 = 1

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<predefined>:181:9:
const m___x86_64__ = 1

// /usr/include/sys/types.h:186:9:
const m_accmode_t = "__accmode_t"

// /usr/include/stdbool.h:37:9:
const m_bool = "_Bool"

// /usr/include/sys/types.h:146:9:
const m_caddr_t = "__caddr_t"

// /usr/include/sys/types.h:271:9:
const m_devmajor_t = "__devmajor_t"

// /usr/include/sys/types.h:272:9:
const m_devminor_t = "__devminor_t"

// /usr/include/stdbool.h:40:9:
const m_false = 0

// /usr/include/sys/fd_set.h:92:9:
const m_fd_mask = "__fd_mask"

// /usr/include/sys/types.h:134:9:
const m_fsblkcnt_t = "__fsblkcnt_t"

// /usr/include/sys/types.h:139:9:
const m_fsfilcnt_t = "__fsfilcnt_t"

// /usr/include/sys/types.h:162:9:
const m_gid_t = "__gid_t"

// /usr/include/sys/endian.h:61:9:
const m_in_addr_t = "__in_addr_t"

// /usr/include/sys/endian.h:66:9:
const m_in_port_t = "__in_port_t"

// /usr/include/math.h:159:9:
const m_math_errhandling = "MATH_ERREXCEPT"

// /usr/include/sys/types.h:181:9:
const m_mode_t = "__mode_t"

// /usr/include/sys/types.h:193:9:
const m_off_t = "__off_t"

// /usr/include/sys/types.h:198:9:
const m_pid_t = "__pid_t"

// /usr/include/assert.h:110:9:
const m_static_assert = "_Static_assert"

// /usr/include/stdbool.h:39:9:
const m_true = 1

// /usr/include/sys/types.h:207:9:
const m_uid_t = "__uid_t"

// /usr/include/stdarg.h:58:9:
const m_va_arg = "__builtin_va_arg"

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:7:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:7:14:
type t__builtin_va_list = uintptr

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:23:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:23:23:
type t__predefined_size_t = uint64

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:27:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:27:24:
type t__predefined_wchar_t = int32

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:31:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/<builtin>:31:26:
type t__predefined_ptrdiff_t = int64

//
// /usr/include/sys/common_int_types.h:45:1:

// /usr/include/sys/common_int_types.h:45:27:
type t__int8_t = int8

//
// /usr/include/sys/common_int_types.h:46:1:

// /usr/include/sys/common_int_types.h:46:27:
type t__uint8_t = uint8

//
// /usr/include/sys/common_int_types.h:47:1:

// /usr/include/sys/common_int_types.h:47:27:
type t__int16_t = int16

//
// /usr/include/sys/common_int_types.h:48:1:

// /usr/include/sys/common_int_types.h:48:27:
type t__uint16_t = uint16

//
// /usr/include/sys/common_int_types.h:49:1:

// /usr/include/sys/common_int_types.h:49:27:
type t__int32_t = int32

//
// /usr/include/sys/common_int_types.h:50:1:

// /usr/include/sys/common_int_types.h:50:27:
type t__uint32_t = uint32

//
// /usr/include/sys/common_int_types.h:51:1:

// /usr/include/sys/common_int_types.h:51:27:
type t__int64_t = int64

//
// /usr/include/sys/common_int_types.h:52:1:

// /usr/include/sys/common_int_types.h:52:27:
type t__uint64_t = uint64

//
// /usr/include/sys/common_int_types.h:58:1:

// /usr/include/sys/common_int_types.h:58:27:
type t__intptr_t = int64

//
// /usr/include/sys/common_int_types.h:59:1:

// /usr/include/sys/common_int_types.h:59:26:
type t__uintptr_t = uint64

//
// /usr/include/sys/ansi.h:37:1:

// /usr/include/sys/ansi.h:37:14:
type t__caddr_t = uintptr

//
// /usr/include/sys/ansi.h:38:1:

// /usr/include/sys/ansi.h:38:20:
type t__gid_t = uint32

//
// /usr/include/sys/ansi.h:39:1:

// /usr/include/sys/ansi.h:39:20:
type t__in_addr_t = uint32

//
// /usr/include/sys/ansi.h:40:1:

// /usr/include/sys/ansi.h:40:20:
type t__in_port_t = uint16

//
// /usr/include/sys/ansi.h:41:1:

// /usr/include/sys/ansi.h:41:20:
type t__mode_t = uint32

//
// /usr/include/sys/ansi.h:42:1:

// /usr/include/sys/ansi.h:42:20:
type t__accmode_t = uint32

//
// /usr/include/sys/ansi.h:43:1:

// /usr/include/sys/ansi.h:43:19:
type t__off_t = int64

//
// /usr/include/sys/ansi.h:44:1:

// /usr/include/sys/ansi.h:44:19:
type t__pid_t = int32

//
// /usr/include/sys/ansi.h:45:1:

// /usr/include/sys/ansi.h:45:19:
type t__sa_family_t = uint8

//
// /usr/include/sys/ansi.h:46:1:

// /usr/include/sys/ansi.h:46:22:
type t__socklen_t = uint32

//
// /usr/include/sys/ansi.h:47:1:

// /usr/include/sys/ansi.h:47:20:
type t__uid_t = uint32

//
// /usr/include/sys/ansi.h:48:1:

// /usr/include/sys/ansi.h:48:20:
type t__fsblkcnt_t = uint64

//
// /usr/include/sys/ansi.h:49:1:

// /usr/include/sys/ansi.h:49:20:
type t__fsfilcnt_t = uint64

//
// /usr/include/sys/ansi.h:52:1:

// /usr/include/sys/ansi.h:52:32:
type t__wctrans_t = uintptr

//
// /usr/include/sys/ansi.h:55:1:

// /usr/include/sys/ansi.h:55:31:
type t__wctype_t = uintptr

//
// /usr/include/sys/ansi.h:61:1:

// /usr/include/sys/ansi.h:64:3:
type t__mbstate_t = struct {
	F__mbstate8  [0][128]int8
	F__mbstateL  t__int64_t
	F__ccgo_pad2 [120]byte
}

//
// /usr/include/sys/ansi.h:73:1:

// /usr/include/sys/ansi.h:73:27:
type t__va_list = uintptr

//
// /usr/include/stdarg.h:53:1:

// /usr/include/stdarg.h:53:19:
type Tva_list = uintptr

type va_list = Tva_list

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:304:3:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:304:25:
type Tsqlite_int64 = int64

type sqlite_int64 = Tsqlite_int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:305:3:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:305:34:
type Tsqlite_uint64 = uint64

type sqlite_uint64 = Tsqlite_uint64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:307:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:307:22:
type Tsqlite3_int64 = int64

type sqlite3_int64 = Tsqlite3_int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:308:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:308:23:
type Tsqlite3_uint64 = uint64

type sqlite3_uint64 = Tsqlite3_uint64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:364:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:364:13:
type Tsqlite3_callback = uintptr

type sqlite3_callback = Tsqlite3_callback

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:746:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:746:29:
type Tsqlite3_file = struct {
	FpMethods uintptr
}

type sqlite3_file = Tsqlite3_file

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:853:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:853:35:
type Tsqlite3_io_methods = struct {
	FiVersion               int32
	FxClose                 uintptr
	FxRead                  uintptr
	FxWrite                 uintptr
	FxTruncate              uintptr
	FxSync                  uintptr
	FxFileSize              uintptr
	FxLock                  uintptr
	FxUnlock                uintptr
	FxCheckReservedLock     uintptr
	FxFileControl           uintptr
	FxSectorSize            uintptr
	FxDeviceCharacteristics uintptr
	FxShmMap                uintptr
	FxShmLock               uintptr
	FxShmBarrier            uintptr
	FxShmUnmap              uintptr
	FxFetch                 uintptr
	FxUnfetch               uintptr
}

type sqlite3_io_methods = Tsqlite3_io_methods

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1340:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1340:20:
type Tsqlite3_filename = uintptr

type sqlite3_filename = Tsqlite3_filename

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1511:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1511:28:
type Tsqlite3_vfs = struct {
	FiVersion          int32
	FszOsFile          int32
	FmxPathname        int32
	FpNext             uintptr
	FzName             uintptr
	FpAppData          uintptr
	FxOpen             uintptr
	FxDelete           uintptr
	FxAccess           uintptr
	FxFullPathname     uintptr
	FxDlOpen           uintptr
	FxDlError          uintptr
	FxDlSym            uintptr
	FxDlClose          uintptr
	FxRandomness       uintptr
	FxSleep            uintptr
	FxCurrentTime      uintptr
	FxGetLastError     uintptr
	FxCurrentTimeInt64 uintptr
	FxSetSystemCall    uintptr
	FxGetSystemCall    uintptr
	FxNextSystemCall   uintptr
}

type sqlite3_vfs = Tsqlite3_vfs

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1512:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1512:14:
type Tsqlite3_syscall_ptr = uintptr

type sqlite3_syscall_ptr = Tsqlite3_syscall_ptr

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1813:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:1813:36:
type Tsqlite3_mem_methods = struct {
	FxMalloc   uintptr
	FxFree     uintptr
	FxRealloc  uintptr
	FxSize     uintptr
	FxRoundup  uintptr
	FxInit     uintptr
	FxShutdown uintptr
	FpAppData  uintptr
}

type sqlite3_mem_methods = Tsqlite3_mem_methods

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:6426:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:6426:14:
type Tsqlite3_destructor_type = uintptr

type sqlite3_destructor_type = Tsqlite3_destructor_type

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7663:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7663:29:
type Tsqlite3_vtab = struct {
	FpModule uintptr
	FnRef    int32
	FzErrMsg uintptr
}

type sqlite3_vtab = Tsqlite3_vtab

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7664:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7664:35:
type Tsqlite3_index_info = struct {
	FnConstraint      int32
	FaConstraint      uintptr
	FnOrderBy         int32
	FaOrderBy         uintptr
	FaConstraintUsage uintptr
	FidxNum           int32
	FidxStr           uintptr
	FneedToFreeIdxStr int32
	ForderByConsumed  int32
	FestimatedCost    float64
	FestimatedRows    Tsqlite3_int64
	FidxFlags         int32
	FcolUsed          Tsqlite3_uint64
}

type sqlite3_index_info = Tsqlite3_index_info

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7665:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7665:36:
type Tsqlite3_vtab_cursor = struct {
	FpVtab uintptr
}

type sqlite3_vtab_cursor = Tsqlite3_vtab_cursor

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7666:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:7666:31:
type Tsqlite3_module = struct {
	FiVersion      int32
	FxCreate       uintptr
	FxConnect      uintptr
	FxBestIndex    uintptr
	FxDisconnect   uintptr
	FxDestroy      uintptr
	FxOpen         uintptr
	FxClose        uintptr
	FxFilter       uintptr
	FxNext         uintptr
	FxEof          uintptr
	FxColumn       uintptr
	FxRowid        uintptr
	FxUpdate       uintptr
	FxBegin        uintptr
	FxSync         uintptr
	FxCommit       uintptr
	FxRollback     uintptr
	FxFindFunction uintptr
	FxRename       uintptr
	FxSavepoint    uintptr
	FxRelease      uintptr
	FxRollbackTo   uintptr
	FxShadowName   uintptr
	FxIntegrity    uintptr
}

type sqlite3_module = Tsqlite3_module

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8523:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:8523:38:
type Tsqlite3_mutex_methods = struct {
	FxMutexInit    uintptr
	FxMutexEnd     uintptr
	FxMutexAlloc   uintptr
	FxMutexFree    uintptr
	FxMutexEnter   uintptr
	FxMutexTry     uintptr
	FxMutexLeave   uintptr
	FxMutexHeld    uintptr
	FxMutexNotheld uintptr
}

type sqlite3_mutex_methods = Tsqlite3_mutex_methods

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9335:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9335:36:
type Tsqlite3_pcache_page = struct {
	FpBuf   uintptr
	FpExtra uintptr
}

type sqlite3_pcache_page = Tsqlite3_pcache_page

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9500:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9500:40:
type Tsqlite3_pcache_methods2 = struct {
	FiVersion   int32
	FpArg       uintptr
	FxInit      uintptr
	FxShutdown  uintptr
	FxCreate    uintptr
	FxCachesize uintptr
	FxPagecount uintptr
	FxFetch     uintptr
	FxUnpin     uintptr
	FxRekey     uintptr
	FxTruncate  uintptr
	FxDestroy   uintptr
	FxShrink    uintptr
}

type sqlite3_pcache_methods2 = Tsqlite3_pcache_methods2

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9523:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:9523:39:
type Tsqlite3_pcache_methods = struct {
	FpArg       uintptr
	FxInit      uintptr
	FxShutdown  uintptr
	FxCreate    uintptr
	FxCachesize uintptr
	FxPagecount uintptr
	FxFetch     uintptr
	FxUnpin     uintptr
	FxRekey     uintptr
	FxTruncate  uintptr
	FxDestroy   uintptr
}

type sqlite3_pcache_methods = Tsqlite3_pcache_methods

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10951:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:10953:3:
type Tsqlite3_snapshot = struct {
	Fhidden [48]uint8
}

type sqlite3_snapshot = Tsqlite3_snapshot

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11391:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11391:39:
type Tsqlite3_rtree_geometry = struct {
	FpContext uintptr
	FnParam   int32
	FaParam   uintptr
	FpUser    uintptr
	FxDelUser uintptr
}

type sqlite3_rtree_geometry = Tsqlite3_rtree_geometry

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11392:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11392:41:
type Tsqlite3_rtree_query_info = struct {
	FpContext      uintptr
	FnParam        int32
	FaParam        uintptr
	FpUser         uintptr
	FxDelUser      uintptr
	FaCoord        uintptr
	FanQueue       uintptr
	FnCoord        int32
	FiLevel        int32
	FmxLevel       int32
	FiRowid        Tsqlite3_int64
	FrParentScore  Tsqlite3_rtree_dbl
	FeParentWithin int32
	FeWithin       int32
	FrScore        Tsqlite3_rtree_dbl
	FapSqlParam    uintptr
}

type sqlite3_rtree_query_info = Tsqlite3_rtree_query_info

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11400:3:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:11400:18:
type Tsqlite3_rtree_dbl = float64

type sqlite3_rtree_dbl = Tsqlite3_rtree_dbl

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:13627:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:13627:33:
type TFts5ExtensionApi = struct {
	FiVersion           int32
	FxUserData          uintptr
	FxColumnCount       uintptr
	FxRowCount          uintptr
	FxColumnTotalSize   uintptr
	FxTokenize          uintptr
	FxPhraseCount       uintptr
	FxPhraseSize        uintptr
	FxInstCount         uintptr
	FxInst              uintptr
	FxRowid             uintptr
	FxColumnText        uintptr
	FxColumnSize        uintptr
	FxQueryPhrase       uintptr
	FxSetAuxdata        uintptr
	FxGetAuxdata        uintptr
	FxPhraseFirst       uintptr
	FxPhraseNext        uintptr
	FxPhraseFirstColumn uintptr
	FxPhraseNextColumn  uintptr
	FxQueryToken        uintptr
	FxInstToken         uintptr
	FxColumnLocale      uintptr
	FxTokenize_v2       uintptr
}

type Fts5ExtensionApi = TFts5ExtensionApi

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:13629:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:13629:31:
type TFts5PhraseIter = struct {
	Fa uintptr
	Fb uintptr
}

type Fts5PhraseIter = TFts5PhraseIter

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:13631:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:13631:14:
type Tfts5_extension_function = uintptr

type fts5_extension_function = Tfts5_extension_function

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14222:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14222:34:
type Tfts5_tokenizer_v2 = struct {
	FiVersion  int32
	FxCreate   uintptr
	FxDelete   uintptr
	FxTokenize uintptr
}

type fts5_tokenizer_v2 = Tfts5_tokenizer_v2

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14249:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14249:31:
type Tfts5_tokenizer = struct {
	FxCreate   uintptr
	FxDelete   uintptr
	FxTokenize uintptr
}

type fts5_tokenizer = Tfts5_tokenizer

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14286:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libsqlite3/include/sqlite3.h:14286:25:
type Tfts5_api = struct {
	FiVersion            int32
	FxCreateTokenizer    uintptr
	FxFindTokenizer      uintptr
	FxCreateFunction     uintptr
	FxCreateTokenizer_v2 uintptr
	FxFindTokenizer_v2   uintptr
}

type fts5_api = Tfts5_api

//TODO "sys_errlist" // /usr/include/errno.h:61:19:

//
// /usr/include/sys/stdint.h:39:1:

// /usr/include/sys/stdint.h:39:18:
type Tint8_t = int8

type int8_t = Tint8_t

//
// /usr/include/sys/stdint.h:44:1:

// /usr/include/sys/stdint.h:44:19:
type Tuint8_t = uint8

type uint8_t = Tuint8_t

//
// /usr/include/sys/stdint.h:49:1:

// /usr/include/sys/stdint.h:49:19:
type Tint16_t = int16

type int16_t = Tint16_t

//
// /usr/include/sys/stdint.h:54:1:

// /usr/include/sys/stdint.h:54:20:
type Tuint16_t = uint16

type uint16_t = Tuint16_t

//
// /usr/include/sys/stdint.h:59:1:

// /usr/include/sys/stdint.h:59:19:
type Tint32_t = int32

type int32_t = Tint32_t

//
// /usr/include/sys/stdint.h:64:1:

// /usr/include/sys/stdint.h:64:20:
type Tuint32_t = uint32

type uint32_t = Tuint32_t

//
// /usr/include/sys/stdint.h:69:1:

// /usr/include/sys/stdint.h:69:19:
type Tint64_t = int64

type int64_t = Tint64_t

//
// /usr/include/sys/stdint.h:74:1:

// /usr/include/sys/stdint.h:74:20:
type Tuint64_t = uint64

type uint64_t = Tuint64_t

//
// /usr/include/sys/stdint.h:79:1:

// /usr/include/sys/stdint.h:79:20:
type Tintptr_t = int64

type intptr_t = Tintptr_t

//
// /usr/include/sys/stdint.h:84:1:

// /usr/include/sys/stdint.h:84:21:
type Tuintptr_t = uint64

type uintptr_t = Tuintptr_t

//
// /usr/include/sys/common_int_mwgwtypes.h:45:1:

// /usr/include/sys/common_int_mwgwtypes.h:45:32:
type Tint_least8_t = int8

type int_least8_t = Tint_least8_t

//
// /usr/include/sys/common_int_mwgwtypes.h:46:1:

// /usr/include/sys/common_int_mwgwtypes.h:46:32:
type Tuint_least8_t = uint8

type uint_least8_t = Tuint_least8_t

//
// /usr/include/sys/common_int_mwgwtypes.h:47:1:

// /usr/include/sys/common_int_mwgwtypes.h:47:32:
type Tint_least16_t = int16

type int_least16_t = Tint_least16_t

//
// /usr/include/sys/common_int_mwgwtypes.h:48:1:

// /usr/include/sys/common_int_mwgwtypes.h:48:32:
type Tuint_least16_t = uint16

type uint_least16_t = Tuint_least16_t

//
// /usr/include/sys/common_int_mwgwtypes.h:49:1:

// /usr/include/sys/common_int_mwgwtypes.h:49:32:
type Tint_least32_t = int32

type int_least32_t = Tint_least32_t

//
// /usr/include/sys/common_int_mwgwtypes.h:50:1:

// /usr/include/sys/common_int_mwgwtypes.h:50:32:
type Tuint_least32_t = uint32

type uint_least32_t = Tuint_least32_t

//
// /usr/include/sys/common_int_mwgwtypes.h:51:1:

// /usr/include/sys/common_int_mwgwtypes.h:51:32:
type Tint_least64_t = int64

type int_least64_t = Tint_least64_t

//
// /usr/include/sys/common_int_mwgwtypes.h:52:1:

// /usr/include/sys/common_int_mwgwtypes.h:52:32:
type Tuint_least64_t = uint64

type uint_least64_t = Tuint_least64_t

//
// /usr/include/sys/common_int_mwgwtypes.h:55:1:

// /usr/include/sys/common_int_mwgwtypes.h:55:32:
type Tint_fast8_t = int32

type int_fast8_t = Tint_fast8_t

//
// /usr/include/sys/common_int_mwgwtypes.h:56:1:

// /usr/include/sys/common_int_mwgwtypes.h:56:32:
type Tuint_fast8_t = uint32

type uint_fast8_t = Tuint_fast8_t

//
// /usr/include/sys/common_int_mwgwtypes.h:57:1:

// /usr/include/sys/common_int_mwgwtypes.h:57:32:
type Tint_fast16_t = int32

type int_fast16_t = Tint_fast16_t

//
// /usr/include/sys/common_int_mwgwtypes.h:58:1:

// /usr/include/sys/common_int_mwgwtypes.h:58:32:
type Tuint_fast16_t = uint32

type uint_fast16_t = Tuint_fast16_t

//
// /usr/include/sys/common_int_mwgwtypes.h:59:1:

// /usr/include/sys/common_int_mwgwtypes.h:59:32:
type Tint_fast32_t = int32

type int_fast32_t = Tint_fast32_t

//
// /usr/include/sys/common_int_mwgwtypes.h:60:1:

// /usr/include/sys/common_int_mwgwtypes.h:60:32:
type Tuint_fast32_t = uint32

type uint_fast32_t = Tuint_fast32_t

//
// /usr/include/sys/common_int_mwgwtypes.h:61:1:

// /usr/include/sys/common_int_mwgwtypes.h:61:32:
type Tint_fast64_t = int64

type int_fast64_t = Tint_fast64_t

//
// /usr/include/sys/common_int_mwgwtypes.h:62:1:

// /usr/include/sys/common_int_mwgwtypes.h:62:32:
type Tuint_fast64_t = uint64

type uint_fast64_t = Tuint_fast64_t

//
// /usr/include/sys/common_int_mwgwtypes.h:66:1:

// /usr/include/sys/common_int_mwgwtypes.h:66:33:
type Tintmax_t = int64

type intmax_t = Tintmax_t

//
// /usr/include/sys/common_int_mwgwtypes.h:67:1:

// /usr/include/sys/common_int_mwgwtypes.h:67:32:
type Tuintmax_t = uint64

type uintmax_t = Tuintmax_t

//
// /usr/include/inttypes.h:41:1:

// /usr/include/inttypes.h:41:23:
type Twchar_t = int32

type wchar_t = Twchar_t

//
// /usr/include/inttypes.h:57:1:

// /usr/include/inttypes.h:60:3:
type Timaxdiv_t = struct {
	Fquot Tintmax_t
	Frem  Tintmax_t
}

type imaxdiv_t = Timaxdiv_t

//
// /usr/include/inttypes.h:66:1:

// /usr/include/inttypes.h:66:25:
type Tlocale_t = uintptr

type locale_t = Tlocale_t

// /usr/include/math.h:24:1:
type t__float_u = struct {
	F__val   [0]float32
	F__dummy [4]uint8
}

// /usr/include/math.h:29:1:
type t__double_u = struct {
	F__val   [0]float64
	F__dummy [8]uint8
}

// /usr/include/math.h:34:1:
type t__long_double_u = struct {
	F__val   [0]float64
	F__dummy [8]uint8
}

//
// /usr/include/math.h:52:1:

// /usr/include/math.h:52:16:
type Tdouble_t = float64

type double_t = Tdouble_t

//
// /usr/include/math.h:53:1:

// /usr/include/math.h:53:15:
type Tfloat_t = float32

type float_t = Tfloat_t

// /usr/include/math.h:187:1:
type _fdversion = int32

const
// /usr/include/math.h:187:17:
_fdlibm_ieee = -1
const
// /usr/include/math.h:187:35:
_fdlibm_svid = 0
const
// /usr/include/math.h:187:48:
_fdlibm_xopen = 1
const
// /usr/include/math.h:187:62:
_fdlibm_posix = 2

// /usr/include/math.h:207:1:
type Texception = struct {
	Ftype1  int32
	Fname   uintptr
	Farg1   float64
	Farg2   float64
	Fretval float64
}

type exception = Texception

//
// /usr/include/machine/types.h:68:1:

// /usr/include/machine/types.h:68:19:
type t__register_t = int64

//
// /usr/include/machine/types.h:69:1:

// /usr/include/machine/types.h:69:24:
type t__cpu_simple_lock_nv_t = uint8

//
// /usr/include/sys/types.h:93:1:

// /usr/include/sys/types.h:93:18:
type Tu_int8_t = uint8

type u_int8_t = Tu_int8_t

//
// /usr/include/sys/types.h:94:1:

// /usr/include/sys/types.h:94:18:
type Tu_int16_t = uint16

type u_int16_t = Tu_int16_t

//
// /usr/include/sys/types.h:95:1:

// /usr/include/sys/types.h:95:18:
type Tu_int32_t = uint32

type u_int32_t = Tu_int32_t

//
// /usr/include/sys/types.h:96:1:

// /usr/include/sys/types.h:96:18:
type Tu_int64_t = uint64

type u_int64_t = Tu_int64_t

//
// /usr/include/sys/endian.h:60:1:

// /usr/include/sys/endian.h:60:21:
type Tin_addr_t = uint32

type in_addr_t = Tin_addr_t

//
// /usr/include/sys/endian.h:65:1:

// /usr/include/sys/endian.h:65:21:
type Tin_port_t = uint16

type in_port_t = Tin_port_t

//
// /usr/include/sys/types.h:101:1:

// /usr/include/sys/types.h:101:23:
type Tu_char = uint8

type u_char = Tu_char

//
// /usr/include/sys/types.h:102:1:

// /usr/include/sys/types.h:102:24:
type Tu_short = uint16

type u_short = Tu_short

//
// /usr/include/sys/types.h:103:1:

// /usr/include/sys/types.h:103:22:
type Tu_int = uint32

type u_int = Tu_int

//
// /usr/include/sys/types.h:104:1:

// /usr/include/sys/types.h:104:23:
type Tu_long = uint64

type u_long = Tu_long

//
// /usr/include/sys/types.h:106:1:

// /usr/include/sys/types.h:106:23:
type Tunchar = uint8

type unchar = Tunchar

//
// /usr/include/sys/types.h:107:1:

// /usr/include/sys/types.h:107:24:
type Tushort = uint16

type ushort = Tushort

//
// /usr/include/sys/types.h:108:1:

// /usr/include/sys/types.h:108:22:
type Tuint = uint32

type uint = Tuint

//
// /usr/include/sys/types.h:109:1:

// /usr/include/sys/types.h:109:23:
type Tulong = uint64

type ulong = Tulong

//
// /usr/include/sys/types.h:112:1:

// /usr/include/sys/types.h:112:18:
type Tu_quad_t = uint64

type u_quad_t = Tu_quad_t

//
// /usr/include/sys/types.h:113:1:

// /usr/include/sys/types.h:113:18:
type Tquad_t = int64

type quad_t = Tquad_t

//
// /usr/include/sys/types.h:114:1:

// /usr/include/sys/types.h:114:16:
type Tqaddr_t = uintptr

type qaddr_t = Tqaddr_t

//
// /usr/include/sys/types.h:126:1:

// /usr/include/sys/types.h:126:18:
type Tlonglong_t = int64

type longlong_t = Tlonglong_t

//
// /usr/include/sys/types.h:127:1:

// /usr/include/sys/types.h:127:18:
type Tu_longlong_t = uint64

type u_longlong_t = Tu_longlong_t

//
// /usr/include/sys/types.h:129:1:

// /usr/include/sys/types.h:129:18:
type Tblkcnt_t = int64

type blkcnt_t = Tblkcnt_t

//
// /usr/include/sys/types.h:130:1:

// /usr/include/sys/types.h:130:18:
type Tblksize_t = int32

type blksize_t = Tblksize_t

//
// /usr/include/sys/types.h:133:1:

// /usr/include/sys/types.h:133:22:
type Tfsblkcnt_t = uint64

type fsblkcnt_t = Tfsblkcnt_t

//
// /usr/include/sys/types.h:138:1:

// /usr/include/sys/types.h:138:22:
type Tfsfilcnt_t = uint64

type fsfilcnt_t = Tfsfilcnt_t

//
// /usr/include/sys/types.h:145:1:

// /usr/include/sys/types.h:145:19:
type Tcaddr_t = uintptr

type caddr_t = Tcaddr_t

//
// /usr/include/sys/types.h:154:1:

// /usr/include/sys/types.h:154:18:
type Tdaddr_t = int64

type daddr_t = Tdaddr_t

//
// /usr/include/sys/types.h:157:1:

// /usr/include/sys/types.h:157:18:
type Tdev_t = uint64

type dev_t = Tdev_t

//
// /usr/include/sys/types.h:158:1:

// /usr/include/sys/types.h:158:18:
type Tfixpt_t = uint32

type fixpt_t = Tfixpt_t

//
// /usr/include/sys/types.h:161:1:

// /usr/include/sys/types.h:161:18:
type Tgid_t = uint32

type gid_t = Tgid_t

//
// /usr/include/sys/types.h:165:1:

// /usr/include/sys/types.h:165:18:
type Tid_t = uint32

type id_t = Tid_t

//
// /usr/include/sys/types.h:175:1:

// /usr/include/sys/types.h:175:18:
type Tino_t = uint64

type ino_t = Tino_t

//
// /usr/include/sys/types.h:177:1:

// /usr/include/sys/types.h:177:15:
type Tkey_t = int64

type key_t = Tkey_t

//
// /usr/include/sys/types.h:180:1:

// /usr/include/sys/types.h:180:18:
type Tmode_t = uint32

type mode_t = Tmode_t

//
// /usr/include/sys/types.h:185:1:

// /usr/include/sys/types.h:185:21:
type Taccmode_t = uint32

type accmode_t = Taccmode_t

//
// /usr/include/sys/types.h:189:1:

// /usr/include/sys/types.h:189:18:
type Tnlink_t = uint32

type nlink_t = Tnlink_t

//
// /usr/include/sys/types.h:192:1:

// /usr/include/sys/types.h:192:18:
type Toff_t = int64

type off_t = Toff_t

//
// /usr/include/sys/types.h:197:1:

// /usr/include/sys/types.h:197:18:
type Tpid_t = int32

type pid_t = Tpid_t

//
// /usr/include/sys/types.h:200:1:

// /usr/include/sys/types.h:200:18:
type Tlwpid_t = int32

type lwpid_t = Tlwpid_t

//
// /usr/include/sys/types.h:201:1:

// /usr/include/sys/types.h:201:18:
type Trlim_t = uint64

type rlim_t = Trlim_t

//
// /usr/include/sys/types.h:202:1:

// /usr/include/sys/types.h:202:18:
type Tsegsz_t = int32

type segsz_t = Tsegsz_t

//
// /usr/include/sys/types.h:203:1:

// /usr/include/sys/types.h:203:18:
type Tswblk_t = int32

type swblk_t = Tswblk_t

//
// /usr/include/sys/types.h:206:1:

// /usr/include/sys/types.h:206:18:
type Tuid_t = uint32

type uid_t = Tuid_t

//
// /usr/include/sys/types.h:210:1:

// /usr/include/sys/types.h:210:14:
type Tmqd_t = int32

type mqd_t = Tmqd_t

//
// /usr/include/sys/types.h:212:1:

// /usr/include/sys/types.h:212:23:
type Tcpuid_t = uint64

type cpuid_t = Tcpuid_t

//
// /usr/include/sys/types.h:214:1:

// /usr/include/sys/types.h:214:14:
type Tpsetid_t = int32

type psetid_t = Tpsetid_t

//
// /usr/include/sys/types.h:216:1:

// /usr/include/sys/types.h:216:41:
type t__cpu_simple_lock_t = uint8

//
// /usr/include/sys/types.h:270:1:

// /usr/include/sys/types.h:270:17:
type t__devmajor_t = int32

// /usr/include/sys/types.h:270:31:
type t__devminor_t = int32

//
// /usr/include/sys/types.h:283:1:

// /usr/include/sys/types.h:283:24:
type Tclock_t = uint32

type clock_t = Tclock_t

//
// /usr/include/sys/types.h:288:1:

// /usr/include/sys/types.h:288:26:
type Tptrdiff_t = int64

type ptrdiff_t = Tptrdiff_t

//
// /usr/include/sys/types.h:293:1:

// /usr/include/sys/types.h:293:23:
type Tsize_t = uint64

type size_t = Tsize_t

//
// /usr/include/sys/types.h:299:1:

// /usr/include/sys/types.h:299:24:
type Tssize_t = int64

type ssize_t = Tssize_t

//
// /usr/include/sys/types.h:304:1:

// /usr/include/sys/types.h:304:23:
type Ttime_t = int64

type time_t = Ttime_t

//
// /usr/include/sys/types.h:309:1:

// /usr/include/sys/types.h:309:26:
type Tclockid_t = int32

type clockid_t = Tclockid_t

//
// /usr/include/sys/types.h:314:1:

// /usr/include/sys/types.h:314:24:
type Ttimer_t = int32

type timer_t = Ttimer_t

//
// /usr/include/sys/types.h:319:1:

// /usr/include/sys/types.h:319:27:
type Tsuseconds_t = int32

type suseconds_t = Tsuseconds_t

//
// /usr/include/sys/types.h:324:1:

// /usr/include/sys/types.h:324:26:
type Tuseconds_t = uint32

type useconds_t = Tuseconds_t

//
// /usr/include/sys/fd_set.h:46:1:

// /usr/include/sys/fd_set.h:46:20:
type t__fd_mask = uint32

//
// /usr/include/sys/fd_set.h:66:1:

// /usr/include/sys/fd_set.h:68:3:
type Tfd_set = struct {
	Ffds_bits [8]t__fd_mask
}

type fd_set = Tfd_set

//
// /usr/include/sys/types.h:333:1:

// /usr/include/sys/types.h:333:27:
type Tkauth_cred_t = uintptr

type kauth_cred_t = Tkauth_cred_t

//
// /usr/include/sys/types.h:335:1:

// /usr/include/sys/types.h:335:13:
type Tpri_t = int32

type pri_t = Tpri_t

//
// /usr/include/pthread_types.h:43:1:

// /usr/include/pthread_types.h:43:29:
type Tpthread_spin_t = uint8

type pthread_spin_t = Tpthread_spin_t

//
// /usr/include/pthread_types.h:48:1:

// /usr/include/pthread_types.h:48:24:
type t__pthread_spin_t = uint8

// /usr/include/pthread_types.h:61:1:
type Tpthread_queue_struct_t = struct {
	Fptqh_first uintptr
	Fptqh_last  uintptr
}

type pthread_queue_struct_t = Tpthread_queue_struct_t

//
// /usr/include/pthread_types.h:62:1:

// /usr/include/pthread_types.h:62:39:
type Tpthread_queue_t = struct {
	Fptqh_first uintptr
	Fptqh_last  uintptr
}

type pthread_queue_t = Tpthread_queue_t

// /usr/include/pthread_types.h:65:1:
type t__pthread_attr_st = struct {
	Fpta_magic   uint32
	Fpta_flags   int32
	Fpta_private uintptr
}

// /usr/include/pthread_types.h:66:1:
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

// /usr/include/pthread_types.h:67:1:
type t__pthread_mutexattr_st = struct {
	Fptma_magic   uint32
	Fptma_private uintptr
}

// /usr/include/pthread_types.h:68:1:
type t__pthread_cond_st = struct {
	Fptc_magic   uint32
	Fptc_lock    t__pthread_spin_t
	Fptc_waiters uintptr
	Fptc_spare   uintptr
	Fptc_mutex   uintptr
	Fptc_private uintptr
}

// /usr/include/pthread_types.h:69:1:
type t__pthread_condattr_st = struct {
	Fptca_magic   uint32
	Fptca_private uintptr
}

// /usr/include/pthread_types.h:71:1:
type t__pthread_rwlock_st = struct {
	Fptr_magic     uint32
	Fptr_interlock t__pthread_spin_t
	Fptr_rblocked  Tpthread_queue_t
	Fptr_wblocked  Tpthread_queue_t
	Fptr_nreaders  uint32
	Fptr_owner     Tpthread_t
	Fptr_private   uintptr
}

// /usr/include/pthread_types.h:72:1:
type t__pthread_rwlockattr_st = struct {
	Fptra_magic   uint32
	Fptra_private uintptr
}

// /usr/include/pthread_types.h:73:1:
type t__pthread_barrier_st = struct {
	Fptb_magic      uint32
	Fptb_lock       Tpthread_spin_t
	Fptb_waiters    Tpthread_queue_t
	Fptb_initcount  uint32
	Fptb_curcount   uint32
	Fptb_generation uint32
	Fptb_private    uintptr
}

// /usr/include/pthread_types.h:74:1:
type t__pthread_barrierattr_st = struct {
	Fptba_magic   uint32
	Fptba_private uintptr
}

//
// /usr/include/pthread_types.h:76:1:

// /usr/include/pthread_types.h:76:29:
type Tpthread_t = uintptr

type pthread_t = Tpthread_t

//
// /usr/include/pthread_types.h:77:1:

// /usr/include/pthread_types.h:77:34:
type Tpthread_attr_t = struct {
	Fpta_magic   uint32
	Fpta_flags   int32
	Fpta_private uintptr
}

type pthread_attr_t = Tpthread_attr_t

//
// /usr/include/pthread_types.h:78:1:

// /usr/include/pthread_types.h:78:35:
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

type pthread_mutex_t = Tpthread_mutex_t

//
// /usr/include/pthread_types.h:79:1:

// /usr/include/pthread_types.h:79:39:
type Tpthread_mutexattr_t = struct {
	Fptma_magic   uint32
	Fptma_private uintptr
}

type pthread_mutexattr_t = Tpthread_mutexattr_t

//
// /usr/include/pthread_types.h:80:1:

// /usr/include/pthread_types.h:80:34:
type Tpthread_cond_t = struct {
	Fptc_magic   uint32
	Fptc_lock    t__pthread_spin_t
	Fptc_waiters uintptr
	Fptc_spare   uintptr
	Fptc_mutex   uintptr
	Fptc_private uintptr
}

type pthread_cond_t = Tpthread_cond_t

//
// /usr/include/pthread_types.h:81:1:

// /usr/include/pthread_types.h:81:38:
type Tpthread_condattr_t = struct {
	Fptca_magic   uint32
	Fptca_private uintptr
}

type pthread_condattr_t = Tpthread_condattr_t

//
// /usr/include/pthread_types.h:82:1:

// /usr/include/pthread_types.h:82:34:
type Tpthread_once_t = struct {
	Fpto_mutex Tpthread_mutex_t
	Fpto_done  int32
}

type pthread_once_t = Tpthread_once_t

//
// /usr/include/pthread_types.h:82:1:

// /usr/include/pthread_types.h:82:34:
type t__pthread_once_st = Tpthread_once_t

//
// /usr/include/pthread_types.h:83:1:

// /usr/include/pthread_types.h:83:38:
type Tpthread_spinlock_t = struct {
	Fpts_magic uint32
	Fpts_spin  t__pthread_spin_t
	Fpts_flags int32
}

type pthread_spinlock_t = Tpthread_spinlock_t

//
// /usr/include/pthread_types.h:83:1:

// /usr/include/pthread_types.h:83:38:
type t__pthread_spinlock_st = Tpthread_spinlock_t

//
// /usr/include/pthread_types.h:84:1:

// /usr/include/pthread_types.h:84:36:
type Tpthread_rwlock_t = struct {
	Fptr_magic     uint32
	Fptr_interlock t__pthread_spin_t
	Fptr_rblocked  Tpthread_queue_t
	Fptr_wblocked  Tpthread_queue_t
	Fptr_nreaders  uint32
	Fptr_owner     Tpthread_t
	Fptr_private   uintptr
}

type pthread_rwlock_t = Tpthread_rwlock_t

//
// /usr/include/pthread_types.h:85:1:

// /usr/include/pthread_types.h:85:40:
type Tpthread_rwlockattr_t = struct {
	Fptra_magic   uint32
	Fptra_private uintptr
}

type pthread_rwlockattr_t = Tpthread_rwlockattr_t

//
// /usr/include/pthread_types.h:86:1:

// /usr/include/pthread_types.h:86:37:
type Tpthread_barrier_t = struct {
	Fptb_magic      uint32
	Fptb_lock       Tpthread_spin_t
	Fptb_waiters    Tpthread_queue_t
	Fptb_initcount  uint32
	Fptb_curcount   uint32
	Fptb_generation uint32
	Fptb_private    uintptr
}

type pthread_barrier_t = Tpthread_barrier_t

//
// /usr/include/pthread_types.h:87:1:

// /usr/include/pthread_types.h:87:41:
type Tpthread_barrierattr_t = struct {
	Fptba_magic   uint32
	Fptba_private uintptr
}

type pthread_barrierattr_t = Tpthread_barrierattr_t

//
// /usr/include/pthread_types.h:88:1:

// /usr/include/pthread_types.h:88:13:
type Tpthread_key_t = int32

type pthread_key_t = Tpthread_key_t

//
// /usr/include/stdlib.h:56:1:

// /usr/include/stdlib.h:59:3:
type Tdiv_t = struct {
	Fquot int32
	Frem  int32
}

type div_t = Tdiv_t

//
// /usr/include/stdlib.h:61:1:

// /usr/include/stdlib.h:64:3:
type Tldiv_t = struct {
	Fquot int64
	Frem  int64
}

type ldiv_t = Tldiv_t

//
// /usr/include/stdlib.h:68:1:

// /usr/include/stdlib.h:73:3:
type Tlldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = Tlldiv_t

//
// /usr/include/stdlib.h:77:1:

// /usr/include/stdlib.h:80:3:
type Tqdiv_t = struct {
	Fquot Tquad_t
	Frem  Tquad_t
}

type qdiv_t = Tqdiv_t

//
// /usr/include/stdio.h:69:1:

// /usr/include/stdio.h:72:3:
type Tfpos_t = struct {
	F_pos         t__off_t
	F_mbstate_in  t__mbstate_t
	F_mbstate_out t__mbstate_t
}

type fpos_t = Tfpos_t

//
// /usr/include/stdio.h:69:1:

// /usr/include/stdio.h:72:3:
type t__sfpos = Tfpos_t

// /usr/include/stdio.h:83:1:
type t__sbuf = struct {
	F_base uintptr
	F_size int32
}

//
// /usr/include/stdio.h:113:1:

// /usr/include/stdio.h:147:3:
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

type FILE = TFILE

//
// /usr/include/stdio.h:113:1:

// /usr/include/stdio.h:147:3:
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

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:76:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:76:16:
type Ti8 = int8

type i8 = Ti8

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:77:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:77:17:
type Tu8 = uint8

type u8 = Tu8

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:78:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:78:17:
type Ti16 = int16

type i16 = Ti16

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:79:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:79:17:
type Ti32 = int32

type i32 = Ti32

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:80:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:80:23:
type Ti64 = int64

type i64 = Ti64

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:81:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:81:18:
type Tu32 = uint32

type u32 = Tu32

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:82:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:82:18:
type Tu64 = uint64

type u64 = Tu64

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:83:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:83:15:
type Tf32 = float32

type f32 = Tf32

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:84:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:84:16:
type Tusize = uint64

type usize = Tusize

// sqlite3_vtab_in() was added in SQLite version 3.38 (2022-02-22)
// https://www.sqlite.org/changes.html#version_3_38_0

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:115:1:
type _VectorElementType = int32

const
// clang-format off

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:117:3:
_SQLITE_VEC_ELEMENT_TYPE_FLOAT32 = 223
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:118:3:
_SQLITE_VEC_ELEMENT_TYPE_BIT = 224
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:119:3:
_SQLITE_VEC_ELEMENT_TYPE_INT8 = 225

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:361:12:
func _l2_sqr_float(tls *libc.TLS, pVect1v uintptr, pVect2v uintptr, qty_ptr uintptr) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:362:46:
	var i, qty Tsize_t
	var pVect1, pVect2 uintptr
	var res, t Tf32
	_, _, _, _, _, _ = i, pVect1, pVect2, qty, res, t
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:363:7:
	pVect1 = pVect1v
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:364:7:
	pVect2 = pVect2v
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:365:10:
	qty = **(**Tsize_t)(__ccgo_up(qty_ptr))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:367:7:
	res = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:368:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:368:15:
	i = uint64(0)
	for {
		if !(i < qty) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:369:9:
		t = **(**Tf32)(__ccgo_up(pVect1)) - **(**Tf32)(__ccgo_up(pVect2))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:370:5:
		pVect1 += 4
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:371:5:
		pVect2 += 4
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:372:5:
		res = res + Tf32(t*t)
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:374:3:
	return float32(libc.Xsqrt(tls, float64(res)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:377:12:
func _l2_sqr_int8(tls *libc.TLS, pA uintptr, pB uintptr, pD uintptr) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:377:72:
	var a, b uintptr
	var d, i Tsize_t
	var res, t Tf32
	_, _, _, _, _, _ = a, b, d, i, res, t
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:378:6:
	a = pA
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:379:6:
	b = pB
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:380:10:
	d = **(**Tsize_t)(__ccgo_up(pD))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:382:7:
	res = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:383:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:383:15:
	i = uint64(0)
	for {
		if !(i < d) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:384:9:
		t = float32(int32(**(**Ti8)(__ccgo_up(a))) - int32(**(**Ti8)(__ccgo_up(b))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:385:5:
		a = a + 1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:386:5:
		b = b + 1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:387:5:
		res = res + Tf32(t*t)
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:389:3:
	return float32(libc.Xsqrt(tls, float64(res)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:392:12:
func _distance_l2_sqr_float(tls *libc.TLS, a uintptr, b uintptr, d uintptr) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:392:79:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:403:3:
	return _l2_sqr_float(tls, a, b, d)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:406:12:
func _distance_l2_sqr_int8(tls *libc.TLS, a uintptr, b uintptr, d uintptr) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:406:78:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:412:3:
	return _l2_sqr_int8(tls, a, b, d)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:415:12:
func _l1_int8(tls *libc.TLS, pA uintptr, pB uintptr, pD uintptr) (r Ti32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:415:68:
	var a, b uintptr
	var d, i Tsize_t
	var res Ti32
	_, _, _, _, _ = a, b, d, i, res
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:416:6:
	a = pA
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:417:6:
	b = pB
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:418:10:
	d = **(**Tsize_t)(__ccgo_up(pD))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:420:7:
	res = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:421:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:421:15:
	i = uint64(0)
	for {
		if !(i < d) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:422:5:
		res = res + libc.Xabs(tls, int32(**(**Ti8)(__ccgo_up(a)))-int32(**(**Ti8)(__ccgo_up(b))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:423:5:
		a = a + 1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:424:5:
		b = b + 1
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:427:3:
	return res
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:430:12:
func _distance_l1_int8(tls *libc.TLS, a uintptr, b uintptr, d uintptr) (r Ti32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:430:74:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:436:3:
	return _l1_int8(tls, a, b, d)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:439:15:
func _l1_f32(tls *libc.TLS, pA uintptr, pB uintptr, pD uintptr) (r float64) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:439:70:
	var a, b uintptr
	var d, i Tsize_t
	var res float64
	_, _, _, _, _ = a, b, d, i, res
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:440:7:
	a = pA
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:441:7:
	b = pB
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:442:10:
	d = **(**Tsize_t)(__ccgo_up(pD))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:444:10:
	res = libc.Float64FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:445:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:445:15:
	i = uint64(0)
	for {
		if !(i < d) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:446:5:
		res = res + libc.Xfabs(tls, float64(**(**Tf32)(__ccgo_up(a)))-float64(**(**Tf32)(__ccgo_up(b))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:447:5:
		a += 4
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:448:5:
		b += 4
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:451:3:
	return res
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:454:15:
func _distance_l1_f32(tls *libc.TLS, a uintptr, b uintptr, d uintptr) (r float64) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:454:76:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:460:3:
	return _l1_f32(tls, a, b, d)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:463:12:
func _distance_cosine_float(tls *libc.TLS, pVect1v uintptr, pVect2v uintptr, qty_ptr uintptr) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:464:55:
	var aMag, bMag, dot Tf32
	var i, qty Tsize_t
	var pVect1, pVect2 uintptr
	_, _, _, _, _, _, _ = aMag, bMag, dot, i, pVect1, pVect2, qty
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:465:7:
	pVect1 = pVect1v
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:466:7:
	pVect2 = pVect2v
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:467:10:
	qty = **(**Tsize_t)(__ccgo_up(qty_ptr))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:469:7:
	dot = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:470:7:
	aMag = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:471:7:
	bMag = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:472:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:472:15:
	i = uint64(0)
	for {
		if !(i < qty) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:473:5:
		dot = dot + Tf32(**(**Tf32)(__ccgo_up(pVect1))***(**Tf32)(__ccgo_up(pVect2)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:474:5:
		aMag = aMag + Tf32(**(**Tf32)(__ccgo_up(pVect1))***(**Tf32)(__ccgo_up(pVect1)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:475:5:
		bMag = bMag + Tf32(**(**Tf32)(__ccgo_up(pVect2))***(**Tf32)(__ccgo_up(pVect2)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:476:5:
		pVect1 += 4
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:477:5:
		pVect2 += 4
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:479:3:
	return float32(libc.Float64FromInt32(1) - float64(dot)/float64(libc.Xsqrt(tls, float64(aMag))*libc.Xsqrt(tls, float64(bMag))))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:481:12:
func _distance_cosine_int8(tls *libc.TLS, pA uintptr, pB uintptr, pD uintptr) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:482:49:
	var a, b uintptr
	var aMag, bMag, dot Tf32
	var d, i Tsize_t
	_, _, _, _, _, _, _ = a, aMag, b, bMag, d, dot, i
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:483:6:
	a = pA
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:484:6:
	b = pB
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:485:10:
	d = **(**Tsize_t)(__ccgo_up(pD))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:487:7:
	dot = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:488:7:
	aMag = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:489:7:
	bMag = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:490:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:490:15:
	i = uint64(0)
	for {
		if !(i < d) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:491:5:
		dot = dot + Tf32(int32(**(**Ti8)(__ccgo_up(a)))*int32(**(**Ti8)(__ccgo_up(b))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:492:5:
		aMag = aMag + Tf32(int32(**(**Ti8)(__ccgo_up(a)))*int32(**(**Ti8)(__ccgo_up(a))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:493:5:
		bMag = bMag + Tf32(int32(**(**Ti8)(__ccgo_up(b)))*int32(**(**Ti8)(__ccgo_up(b))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:494:5:
		a = a + 1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:495:5:
		b = b + 1
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:497:3:
	return float32(libc.Float64FromInt32(1) - float64(dot)/float64(libc.Xsqrt(tls, float64(aMag))*libc.Xsqrt(tls, float64(bMag))))
}

// C documentation
//
//	// https://github.com/facebookresearch/faiss/blob/77e2e79cd0a680adc343b9840dd865da724c579e/faiss/utils/hamming_distance/common.h#L34
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:501:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:501:11:
var _hamdist_table = [256]Tu8{
	1:   uint8(1),
	2:   uint8(1),
	3:   uint8(2),
	4:   uint8(1),
	5:   uint8(2),
	6:   uint8(2),
	7:   uint8(3),
	8:   uint8(1),
	9:   uint8(2),
	10:  uint8(2),
	11:  uint8(3),
	12:  uint8(2),
	13:  uint8(3),
	14:  uint8(3),
	15:  uint8(4),
	16:  uint8(1),
	17:  uint8(2),
	18:  uint8(2),
	19:  uint8(3),
	20:  uint8(2),
	21:  uint8(3),
	22:  uint8(3),
	23:  uint8(4),
	24:  uint8(2),
	25:  uint8(3),
	26:  uint8(3),
	27:  uint8(4),
	28:  uint8(3),
	29:  uint8(4),
	30:  uint8(4),
	31:  uint8(5),
	32:  uint8(1),
	33:  uint8(2),
	34:  uint8(2),
	35:  uint8(3),
	36:  uint8(2),
	37:  uint8(3),
	38:  uint8(3),
	39:  uint8(4),
	40:  uint8(2),
	41:  uint8(3),
	42:  uint8(3),
	43:  uint8(4),
	44:  uint8(3),
	45:  uint8(4),
	46:  uint8(4),
	47:  uint8(5),
	48:  uint8(2),
	49:  uint8(3),
	50:  uint8(3),
	51:  uint8(4),
	52:  uint8(3),
	53:  uint8(4),
	54:  uint8(4),
	55:  uint8(5),
	56:  uint8(3),
	57:  uint8(4),
	58:  uint8(4),
	59:  uint8(5),
	60:  uint8(4),
	61:  uint8(5),
	62:  uint8(5),
	63:  uint8(6),
	64:  uint8(1),
	65:  uint8(2),
	66:  uint8(2),
	67:  uint8(3),
	68:  uint8(2),
	69:  uint8(3),
	70:  uint8(3),
	71:  uint8(4),
	72:  uint8(2),
	73:  uint8(3),
	74:  uint8(3),
	75:  uint8(4),
	76:  uint8(3),
	77:  uint8(4),
	78:  uint8(4),
	79:  uint8(5),
	80:  uint8(2),
	81:  uint8(3),
	82:  uint8(3),
	83:  uint8(4),
	84:  uint8(3),
	85:  uint8(4),
	86:  uint8(4),
	87:  uint8(5),
	88:  uint8(3),
	89:  uint8(4),
	90:  uint8(4),
	91:  uint8(5),
	92:  uint8(4),
	93:  uint8(5),
	94:  uint8(5),
	95:  uint8(6),
	96:  uint8(2),
	97:  uint8(3),
	98:  uint8(3),
	99:  uint8(4),
	100: uint8(3),
	101: uint8(4),
	102: uint8(4),
	103: uint8(5),
	104: uint8(3),
	105: uint8(4),
	106: uint8(4),
	107: uint8(5),
	108: uint8(4),
	109: uint8(5),
	110: uint8(5),
	111: uint8(6),
	112: uint8(3),
	113: uint8(4),
	114: uint8(4),
	115: uint8(5),
	116: uint8(4),
	117: uint8(5),
	118: uint8(5),
	119: uint8(6),
	120: uint8(4),
	121: uint8(5),
	122: uint8(5),
	123: uint8(6),
	124: uint8(5),
	125: uint8(6),
	126: uint8(6),
	127: uint8(7),
	128: uint8(1),
	129: uint8(2),
	130: uint8(2),
	131: uint8(3),
	132: uint8(2),
	133: uint8(3),
	134: uint8(3),
	135: uint8(4),
	136: uint8(2),
	137: uint8(3),
	138: uint8(3),
	139: uint8(4),
	140: uint8(3),
	141: uint8(4),
	142: uint8(4),
	143: uint8(5),
	144: uint8(2),
	145: uint8(3),
	146: uint8(3),
	147: uint8(4),
	148: uint8(3),
	149: uint8(4),
	150: uint8(4),
	151: uint8(5),
	152: uint8(3),
	153: uint8(4),
	154: uint8(4),
	155: uint8(5),
	156: uint8(4),
	157: uint8(5),
	158: uint8(5),
	159: uint8(6),
	160: uint8(2),
	161: uint8(3),
	162: uint8(3),
	163: uint8(4),
	164: uint8(3),
	165: uint8(4),
	166: uint8(4),
	167: uint8(5),
	168: uint8(3),
	169: uint8(4),
	170: uint8(4),
	171: uint8(5),
	172: uint8(4),
	173: uint8(5),
	174: uint8(5),
	175: uint8(6),
	176: uint8(3),
	177: uint8(4),
	178: uint8(4),
	179: uint8(5),
	180: uint8(4),
	181: uint8(5),
	182: uint8(5),
	183: uint8(6),
	184: uint8(4),
	185: uint8(5),
	186: uint8(5),
	187: uint8(6),
	188: uint8(5),
	189: uint8(6),
	190: uint8(6),
	191: uint8(7),
	192: uint8(2),
	193: uint8(3),
	194: uint8(3),
	195: uint8(4),
	196: uint8(3),
	197: uint8(4),
	198: uint8(4),
	199: uint8(5),
	200: uint8(3),
	201: uint8(4),
	202: uint8(4),
	203: uint8(5),
	204: uint8(4),
	205: uint8(5),
	206: uint8(5),
	207: uint8(6),
	208: uint8(3),
	209: uint8(4),
	210: uint8(4),
	211: uint8(5),
	212: uint8(4),
	213: uint8(5),
	214: uint8(5),
	215: uint8(6),
	216: uint8(4),
	217: uint8(5),
	218: uint8(5),
	219: uint8(6),
	220: uint8(5),
	221: uint8(6),
	222: uint8(6),
	223: uint8(7),
	224: uint8(3),
	225: uint8(4),
	226: uint8(4),
	227: uint8(5),
	228: uint8(4),
	229: uint8(5),
	230: uint8(5),
	231: uint8(6),
	232: uint8(4),
	233: uint8(5),
	234: uint8(5),
	235: uint8(6),
	236: uint8(5),
	237: uint8(6),
	238: uint8(6),
	239: uint8(7),
	240: uint8(4),
	241: uint8(5),
	242: uint8(5),
	243: uint8(6),
	244: uint8(5),
	245: uint8(6),
	246: uint8(6),
	247: uint8(7),
	248: uint8(5),
	249: uint8(6),
	250: uint8(6),
	251: uint8(7),
	252: uint8(6),
	253: uint8(7),
	254: uint8(7),
	255: uint8(8),
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:514:12:
func _distance_hamming_u8(tls *libc.TLS, a uintptr, b uintptr, n Tsize_t) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:514:56:
	var i uint64
	var same int32
	_, _ = i, same
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:515:7:
	same = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:516:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:516:22:
	i = uint64(0)
	for {
		if !(i < n) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:517:5:
		same = same + libc.Int32FromUint8(_hamdist_table[libc.Int32FromUint8(**(**Tu8)(__ccgo_up(a + uintptr(i))))^libc.Int32FromUint8(**(**Tu8)(__ccgo_up(b + uintptr(i))))])
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:519:3:
	return float32(same)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:540:12:
func _distance_hamming_u64(tls *libc.TLS, a uintptr, b uintptr, n Tsize_t) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:540:59:
	var i uint64
	var same int32
	_, _ = i, same
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:541:7:
	same = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:542:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:542:22:
	i = uint64(0)
	for {
		if !(i < n) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:543:5:
		same = same + libc.X__builtin_popcountl(tls, **(**Tu64)(__ccgo_up(a + uintptr(i)*8))^**(**Tu64)(__ccgo_up(b + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:545:3:
	return float32(same)
}

// C documentation
//
//	/**
//	 * @brief Calculate the hamming distance between two bitvectors.
//	 *
//	 * @param a - first bitvector, MUST have d dimensions
//	 * @param b - second bitvector, MUST have d dimensions
//	 * @param d - pointer to size_t, MUST be divisible by CHAR_BIT
//	 * @return f32
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:556:12:
func _distance_hamming(tls *libc.TLS, a uintptr, b uintptr, d uintptr) (r Tf32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:556:74:
	var dimensions Tsize_t
	_ = dimensions
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:557:10:
	dimensions = **(**Tsize_t)(__ccgo_up(d))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:559:3:
	if dimensions%uint64(64) == uint64(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:560:5:
		return _distance_hamming_u64(tls, a, b, dimensions/uint64(8)/uint64(m_CHAR_BIT))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:562:3:
	return _distance_hamming_u8(tls, a, b, dimensions/uint64(m_CHAR_BIT))
}

// C documentation
//
//	// from SQLite source:
//	// https://github.com/sqlite/sqlite/blob/a509a90958ddb234d1785ed7801880ccb18b497e/src/json.c#L153
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:579:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:579:19:
var _vecJsonIsSpaceX = [256]int8{
	9:  int8(1),
	10: int8(1),
	13: int8(1),
	32: int8(1),
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:597:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:597:14:
type Tvector_cleanup = uintptr

type vector_cleanup = Tvector_cleanup

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:599:6:
func Xvector_cleanup_noop(tls *libc.TLS, _1 uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:599:35:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:599:37:
	_ = _1
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:603:6:
func Xvtab_set_error(tls *libc.TLS, pVTab uintptr, zFormat uintptr, va uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:603:68:
	var args Tva_list
	_ = args
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:605:3:
	libsqlite3.Xsqlite3_free(tls, (*Tsqlite3_vtab)(unsafe.Pointer(pVTab)).FzErrMsg)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:606:3:
	args = va
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:607:3:
	(*Tsqlite3_vtab)(unsafe.Pointer(pVTab)).FzErrMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, args)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:608:3:
	_ = args
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:610:1:
type TArray = struct {
	Felement_size Tsize_t
	Flength       Tsize_t
	Fcapacity     Tsize_t
	Fz            uintptr
}

type Array = TArray

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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:626:5:
func Xarray_init(tls *libc.TLS, array uintptr, element_size Tsize_t, init_capacity Tsize_t) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:626:80:
	var sz int32
	var z uintptr
	_, _ = sz, z
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:627:7:
	sz = libc.Int32FromUint64(element_size * init_capacity)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:628:8:
	z = libsqlite3.Xsqlite3_malloc(tls, sz)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:629:3:
	if !(z != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:630:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:632:3:
	libc.Xmemset(tls, z, 0, libc.Uint64FromInt32(sz))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:634:3:
	(*TArray)(unsafe.Pointer(array)).Felement_size = element_size
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:635:3:
	(*TArray)(unsafe.Pointer(array)).Flength = uint64(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:636:3:
	(*TArray)(unsafe.Pointer(array)).Fcapacity = init_capacity
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:637:3:
	(*TArray)(unsafe.Pointer(array)).Fz = z
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:638:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:641:5:
func Xarray_append(tls *libc.TLS, array uintptr, element uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:641:60:
	var new_capacity Tsize_t
	var z uintptr
	_, _ = new_capacity, z
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:642:3:
	if (*TArray)(unsafe.Pointer(array)).Flength == (*TArray)(unsafe.Pointer(array)).Fcapacity {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:643:12:
		new_capacity = (*TArray)(unsafe.Pointer(array)).Fcapacity*uint64(2) + uint64(100)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:644:10:
		z = libsqlite3.Xsqlite3_realloc64(tls, (*TArray)(unsafe.Pointer(array)).Fz, (*TArray)(unsafe.Pointer(array)).Felement_size*new_capacity)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:645:5:
		if z != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:646:7:
			(*TArray)(unsafe.Pointer(array)).Fcapacity = new_capacity
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:647:7:
			(*TArray)(unsafe.Pointer(array)).Fz = z
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:649:7:
			return int32(m_SQLITE_NOMEM)
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:652:3:
	libc.Xmemcpy(tls, (*TArray)(unsafe.Pointer(array)).Fz+uintptr((*TArray)(unsafe.Pointer(array)).Flength*(*TArray)(unsafe.Pointer(array)).Felement_size), element, (*TArray)(unsafe.Pointer(array)).Felement_size)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:654:3:
	(*TArray)(unsafe.Pointer(array)).Flength = (*TArray)(unsafe.Pointer(array)).Flength + 1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:655:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:658:6:
func Xarray_cleanup(tls *libc.TLS, array uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:658:41:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:659:3:
	if !(array != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:660:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:661:3:
	(*TArray)(unsafe.Pointer(array)).Felement_size = uint64(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:662:3:
	(*TArray)(unsafe.Pointer(array)).Flength = uint64(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:663:3:
	(*TArray)(unsafe.Pointer(array)).Fcapacity = uint64(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:664:3:
	libsqlite3.Xsqlite3_free(tls, (*TArray)(unsafe.Pointer(array)).Fz)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:665:3:
	(*TArray)(unsafe.Pointer(array)).Fz = libc.UintptrFromInt32(0)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:668:6:
func Xvector_subtype_name(tls *libc.TLS, subtype int32) (r uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:668:40:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:669:3:
	switch subtype {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:670:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:671:5:
		return __ccgo_ts
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:672:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:673:5:
		return __ccgo_ts + 8
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:674:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:675:5:
		return __ccgo_ts + 13
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:677:3:
	return __ccgo_ts + 17
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:679:6:
func Xtype_name(tls *libc.TLS, type1 int32) (r uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:679:27:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:680:3:
	switch type1 {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:681:3:
	case int32(m_SQLITE_INTEGER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:682:5:
		return __ccgo_ts + 18
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:683:3:
		fallthrough
	case int32(m_SQLITE_BLOB):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:684:5:
		return __ccgo_ts + 26
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:685:3:
		fallthrough
	case int32(m_SQLITE_TEXT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:686:5:
		return __ccgo_ts + 31
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:687:3:
		fallthrough
	case int32(m_SQLITE_FLOAT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:688:5:
		return __ccgo_ts + 36
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:689:3:
		fallthrough
	case int32(m_SQLITE_NULL):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:690:5:
		return __ccgo_ts + 42
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:692:3:
	return __ccgo_ts + 17
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:695:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:695:14:
type Tfvec_cleanup = uintptr

type fvec_cleanup = Tfvec_cleanup

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:697:6:
func Xfvec_cleanup_noop(tls *libc.TLS, _1 uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:697:33:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:697:35:
	_ = _1
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:699:12:
func _fvec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:701:42:
	var blob, buf, ptr, source uintptr
	var bytes, i, offset, rc, source_len, value_type int32
	var result float64
	var _ /* endptr at bp+32 */ uintptr
	var _ /* res at bp+40 */ Tf32
	var _ /* x at bp+0 */ TArray
	_, _, _, _, _, _, _, _, _, _, _ = blob, buf, bytes, i, offset, ptr, rc, result, source, source_len, value_type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:702:7:
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:704:3:
	if value_type == int32(m_SQLITE_BLOB) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:705:16:
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:706:9:
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:707:5:
		if bytes == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:708:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:709:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:711:5:
		if libc.Uint64FromInt32(bytes)%uint64(4) != uint64(0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:712:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+86, libc.VaList(bp+56, uint64(4), bytes))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:715:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:717:9:
		buf = libsqlite3.Xsqlite3_malloc(tls, bytes)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:718:5:
		if !(buf != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:719:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+156, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:720:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:722:5:
		libc.Xmemcpy(tls, buf, blob, libc.Uint64FromInt32(bytes))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:723:5:
		**(**uintptr)(__ccgo_up(vector)) = buf
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:724:5:
		**(**Tsize_t)(__ccgo_up(dimensions)) = libc.Uint64FromInt32(bytes) / uint64(4)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:725:5:
		**(**Tfvec_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(libsqlite3.Xsqlite3_free)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:726:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:729:3:
	if value_type == int32(m_SQLITE_TEXT) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:730:16:
		source = libsqlite3.Xsqlite3_value_text(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:731:9:
		source_len = libsqlite3.Xsqlite3_value_bytes(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:732:5:
		if source_len == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:733:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:734:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:736:9:
		i = 0
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:739:9:
		rc = Xarray_init(tls, bp, uint64(4), uint64(libc.Xceil(tls, float64(source_len)/float64(2))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:740:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:741:7:
			return rc
		}
		// advance leading whitespace to first '['
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:745:5:
		for i < source_len {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:746:7:
			if _vecJsonIsSpaceX[libc.Uint8FromInt8(**(**int8)(__ccgo_up(source + uintptr(i))))] != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:747:9:
				i = i + 1
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:748:9:
				continue
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:750:7:
			if int32(**(**int8)(__ccgo_up(source + uintptr(i)))) == int32('[') {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:751:9:
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:754:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:756:7:
			Xarray_cleanup(tls, bp)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:757:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:759:5:
		if int32(**(**int8)(__ccgo_up(source + uintptr(i)))) != int32('[') {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:760:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:762:7:
			Xarray_cleanup(tls, bp)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:763:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:765:9:
		offset = i + int32(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:767:5:
		for offset < source_len {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:768:12:
			ptr = source + uintptr(offset)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:771:7:
			**(**int32)(__ccgo_up(libc.X__errno(tls))) = 0
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:772:14:
			result = libc.Xstrtod(tls, ptr, bp+32)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:773:7:
			if **(**int32)(__ccgo_up(libc.X__errno(tls))) != 0 && result == libc.Float64FromInt32(0) || **(**int32)(__ccgo_up(libc.X__errno(tls))) == int32(m_ERANGE) && (result == libc.X__builtin_huge_val(tls) || result == -libc.X__builtin_huge_val(tls)) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:777:9:
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:778:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:779:9:
				return int32(m_SQLITE_ERROR)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:782:7:
			if **(**uintptr)(__ccgo_up(bp + 32)) == ptr {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:783:9:
				if int32(**(**int8)(__ccgo_up(ptr))) != int32(']') {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:784:11:
					libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:785:11:
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:786:11:
					return int32(m_SQLITE_ERROR)
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:788:9:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:791:11:
			**(**Tf32)(__ccgo_up(bp + 40)) = float32(result)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:792:7:
			Xarray_append(tls, bp, bp+40)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:794:7:
			offset = int32(int64(offset) + (int64(**(**uintptr)(__ccgo_up(bp + 32))) - int64(ptr)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:795:7:
			for offset < source_len {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:796:9:
				if _vecJsonIsSpaceX[libc.Uint8FromInt8(**(**int8)(__ccgo_up(source + uintptr(offset))))] != 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:797:11:
					offset = offset + 1
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:798:11:
					continue
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:800:9:
				if int32(**(**int8)(__ccgo_up(source + uintptr(offset)))) == int32(',') {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:801:11:
					offset = offset + 1
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:802:11:
					continue
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:804:9:
				if int32(**(**int8)(__ccgo_up(source + uintptr(offset)))) == int32(']') {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:805:11:
					goto done
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:806:9:
				break
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:810:3:
		goto done
	done:
		;
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:812:5:
		if (**(**TArray)(__ccgo_up(bp))).Flength > uint64(0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:813:7:
			**(**uintptr)(__ccgo_up(vector)) = (**(**TArray)(__ccgo_up(bp))).Fz
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:814:7:
			**(**Tsize_t)(__ccgo_up(dimensions)) = (**(**TArray)(__ccgo_up(bp))).Flength
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:815:7:
			**(**Tfvec_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(libsqlite3.Xsqlite3_free)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:816:7:
			return m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:818:5:
		libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:819:5:
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:820:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:823:3:
	**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+245, libc.VaList(bp+56, Xtype_name(tls, value_type)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:826:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:829:12:
func _bitvec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:831:44:
	var blob uintptr
	var bytes, value_type int32
	_, _, _ = blob, bytes, value_type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:832:7:
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:833:3:
	if value_type == int32(m_SQLITE_BLOB) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:834:16:
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:835:9:
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:836:5:
		if bytes == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:837:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:838:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:840:5:
		**(**uintptr)(__ccgo_up(vector)) = blob
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:841:5:
		**(**Tsize_t)(__ccgo_up(dimensions)) = libc.Uint64FromInt32(bytes * int32(m_CHAR_BIT))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:842:5:
		**(**Tvector_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(Xvector_cleanup_noop)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:843:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:845:3:
	**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+313, 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:846:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:849:12:
func _int8_vec_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, __ccgo_fp_cleanup uintptr, pzErr uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:851:46:
	var blob, ptr, source uintptr
	var bytes, i, offset, rc, source_len, value_type int32
	var result int64
	var _ /* endptr at bp+32 */ uintptr
	var _ /* res at bp+40 */ Ti8
	var _ /* x at bp+0 */ TArray
	_, _, _, _, _, _, _, _, _, _ = blob, bytes, i, offset, ptr, rc, result, source, source_len, value_type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:852:7:
	value_type = libsqlite3.Xsqlite3_value_type(tls, value)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:853:3:
	if value_type == int32(m_SQLITE_BLOB) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:854:16:
		blob = libsqlite3.Xsqlite3_value_blob(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:855:9:
		bytes = libsqlite3.Xsqlite3_value_bytes(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:856:5:
		if bytes == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:857:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:858:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:860:5:
		**(**uintptr)(__ccgo_up(vector)) = blob
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:861:5:
		**(**Tsize_t)(__ccgo_up(dimensions)) = libc.Uint64FromInt32(bytes)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:862:5:
		**(**Tvector_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(Xvector_cleanup_noop)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:863:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:866:3:
	if value_type == int32(m_SQLITE_TEXT) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:867:16:
		source = libsqlite3.Xsqlite3_value_text(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:868:9:
		source_len = libsqlite3.Xsqlite3_value_bytes(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:869:9:
		i = 0
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:871:5:
		if source_len == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:872:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:873:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:877:9:
		rc = Xarray_init(tls, bp, uint64(1), uint64(libc.Xceil(tls, float64(source_len)/float64(2))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:878:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:879:7:
			return rc
		}
		// advance leading whitespace to first '['
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:883:5:
		for i < source_len {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:884:7:
			if _vecJsonIsSpaceX[libc.Uint8FromInt8(**(**int8)(__ccgo_up(source + uintptr(i))))] != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:885:9:
				i = i + 1
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:886:9:
				continue
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:888:7:
			if int32(**(**int8)(__ccgo_up(source + uintptr(i)))) == int32('[') {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:889:9:
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:892:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:894:7:
			Xarray_cleanup(tls, bp)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:895:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:897:5:
		if int32(**(**int8)(__ccgo_up(source + uintptr(i)))) != int32('[') {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:898:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+170, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:900:7:
			Xarray_cleanup(tls, bp)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:901:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:903:9:
		offset = i + int32(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:905:5:
		for offset < source_len {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:906:12:
			ptr = source + uintptr(offset)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:909:7:
			**(**int32)(__ccgo_up(libc.X__errno(tls))) = 0
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:910:12:
			result = libc.Xstrtol(tls, ptr, bp+32, int32(10))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:911:7:
			if **(**int32)(__ccgo_up(libc.X__errno(tls))) != 0 && result == 0 || **(**int32)(__ccgo_up(libc.X__errno(tls))) == int32(m_ERANGE) && (result == int64(0x7fffffffffffffff) || result == -libc.Int64FromInt64(0x7fffffffffffffff)-libc.Int64FromInt32(1)) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:913:9:
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:914:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:915:9:
				return int32(m_SQLITE_ERROR)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:918:7:
			if **(**uintptr)(__ccgo_up(bp + 32)) == ptr {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:919:9:
				if int32(**(**int8)(__ccgo_up(ptr))) != int32(']') {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:920:11:
					libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:921:11:
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+226, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:922:11:
					return int32(m_SQLITE_ERROR)
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:924:9:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:927:7:
			if result < int64(-libc.Int32FromInt32(m___INT8_MAX__)-libc.Int32FromInt32(1)) || result > int64(m___INT8_MAX__) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:928:9:
				libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:929:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+341, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:931:9:
				return int32(m_SQLITE_ERROR)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:934:10:
			**(**Ti8)(__ccgo_up(bp + 40)) = int8(result)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:935:7:
			Xarray_append(tls, bp, bp+40)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:937:7:
			offset = int32(int64(offset) + (int64(**(**uintptr)(__ccgo_up(bp + 32))) - int64(ptr)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:938:7:
			for offset < source_len {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:939:9:
				if _vecJsonIsSpaceX[libc.Uint8FromInt8(**(**int8)(__ccgo_up(source + uintptr(offset))))] != 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:940:11:
					offset = offset + 1
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:941:11:
					continue
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:943:9:
				if int32(**(**int8)(__ccgo_up(source + uintptr(offset)))) == int32(',') {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:944:11:
					offset = offset + 1
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:945:11:
					continue
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:947:9:
				if int32(**(**int8)(__ccgo_up(source + uintptr(offset)))) == int32(']') {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:948:11:
					goto done
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:949:9:
				break
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:953:3:
		goto done
	done:
		;
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:955:5:
		if (**(**TArray)(__ccgo_up(bp))).Flength > uint64(0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:956:7:
			**(**uintptr)(__ccgo_up(vector)) = (**(**TArray)(__ccgo_up(bp))).Fz
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:957:7:
			**(**Tsize_t)(__ccgo_up(dimensions)) = (**(**TArray)(__ccgo_up(bp))).Flength
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:958:7:
			**(**Tvector_cleanup)(__ccgo_up(__ccgo_fp_cleanup)) = __ccgo_fp(libsqlite3.Xsqlite3_free)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:959:7:
			return m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:961:5:
		libsqlite3.Xsqlite3_free(tls, (**(**TArray)(__ccgo_up(bp))).Fz)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:962:5:
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+47, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:963:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:966:3:
	**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+389, 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:967:3:
	return int32(m_SQLITE_ERROR)
}

type t__ccgo_fp__Xvector_from_value_4 = func(*libc.TLS, uintptr)

// C documentation
//
//	/**
//	 * @brief Extract a vector from a sqlite3_value. Can be a float32, int8, or bit
//	 * vector.
//	 *
//	 * @param value: the sqlite3_value to read from.
//	 * @param vector: Output pointer to vector data.
//	 * @param dimensions: Output number of dimensions
//	 * @param dimensions: Output vector element type
//	 * @param cleanup
//	 * @param pzErrorMessage
//	 * @return int SQLITE_OK on success, error code otherwise
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:982:5:
func Xvector_from_value(tls *libc.TLS, value uintptr, vector uintptr, dimensions uintptr, element_type uintptr, __ccgo_fp_cleanup uintptr, pzErrorMessage uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:984:71:
	var rc, rc1, rc2, subtype int32
	_, _, _, _ = rc, rc1, rc2, subtype
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:985:7:
	subtype = libc.Int32FromUint32(libsqlite3.Xsqlite3_value_subtype(tls, value))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:986:3:
	if !(subtype != 0) || subtype == int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32) || subtype == int32(m_JSON_SUBTYPE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:988:9:
		rc = _fvec_from_value(tls, value, vector, dimensions, __ccgo_fp_cleanup, pzErrorMessage)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:990:5:
		if rc == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:991:7:
			**(**_VectorElementType)(__ccgo_up(element_type)) = int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:993:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:996:3:
	if subtype == int32(_SQLITE_VEC_ELEMENT_TYPE_BIT) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:997:9:
		rc1 = _bitvec_from_value(tls, value, vector, dimensions, __ccgo_fp_cleanup, pzErrorMessage)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:999:5:
		if rc1 == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1000:7:
			**(**_VectorElementType)(__ccgo_up(element_type)) = int32(_SQLITE_VEC_ELEMENT_TYPE_BIT)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1002:5:
		return rc1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1004:3:
	if subtype == int32(_SQLITE_VEC_ELEMENT_TYPE_INT8) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1005:9:
		rc2 = _int8_vec_from_value(tls, value, vector, dimensions, __ccgo_fp_cleanup, pzErrorMessage)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1007:5:
		if rc2 == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1008:7:
			**(**_VectorElementType)(__ccgo_up(element_type)) = int32(_SQLITE_VEC_ELEMENT_TYPE_INT8)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1010:5:
		return rc2
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1012:3:
	**(**uintptr)(__ccgo_up(pzErrorMessage)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+419, libc.VaList(bp+8, subtype))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1013:3:
	return int32(m_SQLITE_ERROR)
}

type t__ccgo_fp__Xensure_vector_match_6 = func(*libc.TLS, uintptr)

type t__ccgo_fp__Xensure_vector_match_7 = func(*libc.TLS, uintptr)

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1016:5:
func Xensure_vector_match(tls *libc.TLS, aValue uintptr, bValue uintptr, a uintptr, b uintptr, element_type uintptr, dimensions uintptr, __ccgo_fp_outACleanup uintptr, __ccgo_fp_outBCleanup uintptr, outError uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1019:71:
	var rc int32
	var _ /* aCleanup at bp+32 */ Tvector_cleanup
	var _ /* aDims at bp+8 */ Tsize_t
	var _ /* aType at bp+0 */ _VectorElementType
	var _ /* bCleanup at bp+40 */ Tvector_cleanup
	var _ /* bDims at bp+16 */ Tsize_t
	var _ /* bType at bp+4 */ _VectorElementType
	var _ /* error at bp+24 */ uintptr
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1023:8:
	**(**uintptr)(__ccgo_up(bp + 24)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1026:3:
	rc = Xvector_from_value(tls, aValue, a, bp+8, bp, bp+32, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1027:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1028:5:
		**(**uintptr)(__ccgo_up(outError)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+439, libc.VaList(bp+56, **(**uintptr)(__ccgo_up(bp + 24))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1029:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1030:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1033:3:
	rc = Xvector_from_value(tls, bValue, b, bp+16, bp+4, bp+40, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1034:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1035:5:
		**(**uintptr)(__ccgo_up(outError)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+468, libc.VaList(bp+56, **(**uintptr)(__ccgo_up(bp + 24))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1036:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1037:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(a)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1038:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1041:3:
	if **(**_VectorElementType)(__ccgo_up(bp)) != **(**_VectorElementType)(__ccgo_up(bp + 4)) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1042:5:
		**(**uintptr)(__ccgo_up(outError)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+497, libc.VaList(bp+56, Xvector_subtype_name(tls, **(**_VectorElementType)(__ccgo_up(bp))), Xvector_subtype_name(tls, **(**_VectorElementType)(__ccgo_up(bp + 4)))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1046:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(a)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1047:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 40)))(tls, **(**uintptr)(__ccgo_up(b)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1048:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1050:3:
	if **(**Tsize_t)(__ccgo_up(bp + 8)) != **(**Tsize_t)(__ccgo_up(bp + 16)) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1051:5:
		**(**uintptr)(__ccgo_up(outError)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+576, libc.VaList(bp+56, **(**Tsize_t)(__ccgo_up(bp + 8)), **(**Tsize_t)(__ccgo_up(bp + 16))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1055:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(a)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1056:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 40)))(tls, **(**uintptr)(__ccgo_up(b)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1057:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1059:3:
	**(**_VectorElementType)(__ccgo_up(element_type)) = **(**_VectorElementType)(__ccgo_up(bp))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1060:3:
	**(**Tsize_t)(__ccgo_up(dimensions)) = **(**Tsize_t)(__ccgo_up(bp + 8))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1061:3:
	**(**Tvector_cleanup)(__ccgo_up(__ccgo_fp_outACleanup)) = **(**Tvector_cleanup)(__ccgo_up(bp + 32))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1062:3:
	**(**Tvector_cleanup)(__ccgo_up(__ccgo_fp_outBCleanup)) = **(**Tvector_cleanup)(__ccgo_up(bp + 40))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1063:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1066:5:
func X_cmp(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1066:40:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1066:42:
	return int32(**(**Ti64)(__ccgo_up(a)) - **(**Ti64)(__ccgo_up(b)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1068:1:
type TVecNpyFile = struct {
	Fpath       uintptr
	FpathLength Tsize_t
}

type VecNpyFile = TVecNpyFile

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1075:13:
func _vec_npy_file(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1076:48:
	var f, path uintptr
	var pathLength Tsize_t
	_, _, _ = f, path, pathLength
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1077:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1078:8:
	path = libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1079:10:
	pathLength = libc.Uint64FromInt32(libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1082:3:
	f = libsqlite3.Xsqlite3_malloc(tls, int32(16))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1083:3:
	if !(f != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1084:5:
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1085:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1087:3:
	libc.Xmemset(tls, f, 0, uint64(16))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1089:3:
	(*TVecNpyFile)(unsafe.Pointer(f)).Fpath = path
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1090:3:
	(*TVecNpyFile)(unsafe.Pointer(f)).FpathLength = pathLength
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1091:3:
	libsqlite3.Xsqlite3_result_pointer(tls, context, f, __ccgo_ts+674, __ccgo_fp(libsqlite3.Xsqlite3_free))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1096:13:
func _vec_f32(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1096:79:
	var rc int32
	var _ /* cleanup at bp+16 */ Tfvec_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* errmsg at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1097:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1099:7:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1103:3:
	rc = _fvec_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1104:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1105:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1106:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1107:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1109:3:
	libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8))*uint64(4)), **(**Tfvec_cleanup)(__ccgo_up(bp + 16)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1111:3:
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1114:13:
func _vec_bit(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1114:79:
	var rc int32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* errmsg at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1115:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1121:3:
	rc = _bitvec_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1122:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1123:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1124:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1125:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1127:3:
	libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8))/uint64(m_CHAR_BIT)), uintptr(-libc.Int32FromInt32(1)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1128:3:
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_BIT))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1129:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1131:13:
func _vec_int8(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1131:80:
	var rc int32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* errmsg at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1132:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1138:3:
	rc = _int8_vec_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1139:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1140:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1141:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1142:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1144:3:
	libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8))), uintptr(-libc.Int32FromInt32(1)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1145:3:
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1146:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1149:13:
func _vec_length(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1150:46:
	var rc int32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* errmsg at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1151:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1158:3:
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1160:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1161:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1162:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1163:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1165:3:
	libsqlite3.Xsqlite3_result_int64(tls, context, libc.Int64FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1166:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1169:13:
func _vec_distance_cosine(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1170:55:
	var rc int32
	var result, result1 Tf32
	var _ /* a at bp+0 */ uintptr
	var _ /* aCleanup at bp+24 */ Tvector_cleanup
	var _ /* b at bp+8 */ uintptr
	var _ /* bCleanup at bp+32 */ Tvector_cleanup
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+48 */ _VectorElementType
	var _ /* error at bp+40 */ uintptr
	_, _, _ = rc, result, result1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1171:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1173:8:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1173:19:
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1178:3:
	rc = Xensure_vector_match(tls, **(**uintptr)(__ccgo_up(argv)), **(**uintptr)(__ccgo_up(argv + 1*8)), bp, bp+8, bp+48, bp+16, bp+24, bp+32, bp+40)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1180:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1181:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 40)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1182:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 40)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1183:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1186:3:
	switch **(**_VectorElementType)(__ccgo_up(bp + 48)) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1187:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1188:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+688, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1191:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1193:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1194:9:
		result = _distance_cosine_float(tls, **(**uintptr)(__ccgo_up(bp)), **(**uintptr)(__ccgo_up(bp + 8)), bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1195:5:
		libsqlite3.Xsqlite3_result_double(tls, context, float64(result))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1196:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1198:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1199:9:
		result1 = _distance_cosine_int8(tls, **(**uintptr)(__ccgo_up(bp)), **(**uintptr)(__ccgo_up(bp + 8)), bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1200:5:
		libsqlite3.Xsqlite3_result_double(tls, context, float64(result1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1201:5:
		goto finish
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1205:1:
	goto finish
finish:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1206:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1207:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1208:3:
	return
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1211:13:
func _vec_distance_l2(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1212:51:
	var rc int32
	var result, result1 Tf32
	var _ /* a at bp+0 */ uintptr
	var _ /* aCleanup at bp+24 */ Tvector_cleanup
	var _ /* b at bp+8 */ uintptr
	var _ /* bCleanup at bp+32 */ Tvector_cleanup
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+48 */ _VectorElementType
	var _ /* error at bp+40 */ uintptr
	_, _, _ = rc, result, result1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1213:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1215:8:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1215:19:
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1220:3:
	rc = Xensure_vector_match(tls, **(**uintptr)(__ccgo_up(argv)), **(**uintptr)(__ccgo_up(argv + 1*8)), bp, bp+8, bp+48, bp+16, bp+24, bp+32, bp+40)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1222:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1223:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 40)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1224:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 40)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1225:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1228:3:
	switch **(**_VectorElementType)(__ccgo_up(bp + 48)) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1229:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1230:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+745, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1232:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1234:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1235:9:
		result = _distance_l2_sqr_float(tls, **(**uintptr)(__ccgo_up(bp)), **(**uintptr)(__ccgo_up(bp + 8)), bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1236:5:
		libsqlite3.Xsqlite3_result_double(tls, context, float64(result))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1237:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1239:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1240:9:
		result1 = _distance_l2_sqr_int8(tls, **(**uintptr)(__ccgo_up(bp)), **(**uintptr)(__ccgo_up(bp + 8)), bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1241:5:
		libsqlite3.Xsqlite3_result_double(tls, context, float64(result1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1242:5:
		goto finish
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1246:1:
	goto finish
finish:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1247:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1248:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1249:3:
	return
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1252:13:
func _vec_distance_l1(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1253:51:
	var rc int32
	var result float64
	var result1 Ti64
	var _ /* a at bp+0 */ uintptr
	var _ /* aCleanup at bp+24 */ Tvector_cleanup
	var _ /* b at bp+8 */ uintptr
	var _ /* bCleanup at bp+32 */ Tvector_cleanup
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+48 */ _VectorElementType
	var _ /* error at bp+40 */ uintptr
	_, _, _ = rc, result, result1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1254:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1261:3:
	rc = Xensure_vector_match(tls, **(**uintptr)(__ccgo_up(argv)), **(**uintptr)(__ccgo_up(argv + 1*8)), bp, bp+8, bp+48, bp+16, bp+24, bp+32, bp+40)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1263:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1264:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 40)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1265:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 40)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1266:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1269:3:
	switch **(**_VectorElementType)(__ccgo_up(bp + 48)) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1270:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1271:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+798, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1273:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1275:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1276:12:
		result = _distance_l1_f32(tls, **(**uintptr)(__ccgo_up(bp)), **(**uintptr)(__ccgo_up(bp + 8)), bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1277:5:
		libsqlite3.Xsqlite3_result_double(tls, context, result)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1278:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1280:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1281:9:
		result1 = int64(_distance_l1_int8(tls, **(**uintptr)(__ccgo_up(bp)), **(**uintptr)(__ccgo_up(bp + 8)), bp+16))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1282:5:
		libsqlite3.Xsqlite3_result_int(tls, context, int32(result1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1283:5:
		goto finish
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1287:1:
	goto finish
finish:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1288:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1289:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1290:3:
	return
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1293:13:
func _vec_distance_hamming(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1294:56:
	var rc int32
	var _ /* a at bp+0 */ uintptr
	var _ /* aCleanup at bp+24 */ Tvector_cleanup
	var _ /* b at bp+8 */ uintptr
	var _ /* bCleanup at bp+32 */ Tvector_cleanup
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+48 */ _VectorElementType
	var _ /* error at bp+40 */ uintptr
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1295:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1297:8:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1297:19:
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1302:3:
	rc = Xensure_vector_match(tls, **(**uintptr)(__ccgo_up(argv)), **(**uintptr)(__ccgo_up(argv + 1*8)), bp, bp+8, bp+48, bp+16, bp+24, bp+32, bp+40)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1304:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1305:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 40)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1306:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 40)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1307:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1310:3:
	switch **(**_VectorElementType)(__ccgo_up(bp + 48)) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1311:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1312:5:
		libsqlite3.Xsqlite3_result_double(tls, context, float64(_distance_hamming(tls, **(**uintptr)(__ccgo_up(bp)), **(**uintptr)(__ccgo_up(bp + 8)), bp+16)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1313:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1315:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1316:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+851, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1319:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1321:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1322:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+914, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1325:5:
		goto finish
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1329:1:
	goto finish
finish:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1330:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1331:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1332:3:
	return
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1335:6:
func Xvec_type_name(tls *libc.TLS, elementType _VectorElementType) (r uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1335:57:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1336:3:
	switch elementType {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1337:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1338:5:
		return __ccgo_ts
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1339:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1340:5:
		return __ccgo_ts + 8
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1341:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1342:5:
		return __ccgo_ts + 13
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1344:3:
	return __ccgo_ts + 17
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1347:13:
func _vec_type(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1347:80:
	var rc int32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* pzError at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1348:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1354:7:
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1356:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1357:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1358:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1359:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1361:3:
	libsqlite3.Xsqlite3_result_text(tls, context, Xvec_type_name(tls, **(**_VectorElementType)(__ccgo_up(bp + 32))), -int32(1), libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1362:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1364:13:
func _vec_quantize_binary(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1365:55:
	var i, i1 Tsize_t
	var out, v2 uintptr
	var rc, res, res1, sz int32
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* pzError at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	var _ /* vectorCleanup at bp+16 */ Tvector_cleanup
	_, _, _, _, _, _, _, _ = i, i1, out, rc, res, res1, sz, v2
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1366:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1372:7:
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1374:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1375:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1376:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1377:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1380:3:
	if **(**Tsize_t)(__ccgo_up(bp + 8)) <= uint64(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1381:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+974, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1382:5:
		goto cleanup
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1383:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1385:3:
	if **(**Tsize_t)(__ccgo_up(bp + 8))%uint64(m_CHAR_BIT) != uint64(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1386:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1013, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1390:5:
		goto cleanup
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1391:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1394:7:
	sz = libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8)) / uint64(m_CHAR_BIT))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1395:6:
	out = libsqlite3.Xsqlite3_malloc(tls, sz)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1396:3:
	if !(out != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1397:5:
		libsqlite3.Xsqlite3_result_error_code(tls, context, int32(m_SQLITE_NOMEM))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1398:5:
		goto cleanup
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1399:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1401:3:
	libc.Xmemset(tls, out, 0, libc.Uint64FromInt32(sz))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1403:3:
	switch **(**_VectorElementType)(__ccgo_up(bp + 32)) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1404:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1406:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1406:17:
		i = uint64(0)
		for {
			if !(i < **(**Tsize_t)(__ccgo_up(bp + 8))) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1407:11:
			res = libc.BoolInt32(float64(**(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4))) > float64(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1408:7:
			v2 = out + uintptr(i/uint64(8))
			*(*Tu8)(unsafe.Pointer(v2)) = Tu8(int32(*(*Tu8)(unsafe.Pointer(v2))) | res<<(i%libc.Uint64FromInt32(8)))
			goto _1
		_1:
			;
			i = i + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1410:5:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1412:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1413:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1413:17:
		i1 = uint64(0)
		for {
			if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 8))) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1414:11:
			res1 = libc.BoolInt32(int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i1)))) > 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1415:7:
			v2 = out + uintptr(i1/uint64(8))
			*(*Tu8)(unsafe.Pointer(v2)) = Tu8(int32(*(*Tu8)(unsafe.Pointer(v2))) | res1<<(i1%libc.Uint64FromInt32(8)))
			goto _3
		_3:
			;
			i1 = i1 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1417:5:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1419:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1420:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1079, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1422:5:
		libsqlite3.Xsqlite3_free(tls, out)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1423:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1426:3:
	libsqlite3.Xsqlite3_result_blob(tls, context, out, sz, __ccgo_fp(libsqlite3.Xsqlite3_free))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1427:3:
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_BIT))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1429:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1430:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1433:13:
func _vec_quantize_int8(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1434:53:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1435:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1440:6:
	out = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1441:7:
	rc = _fvec_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1442:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1443:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1444:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1445:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1448:7:
	sz = libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8)) * uint64(1))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1449:3:
	out = libsqlite3.Xsqlite3_malloc(tls, sz)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1450:3:
	if !(out != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1451:5:
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1452:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1454:3:
	libc.Xmemset(tls, out, 0, libc.Uint64FromInt32(sz))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1456:3:
	if libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv + 1*8))) != int32(m_SQLITE_TEXT) || libc.Uint64FromInt32(libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv + 1*8)))) != libc.Xstrlen(tls, __ccgo_ts+1126) || libsqlite3.Xsqlite3_stricmp(tls, libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv + 1*8))), __ccgo_ts+1126) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1460:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1131, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1462:5:
		libsqlite3.Xsqlite3_free(tls, out)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1463:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1465:7:
	step = float32((float64(1) - -libc.Float64FromFloat64(1)) / libc.Float64FromInt32(255))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1466:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1466:15:
	i = uint64(0)
	for {
		if !(i < **(**Tsize_t)(__ccgo_up(bp + 8))) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1467:12:
		val = (float64(**(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4))) - -libc.Float64FromFloat64(1))/float64(step) - libc.Float64FromInt32(128)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1468:5:
		if !(val <= libc.Float64FromFloat64(127)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1468:26:
			val = float64(127)
		} /* also clamps NaN */
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1469:5:
		if !(val >= -libc.Float64FromFloat64(128)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1469:27:
			val = -libc.Float64FromFloat64(128)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1470:5:
		**(**Ti8)(__ccgo_up(out + uintptr(i))) = int8(val)
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1473:3:
	libsqlite3.Xsqlite3_result_blob(tls, context, out, libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8))*uint64(1)), __ccgo_fp(libsqlite3.Xsqlite3_free))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1474:3:
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1476:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1477:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1480:13:
func _vec_add(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1480:79:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1481:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1483:8:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1483:19:
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1488:3:
	rc = Xensure_vector_match(tls, **(**uintptr)(__ccgo_up(argv)), **(**uintptr)(__ccgo_up(argv + 1*8)), bp, bp+8, bp+48, bp+16, bp+24, bp+32, bp+40)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1490:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1491:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 40)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1492:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 40)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1493:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1496:3:
	switch **(**_VectorElementType)(__ccgo_up(bp + 48)) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1497:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1498:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1183, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1499:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1501:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1502:12:
		outSize = **(**Tsize_t)(__ccgo_up(bp + 16)) * uint64(4)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1503:9:
		out = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(outSize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1504:5:
		if !(out != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1505:7:
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1506:7:
			goto finish
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1508:5:
		libc.Xmemset(tls, out, 0, outSize)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1509:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1509:17:
		i = uint64(0)
		for {
			if !(i < **(**Tsize_t)(__ccgo_up(bp + 16))) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1510:7:
			**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4)) + **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i)*4))
			goto _1
		_1:
			;
			i = i + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1512:5:
		libsqlite3.Xsqlite3_result_blob(tls, context, out, libc.Int32FromUint64(outSize), __ccgo_fp(libsqlite3.Xsqlite3_free))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1513:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1514:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1516:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1517:12:
		outSize1 = **(**Tsize_t)(__ccgo_up(bp + 16)) * uint64(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1518:8:
		out1 = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(outSize1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1519:5:
		if !(out1 != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1520:7:
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1521:7:
			goto finish
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1523:5:
		libc.Xmemset(tls, out1, 0, outSize1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1524:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1524:17:
		i1 = uint64(0)
		for {
			if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 16))) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1525:7:
			**(**Ti8)(__ccgo_up(out1 + uintptr(i1))) = int8(int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i1)))) + int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i1)))))
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1527:5:
		libsqlite3.Xsqlite3_result_blob(tls, context, out1, libc.Int32FromUint64(outSize1), __ccgo_fp(libsqlite3.Xsqlite3_free))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1528:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1529:5:
		goto finish
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1532:1:
	goto finish
finish:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1533:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1534:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1535:3:
	return
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1537:13:
func _vec_sub(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1537:79:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1538:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1540:8:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1540:19:
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1545:3:
	rc = Xensure_vector_match(tls, **(**uintptr)(__ccgo_up(argv)), **(**uintptr)(__ccgo_up(argv + 1*8)), bp, bp+8, bp+48, bp+16, bp+24, bp+32, bp+40)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1547:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1548:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 40)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1549:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 40)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1550:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1553:3:
	switch **(**_VectorElementType)(__ccgo_up(bp + 48)) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1554:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1555:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1219, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1557:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1559:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1560:12:
		outSize = **(**Tsize_t)(__ccgo_up(bp + 16)) * uint64(4)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1561:9:
		out = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(outSize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1562:5:
		if !(out != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1563:7:
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1564:7:
			goto finish
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1566:5:
		libc.Xmemset(tls, out, 0, outSize)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1567:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1567:17:
		i = uint64(0)
		for {
			if !(i < **(**Tsize_t)(__ccgo_up(bp + 16))) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1568:7:
			**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4)) - **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i)*4))
			goto _1
		_1:
			;
			i = i + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1570:5:
		libsqlite3.Xsqlite3_result_blob(tls, context, out, libc.Int32FromUint64(outSize), __ccgo_fp(libsqlite3.Xsqlite3_free))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1571:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1572:5:
		goto finish
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1574:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1575:12:
		outSize1 = **(**Tsize_t)(__ccgo_up(bp + 16)) * uint64(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1576:8:
		out1 = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(outSize1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1577:5:
		if !(out1 != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1578:7:
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1579:7:
			goto finish
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1581:5:
		libc.Xmemset(tls, out1, 0, outSize1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1582:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1582:17:
		i1 = uint64(0)
		for {
			if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 16))) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1583:7:
			**(**Ti8)(__ccgo_up(out1 + uintptr(i1))) = int8(int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i1)))) - int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)) + uintptr(i1)))))
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1585:5:
		libsqlite3.Xsqlite3_result_blob(tls, context, out1, libc.Int32FromUint64(outSize1), __ccgo_fp(libsqlite3.Xsqlite3_free))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1586:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1587:5:
		goto finish
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1590:1:
	goto finish
finish:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1591:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 24)))(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1592:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1593:3:
	return
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1595:13:
func _vec_slice(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1596:45:
	var end, outSize, outSize1, outSize2, rc, start int32
	var i, i1, i2, n Tsize_t
	var out, out1, out2 uintptr
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* err at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = end, i, i1, i2, n, out, out1, out2, outSize, outSize1, outSize2, rc, start
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1597:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1605:7:
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1607:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1608:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1609:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1610:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1613:7:
	start = libsqlite3.Xsqlite3_value_int(tls, **(**uintptr)(__ccgo_up(argv + 1*8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1614:7:
	end = libsqlite3.Xsqlite3_value_int(tls, **(**uintptr)(__ccgo_up(argv + 2*8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1616:3:
	if start < 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1617:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1260, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1619:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1621:3:
	if end < 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1622:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1306, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1624:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1626:3:
	if libc.Uint64FromInt32(start) > **(**Tsize_t)(__ccgo_up(bp + 8)) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1627:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1350, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1630:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1632:3:
	if libc.Uint64FromInt32(end) > **(**Tsize_t)(__ccgo_up(bp + 8)) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1633:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1411, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1636:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1638:3:
	if start > end {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1639:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1470, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1641:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1643:3:
	if start == end {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1644:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1518, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1648:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1650:10:
	n = libc.Uint64FromInt32(end - start)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1652:3:
	switch **(**_VectorElementType)(__ccgo_up(bp + 32)) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1653:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1654:9:
		outSize = libc.Int32FromUint64(n * uint64(4))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1655:9:
		out = libsqlite3.Xsqlite3_malloc(tls, outSize)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1656:5:
		if !(out != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1657:7:
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1658:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1660:5:
		libc.Xmemset(tls, out, 0, libc.Uint64FromInt32(outSize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1661:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1661:17:
		i = uint64(0)
		for {
			if !(i < n) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1662:7:
			**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(libc.Uint64FromInt32(start)+i)*4))
			goto _1
		_1:
			;
			i = i + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1664:5:
		libsqlite3.Xsqlite3_result_blob(tls, context, out, outSize, __ccgo_fp(libsqlite3.Xsqlite3_free))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1665:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1666:5:
		goto done
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1668:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1669:9:
		outSize1 = libc.Int32FromUint64(n * uint64(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1670:8:
		out1 = libsqlite3.Xsqlite3_malloc(tls, outSize1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1671:5:
		if !(out1 != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1672:7:
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1673:7:
			return
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1675:5:
		libc.Xmemset(tls, out1, 0, libc.Uint64FromInt32(outSize1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1676:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1676:17:
		i1 = uint64(0)
		for {
			if !(i1 < n) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1677:7:
			**(**Ti8)(__ccgo_up(out1 + uintptr(i1))) = **(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(libc.Uint64FromInt32(start)+i1)))
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1679:5:
		libsqlite3.Xsqlite3_result_blob(tls, context, out1, outSize1, __ccgo_fp(libsqlite3.Xsqlite3_free))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1680:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_INT8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1681:5:
		goto done
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1683:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1684:5:
		if start%int32(m_CHAR_BIT) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1685:7:
			libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1601, -int32(1))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1686:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1688:5:
		if end%int32(m_CHAR_BIT) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1689:7:
			libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1637, -int32(1))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1690:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1692:9:
		outSize2 = libc.Int32FromUint64(n / uint64(m_CHAR_BIT))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1693:8:
		out2 = libsqlite3.Xsqlite3_malloc(tls, outSize2)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1694:5:
		if !(out2 != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1695:7:
			libsqlite3.Xsqlite3_result_error_nomem(tls, context)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1696:7:
			return
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1698:5:
		libc.Xmemset(tls, out2, 0, libc.Uint64FromInt32(outSize2))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1699:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1699:17:
		i2 = uint64(0)
		for {
			if !(i2 < n/uint64(m_CHAR_BIT)) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1700:7:
			**(**Tu8)(__ccgo_up(out2 + uintptr(i2))) = **(**Tu8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(libc.Uint64FromInt32(start/libc.Int32FromInt32(m_CHAR_BIT))+i2)))
			goto _3
		_3:
			;
			i2 = i2 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1702:5:
		libsqlite3.Xsqlite3_result_blob(tls, context, out2, outSize2, __ccgo_fp(libsqlite3.Xsqlite3_free))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1703:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_BIT))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1704:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1707:1:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1708:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1711:13:
func _vec_to_json(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1712:47:
	var b Tu8
	var i Tsize_t
	var len1, rc int32
	var s, str uintptr
	var value Tf32
	var _ /* cleanup at bp+16 */ Tvector_cleanup
	var _ /* dimensions at bp+8 */ Tsize_t
	var _ /* elementType at bp+32 */ _VectorElementType
	var _ /* err at bp+24 */ uintptr
	var _ /* vector at bp+0 */ uintptr
	_, _, _, _, _, _, _ = b, i, len1, rc, s, str, value
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1713:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1720:7:
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1722:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1723:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1724:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1725:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1728:15:
	str = libsqlite3.Xsqlite3_str_new(tls, libsqlite3.Xsqlite3_context_db_handle(tls, context))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1729:3:
	libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1671)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1730:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1730:15:
	i = uint64(0)
	for {
		if !(i < **(**Tsize_t)(__ccgo_up(bp + 8))) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1731:5:
		if i != uint64(0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1732:7:
			libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1673)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1734:5:
		if **(**_VectorElementType)(__ccgo_up(bp + 32)) == int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1735:11:
			value = **(**Tf32)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i)*4))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1736:7:
			if libc.X__isnanf(tls, value) != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1737:9:
				libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1675)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1739:9:
				libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1680, libc.VaList(bp+48, float64(value)))
			}
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1742:12:
			if **(**_VectorElementType)(__ccgo_up(bp + 32)) == int32(_SQLITE_VEC_ELEMENT_TYPE_INT8) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1743:7:
				libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1683, libc.VaList(bp+48, int32(**(**Ti8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i))))))
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1744:12:
				if **(**_VectorElementType)(__ccgo_up(bp + 32)) == int32(_SQLITE_VEC_ELEMENT_TYPE_BIT) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1745:10:
					b = libc.Uint8FromInt32(libc.Int32FromUint8(**(**Tu8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + uintptr(i/uint64(8))))) >> (i % uint64(m_CHAR_BIT)) & int32(1))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1746:7:
					libsqlite3.Xsqlite3_str_appendf(tls, str, __ccgo_ts+1683, libc.VaList(bp+48, libc.Int32FromUint8(b)))
				}
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1749:3:
	libsqlite3.Xsqlite3_str_appendall(tls, str, __ccgo_ts+1686)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1750:7:
	len1 = libsqlite3.Xsqlite3_str_length(tls, str)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1751:8:
	s = libsqlite3.Xsqlite3_str_finish(tls, str)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1752:3:
	if s != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1753:5:
		libsqlite3.Xsqlite3_result_text(tls, context, s, len1, __ccgo_fp(libsqlite3.Xsqlite3_free))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1754:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(m_JSON_SUBTYPE))
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1756:5:
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1758:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1761:13:
func _vec_normalize(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1762:49:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1763:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1770:7:
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+32, bp+16, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1772:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1773:5:
		libsqlite3.Xsqlite3_result_error(tls, context, **(**uintptr)(__ccgo_up(bp + 24)), -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1774:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1775:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1778:3:
	if **(**_VectorElementType)(__ccgo_up(bp + 32)) != int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1779:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+1688, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1781:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1782:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1785:7:
	outSize = libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8)) * uint64(4))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1786:7:
	out = libsqlite3.Xsqlite3_malloc(tls, outSize)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1787:3:
	if !(out != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1788:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1789:5:
		libsqlite3.Xsqlite3_result_error_code(tls, context, int32(m_SQLITE_NOMEM))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1790:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1792:3:
	libc.Xmemset(tls, out, 0, libc.Uint64FromInt32(outSize))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1794:7:
	v = **(**uintptr)(__ccgo_up(bp))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1796:7:
	norm = libc.Float32FromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1797:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1797:15:
	i = uint64(0)
	for {
		if !(i < **(**Tsize_t)(__ccgo_up(bp + 8))) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1798:5:
		norm = norm + Tf32(**(**Tf32)(__ccgo_up(v + uintptr(i)*4))***(**Tf32)(__ccgo_up(v + uintptr(i)*4)))
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1800:3:
	norm = float32(libc.Xsqrt(tls, float64(norm)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1801:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1801:15:
	i1 = uint64(0)
	for {
		if !(i1 < **(**Tsize_t)(__ccgo_up(bp + 8))) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1802:5:
		**(**Tf32)(__ccgo_up(out + uintptr(i1)*4)) = **(**Tf32)(__ccgo_up(v + uintptr(i1)*4)) / norm
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1805:3:
	libsqlite3.Xsqlite3_result_blob(tls, context, out, libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 8))*uint64(4)), __ccgo_fp(libsqlite3.Xsqlite3_free))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1806:3:
	libsqlite3.Xsqlite3_result_subtype(tls, context, uint32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1807:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 16)))(tls, **(**uintptr)(__ccgo_up(bp)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1810:13:
func __static_text_func(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1811:53:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1812:3:
	_ = argc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1813:3:
	_ = argv
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1814:3:
	libsqlite3.Xsqlite3_result_text(tls, context, libsqlite3.Xsqlite3_user_data(tls, context), -int32(1), libc.UintptrFromInt32(0))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1819:1:
type _Vec0TokenType = int32

const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1820:3:
_TOKEN_TYPE_IDENTIFIER = 0
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1821:3:
_TOKEN_TYPE_DIGIT = 1
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1822:3:
_TOKEN_TYPE_LBRACKET = 2
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1823:3:
_TOKEN_TYPE_RBRACKET = 3
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1824:3:
_TOKEN_TYPE_PLUS = 4
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1825:3:
_TOKEN_TYPE_EQ = 5
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1826:3:
_TOKEN_TYPE_LPAREN = 6
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1827:3:
_TOKEN_TYPE_RPAREN = 7
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1828:3:
_TOKEN_TYPE_COMMA = 8

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1830:1:
type TVec0Token = struct {
	Ftoken_type _Vec0TokenType
	Fstart      uintptr
	Fend        uintptr
}

type Vec0Token = TVec0Token

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1836:5:
func Xis_alpha(tls *libc.TLS, x int8) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1836:22:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1837:3:
	return libc.BoolInt32(int32(x) >= int32('a') && int32(x) <= int32('z') || int32(x) >= int32('A') && int32(x) <= int32('Z'))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1839:5:
func Xis_digit(tls *libc.TLS, x int8) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1839:22:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1839:24:
	return libc.BoolInt32(int32(x) >= int32('0') && int32(x) <= int32('9'))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1840:5:
func Xis_whitespace(tls *libc.TLS, x int8) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1840:27:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1841:3:
	return libc.BoolInt32(int32(x) == int32(' ') || int32(x) == int32('\t') || int32(x) == int32('\n') || int32(x) == int32('\r'))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1848:5:
func Xvec0_token_next(tls *libc.TLS, start uintptr, end uintptr, out uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1848:68:
	var curr int8
	var ptr, start1, start2 uintptr
	_, _, _, _ = curr, ptr, start1, start2
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1849:8:
	ptr = start
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1850:3:
	for ptr < end {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1851:10:
		curr = **(**int8)(__ccgo_up(ptr))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1852:5:
		if Xis_whitespace(tls, curr) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1853:7:
			ptr = ptr + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1854:7:
			continue
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1855:12:
			if int32(curr) == int32('+') {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1856:7:
				ptr = ptr + 1
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1857:7:
				(*TVec0Token)(unsafe.Pointer(out)).Fstart = ptr
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1858:7:
				(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1859:7:
				(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_PLUS)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1860:7:
				return int32(m_VEC0_TOKEN_RESULT_SOME)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1861:12:
				if int32(curr) == int32('[') {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1862:7:
					ptr = ptr + 1
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1863:7:
					(*TVec0Token)(unsafe.Pointer(out)).Fstart = ptr
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1864:7:
					(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1865:7:
					(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_LBRACKET)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1866:7:
					return int32(m_VEC0_TOKEN_RESULT_SOME)
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1867:12:
					if int32(curr) == int32(']') {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1868:7:
						ptr = ptr + 1
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1869:7:
						(*TVec0Token)(unsafe.Pointer(out)).Fstart = ptr
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1870:7:
						(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1871:7:
						(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_RBRACKET)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1872:7:
						return int32(m_VEC0_TOKEN_RESULT_SOME)
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1873:12:
						if int32(curr) == int32('=') {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1874:7:
							ptr = ptr + 1
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1875:7:
							(*TVec0Token)(unsafe.Pointer(out)).Fstart = ptr
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1876:7:
							(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1877:7:
							(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_EQ)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1878:7:
							return int32(m_VEC0_TOKEN_RESULT_SOME)
						} else {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1879:12:
							if int32(curr) == int32('(') {
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1880:7:
								ptr = ptr + 1
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1881:7:
								(*TVec0Token)(unsafe.Pointer(out)).Fstart = ptr
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1882:7:
								(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1883:7:
								(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_LPAREN)
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1884:7:
								return int32(m_VEC0_TOKEN_RESULT_SOME)
							} else {
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1885:12:
								if int32(curr) == int32(')') {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1886:7:
									ptr = ptr + 1
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1887:7:
									(*TVec0Token)(unsafe.Pointer(out)).Fstart = ptr
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1888:7:
									(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1889:7:
									(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_RPAREN)
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1890:7:
									return int32(m_VEC0_TOKEN_RESULT_SOME)
								} else {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1891:12:
									if int32(curr) == int32(',') {
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1892:7:
										ptr = ptr + 1
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1893:7:
										(*TVec0Token)(unsafe.Pointer(out)).Fstart = ptr
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1894:7:
										(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1895:7:
										(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_COMMA)
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1896:7:
										return int32(m_VEC0_TOKEN_RESULT_SOME)
									} else {
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1897:12:
										if Xis_alpha(tls, curr) != 0 {
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1898:12:
											start1 = ptr
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1899:7:
											for ptr < end && (Xis_alpha(tls, **(**int8)(__ccgo_up(ptr))) != 0 || Xis_digit(tls, **(**int8)(__ccgo_up(ptr))) != 0 || int32(**(**int8)(__ccgo_up(ptr))) == int32('_')) {
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1900:9:
												ptr = ptr + 1
											}
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1902:7:
											(*TVec0Token)(unsafe.Pointer(out)).Fstart = start1
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1903:7:
											(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1904:7:
											(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_IDENTIFIER)
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1905:7:
											return int32(m_VEC0_TOKEN_RESULT_SOME)
										} else {
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1906:12:
											if Xis_digit(tls, curr) != 0 {
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1907:12:
												start2 = ptr
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1908:7:
												for ptr < end && Xis_digit(tls, **(**int8)(__ccgo_up(ptr))) != 0 {
													// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1909:9:
													ptr = ptr + 1
												}
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1911:7:
												(*TVec0Token)(unsafe.Pointer(out)).Fstart = start2
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1912:7:
												(*TVec0Token)(unsafe.Pointer(out)).Fend = ptr
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1913:7:
												(*TVec0Token)(unsafe.Pointer(out)).Ftoken_type = int32(_TOKEN_TYPE_DIGIT)
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1914:7:
												return int32(m_VEC0_TOKEN_RESULT_SOME)
											} else {
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1916:7:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1919:3:
	return int32(m_VEC0_TOKEN_RESULT_EOF)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1922:1:
type TVec0Scanner = struct {
	Fstart uintptr
	Fend   uintptr
	Fptr   uintptr
}

type Vec0Scanner = TVec0Scanner

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1928:6:
func Xvec0_scanner_init(tls *libc.TLS, scanner uintptr, source uintptr, source_length int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1929:43:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1930:3:
	(*TVec0Scanner)(unsafe.Pointer(scanner)).Fstart = source
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1931:3:
	(*TVec0Scanner)(unsafe.Pointer(scanner)).Fend = source + uintptr(source_length)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1932:3:
	(*TVec0Scanner)(unsafe.Pointer(scanner)).Fptr = source
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1934:5:
func Xvec0_scanner_next(tls *libc.TLS, scanner uintptr, out uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1934:75:
	var rc int32
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1935:7:
	rc = Xvec0_token_next(tls, (*TVec0Scanner)(unsafe.Pointer(scanner)).Fstart, (*TVec0Scanner)(unsafe.Pointer(scanner)).Fend, out)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1936:3:
	if rc == int32(m_VEC0_TOKEN_RESULT_SOME) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1937:5:
		(*TVec0Scanner)(unsafe.Pointer(scanner)).Fstart = (*TVec0Token)(unsafe.Pointer(out)).Fend
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1939:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1942:5:
func Xvec0_parse_table_option(tls *libc.TLS, source uintptr, source_length int32, out_key uintptr, out_key_length uintptr, out_value uintptr, out_value_length uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1944:70:
	var key, value uintptr
	var keyLength, rc, valueLength int32
	var _ /* scanner at bp+0 */ TVec0Scanner
	var _ /* token at bp+24 */ TVec0Token
	_, _, _, _, _ = key, keyLength, rc, value, valueLength
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1952:3:
	Xvec0_scanner_init(tls, bp, source, source_length)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1954:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1955:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1957:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1959:3:
	key = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1960:3:
	keyLength = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1962:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1963:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_EQ) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1964:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1967:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1968:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && !((**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type == int32(_TOKEN_TYPE_IDENTIFIER) || (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type == int32(_TOKEN_TYPE_DIGIT)) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1971:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1973:3:
	value = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1974:3:
	valueLength = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1976:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1977:3:
	if rc == int32(m_VEC0_TOKEN_RESULT_EOF) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1978:5:
		**(**uintptr)(__ccgo_up(out_key)) = key
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1979:5:
		**(**int32)(__ccgo_up(out_key_length)) = keyLength
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1980:5:
		**(**uintptr)(__ccgo_up(out_value)) = value
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1981:5:
		**(**int32)(__ccgo_up(out_value_length)) = valueLength
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1982:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1984:3:
	return int32(m_SQLITE_ERROR)
}

// C documentation
//
//	/**
//	 * @brief Parse an argv[i] entry of a vec0 virtual table definition, and see if
//	 * it's a PARTITION KEY definition.
//	 *
//	 * @param source: argv[i] source string
//	 * @param source_length: length of the source string
//	 * @param out_column_name: If it is a partition key, the output column name. Same lifetime
//	 * as source, points to specific char *
//	 * @param out_column_name_length: Length of out_column_name in bytes
//	 * @param out_column_type: SQLITE_TEXT or SQLITE_INTEGER.
//	 * @return int: SQLITE_EMPTY if not a PK, SQLITE_OK if it is.
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:1998:5:
func Xvec0_parse_partition_key_definition(tls *libc.TLS, source uintptr, source_length int32, out_column_name uintptr, out_column_name_length uintptr, out_column_type uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2001:56:
	var column_name uintptr
	var column_name_length, column_type, rc int32
	var _ /* scanner at bp+0 */ TVec0Scanner
	var _ /* token at bp+24 */ TVec0Token
	_, _, _, _ = column_name, column_name_length, column_type, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2007:3:
	Xvec0_scanner_init(tls, bp, source, source_length)
	// Check first token is identifier, will be the column name
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2010:7:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2011:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2013:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2016:3:
	column_name = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2017:3:
	column_name_length = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// Check the next token matches "text" or "integer", as column type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2020:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2021:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2023:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2025:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1740, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2026:5:
		column_type = int32(m_SQLITE_TEXT)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2027:10:
		if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1745, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1749, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2031:5:
			column_type = int32(m_SQLITE_INTEGER)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2033:5:
			return int32(m_SQLITE_EMPTY)
		}
	}
	// Check the next token is identifier and matches "partition"
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2037:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2038:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2040:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2042:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1757, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2043:5:
		return int32(m_SQLITE_EMPTY)
	}
	// Check the next token is identifier and matches "key"
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2047:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2048:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2050:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2052:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1767, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2053:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2056:3:
	**(**uintptr)(__ccgo_up(out_column_name)) = column_name
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2057:3:
	**(**int32)(__ccgo_up(out_column_name_length)) = column_name_length
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2058:3:
	**(**int32)(__ccgo_up(out_column_type)) = column_type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2060:3:
	return m_SQLITE_OK
}

// C documentation
//
//	/**
//	 * @brief Parse an argv[i] entry of a vec0 virtual table definition, and see if
//	 * it's an auxiliar column definition, ie `+[name] [type]` like `+contents text`
//	 *
//	 * @param source: argv[i] source string
//	 * @param source_length: length of the source string
//	 * @param out_column_name: If it is a partition key, the output column name. Same lifetime
//	 * as source, points to specific char *
//	 * @param out_column_name_length: Length of out_column_name in bytes
//	 * @param out_column_type: SQLITE_TEXT, SQLITE_INTEGER, SQLITE_FLOAT, or SQLITE_BLOB.
//	 * @return int: SQLITE_EMPTY if not an aux column, SQLITE_OK if it is.
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2075:5:
func Xvec0_parse_auxiliary_column_definition(tls *libc.TLS, source uintptr, source_length int32, out_column_name uintptr, out_column_name_length uintptr, out_column_type uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2078:56:
	var column_name uintptr
	var column_name_length, column_type, rc int32
	var _ /* scanner at bp+0 */ TVec0Scanner
	var _ /* token at bp+24 */ TVec0Token
	_, _, _, _ = column_name, column_name_length, column_type, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2084:3:
	Xvec0_scanner_init(tls, bp, source, source_length)
	// Check first token is '+', which denotes aux columns
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2087:7:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2088:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_PLUS) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2090:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2093:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2094:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2096:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2099:3:
	column_name = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2100:3:
	column_name_length = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// Check the next token matches "text" or "integer", as column type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2103:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2104:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2106:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2108:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1740, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2109:5:
		column_type = int32(m_SQLITE_TEXT)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2110:10:
		if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1745, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1749, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2114:5:
			column_type = int32(m_SQLITE_INTEGER)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2115:10:
			if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1771, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1777, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2119:5:
				column_type = int32(m_SQLITE_FLOAT)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2120:10:
				if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1784, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2121:5:
					column_type = int32(m_SQLITE_BLOB)
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2123:5:
					return int32(m_SQLITE_EMPTY)
				}
			}
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2126:3:
	**(**uintptr)(__ccgo_up(out_column_name)) = column_name
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2127:3:
	**(**int32)(__ccgo_up(out_column_name_length)) = column_name_length
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2128:3:
	**(**int32)(__ccgo_up(out_column_type)) = column_type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2130:3:
	return m_SQLITE_OK
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2133:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2139:3:
type Tvec0_metadata_column_kind = int32

type vec0_metadata_column_kind = Tvec0_metadata_column_kind

const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2134:3:
_VEC0_METADATA_COLUMN_KIND_BOOLEAN = 0
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2135:3:
_VEC0_METADATA_COLUMN_KIND_INTEGER = 1
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2136:3:
_VEC0_METADATA_COLUMN_KIND_FLOAT = 2
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2137:3:
_VEC0_METADATA_COLUMN_KIND_TEXT = 3

// C documentation
//
//	/**
//	 * @brief Parse an argv[i] entry of a vec0 virtual table definition, and see if
//	 * it's an metadata column definition, ie `[name] [type]` like `is_released boolean`
//	 *
//	 * @param source: argv[i] source string
//	 * @param source_length: length of the source string
//	 * @param out_column_name: If it is a metadata column, the output column name. Same lifetime
//	 * as source, points to specific char *
//	 * @param out_column_name_length: Length of out_column_name in bytes
//	 * @param out_column_type: one of vec0_metadata_column_kind
//	 * @return int: SQLITE_EMPTY if not an metadata column, SQLITE_OK if it is.
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2153:5:
func Xvec0_parse_metadata_column_definition(tls *libc.TLS, source uintptr, source_length int32, out_column_name uintptr, out_column_name_length uintptr, out_column_type uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2156:78:
	var column_name, t uintptr
	var column_name_length, n, rc int32
	var column_type Tvec0_metadata_column_kind
	var _ /* scanner at bp+0 */ TVec0Scanner
	var _ /* token at bp+24 */ TVec0Token
	_, _, _, _, _, _ = column_name, column_name_length, column_type, n, rc, t
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2163:3:
	Xvec0_scanner_init(tls, bp, source, source_length)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2165:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2166:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2168:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2171:3:
	column_name = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2172:3:
	column_name_length = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// Check the next token matches a valid metadata type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2175:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2176:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2178:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2180:8:
	t = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2181:7:
	n = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2182:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1789, n) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1797, n) == 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2183:5:
		column_type = int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2184:9:
		if libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1802, n) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1808, n) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1749, n) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1745, n) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2185:5:
			column_type = int32(_VEC0_METADATA_COLUMN_KIND_INTEGER)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2186:9:
			if libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1771, n) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1777, n) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1818, n) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1826, n) == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2187:5:
				column_type = int32(_VEC0_METADATA_COLUMN_KIND_FLOAT)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2188:10:
				if libsqlite3.Xsqlite3_strnicmp(tls, t, __ccgo_ts+1740, n) == 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2189:5:
					column_type = int32(_VEC0_METADATA_COLUMN_KIND_TEXT)
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2191:5:
					return int32(m_SQLITE_EMPTY)
				}
			}
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2194:3:
	**(**uintptr)(__ccgo_up(out_column_name)) = column_name
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2195:3:
	**(**int32)(__ccgo_up(out_column_name_length)) = column_name_length
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2196:3:
	**(**Tvec0_metadata_column_kind)(__ccgo_up(out_column_type)) = column_type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2198:3:
	return m_SQLITE_OK
}

// C documentation
//
//	/**
//	 * @brief Parse an argv[i] entry of a vec0 virtual table definition, and see if
//	 * it's a PRIMARY KEY definition.
//	 *
//	 * @param source: argv[i] source string
//	 * @param source_length: length of the source string
//	 * @param out_column_name: If it is a PK, the output column name. Same lifetime
//	 * as source, points to specific char *
//	 * @param out_column_name_length: Length of out_column_name in bytes
//	 * @param out_column_type: SQLITE_TEXT or SQLITE_INTEGER.
//	 * @return int: SQLITE_EMPTY if not a PK, SQLITE_OK if it is.
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2213:5:
func Xvec0_parse_primary_key_definition(tls *libc.TLS, source uintptr, source_length int32, out_column_name uintptr, out_column_name_length uintptr, out_column_type uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2216:56:
	var column_name uintptr
	var column_name_length, column_type, rc int32
	var _ /* scanner at bp+0 */ TVec0Scanner
	var _ /* token at bp+24 */ TVec0Token
	_, _, _, _ = column_name, column_name_length, column_type, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2222:3:
	Xvec0_scanner_init(tls, bp, source, source_length)
	// Check first token is identifier, will be the column name
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2225:7:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2226:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2228:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2231:3:
	column_name = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2232:3:
	column_name_length = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// Check the next token matches "text" or "integer", as column type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2235:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2236:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2238:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2240:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1740, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2241:5:
		column_type = int32(m_SQLITE_TEXT)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2242:10:
		if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1745, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1749, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2246:5:
			column_type = int32(m_SQLITE_INTEGER)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2248:5:
			return int32(m_SQLITE_EMPTY)
		}
	}
	// Check the next token is identifier and matches "primary"
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2252:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2253:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2255:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2257:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1830, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2258:5:
		return int32(m_SQLITE_EMPTY)
	}
	// Check the next token is identifier and matches "key"
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2262:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2263:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2265:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2267:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1767, int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend)-int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2268:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2271:3:
	**(**uintptr)(__ccgo_up(out_column_name)) = column_name
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2272:3:
	**(**int32)(__ccgo_up(out_column_name_length)) = column_name_length
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2273:3:
	**(**int32)(__ccgo_up(out_column_type)) = column_type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2275:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2278:1:
type _Vec0DistanceMetrics = int32

const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2279:3:
_VEC0_DISTANCE_METRIC_L2 = 1
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2280:3:
_VEC0_DISTANCE_METRIC_COSINE = 2
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2281:3:
_VEC0_DISTANCE_METRIC_L1 = 3

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2284:1:
type TVectorColumnDefinition = struct {
	Fname            uintptr
	Fname_length     int32
	Fdimensions      Tsize_t
	Felement_type    _VectorElementType
	Fdistance_metric _Vec0DistanceMetrics
}

type VectorColumnDefinition = TVectorColumnDefinition

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2292:1:
type TVec0PartitionColumnDefinition = struct {
	Ftype1       int32
	Fname        uintptr
	Fname_length int32
}

type Vec0PartitionColumnDefinition = TVec0PartitionColumnDefinition

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2298:1:
type TVec0AuxiliaryColumnDefinition = struct {
	Ftype1       int32
	Fname        uintptr
	Fname_length int32
}

type Vec0AuxiliaryColumnDefinition = TVec0AuxiliaryColumnDefinition

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2303:1:
type TVec0MetadataColumnDefinition = struct {
	Fkind        Tvec0_metadata_column_kind
	Fname        uintptr
	Fname_length int32
}

type Vec0MetadataColumnDefinition = TVec0MetadataColumnDefinition

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2309:8:
func Xvector_byte_size(tls *libc.TLS, element_type _VectorElementType, dimensions Tsize_t) (r Tsize_t) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2310:44:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2311:3:
	switch element_type {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2312:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2313:5:
		return dimensions * uint64(4)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2314:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2315:5:
		return dimensions * uint64(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2316:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2317:5:
		return dimensions / uint64(m_CHAR_BIT)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2319:3:
	return uint64(0)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2322:8:
func Xvector_column_byte_size(tls *libc.TLS, column TVectorColumnDefinition) (r Tsize_t) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2322:70:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2323:3:
	return Xvector_byte_size(tls, column.Felement_type, column.Fdimensions)
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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2336:5:
func Xvec0_parse_vector_column(tls *libc.TLS, source uintptr, source_length int32, outColumn uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2337:67:
	var dimensions, keyLength, nameLength, rc, valueLength int32
	var distanceMetric _Vec0DistanceMetrics
	var elementType _VectorElementType
	var key, name, value uintptr
	var _ /* scanner at bp+0 */ TVec0Scanner
	var _ /* token at bp+24 */ TVec0Token
	_, _, _, _, _, _, _, _, _, _ = dimensions, distanceMetric, elementType, key, keyLength, name, nameLength, rc, value, valueLength
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2348:28:
	distanceMetric = int32(_VEC0_DISTANCE_METRIC_L2)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2351:3:
	Xvec0_scanner_init(tls, bp, source, source_length)
	// starts with an identifier
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2354:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2356:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2358:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2361:3:
	name = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2362:3:
	nameLength = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
	// vector column type comes next: float, int, or bit
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2365:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2367:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2369:5:
		return int32(m_SQLITE_EMPTY)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2371:3:
	if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1771, int32(5)) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1838, int32(3)) == 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2373:5:
		elementType = int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2374:10:
		if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+8, int32(4)) == 0 || libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+1842, int32(2)) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2376:5:
			elementType = int32(_SQLITE_VEC_ELEMENT_TYPE_INT8)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2377:10:
			if libsqlite3.Xsqlite3_strnicmp(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+13, int32(3)) == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2378:5:
				elementType = int32(_SQLITE_VEC_ELEMENT_TYPE_BIT)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2380:5:
				return int32(m_SQLITE_EMPTY)
			}
		}
	}
	// left '[' bracket
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2384:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2385:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_LBRACKET) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2386:5:
		return int32(m_SQLITE_EMPTY)
	}
	// digit, for vector dimension length
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2390:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2391:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_DIGIT) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2392:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2394:3:
	dimensions = libc.Xatoi(tls, (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2395:3:
	if dimensions <= 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2396:5:
		return int32(m_SQLITE_ERROR)
	}
	// // right ']' bracket
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2400:3:
	rc = Xvec0_scanner_next(tls, bp, bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2401:3:
	if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_RBRACKET) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2402:5:
		return int32(m_SQLITE_ERROR)
	}
	// any other tokens left should be column-level options , ex `key=value`
	// ex `distance_metric=L2 distance_metric=cosine` should error
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2407:3:
	for int32(1) != 0 {
		// should be EOF or identifier (option key)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2409:5:
		rc = Xvec0_scanner_next(tls, bp, bp+24)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2410:5:
		if rc == int32(m_VEC0_TOKEN_RESULT_EOF) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2411:7:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2414:5:
		if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2416:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2419:10:
		key = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2420:9:
		keyLength = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2422:5:
		if libsqlite3.Xsqlite3_strnicmp(tls, key, __ccgo_ts+1845, keyLength) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2424:7:
			if elementType == int32(_SQLITE_VEC_ELEMENT_TYPE_BIT) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2425:9:
				return int32(m_SQLITE_ERROR)
			}
			// ensure equal sign after distance_metric
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2428:7:
			rc = Xvec0_scanner_next(tls, bp, bp+24)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2429:7:
			if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_EQ) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2430:9:
				return int32(m_SQLITE_ERROR)
			}
			// distance_metric value, an identifier (L2, cosine, etc)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2434:7:
			rc = Xvec0_scanner_next(tls, bp, bp+24)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2435:7:
			if rc != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TVec0Token)(__ccgo_up(bp + 24))).Ftoken_type != int32(_TOKEN_TYPE_IDENTIFIER) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2437:9:
				return int32(m_SQLITE_ERROR)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2440:12:
			value = (**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2441:11:
			valueLength = int32(int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fend) - int64((**(**TVec0Token)(__ccgo_up(bp + 24))).Fstart))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2442:7:
			if libsqlite3.Xsqlite3_strnicmp(tls, value, __ccgo_ts+1861, valueLength) == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2443:9:
				distanceMetric = int32(_VEC0_DISTANCE_METRIC_L2)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2444:14:
				if libsqlite3.Xsqlite3_strnicmp(tls, value, __ccgo_ts+1864, valueLength) == 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2445:9:
					distanceMetric = int32(_VEC0_DISTANCE_METRIC_L1)
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2446:14:
					if libsqlite3.Xsqlite3_strnicmp(tls, value, __ccgo_ts+1867, valueLength) == 0 {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2447:9:
						distanceMetric = int32(_VEC0_DISTANCE_METRIC_COSINE)
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2449:9:
						return int32(m_SQLITE_ERROR)
					}
				}
			}
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2454:7:
			return int32(m_SQLITE_ERROR)
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2458:3:
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+56, nameLength, name))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2459:3:
	if !((*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fname != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2460:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2462:3:
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fname_length = nameLength
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2463:3:
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fdistance_metric = distanceMetric
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2464:3:
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Felement_type = elementType
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2465:3:
	(*TVectorColumnDefinition)(unsafe.Pointer(outColumn)).Fdimensions = libc.Uint64FromInt32(dimensions)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2466:3:
	return m_SQLITE_OK
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2471:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2471:30:
type Tvec_each_vtab = struct {
	Fbase Tsqlite3_vtab
}

type vec_each_vtab = Tvec_each_vtab

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2476:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2476:32:
type Tvec_each_cursor = struct {
	Fbase        Tsqlite3_vtab_cursor
	FiRowid      Ti64
	Fvector_type _VectorElementType
	Fvector      uintptr
	Fdimensions  Tsize_t
	Fcleanup     Tvector_cleanup
}

type vec_each_cursor = Tvec_each_cursor

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2486:12:
func _vec_eachConnect(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2488:42:
	var pNew uintptr
	var rc int32
	_, _ = pNew, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2489:3:
	_ = pAux
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2490:3:
	_ = argc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2491:3:
	_ = argv
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2492:3:
	_ = pzErr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2496:3:
	rc = libsqlite3.Xsqlite3_declare_vtab(tls, db, __ccgo_ts+1879)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2499:3:
	if rc == m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2500:5:
		pNew = libsqlite3.Xsqlite3_malloc(tls, int32(24))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2501:5:
		**(**uintptr)(__ccgo_up(ppVtab)) = pNew
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2502:5:
		if pNew == uintptr(0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2503:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2504:5:
		libc.Xmemset(tls, pNew, 0, uint64(24))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2506:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2509:12:
func _vec_eachDisconnect(tls *libc.TLS, pVtab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2509:52:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2510:17:
	p = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2511:3:
	libsqlite3.Xsqlite3_free(tls, p)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2512:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2515:12:
func _vec_eachOpen(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2515:74:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2516:3:
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2518:3:
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(48))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2519:3:
	if pCur == uintptr(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2520:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2521:3:
	libc.Xmemset(tls, pCur, 0, uint64(48))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2522:3:
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2523:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2526:12:
func _vec_eachClose(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2526:52:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2527:19:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2528:3:
	if (*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2529:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fcleanup})))(tls, (*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2531:3:
	libsqlite3.Xsqlite3_free(tls, pCur)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2532:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2535:12:
func _vec_eachBestIndex(tls *libc.TLS, pVTab uintptr, pIdxInfo uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2536:60:
	var hasVector, i int32
	var pCons uintptr
	_, _, _ = hasVector, i, pCons
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2537:3:
	_ = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2538:7:
	hasVector = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2539:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2539:12:
	i = 0
	for {
		if !(i < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2540:43:
		pCons = (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12
		// printf("i=%d iColumn=%d, op=%d, usable=%d\n", i, pCons->iColumn,
		// pCons->op, pCons->usable);
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2543:5:
		switch (*Tsqlite3_index_constraint)(unsafe.Pointer(pCons)).FiColumn {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2544:5:
		case int32(m_VEC_EACH_COLUMN_VECTOR):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2545:7:
			if libc.Int32FromUint8((*Tsqlite3_index_constraint)(unsafe.Pointer(pCons)).Fop) == int32(m_SQLITE_INDEX_CONSTRAINT_EQ) && (*Tsqlite3_index_constraint)(unsafe.Pointer(pCons)).Fusable != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2546:9:
				hasVector = int32(1)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2547:9:
				(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i)*8))).FargvIndex = int32(1)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2548:9:
				(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i)*8))).Fomit = uint8(1)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2550:7:
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2554:3:
	if !(hasVector != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2555:5:
		return int32(m_SQLITE_CONSTRAINT)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2558:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = libc.Float64FromInt32(100000)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2559:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(100000)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2561:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2564:12:
func _vec_eachFilter(tls *libc.TLS, pVtabCursor uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2565:79:
	var pCur uintptr
	var rc int32
	var _ /* pzErrMsg at bp+0 */ uintptr
	_, _ = pCur, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2566:3:
	_ = idxNum
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2567:3:
	_ = idxStr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2568:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2569:19:
	pCur = pVtabCursor
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2571:3:
	if (*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2572:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fcleanup})))(tls, (*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2573:5:
		(*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2577:7:
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), pCur+24, pCur+32, pCur+16, pCur+40, bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2579:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2580:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2581:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2583:3:
	(*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2584:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2587:12:
func _vec_eachRowid(tls *libc.TLS, cur uintptr, pRowid uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2587:74:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2588:19:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2589:3:
	**(**Tsqlite_int64)(__ccgo_up(pRowid)) = (*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2590:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2593:12:
func _vec_eachEof(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2593:50:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2594:19:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2595:3:
	return libc.BoolInt32((*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid >= libc.Int64FromUint64((*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fdimensions))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2598:12:
func _vec_eachNext(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2598:51:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2599:19:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2600:3:
	(*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid = (*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid + 1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2601:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2604:12:
func _vec_eachColumn(tls *libc.TLS, cur uintptr, context uintptr, i int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2605:34:
	var pCur uintptr
	var x Tu8
	_, _ = pCur, x
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2606:19:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2607:3:
	switch i {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2608:3:
	case m_VEC_EACH_COLUMN_VALUE:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2609:5:
		switch (*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector_type {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2610:5:
		case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2611:7:
			libsqlite3.Xsqlite3_result_double(tls, context, float64(**(**Tf32)(__ccgo_up((*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector + uintptr((*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid)*4))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2612:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2614:5:
			fallthrough
		case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2615:10:
			x = **(**Tu8)(__ccgo_up((*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector + uintptr((*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid/int64(m_CHAR_BIT))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2616:7:
			libsqlite3.Xsqlite3_result_int(tls, context, libc.BoolInt32(libc.Int32FromUint8(x)&(int32(0b10000000)>>((*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid%int64(m_CHAR_BIT))) > 0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2618:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2620:5:
			fallthrough
		case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2621:7:
			libsqlite3.Xsqlite3_result_int(tls, context, int32(**(**Ti8)(__ccgo_up((*Tvec_each_cursor)(unsafe.Pointer(pCur)).Fvector + uintptr((*Tvec_each_cursor)(unsafe.Pointer(pCur)).FiRowid)))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2622:7:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2626:5:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2628:3:
	return m_SQLITE_OK
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2631:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2631:23:
var _vec_eachModule = Tsqlite3_module{}

func init() {
	p := unsafe.Pointer(&_vec_eachModule)
	*(*uintptr)(unsafe.Add(p, 16)) = __ccgo_fp(_vec_eachConnect)
	*(*uintptr)(unsafe.Add(p, 24)) = __ccgo_fp(_vec_eachBestIndex)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(_vec_eachDisconnect)
	*(*uintptr)(unsafe.Add(p, 48)) = __ccgo_fp(_vec_eachOpen)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_vec_eachClose)
	*(*uintptr)(unsafe.Add(p, 64)) = __ccgo_fp(_vec_eachFilter)
	*(*uintptr)(unsafe.Add(p, 72)) = __ccgo_fp(_vec_eachNext)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(_vec_eachEof)
	*(*uintptr)(unsafe.Add(p, 88)) = __ccgo_fp(_vec_eachColumn)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(_vec_eachRowid)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2665:1:
type _NpyTokenType = int32

const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2666:3:
_NPY_TOKEN_TYPE_IDENTIFIER = 0
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2667:3:
_NPY_TOKEN_TYPE_NUMBER = 1
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2668:3:
_NPY_TOKEN_TYPE_LPAREN = 2
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2669:3:
_NPY_TOKEN_TYPE_RPAREN = 3
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2670:3:
_NPY_TOKEN_TYPE_LBRACE = 4
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2671:3:
_NPY_TOKEN_TYPE_RBRACE = 5
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2672:3:
_NPY_TOKEN_TYPE_COLON = 6
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2673:3:
_NPY_TOKEN_TYPE_COMMA = 7
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2674:3:
_NPY_TOKEN_TYPE_STRING = 8
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2675:3:
_NPY_TOKEN_TYPE_FALSE = 9

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2678:1:
type TNpyToken = struct {
	Ftoken_type _NpyTokenType
	Fstart      uintptr
	Fend        uintptr
}

type NpyToken = TNpyToken

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2684:5:
func Xnpy_token_next(tls *libc.TLS, start uintptr, end uintptr, out uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2685:42:
	var curr uint8
	var ptr, start1, start2, v1 uintptr
	_, _, _, _, _ = curr, ptr, start1, start2, v1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2686:17:
	ptr = start
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2687:3:
	for ptr < end {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2688:19:
		curr = **(**uint8)(__ccgo_up(ptr))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2689:5:
		if Xis_whitespace(tls, libc.Int8FromUint8(curr)) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2690:7:
			ptr = ptr + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2691:7:
			continue
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2692:12:
			if libc.Int32FromUint8(curr) == int32('(') {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2693:7:
				v1 = ptr
				ptr = ptr + 1
				(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2694:7:
				(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2695:7:
				(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_LPAREN)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2696:7:
				return int32(m_VEC0_TOKEN_RESULT_SOME)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2697:12:
				if libc.Int32FromUint8(curr) == int32(')') {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2698:7:
					v1 = ptr
					ptr = ptr + 1
					(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2699:7:
					(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2700:7:
					(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_RPAREN)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2701:7:
					return int32(m_VEC0_TOKEN_RESULT_SOME)
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2702:12:
					if libc.Int32FromUint8(curr) == int32('{') {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2703:7:
						v1 = ptr
						ptr = ptr + 1
						(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2704:7:
						(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2705:7:
						(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_LBRACE)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2706:7:
						return int32(m_VEC0_TOKEN_RESULT_SOME)
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2707:12:
						if libc.Int32FromUint8(curr) == int32('}') {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2708:7:
							v1 = ptr
							ptr = ptr + 1
							(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2709:7:
							(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2710:7:
							(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_RBRACE)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2711:7:
							return int32(m_VEC0_TOKEN_RESULT_SOME)
						} else {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2712:12:
							if libc.Int32FromUint8(curr) == int32(':') {
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2713:7:
								v1 = ptr
								ptr = ptr + 1
								(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2714:7:
								(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2715:7:
								(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_COLON)
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2716:7:
								return int32(m_VEC0_TOKEN_RESULT_SOME)
							} else {
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2717:12:
								if libc.Int32FromUint8(curr) == int32(',') {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2718:7:
									v1 = ptr
									ptr = ptr + 1
									(*TNpyToken)(unsafe.Pointer(out)).Fstart = v1
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2719:7:
									(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2720:7:
									(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_COMMA)
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2721:7:
									return int32(m_VEC0_TOKEN_RESULT_SOME)
								} else {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2722:12:
									if libc.Int32FromUint8(curr) == int32('\'') {
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2723:21:
										start1 = ptr
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2724:7:
										ptr = ptr + 1
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2725:7:
										for ptr < end {
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2726:9:
											if libc.Int32FromUint8(**(**uint8)(__ccgo_up(ptr))) == int32('\'') {
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2727:11:
												break
											}
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2729:9:
											ptr = ptr + 1
										}
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2731:7:
										if ptr >= end || libc.Int32FromUint8(**(**uint8)(__ccgo_up(ptr))) != int32('\'') {
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2732:9:
											return int32(m_VEC0_TOKEN_RESULT_ERROR)
										}
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2734:7:
										(*TNpyToken)(unsafe.Pointer(out)).Fstart = start1
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2735:7:
										ptr = ptr + 1
										v1 = ptr
										(*TNpyToken)(unsafe.Pointer(out)).Fend = v1
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2736:7:
										(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_STRING)
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2737:7:
										return int32(m_VEC0_TOKEN_RESULT_SOME)
									} else {
										// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2738:12:
										if libc.Int32FromUint8(curr) == int32('F') && libc.Xstrncmp(tls, ptr, __ccgo_ts+1916, libc.Xstrlen(tls, __ccgo_ts+1916)) == 0 {
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2740:7:
											(*TNpyToken)(unsafe.Pointer(out)).Fstart = ptr
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2741:7:
											(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr + uintptr(libc.Int32FromUint64(libc.Xstrlen(tls, __ccgo_ts+1916)))
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2742:7:
											ptr = (*TNpyToken)(unsafe.Pointer(out)).Fend
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2743:7:
											(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_FALSE)
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2744:7:
											return int32(m_VEC0_TOKEN_RESULT_SOME)
										} else {
											// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2745:12:
											if Xis_digit(tls, libc.Int8FromUint8(curr)) != 0 {
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2746:21:
												start2 = ptr
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2747:7:
												for ptr < end && Xis_digit(tls, libc.Int8FromUint8(**(**uint8)(__ccgo_up(ptr)))) != 0 {
													// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2748:9:
													ptr = ptr + 1
												}
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2750:7:
												(*TNpyToken)(unsafe.Pointer(out)).Fstart = start2
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2751:7:
												(*TNpyToken)(unsafe.Pointer(out)).Fend = ptr
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2752:7:
												(*TNpyToken)(unsafe.Pointer(out)).Ftoken_type = int32(_NPY_TOKEN_TYPE_NUMBER)
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2753:7:
												return int32(m_VEC0_TOKEN_RESULT_SOME)
											} else {
												// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2755:7:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2758:3:
	return int32(m_VEC0_TOKEN_RESULT_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2761:1:
type TNpyScanner = struct {
	Fstart uintptr
	Fend   uintptr
	Fptr   uintptr
}

type NpyScanner = TNpyScanner

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2767:6:
func Xnpy_scanner_init(tls *libc.TLS, scanner uintptr, source uintptr, source_length int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2768:42:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2769:3:
	(*TNpyScanner)(unsafe.Pointer(scanner)).Fstart = source
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2770:3:
	(*TNpyScanner)(unsafe.Pointer(scanner)).Fend = source + uintptr(source_length)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2771:3:
	(*TNpyScanner)(unsafe.Pointer(scanner)).Fptr = source
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2774:5:
func Xnpy_scanner_next(tls *libc.TLS, scanner uintptr, out uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2774:72:
	var rc int32
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2775:7:
	rc = Xnpy_token_next(tls, (*TNpyScanner)(unsafe.Pointer(scanner)).Fstart, (*TNpyScanner)(unsafe.Pointer(scanner)).Fend, out)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2776:3:
	if rc == int32(m_VEC0_TOKEN_RESULT_SOME) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2777:5:
		(*TNpyScanner)(unsafe.Pointer(scanner)).Fstart = (*TNpyToken)(unsafe.Pointer(out)).Fend
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2779:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2783:5:
func Xparse_npy_header(tls *libc.TLS, pVTab uintptr, header uintptr, headerLength Tsize_t, out_element_type uintptr, fortran_order uintptr, numElements uintptr, numDimensions uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2787:45:
	var first Tsize_t
	var key uintptr
	var rc, v1 int32
	var _ /* scanner at bp+0 */ TNpyScanner
	var _ /* token at bp+24 */ TNpyToken
	_, _, _, _ = first, key, rc, v1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2792:3:
	Xnpy_scanner_init(tls, bp, header, libc.Int32FromUint64(headerLength))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2794:3:
	if Xnpy_scanner_next(tls, bp, bp+24) != int32(m_VEC0_TOKEN_RESULT_SOME) && (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_LBRACE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2796:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+1922, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2798:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2800:3:
	for int32(1) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2801:5:
		rc = Xnpy_scanner_next(tls, bp, bp+24)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2802:5:
		if rc != int32(m_VEC0_TOKEN_RESULT_SOME) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2803:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+1985, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2804:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2807:5:
		if (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type == int32(_NPY_TOKEN_TYPE_RBRACE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2808:7:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2810:5:
		if (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_STRING) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2811:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+2041, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2813:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2815:19:
		key = (**(**TNpyToken)(__ccgo_up(bp + 24))).Fstart
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2817:5:
		rc = Xnpy_scanner_next(tls, bp, bp+24)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2818:5:
		if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_COLON) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2820:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+2109, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2822:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2825:5:
		if libc.Xstrncmp(tls, key, __ccgo_ts+2177, libc.Xstrlen(tls, __ccgo_ts+2177)) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2826:7:
			rc = Xnpy_scanner_next(tls, bp, bp+24)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2827:7:
			if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_STRING) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2829:9:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+2185, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2831:9:
				return int32(m_SQLITE_ERROR)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2833:7:
			if libc.Xstrncmp(tls, (**(**TNpyToken)(__ccgo_up(bp + 24))).Fstart, __ccgo_ts+2254, libc.Xstrlen(tls, __ccgo_ts+2254)) != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2834:9:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+2260, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2837:9:
				return int32(m_SQLITE_ERROR)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2839:7:
			**(**_VectorElementType)(__ccgo_up(out_element_type)) = int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2840:12:
			if libc.Xstrncmp(tls, key, __ccgo_ts+2349, libc.Xstrlen(tls, __ccgo_ts+2349)) == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2842:7:
				rc = Xnpy_scanner_next(tls, bp, bp+24)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2843:7:
				if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_FALSE) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2845:9:
					Xvtab_set_error(tls, pVTab, __ccgo_ts+2365, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2848:9:
					return int32(m_SQLITE_ERROR)
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2850:7:
				**(**int32)(__ccgo_up(fortran_order)) = 0
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2851:12:
				if libc.Xstrncmp(tls, key, __ccgo_ts+2462, libc.Xstrlen(tls, __ccgo_ts+2462)) == 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2854:7:
					rc = Xnpy_scanner_next(tls, bp, bp+24)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2855:7:
					if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_LPAREN) {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2857:9:
						Xvtab_set_error(tls, pVTab, __ccgo_ts+2470, 0)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2859:9:
						return int32(m_SQLITE_ERROR)
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2862:7:
					rc = Xnpy_scanner_next(tls, bp, bp+24)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2863:7:
					if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_NUMBER) {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2865:9:
						Xvtab_set_error(tls, pVTab, __ccgo_ts+2543, 0)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2867:9:
						return int32(m_SQLITE_ERROR)
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2869:7:
					first = libc.Uint64FromInt64(libc.Xstrtol(tls, (**(**TNpyToken)(__ccgo_up(bp + 24))).Fstart, libc.UintptrFromInt32(0), int32(10)))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2871:7:
					rc = Xnpy_scanner_next(tls, bp, bp+24)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2872:7:
					if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_COMMA) {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2874:9:
						Xvtab_set_error(tls, pVTab, __ccgo_ts+2612, 0)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2876:9:
						return int32(m_SQLITE_ERROR)
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2879:7:
					rc = Xnpy_scanner_next(tls, bp, bp+24)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2880:7:
					if rc != int32(m_VEC0_TOKEN_RESULT_SOME) {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2881:9:
						Xvtab_set_error(tls, pVTab, __ccgo_ts+2678, 0)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2883:9:
						return int32(m_SQLITE_ERROR)
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2885:7:
					if (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type == int32(_NPY_TOKEN_TYPE_NUMBER) {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2886:9:
						**(**Tsize_t)(__ccgo_up(numElements)) = first
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2887:9:
						**(**Tsize_t)(__ccgo_up(numDimensions)) = libc.Uint64FromInt64(libc.Xstrtol(tls, (**(**TNpyToken)(__ccgo_up(bp + 24))).Fstart, libc.UintptrFromInt32(0), int32(10)))
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2888:9:
						rc = Xnpy_scanner_next(tls, bp, bp+24)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2889:9:
						if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_RPAREN) {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2891:11:
							Xvtab_set_error(tls, pVTab, __ccgo_ts+2747, 0)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2893:11:
							return int32(m_SQLITE_ERROR)
						}
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2895:14:
						if (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type == int32(_NPY_TOKEN_TYPE_RPAREN) {
							// '(0,)' means an empty array!
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2897:9:
							if first != 0 {
								v1 = int32(1)
							} else {
								v1 = 0
							}
							**(**Tsize_t)(__ccgo_up(numElements)) = libc.Uint64FromInt32(v1)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2898:9:
							**(**Tsize_t)(__ccgo_up(numDimensions)) = first
						} else {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2900:9:
							Xvtab_set_error(tls, pVTab, __ccgo_ts+2819, 0)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2901:9:
							return int32(m_SQLITE_ERROR)
						}
					}
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2904:7:
					Xvtab_set_error(tls, pVTab, __ccgo_ts+2874, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2905:7:
					return int32(m_SQLITE_ERROR)
				}
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2908:5:
		rc = Xnpy_scanner_next(tls, bp, bp+24)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2909:5:
		if rc != int32(m_VEC0_TOKEN_RESULT_SOME) || (**(**TNpyToken)(__ccgo_up(bp + 24))).Ftoken_type != int32(_NPY_TOKEN_TYPE_COMMA) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2911:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+2929, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2912:7:
			return int32(m_SQLITE_ERROR)
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2916:3:
	return m_SQLITE_OK
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2919:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2919:34:
type Tvec_npy_each_vtab = struct {
	Fbase Tsqlite3_vtab
}

type vec_npy_each_vtab = Tvec_npy_each_vtab

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2924:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2927:3:
type Tvec_npy_each_input_type = int32

type vec_npy_each_input_type = Tvec_npy_each_input_type

const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2925:3:
_VEC_NPY_EACH_INPUT_BUFFER = 0
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2926:3:
_VEC_NPY_EACH_INPUT_FILE = 1

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2929:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2929:36:
type Tvec_npy_each_cursor = struct {
	Fbase              Tsqlite3_vtab_cursor
	FiRowid            Ti64
	FelementType       _VectorElementType
	FnElements         Tsize_t
	FnDimensions       Tsize_t
	Finput_type        Tvec_npy_each_input_type
	Fvector            uintptr
	Ffile              uintptr
	FchunksBuffer      uintptr
	FchunksBufferSize  Tsize_t
	FmaxChunks         Tsize_t
	FcurrentChunkIndex Tsize_t
	FcurrentChunkSize  Tsize_t
	Feof               int32
}

type vec_npy_each_cursor = Tvec_npy_each_cursor

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2976:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2976:22:
var _NPY_MAGIC = [6]uint8{147, 'N', 'U', 'M', 'P', 'Y'}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2979:5:
func Xparse_npy_file(tls *libc.TLS, pVTab uintptr, file uintptr, pCur uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2979:80:
	var dataSize, expectedDataSize Ti32
	var fileSize int64
	var headerX uintptr
	var major, minor Tu8
	var n, rc int32
	var totalHeaderLength Tsize_t
	var _ /* element_type at bp+16 */ _VectorElementType
	var _ /* fortran_order at bp+12 */ int32
	var _ /* header at bp+0 */ [10]uint8
	var _ /* headerLength at bp+10 */ Tuint16_t
	var _ /* numDimensions at bp+32 */ Tsize_t
	var _ /* numElements at bp+24 */ Tsize_t
	_, _, _, _, _, _, _, _, _ = dataSize, expectedDataSize, fileSize, headerX, major, minor, n, rc, totalHeaderLength
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2981:3:
	libc.Xfseek(tls, file, 0, int32(m_SEEK_END))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2982:8:
	fileSize = libc.Xftell(tls, file)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2984:3:
	libc.Xfseek(tls, file, 0, m_SEEK_SET)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2987:3:
	n = libc.Int32FromUint64(libc.Xfread(tls, bp, uint64(1), uint64(10), file))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2988:3:
	if n != int32(10) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2989:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+2988, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2990:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2993:3:
	if libc.Xmemcmp(tls, uintptr(unsafe.Pointer(&_NPY_MAGIC)), bp, uint64(6)) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2994:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3015, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2996:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:2999:6:
	major = (**(**[10]uint8)(__ccgo_up(bp)))[int32(6)]
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3000:6:
	minor = (**(**[10]uint8)(__ccgo_up(bp)))[int32(7)]
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3001:12:
	**(**Tuint16_t)(__ccgo_up(bp + 10)) = uint16(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3002:3:
	libc.Xmemcpy(tls, bp+10, bp+8, uint64(2))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3004:10:
	totalHeaderLength = libc.Uint64FromInt64(6) + libc.Uint64FromInt64(1) + libc.Uint64FromInt64(1) + libc.Uint64FromInt64(2) + uint64(**(**Tuint16_t)(__ccgo_up(bp + 10)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3006:7:
	dataSize = libc.Int32FromUint64(libc.Uint64FromInt64(fileSize) - totalHeaderLength)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3007:3:
	if dataSize < 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3008:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3068, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3009:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3012:17:
	headerX = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint16(**(**Tuint16_t)(__ccgo_up(bp + 10))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3013:3:
	if **(**Tuint16_t)(__ccgo_up(bp + 10)) != 0 && !(headerX != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3014:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3017:3:
	n = libc.Int32FromUint64(libc.Xfread(tls, headerX, uint64(1), uint64(**(**Tuint16_t)(__ccgo_up(bp + 10))), file))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3018:3:
	if n != libc.Int32FromUint16(**(**Tuint16_t)(__ccgo_up(bp + 10))) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3019:5:
		libsqlite3.Xsqlite3_free(tls, headerX)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3020:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3068, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3021:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3028:7:
	rc = Xparse_npy_header(tls, pVTab, headerX, uint64(**(**Tuint16_t)(__ccgo_up(bp + 10))), bp+16, bp+12, bp+24, bp+32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3030:3:
	libsqlite3.Xsqlite3_free(tls, headerX)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3031:3:
	if rc != m_SQLITE_OK {
		// parse_npy_header already attackes an error emssage
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3033:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3036:7:
	expectedDataSize = libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(bp + 24)) * Xvector_byte_size(tls, **(**_VectorElementType)(__ccgo_up(bp + 16)), **(**Tsize_t)(__ccgo_up(bp + 32))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3038:3:
	if expectedDataSize != dataSize {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3039:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3110, libc.VaList(bp+48, expectedDataSize, dataSize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3042:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3045:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FmaxChunks = uint64(1024)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3046:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBufferSize = Xvector_byte_size(tls, **(**_VectorElementType)(__ccgo_up(bp + 16)), **(**Tsize_t)(__ccgo_up(bp + 32))) * (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FmaxChunks
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3048:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBufferSize))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3049:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBufferSize != 0 && !((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3050:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3053:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkSize = libc.Xfread(tls, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer, Xvector_byte_size(tls, **(**_VectorElementType)(__ccgo_up(bp + 16)), **(**Tsize_t)(__ccgo_up(bp + 32))), (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FmaxChunks, file)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3057:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkIndex = uint64(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3058:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType = **(**_VectorElementType)(__ccgo_up(bp + 16))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3059:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnElements = **(**Tsize_t)(__ccgo_up(bp + 24))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3060:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions = **(**Tsize_t)(__ccgo_up(bp + 32))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3061:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Finput_type = int32(_VEC_NPY_EACH_INPUT_FILE)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3063:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Feof = libc.BoolInt32((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkSize == uint64(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3064:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile = file
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3065:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3069:5:
func Xparse_npy_buffer(tls *libc.TLS, pVTab uintptr, buffer uintptr, bufferLength int32, data uintptr, numElements uintptr, numDimensions uintptr, element_type uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3072:60:
	var dataSize, expectedDataSize, totalHeaderLength Ti32
	var header uintptr
	var major, minor Tu8
	var rc int32
	var _ /* fortran_order at bp+4 */ int32
	var _ /* headerLength at bp+0 */ Tuint16_t
	_, _, _, _, _, _, _ = dataSize, expectedDataSize, header, major, minor, rc, totalHeaderLength
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3074:3:
	if bufferLength < int32(10) {
		// IMP: V03312_20150
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3076:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3171, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3077:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3079:3:
	if libc.Xmemcmp(tls, uintptr(unsafe.Pointer(&_NPY_MAGIC)), buffer, uint64(6)) != 0 {
		// V11954_28792
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3081:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3193, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3082:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3085:6:
	major = **(**uint8)(__ccgo_up(buffer + 6))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3086:6:
	minor = **(**uint8)(__ccgo_up(buffer + 7))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3087:12:
	**(**Tuint16_t)(__ccgo_up(bp)) = uint16(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3088:3:
	libc.Xmemcpy(tls, bp, buffer+8, uint64(2))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3090:7:
	totalHeaderLength = libc.Int32FromUint64(libc.Uint64FromInt64(6) + libc.Uint64FromInt64(1) + libc.Uint64FromInt64(1) + libc.Uint64FromInt64(2) + uint64(**(**Tuint16_t)(__ccgo_up(bp))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3092:7:
	dataSize = bufferLength - totalHeaderLength
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3094:3:
	if dataSize < 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3095:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3241, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3096:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3099:23:
	header = buffer + 10
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3102:7:
	rc = Xparse_npy_header(tls, pVTab, header, uint64(**(**Tuint16_t)(__ccgo_up(bp))), element_type, bp+4, numElements, numDimensions)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3104:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3105:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3108:7:
	expectedDataSize = libc.Int32FromUint64(**(**Tsize_t)(__ccgo_up(numElements)) * Xvector_byte_size(tls, **(**_VectorElementType)(__ccgo_up(element_type)), **(**Tsize_t)(__ccgo_up(numDimensions))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3110:3:
	if expectedDataSize != dataSize {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3111:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+3278, libc.VaList(bp+16, expectedDataSize, dataSize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3114:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3117:3:
	**(**uintptr)(__ccgo_up(data)) = buffer + uintptr(totalHeaderLength)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3118:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3121:12:
func _vec_npy_eachConnect(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3123:46:
	var pNew uintptr
	var rc int32
	_, _ = pNew, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3124:3:
	_ = pAux
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3125:3:
	_ = argc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3126:3:
	_ = argv
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3127:3:
	_ = pzErr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3131:3:
	rc = libsqlite3.Xsqlite3_declare_vtab(tls, db, __ccgo_ts+3334)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3134:3:
	if rc == m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3135:5:
		pNew = libsqlite3.Xsqlite3_malloc(tls, int32(24))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3136:5:
		**(**uintptr)(__ccgo_up(ppVtab)) = pNew
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3137:5:
		if pNew == uintptr(0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3138:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3139:5:
		libc.Xmemset(tls, pNew, 0, uint64(24))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3141:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3144:12:
func _vec_npy_eachDisconnect(tls *libc.TLS, pVtab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3144:56:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3145:21:
	p = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3146:3:
	libsqlite3.Xsqlite3_free(tls, p)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3147:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3150:12:
func _vec_npy_eachOpen(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3150:78:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3151:3:
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3153:3:
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(112))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3154:3:
	if pCur == uintptr(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3155:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3156:3:
	libc.Xmemset(tls, pCur, 0, uint64(112))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3157:3:
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3158:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3161:12:
func _vec_npy_eachClose(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3161:56:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3162:23:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3164:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3165:5:
		libc.Xfclose(tls, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3166:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3169:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3170:5:
		libsqlite3.Xsqlite3_free(tls, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3171:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3173:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Fvector != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3174:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Fvector = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3176:3:
	libsqlite3.Xsqlite3_free(tls, pCur)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3177:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3180:12:
func _vec_npy_eachBestIndex(tls *libc.TLS, pVTab uintptr, pIdxInfo uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3181:64:
	var hasInput, i int32
	var pCons uintptr
	_, _, _ = hasInput, i, pCons
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3183:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3183:12:
	i = 0
	for {
		if !(i < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3184:43:
		pCons = (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12
		// printf("i=%d iColumn=%d, op=%d, usable=%d\n", i, pCons->iColumn,
		// pCons->op, pCons->usable);
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3187:5:
		switch (*Tsqlite3_index_constraint)(unsafe.Pointer(pCons)).FiColumn {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3188:5:
		case int32(m_VEC_NPY_EACH_COLUMN_INPUT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3189:7:
			if libc.Int32FromUint8((*Tsqlite3_index_constraint)(unsafe.Pointer(pCons)).Fop) == int32(m_SQLITE_INDEX_CONSTRAINT_EQ) && (*Tsqlite3_index_constraint)(unsafe.Pointer(pCons)).Fusable != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3190:9:
				hasInput = int32(1)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3191:9:
				(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i)*8))).FargvIndex = int32(1)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3192:9:
				(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i)*8))).Fomit = uint8(1)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3194:7:
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3198:3:
	if !(hasInput != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3199:5:
		(*Tsqlite3_vtab)(unsafe.Pointer(pVTab)).FzErrMsg = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+3371, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3200:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3203:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = libc.Float64FromInt32(100000)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3204:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(100000)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3206:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3209:12:
func _vec_npy_eachFilter(tls *libc.TLS, pVtabCursor uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3211:53:
	var f, file, input, pCur, v1 uintptr
	var inputLength, rc int32
	var _ /* data at bp+0 */ uintptr
	var _ /* element_type at bp+24 */ _VectorElementType
	var _ /* numDimensions at bp+16 */ Tsize_t
	var _ /* numElements at bp+8 */ Tsize_t
	_, _, _, _, _, _, _ = f, file, input, inputLength, pCur, rc, v1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3212:3:
	_ = idxNum
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3213:3:
	_ = idxStr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3214:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3217:23:
	pCur = pVtabCursor
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3220:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3221:5:
		libc.Xfclose(tls, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3222:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3225:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3226:5:
		libsqlite3.Xsqlite3_free(tls, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3227:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3229:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Fvector != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3230:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Fvector = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3234:21:
	f = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3235:3:
	v1 = libsqlite3.Xsqlite3_value_pointer(tls, **(**uintptr)(__ccgo_up(argv)), __ccgo_ts+674)
	f = v1
	if v1 != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3236:10:
		file = libc.Xfopen(tls, (*TVecNpyFile)(unsafe.Pointer(f)).Fpath, __ccgo_ts+3398)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3237:5:
		if !(file != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3238:7:
			Xvtab_set_error(tls, (*Tsqlite3_vtab_cursor)(unsafe.Pointer(pVtabCursor)).FpVtab, __ccgo_ts+3400, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3239:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3242:5:
		rc = Xparse_npy_file(tls, (*Tsqlite3_vtab_cursor)(unsafe.Pointer(pVtabCursor)).FpVtab, file, pCur)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3243:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3245:7:
			libc.Xfclose(tls, file)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3247:7:
			return rc
		}
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3254:25:
		input = libsqlite3.Xsqlite3_value_blob(tls, **(**uintptr)(__ccgo_up(argv)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3255:9:
		inputLength = libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3261:5:
		rc = Xparse_npy_buffer(tls, (*Tsqlite3_vtab_cursor)(unsafe.Pointer(pVtabCursor)).FpVtab, input, inputLength, bp, bp+8, bp+16, bp+24)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3263:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3264:7:
			return rc
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3267:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Fvector = **(**uintptr)(__ccgo_up(bp))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3268:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType = **(**_VectorElementType)(__ccgo_up(bp + 24))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3269:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnElements = **(**Tsize_t)(__ccgo_up(bp + 8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3270:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions = **(**Tsize_t)(__ccgo_up(bp + 16))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3271:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Finput_type = int32(_VEC_NPY_EACH_INPUT_BUFFER)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3274:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FiRowid = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3275:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3278:12:
func _vec_npy_eachRowid(tls *libc.TLS, cur uintptr, pRowid uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3278:78:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3279:23:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3280:3:
	**(**Tsqlite_int64)(__ccgo_up(pRowid)) = (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FiRowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3281:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3284:12:
func _vec_npy_eachEof(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3284:54:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3285:23:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3286:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Finput_type == int32(_VEC_NPY_EACH_INPUT_BUFFER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3287:5:
		return libc.BoolInt32(!((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnElements != 0) || libc.Uint64FromInt64((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FiRowid) >= (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnElements)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3289:3:
	return (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Feof
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3292:12:
func _vec_npy_eachNext(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3292:55:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3293:23:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3294:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FiRowid = (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FiRowid + 1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3295:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Finput_type == int32(_VEC_NPY_EACH_INPUT_BUFFER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3296:5:
		return m_SQLITE_OK
	}
	// else: input is a file
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3301:3:
	(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkIndex = (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkIndex + 1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3302:3:
	if (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkIndex >= (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkSize {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3303:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkSize = libc.Xfread(tls, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer, Xvector_byte_size(tls, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions), (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FmaxChunks, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Ffile)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3307:5:
		if !((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkSize != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3308:7:
			(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Feof = int32(1)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3310:5:
		(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkIndex = uint64(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3313:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3316:12:
func _vec_npy_eachColumnBuffer(tls *libc.TLS, pCur uintptr, context uintptr, i int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3317:70:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3318:3:
	switch i {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3319:3:
	case m_VEC_NPY_EACH_COLUMN_VECTOR:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3320:5:
		libsqlite3.Xsqlite3_result_subtype(tls, context, libc.Uint32FromInt32((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3321:5:
		switch (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3322:5:
		case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3323:7:
			libsqlite3.Xsqlite3_result_blob(tls, context, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Fvector+uintptr(libc.Uint64FromInt64((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FiRowid)*(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions*uint64(4)), libc.Int32FromUint64((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions*uint64(4)), uintptr(-libc.Int32FromInt32(1)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3329:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3331:5:
			fallthrough
		case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3332:5:
			fallthrough
		case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
			// https://github.com/asg017/sqlite-vec/issues/42
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3334:7:
			libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+3426, -int32(1))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3336:7:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3340:5:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3343:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3345:12:
func _vec_npy_eachColumnFile(tls *libc.TLS, pCur uintptr, context uintptr, i int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3346:68:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3347:3:
	switch i {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3348:3:
	case m_VEC_NPY_EACH_COLUMN_VECTOR:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3349:5:
		switch (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FelementType {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3350:5:
		case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3351:7:
			libsqlite3.Xsqlite3_result_blob(tls, context, (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FchunksBuffer+uintptr((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FcurrentChunkIndex*(*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions*uint64(4)), libc.Int32FromUint64((*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).FnDimensions*uint64(4)), uintptr(-libc.Int32FromInt32(1)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3357:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3359:5:
			fallthrough
		case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3360:5:
			fallthrough
		case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
			// https://github.com/asg017/sqlite-vec/issues/42
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3362:7:
			libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+3426, -int32(1))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3364:7:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3367:5:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3370:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3372:12:
func _vec_npy_eachColumn(tls *libc.TLS, cur uintptr, context uintptr, i int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3373:64:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3374:23:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3375:3:
	switch (*Tvec_npy_each_cursor)(unsafe.Pointer(pCur)).Finput_type {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3376:3:
	case int32(_VEC_NPY_EACH_INPUT_BUFFER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3377:5:
		return _vec_npy_eachColumnBuffer(tls, pCur, context, i)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3378:3:
		fallthrough
	case int32(_VEC_NPY_EACH_INPUT_FILE):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3379:5:
		return _vec_npy_eachColumnFile(tls, pCur, context, i)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3381:3:
	return int32(m_SQLITE_ERROR)
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3384:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3384:23:
var _vec_npy_eachModule = Tsqlite3_module{}

func init() {
	p := unsafe.Pointer(&_vec_npy_eachModule)
	*(*uintptr)(unsafe.Add(p, 16)) = __ccgo_fp(_vec_npy_eachConnect)
	*(*uintptr)(unsafe.Add(p, 24)) = __ccgo_fp(_vec_npy_eachBestIndex)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(_vec_npy_eachDisconnect)
	*(*uintptr)(unsafe.Add(p, 48)) = __ccgo_fp(_vec_npy_eachOpen)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_vec_npy_eachClose)
	*(*uintptr)(unsafe.Add(p, 64)) = __ccgo_fp(_vec_npy_eachFilter)
	*(*uintptr)(unsafe.Add(p, 72)) = __ccgo_fp(_vec_npy_eachNext)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(_vec_npy_eachEof)
	*(*uintptr)(unsafe.Add(p, 88)) = __ccgo_fp(_vec_npy_eachColumn)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(_vec_npy_eachRowid)
}

/// 1) schema, 2) original vtab table name

/// 1) schema, 2) original vtab table name

// vec0 tables with a text primary keys are still backed by int64 primary keys,
// since a fixed-length rowid is required for vec0 chunks. But we add a new 'id
// text unique' column to emulate a text primary key interface.

/// 1) schema, 2) original vtab table name

/// 1) schema, 2) original vtab table name
//
// IMPORTANT: "rowid" is declared as PRIMARY KEY but WITHOUT the INTEGER type.
// This means it is NOT a true SQLite rowid alias — the user-defined "rowid"
// column and the internal SQLite rowid (_rowid_) are two separate values.
// When inserting, both must be set explicitly to keep them in sync. See the
// _rowid_ bindings in vec0_new_chunk() and the explanation in
// SHADOW_TABLE_ROWID_QUIRK below.

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3481:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3481:26:
type Tvec0_vtab = struct {
	Fbase                       Tsqlite3_vtab
	Fdb                         uintptr
	FpkIsText                   int32
	FnumVectorColumns           int32
	FnumPartitionColumns        int32
	FnumAuxiliaryColumns        int32
	FnumMetadataColumns         int32
	FschemaName                 uintptr
	FtableName                  uintptr
	FshadowRowidsName           uintptr
	FshadowChunksName           uintptr
	Fuser_column_kinds          [52]Tvec0_user_column_kind
	Fuser_column_idxs           [52]Tuint8_t
	FshadowVectorChunksNames    [16]uintptr
	FshadowMetadataChunksNames  [16]uintptr
	Fvector_columns             [16]TVectorColumnDefinition
	Fparitition_columns         [4]TVec0PartitionColumnDefinition
	Fauxiliary_columns          [16]TVec0AuxiliaryColumnDefinition
	Fmetadata_columns           [16]TVec0MetadataColumnDefinition
	Fchunk_size                 int32
	FstmtLatestChunk            uintptr
	FstmtRowidsInsertRowid      uintptr
	FstmtRowidsInsertId         uintptr
	FstmtRowidsUpdatePosition   uintptr
	FstmtRowidsGetChunkPosition uintptr
}

type vec0_vtab = Tvec0_vtab

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3492:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3504:3:
type Tvec0_user_column_kind = int32

type vec0_user_column_kind = Tvec0_user_column_kind

const
// vector column, ie "contents_embedding float[1024]"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3494:3:
_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR = 1
const

// partition key column, ie "user_id integer partition key"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3497:3:
_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION = 2
const

//

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3500:3:
_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY = 3
const

// metadata column that can be filtered, ie "genre text"

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3503:3:
_SQLITE_VEC0_USER_COLUMN_KIND_METADATA = 4

// C documentation
//
//	/**
//	 * @brief Finalize all the sqlite3_stmt members in a vec0_vtab.
//	 *
//	 * @param p vec0_vtab pointer
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3632:6:
func Xvec0_free_resources(tls *libc.TLS, p uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3632:40:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3633:3:
	libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3634:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3635:3:
	libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3636:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3637:3:
	libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3638:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3639:3:
	libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3640:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3641:3:
	libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3642:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition = libc.UintptrFromInt32(0)
}

// C documentation
//
//	/**
//	 * @brief Free all memory and sqlite3_stmt members of a vec0_vtab
//	 *
//	 * @param p vec0_vtab pointer
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3650:6:
func Xvec0_free(tls *libc.TLS, p uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3650:30:
	var i, i1, i2, i3 int32
	_, _, _, _ = i, i1, i2, i3
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3651:3:
	Xvec0_free_resources(tls, p)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3653:3:
	libsqlite3.Xsqlite3_free(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3654:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3655:3:
	libsqlite3.Xsqlite3_free(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3656:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FtableName = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3657:3:
	libsqlite3.Xsqlite3_free(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3658:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3659:3:
	libsqlite3.Xsqlite3_free(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowRowidsName)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3660:3:
	(*Tvec0_vtab)(unsafe.Pointer(p)).FshadowRowidsName = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3662:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3662:12:
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3663:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3664:5:
		**(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)) = libc.UintptrFromInt32(0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3666:5:
		libsqlite3.Xsqlite3_free(tls, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fname)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3667:5:
		(**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fname = libc.UintptrFromInt32(0)
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3670:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3670:12:
	i1 = 0
	for {
		if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3671:5:
		libsqlite3.Xsqlite3_free(tls, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(i1)*24))).Fname)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3672:5:
		(**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(i1)*24))).Fname = libc.UintptrFromInt32(0)
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3675:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3675:12:
	i2 = 0
	for {
		if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3676:5:
		libsqlite3.Xsqlite3_free(tls, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(i2)*24))).Fname)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3677:5:
		(**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(i2)*24))).Fname = libc.UintptrFromInt32(0)
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3680:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3680:12:
	i3 = 0
	for {
		if !(i3 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumMetadataColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3681:5:
		libsqlite3.Xsqlite3_free(tls, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(i3)*24))).Fname)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3682:5:
		(**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(i3)*24))).Fname = libc.UintptrFromInt32(0)
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3686:5:
func Xvec0_num_defined_user_columns(tls *libc.TLS, p uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3686:49:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3687:3:
	return (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns + (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns + (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns + (*Tvec0_vtab)(unsafe.Pointer(p)).FnumMetadataColumns
}

// C documentation
//
//	/**
//	 * @brief Returns the index of the distance hidden column for the given vec0
//	 * table.
//	 *
//	 * @param p vec0 table
//	 * @return int
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3697:5:
func Xvec0_column_distance_idx(tls *libc.TLS, p uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3697:44:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3698:3:
	return int32(m_VEC0_COLUMN_USERN_START) + (Xvec0_num_defined_user_columns(tls, p) - int32(1)) + int32(m_VEC0_COLUMN_OFFSET_DISTANCE)
}

// C documentation
//
//	/**
//	 * @brief Returns the index of the k hidden column for the given vec0 table.
//	 *
//	 * @param p vec0 table
//	 * @return int k column index
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3708:5:
func Xvec0_column_k_idx(tls *libc.TLS, p uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3708:37:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3709:3:
	return int32(m_VEC0_COLUMN_USERN_START) + (Xvec0_num_defined_user_columns(tls, p) - int32(1)) + int32(m_VEC0_COLUMN_OFFSET_K)
}

// C documentation
//
//	/**
//	 * Returns 1 if the given column-based index is a valid vector column,
//	 * 0 otherwise.
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3717:5:
func Xvec0_column_idx_is_vector(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3717:65:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3718:3:
	return libc.BoolInt32(column_idx >= int32(m_VEC0_COLUMN_USERN_START) && column_idx <= int32(m_VEC0_COLUMN_USERN_START)+Xvec0_num_defined_user_columns(tls, pVtab)-int32(1) && **(**Tvec0_user_column_kind)(__ccgo_up(pVtab + 88 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START))*4)) == int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR))
}

// C documentation
//
//	/**
//	 * Returns the vector index of the given user column index.
//	 * ONLY call if validated with vec0_column_idx_is_vector before
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3727:5:
func Xvec0_column_idx_to_vector_idx(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3727:69:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3728:3:
	_ = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3729:3:
	return libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(pVtab + 296 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START)))))
}

// C documentation
//
//	/**
//	 * Returns 1 if the given column-based index is a "partition key" column,
//	 * 0 otherwise.
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3735:5:
func Xvec0_column_idx_is_partition(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3735:68:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3736:3:
	return libc.BoolInt32(column_idx >= int32(m_VEC0_COLUMN_USERN_START) && column_idx <= int32(m_VEC0_COLUMN_USERN_START)+Xvec0_num_defined_user_columns(tls, pVtab)-int32(1) && **(**Tvec0_user_column_kind)(__ccgo_up(pVtab + 88 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START))*4)) == int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION))
}

// C documentation
//
//	/**
//	 * Returns the partition column index of the given user column index.
//	 * ONLY call if validated with vec0_column_idx_is_vector before
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3745:5:
func Xvec0_column_idx_to_partition_idx(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3745:72:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3746:3:
	_ = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3747:3:
	return libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(pVtab + 296 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START)))))
}

// C documentation
//
//	/**
//	 * Returns 1 if the given column-based index is a auxiliary column,
//	 * 0 otherwise.
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3754:5:
func Xvec0_column_idx_is_auxiliary(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3754:68:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3755:3:
	return libc.BoolInt32(column_idx >= int32(m_VEC0_COLUMN_USERN_START) && column_idx <= int32(m_VEC0_COLUMN_USERN_START)+Xvec0_num_defined_user_columns(tls, pVtab)-int32(1) && **(**Tvec0_user_column_kind)(__ccgo_up(pVtab + 88 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START))*4)) == int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY))
}

// C documentation
//
//	/**
//	 * Returns the auxiliary column index of the given user column index.
//	 * ONLY call if validated with vec0_column_idx_to_partition_idx before
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3764:5:
func Xvec0_column_idx_to_auxiliary_idx(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3764:72:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3765:3:
	_ = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3766:3:
	return libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(pVtab + 296 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START)))))
}

// C documentation
//
//	/**
//	 * Returns 1 if the given column-based index is a metadata column,
//	 * 0 otherwise.
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3773:5:
func Xvec0_column_idx_is_metadata(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3773:67:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3774:3:
	return libc.BoolInt32(column_idx >= int32(m_VEC0_COLUMN_USERN_START) && column_idx <= int32(m_VEC0_COLUMN_USERN_START)+Xvec0_num_defined_user_columns(tls, pVtab)-int32(1) && **(**Tvec0_user_column_kind)(__ccgo_up(pVtab + 88 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START))*4)) == int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA))
}

// C documentation
//
//	/**
//	 * Returns the metadata column index of the given user column index.
//	 * ONLY call if validated with vec0_column_idx_is_metadata before
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3783:5:
func Xvec0_column_idx_to_metadata_idx(tls *libc.TLS, pVtab uintptr, column_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3783:71:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3784:3:
	_ = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3785:3:
	return libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(pVtab + 296 + uintptr(column_idx-int32(m_VEC0_COLUMN_USERN_START)))))
}

// C documentation
//
//	/**
//	 * @brief Retrieve the chunk_id, chunk_offset, and possible "id" value
//	 * of a vec0_vtab row with the provided rowid
//	 *
//	 * @param p vec0_vtab
//	 * @param rowid the rowid of the row to query
//	 * @param id output, optional sqlite3_value to provide the id.
//	 *            Useful for text PK rows. Must be freed with sqlite3_value_free()
//	 * @param chunk_id output, the chunk_id the row belongs to
//	 * @param chunk_offset  output, the offset within the chunk the row belongs to
//	 * @return SQLITE_ROW on success, error code otherwise. SQLITE_EMPTY if row DNE
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3800:5:
func Xvec0_get_chunk_position(tls *libc.TLS, p uintptr, rowid Ti64, id uintptr, chunk_id uintptr, chunk_offset uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3801:63:
	var rc int32
	var value, zSql uintptr
	_, _, _ = rc, value, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3804:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3805:16:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+3469, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3809:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3810:7:
			rc = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3811:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3813:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), p+2024, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3814:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3815:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3816:7:
			Xvtab_set_error(tls, p, __ccgo_ts+3541, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3819:7:
			goto cleanup
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3823:3:
	libsqlite3.Xsqlite3_bind_int64(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition, int32(1), rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3824:3:
	rc = libsqlite3.Xsqlite3_step(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition)
	// special case: when no results, return SQLITE_EMPTY to convey "that chunk
	// position doesnt exist"
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3827:3:
	if rc == int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3828:5:
		rc = int32(m_SQLITE_EMPTY)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3829:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3831:3:
	if rc != int32(m_SQLITE_ROW) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3832:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3835:3:
	if id != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3836:19:
		value = libsqlite3.Xsqlite3_column_value(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3838:5:
		**(**uintptr)(__ccgo_up(id)) = libsqlite3.Xsqlite3_value_dup(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3839:5:
		if !(**(**uintptr)(__ccgo_up(id)) != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3840:7:
			rc = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3841:7:
			goto cleanup
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3845:3:
	if chunk_id != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3846:5:
		**(**Ti64)(__ccgo_up(chunk_id)) = libsqlite3.Xsqlite3_column_int64(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition, int32(1))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3848:3:
	if chunk_offset != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3849:5:
		**(**Ti64)(__ccgo_up(chunk_offset)) = libsqlite3.Xsqlite3_column_int64(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition, int32(2))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3852:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3854:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3855:3:
	libsqlite3.Xsqlite3_reset(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3856:3:
	libsqlite3.Xsqlite3_clear_bindings(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3857:3:
	return rc
}

// C documentation
//
//	/**
//	 * @brief Return the id value from the _rowids table where _rowids.rowid =
//	 * rowid.
//	 *
//	 * @param pVtab: vec0 table to query
//	 * @param rowid: rowid of the row to query.
//	 * @param out: A dup'ed sqlite3_value of the id column. Might be null.
//	 *                         Must be cleaned up with sqlite3_value_free().
//	 * @returns SQLITE_OK on success, error code on failure
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3870:5:
func Xvec0_get_id_value_from_rowid(tls *libc.TLS, pVtab uintptr, rowid Ti64, out uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3871:55:
	// PERF: different strategy than get_chunk_position?
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3873:3:
	return Xvec0_get_chunk_position(tls, pVtab, rowid, out, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3876:5:
func Xvec0_rowid_from_id(tls *libc.TLS, p uintptr, valueId uintptr, rowid uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3876:74:
	var rc int32
	var zSql uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _ = rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3877:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3880:3:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+3627, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3883:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3884:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3885:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3887:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3888:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3889:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3890:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3892:3:
	libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), valueId)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3893:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3894:3:
	if rc == int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3895:5:
		rc = int32(m_SQLITE_EMPTY)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3896:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3898:3:
	if rc != int32(m_SQLITE_ROW) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3899:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3901:3:
	**(**Ti64)(__ccgo_up(rowid)) = libsqlite3.Xsqlite3_column_int64(tls, **(**uintptr)(__ccgo_up(bp)), 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3902:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3903:3:
	if rc != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3904:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3907:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3909:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3910:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3911:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3914:5:
func Xvec0_result_id(tls *libc.TLS, p uintptr, context uintptr, rowid Ti64) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3914:71:
	var rc int32
	var _ /* valueId at bp+0 */ uintptr
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3915:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(p)).FpkIsText != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3916:5:
		libsqlite3.Xsqlite3_result_int64(tls, context, rowid)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3917:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3920:7:
	rc = Xvec0_get_id_value_from_rowid(tls, p, rowid, bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3921:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3922:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3924:3:
	if !(**(**uintptr)(__ccgo_up(bp)) != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3925:5:
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3927:5:
		libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3928:5:
		libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp)))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3930:3:
	return m_SQLITE_OK
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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3945:5:
func Xvec0_get_vector_data(tls *libc.TLS, pVtab uintptr, rowid Ti64, vector_column_idx int32, outVector uintptr, outVectorSize uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3946:64:
	var blobOffset, brc, rc int32
	var buf, p uintptr
	var size Tsize_t
	var _ /* chunk_id at bp+0 */ Ti64
	var _ /* chunk_offset at bp+8 */ Ti64
	var _ /* vectorBlob at bp+16 */ uintptr
	_, _, _, _, _, _ = blobOffset, brc, buf, p, rc, size
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3947:13:
	p = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3952:8:
	buf = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3954:16:
	**(**uintptr)(__ccgo_up(bp + 16)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3955:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3958:3:
	rc = Xvec0_get_chunk_position(tls, pVtab, rowid, libc.UintptrFromInt32(0), bp, bp+8)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3959:3:
	if rc == int32(m_SQLITE_EMPTY) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3960:5:
		Xvtab_set_error(tls, pVtab, __ccgo_ts+3675, libc.VaList(bp+32, rowid))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3961:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3963:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3964:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3967:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(vector_column_idx)*8)), __ccgo_ts+3712, **(**Ti64)(__ccgo_up(bp)), 0, bp+16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3971:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3972:5:
		Xvtab_set_error(tls, pVtab, __ccgo_ts+3720, libc.VaList(bp+32, rowid))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3975:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3976:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3979:3:
	size = Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_column_idx)*32)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3980:3:
	blobOffset = libc.Int32FromUint64(libc.Uint64FromInt64(**(**Ti64)(__ccgo_up(bp + 8))) * size)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3982:3:
	buf = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(size))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3983:3:
	if !(buf != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3984:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3985:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3988:3:
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), buf, libc.Int32FromUint64(size), blobOffset)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3989:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3990:5:
		libsqlite3.Xsqlite3_free(tls, buf)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3991:5:
		buf = libc.UintptrFromInt32(0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3992:5:
		Xvtab_set_error(tls, pVtab, __ccgo_ts+3778, libc.VaList(bp+32, rowid))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3996:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:3997:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4000:3:
	**(**uintptr)(__ccgo_up(outVector)) = buf
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4001:3:
	if outVectorSize != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4002:5:
		**(**int32)(__ccgo_up(outVectorSize)) = libc.Int32FromUint64(size)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4004:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4006:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4007:3:
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 16)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4008:3:
	if rc == m_SQLITE_OK && brc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4009:5:
		Xvtab_set_error(tls, p, __ccgo_ts+3841, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4012:5:
		return brc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4015:3:
	return rc
}

// C documentation
//
//	/**
//	 * @brief Retrieve the sqlite3_value of the i'th partition value for the given row.
//	 *
//	 * @param pVtab - the vec0_vtab in questions
//	 * @param rowid - rowid of target row
//	 * @param partition_idx - which partition column to retrieve
//	 * @param outValue - output sqlite3_value
//	 * @return int - SQLITE_OK on success, otherwise error code
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4027:5:
func Xvec0_get_partition_value_for_rowid(tls *libc.TLS, pVtab uintptr, rowid Ti64, partition_idx int32, outValue uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4027:115:
	var rc int32
	var zSql uintptr
	var _ /* chunk_id at bp+0 */ Ti64
	var _ /* chunk_offset at bp+8 */ Ti64
	var _ /* stmt at bp+16 */ uintptr
	_, _ = rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4031:3:
	rc = Xvec0_get_chunk_position(tls, pVtab, rowid, libc.UintptrFromInt32(0), bp, bp+8)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4032:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4033:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4035:16:
	**(**uintptr)(__ccgo_up(bp + 16)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4036:8:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+3933, libc.VaList(bp+32, partition_idx, (*Tvec0_vtab)(unsafe.Pointer(pVtab)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pVtab)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4037:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4038:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4040:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(pVtab)).Fdb, zSql, -int32(1), bp+16, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4041:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4042:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4043:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4045:3:
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 16)), int32(1), **(**Ti64)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4046:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 16)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4047:3:
	if rc != int32(m_SQLITE_ROW) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4048:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4049:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4051:3:
	**(**uintptr)(__ccgo_up(outValue)) = libsqlite3.Xsqlite3_value_dup(tls, libsqlite3.Xsqlite3_column_value(tls, **(**uintptr)(__ccgo_up(bp + 16)), 0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4052:3:
	if !(**(**uintptr)(__ccgo_up(outValue)) != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4053:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4054:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4056:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4058:3:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4059:5:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 16)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4060:5:
	return rc
}

// C documentation
//
//	/**
//	 * @brief Get the value of an auxiliary column for the given rowid
//	 *
//	 * @param pVtab vec0_vtab
//	 * @param rowid the rowid of the row to lookup
//	 * @param auxiliary_idx aux index of the column we care about
//	 * @param outValue Output sqlite3_value to store
//	 * @return int SQLITE_OK on success, error code otherwise
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4073:5:
func Xvec0_get_auxiliary_value_for_rowid(tls *libc.TLS, pVtab uintptr, rowid Ti64, auxiliary_idx int32, outValue uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4073:115:
	var rc int32
	var zSql uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _ = rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4075:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4076:8:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+3995, libc.VaList(bp+16, auxiliary_idx, (*Tvec0_vtab)(unsafe.Pointer(pVtab)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pVtab)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4077:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4078:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4080:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(pVtab)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4081:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4082:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4083:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4085:3:
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4086:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4087:3:
	if rc != int32(m_SQLITE_ROW) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4088:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4089:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4091:3:
	**(**uintptr)(__ccgo_up(outValue)) = libsqlite3.Xsqlite3_value_dup(tls, libsqlite3.Xsqlite3_column_value(tls, **(**uintptr)(__ccgo_up(bp)), 0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4092:3:
	if !(**(**uintptr)(__ccgo_up(outValue)) != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4093:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4094:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4096:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4098:3:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4099:5:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4100:5:
	return rc
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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4113:5:
func Xvec0_result_metadata_value_for_rowid(tls *libc.TLS, p uintptr, rowid Ti64, metadata_idx int32, context uintptr) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4113:112:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4117:3:
	rc = Xvec0_get_chunk_position(tls, p, rowid, libc.UintptrFromInt32(0), bp, bp+8)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4118:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4119:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4122:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 480 + uintptr(metadata_idx)*8)), __ccgo_ts+4053, **(**Ti64)(__ccgo_up(bp)), 0, bp+16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4123:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4124:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4127:3:
	switch (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4128:5:
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4130:7:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), bp+24, int32(1), int32(**(**Ti64)(__ccgo_up(bp + 8))/int64(m_CHAR_BIT)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4131:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4132:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4134:11:
		value = libc.Int32FromUint8(**(**Tu8)(__ccgo_up(bp + 24))) >> (**(**Ti64)(__ccgo_up(bp + 8)) % int64(m_CHAR_BIT)) & int32(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4135:7:
		libsqlite3.Xsqlite3_result_int(tls, context, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4136:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4138:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4140:7:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), bp+32, int32(8), libc.Int32FromUint64(libc.Uint64FromInt64(**(**Ti64)(__ccgo_up(bp + 8)))*uint64(8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4141:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4142:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4144:7:
		libsqlite3.Xsqlite3_result_int64(tls, context, **(**Ti64)(__ccgo_up(bp + 32)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4145:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4147:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4149:7:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), bp+40, int32(8), libc.Int32FromUint64(libc.Uint64FromInt64(**(**Ti64)(__ccgo_up(bp + 8)))*uint64(8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4150:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4151:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4153:7:
		libsqlite3.Xsqlite3_result_double(tls, context, **(**float64)(__ccgo_up(bp + 40)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4154:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4156:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4158:7:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 16)), bp+48, int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH), int32(**(**Ti64)(__ccgo_up(bp + 8))*int64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4159:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4160:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4162:11:
		length = **(**int32)(__ccgo_up(bp + 48))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4163:7:
		if length <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4164:9:
			libsqlite3.Xsqlite3_result_text(tls, context, bp+48+libc.UintptrFromInt32(4), length, uintptr(-libc.Int32FromInt32(1)))
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4168:20:
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+4058, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4169:9:
			if !(zSql != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4170:11:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4171:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4173:9:
			rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+64, libc.UintptrFromInt32(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4174:9:
			libsqlite3.Xsqlite3_free(tls, zSql)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4175:9:
			if rc != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4176:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4178:9:
			libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 64)), int32(1), rowid)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4179:9:
			rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 64)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4180:9:
			if rc != int32(m_SQLITE_ROW) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4181:11:
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 64)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4182:11:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4183:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4185:9:
			libsqlite3.Xsqlite3_result_value(tls, context, libsqlite3.Xsqlite3_column_value(tls, **(**uintptr)(__ccgo_up(bp + 64)), 0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4186:9:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 64)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4187:9:
			rc = m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4189:7:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4192:3:
	goto done
done:
	;
	// blobValue is read-only, will not fail on close
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4194:5:
	libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 16)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4195:5:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4199:5:
func Xvec0_get_latest_chunk_rowid(tls *libc.TLS, p uintptr, chunk_rowid uintptr, partitionKeyValues uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4199:102:
	var i, i1, rc int32
	var s, zSql uintptr
	_, _, _, _, _ = i, i1, rc, s, zSql
	// lazy initialize stmtLatestChunk when needed. May be cleared during xSync()
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4203:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4204:5:
		if (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns > 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4205:19:
			s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4206:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+4118, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4209:7:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4209:15:
			i = 0
			for {
				if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4210:9:
				if i != 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4211:11:
					libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+4165)
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4213:9:
				libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+4171, libc.VaList(bp+8, i))
				goto _1
			_1:
				;
				i = i + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4215:7:
			zSql = libsqlite3.Xsqlite3_str_finish(tls, s)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4217:7:
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+4191, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4221:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4222:7:
			rc = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4223:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4225:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), p+1992, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4226:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4227:5:
		if rc != m_SQLITE_OK {
			// IMP: V21406_05476
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4229:7:
			Xvtab_set_error(tls, p, __ccgo_ts+4231, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4231:7:
			goto cleanup
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4235:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4235:11:
	i1 = 0
	for {
		if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4236:5:
		libsqlite3.Xsqlite3_bind_value(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk, i1+int32(1), **(**uintptr)(__ccgo_up(partitionKeyValues + uintptr(i1)*8)))
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4239:3:
	rc = libsqlite3.Xsqlite3_step(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4240:3:
	if rc != int32(m_SQLITE_ROW) {
		// IMP: V31559_15629
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4242:5:
		Xvtab_set_error(tls, p, __ccgo_ts+4304, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4243:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4244:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4246:3:
	if libsqlite3.Xsqlite3_column_type(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk, 0) == int32(m_SQLITE_NULL) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4247:5:
		rc = int32(m_SQLITE_EMPTY)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4248:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4250:3:
	**(**Ti64)(__ccgo_up(chunk_rowid)) = libsqlite3.Xsqlite3_column_int64(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk, 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4251:3:
	rc = libsqlite3.Xsqlite3_step(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4252:3:
	if rc != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4253:5:
		Xvtab_set_error(tls, p, __ccgo_ts+4359, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4258:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4260:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4262:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4263:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4264:5:
		libsqlite3.Xsqlite3_reset(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4265:5:
		libsqlite3.Xsqlite3_clear_bindings(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4267:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4270:5:
func Xvec0_rowids_insert_rowid(tls *libc.TLS, p uintptr, rowid Ti64) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4270:55:
	var entered, rc int32
	var zSql uintptr
	_, _, _ = entered, rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4271:7:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4272:7:
	entered = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4273:3:
	_ = entered // temporary
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4274:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4275:16:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+4510, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4279:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4280:7:
			rc = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4281:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4283:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), p+2000, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4284:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4285:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4286:7:
			Xvtab_set_error(tls, p, __ccgo_ts+4557, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4288:7:
			goto cleanup
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4298:3:
	libsqlite3.Xsqlite3_bind_int64(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid, int32(1), rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4299:3:
	rc = libsqlite3.Xsqlite3_step(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4301:3:
	if rc != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4302:5:
		if libsqlite3.Xsqlite3_extended_errcode(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb) == libc.Int32FromInt32(m_SQLITE_CONSTRAINT)|libc.Int32FromInt32(6)<<libc.Int32FromInt32(8) {
			// IMP: V17090_01160
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4304:7:
			Xvtab_set_error(tls, p, __ccgo_ts+4631, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		} else {
			// IMP: V04679_21517
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4308:7:
			Xvtab_set_error(tls, p, __ccgo_ts+4674, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, libsqlite3.Xsqlite3_db_handle(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId))))
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4312:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4313:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4316:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4318:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4319:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4320:5:
		libsqlite3.Xsqlite3_reset(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4321:5:
		libsqlite3.Xsqlite3_clear_bindings(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4329:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4332:5:
func Xvec0_rowids_insert_id(tls *libc.TLS, p uintptr, idValue uintptr, rowid uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4332:77:
	var entered, rc int32
	var zSql uintptr
	_, _, _ = entered, rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4333:7:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4334:7:
	entered = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4335:3:
	_ = entered // temporary
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4336:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4337:16:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+4725, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4341:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4342:7:
			rc = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4343:7:
			goto complete
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4345:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), p+2008, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4346:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4347:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4348:7:
			Xvtab_set_error(tls, p, __ccgo_ts+4769, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4350:7:
			goto complete
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4361:3:
	if idValue != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4362:5:
		libsqlite3.Xsqlite3_bind_value(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId, int32(1), idValue)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4364:3:
	rc = libsqlite3.Xsqlite3_step(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4366:3:
	if rc != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4367:5:
		if libsqlite3.Xsqlite3_extended_errcode(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb) == libc.Int32FromInt32(m_SQLITE_CONSTRAINT)|libc.Int32FromInt32(8)<<libc.Int32FromInt32(8) {
			// IMP: V20497_04568
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4369:7:
			Xvtab_set_error(tls, p, __ccgo_ts+4631, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		} else {
			// IMP: V24016_08086
			// IMP: V15177_32015
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4374:7:
			Xvtab_set_error(tls, p, __ccgo_ts+4846, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, libsqlite3.Xsqlite3_db_handle(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId))))
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4378:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4379:5:
		goto complete
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4382:3:
	**(**Ti64)(__ccgo_up(rowid)) = libsqlite3.Xsqlite3_last_insert_rowid(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4383:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4385:1:
	goto complete
complete:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4386:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4387:5:
		libsqlite3.Xsqlite3_reset(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4388:5:
		libsqlite3.Xsqlite3_clear_bindings(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4396:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4399:5:
func Xvec0_metadata_chunk_size(tls *libc.TLS, kind Tvec0_metadata_column_kind, chunk_size int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4399:78:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4400:3:
	switch kind {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4401:5:
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4402:7:
		return chunk_size / int32(8)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4403:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4404:7:
		return libc.Int32FromUint64(libc.Uint64FromInt32(chunk_size) * uint64(8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4405:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4406:7:
		return libc.Int32FromUint64(libc.Uint64FromInt32(chunk_size) * uint64(8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4407:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4408:7:
		return chunk_size * int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4410:3:
	return 0
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4413:5:
func Xvec0_rowids_update_position(tls *libc.TLS, p uintptr, rowid Ti64, chunk_rowid Ti64, chunk_offset Ti64) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4414:51:
	var rc int32
	var zSql uintptr
	_, _ = rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4415:7:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4417:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4418:16:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+4894, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4422:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4423:7:
			rc = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4424:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4426:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), p+2016, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4427:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4428:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4429:7:
			Xvtab_set_error(tls, p, __ccgo_ts+4970, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4431:7:
			goto cleanup
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4435:3:
	libsqlite3.Xsqlite3_bind_int64(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition, int32(1), chunk_rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4436:3:
	libsqlite3.Xsqlite3_bind_int64(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition, int32(2), chunk_offset)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4437:3:
	libsqlite3.Xsqlite3_bind_int64(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition, int32(3), rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4439:3:
	rc = libsqlite3.Xsqlite3_step(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4440:3:
	if rc != int32(m_SQLITE_DONE) {
		// IMP: V21925_05995
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4442:5:
		Xvtab_set_error(tls, p, __ccgo_ts+5053, libc.VaList(bp+8, rowid, chunk_rowid, chunk_offset))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4447:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4448:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4450:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4452:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4453:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4454:5:
		libsqlite3.Xsqlite3_reset(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4455:5:
		libsqlite3.Xsqlite3_clear_bindings(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4458:3:
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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4474:5:
func Xvec0_new_chunk(tls *libc.TLS, p uintptr, partitionKeyValues uintptr, chunk_rowid uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4474:89:
	var failed, i, i1, i2, i3, i4, metadata_column_idx, rc, vector_column_idx int32
	var rowid, vectorsSize Ti64
	var s, zSql uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = failed, i, i1, i2, i3, i4, metadata_column_idx, rc, rowid, s, vector_column_idx, vectorsSize, zSql
	// Step 1: Insert a new row in _chunks, capture that new rowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4481:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns > 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4482:17:
		s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4483:5:
		libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+5165, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4484:5:
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5194)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4485:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4485:13:
		i = 0
		for {
			if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4486:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+5218, libc.VaList(bp+16, i))
			goto _1
		_1:
			;
			i = i + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4488:5:
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5234)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4489:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4489:13:
		i1 = 0
		for {
			if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4490:7:
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5252)
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4492:5:
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5256)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4494:5:
		zSql = libsqlite3.Xsqlite3_str_finish(tls, s)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4496:5:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5258, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4502:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4503:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4505:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4506:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4507:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4508:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4509:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4518:3:
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)) // size
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4519:3:
	libsqlite3.Xsqlite3_bind_zeroblob(tls, **(**uintptr)(__ccgo_up(bp)), int32(2), (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m_CHAR_BIT)) // validity bitmap
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4520:3:
	libsqlite3.Xsqlite3_bind_zeroblob(tls, **(**uintptr)(__ccgo_up(bp)), int32(3), libc.Int32FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8))) // rowids
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4522:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4522:11:
	i2 = 0
	for {
		if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumPartitionColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4523:5:
		libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(bp)), int32(4)+i2, **(**uintptr)(__ccgo_up(partitionKeyValues + uintptr(i2)*8)))
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4526:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4527:7:
	failed = libc.BoolInt32(rc != int32(m_SQLITE_DONE))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4528:3:
	rowid = libsqlite3.Xsqlite3_last_insert_rowid(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4534:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4535:3:
	if failed != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4536:5:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4556:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4556:12:
	i3 = 0
	for {
		if !(i3 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4557:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i3)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4558:7:
			goto _4
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4560:9:
		vector_column_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i3))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4561:9:
		vectorsSize = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))))
		// See SHADOW_TABLE_ROWID_QUIRK above for why _rowid_ and rowid are both set.
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4565:5:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5329, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, vector_column_idx))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4569:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4570:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4572:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4573:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4575:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4576:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4577:7:
			return rc
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4580:5:
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), rowid) // _rowid_ (internal SQLite rowid)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4581:5:
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(2), rowid) // rowid   (user-defined column)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4582:5:
		libsqlite3.Xsqlite3_bind_zeroblob64(tls, **(**uintptr)(__ccgo_up(bp)), int32(3), libc.Uint64FromInt64(vectorsSize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4584:5:
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4585:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4586:5:
		if rc != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4587:7:
			return rc
		}
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
	// Step 3: Create new metadata chunks for each metadata column
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4592:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4592:12:
	i4 = 0
	for {
		if !(i4 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4593:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i4)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4594:7:
			goto _5
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4596:9:
		metadata_column_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i4))))
		// See SHADOW_TABLE_ROWID_QUIRK above for why _rowid_ and rowid are both set.
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4598:5:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5410, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4602:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4603:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4605:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4606:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4608:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4609:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4610:7:
			return rc
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4613:5:
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), rowid) // _rowid_ (internal SQLite rowid)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4614:5:
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(2), rowid) // rowid   (user-defined column)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4615:5:
		libsqlite3.Xsqlite3_bind_zeroblob64(tls, **(**uintptr)(__ccgo_up(bp)), int32(3), libc.Uint64FromInt32(Xvec0_metadata_chunk_size(tls, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_column_idx)*24))).Fkind, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4617:5:
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4618:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4619:5:
		if rc != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4620:7:
			return rc
		}
		goto _5
	_5:
		;
		i4 = i4 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4625:3:
	if chunk_rowid != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4626:5:
		**(**Ti64)(__ccgo_up(chunk_rowid)) = rowid
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4629:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4632:1:
type Tvec0_query_fullscan_data = struct {
	Frowids_stmt uintptr
	Fdone        Ti8
}

type vec0_query_fullscan_data = Tvec0_query_fullscan_data

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4636:6:
func Xvec0_query_fullscan_data_clear(tls *libc.TLS, fullscan_data uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4637:53:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4638:3:
	if !(fullscan_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4639:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4641:3:
	if (*Tvec0_query_fullscan_data)(unsafe.Pointer(fullscan_data)).Frowids_stmt != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4642:5:
		libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_query_fullscan_data)(unsafe.Pointer(fullscan_data)).Frowids_stmt)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4643:5:
		(*Tvec0_query_fullscan_data)(unsafe.Pointer(fullscan_data)).Frowids_stmt = libc.UintptrFromInt32(0)
	}
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4647:1:
type Tvec0_query_knn_data = struct {
	Fk           Ti64
	Fk_used      Ti64
	Frowids      uintptr
	Fdistances   uintptr
	Fcurrent_idx Ti64
}

type vec0_query_knn_data = Tvec0_query_knn_data

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4656:6:
func Xvec0_query_knn_data_clear(tls *libc.TLS, knn_data uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4656:70:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4657:3:
	if !(knn_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4658:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4660:3:
	if (*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Frowids != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4661:5:
		libsqlite3.Xsqlite3_free(tls, (*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Frowids)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4662:5:
		(*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Frowids = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4664:3:
	if (*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4665:5:
		libsqlite3.Xsqlite3_free(tls, (*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4666:5:
		(*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances = libc.UintptrFromInt32(0)
	}
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4670:1:
type Tvec0_query_point_data = struct {
	Frowid   Ti64
	Fvectors [16]uintptr
	Fdone    int32
}

type vec0_query_point_data = Tvec0_query_point_data

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4675:6:
func Xvec0_query_point_data_clear(tls *libc.TLS, point_data uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4675:76:
	var i int32
	_ = i
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4676:3:
	if !(point_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4677:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4678:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4678:12:
	i = 0
	for {
		if !(i < int32(m_VEC0_MAX_VECTOR_COLUMNS)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4679:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(point_data + 8 + uintptr(i)*8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4680:5:
		**(**uintptr)(__ccgo_up(point_data + 8 + uintptr(i)*8)) = libc.UintptrFromInt32(0)
		goto _1
	_1:
		;
		i = i + 1
	}
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4684:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4690:3:
type Tvec0_query_plan = int32

type vec0_query_plan = Tvec0_query_plan

const
// If any values are updated, please update the ARCHITECTURE.md docs accordingly!

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4687:2:
_VEC0_QUERY_PLAN_FULLSCAN = 49
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4688:2:
_VEC0_QUERY_PLAN_POINT = 50
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4689:2:
_VEC0_QUERY_PLAN_KNN = 51

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4692:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4692:28:
type Tvec0_cursor = struct {
	Fbase          Tsqlite3_vtab_cursor
	Fquery_plan    Tvec0_query_plan
	Ffullscan_data uintptr
	Fknn_data      uintptr
	Fpoint_data    uintptr
}

type vec0_cursor = Tvec0_cursor

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4702:6:
func Xvec0_cursor_clear(tls *libc.TLS, pCur uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4702:43:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4703:3:
	if (*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4704:5:
		Xvec0_query_fullscan_data_clear(tls, (*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4705:5:
		libsqlite3.Xsqlite3_free(tls, (*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4706:5:
		(*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4708:3:
	if (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4709:5:
		Xvec0_query_knn_data_clear(tls, (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4710:5:
		libsqlite3.Xsqlite3_free(tls, (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4711:5:
		(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4713:3:
	if (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4714:5:
		Xvec0_query_point_data_clear(tls, (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4715:5:
		libsqlite3.Xsqlite3_free(tls, (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4716:5:
		(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data = libc.UintptrFromInt32(0)
	}
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4721:12:
func _vec0_init(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr, isCreate uint8) (r int32) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4722:74:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4723:3:
	_ = pAux
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4728:3:
	pNew = libsqlite3.Xsqlite3_malloc(tls, int32(2032))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4729:3:
	if pNew == uintptr(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4730:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4731:3:
	libc.Xmemset(tls, pNew, 0, uint64(2032))
	// Declared chunk_size=N for entire table.
	// -1 to use the defualt, otherwise will get re-assigned on `chunk_size=N`
	// option
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4736:7:
	chunk_size = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4737:7:
	numVectorColumns = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4738:7:
	numPartitionColumns = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4739:7:
	numAuxiliaryColumns = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4740:7:
	numMetadataColumns = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4741:7:
	user_column_idx = 0
	// track if a "primary key" column is defined
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4744:8:
	pkColumnName = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4746:7:
	pkColumnType = int32(m_SQLITE_INTEGER)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4748:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4748:12:
	i = int32(3)
	for {
		if !(i < argc) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4753:10:
		**(**uintptr)(__ccgo_up(bp + 104)) = libc.UintptrFromInt32(0)
		// Scenario #1: Constructor argument is a vector column definition, ie `foo float[1024]`
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4758:5:
		rc = Xvec0_parse_vector_column(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), libc.Int32FromUint64(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4759:5:
		if rc == int32(m_SQLITE_ERROR) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4760:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5489, libc.VaList(bp+176, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4762:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4764:5:
		if rc == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4765:7:
			if numVectorColumns >= int32(m_VEC0_MAX_VECTOR_COLUMNS) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4766:9:
				libsqlite3.Xsqlite3_free(tls, (**(**TVectorColumnDefinition)(__ccgo_up(bp))).Fname)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4767:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5548, libc.VaList(bp+176, int32(m_VEC0_MAX_VECTOR_COLUMNS)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4770:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4773:7:
			if (**(**TVectorColumnDefinition)(__ccgo_up(bp))).Fdimensions > uint64(m_SQLITE_VEC_VEC0_MAX_DIMENSIONS) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4774:9:
				libsqlite3.Xsqlite3_free(tls, (**(**TVectorColumnDefinition)(__ccgo_up(bp))).Fname)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4775:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5617, libc.VaList(bp+176, libc.Int64FromUint64((**(**TVectorColumnDefinition)(__ccgo_up(bp))).Fdimensions), int32(m_SQLITE_VEC_VEC0_MAX_DIMENSIONS)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4779:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4781:7:
			**(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(user_column_idx)*4)) = int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4782:7:
			**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(user_column_idx))) = libc.Uint8FromInt32(numVectorColumns)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4783:7:
			libc.Xmemcpy(tls, pNew+608+uintptr(numVectorColumns)*32, bp, uint64(32))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4784:7:
			numVectorColumns = numVectorColumns + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4785:7:
			(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumVectorColumns = numVectorColumns
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4786:7:
			user_column_idx = user_column_idx + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4788:7:
			goto _1
		}
		// Scenario #2: Constructor argument is a partition key column definition, ie `user_id text partition key`
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4792:5:
		rc = Xvec0_parse_partition_key_definition(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), libc.Int32FromUint64(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+104, bp+112, bp+116)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4794:5:
		if rc == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4795:7:
			if numPartitionColumns >= int32(m_VEC0_MAX_PARTITION_COLUMNS) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4796:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5707, libc.VaList(bp+176, int32(m_VEC0_MAX_PARTITION_COLUMNS)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4800:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4802:7:
			(**(**TVec0PartitionColumnDefinition)(__ccgo_up(bp + 32))).Ftype1 = **(**int32)(__ccgo_up(bp + 116))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4803:7:
			(**(**TVec0PartitionColumnDefinition)(__ccgo_up(bp + 32))).Fname_length = **(**int32)(__ccgo_up(bp + 112))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4804:7:
			(**(**TVec0PartitionColumnDefinition)(__ccgo_up(bp + 32))).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+176, **(**int32)(__ccgo_up(bp + 112)), **(**uintptr)(__ccgo_up(bp + 104))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4805:7:
			if !((**(**TVec0PartitionColumnDefinition)(__ccgo_up(bp + 32))).Fname != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4806:9:
				rc = int32(m_SQLITE_NOMEM)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4807:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4810:7:
			**(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(user_column_idx)*4)) = int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4811:7:
			**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(user_column_idx))) = libc.Uint8FromInt32(numPartitionColumns)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4812:7:
			libc.Xmemcpy(tls, pNew+1120+uintptr(numPartitionColumns)*24, bp+32, uint64(24))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4813:7:
			numPartitionColumns = numPartitionColumns + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4814:7:
			(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumPartitionColumns = numPartitionColumns
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4815:7:
			user_column_idx = user_column_idx + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4816:7:
			goto _1
		}
		// Scenario #3: Constructor argument is a primary key column definition, ie `article_id text primary key`
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4820:5:
		rc = Xvec0_parse_primary_key_definition(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), libc.Int32FromUint64(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+104, bp+112, bp+116)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4822:5:
		if rc == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4823:7:
			if pkColumnName != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4824:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5780, libc.VaList(bp+176, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4829:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4831:7:
			pkColumnName = **(**uintptr)(__ccgo_up(bp + 104))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4832:7:
			pkColumnNameLength = **(**int32)(__ccgo_up(bp + 112))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4833:7:
			pkColumnType = **(**int32)(__ccgo_up(bp + 116))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4834:7:
			goto _1
		}
		// Scenario #4: Constructor argument is a auxiliary column definition, ie `+contents text`
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4838:5:
		rc = Xvec0_parse_auxiliary_column_definition(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), libc.Int32FromUint64(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+104, bp+112, bp+116)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4840:5:
		if rc == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4841:7:
			if numAuxiliaryColumns >= int32(m_VEC0_MAX_AUXILIARY_COLUMNS) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4842:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5901, libc.VaList(bp+176, int32(m_VEC0_MAX_AUXILIARY_COLUMNS)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4846:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4848:7:
			(**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(bp + 56))).Ftype1 = **(**int32)(__ccgo_up(bp + 116))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4849:7:
			(**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(bp + 56))).Fname_length = **(**int32)(__ccgo_up(bp + 112))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4850:7:
			(**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(bp + 56))).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+176, **(**int32)(__ccgo_up(bp + 112)), **(**uintptr)(__ccgo_up(bp + 104))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4851:7:
			if !((**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(bp + 56))).Fname != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4852:9:
				rc = int32(m_SQLITE_NOMEM)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4853:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4856:7:
			**(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(user_column_idx)*4)) = int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4857:7:
			**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(user_column_idx))) = libc.Uint8FromInt32(numAuxiliaryColumns)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4858:7:
			libc.Xmemcpy(tls, pNew+1216+uintptr(numAuxiliaryColumns)*24, bp+56, uint64(24))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4859:7:
			numAuxiliaryColumns = numAuxiliaryColumns + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4860:7:
			(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumAuxiliaryColumns = numAuxiliaryColumns
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4861:7:
			user_column_idx = user_column_idx + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4862:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4866:5:
		rc = Xvec0_parse_metadata_column_definition(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), libc.Int32FromUint64(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+104, bp+112, bp+120)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4868:5:
		if rc == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4869:7:
			if numMetadataColumns >= int32(m_VEC0_MAX_METADATA_COLUMNS) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4870:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+5970, libc.VaList(bp+176, int32(m_VEC0_MAX_METADATA_COLUMNS)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4874:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4876:7:
			(**(**TVec0MetadataColumnDefinition)(__ccgo_up(bp + 80))).Fkind = **(**Tvec0_metadata_column_kind)(__ccgo_up(bp + 120))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4877:7:
			(**(**TVec0MetadataColumnDefinition)(__ccgo_up(bp + 80))).Fname_length = **(**int32)(__ccgo_up(bp + 112))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4878:7:
			(**(**TVec0MetadataColumnDefinition)(__ccgo_up(bp + 80))).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+176, **(**int32)(__ccgo_up(bp + 112)), **(**uintptr)(__ccgo_up(bp + 104))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4879:7:
			if !((**(**TVec0MetadataColumnDefinition)(__ccgo_up(bp + 80))).Fname != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4880:9:
				rc = int32(m_SQLITE_NOMEM)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4881:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4884:7:
			**(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(user_column_idx)*4)) = int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4885:7:
			**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(user_column_idx))) = libc.Uint8FromInt32(numMetadataColumns)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4886:7:
			libc.Xmemcpy(tls, pNew+1600+uintptr(numMetadataColumns)*24, bp+80, uint64(24))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4887:7:
			numMetadataColumns = numMetadataColumns + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4888:7:
			(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumMetadataColumns = numMetadataColumns
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4889:7:
			user_column_idx = user_column_idx + 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4890:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4898:5:
		rc = Xvec0_parse_table_option(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), libc.Int32FromUint64(libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)))), bp+128, bp+144, bp+136, bp+148)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4900:5:
		if rc == int32(m_SQLITE_ERROR) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4901:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6038, libc.VaList(bp+176, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4903:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4905:5:
		if rc == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4906:7:
			if libsqlite3.Xsqlite3_strnicmp(tls, **(**uintptr)(__ccgo_up(bp + 128)), __ccgo_ts+6096, **(**int32)(__ccgo_up(bp + 144))) == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4907:9:
				chunk_size = libc.Xatoi(tls, **(**uintptr)(__ccgo_up(bp + 136)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4908:9:
				if chunk_size <= 0 {
					// IMP: V01931_18769
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4910:11:
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6107, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4913:11:
					goto error
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4915:9:
				if chunk_size%int32(8) != 0 {
					// IMP: V14110_30948
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4917:11:
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6178, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4919:11:
					goto error
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4922:9:
				if chunk_size > int32(m_SQLITE_VEC_CHUNK_SIZE_MAX) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4923:11:
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6236, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4925:11:
					goto error
				}
			} else {
				// IMP: V27642_11712
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4929:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6281, libc.VaList(bp+176, **(**int32)(__ccgo_up(bp + 144)), **(**uintptr)(__ccgo_up(bp + 128))))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4931:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4933:7:
			goto _1
		}
		// Scenario #5: Unknown constructor argument
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4937:5:
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6332, libc.VaList(bp+176, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4939:5:
		goto error
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4942:3:
	if chunk_size < 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4943:5:
		chunk_size = int32(1024)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4946:3:
	if numVectorColumns <= 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4947:5:
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6377, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4949:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4952:15:
	createStr = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4953:3:
	libsqlite3.Xsqlite3_str_appendall(tls, createStr, __ccgo_ts+6440)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4954:3:
	if pkColumnName != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4955:5:
		libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6456, libc.VaList(bp+176, pkColumnNameLength, pkColumnName))
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4958:5:
		libsqlite3.Xsqlite3_str_appendall(tls, createStr, __ccgo_ts+6477)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4960:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4960:12:
	i1 = 0
	for {
		if !(i1 < numVectorColumns+numPartitionColumns+numAuxiliaryColumns+numMetadataColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4961:5:
		switch **(**Tvec0_user_column_kind)(__ccgo_up(pNew + 88 + uintptr(i1)*4)) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4962:7:
		case int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4963:13:
			vector_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(i1))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4964:9:
			libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6485, libc.VaList(bp+176, (**(**TVectorColumnDefinition)(__ccgo_up(pNew + 608 + uintptr(vector_idx)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(pNew + 608 + uintptr(vector_idx)*32))).Fname))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4967:9:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4969:7:
			fallthrough
		case int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4970:13:
			partition_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(i1))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4971:9:
			libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6485, libc.VaList(bp+176, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(pNew + 1120 + uintptr(partition_idx)*24))).Fname_length, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(pNew + 1120 + uintptr(partition_idx)*24))).Fname))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4974:9:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4976:7:
			fallthrough
		case int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4977:13:
			auxiliary_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(i1))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4978:9:
			libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6485, libc.VaList(bp+176, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(pNew + 1216 + uintptr(auxiliary_idx)*24))).Fname_length, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(pNew + 1216 + uintptr(auxiliary_idx)*24))).Fname))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4981:9:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4983:7:
			fallthrough
		case int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4984:13:
			metadata_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(pNew + 296 + uintptr(i1))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4985:9:
			libsqlite3.Xsqlite3_str_appendf(tls, createStr, __ccgo_ts+6485, libc.VaList(bp+176, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pNew + 1600 + uintptr(metadata_idx)*24))).Fname_length, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pNew + 1600 + uintptr(metadata_idx)*24))).Fname))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4988:9:
			break
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4993:3:
	libsqlite3.Xsqlite3_str_appendall(tls, createStr, __ccgo_ts+6494)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4994:3:
	if pkColumnName != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4995:5:
		libsqlite3.Xsqlite3_str_appendall(tls, createStr, __ccgo_ts+6523)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4997:3:
	zSql = libsqlite3.Xsqlite3_str_finish(tls, createStr)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4998:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:4999:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5001:3:
	rc = libsqlite3.Xsqlite3_declare_vtab(tls, db, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5002:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5003:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5004:5:
		**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6538, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5007:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5010:14:
	schemaName = **(**uintptr)(__ccgo_up(argv + 1*8))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5011:14:
	tableName = **(**uintptr)(__ccgo_up(argv + 2*8))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5013:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).Fdb = db
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5014:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FpkIsText = libc.BoolInt32(pkColumnType == int32(m_SQLITE_TEXT))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5015:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6600, libc.VaList(bp+176, schemaName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5016:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5017:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5019:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6600, libc.VaList(bp+176, tableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5020:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5021:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5023:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FshadowRowidsName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6603, libc.VaList(bp+176, tableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5024:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(pNew)).FshadowRowidsName != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5025:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5027:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FshadowChunksName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6613, libc.VaList(bp+176, tableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5028:3:
	if !((*Tvec0_vtab)(unsafe.Pointer(pNew)).FshadowChunksName != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5029:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5031:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumVectorColumns = numVectorColumns
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5032:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumPartitionColumns = numPartitionColumns
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5033:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumAuxiliaryColumns = numAuxiliaryColumns
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5034:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumMetadataColumns = numMetadataColumns
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5036:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5036:12:
	i2 = 0
	for {
		if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumVectorColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5037:5:
		**(**uintptr)(__ccgo_up(pNew + 352 + uintptr(i2)*8)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6623, libc.VaList(bp+176, tableName, i2))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5039:5:
		if !(**(**uintptr)(__ccgo_up(pNew + 352 + uintptr(i2)*8)) != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5040:7:
			goto error
		}
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5043:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5043:12:
	i3 = 0
	for {
		if !(i3 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumMetadataColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5044:5:
		**(**uintptr)(__ccgo_up(pNew + 480 + uintptr(i3)*8)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6644, libc.VaList(bp+176, tableName, i3))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5046:5:
		if !(**(**uintptr)(__ccgo_up(pNew + 480 + uintptr(i3)*8)) != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5047:7:
			goto error
		}
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5050:3:
	(*Tvec0_vtab)(unsafe.Pointer(pNew)).Fchunk_size = chunk_size
	// if xCreate, then create the necessary shadow tables
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5053:3:
	if isCreate != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5057:10:
		zCreateInfo = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6666, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5058:5:
		if !(zCreateInfo != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5059:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5061:5:
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zCreateInfo, -int32(1), bp+152, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5063:5:
		libsqlite3.Xsqlite3_free(tls, zCreateInfo)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5064:5:
		if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
			// TODO(IMP)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5066:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5067:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6728, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5069:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5071:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5073:10:
		zSeedInfo = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6770, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5078:5:
		if !(zSeedInfo != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5079:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5081:5:
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSeedInfo, -int32(1), bp+152, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5082:5:
		libsqlite3.Xsqlite3_free(tls, zSeedInfo)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5083:5:
		if rc1 != m_SQLITE_OK {
			// TODO(IMP)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5085:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5086:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6856, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5088:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5090:5:
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(1), __ccgo_ts+6896, -int32(1), libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5091:5:
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(2), __ccgo_ts+6911, -int32(1), libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5092:5:
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(3), __ccgo_ts+6918, -int32(1), libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5093:5:
		libsqlite3.Xsqlite3_bind_int(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(4), m_SQLITE_VEC_VERSION_MAJOR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5094:5:
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(5), __ccgo_ts+6939, -int32(1), libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5095:5:
		libsqlite3.Xsqlite3_bind_int(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(6), int32(m_SQLITE_VEC_VERSION_MINOR))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5096:5:
		libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(7), __ccgo_ts+6960, -int32(1), libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5097:5:
		libsqlite3.Xsqlite3_bind_int(tls, **(**uintptr)(__ccgo_up(bp + 152)), int32(8), int32(m_SQLITE_VEC_VERSION_PATCH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5099:5:
		if libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
			// TODO(IMP)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5101:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5102:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6856, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5104:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5106:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
		// create the _chunks shadow table
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5111:10:
		zCreateShadowChunks = libc.UintptrFromInt32(0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5112:5:
		if (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumPartitionColumns != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5113:19:
			s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5114:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+6981, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5115:7:
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+7012)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5116:7:
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+7078)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5117:7:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5117:15:
			i4 = 0
			for {
				if !(i4 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumPartitionColumns) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5118:9:
				libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+7099, libc.VaList(bp+176, i4))
				goto _5
			_5:
				;
				i4 = i4 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5120:7:
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+7114)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5121:7:
			zCreateShadowChunks = libsqlite3.Xsqlite3_str_finish(tls, s)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5123:7:
			zCreateShadowChunks = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7161, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5126:5:
		if !(zCreateShadowChunks != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5127:9:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5129:5:
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zCreateShadowChunks, -int32(1), bp+152, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5130:5:
		libsqlite3.Xsqlite3_free(tls, zCreateShadowChunks)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5131:5:
		if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
			// IMP: V17740_01811
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5133:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5134:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7302, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5136:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5138:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5142:5:
		if (*Tvec0_vtab)(unsafe.Pointer(pNew)).FpkIsText != 0 {
			// adds a "text unique not null" constraint to the id column
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5144:7:
			zCreateShadowRowids = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7346, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5147:7:
			zCreateShadowRowids = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7480, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5150:5:
		if !(zCreateShadowRowids != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5151:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5153:5:
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zCreateShadowRowids, -int32(1), bp+152, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5154:5:
		libsqlite3.Xsqlite3_free(tls, zCreateShadowRowids)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5155:5:
		if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
			// IMP: V11631_28470
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5157:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5158:7:
			**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7593, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5160:7:
			goto error
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5162:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5164:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5164:14:
		i5 = 0
		for {
			if !(i5 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumVectorColumns) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5165:12:
			zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7637, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName, i5))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5167:7:
			if !(zSql1 != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5168:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5170:7:
			rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSql1, -int32(1), bp+152, uintptr(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5171:7:
			libsqlite3.Xsqlite3_free(tls, zSql1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5172:7:
			if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
				// IMP: V25919_09989
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5174:9:
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5175:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7720, libc.VaList(bp+176, i5, libsqlite3.Xsqlite3_errmsg(tls, db)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5178:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5180:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			goto _6
		_6:
			;
			i5 = i5 + 1
		}
		// See SHADOW_TABLE_ROWID_QUIRK in vec0_new_chunk() — same "rowid PRIMARY KEY"
		// without INTEGER type issue applies here.
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5185:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5185:14:
		i6 = 0
		for {
			if !(i6 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumMetadataColumns) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5186:12:
			zSql2 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7775, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName, i6))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5188:7:
			if !(zSql2 != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5189:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5191:7:
			rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSql2, -int32(1), bp+152, uintptr(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5192:7:
			libsqlite3.Xsqlite3_free(tls, zSql2)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5193:7:
			if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5194:9:
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5195:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7857, libc.VaList(bp+176, i6, libsqlite3.Xsqlite3_errmsg(tls, db)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5198:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5200:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5202:7:
			if (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pNew + 1600 + uintptr(i6)*24))).Fkind == int32(_VEC0_METADATA_COLUMN_KIND_TEXT) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5203:14:
				zSql3 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7912, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName, i6))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5205:9:
				if !(zSql3 != 0) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5206:11:
					goto error
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5208:9:
				rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSql3, -int32(1), bp+152, uintptr(0))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5209:9:
				libsqlite3.Xsqlite3_free(tls, zSql3)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5210:9:
				if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 152))) != int32(m_SQLITE_DONE) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5211:11:
					libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5212:11:
					**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+7983, libc.VaList(bp+176, i6, libsqlite3.Xsqlite3_errmsg(tls, db)))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5215:11:
					goto error
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5217:9:
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 152)))
			}
			goto _7
		_7:
			;
			i6 = i6 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5222:5:
		if (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumAuxiliaryColumns > 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5224:19:
			s1 = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5225:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s1, __ccgo_ts+8037, libc.VaList(bp+176, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(pNew)).FtableName))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5226:7:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5226:15:
			i7 = 0
			for {
				if !(i7 < (*Tvec0_vtab)(unsafe.Pointer(pNew)).FnumAuxiliaryColumns) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5227:9:
				libsqlite3.Xsqlite3_str_appendf(tls, s1, __ccgo_ts+8098, libc.VaList(bp+176, i7))
				goto _8
			_8:
				;
				i7 = i7 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5229:7:
			libsqlite3.Xsqlite3_str_appendall(tls, s1, __ccgo_ts+5256)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5230:12:
			zSql4 = libsqlite3.Xsqlite3_str_finish(tls, s1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5231:7:
			if !(zSql4 != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5232:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5234:7:
			rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, db, zSql4, -int32(1), bp+160, libc.UintptrFromInt32(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5235:7:
			if rc1 != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 160))) != int32(m_SQLITE_DONE) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5236:9:
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 160)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5237:9:
				**(**uintptr)(__ccgo_up(pzErr)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8110, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, db)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5241:9:
				goto error
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5243:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 160)))
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5247:3:
	**(**uintptr)(__ccgo_up(ppVtab)) = pNew
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5248:3:
	return m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5250:1:
	goto error
error:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5251:3:
	Xvec0_free(tls, pNew)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5252:3:
	libsqlite3.Xsqlite3_free(tls, pNew)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5253:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5256:12:
func _vec0Create(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5258:37:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5259:3:
	return _vec0_init(tls, db, pAux, argc, argv, ppVtab, pzErr, libc.BoolUint8(m_true != 0))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5261:12:
func _vec0Connect(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5263:38:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5264:3:
	return _vec0_init(tls, db, pAux, argc, argv, ppVtab, pzErr, libc.BoolUint8(m_false != 0))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5267:12:
func _vec0Disconnect(tls *libc.TLS, pVtab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5267:48:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5268:13:
	p = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5269:3:
	Xvec0_free(tls, p)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5270:3:
	libsqlite3.Xsqlite3_free(tls, p)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5271:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5273:12:
func _vec0Destroy(tls *libc.TLS, pVtab uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5273:45:
	var i, i1, rc int32
	var p, zSql uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _, _, _, _ = i, i1, p, rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5274:13:
	p = pVtab
	// Free up any sqlite3_stmt, otherwise DROPs on those tables will fail
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5280:3:
	Xvec0_free_resources(tls, p)
	// TODO(test) later: can't evidence-of here, bc always gives "SQL logic error" instead of
	// provided error
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5284:3:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8154, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5286:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, uintptr(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5287:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5288:3:
	if rc != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp))) != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5289:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5290:5:
		Xvtab_set_error(tls, pVtab, __ccgo_ts+8182, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5291:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5293:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5295:3:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8217, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5297:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, uintptr(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5298:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5299:3:
	if rc != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp))) != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5300:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5301:5:
		Xvtab_set_error(tls, pVtab, __ccgo_ts+8243, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5302:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5304:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5306:3:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8276, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5308:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, uintptr(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5309:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5310:3:
	if rc != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp))) != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5311:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5312:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5314:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5316:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5316:12:
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5317:5:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8304, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5319:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5320:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5321:5:
		if rc != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp))) != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5322:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5323:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5325:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5328:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns > 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5329:5:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8325, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5330:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5331:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5332:5:
		if rc != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp))) != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5333:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5334:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5336:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5340:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5340:12:
	i1 = 0
	for {
		if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumMetadataColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5341:5:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8356, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, i1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5342:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, uintptr(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5343:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5344:5:
		if rc != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp))) != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5345:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5346:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5348:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5350:5:
		if (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(i1)*24))).Fkind == int32(_VEC0_METADATA_COLUMN_KIND_TEXT) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5351:7:
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+8396, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, i1))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5352:7:
			rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, uintptr(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5353:7:
			libsqlite3.Xsqlite3_free(tls, zSql)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5354:7:
			if rc != m_SQLITE_OK || libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp))) != int32(m_SQLITE_DONE) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5355:9:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5356:9:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5358:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5362:3:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5363:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5365:1:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5366:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5367:3:
	Xvec0_free(tls, p)
	// If there was an error
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5369:3:
	if rc == m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5370:5:
		libsqlite3.Xsqlite3_free(tls, p)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5372:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5375:12:
func _vec0Open(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5375:70:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5376:3:
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5378:3:
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(40))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5379:3:
	if pCur == uintptr(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5380:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5381:3:
	libc.Xmemset(tls, pCur, 0, uint64(40))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5382:3:
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5383:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5386:12:
func _vec0Close(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5386:48:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5387:15:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5388:3:
	Xvec0_cursor_clear(tls, pCur)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5389:3:
	libsqlite3.Xsqlite3_free(tls, pCur)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5390:3:
	return m_SQLITE_OK
}

// C documentation
//
//	// All the different type of "values" provided to argv/argc in vec0Filter.
//	// These enums denote the use and purpose of all of them.
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5395:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5414:3:
type Tvec0_idxstr_kind = int32

type vec0_idxstr_kind = Tvec0_idxstr_kind

const
// If any values are updated, please update the ARCHITECTURE.md docs accordingly!

// ~~~ KNN QUERIES ~~~ //

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5399:3:
_VEC0_IDXSTR_KIND_KNN_MATCH = 123
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5400:3:
_VEC0_IDXSTR_KIND_KNN_K = 125
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5401:3:
_VEC0_IDXSTR_KIND_KNN_ROWID_IN = 91
const
// argv[i] is a constraint on a PARTITON KEY column in a KNN query
//

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5404:3:
_VEC0_IDXSTR_KIND_KNN_PARTITON_CONSTRAINT = 93
const

// argv[i] is a constraint on the distance column in a KNN query

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5407:3:
_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT = 42
const

// ~~~ POINT QUERIES ~~~ //

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5410:3:
_VEC0_IDXSTR_KIND_POINT_ID = 33
const

// ~~~ ??? ~~~ //

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5413:3:
_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT = 38

// C documentation
//
//	// The different SQLITE_INDEX_CONSTRAINT values that vec0 partition key columns
//	// support, but as characters that fit nicely in idxstr.
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5418:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5438:3:
type Tvec0_partition_operator = int32

type vec0_partition_operator = Tvec0_partition_operator

const
// If any values are updated, please update the ARCHITECTURE.md docs accordingly!

// Equality constraint on a PARTITON KEY column, ex `user_id = 123`

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5422:3:
_VEC0_PARTITION_OPERATOR_EQ = 97
const

// "Greater than" constraint on a PARTITON KEY column, ex `year > 2024`

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5425:3:
_VEC0_PARTITION_OPERATOR_GT = 98
const

// "Less than or equal to" constraint on a PARTITON KEY column, ex `year <= 2024`

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5428:3:
_VEC0_PARTITION_OPERATOR_LE = 99
const

// "Less than" constraint on a PARTITON KEY column, ex `year < 2024`

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5431:3:
_VEC0_PARTITION_OPERATOR_LT = 100
const

// "Greater than or equal to" constraint on a PARTITON KEY column, ex `year >= 2024`

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5434:3:
_VEC0_PARTITION_OPERATOR_GE = 101
const

// "Not equal to" constraint on a PARTITON KEY column, ex `year != 2024`

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5437:3:
_VEC0_PARTITION_OPERATOR_NE = 102

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5439:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5447:3:
type Tvec0_metadata_operator = int32

type vec0_metadata_operator = Tvec0_metadata_operator

const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5440:3:
_VEC0_METADATA_OPERATOR_EQ = 97
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5441:3:
_VEC0_METADATA_OPERATOR_GT = 98
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5442:3:
_VEC0_METADATA_OPERATOR_LE = 99
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5443:3:
_VEC0_METADATA_OPERATOR_LT = 100
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5444:3:
_VEC0_METADATA_OPERATOR_GE = 101
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5445:3:
_VEC0_METADATA_OPERATOR_NE = 102
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5446:3:
_VEC0_METADATA_OPERATOR_IN = 103

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5450:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5456:3:
type Tvec0_distance_constraint_operator = int32

type vec0_distance_constraint_operator = Tvec0_distance_constraint_operator

const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5452:3:
_VEC0_DISTANCE_CONSTRAINT_GT = 97
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5453:3:
_VEC0_DISTANCE_CONSTRAINT_GE = 98
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5454:3:
_VEC0_DISTANCE_CONSTRAINT_LT = 99
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5455:3:
_VEC0_DISTANCE_CONSTRAINT_LE = 100

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5458:12:
func _vec0BestIndex(tls *libc.TLS, pVTab uintptr, pIdxInfo uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5458:77:
	var argvIndex, hasAuxConstraint, i, i1, i2, i3, iColumn, iColumn1, iColumn2, iColumn3, iKTerm, iLimitTerm, iMatchTerm, iMatchVectorTerm, iRowidInTerm, iRowidTerm, metadata_idx, op, op1, op2, op3, partition_idx, rc, vtabIn1, v2 int32
	var idxStr, p uintptr
	var value, value1, value2 int8
	var vtabIn Tu8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = argvIndex, hasAuxConstraint, i, i1, i2, i3, iColumn, iColumn1, iColumn2, iColumn3, iKTerm, iLimitTerm, iMatchTerm, iMatchVectorTerm, iRowidInTerm, iRowidTerm, idxStr, metadata_idx, op, op1, op2, op3, p, partition_idx, rc, value, value1, value2, vtabIn, vtabIn1, v2
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5459:13:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5472:7:
	iMatchTerm = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5473:7:
	iMatchVectorTerm = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5474:7:
	iLimitTerm = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5475:7:
	iRowidTerm = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5476:7:
	iKTerm = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5477:7:
	iRowidInTerm = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5478:7:
	hasAuxConstraint = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5484:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5484:12:
	i = 0
	for {
		if !(i < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5485:8:
		vtabIn = uint8(0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5488:5:
		if libsqlite3.Xsqlite3_libversion_number(tls) >= int32(3038000) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5489:7:
			vtabIn = libc.Uint8FromInt32(libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, i, -int32(1)))
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5498:5:
		if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).Fusable != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5499:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5501:9:
		iColumn = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).FiColumn
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5502:9:
		op = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).Fop)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5504:5:
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5505:7:
			iLimitTerm = i
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5507:5:
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_MATCH) && Xvec0_column_idx_is_vector(tls, p, iColumn) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5509:7:
			if iMatchTerm > -int32(1) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5510:9:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+8434, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5512:9:
				return int32(m_SQLITE_ERROR)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5514:7:
			iMatchTerm = i
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5515:7:
			iMatchVectorTerm = Xvec0_column_idx_to_vector_idx(tls, p, iColumn)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5517:5:
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_EQ) && iColumn == m_VEC0_COLUMN_ID {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5518:7:
			if vtabIn != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5519:9:
				if iRowidInTerm != -int32(1) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5520:11:
					Xvtab_set_error(tls, pVTab, __ccgo_ts+8490, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5522:11:
					return int32(m_SQLITE_ERROR)
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5524:9:
				iRowidInTerm = i
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5527:9:
				iRowidTerm = i
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5530:5:
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_EQ) && iColumn == Xvec0_column_k_idx(tls, p) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5531:7:
			iKTerm = i
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5533:5:
		if op != int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) && op != int32(m_SQLITE_INDEX_CONSTRAINT_OFFSET) && Xvec0_column_idx_is_auxiliary(tls, p, iColumn) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5536:9:
			hasAuxConstraint = int32(1)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5540:15:
	idxStr = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5543:3:
	if iMatchTerm >= 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5544:5:
		if iLimitTerm < 0 && iKTerm < 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5545:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+8556, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5548:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5549:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5551:5:
		if iLimitTerm >= 0 && iKTerm >= 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5552:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+8619, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5553:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5554:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5557:5:
		if (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnOrderBy != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5558:7:
			if (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnOrderBy > int32(1) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5559:9:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+8666, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5561:9:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5562:7:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5564:7:
			if (**(**Tsqlite3_index_orderby)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaOrderBy))).FiColumn != Xvec0_column_distance_idx(tls, p) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5565:9:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+8738, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5568:9:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5569:7:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5571:7:
			if (**(**Tsqlite3_index_orderby)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaOrderBy))).Fdesc != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5572:9:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+8832, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5575:9:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5576:7:
				goto done
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5580:5:
		if hasAuxConstraint != 0 {
			// IMP: V25623_09693
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5582:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+8916, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5583:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5584:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5587:5:
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_QUERY_PLAN_KNN))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5589:9:
		argvIndex = int32(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5590:5:
		v2 = argvIndex
		argvIndex = argvIndex + 1
		(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iMatchTerm)*8))).FargvIndex = v2
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5591:5:
		(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iMatchTerm)*8))).Fomit = uint8(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5592:5:
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_MATCH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5593:5:
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(3), int8('_'))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5595:5:
		if iLimitTerm >= 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5596:7:
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iLimitTerm)*8))).FargvIndex = v2
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5597:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iLimitTerm)*8))).Fomit = uint8(1)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5599:7:
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iKTerm)*8))).FargvIndex = v2
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5600:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iKTerm)*8))).Fomit = uint8(1)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5602:5:
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_K))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5603:5:
		libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(3), int8('_'))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5606:5:
		if iRowidInTerm >= 0 {
			// already validated as  >= SQLite 3.38 bc iRowidInTerm is only >= 0 when
			// vtabIn == 1
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5609:7:
			libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, iRowidInTerm, int32(1))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5610:7:
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iRowidInTerm)*8))).FargvIndex = v2
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5611:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iRowidInTerm)*8))).Fomit = uint8(1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5612:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_ROWID_IN))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5613:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(3), int8('_'))
		}
		// find any PARTITION KEY column constraints
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5618:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5618:14:
		i1 = 0
		for {
			if !(i1 < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5619:7:
			if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i1)*12))).Fusable != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5620:9:
				goto _6
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5622:11:
			iColumn1 = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i1)*12))).FiColumn
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5623:11:
			op1 = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i1)*12))).Fop)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5624:7:
			if op1 == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) || op1 == int32(m_SQLITE_INDEX_CONSTRAINT_OFFSET) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5625:9:
				goto _6
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5627:7:
			if !(Xvec0_column_idx_is_partition(tls, p, iColumn1) != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5628:9:
				goto _6
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5631:11:
			partition_idx = Xvec0_column_idx_to_partition_idx(tls, p, iColumn1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5632:12:
			value = 0
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5634:7:
			switch op1 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5635:9:
			case int32(m_SQLITE_INDEX_CONSTRAINT_EQ):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5636:11:
				value = int8(_VEC0_PARTITION_OPERATOR_EQ)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5637:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5639:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_GT):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5640:11:
				value = int8(_VEC0_PARTITION_OPERATOR_GT)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5641:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5643:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_LE):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5644:11:
				value = int8(_VEC0_PARTITION_OPERATOR_LE)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5645:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5647:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_LT):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5648:11:
				value = int8(_VEC0_PARTITION_OPERATOR_LT)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5649:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5651:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_GE):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5652:11:
				value = int8(_VEC0_PARTITION_OPERATOR_GE)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5653:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5655:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_NE):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5656:11:
				value = int8(_VEC0_PARTITION_OPERATOR_NE)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5657:11:
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5661:7:
			if value != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5662:9:
				v2 = argvIndex
				argvIndex = argvIndex + 1
				(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i1)*8))).FargvIndex = v2
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5663:9:
				(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i1)*8))).Fomit = uint8(1)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5664:9:
				libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_PARTITON_CONSTRAINT))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5665:9:
				libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(int32('A')+partition_idx))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5666:9:
				libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), value)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5667:9:
				libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8('_'))
			}
			goto _6
		_6:
			;
			i1 = i1 + 1
		}
		// find any metadata column constraints
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5673:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5673:14:
		i2 = 0
		for {
			if !(i2 < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5674:7:
			if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i2)*12))).Fusable != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5675:9:
				goto _8
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5677:11:
			iColumn2 = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i2)*12))).FiColumn
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5678:11:
			op2 = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i2)*12))).Fop)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5679:7:
			if op2 == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) || op2 == int32(m_SQLITE_INDEX_CONSTRAINT_OFFSET) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5680:9:
				goto _8
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5682:7:
			if !(Xvec0_column_idx_is_metadata(tls, p, iColumn2) != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5683:9:
				goto _8
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5686:11:
			metadata_idx = Xvec0_column_idx_to_metadata_idx(tls, p, iColumn2)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5687:12:
			value1 = 0
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5689:7:
			switch op2 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5690:9:
			case int32(m_SQLITE_INDEX_CONSTRAINT_EQ):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5691:15:
				vtabIn1 = 0
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5693:11:
				if libsqlite3.Xsqlite3_libversion_number(tls) >= int32(3038000) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5694:13:
					vtabIn1 = libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, i2, -int32(1))
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5696:11:
				if vtabIn1 != 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5697:13:
					switch (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5698:15:
					case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5699:15:
						fallthrough
					case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
						// IMP: V15248_32086
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5701:17:
						rc = int32(m_SQLITE_ERROR)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5702:17:
						Xvtab_set_error(tls, pVTab, __ccgo_ts+9000, 0)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5703:17:
						goto done
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5704:17:
						break
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5706:15:
						fallthrough
					case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5707:15:
						fallthrough
					case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5708:17:
						break
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5711:13:
					value1 = int8(_VEC0_METADATA_OPERATOR_IN)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5712:13:
					libsqlite3.Xsqlite3_vtab_in(tls, pIdxInfo, i2, int32(1))
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5716:13:
					value1 = int8(_VEC0_PARTITION_OPERATOR_EQ)
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5718:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5720:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_GT):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5721:11:
				value1 = int8(_VEC0_METADATA_OPERATOR_GT)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5722:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5724:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_LE):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5725:11:
				value1 = int8(_VEC0_METADATA_OPERATOR_LE)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5726:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5728:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_LT):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5729:11:
				value1 = int8(_VEC0_METADATA_OPERATOR_LT)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5730:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5732:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_GE):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5733:11:
				value1 = int8(_VEC0_METADATA_OPERATOR_GE)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5734:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5736:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_NE):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5737:11:
				value1 = int8(_VEC0_METADATA_OPERATOR_NE)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5738:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5740:9:
				fallthrough
			default:
				// IMP: V16511_00582
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5742:11:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5743:11:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+9070, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5747:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5751:7:
			if (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind == int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5752:9:
				if !(int32(value1) == int32(_VEC0_METADATA_OPERATOR_EQ) || int32(value1) == int32(_VEC0_METADATA_OPERATOR_NE)) {
					// IMP: V10145_26984
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5754:11:
					rc = int32(m_SQLITE_ERROR)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5755:11:
					Xvtab_set_error(tls, pVTab, __ccgo_ts+9264, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5756:11:
					goto done
				}
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5760:7:
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i2)*8))).FargvIndex = v2
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5761:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i2)*8))).Fomit = uint8(1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5762:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5763:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(int32('A')+metadata_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5764:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), value1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5765:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8('_'))
			goto _8
		_8:
			;
			i2 = i2 + 1
		}
		// find any distance column constraints
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5770:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5770:14:
		i3 = 0
		for {
			if !(i3 < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5771:7:
			if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i3)*12))).Fusable != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5772:9:
				goto _10
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5774:11:
			iColumn3 = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i3)*12))).FiColumn
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5775:11:
			op3 = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i3)*12))).Fop)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5776:7:
			if op3 == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) || op3 == int32(m_SQLITE_INDEX_CONSTRAINT_OFFSET) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5777:9:
				goto _10
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5779:7:
			if Xvec0_column_distance_idx(tls, p) != iColumn3 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5780:9:
				goto _10
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5783:12:
			value2 = 0
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5784:7:
			switch op3 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5785:9:
			case int32(m_SQLITE_INDEX_CONSTRAINT_GT):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5786:11:
				value2 = int8(_VEC0_DISTANCE_CONSTRAINT_GT)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5787:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5789:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_GE):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5790:11:
				value2 = int8(_VEC0_DISTANCE_CONSTRAINT_GE)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5791:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5793:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_LT):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5794:11:
				value2 = int8(_VEC0_DISTANCE_CONSTRAINT_LT)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5795:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5797:9:
				fallthrough
			case int32(m_SQLITE_INDEX_CONSTRAINT_LE):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5798:11:
				value2 = int8(_VEC0_DISTANCE_CONSTRAINT_LE)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5799:11:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5801:9:
				fallthrough
			default:
				// IMP TODO
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5803:11:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5804:11:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+9350, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5809:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5813:7:
			v2 = argvIndex
			argvIndex = argvIndex + 1
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i3)*8))).FargvIndex = v2
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5814:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(i3)*8))).Fomit = uint8(1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5815:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5816:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), value2)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5817:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8('_'))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5818:7:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8('_'))
			goto _10
		_10:
			;
			i3 = i3 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5823:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxNum = iMatchVectorTerm
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5824:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = float64(30)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5825:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(10)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5827:10:
		if iRowidTerm >= 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5828:5:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_QUERY_PLAN_POINT))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5829:5:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iRowidTerm)*8))).FargvIndex = int32(1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5830:5:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iRowidTerm)*8))).Fomit = uint8(1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5831:5:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_IDXSTR_KIND_POINT_ID))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5832:5:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(3), int8('_'))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5833:5:
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxNum = libc.Int32FromUint64((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FcolUsed)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5834:5:
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = float64(10)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5835:5:
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(1)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5837:5:
			libsqlite3.Xsqlite3_str_appendchar(tls, idxStr, int32(1), int8(_VEC0_QUERY_PLAN_FULLSCAN))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5838:5:
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = float64(3e+06)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5839:5:
			(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(100000)
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5841:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxStr = libsqlite3.Xsqlite3_str_finish(tls, idxStr)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5842:3:
	idxStr = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5843:3:
	if !((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxStr != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5844:5:
		rc = m_SQLITE_OK
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5845:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5847:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FneedToFreeIdxStr = int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5849:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5851:3:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5852:5:
	if idxStr != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5853:7:
		libsqlite3.Xsqlite3_str_finish(tls, idxStr)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5855:5:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5861:6:
func Xmerge_sorted_lists(tls *libc.TLS, a uintptr, a_rowids uintptr, a_length Ti64, b uintptr, b_rowids uintptr, b_top_idxs uintptr, b_length Ti64, out uintptr, out_rowids uintptr, out_length Ti64, out_used uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5863:73:
	var i int32
	var ptrA, ptrB Ti64
	_, _, _ = i, ptrA, ptrB
	// assert((a_length >= out_length) || (b_length >= out_length));
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5865:7:
	ptrA = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5866:7:
	ptrB = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5867:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5867:12:
	i = 0
	for {
		if !(int64(i) < out_length) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5868:5:
		if ptrA >= a_length && ptrB >= b_length {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5869:7:
			**(**Ti64)(__ccgo_up(out_used)) = int64(i)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5870:7:
			return
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5872:5:
		if ptrA >= a_length {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5873:7:
			**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(b + uintptr(**(**Ti32)(__ccgo_up(b_top_idxs + uintptr(ptrB)*4)))*4))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5874:7:
			**(**Ti64)(__ccgo_up(out_rowids + uintptr(i)*8)) = **(**Ti64)(__ccgo_up(b_rowids + uintptr(**(**Ti32)(__ccgo_up(b_top_idxs + uintptr(ptrB)*4)))*8))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5875:7:
			ptrB = ptrB + 1
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5876:12:
			if ptrB >= b_length {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5877:7:
				**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(a + uintptr(ptrA)*4))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5878:7:
				**(**Ti64)(__ccgo_up(out_rowids + uintptr(i)*8)) = **(**Ti64)(__ccgo_up(a_rowids + uintptr(ptrA)*8))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5879:7:
				ptrA = ptrA + 1
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5881:7:
				if **(**Tf32)(__ccgo_up(a + uintptr(ptrA)*4)) <= **(**Tf32)(__ccgo_up(b + uintptr(**(**Ti32)(__ccgo_up(b_top_idxs + uintptr(ptrB)*4)))*4)) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5882:9:
					**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(a + uintptr(ptrA)*4))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5883:9:
					**(**Ti64)(__ccgo_up(out_rowids + uintptr(i)*8)) = **(**Ti64)(__ccgo_up(a_rowids + uintptr(ptrA)*8))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5884:9:
					ptrA = ptrA + 1
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5886:9:
					**(**Tf32)(__ccgo_up(out + uintptr(i)*4)) = **(**Tf32)(__ccgo_up(b + uintptr(**(**Ti32)(__ccgo_up(b_top_idxs + uintptr(ptrB)*4)))*4))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5887:9:
					**(**Ti64)(__ccgo_up(out_rowids + uintptr(i)*8)) = **(**Ti64)(__ccgo_up(b_rowids + uintptr(**(**Ti32)(__ccgo_up(b_top_idxs + uintptr(ptrB)*4)))*8))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5888:9:
					ptrB = ptrB + 1
				}
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5893:3:
	**(**Ti64)(__ccgo_up(out_used)) = out_length
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5896:4:
func Xbitmap_new(tls *libc.TLS, n Ti32) (r uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5896:23:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5897:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5898:6:
	p = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt32(n)*uint64(1)/uint64(m_CHAR_BIT)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5899:3:
	if p != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5900:5:
		libc.Xmemset(tls, p, 0, libc.Uint64FromInt32(n)*uint64(1)/uint64(m_CHAR_BIT))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5902:3:
	return p
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5904:4:
func Xbitmap_new_from(tls *libc.TLS, n Ti32, from uintptr) (r uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5904:38:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5905:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5906:6:
	p = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt32(n)*uint64(1)/uint64(m_CHAR_BIT)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5907:3:
	if p != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5908:5:
		libc.Xmemcpy(tls, p, from, libc.Uint64FromInt32(n/int32(m_CHAR_BIT)))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5910:3:
	return p
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5913:6:
func Xbitmap_copy(tls *libc.TLS, base uintptr, from uintptr, n Ti32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5913:45:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5914:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5915:3:
	libc.Xmemcpy(tls, base, from, libc.Uint64FromInt32(n/int32(m_CHAR_BIT)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5918:6:
func Xbitmap_and_inplace(tls *libc.TLS, base uintptr, other uintptr, n Ti32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5918:53:
	var i int32
	_ = i
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5919:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5920:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5920:12:
	i = 0
	for {
		if !(i < n/int32(m_CHAR_BIT)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5921:5:
		**(**Tu8)(__ccgo_up(base + uintptr(i))) = libc.Uint8FromInt32(libc.Int32FromUint8(**(**Tu8)(__ccgo_up(base + uintptr(i)))) & libc.Int32FromUint8(**(**Tu8)(__ccgo_up(other + uintptr(i)))))
		goto _1
	_1:
		;
		i = i + 1
	}
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5925:6:
func Xbitmap_set(tls *libc.TLS, bitmap uintptr, position Ti32, value int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5925:54:
	var v1 uintptr
	_ = v1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5926:3:
	if value != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5927:5:
		v1 = bitmap + uintptr(position/int32(m_CHAR_BIT))
		*(*Tu8)(unsafe.Pointer(v1)) = Tu8(int32(*(*Tu8)(unsafe.Pointer(v1))) | libc.Int32FromInt32(1)<<(position%libc.Int32FromInt32(m_CHAR_BIT)))
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5929:5:
		v1 = bitmap + uintptr(position/int32(m_CHAR_BIT))
		*(*Tu8)(unsafe.Pointer(v1)) = Tu8(int32(*(*Tu8)(unsafe.Pointer(v1))) & ^(libc.Int32FromInt32(1) << (position % libc.Int32FromInt32(m_CHAR_BIT))))
	}
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5933:5:
func Xbitmap_get(tls *libc.TLS, bitmap uintptr, position Ti32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5933:42:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5934:3:
	return libc.Int32FromUint8(**(**Tu8)(__ccgo_up(bitmap + uintptr(position/int32(m_CHAR_BIT))))) >> (position % int32(m_CHAR_BIT)) & int32(1)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5937:6:
func Xbitmap_clear(tls *libc.TLS, bitmap uintptr, n Ti32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5937:38:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5938:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5939:3:
	libc.Xmemset(tls, bitmap, 0, libc.Uint64FromInt32(n/int32(m_CHAR_BIT)))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5942:6:
func Xbitmap_fill(tls *libc.TLS, bitmap uintptr, n Ti32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5942:37:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5943:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5944:3:
	libc.Xmemset(tls, bitmap, int32(0xFF), libc.Uint64FromInt32(n/int32(m_CHAR_BIT)))
}

// C documentation
//
//	/**
//	 * @brief Finds the minimum k items in distances, and writes the indicies to
//	 * out.
//	 *
//	 * @param distances input f32 array of size n, the items to consider.
//	 * @param n: size of distances array.
//	 * @param out: Output array of size k, will contain at most k element indicies
//	 * @param k: Size of output array
//	 * @return int
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5957:5:
func Xmin_idx(tls *libc.TLS, distances uintptr, n Ti32, candidates uintptr, out uintptr, k Ti32, bTaken uintptr, k_used uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5958:38:
	var i, ik, min_idx int32
	_, _, _ = i, ik, min_idx
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5959:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5960:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5962:3:
	Xbitmap_clear(tls, bTaken, n)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5964:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5964:12:
	ik = 0
	for {
		if !(ik < k) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5965:9:
		min_idx = 0
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5966:5:
		for min_idx < n && (Xbitmap_get(tls, bTaken, min_idx) != 0 || !(Xbitmap_get(tls, candidates, min_idx) != 0)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5968:7:
			min_idx = min_idx + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5970:5:
		if min_idx >= n {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5971:7:
			**(**Ti32)(__ccgo_up(k_used)) = ik
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5972:7:
			return m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5975:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5975:14:
		i = 0
		for {
			if !(i < n) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5976:7:
			if **(**Tf32)(__ccgo_up(distances + uintptr(i)*4)) <= **(**Tf32)(__ccgo_up(distances + uintptr(min_idx)*4)) && !(Xbitmap_get(tls, bTaken, i) != 0) && Xbitmap_get(tls, candidates, i) != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5978:9:
				min_idx = i
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5982:5:
		**(**Ti32)(__ccgo_up(out + uintptr(ik)*4)) = min_idx
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5983:5:
		Xbitmap_set(tls, bTaken, min_idx, int32(1))
		goto _1
	_1:
		;
		ik = ik + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5985:3:
	**(**Ti32)(__ccgo_up(k_used)) = k
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5986:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5989:5:
func Xvec0_get_metadata_text_long_value(tls *libc.TLS, p uintptr, stmt uintptr, metadata_idx int32, rowid Ti64, n uintptr, s uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5995:14:
	var rc int32
	var zSql uintptr
	_, _ = rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5997:3:
	if !(**(**uintptr)(__ccgo_up(stmt)) != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5998:16:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+9462, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_idx))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:5999:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6000:7:
			rc = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6001:7:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6003:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), stmt, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6004:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6005:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6006:7:
			goto done
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6010:3:
	libsqlite3.Xsqlite3_reset(tls, **(**uintptr)(__ccgo_up(stmt)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6011:3:
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(stmt)), int32(1), rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6012:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(stmt)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6013:3:
	if rc != int32(m_SQLITE_ROW) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6014:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6015:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6017:3:
	**(**uintptr)(__ccgo_up(s)) = libsqlite3.Xsqlite3_column_text(tls, **(**uintptr)(__ccgo_up(stmt)), 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6018:3:
	**(**int32)(__ccgo_up(n)) = libsqlite3.Xsqlite3_column_bytes(tls, **(**uintptr)(__ccgo_up(stmt)), 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6019:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6020:3:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6021:5:
	return rc
	return r
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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6038:5:
func Xvec0_chunks_iter(tls *libc.TLS, p uintptr, idxStr uintptr, argc int32, argv uintptr, outStmt uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6038:115:
	var appendedWhere, i, i1, idx, idx1, idxStrLength, n, numValueEntries, operator, partition_idx, rc, v3 int32
	var kind, kind1 int8
	var s, zSql, zSql1 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = appendedWhere, i, i1, idx, idx1, idxStrLength, kind, kind1, n, numValueEntries, operator, partition_idx, rc, s, zSql, zSql1, v3
	// always null terminated, enforced by SQLite
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6040:7:
	idxStrLength = libc.Int32FromUint64(libc.Xstrlen(tls, idxStr))
	// "1" refers to the initial vec0_query_plan char, 4 is the number of chars per "element"
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6042:7:
	numValueEntries = (idxStrLength - int32(1)) / int32(4)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6043:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6046:15:
	s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6047:3:
	libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9522, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6051:7:
	appendedWhere = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6052:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6052:11:
	i = 0
	for {
		if !(i < numValueEntries) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6053:9:
		idx = int32(1) + i*int32(4)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6054:10:
		kind = **(**int8)(__ccgo_up(idxStr + uintptr(idx+0)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6055:5:
		if int32(kind) != int32(_VEC0_IDXSTR_KIND_KNN_PARTITON_CONSTRAINT) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6056:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6059:9:
		partition_idx = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx+int32(1))))) - int32('A')
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6060:9:
		operator = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx+int32(2)))))
		// idxStr[idx + 3] is just null, a '_' placeholder
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6063:5:
		if !(appendedWhere != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6064:7:
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+9579)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6065:7:
			appendedWhere = int32(1)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6067:7:
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+4165)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6069:5:
		switch operator {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6070:6:
		case int32(_VEC0_PARTITION_OPERATOR_EQ):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6071:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+4171, libc.VaList(bp+8, partition_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6072:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6073:6:
			fallthrough
		case int32(_VEC0_PARTITION_OPERATOR_GT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6074:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9587, libc.VaList(bp+8, partition_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6075:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6076:6:
			fallthrough
		case int32(_VEC0_PARTITION_OPERATOR_LE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6077:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9607, libc.VaList(bp+8, partition_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6078:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6079:6:
			fallthrough
		case int32(_VEC0_PARTITION_OPERATOR_LT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6080:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9628, libc.VaList(bp+8, partition_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6081:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6082:6:
			fallthrough
		case int32(_VEC0_PARTITION_OPERATOR_GE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6083:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9648, libc.VaList(bp+8, partition_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6084:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6085:6:
			fallthrough
		case int32(_VEC0_PARTITION_OPERATOR_NE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6086:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+9669, libc.VaList(bp+8, partition_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6087:7:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6088:6:
			fallthrough
		default:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6089:12:
			zSql = libsqlite3.Xsqlite3_str_finish(tls, s)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6090:7:
			libsqlite3.Xsqlite3_free(tls, zSql)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6091:7:
			return int32(m_SQLITE_ERROR)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6098:8:
	zSql1 = libsqlite3.Xsqlite3_str_finish(tls, s)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6099:3:
	if !(zSql1 != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6100:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6103:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql1, -int32(1), outStmt, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6104:3:
	libsqlite3.Xsqlite3_free(tls, zSql1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6105:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6106:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6109:7:
	n = int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6110:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6110:11:
	i1 = 0
	for {
		if !(i1 < numValueEntries) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6111:9:
		idx1 = int32(1) + i1*int32(4)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6112:10:
		kind1 = **(**int8)(__ccgo_up(idxStr + uintptr(idx1+0)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6113:5:
		if int32(kind1) != int32(_VEC0_IDXSTR_KIND_KNN_PARTITON_CONSTRAINT) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6114:7:
			goto _2
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6116:5:
		v3 = n
		n = n + 1
		libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(outStmt)), v3, **(**uintptr)(__ccgo_up(argv + uintptr(i1)*8)))
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6119:3:
	return rc
}

// a single `xxx in (...)` constraint on a metadata column. TEXT or INTEGER only for now.

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6123:1:
type TVec0MetadataIn = struct {
	Fargv_idx     int32
	Fmetadata_idx int32
	Farray        TArray
}

type Vec0MetadataIn = TVec0MetadataIn

// Array elements for `xxx in (...)` values for a text column. basically just a string

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6133:1:
type TVec0MetadataInTextEntry = struct {
	Fn       int32
	FzString uintptr
}

type Vec0MetadataInTextEntry = TVec0MetadataInTextEntry

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6139:5:
func Xvec0_metadata_filter_text(tls *libc.TLS, p uintptr, value uintptr, buffer uintptr, size int32, op Tvec0_metadata_operator, b uintptr, metadata_idx int32, chunk_rowid int32, aMetadataIn uintptr, argv_idx int32) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6139:211:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6141:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6142:7:
	rowids = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6144:14:
	sTarget = libsqlite3.Xsqlite3_value_text(tls, value)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6145:7:
	nTarget = libsqlite3.Xsqlite3_value_bytes(tls, value)
	// TODO(perf): only text metadata news the rowids BLOB. Make it so that
	// rowids BLOB is re-used when multiple fitlers on text columns,
	// ex "name BETWEEN 'a' and 'b'""
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6151:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+9690, int64(chunk_rowid), 0, bp+8)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6152:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6153:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6155:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6156:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6158:3:
	rowids = libsqlite3.Xsqlite3_malloc(tls, libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp + 8))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6159:3:
	if !(rowids != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6160:5:
		libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6161:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6164:3:
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp + 8)), rowids, libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp + 8))), 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6165:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6166:5:
		libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6167:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6169:3:
	libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6171:3:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6177:5:
_1:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6178:7:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6178:15:
	i = 0
	for {
		if !(i < size) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6179:9:
		view = buffer + uintptr(i*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6180:9:
		nPrefix = **(**int32)(__ccgo_up(view))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6181:9:
		sPrefix = view + 4
		// for EQ the text lengths must match
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6184:9:
		if nPrefix != nTarget {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6185:11:
			Xbitmap_set(tls, b, i, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6186:11:
			goto _9
		}
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			v10 = nPrefix
		} else {
			v10 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6188:13:
		cmpPrefix = libc.Xstrncmp(tls, sPrefix, sTarget, libc.Uint64FromInt32(v10))
		// for short strings, use the prefix comparison direclty
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6191:9:
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6192:11:
			Xbitmap_set(tls, b, i, libc.BoolInt32(cmpPrefix == 0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6193:11:
			goto _9
		}
		// for EQ on longs strings, the prefix must match
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6196:9:
		if cmpPrefix != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6197:11:
			Xbitmap_set(tls, b, i, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6198:11:
			goto _9
		}
		// consult the full string
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6201:9:
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i)*8)), bp+24, bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6202:9:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6203:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6205:9:
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6206:11:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6207:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6209:9:
		Xbitmap_set(tls, b, i, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, libc.Uint64FromInt32(**(**int32)(__ccgo_up(bp + 24)))) == 0))
		goto _9
	_9:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6211:7:
	goto _8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6213:5:
_2:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6214:7:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6214:15:
	i1 = 0
	for {
		if !(i1 < size) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6215:9:
		view = buffer + uintptr(i1*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6216:9:
		nPrefix = **(**int32)(__ccgo_up(view))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6217:9:
		sPrefix = view + 4
		// for NE if text lengths dont match, it never will
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6220:9:
		if nPrefix != nTarget {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6221:11:
			Xbitmap_set(tls, b, i1, int32(1))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6222:11:
			goto _11
		}
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			v10 = nPrefix
		} else {
			v10 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6225:13:
		cmpPrefix1 = libc.Xstrncmp(tls, sPrefix, sTarget, libc.Uint64FromInt32(v10))
		// for short strings, use the prefix comparison direclty
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6228:9:
		if nPrefix <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6229:11:
			Xbitmap_set(tls, b, i1, libc.BoolInt32(cmpPrefix1 != 0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6230:11:
			goto _11
		}
		// for NE on longs strings, if prefixes dont match, then long string wont
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6233:9:
		if cmpPrefix1 != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6234:11:
			Xbitmap_set(tls, b, i1, int32(1))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6235:11:
			goto _11
		}
		// consult the full string
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6238:9:
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i1)*8)), bp+24, bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6239:9:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6240:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6242:9:
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6243:11:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6244:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6246:9:
		Xbitmap_set(tls, b, i1, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, libc.Uint64FromInt32(**(**int32)(__ccgo_up(bp + 24)))) != 0))
		goto _11
	_11:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6248:7:
	goto _8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6250:5:
_3:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6251:7:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6251:15:
	i2 = 0
	for {
		if !(i2 < size) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6252:9:
		view = buffer + uintptr(i2*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6253:9:
		nPrefix = **(**int32)(__ccgo_up(view))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6254:9:
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
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6255:13:
		cmpPrefix2 = libc.Xstrncmp(tls, sPrefix, sTarget, libc.Uint64FromInt32(v10))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6257:9:
		if nPrefix < int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// if prefix match, check which is longer
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6259:11:
			if cmpPrefix2 == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6260:13:
				Xbitmap_set(tls, b, i2, libc.BoolInt32(nPrefix > nTarget))
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6263:13:
				Xbitmap_set(tls, b, i2, libc.BoolInt32(cmpPrefix2 > 0))
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6265:11:
			goto _13
		}
		// TODO(perf): may not need to compare full text in some cases
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6269:9:
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i2)*8)), bp+24, bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6270:9:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6271:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6273:9:
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6274:11:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6275:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6277:9:
		Xbitmap_set(tls, b, i2, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, libc.Uint64FromInt32(**(**int32)(__ccgo_up(bp + 24)))) > 0))
		goto _13
	_13:
		;
		i2 = i2 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6279:7:
	goto _8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6281:5:
_4:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6282:7:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6282:15:
	i3 = 0
	for {
		if !(i3 < size) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6283:9:
		view = buffer + uintptr(i3*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6284:9:
		nPrefix = **(**int32)(__ccgo_up(view))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6285:9:
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
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6286:13:
		cmpPrefix3 = libc.Xstrncmp(tls, sPrefix, sTarget, libc.Uint64FromInt32(v10))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6288:9:
		if nPrefix < int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// if prefix match, check which is longer
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6290:11:
			if cmpPrefix3 == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6291:13:
				Xbitmap_set(tls, b, i3, libc.BoolInt32(nPrefix >= nTarget))
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6294:13:
				Xbitmap_set(tls, b, i3, libc.BoolInt32(cmpPrefix3 >= 0))
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6296:11:
			goto _17
		}
		// TODO(perf): may not need to compare full text in some cases
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6300:9:
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i3)*8)), bp+24, bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6301:9:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6302:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6304:9:
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6305:11:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6306:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6308:9:
		Xbitmap_set(tls, b, i3, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, libc.Uint64FromInt32(**(**int32)(__ccgo_up(bp + 24)))) >= 0))
		goto _17
	_17:
		;
		i3 = i3 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6310:7:
	goto _8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6312:5:
_5:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6313:7:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6313:15:
	i4 = 0
	for {
		if !(i4 < size) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6314:9:
		view = buffer + uintptr(i4*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6315:9:
		nPrefix = **(**int32)(__ccgo_up(view))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6316:9:
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
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6317:13:
		cmpPrefix4 = libc.Xstrncmp(tls, sPrefix, sTarget, libc.Uint64FromInt32(v10))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6319:9:
		if nPrefix < int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// if prefix match, check which is longer
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6321:11:
			if cmpPrefix4 == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6322:13:
				Xbitmap_set(tls, b, i4, libc.BoolInt32(nPrefix <= nTarget))
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6325:13:
				Xbitmap_set(tls, b, i4, libc.BoolInt32(cmpPrefix4 <= 0))
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6327:11:
			goto _21
		}
		// TODO(perf): may not need to compare full text in some cases
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6331:9:
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i4)*8)), bp+24, bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6332:9:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6333:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6335:9:
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6336:11:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6337:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6339:9:
		Xbitmap_set(tls, b, i4, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, libc.Uint64FromInt32(**(**int32)(__ccgo_up(bp + 24)))) <= 0))
		goto _21
	_21:
		;
		i4 = i4 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6341:7:
	goto _8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6343:5:
_6:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6344:7:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6344:15:
	i5 = 0
	for {
		if !(i5 < size) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6345:9:
		view = buffer + uintptr(i5*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6346:9:
		nPrefix = **(**int32)(__ccgo_up(view))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6347:9:
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
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6348:13:
		cmpPrefix5 = libc.Xstrncmp(tls, sPrefix, sTarget, libc.Uint64FromInt32(v10))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6350:9:
		if nPrefix < int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// if prefix match, check which is longer
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6352:11:
			if cmpPrefix5 == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6353:13:
				Xbitmap_set(tls, b, i5, libc.BoolInt32(nPrefix < nTarget))
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6356:13:
				Xbitmap_set(tls, b, i5, libc.BoolInt32(cmpPrefix5 < 0))
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6358:11:
			goto _25
		}
		// TODO(perf): may not need to compare full text in some cases
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6362:9:
		rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i5)*8)), bp+24, bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6363:9:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6364:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6366:9:
		if nPrefix != **(**int32)(__ccgo_up(bp + 24)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6367:11:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6368:11:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6370:9:
		Xbitmap_set(tls, b, i5, libc.BoolInt32(libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 16)), sTarget, libc.Uint64FromInt32(**(**int32)(__ccgo_up(bp + 24)))) < 0))
		goto _25
	_25:
		;
		i5 = i5 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6372:7:
	goto _8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6375:5:
_7:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6376:14:
	metadataInIdx = libc.Uint64FromInt32(-libc.Int32FromInt32(1))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6377:7:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6377:18:
	i6 = uint64(0)
	for {
		if !(i6 < (*TArray)(unsafe.Pointer(aMetadataIn)).Flength) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6378:31:
		metadataIn = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(i6)*40
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6379:9:
		if (*TVec0MetadataIn)(unsafe.Pointer(metadataIn)).Fargv_idx == argv_idx {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6380:11:
			metadataInIdx = i6
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6381:11:
			break
		}
		goto _29
	_29:
		;
		i6 = i6 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6384:7:
	if metadataInIdx < uint64(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6385:9:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6386:9:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6389:29:
	metadataIn1 = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(metadataInIdx)*40
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6390:20:
	aTarget = metadataIn1 + 8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6398:7:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6398:15:
	i7 = 0
	for {
		if !(i7 < size) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6399:9:
		view1 = buffer + uintptr(i7*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6400:9:
		nPrefix1 = **(**int32)(__ccgo_up(view1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6401:9:
		sPrefix1 = view1 + 4
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6402:9:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6402:20:
		target_idx = uint64(0)
		for {
			if !(target_idx < (*TArray)(unsafe.Pointer(aTarget)).Flength) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6403:42:
			entry = (*TArray)(unsafe.Pointer(aTarget)).Fz + uintptr(target_idx)*16
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6404:11:
			if (*TVec0MetadataInTextEntry)(unsafe.Pointer(entry)).Fn != nPrefix1 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6405:13:
				goto _31
			}
			if nPrefix1 <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				v10 = nPrefix1
			} else {
				v10 = int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6407:15:
			cmpPrefix6 = libc.Xstrncmp(tls, sPrefix1, (*TVec0MetadataInTextEntry)(unsafe.Pointer(entry)).FzString, libc.Uint64FromInt32(v10))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6408:11:
			if nPrefix1 <= int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6409:13:
				if cmpPrefix6 == 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6410:15:
					Xbitmap_set(tls, b, i7, int32(1))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6411:15:
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6413:13:
				goto _31
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6415:11:
			if cmpPrefix6 != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6416:13:
				goto _31
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6419:11:
			rc = Xvec0_get_metadata_text_long_value(tls, p, bp, metadata_idx, **(**Ti64)(__ccgo_up(rowids + uintptr(i7)*8)), bp+40, bp+32)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6420:11:
			if rc != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6421:13:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6423:11:
			if nPrefix1 != **(**int32)(__ccgo_up(bp + 40)) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6424:13:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6425:13:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6427:11:
			if libc.Xstrncmp(tls, **(**uintptr)(__ccgo_up(bp + 32)), (*TVec0MetadataInTextEntry)(unsafe.Pointer(entry)).FzString, libc.Uint64FromInt32(**(**int32)(__ccgo_up(bp + 40)))) == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6428:13:
				Xbitmap_set(tls, b, i7, int32(1))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6429:13:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6433:7:
	goto _8
_8:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6437:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6439:3:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6440:5:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6441:5:
	libsqlite3.Xsqlite3_free(tls, rowids)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6442:5:
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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6458:5:
func Xvec0_set_metadata_filter_bitmap(tls *libc.TLS, p uintptr, metadata_idx int32, op Tvec0_metadata_operator, value uintptr, blob uintptr, chunk_rowid Ti64, b uintptr, size int32, aMetadataIn uintptr, argv_idx int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6467:45:
	var aTarget, array, array1, buffer, metadataIn, metadataIn1 uintptr
	var blobSize, i, i1, i10, i11, i12, i13, i14, i15, i2, i3, i4, i5, i6, i7, i9, metadataInIdx, rc, szMatch, target int32
	var i8, target_idx Tsize_t
	var kind Tvec0_metadata_column_kind
	var target1 Ti64
	var target2 float64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aTarget, array, array1, blobSize, buffer, i, i1, i10, i11, i12, i13, i14, i15, i2, i3, i4, i5, i6, i7, i8, i9, kind, metadataIn, metadataIn1, metadataInIdx, rc, szMatch, target, target1, target2, target_idx
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6471:3:
	rc = libsqlite3.Xsqlite3_blob_reopen(tls, blob, chunk_rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6472:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6473:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6476:29:
	kind = (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6477:7:
	szMatch = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6478:7:
	blobSize = libsqlite3.Xsqlite3_blob_bytes(tls, blob)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6479:3:
	switch kind {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6480:5:
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6481:7:
		szMatch = libc.BoolInt32(blobSize == size/int32(m_CHAR_BIT))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6482:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6484:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6485:7:
		szMatch = libc.BoolInt32(libc.Uint64FromInt32(blobSize) == libc.Uint64FromInt32(size)*uint64(8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6486:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6488:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6489:7:
		szMatch = libc.BoolInt32(libc.Uint64FromInt32(blobSize) == libc.Uint64FromInt32(size)*uint64(8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6490:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6492:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6493:7:
		szMatch = libc.BoolInt32(blobSize == size*int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6494:7:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6497:3:
	if !(szMatch != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6498:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6500:8:
	buffer = libsqlite3.Xsqlite3_malloc(tls, blobSize)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6501:3:
	if !(buffer != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6502:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6504:3:
	rc = libsqlite3.Xsqlite3_blob_read(tls, blob, buffer, blobSize, 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6505:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6506:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6508:3:
	switch kind {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6509:5:
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6510:11:
		target = libsqlite3.Xsqlite3_value_int(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6511:7:
		if target != 0 && op == int32(_VEC0_METADATA_OPERATOR_EQ) || !(target != 0) && op == int32(_VEC0_METADATA_OPERATOR_NE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6512:9:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6512:17:
			i = 0
			for {
				if !(i < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6512:41:
				Xbitmap_set(tls, b, i, Xbitmap_get(tls, buffer, i))
				goto _1
			_1:
				;
				i = i + 1
			}
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6515:9:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6515:17:
			i1 = 0
			for {
				if !(i1 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6515:41:
				Xbitmap_set(tls, b, i1, libc.BoolInt32(!(Xbitmap_get(tls, buffer, i1) != 0)))
				goto _2
			_2:
				;
				i1 = i1 + 1
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6517:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6519:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6520:11:
		array = buffer
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6521:11:
		target1 = libsqlite3.Xsqlite3_value_int64(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6522:7:
		switch op {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6523:9:
		case int32(_VEC0_METADATA_OPERATOR_EQ):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6524:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6524:19:
			i2 = 0
			for {
				if !(i2 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6524:43:
				Xbitmap_set(tls, b, i2, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i2)*8)) == target1))
				goto _3
			_3:
				;
				i2 = i2 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6525:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6527:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_GT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6528:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6528:19:
			i3 = 0
			for {
				if !(i3 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6528:43:
				Xbitmap_set(tls, b, i3, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i3)*8)) > target1))
				goto _4
			_4:
				;
				i3 = i3 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6529:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6531:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_LE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6532:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6532:19:
			i4 = 0
			for {
				if !(i4 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6532:43:
				Xbitmap_set(tls, b, i4, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i4)*8)) <= target1))
				goto _5
			_5:
				;
				i4 = i4 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6533:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6535:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_LT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6536:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6536:19:
			i5 = 0
			for {
				if !(i5 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6536:43:
				Xbitmap_set(tls, b, i5, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i5)*8)) < target1))
				goto _6
			_6:
				;
				i5 = i5 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6537:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6539:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_GE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6540:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6540:19:
			i6 = 0
			for {
				if !(i6 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6540:43:
				Xbitmap_set(tls, b, i6, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i6)*8)) >= target1))
				goto _7
			_7:
				;
				i6 = i6 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6541:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6543:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_NE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6544:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6544:19:
			i7 = 0
			for {
				if !(i7 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6544:43:
				Xbitmap_set(tls, b, i7, libc.BoolInt32(**(**Ti64)(__ccgo_up(array + uintptr(i7)*8)) != target1))
				goto _8
			_8:
				;
				i7 = i7 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6545:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6547:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_IN):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6548:15:
			metadataInIdx = -int32(1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6549:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6549:22:
			i8 = uint64(0)
			for {
				if !(i8 < (*TArray)(unsafe.Pointer(aMetadataIn)).Flength) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6550:35:
				metadataIn = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(i8)*40
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6551:13:
				if (*TVec0MetadataIn)(unsafe.Pointer(metadataIn)).Fargv_idx == argv_idx {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6552:15:
					metadataInIdx = libc.Int32FromUint64(i8)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6553:15:
					break
				}
				goto _9
			_9:
				;
				i8 = i8 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6556:11:
			if metadataInIdx < 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6557:13:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6558:13:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6560:33:
			metadataIn1 = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(metadataInIdx)*40
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6561:24:
			aTarget = metadataIn1 + 8
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6563:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6563:19:
			i9 = 0
			for {
				if !(i9 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6564:13:
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6564:24:
				target_idx = uint64(0)
				for {
					if !(target_idx < (*TArray)(unsafe.Pointer(aTarget)).Flength) {
						break
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6565:15:
					if **(**Ti64)(__ccgo_up((*TArray)(unsafe.Pointer(aTarget)).Fz + uintptr(target_idx)*8)) == **(**Ti64)(__ccgo_up(array + uintptr(i9)*8)) {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6566:17:
						Xbitmap_set(tls, b, i9, int32(1))
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6567:17:
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
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6571:11:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6574:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6576:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6577:14:
		array1 = buffer
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6578:14:
		target2 = libsqlite3.Xsqlite3_value_double(tls, value)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6579:7:
		switch op {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6580:9:
		case int32(_VEC0_METADATA_OPERATOR_EQ):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6581:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6581:19:
			i10 = 0
			for {
				if !(i10 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6581:43:
				Xbitmap_set(tls, b, i10, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i10)*8)) == target2))
				goto _12
			_12:
				;
				i10 = i10 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6582:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6584:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_GT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6585:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6585:19:
			i11 = 0
			for {
				if !(i11 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6585:43:
				Xbitmap_set(tls, b, i11, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i11)*8)) > target2))
				goto _13
			_13:
				;
				i11 = i11 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6586:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6588:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_LE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6589:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6589:19:
			i12 = 0
			for {
				if !(i12 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6589:43:
				Xbitmap_set(tls, b, i12, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i12)*8)) <= target2))
				goto _14
			_14:
				;
				i12 = i12 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6590:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6592:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_LT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6593:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6593:19:
			i13 = 0
			for {
				if !(i13 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6593:43:
				Xbitmap_set(tls, b, i13, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i13)*8)) < target2))
				goto _15
			_15:
				;
				i13 = i13 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6594:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6596:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_GE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6597:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6597:19:
			i14 = 0
			for {
				if !(i14 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6597:43:
				Xbitmap_set(tls, b, i14, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i14)*8)) >= target2))
				goto _16
			_16:
				;
				i14 = i14 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6598:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6600:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_NE):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6601:11:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6601:19:
			i15 = 0
			for {
				if !(i15 < size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6601:43:
				Xbitmap_set(tls, b, i15, libc.BoolInt32(**(**float64)(__ccgo_up(array1 + uintptr(i15)*8)) != target2))
				goto _17
			_17:
				;
				i15 = i15 + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6602:11:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6604:9:
			fallthrough
		case int32(_VEC0_METADATA_OPERATOR_IN):
			// should never be reached
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6606:11:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6609:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6611:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6612:7:
		rc = Xvec0_metadata_filter_text(tls, p, value, buffer, size, op, b, metadata_idx, int32(chunk_rowid), aMetadataIn, argv_idx)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6613:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6614:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6616:7:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6619:3:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6620:5:
	libsqlite3.Xsqlite3_free(tls, buffer)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6621:5:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6624:5:
func Xvec0Filter_knn_chunks_iter(tls *libc.TLS, p uintptr, stmtChunks uintptr, vector_column uintptr, vectorColumnIdx int32, arrayRowidsIn uintptr, aMetadataIn uintptr, idxStr uintptr, argc int32, argv uintptr, queryVector uintptr, k Ti64, out_topk_rowids uintptr, out_topk_distances uintptr, out_used uintptr) (r int32) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6630:73:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6635:7:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6636:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6638:8:
	baseVectors = libc.UintptrFromInt32(0) // memory: chunk_size * dimensions * element_size
	// OWNED BY CALLER ON SUCCESS
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6641:7:
	topk_rowids = libc.UintptrFromInt32(0) // memory: k * 4
	// OWNED BY CALLER ON SUCCESS
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6643:7:
	topk_distances = libc.UintptrFromInt32(0) // memory: k * 4
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6645:7:
	tmp_topk_rowids = libc.UintptrFromInt32(0) // memory: k * 4
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6646:7:
	tmp_topk_distances = libc.UintptrFromInt32(0) // memory: k * 4
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6647:7:
	chunk_distances = libc.UintptrFromInt32(0) // memory: chunk_size * 4
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6648:6:
	b = libc.UintptrFromInt32(0) // memory: chunk_size / 8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6649:6:
	bTaken = libc.UintptrFromInt32(0) // memory: chunk_size / 8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6650:7:
	chunk_topk_idxs = libc.UintptrFromInt32(0) // memory: k * 4
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6651:6:
	bmRowids = libc.UintptrFromInt32(0) // memory: chunk_size / 8
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6652:6:
	bmMetadata = libc.UintptrFromInt32(0) // memory: chunk_size / 8
	//                        // total: a lot???
	// 6 * (k * 4) + (k * 2) + (chunk_size / 8) + (chunk_size * dimensions * 4)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6657:3:
	topk_rowids = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6658:3:
	if !(topk_rowids != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6659:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6660:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6662:3:
	libc.Xmemset(tls, topk_rowids, 0, uint64(libc.Uint64FromInt64(k)*uint64(8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6664:3:
	topk_distances = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(4)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6665:3:
	if !(topk_distances != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6666:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6667:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6669:3:
	libc.Xmemset(tls, topk_distances, 0, uint64(libc.Uint64FromInt64(k)*uint64(4)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6671:3:
	tmp_topk_rowids = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6672:3:
	if !(tmp_topk_rowids != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6673:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6674:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6676:3:
	libc.Xmemset(tls, tmp_topk_rowids, 0, uint64(libc.Uint64FromInt64(k)*uint64(8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6678:3:
	tmp_topk_distances = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(4)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6679:3:
	if !(tmp_topk_distances != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6680:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6681:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6683:3:
	libc.Xmemset(tls, tmp_topk_distances, 0, uint64(libc.Uint64FromInt64(k)*uint64(4)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6685:7:
	k_used = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6686:7:
	baseVectorsSize = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(vector_column))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6687:3:
	baseVectors = libsqlite3.Xsqlite3_malloc(tls, int32(baseVectorsSize))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6688:3:
	if !(baseVectors != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6689:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6690:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6693:3:
	chunk_distances = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(4)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6694:3:
	if !(chunk_distances != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6695:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6696:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6699:3:
	b = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6700:3:
	if !(b != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6701:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6702:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6705:3:
	bTaken = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6706:3:
	if !(bTaken != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6707:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6708:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6711:3:
	chunk_topk_idxs = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(4)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6712:3:
	if !(chunk_topk_idxs != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6713:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6714:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6717:3:
	if arrayRowidsIn != 0 {
		v1 = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	bmRowids = v1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6718:3:
	if arrayRowidsIn != 0 && !(bmRowids != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6719:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6720:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6724:3:
	libc.Xmemset(tls, bp+8, 0, libc.Uint64FromInt64(8)*libc.Uint64FromInt32(m_VEC0_MAX_METADATA_COLUMNS))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6726:3:
	bmMetadata = Xbitmap_new(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6727:3:
	if !(bmMetadata != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6728:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6729:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6732:7:
	idxStrLength = libc.Int32FromUint64(libc.Xstrlen(tls, idxStr))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6733:7:
	numValueEntries = (idxStrLength - int32(1)) / int32(4)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6734:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6735:7:
	hasMetadataFilters = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6736:7:
	hasDistanceConstraints = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6737:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6737:11:
	i = 0
	for {
		if !(i < argc) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6738:9:
		idx = int32(1) + i*int32(4)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6739:10:
		kind = **(**int8)(__ccgo_up(idxStr + uintptr(idx+0)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6740:5:
		if int32(kind) == int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6741:7:
			hasMetadataFilters = int32(1)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6743:10:
			if int32(kind) == int32(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6744:7:
				hasDistanceConstraints = int32(1)
			}
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6748:3:
	for int32(m_true) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6749:5:
		rc = libsqlite3.Xsqlite3_step(tls, stmtChunks)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6750:5:
		if rc == int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6751:7:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6753:5:
		if rc != int32(m_SQLITE_ROW) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6754:7:
			Xvtab_set_error(tls, p, __ccgo_ts+9697, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6755:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6756:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6758:5:
		libc.Xmemset(tls, chunk_distances, 0, libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(4))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6759:5:
		libc.Xmemset(tls, chunk_topk_idxs, 0, uint64(libc.Uint64FromInt64(k)*uint64(4)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6760:5:
		Xbitmap_clear(tls, b, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6762:9:
		chunk_id = libsqlite3.Xsqlite3_column_int64(tls, stmtChunks, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6763:19:
		chunkValidity = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6765:9:
		validitySize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(1)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6766:5:
		if validitySize != int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m_CHAR_BIT)) {
			// IMP: V05271_22109
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6768:7:
			Xvtab_set_error(tls, p, __ccgo_ts+9715, libc.VaList(bp+168, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m_CHAR_BIT), validitySize))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6772:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6773:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6776:9:
		chunkRowids = libsqlite3.Xsqlite3_column_blob(tls, stmtChunks, int32(2))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6777:9:
		rowidsSize = int64(libsqlite3.Xsqlite3_column_bytes(tls, stmtChunks, int32(2)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6778:5:
		if libc.Uint64FromInt64(rowidsSize) != uint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8)) {
			// IMP: V02796_19635
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6780:7:
			Xvtab_set_error(tls, p, __ccgo_ts+9777, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6781:7:
			Xvtab_set_error(tls, p, __ccgo_ts+9803, libc.VaList(bp+168, libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)*uint64(8), rowidsSize))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6785:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6786:7:
			goto cleanup
		}
		// open the vector chunk blob for the current chunk
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6790:5:
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(vectorColumnIdx)*8)), __ccgo_ts+3712, chunk_id, 0, bp)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6793:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6794:7:
			Xvtab_set_error(tls, p, __ccgo_ts+9863, libc.VaList(bp+168, chunk_id))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6796:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6797:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6800:9:
		currentBaseVectorsSize = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6801:9:
		expectedBaseVectorsSize = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(vector_column))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6803:5:
		if currentBaseVectorsSize != expectedBaseVectorsSize {
			// IMP: V16465_00535
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6805:7:
			Xvtab_set_error(tls, p, __ccgo_ts+9906, libc.VaList(bp+168, expectedBaseVectorsSize, currentBaseVectorsSize))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6809:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6810:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6812:5:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), baseVectors, int32(currentBaseVectorsSize), 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6814:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6815:7:
			Xvtab_set_error(tls, p, __ccgo_ts+9966, libc.VaList(bp+168, chunk_id))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6816:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6817:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6820:5:
		Xbitmap_copy(tls, b, chunkValidity, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6821:5:
		if arrayRowidsIn != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6822:7:
			Xbitmap_clear(tls, bmRowids, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6824:7:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6824:16:
			i1 = 0
			for {
				if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6825:9:
				if !(Xbitmap_get(tls, chunkValidity, i1) != 0) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6826:11:
					goto _3
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6828:13:
				**(**Ti64)(__ccgo_up(bp + 136)) = **(**Ti64)(__ccgo_up(chunkRowids + uintptr(i1)*8))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6829:14:
				in = libc.Xbsearch(tls, bp+136, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Fz, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Flength, uint64(8), __ccgo_fp(X_cmp))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6831:9:
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
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6833:7:
			Xbitmap_and_inplace(tls, b, bmRowids, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6836:5:
		if hasMetadataFilters != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6837:7:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6837:15:
			i2 = 0
			for {
				if !(i2 < argc) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6838:13:
				idx1 = int32(1) + i2*int32(4)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6839:14:
				kind1 = **(**int8)(__ccgo_up(idxStr + uintptr(idx1+0)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6840:9:
				if int32(kind1) != int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6841:11:
					goto _5
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6843:13:
				metadata_idx = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx1+int32(1))))) - int32('A')
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6844:13:
				operator = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx1+int32(2)))))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6846:9:
				if !((**(**[16]uintptr)(__ccgo_up(bp + 8)))[metadata_idx] != 0) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6847:11:
					rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 480 + uintptr(metadata_idx)*8)), __ccgo_ts+4053, chunk_id, 0, bp+8+uintptr(metadata_idx)*8)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6848:11:
					Xvtab_set_error(tls, p, __ccgo_ts+9999, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6849:11:
					if rc != m_SQLITE_OK {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6850:13:
						goto cleanup
					}
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6854:9:
				Xbitmap_clear(tls, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6855:9:
				rc = Xvec0_set_metadata_filter_bitmap(tls, p, metadata_idx, operator, **(**uintptr)(__ccgo_up(argv + uintptr(i2)*8)), (**(**[16]uintptr)(__ccgo_up(bp + 8)))[metadata_idx], chunk_id, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size, aMetadataIn, i2)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6856:9:
				if rc != m_SQLITE_OK {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6857:11:
					Xvtab_set_error(tls, p, __ccgo_ts+10028, 0)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6858:11:
					if rc != m_SQLITE_OK {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6859:13:
						goto cleanup
					}
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6862:9:
				Xbitmap_and_inplace(tls, b, bmMetadata, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
				goto _5
			_5:
				;
				i2 = i2 + 1
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6867:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6867:14:
		i3 = 0
		for {
			if !(i3 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6868:7:
			if !(Xbitmap_get(tls, b, i3) != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6869:9:
				goto _6
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6870:8:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6873:7:
			switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Felement_type {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6874:7:
			case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6875:19:
				base_i = baseVectors + uintptr(libc.Uint64FromInt32(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)*4
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6877:9:
				switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdistance_metric {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6878:9:
				case int32(_VEC0_DISTANCE_METRIC_L2):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6879:11:
					result = _distance_l2_sqr_float(tls, base_i, queryVector, vector_column+16)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6881:11:
					break
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6883:9:
					fallthrough
				case int32(_VEC0_DISTANCE_METRIC_L1):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6884:11:
					result = float32(_distance_l1_f32(tls, base_i, queryVector, vector_column+16))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6886:11:
					break
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6888:9:
					fallthrough
				case int32(_VEC0_DISTANCE_METRIC_COSINE):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6889:11:
					result = _distance_cosine_float(tls, base_i, queryVector, vector_column+16)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6891:11:
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6894:9:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6896:7:
				fallthrough
			case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6897:18:
				base_i1 = baseVectors + uintptr(libc.Uint64FromInt32(i3)*(*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6899:9:
				switch (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdistance_metric {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6900:9:
				case int32(_VEC0_DISTANCE_METRIC_L2):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6901:11:
					result = _distance_l2_sqr_int8(tls, base_i1, queryVector, vector_column+16)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6903:11:
					break
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6905:9:
					fallthrough
				case int32(_VEC0_DISTANCE_METRIC_L1):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6906:11:
					result = float32(_distance_l1_int8(tls, base_i1, queryVector, vector_column+16))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6908:11:
					break
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6910:9:
					fallthrough
				case int32(_VEC0_DISTANCE_METRIC_COSINE):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6911:11:
					result = _distance_cosine_int8(tls, base_i1, queryVector, vector_column+16)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6913:11:
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6917:9:
				break
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6919:7:
				fallthrough
			case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6920:18:
				base_i2 = baseVectors + uintptr(libc.Uint64FromInt32(i3)*((*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions/libc.Uint64FromInt32(m_CHAR_BIT)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6922:9:
				result = _distance_hamming(tls, base_i2, queryVector, vector_column+16)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6924:9:
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6928:7:
			**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i3)*4)) = result
			goto _6
		_6:
			;
			i3 = i3 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6931:5:
		if hasDistanceConstraints != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6932:7:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6932:15:
			i4 = 0
			for {
				if !(i4 < argc) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6933:13:
				idx2 = int32(1) + i4*int32(4)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6934:14:
				kind2 = **(**int8)(__ccgo_up(idxStr + uintptr(idx2+0)))
				// TODO casts f64 to f32, is that a problem?
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6936:13:
				target = float32(libsqlite3.Xsqlite3_value_double(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i4)*8))))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6938:9:
				if int32(kind2) != int32(_VEC0_IDXSTR_KIND_KNN_DISTANCE_CONSTRAINT) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6939:11:
					goto _7
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6941:43:
				op = int32(**(**int8)(__ccgo_up(idxStr + uintptr(idx2+int32(1)))))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6943:9:
				switch op {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6944:11:
				case int32(_VEC0_DISTANCE_CONSTRAINT_GE):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6945:13:
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6945:21:
					i5 = 0
					for {
						if !(i5 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6946:15:
						if Xbitmap_get(tls, b, i5) != 0 && !(**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i5)*4)) >= target) {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6947:17:
							Xbitmap_set(tls, b, i5, 0)
						}
						goto _8
					_8:
						;
						i5 = i5 + 1
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6950:13:
					break
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6952:11:
					fallthrough
				case int32(_VEC0_DISTANCE_CONSTRAINT_GT):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6953:13:
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6953:21:
					i6 = 0
					for {
						if !(i6 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6954:15:
						if Xbitmap_get(tls, b, i6) != 0 && !(**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i6)*4)) > target) {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6955:17:
							Xbitmap_set(tls, b, i6, 0)
						}
						goto _9
					_9:
						;
						i6 = i6 + 1
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6958:13:
					break
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6960:11:
					fallthrough
				case int32(_VEC0_DISTANCE_CONSTRAINT_LE):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6961:13:
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6961:21:
					i7 = 0
					for {
						if !(i7 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6962:15:
						if Xbitmap_get(tls, b, i7) != 0 && !(**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i7)*4)) <= target) {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6963:17:
							Xbitmap_set(tls, b, i7, 0)
						}
						goto _10
					_10:
						;
						i7 = i7 + 1
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6966:13:
					break
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6968:11:
					fallthrough
				case int32(_VEC0_DISTANCE_CONSTRAINT_LT):
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6969:13:
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6969:21:
					i8 = 0
					for {
						if !(i8 < (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
							break
						}
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6970:15:
						if Xbitmap_get(tls, b, i8) != 0 && !(**(**Tf32)(__ccgo_up(chunk_distances + uintptr(i8)*4)) < target) {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6971:17:
							Xbitmap_set(tls, b, i8, 0)
						}
						goto _11
					_11:
						;
						i8 = i8 + 1
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6974:13:
					break
				}
				goto _7
			_7:
				;
				i4 = i4 + 1
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6981:5:
		if k <= int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) {
			v12 = k
		} else {
			v12 = int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size)
		}
		Xmin_idx(tls, chunk_distances, (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size, b, chunk_topk_idxs, int32(v12), bTaken, bp+144)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6985:5:
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
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6990:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6990:14:
		i9 = 0
		for {
			if !(int64(i9) < **(**Ti64)(__ccgo_up(bp + 152))) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6991:7:
			**(**Ti64)(__ccgo_up(topk_rowids + uintptr(i9)*8)) = **(**Ti64)(__ccgo_up(tmp_topk_rowids + uintptr(i9)*8))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6992:7:
			**(**Tf32)(__ccgo_up(topk_distances + uintptr(i9)*4)) = **(**Tf32)(__ccgo_up(tmp_topk_distances + uintptr(i9)*4))
			goto _16
		_16:
			;
			i9 = i9 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6994:5:
		k_used = **(**Ti64)(__ccgo_up(bp + 152))
		// blobVectors is always opened with read-only permissions, so this never
		// fails.
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6997:5:
		libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:6998:5:
		**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7001:3:
	**(**uintptr)(__ccgo_up(out_topk_rowids)) = topk_rowids
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7002:3:
	**(**uintptr)(__ccgo_up(out_topk_distances)) = topk_distances
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7003:3:
	**(**Ti64)(__ccgo_up(out_used)) = k_used
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7004:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7006:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7007:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7008:5:
		libsqlite3.Xsqlite3_free(tls, topk_rowids)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7009:5:
		libsqlite3.Xsqlite3_free(tls, topk_distances)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7011:3:
	libsqlite3.Xsqlite3_free(tls, chunk_topk_idxs)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7012:3:
	libsqlite3.Xsqlite3_free(tls, tmp_topk_rowids)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7013:3:
	libsqlite3.Xsqlite3_free(tls, tmp_topk_distances)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7014:3:
	libsqlite3.Xsqlite3_free(tls, b)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7015:3:
	libsqlite3.Xsqlite3_free(tls, bTaken)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7016:3:
	libsqlite3.Xsqlite3_free(tls, bmRowids)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7017:3:
	libsqlite3.Xsqlite3_free(tls, baseVectors)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7018:3:
	libsqlite3.Xsqlite3_free(tls, chunk_distances)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7019:3:
	libsqlite3.Xsqlite3_free(tls, bmMetadata)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7020:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7020:11:
	i10 = 0
	for {
		if !(i10 < int32(m_VEC0_MAX_METADATA_COLUMNS)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7021:5:
		libsqlite3.Xsqlite3_blob_close(tls, (**(**[16]uintptr)(__ccgo_up(bp + 8)))[i10])
		goto _17
	_17:
		;
		i10 = i10 + 1
	}
	// blobVectors is always opened with read-only permissions, so this never
	// fails.
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7025:3:
	libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7026:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7029:5:
func Xvec0Filter_knn(tls *libc.TLS, pCur uintptr, p uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7030:72:
	var aMetadataIn, arrayRowidsIn, item2, knn_data, s, vector_column uintptr
	var entry3 TVec0MetadataInTextEntry
	var i, i1, k_idx, metadata_idx, n, query_idx, rc, rc1, rowid_in_idx, vectorColumnIdx int32
	var i2, j Tsize_t
	var k Ti64
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+24 */ _VectorElementType
	var _ /* entry at bp+104 */ uintptr
	var _ /* entry at bp+120 */ uintptr
	var _ /* entry at bp+128 */ TVec0MetadataInTextEntry
	var _ /* item at bp+48 */ uintptr
	var _ /* item at bp+64 */ TVec0MetadataIn
	var _ /* k_used at bp+160 */ Ti64
	var _ /* pzError at bp+40 */ uintptr
	var _ /* queryVector at bp+8 */ uintptr
	var _ /* queryVectorCleanup at bp+32 */ Tvector_cleanup
	var _ /* rowid at bp+56 */ Ti64
	var _ /* stmtChunks at bp+0 */ uintptr
	var _ /* topk_distances at bp+152 */ uintptr
	var _ /* topk_rowids at bp+144 */ uintptr
	var _ /* v at bp+112 */ Ti64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aMetadataIn, arrayRowidsIn, entry3, i, i1, i2, item2, j, k, k_idx, knn_data, metadata_idx, n, query_idx, rc, rc1, rowid_in_idx, s, vectorColumnIdx, vector_column
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7031:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7035:7:
	vectorColumnIdx = idxNum
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7036:33:
	vector_column = p + 608 + uintptr(vectorColumnIdx)*32
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7039:16:
	arrayRowidsIn = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7040:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7044:18:
	**(**Tvector_cleanup)(__ccgo_up(bp + 32)) = __ccgo_fp(Xvector_cleanup_noop)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7046:3:
	knn_data = libsqlite3.Xsqlite3_malloc(tls, int32(40))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7047:3:
	if !(knn_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7048:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7050:3:
	libc.Xmemset(tls, knn_data, 0, uint64(40))
	// array of `struct Vec0MetadataIn`, IF there are any `xxx in (...)` metadata constraints
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7052:16:
	aMetadataIn = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7054:7:
	query_idx = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7055:7:
	k_idx = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7056:7:
	rowid_in_idx = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7057:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7057:11:
	i = 0
	for {
		if !(i < argc) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7058:5:
		if int32(**(**int8)(__ccgo_up(idxStr + uintptr(int32(1)+i*int32(4))))) == int32(_VEC0_IDXSTR_KIND_KNN_MATCH) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7059:7:
			query_idx = i
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7061:5:
		if int32(**(**int8)(__ccgo_up(idxStr + uintptr(int32(1)+i*int32(4))))) == int32(_VEC0_IDXSTR_KIND_KNN_K) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7062:7:
			k_idx = i
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7064:5:
		if int32(**(**int8)(__ccgo_up(idxStr + uintptr(int32(1)+i*int32(4))))) == int32(_VEC0_IDXSTR_KIND_KNN_ROWID_IN) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7065:7:
			rowid_in_idx = i
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7068:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7069:3:
	// make sure the query vector matches the vector column (type dimensions etc.)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7072:3:
	rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv + uintptr(query_idx)*8)), bp+8, bp+16, bp+24, bp+32, bp+40)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7075:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7076:5:
		Xvtab_set_error(tls, p, __ccgo_ts+10061, libc.VaList(bp+176, (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fname_length, (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fname, **(**uintptr)(__ccgo_up(bp + 40))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7079:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7080:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7082:3:
	if **(**_VectorElementType)(__ccgo_up(bp + 24)) != (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Felement_type {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7083:5:
		Xvtab_set_error(tls, p, __ccgo_ts+10110, libc.VaList(bp+176, (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fname_length, (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fname, Xvector_subtype_name(tls, (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Felement_type), Xvector_subtype_name(tls, **(**_VectorElementType)(__ccgo_up(bp + 24)))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7090:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7091:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7093:3:
	if **(**Tsize_t)(__ccgo_up(bp + 16)) != (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7094:5:
		Xvtab_set_error(tls, p, __ccgo_ts+10205, libc.VaList(bp+176, (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fname_length, (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fname, (*TVectorColumnDefinition)(unsafe.Pointer(vector_column)).Fdimensions, **(**Tsize_t)(__ccgo_up(bp + 16))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7100:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7101:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7104:7:
	k = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + uintptr(k_idx)*8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7105:3:
	if k < 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7106:5:
		Xvtab_set_error(tls, p, __ccgo_ts+10304, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7108:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7109:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7112:3:
	if k > int64(m_SQLITE_VEC_VEC0_K_MAX) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7113:5:
		Xvtab_set_error(tls, p, __ccgo_ts+10363, libc.VaList(bp+176, k, int32(m_SQLITE_VEC_VEC0_K_MAX)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7117:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7118:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7121:3:
	if k == 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7122:5:
		(*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Fk = 0
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7123:5:
		(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data = knn_data
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7124:5:
		(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC0_QUERY_PLAN_KNN)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7125:5:
		rc = m_SQLITE_OK
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7126:5:
		goto cleanup
	}
	// handle when a `rowid in (...)` operation was provided
	// Array of all the rowids that appear in any `rowid in (...)` constraint.
	// NULL if none were provided, which means a "full" scan.
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7133:3:
	if rowid_in_idx >= 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7136:5:
		arrayRowidsIn = libsqlite3.Xsqlite3_malloc(tls, int32(32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7137:5:
		if !(arrayRowidsIn != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7138:7:
			rc1 = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7139:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7141:5:
		libc.Xmemset(tls, arrayRowidsIn, 0, uint64(32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7143:5:
		rc1 = Xarray_init(tls, arrayRowidsIn, uint64(8), uint64(32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7144:5:
		if rc1 != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7145:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7147:5:
		rc1 = libsqlite3.Xsqlite3_vtab_in_first(tls, **(**uintptr)(__ccgo_up(argv + uintptr(rowid_in_idx)*8)), bp+48)
		for {
			if !(rc1 == m_SQLITE_OK && **(**uintptr)(__ccgo_up(bp + 48)) != 0) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7150:7:
			if (*Tvec0_vtab)(unsafe.Pointer(p)).FpkIsText != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7151:9:
				rc1 = Xvec0_rowid_from_id(tls, p, **(**uintptr)(__ccgo_up(bp + 48)), bp+56)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7152:9:
				if rc1 != m_SQLITE_OK {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7153:11:
					goto cleanup
				}
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7156:9:
				**(**Ti64)(__ccgo_up(bp + 56)) = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(bp + 48)))
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7158:7:
			rc1 = Xarray_append(tls, arrayRowidsIn, bp+56)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7159:7:
			if rc1 != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7160:9:
				goto cleanup
			}
			goto _2
		_2:
			;
			rc1 = libsqlite3.Xsqlite3_vtab_in_next(tls, **(**uintptr)(__ccgo_up(argv + uintptr(rowid_in_idx)*8)), bp+48)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7163:5:
		if rc1 != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7164:7:
			Xvtab_set_error(tls, p, __ccgo_ts+10431, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7165:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7167:5:
		libc.Xqsort(tls, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Fz, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Flength, (*TArray)(unsafe.Pointer(arrayRowidsIn)).Felement_size, __ccgo_fp(X_cmp))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7173:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7173:11:
	i1 = 0
	for {
		if !(i1 < argc) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7174:5:
		if !(int32(**(**int8)(__ccgo_up(idxStr + uintptr(int32(1)+i1*int32(4))))) == int32(_VEC0_IDXSTR_KIND_METADATA_CONSTRAINT) && int32(**(**int8)(__ccgo_up(idxStr + uintptr(int32(1)+i1*int32(4)+int32(2))))) == int32(_VEC0_METADATA_OPERATOR_IN)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7175:7:
			goto _3
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7177:9:
		metadata_idx = int32(**(**int8)(__ccgo_up(idxStr + uintptr(int32(1)+i1*int32(4)+int32(1))))) - int32('A')
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7178:5:
		if !(aMetadataIn != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7179:7:
			aMetadataIn = libsqlite3.Xsqlite3_malloc(tls, int32(32))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7180:7:
			if !(aMetadataIn != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7181:9:
				rc = int32(m_SQLITE_NOMEM)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7182:9:
				goto cleanup
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7184:7:
			libc.Xmemset(tls, aMetadataIn, 0, uint64(32))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7185:7:
			rc = Xarray_init(tls, aMetadataIn, uint64(40), uint64(8))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7186:7:
			if rc != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7187:9:
				goto cleanup
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7192:5:
		libc.Xmemset(tls, bp+64, 0, uint64(40))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7193:5:
		(**(**TVec0MetadataIn)(__ccgo_up(bp + 64))).Fmetadata_idx = metadata_idx
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7194:5:
		(**(**TVec0MetadataIn)(__ccgo_up(bp + 64))).Fargv_idx = i1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7196:5:
		switch (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7197:7:
		case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7198:9:
			rc = Xarray_init(tls, bp+64+8, uint64(8), uint64(16))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7199:9:
			if rc != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7200:11:
				goto cleanup
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7203:9:
			rc = libsqlite3.Xsqlite3_vtab_in_first(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i1)*8)), bp+104)
			for {
				if !(rc == m_SQLITE_OK && **(**uintptr)(__ccgo_up(bp + 104)) != 0) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7204:15:
				**(**Ti64)(__ccgo_up(bp + 112)) = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(bp + 104)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7205:11:
				rc = Xarray_append(tls, bp+64+8, bp+112)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7206:11:
				if rc != m_SQLITE_OK {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7207:13:
					goto cleanup
				}
				goto _4
			_4:
				;
				rc = libsqlite3.Xsqlite3_vtab_in_next(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i1)*8)), bp+104)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7211:9:
			if rc != int32(m_SQLITE_DONE) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7212:11:
				Xvtab_set_error(tls, p, __ccgo_ts+10469, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7213:11:
				goto cleanup
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7216:9:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7218:7:
			fallthrough
		case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7219:9:
			rc = Xarray_init(tls, bp+64+8, uint64(16), uint64(16))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7220:9:
			if rc != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7221:11:
				goto cleanup
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7224:9:
			rc = libsqlite3.Xsqlite3_vtab_in_first(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i1)*8)), bp+120)
			for {
				if !(rc == m_SQLITE_OK && **(**uintptr)(__ccgo_up(bp + 120)) != 0) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7225:22:
				s = libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(bp + 120)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7226:15:
				n = libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(bp + 120)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7229:11:
				(**(**TVec0MetadataInTextEntry)(__ccgo_up(bp + 128))).FzString = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1874, libc.VaList(bp+176, n, s))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7230:11:
				if !((**(**TVec0MetadataInTextEntry)(__ccgo_up(bp + 128))).FzString != 0) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7231:13:
					rc = int32(m_SQLITE_NOMEM)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7232:13:
					goto cleanup
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7234:11:
				(**(**TVec0MetadataInTextEntry)(__ccgo_up(bp + 128))).Fn = n
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7235:11:
				rc = Xarray_append(tls, bp+64+8, bp+128)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7236:11:
				if rc != m_SQLITE_OK {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7237:13:
					goto cleanup
				}
				goto _5
			_5:
				;
				rc = libsqlite3.Xsqlite3_vtab_in_next(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i1)*8)), bp+120)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7241:9:
			if rc != int32(m_SQLITE_DONE) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7242:11:
				Xvtab_set_error(tls, p, __ccgo_ts+10530, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7243:11:
				goto cleanup
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7246:9:
			break
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7248:7:
			fallthrough
		default:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7249:9:
			Xvtab_set_error(tls, p, __ccgo_ts+10588, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7250:9:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7254:5:
		rc = Xarray_append(tls, aMetadataIn, bp+64)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7255:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7256:7:
			goto cleanup
		}
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7261:3:
	rc = Xvec0_chunks_iter(tls, p, idxStr, argc, argv, bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7262:3:
	if rc != m_SQLITE_OK {
		// IMP: V06942_23781
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7264:5:
		Xvtab_set_error(tls, p, __ccgo_ts+10614, libc.VaList(bp+176, libsqlite3.Xsqlite3_errmsg(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7266:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7269:7:
	**(**uintptr)(__ccgo_up(bp + 144)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7270:7:
	**(**uintptr)(__ccgo_up(bp + 152)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7271:7:
	**(**Ti64)(__ccgo_up(bp + 160)) = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7272:3:
	rc = Xvec0Filter_knn_chunks_iter(tls, p, **(**uintptr)(__ccgo_up(bp)), vector_column, vectorColumnIdx, arrayRowidsIn, aMetadataIn, idxStr, argc, argv, **(**uintptr)(__ccgo_up(bp + 8)), k, bp+144, bp+152, bp+160)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7275:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7276:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7279:3:
	(*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Fcurrent_idx = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7280:3:
	(*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Fk = k
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7281:3:
	(*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Frowids = **(**uintptr)(__ccgo_up(bp + 144))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7282:3:
	(*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances = **(**uintptr)(__ccgo_up(bp + 152))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7283:3:
	(*Tvec0_query_knn_data)(unsafe.Pointer(knn_data)).Fk_used = **(**Ti64)(__ccgo_up(bp + 160))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7285:3:
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data = knn_data
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7286:3:
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC0_QUERY_PLAN_KNN)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7287:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7289:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7290:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7291:3:
	Xarray_cleanup(tls, arrayRowidsIn)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7292:3:
	libsqlite3.Xsqlite3_free(tls, arrayRowidsIn)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7293:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 32)))(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7294:3:
	if aMetadataIn != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7295:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7295:16:
		i2 = uint64(0)
		for {
			if !(i2 < (*TArray)(unsafe.Pointer(aMetadataIn)).Flength) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7296:28:
			item2 = (*TArray)(unsafe.Pointer(aMetadataIn)).Fz + uintptr(i2)*40
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7297:7:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7297:18:
			j = uint64(0)
			for {
				if !(j < (*TVec0MetadataIn)(unsafe.Pointer(item2)).Farray.Flength) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7298:9:
				if (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr((*TVec0MetadataIn)(unsafe.Pointer(item2)).Fmetadata_idx)*24))).Fkind == int32(_VEC0_METADATA_COLUMN_KIND_TEXT) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7299:42:
					entry3 = **(**TVec0MetadataInTextEntry)(__ccgo_up((*TVec0MetadataIn)(unsafe.Pointer(item2)).Farray.Fz + uintptr(j)*16))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7300:11:
					libsqlite3.Xsqlite3_free(tls, entry3.FzString)
				}
				goto _7
			_7:
				;
				j = j + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7303:7:
			Xarray_cleanup(tls, item2+8)
			goto _6
		_6:
			;
			i2 = i2 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7305:5:
		Xarray_cleanup(tls, aMetadataIn)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7308:3:
	libsqlite3.Xsqlite3_free(tls, aMetadataIn)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7310:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7311:5:
		libsqlite3.Xsqlite3_free(tls, knn_data)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7314:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7317:5:
func Xvec0Filter_fullscan(tls *libc.TLS, p uintptr, pCur uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7317:58:
	var fullscan_data, zSql uintptr
	var rc int32
	_, _, _ = fullscan_data, rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7322:3:
	fullscan_data = libsqlite3.Xsqlite3_malloc(tls, int32(16))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7323:3:
	if !(fullscan_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7324:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7326:3:
	libc.Xmemset(tls, fullscan_data, 0, uint64(16))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7328:3:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+10644, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7332:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7333:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7334:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7336:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), fullscan_data, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7337:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7338:3:
	if rc != m_SQLITE_OK {
		// IMP: V09901_26739
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7340:5:
		Xvtab_set_error(tls, p, __ccgo_ts+10714, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7342:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7345:3:
	rc = libsqlite3.Xsqlite3_step(tls, (*Tvec0_query_fullscan_data)(unsafe.Pointer(fullscan_data)).Frowids_stmt)
	// DONE when there's no rowids, ROW when there are, both "success"
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7348:3:
	if !(rc == int32(m_SQLITE_ROW) || rc == int32(m_SQLITE_DONE)) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7349:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7352:3:
	(*Tvec0_query_fullscan_data)(unsafe.Pointer(fullscan_data)).Fdone = libc.BoolInt8(rc == int32(m_SQLITE_DONE))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7353:3:
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC0_QUERY_PLAN_FULLSCAN)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7354:3:
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data = fullscan_data
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7355:3:
	return m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7357:1:
	goto error
error:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7358:3:
	Xvec0_query_fullscan_data_clear(tls, fullscan_data)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7359:3:
	libsqlite3.Xsqlite3_free(tls, fullscan_data)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7360:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7363:5:
func Xvec0Filter_point(tls *libc.TLS, pCur uintptr, p uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7364:44:
	var i, rc int32
	var point_data uintptr
	var _ /* rowid at bp+0 */ Ti64
	_, _, _ = i, point_data, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7366:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7368:32:
	point_data = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7370:3:
	point_data = libsqlite3.Xsqlite3_malloc(tls, int32(144))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7371:3:
	if !(point_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7372:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7373:5:
		goto error
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7375:3:
	libc.Xmemset(tls, point_data, 0, uint64(144))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7377:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FpkIsText != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7378:5:
		rc = Xvec0_rowid_from_id(tls, p, **(**uintptr)(__ccgo_up(argv)), bp)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7379:5:
		if rc == int32(m_SQLITE_EMPTY) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7380:7:
			goto eof
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7382:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7383:7:
			goto error
		}
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7386:5:
		**(**Ti64)(__ccgo_up(bp)) = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv)))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7389:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7389:12:
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7390:5:
		rc = Xvec0_get_vector_data(tls, p, **(**Ti64)(__ccgo_up(bp)), i, point_data+8+uintptr(i)*8, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7391:5:
		if rc == int32(m_SQLITE_EMPTY) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7392:7:
			goto eof
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7394:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7395:7:
			goto error
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7399:3:
	(*Tvec0_query_point_data)(unsafe.Pointer(point_data)).Frowid = **(**Ti64)(__ccgo_up(bp))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7400:3:
	(*Tvec0_query_point_data)(unsafe.Pointer(point_data)).Fdone = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7401:3:
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data = point_data
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7402:3:
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC0_QUERY_PLAN_POINT)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7403:3:
	return m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7405:1:
	goto eof
eof:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7406:3:
	(*Tvec0_query_point_data)(unsafe.Pointer(point_data)).Frowid = **(**Ti64)(__ccgo_up(bp))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7407:3:
	(*Tvec0_query_point_data)(unsafe.Pointer(point_data)).Fdone = int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7408:3:
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data = point_data
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7409:3:
	(*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC0_QUERY_PLAN_POINT)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7410:3:
	return m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7412:1:
	goto error
error:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7413:3:
	Xvec0_query_point_data_clear(tls, point_data)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7414:3:
	libsqlite3.Xsqlite3_free(tls, point_data)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7415:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7418:12:
func _vec0Filter(tls *libc.TLS, pVtabCursor uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7419:75:
	var idxStrLength, numValueEntries int32
	var p, pCur uintptr
	var query_plan int8
	_, _, _, _, _ = idxStrLength, numValueEntries, p, pCur, query_plan
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7420:13:
	p = (*Tsqlite3_vtab_cursor)(unsafe.Pointer(pVtabCursor)).FpVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7421:15:
	pCur = pVtabCursor
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7422:3:
	Xvec0_cursor_clear(tls, pCur)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7424:7:
	idxStrLength = libc.Int32FromUint64(libc.Xstrlen(tls, idxStr))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7425:3:
	if idxStrLength <= 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7426:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7428:3:
	if (idxStrLength-int32(1))%int32(4) != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7429:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7431:7:
	numValueEntries = (idxStrLength - int32(1)) / int32(4)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7432:3:
	if numValueEntries != argc {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7433:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7436:8:
	query_plan = **(**int8)(__ccgo_up(idxStr))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7437:3:
	switch int32(query_plan) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7438:5:
	case int32(_VEC0_QUERY_PLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7439:7:
		return Xvec0Filter_fullscan(tls, p, pCur)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7440:5:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7441:7:
		return Xvec0Filter_knn(tls, pCur, p, idxNum, idxStr, argc, argv)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7442:5:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_POINT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7443:7:
		return Xvec0Filter_point(tls, pCur, p, argc, argv)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7444:5:
		fallthrough
	default:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7445:7:
		Xvtab_set_error(tls, (*Tsqlite3_vtab_cursor)(unsafe.Pointer(pVtabCursor)).FpVtab, __ccgo_ts+10745, libc.VaList(bp+8, idxStr))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7446:7:
		return int32(m_SQLITE_ERROR)
	}
	return r
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7450:12:
func _vec0Rowid(tls *libc.TLS, cur uintptr, pRowid uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7450:70:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7451:15:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7452:3:
	switch (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7453:3:
	case int32(_VEC0_QUERY_PLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7454:5:
		**(**Tsqlite_int64)(__ccgo_up(pRowid)) = libsqlite3.Xsqlite3_column_int64(tls, (*Tvec0_query_fullscan_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data)).Frowids_stmt, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7455:5:
		return m_SQLITE_OK
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7457:3:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_POINT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7458:5:
		**(**Tsqlite_int64)(__ccgo_up(pRowid)) = (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7459:5:
		return m_SQLITE_OK
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7461:3:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7462:5:
		Xvtab_set_error(tls, (*Tsqlite3_vtab_cursor)(unsafe.Pointer(cur)).FpVtab, __ccgo_ts+10765, libc.VaList(bp+8, (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7466:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7469:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7472:12:
func _vec0Next(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7472:47:
	var pCur uintptr
	var rc int32
	_, _ = pCur, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7473:15:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7474:3:
	switch (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7475:3:
	case int32(_VEC0_QUERY_PLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7476:5:
		if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7477:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7479:9:
		rc = libsqlite3.Xsqlite3_step(tls, (*Tvec0_query_fullscan_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data)).Frowids_stmt)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7480:5:
		if rc == int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7481:7:
			(*Tvec0_query_fullscan_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data)).Fdone = int8(1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7482:7:
			return m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7484:5:
		if rc == int32(m_SQLITE_ROW) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7485:7:
			return m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7487:5:
		return int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7489:3:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7490:5:
		if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7491:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7494:5:
		(*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx = (*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx + 1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7495:5:
		return m_SQLITE_OK
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7497:3:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_POINT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7498:5:
		if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7499:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7501:5:
		(*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Fdone = int32(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7502:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7505:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7508:12:
func _vec0Eof(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7508:46:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7509:15:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7510:3:
	switch (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7511:3:
	case int32(_VEC0_QUERY_PLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7512:5:
		if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7513:7:
			return int32(1)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7515:5:
		return int32((*Tvec0_query_fullscan_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data)).Fdone)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7517:3:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7518:5:
		if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7519:7:
			return int32(1)
		}
		// return (pCur->knn_data->current_idx >= pCur->knn_data->k) ||
		// (pCur->knn_data->distances[pCur->knn_data->current_idx] == FLT_MAX);
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7523:5:
		return libc.BoolInt32((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx >= (*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fk_used)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7525:3:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_POINT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7526:5:
		if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7527:7:
			return int32(1)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7529:5:
		return (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Fdone
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7532:3:
	return int32(1)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7535:12:
func _vec0Column_fullscan(tls *libc.TLS, pVtab uintptr, pCur uintptr, context uintptr, i int32) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7536:65:
	var auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rc3, vector_idx int32
	var rowid Ti64
	var zErr uintptr
	var _ /* sz at bp+8 */ int32
	var _ /* v at bp+0 */ uintptr
	var _ /* v at bp+16 */ uintptr
	var _ /* v at bp+24 */ uintptr
	_, _, _, _, _, _, _, _, _, _ = auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rc3, rowid, vector_idx, zErr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7537:3:
	if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7538:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+10841, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7540:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7542:7:
	rowid = libsqlite3.Xsqlite3_column_int64(tls, (*Tvec0_query_fullscan_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Ffullscan_data)).Frowids_stmt, 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7543:3:
	if i == m_VEC0_COLUMN_ID {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7544:5:
		return Xvec0_result_id(tls, pVtab, context, rowid)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7546:8:
		if Xvec0_column_idx_is_vector(tls, pVtab, i) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7549:9:
			vector_idx = Xvec0_column_idx_to_vector_idx(tls, pVtab, i)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7550:9:
			rc = Xvec0_get_vector_data(tls, pVtab, rowid, vector_idx, bp, bp+8)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7551:5:
			if rc != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7552:7:
				return rc
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7554:5:
			libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), **(**int32)(__ccgo_up(bp + 8)), __ccgo_fp(libsqlite3.Xsqlite3_free))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7555:5:
			libsqlite3.Xsqlite3_result_subtype(tls, context, libc.Uint32FromInt32((**(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_idx)*32))).Felement_type))
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7559:8:
			if i == Xvec0_column_distance_idx(tls, pVtab) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7560:5:
				libsqlite3.Xsqlite3_result_null(tls, context)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7562:8:
				if Xvec0_column_idx_is_partition(tls, pVtab, i) != 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7563:9:
					partition_idx = Xvec0_column_idx_to_partition_idx(tls, pVtab, i)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7565:9:
					rc1 = Xvec0_get_partition_value_for_rowid(tls, pVtab, rowid, partition_idx, bp+16)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7566:5:
					if rc1 == m_SQLITE_OK {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7567:7:
						libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 16)))
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7568:7:
						libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 16)))
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7570:7:
						libsqlite3.Xsqlite3_result_error_code(tls, context, rc1)
					}
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7573:8:
					if Xvec0_column_idx_is_auxiliary(tls, pVtab, i) != 0 {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7574:9:
						auxiliary_idx = Xvec0_column_idx_to_auxiliary_idx(tls, pVtab, i)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7576:9:
						rc2 = Xvec0_get_auxiliary_value_for_rowid(tls, pVtab, rowid, auxiliary_idx, bp+24)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7577:5:
						if rc2 == m_SQLITE_OK {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7578:7:
							libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 24)))
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7579:7:
							libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
						} else {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7581:7:
							libsqlite3.Xsqlite3_result_error_code(tls, context, rc2)
						}
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7585:8:
						if Xvec0_column_idx_is_metadata(tls, pVtab, i) != 0 {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7586:5:
							if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7587:7:
								return m_SQLITE_OK
							}
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7589:9:
							metadata_idx = Xvec0_column_idx_to_metadata_idx(tls, pVtab, i)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7590:9:
							rc3 = Xvec0_result_metadata_value_for_rowid(tls, pVtab, rowid, metadata_idx, context)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7591:5:
							if rc3 != m_SQLITE_OK {
								// IMP: V15466_32305
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7593:18:
								zErr = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+10891, libc.VaList(bp+40, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname_length, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname, rowid))
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7598:7:
								if zErr != 0 {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7599:9:
									libsqlite3.Xsqlite3_result_error(tls, context, zErr, -int32(1))
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7600:9:
									libsqlite3.Xsqlite3_free(tls, zErr)
								} else {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7602:9:
									libsqlite3.Xsqlite3_result_error_nomem(tls, context)
								}
							}
						}
					}
				}
			}
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7607:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7610:12:
func _vec0Column_point(tls *libc.TLS, pVtab uintptr, pCur uintptr, context uintptr, i int32) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7611:62:
	var auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, vector_idx int32
	var rowid, rowid1, rowid2 Ti64
	var zErr uintptr
	var _ /* v at bp+0 */ uintptr
	var _ /* v at bp+8 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _ = auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rowid, rowid1, rowid2, vector_idx, zErr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7612:3:
	if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7613:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+10954, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7615:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7617:3:
	if i == m_VEC0_COLUMN_ID {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7618:5:
		return Xvec0_result_id(tls, pVtab, context, (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7620:8:
		if i == Xvec0_column_distance_idx(tls, pVtab) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7621:5:
			libsqlite3.Xsqlite3_result_null(tls, context)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7622:5:
			return m_SQLITE_OK
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7624:8:
			if Xvec0_column_idx_is_vector(tls, pVtab, i) != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7625:5:
				if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7626:7:
					libsqlite3.Xsqlite3_result_null(tls, context)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7627:7:
					return m_SQLITE_OK
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7629:9:
				vector_idx = Xvec0_column_idx_to_vector_idx(tls, pVtab, i)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7630:5:
				libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data + 8 + uintptr(vector_idx)*8)), libc.Int32FromUint64(Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_idx)*32)))), uintptr(-libc.Int32FromInt32(1)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7634:5:
				libsqlite3.Xsqlite3_result_subtype(tls, context, libc.Uint32FromInt32((**(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_idx)*32))).Felement_type))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7636:5:
				return m_SQLITE_OK
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7638:8:
				if Xvec0_column_idx_is_partition(tls, pVtab, i) != 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7639:5:
					if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7640:7:
						return m_SQLITE_OK
					}
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7642:9:
					partition_idx = Xvec0_column_idx_to_partition_idx(tls, pVtab, i)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7643:9:
					rowid = (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7645:9:
					rc = Xvec0_get_partition_value_for_rowid(tls, pVtab, rowid, partition_idx, bp)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7646:5:
					if rc == m_SQLITE_OK {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7647:7:
						libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp)))
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7648:7:
						libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp)))
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7650:7:
						libsqlite3.Xsqlite3_result_error_code(tls, context, rc)
					}
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7653:8:
					if Xvec0_column_idx_is_auxiliary(tls, pVtab, i) != 0 {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7654:5:
						if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7655:7:
							return m_SQLITE_OK
						}
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7657:9:
						rowid1 = (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7658:9:
						auxiliary_idx = Xvec0_column_idx_to_auxiliary_idx(tls, pVtab, i)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7660:9:
						rc1 = Xvec0_get_auxiliary_value_for_rowid(tls, pVtab, rowid1, auxiliary_idx, bp+8)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7661:5:
						if rc1 == m_SQLITE_OK {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7662:7:
							libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 8)))
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7663:7:
							libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 8)))
						} else {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7665:7:
							libsqlite3.Xsqlite3_result_error_code(tls, context, rc1)
						}
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7669:8:
						if Xvec0_column_idx_is_metadata(tls, pVtab, i) != 0 {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7670:5:
							if libsqlite3.Xsqlite3_vtab_nochange(tls, context) != 0 {
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7671:7:
								return m_SQLITE_OK
							}
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7673:9:
							rowid2 = (*Tvec0_query_point_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fpoint_data)).Frowid
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7674:9:
							metadata_idx = Xvec0_column_idx_to_metadata_idx(tls, pVtab, i)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7675:9:
							rc2 = Xvec0_result_metadata_value_for_rowid(tls, pVtab, rowid2, metadata_idx, context)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7676:5:
							if rc2 != m_SQLITE_OK {
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7677:18:
								zErr = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+10891, libc.VaList(bp+24, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname_length, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname, rowid2))
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7682:7:
								if zErr != 0 {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7683:9:
									libsqlite3.Xsqlite3_result_error(tls, context, zErr, -int32(1))
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7684:9:
									libsqlite3.Xsqlite3_free(tls, zErr)
								} else {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7686:9:
									libsqlite3.Xsqlite3_result_error_nomem(tls, context)
								}
							}
						}
					}
				}
			}
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7691:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7694:12:
func _vec0Column_knn(tls *libc.TLS, pVtab uintptr, pCur uintptr, context uintptr, i int32) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7695:60:
	var auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rc3, vector_idx int32
	var rowid, rowid1, rowid2, rowid3 Ti64
	var zErr uintptr
	var _ /* out at bp+0 */ uintptr
	var _ /* sz at bp+8 */ int32
	var _ /* v at bp+16 */ uintptr
	var _ /* v at bp+24 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = auxiliary_idx, metadata_idx, partition_idx, rc, rc1, rc2, rc3, rowid, rowid1, rowid2, rowid3, vector_idx, zErr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7696:3:
	if !((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7697:5:
		libsqlite3.Xsqlite3_result_error(tls, context, __ccgo_ts+11001, -int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7699:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7701:3:
	if i == m_VEC0_COLUMN_ID {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7702:9:
		rowid = **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7703:5:
		return Xvec0_result_id(tls, pVtab, context, rowid)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7705:8:
		if i == Xvec0_column_distance_idx(tls, pVtab) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7706:5:
			libsqlite3.Xsqlite3_result_double(tls, context, float64(**(**Tf32)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fdistances + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*4))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7708:5:
			return m_SQLITE_OK
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7710:8:
			if Xvec0_column_idx_is_vector(tls, pVtab, i) != 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7713:9:
				vector_idx = Xvec0_column_idx_to_vector_idx(tls, pVtab, i)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7714:9:
				rc = Xvec0_get_vector_data(tls, pVtab, **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8)), vector_idx, bp, bp+8)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7717:5:
				if rc != m_SQLITE_OK {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7718:7:
					return rc
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7720:5:
				libsqlite3.Xsqlite3_result_blob(tls, context, **(**uintptr)(__ccgo_up(bp)), **(**int32)(__ccgo_up(bp + 8)), __ccgo_fp(libsqlite3.Xsqlite3_free))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7721:5:
				libsqlite3.Xsqlite3_result_subtype(tls, context, libc.Uint32FromInt32((**(**TVectorColumnDefinition)(__ccgo_up(pVtab + 608 + uintptr(vector_idx)*32))).Felement_type))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7723:5:
				return m_SQLITE_OK
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7725:8:
				if Xvec0_column_idx_is_partition(tls, pVtab, i) != 0 {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7726:9:
					partition_idx = Xvec0_column_idx_to_partition_idx(tls, pVtab, i)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7727:9:
					rowid1 = **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7729:9:
					rc1 = Xvec0_get_partition_value_for_rowid(tls, pVtab, rowid1, partition_idx, bp+16)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7730:5:
					if rc1 == m_SQLITE_OK {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7731:7:
						libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 16)))
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7732:7:
						libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 16)))
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7734:7:
						libsqlite3.Xsqlite3_result_error_code(tls, context, rc1)
					}
				} else {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7737:8:
					if Xvec0_column_idx_is_auxiliary(tls, pVtab, i) != 0 {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7738:9:
						auxiliary_idx = Xvec0_column_idx_to_auxiliary_idx(tls, pVtab, i)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7739:9:
						rowid2 = **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8))
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7741:9:
						rc2 = Xvec0_get_auxiliary_value_for_rowid(tls, pVtab, rowid2, auxiliary_idx, bp+24)
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7742:5:
						if rc2 == m_SQLITE_OK {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7743:7:
							libsqlite3.Xsqlite3_result_value(tls, context, **(**uintptr)(__ccgo_up(bp + 24)))
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7744:7:
							libsqlite3.Xsqlite3_value_free(tls, **(**uintptr)(__ccgo_up(bp + 24)))
						} else {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7746:7:
							libsqlite3.Xsqlite3_result_error_code(tls, context, rc2)
						}
					} else {
						// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7750:8:
						if Xvec0_column_idx_is_metadata(tls, pVtab, i) != 0 {
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7751:9:
							metadata_idx = Xvec0_column_idx_to_metadata_idx(tls, pVtab, i)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7752:9:
							rowid3 = **(**Ti64)(__ccgo_up((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tvec0_query_knn_data)(unsafe.Pointer((*Tvec0_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*8))
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7753:9:
							rc3 = Xvec0_result_metadata_value_for_rowid(tls, pVtab, rowid3, metadata_idx, context)
							// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7754:5:
							if rc3 != m_SQLITE_OK {
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7755:18:
								zErr = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+10891, libc.VaList(bp+40, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname_length, (**(**TVec0MetadataColumnDefinition)(__ccgo_up(pVtab + 1600 + uintptr(metadata_idx)*24))).Fname, rowid3))
								// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7760:7:
								if zErr != 0 {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7761:9:
									libsqlite3.Xsqlite3_result_error(tls, context, zErr, -int32(1))
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7762:9:
									libsqlite3.Xsqlite3_free(tls, zErr)
								} else {
									// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7764:9:
									libsqlite3.Xsqlite3_result_error_nomem(tls, context)
								}
							}
						}
					}
				}
			}
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7769:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7772:12:
func _vec0Column(tls *libc.TLS, cur uintptr, context uintptr, i int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7773:30:
	var pCur, pVtab uintptr
	_, _ = pCur, pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7774:15:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7775:13:
	pVtab = (*Tsqlite3_vtab_cursor)(unsafe.Pointer(cur)).FpVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7776:3:
	switch (*Tvec0_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7777:3:
	case int32(_VEC0_QUERY_PLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7778:5:
		return _vec0Column_fullscan(tls, pVtab, pCur, context, i)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7780:3:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7781:5:
		return _vec0Column_knn(tls, pVtab, pCur, context, i)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7783:3:
		fallthrough
	case int32(_VEC0_QUERY_PLAN_POINT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7784:5:
		return _vec0Column_point(tls, pVtab, pCur, context, i)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7787:3:
	return m_SQLITE_OK
}

// C documentation
//
//	/**
//	 * @brief Handles the "insert rowid" step of a row insert operation of a vec0
//	 * table.
//	 *
//	 * This function will insert a new row into the _rowids vec0 shadow table.
//	 *
//	 * @param p: virtual table
//	 * @param idValue: Value containing the inserted rowid/id value.
//	 * @param rowid: Output rowid, will point to the "real" i64 rowid
//	 * value that was inserted
//	 * @return int SQLITE_OK on success, error code on failure
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7802:5:
func Xvec0Update_InsertRowidStep(tls *libc.TLS, p uintptr, idValue uintptr, rowid uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7803:44:
	var rc int32
	var suppliedRowid Ti64
	_, _ = rc, suppliedRowid
	// Option 3: vtab has a user-defined TEXT primary key, so ensure a text value
	// is provided.
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7816:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FpkIsText != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7817:5:
		if libsqlite3.Xsqlite3_value_type(tls, idValue) != int32(m_SQLITE_TEXT) {
			// IMP: V04200_21039
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7819:7:
			Xvtab_set_error(tls, p, __ccgo_ts+11046, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7823:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7826:5:
		return Xvec0_rowids_insert_id(tls, p, idValue, rowid)
	}
	// Option 1: User supplied a i64 rowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7830:3:
	if libsqlite3.Xsqlite3_value_type(tls, idValue) == int32(m_SQLITE_INTEGER) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7831:9:
		suppliedRowid = libsqlite3.Xsqlite3_value_int64(tls, idValue)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7832:5:
		rc = Xvec0_rowids_insert_rowid(tls, p, suppliedRowid)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7833:5:
		if rc == m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7834:7:
			**(**Ti64)(__ccgo_up(rowid)) = suppliedRowid
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7836:5:
		return rc
	}
	// Option 2: User did not suppled a rowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7841:3:
	if libsqlite3.Xsqlite3_value_type(tls, idValue) != int32(m_SQLITE_NULL) {
		// IMP: V30855_14925
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7843:5:
		Xvtab_set_error(tls, p, __ccgo_ts+11153, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7846:5:
		return int32(m_SQLITE_ERROR)
	}
	// NULL to get next auto-incremented value
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7849:3:
	return Xvec0_rowids_insert_id(tls, p, libc.UintptrFromInt32(0), rowid)
}

// C documentation
//
//	/**
//	 * @brief Determines the "next available" chunk position for a newly inserted
//	 * vec0 row.
//	 *
//	 * This operation may insert a new "blank" chunk the _chunks table, if there is
//	 * no more space in previous chunks.
//	 *
//	 * @param p: virtual table
//	 * @param partitionKeyValues: array of partition key column values, to constrain
//	 * against any partition key columns.
//	 * @param chunk_rowid: Output rowid of the chunk in the _chunks virtual table
//	 * that has the avialabiity.
//	 * @param chunk_offset: Output the index of the available space insert the
//	 * chunk, based on the index of the first available validity bit.
//	 * @param pBlobValidity: Output blob of the validity column of the available
//	 * chunk. Will be opened with read/write permissions.
//	 * @param pValidity: Output buffer of the original chunk's validity column.
//	 *    Needs to be cleaned up with sqlite3_free().
//	 * @return int SQLITE_OK on success, error code on failure
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7872:5:
func Xvec0Update_InsertNextAvailableStep(tls *libc.TLS, p uintptr, partitionKeyValues uintptr, chunk_rowid uintptr, chunk_offset uintptr, blobChunksValidity uintptr, bufferChunksValidity uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7877:49:
	var i, j, rc int32
	var validitySize Ti64
	_, _, _, _ = i, j, rc, validitySize
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7881:3:
	**(**Ti64)(__ccgo_up(chunk_offset)) = int64(-int32(1))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7883:3:
	rc = Xvec0_get_latest_chunk_rowid(tls, p, chunk_rowid, partitionKeyValues)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7884:3:
	if rc == int32(m_SQLITE_EMPTY) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7885:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7887:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7888:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7891:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+11207, **(**Ti64)(__ccgo_up(chunk_rowid)), int32(1), blobChunksValidity)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7893:3:
	if rc != m_SQLITE_OK {
		// IMP: V22053_06123
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7895:5:
		Xvtab_set_error(tls, p, __ccgo_ts+11216, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, **(**Ti64)(__ccgo_up(chunk_rowid))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7899:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7902:3:
	validitySize = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(blobChunksValidity))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7903:3:
	if validitySize != int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m_CHAR_BIT)) {
		// IMP: V29362_13432
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7905:5:
		Xvtab_set_error(tls, p, __ccgo_ts+11286, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, **(**Ti64)(__ccgo_up(chunk_rowid)), int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/libc.Int32FromInt32(m_CHAR_BIT)), validitySize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7911:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7912:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7915:3:
	**(**uintptr)(__ccgo_up(bufferChunksValidity)) = libsqlite3.Xsqlite3_malloc(tls, int32(validitySize))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7916:3:
	if !(**(**uintptr)(__ccgo_up(bufferChunksValidity)) != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7917:5:
		Xvtab_set_error(tls, p, __ccgo_ts+11389, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7919:5:
		rc = int32(m_SQLITE_NOMEM)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7920:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7923:3:
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(blobChunksValidity)), **(**uintptr)(__ccgo_up(bufferChunksValidity)), int32(validitySize), 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7926:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7927:5:
		Xvtab_set_error(tls, p, __ccgo_ts+11462, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, **(**Ti64)(__ccgo_up(chunk_rowid))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7931:5:
		goto cleanup
	}
	// find the next available offset, ie first `0` in the bitmap.
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7935:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7935:12:
	i = 0
	for {
		if !(int64(i) < validitySize) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7936:5:
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(**(**uintptr)(__ccgo_up(bufferChunksValidity)) + uintptr(i)))) == int32(0b11111111) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7937:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7938:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7938:14:
		j = 0
		for {
			if !(j < int32(m_CHAR_BIT)) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7939:7:
			if libc.Int32FromUint8(**(**uint8)(__ccgo_up(**(**uintptr)(__ccgo_up(bufferChunksValidity)) + uintptr(i))))>>j&int32(1) == 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7940:9:
				**(**Ti64)(__ccgo_up(chunk_offset)) = int64(i*int32(m_CHAR_BIT) + j)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7941:9:
				goto done
			}
			goto _2
		_2:
			;
			j = j + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7946:1:
	goto done
done:
	;
	// latest chunk was full, so need to create a new one
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7948:3:
	if **(**Ti64)(__ccgo_up(chunk_offset)) == int64(-int32(1)) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7949:5:
		rc = Xvec0_new_chunk(tls, p, partitionKeyValues, chunk_rowid)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7950:5:
		if rc != m_SQLITE_OK {
			// IMP: V08441_25279
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7952:7:
			Xvtab_set_error(tls, p, __ccgo_ts+11535, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7954:7:
			rc = int32(m_SQLITE_ERROR) // otherwise raises a DatabaseError and not operational
			// error?
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7956:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7958:5:
		**(**Ti64)(__ccgo_up(chunk_offset)) = 0
		// blobChunksValidity and pValidity are stale, pointing to the previous
		// (full) chunk. to re-assign them
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7962:5:
		rc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(blobChunksValidity)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7963:5:
		libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bufferChunksValidity)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7964:5:
		**(**uintptr)(__ccgo_up(blobChunksValidity)) = libc.UintptrFromInt32(0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7965:5:
		**(**uintptr)(__ccgo_up(bufferChunksValidity)) = libc.UintptrFromInt32(0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7966:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7967:7:
			Xvtab_set_error(tls, p, __ccgo_ts+11598, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7970:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7971:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7974:5:
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+11207, **(**Ti64)(__ccgo_up(chunk_rowid)), int32(1), blobChunksValidity)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7976:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7977:7:
			Xvtab_set_error(tls, p, __ccgo_ts+11702, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, **(**Ti64)(__ccgo_up(chunk_rowid))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7982:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7984:5:
		validitySize = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(blobChunksValidity))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7985:5:
		if validitySize != int64((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m_CHAR_BIT)) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7986:7:
			Xvtab_set_error(tls, p, __ccgo_ts+11793, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, **(**Ti64)(__ccgo_up(chunk_rowid)), (*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size/int32(m_CHAR_BIT), validitySize))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7992:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7994:5:
		**(**uintptr)(__ccgo_up(bufferChunksValidity)) = libsqlite3.Xsqlite3_malloc(tls, int32(validitySize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7995:5:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(blobChunksValidity)), **(**uintptr)(__ccgo_up(bufferChunksValidity)), int32(validitySize), 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7997:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:7998:7:
			Xvtab_set_error(tls, p, __ccgo_ts+11908, libc.VaList(bp+8, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, **(**Ti64)(__ccgo_up(chunk_rowid))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8003:7:
			goto cleanup
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8007:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8009:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8010:3:
	return rc
	return r
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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8026:1:
func _vec0_write_vector_to_vector_blob(tls *libc.TLS, blobVectors uintptr, chunk_offset Ti64, bVector uintptr, dimensions Tsize_t, element_type _VectorElementType) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8028:71:
	var n, offset int32
	_, _ = n, offset
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8032:3:
	switch element_type {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8033:3:
	case int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8034:5:
		n = libc.Int32FromUint64(dimensions * uint64(4))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8035:5:
		offset = libc.Int32FromUint64(libc.Uint64FromInt64(chunk_offset) * dimensions * uint64(4))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8036:5:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8037:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_INT8):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8038:5:
		n = libc.Int32FromUint64(dimensions * uint64(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8039:5:
		offset = libc.Int32FromUint64(libc.Uint64FromInt64(chunk_offset) * dimensions * uint64(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8040:5:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8041:3:
		fallthrough
	case int32(_SQLITE_VEC_ELEMENT_TYPE_BIT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8042:5:
		n = libc.Int32FromUint64(dimensions / uint64(m_CHAR_BIT))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8043:5:
		offset = libc.Int32FromUint64(libc.Uint64FromInt64(chunk_offset) * dimensions / uint64(m_CHAR_BIT))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8044:5:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8047:3:
	return libsqlite3.Xsqlite3_blob_write(tls, blobVectors, bVector, n, offset)
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
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8063:5:
func Xvec0Update_InsertWriteFinalStep(tls *libc.TLS, p uintptr, chunk_rowid Ti64, chunk_offset Ti64, _rowid Ti64, vectorDatas uintptr, blobChunksValidity uintptr, bufferChunksValidity uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*Ti64)(unsafe.Pointer(bp)) = _rowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8067:80:
	var actual, actual1, expected, expected1 Ti64
	var brc, i, rc int32
	var _ /* blobChunksRowids at bp+8 */ uintptr
	var _ /* blobVectors at bp+24 */ uintptr
	var _ /* bx at bp+16 */ uint8
	_, _, _, _, _, _, _ = actual, actual1, brc, expected, expected1, i, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8069:16:
	**(**uintptr)(__ccgo_up(bp + 8)) = libc.UintptrFromInt32(0)
	// mark the validity bit for this row in the chunk's validity bitmap
	// Get the byte offset of the bitmap
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8073:17:
	**(**uint8)(__ccgo_up(bp + 16)) = **(**uint8)(__ccgo_up(bufferChunksValidity + uintptr(chunk_offset/int64(m_CHAR_BIT))))
	// set the bit at the chunk_offset position inside that byte
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8075:3:
	**(**uint8)(__ccgo_up(bp + 16)) = libc.Uint8FromInt32(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 16))) | int32(1)<<(chunk_offset%int64(m_CHAR_BIT)))
	// write that 1 byte
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8077:3:
	rc = libsqlite3.Xsqlite3_blob_write(tls, blobChunksValidity, bp+16, int32(1), int32(chunk_offset/int64(m_CHAR_BIT)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8078:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8079:5:
		Xvtab_set_error(tls, p, __ccgo_ts+11995, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8080:5:
		return rc
	}
	// Go insert the vector data into the vector chunk shadow tables
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8084:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8084:12:
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8086:5:
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), __ccgo_ts+3712, chunk_rowid, int32(1), bp+24)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8088:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8089:7:
			Xvtab_set_error(tls, p, __ccgo_ts+12051, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_rowid))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8091:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8094:9:
		expected = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8096:9:
		actual = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp + 24))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8098:5:
		if actual != expected {
			// IMP: V16386_00456
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8100:7:
			Xvtab_set_error(tls, p, __ccgo_ts+12091, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_rowid, expected, actual))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8106:7:
			rc = int32(m_SQLITE_ERROR)
			// already error, can ignore result code
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8108:7:
			libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 24)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8109:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8110:6:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8112:5:
		rc = _vec0_write_vector_to_vector_blob(tls, **(**uintptr)(__ccgo_up(bp + 24)), chunk_offset, **(**uintptr)(__ccgo_up(vectorDatas + uintptr(i)*8)), (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fdimensions, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Felement_type)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8115:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8116:7:
			Xvtab_set_error(tls, p, __ccgo_ts+12186, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_rowid))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8120:7:
			rc = int32(m_SQLITE_ERROR)
			// already error, can ignore result code
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8122:7:
			libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 24)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8123:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8125:5:
		rc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 24)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8126:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8127:7:
			Xvtab_set_error(tls, p, __ccgo_ts+12255, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_rowid))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8131:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8132:7:
			goto cleanup
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// write the new rowid to the rowids column of the _chunks table
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8137:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+9690, chunk_rowid, int32(1), bp+8)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8139:3:
	if rc != m_SQLITE_OK {
		// IMP: V09221_26060
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8141:5:
		Xvtab_set_error(tls, p, __ccgo_ts+12324, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_rowid))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8144:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8146:7:
	expected1 = libc.Int64FromUint64(libc.Uint64FromInt32((*Tvec0_vtab)(unsafe.Pointer(p)).Fchunk_size) * uint64(8))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8147:7:
	actual1 = int64(libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp + 8))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8148:3:
	if expected1 != actual1 {
		// IMP: V12779_29618
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8150:5:
		Xvtab_set_error(tls, p, __ccgo_ts+12392, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_rowid, expected1, actual1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8155:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8156:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8158:3:
	rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp + 8)), bp, int32(8), libc.Int32FromUint64(libc.Uint64FromInt64(chunk_offset)*uint64(8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8160:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8161:5:
		Xvtab_set_error(tls, p, __ccgo_ts+12487, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_rowid))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8164:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8165:5:
		goto cleanup
	}
	// Now with all the vectors inserted, go back and update the _rowids table
	// with the new chunk_rowid/chunk_offset values
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8170:3:
	rc = Xvec0_rowids_update_position(tls, p, **(**Ti64)(__ccgo_up(bp)), chunk_rowid, chunk_offset)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8172:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8173:3:
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8174:3:
	if rc == m_SQLITE_OK && brc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8175:5:
		Xvtab_set_error(tls, p, __ccgo_ts+12556, libc.VaList(bp+40, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_rowid))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8178:5:
		return brc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8180:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8183:5:
func Xvec0_write_metadata_value(tls *libc.TLS, p uintptr, metadata_column_idx int32, rowid Ti64, chunk_id Ti64, chunk_offset Ti64, v uintptr, isupdate int32) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8183:146:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8185:39:
	metadata_column = p + 1600 + uintptr(metadata_column_idx)*24
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8186:29:
	kind = (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fkind
	// verify input value matches column type
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8189:3:
	switch kind {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8190:5:
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8191:7:
		if libsqlite3.Xsqlite3_value_type(tls, v) != int32(m_SQLITE_INTEGER) || libsqlite3.Xsqlite3_value_int(tls, v) != 0 && libsqlite3.Xsqlite3_value_int(tls, v) != int32(1) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8192:9:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8193:9:
			Xvtab_set_error(tls, p, __ccgo_ts+12625, libc.VaList(bp+80, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname_length, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8194:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8196:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8198:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8199:7:
		if libsqlite3.Xsqlite3_value_type(tls, v) != int32(m_SQLITE_INTEGER) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8200:9:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8201:9:
			Xvtab_set_error(tls, p, __ccgo_ts+12674, libc.VaList(bp+80, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname_length, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname, Xtype_name(tls, libsqlite3.Xsqlite3_value_type(tls, v))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8202:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8204:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8206:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8207:7:
		if libsqlite3.Xsqlite3_value_type(tls, v) != int32(m_SQLITE_FLOAT) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8208:9:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8209:9:
			Xvtab_set_error(tls, p, __ccgo_ts+12737, libc.VaList(bp+80, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname_length, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname, Xtype_name(tls, libsqlite3.Xsqlite3_value_type(tls, v))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8210:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8212:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8214:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8215:7:
		if libsqlite3.Xsqlite3_value_type(tls, v) != int32(m_SQLITE_TEXT) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8216:9:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8217:9:
			Xvtab_set_error(tls, p, __ccgo_ts+12796, libc.VaList(bp+80, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname_length, (*TVec0MetadataColumnDefinition)(unsafe.Pointer(metadata_column)).Fname, Xtype_name(tls, libsqlite3.Xsqlite3_value_type(tls, v))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8218:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8220:7:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8224:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8225:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 480 + uintptr(metadata_column_idx)*8)), __ccgo_ts+4053, chunk_id, int32(1), bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8226:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8227:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8230:3:
	switch kind {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8231:5:
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8233:11:
		value = libsqlite3.Xsqlite3_value_int(tls, v)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8234:7:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), int32(chunk_offset/libc.Int64FromInt32(m_CHAR_BIT)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8235:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8236:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8239:7:
		if value != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8240:9:
			**(**Tu8)(__ccgo_up(bp + 8)) = libc.Uint8FromInt32(int32(**(**Tu8)(__ccgo_up(bp + 8))) | libc.Int32FromInt32(1)<<(chunk_offset%libc.Int64FromInt32(m_CHAR_BIT)))
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8242:9:
			**(**Tu8)(__ccgo_up(bp + 8)) = libc.Uint8FromInt32(int32(**(**Tu8)(__ccgo_up(bp + 8))) & ^(libc.Int32FromInt32(1) << (chunk_offset % libc.Int64FromInt32(m_CHAR_BIT))))
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8245:7:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), int32(chunk_offset/int64(m_CHAR_BIT)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8246:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8248:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8249:11:
		**(**Ti64)(__ccgo_up(bp + 16)) = libsqlite3.Xsqlite3_value_int64(tls, v)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8250:7:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+16, int32(8), libc.Int32FromUint64(libc.Uint64FromInt64(chunk_offset)*uint64(8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8251:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8253:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8254:14:
		**(**float64)(__ccgo_up(bp + 24)) = libsqlite3.Xsqlite3_value_double(tls, v)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8255:7:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+24, int32(8), libc.Int32FromUint64(libc.Uint64FromInt64(chunk_offset)*uint64(8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8256:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8258:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8260:7:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+32, int32(4), int32(chunk_offset*int64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8261:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8262:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8265:18:
		s = libsqlite3.Xsqlite3_value_text(tls, v)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8266:11:
		**(**int32)(__ccgo_up(bp + 36)) = libsqlite3.Xsqlite3_value_bytes(tls, v)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8268:7:
		libc.Xmemset(tls, bp+40, 0, uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8269:7:
		libc.Xmemcpy(tls, bp+40, bp+36, uint64(4))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8270:7:
		if **(**int32)(__ccgo_up(bp + 36)) <= libc.Int32FromInt32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)-libc.Int32FromInt32(4) {
			v1 = **(**int32)(__ccgo_up(bp + 36))
		} else {
			v1 = libc.Int32FromInt32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH) - libc.Int32FromInt32(4)
		}
		libc.Xmemcpy(tls, bp+40+uintptr(4), s, libc.Uint64FromInt32(v1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8272:7:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+40, int32(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH), int32(chunk_offset*int64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8273:7:
		if **(**int32)(__ccgo_up(bp + 36)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8276:9:
			if isupdate != 0 && **(**int32)(__ccgo_up(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8277:11:
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12853, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8279:11:
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12918, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8281:9:
			if !(zSql != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8282:11:
				rc = int32(m_SQLITE_NOMEM)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8283:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8286:9:
			rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+56, libc.UintptrFromInt32(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8287:9:
			if rc != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8288:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8290:9:
			libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 56)), int32(1), rowid)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8291:9:
			libsqlite3.Xsqlite3_bind_text(tls, **(**uintptr)(__ccgo_up(bp + 56)), int32(2), s, **(**int32)(__ccgo_up(bp + 36)), libc.UintptrFromInt32(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8292:9:
			rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 56)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8293:9:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 56)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8295:9:
			if rc != int32(m_SQLITE_DONE) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8296:11:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8297:11:
				goto done
			}
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8300:12:
			if **(**int32)(__ccgo_up(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8301:20:
				zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12987, libc.VaList(bp+80, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_column_idx))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8302:9:
				if !(zSql1 != 0) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8303:11:
					rc = int32(m_SQLITE_NOMEM)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8304:11:
					goto done
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8307:9:
				rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql1, -int32(1), bp+64, libc.UintptrFromInt32(0))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8308:9:
				if rc != m_SQLITE_OK {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8309:11:
					goto done
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8311:9:
				libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 64)), int32(1), rowid)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8312:9:
				rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 64)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8313:9:
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 64)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8315:9:
				if rc != int32(m_SQLITE_DONE) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8316:11:
					rc = int32(m_SQLITE_ERROR)
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8317:11:
					goto done
				}
			}
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8320:7:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8324:3:
	if rc != m_SQLITE_OK {
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8327:3:
	rc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8328:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8329:5:
		goto done
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8332:3:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8333:5:
	return rc
	return r
}

// C documentation
//
//	/**
//	 * @brief Handles INSERT INTO operations on a vec0 table.
//	 *
//	 * @return int SQLITE_OK on success, otherwise error code on failure
//	 */
//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8342:5:
func Xvec0Update_Insert(tls *libc.TLS, pVTab uintptr, argc int32, argv uintptr, pRowid uintptr) (r int32) {
	bp := tls.Alloc(400)
	defer tls.Free(400)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8343:45:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8344:3:
	_ = argc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8345:13:
	p = pVTab
	// a write-able blob of the validity column for the given chunk. Used to mark
	// validity bit
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8367:16:
	**(**uintptr)(__ccgo_up(bp + 312)) = libc.UintptrFromInt32(0)
	// buffer for the valididty column for the given chunk. Maybe not needed here?
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8369:23:
	**(**uintptr)(__ccgo_up(bp + 320)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8370:7:
	numReadVectors = 0
	// Read all provided partition key values into partitionKeyValues
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8373:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8373:12:
	i = 0
	for {
		if !(i < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8374:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8375:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8377:9:
		partition_key_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8378:5:
		(**(**[4]uintptr)(__ccgo_up(bp + 264)))[partition_key_idx] = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i)*8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8380:9:
		new_value_type = libsqlite3.Xsqlite3_value_type(tls, (**(**[4]uintptr)(__ccgo_up(bp + 264)))[partition_key_idx])
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8381:5:
		if new_value_type != int32(m_SQLITE_NULL) && new_value_type != (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(partition_key_idx)*24))).Ftype1 {
			// IMP: V11454_28292
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8383:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+13042, libc.VaList(bp+368, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(partition_key_idx)*24))).Fname_length, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(partition_key_idx)*24))).Fname, Xtype_name(tls, (**(**TVec0PartitionColumnDefinition)(__ccgo_up(p + 1120 + uintptr(partition_key_idx)*24))).Ftype1), Xtype_name(tls, new_value_type)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8391:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8392:7:
			goto cleanup
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// read all the inserted vectors  into vectorDatas, validate their lengths.
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8397:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8397:12:
	i1 = 0
	for {
		if !(i1 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8398:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i1)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8399:7:
			goto _2
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8401:9:
		vector_column_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i1))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8402:19:
		valueVector = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i1)*8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8407:5:
		rc = Xvector_from_value(tls, valueVector, bp+8+uintptr(vector_column_idx)*8, bp+328, bp+344, bp+136+uintptr(vector_column_idx)*8, bp+336)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8409:5:
		if rc != m_SQLITE_OK {
			// IMP: V06519_23358
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8411:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+13134, libc.VaList(bp+368, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fname, **(**uintptr)(__ccgo_up(bp + 336))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8414:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8415:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8418:5:
		numReadVectors = numReadVectors + 1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8419:5:
		if **(**_VectorElementType)(__ccgo_up(bp + 344)) != (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Felement_type {
			// IMP: V08221_25059
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8421:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+13187, libc.VaList(bp+368, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i1)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i1)*32))).Fname, Xvector_subtype_name(tls, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i1)*32))).Felement_type), Xvector_subtype_name(tls, **(**_VectorElementType)(__ccgo_up(bp + 344)))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8428:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8429:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8432:5:
		if **(**Tsize_t)(__ccgo_up(bp + 328)) != (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fdimensions {
			// IMP: V01145_17984
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8434:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+13285, libc.VaList(bp+368, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fname, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(vector_column_idx)*32))).Fdimensions, **(**Tsize_t)(__ccgo_up(bp + 328))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8440:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8441:7:
			goto cleanup
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// Cannot insert a value in the hidden "distance" column
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8446:3:
	if libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv + uintptr(int32(2)+Xvec0_column_distance_idx(tls, p))*8))) != int32(m_SQLITE_NULL) {
		// IMP: V24228_08298
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8449:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+13387, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8451:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8452:5:
		goto cleanup
	}
	// Cannot insert a value in the hidden "k" column
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8455:3:
	if libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv + uintptr(int32(2)+Xvec0_column_k_idx(tls, p))*8))) != int32(m_SQLITE_NULL) {
		// IMP: V11875_28713
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8457:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+13442, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8458:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8459:5:
		goto cleanup
	}
	// Step #1: Insert/get a rowid for this row, from the _rowids table.
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8463:3:
	rc = Xvec0Update_InsertRowidStep(tls, p, **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_ID))*8)), bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8464:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8465:5:
		goto cleanup
	}
	// Step #2: Find the next "available" position in the _chunks table for this
	// row.
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8470:3:
	rc = Xvec0Update_InsertNextAvailableStep(tls, p, bp+264, bp+296, bp+304, bp+312, bp+320)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8474:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8475:5:
		goto cleanup
	}
	// Step #3: With the next available chunk position, write out all the vectors
	//          to their specified location.
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8480:3:
	rc = Xvec0Update_InsertWriteFinalStep(tls, p, **(**Ti64)(__ccgo_up(bp + 296)), **(**Ti64)(__ccgo_up(bp + 304)), **(**Ti64)(__ccgo_up(bp)), bp+8, **(**uintptr)(__ccgo_up(bp + 312)), **(**uintptr)(__ccgo_up(bp + 320)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8483:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8484:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8487:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns > 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8489:17:
		s = libsqlite3.Xsqlite3_str_new(tls, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8490:5:
		libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+13490, libc.VaList(bp+368, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8491:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8491:13:
		i2 = 0
		for {
			if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8492:7:
			libsqlite3.Xsqlite3_str_appendf(tls, s, __ccgo_ts+8098, libc.VaList(bp+368, i2))
			goto _3
		_3:
			;
			i2 = i2 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8494:5:
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+13529)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8495:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8495:13:
		i3 = 0
		for {
			if !(i3 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8496:7:
			libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5252)
			goto _4
		_4:
			;
			i3 = i3 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8498:5:
		libsqlite3.Xsqlite3_str_appendall(tls, s, __ccgo_ts+5256)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8499:10:
		zSql = libsqlite3.Xsqlite3_str_finish(tls, s)
		// TODO double check error handling ehre
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8501:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8502:7:
			rc = int32(m_SQLITE_NOMEM)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8503:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8505:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+352, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8506:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8507:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8509:5:
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 352)), int32(1), **(**Ti64)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8511:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8511:14:
		i4 = 0
		for {
			if !(i4 < Xvec0_num_defined_user_columns(tls, p)) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8512:7:
			if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i4)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8513:9:
				goto _5
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8515:11:
			auxiliary_key_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i4))))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8516:21:
			v = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i4)*8))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8517:11:
			v_type = libsqlite3.Xsqlite3_value_type(tls, v)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8518:7:
			if v_type != int32(m_SQLITE_NULL) && v_type != (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(auxiliary_key_idx)*24))).Ftype1 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8519:9:
				libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 352)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8520:9:
				rc = int32(m_SQLITE_CONSTRAINT)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8521:9:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+13542, libc.VaList(bp+368, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(auxiliary_key_idx)*24))).Fname_length, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(auxiliary_key_idx)*24))).Fname, Xtype_name(tls, (**(**TVec0AuxiliaryColumnDefinition)(__ccgo_up(p + 1216 + uintptr(auxiliary_key_idx)*24))).Ftype1), Xtype_name(tls, v_type)))
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8529:9:
				goto cleanup
			}
			// first 1 is for 1-based indexing on sqlite3_bind_*, second 1 is to account for initial rowid parameter
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8532:7:
			libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(bp + 352)), libc.Int32FromInt32(1)+libc.Int32FromInt32(1)+auxiliary_key_idx, v)
			goto _5
		_5:
			;
			i4 = i4 + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8535:5:
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 352)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8536:5:
		if rc != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8537:7:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 352)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8538:7:
			rc = int32(m_SQLITE_ERROR)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8539:7:
			goto cleanup
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8541:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 352)))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8545:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8545:11:
	i5 = 0
	for {
		if !(i5 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8546:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i5)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8547:7:
			goto _6
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8549:9:
		metadata_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i5))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8550:19:
		v1 = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i5)*8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8551:5:
		rc = Xvec0_write_metadata_value(tls, p, metadata_idx, **(**Ti64)(__ccgo_up(bp)), **(**Ti64)(__ccgo_up(bp + 296)), **(**Ti64)(__ccgo_up(bp + 304)), v1, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8552:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8553:7:
			goto cleanup
		}
		goto _6
	_6:
		;
		i5 = i5 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8557:3:
	**(**Tsqlite_int64)(__ccgo_up(pRowid)) = **(**Ti64)(__ccgo_up(bp))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8558:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8560:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8561:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8561:12:
	i6 = 0
	for {
		if !(i6 < numReadVectors) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8562:5:
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(**(**[16]Tvector_cleanup)(__ccgo_up(bp + 136)))[i6]})))(tls, (**(**[16]uintptr)(__ccgo_up(bp + 8)))[i6])
		goto _7
	_7:
		;
		i6 = i6 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8564:3:
	libsqlite3.Xsqlite3_free(tls, **(**uintptr)(__ccgo_up(bp + 320)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8565:7:
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp + 312)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8566:3:
	if rc == m_SQLITE_OK && brc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8567:5:
		Xvtab_set_error(tls, p, __ccgo_ts+13634, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8570:5:
		return brc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8572:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8575:5:
func Xvec0Update_Delete_ClearValidity(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8576:55:
	var brc, rc, validityOffset int32
	var mask uint8
	var _ /* blobChunksValidity at bp+0 */ uintptr
	var _ /* bx at bp+8 */ uint8
	var _ /* result at bp+9 */ int8
	_, _, _, _ = brc, mask, rc, validityOffset
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8578:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8580:7:
	validityOffset = libc.Int32FromUint64(chunk_offset / uint64(m_CHAR_BIT))
	// 2. ensure chunks.validity bit is 1, then set to 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8583:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+11207, chunk_id, int32(1), bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8585:3:
	if rc != m_SQLITE_OK {
		// IMP: V26002_10073
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8587:5:
		Xvtab_set_error(tls, p, __ccgo_ts+13737, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8589:5:
		return int32(m_SQLITE_ERROR)
	}
	// will skip the sqlite3_blob_bytes(blobChunksValidity) check for now,
	// the read below would catch it
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8594:3:
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), validityOffset)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8595:3:
	if rc != m_SQLITE_OK {
		// IMP: V21193_05263
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8597:5:
		Xvtab_set_error(tls, p, __ccgo_ts+13781, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8600:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8602:3:
	if !(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 8)))>>(chunk_offset%libc.Uint64FromInt32(m_CHAR_BIT)) != 0) {
		// IMP: V21193_05263
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8604:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8605:5:
		Xvtab_set_error(tls, p, __ccgo_ts+13831, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8609:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8611:17:
	mask = libc.Uint8FromInt32(^(libc.Int32FromInt32(1) << (chunk_offset % libc.Uint64FromInt32(m_CHAR_BIT))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8612:8:
	**(**int8)(__ccgo_up(bp + 9)) = int8(libc.Int32FromUint8(**(**uint8)(__ccgo_up(bp + 8))) & libc.Int32FromUint8(mask))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8613:3:
	rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+9, int32(1), validityOffset)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8615:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8616:5:
		Xvtab_set_error(tls, p, __ccgo_ts+13897, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8619:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8622:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8624:3:
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8625:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8626:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8627:3:
	if brc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8628:5:
		Xvtab_set_error(tls, p, __ccgo_ts+13951, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, validityOffset))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8633:5:
		return brc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8635:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8638:5:
func Xvec0Update_Delete_ClearRowid(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8639:53:
	var brc, rc int32
	var _ /* blobChunksRowids at bp+0 */ uintptr
	var _ /* zero at bp+8 */ Ti64
	_, _ = brc, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8641:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8642:7:
	**(**Ti64)(__ccgo_up(bp + 8)) = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8644:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+9690, chunk_id, int32(1), bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8646:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8647:5:
		Xvtab_set_error(tls, p, __ccgo_ts+14034, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8649:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8652:3:
	rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(8), libc.Int32FromUint64(chunk_offset*uint64(8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8654:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8655:5:
		Xvtab_set_error(tls, p, __ccgo_ts+14076, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, chunk_offset))
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8660:3:
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8661:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8662:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8663:3:
	if brc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8664:5:
		Xvtab_set_error(tls, p, __ccgo_ts+14130, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, chunk_id, chunk_offset))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8668:5:
		return brc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8670:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8673:5:
func Xvec0Update_Delete_ClearVectors(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8674:55:
	var brc, i, rc int32
	var n Tsize_t
	var zeroBuf uintptr
	var _ /* blobVectors at bp+0 */ uintptr
	_, _, _, _, _ = brc, i, n, rc, zeroBuf
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8676:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8676:12:
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8677:18:
		**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8678:12:
		n = Xvector_column_byte_size(tls, **(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8680:5:
		rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), __ccgo_ts+3712, chunk_id, int32(1), bp)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8683:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8684:7:
			Xvtab_set_error(tls, p, __ccgo_ts+14213, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id, i))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8687:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8690:10:
		zeroBuf = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(n))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8691:5:
		if !(zeroBuf != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8692:7:
			libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8693:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8695:5:
		libc.Xmemset(tls, zeroBuf, 0, n)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8697:5:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), zeroBuf, libc.Int32FromUint64(n), libc.Int32FromUint64(chunk_offset*n))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8698:5:
		libsqlite3.Xsqlite3_free(tls, zeroBuf)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8699:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8700:7:
			Xvtab_set_error(tls, p, __ccgo_ts+14265, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id, chunk_offset, i))
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8707:5:
		brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8708:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8709:7:
			return rc
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8710:5:
		if brc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8711:7:
			Xvtab_set_error(tls, p, __ccgo_ts+14329, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id, i))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8715:7:
			return brc
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8718:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8721:5:
func Xvec0Update_Delete_DeleteChunkIfEmpty(tls *libc.TLS, p uintptr, chunk_id Ti64, deleted uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8722:57:
	var allZero, brc, i, i1, i2, rc, validitySize int32
	var validityBuf, zSql uintptr
	var _ /* blobValidity at bp+0 */ uintptr
	var _ /* stmt at bp+8 */ uintptr
	_, _, _, _, _, _, _, _, _ = allZero, brc, i, i1, i2, rc, validityBuf, validitySize, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8724:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8725:3:
	**(**int32)(__ccgo_up(deleted)) = 0
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8727:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FshadowChunksName, __ccgo_ts+11207, chunk_id, 0, bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8729:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8730:5:
		Xvtab_set_error(tls, p, __ccgo_ts+14414, libc.VaList(bp+24, chunk_id))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8732:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8735:7:
	validitySize = libsqlite3.Xsqlite3_blob_bytes(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8736:17:
	validityBuf = libsqlite3.Xsqlite3_malloc(tls, validitySize)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8737:3:
	if !(validityBuf != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8738:5:
		libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8739:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8742:3:
	rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), validityBuf, validitySize, 0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8743:3:
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8744:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8745:5:
		libsqlite3.Xsqlite3_free(tls, validityBuf)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8746:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8748:3:
	if brc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8749:5:
		libsqlite3.Xsqlite3_free(tls, validityBuf)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8750:5:
		return brc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8753:7:
	allZero = int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8754:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8754:12:
	i = 0
	for {
		if !(i < validitySize) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8755:5:
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(validityBuf + uintptr(i)))) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8756:7:
			allZero = 0
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8757:7:
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8760:3:
	libsqlite3.Xsqlite3_free(tls, validityBuf)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8762:3:
	if !(allZero != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8763:5:
		return m_SQLITE_OK
	}
	// Delete from _chunks
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8771:3:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14458, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8774:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8775:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8776:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+8, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8777:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8778:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8779:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8780:3:
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 8)), int32(1), chunk_id)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8781:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8782:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 8)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8783:3:
	if rc != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8784:5:
		return int32(m_SQLITE_ERROR)
	}
	// Delete from each _vector_chunksNN
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8787:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8787:12:
	i1 = 0
	for {
		if !(i1 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumVectorColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8788:5:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14503, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, i1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8791:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8792:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8793:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+8, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8794:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8795:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8796:7:
			return rc
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8797:5:
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 8)), int32(1), chunk_id)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8798:5:
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8799:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8800:5:
		if rc != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8801:7:
			return int32(m_SQLITE_ERROR)
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// Delete from each _metadatachunksNN
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8805:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8805:12:
	i2 = 0
	for {
		if !(i2 < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumMetadataColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8806:5:
		zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14559, libc.VaList(bp+24, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, i2))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8809:5:
		if !(zSql != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8810:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8811:5:
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+8, libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8812:5:
		libsqlite3.Xsqlite3_free(tls, zSql)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8813:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8814:7:
			return rc
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8815:5:
		libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 8)), int32(1), chunk_id)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8816:5:
		rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8817:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8818:5:
		if rc != int32(m_SQLITE_DONE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8819:7:
			return int32(m_SQLITE_ERROR)
		}
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	// Invalidate cached stmtLatestChunk so it gets re-prepared on next insert
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8823:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8824:5:
		libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8825:5:
		(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8828:3:
	**(**int32)(__ccgo_up(deleted)) = int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8829:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8832:5:
func Xvec0Update_Delete_DeleteRowids(tls *libc.TLS, p uintptr, rowid Ti64) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8832:61:
	var rc int32
	var zSql uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _ = rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8834:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8836:8:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14616, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8839:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8840:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8843:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8844:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8845:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8846:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8848:3:
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8849:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8850:3:
	if rc != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8851:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8853:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8855:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8856:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8857:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8860:5:
func Xvec0Update_Delete_DeleteAux(tls *libc.TLS, p uintptr, rowid Ti64) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8860:58:
	var rc int32
	var zSql uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _ = rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8862:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8864:8:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14661, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8867:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8868:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8871:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8872:3:
	libsqlite3.Xsqlite3_free(tls, zSql)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8873:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8874:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8876:3:
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8877:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8878:3:
	if rc != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8879:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8881:3:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8883:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8884:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8885:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8888:5:
func Xvec0Update_Delete_ClearMetadata(tls *libc.TLS, p uintptr, metadata_idx int32, rowid Ti64, chunk_id Ti64, chunk_offset Tu64) (r int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8889:55:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8892:29:
	kind = (**(**TVec0MetadataColumnDefinition)(__ccgo_up(p + 1600 + uintptr(metadata_idx)*24))).Fkind
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8893:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 480 + uintptr(metadata_idx)*8)), __ccgo_ts+4053, chunk_id, int32(1), bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8894:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8895:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8898:3:
	switch kind {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8899:5:
	case int32(_VEC0_METADATA_COLUMN_KIND_BOOLEAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8901:7:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), libc.Int32FromUint64(chunk_offset/libc.Uint64FromInt32(m_CHAR_BIT)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8902:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8903:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8906:7:
		**(**Tu8)(__ccgo_up(bp + 8)) = libc.Uint8FromInt32(int32(**(**Tu8)(__ccgo_up(bp + 8))) & ^(libc.Int32FromInt32(1) << (chunk_offset % libc.Uint64FromInt32(m_CHAR_BIT))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8907:7:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+8, int32(1), libc.Int32FromUint64(chunk_offset/uint64(m_CHAR_BIT)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8908:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8910:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_INTEGER):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8911:11:
		**(**Ti64)(__ccgo_up(bp + 16)) = 0
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8912:7:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+16, int32(8), libc.Int32FromUint64(chunk_offset*uint64(8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8913:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8915:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_FLOAT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8916:14:
		**(**float64)(__ccgo_up(bp + 24)) = libc.Float64FromInt32(0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8917:7:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+24, int32(8), libc.Int32FromUint64(chunk_offset*uint64(8)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8918:7:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8920:5:
		fallthrough
	case int32(_VEC0_METADATA_COLUMN_KIND_TEXT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8922:7:
		rc = libsqlite3.Xsqlite3_blob_read(tls, **(**uintptr)(__ccgo_up(bp)), bp+32, int32(4), libc.Int32FromUint64(chunk_offset*uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8923:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8924:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8928:7:
		libc.Xmemset(tls, bp+36, 0, uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8929:7:
		rc = libsqlite3.Xsqlite3_blob_write(tls, **(**uintptr)(__ccgo_up(bp)), bp+36, int32(16), libc.Int32FromUint64(chunk_offset*uint64(m_VEC0_METADATA_TEXT_VIEW_BUFFER_LENGTH)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8930:7:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8931:9:
			goto done
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8934:7:
		if **(**int32)(__ccgo_up(bp + 32)) > int32(m_VEC0_METADATA_TEXT_VIEW_DATA_LENGTH) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8935:20:
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+12987, libc.VaList(bp+72, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, metadata_idx))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8936:9:
			if !(zSql != 0) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8937:11:
				rc = int32(m_SQLITE_NOMEM)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8938:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8941:9:
			rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp+56, libc.UintptrFromInt32(0))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8942:9:
			if rc != m_SQLITE_OK {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8943:11:
				goto done
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8945:9:
			libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp + 56)), int32(1), rowid)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8946:9:
			rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp + 56)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8947:9:
			libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp + 56)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8948:9:
			if rc != int32(m_SQLITE_DONE) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8949:11:
				rc = int32(m_SQLITE_ERROR)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8950:11:
				goto done
			}
			// Fix for https://github.com/asg017/sqlite-vec/issues/274
			// sqlite3_step returns SQLITE_DONE (101) on DML success, but the
			// `done:` epilogue treats anything other than SQLITE_OK as an error.
			// Without this, SQLITE_DONE propagates up to vec0Update_Delete,
			// which aborts the DELETE scan and silently drops remaining rows.
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8957:9:
			rc = m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8959:7:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8963:3:
	goto done
done:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8964:3:
	rc2 = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8965:3:
	if rc == m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8966:5:
		return rc2
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8968:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8971:5:
func Xvec0Update_Delete(tls *libc.TLS, pVTab uintptr, idValue uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8971:68:
	var i, rc int32
	var p uintptr
	var _ /* chunkDeleted at bp+24 */ int32
	var _ /* chunk_id at bp+8 */ Ti64
	var _ /* chunk_offset at bp+16 */ Ti64
	var _ /* rowid at bp+0 */ Ti64
	_, _, _ = i, p, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8972:13:
	p = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8978:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FpkIsText != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8979:5:
		rc = Xvec0_rowid_from_id(tls, p, idValue, bp)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8980:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8981:7:
			return rc
		}
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8984:5:
		**(**Ti64)(__ccgo_up(bp)) = libsqlite3.Xsqlite3_value_int64(tls, idValue)
	}
	// 1. Find chunk position for given rowid
	// 2. Ensure that validity bit for position is 1, then set to 0
	// 3. Zero out rowid in chunks.rowid
	// 4. Zero out vector data in all vector column chunks
	// 5. Delete value in _rowids table
	// 1. get chunk_id and chunk_offset from _rowids
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8994:3:
	rc = Xvec0_get_chunk_position(tls, p, **(**Ti64)(__ccgo_up(bp)), libc.UintptrFromInt32(0), bp+8, bp+16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8995:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:8996:5:
		return rc
	}
	// 2. clear validity bit
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9000:3:
	rc = Xvec0Update_Delete_ClearValidity(tls, p, **(**Ti64)(__ccgo_up(bp + 8)), libc.Uint64FromInt64(**(**Ti64)(__ccgo_up(bp + 16))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9001:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9002:5:
		return rc
	}
	// 3. zero out rowid in chunks.rowids
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9006:3:
	rc = Xvec0Update_Delete_ClearRowid(tls, p, **(**Ti64)(__ccgo_up(bp + 8)), libc.Uint64FromInt64(**(**Ti64)(__ccgo_up(bp + 16))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9007:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9008:5:
		return rc
	}
	// 4. zero out any data in vector chunks tables
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9012:3:
	rc = Xvec0Update_Delete_ClearVectors(tls, p, **(**Ti64)(__ccgo_up(bp + 8)), libc.Uint64FromInt64(**(**Ti64)(__ccgo_up(bp + 16))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9013:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9014:5:
		return rc
	}
	// 5. delete from _rowids table
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9018:3:
	rc = Xvec0Update_Delete_DeleteRowids(tls, p, **(**Ti64)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9019:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9020:5:
		return rc
	}
	// 6. delete any auxiliary rows
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9024:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FnumAuxiliaryColumns > 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9025:5:
		rc = Xvec0Update_Delete_DeleteAux(tls, p, **(**Ti64)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9026:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9027:7:
			return rc
		}
	}
	// 7. delete metadata
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9032:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9032:11:
	i = 0
	for {
		if !(i < (*Tvec0_vtab)(unsafe.Pointer(p)).FnumMetadataColumns) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9033:5:
		rc = Xvec0Update_Delete_ClearMetadata(tls, p, i, **(**Ti64)(__ccgo_up(bp)), **(**Ti64)(__ccgo_up(bp + 8)), libc.Uint64FromInt64(**(**Ti64)(__ccgo_up(bp + 16))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9034:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9035:7:
			return rc
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// 8. reclaim chunk if fully empty
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9042:5:
	rc = Xvec0Update_Delete_DeleteChunkIfEmpty(tls, p, **(**Ti64)(__ccgo_up(bp + 8)), bp+24)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9043:5:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9044:7:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9048:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9051:5:
func Xvec0Update_UpdateAuxColumn(tls *libc.TLS, p uintptr, auxiliary_column_idx int32, value uintptr, rowid Ti64) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9051:106:
	var rc int32
	var zSql uintptr
	var _ /* stmt at bp+0 */ uintptr
	_, _ = rc, zSql
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9054:14:
	zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+14709, libc.VaList(bp+16, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, (*Tvec0_vtab)(unsafe.Pointer(p)).FtableName, auxiliary_column_idx))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9055:3:
	if !(zSql != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9056:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9058:3:
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, zSql, -int32(1), bp, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9059:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9060:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9062:3:
	libsqlite3.Xsqlite3_bind_value(tls, **(**uintptr)(__ccgo_up(bp)), int32(1), value)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9063:3:
	libsqlite3.Xsqlite3_bind_int64(tls, **(**uintptr)(__ccgo_up(bp)), int32(2), rowid)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9064:3:
	rc = libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9065:3:
	if rc != int32(m_SQLITE_DONE) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9066:5:
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9067:5:
		return int32(m_SQLITE_ERROR)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9069:3:
	libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9070:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9073:5:
func Xvec0Update_UpdateVectorColumn(tls *libc.TLS, p uintptr, chunk_id Ti64, chunk_offset Ti64, i int32, valueVector uintptr) (r int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9074:70:
	var brc, rc int32
	var _ /* blobVectors at bp+0 */ uintptr
	var _ /* cleanup at bp+40 */ Tvector_cleanup
	var _ /* dimensions at bp+16 */ Tsize_t
	var _ /* elementType at bp+24 */ _VectorElementType
	var _ /* pzError at bp+8 */ uintptr
	var _ /* vector at bp+32 */ uintptr
	_, _ = brc, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9077:16:
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9083:18:
	**(**Tvector_cleanup)(__ccgo_up(bp + 40)) = __ccgo_fp(Xvector_cleanup_noop)
	// https://github.com/asg017/sqlite-vec/issues/53
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9085:3:
	rc = Xvector_from_value(tls, valueVector, bp+32, bp+16, bp+24, bp+40, bp+8)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9087:3:
	if rc != m_SQLITE_OK {
		// IMP: V15203_32042
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9089:5:
		Xvtab_set_error(tls, p, __ccgo_ts+14770, libc.VaList(bp+56, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fname, **(**uintptr)(__ccgo_up(bp + 8))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9092:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9093:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9095:3:
	if **(**_VectorElementType)(__ccgo_up(bp + 24)) != (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Felement_type {
		// IMP: V03643_20481
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9097:5:
		Xvtab_set_error(tls, p, __ccgo_ts+14822, libc.VaList(bp+56, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fname, Xvector_subtype_name(tls, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Felement_type), Xvector_subtype_name(tls, **(**_VectorElementType)(__ccgo_up(bp + 24)))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9104:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9105:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9107:3:
	if **(**Tsize_t)(__ccgo_up(bp + 16)) != (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fdimensions {
		// IMP: V25739_09810
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9109:5:
		Xvtab_set_error(tls, p, __ccgo_ts+14919, libc.VaList(bp+56, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fname_length, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fname, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fdimensions, **(**Tsize_t)(__ccgo_up(bp + 16))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9115:5:
		rc = int32(m_SQLITE_ERROR)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9116:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9119:3:
	rc = libsqlite3.Xsqlite3_blob_open(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).Fdb, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), __ccgo_ts+3712, chunk_id, int32(1), bp)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9121:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9122:5:
		Xvtab_set_error(tls, p, __ccgo_ts+15024, libc.VaList(bp+56, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9124:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9126:3:
	rc = _vec0_write_vector_to_vector_blob(tls, **(**uintptr)(__ccgo_up(bp)), chunk_offset, **(**uintptr)(__ccgo_up(bp + 32)), (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Fdimensions, (**(**TVectorColumnDefinition)(__ccgo_up(p + 608 + uintptr(i)*32))).Felement_type)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9129:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9130:5:
		Xvtab_set_error(tls, p, __ccgo_ts+15067, libc.VaList(bp+56, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9132:5:
		goto cleanup
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9135:1:
	goto cleanup
cleanup:
	;
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9136:3:
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(bp + 40)))(tls, **(**uintptr)(__ccgo_up(bp + 32)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9137:7:
	brc = libsqlite3.Xsqlite3_blob_close(tls, **(**uintptr)(__ccgo_up(bp)))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9138:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9139:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9141:3:
	if brc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9142:5:
		Xvtab_set_error(tls, p, __ccgo_ts+15114, libc.VaList(bp+56, (*Tvec0_vtab)(unsafe.Pointer(p)).FschemaName, **(**uintptr)(__ccgo_up(p + 352 + uintptr(i)*8)), chunk_id))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9146:5:
		return brc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9148:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9151:5:
func Xvec0Update_Update(tls *libc.TLS, pVTab uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9151:76:
	var a, b, p, value, value1, value2, valueVector uintptr
	var auxiliary_column_idx, i, i1, i2, i3, metadata_column_idx, rc, vector_idx int32
	var _ /* chunk_id at bp+0 */ Ti64
	var _ /* chunk_offset at bp+8 */ Ti64
	var _ /* rowid at bp+16 */ Ti64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a, auxiliary_column_idx, b, i, i1, i2, i3, metadata_column_idx, p, rc, value, value1, value2, valueVector, vector_idx
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9152:3:
	_ = argc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9153:13:
	p = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9159:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FpkIsText != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9160:16:
		a = libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9161:16:
		b = libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv + 1*8)))
		// IMP: V08886_25725
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9163:5:
		if libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv))) != libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv + 1*8))) || libc.Xstrncmp(tls, a, b, libc.Uint64FromInt32(libsqlite3.Xsqlite3_value_bytes(tls, **(**uintptr)(__ccgo_up(argv))))) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9165:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+15180, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9167:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9169:5:
		rc = Xvec0_rowid_from_id(tls, p, **(**uintptr)(__ccgo_up(argv)), bp+16)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9170:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9171:7:
			return rc
		}
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9174:5:
		**(**Ti64)(__ccgo_up(bp + 16)) = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv)))
	}
	// 1) get chunk_id and chunk_offset from _rowids
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9178:3:
	rc = Xvec0_get_chunk_position(tls, p, **(**Ti64)(__ccgo_up(bp + 16)), libc.UintptrFromInt32(0), bp, bp+8)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9179:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9180:5:
		return rc
	}
	// 2) update any partition key values
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9184:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9184:12:
	i = 0
	for {
		if !(i < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9185:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_PARTITION) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9186:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9188:19:
		value = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i)*8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9189:5:
		if libsqlite3.Xsqlite3_value_nochange(tls, value) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9190:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9192:5:
		Xvtab_set_error(tls, pVTab, __ccgo_ts+15232, 0)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9193:5:
		return int32(m_SQLITE_ERROR)
		goto _1
	_1:
		;
		i = i + 1
	}
	// 3) handle auxiliary column updates
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9197:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9197:12:
	i1 = 0
	for {
		if !(i1 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9198:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i1)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_AUXILIARY) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9199:7:
			goto _2
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9201:9:
		auxiliary_column_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i1))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9202:19:
		value1 = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i1)*8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9203:5:
		if libsqlite3.Xsqlite3_value_nochange(tls, value1) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9204:7:
			goto _2
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9206:5:
		rc = Xvec0Update_UpdateAuxColumn(tls, p, auxiliary_column_idx, value1, **(**Ti64)(__ccgo_up(bp + 16)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9207:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9208:7:
			return int32(m_SQLITE_ERROR)
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// 4) handle metadata column updates
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9213:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9213:12:
	i2 = 0
	for {
		if !(i2 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9214:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i2)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_METADATA) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9215:7:
			goto _3
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9217:9:
		metadata_column_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i2))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9218:19:
		value2 = **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC0_COLUMN_USERN_START)+i2)*8))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9219:5:
		if libsqlite3.Xsqlite3_value_nochange(tls, value2) != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9220:7:
			goto _3
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9222:5:
		rc = Xvec0_write_metadata_value(tls, p, metadata_column_idx, **(**Ti64)(__ccgo_up(bp + 16)), **(**Ti64)(__ccgo_up(bp)), **(**Ti64)(__ccgo_up(bp + 8)), value2, int32(1))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9223:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9224:7:
			return rc
		}
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	// 5) iterate over all new vectors, update the vectors
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9229:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9229:12:
	i3 = 0
	for {
		if !(i3 < Xvec0_num_defined_user_columns(tls, p)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9230:5:
		if **(**Tvec0_user_column_kind)(__ccgo_up(p + 88 + uintptr(i3)*4)) != int32(_SQLITE_VEC0_USER_COLUMN_KIND_VECTOR) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9231:7:
			goto _4
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9233:9:
		vector_idx = libc.Int32FromUint8(**(**Tuint8_t)(__ccgo_up(p + 296 + uintptr(i3))))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9234:19:
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
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9245:5:
		if libsqlite3.Xsqlite3_value_type(tls, valueVector) == int32(m_SQLITE_NULL) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9246:7:
			goto _4
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9249:5:
		rc = Xvec0Update_UpdateVectorColumn(tls, p, **(**Ti64)(__ccgo_up(bp)), **(**Ti64)(__ccgo_up(bp + 8)), vector_idx, valueVector)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9251:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9252:7:
			return int32(m_SQLITE_ERROR)
		}
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9256:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9259:12:
func _vec0Update(tls *libc.TLS, pVTab uintptr, argc int32, argv uintptr, pRowid uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9260:45:
	// DELETE operation
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9262:3:
	if argc == int32(1) && libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv))) != int32(m_SQLITE_NULL) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9263:5:
		return Xvec0Update_Delete(tls, pVTab, **(**uintptr)(__ccgo_up(argv)))
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9266:8:
		if argc > int32(1) && libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv))) == int32(m_SQLITE_NULL) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9267:5:
			return Xvec0Update_Insert(tls, pVTab, argc, argv, pRowid)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9270:8:
			if argc > int32(1) && libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv))) != int32(m_SQLITE_NULL) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9271:5:
				return Xvec0Update_Update(tls, pVTab, argc, argv)
			} else {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9273:5:
				Xvtab_set_error(tls, pVTab, __ccgo_ts+15288, 0)
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9274:5:
				return int32(m_SQLITE_ERROR)
			}
		}
	}
	return r
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9278:12:
func _vec0ShadowName(tls *libc.TLS, zName uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9278:46:
	var i Tsize_t
	_ = i
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9320:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9320:15:
	i = uint64(0)
	for {
		if !(i < libc.Uint64FromInt64(288)/libc.Uint64FromInt64(8)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9321:5:
		if libsqlite3.Xsqlite3_stricmp(tls, zName, _azName[i]) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9322:7:
			return int32(1)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	//for(size_t i = 0; i < )"vector_chunks", "metadatachunks"
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9325:3:
	return 0
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9279:21:
var _azName = [36]uintptr{
	0:  __ccgo_ts + 9690,
	1:  __ccgo_ts + 15338,
	2:  __ccgo_ts + 15345,
	3:  __ccgo_ts + 15355,
	4:  __ccgo_ts + 15360,
	5:  __ccgo_ts + 15377,
	6:  __ccgo_ts + 15394,
	7:  __ccgo_ts + 15411,
	8:  __ccgo_ts + 15428,
	9:  __ccgo_ts + 15445,
	10: __ccgo_ts + 15462,
	11: __ccgo_ts + 15479,
	12: __ccgo_ts + 15496,
	13: __ccgo_ts + 15513,
	14: __ccgo_ts + 15530,
	15: __ccgo_ts + 15547,
	16: __ccgo_ts + 15564,
	17: __ccgo_ts + 15581,
	18: __ccgo_ts + 15598,
	19: __ccgo_ts + 15615,
	20: __ccgo_ts + 15632,
	21: __ccgo_ts + 15647,
	22: __ccgo_ts + 15662,
	23: __ccgo_ts + 15677,
	24: __ccgo_ts + 15692,
	25: __ccgo_ts + 15707,
	26: __ccgo_ts + 15722,
	27: __ccgo_ts + 15737,
	28: __ccgo_ts + 15752,
	29: __ccgo_ts + 15767,
	30: __ccgo_ts + 15782,
	31: __ccgo_ts + 15797,
	32: __ccgo_ts + 15812,
	33: __ccgo_ts + 15827,
	34: __ccgo_ts + 15842,
	35: __ccgo_ts + 15857,
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9328:12:
func _vec0Begin(tls *libc.TLS, pVTab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9328:43:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9329:3:
	_ = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9330:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9332:12:
func _vec0Sync(tls *libc.TLS, pVTab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9332:42:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9333:3:
	_ = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9334:13:
	p = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9335:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9336:5:
		libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9337:5:
		(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtLatestChunk = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9339:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9340:5:
		libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9341:5:
		(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertRowid = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9343:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9344:5:
		libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9345:5:
		(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsInsertId = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9347:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9348:5:
		libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9349:5:
		(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsUpdatePosition = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9351:3:
	if (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9352:5:
		libsqlite3.Xsqlite3_finalize(tls, (*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9353:5:
		(*Tvec0_vtab)(unsafe.Pointer(p)).FstmtRowidsGetChunkPosition = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9355:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9357:12:
func _vec0Commit(tls *libc.TLS, pVTab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9357:44:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9358:3:
	_ = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9359:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9361:12:
func _vec0Rollback(tls *libc.TLS, pVTab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9361:46:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9362:3:
	_ = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9363:3:
	return m_SQLITE_OK
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9366:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9366:23:
var _vec0Module = Tsqlite3_module{
	FiVersion: int32(3),
}

func init() {
	p := unsafe.Pointer(&_vec0Module)
	*(*uintptr)(unsafe.Add(p, 8)) = __ccgo_fp(_vec0Create)
	*(*uintptr)(unsafe.Add(p, 16)) = __ccgo_fp(_vec0Connect)
	*(*uintptr)(unsafe.Add(p, 24)) = __ccgo_fp(_vec0BestIndex)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(_vec0Disconnect)
	*(*uintptr)(unsafe.Add(p, 40)) = __ccgo_fp(_vec0Destroy)
	*(*uintptr)(unsafe.Add(p, 48)) = __ccgo_fp(_vec0Open)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_vec0Close)
	*(*uintptr)(unsafe.Add(p, 64)) = __ccgo_fp(_vec0Filter)
	*(*uintptr)(unsafe.Add(p, 72)) = __ccgo_fp(_vec0Next)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(_vec0Eof)
	*(*uintptr)(unsafe.Add(p, 88)) = __ccgo_fp(_vec0Column)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(_vec0Rowid)
	*(*uintptr)(unsafe.Add(p, 104)) = __ccgo_fp(_vec0Update)
	*(*uintptr)(unsafe.Add(p, 112)) = __ccgo_fp(_vec0Begin)
	*(*uintptr)(unsafe.Add(p, 120)) = __ccgo_fp(_vec0Sync)
	*(*uintptr)(unsafe.Add(p, 128)) = __ccgo_fp(_vec0Commit)
	*(*uintptr)(unsafe.Add(p, 136)) = __ccgo_fp(_vec0Rollback)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(_vec0ShadowName)
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9397:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9397:13:
var _POINTER_NAME_STATIC_BLOB_DEF = __ccgo_ts + 15872

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9398:1:
type Tstatic_blob_definition = struct {
	Fp            uintptr
	Fdimensions   Tsize_t
	Fnvectors     Tsize_t
	Felement_type _VectorElementType
}

type static_blob_definition = Tstatic_blob_definition

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9404:13:
func _vec_static_blob_from_raw(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9405:60:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9407:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9409:3:
	p = libsqlite3.Xsqlite3_malloc(tls, int32(32))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9410:3:
	if !(p != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9411:5:
		libsqlite3.Xsqlite3_result_error_nomem(tls, context)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9412:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9414:3:
	libc.Xmemset(tls, p, 0, uint64(32))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9415:3:
	(*Tstatic_blob_definition)(unsafe.Pointer(p)).Fp = uintptr(libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9416:3:
	(*Tstatic_blob_definition)(unsafe.Pointer(p)).Felement_type = int32(_SQLITE_VEC_ELEMENT_TYPE_FLOAT32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9417:3:
	(*Tstatic_blob_definition)(unsafe.Pointer(p)).Fdimensions = libc.Uint64FromInt64(libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + 2*8))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9418:3:
	(*Tstatic_blob_definition)(unsafe.Pointer(p)).Fnvectors = libc.Uint64FromInt64(libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + 3*8))))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9419:3:
	libsqlite3.Xsqlite3_result_pointer(tls, context, p, _POINTER_NAME_STATIC_BLOB_DEF, __ccgo_fp(libsqlite3.Xsqlite3_free))
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9426:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9426:28:
type Tstatic_blob = struct {
	Fname         uintptr
	Fp            uintptr
	Fdimensions   Tsize_t
	Fnvectors     Tsize_t
	Felement_type _VectorElementType
}

type static_blob = Tstatic_blob

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9435:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9435:37:
type Tvec_static_blob_data = struct {
	Fstatic_blobs [16]Tstatic_blob
}

type vec_static_blob_data = Tvec_static_blob_data

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9440:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9440:38:
type Tvec_static_blobs_vtab = struct {
	Fbase Tsqlite3_vtab
	Fdata uintptr
}

type vec_static_blobs_vtab = Tvec_static_blobs_vtab

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9446:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9446:40:
type Tvec_static_blobs_cursor = struct {
	Fbase   Tsqlite3_vtab_cursor
	FiRowid Tsqlite3_int64
}

type vec_static_blobs_cursor = Tvec_static_blobs_cursor

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9452:12:
func _vec_static_blobsConnect(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9454:73:
	var pNew uintptr
	var rc int32
	_, _ = pNew, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9455:3:
	_ = argc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9456:3:
	_ = argv
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9457:3:
	_ = pzErr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9464:7:
	rc = libsqlite3.Xsqlite3_declare_vtab(tls, db, __ccgo_ts+15893)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9466:3:
	if rc == m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9467:5:
		pNew = libsqlite3.Xsqlite3_malloc(tls, int32(32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9468:5:
		**(**uintptr)(__ccgo_up(ppVtab)) = pNew
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9469:5:
		if pNew == uintptr(0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9470:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9471:5:
		libc.Xmemset(tls, pNew, 0, uint64(32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9472:5:
		(*Tvec_static_blobs_vtab)(unsafe.Pointer(pNew)).Fdata = pAux
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9474:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9477:12:
func _vec_static_blobsDisconnect(tls *libc.TLS, pVtab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9477:60:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9478:25:
	p = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9479:3:
	libsqlite3.Xsqlite3_free(tls, p)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9480:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9483:12:
func _vec_static_blobsUpdate(tls *libc.TLS, pVTab uintptr, argc int32, argv uintptr, pRowid uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9484:79:
	var def, key, p uintptr
	var i, idx int32
	_, _, _, _, _ = def, i, idx, key, p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9485:3:
	_ = pRowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9486:25:
	p = pVTab
	// DELETE operation
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9488:3:
	if argc == int32(1) && libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv))) != int32(m_SQLITE_NULL) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9489:5:
		return int32(m_SQLITE_ERROR)
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9492:8:
		if argc > int32(1) && libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv))) == int32(m_SQLITE_NULL) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9493:16:
			key = libsqlite3.Xsqlite3_value_text(tls, **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC_STATIC_BLOBS_NAME))*8)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9495:9:
			idx = -int32(1)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9496:5:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9496:14:
			i = 0
			for {
				if !(i < int32(m_MAX_STATIC_BLOBS)) {
					break
				}
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9497:7:
				if !((**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr(i)*40))).Fname != 0) {
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9498:9:
					(**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr(i)*40))).Fname = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+6600, libc.VaList(bp+8, key))
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9499:9:
					idx = i
					// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9500:9:
					break
				}
				goto _1
			_1:
				;
				i = i + 1
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9503:5:
			if idx < 0 {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9504:7:
				libc.Xabort(tls)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9505:35:
			def = libsqlite3.Xsqlite3_value_pointer(tls, **(**uintptr)(__ccgo_up(argv + uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(m_VEC_STATIC_BLOBS_DATA))*8)), _POINTER_NAME_STATIC_BLOB_DEF)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9507:5:
			(**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr(idx)*40))).Fp = (*Tstatic_blob_definition)(unsafe.Pointer(def)).Fp
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9508:5:
			(**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr(idx)*40))).Fdimensions = (*Tstatic_blob_definition)(unsafe.Pointer(def)).Fdimensions
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9509:5:
			(**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr(idx)*40))).Fnvectors = (*Tstatic_blob_definition)(unsafe.Pointer(def)).Fnvectors
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9510:5:
			(**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr(idx)*40))).Felement_type = (*Tstatic_blob_definition)(unsafe.Pointer(def)).Felement_type
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9512:5:
			return m_SQLITE_OK
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9515:8:
			if argc > int32(1) && libsqlite3.Xsqlite3_value_type(tls, **(**uintptr)(__ccgo_up(argv))) != int32(m_SQLITE_NULL) {
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9516:5:
				return int32(m_SQLITE_ERROR)
			}
		}
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9518:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9521:12:
func _vec_static_blobsOpen(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9522:65:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9523:3:
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9525:3:
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(16))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9526:3:
	if pCur == uintptr(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9527:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9528:3:
	libc.Xmemset(tls, pCur, 0, uint64(16))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9529:3:
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9530:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9533:12:
func _vec_static_blobsClose(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9533:60:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9534:27:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9535:3:
	libsqlite3.Xsqlite3_free(tls, pCur)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9536:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9539:12:
func _vec_static_blobsBestIndex(tls *libc.TLS, pVTab uintptr, pIdxInfo uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9540:68:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9541:3:
	_ = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9542:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxNum = int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9543:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = libc.Float64FromInt32(10)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9544:3:
	(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(10)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9545:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9549:12:
func _vec_static_blobsFilter(tls *libc.TLS, pVtabCursor uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9551:57:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9552:3:
	_ = idxNum
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9553:3:
	_ = idxStr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9554:3:
	_ = argc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9555:3:
	_ = argv
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9556:27:
	pCur = pVtabCursor
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9557:3:
	(*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid = int64(-int32(1))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9558:3:
	_vec_static_blobsNext(tls, pVtabCursor)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9559:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9562:12:
func _vec_static_blobsRowid(tls *libc.TLS, cur uintptr, pRowid uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9563:56:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9564:27:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9565:3:
	**(**Tsqlite_int64)(__ccgo_up(pRowid)) = (*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9566:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9569:12:
func _vec_static_blobsNext(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9569:59:
	var p, pCur uintptr
	_, _ = p, pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9570:27:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9571:25:
	p = (*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).Fbase.FpVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9572:3:
	(*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid = (*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid + 1
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9573:3:
	for (*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid < int64(m_MAX_STATIC_BLOBS) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9574:5:
		if (**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr((*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid)*40))).Fname != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9575:7:
			return m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9577:5:
		(*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid = (*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9579:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9582:12:
func _vec_static_blobsEof(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9582:58:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9583:27:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9584:3:
	return libc.BoolInt32((*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid >= int64(m_MAX_STATIC_BLOBS))
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9587:12:
func _vec_static_blobsColumn(tls *libc.TLS, cur uintptr, context uintptr, i int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9588:68:
	var p, pCur uintptr
	_, _ = p, pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9589:27:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9590:25:
	p = (*Tsqlite3_vtab_cursor)(unsafe.Pointer(cur)).FpVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9591:3:
	switch i {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9592:3:
	case m_VEC_STATIC_BLOBS_NAME:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9593:5:
		libsqlite3.Xsqlite3_result_text(tls, context, (**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr((*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid)*40))).Fname, -int32(1), uintptr(-libc.Int32FromInt32(1)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9595:5:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9596:3:
		fallthrough
	case int32(m_VEC_STATIC_BLOBS_DATA):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9597:5:
		libsqlite3.Xsqlite3_result_null(tls, context)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9598:5:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9599:3:
		fallthrough
	case int32(m_VEC_STATIC_BLOBS_DIMENSIONS):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9600:5:
		libsqlite3.Xsqlite3_result_int64(tls, context, libc.Int64FromUint64((**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr((*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid)*40))).Fdimensions))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9602:5:
		break
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9603:3:
		fallthrough
	case int32(m_VEC_STATIC_BLOBS_COUNT):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9604:5:
		libsqlite3.Xsqlite3_result_int64(tls, context, libc.Int64FromUint64((**(**Tstatic_blob)(__ccgo_up((*Tvec_static_blobs_vtab)(unsafe.Pointer(p)).Fdata + uintptr((*Tvec_static_blobs_cursor)(unsafe.Pointer(pCur)).FiRowid)*40))).Fnvectors))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9605:5:
		break
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9607:3:
	return m_SQLITE_OK
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9610:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9610:23:
var _vec_static_blobsModule = Tsqlite3_module{
	FiVersion: int32(3),
}

func init() {
	p := unsafe.Pointer(&_vec_static_blobsModule)
	*(*uintptr)(unsafe.Add(p, 16)) = __ccgo_fp(_vec_static_blobsConnect)
	*(*uintptr)(unsafe.Add(p, 24)) = __ccgo_fp(_vec_static_blobsBestIndex)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(_vec_static_blobsDisconnect)
	*(*uintptr)(unsafe.Add(p, 48)) = __ccgo_fp(_vec_static_blobsOpen)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_vec_static_blobsClose)
	*(*uintptr)(unsafe.Add(p, 64)) = __ccgo_fp(_vec_static_blobsFilter)
	*(*uintptr)(unsafe.Add(p, 72)) = __ccgo_fp(_vec_static_blobsNext)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(_vec_static_blobsEof)
	*(*uintptr)(unsafe.Add(p, 88)) = __ccgo_fp(_vec_static_blobsColumn)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(_vec_static_blobsRowid)
	*(*uintptr)(unsafe.Add(p, 104)) = __ccgo_fp(_vec_static_blobsUpdate)
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9643:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9643:45:
type Tvec_static_blob_entries_vtab = struct {
	Fbase Tsqlite3_vtab
	Fblob uintptr
}

type vec_static_blob_entries_vtab = Tvec_static_blob_entries_vtab

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9648:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9651:3:
type Tvec_sbe_query_plan = int32

type vec_sbe_query_plan = Tvec_sbe_query_plan

const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9649:3:
_VEC_SBE__QUERYPLAN_FULLSCAN = 1
const

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9650:3:
_VEC_SBE__QUERYPLAN_KNN = 2

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9653:1:
type Tsbe_query_knn_data = struct {
	Fk           Ti64
	Fk_used      Ti64
	Frowids      uintptr
	Fdistances   uintptr
	Fcurrent_idx Ti64
}

type sbe_query_knn_data = Tsbe_query_knn_data

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9662:6:
func Xsbe_query_knn_data_clear(tls *libc.TLS, knn_data uintptr) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9662:68:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9663:3:
	if !(knn_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9664:5:
		return
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9666:3:
	if (*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Frowids != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9667:5:
		libsqlite3.Xsqlite3_free(tls, (*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Frowids)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9668:5:
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Frowids = libc.UintptrFromInt32(0)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9670:3:
	if (*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances != 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9671:5:
		libsqlite3.Xsqlite3_free(tls, (*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9672:5:
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances = libc.UintptrFromInt32(0)
	}
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9676:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9676:47:
type Tvec_static_blob_entries_cursor = struct {
	Fbase       Tsqlite3_vtab_cursor
	FiRowid     Tsqlite3_int64
	Fquery_plan Tvec_sbe_query_plan
	Fknn_data   uintptr
}

type vec_static_blob_entries_cursor = Tvec_static_blob_entries_cursor

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9684:12:
func _vec_static_blob_entriesConnect(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9686:80:
	var blob_data, pNew uintptr
	var i, idx, rc int32
	_, _, _, _, _ = blob_data, i, idx, pNew, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9687:3:
	_ = argc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9688:3:
	_ = argv
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9689:3:
	_ = pzErr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9690:24:
	blob_data = pAux
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9691:7:
	idx = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9692:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9692:12:
	i = 0
	for {
		if !(i < int32(m_MAX_STATIC_BLOBS)) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9693:5:
		if !((**(**Tstatic_blob)(__ccgo_up(blob_data + uintptr(i)*40))).Fname != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9694:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9695:5:
		if libc.Xstrncmp(tls, (**(**Tstatic_blob)(__ccgo_up(blob_data + uintptr(i)*40))).Fname, **(**uintptr)(__ccgo_up(argv + 3*8)), libc.Xstrlen(tls, (**(**Tstatic_blob)(__ccgo_up(blob_data + uintptr(i)*40))).Fname)) == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9697:7:
			idx = i
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9698:7:
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9701:3:
	if idx < 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9702:5:
		libc.Xabort(tls)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9707:7:
	rc = libsqlite3.Xsqlite3_declare_vtab(tls, db, __ccgo_ts+15953)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9709:3:
	if rc == m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9710:5:
		pNew = libsqlite3.Xsqlite3_malloc(tls, int32(32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9711:5:
		**(**uintptr)(__ccgo_up(ppVtab)) = pNew
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9712:5:
		if pNew == uintptr(0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9713:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9714:5:
		libc.Xmemset(tls, pNew, 0, uint64(32))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9715:5:
		(*Tvec_static_blob_entries_vtab)(unsafe.Pointer(pNew)).Fblob = blob_data + uintptr(idx)*40
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9717:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9720:12:
func _vec_static_blob_entriesCreate(tls *libc.TLS, db uintptr, pAux uintptr, argc int32, argv uintptr, ppVtab uintptr, pzErr uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9722:79:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9723:3:
	return _vec_static_blob_entriesConnect(tls, db, pAux, argc, argv, ppVtab, pzErr)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9726:12:
func _vec_static_blob_entriesDisconnect(tls *libc.TLS, pVtab uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9726:67:
	var p uintptr
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9727:32:
	p = pVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9728:3:
	libsqlite3.Xsqlite3_free(tls, p)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9729:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9732:12:
func _vec_static_blob_entriesOpen(tls *libc.TLS, p uintptr, ppCursor uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9733:72:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9734:3:
	_ = p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9736:3:
	pCur = libsqlite3.Xsqlite3_malloc(tls, int32(32))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9737:3:
	if pCur == uintptr(0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9738:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9739:3:
	libc.Xmemset(tls, pCur, 0, uint64(32))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9740:3:
	**(**uintptr)(__ccgo_up(ppCursor)) = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9741:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9744:12:
func _vec_static_blob_entriesClose(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9744:67:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9745:34:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9746:3:
	libsqlite3.Xsqlite3_free(tls, (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9747:3:
	libsqlite3.Xsqlite3_free(tls, pCur)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9748:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9751:12:
func _vec_static_blob_entriesBestIndex(tls *libc.TLS, pVTab uintptr, pIdxInfo uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9752:75:
	var i, iColumn, iKTerm, iLimitTerm, iMatchTerm, op int32
	var p uintptr
	_, _, _, _, _, _, _ = i, iColumn, iKTerm, iLimitTerm, iMatchTerm, op, p
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9753:32:
	p = pVTab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9754:7:
	iMatchTerm = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9755:7:
	iLimitTerm = -int32(1)
	// int iRowidTerm = -1; // https://github.com/asg017/sqlite-vec/issues/47
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9757:7:
	iKTerm = -int32(1)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9759:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9759:12:
	i = 0
	for {
		if !(i < (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnConstraint) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9760:5:
		if !((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).Fusable != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9761:7:
			goto _1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9763:9:
		iColumn = (**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).FiColumn
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9764:9:
		op = libc.Int32FromUint8((**(**Tsqlite3_index_constraint)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraint + uintptr(i)*12))).Fop)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9765:5:
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_MATCH) && iColumn == m_VEC_STATIC_BLOB_ENTRIES_VECTOR {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9767:7:
			if iMatchTerm > -int32(1) {
				// https://github.com/asg017/sqlite-vec/issues/51
				// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9769:9:
				return int32(m_SQLITE_ERROR)
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9771:7:
			iMatchTerm = i
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9773:5:
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_LIMIT) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9774:7:
			iLimitTerm = i
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9776:5:
		if op == int32(m_SQLITE_INDEX_CONSTRAINT_EQ) && iColumn == int32(m_VEC_STATIC_BLOB_ENTRIES_K) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9778:7:
			iKTerm = i
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9781:3:
	if iMatchTerm >= 0 {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9782:5:
		if iLimitTerm < 0 && iKTerm < 0 {
			// https://github.com/asg017/sqlite-vec/issues/51
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9784:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9786:5:
		if iLimitTerm >= 0 && iKTerm >= 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9787:7:
			return int32(m_SQLITE_ERROR) // limit or k, not both
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9789:5:
		if (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnOrderBy < int32(1) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9790:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+16003, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9791:7:
			return int32(m_SQLITE_CONSTRAINT)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9793:5:
		if (*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FnOrderBy > int32(1) {
			// https://github.com/asg017/sqlite-vec/issues/51
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9795:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+16030, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9796:7:
			return int32(m_SQLITE_CONSTRAINT)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9798:5:
		if (**(**Tsqlite3_index_orderby)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaOrderBy))).FiColumn != int32(m_VEC_STATIC_BLOB_ENTRIES_DISTANCE) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9799:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+16067, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9800:7:
			return int32(m_SQLITE_CONSTRAINT)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9802:5:
		if (**(**Tsqlite3_index_orderby)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaOrderBy))).Fdesc != 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9803:7:
			Xvtab_set_error(tls, pVTab, __ccgo_ts+8832, 0)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9806:7:
			return int32(m_SQLITE_CONSTRAINT)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9809:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxNum = int32(_VEC_SBE__QUERYPLAN_KNN)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9810:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = libc.Float64FromInt32(10)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9811:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = int64(10)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9813:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).ForderByConsumed = int32(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9814:5:
		(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iMatchTerm)*8))).FargvIndex = int32(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9815:5:
		(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iMatchTerm)*8))).Fomit = uint8(1)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9816:5:
		if iLimitTerm >= 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9817:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iLimitTerm)*8))).FargvIndex = int32(2)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9818:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iLimitTerm)*8))).Fomit = uint8(1)
		} else {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9820:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iKTerm)*8))).FargvIndex = int32(2)
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9821:7:
			(**(**Tsqlite3_index_constraint_usage)(__ccgo_up((*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FaConstraintUsage + uintptr(iKTerm)*8))).Fomit = uint8(1)
		}
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9825:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FidxNum = int32(_VEC_SBE__QUERYPLAN_FULLSCAN)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9826:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedCost = float64((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9827:5:
		(*Tsqlite3_index_info)(unsafe.Pointer(pIdxInfo)).FestimatedRows = libc.Int64FromUint64((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9829:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9832:12:
func _vec_static_blob_entriesFilter(tls *libc.TLS, pVtabCursor uintptr, idxNum int32, idxStr uintptr, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9834:74:
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
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9835:3:
	_ = idxStr
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9836:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9837:34:
	pCur = pVtabCursor
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9839:32:
	p = (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fbase.FpVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9842:3:
	if idxNum == int32(_VEC_SBE__QUERYPLAN_KNN) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9843:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9844:5:
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC_SBE__QUERYPLAN_KNN)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9846:5:
		knn_data = libsqlite3.Xsqlite3_malloc(tls, int32(40))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9847:5:
		if !(knn_data != 0) {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9848:7:
			return int32(m_SQLITE_NOMEM)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9850:5:
		libc.Xmemset(tls, knn_data, 0, uint64(40))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9857:9:
		rc = Xvector_from_value(tls, **(**uintptr)(__ccgo_up(argv)), bp, bp+8, bp+16, bp+24, bp+32)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9859:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9860:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9862:5:
		if **(**_VectorElementType)(__ccgo_up(bp + 16)) != (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Felement_type {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9863:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9865:5:
		if **(**Tsize_t)(__ccgo_up(bp + 8)) != (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9866:7:
			return int32(m_SQLITE_ERROR)
		}
		if libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + 1*8))) <= libc.Int64FromUint64((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors) {
			v1 = libsqlite3.Xsqlite3_value_int64(tls, **(**uintptr)(__ccgo_up(argv + 1*8)))
		} else {
			v1 = libc.Int64FromUint64((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9869:9:
		k = v1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9870:5:
		if k < 0 {
			// HANDLE https://github.com/asg017/sqlite-vec/issues/55
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9872:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9874:5:
		if k == 0 {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9875:7:
			(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fk = 0
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9876:7:
			(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data = knn_data
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9877:7:
			return m_SQLITE_OK
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9880:12:
		bsize = ((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors + uint64(7)) & libc.Uint64FromInt32(^libc.Int32FromInt32(7))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9882:9:
		topk_rowids = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(libc.Uint64FromInt64(k)*uint64(4)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9883:5:
		if !(topk_rowids != 0) {
			// HANDLE https://github.com/asg017/sqlite-vec/issues/55
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9885:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9887:9:
		distances = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(bsize*uint64(4)))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9888:5:
		if !(distances != 0) {
			// HANDLE https://github.com/asg017/sqlite-vec/issues/55
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9890:7:
			return int32(m_SQLITE_ERROR)
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9893:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9893:17:
		i = uint64(0)
		for {
			if !(i < (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors) {
				break
			}
			// https://github.com/asg017/sqlite-vec/issues/52
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9895:13:
			v = (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fp + uintptr(i*(*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions)*4
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9896:7:
			**(**Tf32)(__ccgo_up(distances + uintptr(i)*4)) = _distance_l2_sqr_float(tls, v, **(**uintptr)(__ccgo_up(bp)), (*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob+16)
			goto _2
		_2:
			;
			i = i + 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9899:8:
		candidates = Xbitmap_new(tls, libc.Int32FromUint64(bsize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9900:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9902:8:
		taken = Xbitmap_new(tls, libc.Int32FromUint64(bsize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9903:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9905:5:
		Xbitmap_fill(tls, candidates, libc.Int32FromUint64(bsize))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9906:5:
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9906:17:
		i1 = bsize
		for {
			if !(i1 >= (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors) {
				break
			}
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9907:7:
			Xbitmap_set(tls, candidates, libc.Int32FromUint64(i1), 0)
			goto _3
		_3:
			;
			i1 = i1 - 1
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9909:9:
		**(**Ti32)(__ccgo_up(bp + 40)) = 0
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9910:5:
		Xmin_idx(tls, distances, libc.Int32FromUint64(bsize), candidates, topk_rowids, int32(k), taken, bp+40)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9911:5:
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fcurrent_idx = 0
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9912:5:
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fdistances = distances
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9913:5:
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Fk = k
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9914:5:
		(*Tsbe_query_knn_data)(unsafe.Pointer(knn_data)).Frowids = topk_rowids
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9916:5:
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data = knn_data
	} else {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9918:5:
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan = int32(_VEC_SBE__QUERYPLAN_FULLSCAN)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9919:5:
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid = 0
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9922:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9925:12:
func _vec_static_blob_entriesRowid(tls *libc.TLS, cur uintptr, pRowid uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9926:63:
	var pCur uintptr
	var rowid Ti32
	_, _ = pCur, rowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9927:34:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9928:3:
	switch (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9929:3:
	case int32(_VEC_SBE__QUERYPLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9930:5:
		**(**Tsqlite_int64)(__ccgo_up(pRowid)) = (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9931:5:
		return m_SQLITE_OK
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9933:3:
		fallthrough
	case int32(_VEC_SBE__QUERYPLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9934:9:
		rowid = **(**Ti32)(__ccgo_up((*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*4))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9935:5:
		**(**Tsqlite_int64)(__ccgo_up(pRowid)) = int64(rowid)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9936:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9939:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9942:12:
func _vec_static_blob_entriesNext(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9942:66:
	var pCur uintptr
	_ = pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9943:34:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9944:3:
	switch (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9945:3:
	case int32(_VEC_SBE__QUERYPLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9946:5:
		(*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid = (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid + 1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9947:5:
		return m_SQLITE_OK
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9949:3:
		fallthrough
	case int32(_VEC_SBE__QUERYPLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9950:5:
		(*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx = (*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx + 1
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9951:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9954:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9957:12:
func _vec_static_blob_entriesEof(tls *libc.TLS, cur uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9957:65:
	var p, pCur uintptr
	_, _ = p, pCur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9958:34:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9959:32:
	p = (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fbase.FpVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9961:3:
	switch (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9962:3:
	case int32(_VEC_SBE__QUERYPLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9963:5:
		return libc.BoolInt32(libc.Uint64FromInt64((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid) >= (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fnvectors)
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9965:3:
		fallthrough
	case int32(_VEC_SBE__QUERYPLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9966:5:
		return libc.BoolInt32((*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx >= (*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fk)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9969:3:
	return int32(m_SQLITE_ERROR)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9972:12:
func _vec_static_blob_entriesColumn(tls *libc.TLS, cur uintptr, context uintptr, i int32) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9973:75:
	var p, pCur uintptr
	var rowid Ti32
	_, _, _ = p, pCur, rowid
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9974:34:
	pCur = cur
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9975:32:
	p = (*Tsqlite3_vtab_cursor)(unsafe.Pointer(cur)).FpVtab
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9977:3:
	switch (*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fquery_plan {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9978:3:
	case int32(_VEC_SBE__QUERYPLAN_FULLSCAN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9979:5:
		switch i {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9980:5:
		case m_VEC_STATIC_BLOB_ENTRIES_VECTOR:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9982:7:
			libsqlite3.Xsqlite3_result_blob(tls, context, (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fp+uintptr(libc.Uint64FromInt64((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).FiRowid)*(*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions*libc.Uint64FromInt64(4)), libc.Int32FromUint64((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions*uint64(4)), uintptr(-libc.Int32FromInt32(1)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9987:7:
			libsqlite3.Xsqlite3_result_subtype(tls, context, libc.Uint32FromInt32((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Felement_type))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9988:7:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9990:5:
		return m_SQLITE_OK
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9992:3:
		fallthrough
	case int32(_VEC_SBE__QUERYPLAN_KNN):
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9993:5:
		switch i {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9994:5:
		case m_VEC_STATIC_BLOB_ENTRIES_VECTOR:
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9995:11:
			rowid = **(**Ti32)(__ccgo_up((*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Frowids + uintptr((*Tsbe_query_knn_data)(unsafe.Pointer((*Tvec_static_blob_entries_cursor)(unsafe.Pointer(pCur)).Fknn_data)).Fcurrent_idx)*4))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:9996:7:
			libsqlite3.Xsqlite3_result_blob(tls, context, (*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fp+uintptr(libc.Uint64FromInt32(rowid)*(*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions*libc.Uint64FromInt64(4)), libc.Int32FromUint64((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Fdimensions*uint64(4)), uintptr(-libc.Int32FromInt32(1)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10001:7:
			libsqlite3.Xsqlite3_result_subtype(tls, context, libc.Uint32FromInt32((*Tstatic_blob)(unsafe.Pointer((*Tvec_static_blob_entries_vtab)(unsafe.Pointer(p)).Fblob)).Felement_type))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10002:7:
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10005:5:
		return m_SQLITE_OK
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10008:3:
	return int32(m_SQLITE_ERROR)
}

//
// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10011:1:

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10011:23:
var _vec_static_blob_entriesModule = Tsqlite3_module{
	FiVersion: int32(3),
}

func init() {
	p := unsafe.Pointer(&_vec_static_blob_entriesModule)
	*(*uintptr)(unsafe.Add(p, 8)) = __ccgo_fp(_vec_static_blob_entriesCreate)
	*(*uintptr)(unsafe.Add(p, 16)) = __ccgo_fp(_vec_static_blob_entriesConnect)
	*(*uintptr)(unsafe.Add(p, 24)) = __ccgo_fp(_vec_static_blob_entriesBestIndex)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(_vec_static_blob_entriesDisconnect)
	*(*uintptr)(unsafe.Add(p, 40)) = __ccgo_fp(_vec_static_blob_entriesDisconnect)
	*(*uintptr)(unsafe.Add(p, 48)) = __ccgo_fp(_vec_static_blob_entriesOpen)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_vec_static_blob_entriesClose)
	*(*uintptr)(unsafe.Add(p, 64)) = __ccgo_fp(_vec_static_blob_entriesFilter)
	*(*uintptr)(unsafe.Add(p, 72)) = __ccgo_fp(_vec_static_blob_entriesNext)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(_vec_static_blob_entriesEof)
	*(*uintptr)(unsafe.Add(p, 88)) = __ccgo_fp(_vec_static_blob_entriesColumn)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(_vec_static_blob_entriesRowid)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10064:20:
func Xsqlite3_vec_init(tls *libc.TLS, db uintptr, pzErrMsg uintptr, pApi uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10065:71:
	var i, i1 uint64
	var rc int32
	_, _, _ = i, i1, rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10069:7:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10073:3:
	rc = libsqlite3.Xsqlite3_create_function_v2(tls, db, __ccgo_ts+16107, 0, libc.Int32FromInt32(m_SQLITE_UTF8)|libc.Int32FromInt32(m_SQLITE_INNOCUOUS)|libc.Int32FromInt32(m_SQLITE_DETERMINISTIC), __ccgo_ts+6911, __ccgo_fp(__static_text_func), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10076:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10077:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10079:3:
	rc = libsqlite3.Xsqlite3_create_function_v2(tls, db, __ccgo_ts+16119, 0, libc.Int32FromInt32(m_SQLITE_UTF8)|libc.Int32FromInt32(m_SQLITE_INNOCUOUS)|libc.Int32FromInt32(m_SQLITE_DETERMINISTIC), __ccgo_ts+16129, __ccgo_fp(__static_text_func), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10082:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10083:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10125:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10125:22:
	i = uint64(0)
	for {
		if !(i < libc.Uint64FromInt64(384)/libc.Uint64FromInt64(24) && rc == m_SQLITE_OK) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10126:5:
		rc = libsqlite3.Xsqlite3_create_function_v2(tls, db, _aFunc[i].FzFName, _aFunc[i].FnArg, _aFunc[i].Fflags, libc.UintptrFromInt32(0), _aFunc[i].FxFunc, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10129:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10130:7:
			**(**uintptr)(__ccgo_up(pzErrMsg)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+16398, libc.VaList(bp+8, _aFunc[i].FzFName, libsqlite3.Xsqlite3_errmsg(tls, db)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10132:7:
			return rc
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10136:3:
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10136:22:
	i1 = uint64(0)
	for {
		if !(i1 < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(32) && rc == m_SQLITE_OK) {
			break
		}
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10137:5:
		rc = libsqlite3.Xsqlite3_create_module_v2(tls, db, _aMod[i1].Fname, _aMod[i1].Fmodule, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10138:5:
		if rc != m_SQLITE_OK {
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10139:7:
			**(**uintptr)(__ccgo_up(pzErrMsg)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+16429, libc.VaList(bp+8, _aMod[i1].Fname, libsqlite3.Xsqlite3_errmsg(tls, db)))
			// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10141:7:
			return rc
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10145:3:
	return m_SQLITE_OK
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10090:5:
var _aFunc = [16]struct {
	FzFName uintptr
	FxFunc  uintptr
	FnArg   int32
	Fflags  int32
}{
	0: {
		FzFName: __ccgo_ts + 16176,
		FnArg:   int32(2),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE),
	},
	1: {
		FzFName: __ccgo_ts + 16192,
		FnArg:   int32(2),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE),
	},
	2: {
		FzFName: __ccgo_ts + 16208,
		FnArg:   int32(2),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE),
	},
	3: {
		FzFName: __ccgo_ts + 16229,
		FnArg:   int32(2),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE),
	},
	4: {
		FzFName: __ccgo_ts + 16249,
		FnArg:   int32(1),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE),
	},
	5: {
		FzFName: __ccgo_ts + 16260,
		FnArg:   int32(1),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC),
	},
	6: {
		FzFName: __ccgo_ts + 16269,
		FnArg:   int32(1),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	7: {
		FzFName: __ccgo_ts + 16281,
		FnArg:   int32(2),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	8: {
		FzFName: __ccgo_ts + 16289,
		FnArg:   int32(2),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	9: {
		FzFName: __ccgo_ts + 16297,
		FnArg:   int32(3),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	10: {
		FzFName: __ccgo_ts + 16307,
		FnArg:   int32(1),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	11: {
		FzFName: __ccgo_ts + 16321,
		FnArg:   int32(1),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	12: {
		FzFName: __ccgo_ts + 16329,
		FnArg:   int32(1),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	13: {
		FzFName: __ccgo_ts + 16337,
		FnArg:   int32(1),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	14: {
		FzFName: __ccgo_ts + 16346,
		FnArg:   int32(2),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
	15: {
		FzFName: __ccgo_ts + 16364,
		FnArg:   int32(1),
		Fflags:  libc.Int32FromInt32(m_SQLITE_UTF8) | libc.Int32FromInt32(m_SQLITE_INNOCUOUS) | libc.Int32FromInt32(m_SQLITE_DETERMINISTIC) | libc.Int32FromInt32(m_SQLITE_SUBTYPE) | libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE),
	},
}

func init() {
	p := unsafe.Pointer(&_aFunc)
	*(*uintptr)(unsafe.Add(p, 8)) = __ccgo_fp(_vec_distance_l2)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(_vec_distance_l1)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_vec_distance_hamming)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(_vec_distance_cosine)
	*(*uintptr)(unsafe.Add(p, 104)) = __ccgo_fp(_vec_length)
	*(*uintptr)(unsafe.Add(p, 128)) = __ccgo_fp(_vec_type)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(_vec_to_json)
	*(*uintptr)(unsafe.Add(p, 176)) = __ccgo_fp(_vec_add)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(_vec_sub)
	*(*uintptr)(unsafe.Add(p, 224)) = __ccgo_fp(_vec_slice)
	*(*uintptr)(unsafe.Add(p, 248)) = __ccgo_fp(_vec_normalize)
	*(*uintptr)(unsafe.Add(p, 272)) = __ccgo_fp(_vec_f32)
	*(*uintptr)(unsafe.Add(p, 296)) = __ccgo_fp(_vec_bit)
	*(*uintptr)(unsafe.Add(p, 320)) = __ccgo_fp(_vec_int8)
	*(*uintptr)(unsafe.Add(p, 344)) = __ccgo_fp(_vec_quantize_int8)
	*(*uintptr)(unsafe.Add(p, 368)) = __ccgo_fp(_vec_quantize_binary)
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10118:5:
var _aMod = [2]struct {
	Fname     uintptr
	Fmodule   uintptr
	Fp        uintptr
	FxDestroy uintptr
}{
	0: {
		Fname:   __ccgo_ts + 16384,
		Fmodule: uintptr(unsafe.Pointer(&_vec0Module)),
	},
	1: {
		Fname:   __ccgo_ts + 16389,
		Fmodule: uintptr(unsafe.Pointer(&_vec_eachModule)),
	},
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10149:20:
func Xsqlite3_vec_numpy_init(tls *libc.TLS, db uintptr, pzErrMsg uintptr, pApi uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10150:79:
	var rc int32
	_ = rc
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10151:3:
	_ = pzErrMsg
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10155:7:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10156:3:
	rc = libsqlite3.Xsqlite3_create_function_v2(tls, db, __ccgo_ts+16458, int32(1), int32(m_SQLITE_RESULT_SUBTYPE), libc.UintptrFromInt32(0), __ccgo_fp(_vec_npy_file), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10158:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10159:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10161:3:
	rc = libsqlite3.Xsqlite3_create_module_v2(tls, db, __ccgo_ts+16471, uintptr(unsafe.Pointer(&_vec_npy_eachModule)), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10162:3:
	return rc
}

// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10167:1:
func Xsqlite3_vec_static_blobs_init(tls *libc.TLS, db uintptr, pzErrMsg uintptr, pApi uintptr) (r int32) {
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10168:65:
	var rc int32
	var static_blob_data uintptr
	_, _ = rc, static_blob_data
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10169:3:
	_ = pzErrMsg
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10174:7:
	rc = m_SQLITE_OK
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10176:3:
	static_blob_data = libsqlite3.Xsqlite3_malloc(tls, int32(640))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10177:3:
	if !(static_blob_data != 0) {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10178:5:
		return int32(m_SQLITE_NOMEM)
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10180:3:
	libc.Xmemset(tls, static_blob_data, 0, uint64(640))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10182:3:
	rc = libsqlite3.Xsqlite3_create_function_v2(tls, db, __ccgo_ts+16484, int32(4), libc.Int32FromInt32(m_SQLITE_UTF8)|libc.Int32FromInt32(m_SQLITE_INNOCUOUS)|libc.Int32FromInt32(m_SQLITE_DETERMINISTIC)|libc.Int32FromInt32(m_SQLITE_SUBTYPE)|libc.Int32FromInt32(m_SQLITE_RESULT_SUBTYPE), libc.UintptrFromInt32(0), __ccgo_fp(_vec_static_blob_from_raw), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10186:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10187:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10189:3:
	rc = libsqlite3.Xsqlite3_create_module_v2(tls, db, __ccgo_ts+16509, uintptr(unsafe.Pointer(&_vec_static_blobsModule)), static_blob_data, __ccgo_fp(libsqlite3.Xsqlite3_free))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10191:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10192:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10193:3:
	rc = libsqlite3.Xsqlite3_create_module_v2(tls, db, __ccgo_ts+16526, uintptr(unsafe.Pointer(&_vec_static_blob_entriesModule)), static_blob_data, libc.UintptrFromInt32(0))
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10196:3:
	if rc != m_SQLITE_OK {
		// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10197:5:
		return rc
	}
	// /tmp/sqlite_vec/sqlite-vec-0.1.9/sqlite-vec.c:10198:3:
	return rc
}

// -:
type Tsqlite3_index_constraint = struct {
	FiColumn     int32
	Fop          uint8
	Fusable      uint8
	FiTermOffset int32
}

type sqlite3_index_constraint = Tsqlite3_index_constraint

// -:
type Tsqlite3_index_constraint_usage = struct {
	FargvIndex int32
	Fomit      uint8
}

type sqlite3_index_constraint_usage = Tsqlite3_index_constraint_usage

// -:
type Tsqlite3_index_orderby = struct {
	FiColumn int32
	Fdesc    uint8
}

type sqlite3_index_orderby = Tsqlite3_index_orderby

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "float32\x00int8\x00bit\x00\x00INTEGER\x00BLOB\x00TEXT\x00FLOAT\x00NULL\x00zero-length vectors are not supported.\x00invalid float32 vector BLOB length. Must be divisible by %d, found %d\x00out of memory\x00JSON array parsing error: Input does not start with '['\x00JSON parsing error\x00Input must have type BLOB (compact format) or TEXT (JSON), found %s\x00Unknown type for bitvector.\x00JSON parsing error: value out of range for int8\x00Unknown type for int8 vector.\x00Unknown subtype: %d\x00Error reading 1st vector: %s\x00Error reading 2nd vector: %s\x00Vector type mistmatch. First vector has type %s, while the second has type %s.\x00Vector dimension mistmatch. First vector has %ld dimensions, while the second has %ld dimensions.\x00vec0-npy-file\x00Cannot calculate cosine distance between two bitvectors.\x00Cannot calculate L2 distance between two bitvectors.\x00Cannot calculate L1 distance between two bitvectors.\x00Cannot calculate hamming distance between two float32 vectors.\x00Cannot calculate hamming distance between two int8 vectors.\x00Zero length vectors are not supported.\x00Binary quantization requires vectors with a length divisible by 8\x00Can only binary quantize float or int8 vectors\x00unit\x002nd argument to vec_quantize_int8() must be 'unit'.\x00Cannot add two bitvectors together.\x00Cannot subtract two bitvectors together.\x00slice 'start' index must be a postive number.\x00slice 'end' index must be a postive number.\x00slice 'start' index is greater than the number of dimensions\x00slice 'end' index is greater than the number of dimensions\x00slice 'start' index is greater than 'end' index\x00slice 'start' index is equal to the 'end' index, vectors must have non-zero length\x00start index must be divisible by 8.\x00end index must be divisible by 8.\x00[\x00,\x00null\x00%f\x00%d\x00]\x00only float32 vectors are supported when normalizing\x00text\x00int\x00integer\x00partition\x00key\x00float\x00double\x00blob\x00boolean\x00bool\x00int64\x00integer64\x00float64\x00f64\x00primary\x00f32\x00i8\x00distance_metric\x00l2\x00l1\x00cosine\x00%.*s\x00CREATE TABLE x(value, vector hidden)\x00False\x00Error parsing numpy array: numpy header did not start with '{'\x00Error parsing numpy array: expected key in numpy header\x00Error parsing numpy array: expected a string as key in numpy header\x00Error parsing numpy array: expected a ':' after key in numpy header\x00'descr'\x00Error parsing numpy array: expected a string value after 'descr' key\x00'<f4'\x00Error parsing numpy array: Only '<f4' values are supported in sqlite-vec numpy functions\x00'fortran_order'\x00Error parsing numpy array: Only fortran_order = False is supported in sqlite-vec numpy functions\x00'shape'\x00Error parsing numpy array: Expected left parenthesis '(' after shape key\x00Error parsing numpy array: Expected an initial number in shape value\x00Error parsing numpy array: Expected comma after first shape value\x00Error parsing numpy array: unexpected header EOF while parsing shape\x00Error parsing numpy array: expected right parenthesis after shape value\x00Error parsing numpy array: unknown type in shape value\x00Error parsing numpy array: unknown key in numpy header\x00Error parsing numpy array: unknown extra token after value\x00numpy array file too short\x00numpy array file does not contain the 'magic' header\x00numpy array file header length is invalid\x00numpy array file error: Expected a data size of %d, found %d\x00numpy array too short\x00numpy array does not contain the 'magic' header\x00numpy array header length is invalid\x00numpy array error: Expected a data size of %d, found %d\x00CREATE TABLE x(vector, input hidden)\x00input argument is required\x00r\x00Could not open numpy file\x00vec_npy_each only supports float32 vectors\x00SELECT id, chunk_id, chunk_offset FROM \"%w\".\"%w_rowids\" WHERE rowid = ?\x00Internal sqlite-vec error: could not initialize 'rowids get chunk position' statement\x00SELECT rowid FROM \"%w\".\"%w_rowids\" WHERE id = ?\x00Could not find a row with rowid %lld\x00vectors\x00Could not fetch vector data for %lld, opening blob failed\x00Could not fetch vector data for %lld, reading from blob failed\x00Internal sqlite-vec error: unknown error, could not close vector blob, please file an issue\x00SELECT partition%02d FROM \"%w\".\"%w_chunks\" WHERE chunk_id = ?\x00SELECT value%02d FROM \"%w\".\"%w_auxiliary\" WHERE rowid = ?\x00data\x00SELECT data FROM \"%w\".\"%w_metadatatext%02d\" WHERE rowid = ?\x00SELECT max(rowid) FROM \"%w\".\"%w_chunks\" WHERE \x00 AND \x00 partition%02d = ? \x00SELECT max(rowid) FROM \"%w\".\"%w_chunks\"\x00Internal sqlite-vec error: could not initialize 'latest chunk' statement\x00Internal sqlite-vec error: Could not find latest chunk\x00Internal sqlite-vec error: unknown result code when closing out stmtLatestChunk. Please file an issue: https://github.com/asg017/sqlite-vec/issues/new\x00INSERT INTO \"%w\".\"%w_rowids\"(rowid)VALUES (?);\x00Internal sqlite-vec error: could not initialize 'insert rowids' statement\x00UNIQUE constraint failed on %s primary key\x00Error inserting rowid into rowids shadow table: %s\x00INSERT INTO \"%w\".\"%w_rowids\"(id)VALUES (?);\x00Internal sqlite-vec error: could not initialize 'insert rowids id' statement\x00Error inserting id into rowids shadow table: %s\x00 UPDATE \"%w\".\"%w_rowids\" SET chunk_id = ?, chunk_offset = ? WHERE rowid = ?\x00Internal sqlite-vec error: could not initialize 'update rowids position' statement\x00Internal sqlite-vec error: could not update rowids position for rowid=%lld, chunk_rowid=%lld, chunk_offset=%lld\x00INSERT INTO \"%w\".\"%w_chunks\"\x00(size, validity, rowids\x00, partition%02d\x00) VALUES (?, ?, ?\x00, ?\x00)\x00INSERT INTO \"%w\".\"%w_chunks\"(size, validity, rowids) VALUES (?, ?, ?);\x00INSERT INTO \"%w\".\"%w_vector_chunks%02d\"(_rowid_, rowid, vectors)VALUES (?, ?, ?)\x00INSERT INTO \"%w\".\"%w_metadatachunks%02d\"(_rowid_, rowid, data)VALUES (?, ?, ?)\x00vec0 constructor error: could not parse vector column '%s'\x00vec0 constructor error: Too many provided vector columns, maximum %d\x00vec0 constructor error: Dimension on vector column too large, provided %lld, maximum %lld\x00vec0 constructor error: More than %d partition key columns were provided\x00vec0 constructor error: More than one primary key definition was provided, vec0 only suports a single primary key column\x00vec0 constructor error: More than %d auxiliary columns were provided\x00vec0 constructor error: More than %d metadata columns were provided\x00vec0 constructor error: could not parse table option '%s'\x00chunk_size\x00vec0 constructor error: chunk_size must be a non-zero positive integer\x00vec0 constructor error: chunk_size must be divisible by 8\x00vec0 constructor error: chunk_size too large\x00vec0 constructor error: Unknown table option: %.*s\x00vec0 constructor error: Could not parse '%s'\x00vec0 constructor error: At least one vector column is required\x00CREATE TABLE x(\x00\"%.*w\" primary key, \x00rowid, \x00\"%.*w\", \x00 distance hidden, k hidden) \x00without rowid \x00vec0 constructor error: could not declare virtual table, '%s'\x00%s\x00%s_rowids\x00%s_chunks\x00%s_vector_chunks%02d\x00%s_metadatachunks%02d\x00CREATE TABLE \"%w\".\"%w_info\" (key text primary key, value any)\x00Could not create '_info' shadow table: %s\x00INSERT INTO \"%w\".\"%w_info\"(key, value) VALUES (?1, ?2), (?3, ?4), (?5, ?6), (?7, ?8) \x00Could not seed '_info' shadow table: %s\x00CREATE_VERSION\x00v0.1.9\x00CREATE_VERSION_MAJOR\x00CREATE_VERSION_MINOR\x00CREATE_VERSION_PATCH\x00CREATE TABLE \"%w\".\"%w_chunks\"(\x00chunk_id INTEGER PRIMARY KEY AUTOINCREMENT,size INTEGER NOT NULL,\x00sequence_id integer,\x00partition%02d,\x00validity BLOB NOT NULL, rowids BLOB NOT NULL);\x00CREATE TABLE \"%w\".\"%w_chunks\"(chunk_id INTEGER PRIMARY KEY AUTOINCREMENT,size INTEGER NOT NULL,validity BLOB NOT NULL,rowids BLOB NOT NULL);\x00Could not create '_chunks' shadow table: %s\x00CREATE TABLE \"%w\".\"%w_rowids\"(rowid INTEGER PRIMARY KEY AUTOINCREMENT,id TEXT UNIQUE NOT NULL,chunk_id INTEGER,chunk_offset INTEGER);\x00CREATE TABLE \"%w\".\"%w_rowids\"(rowid INTEGER PRIMARY KEY AUTOINCREMENT,id,chunk_id INTEGER,chunk_offset INTEGER);\x00Could not create '_rowids' shadow table: %s\x00CREATE TABLE \"%w\".\"%w_vector_chunks%02d\"(rowid PRIMARY KEY,vectors BLOB NOT NULL);\x00Could not create '_vector_chunks%02d' shadow table: %s\x00CREATE TABLE \"%w\".\"%w_metadatachunks%02d\"(rowid PRIMARY KEY, data BLOB NOT NULL);\x00Could not create '_metata_chunks%02d' shadow table: %s\x00CREATE TABLE \"%w\".\"%w_metadatatext%02d\"(rowid PRIMARY KEY, data TEXT);\x00Could not create '_metadatatext%02d' shadow table: %s\x00CREATE TABLE \"%w\".\"%w_auxiliary\"( rowid integer PRIMARY KEY \x00, value%02d\x00Could not create auxiliary shadow table: %s\x00DROP TABLE \"%w\".\"%w_chunks\"\x00could not drop chunks shadow table\x00DROP TABLE \"%w\".\"%w_info\"\x00could not drop info shadow table\x00DROP TABLE \"%w\".\"%w_rowids\"\x00DROP TABLE \"%w\".\"%w\"\x00DROP TABLE \"%w\".\"%w_auxiliary\"\x00DROP TABLE \"%w\".\"%w_metadatachunks%02d\"\x00DROP TABLE \"%w\".\"%w_metadatatext%02d\"\x00only 1 MATCH operator is allowed in a single vec0 query\x00only 1 'rowid in (..)' operator is allowed in a single vec0 query\x00A LIMIT or 'k = ?' constraint is required on vec0 knn queries.\x00Only LIMIT or 'k =?' can be provided, not both\x00Only a single 'ORDER BY distance' clause is allowed on vec0 KNN queries\x00Only a single 'ORDER BY distance' clause is allowed on vec0 KNN queries, not on other columns\x00Only ascending in ORDER BY distance clause is supported, DESC is not supported yet.\x00An illegal WHERE constraint was provided on a vec0 auxiliary column in a KNN query.\x00'xxx in (...)' is only available on INTEGER or TEXT metadata columns.\x00An illegal WHERE constraint was provided on a vec0 metadata column in a KNN query. Only one of EQUALS, GREATER_THAN, LESS_THAN_OR_EQUAL, LESS_THAN, GREATER_THAN_OR_EQUAL, NOT_EQUALS is allowed.\x00ONLY EQUALS (=) or NOT_EQUALS (!=) operators are allowed on boolean metadata columns.\x00Illegal WHERE constraint on distance column in a KNN query. Only one of GT, GE, LT, LE constraints are allowed.\x00select data from \"%w\".\"%w_metadatatext%02d\" where rowid = ?\x00select chunk_id, validity, rowids  from \"%w\".\"%w_chunks\"\x00 WHERE \x00 partition%02d > ? \x00 partition%02d <= ? \x00 partition%02d < ? \x00 partition%02d >= ? \x00 partition%02d != ? \x00rowids\x00chunks iter error\x00chunk validity size doesn't match - expected %lld, found %lld\x00rowids size doesn't match\x00chunk rowids size doesn't match - expected %lld, found %lld\x00could not open vectors blob for chunk %lld\x00vectors blob size doesn't match - expected %lld, found %lld\x00vectors blob read error for %lld\x00Could not open metadata blob\x00Could not filter metadata fields\x00Query vector on the \"%.*s\" column is invalid: %z\x00Query vector for the \"%.*s\" column is expected to be of type %s, but a %s vector was provided.\x00Dimension mismatch for query vector for the \"%.*s\" column. Expected %d dimensions but received %d.\x00k value in knn queries must be greater than or equal to 0.\x00k value in knn query too large, provided %lld and the limit is %lld\x00error processing rowid in (...) array\x00Error fetching next value in `x in (...)` integer expression\x00Error fetching next value in `x in (...)` text expression\x00Internal sqlite-vec error\x00Error preparing stmtChunk: %s\x00 SELECT rowid  FROM \"%w\".\"%w_rowids\" ORDER by chunk_id, chunk_offset \x00Error preparing rowid scan: %s\x00unknown idxStr '%s'\x00Internal sqlite-vec error: expected point query plan in vec0Rowid, found %d\x00Internal sqlite-vec error: fullscan_data is NULL.\x00Could not extract metadata value for column %.*s at rowid %lld\x00Internal sqlite-vec error: point_data is NULL.\x00Internal sqlite-vec error: knn_data is NULL.\x00The %s virtual table was declared with a TEXT primary key, but a non-TEXT value was provided in an INSERT.\x00Only integers are allows for primary key values on %s\x00validity\x00Internal sqlite-vec error: could not open validity blob on %s.%s.%lld\x00Internal sqlite-vec error: validity blob size mismatch on %s.%s.%lld, expected %lld but received %lld.\x00Internal sqlite-vec error: Could not allocate memory for validity bitmap\x00Internal sqlite-vec error: Could not read validity bitmap for %s.%s.%lld\x00Internal sqlite-vec error: Could not insert a new vector chunk\x00Internal sqlite-vec error: unknown error, blobChunksValidity could not be closed, please file an issue.\x00Internal sqlite-vec error: Could not open validity blob for newly created chunk %s.%s.%lld\x00Internal sqlite-vec error: validity blob size mismatch for newly created chunk %s.%s.%lld. Exepcted %lld, got %lld\x00Internal sqlite-vec error: could not read validity blob newly created chunk %s.%s.%lld\x00Internal sqlite-vec error: could not mark validity bit \x00Error opening vector blob at %s.%s.%lld\x00Internal sqlite-vec error: vector blob size mismatch on %s.%s.%lld. Expected %lld, actual %lld\x00Internal sqlite-vec error: could not write vector blob on %s.%s.%lld\x00Internal sqlite-vec error: could not close vector blob on %s.%s.%lld\x00Internal sqlite-vec error: could not open rowids blob on %s.%s.%lld\x00Internal sqlite-vec error: rowids blob size mismatch on %s.%s.%lld. Expected %lld, actual %lld\x00Internal sqlite-vec error: could not write rowids blob on %s.%s.%lld\x00Internal sqlite-vec error: could not close rowids blob on %s.%s.%lld\x00Expected 0 or 1 for BOOLEAN metadata column %.*s\x00Expected integer for INTEGER metadata column %.*s, received %s\x00Expected float for FLOAT metadata column %.*s, received %s\x00Expected text for TEXT metadata column %.*s, received %s\x00UPDATE \"%w\".\"%w_metadatatext%02d\" SET data = ?2 WHERE rowid = ?1\x00INSERT INTO \"%w\".\"%w_metadatatext%02d\" (rowid, data) VALUES (?1, ?2)\x00DELETE FROM \"%w\".\"%w_metadatatext%02d\" WHERE rowid = ?\x00Parition key type mismatch: The partition key column %.*s has type %s, but %s was provided.\x00Inserted vector for the \"%.*s\" column is invalid: %z\x00Inserted vector for the \"%.*s\" column is expected to be of type %s, but a %s vector was provided.\x00Dimension mismatch for inserted vector for the \"%.*s\" column. Expected %d dimensions but received %d.\x00A value was provided for the hidden \"distance\" column.\x00A value was provided for the hidden \"k\" column.\x00INSERT INTO \"%w\".\"%w_auxiliary\"(rowid \x00) VALUES (? \x00Auxiliary column type mismatch: The auxiliary column %.*s has type %s, but %s was provided.\x00Internal sqlite-vec error: unknown error, blobChunksValidity could not be closed, please file an issue\x00could not open validity blob for %s.%s.%lld\x00could not read validity blob for %s.%s.%lld at %d\x00vec0 deletion error: validity bit is not set for %s.%s.%lld at %d\x00could not write to validity blob for %s.%s.%lld at %d\x00vec0 deletion error: Error commiting validity blob transaction on %s.%s.%lld at %d\x00could not open rowids blob for %s.%s.%lld\x00could not write to rowids blob for %s.%s.%lld at %llu\x00vec0 deletion error: Error commiting rowids blob transaction on %s.%s.%lld at %llu\x00could not open vector blob for %s.%s.%lld column %d\x00could not write to vector blob for %s.%s.%lld at %llu column %d\x00vec0 deletion error: Error commiting vector blob transaction on %s.%s.%lld column %d\x00could not open validity blob for chunk %lld\x00DELETE FROM \"%w\".\"%w_chunks\" WHERE rowid = ?\x00DELETE FROM \"%w\".\"%w_vector_chunks%02d\" WHERE rowid = ?\x00DELETE FROM \"%w\".\"%w_metadatachunks%02d\" WHERE rowid = ?\x00DELETE FROM \"%w\".\"%w_rowids\" WHERE rowid = ?\x00DELETE FROM \"%w\".\"%w_auxiliary\" WHERE rowid = ?\x00UPDATE \"%w\".\"%w_auxiliary\" SET value%02d = ? WHERE rowid = ?\x00Updated vector for the \"%.*s\" column is invalid: %z\x00Updated vector for the \"%.*s\" column is expected to be of type %s, but a %s vector was provided.\x00Dimension mismatch for new updated vector for the \"%.*s\" column. Expected %d dimensions but received %d.\x00Could not open vectors blob for %s.%s.%lld\x00Could not write to vectors blob for %s.%s.%lld\x00Could not commit blob transaction for vectors blob for %s.%s.%lld\x00UPDATEs on vec0 primary key values are not allowed.\x00UPDATE on partition key columns are not supported yet. \x00Unrecognized xUpdate operation provided for vec0.\x00chunks\x00auxiliary\x00info\x00metadatachunks00\x00metadatachunks01\x00metadatachunks02\x00metadatachunks03\x00metadatachunks04\x00metadatachunks05\x00metadatachunks06\x00metadatachunks07\x00metadatachunks08\x00metadatachunks09\x00metadatachunks10\x00metadatachunks11\x00metadatachunks12\x00metadatachunks13\x00metadatachunks14\x00metadatachunks15\x00metadatatext00\x00metadatatext01\x00metadatatext02\x00metadatatext03\x00metadatatext04\x00metadatatext05\x00metadatatext06\x00metadatatext07\x00metadatatext08\x00metadatatext09\x00metadatatext10\x00metadatatext11\x00metadatatext12\x00metadatatext13\x00metadatatext14\x00metadatatext15\x00vec0-static_blob_def\x00CREATE TABLE x(name, data, dimensions hidden, count hidden)\x00CREATE TABLE x(vector, distance hidden, k hidden)\x00ORDER BY distance required\x00more than 1 ORDER BY clause provided\x00ORDER BY must be on the distance column\x00vec_version\x00vec_debug\x00Version: v0.1.9\nDate: \nCommit: \nBuild flags:  \x00vec_distance_l2\x00vec_distance_l1\x00vec_distance_hamming\x00vec_distance_cosine\x00vec_length\x00vec_type\x00vec_to_json\x00vec_add\x00vec_sub\x00vec_slice\x00vec_normalize\x00vec_f32\x00vec_bit\x00vec_int8\x00vec_quantize_int8\x00vec_quantize_binary\x00vec0\x00vec_each\x00Error creating function %s: %s\x00Error creating module %s: %s\x00vec_npy_file\x00vec_npy_each\x00vec_static_blob_from_raw\x00vec_static_blobs\x00vec_static_blob_entries\x00"
