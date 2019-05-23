// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/riscv"
)

// FIXME: This is incredibly inefficient, but nice and simple. Optimize.
// See other zerorange implementations for ideas.
func zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	// Loop, zeroing one byte at a time.
	// ADD	$(frame+lo), SP, T0
	// ADD	$(cnt), T0, T1
	// loop:
	// 	MOVB	ZERO, (T0)
	// 	ADD	$1, T0
	//	BNE	T0, T1, loop
	p = pp.Appendpp(p, riscv.AADD, obj.TYPE_CONST, 0, 8+off, obj.TYPE_REG, riscv.REG_T0, 0)
	p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: riscv.REG_SP})
	p = pp.Appendpp(p, riscv.AADD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, riscv.REG_T1, 0)
	p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: riscv.REG_T0})
	p = pp.Appendpp(p, riscv.AMOVB, obj.TYPE_REG, riscv.REG_ZERO, 0, obj.TYPE_MEM, riscv.REG_T0, 0)
	loop := p
	p = pp.Appendpp(p, riscv.AADD, obj.TYPE_CONST, 0, 1, obj.TYPE_REG, riscv.REG_T0, 0)
	p = pp.Appendpp(p, riscv.ABNE, obj.TYPE_REG, riscv.REG_T0, 0, obj.TYPE_BRANCH, 0, 0)
	p.Reg = riscv.REG_T1
	gc.Patch(p, loop)
	return p
}

func zeroAuto(pp *gc.Progs, n *gc.Node) {
	// Note: this code must not clobber any registers.
	sym := n.Sym.Linksym()
	size := n.Type.Size()
	for i := int64(0); i < size; i += 8 {
		p := pp.Prog(riscv.AMOV)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = riscv.REG_ZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = riscv.REG_SP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}
