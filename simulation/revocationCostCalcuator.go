package simulation

import (
	"encoding/hex"
	"fmt"
	"github.com/praveensankar/Revocation-Service/blockchain"
	"github.com/praveensankar/Revocation-Service/common"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/revocation_service"
	"github.com/praveensankar/Revocation-Service/techniques"
	"go.uber.org/zap"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

func RevocationCostCalculator(conf config.Config){
	falsePositiveRates:= []float64{0.1,0.01,0.001,0.0001}
	mtHeights:= []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	totalRevokedVCs:=100

	bfSize := 10000
	var RevocationResults []common.RevocationCostResults

	for exp:=1;exp<=3;exp++ {
		counter := 1
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < len(falsePositiveRates); i++ {
			for j := 0; j < len(mtHeights); j++ {

				zap.S().Infoln("exp: ", counter, "/", len(falsePositiveRates)*len(mtHeights), " running")
				counter++
				address, _, err := blockchain.DeployContract(conf, 0)

				if err != nil {
					zap.S().Errorln("error deploying contract")
				}

				conf.SmartContractAddress = address
				rs := revocation_service.CreateRevocationService(conf)
				bf := techniques.CreateBloomFilter(uint(bfSize), falsePositiveRates[i])

				result := common.CreateRevocationCostResults()
				result.RevokedVCs = totalRevokedVCs
				result.FalsePositiveRate = falsePositiveRates[i]
				result.MtLevelInDLT = mtHeights[j]
				for k := 0; k < totalRevokedVCs; k++ {
					bfIndexes := bf.GetIndexes(strconv.Itoa(k + 1))
					var mtIndexes []*big.Int
					var mtValuesInBytes [][32]byte
					for x := 0; x <= mtHeights[j]; x++ {
						s := strconv.Itoa(rand.Int())
						h, _ := hex.DecodeString(s)
						byteRepr := [32]byte{}
						copy(byteRepr[:], h[:])
						mtValuesInBytes = append(mtValuesInBytes, byteRepr)
						mtIndexes = append(mtIndexes, big.NewInt(int64(x)))
					}
					gasUsed, err := rs.RevocationCostCalculator(bfIndexes, mtIndexes, mtValuesInBytes)
					if err != nil {
						zap.S().Infoln("error revoking in the smart contract")
					}
					result.AddRevocationCostPerBatch(gasUsed)
				}
				RevocationResults = append(RevocationResults, *result)
			}
		}
	}
	filename := fmt.Sprintf("results/results_revocation_cost.json")
	common.WriteRevocationCostResultsToFile(filename, RevocationResults)

}
