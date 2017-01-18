// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

func TestEqual(t *testing.T) {
	b := make(Bitset, 10)
	b1 := make(Bitset, len(b))

	if !b.Equal(b1) {
		t.Error("Equal failed")
	}

	b.Set(10)

	if b.Equal(b1) {
		t.Error("Equal failed")
	}

	b.Clear(10)

	if !b.Equal(b1) {
		t.Error("Equal failed")
	}
}
