// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

func TestComplement(t *testing.T) {
	b := New(80)
	b.Complement(b)

	if !b.All() {
		t.Error("Complement failed, All should have returned true")
	}

	b.Complement(b)

	if !b.None() {
		t.Error("Complement failed, None should have returned true")
	}
}

func TestUnion(t *testing.T) {
	b, b1 := New(80), New(80)

	b.SetRange(0, 40)
	b1.SetRange(40, 80)

	b.Union(b, b1)

	if !b.All() {
		t.Error("Union failed")
	}
}

func TestIntersection(t *testing.T) {
	b, b1 := New(80), New(80)

	b.SetAll()
	b1.SetRange(40, 80)

	b.Intersection(b, b1)

	if !b.IsRangeClear(0, 40) {
		t.Error("Intersection failed")
	}

	if !b.IsRangeSet(40, 80) {
		t.Error("Intersection failed")
	}
}

func TestDifference(t *testing.T) {
	b, b1 := New(80), New(80)

	b.SetAll()
	b1.SetRange(40, 80)

	b.Difference(b, b1)

	if !b.IsRangeSet(0, 40) {
		t.Error("Difference failed")
	}

	if !b.IsRangeClear(40, 80) {
		t.Error("Difference failed")
	}
}

func TestSymmetricDifference(t *testing.T) {
	b, b1 := New(80), New(80)

	b.SetRange(20, 60)
	b1.SetRange(40, 80)

	b.SymmetricDifference(b, b1)

	if !b.IsRangeClear(0, 20) {
		t.Error("SymmetricDifference failed")
	}

	if !b.IsRangeSet(20, 40) {
		t.Error("SymmetricDifference failed")
	}

	if !b.IsRangeClear(40, 60) {
		t.Error("SymmetricDifference failed")
	}

	if !b.IsRangeSet(60, 80) {
		t.Error("SymmetricDifference failed")
	}
}
