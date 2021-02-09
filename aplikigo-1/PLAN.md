# Regen Network - Aplikiĝo Testnet Planning

## Disclaimer
- Everything is subject to change. Pay attention to the discord [dev-validator announcement channel] (https://discord.gg/ePFC8edB) for announcements. We are testing complex, opensource beta software. Expect things to break. If the dcoumentation is inadequit, make a PR. 
## Overview
- 3 Phases (min)
- 4 governance proposals, 2 Upgrades, 2 community spend proposals 
- Rewards: 1800+ points possible per participant with a total of 500,000 Regen tokens available for winners. 
- Schedule: Feb 8th - Mar 8th

## Phase - 1: Genesis Phase - max 400 points
- Gentx submission : 5th-7th Feb, 2021 1500 UTC. Up to 125 Gentx will be accepted.
- Genesis file release: 7th Feb, 2021 1600 UTC
- Network start time: 8th Feb, 2021 1500 UTC - 50 points for everyone who brings up their nodes in first 10 blocks (votes).
- Celebrate Stargate: - 100 points
    - Send 1TREE to others - 10 points
    - Delegate some TREEs to others - 10 points
    - Tweet url linked in a memo - 30 points
    - LinkedIn post url linked in a memo - 30 points
    - memo with your discord handle - 20 points (come to discord, leave telegram behind!)
    - Blogpost describing about Regen  linked in memo
        - 20 teams will get 50 each (Regen network team will review the blogs and finalize top-20)
    - Participants get to raise a PR with the details onto regen-ledger/testnets repo
        - Time: Feb 9th, 1500UTC - Feb 11th, 1200 UTC (Only the txs in this time period are considered to be valid)
        - Instructions to submit the PR:
            - Clone regen-network/testnets repo,
            ```sh
            $ git clone https://github.com/regen-network/testnets
            $ cd testnets
            $ git pull origin master
            $ cd aplikigo-1/tasks/phase-1/1-celebrate
            $ cp sample.json <your_moniker>.json
            ```
            - Add/Update the details
            - Push to the repo and create a PR

- Network security and performance testing - upto 100 points:
    - Make transactions for continuous 3 hours from 1500UTC to 1800UTC 11 Feb, 2021
    - Top 10 teams will get 100 points each
    - Next 10 (11 to 20) will get 60 points each
    - 21 to 50 teams will get 50 points each
- Upgrade-1 Proposal: 15th Feb, 2021 ~1500 UTC
- Voting Period: 15th Feb-17th Feb, ~1500 UTC
- Upgrade Time: 18th UTC
    - First 10 teams to sign the upgrade block will get 100points each (upgrade height + 1's consensus state dump + prevotes will be used).
    - Next 20 teams will get 60 points each
    - Next 20 teams will get 40 points each
    - Other teams who are part of the upgrade will share 20 points each
    - Note: Teams are eligible for this reward only if they vote on the proposal

## Phase - 2: Regen Times  - max 600 points [TBD]
- Upgrade proposal to enable IBC transfers
- Test IBC Transfers of $REGEN token
- Test eco-credit module
- Deploy eco-credit token contract
- Deploy cosmwasm (?) contract

## Phase - 3: Internet of Regeneration - 800 points
- Test Groups Module
- Test Data Module
- Run solo machine and send token to Regen Ledger
- Run Small PoA PoS chain and transact with Regen Ledger 

## Bonus Challenges
- Uptime - 4000points
    - 2000 points will be distributed among who never misses a block (max 200 points per validator)
    - uptime >= 99% 1000 points will be distributed equally among all the eligible validators  (max 100 points per validator)
    - 99% <= uptime >= 98% 1000 points  (max 60 points per validator)
- Never jailed validators - 2000 points
- Max txs in a block (min 5000txs/block) - 100 points
    - If multiple blocks has same number of txs, only the txs from single or list of accounts owned/organized by the validator will be counted. Txs from others in the block will not be considered.
- Max msgs in a transaction - 100 points
    - If multiple blocks has same number of messages in a tx, first 2 teams will share the reward (50 points each)
- P2P/Mempool attacks - Bringing 60% nodes down (Should publish a blogpost with details and possible proofs) - 300 points
- Exploits/bug bounty for x/groups and x/authz modules (Regen team will review and categorize the criticality into following categories)
    - Low - 20 points
    - Medium - 50 points
    - High - 100 points
    - Critical - 200points
- Community reward
    - Top - 10 teams/individuals will receive 50points each for their contributions for the community. technical docs, helping/resolving community issues, etc.

## NOTE:
Regen Network is committed to build a strong community. We would like to extend our thanks to our  early adapters & supporters and they are given priority over new participants wherever there's a tie. Here are a possible cases (but not limited to) where we extend our support and prioritize early community members (participants from previous testnets, our partners and investors)
- Allocate 10% extra tokens on genesis (gentxs)
- Ranking boost (whenever there's a tie)

## Code Of Conduct

- Testnet tokens are limited and valuable assets for the testnet. It is restricted/not allowed to receive external delegations. One should not try to increase their voting power by spamming the faucet. But it is encouraged to increase their stake by re-staking their rewards.  There are no incentives related to the number of tokens you attain.  The faucet should be used as a utility.  
- We expect formal and professional behaviour from the participants and encourage a healthy competition, as well as healthy cooperation. Any misbehaviour will potentially disqualify one from the contest. This includes trolling channels, and being dismissive or rude to teammembers or other testnet participants.  
- Participating as a group or running validators from multiple accounts should be strictly avoided and any hint on such activities would disqualify users staraight away. Everyone involved would get ZERO rewards. We are not running this testnet to uncover cartel behavior.  There are testnets where that is perfectly acceptable and encouraged. Find those test nets. This testnet is explicitly to test Regen Ledger functionality and performance and create an opportunity for a livlihood helping to secure and govern the worlds firs tpublic ecological ledger.  This is not an advisarial testnet in which the basic assumptions of PoS systems are being tested. Don't encourage your friends to spin up nodes just for the rewards sake.
- If a user is flagged or blocked on Discord/Github/Twitter will be considered ineligible. 
- Same account/name should be used throughout the testnet. It's not allowed to raise PRs from different github accounts.
