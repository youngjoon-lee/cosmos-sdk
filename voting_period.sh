#!/usr/bin/env bash

# gsed is like sed, but for macos
# - if you're on linux: replace gsed with sed
# - if you're on macos: install gsed, of figure out macos's sed directly
gsed -i 's/"voting_period": "172800s"/"voting_period": "60s"/' build/node0/simd/config/genesis.json
gsed -i 's/"voting_period": "172800s"/"voting_period": "60s"/' build/node1/simd/config/genesis.json
gsed -i 's/"voting_period": "172800s"/"voting_period": "60s"/' build/node2/simd/config/genesis.json
gsed -i 's/"voting_period": "172800s"/"voting_period": "60s"/' build/node3/simd/config/genesis.json
