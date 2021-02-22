# The Gir Upgrade

## How to vote

Here is the CLI command to vote on the proposal:


    regen tx gov vote <proposalID> <option> --from <your_key_or_acc_address> [flags]


Eg:

``` 
regen tx gov vote 3 yes --from <your_address> --gas auto -y --fees=2000utree --chain-id aplikigo-1
```  


## How to upgrade

## Option 1:- Setting up Cosmovisor (recommended)

 - ### Creating directories for cosmovisor
 ```
 $ mkdir -p ${HOME}/.regen/cosmovisor/genesis/bin
 $ mkdir -p ${HOME}/.regen/cosmovisor/upgrades/Gir/bin
 ```
 - ### Clone and build Cosmovisor
 ```
 $ cd ~
 $ git clone https://github.com/cosmos/cosmos-sdk
 $ cd cosmos-sdk/cosmovisor
 $ git checkout v0.41.0
 $ make cosmovisor
 $ mv cosmovisor $GOBIN
 ```
 
 - ### Build the regen binaries
 ```
 $ cd $GOPATH/src/github.com/regen/regen-ledger
 $ git fetch && git checkout v0.6.0
 $ make build
 $ mv build/regen ${HOME}/.regen/cosmovisor/genesis/bin
 $ git checkout v0.6.1
 $ EXPERIMENTAL=true make build
 ```
 Verify the version of new binary
 ```
 $ ./build/regen version --long
 ```
 It will display the version of regen built:
 ```
 TBA
 ```
 
 Move the newly built binary to it's directory
 ```
 $ mv build/regen ${HOME}/.regen/cosmovisor/upgrades/Gir/bin
 ```
 
 Setup cosmovisor current version link
```
$ ln -s -T ${HOME}/.regen/cosmovisor/genesis ${HOME}/.regen/cosmovisor/current
```

Setup cosmovisor systemd service
```
echo "[Unit]
Description=Regen Node
After=network-online.target
[Service]
User=${USER}
Environment=DAEMON_NAME=regen
Environment=DAEMON_RESTART_AFTER_UPGRADE=true
Environment=DAEMON_HOME=${HOME}/.regen
ExecStart=$(which cosmovisor) start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```

```
$ sudo mv cosmovisor.service /lib/systemd/system/
$ sudo systemctl daemon-reload
$ sudo systemctl stop regen.service && sudo systemctl disable regen.service 
$ sudo systemctl enable cosmovisor.service && sudo systemctl start cosmovisor.service
```

Check logs

```
$ sudo journalctl -u cosmovisor -f
``` 
## Option 2:- Manual 

- Step-1: Stop your regen service after you see this in your logs ` ERR UPGRADE "Gir" NEEDED at height:`
    ```
    sudo service regen stop
    ```

- Update the binary
    ```
    cd $GOPATH/src/github.com/regen/regen-ledger
    git fetch
    git checkout v0.6.1
    EXPERIMENTAL=true make install
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


