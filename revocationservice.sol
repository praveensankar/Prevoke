//SPDX-License-Identifier: MIT

pragma solidity >=0.4.16 <0.9.0;
import "hardhat/console.sol";

contract RevocationService{
    
    
    // bloom filter
    uint private constant numberOfHashFunctions = 3;
    // const private numberofVCs = 10000;
    
    // // BF size is set for 10% false positive
    // const private bfSize = 50000;

    // storing bloom filters as maps
    mapping(uint256=>bool) public bloomFilter;

    // merkle tree accumulator

    // issuer is the owner of the contract
    address public issuer;

    // sets the issuer - contract creator is the issuer
    constructor(){
        issuer = msg.sender;
    }

    /*
    This function is used to register new issuers.
    Register did of issuers and public keys. (maybe in the form of DID Docs).
    input: did doc
    */
    function registerIssuers() public{

    }



    function revokeInBloomFilter(uint256[numberOfHashFunctions] memory indexes) public{

        //only issuer can perform the revocation
        require(msg.sender==issuer);
        
        // sets the indexes to true
        for (uint i = 0; i < indexes.length; i++) {
            bloomFilter[indexes[i]] = true; 
        }
    }

    // if it returns true then the VC is probably revoked. go to the phase 2.
    function checkRevocationStatusInBloomFilter(uint256[numberOfHashFunctions] memory indexes) public view returns(bool){

        bool isRevoked = false;
        for (uint i = 0; i < indexes.length; i++) {
            if(bloomFilter[indexes[i]]==true){
                isRevoked = true;
                break;
            }
        }
        return isRevoked;
    }   

    function testRevocation() public{
       uint256[numberOfHashFunctions] memory vc1 = [uint256(1),uint256(2),uint256(3)];
       uint256[numberOfHashFunctions] memory vc2 = [uint256(4),uint256(5),uint256(6)];
       uint256[numberOfHashFunctions] memory vc3 = [uint256(7),uint256(8),uint256(9)];
       revokeInBloomFilter(vc1);
       revokeInBloomFilter(vc2);
       string memory vc1Status = checkRevocationStatusInBloomFilter(vc1) ? "revoked" : "not revoked";
       string memory vc2Status = checkRevocationStatusInBloomFilter(vc2) ? "revoked" : "not revoked";
       string memory vc3Status = checkRevocationStatusInBloomFilter(vc3) ? "revoked" : "not revoked";
        console.log(
            "vc1 status : %s \n vc2 status : %s\n vc3 status : %s", vc1Status, vc2Status, vc3Status
        );
    }

   


}