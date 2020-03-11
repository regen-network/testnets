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
$ xrnd init --chain-id=kontraua <your_moniker>
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
