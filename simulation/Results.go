package simulation

import "fmt"

type Results struct {
	TotalVCs int `json:"total_issued_vcs"`
	RevokedVCs int `json:"total_revoked_vcs"`
	FalsePositiveRate float64 `json:"false_positive_rate"`
	MtLevelInDLT int `json:"mt_level_in_dlt"`
	MTHeight int `json:"mt_height"`
	NumberOfFalsePositives int `json:"number_of_false_positives"`
	NumberOfVCsRetrievedWitnessFromIssuer int `json:"number_of_vcs_retrieved_witness_from_issuer"`
	NumberOfWitnessUpdatesForMT int `json:"number_of_witness_updates_MT_accumulator"`
	AmountPaid int64 `json:"mt_accumulator_per_update_cost_in_gwei"`
	NumberOfWitnessUpdatesSaved int `json:"number_of_witness_updates_saved"`
	BloomFilterSize int `json:"bloom_filter_size"`
	BloomFilterIndexesPerEntry int  `json:"bloom_filter_indexes_per_entry"`
}


func (r Results) String() string{
	var response string
	response = response + "Total VCs : "+fmt.Sprintf("%d",r.TotalVCs)+ "\n"
	response = response + "Total Revoked VCs : "+fmt.Sprintf("%d",r.RevokedVCs)+ "\n"
	response = response + "False Positive Rate : "+fmt.Sprintf("%f",r.FalsePositiveRate)+ "\n"
	response = response + "Bloom Filter Size : "+fmt.Sprintf("%d",r.BloomFilterSize)+ "\n"
	response = response + "Bloom Filter indexes per entry (no of hash functions) : "+fmt.Sprintf("%d",r.BloomFilterIndexesPerEntry)+ "\n"
	response = response + "Merkle Tree Accumulator height : "+fmt.Sprintf("%d",r.MTHeight)+ "\n"
	response = response + "Merkle Tree Accumulator Level Stored in DLT : "+fmt.Sprintf("%d",r.MtLevelInDLT)+ "\n \n"
	response = response + "Number of False Positives : "+ fmt.Sprintf("%d",r.NumberOfFalsePositives) + "\n"
	response = response + "Number of VCS that ended up updating witnesses from issuer: "+fmt.Sprintf("%d",r.NumberOfVCsRetrievedWitnessFromIssuer)+ "\n"
	response = response + "Amount (in gwei) paid per revocation: "+fmt.Sprintf("%d",r.AmountPaid)+ "\n"
	response = response + "Number of witness that are affected by revocation and require witness update in Merkle Tree Accumulator: "+ fmt.Sprintf("%d",r.NumberOfWitnessUpdatesForMT) + "\n"
	response = response + "Number of witness updates that we saved : "+fmt.Sprintf("%d",r.NumberOfWitnessUpdatesSaved)+ "\n"
	return response
}


