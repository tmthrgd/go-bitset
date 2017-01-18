// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitwise

import (
	"bytes"
	"math/rand"
	"testing"
	"testing/quick"

	"github.com/tmthrgd/go-bitwise"
	"github.com/tmthrgd/go-memset"
)

type testVector struct {
	res bool
	a   []byte
	b   []byte
}

var andEqTestVectors = []testVector{
	{
		true,
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		false,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		true,
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		false,
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
	},
	{
		true,
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
}

func testAndEqBytes(a, b []byte) bool {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}

	for i := 0; i < n; i++ {
		if a[i] & b[i] != b[i] {
			return false
		}
	}

	return true
}

func andEqBytesOther(a, b []byte) bool {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}

	// a & b == b
	//  -> c = a & b; c == b

	c := make([]byte, n)
	bitwise.And(c, a, b)
	return bytes.Equal(c, b[:n])
}

func TestAndEq(t *testing.T) {
	for i, vector := range andEqTestVectors {
		if AndEq(vector.a, vector.b) != vector.res {
			t.Errorf("test case #%d failed, expected %t, got %t", i, vector.res, !vector.res)
		}
	}

	for alignP := 0; alignP < 2; alignP++ {
		for alignQ := 0; alignQ < 2; alignQ++ {
			for alignD := 0; alignD < 2; alignD++ {
				p := make([]byte, 1024)[alignP:]
				rand.Read(p)

				q := make([]byte, 1024)[alignQ:]
				rand.Read(q)

				if AndEq(p, q) != testAndEqBytes(p, q) {
					t.Error("not equal")
				}
			}
		}
	}

	if err := quick.CheckEqual(AndEq, testAndEqBytes, &quick.Config{
		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}

	if err := quick.CheckEqual(AndEq, andEqBytesOther, &quick.Config{
		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}
}

var benchSizes = []struct {
	name string
	l    int
}{
	{"15", 15},
	{"32", 32},
	{"128", 128},
	{"1K", 1 * 1024},
	{"16K", 16 * 1024},
	{"128K", 128 * 1024},
	{"1M", 1024 * 1024},
	{"16M", 16 * 1024 * 1024},
	{"128M", 128 * 1024 * 1024},
}

func benchmarkThree(b *testing.B, testFn func(a, b []byte) bool) {
	maxSize := benchSizes[len(benchSizes)-1]

	p, q := make([]byte, maxSize.l), make([]byte, maxSize.l)
	memset.Memset(p, 0xff)
	rand.Read(q)

	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			b.SetBytes(int64(size.l))

			p, q := p[:size.l], q[:size.l]

			for i := 0; i < b.N; i++ {
				testFn(p, q)
			}
		})
	}
}

func BenchmarkAndEq(b *testing.B) {
	benchmarkThree(b, AndEq)
}

func BenchmarkAndEqGo(b *testing.B) {
	benchmarkThree(b, testAndEqBytes)
}

func BenchmarkAndEqOther(b *testing.B) {
	benchmarkThree(b, andEqBytesOther)
}
