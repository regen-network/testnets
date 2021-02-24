# Regen Network Testnet: Aplikigo-1 Testnet

## Quick Links
Genesis: [genesis.json](https://raw.githubusercontent.com/regen-network/testnets/master/aplikigo-1/genesis.json)

Git tag: [v0.6.0](https://github.com/regen-network/regen-ledger/releases/tag/v0.6.0)

Block explorer: [aplikigo.regen.aneka.io](https://aplikigo.regen.aneka.io)

Faucet: [regen.vitwit.com/faucet](https://regen.vitwit.com/faucet)

Seeds: TBA

## Hardware Requirements
Here are the minimal hardware configs required for running a validator/sentry node
 - 8GB RAM
 - 4vCPUs
 - 200GB Disk space

## Software Requirements

### Install deps
```
sudo apt-get install build-essential jq
```

### Install Regen
You can install Regen by downloading the binary (simple) or compiling from source.

#### Option-1: Build from source (Recommended)

***Go 1.15+ is required***

You can use following commands to install go-1.15.5
```sh
$ wget https://raw.githubusercontent.com/jim380/node_tooling/master/cosmos/scripts/go_install.sh
$ chmod +x go_install.sh
$ ./go_install.sh -v 1.15.5
$ go version # this should output `go version go1.15.5 ...`
```

```sh
$ mkdir -p $GOPATH/src/github.com/regen
$ cd $GOPATH/src/github.com/regen
$ git clone https://github.com/regen-network/regen-ledger.git && cd regen-ledger
$ git checkout v0.6.0
$ make install
```

To verify if the installation was successful, execute the following command:
```sh
$ regen version --long
```

It will display the version of regen currently installed:
```sh
name: regen
server_name: regen
version: v0.6.0
commit: 6f743dce1e62b62c97b01ad042b49b0db31df5ab
build_tags: netgo,ledger
go: go version go1.15.5 linux/amd64
build_deps:
- github.com/99designs/keyring@v1.1.6
- github.com/ChainSafe/go-schnorrkel@v0.0.0-20200405005733-88cbf1b4c40d
...
```

#### Option 2: Download binary

1. Download the binary for your platform: [releases](https://github.com/regen-network/regen-ledger/releases/tag/v0.6.0).
2. Copy it to a location in your PATH, i.e: `/usr/local/bin` or `$HOME/bin`.

i.e:
```sh
$ wget https://github.com/regen-network/regen-ledger/releases/download/v0.6.0/regen_0.6.0_linux_arm64.tar.gz
$ sudo tar -C /usr/local/bin -zxvf regen_0.6.0_linux_arm64.tar.gz
```

## How to become a validator

### Setting Up a Validator Node
```
$ regen init --chain-id=aplikigo-1 <your_moniker>
$ regen keys add <your_wallet_name>
```
**Make sure you back up the mnemonics !!!**

### Become a Genesis validator

If you are looking to join the testnet after the genesis, please check [Start your validator](#start-your-validator)

*This section applies ONLY if you are wishing to validate from the genesis block.

#### Generate Genesis Transaction 
```
$ curl -s https://raw.githubusercontent.com/regen-network/testnets/master/aplikigo-1/genesis.json > ~/.regen/config/genesis.json
$ regen add-genesis-account <your_wallet_name> 100000000000utree   # other values will be removed.
$ regen gentx <your_wallet_name> 90000000000utree --chain-id aplikigo-1
```
If all goes well, you will see the following message:
```
Genesis transaction written to "/home/user/.regen/config/gentx/gentx-f8038a89034kl987ebd493b85a125624d5f4770.json"
```
#### Submit Gentx
Submit your gentx in a PR [here](https://github.com/regen-network/testnets)

- Fork the testnets repo to your github account (You can fork any repo by clicking the fork button in the upper right hand corner of a repo page.)

- Clone your repo using

```
git clone https://github.com/<your-github-username>/testnets
```

- Copy the generated gentx json file to <repo_path>/aplikigo-1/gentxs/<your_gentx_file.json>

- Commit and push to your repo
```
git add <repo_path>/aplikigo-1/gentxs/<your_gentx_file.json>
git commit -m 'add gentx'
git push origin master
```
- Create a PR into https://github.com/regen-network/testnets


## Start your validator

**Note**: This section is applicable only after the genesis is released.

### Configure your validator

If you are a Genesis Validator, skip to [Genesis & Seeds](#genesis--seeds)

If you are not part of the genesis validators, please request some free tokens here: [Aplikigo-1 Faucet](https://regen.vitwit.com/faucet)

#### Create validator
```
regen tx staking create-validator \
  --amount=9000000utree \
  --pubkey=$(regen tendermint show-validator) \
  --moniker="<your_moniker>" \
  --chain-id=aplikigo-1 \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --from=<your_wallet_name>
```

#### Genesis & Seeds
Fetch `genesis.json` into `regen`'s `config` directory.
```
$ curl https://raw.githubusercontent.com/regen-network/testnets/master/aplikigo-1/genesis.json > $HOME/.regen/config/genesis.json
```

Add seed nodes in `config.toml`.
```
$ nano $HOME/.regen/config/config.toml
```
Find the following section and add the seed nodes.
```
# Comma separated list of seed nodes to connect to
seeds = "9082e4d408b0794f884f8c1733d7d11ffb010e38@161.35.51.84:26656"
```
```
# Comma separated list of persistent peers to connect to
persistent_peers = "35d64042e1a5f6466b2b1540fa2e859dfc49666e@public-rpc1.regen.vitwit.com:26656,9082e4d408b0794f884f8c1733d7d11ffb010e38@161.35.51.84:26656,937bbb7a231a870452dd595a442a191d609a26a5@116.203.20.144:26656"
```

#### Set validator gas fees

You can set the minimum gas prices for transactions to be accepted into your node's mempool. This sets a lower bound on gas prices, preventing spam. Stakebird can accept gas in *any* currency. To accept both ATOM and EGG for example, set `minimum-gas-prices` in `app.toml`.

```sh
$ nano $HOME/.regen/config/app.toml
```

```sh
minimum-gas-prices = "0.025utree"
```

## Start Your Node

### **Option 1** - With `systemd`

#### Make `regen` a System Service

```
echo "[Unit]
Description=Regen Node
After=network-online.target
[Service]
User=${USER}
ExecStart=$(which regen) start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >regen.service
```

```
$ sudo mv regen.service /lib/systemd/system/
$ sudo systemctl daemon-reload
$ sudo systemctl enable regen.service
$ sudo systemctl start regen.service
```
Check node status
```
$ regen status
```
Check logs
```
$ sudo journalctl -u regen -f
```

### **Option 2** - Without `systemd` (using screen, etc)
```
$ regen start 
```
Check node status
```
$ regen status
```

### Troubleshooting
If you get an error like: 
```
Error: error during handshake: error on replay: validator set is nil in genesis and still empty after InitChain
```
Please run the following command and restart the regen service
```
regen unsafe-reset-all
sudo systemctl restart regen.service
``` 



## More about validator

Please refer to the Cosmos Hub documentation on validators for a general overview of running a validator. We are using the exact same validator model and software, but with slightly different parameters and other functionality specific to Regen Network.

* [Run a Validator](https://hub.cosmos.network/main/validators/validator-setup.html)
* [Validators Overview](https://hub.cosmos.network/main/validators/overview.html)
* [Validator Security](https://hub.cosmos.network/main/validators/security.html)
* [Validator FAQ](https://hub.cosmos.network/main/validators/validator-faq.html)
