# Regen Ledger Testnets

Testnets for [Regen Ledger](https://github.com/regen-network/regen-ledger)

## Active Testnet

### regen-redwood-1

`regen-redwood-1` is an active testnet publicly available for the community. It runs the same software as mainnet and is useful for testing third-party integrations and also for validators who are willing to integrate their setups on Regen Network Mainnet, and want a stable testing environment. It will be used as a platform for testing mainnet upgrades in the future.

- [Faucet](./redwood-testnet/faucet.txt)
- [Explorer](https://redwood.regen.aneka.io/)

Here are the instructions to run a validator for `regen-redwood-1`:

1. Stop your existing regen validator (if any)
```shell script
sudo service regen stop
```
2. Run the latest setup script
```sh
git clone https://github.com/regen-network/testnets
cd testnets
git pull

chmod +x scripts/testnet-val-setup.sh
./scripts/testnet-val-setup.sh <your_key_name> <your_validator_moniker_name>
```

You can find the genesis file and all the relevant information in [redwood-testnet](./redwood-testnet) directory

To get tokens from the faucet use the following curl command
```
curl http://<faucet-url>/faucet/<your-regen-address>
```

## Regen Devnet

### regen-hambach-1

`regen-hambach-1` is an active devnet with experimental modules enabled running the latest tagged [release](https://github.com/regen-network/regen-ledger/releases). It is being used for integrations and testing by the Registry development team, as well as 3rd parties to test out new features like the ecocredit module.

- [Faucet](./hambach-devnet/faucet.txt)
- [Explorer](https://hambach.regen.aneka.io/)

You can find the genesis file and all the relevant information in [hambach-devnet](./hambach-devnet) directory

To get tokens from the faucet use the following curl command
```
curl http://<faucet-url>/faucet/<your-regen-address>
```
## Temporary testnet

### regen-devnet-6

`regen-devnet-6-9` is an internal testnet which is being used to host the latest code from `master` branch of regen-ledger. It is intended for rapid teardown & periodic restarts as bugs are identified and fixed for internal audit of the code.

 - RPC:- http://167.71.178.81:26657/

## Historic Testnets (not in use)

The testnets listed below are no longer active but are retained here for posterity. Do not waste your time trying to join them :)

### Regen Network Testnet 4000: Aplikiĝo Testnet

**Focus**: Application specific testing and simulation of ecosystem service credit creation and trading with production ready MVP blockchain.

*Minimum points allocated: 1800*

*$REGEN staking token rewards: 500,000*

Aplikigo-1 Testnet ran from 8th Feb, 2021 to . More details [here](./aplikigo-1)

Blog Post: https://medium.com/regen-network/apliki%C4%9Do-regen-networks-final-pre-launch-incentivized-testnet-2e353dffb4b6

Testnet Plan: [Aplikigo-1 Testnet Plan](./aplikigo-1/PLAN.md)


### Bigbang Stargate testnet
We also supported the [BigBang-1 Stargate testnet](https://github.com/cosmos/testnets/tree/master/bigbang-1)

### Regen Network Testnet 3000: COSMWASM Kontraŭa Testnet

**Focus**: Adversarial testnet and network load testing with Regen Ledger running CosmWASM. This testnet may also morph into a Game of Zones testnet, as we are sensitive to the larger community opportunity.

* Testnet schedule: 13th March - 24th Apr, 2020 (6 weeks)
* Total points allocated: 1100+

### [`Algradigon-1`](./archive/algradigon-1)

`Algradigon-1`, an incentivized testnet was launched on 23rd Jan at 17:00UTC with 38 validators signing on the genesis block. The validator set was increased to a total of 49 validators and total of 4 upgrades were executed in a span of two weeks by 7th Feb.

Total points allocated: 900

**Note**
The incentive program for the testnet was ended on 12th Feb 2020 at 12:00UTC

**Important links**

* Postmortem report of the testnet: https://medium.com/regen-network/postmortem-of-algradigon-1-eb2dc7652850
* Testnet results: https://medium.com/@gregorylandua/algradigon-1-testnet-results-d32d8cfca615

### [`congo-1`](./archive/congo-1)

`congo-1` was the first community-driven testnet between 7th November 2019 and 17th January 2020, at which point it was abandoned in order to prepare for the next incentivised testnet.

### [`regen-test-1001`](./archive/regen-test-1001)

`regen-test-1001` ran between August and October 2019. It is now defunct, having discovered a cosmos-sdk bug in governance-driven parameter updates logic.

### [`regen-test-1000`](./archive/regen-test-1000)

`regen-test-1000` hit some weird consensus error on app state at block 2.

#### [`xrn-test-3`](./archive/xrn-test-3)

Testnet `xrn-test-3` started producing blocks at `2019-03-29T19:44:44.571815638Z` and is now defunct.


#### [`xrn-test-2`](./archive/xrn-test-2)

Deployed at `2018-12-19T20:40:06.463846Z`.

#### [`xrn-1`](./archive/xrn-1)

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
$ git checkout v1.0.0
$ make install
```

Step-3: Create your key (keyname can be anything):

```sh
$ regen keys add <your_key_name>
```
Use the address generated from above command to fill your KYC
