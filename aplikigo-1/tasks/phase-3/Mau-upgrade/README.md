# The Mau Upgrade

 


## How to upgrade

## Option 1:- Cosmovisor (recommended)

 - ### Creating directory for Mau upgrade
 ```
 $ mkdir -p ${HOME}/.regen/cosmovisor/upgrades/Mau/bin
 ```

 
 - ### Build the regen binaries
 ```
 $ cd $GOPATH/src/github.com/regen/regen-ledger
 $ git fetch && git checkout v0.6.2
 $ EXPERIMENTAL=true make build
 ```
 Verify the version of new binary
 ```
 $ ./build/regen version --long
 ```
 It will display the version of regen built:
 ```
 name: regen
server_name: regen
version: v0.6.2
commit: d691f09fce43cdd7bb6067ab12b38cdd875d867c
build_tags: netgo,experimental,ledger
go: go version go1.15.5 linux/amd64
build_deps:
- github.com/99designs/keyring@v1.1.6
- github.com/ChainSafe/go-schnorrkel@v0.0.0-20200405005733-88cbf1b4c40d
- github.com/CosmWasm/wasmd@v0.14.0
- github.com/CosmWasm/wasmvm@v0.12.0
.....
 ```
 
 Move the newly built binary to it's directory
 ```
 $ mv build/regen ${HOME}/.regen/cosmovisor/upgrades/Mau/bin
 ```
 
 
## Option 2:- Manual 

- Step-1: Stop your regen service after you see this in your logs ` ERR UPGRADE "Mau" NEEDED at height:`
    ```
    sudo service regen stop
    ```

- Update the binary
    ```
    cd $GOPATH/src/github.com/regen/regen-ledger
    git fetch
    git checkout v0.6.2
    EXPERIMENTAL=true make install
    ```
- Verify the binary
    ```
    regen version --long 
    ```
    It shoud show following output:
    ```sh
    name: regen
    server_name: regen
    version: v0.6.2
    commit: d691f09fce43cdd7bb6067ab12b38cdd875d867c
    build_tags: netgo,experimental,ledger
    go: go version go1.15.5 linux/amd64
    build_deps:
    - github.com/99designs/keyring@v1.1.6
    - github.com/ChainSafe/go-schnorrkel@v0.0.0-20200405005733-88cbf1b4c40d
    - github.com/CosmWasm/wasmd@v0.14.0
    - github.com/CosmWasm/wasmvm@v0.12.0
    ........
    ```
- Restart your validator
    ```
    sudo service regen start
    ```



