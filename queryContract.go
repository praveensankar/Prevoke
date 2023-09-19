package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// generate bidings for revocation service

//  solc --abi revocationservice.sol -o build
// solc --bin revocationservice.sol -o build
//  abigen --abi RevocationService.abi --pkg main --type RevocationService --out RevocationService.go

const key = `<<json object from keystore>>`

func QueryContract(config Config) {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	ethClient, err := ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}


	revocationService, err := NewRevocationService(common.HexToAddress(config.SmartContractAddress), ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}

	var status bool
	var field int64 = 30
	status, err = revocationService.BloomFilter(nil, big.NewInt(field))
	fmt.Printf("status at field %d is %s", field, status)


}