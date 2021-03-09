# Regen Ledger Testnets

Testnets for [Regen Ledger](https://github.com/regen-network/regen-ledger)

## Active Testnets

### 1. Regen Network Testnet 4000: Aplikiĝo Testnet

**Focus**: Application specific testing and simulation of ecosystem service credit creation and trading with production ready MVP blockchain.

*Estimated Dates: Feb 8th — March 15th 2021

*Minimum points to be allocated: 1800*
*$REGEN staking token rewards: 500,000*

Aplikigo-1 Testnet is scheduled to start on 8th Feb, 2021. More details [here](./aplikigo-1)

Blog Post: https://medium.com/regen-network/apliki%C4%9Do-regen-networks-final-pre-launch-incentivized-testnet-2e353dffb4b6

Testnet Plan: [Aplikigo-1 Testnet Plan](./aplikigo-1/PLAN.md)

**TL;DR**
- Explorer: https://aplikigo.regen.aneka.io
- RPC: http://public-rpc1.regen.vitwit.com:26657
- LCD: http://public-rpc1.regen.vitwit.com:1317
- Persistent Peers (Feel free to add yours and create a PR)
```
35d64042e1a5f6466b2b1540fa2e859dfc49666e@public-rpc1.regen.vitwit.com:26656,9082e4d408b0794f884f8c1733d7d11ffb010e38@161.35.51.84:26656,937bbb7a231a870452dd595a442a191d609a26a5@116.203.20.144:26656,
```

### 2. Bigbang Stargate testnet
We are also supporting the [BigBang-1 Stargate testnet](https://github.com/cosmos/testnets/tree/master/bigbang-1)


## Regen Devnets

### regen-devnet-5

`regen-devnet-5` is active now and here are some important details:

- Explorer: https://devnet.regen.aneka.io
- Faucet: https://faucet.devnet.regen.vitwit.com
- RPC: http://18.220.101.192:26657
- LCD: http://18.220.101.192:1317

Persistent Peer: `b2679a74d6bd9f89a3c294c447d6930293255e6b@18.220.101.192:26656`

Here are the instructions to run a validator for `regen-devnet-5`:

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


## KYC Utils

### Generate Validator Keys
Linking a validator address to your identity is how we ensure the right validator is rewarded for the hard work of participating in our incentivized testnets.  Please back up your keys and maintain the same keys thorughout testnet operations if possible.  If you have a key management issue, please use the same Moniker in generation of new keys and notify the team.

The same validator keys can be used for different testnets, and even for main net (as long as you practice good key management). 

### Handy Script
Here are instructions for generating keys for regen ledger. This is well-tested on Ubuntu 18.04, if you are using different arch and running into issues, please use manual key generation instructions below
```sh
git clone https://github.com/regen-network/testnets
cd testnets
git pull

chmod +x scripts/gen_val_keys.sh
./scripts/gen_val_keys.sh <your_key_name>
```
### Generate regen keys manually
Step-1: Install Go 1.15.x (Optional)
```sh
  $ sudo apt update
  $ sudo apt install build-essential jq -y

  $ wget https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz
  $ tar -xvf go1.15.2.linux-amd64.tar.gz
  $ sudo mv go /usr/local

  $ echo "" >> ~/.bashrc
  $ echo 'export GOPATH=$HOME/go' >> ~/.bashrc
  $ echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
  $ echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
  $ echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc

  $ source ~/.bashrc
  $ go version # should print 1.15.2
```

Step-2: Install Regen (from source)
```sh
$ git clone https://github.com/regen-network/regen-ledger
$ cd regen-ledger
$ git checkout v0.6.0-alpha6
$ make install
```

Step-3: Create your key (keyname can be anything):

```sh
$ regen keys add <your_key_name>
```
Use the address generated from above command to fill your KYC
