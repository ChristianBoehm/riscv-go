// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build riscv64

/*
 * void crosscall1(void (*fn)(void))
 *
 * Calling into the go tool chain, where all registers are caller save.
 * Called from standard riscv64 ABI, where x2, x8, x9, x18–27
 * are callee-save so they must be saved explicitly. Also x1 (RA).
 * For FP registrers,  f8–9 and f18–27 must be saved.
 */
.globl crosscall1
crosscall1:
	addi sp, sp, -120
	sd	x1, 0(sp)
	sd	x18, 8(sp)
	sd	x19, 16(sp)
	sd	x20, 24(sp)
	sd	x21, 32(sp)
	sd	x22, 40(sp)
	sd	x23, 48(sp)
	sd	x24, 56(sp)
	sd	x25, 64(sp)
	sd	x26, 72(sp)
	sd	x27, 80(sp)
	sd	x3, 88(sp)
	sd	x8, 96(sp)
	sd	x9, 104(sp)
	sd	x2, 112(sp)

	move x18, x10
	move x19, x11
	move x10, x12

	jalr x1, x19	    // call setg_gcc
	jalr x1, x18	    // call fn

	ld	x18, 8(sp)
	ld	x19, 16(sp)
	ld	x20, 24(sp)
	ld	x21, 32(sp)
	ld	x22, 40(sp)
	ld	x23, 48(sp)
	ld	x24, 56(sp)
	ld	x25, 64(sp)
	ld	x26, 72(sp)
	ld	x27, 80(sp)
	ld	x3,  88(sp)
	ld	x8,  96(sp)
	ld	x9,  104(sp)
	ld	x2, 208(sp)
	ld	x1,  0(sp)
	addi sp, sp, 120
	jalr x0, x1

#ifdef __ELF__
.section .note.GNU-stack,"",@progbits
#endif
