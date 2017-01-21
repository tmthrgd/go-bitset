// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"testing"
	"testing/quick"
)

func TestCopy(t *testing.T) {
	b, b1 := New(80), New(80)
	b1.SetAll()

	b.Copy(b1)

	if !b.All() {
		t.Error("Copy failed")
	}
}

func TestCopyRange(t *testing.T) {
	if err := quick.CheckEqual(func(b, b1 Bitset, start, end uint) []byte {
		b = b.Clone()

		for i := start; i < end; i++ {
			b.SetTo(i, b1.IsSet(i))
		}

		return b
	}, func(b, b1 Bitset, start, end uint) []byte {
		b = b.Clone()
		b.CopyRange(b1, start, end)
		return b
	}, &quick.Config{
		Values:        rangeTestValues2,
		MaxCountScale: 100,
	}); err != nil {
		t.Error(err)
	}
}
