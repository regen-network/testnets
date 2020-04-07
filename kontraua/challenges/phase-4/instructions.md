# Phase-4 Instructions: Deploy Escrow token smart contract

This guide helps you to deploy an escrow token smart contract on [Kontraua testnet](https://regen.wasm.glass/)

Note: This document borrows most of the instructions from [COSMWASM official docs](https://www.cosmwasm.com/docs/getting-started/intro), thanks to **Ethan Frey** and team.

Don't be in a hurry, please read all the instructions before proceeding. We appreciate your time!

## Pre-requisites

### Install Rust

Please feel free to refer [rust basics](https://www.cosmwasm.com/docs/getting-started/rust-basics) for more details.

`rustup` is an installer for the systems programming language [Rust](https://www.rust-lang.org/)

Run the following in your terminal, then follow the onscreen instructions.

```
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Once installed, make sure you have the wasm32 target, both on stable and nightly:
```
$ rustup default stable
$ rustup target list --installed
$ rustup target add wasm32-unknown-unknown

$ rustup install nightly
$ rustup target add wasm32-unknown-unknown --toolchain nightly
```

## Download and edit/update the contract

### Get the code

```
$ git clone https://github.com/cosmwasm/cosmwasm-examples
$ cd cosmwasm-examples/escrow
$ git checkout escrow-0.3.0
```

### Compile the wasm contract with stable toolchain
```
rustup default stable
cargo wasm
```
After this compiles, it should produce a file at  `target/wasm32-unknown-unknown/release/escrow.wasm`. A quick `ls -l` should show around 1.5MB. This is a release build, but not stripped of all unneeded code

### Compiling for Production
You can check the size of the contract file by running:
```
$ du -h target/wasm32-unknown-unknown/release/cw_escrow.wasm

// Outputs
1.9M    target/wasm32-unknown-unknown/release/cw_escrow.wasm
```
This works, but is huge for a blockchain transaction. Let's try to make it smaller. Turns out there is a linker flag to strip off debug information:

```
$ RUSTFLAGS='-C link-arg=-s' cargo wasm
$ du -h target/wasm32-unknown-unknown/release/cw_escrow.wasm

// Outputs
128K    target/wasm32-unknown-unknown/release/cw_escrow.wasm
```
This is looking much better in size.

Those who wants to experiment with other cool features can try out: [reproduceable builds](https://www.cosmwasm.com/docs/getting-started/editing-escrow-contract#reproduceable-builds)

## Deploy your escrow contract

All the wasm related commands can be found at:
```
xrncli tx wasm -h
```

And, related queries at: 
```
xrncli query wasm -h
```

NOTE: Please feel free to request tokens in DVD channel if your validator account doesn't have available tokens

### Step - 1 : Upload your contract

To upload a new contract,

```
xrncli tx wasm store cw_escrow.wasm --gas auto --from <key_name> --node <rpc_endpoint> -y

#Add keys for arbiter, recipient and thief

xrncli keys add fred
xrncli keys add bob
xrncli keys add thief
```

You can check your code id at (by your account address): https://regen.wasm.glass/codes

For more details about uploading contract, check the details here: https://www.cosmwasm.com/docs/getting-started/first-demo#uploading-the-code

### Step - 2: Instantiating the Contract


```
# Please make sure to add keys of bob, fred and thief to your keyring

INIT="{\"arbiter\":\"$(xrncli keys show fred -a)\", \"recipient\":\"$(xrncli keys show bob -a)\", \"end_time\":0, \"end_height\":0}"

xrncli tx wasm instantiate <code_id> "$INIT" --from <key_name> --label "escrow 1 <moniker>" --node <rpc_endpoint> --chain-id kontraua --amount 5000utree -y 
```

Verify your code instance:
```
# check the contract state (and account balance)
xrncli query wasm list-contract-by-code <code_id>  --node <rpc_endpoint> --chain-id kontraua

# contracts ids (like code ids) are based on an auto-gen sequence
# if this is the first contract in the devnet, it will have this address (otherwise, use the result from list-contract-by-code)
CONTRACT=xrn:10pyejy66429refv3g35g2t7am0was7ya75d7y2

# query contract to verify init message and balances
xrncli query wasm contract $CONTRACT --node <rpc_endpoint> --chain-id kontraua --trust-node

# you can query contract address as normal account
xrncli query account $CONTRACT --node <rpc_endpoint> --chain-id kontraua --trust-node

# you can dump entire contract state
xrncli query wasm contract-state all $CONTRACT

# note that we prefix the key "config" with two bytes indicating it's length
# echo -n config | xxd -ps
# gives 636f6e666967
# thus we have a key 0006636f6e666967

# you can also query one key directly
xrncli query wasm contract-state raw $CONTRACT 0006636f6e666967 --hex

# Note that keys are hex encoded, and val is base64 encoded.
# To view the returned data (assuming it is ascii), try something like:
# (Note that in many cases the binary data returned is non in ascii format, thus the encoding)
xrncli query wasm contract-state all $CONTRACT | jq -r .[0].key | xxd -r -ps
xrncli query wasm contract-state all $CONTRACT | jq -r .[0].val | base64 -d

```

### Step-3: Execute contract functions

Set the Approve msg:
```
APPROVE='{"approve":{"quantity":[{"amount":"2000","denom":"utree"}]}}'
```

Execute the approve function 
```
xrncli tx wasm execute $CONTRACT "$APPROVE" --from <key-name> -y --chain-id kontraua
# looking at the logs should show: "execute wasm contract failed: Unauthorized"
# and bob should still be broke (and broken showing the account does not exist Error)
xrncli query account $(xrncli keys show bob -a)

# but succeeds when fred tries
xrncli tx wasm execute $CONTRACT "$APPROVE" --from fred -y
xrncli query account $(xrncli keys show bob -a)
xrncli query account $CONTRACT
```

Set the steal msg:
```
STEAL="{\"steal\":{\"destination\":\"$(xrncli keys show thief -a)\"}}"
```
Execute the steal function
```
xrncli tx wasm execute $CONTRACT "$STEAL" --from thief -y

xrncli query account $(xrncli keys show thief -a)

xrncli query account $CONTRACT
```



## What is expected from validators?

1. Upload escrow code and Copy the TX HASH
2. Instantiate contract with validator label and copy TX Hash, Contract address
3. Fork testnets repo and clone: github.com/regen-network/testnets
4. cd testnets/kontraua/challenges/phase-4/
5. cp sample.json <your_validator_moniker>.json
6. Edit the details
7. Commit the changes to your repo
. Raise PR with title: "Phase-4: <Validator_moniker>"

## Points criteria
- 50 points for successful deployment
- 50 points for executing the escrow
- A total of 1000 bonus points are shared among first 20 validators to complete these tasks
    - First 5 - 100 each
    - 6 to 10 : 50 each
    - 11 to 20: 25 each

**Note:** 
- There will be a special bonus of 100 points for each bug/vulnerability reported (non-duplicate), malfunctioning the network.

- All the edited contracts must be deployed using your validator owner account.

# Important Links

- Code explorer: https://regen.wasm.glass/
- Network explorer: https://explorer.regen.vitwit.com
- Testnet plan: https://medium.com/regen-network/cosmwasm-kontra%C5%ADa-testnet-plan-2756490ccdf4
- Regen Network DVD Channel: https://t.me/joinchat/FJGNSxOpjJcgrUGwAAOKUg
- Testnet instructions https://github.com/regen-network/testnets/blob/master/kontraua/README.md
