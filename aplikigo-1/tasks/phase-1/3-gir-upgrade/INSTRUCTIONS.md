# The Gir Upgrade

## How to vote

Here is the CLI command to vote on the proposal
    ```sh
    regen tx gov vote <proposalID> <option> --from <your_key_or_acc_address>
    ```
    ex:

    ```sh
    regen tx gov vote 1 YES --from regen:1gpgpkmz3swt5ucdjxvrszk6yvjxjylsh4eur9d
    ```

## How to upgrade
 
- Upgrade instructions:
    - Step-1: Configure halt-height to `138650` in ~/.regen/config/app.toml
    - After the halt-height, export the state using:
    `regen export --for-zero-height --height 138650` > `060_exported_state.json`