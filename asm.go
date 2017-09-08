package main

import (
	"fmt"
	"math"
)

func fma(x, y, z float64) float64

func normal(x, y, z float64) float64

func myxatan(x float64) float64

func testxatan(x float64) float64

func main() {
	fmt.Println(math.Atan2(-479, 123))
	fmt.Println(myatan2(-479, 123))
	// fmt.Println(myxatan(123))
	// fmt.Println(test(3))
	// fmt.Println(xatan(123))
}

func test(x float64) float64 {
	return 45 * x
}

// xatan evaluates a series valid in the range [0, 0.66].
func xatan(x float64) float64 {
	const (
		P0 = -8.750608600031904122785e-01
		P1 = -1.615753718733365076637e+01
		P2 = -7.500855792314704667340e+01
		P3 = -1.228866684490136173410e+02
		P4 = -6.485021904942025371773e+01
		Q0 = +2.485846490142306297962e+01
		Q1 = +1.650270098316988542046e+02
		Q2 = +4.328810604912902668951e+02
		Q3 = +4.853903996359136964868e+02
		Q4 = +1.945506571482613964425e+02
	)
	z := x * x
	z = z * ((((P0*z+P1)*z+P2)*z+P3)*z + P4) / (((((z+Q0)*z+Q1)*z+Q2)*z+Q3)*z + Q4)
	z = x*z + x
	return z
}

/*
TEXT ·fma(SB),$0
	FMOVD x+0(FP), F0
	FMOVD F0, F1
	FMOVD y+8(FP), F0
	FADDD F1, F0
	FMOVD F0, F1
	FMOVD z+16(FP), F0
	FADDD F1, F0
	FMOVD F0, ret+24(FP)
	RET

// func fma(x, y, z) float64
TEXT ·fma(SB),NOSPLIT,$0
	MOVSD x+0(FP), X0
	MOVSD y+8(FP), X1
	ADDSD X1, X0
	MOVSD z+16(FP), X1
	ADDSD X1, X0
	// VFMADD213PD X0, X1, xmm2/m128
	// BYTE $0x0f; BYTE $0x38; BYTE $0xa8; hex codes for X0, X1??
	MOVSD X0, ret+24(FP)
	RET

// func fma(x, y, z) float64
TEXT ·fma(SB),NOSPLIT,$0
	MOVSD x+0(FP), X0
	MOVSD y+8(FP), X1
	MOVSD z+16(FP), X2
	// VFMADD213PD X0, X1, X2
	BYTE $0xc4; BYTE $0xe2; BYTE $0xF1; BYTE $0xa8; BYTE $0xC2
	MOVSD X0, ret+24(FP)
	RET

*/
