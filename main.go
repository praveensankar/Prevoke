package main

import (
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/techniques"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setupLogger(conf config.Config){
	zapConfig := &zap.Config{
		Encoding: "console",
		Level: zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level",
			MessageKey: "***",
		},
	}
	if conf.LoggerType == "dev"{
		zap.ReplaceGlobals(zap.Must(zapConfig.Build()))
	} else if conf.LoggerType=="prod"{
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
	} else{
		zap.ReplaceGlobals(zap.Must(zapConfig.Build()))
	}



}

func setupConfig(){
	viper.SetConfigFile("config.json")// name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
}

func initialize() {
	setupConfig()
	conf, _ := config.ParseConfig()
	setupLogger(conf)
}

func main()  {
	//testAries()
	initialize()
	_, _ = config.ParseConfig()
	//blockchain.TestConnectionToBlockchain(conf)
	//contracts.TestSmartContract(conf)
	techniques.TestMerkleTree()
	//testIssuer(config)
}


