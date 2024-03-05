package Results

import (
	"encoding/json"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	mapset "github.com/deckarep/golang-set"
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"os"
	"time"
)

type Results struct {
	TotalVCs int `json:"total_issued_vcs"`
	RevokedVCs int `json:"total_revoked_vcs"`
	FalsePositiveRate float64 `json:"false_positive_rate"`
	MtLevelInDLT int `json:"mt_level_in_dlt"`
	MTHeight int `json:"mt_height"`
	NumberOfFalsePositives int `json:"number_of_false_positives"`
	NumberOfVCsRetrievedWitnessFromIssuer int `json:"number_of_vcs_retrieved_witness_from_issuer"`
	NumberOfVCsRetrievedWitnessFromDLT int `json:"number_of_vcs_retrieved_witness_from_dlt"`
	AmountPaid int64 `json:"mt_accumulator_per_update_cost_in_gwei"`
	BloomFilterSize int `json:"bloom_filter_size"`
	BloomFilterIndexesPerEntry int  `json:"bloom_filter_indexes_per_entry"`
	MerkleTreeSizeTotal int `json:"merkle_tree_size_total"`
	MerkleTreeSizeInDLT int `json:"merkle_tree_size_in_dlt"`
	MerkleTreeNodesCountTotal int `json:"merkle_tree_nodes_count_total"`
	MerkleTreeNodesCountInDLT int `json:"merkle_tree_nodes_count_in_dlt"`
	RevocationTimeperBatch float64 `json:"revocation_timeper_vc"`
	RevocationBatchSize int `json:"revocation_batch_size"`
	RevocationTimeTotal float64 `json:"revocation_time_total"`
	VerificationTimeTotalValidVCs float64 `json:"verification_time_total_valid_vcs"`
	VerificationTimeTotalRevokedorFalsePositiveVCs float64 `json:"verification_time_total_false_positive_and_revoked_vcs"`
	VerificationTimePerValidVC float64 `json:"verification_time_per_valid_vc"`
	BBSVerificationTimePerVP float64`json:"bbs_verification_time"`
	VerificationTimePerRevokedorFalsePositiveVC float64 `json:"verification_time_per_false_positive_or_revoked_vc"`
	VerificationTimeTotal float64 `json:"verification_time_total"`
	SimulationTime float64 `json:"simulation_time"`
	AffectedIndexes mapset.Set
	FalsePositiveResults mapset.Set
	FetchedWitnessesFromIssuers mapset.Set
}

func CreateResult() *Results {
	result := &Results{}
	result.AffectedIndexes = mapset.NewSet()
	result.FalsePositiveResults = mapset.NewSet()
	result.FetchedWitnessesFromIssuers = mapset.NewSet()
	return result
}

func (r *Results) AddVerificationTimeTotal(vTime float64){
	r.VerificationTimeTotal = r.VerificationTimeTotal+vTime
}
func (r *Results) AddVerificationTimeTotalValidVCs(phase1Time float64){
	r.VerificationTimeTotalValidVCs = r.VerificationTimeTotalValidVCs+phase1Time
}

func (r *Results) AddVerificationTimePerValidVC(phase1Time float64){
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
	if r.VerificationTimePerRevokedorFalsePositiveVC==0.0{
		r.VerificationTimePerRevokedorFalsePositiveVC = r.VerificationTimePerRevokedorFalsePositiveVC + phase2Time
	} else {
		r.VerificationTimePerRevokedorFalsePositiveVC = r.VerificationTimePerRevokedorFalsePositiveVC + phase2Time
		r.VerificationTimePerRevokedorFalsePositiveVC = r.VerificationTimePerRevokedorFalsePositiveVC / 2
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

func (r *Results) IncrementNumberofVCsRetrievedWitnessesFromIssuer(){
	r.NumberOfVCsRetrievedWitnessFromIssuer++
}

func (r *Results) IncrementNumberofVCsRetrievedWitnessesFromDLT(){
	r.NumberOfVCsRetrievedWitnessFromDLT++
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
	response = response + "verification time total valid VCs : "+fmt.Sprintf("%f",r.VerificationTimeTotalValidVCs)+ "\n"
	response = response + "verification time total false positive and revoked VC : "+fmt.Sprintf("%f",r.VerificationTimeTotalRevokedorFalsePositiveVCs)+ "\n"
	response = response + "BBS verification time per valid VP : "+fmt.Sprintf("%f",r.BBSVerificationTimePerVP)+ "\n"
	response = response + "verification time per valid VC : "+fmt.Sprintf("%f",r.VerificationTimePerValidVC)+ "\n"
	response = response + "verification time per false positive and revoked VC : "+fmt.Sprintf("%f",r.VerificationTimePerRevokedorFalsePositiveVC)+ "\n"
	response = response + "verification time Total : "+fmt.Sprintf("%f",r.VerificationTimeTotal)+ "\n"

	response = response + "Amount (in gwei) paid per revocation: "+fmt.Sprintf("%d",r.AmountPaid)+ "\n"
	response = response + "Number of False Positives : "+ fmt.Sprintf("%d",r.NumberOfFalsePositives) + "\n"
	response = response + "Number of VCS that ended up updating witnesses from issuer: "+fmt.Sprintf("%d",r.NumberOfVCsRetrievedWitnessFromIssuer)+ "\n"
	response = response + "Number of VCS that updated witnesses from smart contract: "+fmt.Sprintf("%d",r.NumberOfVCsRetrievedWitnessFromDLT)+ "\n"
	response = response + "time to run the experiment (in seconds) : "+fmt.Sprintf("%f",r.SimulationTime)+ " \n"
	return response
}


func  WriteToFile( filename string, result Results) {

	var results []Results
	jsonFile, err := os.Open(filename)
	if err != nil {
		zap.S().Errorln("ERROR - results.json file open error")
	}
	resJson, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(resJson, &results)
	results = append(results, result)
	jsonRes, _ := json.MarshalIndent(results,"","")
	//filename := fmt.Sprintf("Simulation/results/result_%v_%v_%v.json",numberOfVcs, numberOfRevokedVcs, mtLevelInDLT)
	err = ioutil.WriteFile(filename, jsonRes, 0644)
	if err != nil {
		zap.S().Errorln("unable to write results to file")
	}
	zap.S().Errorln("RESULTS - successfully written to the file")

}

func BloomFilterConfigurationGenerators(totalNumberOfVCs uint, falsePositiveRate float64) (uint, uint) {
	size, numberOfIndexesPerEntry := bloom.EstimateParameters(totalNumberOfVCs, falsePositiveRate)
	return size, numberOfIndexesPerEntry
}

func ConstructResults(config config.Config, start  time.Time, result *Results){
	zap.S().Infoln("RESULT - \t indexes of VCs that retrieved witnesses from entities: ", result.FetchedWitnessesFromIssuers)
	size, k :=BloomFilterConfigurationGenerators(config.ExpectedNumberofRevokedVCs,config.FalsePositiveRate)
	// Code to measure
	end := time.Since(start)
	zap.S().Infof("SIMULATOR : \t total time to run the experiment: %f", end.Seconds())

	result.SimulationTime = end.Seconds()
	result.TotalVCs = int(config.ExpectedNumberOfTotalVCs)
	result.RevokedVCs =  int(config.ExpectedNumberofRevokedVCs)
	result.FalsePositiveRate = config.FalsePositiveRate
	result.MTHeight = int(config.MTHeight)
	result.MtLevelInDLT = int(config.MtLevelInDLT)
	result.BloomFilterSize = int(size)
	result.BloomFilterIndexesPerEntry = int(k)
	result.MerkleTreeNodesCountTotal = int(math.Pow(2, float64(config.MTHeight+1)))-1
	result.MerkleTreeNodesCountInDLT = int(math.Pow(2, float64(config.MtLevelInDLT+1)))-1
	zap.S().Infoln("SIMULATOR : \t results: ", result.String())
}


func (r *Results) Json() ([]byte, error){
	return json.MarshalIndent(r, "","    ")
}



func JsonToResults(jsonObj []byte) *Results{
	res := Results{}
	json.Unmarshal(jsonObj, &res)
	return &res
}
