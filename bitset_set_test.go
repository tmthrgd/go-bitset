// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

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

func TestSet(t *testing.T) {
	b := make(Bitset, 10)

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
	b := make(Bitset, 10)

	for i := range b {
		b[i] = 0xff
	}

	b.Clear(50)

	if !isOne(b[:6]) || b[6] != 0xfb || !isOne(b[7:]) {
		t.Error("Clear failed")
	}
}

func TestInvert(t *testing.T) {
	b := make(Bitset, 10)

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
	b := make(Bitset, 10)
	b1 := make(Bitset, len(b))

	b.SetRange(60, 70)

	for i := uint(60); i < 70; i++ {
		b1.Set(i)
	}

	if !b.Equal(b1) {
		t.Error("SetRange failed")
	}
}

func TestClearRange(t *testing.T) {
	b := make(Bitset, 10)
	b1 := make(Bitset, len(b))

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
	b := make(Bitset, 10)
	b1 := make(Bitset, len(b))

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

func BenchmarkSet(b *testing.B) {
	bs := make(Bitset, 10)

	for i := 0; i < b.N; i++ {
		bs.Set(50)
	}
}

func BenchmarkClear(b *testing.B) {
	bs := make(Bitset, 10)

	for i := 0; i < b.N; i++ {
		bs.Clear(50)
	}
}

func BenchmarkSetRange(b *testing.B) {
	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			bs := make(Bitset, size.l)
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				bs.SetRange(1, l-1)
			}
		})
	}
}

func BenchmarkClearRange(b *testing.B) {
	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			bs := make(Bitset, size.l)
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				bs.ClearRange(1, l-1)
			}
		})
	}
}
