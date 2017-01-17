// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"errors"

	"github.com/tmthrgd/go-bitwise"
	"github.com/tmthrgd/go-memset"
)

var (
	errEndLessThanStart = errors.New("cannot range backwards")
	errOutOfRange       = errors.New("out of range")
)

type Bitset []byte

func New(size uint) Bitset {
	size = (size + 7) &^ 7
	return make([]byte, size>>3)
}

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

func (b Bitset) Len() int {
	return len(b) * 8
}

func (b Bitset) ByteLen() int {
	return len(b)
}

func (b Bitset) Subset(start, end uint) Bitset {
	if start > end {
		panic(errEndLessThanStart)
	}

	if end > uint(b.Len()) {
		panic(errOutOfRange)
	}

	if start&7 != 0 || end&7 != 0 {
		panic(errors.New("cannot take subset accross byte boundary"))
	}

	return b[start>>3 : end>>3]
}
