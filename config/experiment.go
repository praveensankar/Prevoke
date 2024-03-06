package config

import (
	"encoding/json"
	"fmt"
)

type Experiment struct {
	TotalVCs int `json:"totalVCs"`
	RevokedVCs int `json:"revokedVCs"`
	FalsePositiveRate float64 `json:"falsePositiveRate"`
	MtLevelInDLT int `json:"mtLevelInDLT"`
	MtHeight int `json:"mtHeight"`
	RevocationBatchSize int `json:"revocation_batch_size"`

}

func (e Experiment) String() string  {

	var response string
	response = response + fmt.Sprintf("Total VCs: %v", e.TotalVCs)+"\n"
	response = response + fmt.Sprintf("Revoked VCs: %v", e.RevokedVCs)+"\n"
	response = response + fmt.Sprintf("Revocation Batch Size: %v", e.RevocationBatchSize)+"\n"
	response = response + fmt.Sprintf("False Positive Rate: %v", e.FalsePositiveRate)+"\n"
	response = response + fmt.Sprintf("MT Height: %v", e.MtHeight)+"\n"
	response = response + fmt.Sprintf("MT Level in DLT: %v", e.MtLevelInDLT)+"\n"
	return response
}

func (e *Experiment) Json() ([]byte, error){
	return json.MarshalIndent(e, "","  ")
}



func JsonToExperiment(jsonObj []byte) *Experiment {
	exp := Experiment{}
	json.Unmarshal(jsonObj, &exp)
	return &exp
}