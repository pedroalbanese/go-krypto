//go:build amd64 && gc && !purego

package lsh512

import (
	"hash"
	"testing"
)

func Benchmark_Hash_8_SSSE3(b *testing.B)  { benchmarkSize(b, newSSSE3, 8, hasSSSE3) }
func Benchmark_Hash_1K_SSSE3(b *testing.B) { benchmarkSize(b, newSSSE3, 1024, hasSSSE3) }
func Benchmark_Hash_8K_SSSE3(b *testing.B) { benchmarkSize(b, newSSSE3, 8192, hasSSSE3) }

func newSSSE3(size int) hash.Hash {
	return newContextAsm(size, simdSetSSSE3)
}