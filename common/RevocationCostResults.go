package common

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

type RevocationCostResults struct {
	RevokedVCs int `json:"total_revoked_vcs"`
	RevocationBatchSize int `json:"revocation_batch_size"`
	MtLevelInDLT int `json:"mt_level_in_dlt"`
	FalsePositiveRate float64 `json:"false_positive_rate"`
	RevocationCost int64 `json:"revocation_cost_in_wei"`
	RevocationCostRawData []int64 `json:"revocation_cost_raw_data"`
}

func CreateRevocationCostResults() *RevocationCostResults {
	result := &RevocationCostResults{}
	return result
}




func (r *RevocationCostResults) AddRevocationCostPerBatch(revocationCost int64){
	r.RevocationCostRawData = append(r.RevocationCostRawData, revocationCost)
	if r.RevocationCost==0{
		r.RevocationCost = r.RevocationCost + revocationCost
	} else {
		r.RevocationCost = r.RevocationCost + revocationCost
		r.RevocationCost = r.RevocationCost / 2
	}
}


func (r RevocationCostResults) String() string{
	var response string
	response = response + "Total Revoked VCs : "+fmt.Sprintf("%d",r.RevokedVCs)+ "\n"
	response = response + "revocation Batch size : "+fmt.Sprintf("%d",r.RevocationBatchSize)+ "\n"
	response = response + "False Positive Rate : "+fmt.Sprintf("%f",r.FalsePositiveRate)+ "\n"
	response = response + "Merkle Tree Accumulator Level Stored in DLT : "+fmt.Sprintf("%d",r.MtLevelInDLT)+ "\n \n"
	response = response + "revocation cost (in unit of gas) per revocation: "+fmt.Sprintf("%d",r.RevocationCost)+ "\n"


	return response
}


func  WriteRevocationCostResultsToFile(filename string, result []RevocationCostResults) {

	var results []RevocationCostResults


	jsonFile, err := os.Open(filename)
	if err != nil {
		jsonFile2, err2 := os.Create(filename)
		if err2 != nil {
			zap.S().Errorln("ERROR - results.json file creation error")
		}
		resJson, _ := ioutil.ReadAll(jsonFile2)
		json.Unmarshal(resJson, &results)
		results = append(results, result...)
	} else{
		resJson, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(resJson, &results)
		results = append(results, result...)
	}

	jsonRes, err3 := json.MarshalIndent(results,"","")
	if err3 != nil {
		zap.S().Errorln("ERROR - marshalling the results")
	}

	//filename := fmt.Sprintf("Simulation/results/result_%v_%v_%v.json",numberOfVcs, numberOfRevokedVcs, mtLevelInDLT)
	err = ioutil.WriteFile(filename, jsonRes, 0644)
	if err != nil {
		zap.S().Errorln("unable to write results to file")
	}
	//zap.S().Errorln("RESULTS - successfully written to the file")

}



func (r *RevocationCostResults) Json() ([]byte, error){
	//return json.MarshalIndent(r, "","    ")
	return json.Marshal(r)
}



func JsonToRevocationCostResults(jsonObj []byte) *RevocationCostResults {
	res := RevocationCostResults{}
	json.Unmarshal(jsonObj, &res)
	return &res
}
