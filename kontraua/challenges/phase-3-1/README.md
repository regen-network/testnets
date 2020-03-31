# Phase - 3.1: Skip upgrade (**twilight drama**)

## Context - Why SKIP UPGRADE?
There could be cases where the planned upgrades might need to be skipped due to issues in new binaries or decision changes after the proposal goes through. `Upgrade` module provides us a way to handle such cases with **SKIP UPGRADE** functionality. This phase is about testing SKIP UPGRADE feature from `upgrade` module.

## Schedule
- Upgrade proposal time: 30 March, 1411 UTC
- Binary release (Dummy release) - 31 March, 1600 UTC
- Voting Period : 30, March - 1 Apr, 1417 UTC
- Upgrade Height: 288888

## Proposal Details
|    |            |
|----------|:-------------:|
| Proposal ID |  2 |
| Name |    twilight-drama   |
| Title | Twilight Drama | 
| Description | Twilight Drama Upgrade Proposal to test SKIP UPGRADE functionality from `upgrade` module
| Proposal Time | 2020-03-30 1411 UTC |
| Voting Start Time | 2020-03-30 1417 UTC |
| Voting End Time | 2020-04-01 1417 UTC |
| Upgrade Height | 288888 |
| Link | https://regen-lcd.vitwit.com/gov/proposals/2 |   

## Voting for proposal

Use the following command to vote on the proposal.
```sh
xrncli tx gov vote 2 yes --chain-id kontraua --node http://<ip>:26657 --from <key-name>
```

Though you have `yes`/`no`/`abstain`/`no_with_veto` options to vote, it is recommended to choose only `yes` as the proposal vote as this is not about testing the upgrade and there won't be any changes on the network after this phase.

## How to SKIP UPGRADE

Upgrade module has a handy SKIP UPGRADE option. You can skip any known upgrade with its HEIGHT. It can be done anytime after the upgrade proposal goes through, even after the binary panics with UPGRADE NEEDED message.

Here is the general SKIP UPGRADE usage
```sh
xrnd start --unsafe-skip-upgrades <upgrade1_height> <upgrade2_height> <upgrade3_height>
```

To skip **Twilight Drama** upgrade, just restart your `xrnd` instance with the `--unsafe-skip-upgrades` flag

### Using systemd

1. Stop your xrnd
```sh
sudo service xrnd stop
```
2. Edit your xrnd service with your favorite editor and add `--unsafe-skip-upgrades` flag to the start command
```
$ sudo nano /lib/systemd/system/xrnd.service
```
Edit in the following line to add `--unsafe-skip-upgrades`

```
xrnd start --pruning nothing --unsafe-skip-upgrades 288888
```

3. Update systemd
```sh
$ systemctl daemon-reload
```

4. Start your xrnd
```sh
sudo service xrnd start
```

### Manual
This is very simple, just stop the binary and restart it with the following command
```sh
xrnd start --pruning nothing --unsafe-skip-upgrades 288888
```



## Tweet bonus

Those who tweeted about network upgrade proposal, please raise a PR.
```
$ cd <path/to/testnets>/kontraua/challenges/phase-3-1
$ cp sample.json <your_moniker.json>
```

Update the required details and raise a PR with title: `Phase-3-1_<your_moniker>`
