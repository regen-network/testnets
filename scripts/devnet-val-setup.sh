echo "Install dependencies"
sudo apt update
sudo apt install build-essential jq -y

wget https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz
tar -xvf go1.15.2.linux-amd64.tar.gz
sudo mv go /usr/local

echo "" >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc

source ~/.bashrc

go version


-- Install Regen-ledger and setup the node -----

YOUR_KEY_NAME=mykey
YOUR_NAME=myname
DAEMON=regen
DENOM=utree
CHAIN_ID=regen-devnet-1
PERSISTENT_PEERS="4763e95d731a3402c1f06fbac1535fc38a2e439f@18.220.101.192:26656"

echo "install regen-ledger:master"
git clone https://github.com/regen-network/regen-ledger $GOPATH/src/github.com/regen-network/regen-ledger
cd $GOPATH/src/github.com/regen-network/regen-ledger
make install

echo "Creating keys"
$DAEMON keys add $YOUR_KEY_NAME

echo "Setting up your validator"
$DAEMON init --chain-id $CHAIN_ID $YOUR_NAME
curl http://18.220.101.192:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
#sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.$DAEMON/config/genesis.json

echo "----------Setting config for seed node---------"
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.$DAEMON/config/config.toml
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.$DAEMON/config/config.toml

DAEMON_PATH=$(which $DAEMON)


echo "---------Creating system file---------"

echo "[Unit]
Description=${DAEMON} daemon
After=network-online.target
[Service]
User=${USER}
ExecStart=${DAEMON_PATH} start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >$DAEMON.service

sudo mv $DAEMON.service /lib/systemd/system/$DAEMON.service
sudo -S systemctl daemon-reload
sudo -S systemctl start $DAEMON


echo "Your node setup is done. You would need some tokens to start your validator. You can get some tokens from the faucet: https://faucet.devnet.regen.vitwit.com"

echo

echo "Your account address is :"
$DAEMON keys show $YOUR_KEY_NAME -a

echo
echo

echo "After receiving tokens, you can create your validator by running"

echo '$DAEMON tx staking create-validator --amount 90000000000$DENOM --commission-max-change-rate "0.1" --commission-max-rate "0.20" --commission-rate "0.1" --details "Some details about yourvalidator" --from $YOUR_KEY_NAME   --pubkey=$($DAEMON tendermint show-validator) --moniker $YOUR_NAME --min-self-delegation "1" --chain-id $CHAIN_ID'