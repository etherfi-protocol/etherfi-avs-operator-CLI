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

# AVS Registration

### Step 1: Request ether.fi team to be registered as a Delegated AVS operator

Request to ether.fi team to register your wallet address as delegated AVS operator.
Then, ether.fi team will register submitted account to `EtherFiAvsOperatorManager` contract.

After ether.fi team register the information, delegted AVS operator should have the following information.
- `operatorId`: AVS operator ID assigned by ether.fi team.
- `operatorAddress`: Eigenlayer operator address, which is managed by ether.fi team.

### Step 2: Follow the instructions for the specific AVS you are registering for
* [Witness Chain](https://github.com/etherfi-protocol/etherfi-avs-operator-CLI/blob/witness-chain/README.md#witness-chain)

# Witness Chain

## Operators
In order to run the witnesschain node software you will need to register a watchtower on both mainnet and their L2

### registering operator + watchtower on L1

1. Generate a new ECDSA keypair that will be associated with a witness chain "Watchtower"
2. Sign required inputs for registering watchtower

        // Expose the private key generated above as an environment variable
        export WATCHTOWER_PRIVATE_KEY={MY_PRIVATE_KEY}
        
        // Sign 
        ./avs-cli witness-chain prepare-registration --rpc-url $RPC_URL --operator-id {operator_id} --watchtower-address {address_of_generated_watchtower_key}

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

           ./avs-cli witness-chain register --registration-input witness-input.json --rpc-url $RPC_URL

           // submit resulting output as a gnosis TX via AVS admin gnosis
   
5. Register the watchtower on L1
   
           ./avs-cli witness-chain register-watchtower --registration-input witness-input.json --rpc-url $RPC_URL

           // submit resulting output as a gnosis TX via AVS admin gnosis




## Contracts
- Code
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperatorsManager.sol
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperator.sol
- Deployment
  - Mainnet: 0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a
  - Holesky: 0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4


## Ref.
- [on-chain “operator” as a contract](https://etherfi.notion.site/Node-Operator-on-chain-operator-as-a-contract-9e86d3390a9e45df8c088d0c283a7dd1)
