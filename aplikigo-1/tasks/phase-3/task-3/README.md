# Groups Module

## Tasks
### Create Group
Create a group which is an aggregation of member accounts with associated weights and an administrator account. Note, the '--from' flag is ignored as it is implied from [admin].

```sh
$ regen tx group create-group [admin] [metadata] [members-json-file]
```
Where
- `admin` is your account address
- `metadata` is the description for your group, should be a base64 encoded string
- `members.json` contains:
```sh
{
	"members": [
		{
			"address": "addr1",
			"weight": "1",
			"metadata": "<some base64 encoded string here>"
		},
		{
			"address": "addr2",
			"weight": "1",
			"metadata": "<some base64 encoded string here>"
		},
		{
			"address": "addr3",
			"weight": "1",
			"metadata": "<some base64 encoded string here>"
		}
	]
}
```

### Query groups by admin
```
regen q group groups-by-admin <admin_account_address> --chain-id aplikigo-1
```

### Create Group Account
Create a group account which is an account associated with a group and a decision policy. Note, the '--from' flag is ignored as it is implied from [admin]. Metadata should be a base64 encoded string

Example:
```sh
$ regen tx group create-group-account [admin] [group-id] [metadata] \
'{"@type":"/regen.group.v1alpha1.ThresholdDecisionPolicy", "threshold":"2", "timeout":"1000s"}'
```

### Query group account by admin
```
regen q group group-accounts-by-admin <admin_account_address> --chain-id aplikigo-1
```

### Send some funds to group account
```
$ regen tx bank send <from_account> <group_account_address> 10000000utree --chain-id aplikigo-1 --from <from_account>
```
Query group account balances to cross check

### Create a sample Proposal to send tokens to some receipient
```sh
regen tx bank send <group_account_address> <recipient_address> 1000utree --from <group_account_address> --chain-id aplikigo-1 --generate-only > msg_tx.json
```

```sh
$  regen tx group create-proposal [group-account_address] [proposer_address] msg_tx.json [metadata] --from <proposer_address>
```

here `proposer_address` should be one of the group members and metadata should be a base64 encoded string

### Vote for the proposal
```sh
$  regen tx group vote [proposal-id] [voter] [choice] [metadata] [flags]
```

```sh
regen tx group vote -h
Vote on a proposal.

Parameters:
			proposal-id: unique ID of the proposal
			voter: voter account addresses.
			choice: choice of the voter(s)
				CHOICE_UNSPECIFIED: no-op
				CHOICE_NO: no
				CHOICE_YES: yes
				CHOICE_ABSTAIN: abstain
				CHOICE_VETO: veto
			Metadata: metadata for the vote
```
### Execute Proposal
- Query group account balances
- Execute the proposal
```sh
 regen tx group exec 4 --from <your_account>
```
- Query group account balances

## What is expected from validators?
1. Execute all the group related txs
2. Fork testnets repo and clone: githbu.com/regen-network/testnets
3. cd aplikigo-1/phase-3/task-3/
4. cp sample.json <your_validator_moniker>.json
5. Edit the details
6. Commit the changes to your repo
7. Raise PR with title: "Phase-3 | Task-3 : <Validator_moniker>"