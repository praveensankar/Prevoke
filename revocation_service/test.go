package revocation_service

import (
	"github.com/praveensankar/Revocation-Service/blockchain"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/signature"
	"go.uber.org/zap"
	"math"
)

func DeployContract(conf *config.Config){
	address, err := blockchain.DeployContract(*conf)

	if err != nil {
		zap.S().Errorln("error deploying contract")
	}

	conf.SmartContractAddress = address
}


func TestRevocationService(config config.Config) {

	DeployContract(&config)

	rs := CreateRevocationService(config)

	remainingSpace := int(math.Pow(2, float64(config.MTHeight)))-int(config.ExpectedNumberOfTotalVCs)
	var vcIDs []string

	for i:=0; i<(int(config.ExpectedNumberOfTotalVCs)+remainingSpace);i++{
			vcIDs=append(vcIDs, string(i))
	}
	rs.IssueVCsInBulk(vcIDs)
	rs.RevokeVC(vcIDs[3])
	rs.RevokeVC(vcIDs[6])
	rs.PrintMerkleTree()
	keyPair1 := signature.GenerateKeyPair()
	keyPair2 := signature.GenerateKeyPair()
	publicKey1, _ := keyPair1.PublicKey.Marshal()
	publicKey2, _ := keyPair2.PublicKey.Marshal()
	keys := make([][]byte, 2)
	keys[0]=publicKey1
	keys[1]=publicKey2
	rs.AddPublicKeys(keys)
	publicKeys := rs.FetchPublicKeys()
	zap.S().Infoln("REVOCATION SERVICE  TEST - \t  issuers public keys: ", publicKeys)

}
