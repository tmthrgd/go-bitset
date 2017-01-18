// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bitset

import (
	"math/rand"
	"testing"
)

func TestCopy(t *testing.T) {
	b := make(Bitset, 10)

	b1 := b.Clone()
	rand.Read(b1)

	if b.Equal(b1) {
		panic("Clone failed")
	}

	b.Copy(b1)

	if !b.Equal(b1) {
		t.Error("Copy failed")
	}
}

func TestCopyRange(t *testing.T) {
	b := make(Bitset, 10)

	b1 := b.Clone()
	rand.Read(b1)

	if b.Equal(b1) {
		panic("Clone failed")
	}

	b.CopyRange(b1, 7, 63)

	if !b.IsRangeClear(0, 8) {
		t.Error("Copy failed")
	}

	if !b.EqualRange(b1, 7, 64) {
		t.Error("Copy failed")
	}

	if !b.IsRangeClear(63, b.Len()) {
		t.Error("Copy failed")
	}
}
