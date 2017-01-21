// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

var benchSizes = []struct {
	name string
	l    int
}{
	{"16", 16},
	{"32", 32},
	{"128", 128},
	{"1K", 1 * 1024},
	{"16K", 16 * 1024},
	{"128K", 128 * 1024},
	{"1M", 1024 * 1024},
	{"16M", 16 * 1024 * 1024},
	{"128M", 128 * 1024 * 1024},
}

func rangeTestValues(args []reflect.Value, rand *rand.Rand) {
	size := 1 + rand.Intn(4096)

	start := rand.Intn(size)
	end := start + rand.Intn(size-start+1)

	args[0] = reflect.ValueOf(uint(size))
	args[1] = reflect.ValueOf(uint(start))
	args[2] = reflect.ValueOf(uint(end))
}

func rangeTestValues2(args []reflect.Value, rand *rand.Rand) {
	size := 1 + rand.Intn(4096)

	start := rand.Intn(size)
	end := start + rand.Intn(size-start+1)

	b, b1 := New(uint(size)), New(uint(size))
	rand.Read(b)
	rand.Read(b1)

	args[0] = reflect.ValueOf(b)
	args[1] = reflect.ValueOf(b1)
	args[2] = reflect.ValueOf(uint(start))
	args[3] = reflect.ValueOf(uint(end))
}

func TestNew(t *testing.T) {
	for _, v := range []struct {
		size, expected uint
	}{
		{0, 0},
		{1, 8}, {7, 8}, {8, 8},
		{9, 16}, {12, 16}, {15, 16}, {16, 16},
		{100, 104},
	} {
		if l := New(v.size).Len(); l != v.expected {
			t.Errorf("New failed for size %d, expected Len of %d, got %d", v.size, v.expected, l)
		}
	}
}

func TestLen(t *testing.T) {
	b := New(80)

	if b.Len() != 80 {
		t.Errorf("invalid length, expected 80, got %d", b.Len())
	}
}

func TestByteLen(t *testing.T) {
	b := make(Bitset, 10)

	if b.ByteLen() != 10 {
		t.Errorf("invalid length, expected 10, got %d", b.ByteLen())
	}
}

func TestSlice(t *testing.T) {
	b := New(80)
	rand.Read(b)

	if !b.Slice(8, 64).Equal(b[1:8]) {
		t.Error("Subset failed")
	}

	defer func() {
		if recover() == nil {
			t.Error("Slice did not panic for invalid range")
		}
	}()

	b.Slice(7, 63)
}

func TestClone(t *testing.T) {
	b := New(80)

	if !b.Equal(b.Clone()) {
		t.Error("Clone failed")
	}

	b.Set(10)

	if !b.Equal(b.Clone()) {
		t.Error("Clone failed")
	}

	b.Clear(10)

	if !b.Equal(b.Clone()) {
		t.Error("Clone failed")
	}
}

func TestCloneRange(t *testing.T) {
	if err := quick.CheckEqual(func(b, _ Bitset, start, end uint) []byte {
		b1 := New(end - start)

		for i := start; i < end; i++ {
			b1.SetTo(i-start, b.IsSet(i))
		}

		return b1
	}, func(b, _ Bitset, start, end uint) []byte {
		return b.CloneRange(start, end)
	}, &quick.Config{
		Values:        rangeTestValues2,
		MaxCountScale: 100,
	}); err != nil {
		t.Error(err)
	}
}

func TestString(t *testing.T) {
	b := New(80)

	if exp, got := "Bitset{00000000000000000000}", b.String(); exp != got {
		t.Errorf("String failed, expected %s, got %s", exp, got)
	}

	b.Set(0)

	if exp, got := "Bitset{01000000000000000000}", b.String(); exp != got {
		t.Errorf("String failed, expected %s, got %s", exp, got)
	}

	b.SetRange(0, b.Len())

	if exp, got := "Bitset{ffffffffffffffffffff}", b.String(); exp != got {
		t.Errorf("String failed, expected %s, got %s", exp, got)
	}

	b.Clear(b.Len() - 1)

	if exp, got := "Bitset{ffffffffffffffffff7f}", b.String(); exp != got {
		t.Errorf("String failed, expected %s, got %s", exp, got)
	}

	b = make(Bitset, 256)

	x := "0000000000000000000000000000000000000000000000000000000000000000"
	if exp, got := "Bitset{"+x+x+x+x+"...}", b.String(); exp != got {
		t.Errorf("String failed, expected %s, got %s", exp, got)
	}
}

func BenchmarkLen(b *testing.B) {
	bs := New(80)

	for i := 0; i < b.N; i++ {
		var _ = bs.Len()
	}
}

func BenchmarkByteLen(b *testing.B) {
	bs := New(80)

	for i := 0; i < b.N; i++ {
		var _ = bs.ByteLen()
	}
}

func BenchmarkSlice(b *testing.B) {
	bs := New(80)

	for i := 0; i < b.N; i++ {
		var _ = bs.Slice(8, 64)
	}
}

func BenchmarkString(b *testing.B) {
	bs := make(Bitset, 132)

	for i := 0; i < b.N; i++ {
		var _ = bs.String()
	}
}
