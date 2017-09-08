// +build amd64
#include "textflag.h"

DATA xatandata<>+0(SB)/8, $-8.750608600031904122785e-01
DATA xatandata<>+8(SB)/8, $-1.615753718733365076637e+01
DATA xatandata<>+16(SB)/8, $-7.500855792314704667340e+01
DATA xatandata<>+24(SB)/8, $-1.228866684490136173410e+02
DATA xatandata<>+32(SB)/8, $-6.485021904942025371773e+01
DATA xatandata<>+40(SB)/8, $+2.485846490142306297962e+01
DATA xatandata<>+48(SB)/8, $+1.650270098316988542046e+02
DATA xatandata<>+56(SB)/8, $+4.328810604912902668951e+02
DATA xatandata<>+64(SB)/8, $+4.853903996359136964868e+02
DATA xatandata<>+72(SB)/8, $+1.945506571482613964425e+02
GLOBL xatandata<>+0(SB), RODATA, $80

// func fma(x, y, z) float64
TEXT ·fma(SB),NOSPLIT,$0
	MOVSD x+0(FP), X0
	MOVSD y+8(FP), X2
	MOVSD z+16(FP), X3
	// VFMADD213PD X0, X2, X3
	// BYTE $0xC4; BYTE $0xE2; BYTE $0xE9; BYTE $0xA8; BYTE $0xC3
	// VFMADD213PD X3, X2, X4
	BYTE $0xC4; BYTE $0xE2; BYTE $0xE9; BYTE $0xA8; BYTE $0xDC
	MOVSD X0, ret+24(FP)
	RET

// func normal(x, y, z) float64
TEXT ·normal(SB),NOSPLIT,$0
	MOVSD   x+0(FP), X0
	MOVSD   y+8(FP), X1
	MULSD   X1, X0
	MOVSD   z+16(FP), X1
	SUBSD   X0, X1
	MOVSD   X1, ret+24(FP)
	RET

// func myxatan(x) float64
TEXT ·myxatan(SB),NOSPLIT,$0-16
	// VMOVSD   x+0(FP), X2
	BYTE $0xC5; BYTE $0xFB; BYTE $0x10; BYTE $0x54; BYTE $0x24; BYTE $0x08

	// Storing x into xmm1
	MOVUPS  X2, X1
	// Storing z into xmm2
	// VMULSD  X2, X2, X2
	BYTE $0xC5; BYTE $0xEB; BYTE $0x59; BYTE $0xD2

	// Calculating numerator
	MOVSD   xatandata<>+0(SB), X0
	VFMADD213SD xatandata<>+8(SB), X2, X0
	VFMADD213SD xatandata<>+16(SB), X2, X0
	VFMADD213SD xatandata<>+24(SB), X2, X0
	VFMADD213SD xatandata<>+32(SB), X2, X0

	BYTE $0xC5; BYTE $0xEB; BYTE $0x59; BYTE $0xC0 // VMULSD X0, X2, X0 // storing numerator in X0

	// Calculating denominator
	VADDSD xatandata<>+40(SB), X2, X3
	VFMADD213SD xatandata<>+48(SB), X2, X3
	VFMADD213SD xatandata<>+56(SB), X2, X3
	VFMADD213SD xatandata<>+64(SB), X2, X3
	VFMADD213SD xatandata<>+72(SB), X2, X3

	// Dividing numerator / denominator
	BYTE $0xC5; BYTE $0xFB; BYTE $0x5E; BYTE $0xC3 // VDIVSD xmm0, xmm0, xmm3

	// Doing final x*z+x
	VFMADD213SD X1, X0, X1
	MOVSD   X1, ret+8(FP)
	RET

