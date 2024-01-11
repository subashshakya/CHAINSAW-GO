#!bin/bash

# adding path for fabric tools
pushd ../fabric-samples/test-network/
export PATH=${PWD}/../bin/:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
popd

# bringing up the network with certificate authorities (use when need of cryptography)
pushd ../fabric-samples/test-network/../bin/
cryptogen generate --config=./organizations/cryptogen/crypto-config-org1.yaml --output=organizations org1.example.com
cryptogen generate --config=./organizations/cryptogen/crypto-config-org2.yaml --output=organizations org2.example.com
cryptogen generate --config=./organizations/cryptogen/crypto-config-orderer.yaml --output=organizations
popd