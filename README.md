# Regen Ledger Testnets

Testnets for [Regen Ledger](https://github.com/regen-network/regen-ledger)

## Active Testnets

### Regen Network Testnet 3000: COSMWASM Kontraŭa Testnet

**Focus**: Adversarial testnet and network load testing with Regen Ledger running CosmWASM. This testnet may also morph into a Game of Zones testnet, as we are sensitive to the larger community opportunity.

* Testnet schedule: 13th March - 17th Apr, 2020 (5 weeks)
* Total points to be allocated: 1100+
* Gentx submissions: 9th March 1500UTC (start date) - 12th March 0900 UTC (end date)
* Genesis release time: 12th March, 1600UTC (23 hours before genesis time)
* Network start time: 13th March, 1500UTC

#### [Click here to join the testnet](./kontraua/README.md)


## Upcoming Testnets

Here are the details for upcoming testnets. Please check our blog post [Regen Network 2020 testnet roadmap](https://link.medium.com/vVBNDosMr4) for more details

### Regen Network Testnet 4000: Aplikiĝo Testnet

**Focus**: Application specific testing and simulation of ecosystem service credit creation and trading with production ready MVP blockchain.

*Estimated Dates: April 4th week*

*Total points to be allocated: 800*

## Historic Testnets (not in use)

The testnets listed below are no longer active but are retained here for posterity. Do not waste your time trying to join them :)

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
