package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strconv"
)

type Config struct{
	SmartContractAddress string
	GasLimit uint64
	BlockchainRpcEndpoint string
	BlockchainWebSocketEndPoint string
	SenderAddress string
	PrivateKey    string
	passPhrase    string
	OtherAccounts []string
	LoggerType string
	IssuerName string
	ExpectedNumberOfTotalVCs uint
	ExpectedNumberofRevokedVCs uint
	MtLevelInDLT uint
	FalsePositiveRate float64
	DEBUG bool
}

func (config Config) printConfig()  {
	zap.L().Info("\n\n--------------------------------------------------------printing blockchain related configuration--------------------------------------------------")
	zap.L().Info("smart contract address:"+config.SmartContractAddress)
	zap.L().Info("blockchain rpc endpoint: "+config.BlockchainRpcEndpoint)
	zap.L().Info("gas limit: "+string(config.GasLimit))
	zap.L().Info("sender address: "+config.SenderAddress)
	zap.L().Info("sender private key: "+config.PrivateKey)
	zap.L().Info("sender pass phrase: "+config.passPhrase)
	zap.S().Infoln("other accounts in ganache test network : ",config.OtherAccounts)
	zap.L().Info("logger environment: "+config.LoggerType)
	zap.L().Info("********************************************************************************************************************************\n")
	zap.L().Info("\n\n--------------------------------------------------------printing issuer configuration--------------------------------------------------")
	zap.L().Info("issuer name:"+config.IssuerName)
	zap.L().Info("total of VCs would be issued:"+ strconv.Itoa(int(config.ExpectedNumberOfTotalVCs)))
	zap.L().Info("total of VCs would be revoked:"+ strconv.Itoa(int(config.ExpectedNumberofRevokedVCs)))
	zap.S().Infoln("bloom filter false positive rate: ",config.FalsePositiveRate)
	zap.S().Infoln("merkle tree accumulator level in DLT: ", config.MtLevelInDLT)
	zap.L().Info("********************************************************************************************************************************\n")

}


func ParseConfig() (Config, error){



	err := viper.ReadInConfig()
	if err!=nil{
		zap.S().Fatalln("error reading config file for viper.\t", err)
	}
	config := Config{}
	config.SmartContractAddress = viper.GetString("contract.address")
	config.GasLimit = viper.GetUint64("contract.gasLimit")
	config.BlockchainRpcEndpoint = viper.GetString("blockchain.rpcEndpoint")
	config.BlockchainWebSocketEndPoint = viper.GetString("blockchain.wsEndPoint")
	config.SenderAddress = viper.GetString("account.address")
	config.PrivateKey = viper.GetString("account.privateKey")
	config.passPhrase = viper.GetString("account.passphrase")
	config.OtherAccounts = viper.GetStringSlice("otherAccounts")
	config.LoggerType = viper.GetString("logger.env")
	config.ExpectedNumberOfTotalVCs = viper.GetUint("issuer.totalVCs")
	config.ExpectedNumberofRevokedVCs = viper.GetUint("issuer.revokedVCs")
	config.FalsePositiveRate = viper.GetFloat64("issuer.falsePositiveRate")
	config.MtLevelInDLT = viper.GetUint("issuer.mtLevelInDLT")
	config.IssuerName = viper.GetString("issuer.name")
	config.DEBUG = viper.GetBool("mode.debug")
	//"account1" :  "0xB97F44Ce8dA7E824F7aBD0068F92D08438E3405A",
	//	"account2" : "0x6C3d120Ee76E635d7b221a996718a8277BeA973f",
	//	"account3" : "0xF82407B704B5FF6AB71894ec0f1d78f514c3A13A",
	//	"_comment": " three other accounts in the ganache test network"
	config.printConfig()
	return config, nil
}