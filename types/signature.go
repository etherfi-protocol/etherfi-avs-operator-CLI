package types

// AVSBLSSignature is the signature of a message which requires to opt-in to the AVS
type AVSBLSSignature struct {
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
