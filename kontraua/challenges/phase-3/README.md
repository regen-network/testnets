# Himalaya Upgrade Instructions

Upgrading your node is simple, you just need to replace old binary (v0.7.1) with new binary (v0.7.2) and start your node.

### Get the latest release (v0.7.2)
```
$ mkdir -p $GOPATH/src/github.com/regen
$ cd $GOPATH/src/github.com/regen
$ git clone https://github.com/regen-network/wasmd && cd wasmd
$ git checkout v0.7.2
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
version: 0.7.2
commit: c5ffddfc6d960e8140425f4d7887749ac50aab5e
build_tags: netgo,ledger
go: go version go1.13.3 linux/amd64
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


NOTE: You need to start your new binary after reaching the upgrade time i.e., 2020-03-26 12:00:00UTC (Thursday)

### Tweet bonus

Those who tweeted about network upgrade proposal, please raise a PR.
```
$ cd testnets/kontraua/challenges/phase-3
$ cp sample.json <your_moniker.json>
```

Update the required details and raise a PR with title: `Phase-3_<your_moniker>`
