// Code generated by command: go run lea.go -out ../../lea_amd64.s -stubs ../../lea_amd64_stubs.go -pkg lea. DO NOT EDIT.

//go:build amd64 && gc && !purego

package lea

func leaEnc4SSE2(ctx *leaContext, dst []byte, src []byte)

func leaDec4SSE2(ctx *leaContext, dst []byte, src []byte)

func leaEnc8AVX2(ctx *leaContext, dst []byte, src []byte)

func leaDec8AVX2(ctx *leaContext, dst []byte, src []byte)
