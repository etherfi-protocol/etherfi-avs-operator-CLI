package othentic

import (
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
)

// XXX: The othentic signing server performs an eip-1271 signature check before returning
// the token so this test will only work if provided the real currently configured
// signing key for the operator contracts
func TestAcquireAuthToken(t *testing.T) {

	// load eip-1271 admin signing key
	signingKey, err := crypto.HexToECDSA(os.Getenv("ADMIN_1271_SIGNING_KEY"))
	if err != nil {
		t.Fatal(err)
	}

	operator := etherfi.Operator{Address: common.HexToAddress("0xfB487f216CA24162119C0C6Ae015d680D7569C2f")} // operator #1
	//avsGovernance := common.HexToAddress("0xB3e069FD6dDA251AcBDE09eDa547e0AB207016ee")                       // ungate AVS
	avsGovernance := common.HexToAddress("0x6f943318b05AD7c6EE596A220510A6D64B518dd8") // redstone AVS

	token, err := AcquireSignedAuthToken(signingKey, &operator, avsGovernance)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("token: %s\n", hex.EncodeToString(token))
}
