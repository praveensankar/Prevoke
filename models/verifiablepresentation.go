package models


type IVerifiablePresentation interface {
	Present() IVerifiablePresentation
}


type VerifiablePresentation struct{
	Messages interface{}
	Proof   []byte
}
