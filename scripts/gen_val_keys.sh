#!/bin/bash

command_exists () {
    type "$1" &> /dev/null ;
}

if command_exists go ; then
    echo "Golang is already installed"
else
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

  #source ~/.bashrc
  . ~/.bashrc

  go version
fi

echo "-- Clear old regen data and install Regen-ledger and setup the node --"

rm -rf ~/.regen
rm -rf $GOPATH/src/github.com/regen-network/regen-ledger

YOUR_KEY_NAME=$1
DAEMON=regen
DENOM=utree
CHAIN_ID=regen-testnet
PERSISTENT_PEERS="f864b879f59141d0ad3828ee17ea0644bdd10e9b@18.220.101.192:26656"

echo "install regen-ledger:v0.6.0-alpha2"
git clone https://github.com/regen-network/regen-ledger $GOPATH/src/github.com/regen-network/regen-ledger
cd $GOPATH/src/github.com/regen-network/regen-ledger
git checkout v0.6.0-alpha2
make install

echo "Creating key $YOUR_KEY_NAME"
$DAEMON keys add $YOUR_KEY_NAME

echo "Generating your validator keys"
$DAEMON init --chain-id $CHAIN_ID $YOUR_KEY_NAME

echo "Here is your validator operator key"
$DAEMON keys show $YOUR_KEY_NAME --bech val
