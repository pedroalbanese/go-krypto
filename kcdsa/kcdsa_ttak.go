package kcdsa

import (
	"io"

	"github.com/pedroalbanese/go-krypto/internal"
	"github.com/pedroalbanese/go-krypto/kcdsa/kcdsattak"
)

// Generate the paramters
// using the prime number generator used in krypto/kcdsa/kcdsattak package.
func GenerateParametersTTAK(params *Parameters, rand io.Reader, sizes ParameterSizes) (seed []byte, count int, err error) {
	domain, err := sizes.domain()
	if err != nil {
		return nil, 0, err
	}
	params.Sizes = sizes

	// p. 13
	for {
		seed, err := internal.ReadBits(seed[:0], rand, domain.B)
		if err != nil {
			return nil, 0, err
		}

		J, err := kcdsattak.GenerateJ(seed, domain.Domain)
		if err != nil {
			continue
		}

		P, Q, count, err := kcdsattak.GeneratePQ(J, seed, domain.Domain)
		if err != nil {
			continue
		}

		_, G, err := kcdsattak.GenerateHG(rand, P, J)
		if err != nil {
			continue
		}

		params.P = P
		params.Q = Q
		params.G = G
		params.Sizes = sizes
		return seed, count, nil
	}
}

// Generate PublicKey and PrivateKey
// using krypto/kcdsa/kcdsattak package.
func GenerateKeyTTAK(priv *PrivateKey, rand io.Reader, userProvidedRandomInput []byte) error {
	if priv.P == nil || priv.Q == nil || priv.G == nil {
		return ErrParametersNotSetUp
	}
	domain, err := priv.Sizes.domain()
	if err != nil {
		return err
	}

	// p.16
	xkey, err := internal.ReadBits(nil, rand, domain.B)
	if err != nil {
		return err
	}

	X, Y, _, _, err := kcdsattak.GenerateXYZ(priv.P, priv.Q, priv.G, userProvidedRandomInput, xkey, domain.Domain)
	if err != nil {
		return err
	}

	priv.X = X
	priv.Y = Y
	return nil
}
