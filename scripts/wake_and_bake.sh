#!/bin/bash

set -e

ORIGINAL_DIR=$(pwd)
CHAINCODE_NAME="REPORT-HUB"
CHAINCODE_LANGUAGE="go"
CHAINCODE_PATH="../../chaincode-sm/"

echo REMOVING EXISTING WALLETS
rm -rf javascript/wallet/*
rm -rf go/wallet/*
rm -rf java/wallet/*
rm -rf typescript/wallet/*
echo COMPLETE

pushd ../fabric-samples/test-network/

./network.sh down
./network.sh up createChannel -c report-hub-channel

# optional for invoking a peer to invoke a chaincode function
export PATH=${PWD}/../bin:$PATH

echo "<<============ PACKAGING CHAINCODE =============>>"
peer lifecycle chaincode package reports.tar.gz --path ${CHAINCODE_PATH} --lang golang --label report-hub_1.0

echo "<<========== INSTALLING CHAINCODE TO PEERS ==========>>"
# environment variables for Org1
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# installing chaincode to peer 1 of org1
peer lifecycle chaincode install reports.tar.gz

# environment variables for Org2
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

# installing chaincode for peers of org2
peer lifecycle chaincode install reports.tar.gz

# checking for completion of installation in peers
echo "<<=========== INSPECTION OF INSTALLATION ===========>"
OUTPUT=$(peer lifecycle chaincode queryinstalled)
PACKAGE_ID=$(sed -n 's/.*Package ID: \(.*\), Label:.*/\1/p' <<< "$OUTPUT")

echo "CHAINCODE PACKAGE INSTALLED WITH PACKAGE ID:"
echo $PACKAGE_ID

# setting env variables for approving chaincode in org1 peer
export CC_PACKAGE_ID=$PACKAGE_ID
echo "<============================= CORE PEER MSP CONFIG PATH ========================================>"

export FABRIC_CFG_PATH=$PWD/configtx/
echo $FABRIC_CFG_PATH

echo "<==================== CONFIGTX PATH =======================>"

echo ${PWD}/configtx/
configtxgen -configPath ${PWD}/configtx/

echo "<================================ CREATING GENESIS BLOCK ================================>"
configtxgen -outputBlock report-hub-channel.block -profile ChannelUsingRaft -channelID report-hub-channel
configtxgen -profile ChannelUsingRaft -channelID report-hub-channel -outputCreateChannelTx report-hub-channel.tx

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.key

osnadmin channel join --channelID report-hub-channel --config-block ./report-hub-channel.block -o localhost:7053 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"

osnadmin channel list -o localhost:7053 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# setting anchor peer
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
echo $CORE_PEER_MSPCONFIGPATH
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051

echo ${PWD}/../config/orderer.yaml

./network.sh deployCC -c report-hub-channel -ccn ${CHAINCODE_NAME} -ccp ${CHAINCODE_PATH} -ccl ${CHAINCODE_LANGUAGE}

popd

./get_user_msp.sh