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

	for i := uint(0); i < 100; i++ {
		a.Set(i)
	}

	for i := uint(50); i < 150; i++ {
		b.Set(i)
	}

	for i := uint(0); i < 200; i++ {
		c.Set(i)
	}

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
