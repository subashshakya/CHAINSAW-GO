#!/bin/bash
set -e

export MSYS_NO_PATHCONV=1
starttime=$(date +%s)

USER_MSP_SOURCE_PATH="../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp"
USER_MSP_DESTINATION_PATH="../rest-api/"
cp ${USER_MSP_SOURCE_PATH}/cacerts/* ${USER_MSP_DESTINATION_PATH}/cacert.pem
cp ${USER_MSP_SOURCE_PATH}/keystore/* ${USER_MSP_DESTINATION_PATH}/privatekey.pem
cp ${USER_MSP_SOURCE_PATH}/signcerts/* ${USER_MSP_DESTINATION_PATH}/signcert.pem

echo COPIED CA CERTIFICATES TO cacert.pem
echo COPIED PRIVATE KEY TO privatekey.pem
echo COPIED SIGNED CERTIFICATES TO signcert.pem