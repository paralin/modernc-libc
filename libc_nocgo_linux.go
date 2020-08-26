// Copyright 2020 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !cgo

// Package libc provides a subset of functionality of C libc.
package libc // import "modernc.org/libc"

import (
	"bufio"
	"os"
	"strings"
	"unsafe"

	"golang.org/x/sys/unix"
	"modernc.org/libc/errno"
	"modernc.org/libc/langinfo"
	"modernc.org/libc/netinet/in"
	"modernc.org/libc/stdio"
	"modernc.org/libc/sys/socket"
	"modernc.org/libc/sys/types"
	"modernc.org/libc/termios"
)

type file uintptr

func (f file) fd() int      { return int((*stdio.FILE)(unsafe.Pointer(f)).F_fileno) }
func (f file) setFd(fd int) { (*stdio.FILE)(unsafe.Pointer(f)).F_fileno = int32(fd) }

func (f file) close(t *TLS) int32 {
	r := Xclose(t, int32(f.fd()))
	Xfree(t, uintptr(f))
	return r
}

func newFile(t *TLS, fd int) uintptr {
	p := Xcalloc(t, 1, types.Size_t(unsafe.Sizeof(stdio.FILE{})))
	if p == 0 {
		return 0
	}

	file(p).setFd(fd)
	return p
}

// int * __errno_location(void);
func X__errno_location(t *TLS) uintptr {
	return t.errnop
}

func (t *TLS) setErrno(err interface{}) { //TODO -> etc.go
again:
	switch x := err.(type) {
	case int:
		*(*int32)(unsafe.Pointer(X__errno_location(t))) = int32(x)
	case int32:
		*(*int32)(unsafe.Pointer(X__errno_location(t))) = x
	case *os.PathError:
		err = x.Err
		goto again
	case unix.Errno:
		*(*int32)(unsafe.Pointer(X__errno_location(t))) = int32(x)
	case *os.SyscallError:
		err = x.Err
		goto again
	default:
		panic(todo("%T", x))
	}
}

func Environ() uintptr {
	return Xenviron
}

func EnvironP() uintptr {
	return uintptr(unsafe.Pointer(&Xenviron))
}

func Xabort(t *TLS) {
	panic(todo(""))
}

func Xexit(t *TLS, status int32) {
	if len(Covered) != 0 { //TODO -> etc.go
		buf := bufio.NewWriter(os.Stdout)
		CoverReport(buf)
		buf.Flush()
	}
	for _, v := range atExit {
		v()
	}
	os.Exit(int(status))
}

// FILE *fdopen(int fd, const char *mode);
func Xfdopen(t *TLS, fd int32, mode uintptr) uintptr {
	panic(todo(""))
}

// int fseek(FILE *stream, long offset, int whence);
func Xfseek(t *TLS, stream uintptr, offset long, whence int32) int32 {
	n := Xlseek(t, int32(file(stream).fd()), offset, whence)
	if n < 0 {
		return -1
	}

	return int32(n)
}

// int fputc(int c, FILE *stream);
func Xfputc(t *TLS, c int32, stream uintptr) int32 {
	panic(todo(""))
}

// int fflush(FILE *stream);
func Xfflush(t *TLS, stream uintptr) int32 {
	return 0 //TODO
}

func Xtzset(t *TLS) {
	panic(todo(""))
}

// void *dlopen(const char *filename, int flags);
func Xdlopen(t *TLS, filename uintptr, flags int32) uintptr {
	panic(todo(""))
}

// char *dlerror(void);
func Xdlerror(t *TLS) uintptr {
	panic(todo(""))
}

// int dlclose(void *handle);
func Xdlclose(t *TLS, handle uintptr) int32 {
	panic(todo(""))
}

// void *dlsym(void *handle, const char *symbol);
func Xdlsym(t *TLS, handle, symbol uintptr) uintptr {
	panic(todo(""))
}

// struct servent *getservbyname(const char *name, const char *proto);
func Xgetservbyname(t *TLS, name, proto uintptr) uintptr {
	panic(todo(""))
}

// int getaddrinfo(const char *node, const char *service, const struct addrinfo *hints, struct addrinfo **res);
func Xgetaddrinfo(t *TLS, node, service, hints, res uintptr) int32 { //TODO not needed by sqlite
	panic(todo(""))
}

// const char *gai_strerror(int errcode);
func Xgai_strerror(t *TLS, errcode int32) uintptr {
	panic(todo(""))
}

// char *strerror(int errnum);
func Xstrerror(t *TLS, errnum int32) uintptr {
	panic(todo(""))
}

// int tcgetattr(int fd, struct termios *termios_p);
func Xtcgetattr(t *TLS, fd int32, termios_p uintptr) int32 {
	panic(todo(""))
}

// int tcsetattr(int fd, int optional_actions, const struct termios *termios_p);
func Xtcsetattr(t *TLS, fd, optional_actions int32, termios_p uintptr) int32 {
	panic(todo(""))
}

// ssize_t readlink(const char *restrict path, char *restrict buf, size_t bufsize);
func Xreadlink(t *TLS, path, buf uintptr, bufsize types.Size_t) types.Ssize_t {
	panic(todo(""))
}

// speed_t cfgetospeed(const struct termios *termios_p);
func Xcfgetospeed(t *TLS, termios_p uintptr) termios.Speed_t {
	panic(todo(""))
}

// int cfsetospeed(struct termios *termios_p, speed_t speed);
func Xcfsetospeed(t *TLS, termios_p uintptr, speed uint32) int32 {
	panic(todo(""))
}

// int cfsetispeed(struct termios *termios_p, speed_t speed);
func Xcfsetispeed(t *TLS, termios_p uintptr, speed uint32) int32 {
	panic(todo(""))
}

// char *realpath(const char *path, char *resolved_path);
func Xrealpath(t *TLS, path, resolved_path uintptr) uintptr {
	panic(todo(""))
}

// DIR *opendir(const char *name);
func Xopendir(t *TLS, name uintptr) uintptr {
	panic(todo(""))
}

// int closedir(DIR *dirp);
func Xclosedir(t *TLS, dir uintptr) int32 {
	panic(todo(""))
}

// struct dirent *readdir(DIR *dirp);
func Xreaddir64(t *TLS, dirp uintptr) uintptr {
	panic(todo(""))
}

// FTS *fts_open(char * const *path_argv, int options, int (*compar)(const FTSENT **, const FTSENT **));
func Xfts_open(t *TLS, path_argv uintptr, options int32, compar uintptr) uintptr {
	panic(todo(""))
}

// FTSENT *fts_read(FTS *ftsp);
func Xfts_read(t *TLS, ftsp uintptr) uintptr {
	panic(todo(""))
}

// int fts_close(FTS *ftsp);
func Xfts_close(t *TLS, ftsp uintptr) int32 {
	panic(todo(""))
}

// int mkstemps(char *template, int suffixlen);
func Xmkstemps(t *TLS, template uintptr, suffixlen int32) int32 {
	panic(todo(""))
}

// int mkstemp(char *template);
func Xmkstemp(t *TLS, template uintptr) int32 {
	panic(todo(""))
}

// void freeaddrinfo(struct addrinfo *res);
func Xfreeaddrinfo(t *TLS, res uintptr) {
	panic(todo(""))
}

func X__ccgo_in6addr_anyp(t *TLS) uintptr {
	panic(todo(""))
}

// int getnameinfo(const struct sockaddr *addr, socklen_t addrlen, char *host, socklen_t hostlen, char *serv, socklen_t servlen, int flags);
func Xgetnameinfo(t *TLS, addr uintptr, addrlen socket.Socklen_t, host uintptr, hostlen socket.Socklen_t, serv uintptr, servlen socket.Socklen_t, flags int32) int32 {
	panic(todo(""))
}

// struct tm *gmtime_r(const time_t *timep, struct tm *result);
func Xgmtime_r(t *TLS, timep, result uintptr) uintptr {
	panic(todo(""))
}

// char *setlocale(int category, const char *locale);
func Xsetlocale(t *TLS, category int32, locale uintptr) uintptr {
	panic(todo(""))
}

// char *nl_langinfo(nl_item item);
func Xnl_langinfo(t *TLS, item langinfo.Nl_item) uintptr {
	panic(todo(""))
}

// struct hostent *gethostbyname(const char *name);
func Xgethostbyname(t *TLS, name uintptr) uintptr {
	panic(todo(""))
}

// struct hostent *gethostbyaddr(const void *addr, socklen_t len, int type);
func Xgethostbyaddr(t *TLS, addr uintptr, len socket.Socklen_t, type1 int32) uintptr {
	panic(todo(""))
}

// uint32_t htonl(uint32_t hostlong);
func Xhtonl(t *TLS, hostlong uint32) uint32 {
	panic(todo(""))
}

// char *inet_ntoa(struct in_addr in);
func Xinet_ntoa(t *TLS, in1 in.In_addr) uintptr {
	panic(todo(""))
}

// int fclose(FILE *stream);
func Xfclose(t *TLS, stream uintptr) int32 {
	return file(stream).close(t)
}

// long ftell(FILE *stream);
func Xftell(t *TLS, stream uintptr) long {
	n := Xlseek(t, int32(file(stream).fd()), 0, stdio.SEEK_CUR)
	if n < 0 {
		return -1
	}

	return long(n)
}

// size_t fread(void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfread(t *TLS, ptr uintptr, size, nmemb types.Size_t, stream uintptr) types.Size_t {
	fd := file(stream).fd()
	m := Xread(t, int32(fd), ptr, size*nmemb)
	if m < 0 {
		return 0
	}

	return types.Size_t(m) / size
}

// size_t fwrite(const void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfwrite(t *TLS, ptr uintptr, size, nmemb types.Size_t, stream uintptr) types.Size_t {
	panic(todo(""))
}

// FILE *fopen64(const char *pathname, const char *mode);
func Xfopen64(t *TLS, pathname, mode uintptr) uintptr {
	s := GoString(pathname)
	m := strings.ReplaceAll(GoString(mode), "b", "")
	var fd int
	var err error
	switch m {
	case "r":
		if fd, err = unix.Open(s, os.O_RDONLY, 0666); err != nil {
			t.setErrno(err)
			return 0
		}
	case "r+":
		if fd, err = unix.Open(s, os.O_RDWR, 0666); err != nil {
			t.setErrno(err)
			return 0
		}
	case "w":
		if fd, err = unix.Open(s, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
			t.setErrno(err)
			return 0
		}
	case "w+":
		if fd, err = unix.Open(s, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
			t.setErrno(err)
			return 0
		}
	case "a":
		if fd, err = unix.Open(s, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err != nil {
			t.setErrno(err)
			return 0
		}
	case "a+":
		if fd, err = unix.Open(s, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
			t.setErrno(err)
			return 0
		}
	default:
		panic(m)
	}
	if p := newFile(t, fd); p != 0 {
		return p
	}

	t.setErrno(errno.ENOMEM)
	return 0
}

// int fgetc(FILE *stream);
func Xfgetc(t *TLS, stream uintptr) int32 {
	panic(todo(""))
}

// int ferror(FILE *stream);
func Xferror(t *TLS, stream uintptr) int32 {
	panic(todo(""))
}
