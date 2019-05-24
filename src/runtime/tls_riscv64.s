// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build riscv64

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

// NOTE: mcall() assumes this clobbers only X31 (REGTMP).
TEXT runtime·save_g(SB),NOSPLIT|NOFRAME,$0-0
	MOVB	runtime·iscgo(SB), X31
	BEQ	X31, X0, nocgo

	MOV	T0, X31	// save T0
	MOV	g, runtime·tls_g(SB) // TLS relocation clobbers T0
	MOV	X31, T0	// restore T0

nocgo:
	RET

TEXT runtime·load_g(SB),NOSPLIT|NOFRAME,$0-0
	MOV	runtime·tls_g(SB), g // TLS relocation clobbers T0
	RET

GLOBL runtime·tls_g(SB), TLSBSS, $8
