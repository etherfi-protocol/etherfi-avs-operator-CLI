# etherfi-avs-operator-CLI


## Build

```bash
make build
```

## Run

```bash
./avs-cli
```

---

# AVS Registration

### Prerequisites
- Ethereum RPC endpoint to send transactions.


## Step 1: Request ether.fi team to be registered as a Delegated AVS operator

You will be assigned an operatorID and an operator smart contract that is registered with eigenlayeer
- `operatorId`: AVS operator ID assigned by ether.fi team.
- `operatorAddress`: Eigenlayer operator address, which is managed by ether.fi team.

## Step 2: Follow the instructions for the specific AVS you are registering for
* [Witness Chain](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#witness-chain)
* [EigenDA](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#eigenda)
* [eOracle](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#witness-chain)
* [Brevis](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#witness-chain)
* [Lagrange ZK Coprocessor](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#lagrange-zk-coprocessor)
* [Lagrange State Committees](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#lagrange-state-committees)
* [AltLayer MACH](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#altlayermach)
* [Automata Multi-Prover](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#automata-multi-prover)
* [Hyperlane](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI?tab=readme-ov-file#hyperlane)

---

# Witness Chain

## Operator Flow
In order to run the witnesschain node software you will need to register a watchtower on both mainnet and their L2

### registering operator + watchtower on L1

1. Generate a new ECDSA keypair that will be associated with a witness chain "Watchtower"
2. Sign required inputs for registering watchtower

        // Expose the private key generated above as an environment variable "WATCHTOWER_PRIVATE_KEY"

        // Sign 
        ./avs-cli witness-chain prepare-registration --rpc-url $RPC_URL --operator-id {operator_id}

3. Send the json output of the above command to `restaking@ether.fi`
4. Wait for confirmation from ether.fi team that L1 registration is complete
5. Proceed to L2 watchtower registration below

### registering watchtower on L2

1. Follow the steps at https://docs.witnesschain.com/rollup-watchtower-network-live/for-the-node-operators/watchtower-setup/mainnet-setup#step-3-register-the-operator-on-the-witness-chain-watchtower-network

Generate a separate ECDSA key you control for the value of `operator_private_key`, and request the witnesschain team to whitelist this address on their L2.
Be sure to explicitly set `"eth_rpc_url": ""` (set it to empty value) in your config file so that the cli only registers the watchtower on the L2


2. Notify the ether.fi team that you have completed registration and begin to run witnesschain node software.
Please use the same value as `operator_private_key` in your L1 + L2 config files.


## Ether.fi Admin Flow

1. Request WitnessChain team to whitelist target Operator contract
2. Receive prepared registration json file from target node operator
3. Register the operator contract with witness chain

           ./avs-cli witness-chain register --registration-input witness-input.json --rpc-url $RPC_URL

           // submit resulting output as a gnosis TX via AVS admin gnosis

5. Register the watchtower on L1

           ./avs-cli witness-chain register-watchtower --registration-input witness-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis

---

# Lagrange ZK Coprocessor

## Operator Flow

1. Generate a new key via the lagrange tooling https://docs.lagrange.dev/zk-coprocessor/avs-operators/registration#lagrange-network-avs-key
Be sure to save the `public_key` hex string that will be returned as part of the above command
2. Run the following command and send the output to the ether.fi team via `restaking@ether.fi`

           ./avs-cli lagrangeZK prepare-registration --operator-id {operator_id} --public-key {pubkey_hex}

3. Wait for confirmation from the ether.fi team that your registration is complete
4. Proceed to run the lagrange zk coprocessor node software

## Ether.fi Admin Flow

1. Recieve prepared registration json file from target node operator
2. Ensure target operator contract is whitelisted by lagrange team (Whitelisting was requested for operators 1-12)
3. Register the operator contract with lagrange ZK Coprocessor

           ./avs-cli lagrangeZK register --registration-input lagrangeZK-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis

---

# Lagrange State Committees

## Operator Flow

1. Generate a new BLS and ECSDA key via the lagrange tooling https://docs.lagrange.dev/state-committees/run-node/commands
2. Run the following command and send the output to the ether.fi team via `restaking@ether.fi`

           ./avs-cli lagrangeSC prepare-registration --operator-id {operator_id} --bls-keystore {path_to_keystore} --bls-password {encryption_password} --signer-address {address_of_generated_ecdsa_key}

3. Wait for confirmation from the ether.fi team that your registration is complete
4. Proceed to run the lagrange state committee node software

## Ether.fi Admin Flow

1. Recieve prepared registration json file from target node operator
2. Ensure target operator contract is whitelisted by lagrange team (Whitelisting was requested for operators 1-12)
3. Register the operator contract with lagrange state committees

           ./avs-cli lagrangeSC register --registration-input lagrangeSC-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis

---

# EigenDA

## Operator Flow

1. generate a new BLS keystore using the eigenlayer tooling https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-keys
2. Determine which `quorums` and `socket` you wish to register for
3. Sign digest establishing ownership of your newly generated BLS key

           ./avs-cli eigenda prepare-registration --operator-id {operator_id} --bls-keystore {path_to_keystore} --bls-password {password} --quorums {0,1} --socket {socket}

4. Send the result of the previous command to the ether.fi team via `restaking@ether.fi`
5. Wait for confirmation from the ether.fi team that your registration is complete
6. Proceed to run the eigenDA node software

## Ether.fi Admin Flow

1. Receive prepared registration json file from target node operator
2. Register the operator contract with eigenda

           ./avs-cli eigenda register --registration-input eigenda-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis

---

# Brevis

## Operator Flow

1. generate a new BLS keystore using the eigenlayer tooling https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-keys
2. Determine which `quorums` and `socket` you wish to register for
3. Sign digest establishing ownership of your newly generated BLS key

           ./avs-cli brevis prepare-registration --operator-id {operator_id} --bls-keystore {path_to_keystore} --bls-password {password} --quorums {0,1} --socket {socket}

4. Send the result of the previous command to the ether.fi team via `restaking@ether.fi`
5. Wait for confirmation from the ether.fi team that your registration is complete
6. Proceed to run the brevis node software

## Ether.fi Admin Flow

1. Receive prepared registration json file from target node operator
2. Brevis operates with a strict limit on operator numbers. Prior to registering a new operator, you must organize a time with the brevis team
where they will briefly update the limits. After you get confirmation that they have updated the limits, immediately perform the below registration
3. Register the operator contract with brevis

           ./avs-cli brevis register --registration-input brevis-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis


---

# eOracle

## Operator Flow

1. generate and encrypt a new BLS keystore using the eOracle tooling https://eoracle.gitbook.io/eoracle/operators/registration#generate-a-bls-pair-recommended
2. generate a new ECDSA key pair to serve as your `aliasAddress`
3. Sign digest establishing ownership of your newly generated BLS key

           ./avs-cli eoracle prepare-registration --operator-id {my_operator_id} --bls-keystore {path_to_keystore} --bls-password {keystore_password} --alias-address {alias_address}

4. Send the result of the previous command to the ether.fi team via `restaking@ether.fi`
5. Wait for confirmation from the ether.fi team that your registration is complete
6. Proceed to run the eigenda node software

## Ether.fi Admin Flow

1. Receive prepared registration json file from target node operator
2. Register the operator contract with eoracle

           ./avs-cli eoracle register --registration-input eoracle-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis

3. Ask the eOracle team to manually set the alias address from the input to be associated with the target operator
4. Once the eOracle has confirmed set the alias address, you may proceed to tell the node operator to begin running the node software

---

# AltLayer MACH

## Operator Flow

1. generate and encrypt a BLS keystore using the [EigenLayer CLI](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-and-list-keys)
2. Determine which `quorums` and `socket` you wish to register for
3. Sign digest establishing ownership of your newly generated BLS key

           ./avs-cli altlayer prepare-registration --operator-id {operator_id} --bls-keystore {path_to_keystore} --bls-password {password} --quorums {0,1} --socket {socket}

4. Send the result of the previous command to the ether.fi team via `restaking@ether.fi`
5. Wait for confirmation from the ether.fi team that your registration is complete
6. Follow the [AltLayer Operator Guide](https://docs.altlayer.io/altlayer-documentation/altlayer-facilitated-actively-validated-services/altlayer-mach-avs/operator-guide) skipping the `Opt-in and out of MACH AVS` sections. The `Opt-in` was already handling by the steps above. The `OPERATOR_ECDSA_ADDRESS` is `ecdsaSigner` of ether.fi operator contract you are registering with.

## Ether.fi Admin Flow

1. Receive prepared registration json file from target node operator
2. Register the operator contract with altlayer

           ./avs-cli altlayer register --registration-input eoracle-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis

---

# Automata Multi-Prover

## Operator Flow

1. generate and encrypt a BLS keystore using the [EigenLayer CLI](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-and-list-keys)
2. Sign digest establishing ownership of your newly generated BLS key

           ./avs-cli automata prepare-registration --operator-id {operator_id} --bls-keystore {path_to_keystore} --bls-password {password}

4. Send the result of the previous command to the ether.fi team via `restaking@ether.fi`
5. Wait for confirmation from the ether.fi team that your registration is complete
6. Follow the [Automata Mainnet Runbook](https://github.com/automata-network/multiprover-avs-operator-setup/blob/main/mainnet/README.md) skipping the steps to `Deposit into strategies` and `Opt into Multi-Prover AVS` to run the automata operator node


## Ether.fi Admin Flow

1. Recieve prepared registration json file from target node operator
2. Ensure target operator contract is whitelisted by automata team (Whitelisting was requested for operators 1-12)
3. Register the operator contract with automata

           ./avs-cli automata register --registration-input automata-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis

---

# Hyperlane

## Operator Flow

1. Generate a new ECDSA key using [EigenLayer CLI](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-and-list-keys)
and note the address of the key you generated
2. Run the following command and send the output to the ether.fi team via `restaking@ether.fi`

           ./avs-cli hyperlane prepare-registration --operator-id {operator_id} --avs-signer {address_of_generated_ecdsa_key}

3. Wait for confirmation from the ether.fi team that your registration is complete
4. Proceed to run the hyperlane node software


## Ether.fi Admin Flow

1. Recieve prepared registration json file from target node operator
2. Register the operator contract with Hyperlane

           ./avs-cli hyperlane register --registration-input hyperlane-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis

---

# Adding a new AVS to the CLI

### 1. Understand the complete registration flow for the AVS you are adding
Ether.fi utilizes a contract based operator alongside EIP-1271 signing. Most AVS's do not support this
out of the box. Please confirm that their contracts will be compatible with this scheme.
Many AVS's also utilize different styles of keys/signatures and different numbers of them and even multiple chains.
Figure out which actions need to be taken by the individual node operators and which need to be 
done by an ether.fi admin with the EIP-1271 signing key.
Please take the time to open a PR against https://github.com/etherfi-protocol/avs-smart-contracts/
with a test walking through the entire registration flow. You can find an examples in the `test/` directory of the avs-smart-contracts repo

### 2. Add a new command for your avs to `bin/avs-cli/main.go`
Add your top level command to this file and then implement subcommands in their own package.
For an example see https://github.com/etherfi-protocol/etherfi-avs-operator-CLI/tree/witness-chain/bin/avs-cli/witness-chain
The CLI command should be a simple wrapper that forwards data to the package you implement
in the following step

### 3. Implement core logic in a new package `src/avs/{my_avs}`
Please also place any abi's and generated bindings in this package.
For an example see https://github.com/etherfi-protocol/etherfi-avs-operator-CLI/blob/witness-chain/src/witnesschain/witnesschain.go

### 4. Update readme with registration instructions


---------------------------------------------------------------------


## Contracts
- Code
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperatorsManager.sol
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperator.sol
- Deployment
  - Mainnet: 0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a
  - Holesky: 0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4

## References.
- [on-chain “operator” as a contract](https://etherfi.notion.site/Node-Operator-on-chain-operator-as-a-contract-9e86d3390a9e45df8c088d0c283a7dd1)

