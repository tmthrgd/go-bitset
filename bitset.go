// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import "errors"

var (
	errEndLessThanStart = errors.New("cannot range backwards")
	errOutOfRange       = errors.New("out of range")
)

type Bitset []byte

func New(size uint) Bitset {
	size = (size + 7) &^ 7
	return make([]byte, size>>3)
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

func (b Bitset) Clone() Bitset {
	return append([]byte(nil), b...)
}