package web

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var contract *client.Contract

func Initialize (setup OrgSetup) (*OrgSetup, error) {
	clientConnection := setup.newGrpcConnection()
	id := setup.newIdentity()
	sign := setup.newSign()

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	panicOnError(err)
	setup.Gateway = *gateway
	// defer gateway.Close()
	network := gateway.GetNetwork("report-hub-channel")
	fmt.Print(network)

	contract := network.GetContract("REPORT-HUB")
	fmt.Print(contract)
	log.Println("INITIALIZATION COMPLETED")
	return &setup, nil
}

func (setup OrgSetup) newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(setup.TLSCertPath)
	panicOnError(err)

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportLayerCredentials := credentials.NewClientTLSFromCert(certPool, setup.GatewayPeer)

	connection, err := grpc.Dial(setup.PeerEndpoint, grpc.WithTransportCredentials(transportLayerCredentials))
	panicOnError(err)
	return connection
}

func (setup OrgSetup) newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(setup.CertPath)
	panicOnError(err)

	id, err := identity.NewX509Identity(setup.MSPId, certificate)
	panicOnError(err)
	return id
}

func (setup OrgSetup) newSign() identity.Sign {
	files, err := ioutil.ReadDir(setup.KeyPath)
	panicOnError(err)
	privateKeyPEM, err := ioutil.ReadFile(path.Join(setup.KeyPath, files[0].Name()))
	panicOnError(err)

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	panicOnError(err);

	sign, err := identity.NewPrivateKeySign(privateKey)
	panicOnError(err)

	return sign
}

func loadCertificate(filePath string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filePath)
	panicOnError(err)
	return identity.CertificateFromPEM(certificatePEM)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
