#!/bin/sh
REGEN_HOME="/tmp/regen$(date +%s)"
RANDOM_KEY="randomregenvalidatorkey"
CHAIN_ID=aplikigo-1

GENTX_FILE=$(ls $CHAIN_ID/gentxs | head -1)
LEN_GENTX=$(echo ${#GENTX_FILE})


GENTX_STARTDATE=$(date -d '05-02-2021 15:00:00' '+%d/%m/%Y %H:%M:%S')
GENTX_DEADLINE=$(date -d '07-02-2021 15:00:00' '+%d/%m/%Y %H:%M:%S')
now=$(date +"%d/%m/%Y %H:%M:%S")

# if [ $now -le $GENTX_STARTDATE ]; then
#     echo 'Gentx submission is not open yet. Please close the PR and raise a new PR after 05-Feb-2021 15:00:00'
#     exit 0
# fi

# if [ $now -ge $GENTX_DEADLINE ]; then
#     echo 'Gentx submission is closed'
#     exit 0
# fi

if [ $LEN_GENTX -eq 0 ]; then
    echo "No new gentx file found."
else
    set -e

    ./scripts/check-gentx-amount.sh "./$CHAIN_ID/gentxs/$GENTX_FILE" || exit 1

    echo "...........Init Regen.............."
    curl -L https://github.com/regen-network/regen-ledger/releases/download/v0.6.0/regen_0.6.0_linux_amd64.zip -o regen_linux.zip && unzip regen_linux.zip
    rm regen_linux.zip
    #cd regen_0.6.0_linux_amd64

    ./regen keys add $RANDOM_KEY --keyring-backend test --home $REGEN_HOME

    ./regen init --chain-id $CHAIN_ID validator --home $REGEN_HOME

    echo "..........Fetching genesis......."
    rm -rf $REGEN_HOME/config/genesis.json
    curl -s https://raw.githubusercontent.com/regen-network/testnets/master/$CHAIN_ID/genesis.json > $REGEN_HOME/config/genesis.json

    sed -i '/genesis_time/c\   \"genesis_time\" : \"2021-01-01T00:00:00Z\",' $REGEN_HOME/config/genesis.json

    GENACC=$(cat ../$CHAIN_ID/gentxs/$GENTX_FILE | sed -n 's|.*"delegator_address":"\([^"]*\)".*|\1|p')

    echo $GENACC

    ./regen add-genesis-account $RANDOM_KEY 1000000000000utree --home $REGEN_HOME \
        --keyring-backend test
    ./regen add-genesis-account $GENACC 100000000000utree --home $REGEN_HOME

    ./regen gentx --name $RANDOM_KEY --amount 900000000000utree --home $REGEN_HOME \
        --keyring-backend test
    cp ../$CHAIN_ID/gentxs/$GENTX_FILE $REGEN_HOME/config/gentx/

    echo "..........Collecting gentxs......."
    ./regen collect-gentxs --home $REGEN_HOME
    sed -i '/persistent_peers =/c\persistent_peers = ""' $REGEN_HOME/config/config.toml

    ./regen validate-genesis --home $REGEN_HOME

    echo "..........Starting node......."
    ./regen start --home $REGEN_HOME &

    sleep 5s

    echo "...checking network status.."

    ./regen status --chain-id $CHAIN_ID --node http://localhost:26657

    echo "...Cleaning the stuff..."
    killall regen >/dev/null 2>&1
    rm -rf $REGEN_HOME >/dev/null 2>&1
fi
