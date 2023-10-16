//SPDX-License-Identifier: MIT

pragma solidity >=0.4.16 <0.9.0;
import "hardhat/console.sol";


contract BloomFilter{
    // bloom filter
    uint public constant numberOfHashFunctions = 3;
    // const private numberofVCs = 10000;
    
    // // BF size is set for 10% false positive
    // const private bfSize = 50000;

    // storing bloom filters as mapping
    mapping(uint256=>bool) public bloomFilter;




    function revokeInBloomFilter(uint256[numberOfHashFunctions] memory indexes) public{
        
        // sets the indexes to true
        for (uint i = 0; i < indexes.length; i++) {
            bloomFilter[indexes[i]] = true; 
        }
    }


    // if it returns true then the VC is not revoked. 
    // if it retuns false then the VC is probably revoked.
    function checkRevocationStatusInBloomFilter(uint256[numberOfHashFunctions] memory indexes) public view returns(bool){

        bool isValid = true;
        for (uint i = 0; i < indexes.length; i++) {
            if(bloomFilter[indexes[i]]==true){
                isValid = false;
                break;
            }
        }
        return isValid;
    }


    function testRevocationInBloomFilter() public{
       uint256[numberOfHashFunctions] memory vc1 = [uint256(1),uint256(2),uint256(3)];
       uint256[numberOfHashFunctions] memory vc2 = [uint256(4),uint256(5),uint256(6)];
       uint256[numberOfHashFunctions] memory vc3 = [uint256(7),uint256(8),uint256(9)];
       string memory vc1Status = checkRevocationStatusInBloomFilter(vc1) ? "not revoked" : "revoked";
       string memory vc2Status = checkRevocationStatusInBloomFilter(vc2) ? "not revoked" : "revoked";
       string memory vc3Status = checkRevocationStatusInBloomFilter(vc3) ? "not revoked" : "revoked";
        console.log(
            "before revoking:--------- vc1 status: %s ********* vc2 status: %s  *********  vc3 status: %s", vc1Status, vc2Status, vc3Status
        );
       revokeInBloomFilter(vc1);
       revokeInBloomFilter(vc2);
       vc1Status = checkRevocationStatusInBloomFilter(vc1) ? "not revoked" : "revoked";
       vc2Status = checkRevocationStatusInBloomFilter(vc2) ? "not revoked" : "revoked";
       vc3Status = checkRevocationStatusInBloomFilter(vc3) ? "not revoked" : "revoked";
        console.log(
            "revoked vc1 and vc2:------- vc1 status: %s  *********  vc2 status: %s  *********  vc3 status: %s", vc1Status, vc2Status, vc3Status
        );
    }


}