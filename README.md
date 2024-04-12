# etherfi-avs-operator-CLI

https://etherfi.notion.site/Node-Operator-on-chain-operator-as-a-contract-9e86d3390a9e45df8c088d0c283a7dd1?pvs=4

## Prerequisites

    See sample.env for expected ENV variables

## Register BLS key as Delegated node operator

    ./etherfi register-bls --bls-signature-file keystore/fixtures/fixture.bls.signature.json --operator-id 1 --quorum-numbers 16,32,3 --socket "TEST_SOCKET"

    // Add --broadcast flag to submit transaction onchain

