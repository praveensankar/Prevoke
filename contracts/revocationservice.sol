//SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;
import "./console.sol";
import "./MerkleProof.sol";

contract RevocationService{

    bool private constant DEBUG = true;

    // bloom filter
    uint public constant numberOfHashFunctions = 4;
    // const private numberofVCs = 10000;

    // // BF size is set for 10% false positive
    // const private bfSize = 50000;

    // storing bloom filters as maps
    mapping(uint256=>bool) public bloomFilter;


    /*
    merkle tree
    stores entries in level order
    root is stored at index 0.
    */
    mapping(uint => string) public merkleTree;

    string public merkleRoot;

    // stores the list of indexes present in the merkle tree.
    uint[] private indexes;
    mapping (uint => bool) public isExistInMTAccumulator;

    // issuer is the owner of the contract
    address issuer;


    event Issue(uint[]  _mtIndexes, bytes1  _mtValue1, bytes1  _mtValue2, bytes1  _mtValue3, bytes1  _mtValue4);
    event VerificationPhase2(bytes32 merkleRoot, bytes32 vcLeaf, bytes32[] proof);
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


    /*
    This function is used to set proofs at merkle tree accumulator when one or more VC is issued.
    The merkle tree stores hash of VCs in leaves. Arrays using Level order structure is used to store the merkle tree.
    Every time new VC is issued, update the array.


    Note: The logic for mapping VCs to level order indexes should be done at the issuers side.
    */
    function issueVC(uint[] memory _mtIndexes, string[] memory _mtValues) public{
        //only issuer can perform the revocation
        require(msg.sender==issuer);

        updateMerkleTree(_mtIndexes, _mtValues);
    }

    /*
    Revokes a VC by updating bloom filter and merkle tree accumulator.

    Inputs:
        _bfIndexes: bloom filter indexes that needs to be set
        _mtIndexes: merkle tree indexes that needs to be updated
        _mtValues:  merkle tree values corresponding to the indexes.

    Operations:
        Bloom filter: sets the indexes at bloomfilter to 1
        merkle tree accumulator: update the merkle tree. since we only store 'k' levels, not all revocations might require update of
        merkle tree. if merkle tree doesn't require update, then index to -1

    Note: set merkle tree to -1 if it is not required to update merkle tree.
    */
    function revokeVC(uint256[numberOfHashFunctions] memory _bfIndexes, uint[] memory _mtIndexes, string[] memory _mtValues) public{
        //only issuer can perform the revocation
        require(msg.sender==issuer);

        updateBloomFilter(_bfIndexes);
        updateMerkleTree(_mtIndexes, _mtValues);

    }



    function updateBloomFilter(uint256[numberOfHashFunctions] memory _indexes) public{

        //only issuer can perform the revocation
        require(msg.sender==issuer);

        // sets the indexes to true
        for (uint i = 0; i < _indexes.length; i++) {
            bloomFilter[_indexes[i]] = true;
        }
    }


    // updates the merkle tree values in the speicified indexes
    function updateMerkleTree(uint[] memory _indexes, string[] memory _values) public {

        //only own can replace merkle tree
        require(msg.sender==issuer);

        // indexes and values should have same number of entries
        require(_indexes.length==_values.length);


        // update the merkle tree
        for (uint i = 0; i < _indexes.length; i++) {
            merkleTree[_indexes[i]] = _values[i];

        }


        for (uint i = 0; i < _indexes.length; i++) {
            if (isExistInMTAccumulator[_indexes[i]] == false){
                isExistInMTAccumulator[_indexes[i]] = true;
                indexes.push(_indexes[i]);
            }
        }
        merkleRoot = merkleTree[0];
//        emit Issue(indexes, _values[0][0], _values[0][1], _values[0][2], _values[0][3]);

    }





    /*
    verifies a VC by using two phase verification approach.

    Inputs:
        _bfIndexes: bloom filter indexes that needs to be set
        vcLeaf: leaf value of the corresponding VC in the merkle tree
        proof:  merkle proof

    Operations:
        Phase 1: check the revocation status in bloom filter. if bloom filter returns false, then it means that VC might be revoked.
        phase 2:  check the merkle tree accumulator. if it returns true then it means the proof is valid otherwise invalid.

    Returns:
        True: indicates VC is valid
        False: indicates VC is revoked
    */
    // function verifyVC(uint256[numberOfHashFunctions] memory _bfIndexes) public view returns(bool){

    //     bool statusInBloomFilter = checkRevocationStatusInBloomFilter(_bfIndexes);

    //     if(statusInBloomFilter==true){
    //         return true;
    //     }else{
    //         return checkRevocationStatusInMerkleTreeAccumulator(vcLeaf, proof);
    //     }
    // }

    /*
    verifies a VC by using only bloom filters.

    Inputs:
        vcLeaf: leaf value of the corresponding VC in the merkle tree
        proof:  merkle proof

    Operations:
         phase 2:  check the merkle tree accumulator. if it returns true then it means the proof is valid otherwise invalid.

    Returns:
        True: indicates VC is valid
        False: indicates VC is revoked
    */
    function verificationPhase1(uint256[numberOfHashFunctions] memory _bfIndexes) public view returns(bool){
        return checkRevocationStatusInBloomFilter(_bfIndexes);
    }



    // if it returns true then the VC is not revoked.
    // if it retuns false then the VC is probably revoked.
    function checkRevocationStatusInBloomFilter(uint256[numberOfHashFunctions] memory _indexes) public view returns(bool){

        bool isValid = false;
        for (uint i = 0; i < _indexes.length; i++) {
            if(bloomFilter[_indexes[i]]==false){
                isValid = true;
                break;
            }
        }
        return isValid;
    }


    /*
    Returns the merkle root.
    */
    function verificationPhase2() public view returns(string memory){
        return merkleTree[0];
    }

    function verificationPhase2Test() public returns(string memory){
        return merkleTree[0];
    }





    // updates a non-leaf or leaf node at the specified index
    function updateNode(uint index, string memory value) public{
        //only own can replace merkle tree
        require(msg.sender==issuer);
        merkleTree[index] = value;
    }

    // prints the tree in console
    function printMerkleTree() public view{

        if (DEBUG==true){
            console.log("priting merkle tree");
        }
        for (uint i = 0; i < indexes.length; i++) {
            if (DEBUG==true){
//                console.log("index : %d \t value :",i);
//                console.logBytes(abi.encodePacked(merkleTree[i]));
                console.log("index : %d \t value : %s",i, merkleTree[i]);
            }

        }
    }

    function bytes32tostring(bytes32 _bytes32) private returns(string memory){
        bytes memory bytesArray = new bytes(32);
        for (uint256 i; i < 32; i++) {
            bytesArray[i] = _bytes32[i];
        }
        return string(bytesArray);
    }







}