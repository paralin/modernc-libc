// Code generated by 'ccgo locale/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o locale/locale_openbsd_amd64.go -pkgname locale', DO NOT EDIT.

package locale

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
	LC_ALL              = 0
	LC_ALL_MASK         = 126
	LC_COLLATE          = 1
	LC_COLLATE_MASK     = 2
	LC_CTYPE            = 2
	LC_CTYPE_MASK       = 4
	LC_MESSAGES         = 6
	LC_MESSAGES_MASK    = 64
	LC_MONETARY         = 3
	LC_MONETARY_MASK    = 8
	LC_NUMERIC          = 4
	LC_NUMERIC_MASK     = 16
	LC_TIME             = 5
	LC_TIME_MASK        = 32
	X_FILE_OFFSET_BITS  = 64
	X_LC_LAST           = 7
	X_LOCALE_H_         = 0
	X_LOCALE_T_DEFINED_ = 0
	X_LP64              = 1
	X_MACHINE_CDEFS_H_  = 0
	X_RET_PROTECTOR     = 1
	X_SYS_CDEFS_H_      = 0
	Unix                = 1
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

//	$OpenBSD: locale.h,v 1.11 2017/09/05 03:16:13 schwarze Exp $
//	$NetBSD: locale.h,v 1.6 1994/10/26 00:56:02 cgd Exp $

// Copyright (c) 1991 The Regents of the University of California.
// All rights reserved.
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
//	@(#)locale.h	5.2 (Berkeley) 2/24/91

//	$OpenBSD: _null.h,v 1.2 2016/09/09 22:07:58 millert Exp $

// Written by Todd C. Miller, September 9, 2016
// Public domain.

type Lconv = struct {
	Fdecimal_point      uintptr
	Fthousands_sep      uintptr
	Fgrouping           uintptr
	Fint_curr_symbol    uintptr
	Fcurrency_symbol    uintptr
	Fmon_decimal_point  uintptr
	Fmon_thousands_sep  uintptr
	Fmon_grouping       uintptr
	Fpositive_sign      uintptr
	Fnegative_sign      uintptr
	Fint_frac_digits    int8
	Ffrac_digits        int8
	Fp_cs_precedes      int8
	Fp_sep_by_space     int8
	Fn_cs_precedes      int8
	Fn_sep_by_space     int8
	Fp_sign_posn        int8
	Fn_sign_posn        int8
	Fint_p_cs_precedes  int8
	Fint_p_sep_by_space int8
	Fint_n_cs_precedes  int8
	Fint_n_sep_by_space int8
	Fint_p_sign_posn    int8
	Fint_n_sign_posn    int8
	F__ccgo_pad1        [2]byte
} /* locale.h:40:1 */

//	$OpenBSD: cdefs.h,v 1.43 2018/10/29 17:10:40 guenther Exp $
//	$NetBSD: cdefs.h,v 1.16 1996/04/03 20:46:39 christos Exp $

// Copyright (c) 1991, 1993
//	The Regents of the University of California.  All rights reserved.
//
// This code is derived from software contributed to Berkeley by
// Berkeley Software Design, Inc.
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
//	@(#)cdefs.h	8.7 (Berkeley) 1/21/94

//	$OpenBSD: cdefs.h,v 1.3 2013/03/28 17:30:45 martynas Exp $

// Written by J.T. Conklin <jtc@wimsey.com> 01/17/95.
// Public domain.

// Macro to test if we're using a specific version of gcc or later.

// The __CONCAT macro is used to concatenate parts of symbol names, e.g.
// with "#define OLD(foo) __CONCAT(old,foo)", OLD(foo) produces oldfoo.
// The __CONCAT macro is a bit tricky -- make sure you don't put spaces
// in between its arguments.  Do not use __CONCAT on double-quoted strings,
// such as those from the __STRING macro: to concatenate strings just put
// them next to each other.

// GCC1 and some versions of GCC2 declare dead (non-returning) and
// pure (no side effects) functions using "volatile" and "const";
// unfortunately, these then cause warnings under "-ansi -pedantic".
// GCC >= 2.5 uses the __attribute__((attrs)) style.  All of these
// work for GNU C++ (modulo a slight glitch in the C++ grammar in
// the distribution version of 2.5.5).

// __returns_twice makes the compiler not assume the function
// only returns once.  This affects registerisation of variables:
// even local variables need to be in memory across such a call.
// Example: setjmp()

// __only_inline makes the compiler only use this function definition
// for inlining; references that can't be inlined will be left as
// external references instead of generating a local copy.  The
// matching library should include a simple extern definition for
// the function to handle those references.  c.f. ctype.h

// GNU C version 2.96 adds explicit branch prediction so that
// the CPU back-end can hint the processor and also so that
// code blocks can be reordered such that the predicted path
// sees a more linear flow, thus improving cache behavior, etc.
//
// The following two macros provide us with a way to utilize this
// compiler feature.  Use __predict_true() if you expect the expression
// to evaluate to true, and __predict_false() if you expect the
// expression to evaluate to false.
//
// A few notes about usage:
//
//	* Generally, __predict_false() error condition checks (unless
//	  you have some _strong_ reason to do otherwise, in which case
//	  document it), and/or __predict_true() `no-error' condition
//	  checks, assuming you want to optimize for the no-error case.
//
//	* Other than that, if you don't know the likelihood of a test
//	  succeeding from empirical or other `hard' evidence, don't
//	  make predictions.
//
//	* These are meant to be used in places that are run `a lot'.
//	  It is wasteful to make predictions in code that is run
//	  seldomly (e.g. at subsystem initialization time) as the
//	  basic block reordering that this affects can often generate
//	  larger code.

// Delete pseudo-keywords wherever they are not available or needed.

// The __packed macro indicates that a variable or structure members
// should have the smallest possible alignment, despite any host CPU
// alignment requirements.
//
// The __aligned(x) macro specifies the minimum alignment of a
// variable or structure.
//
// These macros together are useful for describing the layout and
// alignment of messages exchanged with hardware or other systems.

// "The nice thing about standards is that there are so many to choose from."
// There are a number of "feature test macros" specified by (different)
// standards that determine which interfaces and types the header files
// should expose.
//
// Because of inconsistencies in these macros, we define our own
// set in the private name space that end in _VISIBLE.  These are
// always defined and so headers can test their values easily.
// Things can get tricky when multiple feature macros are defined.
// We try to take the union of all the features requested.
//
// The following macros are guaranteed to have a value after cdefs.h
// has been included:
//	__POSIX_VISIBLE
//	__XPG_VISIBLE
//	__ISO_C_VISIBLE
//	__BSD_VISIBLE

// X/Open Portability Guides and Single Unix Specifications.
// _XOPEN_SOURCE				XPG3
// _XOPEN_SOURCE && _XOPEN_VERSION = 4		XPG4
// _XOPEN_SOURCE && _XOPEN_SOURCE_EXTENDED = 1	XPG4v2
// _XOPEN_SOURCE == 500				XPG5
// _XOPEN_SOURCE == 520				XPG5v2
// _XOPEN_SOURCE == 600				POSIX 1003.1-2001 with XSI
// _XOPEN_SOURCE == 700				POSIX 1003.1-2008 with XSI
//
// The XPG spec implies a specific value for _POSIX_C_SOURCE.

// POSIX macros, these checks must follow the XOPEN ones above.
//
// _POSIX_SOURCE == 1		1003.1-1988 (superseded by _POSIX_C_SOURCE)
// _POSIX_C_SOURCE == 1		1003.1-1990
// _POSIX_C_SOURCE == 2		1003.2-1992
// _POSIX_C_SOURCE == 199309L	1003.1b-1993
// _POSIX_C_SOURCE == 199506L   1003.1c-1995, 1003.1i-1995,
//				and the omnibus ISO/IEC 9945-1:1996
// _POSIX_C_SOURCE == 200112L   1003.1-2001
// _POSIX_C_SOURCE == 200809L   1003.1-2008
//
// The POSIX spec implies a specific value for __ISO_C_VISIBLE, though
// this may be overridden by the _ISOC99_SOURCE macro later.

// _ANSI_SOURCE means to expose ANSI C89 interfaces only.
// If the user defines it in addition to one of the POSIX or XOPEN
// macros, assume the POSIX/XOPEN macro(s) should take precedence.

// _ISOC99_SOURCE, _ISOC11_SOURCE, __STDC_VERSION__, and __cplusplus
// override any of the other macros since they are non-exclusive.

// Finally deal with BSD-specific interfaces that are not covered
// by any standards.  We expose these when none of the POSIX or XPG
// macros is defined or if the user explicitly asks for them.

// Default values.

type Locale_t = uintptr /* locale.h:83:14 */

var _ int8 /* gen.c:2:13: */