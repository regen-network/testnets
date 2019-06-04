#!/usr/bin/env bash
cat chat-addrs-cosmos.txt| xargs -L 1 convert_to_xrn_addr > chat-addrs.txt
rm -f all.txt
touch all.txt
cat gentx-addrs.txt >> all.txt
cat chat-addrs.txt >> all.txt
cat gaia-testnet-addrs.txt >> all.txt
sort all.txt | uniq > unique.txt
sed '/^$/d' unique.txt > addrs.txt
rm -f addrs.json-part
touch addrs.json-part
while read addr; do
 echo '{ "address": "'"$addr"'", "coins": [ { "denom": "tree", "amount": "1001000" }, { "denom": "seed", "amount": "1000000000" } ], "sequence_number": "0", "account_number": "0", "original_vesting": null, "delegated_free": [], "delegated_vesting": [], "start_time": "0", "end_time": "0" },' >> addrs.json-part
done <addrs.txt
