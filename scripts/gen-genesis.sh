#!/bin/bash

CHAIN_ID=aplikigo-1
CONFIG=~/.regen/config
FAUCET_ACCOUNTS=("regen:1wuufq6vkl4qmmgzs06mtgatklpxr5zr4qqnk4k" "regen:1anc2w8g3ll9ypr4cdtl5j244eef2nsz5anre7k")

rm -rf ~/.regen

regen init dummy --chain-id $CHAIN_ID

rm -rf $CONFIG/gentx && mkdir $CONFIG/gentx
rm -rf $CONFIG/genesis.json

cp $CHAIN_ID/genesis.json $CONFIG/genesis.json
sed -i "s/\"stake\"/\"utree\"/g" ~/.regen/config/genesis.json

for i in $CHAIN_ID/gentxs/*.json; do
  echo $i
  regen add-genesis-account $(jq -r '.body.messages[0].delegator_address' $i) 100000000000utree
  cp $i $CONFIG/gentx/
done

for addr in "${FAUCET_ACCOUNTS[@]}"; do
    echo "Adding faucet addr: $addr"
    regen add-genesis-account $addr 10000000000000utree
done

regen collect-gentxs

regen validate-genesis

cp $CONFIG/genesis.json $CHAIN_ID
