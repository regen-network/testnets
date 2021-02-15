# The Gir Upgrade

## How to vote

Here is the CLI command to vote on the proposal:


    regen tx gov vote <proposalID> <option> --from <your_key_or_acc_address> [flags]


Eg:

``` 
regen tx gov vote 3 yes --from <your_address> --gas auto -y --fees=2000utree --chain-id aplikigo-1
```  


## How to upgrade
 
Upgrade instructions:
- Step-1: Configure halt-height to `138650` in ~/.regen/config/app.toml and restart the binary
    ```
    sudo service regen restart
    ```

The following instructions should only be executed after the upgrade time
- Stop your regen service
    ```
    sudo service regen stop
    ```
- Back up the blockchain data and validator state (Useful when reverting the upgrade if something goes wrong)
    ```
    cp ~/.regen/data <some_back_up_path/data>
    cp ~/.regen/config <some_back_up_path/config>
    ```
- After the `halt-height`, export the state using:
    ```
    regen export --for-zero-height --height 138650` > `060_exported_state.json
    ```
- Verify the genesis
    ```
    jq -S -c -M '' 060_exported_state.json | shasum -a 256

    [TBD]
    ```
- Reset the state
    ```
    regen unsafe-reset-all
    ```
- Update the binary
    ```
    cd regen-ledger
    git fetch
    git checkout v0.6.1
    make install
    ```
- Verify the binary
    ```
    regen version --long 
    ```
    It shoud show following output:
    ```sh
    [TBD]
    ```
- Restart your validator
    ```
    sudo service regen start
    ```
