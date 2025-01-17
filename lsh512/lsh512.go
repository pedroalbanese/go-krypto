// Package lsh512 implements the LSH-512, LSH-384, LSH-512-256, LSH-512-224 hash algorithms as defined in KS X 3262
package lsh512

import (
	"hash"

	"github.com/RyuaNerin/go-krypto"
)

func init() {
	krypto.RegisterHash(krypto.LSH512, New)
	krypto.RegisterHash(krypto.LSH512_384, New384)
	krypto.RegisterHash(krypto.LSH512_256, New256)
	krypto.RegisterHash(krypto.LSH512_224, New224)
}

var (
	newContext func(size int) hash.Hash               = newContextGo
	sum        func(size int, data []byte) [Size]byte = sumGo
)

const (
	// The size of a LSH-512 checksum in bytes.
	Size = 64
	// The size of a LSH-384 checksum in bytes.
	Size384 = 48
	// The size of a LSH-512-256 checksum in bytes.
	Size256 = 32
	// The size of a LSH-512-224 checksum in bytes.
	Size224 = 28

	// The blocksize of LSH-512, LSH-384, LSH-512-256 and LSH-512-224 in bytes.
	BlockSize = 256
)

// New returns a new hash.Hash computing the LSH-512 checksum.
func New() hash.Hash { return newContext(Size) }

// New384 returns a new hash.Hash computing the LSH-384 checksum.
func New384() hash.Hash { return newContext(Size384) }

// New256 returns a new hash.Hash computing the LSH-512-256 checksum.
func New256() hash.Hash { return newContext(Size256) }

// New224 returns a new hash.Hash computing the LSH-512-224 checksum.
func New224() hash.Hash { return newContext(Size224) }

// Sum512 returns the LSH-512 checksum of the data.
func Sum512(data []byte) (sum512 [Size]byte) {
	return sum(Size, data)
}

// Sum384 returns the LSH-384 checksum of the data.
func Sum384(data []byte) (sum384 [Size384]byte) {
	sum := sum(Size384, data)
	copy(sum384[:], sum[:Size384])
	return
}

// Sum256 returns the LSH-512-256 checksum of the data.
func Sum256(data []byte) (sum256 [Size256]byte) {
	sum := sum(Size256, data)
	copy(sum256[:], sum[:Size256])
	return
}

// Sum224 returns the LSH-512-224 checksum of the data.
func Sum224(data []byte) (sum224 [Size224]byte) {
	sum := sum(Size224, data)
	copy(sum224[:], sum[:Size224])
	return
}
