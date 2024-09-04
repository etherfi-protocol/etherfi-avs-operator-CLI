package types

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
)

type SignatureWithSaltAndExpiry struct {
	Signature []byte   `json:"signature,omitempty"`
	Salt      [32]byte `json:"salt,omitempty"`
	Expiry    *big.Int `json:"expiry,omitempty"`
}

// BLSPubkeyRegistrationParams is the signature of a message which requires to opt-in to the AVS
type BLSPubkeyRegistrationParams struct {
	G1 struct {
		X *big.Int `json:"x,omitempty"`
		Y *big.Int `json:"y,omitempty"`
	} `json:"g1,omitempty"`
	G2 struct {
		X [2]*big.Int `json:"x,omitempty"`
		Y [2]*big.Int `json:"y,omitempty"`
	} `json:"g2,omitempty"`
	Signature struct {
		X *big.Int `json:"x,omitempty"`
		Y *big.Int `json:"y,omitempty"`
	} `json:"signature,omitempty"`
}

func (s *BLSPubkeyRegistrationParams) Load(g1 *bn254.G1Affine, g2 *bn254.G2Affine, sig *bn254.G1Affine) {
	s.G1.X = g1.X.BigInt(new(big.Int))
	s.G1.Y = g1.Y.BigInt(new(big.Int))

	// WARNING: 0 and 1 are reversed
	s.G2.X[1] = g2.X.A0.BigInt(new(big.Int))
	s.G2.X[0] = g2.X.A1.BigInt(new(big.Int))
	s.G2.Y[1] = g2.Y.A0.BigInt(new(big.Int))
	s.G2.Y[0] = g2.Y.A1.BigInt(new(big.Int))

	s.Signature.X = sig.X.BigInt(new(big.Int))
	s.Signature.Y = sig.Y.BigInt(new(big.Int))
}

func (s *BLSPubkeyRegistrationParams) ExportSignature() *bn254.G1Affine {
	sig := new(bn254.G1Affine)
	sig.X.SetBigInt(s.Signature.X)
	sig.Y.SetBigInt(s.Signature.Y)
	return sig
}
