// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package bitset

func (a Atomic) Set(bit uint) {
	if bit > a.Len() {
		panic(errOutOfRange)
	}

	ptr, mask := a.index(bit)
	old := ptr.Load()
	for !ptr.CompareAndSwap(old, old|mask) {
		old = ptr.Load()
	}
}

func (a Atomic) Clear(bit uint) {
	if bit > a.Len() {
		panic(errOutOfRange)
	}

	ptr, mask := a.index(bit)
	old := ptr.Load()
	for !ptr.CompareAndSwap(old, old&^mask) {
		old = ptr.Load()
	}
}

func (a Atomic) Invert(bit uint) {
	if bit > a.Len() {
		panic(errOutOfRange)
	}

	ptr, mask := a.index(bit)
	old := ptr.Load()
	for !ptr.CompareAndSwap(old, old^mask) {
		old = ptr.Load()
	}
}

func (a Atomic) SetTo(bit uint, value bool) {
	if value {
		a.Set(bit)
	} else {
		a.Clear(bit)
	}
}
