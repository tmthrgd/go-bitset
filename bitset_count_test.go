// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"testing"
	"testing/quick"
)

func TestCount(t *testing.T) {
	b := New(80)

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

	b.SetAll()

	if b.Count() != b.Len() {
		t.Errorf("invalid count, expected %d, got %d", b.Len(), b.Count())
	}
}

func TestCountRange(t *testing.T) {
	b := New(80)

	if c := b.CountRange(0, b.Len()); c != 0 {
		t.Errorf("invalid count, expected 0, got %d", c)
	}

	b.Set(0)
	b.Set(1)
	b.Set(10)

	if c := b.CountRange(1, 15); c != 2 {
		t.Errorf("invalid count, expected 2, got %d", c)
	}

	b.SetAll()

	if c := b.CountRange(10, 60); c != 50 {
		t.Errorf("invalid count, expected 50, got %d", c)
	}

	if c := b.CountRange(0, b.Len()); c != b.Len() {
		t.Errorf("invalid count, expected %d, got %d", b.Len(), c)
	}

	if err := quick.CheckEqual(func(b, _ Bitset, start, end uint) (count uint) {
		for i := start; i < end; i++ {
			if b.IsSet(i) {
				count++
			}
		}

		return
	}, func(b, _ Bitset, start, end uint) uint {
		return b.CountRange(start, end)
	}, &quick.Config{
		Values:        rangeTestValues2,
		MaxCountScale: 250,
	}); err != nil {
		t.Error(err)
	}
}

func BenchmarkCount(b *testing.B) {
	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			bs := make(Bitset, size.l)

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
	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			bs := make(Bitset, size.l)
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				var _ = bs.CountRange(1, l-1)
			}
		})
	}
}
