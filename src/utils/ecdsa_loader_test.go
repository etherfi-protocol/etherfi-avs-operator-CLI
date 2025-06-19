package utils

import (
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestLoadECDSAKey_FromEnvironmentVariable(t *testing.T) {
	// Test private key (this is a test key, not for production use)
	testPrivateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	envVar := "TEST_PRIVATE_KEY"

	// Set environment variable
	os.Setenv(envVar, testPrivateKeyHex)
	defer os.Unsetenv(envVar)

	// Create mock command that returns empty strings for keystore flags
	testCmd := &mockCommand{values: map[string]string{}}

	privateKey, err := LoadECDSAKey(testCmd, envVar)

	assert.Nil(t, err)
	assert.NotNil(t, privateKey)

	// Verify the loaded key matches the test key
	expectedKey, _ := crypto.HexToECDSA(testPrivateKeyHex)
	assert.Equal(t, expectedKey.D, privateKey.D)
}

func TestLoadECDSAKey_FromKeystore(t *testing.T) {
	// Create a mock command with keystore flags set
	keystorePath := "../keystore/fixtures/fixture.ecdsa.key.json"
	keystorePassword := "fixture@test1234"
	expectedAddress := "0x287B703F25CE707D7974D26DBE5B78121f70794f"

	testCmd := &mockCommand{values: map[string]string{
		"ecdsa-keystore": keystorePath,
		"ecdsa-password": keystorePassword,
	}}

	privateKey, err := LoadECDSAKey(testCmd, "UNUSED_ENV_VAR")

	assert.Nil(t, err)
	assert.NotNil(t, privateKey)

	// Verify the loaded key generates the expected address
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	assert.Equal(t, strings.ToLower(expectedAddress), strings.ToLower(address.Hex()))
}

func TestLoadECDSAKey_KeystoreWithoutPassword(t *testing.T) {
	keystorePath := "../keystore/fixtures/fixture.ecdsa.key.json"

	// Mock command with keystore but no password
	testCmd := &mockCommand{values: map[string]string{
		"ecdsa-keystore": keystorePath,
	}}

	privateKey, err := LoadECDSAKey(testCmd, "UNUSED_ENV_VAR")

	assert.NotNil(t, err)
	assert.Nil(t, privateKey)
	assert.Contains(t, err.Error(), "--ecdsa-password is required")
}

func TestLoadECDSAKey_InvalidKeystoreFile(t *testing.T) {
	invalidKeystorePath := "../keystore/fixtures/nonexistent.json"
	keystorePassword := "fixture@test1234"

	testCmd := &mockCommand{values: map[string]string{
		"ecdsa-keystore": invalidKeystorePath,
		"ecdsa-password": keystorePassword,
	}}

	privateKey, err := LoadECDSAKey(testCmd, "UNUSED_ENV_VAR")

	assert.NotNil(t, err)
	assert.Nil(t, privateKey)
	assert.Contains(t, err.Error(), "loading ECDSA key from keystore")
}

func TestLoadECDSAKey_InvalidPassword(t *testing.T) {
	keystorePath := "../keystore/fixtures/fixture.ecdsa.key.json"
	wrongPassword := "wrongpassword"

	testCmd := &mockCommand{values: map[string]string{
		"ecdsa-keystore": keystorePath,
		"ecdsa-password": wrongPassword,
	}}

	privateKey, err := LoadECDSAKey(testCmd, "UNUSED_ENV_VAR")

	assert.NotNil(t, err)
	assert.Nil(t, privateKey)
	assert.Contains(t, err.Error(), "loading ECDSA key from keystore")
}

func TestLoadECDSAKey_NoKeystoreNoEnvVar(t *testing.T) {
	envVar := "NONEXISTENT_ENV_VAR"

	// Ensure env var is not set
	os.Unsetenv(envVar)

	testCmd := &mockCommand{values: map[string]string{}}

	privateKey, err := LoadECDSAKey(testCmd, envVar)

	assert.NotNil(t, err)
	assert.Nil(t, privateKey)
	assert.Contains(t, err.Error(), "must provide either --ecdsa-keystore flag or set")
}

func TestLoadECDSAKey_InvalidEnvironmentVariable(t *testing.T) {
	invalidPrivateKeyHex := "invalidhex"
	envVar := "INVALID_PRIVATE_KEY"

	os.Setenv(envVar, invalidPrivateKeyHex)
	defer os.Unsetenv(envVar)

	testCmd := &mockCommand{values: map[string]string{}}

	privateKey, err := LoadECDSAKey(testCmd, envVar)

	assert.NotNil(t, err)
	assert.Nil(t, privateKey)
	assert.Contains(t, err.Error(), "invalid private key")
}

func TestGetECDSAKeystoreFlags(t *testing.T) {
	flags := GetECDSAKeystoreFlags()

	assert.Len(t, flags, 2)

	// Check that the flags have the expected names
	flagNames := make([]string, len(flags))
	for i, flag := range flags {
		flagNames[i] = flag.Names()[0]
	}

	assert.Contains(t, flagNames, "ecdsa-keystore")
	assert.Contains(t, flagNames, "ecdsa-password")
}

// Mock implementation for testing
type mockCommand struct {
	values map[string]string
}

func (mc *mockCommand) String(name string) string {
	if val, exists := mc.values[name]; exists {
		return val
	}
	return ""
}