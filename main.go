package main

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/simulation"
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

/*
This function that returns estimated size of bloom filter and number of hash functions


Inputs:
	TotalNumberofVCs - number of VCs issuer expects to issue in its lifetime
	falsePositiveRate - false positive rate of bloomfilter

Output:
	size - number of entries in bloomfilter
	numberOfIndexesPerEntry - number of indexes per entry
*/
func BloomFilterConfigurationGenerators(totalNumberOfVCs uint, falsePositiveRate float64) (uint, uint) {
	size, numberOfIndexesPerEntry := bloom.EstimateParameters(totalNumberOfVCs, falsePositiveRate)
	return size, numberOfIndexesPerEntry
}

func main()  {
	//testAries()
	initialize()
	conf, _ := config.ParseConfig()
	simulation.Start(conf)
	//size, numberofIndexesPerEntry := BloomFilterConfigurationGenerators(conf.ExpectedNumberofRevokedVCs, conf.FalsePositiveRate)
	//zap.S().Infoln("bloom filter size: ", size, "\t number of hash functions: ", numberofIndexesPerEntry)
	//blockchain.TestConnectionToBlockchain(conf)
	//blockchain.Test(conf)
	//techniques.TestMerkleTree()
	//techniques.TestBloomFilter(100)
	//issuer.TestIssuer(conf)
}


