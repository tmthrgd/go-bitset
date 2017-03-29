// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package bitset

import "testing"

func TestNew(t *testing.T) {
	for _, v := range []struct {
		size, expected uint
	}{
		{0, 0},
		{1, 64}, {8, 64}, {16, 64}, {63, 64}, {64, 64},
		{65, 128}, {100, 128},
	} {
		if l := NewAtomic(v.size).Len(); l != v.expected {
			t.Errorf("NewAtomic failed for size %d, expected Len of %d, got %d", v.size, v.expected, l)
		}
	}
}

func TestLen(t *testing.T) {
	b := NewAtomic(192)

	if b.Len() != 192 {
		t.Errorf("invalid length, expected 80, got %d", b.Len())
	}
}

func TestUint64Len(t *testing.T) {
	b := make(Atomic, 10)

	if b.Uint64Len() != 10 {
		t.Errorf("invalid length, expected 10, got %d", b.Uint64Len())
	}
}

func TestSlice(t *testing.T) {
	b := NewAtomic(192)

	defer func() {
		if recover() == nil {
			t.Error("Slice did not panic for invalid range")
		}
	}()

	b.Slice(63, 127)
}

func BenchmarkLen(b *testing.B) {
	bs := NewAtomic(192)

	for i := 0; i < b.N; i++ {
		var _ = bs.Len()
	}
}

func BenchmarkUint64Len(b *testing.B) {
	bs := NewAtomic(192)

	for i := 0; i < b.N; i++ {
		var _ = bs.Uint64Len()
	}
}

func BenchmarkSlice(b *testing.B) {
	bs := NewAtomic(192)

	for i := 0; i < b.N; i++ {
		var _ = bs.Slice(64, 128)
	}
}

func BenchmarkString(b *testing.B) {
	var bs Atomic

	for i := 0; i < b.N; i++ {
		var _ = bs.String()
	}
}
