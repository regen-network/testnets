# Regen Network Testnet 3000: COSMWASM Kontraŭa Testnet

**Focus**: Adversarial testnet and network load testing with Regen Ledger running CosmWASM. This testnet may also morph into a Game of Zones testnet, as we are sensitive to the larger community opportunity.

* Testnet schedule: 13th March - 17th Apr, 2020 (5 weeks)
* Total points to be allocated: 1100+
* Gentx submissions: 9th March 0900UTC (start date) - 12th March 0900 UTC (end date)
* Genesis release time: 12th March, 1600UTC (23 hours before genesis time)
* Network start time: 13th March, 1500UTC

## Become a Genesis validator


### Install Wasmd
```
$ mkdir -p $GOPATH/src/github.com/regen
$ cd $GOPATH/src/github.com/regen
$ git clone https://github.com/regen-network/wasmd && cd wasmd
$ git checkout v0.7.0
$ make install
```
 To verify if installation was successful:
```
$ wamsd version --long
$ wasmcli version --long
```
### Setting Up a New Node
```
$ wasmd init --chain-id=kontraŭa <your_moniker>
$ wamscli keys add <your_wallet_name>

```
**Make sure you back up the mnemonics !!!**

*This section applies ONLY if you are wishing to validate from the genesis block. This process will close at 0900UTC on 12th March 2020.

#### Generate Genesis Transaction 
```
$ curl -s https://raw.githubusercontent.com/regen-network/testnets/master/kontraŭa/genesis.json > ~/.wamsd/config/genesis.json
$ wamsd add-genesis-account $(wamscli keys show <your_wallet_name> -a) 10000000utree   # other values will be removed.
$ wamsd gentx --name <your_wallet_name> --amount 9000000utree
```
If all goes well, you will see the following message:
```
Genesis transaction written to "/home/user/.wamsd/config/gentx/gentx-f8038a89034kl987ebd493b85a125624d5f4770.json"
```
#### Submit Gentx (optional)
Submit your gentx in a PR [here](https://github.com/regen-network/testnets)

- Clone the repo using

```sh
git clone https://github.com/regen-network/testnets
```

- Copy the generated gentx json file to <repo_path>/algradigon-1/gentxs/<your_gentx_file.json>

- Commit and push to your repo
- Create a PR into https://github.com/regen-network/testnets


