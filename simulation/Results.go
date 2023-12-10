package simulation

import "fmt"

type Results struct {
	TotalVCs int `json:"TotalVCs"`
	RevokedVCs int `json:"TotalRevokedVCs"`
	FalsePositiveRate float64 `json:"FalsePositiveRate"`
	MtLevelInDLT int `json:"MtLevelInDLT"`
	MtDepth int `json:"MtDepth"`
	NumberOfFalsePositives int `json:"NumberOfResultedFalsePositives"`
	NumberOfAffectedVCs int `json:"NumberOfAffectedVCs"`
	NumberOfWitnessUpdatesSaved int `json:"NumberOfWitnessUpdatesSaved"`
}


func (r Results) String() string{
	var response string
	response = response + "Total VCs : "+fmt.Sprintf("%d",r.TotalVCs)+ "\t"
	response = response + "Total Revoked VCs : "+fmt.Sprintf("%d",r.RevokedVCs)+ "\t"
	response = response + "FalsePositiveRate : "+fmt.Sprintf("%f",r.FalsePositiveRate)+ "\t"
	response = response + "Merkle Tree Accumulator depth : "+fmt.Sprintf("%d",r.MtDepth)+ "\t"
	response = response + "Merkle Tree Accumulator Level Stored in DLT : "+fmt.Sprintf("%d",r.MtLevelInDLT)+ "\n"
	response = response + "Number of Affected VCs : "+ fmt.Sprintf("%d",r.NumberOfAffectedVCs) + "\t"
	response = response + "Number of VCS that ended up updating witnesses : "+fmt.Sprintf("%d",r.NumberOfFalsePositives)+ "\t"
	response = response + "Number of witness updates that we saved : "+fmt.Sprintf("%d",r.NumberOfWitnessUpdatesSaved)+ "\t"
	return response
}


