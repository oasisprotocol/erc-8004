# ERC-8004 on Oasis

This repository is a home for on-chain and off-chain tooling for Oasis
ROFL-powered agents that conform to the [ERC-8004] proposal.

[ERC-8004]: https://eips.ethereum.org/EIPS/eip-8004

## Validator Agent

The validator agent listens to validation request events in the validation
registry and copies over current ROFL metadata as `ValidationResponse`.

To compile and run it locally outside TEE:

```shell
cd validator-agent
make
APPD__MOCK=true \
APPD__MOCK_SIGNING_KEY=your_private_key_on_validation_networks \
API_KEY=your_infura_api_key \
./validator-agent  --config config.testnet.yaml
```

You can run tests for Localnet which mimic the whole flow:

```shell
make test
```

## Validator Agent ROFL

Production validator agent runs inside ROFL TEE. To build and deploy it
yourself:

```shell
oasis rofl init --reset
oasis rofl create
oasis rofl build
echo -n "Your Infura API KEY here" | oasis rofl secret set API_KEY -
oasis rofl update
oasis rofl deploy
```

If you change the config file, don't forget to rebuild container images, push
them to publicly accessible URL and update your `compose.yaml`.
