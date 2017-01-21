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

func rangeTestShiftValues(args []reflect.Value, rand *rand.Rand) {
	size := 8 + rand.Intn(4096-8)
	shift := 8 + rand.Intn(size>>3)<<3
	size2 := shift + rand.Intn(4096-shift)

	b := New(uint(size))
	rand.Read(b)

	args[0] = reflect.ValueOf(b)
	args[1] = reflect.ValueOf(uint(size2))
	args[2] = reflect.ValueOf(uint(shift))
}

func TestShiftLeft(t *testing.T) {
	b := make(Bitset, 10)
	b.SetRange(50, 70)

	b.ShiftLeft(b, 10)

	if !b.IsRangeClear(0, 40) {
		t.Fatal("ShiftLeft failed")
	}

	if !b.IsRangeSet(40, 60) {
		t.Fatal("ShiftLeft failed")
	}

	if !b.IsRangeClear(60, b.Len()) {
		t.Fatal("ShiftLeft failed")
	}
}

func TestShiftLeftFastPath(t *testing.T) {
	b := make(Bitset, 10)
	b.SetRange(48, 68)

	b.ShiftLeft(b, 8)

	if !b.IsRangeClear(0, 40) {
		t.Fatal("ShiftLeft failed")
	}

	if !b.IsRangeSet(40, 60) {
		t.Fatal("ShiftLeft failed")
	}

	if !b.IsRangeClear(60, b.Len()) {
		t.Fatal("ShiftLeft failed")
	}
}

func TestShiftLeftByZero(t *testing.T) {
	b := make(Bitset, 10)
	b.SetRange(40, 60)

	b.ShiftLeft(b, 0)

	if !b.IsRangeClear(0, 40) {
		t.Fatal("ShiftLeft failed")
	}

	if !b.IsRangeSet(40, 60) {
		t.Fatal("ShiftLeft failed")
	}

	if !b.IsRangeClear(60, b.Len()) {
		t.Fatal("ShiftLeft failed")
	}
}

func TestShiftLeft2(t *testing.T) {
	useShiftFastPath = false
	defer func() {
		useShiftFastPath = true
	}()

	if err := quick.CheckEqual(func(b Bitset, size, shift uint) []byte {
		b1 := New(size)
		copy(b1, b[shift>>3:])
		return b1
	}, func(b Bitset, size, shift uint) []byte {
		b1 := New(size)
		b1.ShiftLeft(b, shift)
		return b1
	}, &quick.Config{
		Values:        rangeTestShiftValues,
		MaxCountScale: 100,
	}); err != nil {
		t.Error(err)
	}
}

func TestShiftRight(t *testing.T) {
	b := make(Bitset, 10)
	b.SetRange(30, 50)

	b.ShiftRight(b, 10)

	if !b.IsRangeClear(0, 40) {
		t.Fatal("ShiftRight failed")
	}

	if !b.IsRangeSet(40, 60) {
		t.Fatal("ShiftRight failed")
	}

	if !b.IsRangeClear(60, b.Len()) {
		t.Fatal("ShiftRight failed")
	}
}

func TestShiftRightFastPath(t *testing.T) {
	b := make(Bitset, 10)
	b.SetRange(32, 52)

	b.ShiftRight(b, 8)

	if !b.IsRangeClear(0, 40) {
		t.Fatal("ShiftRight failed")
	}

	if !b.IsRangeSet(40, 60) {
		t.Fatal("ShiftRight failed")
	}

	if !b.IsRangeClear(60, b.Len()) {
		t.Fatal("ShiftRight failed")
	}
}

func TestShiftRightByZero(t *testing.T) {
	b := make(Bitset, 10)
	b.SetRange(40, 60)

	b.ShiftRight(b, 0)

	if !b.IsRangeClear(0, 40) {
		t.Fatal("ShiftRight failed")
	}

	if !b.IsRangeSet(40, 60) {
		t.Fatal("ShiftRight failed")
	}

	if !b.IsRangeClear(60, b.Len()) {
		t.Fatal("ShiftRight failed")
	}
}

func TestShiftRight2(t *testing.T) {
	useShiftFastPath = false
	defer func() {
		useShiftFastPath = true
	}()

	if err := quick.CheckEqual(func(b Bitset, size, shift uint) []byte {
		b1 := New(size)
		copy(b1[shift>>3:], b)
		return b1
	}, func(b Bitset, size, shift uint) []byte {
		b1 := New(size)
		b1.ShiftRight(b, shift)
		return b1
	}, &quick.Config{
		Values:        rangeTestShiftValues,
		MaxCountScale: 100,
	}); err != nil {
		t.Error(err)
	}
}
