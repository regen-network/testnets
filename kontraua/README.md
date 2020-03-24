# Regen Network Testnet 3000: COSMWASM Kontraŭa Testnet

**Focus**: Adversarial testnet, network load testing, and interchain smart contracting with Regen Ledger running CosmWASM. This testnet may also morph into a Game of Zones testnet as we are sensitive to the larger community opportunity.

* Testnet schedule: 13th March - 17th Apr 2020 (5 weeks)
* Total points to be allocated: 1100+
* Gentx submissions: 9th March 1500UTC (start date) - 12th March 0900 UTC (end date)
* Genesis release time: 12th March 1600UTC (23 hours before genesis time)
* Network start time: 13th March 1500UTC
* Network Explorer: [Kontraua Explorer](https://explorer.regen.vitwit.com/)

## How to become a validator

Please refer to the Cosmos Hub documentation on validators for a general overview of running a validator. We are using the exact same validator model and software, but with slightly different parameters and other functionality specific to Regen Network.

* [Run a Validator](https://cosmos.network/docs/cosmos-hub/validators/validator-setup.html)
* [Validators Overview](https://cosmos.network/docs/cosmos-hub/validators/overview.html)
* [Validator Security](https://cosmos.network/docs/cosmos-hub/validators/security.html)
* [Validator FAQ](https://cosmos.network/docs/cosmos-hub/validators/validator-faq.html)

### Prerequisites

```sh
sudo apt-get install build-essential
```

***Go 1.13+ is required***

You can use following commands to install go-1.13.3
```sh
$ wget https://raw.githubusercontent.com/jim380/node_tooling/master/Cosmos/CLI/go_install.sh
$ chmod +x go_install.sh
$ ./go_install.sh -v 1.13.3
$ go version # this should output `go version go1.13.3 ...`
```

### Install Cosmwasm
```
$ mkdir -p $GOPATH/src/github.com/regen
$ cd $GOPATH/src/github.com/regen
$ git clone https://github.com/regen-network/wasmd && cd wasmd
$ git checkout v0.7.1
$ make install
```

To verify if the installation was successful, execute the following command:
```
$ xrnd version --long
```
It will display the version of xrnd currently installed:
```
name: wasm
server_name: xrnd
client_name: xrncli
version: 0.7.1
commit: c91f81c25042bfa3ee5890761f818b59914e344b
build_tags: netgo,ledger
go: go version go1.13.3 linux/amd64
```

### Setting Up a Validator Node
```
$ xrnd init --chain-id=kontraua <your_moniker>
$ xrncli keys add <your_wallet_name>
```
**Make sure you back up the mnemonics !!!**

### Become a Genesis validator

If you are looking to join the testnet after the genesis, please check [Start your validator](#start-your-validator)

*This section applies ONLY if you are wishing to validate from the genesis block. This process will close at 0900UTC on 12th March 2020.

#### Generate Genesis Transaction 
```
$ curl -s https://raw.githubusercontent.com/regen-network/testnets/master/kontraua/genesis.json > ~/.xrnd/config/genesis.json
$ xrnd add-genesis-account <your_wallet_name> 10000000utree   # other values will be removed.
$ xrnd gentx --name <your_wallet_name> --amount 9000000utree
```
If all goes well, you will see the following message:
```
Genesis transaction written to "/home/user/.xrnd/config/gentx/gentx-f8038a89034kl987ebd493b85a125624d5f4770.json"
```
#### Submit Gentx
Submit your gentx in a PR [here](https://github.com/regen-network/testnets)

- Fork the testnets repo to your github account 

- Clone your repo using

```sh
git clone https://github.com/<your-github-username>/testnets
```

- Copy the generated gentx json file to <repo_path>/kontraua/gentxs/<your_gentx_file.json>

- Commit and push to your repo
- Create a PR into https://github.com/regen-network/testnets


## Start your validator

This section is applicable only after the genesis is released. Genesis release time is: 12th March, 1600UTC.

### Configure your validator

If you are a Genesis Validator, skip to [Genesis & Seeds](#genesis--seeds)

If you are not part of the genesis validators, please request some free tokens here: [Kontraua Faucet](https://regen.vitwit.com/faucet)

#### Create validator
```sh
xrncli tx staking create-validator \
  --amount=9000000utree \
  --pubkey=$(xrnd tendermint show-validator) \
  --moniker="<your_moniker>" \
  --chain-id=kontraua \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --from=<your_wallet_name>
```

#### Genesis & Seeds
Fetch `genesis.json` into `xrnd`'s `config` directory.
```
$ curl https://raw.githubusercontent.com/regen-network/testnets/master/kontraua/genesis.json > $HOME/.xrnd/config/genesis.json
```

Add seed nodes in `config.toml`.
```
$ nano $HOME/.xrnd/config/config.toml
```
Find the following section and add the seed nodes.
```
# Comma separated list of seed nodes to connect to
seeds = "15ee12ae5fe8256ee94d1065e0000893e52532d9@regen-seed-eu.chorus.one:36656,ca130fd7ca16a957850a96ee9bdb74a351c4929f@regen-seed-us.chorus.one:36656"
```
```
# Comma separated list of persistent peers to connect to
persistent_peers = "4efed1cf14c69aee0695837fc7341036fe7dcc52@173.249.40.87:26656,bf2d9c642d82990ab7e6b5dbc24695e56edfb43e@3.134.55.15:26656"
```

```
# Comma separated list of persistent peers to connect to
persistent_peers = "4efed1cf14c69aee0695837fc7341036fe7dcc52@173.249.40.87:26656,bf2d9c642d82990ab7e6b5dbc24695e56edfb43e@3.134.55.15:26656"
```

## Start Your Node

### **Method 1** - With `systemd`

#### Make `xrnd` a System Service

```
$ sudo nano /lib/systemd/system/xrnd.service
```
Paste in the following:
```
[Unit]
Description=Regen Xrnd
After=network-online.target

[Service]
User=<your_user>
ExecStart=/home/<your_user>/go_workspace/bin/xrnd start --pruning nothing
StandardOutput=file:/var/log/xrnd/xrnd.log
StandardError=file:/var/log/xrnd/xrnd_error.log
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

##### Note: Please make sure to add the ```--pruning ``` flag after the start command

**This tutorial assumes `$HOME/go_workspace` to be your Go workspace. Your actual workspace directory may vary.**

```
$ sudo systemctl enable xrnd
$ sudo systemctl start xrnd
```
Check node status
```
$ xrncli status
```
Check logs
```
$ sudo journalctl -u xrnd -f
```

### **Method 2** - Without `systemd`
```
$ xrnd start --pruning nothing
```
Check node status
```
$ xrncli status
```
##### Note: Please make sure to add the ```--pruning ``` flag after the start command
