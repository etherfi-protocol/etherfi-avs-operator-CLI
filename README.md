# etherfi-avs-operator-CLI

## Contracts
- Code
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperatorsManager.sol
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperator.sol
- Deployment
  - Mainnet: 0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a
  - Holesky: 0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4

## Prerequisites

- Environment file for expected ENV variables(See sample.env).
- ECDSA key file which is a account to be delegated by ether.fi team.
- BLS key file which is used for AVS operation.
- Etherereum RPC endpoint to send transactions.

## Scenario

1. AVS operators prepare their BLS and a wallet (EOA or gnosis safe).
2. ether.fi team registers the AVS operator's wallet address as Delegated AVS operator (or AVS runner).
3. AVS operators register their BLS key to `EtherFiAvsOperator' contract.
4. ether.fi team trigger `registerOperator` function with their ECDSA signature to register the node to the AVS.

## Workflows

### Step 1: Build the binary

Simply run `make build` to build the binary. The binary will be placed in the `dist` directory.

```bash
~$ make build
~$ tree dist
dist
└── etherfi
```

### Step 2: Request to register Delegated AVS operator

Request to ether.fi team to register the ECDSA key as delegated AVS operator.
Then, ether.fi team will register submitted account to `ether.fiAvsOperatorManager` contract.

After ether.fi team register the information, delegted AVS operator should have the following information.
- `operatorId`: AVS operator ID assigned by ether.fi team.
- `operatorAddress`: Eigenlayer operator address, which is managed by ether.fi team.

### Step 3: Create BLS signature

If *delegated AVS operator* decide to opt-in to AVS through ether.fi contract,
delegated AVS operator should submit BLS signature to `ether.fiAvsOperator` contract.

First, the AVS operator should create BLS signature using BLS key file.
To create BLS signaure, *service manager* address is required,
which is provided by [AVS Dashboard](https://app.eigenlayer.xyz/avs/0x870679e138bcdf293b7ff14dd44b70fc97e12fc0).

If all the required information is ready, then run the following command to create BLS signature.

```bash
~$ ./dist/etherfi create-bls \
  --bls-key-file "<Key Filepath>" \
  --bls-key-password "<Key password>" \
  --service-manager "<Service Manager Address of AVS>" \
  --eigenlayer-operator "<Address of Eigenlayer Operator>" \
  --rpc "<Ethereum RPC Endpoint>"
```

Example of the command like below.

```bash
~$ ./dist/etherfi create-bls-signature \
  --bls-key-file "./keystore/fixtures/fixture.bls.json" \
  --bls-key-password "password@1234" \
  --service-manager "0xd4a7e1bd8015057293f0d0a557088c286942e84b" \
  --eigenlayer-operator "0x53F69255A16E6a924665EB839f2730bFF01BE7D8" \
  --rpc https://ethereum-holesky-rpc.publicnode.com
```

If the command is successful, then the BLS signature file will be created in the current directory.

```bash
~$ cat b1585827df842b01be31601d3b09fa094a40a59a2bd67e1a1cf8fad52a8bad86.json | jq .
{
  "g1": {
    "x": "11309115092513905692011766794268611533060860132450390416706713417007017003290",
    "y": "10099038639465945853355841479849128282561908322481975349946661108792552165464"
  },
  "g2": {
    "x": [
      "9999874916423263486182869108140729809350389216313891965609964113333324489295",
      "1748646664750896244090849000611144921186133596533213003220432661350444049387"
    ],
    "y": [
      "12385645521204038946454065367448594647484214586612685207057210041675620555895",
      "11847573064610017387949012476285891282019683730026575197328924744021789701473"
    ]
  },
  "signature": {
    "x": "20492878901239690531093557663058547107183361356438727967873867738966365432584",
    "y": "13238159785824621533432502787105333225248656315234010974670933048183046701631"
  }
}
```

### Step 4: Register BLS signature

The BLS signature file created on the previous step should be registered on-chain.

```bash
~$ ./dist/etherfi register-bls \
		--bls-signature-file "<BLS signature file created on the previous stpe>" \
		--operator-id <Deleted AVS Operator ID assigned by ether.fi> \
		--quorum-numbers <Quorum code, 0> \
		--socket "not yet decided" \
		--broadcast \
		--chain-id <Chain ID, 1: mainnet, 17000: holesky>
```

Example of the command like below.

```bash

#!/bin/bash
source .env

~$ ./dist/etherfi register-bls \
		--bls-signature-file "b1585827df842b01be31601d3b09fa094a40a59a2bd67e1a1cf8fad52a8bad86.json" \
		--operator-id 1 \
		--quorum-numbers 0x0000 \
		--socket "not yet decided" \
		--broadcast \
		--chain-id 17000

~$ bash register_eigenda.sh
raw tx: b9027802f902748242680d8404906e0385010e48ff858301d2cd94df9679e8bfce22ae503fd2726cb1218a18cd8bf480b90204c74e7ae5000000000000000000000000000000000000000000000000000000000000000100000000000000000000000053012c69a189cfa2d9d29eb6f19b32e0a2ea3490000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001c0304bc6600485c4c85e020f5f2ce1bde9fec6df8b02b7cda506cecf2fc9e69cce163cfd445e4da14a229e1fa771dc84577b637aacfbff373fff80f9b29113f47c1a8ce2b5d15f311972826258f59e0edad1d5880f48dfc18f4a9d25d1601ccc240507c97e361ef729d74db6f03762cbdb3b3b2dcb0a4d6fb5af81db0162cc6a650c65fb0713560e5c4aca8f1f113579d80116f12c5b365ab0a1a57ca69b875493238a86dc749e842699a945028e361ff7a648589cee7ae4067fe7224db2d070b727c253813f2db2a1055dbf60745f831dee11890366f5bc5c9518bb642a899ba62e4f6e4f0829b13a3b48de33070ad43f6bcf4e143dc9d76feea05c2e6d0187f700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f6e6f742079657420646563696465640000000000000000000000000000000000c080a087d531c1181ddef62ecaa53a9383502bd9ffa179ed8524ce66cd781beb78bb6fa04f7b6316bea822c764470adba68bd59f27deff91fc63b8bf268e1a8386896fd1
```

### Step 5: registerOperator

After all the above steps are completed,
ether.fi team should trigger `registerOperator` function to register BLS key which is submitted by AVS operator.

## Ref.
- [on-chain “operator” as a contract](https://etherfi.notion.site/Node-Operator-on-chain-operator-as-a-contract-9e86d3390a9e45df8c088d0c283a7dd1)
