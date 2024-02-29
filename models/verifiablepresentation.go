package models

import "encoding/json"

type IVerifiablePresentation interface {
	Present() IVerifiablePresentation
}


type VerifiablePresentation struct{
	Messages interface{}
	Proof   []byte
}

func (vp *VerifiablePresentation) Json() []byte {
	jsonObj,_ := json.MarshalIndent(vp, "","    ")
	return jsonObj
}
