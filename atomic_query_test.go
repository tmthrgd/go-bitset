// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package bitset

import "testing"

func TestAtomicIsSet(t *testing.T) {
	b := NewAtomic(192)

	for i := uint(0); i < b.Len(); i++ {
		if b.IsSet(i) {
			t.Errorf("IsSet failed, should not have found bit #%d", i)
		}

		b.Set(i)
	}

	//b.SetAll()

	for i := uint(0); i < b.Len(); i++ {
		if !b.IsSet(i) {
			t.Errorf("IsSet failed, should have found bit #%d", i)
		}
	}
}

func TestAtomicIsClear(t *testing.T) {
	b := NewAtomic(192)

	for i := uint(0); i < b.Len(); i++ {
		if !b.IsClear(i) {
			t.Errorf("IsClear failed, should not have found bit #%d", i)
		}

		b.Set(i)
	}

	//b.SetAll()

	for i := uint(0); i < b.Len(); i++ {
		if b.IsClear(i) {
			t.Errorf("IsClear failed, should have found bit #%d", i)
		}
	}
}

func BenchmarkAtomicIsSet(b *testing.B) {
	bs := NewAtomic(192)

	for i := 0; i < b.N; i++ {
		var _ = bs.IsSet(50)
	}
}
