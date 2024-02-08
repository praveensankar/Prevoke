package vc

import (
	"fmt"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
)


func CreateEmployementProofCredential(id string) (*verifiable.Credential){

	vc := verifiable.Credential{
		Context:          nil,
		CustomContext:    nil,
		ID:               id,
		Types:            nil,
		Subject:  verifiable.Subject{
			ID:           id,
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


	return &vc

}

//func AddBBSSignature(privatekey *keyset.Handle, vc *verifiable.Credential){
//	issued := time.Now()
//	signerHandle, _ := bbs.NewSigner(privatekey)
//
//	signature, err := signerHandle.Sign(messages)
//	if err != nil{
//		zap.S().Infoln("BBS - error signing: ",err)
//	}
//	vc.AddLinkedDataProof(&verifiable.LinkedDataProofContext{
//		Created:                 &issued,
//		SignatureType:           "BbsBlsSignature2020",
//		Suite:                  bbsblssignature2020.New(suite.WithSigner(signerHandle)),
//		SignatureRepresentation: verifiable.SignatureJWS,
//		VerificationMethod:      "did:example:123456#key1",
//	}, jsonld.WithDocumentLoader(getJSONLDDocumentLoader()))
//}

func testAries(){

	employmentVC := CreateEmployementProofCredential("1")
	fmt.Println(employmentVC)
}
