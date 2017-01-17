// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"github.com/tmthrgd/go-bitwise"
	"github.com/tmthrgd/go-memset"
)

func (b Bitset) Set(bit uint) {
	if bit > uint(b.Len()) {
		panic(errOutOfRange)
	}

	b[bit>>3] |= 1 << (bit & 7)
}

func (b Bitset) Clear(bit uint) {
	if bit > uint(b.Len()) {
		panic(errOutOfRange)
	}

	b[bit>>3] &^= 1 << (bit & 7)
}

func (b Bitset) Invert(bit uint) {
	if bit > uint(b.Len()) {
		panic(errOutOfRange)
	}

	b[bit>>3] ^= 1 << (bit & 7)
}

func (b Bitset) SetRange(start, end uint) {
	if start > end {
		panic(errEndLessThanStart)
	}

	if end > uint(b.Len()) {
		panic(errOutOfRange)
	}

	for ; start&7 != 0 && start < end; start++ {
		b[start>>3] |= 1 << (start & 7)
	}

	memset.Memset(b[start>>3:end>>3], 0xff)

	for start = end &^ 7; start < end; start++ {
		b[start>>3] |= 1 << (start & 7)
	}
}

func (b Bitset) ClearRange(start, end uint) {
	if start > end {
		panic(errEndLessThanStart)
	}

	if end > uint(b.Len()) {
		panic(errOutOfRange)
	}

	for ; start&7 != 0 && start < end; start++ {
		b[start>>3] &^= 1 << (start & 7)
	}

	memset.Memset(b[start>>3:end>>3], 0)

	for start = end &^ 7; start < end; start++ {
		b[start>>3] &^= 1 << (start & 7)

	}
}

func (b Bitset) InvertRange(start, end uint) {
	if start > end {
		panic(errEndLessThanStart)
	}

	if end > uint(b.Len()) {
		panic(errOutOfRange)
	}

	for ; start&7 != 0 && start < end; start++ {
		b[start>>3] ^= 1 << (start & 7)
	}

	bitwise.Not(b[start>>3:end>>3], b[start>>3:end>>3])

	for start = end &^ 7; start < end; start++ {
		b[start>>3] ^= 1 << (start & 7)
	}
}

func (b Bitset) SetTo(bit uint, value bool) {
	if value {
		b.Set(bit)
	} else {
		b.Clear(bit)
	}
}

func (b Bitset) SetRangeTo(start, end uint, value bool) {
	if value {
		b.SetRange(start, end)
	} else {
		b.ClearRange(start, end)
	}
}

func (b Bitset) SetAll() {
	memset.Memset(b, 0xff)
}

func (b Bitset) ClearAll() {
	memset.Memset(b, 0)
}
