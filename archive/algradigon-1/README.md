# Regen Ledger Testnets

Testnets for [Regen Ledger](https://github.com/regen-network/regen-ledger)



## Join `algradigon-1` Incentivised Testnet

`algradigon-1` is the second phase of the Regen Network Incentivised testnet programme. It is due to have it's genesis block set in stone on Thursday January 23rd at 1700 UTC (1200EST and 0900PST; 24/01 at 0200KST). Future upgrades will be staggered so all validators get 'good' upgrade windows and 'bad' upgrade windows!

We will have a working fork of [Lunie](https://github.com/luniehq/lunie) at https://lunie.regen.network
and a fork of [Big Dipper](https://github.com/forbole/big_dipper) at at https://bigdipper.regen.network/

The genesis files are in [./latest](latest). This is an incentivised testnet, and further details can be found on the Regen blog: [algradigon regens testnet](https://medium.com/regen-network/algradigon-regens-testnet-2000-4ea377c4a590)

Gentx submissions must be included in a PR to this repo *by 22nd Jan at 1200UTC*. The genesis file will be released by 1800UTC on 22nd Jan - 23 hours before genesis.

For those wanting to develop against the Regen test network APIs, please use the following highly available service provided by [Chorus One](https://chorus.one):
* **RPC**: https://regen.chorus.one:26657
* **LCD**: https://regen-lcd.chorus.one:1317

## Creating a Genesis Validator

*This section applies ONLY if you are wishing to validate from the genesis block. This process will close at 1200UTC on 22nd January 2020*

#### Generate Genesis Transaction (optional)
```
$ curl -s https://raw.githubusercontent.com/regen-network/testnets/master/algradigon-1/genesis.json > ~/.xrnd/config/genesis.json
$ xrnd add-genesis-account $(xrncli keys show <your_wallet_name> -a) 10000000utree   # other values will be removed.
$ xrnd gentx --name <your_wallet_name> --amount 9000000utree
```
If all goes well, you will see the following message:
```
Genesis transaction written to "/home/user/.xrnd/config/gentx/gentx-f8038a89034kl987ebd493b85a125624d5f4770.json"
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

## How to Run a Testnet Validator

Please refer to the Cosmos Hub documentation on validators for a general overview of running a validator. We are using the exact same validator model and software, but with slightly different parameters and other functionality specific to Regen Network.

* [Run a Validator](https://cosmos.network/docs/cosmos-hub/validators/validator-setup.html)
* [Validators Overview](https://cosmos.network/docs/cosmos-hub/validators/overview.html)
* [Validator Security](https://cosmos.network/docs/cosmos-hub/validators/security.html)
* [Validator FAQ](https://cosmos.network/docs/cosmos-hub/validators/validator-faq.html)


### Prerequisites
```
$ sudo apt-get install gcc g++
```
### Install GO
```
$ wget https://raw.githubusercontent.com/jim380/node_tooling/master/Cosmos/CLI/go_install.sh
$ chmod +x go_install.sh
$ ./go_install.sh -v 1.12.5
```
**Go 1.12+ is required.**
### Install XRN
```
$ mkdir -p $GOPATH/src/github.com/regen
$ cd $GOPATH/src/github.com/regen
$ git clone -b v0.5.0 https://github.com/regen-network/regen-ledger
$ cd regen-ledger
$ make install
```
Find the latest release tags [here](https://github.com/regen-network/regen-ledger/releases). To verify if installation was successful:
```
$ xrnd version --long
$ xrncli version --long
```
### Setting Up a New Node
```
$ xrnd init --chain-id=algradigon-1 <your_moniker>
$ xrncli keys add <your_wallet_name>

##
```
**Make sure you back up the mnemonics !!!**

### Creating a Validator
*If you are joining at genesis scroll down to the section on Creating a Genesis Validator!*

Please follow the documentation provided on [creating a validator for Cosmos hub](https://github.com/cosmos/gaia/blob/master/docs/validators/validator-setup.md#create-your-validator), replacing `gaiad` and `gaiacli` with `xrnd` and `xrncli` respectively. Also our testnet staking token denomination is `tree` and Regen addresses begin with `xrn:` instead of `cosmos`.

### Genesis & Seeds
Fetch `genesis.json` into `xrnd`'s `config` directory.
```
$ curl https://raw.githubusercontent.com/regen-network/testnets/master/algradigon-1/genesis.json > $HOME/.xrnd/config/genesis.json
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
### Make `xrnd` a System Service (optional)
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
ExecStart=/home/<your_user>/go_workspace/bin/xrnd start
StandardOutput=file:/var/log/xrnd/xrnd.log
StandardError=file:/var/log/xrnd/xrnd_error.log
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
**This tutorial assumes `$HOME/go_workspace` to be your Go workspace. Your actual workspace directory may vary.**
#### Start Node
**Method 1** - With `systemd`
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
**Method 2** - Without `systemd`
```
$ xrnd start
```
Check node status
```
$ xrncli status
```

# Historic Testnets (not in use)

The testnets listed below are no longer active but are retained here for posterity. Do not waste your time trying to join them :)

## `congo-1` 

`congo-1` was the first community-driven testnet between 7th November 2019 and 17th January 2020, at which point it was abandoned in order to prepare for the next incentivised testnet.

## `regen-test-1001`

`regen-test-1001` ran between August and October 2019. It is now defunct, having discovered a cosmos-sdk bug in governance-driven parameter updates logic.

## `regen-test-1000` 

`regen-test-1000` hit some weird consensus error on app state at block 2.

### `xrn-test-3`

Testnet `xrn-test-3` started producing blocks at `2019-03-29T19:44:44.571815638Z` and is now defunct.


`xrncli` can be configured to connect to the testnet as follows:

```sh
xrncli init --chain-id xrn-test-2 --node tcp://xrn-us-east-1.regen.network:26657
```

### `xrn-test-2`

Deployed at `2018-12-19T20:40:06.463846Z`.

### `xrn-1`

The initial Regen Ledger testnet `xrn-1` was deployed on 2018-12-19.

```
