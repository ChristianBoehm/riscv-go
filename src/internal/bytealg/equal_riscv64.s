// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "textflag.h"

#define	CTXT	S4

// TODO: share code with memequal?
// func Equal(a, b []byte) bool
TEXT ·Equal(SB),NOSPLIT,$0-49
	MOV	a_len+8(FP), A3
	MOV	b_len+32(FP), A4
	BNE	A3, A4, noteq		// unequal lengths are not equal

	MOV	a+0(FP), A1
	MOV	b+24(FP), A2
	ADD	A1, A3		// end

loop:
	BEQ	A1, A3, equal		// reached the end
	MOVBU	(A1), A6
	ADD	$1, A1
	MOVBU	(A2), A7
	ADD	$1, A2
	BEQ	A6, A7, loop

noteq:
	MOVB	ZERO, ret+48(FP)
	RET

equal:
	MOV	$1, A1
	MOVB	A1, ret+48(FP)
	RET

// func memequal(a, b unsafe.Pointer, size uintptr) bool
TEXT runtime·memequal(SB),NOSPLIT|NOFRAME,$0-25
	MOV	a+0(FP), A1
	MOV	b+8(FP), A2
	BEQ	A1, A2, eq
	MOV	size+16(FP), A3
	ADD	A1, A3, A4
loop:
	BNE	A1, A4, test
	MOV	$1, A1
	MOVB	A1, ret+24(FP)
	RET
test:
	MOVBU	(A1), A6
	ADD	$1, A1
	MOVBU	(A2), A7
	ADD	$1, A2
	BEQ	A6, A7, loop

	MOVB	ZERO, ret+24(FP)
	RET
eq:
	MOV	$1, A1
	MOVB	A1, ret+24(FP)
	RET

// func memequal_varlen(a, b unsafe.Pointer) bool
TEXT runtime·memequal_varlen(SB),NOSPLIT,$40-17
	MOV	a+0(FP), A1
	MOV	b+8(FP), A2
	BEQ	A1, A2, eq
	MOV	8(CTXT), A3    // compiler stores size at offset 8 in the closure
	MOV	A1, 8(X2)
	MOV	A2, 16(X2)
	MOV	A3, 24(X2)
	CALL	runtime·memequal(SB)
	MOVBU	32(X2), A1
	MOVB	A1, ret+16(FP)
	RET
eq:
	MOV	$1, A1
	MOVB	A1, ret+16(FP)
	RET
