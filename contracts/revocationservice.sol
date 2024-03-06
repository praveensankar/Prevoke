//SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;
import "./console.sol";

contract RevocationService{

    bool private constant DEBUG = true;


    // storing bloom filters as maps
    mapping(uint256=>uint256) public bloomFilter;

    /*
    merkle tree
    stores entries in level order
    root is stored at index 0.
    */
    mapping(uint => bytes32) public merkleTree;

    // stores the list of indexes present in the merkle tree.
    uint[] private indexes;
    mapping (uint => bool) private isExistInMTAccumulator;

    // entities is the owner of the contract
    address issuer;

    // stores the public keys used by the entities
    bytes[] public publicKeys;
    mapping (string => bool) private isExistInPublicKeys;

    // sets the entities - contract creator is the entities
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
    This function adds public keys used by the entities.

    input: public keys
    */
    function addPublicKeys(bytes[] memory _publicKeys) public{
        //only entities can perform the revocation
        require(msg.sender==issuer);


        for (uint i = 0; i < _publicKeys.length; i++) {
                publicKeys.push(_publicKeys[i]);
        }
    }

    /*
This function returns the public keys used by the entities

output: public keys
*/
    function RetrievePublicKeys() public  view returns (bytes[] memory){
        return publicKeys;
    }


    /*
    This function is used to set proofs at merkle tree accumulator when one or more VC is issued.
    The merkle tree stores hash of VCs in leaves. Arrays using Level order structure is used to store the merkle tree.
    Every time new VC is issued, update the array.


    Note: The logic for mapping VCs to level order indexes should be done at the issuers side.
    */
    function issueVC(uint[] memory _mtIndexes, bytes32[] memory _mtValues) public{
        //only entities can perform the revocation
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
    function revokeVC(uint256[] memory _bfIndexes, uint[] memory _mtIndexes, bytes32[] memory _mtValues) public{
        //only entities can perform the revocation
        require(msg.sender==issuer);

        updateBloomFilter(_bfIndexes);
        updateMerkleTree(_mtIndexes, _mtValues);

    }



    function updateBloomFilter(uint256[] memory _indexes) public{

        //only entities can perform the revocation
        require(msg.sender==issuer);

        // sets the bits at the specified indexes
        for (uint i = 0; i < _indexes.length; i++) {
            uint256 bucket = _indexes[i] >> 8;
            uint256 mask = 1 << (_indexes[i] & 0xff);
            bloomFilter[bucket] |= mask;
        }
    }


    // updates the merkle tree values in the speicified indexes
    function updateMerkleTree(uint[] memory _indexes, bytes32[] memory _values) public {

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
    }




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
    function verificationPhase1(uint256[] memory _bfIndexes) public view returns(bool){
        return checkRevocationStatusInBloomFilter(_bfIndexes);
    }



    // if it returns true then the VC is not revoked.
    // if it retuns false then the VC is probably revoked.
    function checkRevocationStatusInBloomFilter(uint256[] memory _indexes) public view returns(bool){

        bool isValid = false;
        for (uint i = 0; i < _indexes.length; i++) {
                uint256 bucket = _indexes[i] >> 8;
                uint256 mask = 1 << (_indexes[i] & 0xff);
                bool status = bloomFilter[bucket] & mask != 0;
                if(status==false){
                isValid = true;
                break;
            }
        }
        return isValid;
    }


    /*
    Returns the merkle root.
    */
    function verificationPhase2() public view returns(bytes32){
        return merkleTree[0];
    }

    function verificationPhase2Test() public view returns(bytes32){
        return merkleTree[0];
    }





    // updates a non-leaf or leaf node at the specified index
    function updateNode(uint index, bytes32 value) public{
        //only own can replace merkle tree
        require(msg.sender==issuer);
        merkleTree[index] = value;
    }


    function RetrieveMerkleTree() public view returns (bytes32[] memory){
        bytes32 [] memory mt = new bytes32[](indexes.length);
        for (uint i = 0; i < indexes.length; i++) {
            mt[i] = merkleTree[i];
        }
        return mt;
    }


    function GetMerkleTreeSize() public view returns (uint)  {
        uint mtSize;
        mtSize=0;
        for (uint i = 0; i < indexes.length; i++) {
            mtSize = mtSize + merkleTree[i].length;
        }
        return mtSize;
    }

 

        // prints the tree in console
    function printMerkleTree() public view{

        if (DEBUG==true){
            console.log("priting merkle tree");
        }
        for (uint i = 0; i < indexes.length; i++) {
            if (DEBUG==true){
                               console.log("index : %d \t value :",i);
                               console.logBytes(abi.encodePacked(merkleTree[i]));
                // console.log("index : %d \t value : %s",i, merkleTree[i]);
            }

        }
    }

    function test() public{
        uint[] memory mtIndexes = new uint[](2);
        mtIndexes[0]=0;
        mtIndexes[1]=1;
        bytes32[] memory values = new bytes32[](2);
        values[0] = 0x68656c6c6f000000000000000000000000000000000000000000000000000000;
        values[1] = 0x68656c6c6f000000000000000000000000000000000000000000000000000000;
        issueVC(mtIndexes, values);
        printMerkleTree();
        console.log("merkle tree size: %d ",GetMerkleTreeSize());
    }
}




