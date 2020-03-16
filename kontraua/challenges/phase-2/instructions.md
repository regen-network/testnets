# Phase-2 Instructions: Deploy ERC20 token smart contract

This guide helps you to deploy an ERC20 token smart contract on [Kontraua testnet](https://regen.wasm.glass/)

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
$ cd cosmwasm-examples/erc20
$ git checkout erc20-0.3.0
```

### Compile the wasm contract with stable toolchain
```
rustup default stable
cargo wasm
```
After this compiles, it should produce a file at  `target/wasm32-unknown-unknown/release/cw_erc20.wasm`. A quick `ls -l` should show around 1.5MB. This is a release build, but not stripped of all unneeded code

### Compiling for Production
You can check the size of the contract file by running:
```
$ du -h target/wasm32-unknown-unknown/release/cw_erc20.wasm

// Outputs
1.9M    target/wasm32-unknown-unknown/release/cw_erc20.wasm
```
This works, but is huge for a blockchain transaction. Let's try to make it smaller. Turns out there is a linker flag to strip off debug information:

```
$ RUSTFLAGS='-C link-arg=-s' cargo wasm
$ du -h target/wasm32-unknown-unknown/release/cw_erc20.wasm

// Outputs
128K    target/wasm32-unknown-unknown/release/cw_erc20.wasm
```
This is looking much better in size.

Those who wants to experiment with other cool features can try out: [reproduceable builds](https://www.cosmwasm.com/docs/getting-started/editing-escrow-contract#reproduceable-builds)

## Deploy your erc20 contract

All the wasm related commands can be found at:
```
xrncli tx wasm -h
```

And, related queries at: 
```
xrncli query wasm -h
```

NOTE: Please feel free to request tokens in DVD channel if your validator account doesn't have available tokens

### Step - 1 : Upload your contract (Optional)
You can create/upload a new contract code or instantiate an existing code too (You can use https://regen.wasm.glass/codes/5 in case you chose to instantiate contract directly).

To upload a new contract,

```
xrncli tx wasm store contract.wasm --gas auto --from <key_name> --node <rpc_endpoint> -y
```

You can check your code id at (by your account address): https://regen.wasm.glass/codes

For more details about uploading contract, check the details here: https://www.cosmwasm.com/docs/getting-started/first-demo#uploading-the-code

### Step - 2: Instantiating the Contract

Please edit the **INIT** config with your details. You can add as many account address as you wish. Naming your token name unique is preferred.

```
# Please make sure to add your address(es) in initial_balances

INIT="{\"decimals\":5,\"name\":\"XRN token\",\"symbol\":\"XRN\",\"initial_balances\":[{\"address\":\"xrn:1tay3xlj0jnl8csnvteky8wc0e3206favkg74ue\",\"amount\":\"1000\"},{\"address\":\"xrn:1ntlzxh9y245htk99gz55gslz3n8lzclexenx9m\",\"amount\":\"2000\"},{\"address\":\"xrn:1vkxgpw4xtyeljzvqnxxy84kpa6udqaqw8leqjg\",\"amount\":\"3000\"}]}"

xrncli tx wasm instantiate <code_id> "$INIT" --from <key_name> --label "ERC20 <moniker>" --node <rpc_endpoint> --chain-id kontraua -y 
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
xrncli query wasm contract-state all $CONTRACT --node <rpc_endpoint> --chain-id kontraua
```

### Step-3: Execute contract functions

Set the transfer msg:
```
TRANSFER_MSG="{\"transfer\": {\"recipient\": \"xrn:1ntlzxh9y245htk99gz55gslz3n8lzclexenx9m\",\"amount\": \"2\"}}"
```

Execute the transfer function 
```
xrncli tx wasm execute $CONTRACT "$TRANSFER_MSG" --from <from_key>
```

**Check balance:**

You need to create a smart query with account address for querying balance:
```
BALANCE_QUERY="{\"balance\": {\"address\": \"xrn:1ntlzxh9y245htk99gz55gslz3n8lzclexenx9m\"}}"
```

Command to query balance is
```
xrncli --chain-id kontraua query wasm  contract-state smart $CONTRACT "$BALANCE_QUERY" --node http://173.255.192.172:26657 --chain-id kontraua -o json
```

## What is expected from validators?

1. Upload erc20 code and Copy the TX HASH
2. Instantiate contract with validator label and copy TX Hash, Contract address
3. Transfer tokens to 5 validators and copy hash
4. Fork testnets repo and clone: githbu.com/regen-network/testnets
5. cd testnets/kontraua/challenges/phase-2/
6. cp sample.json <your_validator_moniker>.json
7. Edit the details
8. Commit the changes to your repo
9. Raise PR with title: "Phase-2: <Validator_moniker>"

## Points criteria
- 100 points for successful deployment (code upload is optional, instance creation is mandatory)
- 100 points for transfering contract tokens (at least 5 transfers)
- 50 points for editing the contract to add any custom feature (send extra tokens than allowed,  hardcode recipient address, send text message, request funds, approve fund request, data storage etc)
- 100 bonus points for creating an allowance and use transferFrom to send tokens from a second address (Instructions for this are not available, as it is a bonus)
- A total of 1000 bonus points are shared among first 20 validators to complete these tasks (Eligibility: min 250 points earnings in the phase-2).
First 5–100 each
6 to 10 : 50 each
11 to 20: 25 each

**Note:** 
- There will be a special bonus of 100 points for each bug/vulnerability reported (non-duplicate), malfunctioning the network.

- All the edited contracts must be deployed using your validator owner account.

# Important Links

- Code explorer: https://regen.wasm.glass/
- Network explorer: https://explorer.regen.vitwit.com
- Testnet plan: https://medium.com/regen-network/cosmwasm-kontra%C5%ADa-testnet-plan-2756490ccdf4
- Regen Network DVD Channel: https://t.me/joinchat/FJGNSxOpjJcgrUGwAAOKUg
- Testnet instructions https://github.com/regen-network/testnets/blob/master/kontraua/README.md
