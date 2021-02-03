# Regen Network Testnet: Aplikigo-1 Testnet

Genesis: [genesis.json](https://raw.githubusercontent.com/regen-network/testnets/master/aplikigo-1/genesis.json)

Git tag: [v0.6.0](https://github.com/regen-network/regen-ledger/releases/tag/v0.6.0)

Block explorer: [aplikigo.regen.aneka.io](https://aplikigo.regen.aneka.io)

Faucet: [regen.vitwit.com/faucet](https://regen.vitwit.com/faucet)

Seeds: TBA

### Minimum Hardware Requirements

 - 8GB RAM
 - 4vCPUs
 - 200GB Disk space

## How to become a validator

Please refer to the Cosmos Hub documentation on validators for a general overview of running a validator. We are using the exact same validator model and software, but with slightly different parameters and other functionality specific to Regen Network.

* [Run a Validator](https://cosmos.network/docs/cosmos-hub/validators/validator-setup.html)
* [Validators Overview](https://cosmos.network/docs/cosmos-hub/validators/overview.html)
* [Validator Security](https://cosmos.network/docs/cosmos-hub/validators/security.html)
* [Validator FAQ](https://cosmos.network/docs/cosmos-hub/validators/validator-faq.html)

### Prerequisites

```
sudo apt-get install build-essential jq
```

***Go 1.15+ is required***

You can use following commands to install go-1.15.5
```sh
$ wget https://raw.githubusercontent.com/jim380/node_tooling/master/cosmos/scripts/go_install.sh
$ chmod +x go_install.sh
$ ./go_install.sh -v 1.15.5
$ go version # this should output `go version go1.15.5 ...`
```

### Install Regen
```
$ mkdir -p $GOPATH/src/github.com/regen
$ cd $GOPATH/src/github.com/regen
$ git clone https://github.com/regen-network/regen-ledger.git && cd regen-ledger
$ git checkout v0.6.0
$ make install
```

To verify if the installation was successful, execute the following command:
```
$ regen version --long
```
It will display the version of regen currently installed:
```
TBA
```

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
$ regen add-genesis-account <your_wallet_name> 10000000utree   # other values will be removed.
$ regen gentx <your_wallet_name> --amount 9000000utree --chain-id aplikigo-1
```
If all goes well, you will see the following message:
```
Genesis transaction written to "/home/user/.regen/config/gentx/gentx-f8038a89034kl987ebd493b85a125624d5f4770.json"
```
#### Submit Gentx
Submit your gentx in a PR [here](https://github.com/regen-network/testnets)

- Fork the testnets repo to your github account 

- Clone your repo using

```
git clone https://github.com/<your-github-username>/testnets
```

- Copy the generated gentx json file to <repo_path>/aplikigo-1/gentxs/<your_gentx_file.json>

- Commit and push to your repo
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
seeds = "TBA"
```
```
# Comma separated list of persistent peers to connect to
persistent_peers = "TBA"
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
ExecStart=${GOBIN}/regen start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >regen.service
```


```
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

### **Option 2** - Without `systemd`
```
$ regen start 
```
Check node status
```
$ regen status
```


