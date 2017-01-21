// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

func TestIsSet(t *testing.T) {
	b := New(80)

	for i := uint(0); i < b.Len(); i++ {
		if b.IsSet(i) {
			t.Errorf("IsSet failed, should not have found bit #%d", i)
		}
	}

	b.SetAll()

	for i := uint(0); i < b.Len(); i++ {
		if !b.IsSet(i) {
			t.Errorf("IsSet failed, should have found bit #%d", i)
		}
	}
}

func TestIsClear(t *testing.T) {
	b := New(80)

	for i := uint(0); i < b.Len(); i++ {
		if !b.IsClear(i) {
			t.Errorf("IsClear failed, should not have found bit #%d", i)
		}
	}

	b.SetAll()

	for i := uint(0); i < b.Len(); i++ {
		if b.IsClear(i) {
			t.Errorf("IsClear failed, should have found bit #%d", i)
		}
	}
}

func TestIsRangeSet(t *testing.T) {
	b := New(80)

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
	b := New(80)

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

func TestAll(t *testing.T) {
	b := New(80)

	if b.All() {
		t.Error("All failed")
	}

	b.Set(10)

	if b.All() {
		t.Error("All failed")
	}

	b.SetAll()

	if !b.All() {
		t.Error("All failed")
	}

	b.Clear(17)

	if b.All() {
		t.Error("All failed")
	}
}

func TestNone(t *testing.T) {
	b := New(80)

	if !b.None() {
		t.Error("None failed")
	}

	b.Set(10)

	if b.None() {
		t.Error("None failed")
	}
}

func TestAny(t *testing.T) {
	b := New(80)

	if b.Any() {
		t.Error("Any failed")
	}

	b.Set(10)

	if !b.Any() {
		t.Error("Any failed")
	}
}

func BenchmarkIsSet(b *testing.B) {
	bs := New(80)

	for i := 0; i < b.N; i++ {
		var _ = bs.IsSet(50)
	}
}

func BenchmarkIsRangeSet(b *testing.B) {
	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			bs := make(Bitset, size.l)
			bs.SetAll()
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				if !bs.IsRangeSet(1, l-1) {
					b.Fatal("IsRangeSet failed")
				}
			}
		})
	}
}

func BenchmarkIsRangeClear(b *testing.B) {
	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			bs := make(Bitset, size.l)
			l := bs.Len()

			if size.l > 1024 {
				b.ResetTimer()
			}

			for i := 0; i < b.N; i++ {
				if !bs.IsRangeClear(1, l-1) {
					b.Fatal("IsRangeClear failed")
				}
			}
		})
	}
}
