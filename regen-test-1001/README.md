# Regen Ledger Incentivized Testnets

## Overview

*Our incentive program is inspired by [IRISNet](https://github.com/irisnet/testnets)'s program and borrows liberally from it.*

Regen Network will offer rewards to early testnet participants according to the points they have earned. The exact details regarding conversion of points to tokens will be described in another document pending legal approval.

## Points and Instructions

You need to use keybase to generate your own [pgp fingerprint](https://github.com/irisnet/testnets/blob/master/fuxi/How%20to%20use%20keybase.md) first. 

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Participate in Genesis file generation process | Submit your gen-tx.json and use `name-pgp-fingerprint` as validator's name | Submit the url of your PR                                        | 100    |
| 2    | Run a validator node |  | | 100    |
| 3    | Vote on software upgrade proposal              |  |  | 100    |
| 4    | Upgrade your node according the software upgrade proposal |  |  | 400 minus 1 point for every missed block   |
| 5    | Uptime reward |  points awarded for blocks signed since genesis | points are rewarded based your uptime proportional to the life of the life of the testnet  | max 200, awarded proportionally to up time |

