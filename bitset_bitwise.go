// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "github.com/tmthrgd/go-bitwise"

func (b Bitset) Complement(b1 Bitset) {
	bitwise.Not(b, b1)
}

func (b Bitset) Union(b1, b2 Bitset) {
	bitwise.Or(b, b1, b2)
}

func (b Bitset) Intersection(b1, b2 Bitset) {
	bitwise.And(b, b1, b2)
}

func (b Bitset) Difference(b1, b2 Bitset) {
	bitwise.AndNot(b, b1, b2)
}

func (b Bitset) SymmetricDifference(b1, b2 Bitset) {
	bitwise.XOR(b, b1, b2)
}
