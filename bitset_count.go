// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "github.com/tmthrgd/go-popcount"

func (b Bitset) Count() uint {
	return uint(popcount.CountBytes(b))
}

func (b Bitset) CountRange(start, end uint) uint {
	if start > end {
		panic(errEndLessThanStart)
	}

	if end > b.Len() {
		panic(errOutOfRange)
	}

	var x, total uint64

	if mask := mask1(start, end); mask != 0 {
		x = uint64(b[start>>3] & mask)
	}

	if start := (start + 7) &^ 7; start < end {
		total = popcount.CountBytes(b[start>>3 : end>>3])
	}

	if mask := mask2(start, end); mask != 0 {
		x |= uint64(b[end>>3]&mask) << 8
	}

	return uint(popcount.Count64(x) + total)
}
