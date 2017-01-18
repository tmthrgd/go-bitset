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

func testMask2(end uint) (mask byte) {
	for start := end &^ 7; start < end; start++ {
		mask |= 1 << (start & 7)
	}

	return
}

func TestMask1(t *testing.T) {
	if err := quick.CheckEqual(testMask1, mask1, &quick.Config{
		Values: func(args []reflect.Value, rand *rand.Rand) {
			start := uint(rand.Int())
			end := start + uint(rand.Intn(int(^uint(0)>>1)-int(start)))

			args[0] = reflect.ValueOf(start)
			args[1] = reflect.ValueOf(end)
		},

		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}
}

func TestMask2(t *testing.T) {
	if err := quick.CheckEqual(testMask2, mask2, &quick.Config{
		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}
}
