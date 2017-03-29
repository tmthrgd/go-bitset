// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package bitset

import (
	"testing"
)

func TestSet(t *testing.T) {
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

func TestClear(t *testing.T) {
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

func TestInvert(t *testing.T) {
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

func BenchmarkSet(b *testing.B) {
	bs := NewAtomic(192)

	for i := 0; i < b.N; i++ {
		bs.Set(50)
	}
}

func BenchmarkClear(b *testing.B) {
	bs := NewAtomic(192)

	for i := 0; i < b.N; i++ {
		bs.Clear(50)
	}
}
