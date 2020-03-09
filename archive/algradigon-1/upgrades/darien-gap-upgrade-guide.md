# Instructions to manage Regen-network's "Darien Gap" upgrade

You have 2 options here. 

- Manual Upgrade
- Using cosmosd (Suggested)

## 1. Instructions to manage Darien Gap upgrade manually

#### Step-1:
Download the Darien Gap release from here 
[regen-ledger-v0.5.3-linux-amd64.tar.xz](https://github.com/regen-network/regen-ledger/releases/tag/v0.5.3)

Linux:
```
wget https://github.com/regen-network/regen-ledger/releases/download/v0.5.3/regen-ledger-v0.5.3-linux-amd64.tar.xz

tar -xzvf regen-ledger-v0.5.3-linux-amd64.tar.xz
```


(OR) build from source
```
go get github.com/regen-network/regen-ledger
cd <GOPATH>/regen-network/regen-ledger

git fetch
git checkout v0.5.3
make build
./build/xrnd version #this should print 0.5.3
```

Note: The following instructions should be executed at the time of upgrade

#### Step-2:
- Stop the current xrnd. If you are using systemctl, you can do following:
```
sudo systemctl stop xrnd
```

### Step-3:
- Update the binary

If you have dowloaded the prebuild binary, use this:
```
mv regen-ledger-v0.5.3-linux-amd64/bin/* <GOBIN>/
```

If you built the binary from source:

Option - 1:

```
mv build/* <GOBIN>/
```
Option - 2:

You can just run following from regen-ledger directory:
```
make install
```

#### Step-3:
- Start the latest xrnd
```
xrnd start
```

If you are using xrnd systemctl, you can do following:

```
sudo systemctl start xrnd
```

All done! You have completed the upgrade proces. Sit back and wait for the first block to sync, just to ensure your upgrade is successful.


## Instructions to manage the *Darien Gap Upgrade* with cosmosd

### Note: If you have configured cosmosd already for previous (Papua) upgrade, skip to [Configure Darien Gap Upgrade](#configure-darien-gap-upgrade)

If you are starting the setup for cosmosd now, follow this:
## Setup the required env

#### init env
```
export DAEMON_HOME=$HOME/.cosmosd #you can also use your own setting here
export DAEMON_NAME=xrnd
```

#### create required directories
```
mkdir -p $DAEMON_HOME/upgrade_manager
cd $DAEMON_HOME/upgrade_manager
mkdir -p genesis (optional)
```

## Get a prebuilt copy of cosmosd or build from source:

#### Use prebuild copy

```
curl -L -o $DAEMON_HOME/cosmosd https://github.com/regen-network/cosmosd/releases/download/0.2.0/cosmosd

chmod +x $DAEMON_HOME/cosmosd
```

#### or build from source using:

```
git clone git@github.com/regen-network/cosmosd.git 
cd cosmosd
go build
mv cosmosd $DAEMON_HOME/
chmod +x $DAEMON_HOME/cosmosd
```

## Lets configure `cosmosd` to use `xrnd` `v0.5.2`

### Download papua binary

```
cd $DAEMON_HOME/upgrade_manager
mkdir -p upgrades/papua/bin
cd upgrades/papua

wget -c https://github.com/regen-network/regen-ledger/releases/download/v0.5.2/regen-ledger-v0.5.2-linux-amd64.tar.xz -O - | tar -xz -C $DAEMON_HOME/upgrade_manager/upgrades/papua/ bin

chmod +x bin/xrnd
./bin/xrnd version # this should print 0.5.2
cd ../..
```

### Make sure to set the current link, if not starting from genesis

`ln -s $DAEMON_HOME/upgrade_manager/upgrades/papua current`

### Check cosmosd configuration
```
cosmosd version #this should print 0.5.2
```

## Setup Cosmosd system service (Ubuntu)

### Prerequistes

Increase resource limits for [Tendermint](https://tendermint.com):
```
$ ulimit -n 4096
```

### Make a system service

Creating a system service makes it easy to start and stop xrnd, and view logs.

```
sudo nano /etc/systemd/system/xrnd.service
```

Paste the following (replace `ubuntu` with your username):

```
[Unit]
Description=RegenNetwork Node
After=network-online.target

[Service]
User=ubuntu 
WorkingDirectory=/home/ubuntu/.cosmosd/
ExecStart=/home/ubuntu/.cosmosd/cosmosd start
Environment=DAEMON_HOME=/home/ubuntu/.cosmosd
Environment=DAEMON_NAME=xrnd
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
****
******
# IMPORTANT
### STOP YOUR `xrnd` if it is runnning already and make sure your validator is not running anymore.
********
****

Now enable the service:
```
sudo systemctl enable xrnd
```

Start node:
```
sudo systemctl start xrnd
```

Check logs:
```
journalctl -u xrnd -f --lines 50
```

If you need to Stop the node (not required atm):
```
sudo systemctl stop xrnd
```

Your cosmosd setup is done and your validator should start syncing new blocks by now.

## Configure Darien Gap Upgrade

### Download Darien Gap binary

```
cd $DAEMON_HOME/upgrade_manager
mkdir -p upgrades/darien-gap
cd upgrades/darien-gap

wget -c https://github.com/regen-network/regen-ledger/releases/download/v0.5.3/regen-ledger-v0.5.3-linux-amd64.tar.xz -O - | tar -xz -C $DAEMON_HOME/upgrade_manager/upgrades/darien-gap/ ./bin

chmod +x bin/xrnd
./bin/xrnd version # this should print 0.5.3
cd ../..
```

Now your setup for automatic upgrade is done. `cosmosd` should do the work needed for `upgrade`

## Important Note

Hopefully the setup works as intented, but if it fails, you need to download the binary and run it manually. So, make yourself available for the upgrade and keep an eye on the logs. `Cosmosd` is still under testing for different scenarios. Please help us improving it by sending your valuable feedback. As always, please join the Regen Network DVD telegram group for any development or validator related questions.
