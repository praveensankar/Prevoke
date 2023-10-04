package contracts

/* generate bidings for revocation service

solc --abi revocationservice.sol -o build
solc --bin revocationservice.sol -o build
abigen --abi RevocationService.abi --pkg contracts --type RevocationService --out RevocationService.go

*/

/*

We can automate these steps using solgo in future.


SolGo - Solidity parser built using the Go programming language
SolGo is equipped to parse contract definitions and generate the Application Binary Interface (ABI) for
individual contracts or groups of contracts
 */


