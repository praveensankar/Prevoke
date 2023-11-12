package blockchain

import (
	_ "github.com/ethereum/go-ethereum"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "time"
)
//
//func DeployContract(config config.Config){
//	client, err :=  ethclient.Dial(config.BlockchainRpcEndpoint)
//	if err != nil {
//		zap.S().Fatalln(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
//	if err != nil {
//		zap.S().Fatalln(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		zap.S().Fatalln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
//	}
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		zap.S().Fatalln(err)
//	}
//
//
//	gasLimit := uint64(6721975)                // in units
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		zap.S().Fatalln(err)
//	}
//
//	auth := bind.NewKeyedTransactor(privateKey)
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)
//	auth.GasLimit = gasLimit
//	auth.GasPrice = gasPrice
//
//	//revocationServiceSmartContract := common.HexToAddress(config.SmartContractAddress)
//
//	//revocationservice, err := NewRevocationService(revocationServiceSmartContract, client)
//
//
//	addresss, tx, revocationservice, err  := contracts.DeployRevocationService(auth, client)
//	if err != nil {
//		zap.S().Infof("Failed to deploy contract: %v", err)
//	}
//
//	zap.S().Infoln("deployed smart contract address: ", addresss.String())
//	zap.S().Infoln("tx hash: ", tx)
//	n, _ := revocationservice.NumberOfHashFunctions(nil)
//	zap.S().Infoln("number of hash functions: ", n)
//
//}
