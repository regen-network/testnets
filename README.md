# Regen Ledger Testnets

Testnets for [Regen Ledger](https://github.com/regen-network/regen-ledger)

## Join `regen-test-1000` Public Testnet

We are currently preparing the `regen-test-1000` public testnet launch.

We are intending to airdrop testnet tokens to all current Cosmos Hub validators
and delegators (as of early next week). You are also welcome to submit a genesis
account PR to receive tokens.

If you would like to be in the initial validator set, you are welcome to submit
a `gentx` PR.

Here is the planned timing for the testnet rollout:
- 2019-06-04 21:00UTC (2pm PDT/5pm EDT/11pm CEST) - Finalize `genesis.json` and  [regen-ledger](https://github.com/regen-network/regen-ledger) version
- 2019-06-05 18:00UTC (11am PDT/2pm EDT/8pm CEST) - Genesis time

## Testnet Status

### `xrn-test-3`

Testnet `xrn-test-3` started producing blocks at `2019-03-29T19:44:44.571815638Z` and is live as of this writing.

In this testnet, validator nodes currently have ports 26656, 26657 and 1317 open for testing purposes. In the future,
the testnet will be setup with more security hardening via sentry and seed nodes.

The validator node URL's are as follows:

* [xrn-us-east-1.regen.network](http://xrn-us-east-1.regen.network:26657)
* [xrn-us-west-1.regen.network](http://xrn-us-west-1.regen.network:26657)
* [xrn-eu-central-1.regen.network](http://xrn-eu-central-1.regen.network:26657)

`xrncli` can be configured to connect to the testnet as follows:

```sh
xrncli init --chain-id xrn-test-2 --node tcp://xrn-us-east-1.regen.network:26657
```

### `xrn-test-2`

Deployed at `2018-12-19T20:40:06.463846Z`.

### `xrn-1`

The initial Regen Ledger testnet `xrn-1` was deployed on 2018-12-19.

## Running a full node in the cloud

Node configurations for [NixOS](https://nixos.org) are provided in this repository.

[../module.nix](../module.nix) contains a NixOS module for running a node.

The steps for setting this up on a running NixOS machine are roughly as follows:
1. Clone the git repository into a local folder on the machine
2. Import `module.nix` in `/etc/nixos/configuration`:
```nix
  imports =
    [
      ./hardware-configuration.nix
      /path-to-my-git-repo/module.nix
    ];
```
3. Enable the `xrn` programs in `/etc/nixos/configuration` and run `nixos-rebuild switch):
```nix
  programs.xrn.enable = true;
```
4. Run `xrncli init --home /var/xrnd`
5. Configure node configuration in `/var/xrnd/config` (`genesis.json`, `config.toml`, etc.)
6. Enable the xrnd service in `/etc/nixos/configuration` and run `nixos-rebuild switch`
```nix
  services.xrnd.enable = true;
  services.xrnd.moniker = "my-node-moniker";
```