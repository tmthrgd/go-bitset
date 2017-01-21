// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

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
