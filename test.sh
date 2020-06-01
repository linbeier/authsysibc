#!/bin/bash
rm -rf ~/.appd ~/.appcli
make install
appd init testnode --chain-id testchain
appcli config chain-id testchain
appcli config output json
appcli config indent true
appcli config trust-node true
appcli config keyring-backend test
appcli keys add alice
appcli keys add bob
appcli keys add clare
appcli keys add dog
appcli keys add egg
appcli keys add f
appcli keys add g
appcli keys add h
appcli keys add i
appcli keys add j
appcli keys add k
appd add-genesis-account $(appcli keys show alice -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show bob -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show clare -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show dog -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show egg -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show f -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show g -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show h -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show i -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show j -a) 100000token,1000000000stake
appd add-genesis-account $(appcli keys show k -a) 100000token,1000000000stake
appd gentx --name alice --keyring-backend test
appd collect-gentxs
appd validate-genesis
appd start

