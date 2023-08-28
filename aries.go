package main

import (
	"fmt"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
)


func CreateEmployementProofCredential(id string) (verifiable.Credential){

	vc := verifiable.Credential{
		Context:          nil,
		CustomContext:    nil,
		ID:               id,
		Types:            nil,
		Subject:  verifiable.Subject{
			ID:           "1",
			CustomFields: map[string]interface{}{"employeeID" : 123, "joiningDate": "june 1, 2023"},
		},
		Issuer:           verifiable.Issuer{},
		Issued:           nil,
		Expired:          nil,
		Proofs:           nil,
		Status:           nil,
		Schemas:          nil,
		Evidence:         nil,
		TermsOfUse:       nil,
		RefreshService:   nil,
		JWT:              "",
		SDJWTHashAlg:     "",
		SDJWTDisclosures: nil,
		SDHolderBinding:  "",
		CustomFields:     nil,
	}

	return vc
}

func testAries(){

	employmentVC := CreateEmployementProofCredential("1")
	fmt.Println(employmentVC)
}
