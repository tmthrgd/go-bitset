// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

func mask1(start, end uint) byte {
	/*for ; start&7 != 0 && start < end; start++ {
		mask |= 1 << (start & 7)
	}*/

	return ((0xff << (start & 7)) ^ (0xff << (end - start&^7))) & ((1 >> (start & 7)) - 1)
}

func mask2(end uint) byte {
	/*for start := end &^ 7; start < end; start++ {
		mask |= 1 << (start & 7)
	}*/

	return (1 << (end & 7)) - 1
}
