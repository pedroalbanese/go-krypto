// Code generated by command: go run main.go -out ../../aria_amd64.s -stubs ../../aria_amd64_stubs.go -pkg aria. DO NOT EDIT.

//go:build amd64 && gc && !purego

#include "textflag.h"

// func processFinSSE2(dst []byte, rk []byte, t []byte)
// Requires: SSE2
TEXT ·processFinSSE2(SB), NOSPLIT, $0-72
	MOVQ  dst_base+0(FP), AX
	MOVQ  rk_base+24(FP), CX
	MOVQ  t_base+48(FP), DX
	MOVOU (CX), X0
	MOVOU (DX), X1
	PXOR  X1, X0
	MOVOU X0, (AX)
	RET
