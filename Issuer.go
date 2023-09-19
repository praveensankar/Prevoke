package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
)

type IIsser interface {
	bootstrap()
	issue(credential verifiable.Credential)
	revoke(credential verifiable.Credential)
}


type Issuer struct{
	name string
	credentialStore []verifiable.Credential
	vcCounter int
	blockchainEndPoint *ethclient.Client
	RevocationServiceAddress common.Address
}

func (issuer *Issuer) bootstrap(config Config) {

	// register public keys at the revocation service
	// ideally, this step should be performed before the starting of the issuance process

	// connect to the blockchain network
	var err error
	issuer.blockchainEndPoint, err = ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	issuer.RevocationServiceAddress = common.HexToAddress(config.SmartContractAddress)



}


func (issuer *Issuer) issue() {


	// step 1 - issuer generates new VC
	vc := CreateEmployementProofCredential(string(issuer.vcCounter))

	// when issuer issue new credentials, the credential is created
	issuer.credentialStore = append(issuer.credentialStore, vc)

}

func (issuer *Issuer) revoke(credential verifiable.Credential) {

	// call revocation service
}


func testIssuer(config Config){

	var issuer Issuer = Issuer{
		name:            "employer 1",
		credentialStore: []verifiable.Credential{},
		vcCounter:       0,
	}

	issuer.bootstrap(config)
	issuer.issue()
	for index := range issuer.credentialStore{
		fmt.Println(issuer.credentialStore[index])
	}
}

