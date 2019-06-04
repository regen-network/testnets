# Regen Ledger Testnets

Testnets for [Regen Ledger](https://github.com/regen-network/regen-ledger)

## Join `regen-test-1000` Public Testnet

We are currently preparing the `regen-test-1000` public testnet launch.

Here is the planned timing for the testnet launch:
- 2019-06-05 18:00UTC (11am PDT/2pm EDT/8pm CEST) - Genesis time

The genesis files are in [./latest](latest)

We have airdropped testnet tokens to all participants in the last Cosmos Hub
testnet. If you would like to participate and have not received tokens, you
can open an issue with an address and we'll send you some.

## How to Run a Testnet Validator
Thank you to Jay from CypherCore who wrote this [short and sweet tutorial on setting up a testnet validator](https://www.notion.so/jim380/Regen-Ledger-Node-Set-up-67ede4023f1b45f99724ac52386a1130)
to be included in genesis.

## Testnet Status

### `xrn-test-3`

Testnet `xrn-test-3` started producing blocks at `2019-03-29T19:44:44.571815638Z` and is live as of this writing.

In this testnet, validator nodes currently have ports 26656, 26657 and 1317 open for testing purposes. In the future,
the testnet will be setup with more security hardening via sentry and seed nodes.

The validator node URL's are as follows:

* [xrn-us-east-1.regen.network](http://xrn-us-east-1.regen.network:26657)
* [xrn-us-west-1.regen.network](http://xrn-us-west-1.regen.network:26657)
* [xrn-eu-central-1.regen.network](http://xrn-eu-central-1.regen.network:26657)

`xrncli` can be configured to connect to the testnet as follows:

```sh
xrncli init --chain-id xrn-test-2 --node tcp://xrn-us-east-1.regen.network:26657
```

### `xrn-test-2`

Deployed at `2018-12-19T20:40:06.463846Z`.

### `xrn-1`

The initial Regen Ledger testnet `xrn-1` was deployed on 2018-12-19.

```
