package utils

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/keystore"
	"github.com/urfave/cli/v3"
)

// LoadECDSAKey loads an ECDSA private key from keystore file if provided,
// otherwise falls back to environment variable for backward compatibility
func LoadECDSAKey(cmd *cli.Command, envVarName string) (*ecdsa.PrivateKey, error) {
	keystoreFile := cmd.String("ecdsa-keystore")
	keystorePassword := cmd.String("ecdsa-password")

	// If keystore flags are provided, load from keystore
	if keystoreFile != "" {
		if keystorePassword == "" {
			return nil, fmt.Errorf("--ecdsa-password is required when --ecdsa-keystore is provided")
		}

		ks := keystore.NewKeystoreV3()
		privateKey, err := ks.LoadECDSA(keystoreFile, keystorePassword)
		if err != nil {
			return nil, fmt.Errorf("loading ECDSA key from keystore: %w", err)
		}
		return privateKey, nil
	}

	// Fall back to environment variable for backward compatibility
	privateKeyHex := os.Getenv(envVarName)
	if privateKeyHex == "" {
		return nil, fmt.Errorf("must provide either --ecdsa-keystore flag or set %s environment variable", envVarName)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key in %s: %w", envVarName, err)
	}

	return privateKey, nil
}

// GetECDSAKeystoreFlags returns the standard ECDSA keystore flags
func GetECDSAKeystoreFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "ecdsa-keystore",
			Usage: "Path to ECDSA keystore file (alternative to environment variable)",
		},
		&cli.StringFlag{
			Name:  "ecdsa-password",
			Usage: "Password for ECDSA keystore file",
		},
	}
}