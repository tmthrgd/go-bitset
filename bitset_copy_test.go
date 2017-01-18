// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

func TestCopy(t *testing.T) {
	b := make(Bitset, 10)
	b1 := b.Clone()

	b.Set(10)
	b1.Copy(b)

	if !b.Equal(b1) {
		t.Error("Copy failed")
	}

	b.Clear(10)
	b1.Copy(b)

	if !b.Equal(b1) {
		t.Error("Copy failed")
	}
}
