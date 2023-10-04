package contracts

// generate bidings for revocation service

//  solc --abi revocationservice.sol -o build
// solc --bin revocationservice.sol -o build
//  abigen --abi RevocationService.abi --pkg contracts --type RevocationService --out RevocationService.go



/*

Why we are using solgo for generating bindings?


SolGo - Solidity parser built using the Go programming language
SolGo is equipped to parse contract definitions and generate the Application Binary Interface (ABI) for
individual contracts or groups of contracts
 */

func generateBindingsForSmartContract(contractName string){


}
