// Code generated for darwin/arm64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o vec.go --package-name libsqlite_vec dist/libsqlite_vec0.a -lsqlite3', DO NOT EDIT.

//go:build darwin && arm64

package vec

type Trusage_info_current = struct {
	Fri_uuid                          [16]Tuint8_t
	Fri_user_time                     Tuint64_t
	Fri_system_time                   Tuint64_t
	Fri_pkg_idle_wkups                Tuint64_t
	Fri_interrupt_wkups               Tuint64_t
	Fri_pageins                       Tuint64_t
	Fri_wired_size                    Tuint64_t
	Fri_resident_size                 Tuint64_t
	Fri_phys_footprint                Tuint64_t
	Fri_proc_start_abstime            Tuint64_t
	Fri_proc_exit_abstime             Tuint64_t
	Fri_child_user_time               Tuint64_t
	Fri_child_system_time             Tuint64_t
	Fri_child_pkg_idle_wkups          Tuint64_t
	Fri_child_interrupt_wkups         Tuint64_t
	Fri_child_pageins                 Tuint64_t
	Fri_child_elapsed_abstime         Tuint64_t
	Fri_diskio_bytesread              Tuint64_t
	Fri_diskio_byteswritten           Tuint64_t
	Fri_cpu_time_qos_default          Tuint64_t
	Fri_cpu_time_qos_maintenance      Tuint64_t
	Fri_cpu_time_qos_background       Tuint64_t
	Fri_cpu_time_qos_utility          Tuint64_t
	Fri_cpu_time_qos_legacy           Tuint64_t
	Fri_cpu_time_qos_user_initiated   Tuint64_t
	Fri_cpu_time_qos_user_interactive Tuint64_t
	Fri_billed_system_time            Tuint64_t
	Fri_serviced_system_time          Tuint64_t
	Fri_logical_writes                Tuint64_t
	Fri_lifetime_max_phys_footprint   Tuint64_t
	Fri_instructions                  Tuint64_t
	Fri_cycles                        Tuint64_t
	Fri_billed_energy                 Tuint64_t
	Fri_serviced_energy               Tuint64_t
	Fri_interval_max_phys_footprint   Tuint64_t
	Fri_runnable_time                 Tuint64_t
	Fri_flags                         Tuint64_t
	Fri_user_ptime                    Tuint64_t
	Fri_system_ptime                  Tuint64_t
	Fri_pinstructions                 Tuint64_t
	Fri_pcycles                       Tuint64_t
	Fri_energy_nj                     Tuint64_t
	Fri_penergy_nj                    Tuint64_t
	Fri_secure_time_in_system         Tuint64_t
	Fri_secure_ptime_in_system        Tuint64_t
	Fri_neural_footprint              Tuint64_t
	Fri_lifetime_max_neural_footprint Tuint64_t
	Fri_interval_max_neural_footprint Tuint64_t
	Fri_conclave_footprint            Tuint64_t
	Fri_page_wait_time_mach           Tuint64_t
	Fri_page_cache_hits               Tuint64_t
	Fri_reserved                      [6]Tuint64_t
}

type Trusage_info_v6 = struct {
	Fri_uuid                          [16]Tuint8_t
	Fri_user_time                     Tuint64_t
	Fri_system_time                   Tuint64_t
	Fri_pkg_idle_wkups                Tuint64_t
	Fri_interrupt_wkups               Tuint64_t
	Fri_pageins                       Tuint64_t
	Fri_wired_size                    Tuint64_t
	Fri_resident_size                 Tuint64_t
	Fri_phys_footprint                Tuint64_t
	Fri_proc_start_abstime            Tuint64_t
	Fri_proc_exit_abstime             Tuint64_t
	Fri_child_user_time               Tuint64_t
	Fri_child_system_time             Tuint64_t
	Fri_child_pkg_idle_wkups          Tuint64_t
	Fri_child_interrupt_wkups         Tuint64_t
	Fri_child_pageins                 Tuint64_t
	Fri_child_elapsed_abstime         Tuint64_t
	Fri_diskio_bytesread              Tuint64_t
	Fri_diskio_byteswritten           Tuint64_t
	Fri_cpu_time_qos_default          Tuint64_t
	Fri_cpu_time_qos_maintenance      Tuint64_t
	Fri_cpu_time_qos_background       Tuint64_t
	Fri_cpu_time_qos_utility          Tuint64_t
	Fri_cpu_time_qos_legacy           Tuint64_t
	Fri_cpu_time_qos_user_initiated   Tuint64_t
	Fri_cpu_time_qos_user_interactive Tuint64_t
	Fri_billed_system_time            Tuint64_t
	Fri_serviced_system_time          Tuint64_t
	Fri_logical_writes                Tuint64_t
	Fri_lifetime_max_phys_footprint   Tuint64_t
	Fri_instructions                  Tuint64_t
	Fri_cycles                        Tuint64_t
	Fri_billed_energy                 Tuint64_t
	Fri_serviced_energy               Tuint64_t
	Fri_interval_max_phys_footprint   Tuint64_t
	Fri_runnable_time                 Tuint64_t
	Fri_flags                         Tuint64_t
	Fri_user_ptime                    Tuint64_t
	Fri_system_ptime                  Tuint64_t
	Fri_pinstructions                 Tuint64_t
	Fri_pcycles                       Tuint64_t
	Fri_energy_nj                     Tuint64_t
	Fri_penergy_nj                    Tuint64_t
	Fri_secure_time_in_system         Tuint64_t
	Fri_secure_ptime_in_system        Tuint64_t
	Fri_neural_footprint              Tuint64_t
	Fri_lifetime_max_neural_footprint Tuint64_t
	Fri_interval_max_neural_footprint Tuint64_t
	Fri_conclave_footprint            Tuint64_t
	Fri_page_wait_time_mach           Tuint64_t
	Fri_page_cache_hits               Tuint64_t
	Fri_reserved                      [6]Tuint64_t
}

const m_IOPOL_TYPE_VFS_CONTENT_PROTECTION = 6

const m_IOPOL_TYPE_VFS_IGNORE_CONTENT_PROTECTION = "IOPOL_TYPE_VFS_CONTENT_PROTECTION"

const m_IOPOL_VFS_CONTENT_PROTECTION_STRICT = 2

const m_MAC_OS_VERSION_26_6 = "__MAC_26_6"

const m_MAC_OS_VERSION_27_0 = "__MAC_27_0"

const m_PTHREAD_STACK_MIN = 16384

const m_TARGET_OS_ARROW = 1

const m_TARGET_OS_FIRMWARE = 0

const m_TARGET_OS_KERNELKIT = 0

const m__ARM_SIGNAL_ = 1

const m__DARWIN_FEATURE_ONLY_64_BIT_INODE = 1

const m__DARWIN_FEATURE_ONLY_VERS_1050 = 1

const m___AARCH64_SIMD__ = 1

const m___ARM64_ARCH_8__ = 1

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

const m___BRIDGEOS_10_5 = 100500

const m___BRIDGEOS_10_6 = 100600

const m___DARWIN_ONLY_64_BIT_INO_T = 1

const m___DARWIN_ONLY_VERS_1050 = 1

const m___DARWIN_OPAQUE_ARM_THREAD_STATE64 = 0

const m___DRIVERKIT_25_6 = 250600

const m___DRIVERKIT_27_0 = 270000

const m___ENVIRONMENT_MAC_OS_X_VERSION_MIN_REQUIRED__ = 270000

const m___ENVIRONMENT_OS_VERSION_MIN_REQUIRED__ = 270000

const m___IPHONE_26_6 = 260600

const m___IPHONE_27_0 = 270000

const m___MAC_26_6 = 260600

const m___MAC_27_0 = 270000

const m___OBJC_BOOL_IS_BOOL = 1

const m___TVOS_26_6 = 260600

const m___TVOS_27_0 = 270000

const m___VERSION__ = "Apple LLVM 21.0.0 (clang-2100.3.34.2)"

const m___VISIONOS_26_6 = 260600

const m___VISIONOS_27_0 = 270000

const m___WATCHOS_26_6 = 260600

const m___WATCHOS_27_0 = 270000

const m___apple_build_version__ = 21000334

const m___arm64 = 1

const m___arm64__ = 1

const m___clang_version__ = "21.0.0 (clang-2100.3.34.2)"

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
