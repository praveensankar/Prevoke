package blockchain

import (
	"context"
	"crypto/ecdsa"
	_ "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/contracts"
	"go.uber.org/zap"
	"math/big"
	_ "time"
)
//
func DeployContract(config config.Config) (string, error){
	client, err :=  ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		zap.S().Fatalln(err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		zap.S().Fatalln(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		zap.S().Fatalln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		zap.S().Fatalln(err)
	}


	gasLimit := uint64(80000000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		zap.S().Fatalln(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	auth.GasPrice=big.NewInt(int64(20000000000))



	addresss, tx, revocationservice, err  := contracts.DeployRevocationService(auth, client)
	if err != nil {
		zap.S().Infof("Failed to deploy contract: %v", err)
	}
	zap.L().Info("\n\n------------------------------------------------------- deploying smart contract --------------------------------------------------")

	zap.S().Infoln("BLOCKCHAIN- \t  smart contract address: ", addresss.String())
	zap.S().Infoln("BLOCKCHAIN- \t tx hash: ", tx.Hash())
	_, _ = revocationservice.NumberOfHashFunctions(nil)
	//zap.S().Infoln("number of hash functions: ", n)
	zap.L().Info("********************************************************************************************************************************\n")

	return addresss.String(), err
}
