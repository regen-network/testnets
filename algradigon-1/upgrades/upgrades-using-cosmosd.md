# Instructions to manage Regen-network upgrades using [cosmosd](https://github.com/regen-network/cosmosd)


## Setup the required env

#### init env
```
export DAEMON_HOME=$HOME/.cosmosd
export DAEMON_NAME=xrnd
```

#### clean (optional)
```rm -rf $DAEMON_HOME # clean the DAEMON_HOME```

#### create required directories
```
mkdir -p $DAEMON_HOME/upgrade_manager
cd $DAEMON_HOME/upgrade_manager
mkdir -p genesis
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

## Lets configure `cosmosd` to use `xrnd` `v0.5.0`

```
cd $DAEMON_HOME/upgrade_manager

wget -c https://github.com/regen-network/regen-ledger/releases/download/v0.5.0/regen-ledger-v0.5.0-linux-amd64.tar.xz -O - | tar -xz -C $DAEMON_HOME/upgrade_manager/genesis/ ./bin

chmod +x genesis/bin/xrnd
./genesis/bin/xrnd version # this should print 0.5.0
```


### Make sure to set the current link, if not starting from genesis

`ln -s $DAEMON_HOME/upgrade_manager/genesis current`

# Download the upgrade binaries, so `cosmosd` can handle the required upgrades for you


## Download patagonia binary

```
cd $DAEMON_HOME/upgrade_manager
mkdir -p upgrades/patagonia/bin
cd upgrades/patagonia
curl -L -o bin/xrnd https://github.com/regen-network/regen-ledger/releases/download/v0.5.1/xrnd-v0.5.1
chmod +x bin/xrnd
./bin/xrnd version # this should print 0.5.1
cd ../..
```

## Setup Cosmosd system service (Ubuntu)

### Prerequistes

Increase resource limits for [Tendermint](https://tendermint.com):
```
$ ulimit -n 4096
```

### Make a system service (optional)

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

# *** IMPORTANT ***
### STOP YOUR `xrnd` if it is runnning already.
#### make sure your validator is not running anymore.


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

Now your setup for automatic upgrade is done. `cosmosd` should do the work needed for `upgrade`

## Important Note

Hopefully the setup works as intented, but if it fails, you need to download the binary and run it manually. So, make yourself available for the upgrade and keep an eye on the logs. `Cosmosd` is still under testing for different scenarios. Please help us improving it by sending your valuable feedback. As always, please join the Regen Network DVD telegram group for any development or validator related questions.

