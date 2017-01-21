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
