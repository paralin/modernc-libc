// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"math"
	"os"
	"testing"
	"unsafe"

	"modernc.org/mathutil"
)

const (
	pageSize    = 1 << pageSizeLog
	pageSizeLog = 20
	quota       = 128 << 20
)

var (
	bigMax     = 2 * pageSize
	max        = 2 * osPageSize
	osPageSize = os.Getpagesize()
)

type block struct {
	p    uintptr
	size int
}

func TestAllocator(t *testing.T) {
	t.Run("small1", func(t *testing.T) { testAllocator1(t, max) })
	t.Run("big1", func(t *testing.T) { testAllocator1(t, bigMax) })
	t.Run("small2", func(t *testing.T) { testAllocator2(t, max) })
	t.Run("big2", func(t *testing.T) { testAllocator2(t, bigMax) })
	t.Run("small3", func(t *testing.T) { testAllocator3(t, max) })
	//TODO fail in v0.6.0 t.Run("big3", func(t *testing.T) { testAllocator3(t, bigMax) })
}

func testAllocator1(t *testing.T, max int) {
	tls := NewTLS()

	defer tls.Close()

	rem := quota
	var a []block
	srng, err := mathutil.NewFC32(0, math.MaxInt32, true)
	if err != nil {
		t.Fatal(err)
	}

	vrng, err := mathutil.NewFC32(0, math.MaxInt32, true)
	if err != nil {
		t.Fatal(err)
	}

	// Allocate
	for rem > 0 {
		size := srng.Next()%max + 1
		rem -= size
		p := Xmalloc(tls, uint64(size))
		if p == 0 {
			t.Fatal("Xmalloc failed")
		}

		a = append(a, block{p, size})
		for i := 0; i < size; i++ {
			*(*byte)(unsafe.Pointer(p + uintptr(i))) = byte(vrng.Next())
		}
	}
	srng.Seek(0)
	vrng.Seek(0)
	// Verify
	for i, b := range a {
		if g, e := b.size, srng.Next()%max+1; g != e {
			t.Fatal(i, g, e)
		}

		//TODO N/A v0.6.0
		//TODO if a, b := b.size, UintptrUsableSize(b.p); a > b {
		//TODO 	t.Fatal(i, a, b)
		//TODO }

		for j := 0; j < b.size; j++ {
			g := *(*byte)(unsafe.Pointer(b.p + uintptr(j)))
			if e := byte(vrng.Next()); g != e {
				t.Fatalf("%v,%v %#x: %#02x %#02x", i, j, b.p+uintptr(j), g, e)
			}

			*(*byte)(unsafe.Pointer(b.p + uintptr(j))) = 0
		}
	}
	// Shuffle
	for i := range a {
		j := srng.Next() % len(a)
		a[i], a[j] = a[j], a[i]
	}
	// Free
	for _, b := range a {
		Xfree(tls, b.p)
	}
}

func testAllocator2(t *testing.T, max int) {
	tls := NewTLS()

	defer tls.Close()

	rem := quota
	var a []block
	srng, err := mathutil.NewFC32(0, math.MaxInt32, true)
	if err != nil {
		t.Fatal(err)
	}

	vrng, err := mathutil.NewFC32(0, math.MaxInt32, true)
	if err != nil {
		t.Fatal(err)
	}

	// Allocate
	for rem > 0 {
		size := srng.Next()%max + 1
		rem -= size
		p := Xmalloc(tls, uint64(size))
		if p == 0 {
			t.Fatal("Xmalloc failed")
		}

		a = append(a, block{p, size})
		for i := 0; i < size; i++ {
			*(*byte)(unsafe.Pointer(p + uintptr(i))) = byte(vrng.Next())
		}
	}
	srng.Seek(0)
	vrng.Seek(0)
	// Verify & free
	for i, b := range a {
		if g, e := b.size, srng.Next()%max+1; g != e {
			t.Fatal(i, g, e)
		}

		//TODO N/A v0.6.0
		//TODO if a, b := b.size, UintptrUsableSize(b.p); a > b {
		//TODO 	t.Fatal(i, a, b)
		//TODO }

		for j := 0; j < b.size; j++ {
			g := *(*byte)(unsafe.Pointer(b.p + uintptr(j)))
			if e := byte(vrng.Next()); g != e {
				t.Fatalf("%v,%v %#x: %#02x %#02x", i, j, b.p+uintptr(j), g, e)
			}

			*(*byte)(unsafe.Pointer(b.p + uintptr(j))) = 0
		}
		Xfree(tls, b.p)
	}
}

func testAllocator3(t *testing.T, max int) {
	tls := NewTLS()

	defer tls.Close()

	rem := quota
	m := map[block][]byte{}
	srng, err := mathutil.NewFC32(1, max, true)
	if err != nil {
		t.Fatal(err)
	}

	vrng, err := mathutil.NewFC32(1, max, true)
	if err != nil {
		t.Fatal(err)
	}

	for rem > 0 {
		switch srng.Next() % 3 {
		case 0, 1: // 2/3 allocate
			size := srng.Next()
			rem -= size
			p := Xmalloc(tls, uint64(size))
			if p == 0 {
				t.Fatal("Xmalloc failed")
			}

			b := make([]byte, size)
			for i := range b {
				b[i] = byte(vrng.Next())
				*(*byte)(unsafe.Pointer(p + uintptr(i))) = b[i]
			}
			m[block{p, size}] = append([]byte(nil), b...)
		default: // 1/3 free
			for b, v := range m {
				for i, v := range v {
					if *(*byte)(unsafe.Pointer(b.p + uintptr(i))) != v {
						t.Fatal("corrupted heap")
					}
				}

				//TODO N/A v0.6.0
				//TODO if a, b := b.size, UintptrUsableSize(b.p); a > b {
				//TODO 	t.Fatal(a, b)
				//TODO }

				for j := 0; j < b.size; j++ {
					*(*byte)(unsafe.Pointer(b.p + uintptr(j))) = 0
				}
				rem += b.size
				Xfree(tls, b.p)
				delete(m, b)
				break
			}
		}
	}
	for b, v := range m {
		for i, v := range v {
			if *(*byte)(unsafe.Pointer(b.p + uintptr(i))) != v {
				t.Fatal("corrupted heap")
			}
		}

		//TODO N/A v0.6.0
		//TODO if a, b := b.size, UintptrUsableSize(b.p); a > b {
		//TODO 	t.Fatal(a, b)
		//TODO }

		for j := 0; j < b.size; j++ {
			*(*byte)(unsafe.Pointer(b.p + uintptr(j))) = 0
		}
		Xfree(tls, b.p)
	}
}

func TestAllocatorFree(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	p := Xmalloc(tls, 1)
	if p == 0 {
		t.Fatal("Xmalloc failed")
	}

	Xfree(tls, p)
}

func benchmarkFree(b *testing.B, size int) {
	tls := NewTLS()

	defer tls.Close()

	a := make([]uintptr, b.N)
	for i := range a {
		p := Xmalloc(tls, uint64(size))
		if p == 0 {
			b.Fatal("Xmalloc failed")
		}

		a[i] = p
	}
	b.ResetTimer()
	for _, p := range a {
		Xfree(tls, p)
	}
	b.StopTimer()
}

func BenchmarkFree(b *testing.B) {
	b.Run("16", func(b *testing.B) { benchmarkFree(b, 1<<4) })
	b.Run("32", func(b *testing.B) { benchmarkFree(b, 1<<5) })
	b.Run("64", func(b *testing.B) { benchmarkFree(b, 1<<6) })
}

func benchmarkCalloc(b *testing.B, size int) {
	tls := NewTLS()

	defer tls.Close()

	a := make([]uintptr, b.N)
	b.ResetTimer()
	for i := range a {
		p := Xcalloc(tls, 1, uint64(size))
		if p == 0 {
			b.Fatal("Xcalloc failed")
		}

		a[i] = p
	}
	b.StopTimer()
	for _, p := range a {
		Xfree(tls, p)
	}
}

func BenchmarkCalloc(b *testing.B) {
	b.Run("16", func(b *testing.B) { benchmarkCalloc(b, 1<<4) })
	b.Run("32", func(b *testing.B) { benchmarkCalloc(b, 1<<5) })
	b.Run("64", func(b *testing.B) { benchmarkCalloc(b, 1<<6) })
}

func benchmarkMalloc(b *testing.B, size int) {
	tls := NewTLS()

	defer tls.Close()

	a := make([]uintptr, b.N)
	b.ResetTimer()
	for i := range a {
		p := Xmalloc(tls, uint64(size))
		if p == 0 {
			b.Fatal("Xmalloc failed")
		}

		a[i] = p
	}
	b.StopTimer()
	for _, p := range a {
		Xfree(tls, p)
	}
}

func BenchmarkMalloc(b *testing.B) {
	b.Run("16", func(b *testing.B) { benchmarkMalloc(b, 1<<4) })
	b.Run("32", func(b *testing.B) { benchmarkMalloc(b, 1<<5) })
	b.Run("64", func(b *testing.B) { benchmarkMalloc(b, 1<<6) })
}