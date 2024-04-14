package blockchain

import (
	"context"
	"crypto/ecdsa"
	_ "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	common2 "github.com/Revocation-Service/config"
	"github.com/Revocation-Service/contracts"
	"go.uber.org/zap"
	"math/big"
	_ "time"
)
//
func DeployContract(config common2.Config, counter int) (string, int64, error){
	client, err :=  ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		zap.S().Fatalln("ERROR in deploying contract",err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKeys[counter])
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


	gasLimit := config.GasLimit             // in units
	_, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		zap.S().Fatalln(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = config.GasPrice
	//auth.GasPrice=big.NewInt(gasPrice)

	account := common.HexToAddress(config.SenderAddress)
	startBalance, err := client.BalanceAt(context.Background(), account, nil)
	addresss, tx, _, err  := contracts.DeployRevocationService(auth, client)

	endBalance, err := client.BalanceAt(context.Background(), account, nil)
	gasUsed := (startBalance.Int64()-endBalance.Int64()) / config.GasPrice.Int64()
	if err != nil {
		zap.S().Infof("Failed to deploy contract: %v", err)
	}

	if config.DEBUG==true {
		zap.L().Info("\n\n------------------------------------------------------- deploying smart contract --------------------------------------------------")

		zap.S().Infoln("BLOCKCHAIN- \t  smart contract address: ", addresss.String())
		zap.S().Infoln("BLOCKCHAIN- \t tx hash: ", tx.Hash())
		zap.S().Infoln("BLOCKCHAIN - \t gas used: ", gasUsed)
		zap.L().Info("********************************************************************************************************************************\n")
	}
	return addresss.String(), gasUsed, err
}
