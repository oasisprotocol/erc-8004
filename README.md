# ERC-8004 on Oasis

This repository is a home of the tooling for [Oasis ROFL-powered] agents that
conform to the [ERC-8004] proposal.

[ERC-8004]: https://eips.ethereum.org/EIPS/eip-8004
[Oasis ROFL-powered]: https://docs.oasis.io/build/rofl/

The validation flow of ROFL-powered agents is the following:

![ERC-8004 validation flow of ROFL-powered agents](./docs/flow.svg)

This repository contains:

- the [Validator Agent](#validator-agent)
- the [ROFL-8004 startup image](#rofl-8004-startup-image)
- a simple [example] using the ROFL-8004 startup image to register ROFL as an 
  ERC-8004 agent and request validation through `ValidationRegistry`

[example]: ./example/compose.yaml

## Validator Agent

The Validator Agent is a simple Go application that listens to 
`ValidationRequest` events reported by the validation registry on a 
specified chain. Then, it connects to the corresponding Oasis chain and copies 
over ROFL metadata of an app and the most recent replica inside the 
`ValidationResponse`.

To compile and run the validator agent locally outside of TEE:

```shell
cd validator-agent
make
APPD__MOCK=true \
APPD__MOCK_SIGNING_KEY=your_private_key_on_validation_networks \
INFURA_API_KEY=your_infura_api_key \
./validator-agent  --config config.testnet.yaml
```

### Validator Agent ROFL

Production validator agent runs inside ROFL TEE. Use the [Oasis CLI] to 
verify that the validator agent instance corresponds to the source code in 
this repository by running:

```shell
oasis rofl build --verify
```

To build and deploy your own validator ROFL:

```shell
oasis rofl init --reset
oasis rofl create
oasis rofl build
echo -n "Your Infura API KEY here" | oasis rofl secret set INFURA_API_KEY -
oasis rofl update
oasis rofl deploy
```

The keypair for submitting `ValidationResponse` to the target chain is derived 
inside the TEE. Its public key is reported inside the Validator Agent's Replica
metadata and in the machine logs. Make sure to fund the account otherwise the 
agent will not be able to submit validation reports.

If you change the config file, don't forget to rebuild container images, push
them to a publicly accessible URL and update your `compose.yaml`.

[Oasis CLI]: https://cli.oasis.io

### Testing

Make sure you have `anvil` installed. Then run localnet tests which mimic the
whole flow:

```shell
cd validator-agent
make test
```

## ROFL-8004 startup image

This is a container that simplifies the registration and validation of your 
ROFL. 

On the first startup it will register your ROFL as an ERC-8004 agent if it 
doesn't exist yet and store the app ID along the agent's metadata.

On every startup it will derive configured public key(s) and upsert them to 
the ROFL replica metadata. Then, it will submit a `ValidationRequest` to the 
validator agent to perform the validation of your app.

A minimal setup to use the container is to include the following service in
your ROFL `compose.yaml`:

```yaml
  rofl-8004:
    image: ghcr.io/oasisprotocol/rofl-8004
    platform: linux/amd64
    environment:
      # RPC for ERC-8004 registry. e.g. https://sepolia.infura.io/v3/<YOUR_KEY>
      - RPC_URL=${RPC_URL}
      # Pinata token for storing token URI when registering new agent.
      - PINATA_JWT=${PINATA_JWT}
    volumes:
      - /run/rofl-appd.sock:/run/rofl-appd.sock
```

Oasis gas fees are covered by the ROFL node hosting your application. Ethereum
transactions are signed with a generated keypair by default. Check out your 
ROFL logs to find out which address to fund.

For full functionality consult the [example].