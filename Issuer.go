package main

import (
	"fmt"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
)

type IIsser interface {
	issue(credential verifiable.Credential)
	revoke(credential verifiable.Credential)
}


type Issuer struct{
	name string
	credentialStore []verifiable.Credential
	vcCounter int
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


func testIssuer(){

	var issuer Issuer = Issuer{
		name:            "employer 1",
		credentialStore: []verifiable.Credential{},
		vcCounter:       0,
	}

	issuer.issue()
	for index := range issuer.credentialStore{
		fmt.Println(issuer.credentialStore[index])
	}
}

