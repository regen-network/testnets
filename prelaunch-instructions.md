# Start a node for Prelaunch Chain (`regen-devnet-5)

```sh
sudo apt update
sudo apt install build-essential jq

wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz
tar -xvf go1.15.6.linux-amd64.tar.gz
sudo mv go /usr/local

echo "" >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc

source ~/.bashrc

go version

go get github.com/regen-network/regen-ledger
cd ~/go/src/github.com/regen-network/regen-ledger
git checkout v1.0.0-rc1
make install
regen version

regen init --chain-id regen-devnet-5 
curl -s https://raw.githubusercontent.com/anilCSE/mainnet/main/regen-prelaunch-1/genesis.json > ~/.regen/config/genesis.json

export DAEMON=regen
PERSISTENT_PEERS="1db306fd845db455788fc3211631573b27f5ee0d@138.197.199.25:26656,295849ec0051216fd1dc33a2ee282a8583c52475@138.68.243.146:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.$DAEMON/config/config.toml

echo "---------Creating system file---------"

echo "[Unit]
Description=${DAEMON} daemon
After=network.target
[Service]
Type=simple
User=$USER
ExecStart=$(which $DAEMON) start
Restart=on-failure
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
">$DAEMON.service

sudo mv $DAEMON.service /lib/systemd/system/$DAEMON.service

echo "-------Start $DAEMON service-------"

sudo -S systemctl daemon-reload
sudo -S systemctl start $DAEMON
```

# If you are planning a validator, you can just execute the following command after your node syncs fully.
Check latest height of the network at http://prelaunch.regen.aneka.io

```sh
regen tx staking create-validator --amount 9000000000$DENOM \
  --commission-max-change-rate "0.1" --commission-max-rate "0.20" --commission-rate "0.1" \
  --details "Some details about yourvalidator" \
  --from <YOUR_KEY_NAME> \
  --pubkey=$(regen tendermint show-validator) \
  --moniker <your_moniker>
  --min-self-delegation "1"
  --chain-id regen-devnet-5
  --node http://localhost:26657
```

Note: you can execute this command locally by replacing `--pubkey` with the validator pubkey and appending memo flag with `node-id@ip:26657`

To get validator pubkey: `regen tendermint show-validator`

To get node id: `regen tendermint show-node-id`
