// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "testing"

func TestComplement(t *testing.T) {
	b := make(Bitset, 10)
	b.Complement(b)

	if !b.IsRangeSet(0, b.Len()) {
		t.Errorf("Not failed, should have found range #0-#%d", b.Len())
	}

	b.Complement(b)

	if !b.IsRangeClear(0, b.Len()) {
		t.Errorf("Not failed, should not have found in range #0-#%d", b.Len())
	}
}
