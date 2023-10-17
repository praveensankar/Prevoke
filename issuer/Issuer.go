package issuer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"github.com/praveensankar/Revocation-Service/contracts"
)

type IIsser interface {
	generateDummyVC() verifiable.Credential
	issue(config config.Config, credential verifiable.Credential)
	revoke(config config.Config, credential verifiable.Credential)
	setRevocationService(rs IRevocationService)
}


type Issuer struct{
	name string
	credentialStore []verifiable.Credential
	vcCounter int
	blockchainEndPoint *ethclient.Client
	RevocationService IRevocationService
}

/*
creates new issuer instance.
sets up revocation service and blockchain endpoint
 */
func  CreateIssuer(config config.Config) *Issuer{

	issuer := Issuer{}
	// register public keys at the revocation service
	// ideally, this step should be performed before the starting of the issuance process

	// connect to the blockchain network
	var err error
	issuer.blockchainEndPoint, err = ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	issuer.name = config.IssuerName
	issuer.credentialStore= []verifiable.Credential{}
	issuer.vcCounter = 0

	rs := CreateRevocationService(config)
	issuer.setRevocationService(rs)

	return &issuer
}


func (issuer *Issuer) setRevocationService(rs IRevocationService) {
	issuer.RevocationService = rs
}

func (issuer *Issuer) generateDummyVC() verifiable.Credential {
	// step 1 - issuer generates new VC
	vc := vc.CreateEmployementProofCredential(string(issuer.vcCounter))
	return vc
}

func (issuer *Issuer) issue(config config.Config, vc verifiable.Credential) {
	// when issuer issue new credentials, the credential is created
	issuer.credentialStore = append(issuer.credentialStore, vc)


	tx, err := issuer.RevocationService.IssueVC(vc)
	if err != nil {
		return
	}

}

func (issuer *Issuer) revoke(config config.Config, credential verifiable.Credential) {

	// call revocation service
}


func testIssuer(config config.Config){

	issuer := CreateIssuer(config)
	vc1 := issuer.generateDummyVC()
	issuer.issue(vc1)
	for index := range issuer.credentialStore{
		fmt.Println(issuer.credentialStore[index])
	}
}

