// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

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
	b := make(Bitset, 10)

	if b.Len() != 80 {
		t.Errorf("invalid length, expected 80, got %d", b.Len())
	}
}

func TestClone(t *testing.T) {
	b := make(Bitset, 10)

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

func TestString(t *testing.T) {
	b := make(Bitset, 10)

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
	bs := make(Bitset, 10)

	for i := 0; i < b.N; i++ {
		var _ = bs.Len()
	}
}
