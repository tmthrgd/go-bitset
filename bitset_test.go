// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"testing"
)

func isZero(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}

	return true
}

func isOne(b []byte) bool {
	for _, v := range b {
		if v != 0xff {
			return false
		}
	}

	return true
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
	b := Bitset(make([]byte, 10))

	if b.Len() != 80 {
		t.Errorf("invalid length, expected 80, got %d", b.Len())
	}
}

func TestCount(t *testing.T) {
	b := Bitset(make([]byte, 10))

	if b.Count() != 0 {
		t.Errorf("invalid count, expected 0, got %d", b.Count())
	}

	b.Set(0)

	if b.Count() != 1 {
		t.Errorf("invalid count, expected 1, got %d", b.Count())
	}

	b.Set(1)

	if b.Count() != 2 {
		t.Errorf("invalid count, expected 2, got %d", b.Count())
	}

	b.Set(10)

	for i := range b {
		b[i] = 0xff
	}

	if b.Count() != b.Len() {
		t.Errorf("invalid count, expected %d, got %d", b.Len(), b.Count())
	}
}

func TestCountRange(t *testing.T) {
	b := Bitset(make([]byte, 10))

	if c := b.CountRange(0, b.Len()); c != 0 {
		t.Errorf("invalid count, expected 0, got %d", c)
	}

	b.Set(0)
	b.Set(1)
	b.Set(10)

	if c := b.CountRange(1, 15); c != 2 {
		t.Errorf("invalid count, expected 2, got %d", c)
	}

	for i := range b {
		b[i] = 0xff
	}

	if c := b.CountRange(10, 60); c != 50 {
		t.Errorf("invalid count, expected 50, got %d", c)
	}

	if c := b.CountRange(0, b.Len()); c != b.Len() {
		t.Errorf("invalid count, expected %d, got %d", b.Len(), c)
	}
}

func TestSet(t *testing.T) {
	b := Bitset(make([]byte, 10))

	b.Set(50)

	if !isZero(b[:6]) || b[6] != 0x04 || !isZero(b[7:]) {
		t.Error("Set failed")
	}

	b.Set(60)

	if !isZero(b[:6]) || b[6] != 0x04 || b[7] != 0x10 || !isZero(b[8:]) {
		t.Error("Set failed")
	}
}

func TestClear(t *testing.T) {
	b := Bitset(make([]byte, 10))

	for i := range b {
		b[i] = 0xff
	}

	b.Clear(50)

	if !isOne(b[:6]) || b[6] != 0xfb || !isOne(b[7:]) {
		t.Error("Clear failed")
	}
}

func TestInvert(t *testing.T) {
	b := Bitset(make([]byte, 10))

	b.Invert(50)

	if !isZero(b[:6]) || b[6] != 0x04 || !isZero(b[7:]) {
		t.Error("Invert failed")
	}

	b.Invert(50)

	if !isZero(b) {
		t.Error("Invert failed")
	}
}

func TestSetRange(t *testing.T) {
	b := Bitset(make([]byte, 10))
	b1 := Bitset(make([]byte, len(b)))

	b.SetRange(60, 70)

	for i := uint(60); i < 70; i++ {
		b1.Set(i)
	}

	if !b.Equal(b1) {
		t.Error("SetRange failed")
	}
}

func TestClearRange(t *testing.T) {
	b := Bitset(make([]byte, 10))
	b1 := Bitset(make([]byte, len(b)))

	for i := range b {
		b[i] = 0xff
		b1[i] = 0xff
	}

	b.ClearRange(60, 70)

	for i := uint(60); i < 70; i++ {
		b1.Clear(i)
	}

	if !b.Equal(b1) {
		t.Error("ClearRange failed")
	}
}

func TestInvertRange(t *testing.T) {
	b := Bitset(make([]byte, 10))
	b1 := Bitset(make([]byte, len(b)))

	b.SetRange(63, 67)
	b1.SetRange(63, 67)

	b.InvertRange(60, 70)

	for i := uint(60); i < 70; i++ {
		b1.Invert(i)
	}

	if !b.Equal(b1) {
		t.Error("InvertRange failed")
	}
}

func TestIsSet(t *testing.T) {
	b := Bitset(make([]byte, 10))

	for i := uint(0); i < b.Len(); i++ {
		if b.IsSet(i) {
			t.Errorf("IsSet failed, should not have found bit #%d", i)
		}
	}

	for i := range b {
		b[i] = 0xff
	}

	for i := uint(0); i < b.Len(); i++ {
		if !b.IsSet(i) {
			t.Errorf("IsSet failed, should have found bit #%d", i)
		}
	}
}

func TestIsRangeSet(t *testing.T) {
	b := Bitset(make([]byte, 10))

	if b.IsRangeSet(0, b.Len()) {
		t.Errorf("IsRangeSet failed, should not have found range #0-#%d", b.Len())
	}

	b.SetRange(40, 55)

	if !b.IsRangeSet(40, 55) {
		t.Error("IsRangeSet failed, should have found range #40-#55")
	}

	if b.IsRangeSet(40, 56) {
		t.Error("IsRangeSet failed, should not have found range #40-#56")
	}

	if b.IsRangeSet(0, b.Len()) {
		t.Errorf("IsRangeSet failed, should not have found range #0-#%d", b.Len())
	}
}

func TestIsRangeClear(t *testing.T) {
	b := Bitset(make([]byte, 10))

	if !b.IsRangeClear(0, b.Len()) {
		t.Errorf("IsRangeClear, should not have found in range #0-#%d", b.Len())
	}

	b.SetRange(40, 55)

	if b.IsRangeClear(40, 55) {
		t.Error("IsRangeClear failed, should have found in range #40-#55")
	}

	if b.IsRangeClear(35, 50) {
		t.Error("IsRangeClear failed, should have found in range #35-#50")
	}

	if b.IsRangeClear(0, b.Len()) {
		t.Errorf("IsRangeClear failed, should have found in range #0-#%d", b.Len())
	}
}

func TestComplement(t *testing.T) {
	b := Bitset(make([]byte, 10))
	b.Complement(b)

	if !b.IsRangeSet(0, b.Len()) {
		t.Errorf("Not failed, should have found range #0-#%d", b.Len())
	}

	b.Complement(b)

	if !b.IsRangeClear(0, b.Len()) {
		t.Errorf("Not failed, should not have found in range #0-#%d", b.Len())
	}
}

func TestEqual(t *testing.T) {
	b := Bitset(make([]byte, 10))
	b1 := Bitset(make([]byte, len(b)))

	if !b.Equal(b1) {
		t.Error("Equal failed")
	}

	b.Set(10)

	if b.Equal(b1) {
		t.Error("Equal failed")
	}

	b.Clear(10)

	if !b.Equal(b1) {
		t.Error("Equal failed")
	}
}

func TestAll(t *testing.T) {
	b := Bitset(make([]byte, 10))

	if b.All() {
		t.Error("All failed")
	}

	b.Set(10)

	if b.All() {
		t.Error("All failed")
	}

	for i := range b {
		b[i] = 0xff
	}

	if !b.All() {
		t.Error("All failed")
	}

	b.Clear(17)

	if b.All() {
		t.Error("All failed")
	}
}

func TestNone(t *testing.T) {
	b := Bitset(make([]byte, 10))

	if !b.None() {
		t.Error("None failed")
	}

	b.Set(10)

	if b.None() {
		t.Error("None failed")
	}
}

func TestAny(t *testing.T) {
	b := Bitset(make([]byte, 10))

	if b.Any() {
		t.Error("Any failed")
	}

	b.Set(10)

	if !b.Any() {
		t.Error("Any failed")
	}
}

func TestClone(t *testing.T) {
	b := Bitset(make([]byte, 10))

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

func TestCopy(t *testing.T) {
	b := Bitset(make([]byte, 10))
	b1 := b.Clone()

	b.Set(10)
	b1.Copy(b)

	if !b.Equal(b1) {
		t.Error("Copy failed")
	}

	b.Clear(10)
	b1.Copy(b)

	if !b.Equal(b1) {
		t.Error("Copy failed")
	}
}

func TestString(t *testing.T) {
	b := Bitset(make([]byte, 10))

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

	b = Bitset(make([]byte, 256))

	x := "0000000000000000000000000000000000000000000000000000000000000000"
	if exp, got := "Bitset{"+x+x+x+x+"...}", b.String(); exp != got {
		t.Errorf("String failed, expected %s, got %s", exp, got)
	}
}

func BenchmarkSet(b *testing.B) {
	bs := Bitset(make([]byte, 10))

	for i := 0; i < b.N; i++ {
		bs.Set(50)
	}
}

func BenchmarkClear(b *testing.B) {
	bs := Bitset(make([]byte, 10))

	for i := 0; i < b.N; i++ {
		bs.Clear(50)
	}
}

func BenchmarkIsSet(b *testing.B) {
	bs := Bitset(make([]byte, 10))

	for i := 0; i < b.N; i++ {
		var _ = bs.IsSet(50)
	}
}

func BenchmarkLen(b *testing.B) {
	bs := Bitset(make([]byte, 10))

	for i := 0; i < b.N; i++ {
		var _ = bs.Len()
	}
}

var sizes = []struct {
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

func BenchmarkCount(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			bs := Bitset(make([]byte, size.l))

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				var _ = bs.Count()
			}
		})
	}
}

func BenchmarkCountRange(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			bs := Bitset(make([]byte, size.l))
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				var _ = bs.CountRange(0, l)
			}
		})
	}
}

func BenchmarkSetRange(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			bs := Bitset(make([]byte, size.l))
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				bs.SetRange(0, l)
			}
		})
	}
}

func BenchmarkClearRange(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			bs := Bitset(make([]byte, size.l))
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				bs.ClearRange(0, l)
			}
		})
	}
}

func BenchmarkIsRangeSet(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			bs := Bitset(make([]byte, size.l))
			bs.Complement(bs)
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				if !bs.IsRangeSet(0, l) {
					b.Fatal("IsRangeSet failed")
				}
			}
		})
	}
}

func BenchmarkIsRangeClear(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			bs := Bitset(make([]byte, size.l))
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				if !bs.IsRangeClear(0, l) {
					b.Fatal("IsRangeClear failed")
				}
			}
		})
	}
}
