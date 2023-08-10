package has160

import (
	"bytes"
	"testing"
)

type testCase struct {
	M  []byte
	MD []byte
}

func testGo(t *testing.T, testCases []testCase) {
	h := New()

	out := make([]byte, BlockSize)

	for _, tc := range testCases {
		h.Reset()
		h.Write(tc.M)
		out = h.Sum(out[:0])

		if !bytes.Equal(out, tc.MD) {
			t.Fail()
		}
	}
}
