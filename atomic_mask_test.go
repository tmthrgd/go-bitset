// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"math/rand"
	"testing"
	"testing/quick"
)

func testAtomicMask1(start, end uint) (mask uint64) {
	for ; start&63 != 0 && start < end; start++ {
		mask |= 1 << (start & 63)
	}

	return
}

func testAtomicMask2(start, end uint) (mask uint64) {
	if start > end&^63 {
		return 0
	}

	for start := end &^ 63; start < end; start++ {
		mask |= 1 << (start & 63)
	}

	return
}

func TestAtomicMask1(t *testing.T) {
	for i := uint(0); i < 0x2000; i++ {
		for j := i; j < 0x2000; j++ {
			if x, y := testAtomicMask1(i, j), atomicMask1(i, j); x != y {
				t.Fatalf("testAtomicMask1(%[1]d, %[2]d) = 0x%02x != 0x%02x = atomicMask1(%[1]d, %[2]d)", i, j, x, y)
			}
		}
	}

	if err := quick.CheckEqual(testAtomicMask1, atomicMask1, &quick.Config{
		Values:        maskTestValues,
		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}
}

func TestAtomicMask2(t *testing.T) {
	for i := uint(0); i < 0x2000; i++ {
		for j := i; j < 0x2000; j++ {
			if x, y := testAtomicMask2(i, j), atomicMask2(i, j); x != y {
				t.Fatalf("testAtomicMask2(%[1]d, %[2]d) = 0x%02x != 0x%02x = atomicMask2(%[1]d, %[2]d)", i, j, x, y)
			}
		}
	}

	if err := quick.CheckEqual(testAtomicMask2, atomicMask2, &quick.Config{
		Values:        maskTestValues,
		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}
}

func BenchmarkAtomicMask1(b *testing.B) {
	start := uint(rand.Int())
	end := start + uint(rand.Intn(int(^uint(0)>>1)-int(start)))

	for i := 0; i < b.N; i++ {
		var _ = atomicMask1(start, end)
	}
}

func BenchmarkAtomicMask2(b *testing.B) {
	start := uint(rand.Int())
	end := start + uint(rand.Intn(int(^uint(0)>>1)-int(start)))

	for i := 0; i < b.N; i++ {
		var _ = atomicMask2(start, end)
	}
}
