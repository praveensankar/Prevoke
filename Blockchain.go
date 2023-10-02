package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math"
	"math/big"
	"sync"
)

type BlockchainNode struct {
	rpcEndPoint string
}

func (node BlockchainNode) setRpcEndPoint(rpcEndPoint string)  {
	node.rpcEndPoint = rpcEndPoint;
}



func (node BlockchainNode) connect() {
	// step 1: connect to a blockchain node using RPC endpoint
	_, err := ethclient.Dial(node.rpcEndPoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	} else {
		fmt.Println("connected to blockchain node")
	}
}




func (node BlockchainNode) getBalance(account string){
	// step 1: connect to a blockchain node using RPC endpoint
	ethClient, err := ethclient.Dial(node.rpcEndPoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	} else {
		fmt.Println("connected to blockchain node")
	}

	// step 2: convert smart contract to the requied format
	address := common.HexToAddress(account)
	node.checkAccountType(address.String())
	balance, err := ethClient.BalanceAt(context.Background(), address, nil)
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	balanceInETH := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Printf("%s has balance of %f eth", address.String(), balanceInETH)
}

func (node BlockchainNode) checkAccountType(account string){
	ethClient, err := ethclient.Dial(node.rpcEndPoint)
	address := common.HexToAddress(account)
	bytecode, err := ethClient.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	if isContract == true{
		fmt.Println(account, "is a smart contract")
	} else {
		fmt.Println(account, "is a wallet")
	}

}

func (node BlockchainNode) queryBlock() {
	ethClient, err := ethclient.Dial(node.rpcEndPoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	latestBlock, err := ethClient.BlockByNumber(context.Background(), nil)
	fmt.Println("......printing latest block information.......")
	fmt.Println("block number : ", latestBlock.Number().Uint64())
	fmt.Println("block time : ", latestBlock.Time())
	fmt.Println("block difficulty : ", latestBlock.Difficulty().Uint64())
	fmt.Println("block hash : ", latestBlock.Hash().Hex())
	fmt.Println("number of transactions :", len(latestBlock.Transactions()))
	fmt.Println("transactions : ")
	for _, tx := range latestBlock.Transactions() {
		fmt.Printf("hash : %s \t gas : %d \t gas price : %d \t nonce : %d \n ",tx.Hash().Hex(),
			tx.Gas(), tx.GasPrice().Uint64(), tx.Nonce())
	}
}

func createNewKeyStore() (*keystore.KeyStore)  {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	_, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nnew keystored with the following account: ",ks.Accounts())
	return ks
}


func (node BlockchainNode) doTransaction(privateKeyInString string){
	privateKey, err := crypto.HexToECDSA(privateKeyInString)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	ethClient, err := ethclient.Dial(node.rpcEndPoint)
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d is nonce for the next transaction for the public key : %x \n",nonce, fromAddress)
}

func (node BlockchainNode) sendLegacyTransaction(privateKeyInString string,  toAddressInString string){
	client, err :=  ethclient.Dial(node.rpcEndPoint)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyInString)
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

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(toAddressInString)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func (node BlockchainNode) sendTransaction()  {
	var (
		sk       = crypto.ToECDSAUnsafe(common.FromHex("285e18d537f6da98100a685816f5aa41c20f716d2107bb7fcb662d314355b086"))
		to       = common.HexToAddress("0x0C560e0FE1E6e9d4f73858cb5cF83B377898b10e")
		value    = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))
		sender   = common.HexToAddress("0xA59d0677384A9Fab7dd8eE18421b6A0601968911")
		gasLimit = uint64(6721975)
	)

	cl, err := ethclient.Dial("https://rpc.sepolia.dev")
	// Retrieve the chainid (needed for signer)
	chainid, err := cl.ChainID(context.Background())

	// Retrieve the pending nonce
	nonce, err := cl.PendingNonceAt(context.Background(), sender)

	// Get suggested gas price
	tipCap, _ := cl.SuggestGasTipCap(context.Background())
	feeCap, _ := cl.SuggestGasPrice(context.Background())
	// Create a new transaction
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainid,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     value,
			Data:      nil,
		})
	// Sign the transaction using our keys
	signedTx, _ := types.SignTx(tx, types.NewLondonSigner(chainid), sk)
	// Send the transaction to our node
	err = cl.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %s", signedTx.Hash().Hex())
}


func (node BlockchainNode) listenForNewBlocks(){
	client, err :=  ethclient.Dial(node.rpcEndPoint)
	if err != nil {
		log.Fatal("error connecting to blockchain",err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal("error subscribing to node \t",err)
	}
	fmt.Printf("listening for new blocks:")
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("new block received- \t block number : %d",block.Number())
		}
	}
}


func testConnectionToBlockchain(config Config){

	// set up blockchain node
	node := BlockchainNode{rpcEndPoint: config.BlockchainRpcEndpoint}

	// get balance
	node.getBalance(config.SmartContractAddress)

	node.queryBlock()
	node.doTransaction(config.privateKey)
	//node.sendEther(config.privateKey,1, config.otherAccounts[0])
	//node.sendTransaction()
	node.sendLegacyTransaction(config.privateKey, config.otherAccounts[0])
	_ = createNewKeyStore()

	var waitGroupforBlocksListener sync.WaitGroup

	waitGroupforBlocksListener.Add(1)
	go func(node BlockchainNode) {
		defer waitGroupforBlocksListener.Done()
		node.listenForNewBlocks()
	}(node)


	waitGroupforBlocksListener.Wait()
}

