# Payout Contract for ecological credits

## Introduction
Providing payout contracts for the ecological state was the original reason behind **cosmwasm**. This phase focuses on building a custom payout contract to release tokens to the beneficiary based on the ecological state change.

Implement a payout contract to allow a beneficiary to receive the tokens.

### Payout Contract Requirements

- Contract should implement init with `region`, `ecostate`, `total_tokens`, `oracle`,  `beneficiary` addresses. You can set optional params : `payout_start_height`, `payout_end_height`, `is_locked` as well.
    - `region` is a string, refers to a place/zone
    - `ecostate` is a positive integer, which maintains the percentage of forest cover, without decimals. 56.16% is stored as 5616
    - `total_tokens` is a positive integer, locked in the contract
    - `oracle` is the account address of ecological state data provider (It can be any general account with some tokens for gas cost)
    - `beneficiary` is the account address of beneficiary, which receives token payouts
    - `payout_start_height` is the block height from which the contract is valid
    - `payout_end_height` is the block height until which the contract is valid
    - `is_locked` maintains contract status. It allows the owner to lock the contract in case of any malicious activities.
- Initialize the contract with 100,000 coins (`total_tokens`) with a specified beneficiary
- At some interval (10 minutes lets say), the oracle (a thrid party account in general) provides the percentage of forest cover for the region. To avoid floats in the contract, multiply the percentage with 100 and trim  decimals. For example: 12.39% is inputed as 1239, 1.004% is inputed as 100.
* If the forest cover has decreased since last measurement - no payout. It means, the beneficiary will not get any token rewards.
* If it is more than 1%, pay out 100 coin per 1% increase. So if there's 5.31% increase in the forest cover, beneficiary would get 531 coins.
* If it is the same or less than 1% (0% < change% < 1%), pay out 2 coins for current forest cover above 50% (eg. 0 at 50%, 30 at 65%, 78 at 89%). Round the tokens to next integer if you are not using any decimals. Like 57.6 at 78.8% would result in a payout of 58 tokens

### Incentive Plan

#### Working Payout Contract - 300 Points
To be eligible for the reward, the contract should at least do one sample payout with state update trigger (direct token transfer are not eligible for the reward)
A contract is considered to be a Payout Contract if it meets the following criterias:
- Should be able to initialize with oracle, owner, beneficiary (INIT)
- Oracle (account) should (only) be able to update the ecostate
- Implements at least one payout criteria from the points mentioned above. The contract should be able to receive and process the oracle's data
- At least one successful payout trasaction (PAYOUT)

#### Bonus Payout - 50 points for each of the below features
- Lock & Unlock contract
    - Should stop processing the ecostate data from oracle if the lock is set and thus no payouts should be processed. Unlock should allow the lock to be removed.
- Use Optional validations
    - Start height/time and end height/time. 
        - The contract will funciton only if the height/time fullfills the check `start <= current <= end`
    - Contract status should change to DONE when it releases all the tokens. For the last payout, pay everything available if the required tokens are more than available. If the contract status is DONE, it shouldn't accept/process any oracle ecostate data
- Update state
    - Transfer Ownership
    - Change Beneficiary
    - Change Oracle
- Implemets query contract state (Should return the contract's state data)
- Query balance
- Any custom feature resembling/addressing real PAYOUT CONTRACT scenarios

#### A total of 1000 bonus points are shared among first 20 validators to complete these tasks
- First 5 - 100 each
- 6 to 10 : 50 each
- 11 to 20: 25 each

## Helpful code snippets (This is all optional, you can refer cosmwasm official docs for more details)

### State

```rust=
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub region: String,
    pub beneficiary: CanonicalAddr,
    pub owner: CanonicalAddr,
    pub oracle: CanonicalAddr,
    pub ecostate: i64,
    pub total_tokens: i64,
    pub released_tokens: i64,
    pub payout_start_height: i64,
    pub payout_end_height: i64,
    pub is_locked: i64,
}
```

### Messages
Init Message
```rust=
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {
    pub region: String,
    pub beneficiary: HumanAddr,
    pub oracle: CanonicalAddr,
    pub ecostate: i64,
    pub total_tokens: i64,
    pub payout_start_height: i64,
    pub payout_end_height: i64,
}
```

Handle Message
```rust=
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "lowercase")]
pub enum HandleMsg {
    UpdateEcostate {ecostate: i64},
    Lock {},
    UnLock {},
    ChangeBeneficiary {beneficiary: HumanAddr},
    TransferOwnership {owner: HumanAddr},
}
```

Query Message
```rust=
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "lowercase")]
pub enum QueryMsg {
    State {},
    Balance {address: HumanAddr}
}
```

### Sample Txs to help with contract design

It assumes that, you have beneficiary, owner and  oracle accounts with minimum funds required. Also, your code is successfully compiled and uploaded.
```sh
$ CODE_ID=<your_code_id>
```

NOTE: This is just to give a way to get started. You can change it as you wish.

**Init message:**
```sh
$ INIT = "{\"region\":\"region-1\",\"beneficiary\":\"$(wasmcli keys show beneficiary -a)\",\"oracle\":\"$(wasmcli keys show oracle -a)\",\"ecostate\":2500,\"total_tokens\":100000,\"released_tokens\":0,\"payout_start_height\":460000,\"payout_end_height\":1000000,\"is_locked\":0}"
```

**Instantiate code:**
```sh
$ xrncli tx wasm instantiate $CODE_ID "$INIT" --from owner --label "ecostate 1" -y
```

###### check the contract state
```sh
$ xrncli query wasm list-contract-by-code $CODE_ID
```
###### contracts ids (like code ids) are based on an auto-gen sequence, use the result from list-contract-by-code

```sh
$ CONTRACT=<your_contract_address>
$ xrncli query wasm contract $CONTRACT
```

**Update ecological state (push some data from oracle)**

Let's update teh ecostate to 27.1% i.e., 2710
```sh
$ UPDATE_ECOSTATE='{"update_ecostate":{"ecostate": 2710}'
```

```sh
$ xrncli tx wasm execute $CONTRACT "$UPDATE_ECOSTATE" --from oracle -y
```
- Query contract state response should show the updated `released_tokens` based on the ecostate change

## What is expected from validators?

1. Implement Payout Contract, compile and upload
2. Instantiate contract with validator label and copy Contract address
3. Update the ecostate using oracle's account
4. Copy all the tx hashes
4. Fork testnets repo and clone: githbu.com/regen-network/testnets
5. cd testnets/kontraua/challenges/phase-5/
6. cp sample.json <your_validator_moniker>.json
7. Add tx hashes from 4 in `oracleTxs` section. You can add array of tx hashes.
8. If you have added any extra functionality as stated in [Bonus Payout](#bonus-payout---50-points-for-each-of-the-below-features), add the respective tx hashes and feature title/description in `customTxs` section.
8. Commit the changes to your repo
9. Raise PR with title: "Phase-5: <Validator_moniker>"

**Note:**
- For early submission rewards, only tx time is considered (Last tx from the list). PR time will not play a role in distributing bonus rewards.
- Contract should be unique in-order to be eligible for the reward
- There will be a round of code submission after the deadline.
- There will be a special bonus of 100 points for each bug/vulnerability reported (non-duplicate), malfunctioning the network.
- All the edited contracts must be deployed using your validator owner account.
- Expect chaos : Custom oracle service?

# Important Links
- Cosmwasm docs: https://www.cosmwasm.com/docs/getting-started/intro
- Code explorer: https://regen.wasm.glass/
- Network explorer: https://explorer.regen.vitwit.com
- Testnet plan: https://medium.com/regen-network/cosmwasm-kontra%C5%ADa-testnet-plan-2756490ccdf4
- Regen Network DVD Channel: https://t.me/joinchat/FJGNSxOpjJcgrUGwAAOKUg
- Testnet instructions https://github.com/regen-network/testnets/blob/master/kontraua/README.md
