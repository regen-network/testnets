# Phase - 3.1: Skip upgrade (**twilight drama**)

## Foreword
There could be cases where the planned upgrades might need to be skipped due to issues in new binaries or decision changes after the proposal goes through. Upgrade module allows a way to handle such cases with SKIP UPGRADE funcationaliy. This phase is about testing SKIP UPGRADE functionality.

## Schedule
1. Upgrade proposal : 30 March, 1200 UTC
2. Binary release - 31 March, 1600 UTC
3. Voting Period : 30, March - 1 Apr
4. Upgrade time: 02 Apr, 1600UTC

## Proposal Details
|    |            |
|----------|:-------------:|
| Proposal ID |  2 |
| Name |    twilight-drama   |
| Title | Twilight Drama | 
| Description | Twilight Drama Upgrade Proposal to test SKIP UPGRADE functionality from `upgrade` module
| Proposal Time | 2020-03-25 13:03:27 UTC |
| Voting Start Time | 2020-03-25 13:03:27 UTC |
| Voting End Time | 2020-03-25 13:03:27 UTC |
| Upgrade Height | 288888 |
| Link | https://regen-lcd.vitwit.com/gov/proposals/2 |   



## Voting for proposal

Use the following command to vote on the proposal.
```sh
xrncli tx gov vote 2 yes --chain-id kontraua --node http://<ip>:26657
```

Though you have `yes`/`no`/`abstain`/`no_with_veto` options to vote, it is recommended to chose only `yes` on the proposal as this is not about testing the upgrade and there won't be any changes on the network after this phase.

## How to SKIP UPGRADE

Upgrade module has a handly SKIP UPGRADE option. You can skip any known upgrade with it's HEIGHT.

Here is the general SKIP UPGRADE usage
```sh
xrnd start --unsafe-skip-upgrades <upgrade1_height> <upgrade2_height> <upgrade3_height>
```

To skip **Twilight Drama** upgrade, use the following command

```sh
xrnd start --unsafe-skip-upgrades 288888
```
