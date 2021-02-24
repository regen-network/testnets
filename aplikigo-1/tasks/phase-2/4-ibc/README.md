# Instructions to transfer tokens b/w chains

The following instructions are for transfering the tokens from `aplikigo-1` (source) chain to `regen-devnet-4` (destination) chain and transfer them back to source chain.

## Setting up the realyer

Useful links:
https://docs.cosmos.network/master/ibc/relayer.html

https://github.com/cosmos/relayer#relayer

### Install Relayer:
```sh
export RELAYER=$GOPATH/src/github.com/cosmos/relayer
mkdir -p $(dirname $RELAYER) && git clone https://github.com:cosmos/relayer $RELAYER && cd $RELAYER

git checkout colin/425-refactor-update-msgs

make install
```

### Setup the relayer & make ibc transfer
- **Step-1:** Init relayer config

    ```sh
    rly config init
    ```

- **Step-2:** Write `regen-devnet-4` and `aplikigo-1` chain configs into separate json files.

    Note: Please change gas-prices value based on your rpc-node. Also make sure to update the aplikigo-1 RPC node if you are running this on a different machine.

    Also please update your keyname as wish. `testkey` is being used as an example in this document

    ```sh
    $ echo "{\"key\":\"testkey\",\"chain-id\":\"regen-devnet-4\",\"rpc-addr\":\"http://18.220.101.192:26657\",\"account-prefix\":\"regen:\",\"gas-adjustment\":1.5,\"gas-prices\":\"0.025uregen\",\"trusting-period\":\"336h\"}" > regen-devnet-4.json

    $ echo "{\"key\":\"testkey\",\"chain-id\":\"aplikigo-1\",\"rpc-addr\":\"http://localhost:26657\",\"account-prefix\":\"regen:\",\"gas-adjustment\":1.5,\"gas-prices\":\"0.025utree\",\"trusting-period\":\"336h\"}" > aplikigo-1.json
    ```
- **Step-3:** Add above both chains to relayer
Option-1: Add individual chains
    ```sh
    rly ch a -f regen-devnet-4.json
    rly ch a -f aplikigo-1.json
    ```
Option-2: Add bulk (by copying both json files into one folder)
    ```sh

    rly chains add-dir  <path-to-chains-folder>/
    ```
- **Step-4** Export source and destination chain-ids to variables.

    Note: Please update the path name as you wish
    ```sh
    export SRC=aplikigo-1
    export DST=regen-devnet-4
    export PTH=demo-path
    ```
- **Step-5** Create IBC light clients locally
    ```sh
    rly light init $SRC -f 
    rly l i $DST -f
    ```
- **Step-6** Add/Import keys
Ensure each chain has its appropriate key. Import your keys by using:
    ```sh
    rly keys restore $SRC testkey "{{mnemonic-words}}"
    rly keys restore $DST testkey "{{mnemonic-words}}"
    ```
    show key will return address of chain's default key i.e., testkey
    ```sh
    rly keys show $SRC
    rly keys show $DST
    ```

- **Step-7** Request some funds on `regen-devnet-4` here: http://faucet.devnet.regen.vitwit.com/

- **Step-8** Ensure you have funds on both chains
    ```sh
    rly query bal $SRC
    rly query bal $DST
    ```

    Note: If you don't have funds, you cannot make transactions

- **Step-9** Add path between chains
    ```sh
    $ echo "{\"src\":{\"chain-id\":\"$SRC\",\"port-id\":\"transfer\",\"order\":\"unordered\",\"version\":\"ics20-1\"},\"dst\":{\"chain-id\":\"$DST\",\"port-id\":\"transfer\",\"order\":\"unordered\",\"version\":\"ics20-1\"},\"strategy\":{\"type\":\"naive\"}}" > $PTH.json
    $ rly pth add $SRC $DST $PTH -f $PTH.json
    ```

- **Step-10** Link path (creates client, connections and channels)
    ```sh
    rly tx link $PTH
    ```

- **Step-11** Send some funds back and forth
    ```sh
    rly q bal $SRC
    rly q bal $DST

    # transfer tokens from source chain to dst chain
    rly tx transfer $SRC $DST {{amount}} $(rly ch addr $DST)

    # check balance again
    rly q bal $SRC
    rly q bal $DST
    # relay packets to complete tokens transfer
    rly tx relay-packets $PTH
    rly q bal $SRC
    rly q bal $DST


    # You can send back ibc tokens from dst to src
    rly tx transfer $DST $SRC {{amount}} $(rly ch addr $SRC)
    ```

**BONUS** Transfer `uregen` tokens from `regen-devnet-4` to `aplikigo-1`