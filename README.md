# mantle-avs-operator-cli
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


## Step 1: Request mantle team to be registered as a Delegated AVS operator

- `operatorAddress`: Eigenlayer operator address, which is managed by mantle team.

## Step 2: Follow the instructions for the specific AVS you are registering for

* [EigenDA](https://github.com/mantle-lsp/mantle-avs-operator-CLI?tab=readme-ov-file#eigenda)
---

# EigenDA

## Operator Flow

1. generate a new BLS keystore using the eigenlayer tooling https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-keys
2. Determine which `quorums` and `socket` you wish to register for. Currently, the mantle only supports quorum=0.
3. Sign digest establishing ownership of your newly generated BLS key

           ./avs-cli eigenda prepare-registration --operator-address {operator_address} --bls-keystore {path_to_keystore} --bls-password {password} --quorums {0} --socket {socket}

4. Send the result of the previous command to the mantle team 
5. Wait for confirmation from the mantle team that your registration is complete
6. Proceed to run the eigenDA node software

## Admin Flow

1. Receive prepared registration json file from target node operator
2. Register the operator contract with eigenda

           ./avs-cli eigenda register --registration-input eigenda-input.json

           // submit resulting output as a gnosis TX via AVS admin gnosis