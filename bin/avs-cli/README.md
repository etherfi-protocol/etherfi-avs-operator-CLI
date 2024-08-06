# How to use etherfi avs command

This CLI tool will create `PubkeyRegistrationParams` which is required to opt-in AVS on Eigenlayer.

## Prequisites

- `pubkeyRegistrationMessageHash`

Call `RegisterCoordinator.pubkeyRegistrationMessageHash` by providing Eigenlayer Operator address.
Then, G1Point format value will be returned. This value should be passed to `--reg-msg` option.

[EigenDA RegistryCoordinator](https://holesky.etherscan.io/address/0x53012C69A189cfA2D9d29eb6F19B32e0A2EA3490#readProxyContract)

- BLS key to register

BLS key file required to create `PubkeyRegistrationParams`.
This key file can be created by `eigenlayer` CLI which can be found [eigenlayer-cli](https://github.com/Layr-Labs/eigenlayer-cli)

## Build

Simply type `~$ make build`, then built binary will be created in `dist` directory.

## Run

Run,
```
~$ ./dist/etherfi avs \
    --reg-msg <pubkeyRegistrationMessageHash> \
    --bls-key-file <BLS key file> \
    --bls-key-password <Password of keyfile>

~$ ./dist/etherfi avs \
    --bls-key-file <BLS key file> \
    --bls-key-password <Password of keyfile>
    --service-manager <Contract Address of AVS Service Manager> \
    --eigenlayer-operator <Registered Eigenlayer Operator> \
    --rpc <EL RPC>

```

Example:
```
~$ ./dist/etherfi avs \
    --bls-key-file  ./keyfile.bls.json \
    --bls-key-password thisisyourpassword \
    --service-manager "0xd4a7e1bd8015057293f0d0a557088c286942e84b" \
    --eigenlayer-operator "0x53F69255A16E6a924665EB839f2730bFF01BE7D8" \
    --rpc https://ethereum-holesky-rpc.publicnode.com

```

Then, the new json file will be created.
Example of created json file seems like below

```
{
  "g1": {
    "x": "15218112579367957411928954156670379256187395672498320260043799151083351607878",
    "y": "1829093517312581359150964593061338525079957054057037886068815759612114045222"
  },
  "g2": {
    "x": [
      "2627565466878260585387522540939104292574753779200643920609949107130139616193",
      "7261989353910828479771287209995474709109618234767145072862270353781610454899"
    ],
    "y": [
      "5149090414582510750269480612330833866066911566206397912648971332104940904083",
      "10300083260136129484702287768515135649716536253731847889418957492428969904730"
    ]
  },
  "signature": {
    "x": "17587541414338937490318634396336343655022504812039872978230542318938302397665",
    "y": "3494028518829950357281653456620060411555755988257320462632996122767491526394"
  }
}
```

## TODO

- Retrieve `pubkeyRegistrationMessageHash` from contract.
- Create operator signature.
- Send transaction to store BLS signature.
