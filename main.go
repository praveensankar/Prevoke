package main

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/techniques"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func setupLogger(conf config.Config){

	var filename string

	if len(os.Args) > 1 {
		if os.Args[1] == "size" {
			conf.LoggerOutputMode = "console"
		}
		if os.Args[1] == "simulation" {
			conf.LoggerOutputMode = "file"
		}
	}
	if conf.LoggerOutputMode=="console"{
		filename="stdout"
	}
	if conf.LoggerOutputMode=="file"{
		filename = fmt.Sprintf("logs/%v_%v_%v_%f_%v",conf.LoggerFile, conf.ExpectedNumberOfTotalVCs,
			conf.ExpectedNumberofRevokedVCs, conf.FalsePositiveRate, conf.MtLevelInDLT)

	}
	//OutputPaths: []string{"stdout"},
//OutputPaths: []string{filename},
	zapConfig := &zap.Config{
		Encoding: "console",
		Level: zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{filename},
		ErrorOutputPaths: []string{filename},
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

func TestIndividualComponents(conf config.Config){
	//blockchain.TestConnectionToBlockchain(conf)
	//blockchain.Test(conf)
	//techniques.TestMerkleTree(conf)
	techniques.TestMerkleTreeAccumulator(conf)
	//techniques.TestBloomFilter(100)
	//issuer.TestIssuer(conf)
}

func main()  {
	//testAries()
	initialize()
	conf, _ := config.ParseConfig()

	//if os.Args[1]=="simulation"{
	//	simulation.Start(conf)
	//}
	//
	//if os.Args[1]=="size" {
	//	size, numberofIndexesPerEntry := BloomFilterConfigurationGenerators(conf.ExpectedNumberofRevokedVCs, conf.FalsePositiveRate)
	//	fmt.Println("bloom filter size: ", size, "\t number of hash functions: ", numberofIndexesPerEntry)
	//}

	TestIndividualComponents(conf)

}


