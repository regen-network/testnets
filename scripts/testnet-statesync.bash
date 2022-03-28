#!/bin/bash
# This is the Regen Network State sync file for Redwood Testnet, which is based on the Gaia State sync file, which is based on scripts written by Bitcanna and Microtick.
# http://redwood.regen.network:26657
# http://redwood-sentry.vitwit.com:26657



set -uxe

# set environment variables
export GOPATH=~/go
export PATH=$PATH:~/go/bin
export HOME_DIR=~/.regen

MONIKER=$1
if [ -z $MONIKER ]
then
    MONIKER=test-sync
fi

# MAKE HOME FOLDER AND GET GENESIS
regen init $MONIKER --home $HOME_DIR
wget https://raw.githubusercontent.com/regen-network/testnets/master/redwood-testnet/genesis.json -O $HOME_DIR/config/genesis.json 

INTERVAL=10

# GET TRUST HASH AND TRUST HEIGHT

LATEST_HEIGHT=$(curl -s http://redwood.regen.network:26657/block | jq -r .result.block.header.height);
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL))
TRUST_HASH=$(curl -s "http://redwood.regen.network:26657/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)


# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"


# expor state sync vars
export REGEN_STATESYNC_ENABLE=true
export REGEN_P2P_MAX_NUM_OUTBOUND_PEERS=200
export REGEN_P2P_MAX_NUM_INBOUND_PEERS=200
export REGEN_STATESYNC_RPC_SERVERS="http://redwood.regen.network:26657,http://redwood-sentry.vitwit.com:26657"
export REGEN_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export REGEN_STATESYNC_TRUST_HASH=$TRUST_HASH
export REGEN_P2P_PERSISTENT_PEERS=$(curl -s https://raw.githubusercontent.com/regen-network/testnets/master/redwood-testnet/peer-nodes.txt | paste -sd,)

sed -i '/persistent_peers =/c\persistent_peers = "'"$REGEN_P2P_PERSISTENT_PEERS"'"' $HOME_DIR/config/config.toml
sed -i '/max_num_outbound_peers =/c\max_num_outbound_peers = '$REGEN_P2P_MAX_NUM_OUTBOUND_PEERS'' $HOME_DIR/config/config.toml
sed -i '/max_num_inbound_peers =/c\max_num_inbound_peers = '$REGEN_P2P_MAX_NUM_INBOUND_PEERS'' $HOME_DIR/config/config.toml
sed -i '/enable =/c\enable = true' $HOME_DIR/config/config.toml
sed -i '/rpc_servers =/c\rpc_servers = "'"$REGEN_STATESYNC_RPC_SERVERS"'"' $HOME_DIR/config/config.toml
sed -i '/trust_height =/c\trust_height = '$REGEN_STATESYNC_TRUST_HEIGHT'' $HOME_DIR/config/config.toml
sed -i '/trust_hash =/c\trust_hash = "'"$REGEN_STATESYNC_TRUST_HASH"'"' $HOME_DIR/config/config.toml

regen start --x-crisis-skip-assert-invariants
