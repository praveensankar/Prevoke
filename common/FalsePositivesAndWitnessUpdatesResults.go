package common

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

type FalsePositiveAndWitnessUpdateResults struct {
	TotalVCs int `json:"total_issued_vcs"`
	RevokedVCs int `json:"total_revoked_vcs"`
	RevocationBatchSize int `json:"revocation_batch_size"`
	RevocationMode string `json:"revocation_mode"`
	MtLevelInDLT int `json:"mt_level_in_dlt"`
	MTHeight int `json:"mt_height"`
	FalsePositiveRate float64 `json:"false_positive_rate"`
	NumberOfFalsePositives int `json:"number_of_false_positives"`
	NumberOfVCsRetrievedWitnessFromDLT int `json:"number_of_vcs_retrieved_witness_from_dlt"`
	AffectedVCIDs []string `json:"affected_vcIDs"`
	FalsePositiveResults []string `json:"false_positive_vcIDs"`
	FetchedWitnessesFromDLT []string `json:"vcIDs_fetched_witnesses_from_DLT"`
}

func CreateFalsePositiveAndWitnessUpdateResults() *FalsePositiveAndWitnessUpdateResults {
	result := &FalsePositiveAndWitnessUpdateResults{}
	return result
}




func (r FalsePositiveAndWitnessUpdateResults) String() string{
	var response string
	response = response + "Total VCs : "+fmt.Sprintf("%d",r.TotalVCs)+ "\n"
	response = response + "Total Revoked VCs : "+fmt.Sprintf("%d",r.RevokedVCs)+ "\n"
	response = response + "revocation Batch size : "+fmt.Sprintf("%d",r.RevocationBatchSize)+ "\n"
	response = response + "revocation mode : "+fmt.Sprintf("%s",r.RevocationMode)+ "\n"
	response = response + "False Positive Rate : "+fmt.Sprintf("%f",r.FalsePositiveRate)+ "\n"
	response = response + "Merkle Tree Accumulator height : "+fmt.Sprintf("%d",r.MTHeight)+ "\n"
	response = response + "Merkle Tree Accumulator Level Stored in DLT : "+fmt.Sprintf("%d",r.MtLevelInDLT)+ "\n \n"

	response = response + "Number of False Positives : "+ fmt.Sprintf("%d",r.NumberOfFalsePositives) + "\n"
	response = response + "Number of VCS that updated witnesses from smart contract: "+fmt.Sprintf("%d",r.NumberOfVCsRetrievedWitnessFromDLT)+ "\n"

	return response
}


func  WriteFalsePositiveAndWitnessUpdateResultsToFile(filename string, result []FalsePositiveAndWitnessUpdateResults) {

	var results []FalsePositiveAndWitnessUpdateResults


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


func  WriteFalsePositiveAndWitnessUpdateRawResultsToFile(filename string, result []FalsePositiveAndWitnessUpdateResults) {

	var results []FalsePositiveAndWitnessUpdateResults


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

	jsonRes, err3 := json.Marshal(results)
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


func (r *FalsePositiveAndWitnessUpdateResults) Json() ([]byte, error){
	//return json.MarshalIndent(r, "","    ")
	return json.Marshal(r)
}



func JsonToFalsePositiveAndWitnessUpdateResultsResults(jsonObj []byte) *FalsePositiveAndWitnessUpdateResults {
	res := FalsePositiveAndWitnessUpdateResults{}
	json.Unmarshal(jsonObj, &res)
	return &res
}
