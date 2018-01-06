// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package bitset

import (
	"testing"
	"testing/quick"
)

func TestAtomicSet(t *testing.T) {
	b := NewAtomic(192)

	b.Set(50)

	if b[0].Load() != 1<<50 || b[1].Load() != 0 {
		t.Error("Set failed")
	}

	b.Set(70)

	if b[0].Load() != 1<<50 || b[1].Load() != 1<<(70-64) {
		t.Error("Set failed")
	}
}

func TestAtomicClear(t *testing.T) {
	b := NewAtomic(192)

	//b.SetAll()
	for i := uint(0); i < b.Len(); i++ {
		b.Set(i)
	}

	b.Clear(50)

	if b[0].Load() != ^uint64(0)&^(1<<50) || b[1].Load() != ^uint64(0) {
		t.Error("Clear failed")
	}
}

func TestAtomicInvert(t *testing.T) {
	b := NewAtomic(192)

	b.Invert(50)

	if b[0].Load() != 1<<50 || b[1].Load() != 0 {
		t.Error("Invert failed")
	}

	b.Invert(50)

	if b[0].Load() != 0 || b[1].Load() != 0 {
		t.Error("Invert failed")
	}
}

func TestAtomicSetRange(t *testing.T) {
	if err := quick.Check(func(size, start, end uint) bool {
		a := NewAtomic(size)
		a.SetRange(start, end)

		for i := uint(0); i < start; i++ {
			if a.IsSet(i) {
				return false
			}
		}

		for i := start; i < end; i++ {
			if !a.IsSet(i) {
				return false
			}
		}

		for i := end; i < size; i++ {
			if a.IsSet(i) {
				return false
			}
		}

		return true
	}, &quick.Config{
		Values:        rangeTestValues,
		MaxCountScale: 100,
	}); err != nil {
		t.Error(err)
	}
}

func TestAtomicClearRange(t *testing.T) {
	if err := quick.Check(func(size, start, end uint) bool {
		a := NewAtomic(size)
		a.SetRange(0, size)
		a.ClearRange(start, end)

		for i := uint(0); i < start; i++ {
			if !a.IsSet(i) {
				return false
			}
		}

		for i := start; i < end; i++ {
			if a.IsSet(i) {
				return false
			}
		}

		for i := end; i < size; i++ {
			if !a.IsSet(i) {
				return false
			}
		}

		return true
	}, &quick.Config{
		Values:        rangeTestValues,
		MaxCountScale: 100,
	}); err != nil {
		t.Error(err)
	}
}

func BenchmarkAtomicSet(b *testing.B) {
	bs := NewAtomic(192)

	for i := 0; i < b.N; i++ {
		bs.Set(50)
	}
}

func BenchmarkAtomicClear(b *testing.B) {
	bs := NewAtomic(192)

	for i := 0; i < b.N; i++ {
		bs.Clear(50)
	}
}

func BenchmarkAtomicSetRange(b *testing.B) {
	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			bs := NewAtomic(uint(size.l) * 8)
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

func BenchmarkAtomicClearRange(b *testing.B) {
	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			bs := NewAtomic(uint(size.l) * 8)
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
