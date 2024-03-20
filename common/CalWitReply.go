package common

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type CalWitnessReply struct {
	numberOfFalsePositives string
	NumberOfVCsRetrievingVCsFromDLT string
	Type RequestType
}

func (r *CalWitnessReply) SetFalsePositives(falsePositives string){
	r.numberOfFalsePositives = falsePositives
}

func (r *CalWitnessReply) SetNumberOfVCsRetrievingVCsFromDLT(numberOfVCsRetrievingVCsFromDLT string){
	r.NumberOfVCsRetrievingVCsFromDLT = numberOfVCsRetrievingVCsFromDLT
}

func (r *CalWitnessReply) GetFalsePositives() string {
	return r.numberOfFalsePositives
}

func (r *CalWitnessReply) GetNumberOfVCsRetrievingVCsFromDLT() string {
	return r.NumberOfVCsRetrievingVCsFromDLT
}

func (r *CalWitnessReply) Json() ([]byte, error){
	return json.Marshal(r)
}



func JsonToCalWitnessReply(jsonObj []byte) *CalWitnessReply {
	reply := CalWitnessReply{}
	err := json.Unmarshal(jsonObj, &reply)
	if err!=nil{
		zap.S().Infoln("REQUEST - error unmarshalling witness calculation reply: ", err)
	}
	return &reply
}


func (r CalWitnessReply) String() string{
	var response string
	response = response + "False Positives: "+fmt.Sprintf("%d",r.GetFalsePositives())+ "\t"
	response = response + "Number of VCs retrieve witnesses from DLT: "+fmt.Sprintf("%d",r.GetNumberOfVCsRetrievingVCsFromDLT())+ "\t"
	return response
}


func NewCalWitnessReply() CalWitnessReply {
	r := CalWitnessReply{}
	return r
}
