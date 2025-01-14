package has160

import (
	"strings"
	"testing"

	. "github.com/RyuaNerin/go-krypto/testingutil"
)

func Test_HAS160(t *testing.T) { HT(t, New(), testCases, false) }

// TTAS.KO-12.0011/R2
var testCases = []HashTestCase{
	{
		MsgBytes: []byte(``),
		MD:       `30 79 64 ef 34 15 1d 37 c8 04 7a de c7 ab 50 f4 ff 89 76 2d`,
	},
	{
		MsgBytes: []byte(`a`),
		MD:       `48 72 bc bc 4c d0 f0 a9 dc 7c 2f 70 45 e5 b4 3b 6c 83 0d b8`,
	},
	{
		MsgBytes: []byte(`abc`),
		MD:       `97 5e 81 04 88 cf 2a 3d 49 83 84 78 12 4a fc e4 b1 c7 88 04`,
	},
	{
		MsgBytes: []byte(`message digest`),
		MD:       ` 23 38 db c8 63 8d 31 22 5f 73 08 62 46 ba 52 9f 96 71 0b c6`,
	},
	{
		MsgBytes: []byte(`abcdefghijklmnopqrstuvwxyz`),
		MD:       `59 61 85 c9 ab 67 03 d0 d0 db b9 87 02 bc 0f 57 29 cd 1d 3c`,
	},
	{
		MsgBytes: []byte(`12345678901234567890123456789012345678901234567890123456789012345678901234567890`),
		MD:       `07 f0 5c 8c 07 73 c5 5c a3 a5 a6 95 ce 6a ca 4c 43 89 11 b5`,
	},
	{
		MsgBytes: []byte(strings.Repeat("a", 1_000_000)),
		MD:       `d6 ad 6f 06 08 b8 78 da 9b 87 99 9c 25 25 cc 84 f4 c9 f1 8d`,
	},
	{
		MsgBytes: []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`),
		MD:       `cb 5d 7e fb ca 2f 02 e0 fb 71 67 ca bb 12 3a f5 79 57 64 e5`,
	},
}
