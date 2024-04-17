package signer

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/consensys/gnark-crypto/ecc/bn254"
)

type BLSAvsSigner interface {
	Sign(g1MsgToSign *bn254.G1Affine) (*bn254.G1Affine, error)
	Verify(msg *bn254.G1Affine, signature *bn254.G1Affine) (bool, error)
}

type eigenDABLSSigner struct {
	keypair *bls.KeyPair
}

func (s *eigenDABLSSigner) Sign(g1MsgToSign *bn254.G1Affine) (*bn254.G1Affine, error) {
	sig := s.keypair.SignHashedToCurveMessage(g1MsgToSign)
	return sig.G1Affine, nil
}

func (s *eigenDABLSSigner) Verify(msg *bn254.G1Affine, signature *bn254.G1Affine) (bool, error) {
	var negSig bn254.G1Affine
	negSig.Neg(signature)

	g2Gen := new(bn254.G2Affine)
	g2Gen.X.SetString("10857046999023057135944570762232829481370756359578518086990519993285655852781",
		"11559732032986387107991004021392285783925812861821192530917403151452391805634")
	g2Gen.Y.SetString("8495653923123431417604973247489272438418190587263600148770280649306958101930",
		"4082367875863433681332203403145435568316851327593401208105741076214120093531")

	P := [2]bn254.G1Affine{*msg, negSig}
	Q := [2]bn254.G2Affine{*s.keypair.GetPubKeyG2().G2Affine, *g2Gen}

	return bn254.PairingCheck(P[:], Q[:])
}

func NewAVSSigner(keypair *bls.KeyPair) BLSAvsSigner {
	return &eigenDABLSSigner{keypair: keypair}
}
