// Code generated for darwin/arm64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build darwin && arm64

package vec

const m_TARGET_OS_ARROW = 1

const m__ARM_SIGNAL_ = 1

const m__DARWIN_FEATURE_ONLY_64_BIT_INODE = 1

const m__DARWIN_FEATURE_ONLY_VERS_1050 = 1

const m___AARCH64_SIMD__ = 1

const m___ARM64_ARCH_8__ = 1

const m___ARM_ACLE = 202420

const m___ARM_ARCH_8_2__ = 1

const m___ARM_ARCH_8_3__ = 1

const m___ARM_ARCH_8_4__ = 1

const m___ARM_ARCH_8_5__ = 1

const m___ARM_FEATURE_AES = 1

const m___ARM_FEATURE_ATOMICS = 1

const m___ARM_FEATURE_BTI = 1

const m___ARM_FEATURE_COMPLEX = 1

const m___ARM_FEATURE_CRC32 = 1

const m___ARM_FEATURE_CRYPTO = 1

const m___ARM_FEATURE_DOTPROD = 1

const m___ARM_FEATURE_FP16_FML = 1

const m___ARM_FEATURE_FP16_SCALAR_ARITHMETIC = 1

const m___ARM_FEATURE_FP16_VECTOR_ARITHMETIC = 1

const m___ARM_FEATURE_FRINT = 1

const m___ARM_FEATURE_JCVT = 1

const m___ARM_FEATURE_PAUTH = 1

const m___ARM_FEATURE_QRDMX = 1

const m___ARM_FEATURE_RCPC = 1

const m___ARM_FEATURE_SHA2 = 1

const m___ARM_FEATURE_SHA3 = 1

const m___ARM_FEATURE_SHA512 = 1

const m___ARM_NEON_SVE_BRIDGE = 1

const m___DARWIN_ONLY_64_BIT_INO_T = 1

const m___DARWIN_ONLY_VERS_1050 = 1

const m___DARWIN_OPAQUE_ARM_THREAD_STATE64 = 0

const m___FUNCTION_MULTI_VERSIONING_SUPPORT_LEVEL = 202430

const m___OBJC_BOOL_IS_BOOL = 1

const m___arm64 = 1

const m___arm64__ = 1

type t__arm_legacy_debug_state = struct {
	F__bvr [16]t__uint32_t
	F__bcr [16]t__uint32_t
	F__wvr [16]t__uint32_t
	F__wcr [16]t__uint32_t
}

type t__arm_pagein_state = struct {
	F__pagein_error int32
}

type t__darwin_arm_cpmu_state64 = struct {
	F__ctrs [16]t__uint64_t
}

type t__darwin_arm_debug_state32 = struct {
	F__bvr       [16]t__uint32_t
	F__bcr       [16]t__uint32_t
	F__wvr       [16]t__uint32_t
	F__wcr       [16]t__uint32_t
	F__mdscr_el1 t__uint64_t
}

type t__darwin_arm_debug_state64 = struct {
	F__bvr       [16]t__uint64_t
	F__bcr       [16]t__uint64_t
	F__wvr       [16]t__uint64_t
	F__wcr       [16]t__uint64_t
	F__mdscr_el1 t__uint64_t
}

type t__darwin_arm_exception_state = struct {
	F__exception t__uint32_t
	F__fsr       t__uint32_t
	F__far       t__uint32_t
}

type t__darwin_arm_exception_state64 = struct {
	F__far       t__uint64_t
	F__esr       t__uint32_t
	F__exception t__uint32_t
}

type t__darwin_arm_exception_state64_v2 = struct {
	F__far t__uint64_t
	F__esr t__uint64_t
}

type t__darwin_arm_neon_state = struct {
	F__ccgo_align [0]uint64
	F__v          [16][2]uint64
	F__fpsr       t__uint32_t
	F__fpcr       t__uint32_t
	F__ccgo_pad3  [8]byte
}

type t__darwin_arm_neon_state64 = struct {
	F__ccgo_align [0]uint64
	F__v          [32][2]uint64
	F__fpsr       t__uint32_t
	F__fpcr       t__uint32_t
	F__ccgo_pad3  [8]byte
}

type t__darwin_arm_sme2_state = struct {
	F__ccgo_align [0]uint32
	F__zt0        [64]int8
}

type t__darwin_arm_sme_state = struct {
	F__svcr       t__uint64_t
	F__tpidr2_el0 t__uint64_t
	F__svl_b      t__uint16_t
}

type t__darwin_arm_sme_za_state = struct {
	F__ccgo_align [0]uint32
	F__za         [4096]int8
}

type t__darwin_arm_sve_p_state = struct {
	F__ccgo_align [0]uint32
	F__p          [16][32]int8
}

type t__darwin_arm_sve_z_state = struct {
	F__ccgo_align [0]uint32
	F__z          [16][256]int8
}

type t__darwin_arm_thread_state = struct {
	F__r    [13]t__uint32_t
	F__sp   t__uint32_t
	F__lr   t__uint32_t
	F__pc   t__uint32_t
	F__cpsr t__uint32_t
}

type t__darwin_arm_thread_state64 = struct {
	F__x    [29]t__uint64_t
	F__fp   t__uint64_t
	F__lr   t__uint64_t
	F__sp   t__uint64_t
	F__pc   t__uint64_t
	F__cpsr t__uint32_t
	F__pad  t__uint32_t
}

type t__darwin_arm_vfp_state = struct {
	F__r     [64]t__uint32_t
	F__fpscr t__uint32_t
}

type t__darwin_mcontext32 = struct {
	F__es t__darwin_arm_exception_state
	F__ss t__darwin_arm_thread_state
	F__fs t__darwin_arm_vfp_state
}

type t__darwin_mcontext64 = struct {
	F__ccgo_align [0]uint64
	F__es         t__darwin_arm_exception_state64
	F__ss         t__darwin_arm_thread_state64
	F__ns         t__darwin_arm_neon_state64
}
