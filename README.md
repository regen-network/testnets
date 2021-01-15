# Regen Ledger Testnets

Testnets for [Regen Ledger](https://github.com/regen-network/regen-ledger)

## Active Testnets

We are currently supporting the [BigBang-1 Stargate testnet](https://github.com/cosmos/testnets/tree/master/bigbang-1)

## Regen Devnets

### regen-devnet-2

`regen-devnet-2` is active now and here are some important details:

- Explorer: https://devnet.regen.aneka.io
- Faucet: https://faucet.devnet.regen.vitwit.com
- RPC: http://18.220.101.192:26657
- LCD: http://18.220.101.192:1317

Persistent Peer: `f864b879f59141d0ad3828ee17ea0644bdd10e9b@18.220.101.192:26656`

Here are the instructions to run a validator for `regen-devnet-2`:

1. Stop your existing regen validator (if any)
```shell script
sudo service regen stop
```
2. Run the latest setup script
```sh
git clone https://github.com/regen-network/testnets
cd testnets
git pull

chmod +x scripts/devnet-val-setup.sh
./scripts/devnet-val-setup.sh <your_key_name> <your_validator_moniker_name>
```



## Upcoming Testnets

Here are the details for upcoming testnets. Please check our blog post [Regen Network 2020 testnet roadmap](https://link.medium.com/vVBNDosMr4) for more details

### Regen Network Testnet 4000: Aplikiĝo Testnet

**Focus**: Application specific testing and simulation of ecosystem service credit creation and trading with production ready MVP blockchain.

*Estimated Dates: Feb 8th — March 15th 2021

*Total points to be allocated: 1800*

## Historic Testnets (not in use)

The testnets listed below are no longer active but are retained here for posterity. Do not waste your time trying to join them :)

### Regen Network Testnet 3000: COSMWASM Kontraŭa Testnet

**Focus**: Adversarial testnet and network load testing with Regen Ledger running CosmWASM. This testnet may also morph into a Game of Zones testnet, as we are sensitive to the larger community opportunity.

* Testnet schedule: 13th March - 24th Apr, 2020 (6 weeks)
* Total points to be allocated: 1100+
* Gentx submissions: 9th March 1500UTC (start date) - 12th March 0900 UTC (end date)
* Genesis release time: 12th March, 1600UTC (23 hours before genesis time)
* Network start time: 13th March, 1500UTC

#### [Click here to join the testnet](./kontraua/README.md)

### [`Algradigon-1`](https://github.com/regen-network/testnets/tree/modifications/archive/algradigon-1)

`Algradigon-1`, an incentivized testnet was launched on 23rd Jan at 17:00UTC with 38 validators signing on the genesis block. The validator set was increased to a total of 49 validators and total of 4 upgrades were executed in a span of two weeks by 7th Feb.

Total points allocated: 900

**Note**
The incentive program for the testnet was ended on 12th Feb 2020 at 12:00UTC

**Important links**

* Postmortem report of the testnet: https://medium.com/regen-network/postmortem-of-algradigon-1-eb2dc7652850
* Testnet results: https://medium.com/@gregorylandua/algradigon-1-testnet-results-d32d8cfca615

### [`congo-1`](https://github.com/regen-network/testnets/tree/modifications/archive/congo-1)

`congo-1` was the first community-driven testnet between 7th November 2019 and 17th January 2020, at which point it was abandoned in order to prepare for the next incentivised testnet.

### [`regen-test-1001`](https://github.com/regen-network/testnets/tree/modifications/archive/regen-test-1001)

`regen-test-1001` ran between August and October 2019. It is now defunct, having discovered a cosmos-sdk bug in governance-driven parameter updates logic.

### [`regen-test-1000`](https://github.com/regen-network/testnets/tree/modifications/archive/regen-test-1000)

`regen-test-1000` hit some weird consensus error on app state at block 2.

#### [`xrn-test-3`](https://github.com/regen-network/testnets/tree/modifications/archive/xrn-test-3)

Testnet `xrn-test-3` started producing blocks at `2019-03-29T19:44:44.571815638Z` and is now defunct.


`xrncli` can be configured to connect to the testnet as follows:

```sh
xrncli init --chain-id xrn-test-2 --node tcp://xrn-us-east-1.regen.network:26657
```

#### [`xrn-test-2`](https://github.com/regen-network/testnets/tree/modifications/archive/xrn-test-2)

Deployed at `2018-12-19T20:40:06.463846Z`.

#### [`xrn-1`](https://github.com/regen-network/testnets/tree/modifications/archive/xrn-1)

The initial Regen Ledger testnet `xrn-1` was deployed on 2018-12-19.

```

## KYC Utils

### Generate Validator Keys
```sh
git clone https://github.com/regen-network/testnets
cd testnets
git pull

chmod +x scripts/gen_val_keys.sh
./scripts/gen_val_keys.sh <your_key_name>
```
