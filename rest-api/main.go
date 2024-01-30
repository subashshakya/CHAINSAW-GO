package main

import (
	"fmt"

	"chainsaw-man-api.example.com/web"
)
func main() {
	cryptoPath := "../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com"
	orgConfig := web.OrgSetup{
		OrgName:      "Org1",
		MSPId:        "Org1MSP",
		CertPath:     cryptoPath + "/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem",
		KeyPath:      cryptoPath + "/users/User1@org1.example.com/msp/keystore/",
		TLSCertPath:  cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.org1.example.com",
	}
	config, err := web.Initialize(orgConfig)
	fmt.Println(config)
	if err != nil {
		fmt.Println("Error while initializing the gRPC connection to the gatewway")
	}
	web.Serve(web.OrgSetup(orgConfig))
}
