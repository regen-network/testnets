# COSMWASM Kontraŭa Testnet - Plan

## Dates:  13th March - 17th Apr, 2020

## Total Points to be allocated: 1100+

## Schedule
### Phase - 1: Get of the ground

1. Gentx submissions: 9th March 0900UTC - 12th March 0900 UTC
2. Genesis release time: 12th March, 1600UTC (23 hours before genesis time)
3. Network start time: 13th March, 1500UTC

### Phase - 2: The **"wow, it's monday morning"** task: 16th March, 0500UTC - 19th March, 0500UTC

1. Deploy a token contract (ERC20). Rankings are based on deployment time. Contract's name should be validator's name
2. Give away (ERC20) tokens to 5 validators
3. Special bonus for adding allowance (it's a bonus, go get it mode)

### Phase - 3: Ameliorate network (**High-noon upgrade**) 26th March, 1200UTC
1. Upgrade Name: Himalaya
2. Upgrade proposal : 23rd March, ~1200 UTC
3. Binary release - 23rd March, ~1800 UTC
4. Voting Period : 23rd - 25th March (Relative to proposal time)
5. Upgrade time: 26th March, 1200UTC

Note: Instructions will be shared on 24th March, 1200UTC

### Phase - 3.1: Skip upgrade (**twilight drama**) - 02 Apr, 1600UTC
There might be a case where the upgrade binaries will be available only after the proposal is Passed. What if something is wrong with the new binary? Skipping the upgrade is a best way to handle such cases.
1. Upgrade proposal : 30 March, 1200 UTC
2. Binary release - 31 March, 1600 UTC
3. Voting Period : 30, March - 1 Apr
4. Upgrade time: 02 Apr, 1600UTC

Note: Instructions will be shared on 1 Apr, 1200UTC

### Phase - 4: Deploy contract - 2 (Escrow contract) - 7th Apr, 2100UTC - 9th Apr, 2100UTC
1. Deploy escrow contract
2. Execute escrow

Note: Instructions will be shared on 7 Apr, 1200UTC

### Phase - 5: High skill time - 14 Apr, 0200UTC - 18 Apr, 0200 UTC (Bonus phase)
1. Custom contracts deployment based on a provided criteria
2. Criteria & instructions will be shared on: 13 Apr, 2100UTC

### Phase - 6: Silent blast. Expect chaos - 20th Apr 0900UTC - 24th Apr 0900 UTC 
Intentionally keeping minimum information possible.
1. Tolerate spam transactions
2. DDoS attack

Note: Expect chaos


## Incentive Program

### Phase-1: 50+ points
- Genesis Score: 50 points
    - Gentx submission & approval - 25 points
    - Genesis block validator / 15000th block validator (applicable for only genesis validators) - 25 points
- Gentx submission time: A total of 400 bonus points are shared for first 20 valid gentx submissions (non-edited). Edited gentx submissions are not eligible for this. 
    - First 10 - 25 each
    - 10 to 20: 15 each

Note: Due to seed nodes issue on genesis time, uptime calculations are calculated from 15000th block. 

### Phase-2: 250+ points
- 100 points for successful deployment
- 50 points for adding source url and build tag (Without special instruction docs)
- 50 points for transfering contract tokens
- 50 points for editing the contract to add any custom feature (send text message, request funds, approve fund request, data storage)
- 100 bonus points for creating an allowance and use transferFrom to send tokens from a second address. 
- A total of 1000 bonus points are shared among first 20 validators to complete these tasks
    - First 5 - 100 each
    - 6 to 10 : 50 each
    - 11 to 20: 25 each

### Phase-3: 100+ points
- 25 points for casting vote
- 25 points for tweeting about the proposal and vote (with why?) and putting a link to the tweet in the memo of a transaction. 
- 50 points for successful upgrade
- A total of 1000 bonus points are shared among first 20 validators who upgrade their nodes (A periodic state dump is used to find out the winners)
    - First 200 points
    - 2 to 5: 75 points each
    - 6 to 10: 50 points each
    - 11 to 20: 25 points each 

### Phase-3.1: 100+ points
- 25 points for casting vote
- 25 points for tweeting about the proposal and vote (with why?)
- 50 points for successful upgrade (skip)
- A total of 1000 bonus points are shared among first 20 validators who upgrade their nodes (A periodic state dump is used to find out the winners)
    - First 200 points
    - 2 to 5: 75 points each
    - 6 to 10: 50 points each
    - 11 to 20: 25 points each

### Phase-4: 100+ points
- 50 points for successful deployment
- 50 points for executing the escrow
- A total of 1000 bonus points are shared among first 20 validators to complete these tasks
    - First 5 - 100 each
    - 6 to 10 : 50 each
    - 11 to 20: 25 each

### Phase-5: 300+ points
- 150 points for writing contract
- 100 points for executing the contract
- 50 points for executing transactions on other's contracts (if needed external validators support)

### Phase-6: 200+ points
- 90-100% uptime will share 100 points. 90% uptime will get 0 points, 100% uptime shares 100 points. (Points distribution is linear mostly, more details coming soon)
- 100 points if a validator never miss a block in the phase  single block (highly available - highly secured)
- Special bonus of 200 points each for performing and proving the attacks. The proved attack should take the specific/a group of validators down or create uptime issues.

## Note
At each phase, there will be a special bonus of 100 points for each bug reported, malfunctioning the network. Top 5 active community supporters will receive a special bonus too. 50 bonus points for never jailed validators.
