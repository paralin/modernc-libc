// Copyright 2020 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go.generate echo package libc > ccgo.go
//go:generate go run generate.go
//go:generate go fmt ./...

// Package libc provides run time support for ccgo generated programs and
// implements selected parts of the C standard library.
package libc // import "modernc.org/libc"

//TODO use O_RDONLY etc. from fcntl header

//TODO use t.Alloc/Free where appropriate

import (
	"fmt"
	"math"
	mbits "math/bits"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	gotime "time"
	"unsafe"

	"github.com/mattn/go-isatty"
	"modernc.org/libc/errno"
	"modernc.org/libc/sys/types"
	"modernc.org/libc/time"
	"modernc.org/libc/unistd"
	"modernc.org/mathutil"
	"modernc.org/memory"
)

var (
	allocMu   sync.Mutex
	allocator memory.Allocator
	isWindows bool
)

// Keep these outside of the var block otherwise go generate will miss them.
var Xenviron uintptr
var Xstdin = newFile(nil, unistd.STDIN_FILENO)
var Xstdout = newFile(nil, unistd.STDOUT_FILENO)
var Xstderr = newFile(nil, unistd.STDERR_FILENO)

func Environ() uintptr {
	return Xenviron
}

func EnvironP() uintptr {
	return uintptr(unsafe.Pointer(&Xenviron))
}

func X___errno_location(t *TLS) uintptr {
	return X__errno_location(t)
}

// int * __errno_location(void);
func X__errno_location(t *TLS) uintptr {
	return t.errnop
}

func Start(main func(*TLS, int32, uintptr) int32) {
	runtime.LockOSThread()
	t := NewTLS()
	argv := mustCalloc(t, types.Size_t((len(os.Args)+1)*int(uintptrSize)))
	p := argv
	for _, v := range os.Args {
		s := mustCalloc(t, types.Size_t(len(v)+1))
		copy((*RawMem)(unsafe.Pointer(s))[:len(v):len(v)], v)
		*(*uintptr)(unsafe.Pointer(p)) = s
		p += uintptrSize
	}
	SetEnviron(t, os.Environ())
	Xexit(t, main(t, int32(len(os.Args)), argv))
}

func SetEnviron(t *TLS, env []string) {
	p := mustCalloc(t, types.Size_t((len(env)+1)*(int(uintptrSize))))
	*(*uintptr)(unsafe.Pointer(EnvironP())) = p
	for _, v := range env {
		s := mustCalloc(t, types.Size_t(len(v)+1))
		copy((*(*RawMem)(unsafe.Pointer(s)))[:len(v):len(v)], v)
		*(*uintptr)(unsafe.Pointer(p)) = s
		p += uintptrSize
	}
}

// void setbuf(FILE *stream, char *buf);
func Xsetbuf(t *TLS, stream, buf uintptr) {
	//TODO panic(todo(""))
}

// size_t confstr(int name, char *buf, size_t len);
func Xconfstr(t *TLS, name int32, buf uintptr, len types.Size_t) types.Size_t {
	panic(todo(""))
}

// int puts(const char *s);
func Xputs(t *TLS, s uintptr) int32 {
	panic(todo(""))
}

var (
	randomMu  sync.Mutex
	randomGen = rand.New(rand.NewSource(42))
)

// long int random(void);
func Xrandom(t *TLS) long {
	randomMu.Lock()
	r := randomGen.Int63n(math.MaxInt32 + 1)
	randomMu.Unlock()
	return long(r)
}

// void *malloc(size_t size);
func Xmalloc(t *TLS, n types.Size_t) uintptr {
	if n == 0 {
		return 0
	}

	allocMu.Lock()

	defer allocMu.Unlock()

	p, err := allocator.UintptrMalloc(int(n))
	if err != nil {
		t.setErrno(errno.ENOMEM)
		return 0
	}

	return p
}

// void *calloc(size_t nmemb, size_t size);
func Xcalloc(t *TLS, n, size types.Size_t) uintptr {
	rq := int(n * size)
	if rq == 0 {
		return 0
	}

	allocMu.Lock()

	defer allocMu.Unlock()

	p, err := allocator.UintptrCalloc(int(n * size))
	if err != nil {
		t.setErrno(errno.ENOMEM)
		return 0
	}

	return p
}

// void *realloc(void *ptr, size_t size);
func Xrealloc(t *TLS, ptr uintptr, size types.Size_t) uintptr {
	allocMu.Lock()

	defer allocMu.Unlock()

	p, err := allocator.UintptrRealloc(ptr, int(size))
	if err != nil {
		t.setErrno(errno.ENOMEM)
		return 0
	}

	return p
}

// void free(void *ptr);
func Xfree(t *TLS, p uintptr) {
	if p == 0 {
		return
	}

	allocMu.Lock()

	defer allocMu.Unlock()

	allocator.UintptrFree(p)
}

func write(b []byte) (int, error) {
	if dmesgs {
		dmesg("%v: %s", origin(1), b)
	}
	if _, err := os.Stdout.Write(b); err != nil {
		return -1, err
	}

	return len(b), nil
}

func X__builtin_abort(t *TLS)                                        { Xabort(t) }
func X__builtin_abs(t *TLS, j int32) int32                           { return Xabs(t, j) }
func X__builtin_clzll(t *TLS, n uint64) int32                        { return int32(mbits.LeadingZeros64(n)) }
func X__builtin_constant_p_impl()                                    { panic(todo("internal error: should never be called")) }
func X__builtin_copysign(t *TLS, x, y float64) float64               { return Xcopysign(t, x, y) }
func X__builtin_copysignf(t *TLS, x, y float32) float32              { return Xcopysignf(t, x, y) }
func X__builtin_exit(t *TLS, status int32)                           { Xexit(t, status) }
func X__builtin_expect(t *TLS, exp, c long) long                     { return exp }
func X__builtin_fabs(t *TLS, x float64) float64                      { return Xfabs(t, x) }
func X__builtin_free(t *TLS, ptr uintptr)                            { Xfree(t, ptr) }
func X__builtin_huge_val(t *TLS) float64                             { return math.Inf(1) }
func X__builtin_huge_valf(t *TLS) float32                            { return float32(math.Inf(1)) }
func X__builtin_inf(t *TLS) float64                                  { return math.Inf(1) }
func X__builtin_inff(t *TLS) float32                                 { return float32(math.Inf(1)) }
func X__builtin_malloc(t *TLS, size types.Size_t) uintptr            { return Xmalloc(t, size) }
func X__builtin_memcmp(t *TLS, s1, s2 uintptr, n types.Size_t) int32 { return Xmemcmp(t, s1, s2, n) }
func X__builtin_nanf(t *TLS, s uintptr) float32                      { return float32(math.NaN()) }
func X__builtin_prefetch(t *TLS, addr, args uintptr)                 {}
func X__builtin_printf(t *TLS, s, args uintptr) int32                { return Xprintf(t, s, args) }
func X__builtin_strchr(t *TLS, s uintptr, c int32) uintptr           { return Xstrchr(t, s, c) }
func X__builtin_strcmp(t *TLS, s1, s2 uintptr) int32                 { return Xstrcmp(t, s1, s2) }
func X__builtin_strcpy(t *TLS, dest, src uintptr) uintptr            { return Xstrcpy(t, dest, src) }
func X__builtin_strlen(t *TLS, s uintptr) types.Size_t               { return Xstrlen(t, s) }
func X__builtin_trap(t *TLS)                                         { Xabort(t) }
func X__isnan(t *TLS, arg float64) int32                             { return X__builtin_isnan(t, arg) }
func X__isnanf(t *TLS, arg float32) int32                            { return Xisnanf(t, arg) }
func X__isnanl(t *TLS, arg float64) int32                            { return Xisnanl(t, arg) }
func Xvfprintf(t *TLS, stream, format, ap uintptr) int32             { return Xfprintf(t, stream, format, ap) }

// char * __builtin___strcpy_chk (char *dest, const char *src, size_t os);
func X__builtin___strcpy_chk(t *TLS, dest, src uintptr, os types.Size_t) uintptr {
	return Xstrcpy(t, dest, src)
}

func X__builtin_mmap(t *TLS, addr uintptr, length types.Size_t, prot, flags, fd int32, offset types.Off_t) uintptr {
	return Xmmap(t, addr, length, prot, flags, fd, offset)
}

// uint16_t __builtin_bswap16 (uint32_t x)
func X__builtin_bswap16(t *TLS, x uint16) uint16 {
	return x<<8 |
		x>>8
}

// uint32_t __builtin_bswap32 (uint32_t x)
func X__builtin_bswap32(t *TLS, x uint32) uint32 {
	return x<<24 |
		x&0xff00<<8 |
		x&0xff0000>>8 |
		x>>24
}

// uint64_t __builtin_bswap64 (uint64_t x)
func X__builtin_bswap64(t *TLS, x uint64) uint64 {
	return x<<56 |
		x&0xff00<<40 |
		x&0xff0000<<24 |
		x&0xff000000<<8 |
		x&0xff00000000>>8 |
		x&0xff0000000000>>24 |
		x&0xff000000000000>>40 |
		x>>56
}

// bool __builtin_add_overflow (type1 a, type2 b, type3 *res)
func X__builtin_add_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	r, ovf := mathutil.AddOverflowInt64(a, b)
	*(*int64)(unsafe.Pointer(res)) = r
	return Bool32(ovf)
}

// bool __builtin_add_overflow (type1 a, type2 b, type3 *res)
func X__builtin_add_overflowUint32(t *TLS, a, b uint32, res uintptr) int32 {
	r := a + b
	*(*uint32)(unsafe.Pointer(res)) = r
	return Bool32(r < a)
}

// bool __builtin_add_overflow (type1 a, type2 b, type3 *res)
func X__builtin_add_overflowUint64(t *TLS, a, b uint64, res uintptr) int32 {
	r := a + b
	*(*uint64)(unsafe.Pointer(res)) = r
	return Bool32(r < a)
}

// bool __builtin_sub_overflow (type1 a, type2 b, type3 *res)
func X__builtin_sub_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	r, ovf := mathutil.SubOverflowInt64(a, b)
	*(*int64)(unsafe.Pointer(res)) = r
	return Bool32(ovf)
}

// bool __builtin_mul_overflow (type1 a, type2 b, type3 *res)
func X__builtin_mul_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	r, ovf := mathutil.MulOverflowInt64(a, b)
	*(*int64)(unsafe.Pointer(res)) = r
	return Bool32(ovf)
}

func X__builtin_unreachable(t *TLS) {
	fmt.Fprintf(os.Stderr, "unrechable\n")
	os.Stderr.Sync()
	Xexit(t, 1)
}

func X__builtin_snprintf(t *TLS, str uintptr, size types.Size_t, format, args uintptr) int32 {
	return Xsnprintf(t, str, size, format, args)
}

func X__builtin_sprintf(t *TLS, str, format, args uintptr) (r int32) {
	return Xsprintf(t, str, format, args)
}

func X__builtin_memcpy(t *TLS, dest, src uintptr, n types.Size_t) (r uintptr) {
	return Xmemcpy(t, dest, src, n)
}

// void * __builtin___memcpy_chk (void *dest, const void *src, size_t n, size_t os);
func X__builtin___memcpy_chk(t *TLS, dest, src uintptr, n, os types.Size_t) (r uintptr) {
	if os != ^types.Size_t(0) && n < os {
		Xabort(t)
	}

	return Xmemcpy(t, dest, src, n)
}

func X__builtin_memset(t *TLS, s uintptr, c int32, n types.Size_t) uintptr {
	return Xmemset(t, s, c, n)
}

// void * __builtin___memset_chk (void *s, int c, size_t n, size_t os);
func X__builtin___memset_chk(t *TLS, s uintptr, c int32, n, os types.Size_t) uintptr {
	if os < n {
		Xabort(t)
	}

	return Xmemset(t, s, c, n)
}

// size_t __builtin_object_size (const void * ptr, int type)
func X__builtin_object_size(t *TLS, p uintptr, typ int32) types.Size_t {
	return ^types.Size_t(0) //TODO frontend magic
}

var atomicLoadStore16 sync.Mutex

func AtomicLoadNUint16(ptr uintptr, memorder int16) uint16 {
	atomicLoadStore16.Lock()
	r := *(*uint16)(unsafe.Pointer(ptr))
	atomicLoadStore16.Unlock()
	return r
}

func AtomicStoreNUint16(ptr uintptr, val uint16, memorder int32) {
	atomicLoadStore16.Lock()
	*(*uint16)(unsafe.Pointer(ptr)) = val
	atomicLoadStore16.Unlock()
}

// int sprintf(char *str, const char *format, ...);
func Xsprintf(t *TLS, str, format, args uintptr) (r int32) {
	b := printf(format, args)
	r = int32(len(b))
	copy((*RawMem)(unsafe.Pointer(str))[:r:r], b)
	*(*byte)(unsafe.Pointer(str + uintptr(r))) = 0
	return int32(len(b))
}

// int __builtin___sprintf_chk (char *s, int flag, size_t os, const char *fmt, ...);
func X__builtin___sprintf_chk(t *TLS, s uintptr, flag int32, os types.Size_t, format, args uintptr) (r int32) {
	return Xsprintf(t, s, format, args)
}

// void qsort(void *base, size_t nmemb, size_t size, int (*compar)(const void *, const void *));
func Xqsort(t *TLS, base uintptr, nmemb, size types.Size_t, compar uintptr) {
	sort.Sort(&sorter{
		len:  int(nmemb),
		base: base,
		sz:   uintptr(size),
		f: (*struct {
			f func(*TLS, uintptr, uintptr) int32
		})(unsafe.Pointer(&struct{ uintptr }{compar})).f,
		t: t,
	})
}

// void __assert_fail(const char * assertion, const char * file, unsigned int line, const char * function);
func X__assert_fail(t *TLS, assertion, file uintptr, line uint32, function uintptr) {
	fmt.Fprintf(os.Stderr, "assertion failure: %s:%d.%s: %s\n", GoString(file), line, GoString(function), GoString(assertion))
	os.Stderr.Sync()
	Xexit(t, 1)
}

// int vprintf(const char *format, va_list ap);
func Xvprintf(t *TLS, s, ap uintptr) int32 { return Xprintf(t, s, ap) }

// int __isoc99_sscanf(const char *str, const char *format, ...);
func X__isoc99_sscanf(t *TLS, str, format, va uintptr) int32 {
	r := scanf(strings.NewReader(GoString(str)), format, va)
	if dmesgs {
		dmesg("%v: %q %q: %d", origin(1), GoString(str), GoString(format), r)
	}
	return r
}

// int sscanf(const char *str, const char *format, ...);
func Xsscanf(t *TLS, str, format, va uintptr) int32 {
	r := scanf(strings.NewReader(GoString(str)), format, va)
	if dmesgs {
		dmesg("%v: %q %q: %d", origin(1), GoString(str), GoString(format), r)
	}
	return r
}

// int vsprintf(char *str, const char *format, va_list ap);
func Xvsprintf(t *TLS, str, format, va uintptr) int32 {
	panic(todo(""))
}

// int vsnprintf(char *str, size_t size, const char *format, va_list ap);
func Xvsnprintf(t *TLS, str uintptr, size types.Size_t, format, va uintptr) int32 {
	panic(todo(""))
}

// int obstack_vprintf (struct obstack *obstack, const char *template, va_list ap)
func Xobstack_vprintf(t *TLS, obstack, template, va uintptr) int32 {
	panic(todo(""))
}

// extern void _obstack_newchunk(struct obstack *, int);
func X_obstack_newchunk(t *TLS, obstack uintptr, length int32) int32 {
	panic(todo(""))
}

// int _obstack_begin (struct obstack *h, _OBSTACK_SIZE_T size, _OBSTACK_SIZE_T alignment,	void *(*chunkfun) (size_t),  void (*freefun) (void *))
func X_obstack_begin(t *TLS, obstack uintptr, size, alignment int32, chunkfun, freefun uintptr) int32 {
	panic(todo(""))
}

// void obstack_free (struct obstack *h, void *obj)
func Xobstack_free(t *TLS, obstack, obj uintptr) {
	panic(todo(""))
}

// unsigned int sleep(unsigned int seconds);
func Xsleep(t *TLS, seconds uint32) uint32 {
	gotime.Sleep(gotime.Second * gotime.Duration(seconds))
	return 0
}

// size_t strcspn(const char *s, const char *reject);
func Xstrcspn(t *TLS, s, reject uintptr) (r types.Size_t) {
	bits := newBits(256)
	for {
		c := *(*byte)(unsafe.Pointer(reject))
		if c == 0 {
			break
		}

		reject++
		bits.set(int(c))
	}
	for {
		c := *(*byte)(unsafe.Pointer(s))
		if c == 0 || bits.has(int(c)) {
			return r
		}

		s++
		r++
	}
}

// int printf(const char *format, ...);
func Xprintf(t *TLS, format, args uintptr) int32 {
	n, _ := write(printf(format, args))
	return int32(n)
}

// int snprintf(char *str, size_t size, const char *format, ...);
func Xsnprintf(t *TLS, str uintptr, size types.Size_t, format, args uintptr) (r int32) {
	switch size {
	case 0:
		return 0
	case 1:
		*(*byte)(unsafe.Pointer(str)) = 0
		return 0
	}

	b := printf(format, args)
	if len(b)+1 > int(size) {
		b = b[:size-1]
	}
	r = int32(len(b))
	copy((*RawMem)(unsafe.Pointer(str))[:r:r], b)
	*(*byte)(unsafe.Pointer(str + uintptr(r))) = 0
	return r
}

// int __builtin___snprintf_chk(char * str, size_t maxlen, int flag, size_t strlen, const char * format);
func X__builtin___snprintf_chk(t *TLS, str uintptr, maxlen types.Size_t, flag int32, strlen types.Size_t, format, args uintptr) (r int32) {
	if maxlen != ^types.Size_t(0) && strlen > maxlen {
		Xabort(t)
	}

	return Xsnprintf(t, str, strlen, format, args)
}

// int abs(int j);
func Xabs(t *TLS, j int32) int32 {
	if j >= 0 {
		return j
	}

	return -j
}

func X__builtin_isnan(t *TLS, x float64) int32    { return Bool32(math.IsNaN(x)) }
func Xacos(t *TLS, x float64) float64             { return math.Acos(x) }
func Xasin(t *TLS, x float64) float64             { return math.Asin(x) }
func Xatan(t *TLS, x float64) float64             { return math.Atan(x) }
func Xatan2(t *TLS, x, y float64) float64         { return math.Atan2(x, y) }
func Xceil(t *TLS, x float64) float64             { return math.Ceil(x) }
func Xcopysign(t *TLS, x, y float64) float64      { return math.Copysign(x, y) }
func Xcopysignf(t *TLS, x, y float32) float32     { return float32(math.Copysign(float64(x), float64(y))) }
func Xcos(t *TLS, x float64) float64              { return math.Cos(x) }
func Xcosf(t *TLS, x float32) float32             { return float32(math.Cos(float64(x))) }
func Xcosh(t *TLS, x float64) float64             { return math.Cosh(x) }
func Xexp(t *TLS, x float64) float64              { return math.Exp(x) }
func Xfabs(t *TLS, x float64) float64             { return math.Abs(x) }
func Xfabsf(t *TLS, x float32) float32            { return float32(math.Abs(float64(x))) }
func Xfloor(t *TLS, x float64) float64            { return math.Floor(x) }
func Xfmod(t *TLS, x, y float64) float64          { return math.Mod(x, y) }
func Xhypot(t *TLS, x, y float64) float64         { return math.Hypot(x, y) }
func Xisnan(t *TLS, x float64) int32              { return X__builtin_isnan(t, x) }
func Xisnanf(t *TLS, x float32) int32             { return Bool32(math.IsNaN(float64(x))) }
func Xisnanl(t *TLS, x float64) int32             { return Bool32(math.IsNaN(x)) } // ccgo has to handle long double as double as Go does not support long double.
func Xldexp(t *TLS, x float64, exp int32) float64 { return math.Ldexp(x, int(exp)) }
func Xlog(t *TLS, x float64) float64              { return math.Log(x) }
func Xlog10(t *TLS, x float64) float64            { return math.Log10(x) }
func Xround(t *TLS, x float64) float64            { return math.Round(x) }
func Xsin(t *TLS, x float64) float64              { return math.Sin(x) }
func Xsinf(t *TLS, x float32) float32             { return float32(math.Sin(float64(x))) }
func Xsinh(t *TLS, x float64) float64             { return math.Sinh(x) }
func Xsqrt(t *TLS, x float64) float64             { return math.Sqrt(x) }
func Xtan(t *TLS, x float64) float64              { return math.Tan(x) }
func Xtanh(t *TLS, x float64) float64             { return math.Tanh(x) }

var nextRand = uint64(1)

// int rand(void);
func Xrand(t *TLS) int32 {
	nextRand = nextRand*1103515245 + 12345
	return int32(uint32(nextRand / (math.MaxUint32 + 1) % math.MaxInt32))
}

func Xpow(t *TLS, x, y float64) float64 {
	r := math.Pow(x, y)
	if x > 0 && r == 1 && y >= -1.0000000000000000715e-18 && y < -1e-30 {
		r = 0.9999999999999999
	}
	return r
}

func Xfrexp(t *TLS, x float64, exp uintptr) float64 {
	f, e := math.Frexp(x)
	*(*int32)(unsafe.Pointer(exp)) = int32(e)
	return f
}

func Xmodf(t *TLS, x float64, iptr uintptr) float64 {
	i, f := math.Modf(x)
	*(*float64)(unsafe.Pointer(iptr)) = i
	return f
}

// char *strncpy(char *dest, const char *src, size_t n)
func Xstrncpy(t *TLS, dest, src uintptr, n types.Size_t) (r uintptr) {
	r = dest
	for c := *(*int8)(unsafe.Pointer(src)); c != 0 && n > 0; n-- {
		*(*int8)(unsafe.Pointer(dest)) = c
		dest++
		src++
		c = *(*int8)(unsafe.Pointer(src))
	}
	for ; uintptr(n) > 0; n-- {
		*(*int8)(unsafe.Pointer(dest)) = 0
		dest++
	}
	return r
}

// char * __builtin___strncpy_chk (char *dest, const char *src, size_t n, size_t os);
func X__builtin___strncpy_chk(t *TLS, dest, src uintptr, n, os types.Size_t) (r uintptr) {
	if n != ^types.Size_t(0) && os < n {
		Xabort(t)
	}

	return Xstrncpy(t, dest, src, n)
}

// int strcmp(const char *s1, const char *s2)
func Xstrcmp(t *TLS, s1, s2 uintptr) int32 {
	for {
		ch1 := *(*byte)(unsafe.Pointer(s1))
		s1++
		ch2 := *(*byte)(unsafe.Pointer(s2))
		s2++
		if ch1 != ch2 || ch1 == 0 || ch2 == 0 {
			return int32(ch1) - int32(ch2)
		}
	}
}

// size_t strlen(const char *s)
func Xstrlen(t *TLS, s uintptr) (r types.Size_t) {
	for ; *(*int8)(unsafe.Pointer(s)) != 0; s++ {
		r++
	}
	return r
}

// char *strcat(char *dest, const char *src)
func Xstrcat(t *TLS, dest, src uintptr) (r uintptr) {
	r = dest
	for *(*int8)(unsafe.Pointer(dest)) != 0 {
		dest++
	}
	for {
		c := *(*int8)(unsafe.Pointer(src))
		src++
		*(*int8)(unsafe.Pointer(dest)) = c
		dest++
		if c == 0 {
			return r
		}
	}
}

// char * __builtin___strcat_chk (char *dest, const char *src, size_t os);
func X__builtin___strcat_chk(t *TLS, dest, src uintptr, os types.Size_t) (r uintptr) {
	return Xstrcat(t, dest, src)
}

// int strncmp(const char *s1, const char *s2, size_t n)
func Xstrncmp(t *TLS, s1, s2 uintptr, n types.Size_t) int32 {
	var ch1, ch2 byte
	for ; n != 0; n-- {
		ch1 = *(*byte)(unsafe.Pointer(s1))
		s1++
		ch2 = *(*byte)(unsafe.Pointer(s2))
		s2++
		if ch1 != ch2 {
			return int32(ch1) - int32(ch2)
		}

		if ch1 == 0 {
			return 0
		}
	}
	return 0
}

// char *strcpy(char *dest, const char *src)
func Xstrcpy(t *TLS, dest, src uintptr) (r uintptr) {
	r = dest
	// src0 := src
	for ; ; dest++ {
		c := *(*int8)(unsafe.Pointer(src))
		src++
		*(*int8)(unsafe.Pointer(dest)) = c
		if c == 0 {
			return r
		}
	}
}

// char *strchr(const char *s, int c)
func Xstrchr(t *TLS, s uintptr, c int32) uintptr {
	for {
		ch2 := *(*byte)(unsafe.Pointer(s))
		if ch2 == byte(c) {
			return s
		}

		if ch2 == 0 {
			return 0
		}

		s++
	}
}

// char *strrchr(const char *s, int c)
func Xstrrchr(t *TLS, s uintptr, c int32) (r uintptr) {
	for {
		ch2 := *(*byte)(unsafe.Pointer(s))
		if ch2 == 0 {
			return r
		}

		if ch2 == byte(c) {
			r = s
		}
		s++
	}
}

// void *memset(void *s, int c, size_t n)
func Xmemset(t *TLS, s uintptr, c int32, n types.Size_t) uintptr {
	if n != 0 {
		c := byte(c & 0xff)

		//this will make sure that on platforms where they are not equally alligned
		//we clear out the first few bytes until allignment
		bytesBeforeAllignment := s % unsafe.Alignof(uint64(0))
		if bytesBeforeAllignment > uintptr(n) {
			bytesBeforeAllignment = uintptr(n)
		}
		b := (*RawMem)(unsafe.Pointer(s))[:bytesBeforeAllignment:bytesBeforeAllignment]
		n -= types.Size_t(bytesBeforeAllignment)
		for i := range b {
			b[i] = c
		}
		i64 := uint64(c) + uint64(c)<<8 + uint64(c)<<16 + uint64(c)<<24 + uint64(c)<<32 + uint64(c)<<40 + uint64(c)<<48 + uint64(c)<<56
		b8 := (*RawMem64)(unsafe.Pointer(s + bytesBeforeAllignment))[: n/8 : n/8]
		for i := range b8 {
			b8[i] = i64
		}
		if n%8 != 0 {
			b = (*RawMem)(unsafe.Pointer(s + bytesBeforeAllignment + uintptr(n-n%8)))[: n%8 : n%8]
			for i := range b {
				b[i] = c
			}
		}
	}
	return s
}

// void *memcpy(void *dest, const void *src, size_t n);
func Xmemcpy(t *TLS, dest, src uintptr, n types.Size_t) (r uintptr) {
	s := (*RawMem)(unsafe.Pointer(src))[:n:n]
	d := (*RawMem)(unsafe.Pointer(dest))[:n:n]
	copy(d, s)
	return dest
}

// int memcmp(const void *s1, const void *s2, size_t n);
func Xmemcmp(t *TLS, s1, s2 uintptr, n types.Size_t) int32 {
	for ; n != 0; n-- {
		c1 := *(*byte)(unsafe.Pointer(s1))
		s1++
		c2 := *(*byte)(unsafe.Pointer(s2))
		s2++
		if c1 < c2 {
			return -1
		}

		if c1 > c2 {
			return 1
		}
	}
	return 0
}

// void *memchr(const void *s, int c, size_t n);
func Xmemchr(t *TLS, s uintptr, c int32, n types.Size_t) uintptr {
	for ; n != 0; n-- {
		if *(*byte)(unsafe.Pointer(s)) == byte(c) {
			return s
		}

		s++
	}
	return 0
}

// void *memmove(void *dest, const void *src, size_t n);
func Xmemmove(t *TLS, dest, src uintptr, n types.Size_t) uintptr {
	copy((*RawMem)(unsafe.Pointer(uintptr(dest)))[:n:n], (*RawMem)(unsafe.Pointer(uintptr(src)))[:n:n])
	return dest
}

// void * __builtin___memmove_chk (void *dest, const void *src, size_t n, size_t os);
func X__builtin___memmove_chk(t *TLS, dest, src uintptr, n, os types.Size_t) uintptr {
	if os != ^types.Size_t(0) && os < n {
		Xabort(t)
	}

	return Xmemmove(t, dest, src, n)
}

var getenvOnce sync.Once

// char *getenv(const char *name);
func Xgetenv(t *TLS, name uintptr) uintptr {
	p := Environ()
	if p == 0 {
		getenvOnce.Do(func() {
			SetEnviron(t, os.Environ())
			p = Environ()
		})
	}

	return getenv(p, GoString(name))
}

func getenv(p uintptr, nm string) uintptr {
	for ; ; p += uintptrSize {
		q := *(*uintptr)(unsafe.Pointer(p))
		if q == 0 {
			return 0
		}

		s := GoString(q)
		a := strings.SplitN(s, "=", 2)
		if len(a) != 2 {
			panic(todo("%q %q %q", nm, s, a))
		}

		if a[0] == nm {
			return q + uintptr(len(nm)) + 1
		}
	}
}

// char *strstr(const char *haystack, const char *needle);
func Xstrstr(t *TLS, haystack, needle uintptr) uintptr {
	hs := GoString(haystack)
	nd := GoString(needle)
	if i := strings.Index(hs, nd); i >= 0 {
		r := haystack + uintptr(i)
		return r
	}

	return 0
}

// int putc(int c, FILE *stream);
func Xputc(t *TLS, c int32, fp uintptr) int32 {
	return Xfputc(t, c, fp)
}

// int atoi(const char *nptr);
func Xatoi(t *TLS, nptr uintptr) int32 {
	_, neg, _, n, _ := strToUint64(t, nptr, 10)
	switch {
	case neg:
		return int32(-n)
	default:
		return int32(n)
	}
}

// double atof(const char *nptr);
func Xatof(t *TLS, nptr uintptr) float64 {
	n, _ := strToFloatt64(t, nptr, 64)
	return n
}

// int tolower(int c);
func Xtolower(t *TLS, c int32) int32 {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}

	return c
}

// int toupper(int c);
func Xtoupper(t *TLS, c int32) int32 {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}

	return c
}

// int isatty(int fd);
func Xisatty(t *TLS, fd int32) int32 {
	return Bool32(isatty.IsTerminal(uintptr(fd)))
}

// char *strdup(const char *s);
func Xstrdup(t *TLS, s uintptr) uintptr {
	panic(todo(""))
}

// long atol(const char *nptr);
func Xatol(t *TLS, nptr uintptr) long {
	_, neg, _, n, _ := strToUint64(t, nptr, 10)
	switch {
	case neg:
		return long(-n)
	default:
		return long(n)
	}
}

// time_t mktime(struct tm *tm);
func Xmktime(t *TLS, ptm uintptr) types.Time_t {
	loc := gotime.Local
	if r := getenv(Environ(), "TZ"); r != 0 {
		zone, off := parseZone(GoString(r))
		loc = gotime.FixedZone(zone, off)
	}
	tt := gotime.Date(
		int((*time.Tm)(unsafe.Pointer(ptm)).Ftm_year+1900),
		gotime.Month((*time.Tm)(unsafe.Pointer(ptm)).Ftm_mon+1),
		int((*time.Tm)(unsafe.Pointer(ptm)).Ftm_mday),
		int((*time.Tm)(unsafe.Pointer(ptm)).Ftm_hour),
		int((*time.Tm)(unsafe.Pointer(ptm)).Ftm_min),
		int((*time.Tm)(unsafe.Pointer(ptm)).Ftm_sec),
		0,
		loc,
	)
	(*time.Tm)(unsafe.Pointer(ptm)).Ftm_wday = int32(tt.Weekday())
	(*time.Tm)(unsafe.Pointer(ptm)).Ftm_yday = int32(tt.YearDay() - 1)
	return types.Time_t(tt.Unix())
}

// char *strpbrk(const char *s, const char *accept);
func Xstrpbrk(t *TLS, s, accept uintptr) uintptr {
	bits := newBits(256)
	for {
		b := *(*byte)(unsafe.Pointer(accept))
		if b == 0 {
			break
		}

		bits.set(int(b))
		accept++
	}
	for {
		b := *(*byte)(unsafe.Pointer(s))
		if b == 0 {
			return 0
		}

		if bits.has(int(b)) {
			return s
		}

		s++
	}
}

// int strcasecmp(const char *s1, const char *s2);
func Xstrcasecmp(t *TLS, s1, s2 uintptr) int32 {
	for {
		ch1 := *(*byte)(unsafe.Pointer(s1))
		if ch1 >= 'a' && ch1 <= 'z' {
			ch1 = ch1 - ('a' - 'A')
		}
		s1++
		ch2 := *(*byte)(unsafe.Pointer(s2))
		if ch2 >= 'a' && ch2 <= 'z' {
			ch2 = ch2 - ('a' - 'A')
		}
		s2++
		if ch1 != ch2 || ch1 == 0 || ch2 == 0 {
			r := int32(ch1) - int32(ch2)
			return r
		}
	}
}

var __ctype_b_table = [...]uint16{ //TODO use symbolic constants
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002,
	0x0002, 0x2003, 0x2002, 0x2002, 0x2002, 0x2002, 0x0002, 0x0002,
	0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002,
	0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002, 0x0002,
	0x6001, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004,
	0xc004, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004,
	0xd808, 0xd808, 0xd808, 0xd808, 0xd808, 0xd808, 0xd808, 0xd808,
	0xd808, 0xd808, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004,
	0xc004, 0xd508, 0xd508, 0xd508, 0xd508, 0xd508, 0xd508, 0xc508,
	0xc508, 0xc508, 0xc508, 0xc508, 0xc508, 0xc508, 0xc508, 0xc508,
	0xc508, 0xc508, 0xc508, 0xc508, 0xc508, 0xc508, 0xc508, 0xc508,
	0xc508, 0xc508, 0xc508, 0xc004, 0xc004, 0xc004, 0xc004, 0xc004,
	0xc004, 0xd608, 0xd608, 0xd608, 0xd608, 0xd608, 0xd608, 0xc608,
	0xc608, 0xc608, 0xc608, 0xc608, 0xc608, 0xc608, 0xc608, 0xc608,
	0xc608, 0xc608, 0xc608, 0xc608, 0xc608, 0xc608, 0xc608, 0xc608,
	0xc608, 0xc608, 0xc608, 0xc004, 0xc004, 0xc004, 0xc004, 0x0002,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var ptable = uintptr(unsafe.Pointer(&__ctype_b_table[128]))

// const unsigned short * * __ctype_b_loc (void);
func X__ctype_b_loc(t *TLS) uintptr {
	return uintptr(unsafe.Pointer(&ptable))
}

func Xntohs(t *TLS, netshort uint16) uint16 {
	return uint16((*[2]byte)(unsafe.Pointer(&netshort))[0])<<8 | uint16((*[2]byte)(unsafe.Pointer(&netshort))[1])
}

// uint16_t htons(uint16_t hostshort);
func Xhtons(t *TLS, hostshort uint16) uint16 {
	var a [2]byte
	a[0] = byte(hostshort >> 8)
	a[1] = byte(hostshort)
	return *(*uint16)(unsafe.Pointer(&a))
}

// uint32_t htonl(uint32_t hostlong);
func Xhtonl(t *TLS, hostlong uint32) uint32 {
	var a [4]byte
	a[0] = byte(hostlong >> 24)
	a[1] = byte(hostlong >> 16)
	a[2] = byte(hostlong >> 8)
	a[3] = byte(hostlong)
	return *(*uint32)(unsafe.Pointer(&a))
}

// FILE *fopen(const char *pathname, const char *mode);
func Xfopen(t *TLS, pathname, mode uintptr) uintptr {
	return Xfopen64(t, pathname, mode) //TODO 32 bit
}

// void sqlite3_log(int iErrCode, const char *zFormat, ...);
func X__ccgo_sqlite3_log(t *TLS, iErrCode int32, zFormat uintptr, args uintptr) {
	if dmesgs {
		dmesg("%v: iErrCode: %v, msg: %s\n%s", origin(1), iErrCode, printf(zFormat, args), debug.Stack())
	}
}

// int _IO_putc(int __c, _IO_FILE *__fp);
func X_IO_putc(t *TLS, c int32, fp uintptr) int32 {
	return Xputc(t, c, fp)
}

// int atexit(void (*function)(void));
func Xatexit(t *TLS, function uintptr) int32 {
	panic(todo(""))
}

// int vasprintf(char **strp, const char *fmt, va_list ap);
func Xvasprintf(t *TLS, strp, fmt, ap uintptr) int32 {
	panic(todo(""))
}
