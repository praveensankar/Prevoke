package main

import (
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/contracts"
	"github.com/spf13/viper"
)

func setupConfig(){
	viper.SetConfigFile("config.json")// name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
}

func main()  {
	//testAries()
	setupConfig()
	conf, _ := config.ParseConfig()
	//blockchain.TestConnectionToBlockchain(conf)
	contracts.TestSmartContract(conf)

	//testIssuer(config)
}
