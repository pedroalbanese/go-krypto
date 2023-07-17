package kcdsa

import (
	"crypto/rand"
	"testing"

	"github.com/RyuaNerin/go-krypto/kcdsa/kcdsakisa"
)

func TestGenerateKisaJ(t *testing.T) {
	for _, tc := range testCases {
		domain, _ := tc.Sizes.domain()
		J, err := kcdsakisa.GenerateJ(tc.Seed_, domain.Domain)
		if err != nil {
			t.Error(err)
		}
		if J.Cmp(tc.J) != 0 {
			t.Errorf("GenerateKisaJ failed")
		}
	}
}

func TestGenerateKisaPQ(t *testing.T) {
	for _, tc := range testCases {
		domain, _ := tc.Sizes.domain()
		P, Q, count, err := kcdsakisa.GeneratePQ(tc.J, tc.Seed_, domain.Domain)
		if err != nil {
			t.Error(err)
		}
		if P.Cmp(tc.P) != 0 || Q.Cmp(tc.Q) != 0 || count != tc.Count {
			t.Errorf("GenerateKisaPQ failed")
		}
	}
}

func TestGenerateHG(t *testing.T) {
	for _, tc := range testCases {
		_, _, err := kcdsakisa.GenerateHG(rand.Reader, tc.P, tc.J)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGenerateKisaG(t *testing.T) {
	for _, tc := range testCases {
		G, err := kcdsakisa.GenerateG(tc.P, tc.J, tc.H)
		if err != nil {
			t.Error(err)
		}
		if G.Cmp(tc.G) != 0 {
			t.Errorf("GenerateKisaG failed")
		}
	}
}

func TestGenerateKisaXYZ(t *testing.T) {
	for _, tc := range testCases {
		domain, _ := tc.Sizes.domain()
		X, Y, Z, _, err := kcdsakisa.GenerateXYZ(tc.P, tc.Q, tc.G, UserProvidedRandomInput, tc.XKEY, domain.Domain)
		if err != nil {
			t.Error(err)
		}
		if X.Cmp(tc.X) != 0 || Y.Cmp(tc.Y) != 0 || Z.Cmp(tc.Z) != 0 {
			t.Errorf("GenerateKisaX failed")
		}
	}
}

func TestGenerateParametersKISA(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping parameter generation test in short mode")
	}

	gp := func(params *Parameters, sizes ParameterSizes) error {
		_, _, err := GenerateParametersKISA(params, rand.Reader, sizes)
		return err
	}
	gk := func(priv *PrivateKey) error {
		return GenerateKeyKISA(priv, rand.Reader, UserProvidedRandomInput)
	}

	testKCDSA(t, L2048N224SHA224, 2048, 224, gp, gk)
	testKCDSA(t, L2048N224SHA256, 2048, 224, gp, gk)
	testKCDSA(t, L2048N256SHA256, 2048, 256, gp, gk)
	testKCDSA(t, L3072N256SHA256, 3072, 256, gp, gk)
}