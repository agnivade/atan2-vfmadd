BITS 64

; myvar dd 1.615753718733365076637e+01
; vmovsd xmm2,[rsp+8]
; VZEROUPPER
; VFNMADD231PD xmm0, xmm1, xmm2
VFMADD213PD xmm0, xmm2, [rip+0x41158-1]
; VMOVSD xmm3, [rip+0x710a3]
; VSUBSD xmm1, xmm0, [rip+0x3fe6a]
; c4 e2 f9 a8 0d 1f 11    vfmadd213pd 0x7111f(%rip),%xmm0,%xmm1        # 4c74b5 <exprodata+0x35>
