// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// Copyright 2014 Will Fitzgerald. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitset

import "testing"

func TestIsSuperSet(t *testing.T) {
	a := New(500)
	b := New(300)
	c := New(200)

	// Setup bitsets
	// a and b overlap
	// only c is (strict) super set

	a.SetRange(0, 100)
	b.SetRange(50, 150)
	c.SetRange(0, 200)

	if a.IsSuperSet(b) {
		t.Errorf("IsSuperSet fails")
	}

	if a.IsSuperSet(c) {
		t.Errorf("IsSuperSet fails")
	}

	if b.IsSuperSet(a) {
		t.Errorf("IsSuperSet fails")
	}

	if b.IsSuperSet(c) {
		t.Errorf("IsSuperSet fails")
	}

	if !c.IsSuperSet(a) {
		t.Errorf("IsSuperSet fails")
	}

	if !c.IsSuperSet(b) {
		t.Errorf("IsSuperSet fails")
	}

	if a.IsStrictSuperSet(b) {
		t.Errorf("IsStrictSuperSet fails")
	}

	if a.IsStrictSuperSet(c) {
		t.Errorf("IsStrictSuperSet fails")
	}

	if b.IsStrictSuperSet(a) {
		t.Errorf("IsStrictSuperSet fails")
	}

	if b.IsStrictSuperSet(c) {
		t.Errorf("IsStrictSuperSet fails")
	}

	if !c.IsStrictSuperSet(a) {
		t.Errorf("IsStrictSuperSet fails")
	}

	if !c.IsStrictSuperSet(b) {
		t.Errorf("IsStrictSuperSet fails")
	}
}
