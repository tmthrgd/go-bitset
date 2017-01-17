// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "github.com/tmthrgd/go-popcount"

func (b Bitset) Count() int {
	return int(popcount.CountBytes(b))
}

func (b Bitset) CountRange(start, end uint) int {
	if start > end {
		panic(errEndLessThanStart)
	}

	if end > uint(b.Len()) {
		panic(errOutOfRange)
	}

	var total uint64

	for ; start&7 != 0 && start < end; start++ {
		if b[start>>3]&(1<<(start&7)) != 0 {
			total++
		}
	}

	total += popcount.CountBytes(b[start>>3 : end>>3])

	for start = end &^ 7; start < end; start++ {
		if b[start>>3]&(1<<(start&7)) != 0 {
			total++
		}
	}

	return int(total)
}
