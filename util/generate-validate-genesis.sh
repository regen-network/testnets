#!/bin/bash

CHAIN_DIR=algradigon-1 #chain-id directory
GENTXS_DIR=$CHAIN_DIR/gentxs
CONFIG=~/.xrnd/config #xrnd config path

echo "This script assumes that all the gentx submissions are placed in : $GENTXS_DIR, and the directory contains only valid gentxs"

sleep 5 # sleep for 5 seconds

echo 
# format the gentx files
for i in $GENTXS_DIR/*.json; do
  echo $i
  cat $i | jq '.' > /tmp/gentx.json
  cp /tmp/gentx.json $i
  rm /tmp/gentx.json
done

sleep 5 # sleep for 5 seconds

# collate gentxs

## remove old gentx (if any)
rm -rf $CONFIG/gentx && mkdir $CONFIG/gentx

## remove old genesis (if any)
rm -rf $CHAIN_DIR/genesis.json

echo "add all gentx's accounts to genesis \n"
for i in $GENTXS_DIR/*.json; do
  echo $i
  xrnd add-genesis-account $(jq -r '.value.msg[0].value.delegator_address' $i) $(jq -r '.value.msg[0].value.value.amount' $i)$(jq -r '.value.msg[0].value.value.denom' $i)
  cp $i $CONFIG/gentx/
done

sleep 5 # sleep for 5 seconds

# collect gentxs
echo "collecting gentxs \n"
xrnd collect-gentxs

sleep 5 # sleep for 5 seconds

echo "validate genesis file \n"
xrnd validate-genesis

sleep 5 # sleep for 5 seconds

echo "copy the genesis to $CHAIN_DIR \n"
cp $CONFIG/genesis.json $CHAIN_DIR

sleep 5 # sleep for 5 seconds

echo "Done! Genesis file generated and validated successfully!!!"