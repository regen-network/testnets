# Phase-3 Instructions: Deploy ERC20 token smart contract

This guide helps you to deploy an ERC20 token smart contract on [Aplikigo testnet](https://aplikigo.regen.ankea.io/)

Note: This document borrows most of the instructions from [COSMWASM official docs](https://docs.cosmwasm.com/0.13/getting-started/intro.html).

Don't be in a hurry, please read all the instructions before proceeding. We appreciate your time!

## Pre-requisites

### Install Rust

`rustup` is an installer for the systems programming language [Rust](https://www.rust-lang.org/)

Run the following in your terminal, then follow the onscreen instructions.

```
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Once installed, make sure you have the wasm32 target, both on stable and nightly:
```sh
$ rustup default stable
$ cargo version
# If this is lower than 1.47.0+, update
$ rustup update stable

$ rustup target list --installed
$ rustup target add wasm32-unknown-unknown
```

## Download and edit/update the contract

### Get the code

```
$ git clone https://github.com/CosmWasm/cosmwasm-plus
$ cd cosmwasm-plus
```

### Compiling

To compile all the contracts, run the following in the repo root:

```
docker run --rm -v "$(pwd)":/code \
  --mount type=volume,source="$(basename "$(pwd)")_cache",target=/code/target \
  --mount type=volume,source=registry_cache,target=/usr/local/cargo/registry \
  cosmwasm/workspace-optimizer:0.10.7
```

This will compile all packages in the `contracts` directory and output the
stripped and optimized wasm code under the `artifacts` directory as output,
along with a `checksums.txt` file.

If you hit any issues there and want to debug, you can try to run the 
following in each contract dir:
`RUSTFLAGS="-C link-arg=-s" cargo build --release --target=wasm32-unknown-unknown --locked`

## Deploy your erc20 contract

All the wasm related commands can be found at:
```
regen tx wasm -h
```

And, related queries at: 
```
regen query wasm -h
```

### Step - 1 : Upload your contract (Optional)
You can create/upload a new contract code or instantiate an existing code too.

To upload a new contract,

```
regen tx wasm store artifacts/cw20_base.wasm --chain-id aplikigo-1 --from <key_name> -y
```

For more details about uploading contract, check the details here: https://docs.cosmwasm.com/0.13/getting-started/interact-with-contract.html

### Step - 2: Instantiating the Contract

Note: Query the tx hash from above and get your `CodeID` details

Instantiate the contract using the uploaded code using this command.
```sh
regen tx wasm instantiate <CodeID> '{"name": "Witval Token", "symbol": "UWIT", "decimals": 6, "initial_balances": [{"address": "<any_account_address>", "amount": "1000000"}], "mint": {"minter": "<your_account_address>"}}' --label "cw20base" --from <yourkey> --chain-id=aplikigo-1 --amount=1000000utree
```

Where:
    `amount` is Coins to send to the contract during instantiation

Note: Query the tx with hash and get `contract_address`

### Step-3: Execute contract functions
Mint tokens using the following command
```sh
regen tx wasm execute <your_contract_address> '{"mint": {"recipient": "<your_account_address>", "amount": "1000000"}}' --from <your_key> --chain-id=aplikigo-1
```

### Step-4: Transfer some tokens
```sh
regen tx wasm execute <your_contract_address> '{"transfer": {"recipient": "<your_account_address>", "amount": "1000000"}}' --from <your_key> --chain-id aplikigo-1
```

## What is expected from validators?

1. Upload cw20-base code and Copy the TX HASH
2. Instantiate contract with validator label and copy TX Hash, Contract address
3. Mint tokens and copy the txHash
3. Transfer tokens to any validator and copy hash
4. Fork testnets repo and clone: githbu.com/regen-network/testnets
5. cd aplikigo-1/phase-3/task-1/
6. cp sample.json <your_validator_moniker>.json
7. Edit the details
8. Commit the changes to your repo
9. Raise PR with title: "Phase-3 | Task-1 : <Validator_moniker>"

## Points criteria
- 20 points for successful deployment (code upload is optional, instance creation is mandatory)
- 40 points for transfering contract tokens
- 40 points for minting tokens
- 20 points bonus for executing all other messages (send, burn, burnFrom, trasferFrom, sendFrom, ...). Max 100 points
- 50 points bonus for adding & executing any new message. Max 100 points each

**Note:** 
- There will be a special bonus of 100 points for each bug/vulnerability reported (non-duplicate), malfunctioning the network.
- All the edited contracts must be deployed using your validator owner account.
