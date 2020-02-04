# Darien Gap software upgrade details and voting instructions.

The second software proposal is titled "Darien Gap Upgrade". This upgrade will be applied on 7th Feb, 2020 at 09:00UTC. This proposal proposes to switch the working binary to v0.5.3. This release can be found here:- https://github.com/regen-network/regen-ledger/releases/tag/v0.5.3

To query this proposal using the client, you can use this command on your validator or proxy node:
```
xrncli query gov proposal 5 --chain-id algradigon-1 -o json --node https://regen.chorus.one:26657
```
Note: The ```--node``` flag indicates that we are connecting to Chorus one node. If you are executing this command on your local node/validator, this flag is not needed.

This will return a response that looks something like this:
```
{ 
   "content":{ 
      "type":"cosmos-sdk/SoftwareUpgradeProposal",
      "value":{ 
         "title":"Darien Gap Upgrade",
         "description":"Upgrade to Darien Gap release (v0.5.3)",
         "plan":{ 
            "name":"darien-gap",
            "time":"2020-02-07T09:00:00Z",
            "info":"Release tag v0.5.3"
         }
      }
   },
   "id":"4",
   "proposal_status":"VotingPeriod",
   "final_tally_result":{ 
      "yes":"0",
      "abstain":"0",
      "no":"0",
      "no_with_veto":"0"
   },
   "submit_time":"2020-02-04T13:00:46Z",
   "deposit_end_time":"2020-02-04T13:15:46Z",
   "total_deposit":[ 
      { 
         "denom":"utree",
         "amount":"50000000"
      }
   ],
   "voting_start_time":"2020-02-04T13:15:00Z",
   "voting_end_time":"2020-02-06T13:15:00Z"
}
```

As we can see in the ``` proposal_status```, we are currently in the ```VotingPeriod```. This means that all the token holders can and should vote on the proposal before the voting period ends. Voting period for this proposal is scheduled to end on 6th Feb, 2019 at 13:15UTC. Votes cast after this time will not be registered on-chain and the voters will miss out on points.

## Voting for the Software Upgrade proposal

You can vote for this proposal using the client by using the following command:
```xrncli tx gov vote 5 yes --from <mykey> --chain-id algradigon-1 --node https://regen.chorus.one:26657```

Note:
You could also vote ```no``` or ```abstain``` on the proposal if you're not confident about it.


