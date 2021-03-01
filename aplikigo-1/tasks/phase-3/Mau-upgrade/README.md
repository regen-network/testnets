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
 TBA
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
    [TBD]
    ```
- Restart your validator
    ```
    sudo service regen start
    ```



