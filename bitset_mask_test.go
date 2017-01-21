// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func testMask1(start, end uint) (mask byte) {
	for ; start&7 != 0 && start < end; start++ {
		mask |= 1 << (start & 7)
	}

	return
}

func testMask2(start, end uint) (mask byte) {
	if start > end&^7 {
		return 0
	}

	for start := end &^ 7; start < end; start++ {
		mask |= 1 << (start & 7)
	}

	return
}

func maskTestValues(args []reflect.Value, rand *rand.Rand) {
	start := uint(rand.Int())
	end := start + uint(rand.Intn(int(^uint(0)>>1)-int(start)))

	args[0] = reflect.ValueOf(start)
	args[1] = reflect.ValueOf(end)
}

func TestMask1(t *testing.T) {
	for i := uint(0); i < 0x2000; i++ {
		for j := i; j < 0x2000; j++ {
			if x, y := testMask1(i, j), mask1(i, j); x != y {
				t.Fatalf("testMask1(%[1]d, %[2]d) = 0x%02x != 0x%02x = mask1(%[1]d, %[2]d)", i, j, x, y)
			}
		}
	}

	if err := quick.CheckEqual(testMask1, mask1, &quick.Config{
		Values:        maskTestValues,
		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}
}

func TestMask2(t *testing.T) {
	for i := uint(0); i < 0x2000; i++ {
		for j := i; j < 0x2000; j++ {
			if x, y := testMask2(i, j), mask2(i, j); x != y {
				t.Fatalf("testMask2(%[1]d, %[2]d) = 0x%02x != 0x%02x = mask2(%[1]d, %[2]d)", i, j, x, y)
			}
		}
	}

	if err := quick.CheckEqual(testMask2, mask2, &quick.Config{
		Values:        maskTestValues,
		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}
}

func BenchmarkMask1(b *testing.B) {
	start := uint(rand.Int())
	end := start + uint(rand.Intn(int(^uint(0)>>1)-int(start)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var _ = mask1(start, end)
	}
}

func BenchmarkMask2(b *testing.B) {
	start := uint(rand.Int())
	end := start + uint(rand.Intn(int(^uint(0)>>1)-int(start)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var _ = mask2(start, end)
	}
}
