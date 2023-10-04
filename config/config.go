package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct{
	SmartContractAddress string
	BlockchainRpcEndpoint string
	BlockchainWebSocketEndPoint string
	SenderAddress string
	PrivateKey    string
	passPhrase    string
	OtherAccounts []string
}

func (config Config) printConfig()  {
	fmt.Println("smart contract address: ",config.SmartContractAddress)
	fmt.Println("blockchain rpc endpoint: ",config.BlockchainRpcEndpoint)
	fmt.Println("sender address: ",config.SenderAddress)
	fmt.Println("sender private key: ",config.PrivateKey)
	fmt.Println("sender pass phrase: ",config.passPhrase)
	fmt.Println("other accounts in ganache test network : ", config.OtherAccounts)
}


func ParseConfig() (Config, error){



	err := viper.ReadInConfig()
	if err!=nil{
		log.Fatal("error reading config file for viper.\t", err)
	}
	config := Config{}
	config.SmartContractAddress = viper.GetString("contract.address")
	config.BlockchainRpcEndpoint = viper.GetString("blockchain.rpcEndpoint")
	config.BlockchainWebSocketEndPoint = viper.GetString("blockchain.wsEndPoint")
	config.SenderAddress = viper.GetString("account.address")
	config.PrivateKey = viper.GetString("account.privateKey")
	config.passPhrase = viper.GetString("account.passphrase")
	config.OtherAccounts = viper.GetStringSlice("otherAccounts")

	//"account1" :  "0xB97F44Ce8dA7E824F7aBD0068F92D08438E3405A",
	//	"account2" : "0x6C3d120Ee76E635d7b221a996718a8277BeA973f",
	//	"account3" : "0xF82407B704B5FF6AB71894ec0f1d78f514c3A13A",
	//	"_comment": " three other accounts in the ganache test network"
	config.printConfig()
	return config, nil
}