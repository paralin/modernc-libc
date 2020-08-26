// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"modernc.org/cc/v3"
)

var (
	goos   = runtime.GOOS
	goarch = runtime.GOARCH
)

func main() {
	hostConfigCmd := os.Getenv("GO_GENERATE_CPP")
	var hostConfigOpts []string
	if s := os.Getenv("GO_GENERATE_CPP_OPTS"); s != "" {
		hostConfigOpts = strings.Split(s, ",")
	}
	if s := os.Getenv("TARGET_GOOS"); s != "" {
		goos = s
	}
	if s := os.Getenv("TARGET_GOARCH"); s != "" {
		goarch = s
	}
	// fmt.Printf("%q %q %q\n", hostConfigOpts, goos, goarch)
	_, _, hostSysIncludes, err := cc.HostConfig(hostConfigCmd, hostConfigOpts...)
	if err != nil {
		fail(err)
	}

	g := []string{"etc.go", "libc.go"} //TODO -etc.go
	x, err := filepath.Glob(fmt.Sprintf("*_%s.go", goos))
	if err != nil {
		fail(err)
	}

	g = append(g, x...)
	if x, err = filepath.Glob(fmt.Sprintf("*_%s_%s.go", goos, goarch)); err != nil {
		fail(err)
	}

	g = append(g, x...)
	m := map[string]struct{}{}
	for _, v := range g {
		f, err := os.Open(v)
		if err != nil {
			fail(err)
		}

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			s := sc.Text()
			switch {
			case strings.HasPrefix(s, "func X"):
				s = s[len("func X"):]
				x := strings.IndexByte(s, '(')
				s = s[:x]
			case strings.HasPrefix(s, "var X"):
				s = s[len("var X"):]
				x := strings.IndexByte(s, ' ')
				s = s[:x]
			default:
				continue
			}

			m[s] = struct{}{}
		}
		if err := sc.Err(); err != nil {
			fail(err)
		}

		f.Close()
	}
	var a []string
	for k := range m {
		a = append(a, k)
	}
	sort.Strings(a)
	b := bytes.NewBuffer(nil)
	b.WriteString(`// Code generated by 'go generate' - DO NOT EDIT.

package libc // import "modernc.org/libc"

var CAPI = map[string]struct{}{`)

	for _, v := range a {
		fmt.Fprintf(b, "\n\t%q: {},", v)
	}
	b.WriteString("\n}")
	if err := ioutil.WriteFile(fmt.Sprintf("capi_%s_%s.go", goos, goarch), b.Bytes(), 0660); err != nil {
		fail(err)
	}

	ccgoHelpers()

	if err := libcHeaders(hostSysIncludes); err != nil {
		fail(err)
	}
}

func libcHeaders(paths []string) error {
	const cfile = "gen.c"
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		dir := path
		ok := false
		for _, v := range paths {
			full := filepath.Join(v, dir+".h")
			if fi, err := os.Stat(full); err == nil && !fi.IsDir() {
				ok = true
				break
			}
		}
		if !ok {
			return nil
		}

		src := fmt.Sprintf(`#include <%s.h>
static char _;
`, dir)
		fn := filepath.Join(dir, cfile)
		if err := ioutil.WriteFile(fn, []byte(src), 0660); err != nil {
			return err
		}

		defer os.Remove(fn)

		dest := filepath.Join(path, fmt.Sprintf("%s_%s_%s.go", filepath.Base(path), goos, goarch))
		base := filepath.Base(dir)
		cmd := exec.Command(
			"ccgo", fn,
			"-ccgo-crt-import-path", "",
			"-ccgo-export-defines", "",
			"-ccgo-export-enums", "",
			"-ccgo-export-externs", "X",
			"-ccgo-export-fields", "F",
			"-ccgo-export-structs", "",
			"-ccgo-export-typedefs", "",
			"-ccgo-host-config-cmd", os.Getenv("GO_GENERATE_CPP"),
			"-ccgo-host-config-opts", os.Getenv("GO_GENERATE_CPP_OPTS"),
			"-ccgo-long-double-is-double",
			"-ccgo-pkgname", base,
			"-o", dest,
		)
		out, err := cmd.CombinedOutput()
		sout := strings.TrimSpace(string(out) + "\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s%s\n", path, sout, err)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n%s", path, sout)
		}
		return nil
	})
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func ccgoHelpers() {
	var (
		signed = []string{
			"int8",
			"int16",
			"int32",
			"int64",
		}
		unsigned = []string{
			"uint8",
			"uint16",
			"uint32",
			"uint64",
		}
		ints   = append(signed[:len(signed):len(signed)], unsigned...)
		intptr = append(ints[:len(ints):len(ints)], "uintptr")
		arith  = append(ints[:len(ints):len(ints)], "float32", "float64")
		scalar = append(arith[:len(arith):len(arith)], []string{"uintptr"}...)
		sizes  = []string{"8", "16", "32", "64"}
	)

	b := bytes.NewBuffer(nil)
	b.WriteString(`// Code generated by 'go generate' - DO NOT EDIT.

package libc // import "modernc.org/libc"

import "unsafe"

`)
	for _, v := range scalar {
		fmt.Fprintf(b, "func Assign%s(p *%s, v %[2]s) %[2]s { *p = v; return v }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) = v; return v }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignMul%s(p *%s, v %[2]s) %[2]s { *p *= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignDiv%s(p *%s, v %[2]s) %[2]s { *p /= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignRem%s(p *%s, v %[2]s) %[2]s { *p %%= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignAdd%s(p *%s, v %[2]s) %[2]s { *p += v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignSub%s(p *%s, v %[2]s) %[2]s { *p -= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignAnd%s(p *%s, v %[2]s) %[2]s { *p &= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignXor%s(p *%s, v %[2]s) %[2]s { *p ^= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignOr%s(p *%s, v %[2]s) %[2]s { *p |= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignMulPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) *= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignDivPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) /= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignRemPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) %%= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignAddPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) += v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignSubPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) -= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignAndPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) &= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignXorPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) ^= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignOrPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) |= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignShlPtr%s(p uintptr, v int) %s { *(*%[2]s)(unsafe.Pointer(p)) <<= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignShrPtr%s(p uintptr, v int) %s { *(*%[2]s)(unsafe.Pointer(p)) >>= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignShl%s(p *%s, v int) %[2]s { *p <<= v; return *p }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignShr%s(p *%s, v int) %[2]s { *p >>= v; return *p }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func PreInc%s(p *%s, d %[2]s) %[2]s { *p += d; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func PreDec%s(p *%s, d %[2]s) %[2]s { *p -= d; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func PostInc%s(p *%s, d %[2]s) %[2]s { r := *p; *p += d; return r }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func PostDec%s(p *%s, d %[2]s) %[2]s { r := *p; *p -= d; return r }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		for _, w := range scalar {
			fmt.Fprintf(b, "func %sFrom%s(n %s) %s { return %[4]s(n) }\n", capitalize(v), capitalize(w), w, v)
		}
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func %s(n %s) %[2]s { return n }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func Neg%s(n %s) %[2]s { return -n }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func Cpl%s(n %s) %[2]s { return ^n }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range ints {
		fmt.Fprintf(b, `
func Bool%s(b bool) %s {
	if b {
		return 1
	}
	return 0
}
`, capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, sz := range sizes {
		for _, v := range ints {
			fmt.Fprintf(b, `
func SetBitFieldPtr%s%s(p uintptr, v %s, off int, mask uint%[1]s) {
	*(*uint%[1]s)(unsafe.Pointer(p)) = *(*uint%[1]s)(unsafe.Pointer(p))&^uint%[1]s(mask) | uint%[1]s(v<<off)&mask
}

`, sz, capitalize(v), v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func AssignBitFieldPtr%dInt%d(p uintptr, v int%[2]d, w, off int, mask uint%[1]d) int%[2]d {
	*(*uint%[1]d)(unsafe.Pointer(p)) = *(*uint%[1]d)(unsafe.Pointer(p))&^uint%[1]d(mask) | uint%[1]d(v<<off)&mask
	s := %[2]d - w
	return v << s >> s
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func AssignBitFieldPtr%dUint%d(p uintptr, v uint%[2]d, w, off int, mask uint%[1]d) uint%[2]d {
	*(*uint%[1]d)(unsafe.Pointer(p)) = *(*uint%[1]d)(unsafe.Pointer(p))&^uint%[1]d(mask) | uint%[1]d(v<<off)&mask
	return v & uint%[2]d(mask >> off)
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func PostDecBitFieldPtr%dInt%d(p uintptr, d int%[2]d, w, off int, mask uint%[1]d) (r int%[2]d) {
	x0 := *(*uint%[1]d)(unsafe.Pointer(p))
	s := %[2]d - w
	r = int%[2]d(x0) & int%[2]d(mask) << s >> s
	*(*uint%[1]d)(unsafe.Pointer(p)) = x0&^uint%[1]d(mask) | uint%[1]d(r-d)<<off&mask
	return r
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func PostDecBitFieldPtr%dUint%d(p uintptr, d uint%[2]d, w, off int, mask uint%[1]d) (r uint%[2]d) {
	x0 := *(*uint%[1]d)(unsafe.Pointer(p))
	r = uint%[2]d(x0) & uint%[2]d(mask) >> off
	*(*uint%[1]d)(unsafe.Pointer(p)) = x0&^uint%[1]d(mask) | uint%[1]d(r-d)<<off&mask
	return r
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func PostIncBitFieldPtr%dInt%d(p uintptr, d int%[2]d, w, off int, mask uint%[1]d) (r int%[2]d) {
	x0 := *(*uint%[1]d)(unsafe.Pointer(p))
	s := %[2]d - w
	r = int%[2]d(x0) & int%[2]d(mask) << s >> s
	*(*uint%[1]d)(unsafe.Pointer(p)) = x0&^uint%[1]d(mask) | uint%[1]d(r+d)<<off&mask
	return r
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func PostIncBitFieldPtr%dUint%d(p uintptr, d uint%[2]d, w, off int, mask uint%[1]d) (r uint%[2]d) {
	x0 := *(*uint%[1]d)(unsafe.Pointer(p))
	r = uint%[2]d(x0) & uint%[2]d(mask) >> off
	*(*uint%[1]d)(unsafe.Pointer(p)) = x0&^uint%[1]d(mask) | uint%[1]d(r+d)<<off&mask
	return r
}

`, sz, v)
		}
	}

	b.WriteString("\n")
	if err := ioutil.WriteFile(fmt.Sprintf("ccgo.go"), b.Bytes(), 0660); err != nil {
		fail(err)
	}
}

func capitalize(s string) string { return strings.ToUpper(s[:1]) + s[1:] }
