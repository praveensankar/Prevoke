package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"sync"
	"time"
)


const key = `<<json object from keystore>>`



func ReadFromContract(config Config) {

	// step 1: connect to a blockchain node using RPC endpoint
	ethClient, err := ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// step 2: convert smart contract to the requied format
	revocationServiceSmartContract := common.HexToAddress(config.SmartContractAddress)

	balance, err := ethClient.BalanceAt(context.Background(), revocationServiceSmartContract, nil)

	fmt.Printf("%s has %d balance", revocationServiceSmartContract.String(), balance)

	revocationService, err := NewRevocationService(revocationServiceSmartContract, ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}

	var status bool
	var field int64 = 30
	status, err = revocationService.BloomFilter(nil, big.NewInt(field))
	fmt.Printf("\n status at field %d is %t", field, status)


}


func WriteToContract(config Config){
	client, err :=  ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(config.privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}


	gasLimit := uint64(6721975)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	revocationServiceSmartContract := common.HexToAddress(config.SmartContractAddress)
	revocationService, err := NewRevocationService(revocationServiceSmartContract, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}


	index1 := big.NewInt(505)
	index2 := big.NewInt(605)
	index3 := big.NewInt(705)

	indexes := [3]*big.Int{index1, index2, index3}

	status, err := revocationService.CheckRevocationStatusInBloomFilter(nil, indexes)
	fmt.Printf("\n revocation status[%d, %d, %d]: %t", index1, index2, index3, status)

	fmt.Println("\n Revoking in Bloom Filter")
	tx, err :=revocationService.RevokeInBloomFilter(auth, indexes)
	if err != nil {
		log.Fatal("failed to revoke", err)
	}
	fmt.Printf("tx hash: %s\n", tx.Hash().Hex())

	status, err = revocationService.CheckRevocationStatusInBloomFilter(nil, indexes)
	fmt.Printf("\n revocation status[%d, %d, %d]: %t", index1, index2, index3, status)

}


func SubscribeToEvents(config Config){
	client, err :=  ethclient.Dial(config.BlockchainWebSocketEndPoint)
	if err != nil {
		log.Fatal(err)
	}

	revocationServiceSmartContract := common.HexToAddress(config.SmartContractAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{revocationServiceSmartContract},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	stopListener := time.NewTimer(5 * time.Second)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case <-stopListener.C:
			fmt.Println("\n stopping the listener")
			return
		case vLog := <-logs:
			fmt.Println("new event ",vLog) // pointer to event log
		}
	}
}


func testSmartContract(config Config){
	ReadFromContract(config)
	WriteToContract(config)


	var waitGroupforBlocksListener sync.WaitGroup

	waitGroupforBlocksListener.Add(1)
	go func(config Config) {
		defer waitGroupforBlocksListener.Done()
		SubscribeToEvents(config)
	}(config)


	waitGroupforBlocksListener.Wait()
}