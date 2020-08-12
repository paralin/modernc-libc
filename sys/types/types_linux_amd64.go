// Code generated by 'ccgo /tmp/go-generate-206723672/x.c -ccgo-crt-import-path  -ccgo-export-defines  -ccgo-export-enums  -ccgo-export-externs X -ccgo-export-fields F -ccgo-export-structs  -ccgo-export-typedefs  -ccgo-pkgname types -o sys/types/types_linux_amd64.go', DO NOT EDIT.

package types

import (
	"math"
	"reflect"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ unsafe.Pointer

const (
	BIG_ENDIAN                   = 4321
	BYTE_ORDER                   = 1234
	FD_SETSIZE                   = 1024
	LITTLE_ENDIAN                = 1234
	PDP_ENDIAN                   = 3412
	X_ATFILE_SOURCE              = 1
	X_BITS_BYTESWAP_H            = 1
	X_BITS_PTHREADTYPES_ARCH_H   = 1
	X_BITS_PTHREADTYPES_COMMON_H = 1
	X_BITS_STDINT_INTN_H         = 1
	X_BITS_TYPESIZES_H           = 1
	X_BITS_TYPES_H               = 1
	X_BITS_UINTN_IDENTITY_H      = 1
	X_BSD_SIZE_T_                = 0
	X_BSD_SIZE_T_DEFINED_        = 0
	X_DEFAULT_SOURCE             = 1
	X_ENDIAN_H                   = 1
	X_FEATURES_H                 = 1
	X_GCC_SIZE_T                 = 0
	X_LP64                       = 1
	X_POSIX_C_SOURCE             = 200809
	X_POSIX_SOURCE               = 1
	X_SIZET_                     = 0
	X_SIZE_T                     = 0
	X_SIZE_T_                    = 0
	X_SIZE_T_DECLARED            = 0
	X_SIZE_T_DEFINED             = 0
	X_SIZE_T_DEFINED_            = 0
	X_STDC_PREDEF_H              = 1
	X_STRUCT_TIMESPEC            = 1
	X_SYS_CDEFS_H                = 1
	X_SYS_SELECT_H               = 1
	X_SYS_SIZE_T_H               = 0
	X_SYS_TYPES_H                = 1
	X_THREAD_SHARED_TYPES_H      = 1
	X_T_SIZE                     = 0
	X_T_SIZE_                    = 0
	Linux                        = 1
	Unix                         = 1
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type U_char = uint8                     /* types.h:33:18 */
type U_short = uint16                   /* types.h:34:19 */
type U_int = uint32                     /* types.h:35:17 */
type U_long = uint64                    /* types.h:36:18 */
type Quad_t = int64                     /* types.h:37:18 */
type U_quad_t = uint64                  /* types.h:38:20 */
type Fsid_t = struct{ F__val [2]int32 } /* types.h:39:18 */
type Loff_t = int64                     /* types.h:42:18 */

type Ino_t = uint64 /* types.h:47:17 */

type Dev_t = uint64 /* types.h:59:17 */

type Gid_t = uint32 /* types.h:64:17 */

type Mode_t = uint32 /* types.h:69:18 */

type Nlink_t = uint64 /* types.h:74:19 */

type Uid_t = uint32 /* types.h:79:17 */

type Off_t = int64 /* types.h:85:17 */

type Pid_t = int32 /* types.h:97:17 */

type Id_t = uint32 /* types.h:103:16 */

type Ssize_t = int64 /* types.h:108:19 */

type Daddr_t = int32   /* types.h:114:19 */
type Caddr_t = uintptr /* types.h:115:19 */

type Key_t = int32 /* types.h:121:17 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//   Copyright (C) 2002-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Returned by `clock'.
type Clock_t = int64 /* clock_t.h:7:19 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//   Copyright (C) 2002-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Clock ID used in clock and timer functions.
type Clockid_t = int32 /* clockid_t.h:7:21 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//   Copyright (C) 2002-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Returned by `time'.
type Time_t = int64 /* time_t.h:7:18 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//   Copyright (C) 2002-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Timer ID returned by `timer_create'.
type Timer_t = uintptr /* timer_t.h:7:19 */

// Wide character type.
//   Locale-writers should change this as necessary to
//   be big enough to hold unique values not between 0 and 127,
//   and not (wchar_t) -1, for each defined multibyte character.

// Define this type if we are doing the whole job,
//   or if we want this type in particular.

//  In 4.3bsd-net2, leave these undefined to indicate that size_t, etc.
//    are already defined.
//  BSD/OS 3.1 and FreeBSD [23].x require the MACHINE_ANSI_H check here.
//  NetBSD 5 requires the I386_ANSI_H and X86_64_ANSI_H checks here.

// A null pointer constant.

// Old compatibility names for C types.
type Ulong = uint64  /* types.h:148:27 */
type Ushort = uint16 /* types.h:149:28 */
type Uint = uint32   /* types.h:150:22 */

// These size-specific names are used by some of the inet code.

// Define intN_t types.
//   Copyright (C) 2017-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//   Copyright (C) 2002-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

type Int8_t = int8   /* stdint-intn.h:24:18 */
type Int16_t = int16 /* stdint-intn.h:25:19 */
type Int32_t = int32 /* stdint-intn.h:26:19 */
type Int64_t = int64 /* stdint-intn.h:27:19 */

// These were defined by ISO C without the first `_'.
type U_int8_t = uint8   /* types.h:160:23 */
type U_int16_t = uint16 /* types.h:161:28 */
type U_int32_t = uint32 /* types.h:162:22 */
type U_int64_t = uint64 /* types.h:164:27 */

type Register_t = int32 /* types.h:169:13 */

// A set of signals to be blocked, unblocked, or waited for.
type Sigset_t = struct{ F__val [16]uint64 } /* sigset_t.h:7:20 */

// Get definition of timer specification structures.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//   Copyright (C) 2002-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// A time value that is accurate to the nearest
//   microsecond but also has a range of years.
type Timeval = struct {
	Ftv_sec  int64
	Ftv_usec int64
}

// NB: Include guard matches what <linux/time.h> uses.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//   Copyright (C) 2002-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// POSIX.1b structure for a time value.  This is like a `struct timeval' but
//   has nanoseconds instead of microseconds.
type Timespec = struct {
	Ftv_sec  int64
	Ftv_nsec int64
}

type Suseconds_t = int64 /* select.h:43:23 */

// Some versions of <linux/posix_types.h> define this macros.
// It's easier to assume 8-bit bytes than to get CHAR_BIT.

// fd_set for select and pselect.
type Fd_set = struct{ F__fds_bits [16]int64 } /* select.h:70:5 */

// Maximum number of file descriptors in `fd_set'.

// Sometimes the fd_set member is assumed to have this type.
type Fd_mask = int64 /* select.h:77:19 */

// Define some inlines helping to catch common problems.

type Blksize_t = int64 /* types.h:202:21 */

// Types from the Large File Support interface.
type Blkcnt_t = int64    /* types.h:209:20 */ // Type to count number of disk blocks.
type Fsblkcnt_t = uint64 /* types.h:213:22 */ // Type to count file system blocks.
type Fsfilcnt_t = uint64 /* types.h:217:22 */ // Type to count file system inodes.

// Now add the thread types.
// Declaration of common pthread types for all architectures.
//   Copyright (C) 2017-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// For internal mutex and condition variable definitions.
// Common threading primitives definitions for both POSIX and C11.
//   Copyright (C) 2017-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Arch-specific definitions.  Each architecture must define the following
//   macros to define the expected sizes of pthread data types:
//
//   __SIZEOF_PTHREAD_ATTR_T        - size of pthread_attr_t.
//   __SIZEOF_PTHREAD_MUTEX_T       - size of pthread_mutex_t.
//   __SIZEOF_PTHREAD_MUTEXATTR_T   - size of pthread_mutexattr_t.
//   __SIZEOF_PTHREAD_COND_T        - size of pthread_cond_t.
//   __SIZEOF_PTHREAD_CONDATTR_T    - size of pthread_condattr_t.
//   __SIZEOF_PTHREAD_RWLOCK_T      - size of pthread_rwlock_t.
//   __SIZEOF_PTHREAD_RWLOCKATTR_T  - size of pthread_rwlockattr_t.
//   __SIZEOF_PTHREAD_BARRIER_T     - size of pthread_barrier_t.
//   __SIZEOF_PTHREAD_BARRIERATTR_T - size of pthread_barrierattr_t.
//
//   Also, the following macros must be define for internal pthread_mutex_t
//   struct definitions (struct __pthread_mutex_s):
//
//   __PTHREAD_COMPAT_PADDING_MID   - any additional members after 'kind'
//				    and before '__spin' (for 64 bits) or
//				    '__nusers' (for 32 bits).
//   __PTHREAD_COMPAT_PADDING_END   - any additional members at the end of
//				    the internal structure.
//   __PTHREAD_MUTEX_LOCK_ELISION   - 1 if the architecture supports lock
//				    elision or 0 otherwise.
//   __PTHREAD_MUTEX_NUSERS_AFTER_KIND - control where to put __nusers.  The
//				       preferred value for new architectures
//				       is 0.
//   __PTHREAD_MUTEX_USE_UNION      - control whether internal __spins and
//				    __list will be place inside a union for
//				    linuxthreads compatibility.
//				    The preferred value for new architectures
//				    is 0.
//
//   For a new port the preferred values for the required defines are:
//
//   #define __PTHREAD_COMPAT_PADDING_MID
//   #define __PTHREAD_COMPAT_PADDING_END
//   #define __PTHREAD_MUTEX_LOCK_ELISION         0
//   #define __PTHREAD_MUTEX_NUSERS_AFTER_KIND    0
//   #define __PTHREAD_MUTEX_USE_UNION            0
//
//   __PTHREAD_MUTEX_LOCK_ELISION can be set to 1 if the hardware plans to
//   eventually support lock elision using transactional memory.
//
//   The additional macro defines any constraint for the lock alignment
//   inside the thread structures:
//
//   __LOCK_ALIGNMENT - for internal lock/futex usage.
//
//   Same idea but for the once locking primitive:
//
//   __ONCE_ALIGNMENT - for pthread_once_t/once_flag definition.
//
//   And finally the internal pthread_rwlock_t (struct __pthread_rwlock_arch_t)
//   must be defined.
//
// Copyright (C) 2002-2018 Free Software Foundation, Inc.
//   This file is part of the GNU C Library.
//
//   The GNU C Library is free software; you can redistribute it and/or
//   modify it under the terms of the GNU Lesser General Public
//   License as published by the Free Software Foundation; either
//   version 2.1 of the License, or (at your option) any later version.
//
//   The GNU C Library is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   Lesser General Public License for more details.
//
//   You should have received a copy of the GNU Lesser General Public
//   License along with the GNU C Library; if not, see
//   <http://www.gnu.org/licenses/>.

// Determine the wordsize from the preprocessor defines.

// Both x86-64 and x32 use the 64-bit system call interface.

// Definitions for internal mutex struct.

type X__pthread_rwlock_arch_t = struct {
	F__readers       uint32
	F__writers       uint32
	F__wrphase_futex uint32
	F__writers_futex uint32
	F__pad3          uint32
	F__pad4          uint32
	F__cur_writer    int32
	F__shared        int32
	F__rwelision     int8
	F__pad1          [7]uint8
	F__pad2          uint64
	F__flags         uint32
	_                [4]byte
}

// Common definition of pthread_mutex_t.

type X__pthread_internal_list = struct {
	F__prev uintptr
	F__next uintptr
}

// Lock elision support.

type X__pthread_mutex_s = struct {
	F__lock    int32
	F__count   uint32
	F__owner   int32
	F__nusers  uint32
	F__kind    int32
	F__spins   int16
	F__elision int16
	F__list    struct {
		F__prev uintptr
		F__next uintptr
	}
}

// Common definition of pthread_cond_t.

type X__pthread_cond_s = struct {
	F__0            struct{ F__wseq uint64 }
	F__8            struct{ F__g1_start uint64 }
	F__g_refs       [2]uint32
	F__g_size       [2]uint32
	F__g1_orig_size uint32
	F__wrefs        uint32
	F__g_signals    [2]uint32
}

// Thread identifiers.  The structure of the attribute type is not
//   exposed on purpose.
type Pthread_t = uint64 /* pthreadtypes.h:27:27 */

// Data structures for mutex handling.  The structure of the attribute
//   type is not exposed on purpose.
type Pthread_mutexattr_t = struct {
	_       [0]uint32
	F__size [4]int8
} /* pthreadtypes.h:36:3 */

// Data structure for condition variable handling.  The structure of
//   the attribute type is not exposed on purpose.
type Pthread_condattr_t = struct {
	_       [0]uint32
	F__size [4]int8
} /* pthreadtypes.h:45:3 */

// Keys for thread-specific data
type Pthread_key_t = uint32 /* pthreadtypes.h:49:22 */

// Once-only execution
type Pthread_once_t = int32 /* pthreadtypes.h:53:30 */

type Pthread_attr_t1 = struct {
	_       [0]uint64
	F__size [56]int8
}

type Pthread_attr_t = Pthread_attr_t1 /* pthreadtypes.h:62:30 */

type Pthread_mutex_t = struct{ F__data X__pthread_mutex_s } /* pthreadtypes.h:72:3 */

type Pthread_cond_t = struct{ F__data X__pthread_cond_s } /* pthreadtypes.h:80:3 */

// Data structure for reader-writer lock variable handling.  The
//   structure of the attribute type is deliberately not exposed.
type Pthread_rwlock_t = struct{ F__data X__pthread_rwlock_arch_t } /* pthreadtypes.h:91:3 */

type Pthread_rwlockattr_t = struct {
	_       [0]uint64
	F__size [8]int8
} /* pthreadtypes.h:97:3 */

// POSIX spinlock data type.
type Pthread_spinlock_t = int32 /* pthreadtypes.h:103:22 */

// POSIX barriers data type.  The structure of the type is
//   deliberately not exposed.
type Pthread_barrier_t = struct {
	_       [0]uint64
	F__size [32]int8
} /* pthreadtypes.h:112:3 */

type Pthread_barrierattr_t = struct {
	_       [0]uint32
	F__size [4]int8
} /* pthreadtypes.h:118:3 */

var _ int8 /* x.c:2:13: */
