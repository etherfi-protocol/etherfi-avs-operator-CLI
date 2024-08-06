# etherfi-avs-operator-CLI


## Prerequisites as AVS operator

- Etherereum RPC endpoint to send transactions.

## Build

Simply run `make build` to build the binary. The binary will be placed in the `dist` directory.

```bash
~$ make build
~$ tree dist
dist
└── etherfi
```

# AVS Registration

## Step 1: Request ether.fi team to be registered as a Delegated AVS operator

Request to ether.fi team to register your wallet address as delegated AVS operator.
Then, ether.fi team will register submitted account to `EtherFiAvsOperatorManager` contract.

After ether.fi team register the information, delegted AVS operator should have the following information.
- `operatorId`: AVS operator ID assigned by ether.fi team.
- `operatorAddress`: Eigenlayer operator address, which is managed by ether.fi team.

## Step 2: Follow the instructions for the specific AVS you are registering for

# Witness Chain

## Operators

### Step 1: Generate a new ECDSA keypair that will be associated with a witness chain "Watchtower"

### Step 2: Prepare + Sign inputs required for registering

    // Expose the private key generated above as an environment variable
    export WATCHTOWER_PRIVATE_KEY={MY_PRIVATE_KEY}

    ./avs-cli witness-chain prepare-registration --rpc-url $RPC_URL --operator-id {operator_id} --watchtower-address {address_of_generated_watchtower_key}

### Step 3: Send the json output of the above command to the ether.fi team at restaking@ether.fi

### Step 4: 

## Ether.fi Admin

### Step 1: Request WitnessChain team to whitelist target Operator contract

### Step 2: Recieve prepared registration json file from target node operator



## Contracts
- Code
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperatorsManager.sol
  - https://github.com/etherfi-protocol/smart-contracts/blob/syko/feature/etherfi_avs_operator/src/EtherFiAvsOperator.sol
- Deployment
  - Mainnet: 0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a
  - Holesky: 0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4


## Ref.
- [on-chain “operator” as a contract](https://etherfi.notion.site/Node-Operator-on-chain-operator-as-a-contract-9e86d3390a9e45df8c088d0c283a7dd1)
