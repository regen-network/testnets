#!/bin/bash

CHAIN_ID=aplikigo-1
CONFIG=~/.regen/config
FAUCET_ACCOUNTS=("acc1" "acc2" "acc3")

rm -rf ~/.regen

regen init dummy --chain-id $CHAIN_ID

rm -rf $CONFIG/gentx && mkdir $CONFIG/gentx

sed -i "s/\"stake\"/\"utree\"/g" ~/.regen/config/genesis.json

for i in $NETWORK/gentxs/*.json; do
  echo $i
  regen add-genesis-account $(jq -r '.value.msg[0].value.delegator_address' $i) 100000000000utree
  cp $i $CONFIG/gentx/
done

for addr in "${FAUCET_ACCOUNTS[@]}"; do
    echo "Adding faucet addr: $addr"
    regen add-genesis-account $addr 10000000000000utree
done

regen collect-gentxs

regen validate-genesis

cp $CONFIG/genesis.json $NETWORK
