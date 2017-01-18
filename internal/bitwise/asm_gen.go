// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// +build ignore

package main

import "github.com/tmthrgd/asm"

const header = `// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// This file is auto-generated - do not modify

// +build amd64,!gccgo,!appengine
`

func andeqASM(a *asm.Asm) {
	a.NewFunction("andeqASM")
	a.NoSplit()

	srcA := a.Argument("a", 8)
	srcB := a.Argument("b", 8)
	length := a.Argument("len", 8)
	ret := a.Argument("ret", 8)

	a.Start()

	bigloop := a.NewLabel("bigloop")
	loop := a.NewLabel("loop")
	retLabel := a.NewLabel("ret")
	retTrue := retLabel.Suffix("true")
	retFalse := retLabel.Suffix("false")

	sA, sB, cx := asm.SI, asm.DI, asm.BX

	a.Movq(sA, srcA)
	a.Movq(sB, srcB)
	a.Movq(cx, length)

	a.Cmpq(asm.Constant(16), cx)
	a.Jb(loop)

	a.Label(bigloop)

	a.Movou(asm.X0, asm.Address(sA, cx, asm.SX1, -16))
	a.Movou(asm.X1, asm.Address(sB, cx, asm.SX1, -16))

	a.Pand(asm.X0, asm.X1)

	a.Pxor(asm.X0, asm.X1)
	a.Ptest(asm.X0, asm.X0)
	a.Jnz(retFalse)

	a.Subq(cx, asm.Constant(16))
	a.Jz(retTrue)

	a.Cmpq(asm.Constant(16), cx)
	a.Jae(bigloop)

	a.Label(loop)

	a.Movb(asm.AX, asm.Address(sA, cx, asm.SX1, -1))
	a.Movb(asm.DX, asm.Address(sB, cx, asm.SX1, -1))

	a.Andb(asm.AX, asm.DX)

	a.Cmpb(asm.AX, asm.DX)
	a.Jne(retFalse)

	a.Subq(cx, asm.Constant(1))
	a.Jnz(loop)

	a.Label(retTrue)
	a.Movb(ret, asm.Constant(0x01))
	a.Ret()

	a.Label(retFalse)
	a.Movb(ret, asm.Constant(0x00))
	a.Ret()
}

func main() {
	if err := asm.Do("bitwise_andeq_amd64.s", header, andeqASM); err != nil {
		panic(err)
	}
}
