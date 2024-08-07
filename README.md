# etherfi-avs-operator-CLI


## Prerequisites as AVS operator

- Etherereum RPC endpoint to send transactions.

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

## Step 1: Request ether.fi team to be registered as a Delegated AVS operator

Request to ether.fi team to register your wallet address as delegated AVS operator.
Then, ether.fi team will register submitted account to `EtherFiAvsOperatorManager` contract.

After ether.fi team register the information, delegted AVS operator should have the following information.
- `operatorId`: AVS operator ID assigned by ether.fi team.
- `operatorAddress`: Eigenlayer operator address, which is managed by ether.fi team.

## Step 2: Follow the instructions for the specific AVS you are registering for
* [Witness Chain](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI/blob/witness-chain/README.md#witness-chain)
* [EigenDA](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI/blob/witness-chain/README.md#eigenda)

---

# Witness Chain

## Operators
In order to run the witnesschain node software you will need to register a watchtower on both mainnet and their L2

### registering operator + watchtower on L1

1. Generate a new ECDSA keypair that will be associated with a witness chain "Watchtower"
2. Sign required inputs for registering watchtower

        // Expose the private key generated above as an environment variable
        export WATCHTOWER_PRIVATE_KEY={MY_PRIVATE_KEY}

        // Sign 
        ./avs-cli witness-chain prepare-registration --rpc-url $RPC_URL --operator-id {operator_id}

3. Send the json output of the above command to `restaking@ether.fi`
4. Wait for confirmation from ether.fi team that L1 registration is complete
5. Proceed to L2 watchtower registration below

### registering watchtower on L2

1. Follow the steps at https://docs.witnesschain.com/rollup-watchtower-network-live/for-the-node-operators/watchtower-setup/mainnet-setup#step-3.3-registering-the-watchtowers-on-witnesschain-mainnet-l2

Supply a separate ECDSA key you control for the value of `operator_private_key`

2. Notify the ether.fi team that you have completed registration and begin to run witnesschain node software
    

## Ether.fi Admin

1. Request WitnessChain team to whitelist target Operator contract
2. Recieve prepared registration json file from target node operator
3. Register the operator contract with witness chain

            // Expose the ECSDA signing key as an environment variable
            export PRIVATE_KEY={ECSDSA_SIGNING_KEY}

           ./avs-cli witness-chain register --registration-input witness-input.json --rpc-url $RPC_URL

           // submit resulting output as a gnosis TX via AVS admin gnosis

5. Register the watchtower on L1

           ./avs-cli witness-chain register-watchtower --registration-input witness-input.json --rpc-url $RPC_URL

           // submit resulting output as a gnosis TX via AVS admin gnosis

---

# EigenDA

## Operators

1. generate a new BLS keystore using the eigenlayer tooling https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-keys
2. Determine which `quorums` and `socket` you wish to register for
3. Sign digest establishing ownership of your newly generated BLS key

           ./avs-cli eigenda prepare-registration

4. Send the result of the previous command to the ether.fi team via `restaking@ether.fi`
5. Wait for confirmation from the ether.fi team that your registration is complete
6. Proceed to run the eigenda node software

## Ether.fi Admin

1. Recieve prepared registration json file from target node operator
2. Register the operator contract with eigenda

            // Expose the ECSDA signing key as an environment variable
            export PRIVATE_KEY={ECSDSA_SIGNING_KEY}

           ./avs-cli eigenda register --registration-input eigenda-input.json --rpc-url $RPC_URL

           // submit resulting output as a gnosis TX via AVS admin gnosis   

---


## Contracts
- Code
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperatorsManager.sol
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperator.sol
- Deployment
  - Mainnet: 0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a
  - Holesky: 0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4

---

## References.
- [on-chain “operator” as a contract](https://etherfi.notion.site/Node-Operator-on-chain-operator-as-a-contract-9e86d3390a9e45df8c088d0c283a7dd1)

# Adding a new AVS to the CLI

### 1. Understand the complete registration flow for the AVS you are adding
Ether.fi utilizes a contract based operator alongside EIP-1271 signing. Most AVS's do not support this
out of the box. Please confirm that their contracts will be compatible with this scheme.
Many AVS's also utilize different styles of keys/signatures and different numbers of them and even multiple chains.
Figure out which actions need to be taken by the individual node operators and which need to be 
done by an ether.fi admin with the EIP-1271 signing key.
Please take the time to open a PR against https://github.com/etherfi-protocol/avs-smart-contracts/tree/witness-chain
with a test walking through the entire registration flow. You can find an example here https://github.com/etherfi-protocol/avs-smart-contracts/blob/witness-chain/test/WitnessChain.t.sol

### 2. Add a new command for your avs to `bin/avs-cli/main.go`
Add your top level command to this file and then implement subcommands in their own package.
For an example see https://github.com/etherfi-protocol/etherfi-avs-operator-CLI/tree/witness-chain/bin/avs-cli/witness-chain
The CLI command should be a simple wrapper that forwards data to the package you implement
in the following step

### 3. Implement core logic in a new package `src/{my_avs}`
Please also place any abi's and generated bindings in this package.
For an example see https://github.com/etherfi-protocol/etherfi-avs-operator-CLI/blob/witness-chain/src/witnesschain/witnesschain.go
