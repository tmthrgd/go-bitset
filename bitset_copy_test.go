// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

func TestCopy(t *testing.T) {
	b, b1 := New(80), New(80)
	b1.SetAll()

	b.Copy(b1)

	if !b.All() {
		t.Error("Copy failed")
	}
}

func TestCopyRange(t *testing.T) {
	b, b1 := New(80), New(80)
	b1.SetAll()

	b.CopyRange(b1, 7, 63)

	if !b.IsRangeClear(0, 7) {
		t.Error("CopyRange failed")
	}

	if !b.IsRangeSet(7, 63) {
		t.Error("CopyRange failed")
	}

	if !b.IsRangeClear(63, b.Len()) {
		t.Error("CopyRange failed")
	}
}
