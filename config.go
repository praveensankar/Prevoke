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
}

func (config Config) printConfig()  {
	fmt.Println("smart contract address: ",config.SmartContractAddress)
	fmt.Println("blockchain rpc endpoint: ",config.BlockchainRpcEndpoint)
	fmt.Println("sender address: ",config.senderAddress)
	fmt.Println("sender private key: ",config.privateKey)
	fmt.Println("sender pass phrase: ",config.passPhrase)
}


func parseConfig() (Config, error){

	viper.SetConfigFile("config.json")// name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	fmt.Println(viper.Get("test"))
	config := Config{}
	config.SmartContractAddress = viper.GetString("contract.address")
	config.BlockchainRpcEndpoint = viper.GetString("blockchain.rpcEndpoint")
	config.senderAddress = viper.GetString("account.address")
	config.privateKey = viper.GetString("account.privateKey")
	config.passPhrase = viper.GetString("account.passphrase")
	//config.printConfig()
	return config, nil
}