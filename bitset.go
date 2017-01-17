// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"errors"

	"github.com/tmthrgd/go-hex"
)

var (
	errEndLessThanStart = errors.New("cannot range backwards")
	errOutOfRange       = errors.New("out of range")
)

type Bitset []byte

func New(size uint) Bitset {
	size = (size + 7) &^ 7
	return make(Bitset, size>>3)
}

func (b Bitset) Len() uint {
	return uint(len(b)) << 3
}

func (b Bitset) ByteLen() int {
	return len(b)
}

func (b Bitset) Subset(start, end uint) Bitset {
	if start > end {
		panic(errEndLessThanStart)
	}

	if end > b.Len() {
		panic(errOutOfRange)
	}

	if start&7 != 0 || end&7 != 0 {
		panic(errors.New("cannot take subset accross byte boundary"))
	}

	return b[start>>3 : end>>3]
}

func (b Bitset) Clone() Bitset {
	return append(Bitset(nil), b...)
}

func (b Bitset) CloneRange(start, end uint) Bitset {
	return b.Subset(start, end).Clone()
}

func (b Bitset) String() string {
	const maxSize = 128

	if len(b) > maxSize {
		return "Bitset{" + hex.EncodeToString(b[:maxSize]) + "...}"
	}

	return "Bitset{" + hex.EncodeToString(b) + "}"
}
