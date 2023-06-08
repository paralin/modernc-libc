// Code generated by 'ccgo fts/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -nostdinc -nostdlib -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o fts/fts_linux_amd64.go -pkgname fts -Imusl-fts/ -Imusl/arch/x86_64 -Imusl/arch/generic -Imusl/obj/src/internal -Imusl/src/include -Imusl/src/internal -Imusl/obj/include -Imusl/include', DO NOT EDIT.

package fts

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	BIG_ENDIAN                                               = 4321                 // endian.h:14:1:
	BYTE_ORDER                                               = 1234                 // endian.h:17:1:
	FD_SETSIZE                                               = 1024                 // select.h:18:1:
	FEATURES_H                                               = 0                    // features.h:2:1:
	FTS_AGAIN                                                = 1                    // fts.h:136:1:
	FTS_COMFOLLOW                                            = 0x001                // fts.h:79:1:
	FTS_D                                                    = 1                    // fts.h:115:1:
	FTS_DC                                                   = 2                    // fts.h:116:1:
	FTS_DEFAULT                                              = 3                    // fts.h:117:1:
	FTS_DNR                                                  = 4                    // fts.h:118:1:
	FTS_DONTCHDIR                                            = 0x01                 // fts.h:131:1:
	FTS_DOT                                                  = 5                    // fts.h:119:1:
	FTS_DP                                                   = 6                    // fts.h:120:1:
	FTS_ERR                                                  = 7                    // fts.h:121:1:
	FTS_F                                                    = 8                    // fts.h:122:1:
	FTS_FOLLOW                                               = 2                    // fts.h:137:1:
	FTS_INIT                                                 = 9                    // fts.h:123:1:
	FTS_ISW                                                  = 0x04                 // fts.h:133:1:
	FTS_LOGICAL                                              = 0x002                // fts.h:80:1:
	FTS_NAMEONLY                                             = 0x100                // fts.h:89:1:
	FTS_NOCHDIR                                              = 0x004                // fts.h:81:1:
	FTS_NOINSTR                                              = 3                    // fts.h:138:1:
	FTS_NOSTAT                                               = 0x008                // fts.h:82:1:
	FTS_NS                                                   = 10                   // fts.h:124:1:
	FTS_NSOK                                                 = 11                   // fts.h:125:1:
	FTS_OPTIONMASK                                           = 0x0ff                // fts.h:87:1:
	FTS_PHYSICAL                                             = 0x010                // fts.h:83:1:
	FTS_ROOTLEVEL                                            = 0                    // fts.h:112:1:
	FTS_ROOTPARENTLEVEL                                      = -1                   // fts.h:111:1:
	FTS_SEEDOT                                               = 0x020                // fts.h:84:1:
	FTS_SKIP                                                 = 4                    // fts.h:139:1:
	FTS_SL                                                   = 12                   // fts.h:126:1:
	FTS_SLNONE                                               = 13                   // fts.h:127:1:
	FTS_STOP                                                 = 0x200                // fts.h:90:1:
	FTS_SYMFOLLOW                                            = 0x02                 // fts.h:132:1:
	FTS_W                                                    = 14                   // fts.h:128:1:
	FTS_WHITEOUT                                             = 0x080                // fts.h:86:1:
	FTS_XDEV                                                 = 0x040                // fts.h:85:1:
	INT16_MAX                                                = 32767                // stdint.h:44:1:
	INT16_MIN                                                = -32768               // stdint.h:39:1:
	INT32_MAX                                                = 2147483647           // stdint.h:45:1:
	INT32_MIN                                                = -2147483648          // stdint.h:40:1:
	INT64_MAX                                                = 9223372036854775807  // stdint.h:46:1:
	INT64_MIN                                                = -9223372036854775808 // stdint.h:41:1:
	INT8_MAX                                                 = 127                  // stdint.h:43:1:
	INT8_MIN                                                 = -128                 // stdint.h:38:1:
	INTMAX_MAX                                               = 9223372036854775807  // stdint.h:78:1:
	INTMAX_MIN                                               = -9223372036854775808 // stdint.h:77:1:
	INTPTR_MAX                                               = 9223372036854775807  // stdint.h:16:1:
	INTPTR_MIN                                               = -9223372036854775808 // stdint.h:15:1:
	INT_FAST16_MAX                                           = 2147483647           // stdint.h:9:1:
	INT_FAST16_MIN                                           = -2147483648          // stdint.h:6:1:
	INT_FAST32_MAX                                           = 2147483647           // stdint.h:10:1:
	INT_FAST32_MIN                                           = -2147483648          // stdint.h:7:1:
	INT_FAST64_MAX                                           = 9223372036854775807  // stdint.h:62:1:
	INT_FAST64_MIN                                           = -9223372036854775808 // stdint.h:54:1:
	INT_FAST8_MAX                                            = 127                  // stdint.h:61:1:
	INT_FAST8_MIN                                            = -128                 // stdint.h:53:1:
	INT_LEAST16_MAX                                          = 32767                // stdint.h:65:1:
	INT_LEAST16_MIN                                          = -32768               // stdint.h:57:1:
	INT_LEAST32_MAX                                          = 2147483647           // stdint.h:66:1:
	INT_LEAST32_MIN                                          = -2147483648          // stdint.h:58:1:
	INT_LEAST64_MAX                                          = 9223372036854775807  // stdint.h:67:1:
	INT_LEAST64_MIN                                          = -9223372036854775808 // stdint.h:59:1:
	INT_LEAST8_MAX                                           = 127                  // stdint.h:64:1:
	INT_LEAST8_MIN                                           = -128                 // stdint.h:56:1:
	LITTLE_ENDIAN                                            = 1234                 // endian.h:15:1:
	LOCAL_GOOGLE_HOME_YAIRH_GIT_MODERNC_LIBC_MUSL_FTS_FTS_H_ = 0                    // fts.h:2:1:
	PDP_ENDIAN                                               = 3412                 // endian.h:16:1:
	PTRDIFF_MAX                                              = 9223372036854775807  // stdint.h:19:1:
	PTRDIFF_MIN                                              = -9223372036854775808 // stdint.h:18:1:
	SIG_ATOMIC_MAX                                           = 2147483647           // stdint.h:93:1:
	SIG_ATOMIC_MIN                                           = -2147483648          // stdint.h:92:1:
	SIZE_MAX                                                 = 18446744073709551615 // stdint.h:20:1:
	UINT16_MAX                                               = 65535                // stdint.h:49:1:
	UINT32_MAX                                               = 4294967295           // stdint.h:50:1:
	UINT64_MAX                                               = 18446744073709551615 // stdint.h:51:1:
	UINT8_MAX                                                = 255                  // stdint.h:48:1:
	UINTMAX_MAX                                              = 18446744073709551615 // stdint.h:79:1:
	UINTPTR_MAX                                              = 18446744073709551615 // stdint.h:17:1:
	UINT_FAST16_MAX                                          = 4294967295           // stdint.h:12:1:
	UINT_FAST32_MAX                                          = 4294967295           // stdint.h:13:1:
	UINT_FAST64_MAX                                          = 18446744073709551615 // stdint.h:70:1:
	UINT_FAST8_MAX                                           = 255                  // stdint.h:69:1:
	UINT_LEAST16_MAX                                         = 65535                // stdint.h:73:1:
	UINT_LEAST32_MAX                                         = 4294967295           // stdint.h:74:1:
	UINT_LEAST64_MAX                                         = 18446744073709551615 // stdint.h:75:1:
	UINT_LEAST8_MAX                                          = 255                  // stdint.h:72:1:
	WCHAR_MAX                                                = 2147483647           // stdint.h:88:1:
	WCHAR_MIN                                                = -2147483648          // stdint.h:89:1:
	WINT_MAX                                                 = 4294967295           // stdint.h:82:1:
	WINT_MIN                                                 = 0                    // stdint.h:81:1:
	X_BSD_SOURCE                                             = 1                    // features.h:15:1:
	X_ENDIAN_H                                               = 0                    // endian.h:2:1:
	X_FEATURES_H                                             = 0                    // features.h:2:1:
	X_FILE_OFFSET_BITS                                       = 64                   // <builtin>:25:1:
	X_FTS_H_                                                 = 0                    // fts.h:40:1:
	X_LP64                                                   = 1                    // <predefined>:312:1:
	X_STDC_PREDEF_H                                          = 1                    // <predefined>:174:1:
	X_STDINT_H                                               = 0                    // stdint.h:2:1:
	X_SYS_SELECT_H                                           = 0                    // select.h:2:1:
	X_SYS_TYPES_H                                            = 0                    // types.h:2:1:
	X_XOPEN_SOURCE                                           = 700                  // features.h:16:1:
	Linux                                                    = 1                    // <predefined>:255:1:
	Unix                                                     = 1                    // <predefined>:191:1:
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

//	$NetBSD: fts.h,v 1.19 2009/08/16 19:33:38 christos Exp $

// Copyright (c) 1989, 1993
//	The Regents of the University of California.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)fts.h	8.3 (Berkeley) 8/14/94

type Uintptr_t = uint64 /* alltypes.h:55:24 */

type Intptr_t = int64 /* alltypes.h:70:15 */

type Int8_t = int8 /* alltypes.h:96:25 */

type Int16_t = int16 /* alltypes.h:101:25 */

type Int32_t = int32 /* alltypes.h:106:25 */

type Int64_t = int64 /* alltypes.h:111:25 */

type Intmax_t = int64 /* alltypes.h:116:25 */

type Uint8_t = uint8 /* alltypes.h:121:25 */

type Uint16_t = uint16 /* alltypes.h:126:25 */

type Uint32_t = uint32 /* alltypes.h:131:25 */

type Uint64_t = uint64 /* alltypes.h:136:25 */

type Uintmax_t = uint64 /* alltypes.h:146:25 */

type Int_fast8_t = Int8_t   /* stdint.h:22:16 */
type Int_fast64_t = Int64_t /* stdint.h:23:17 */

type Int_least8_t = Int8_t   /* stdint.h:25:17 */
type Int_least16_t = Int16_t /* stdint.h:26:17 */
type Int_least32_t = Int32_t /* stdint.h:27:17 */
type Int_least64_t = Int64_t /* stdint.h:28:17 */

type Uint_fast8_t = Uint8_t   /* stdint.h:30:17 */
type Uint_fast64_t = Uint64_t /* stdint.h:31:18 */

type Uint_least8_t = Uint8_t   /* stdint.h:33:18 */
type Uint_least16_t = Uint16_t /* stdint.h:34:18 */
type Uint_least32_t = Uint32_t /* stdint.h:35:18 */
type Uint_least64_t = Uint64_t /* stdint.h:36:18 */

type Int_fast16_t = Int32_t   /* stdint.h:1:17 */
type Int_fast32_t = Int32_t   /* stdint.h:2:17 */
type Uint_fast16_t = Uint32_t /* stdint.h:3:18 */
type Uint_fast32_t = Uint32_t /* stdint.h:4:18 */

type Ssize_t = int64 /* alltypes.h:65:15 */

type Register_t = int64 /* alltypes.h:80:14 */

type Time_t = int64 /* alltypes.h:85:16 */

type Suseconds_t = int64 /* alltypes.h:90:16 */

type U_int64_t = uint64 /* alltypes.h:141:25 */

type Mode_t = uint32 /* alltypes.h:152:18 */

type Nlink_t = uint64 /* alltypes.h:157:23 */

type Off_t = int64 /* alltypes.h:162:16 */

type Ino_t = uint64 /* alltypes.h:167:25 */

type Dev_t = uint64 /* alltypes.h:172:25 */

type Blksize_t = int64 /* alltypes.h:177:14 */

type Blkcnt_t = int64 /* alltypes.h:182:16 */

type Fsblkcnt_t = uint64 /* alltypes.h:187:25 */

type Fsfilcnt_t = uint64 /* alltypes.h:192:25 */

type Timer_t = uintptr /* alltypes.h:209:14 */

type Clockid_t = int32 /* alltypes.h:214:13 */

type Clock_t = int64 /* alltypes.h:219:14 */

type Pid_t = int32 /* alltypes.h:235:13 */

type Id_t = uint32 /* alltypes.h:240:18 */

type Uid_t = uint32 /* alltypes.h:245:18 */

type Gid_t = uint32 /* alltypes.h:250:18 */

type Key_t = int32 /* alltypes.h:255:13 */

type Useconds_t = uint32 /* alltypes.h:260:18 */

type Pthread_t = uintptr /* alltypes.h:273:26 */

type Pthread_once_t = int32 /* alltypes.h:279:13 */

type Pthread_key_t = uint32 /* alltypes.h:284:18 */

type Pthread_spinlock_t = int32 /* alltypes.h:289:13 */

type Pthread_mutexattr_t = struct{ F__attr uint32 } /* alltypes.h:294:37 */

type Pthread_condattr_t = struct{ F__attr uint32 } /* alltypes.h:299:37 */

type Pthread_barrierattr_t = struct{ F__attr uint32 } /* alltypes.h:304:37 */

type Pthread_rwlockattr_t = struct{ F__attr [2]uint32 } /* alltypes.h:309:40 */

type Pthread_attr_t = struct {
	F__u struct {
		F__ccgo_pad1 [0]uint64
		F__i         [14]int32
	}
} /* alltypes.h:372:147 */

type Pthread_mutex_t = struct {
	F__u struct {
		F__ccgo_pad1 [0]uint64
		F__i         [10]int32
	}
} /* alltypes.h:377:157 */

type Pthread_cond_t = struct {
	F__u struct {
		F__ccgo_pad1 [0]uint64
		F__i         [12]int32
	}
} /* alltypes.h:387:112 */

type Pthread_rwlock_t = struct {
	F__u struct {
		F__ccgo_pad1 [0]uint64
		F__i         [14]int32
	}
} /* alltypes.h:397:139 */

type Pthread_barrier_t = struct {
	F__u struct {
		F__ccgo_pad1 [0]uint64
		F__i         [8]int32
	}
} /* alltypes.h:402:137 */

type U_int8_t = uint8   /* types.h:60:23 */
type U_int16_t = uint16 /* types.h:61:24 */
type U_int32_t = uint32 /* types.h:62:18 */
type Caddr_t = uintptr  /* types.h:63:14 */
type U_char = uint8     /* types.h:64:23 */
type U_short = uint16   /* types.h:65:24 */
type Ushort = uint16    /* types.h:65:33 */
type U_int = uint32     /* types.h:66:18 */
type Uint = uint32      /* types.h:66:25 */
type U_long = uint64    /* types.h:67:23 */
type Ulong = uint64     /* types.h:67:31 */
type Quad_t = int64     /* types.h:68:19 */
type U_quad_t = uint64  /* types.h:69:28 */

type Timeval = struct {
	Ftv_sec  Time_t
	Ftv_usec Suseconds_t
} /* alltypes.h:224:1 */

type Timespec = struct {
	Ftv_sec  Time_t
	Ftv_nsec int64
} /* alltypes.h:229:1 */

type X__sigset_t = struct{ F__bits [16]uint64 } /* alltypes.h:349:9 */

type Sigset_t = X__sigset_t /* alltypes.h:349:71 */

type Fd_mask = uint64 /* select.h:20:23 */

type Fd_set = struct{ Ffds_bits [16]uint64 } /* select.h:24:3 */

type X_ftsent = struct {
	Ffts_cycle   uintptr
	Ffts_parent  uintptr
	Ffts_link    uintptr
	Ffts_number  Int64_t
	Ffts_pointer uintptr
	Ffts_accpath uintptr
	Ffts_path    uintptr
	Ffts_errno   int32
	Ffts_symfd   int32
	Ffts_pathlen uint32
	Ffts_namelen uint32
	Ffts_ino     Ino_t
	Ffts_dev     Dev_t
	Ffts_nlink   Nlink_t
	Ffts_level   int32
	Ffts_info    uint16
	Ffts_flags   uint16
	Ffts_instr   uint16
	F__ccgo_pad1 [6]byte
	Ffts_statp   uintptr
	Ffts_name    [1]int8
	F__ccgo_pad2 [7]byte
} /* fts.h:68:2 */

type FTS = struct {
	Ffts_cur     uintptr
	Ffts_child   uintptr
	Ffts_array   uintptr
	Ffts_dev     Dev_t
	Ffts_path    uintptr
	Ffts_rfd     int32
	Ffts_pathlen uint32
	Ffts_nitems  uint32
	F__ccgo_pad1 [4]byte
	Ffts_compar  uintptr
	Ffts_options int32
	F__ccgo_pad2 [4]byte
} /* fts.h:92:3 */

type FTSENT = X_ftsent /* fts.h:144:3 */

var _ int8 /* gen.c:2:13: */
