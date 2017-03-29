// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package bitset

import "github.com/tmthrgd/atomics"

func (a Atomic) index(bit uint) (ptr *atomics.Uint64, mask uint64) {
	return &a[bit/64], 1 << (bit & 63)
}
