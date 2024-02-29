package models

import (
	"encoding/json"
	"fmt"
)

/*
URI - it represents URI (rfc3986). eg:- http://example.com
*/
type URI interface{}

type TimeString interface{}

type VCFactory interface {
	CreateCredential() IVerifiableCredential
}


type IVerifiableCredential interface {
	CreateCredential() IVerifiableCredential
	GetId() string
}

type VerifiableCredential struct{
	Metadata Metadata
	Claims   Claims
	Proofs   []Proof
}


type Metadata struct{

	Contexts         interface{}
	Id               URI
	Types            []URI
	Issuer           URI
	IssuanceDate     TimeString
	ExpirationDate   TimeString
	CredentialStatus CredentialStatus `json:"credentialStatus"`
}

type Claims interface {}

type Proof struct {
	Type string
	ProofValue []byte
}



type CredentialStatus struct{
	Id     URI         `json:"id"`
	Method interface{} `json:"method"`
	BfIndexes []string `json:"bfIndexes"`
	MTLeafValue string `json:"mtLeaf"`
}

func (cs CredentialStatus) String() string{
	var response string
	response = response + "proof :  \n"
	response = response + "Id : "+ fmt.Sprintf("%v",cs.Id) + "\n"
	response = response + "method: "+ fmt.Sprintf("%v",cs.Method) + "\n"
	response = response + "bloom filter indexes: "+ fmt.Sprintf("%v",cs.BfIndexes) + "\n"
	response = response + "mt leaf hash: "+ fmt.Sprintf("%v",cs.MTLeafValue) + "\n"
	return response
}

func (proof Proof) String() string{
	var response string
	response = response + " \n"
	response = response + "type : "+ fmt.Sprintf("%v",proof.Type) + "\n"
	response = response + "proof value : "+ fmt.Sprintf("%v",proof.ProofValue) + "\n"
	return response
}


func (metadata Metadata) String() string{

	var response string

	response = response + "context : "+ fmt.Sprintf("%v",metadata.Contexts) + "\n"
	response = response + "vc id : "+ fmt.Sprintf("%v", metadata.Id) + "\n"
	response = response + "type : "+ fmt.Sprintf("%v", metadata.Types) + "\n"
	response = response + "entities : "+ fmt.Sprintf("%v",metadata.Issuer) + "\n"
	response = response + "issuance date : "+ fmt.Sprintf("%v", metadata.IssuanceDate) + "\n"
	response = response + "expiration date: "+ fmt.Sprintf("%v", metadata.ExpirationDate) + "\n"
	response = response + "credential status: "+ metadata.CredentialStatus.String()

	return response
}

func (vc VerifiableCredential) GetId() string  {
	return fmt.Sprintf("%v",vc.Metadata.Id)
}

func (vc VerifiableCredential) String() string  {

	var response string

	response = response + fmt.Sprintf("%v", vc.Metadata)+"\n"
	response = response + fmt.Sprintf("%v", vc.Claims)+"\n"
	response = response + fmt.Sprintf("%v", vc.Proofs)+"\n"
	return response
}

func (vc *VerifiableCredential) Json() []byte {
	jsonObj,_ := json.MarshalIndent(vc, "","    ")
	return jsonObj
}




