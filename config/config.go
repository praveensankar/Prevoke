package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct{
	SmartContractAddress string
	BlockchainRpcEndpoint string
	BlockchainWebSocketEndPoint string
	SenderAddress string
	PrivateKey    string
	passPhrase    string
	OtherAccounts []string
	LoggerType string
}

func (config Config) printConfig()  {
	zap.L().Info("\n\n--------------------------------------------------------printing configuration--------------------------------------------------")
	zap.L().Info("smart contract address:"+config.SmartContractAddress)
	zap.L().Info("blockchain rpc endpoint: "+config.BlockchainRpcEndpoint)
	zap.L().Info("sender address: "+config.SenderAddress)
	zap.L().Info("sender private key: "+config.PrivateKey)
	zap.L().Info("sender pass phrase: "+config.passPhrase)
	zap.S().Infoln("other accounts in ganache test network : ",config.OtherAccounts)
	zap.L().Info("logger environment: "+config.LoggerType)
	zap.L().Info("********************************************************************************************************************************\n")
}


func ParseConfig() (Config, error){



	err := viper.ReadInConfig()
	if err!=nil{
		zap.S().Fatalln("error reading config file for viper.\t", err)
	}
	config := Config{}
	config.SmartContractAddress = viper.GetString("contract.address")
	config.BlockchainRpcEndpoint = viper.GetString("blockchain.rpcEndpoint")
	config.BlockchainWebSocketEndPoint = viper.GetString("blockchain.wsEndPoint")
	config.SenderAddress = viper.GetString("account.address")
	config.PrivateKey = viper.GetString("account.privateKey")
	config.passPhrase = viper.GetString("account.passphrase")
	config.OtherAccounts = viper.GetStringSlice("otherAccounts")
	config.LoggerType = viper.GetString("logger.env")
	//"account1" :  "0xB97F44Ce8dA7E824F7aBD0068F92D08438E3405A",
	//	"account2" : "0x6C3d120Ee76E635d7b221a996718a8277BeA973f",
	//	"account3" : "0xF82407B704B5FF6AB71894ec0f1d78f514c3A13A",
	//	"_comment": " three other accounts in the ganache test network"
	config.printConfig()
	return config, nil
}