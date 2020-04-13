# Payout Contract for ecological credits

## Introduction
Providing payout contracts for the ecological state was the original reason behind **cosmwasm**. This phase focuses on building a custom payout contract to release tokens to the beneficiary based on the ecological state change.

Implement a payout contract to allow a beneficiary to receive the tokens.

### Payout Contract Requirements

- Contract should implement init with `region`, `total_tokens`, `oracle`,  `beneficiary` addresses. You can set optional params : `payout_start_height`, `payout_end_height`
- Initialize the contract with 1000 coins with a specified beneficiary
- At some interval (10minutes lets say), the oracle (a thrid party account in general) provides the percentage of forest cover for the region. To avoid floats in the contract, multiply the percentage with 100 and trim  decimals. For example: 12.39% is inputed as 1239, 1.004% is inputed as 100.
* If the forest cover has decreased since last measurement - no payout. It means, the beneficiary will not get any token rewards.
* If it is more than 1$, pay out 100 coin per 1% increase. So if there's 5.31% increase in the forest cover, beneficiary would get 531 coins.
* If it is the same or less than 1% (0% < change% < 1%), pay out 20 coins for current forest cover above 50% (eg. 0 at 50%, 30 at 65%, 78 at 89%)

### Incentive Plan

#### Working Payout Contract - 300 Points
To be eligible for the reward, the contract should at least do one sample payout with external oracle trigger (not direct transfer of tokens)
A contract is considered to be a Payout Contract if it meets the following criterias:
- Should be able to initialize with oracle, owner, beneficiary (INIT)
- Oracle should (only) be able to send/update ecostate
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

## Helpful code snippets

### State

```rust=
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub region: String,
    pub beneficiary: CanonicalAddr,
    pub owner: CanonicalAddr,
    pub oracle: CanonicalAddr,
    pub ecolstate: i64,
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
$ INIT = "{\"region\":\"region-1\",\"beneficiary\":\"$(wasmcli keys show beneficiary -a)\",\"oracle\":\"$(wasmcli keys show oracle -a)\",\"ecostate\":1070,\"total_tokens\":10000,\"released_tokens\":0,\"payout_start_height\":460000,\"payout_end_height\":1000000,\"is_locked\":0}"
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

Let's update teh ecostate change to 12.1% i.e., 1210
```sh
$ UPDATE_ECOSTATE='{"update_ecostate":{"ecostate": 1210}'
```

```sh
$ xrncli tx wasm execute $CONTRACT "$UPDATE_ECOSTATE" --from oracle -y
```

- Query contract state, `released_tokens` should have been updated based on the ecostate change.