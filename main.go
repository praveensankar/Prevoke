/*
Main package sets up logger, config and then run the simulation

Zap library is used for Logging. The log is stored on file or shown
in the console depending on the setting given in the config file.

The config file contains the parameters needed throughtout the program.

Usage:

The flags are:

    -simulation
        Runs the simulation

	-simulatorTest
		test simulator

	-vcTest
		tests verifiable credential
	-mtTest
		tests merkle tree accumulator

	-bfTest
		tests the bloom filter

	-issuerTest
		tests entities

	-revocationServiceTest
		tests revocation service

	-BBSTest
		test BBS Signature
 */
package main

import (
	"flag"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/Revocation-Service/config"
	"github.com/Revocation-Service/entities"
	"github.com/Revocation-Service/revocation_service"
	"github.com/Revocation-Service/signature"
	"github.com/Revocation-Service/simulation"
	"github.com/Revocation-Service/techniques"
	"github.com/Revocation-Service/vc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
setupLogger sets up the logger.
The logger output mode is retrieved from the config file.
If the output mode is console then the log is shown on console
If the output mode is file then the log is stored in a file
*/
func SetupLogger(conf config.Config, filename string){


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




/*
TestIndividualComponents tests the following components in the project.
1) BloomFilter
2) MerkleTreeAccumulator
3) Issuer
4) Connection to Blockchain
5) simulator
6) vc
7) BBS Signature
 */
func Run(conf config.Config){

	mtTestFlag := flag.Bool("mtTest", false, "a bool")
	issuerTestFlag := flag.Bool("issuerTest", false, "a bool")
	bfTestFlag := flag.Bool("bfTest", false, "a bool")
	vcTestFlag := flag.Bool("vcTest", false, "a bool")
	simulatorTestFlag := flag.Bool("simulatorTest", false, "a bool")
	rsTestFlag := flag.Bool("revocationServiceTest", false, "a bool")
	simulationFlag := flag.Bool("simulation", false, "a bool")
	bbsTestFlag := flag.Bool("bbsTest", false, "a bool")
	holderFlag := flag.Bool("holder", false, "a bool")
	issuerFlag := flag.Bool("issuer", false, "a bool")
	verifierFlag := flag.Bool("verifier", false, "a bool")
	IPFSTestFlag := flag.Bool("ipfsTest", false, "a bool")
	witnessCalculationFlag := flag.Bool("calWitness", false, "a bool")
	revocationScalabilityFlag := flag.Bool("scaleRevocation", false, "a bool")
	revocationCostCalculationFlag := flag.Bool("revocationCostCalculation", false, "a bool")
	flag.Parse()

	var filename string

	if conf.LoggerOutputMode=="console"{
		filename="stdout"
	}


	if *bfTestFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/%v",conf.LoggerFile)
		}
		SetupLogger(conf, filename)
		techniques.TestBloomFilter(100)
	}

	if *vcTestFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/%v",conf.LoggerFile)
		}
		SetupLogger(conf, filename)
		vc.TestVC(conf)
	}

	if *mtTestFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/%v",conf.LoggerFile)
		}
		SetupLogger(conf, filename)
		//techniques.TestMerkleTree(conf)
		techniques.TestMerkleTreeAccumulator(conf)
	}

	if *issuerTestFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/%v",conf.LoggerFile)
		}
		SetupLogger(conf, filename)
		entities.TestIssuer(conf)
	}

	if *rsTestFlag == true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/%v",conf.LoggerFile)
		}
		SetupLogger(conf, filename)
		revocation_service.TestRevocationService(conf)
	}

	if *bbsTestFlag == true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/%v",conf.LoggerFile)
		}
		SetupLogger(conf, filename)
		signature.TestBBS(conf)
	}


	if *simulatorTestFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/%v",conf.LoggerFile)
		}
		SetupLogger(conf, filename)
		simulation.TestSimulator(conf)
	}


	if *simulationFlag==true {
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/%v",conf.LoggerFile)
		}
		SetupLogger(conf, filename)
		simulation.StartExperiments(conf)
	}

	if *holderFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/holder")
		}
		SetupLogger(conf, filename)
		entities.StartHolder(conf)
	}

	if *issuerFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/issuer")
		}
		SetupLogger(conf, filename)
		entities.StartIssuerServer(conf)
	}

	if *verifierFlag == true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/verifier")
		}
		SetupLogger(conf, filename)
		entities.StartVerifierServer(conf)
	}

	if *IPFSTestFlag == true{
		SetupLogger(conf, filename)
		//blockchain.TestIPFS()
	}

	if *witnessCalculationFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/holder")
		}
		SetupLogger(conf, filename)
		simulation.CalculateNumberOfVCsWouldRetrieveWitnessFromDLT(conf)
	}

	if *revocationCostCalculationFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/holder")
		}
		SetupLogger(conf, filename)
		simulation.RevocationCostCalculator(conf)
	}

	if *revocationScalabilityFlag==true{
		if conf.LoggerOutputMode=="file"{
			filename = fmt.Sprintf("logs/holder")
		}
		SetupLogger(conf, filename)
		simulation.CalculateRevocationScalability(conf)
	}
	//blockchain.TestConnectionToBlockchain(conf)
	//blockchain.Test(conf)

}

func main()  {

	//testAries()


	conf, _ := config.ParseConfig()
	Run(conf)


	//if os.Args[1]=="size" {
	//	size, numberofIndexesPerEntry := BloomFilterConfigurationGenerators(conf.ExpectedNumberofRevokedVCs, conf.FalsePositiveRate)
	//	fmt.Println("bloom filter size: ", size, "\t number of hash functions: ", numberofIndexesPerEntry)
	//}

}


/*
This function that returns estimated size of bloom filter and number of hash functions


Inputs:
	TotalNumberofVCs - number of VCs entities expects to issue in its lifetime
	falsePositiveRate - false positive rate of bloomfilter

Output:
	size - number of entries in bloomfilter
	numberOfIndexesPerEntry - number of indexes per entry
*/
func BloomFilterConfigurationGenerators(totalNumberOfVCs uint, falsePositiveRate float64) (uint, uint) {
	size, numberOfIndexesPerEntry := bloom.EstimateParameters(totalNumberOfVCs, falsePositiveRate)
	return size, numberOfIndexesPerEntry
}


