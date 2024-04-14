package common

import (
	"encoding/json"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	mapset "github.com/deckarep/golang-set"
	"github.com/Revocation-Service/config"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"os"
	"time"
)

type Results struct {
	TotalVCs int `json:"total_issued_vcs"`
	RevokedVCs int `json:"total_revoked_vcs"`
	RevocationBatchSize int `json:"revocation_batch_size"`
	MtLevelInDLT int `json:"mt_level_in_dlt"`
	MTHeight int `json:"mt_height"`
	FalsePositiveRate float64 `json:"false_positive_rate"`
	BloomFilterSize int `json:"bloom_filter_size"`
	BloomFilterIndexesPerEntry int  `json:"bloom_filter_indexes_per_entry"`
	MerkleTreeSizeTotal int `json:"merkle_tree_size_total"`
	MerkleTreeNodesCountTotal int `json:"merkle_tree_nodes_count_total"`
	MerkleTreeSizeInDLT int `json:"merkle_tree_size_in_dlt"`
	MerkleTreeNodesCountInDLT int `json:"merkle_tree_nodes_count_in_dlt"`
	NumberOfFalsePositives int `json:"number_of_false_positives"`
	NumberOfVCsRetrievedWitnessFromDLT int `json:"number_of_vcs_retrieved_witness_from_dlt"`
	NumberOfVCsRetrievedWitnessFromIssuer int `json:"number_of_vcs_retrieved_witness_from_issuer"`
	RevocationTimeperBatch float64 `json:"revocation_timeper_vc"`
	RevocationTimeTotal float64 `json:"revocation_time_total"`
	RevocationTimeRawData []float64 `json:"revocation_time_raw_data"`
	VerificationTimePerValidVC float64 `json:"verification_time_per_valid_vc"`
	VerificationTimePerValidVCRawData []float64 `json:"verification_time_per_valid_vc_raw_data"`
	VerificationTimeTotalValidVCs float64 `json:"verification_time_total_valid_vcs"`
	VerificationTimePerRevokedorFalsePositiveVC float64 `json:"verification_time_per_false_positive_or_revoked_vc"`
	VerificationTimePerRevokedorFalsePositiveVCRawData []float64 `json:"verification_time_per_revokedor_false_positive_vc_raw_data"`
	AvgTimeToFetchWitnessFromIssuer float64 `json:"avg_time_to_fetch_witness_from_issuer"`
	AvgTimeToFetchWitnessFromIssuerRawData []float64 `json:"avg_time_to_fetch_witness_from_issuer_raw_data"`
	AvgTimeToFetchWitnessFromSmartContract float64 `json:"avg_time_to_fetch_witness_from_smart_contract"`
	AvgTimeToFetchWitnessFromSmartContractRawData []float64 `json:"avg_time_to_fetch_witness_from_smart_contract_raw_data"`
	AvgTimeToComputeCorrectWitnessAtHolder float64 `json:"avg_time_to_compute_correct_witness_at_holder"`
	AvgTimeToComputeCorrectWitnessAtHolderRawData []float64 `json:"avg_time_to_compute_correct_witness_at_holder_raw_data"`
	VerificationTimeTotalRevokedorFalsePositiveVCs float64 `json:"verification_time_total_false_positive_and_revoked_vcs"`
	VerificationTimeTotal float64 `json:"verification_time_total"`
	BBSProoGenerationTimePerVP float64 `json:"bbs_proof_generation_time"`
	BBSVerificationTimePerVP float64`json:"bbs_verification_time"`
	SimulationTime float64 `json:"simulation_time"`
	ContractDeploymentCost int64 `json:"contract_deployment_cost"`
	BulkIssuanceCost int64 `json:"bulk_issuance_cost"`
	AmountPaid int64 `json:"revocation_cost_in_wei"`
	RevocationCostRawData []int64 `json:"revocation_cost_raw_data"`
	AffectedIndexes mapset.Set `json:"affectedIndexes"`
	AffectedVCIDs []string `json:"affected_vcIDs"`
	FalsePositiveResults mapset.Set `json:"false_positive_vcIDs"`
	FetchedWitnessesFromDLT mapset.Set `json:"vcIDs_fetched_witnesses_from_DLT"`
}

func CreateResult() *Results {
	result := &Results{}
	result.AffectedIndexes = mapset.NewSet()
	result.FalsePositiveResults = mapset.NewSet()
	result.FetchedWitnessesFromDLT = mapset.NewSet()
	return result
}

func (r *Results) AddRevocationTimeTotal(revocationTime float64){
	r.RevocationTimeTotal = r.RevocationTimeTotal+revocationTime
}

func (r *Results) AddRevocationTimePerBatch(revocationTime float64){
	r.RevocationTimeRawData = append(r.RevocationTimeRawData, revocationTime)
	if r.RevocationTimeperBatch==0.0{
		r.RevocationTimeperBatch = r.RevocationTimeperBatch + revocationTime
	} else {
		r.RevocationTimeperBatch = r.RevocationTimeperBatch + revocationTime
		r.RevocationTimeperBatch = r.RevocationTimeperBatch / 2
	}
}

func (r *Results) AddVerificationTimeTotal(vTime float64){
	r.VerificationTimeTotal = r.VerificationTimeTotal+vTime
}
func (r *Results) AddVerificationTimeTotalValidVCs(phase1Time float64){
	r.VerificationTimeTotalValidVCs = r.VerificationTimeTotalValidVCs+phase1Time
}

func (r *Results) AddVerificationTimePerValidVC(phase1Time float64){
	r.VerificationTimePerValidVCRawData = append(r.VerificationTimePerValidVCRawData, phase1Time)
	if r.VerificationTimePerValidVC==0.0{
		r.VerificationTimePerValidVC = r.VerificationTimePerValidVC + phase1Time
	} else {
		r.VerificationTimePerValidVC = r.VerificationTimePerValidVC + phase1Time
		r.VerificationTimePerValidVC = r.VerificationTimePerValidVC / 2
	}
}

func (r *Results) AddVerificationTimeTotalRevokedandFalsePositiveVCs(phase2Time float64){
	r.VerificationTimeTotalRevokedorFalsePositiveVCs = r.VerificationTimeTotalRevokedorFalsePositiveVCs+phase2Time
}

func (r *Results) AddVerificationTimePerRevokedandFalsePositiveVC(phase2Time float64){
	r.VerificationTimePerRevokedorFalsePositiveVCRawData = append(r.VerificationTimePerRevokedorFalsePositiveVCRawData, phase2Time)
	if r.VerificationTimePerRevokedorFalsePositiveVC==0.0{
		r.VerificationTimePerRevokedorFalsePositiveVC = r.VerificationTimePerRevokedorFalsePositiveVC + phase2Time
	} else {
		r.VerificationTimePerRevokedorFalsePositiveVC = r.VerificationTimePerRevokedorFalsePositiveVC + phase2Time
		r.VerificationTimePerRevokedorFalsePositiveVC = r.VerificationTimePerRevokedorFalsePositiveVC / 2
	}
}

func (r *Results) AddAvgTimeToFetchWitnessFromIssuer(timeToFetch float64) {
	r.AvgTimeToFetchWitnessFromIssuerRawData = append(r.AvgTimeToFetchWitnessFromIssuerRawData, timeToFetch)
	if r.AvgTimeToFetchWitnessFromIssuer==0.0{
		r.AvgTimeToFetchWitnessFromIssuer = r.AvgTimeToFetchWitnessFromIssuer + timeToFetch
	} else {
		r.AvgTimeToFetchWitnessFromIssuer = r.AvgTimeToFetchWitnessFromIssuer + timeToFetch
		r.AvgTimeToFetchWitnessFromIssuer = r.AvgTimeToFetchWitnessFromIssuer / 2
	}
}

func (r *Results) AddAvgTimeToFetchWitnessFromSmartContract(timeToFetch float64) {
	r.AvgTimeToFetchWitnessFromSmartContractRawData = append(r.AvgTimeToFetchWitnessFromSmartContractRawData, timeToFetch)
	if r.AvgTimeToFetchWitnessFromSmartContract==0.0{
		r.AvgTimeToFetchWitnessFromSmartContract = r.AvgTimeToFetchWitnessFromSmartContract + timeToFetch
	} else {
		r.AvgTimeToFetchWitnessFromSmartContract = r.AvgTimeToFetchWitnessFromSmartContract + timeToFetch
		r.AvgTimeToFetchWitnessFromSmartContract = r.AvgTimeToFetchWitnessFromSmartContract / 2
	}
}

func (r *Results) AddAvgTimeToComputeCorrectWitnessAtHolder(timeToCompute float64) {
	r.AvgTimeToComputeCorrectWitnessAtHolderRawData = append(r.AvgTimeToComputeCorrectWitnessAtHolderRawData, timeToCompute)
	if r.AvgTimeToComputeCorrectWitnessAtHolder==0.0{
		r.AvgTimeToComputeCorrectWitnessAtHolder = r.AvgTimeToComputeCorrectWitnessAtHolder + timeToCompute
	} else {
		r.AvgTimeToComputeCorrectWitnessAtHolder = r.AvgTimeToComputeCorrectWitnessAtHolder + timeToCompute
		r.AvgTimeToComputeCorrectWitnessAtHolder = r.AvgTimeToComputeCorrectWitnessAtHolder / 2
	}
}

func (r *Results) AddBBSVerificationTimePerVP(bbsTime float64){
	if r.BBSVerificationTimePerVP==0.0{
		r.BBSVerificationTimePerVP = r.BBSVerificationTimePerVP + bbsTime
	} else {
		r.BBSVerificationTimePerVP = r.BBSVerificationTimePerVP + bbsTime
		r.BBSVerificationTimePerVP = r.BBSVerificationTimePerVP / 2
	}
}

func (r *Results) AddBBSProofGenerationTimePerVP(bbsTime float64){
	if r.BBSProoGenerationTimePerVP==0.0{
		r.BBSProoGenerationTimePerVP = r.BBSProoGenerationTimePerVP + bbsTime
	} else {
		r.BBSProoGenerationTimePerVP = r.BBSProoGenerationTimePerVP + bbsTime
		r.BBSProoGenerationTimePerVP = r.BBSProoGenerationTimePerVP / 2
	}
}

func (r *Results) IncrementNumberofVCsRetrievedWitnessesFromIssuer(){
	r.NumberOfVCsRetrievedWitnessFromIssuer++
}

func (r *Results) IncrementNumberofVCsRetrievedWitnessesFromDLT(){
	r.NumberOfVCsRetrievedWitnessFromDLT++
}


func (r *Results) AddRevocationCostPerBatch(revocationCost int64){
	r.RevocationCostRawData = append(r.RevocationCostRawData, revocationCost)
	if r.AmountPaid==0{
		r.AmountPaid = r.AmountPaid + revocationCost
	} else {
		r.AmountPaid = r.AmountPaid + revocationCost
		r.AmountPaid = r.AmountPaid / 2
	}
}


func (r Results) String() string{
	var response string
	response = response + "Total VCs : "+fmt.Sprintf("%d",r.TotalVCs)+ "\n"
	response = response + "Total Revoked VCs : "+fmt.Sprintf("%d",r.RevokedVCs)+ "\n"
	response = response + "revocation Batch size : "+fmt.Sprintf("%d",r.RevocationBatchSize)+ "\n"
	response = response + "False Positive Rate : "+fmt.Sprintf("%f",r.FalsePositiveRate)+ "\n"
	response = response + "Merkle Tree Accumulator height : "+fmt.Sprintf("%d",r.MTHeight)+ "\n"
	response = response + "Merkle Tree Accumulator Level Stored in DLT : "+fmt.Sprintf("%d",r.MtLevelInDLT)+ "\n \n"
	response = response + "Bloom Filter Size (in bytes) : "+fmt.Sprintf("%d",r.BloomFilterSize)+ "\n"
	response = response + "Bloom Filter indexes per entry (no of hash functions) : "+fmt.Sprintf("%d",r.BloomFilterIndexesPerEntry)+ "\n"
	response = response + "merkle tree size total (in bytes) : "+fmt.Sprintf("%d",r.MerkleTreeSizeTotal)+ "\n"
	response = response + "merkle tree size in DLT (in bytes) : "+fmt.Sprintf("%d",r.MerkleTreeSizeInDLT)+ "\n"
	response = response + "merkle tree total nodes count : "+fmt.Sprintf("%d",r.MerkleTreeNodesCountTotal)+ "\n"
	response = response + "merkle tree nodes count in DLT : "+fmt.Sprintf("%d",r.MerkleTreeNodesCountInDLT)+ "\n"

	response = response + "revocation time per Batch : "+fmt.Sprintf("%f",r.RevocationTimeperBatch)+ "\n"
	response = response + "revocation time Total : "+fmt.Sprintf("%f",r.RevocationTimeTotal)+ "\n"
	response = response + "avg BBS proof generation time per VP : "+fmt.Sprintf("%f",r.BBSProoGenerationTimePerVP)+ "\n"
	response = response + "avg BBS verification time per valid VP : "+fmt.Sprintf("%f",r.BBSVerificationTimePerVP)+ "\n"
	response = response + "verification time per valid VC : "+fmt.Sprintf("%f",r.VerificationTimePerValidVC)+ "\n"
	response = response + "verification time per false positive and revoked VC : "+fmt.Sprintf("%f",r.VerificationTimePerRevokedorFalsePositiveVC)+ "\n"
	response = response + "avg time to fetch witness from issuer: "+fmt.Sprintf("%f",r.AvgTimeToFetchWitnessFromIssuer)+ "\n"
	response = response + "avg time to fetch witness from smart contract: "+fmt.Sprintf("%f",r.AvgTimeToFetchWitnessFromSmartContract)+ "\n"
	response = response + "verification time total valid VCs : "+fmt.Sprintf("%f",r.VerificationTimeTotalValidVCs)+ "\n"
	response = response + "verification time total false positive and revoked VC : "+fmt.Sprintf("%f",r.VerificationTimeTotalRevokedorFalsePositiveVCs)+ "\n"
	response = response + "verification time Total : "+fmt.Sprintf("%f",r.VerificationTimeTotal)+ "\n"

	response = response + "Number of False Positives : "+ fmt.Sprintf("%d",r.NumberOfFalsePositives) + "\n"
	response = response + "Number of VCS that ended up updating witnesses from issuer: "+fmt.Sprintf("%d",r.NumberOfVCsRetrievedWitnessFromIssuer)+ "\n"
	response = response + "Number of VCS that updated witnesses from smart contract: "+fmt.Sprintf("%d",r.NumberOfVCsRetrievedWitnessFromDLT)+ "\n"

	response = response + "contract deployment cost: "+fmt.Sprintf("%d",r.ContractDeploymentCost)+ "\n \n"
	response = response + "Amount (in unit of gas) spent per revocation: "+fmt.Sprintf("%d",r.AmountPaid)+ "\n"
	response = response + "Amount (in unit of gas) spent for bulk issuance: "+fmt.Sprintf("%d",r.BulkIssuanceCost)+ "\n"
	response = response + "time to run the experiment (in seconds) : "+fmt.Sprintf("%f",r.SimulationTime)+ " \n"
	return response
}


func  WriteToFile(result Results) {

	var results []Results
	//filename := fmt.Sprintf("results/results_computed.json")
	filename := fmt.Sprintf("results/results_%d_%d_%f_%d.json",result.TotalVCs, result.RevokedVCs, result.FalsePositiveRate, result.MtLevelInDLT)
	jsonFile, err := os.Open(filename)
	if err != nil {
		jsonFile2, err2 := os.Create(filename)
		if err2 != nil {
			zap.S().Errorln("ERROR - results.json file creation error")
		}
		resJson, _ := ioutil.ReadAll(jsonFile2)
		json.Unmarshal(resJson, &results)
		results = append(results, result)
	} else{
		resJson, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(resJson, &results)
		results = append(results, result)
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

func BloomFilterConfigurationGenerators(totalNumberOfVCs uint, falsePositiveRate float64) (uint, uint) {
	size, numberOfIndexesPerEntry := bloom.EstimateParameters(totalNumberOfVCs, falsePositiveRate)
	return size, numberOfIndexesPerEntry
}

func ConstructResults(config config.Config, start  time.Time, result *Results){
	//zap.S().Infoln("RESULT - \t indexes of VCs that retrieved witnesses from entities: ", result.FetchedWitnessesFromIssuers)
	_, k := bloom.EstimateParameters(config.ExpectedNumberofRevokedVCs,config.FalsePositiveRate)
	// Code to measure
	end := time.Since(start)
	zap.S().Infof("SIMULATOR : \t total time to run the experiment: %f", end.Seconds())

	result.SimulationTime = end.Seconds()
	result.TotalVCs = int(config.ExpectedNumberOfTotalVCs)
	result.RevokedVCs =  int(config.ExpectedNumberofRevokedVCs)
	result.FalsePositiveRate = config.FalsePositiveRate
	result.MTHeight = int(config.MTHeight)
	result.MtLevelInDLT = int(config.MtLevelInDLT)
	result.BloomFilterIndexesPerEntry = int(k)
	result.MerkleTreeNodesCountTotal = int(math.Pow(2, float64(config.MTHeight+1)))-1
	result.MerkleTreeNodesCountInDLT = int(math.Pow(2, float64(config.MtLevelInDLT+1)))-1
	//zap.S().Infoln("SIMULATOR : \t results: ", result.String())
}


func (r *Results) Json() ([]byte, error){
	//return json.MarshalIndent(r, "","    ")
	return json.Marshal(r)
}

//func (r Results) MarshalJSON() ([]byte, error) {
//	var response string
//	response = response + "Total VCs : "+fmt.Sprintf("%d",r.TotalVCs)+ "\n"
//	response = response + "Total Revoked VCs : "+fmt.Sprintf("%d",r.RevokedVCs)+ "\n"
//	response = response + "revocation Batch size : "+fmt.Sprintf("%d",r.RevocationBatchSize)+ "\n"
//	response = response + "False Positive Rate : "+fmt.Sprintf("%f",r.FalsePositiveRate)+ "\n"
//	response = response + "Merkle Tree Accumulator height : "+fmt.Sprintf("%d",r.MTHeight)+ "\n"
//	response = response + "Merkle Tree Accumulator Level Stored in DLT : "+fmt.Sprintf("%d",r.MtLevelInDLT)+ "\n \n"
//	response = response + "Bloom Filter Size (in bytes) : "+fmt.Sprintf("%d",r.BloomFilterSize)+ "\n"
//	response = response + "Bloom Filter indexes per entry (no of hash functions) : "+fmt.Sprintf("%d",r.BloomFilterIndexesPerEntry)+ "\n"
//	response = response + "merkle tree size total (in bytes) : "+fmt.Sprintf("%d",r.MerkleTreeSizeTotal)+ "\n"
//	response = response + "merkle tree size in DLT (in bytes) : "+fmt.Sprintf("%d",r.MerkleTreeSizeInDLT)+ "\n"
//	response = response + "merkle tree total nodes count : "+fmt.Sprintf("%d",r.MerkleTreeNodesCountTotal)+ "\n"
//	response = response + "merkle tree nodes count in DLT : "+fmt.Sprintf("%d",r.MerkleTreeNodesCountInDLT)+ "\n"
//
//	response = response + "revocation time per Batch : "+fmt.Sprintf("%f",r.RevocationTimeperBatch)+ "\n"
//	response = response + "revocation time Total : "+fmt.Sprintf("%f",r.RevocationTimeTotal)+ "\n"
//	response = response + "avg BBS proof generation time per VP : "+fmt.Sprintf("%f",r.BBSProoGenerationTimePerVP)+ "\n"
//	response = response + "avg BBS verification time per valid VP : "+fmt.Sprintf("%f",r.BBSVerificationTimePerVP)+ "\n"
//	response = response + "verification time per valid VC : "+fmt.Sprintf("%f",r.VerificationTimePerValidVC)+ "\n"
//	response = response + "verification time per false positive and revoked VC : "+fmt.Sprintf("%f",r.VerificationTimePerRevokedorFalsePositiveVC)+ "\n"
//	response = response + "avg time to fetch witness from issuer: "+fmt.Sprintf("%f",r.AvgTimeToFetchWitnessFromIssuer)+ "\n"
//	response = response + "avg time to fetch witness from smart contract: "+fmt.Sprintf("%f",r.AvgTimeToFetchWitnessFromSmartContract)+ "\n"
//	response = response + "verification time total valid VCs : "+fmt.Sprintf("%f",r.VerificationTimeTotalValidVCs)+ "\n"
//	response = response + "verification time total false positive and revoked VC : "+fmt.Sprintf("%f",r.VerificationTimeTotalRevokedorFalsePositiveVCs)+ "\n"
//	response = response + "verification time Total : "+fmt.Sprintf("%f",r.VerificationTimeTotal)+ "\n"
//
//	response = response + "Number of False Positives : "+ fmt.Sprintf("%d",r.NumberOfFalsePositives) + "\n"
//	response = response + "Number of VCS that ended up updating witnesses from issuer: "+fmt.Sprintf("%d",r.NumberOfVCsRetrievedWitnessFromIssuer)+ "\n"
//	response = response + "Number of VCS that updated witnesses from smart contract: "+fmt.Sprintf("%d",r.NumberOfVCsRetrievedWitnessFromDLT)+ "\n"
//
//	response = response + "contract deployment cost: "+fmt.Sprintf("%d",r.ContractDeploymentCost)+ "\n \n"
//	response = response + "Amount (in unit of gas) spent per revocation: "+fmt.Sprintf("%d",r.AmountPaid)+ "\n"
//	response = response + "Amount (in unit of gas) spent for bulk issuance: "+fmt.Sprintf("%d",r.BulkIssuanceCost)+ "\n"
//
//
//	return []byte(response), nil
//}


func JsonToResults(jsonObj []byte) *Results {
	res := Results{}
	json.Unmarshal(jsonObj, &res)
	return &res
}
