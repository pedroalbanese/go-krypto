# [krypto](https://pkg.go.dev/github.com/RyuaNerin/go-krypto)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/RyuaNerin/go-krypto)](https://pkg.go.dev/github.com/RyuaNerin/go-krypto)

- Golang implementation of cryptographic algorithms designed in Republic of Korea

- It is intended for compatibility with the `crypto` package.

- `krypto` does not required any other C/C++ compiler to use SIMD :\)

    - But, may have not been optimized in compiler level.

## [LICNESE](/LICENSE)

## Supported

- `A` : Not tested.
- `^` : Deprecated algorithm.

- Block cipher

    | Algorithm | Package        | Document           | 128 | 192 | 256 | SIMD Supports      |
    |:---------:|----------------|:------------------:|:---:|:---:|:---:|:------------------:|
    | SEED 128  | `krypto/seed`  | TTAS.KO-12.0004/R1 | O   |     |     |                    |
    | SEED 256  | `krypto/seed`  | ***Unknown***      |     |     | A   |                    |
    | HIGHT     | `krypto/hight` | TTAS.KO-12.0040/R1 | O   |     |     |                    |
    | ARIA      | `krypto/aria`  | KS X 1213-1        | O   | O   | O   |                    |
    | LEA       | `krypto/lea`   | TTAK.KO-12.0223    | O   | O   | O   | x86: `SSE2` `AVX2` |

- Digital Signature

    | Algorithm | Package          | Document           |
    |:---------:|------------------|:------------------:|
    | KCDSA     | `krypto/kcdsa`   | TTAK.KO-12.0001/R4 |
    | EC-KCDSA  | `krypto/eckcdsa` | TTAK.KO-12.0015/R3 |

- Hash

    | Algorithm  | Package         | Document           | 160 | 224 | 256 | 384 | 512 | SIMD Supports              |
    |:----------:|-----------------|:------------------:|:---:|:---:|:---:|:---:|:---:|:--------------------------:|
    | HAS-160`^` | `krypto/has160` | TTAS.KO-12.0011/R2 | O   |     |     |     |     |                            |
    | LSH-256    | `krypto/lsh256` | KS X 3262          |     | O   | O   |     |     | x86: `SSE2` `SSSE3` `AVX2` |
    | LSH-512    | `krypto/lsh512` | KS X 3262          |     | O   | O   | O   | O   | x86: `SSE2` `SSSE3` `AVX2` |

### SIMD Support

- It was based on the below.

    | Algorithm | x86 SIMD Supports       | Reference                                                   |
    |:---------:|-------------------------|:-----------------------------------------------------------:|
    | LEA       | `SSE2` `AVX2`*          | [KISA](https://seed.kisa.or.kr/kisa/Board/20/detailView.do) |
    | LSH-256   | `SSE2` `SSSE3` `AVX2`** | [KISA](https://seed.kisa.or.kr/kisa/Board/22/detailView.do) |
    | LSH-512   | `SSE2` `SSSE3` `AVX2`   | [KISA](https://seed.kisa.or.kr/kisa/Board/22/detailView.do) |

    - *: LEA AVX2 was Disabled in default.

        - SSSE3 is more faster then AVX2 in benchmarks. It seems like a different algorithm is needed.

        - you can enable LEA AVX2 manually by build from `krypto_lea_avx2` tags.

            ```bash
            $ go build -tags=krypto_lea_avx2 .
            ```

    - **: LSH-256 AVX2 with Disabled in default.

        - SSSE3 is more faster then AVX2 in benchmarks. It seems like a different algorithm is needed.

        - but you can enable LSH-256 AVX2 manually by build from `krypto_lsh256_avx2` tags.

            ```bash
            $ go build -tags=krypto_lsh256_avx2 .
            ```

## [Performance](/PERFORMANCE.md)

## Installation

```shell
go get -v "github.com/RyuaNerin/go-krypto"
```

```go
package main

import (
    ...
    krypto "github.com/RyuaNerin/go-krypto"
    ...
)
```

## Usage

Todo

## TODO

- [ ] Optimized SIMD

- [ ] ARMv8 NEON supports

    - [mmcloughlin/avo: Generate x86 Assembly with Go](https://github.com/mmcloughlin/avo) does not support ARM.

- Supoorts Post-Quantum Cryptography
