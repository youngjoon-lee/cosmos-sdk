#!/usr/bin/env bash
set -e

BUILD_CMD=./build/simd
CHAIN_ID=my-chain

rm -rf build
make localnet-start

make build

cat >proposal.json <<EOL
{
    "title": "Staking power reduction param change proposal",
    "description": "Update power reduction",
    "changes": [
        {
            "subspace": "staking",
            "key": "PowerReduction",
            "value": "10000000"
        }
    ],
    "deposit": "10000000stake"
}
EOL

echo
echo "Submitting a proposal from node0..."
echo
$BUILD_CMD tx gov submit-proposal param-change proposal.json --from node0 --home build/node0/simd --keyring-backend test --chain-id $CHAIN_ID --fees="10stake" --yes

sleep 6

echo
echo "Voting phase..."
echo
$BUILD_CMD tx gov vote 1 yes --from node0 --home build/node0/simd --chain-id $CHAIN_ID --keyring-backend test --yes --fees="10stake"
$BUILD_CMD tx gov vote 1 yes --from node1 --home build/node1/simd --chain-id $CHAIN_ID --keyring-backend test --yes --fees="10stake"
$BUILD_CMD tx gov vote 1 yes --from node2 --home build/node2/simd --chain-id $CHAIN_ID --keyring-backend test --yes --fees="10stake"
$BUILD_CMD tx gov vote 1 yes --from node3 --home build/node3/simd --chain-id $CHAIN_ID --keyring-backend test --yes --fees="10stake"

echo
echo "Wait voting_period == 30s..."
echo "In another terminal, you can check the proposal status: $BUILD_CMD q gov proposals"
sleep 35
echo
echo "Make sure the power reduction has been bumped from 1M (default) to 10M (new)!"
echo
$BUILD_CMD q staking params

echo
echo "Taking down simdnode3..."
echo
docker-compose rm -fsv simdnode3

echo "Now check the docker logs: `docker-compose logs -f`. Watch some video on youtube and come back after 10min, there'll be a consensus error at block ~100."