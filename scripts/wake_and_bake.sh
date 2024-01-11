#!/bin/bash

set -e

ORIGINAL_DIR=$(pwd)
CHAINCODE_NAME="REPORT-HUB"
CHAINCODE_LANGUAGE="go"
CHAINCODE_PATH="../chaincode-sm/"

echo REMOVING EXISTING WALLETS
rm -rf javascript/wallet/*
rm -rf go/wallet/*
rm -rf java/wallet/*
rm -rf typescript/wallet/*
echo COMPLETE

pushd ../fabric-samples/test-netork

./network.sh down
./network.sh up

# optional for invoking a peer to invoke a chaincode function
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

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
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --name ${CHAINCODE_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

peer lifecycle chaincode checkcommitreadiness --name ${CHAINCODE_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json

peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --name ${CHAINCODE_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"

# for deploying a channel with specified parameters
./network.sh deployCC -ccn ${CHAINCODE_NAME} -ccp ${CHAINCODE_PATH} -ccl ${CHAINCODE_LANGUAGE}

./get_user_msp.sh