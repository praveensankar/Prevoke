package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct{
	SmartContractAddress string
	BlockchainRpcEndpoint string
	senderAddress string
	privateKey string
	passPhrase string
	otherAccounts []string
}

func (config Config) printConfig()  {
	fmt.Println("smart contract address: ",config.SmartContractAddress)
	fmt.Println("blockchain rpc endpoint: ",config.BlockchainRpcEndpoint)
	fmt.Println("sender address: ",config.senderAddress)
	fmt.Println("sender private key: ",config.privateKey)
	fmt.Println("sender pass phrase: ",config.passPhrase)
	fmt.Println("other accounts in ganache test network : ", config.otherAccounts)
}


func parseConfig() (Config, error){

	viper.SetConfigFile("config.json")// name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	config := Config{}
	config.SmartContractAddress = viper.GetString("contract.address")
	config.BlockchainRpcEndpoint = viper.GetString("blockchain.rpcEndpoint")
	config.senderAddress = viper.GetString("account.address")
	config.privateKey = viper.GetString("account.privateKey")
	config.passPhrase = viper.GetString("account.passphrase")
	config.otherAccounts = viper.GetStringSlice("otherAccounts")

	//"account1" :  "0xB97F44Ce8dA7E824F7aBD0068F92D08438E3405A",
	//	"account2" : "0x6C3d120Ee76E635d7b221a996718a8277BeA973f",
	//	"account3" : "0xF82407B704B5FF6AB71894ec0f1d78f514c3A13A",
	//	"_comment": " three other accounts in the ganache test network"
	config.printConfig()
	return config, nil
}