// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "bytes"

func (b Bitset) Equal(b1 Bitset) bool {
	return bytes.Equal(b, b1)
}

func (b Bitset) EqualRange(b1 Bitset, start, end uint) bool {
	if start > end {
		panic(errEndLessThanStart)
	}

	if end > uint(b.Len()) || end > uint(b1.Len()) {
		panic(errOutOfRange)
	}

	for ; start&7 != 0 && start < end; start++ {
		if b[start>>3]&(1<<(start&7)) != b1[start>>3]&(1<<(start&7)) {
			return false
		}
	}

	if !bytes.Equal(b[start>>3:end>>3], b1[start>>3:end>>3]) {
		return false
	}

	for start = end &^ 7; start < end; start++ {
		if b[start>>3]&(1<<(start&7)) != b1[start>>3]&(1<<(start&7)) {
			return false
		}
	}

	return true
}
