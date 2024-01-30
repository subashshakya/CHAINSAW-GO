package web

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type OrgSetup struct {
	OrgName string
	MSPId string
	CryptoPath string
	CertPath string
	KeyPath string
	TLSCertPath string
	PeerEndpoint string
	GatewayPeer string
	Gateway client.Gateway
}

func Serve(setup OrgSetup) {
	http.HandleFunc("/", setup.Invoke)
	http.HandleFunc("/read-report", setup.ReadReportHandler)
	fmt.Println("Listening on http://127.0.0.1:3000")
	err := http.ListenAndServe(":3000", nil);
	if err != nil {
		panic(err);
	}
}
