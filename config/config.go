package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"math/big"
	"net"
	"strconv"
)

type Config struct{
	SmartContractAddress string
	GasLimit uint64
	GasPrice *big.Int
	BlockchainRpcEndpoint string
	BlockchainWebSocketEndPoint string
	SenderAddress string
	passPhrase    string
	PrivateKeys []string
	LoggerType string
	LoggerFile string
	LoggerOutputMode string
	IssuerName string
	ExpectedNumberOfTotalVCs uint
	ExpectedNumberofRevokedVCs uint
	ExpParamters  map[string]*Experiment
	MtLevelInDLT uint
	MTHeight uint
	RevocationBatchSize uint
	FalsePositiveRate float64
	DEBUG bool
	IssuerAddress string
	HolderName string
	HolderAddress string
	VerifierName string
	VerifierAddress string
	ManagerAddress string
}

func (config Config) printConfig()  {
	zap.L().Info("\n\n--------------------------------------------------------printing blockchain related configuration--------------------------------------------------")
	zap.L().Info("smart contract address:"+config.SmartContractAddress)
	zap.L().Info("blockchain rpc endpoint: "+config.BlockchainRpcEndpoint)
	zap.L().Info("gas limit: "+string(config.GasLimit))
	zap.L().Info("gas price: "+config.GasPrice.String())
	zap.L().Info("sender address: "+config.SenderAddress)
	zap.L().Info("sender pass phrase: "+config.passPhrase)
	zap.S().Infoln("private keys in ganache test network : ",config.PrivateKeys)
	zap.L().Info("logger environment: "+config.LoggerType)
	zap.L().Info("logger output file name: "+config.LoggerFile)
	zap.L().Info("logger output mode: "+config.LoggerOutputMode)
	zap.L().Info("********************************************************************************************************************************\n")
	zap.L().Info("\n\n--------------------------------------------------------printing issuer configuration--------------------------------------------------")
	zap.L().Info("issuer name:"+config.IssuerName)
	zap.L().Info("issuer address:"+config.IssuerAddress)
	zap.L().Info("total of VCs would be issued:"+ strconv.Itoa(int(config.ExpectedNumberOfTotalVCs)))
	zap.L().Info("total of VCs would be revoked:"+ strconv.Itoa(int(config.ExpectedNumberofRevokedVCs)))
	zap.S().Infoln("bloom filter false positive rate: ",config.FalsePositiveRate)
	zap.S().Infoln("merkle tree accumulator level in DLT: ", config.MtLevelInDLT)
	zap.S().Infoln("merkle tree height: ", config.MTHeight)
	zap.S().Infoln("revocation batch size: ", config.RevocationBatchSize)
	zap.L().Info("********************************************************************************************************************************\n")
	zap.L().Info("--------------------------------------------------------printing holder configuration--------------------------------------------------")
	zap.L().Info("holder name:"+config.HolderName)
	zap.L().Info("holder address:"+config.HolderAddress)
	zap.L().Info("********************************************************************************************************************************\n")
	zap.L().Info("--------------------------------------------------------printing verifier configuration--------------------------------------------------")
	zap.L().Info("verifier name:"+config.VerifierName)
	zap.L().Info("verifier address:"+config.VerifierAddress)
	zap.L().Info("********************************************************************************************************************************\n")

	zap.L().Info("\n\n--------------------------------------------------------printing Experiment parameters--------------------------------------------------")
	zap.L().Info("experiment manager address:"+config.ManagerAddress)
	for key, exp := range config.ExpParamters{
		zap.S().Infoln(key, *exp)
	}
	zap.L().Info("********************************************************************************************************************************\n")

}

/*
setupConfig sets up the config file, config file type and config file path
*/
func setupConfig(){
	viper.SetConfigFile("config.json")// name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
}

func ParseConfig() (Config, error){


	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("contractAddress")
	config := Config{}
	err := viper.ReadInConfig()
	if err!=nil{
		zap.S().Fatalln("error reading config file for viper.\t", err)
	}

	config.SmartContractAddress = viper.GetString("contractAddress")

	viper.SetConfigFile("config.json")// name of config file (without extension)
	err = viper.MergeInConfig()
	err = viper.ReadInConfig()
	if err!=nil{
		zap.S().Fatalln("error reading config.json file for viper.\t", err)
	}
	config.GasLimit = viper.GetUint64("contract.gasLimit")
	config.GasPrice = big.NewInt(int64(viper.GetUint64("contract.gasPrice")))
	config.BlockchainRpcEndpoint = viper.GetString("blockchain.rpcEndpoint")
	config.BlockchainWebSocketEndPoint = viper.GetString("blockchain.wsEndPoint")
	config.SenderAddress = viper.GetString("account.address")
	config.passPhrase = viper.GetString("account.passphrase")
	config.PrivateKeys = viper.GetStringSlice("account.privateKeys")
	config.LoggerType = viper.GetString("logger.env")
	config.LoggerOutputMode = viper.GetString("logger.output")
	config.LoggerFile = viper.GetString("logger.filename")
	expParams:=viper.GetStringMap("exp")

	config.ExpParamters = make(map[string]*Experiment)
	for k, v := range expParams {
		exp:=&Experiment{}
		m:= v.(map[string]interface{})
		exp.TotalVCs, _=strconv.Atoi(m["totalvcs"].(string))
		exp.RevokedVCs, _ = strconv.Atoi(m["revokedvcs"].(string))
		exp.MtLevelInDLT, _ = strconv.Atoi(m["mtlevelindlt"].(string))
		exp.MtHeight, _ = strconv.Atoi(m["mtheight"].(string))
		exp.FalsePositiveRate, _ = strconv.ParseFloat(m["falsepositiverate"].(string), 64)
		exp.RevocationBatchSize, _ = strconv.Atoi(m["revocationbatchsize"].(string))
		config.ExpParamters[k]=exp
	}

	config.ExpectedNumberOfTotalVCs = viper.GetUint("issuer.totalVCs")
	config.ExpectedNumberofRevokedVCs = viper.GetUint("issuer.revokedVCs")
	config.FalsePositiveRate = viper.GetFloat64("issuer.falsePositiveRate")
	config.MtLevelInDLT = viper.GetUint("issuer.mtLevelInDLT")
	config.MTHeight = viper.GetUint("issuer.mtHeight")
	config.RevocationBatchSize = viper.GetUint("issuer.revocationBatchSize")
	config.IssuerName = viper.GetString("issuer.name")
	config.IssuerAddress = viper.GetString("issuer.address")

	config.HolderName = viper.GetString("holder.name")
	config.HolderAddress = viper.GetString("holder.address")

	config.VerifierName = viper.GetString("verifier.name")
	config.VerifierAddress = viper.GetString("verifier.address")
	config.ManagerAddress = viper.GetString("manager.address")
	config.DEBUG = viper.GetBool("mode.debug")
	//"account1" :  "0xB97F44Ce8dA7E824F7aBD0068F92D08438E3405A",
	//	"account2" : "0x6C3d120Ee76E635d7b221a996718a8277BeA973f",
	//	"account3" : "0xF82407B704B5FF6AB71894ec0f1d78f514c3A13A",
	//	"_comment": " three other accounts in the ganache test network"


	config.printConfig()
	return config, nil
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}