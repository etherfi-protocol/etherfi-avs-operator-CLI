package types

import (
	"github.com/consensys/gnark-crypto/ecc/bn254"
)

// BLSPubkeyRegistrationParams is the signature of a message which requires to opt-in to the AVS
type BLSPubkeyRegistrationParams struct {
	G1 struct {
		X string `json:"x,omitempty"`
		Y string `json:"y,omitempty"`
	} `json:"g1,omitempty"`
	G2 struct {
		X [2]string `json:"x,omitempty"`
		Y [2]string `json:"y,omitempty"`
	} `json:"g2,omitempty"`
	Signature struct {
		X string `json:"x,omitempty"`
		Y string `json:"y,omitempty"`
	} `json:"signature,omitempty"`
}

func (s *BLSPubkeyRegistrationParams) Load(g1 *bn254.G1Affine, g2 *bn254.G2Affine, sig *bn254.G1Affine) {
	s.G1.X = g1.X.String()
	s.G1.Y = g1.Y.String()

	// WARNING: 0 and 1 are reversed
	s.G2.X[1] = g2.X.A0.String()
	s.G2.X[0] = g2.X.A1.String()
	s.G2.Y[1] = g2.Y.A0.String()
	s.G2.Y[0] = g2.Y.A1.String()

	s.Signature.X = sig.X.String()
	s.Signature.Y = sig.Y.String()
}

func (s *BLSPubkeyRegistrationParams) ExportSignature() *bn254.G1Affine {
	sig := new(bn254.G1Affine)
	sig.X.SetString(s.Signature.X)
	sig.Y.SetString(s.Signature.Y)
	return sig
}
