

//SPDX-License-Identifier: MIT

pragma solidity >=0.4.16 <0.9.0;


import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";


contract MerkleTreeAccumulator{



    

    /* 
    merkle tree
    stores entries in level order 
    root is stored at index 0. 
    */
    mapping(uint => bytes32) public merkleTree;

    // stores the list of indexes present in the merkle tree.
    uint[] private indexes;
    mapping (uint => bool) private isExist;

    //owner of the contract
    address public owner;

    // sets the issuer - contract creator is the issuer
    constructor(){
        owner = msg.sender;
    }

    // updates the merkle tree values in the speicified indexes
    function updateMerkleTree(uint[] memory _indexes, bytes32[] memory _values) public {
        
        //only own can replace merkle tree
        require(msg.sender==owner);

        // indexes and values should have same number of entries
        require(_indexes.length==_values.length);
        
        // update the merkle tree
        for (uint i = 0; i < _indexes.length; i++) {
        merkleTree[_indexes[i]] = _values[i];
        }

       
        for (uint i = 0; i < _indexes.length; i++) {
            if (isExist[_indexes[i]] == false){
                isExist[_indexes[i]] = true;
                indexes.push(_indexes[i]);
            }
        }
        
    }

    // updates a non-leaf or leaf node at the specified index
    function updateNode(uint index, bytes32 value) public{

        //only own can replace merkle tree
        require(msg.sender==owner);
        
        merkleTree[index] = value;
    }

    // prints the tree in console
    function printMerkleTree() public{
         for (uint i = 0; i < indexes.length; i++) {
         console.logBytes(abi.encodePacked(merkleTree[i]));
         }
    }

    // returns 
    // True - if leaf and proof are valid
    // False - if invalid
    function verifyLeaf(bytes32 leaf, bytes32[] memory proof) public returns (bool){
        return MerkleProof.verify(proof, merkleTree[0], leaf);
    }




    function testMerkleTree() public{
        
        // mapping(uint => bytes32) memory tree;
        bytes32 leaf1 = keccak256(abi.encode("leaf 1"));
        bytes32 leaf2 = keccak256(abi.encode("leaf 2"));
        bytes32 leaf3 = keccak256(abi.encode("leaf 3"));
        bytes32 leaf4 = keccak256(abi.encode("leaf 4"));

        bytes32 internal1 = keccak256(abi.encodePacked(leaf1, leaf2));
        bytes32 internal2 = keccak256(abi.encodePacked(leaf3, leaf4));

        bytes32 root = keccak256(abi.encodePacked(internal1, internal2));
        
        
        uint[] memory myindexes = new uint[](7);    
        myindexes[0] = uint(0);
        myindexes[1] = uint(1);
        myindexes[2] = uint(2);
        myindexes[3] = uint(3);
        myindexes[4] = uint(4);
        myindexes[5] = uint(5);
        myindexes[6] = uint(6);
        

        bytes32[] memory values = new bytes32[](7);   
        values[0] = root;
        values[1] = internal1;
        values[2] = internal2;
        values[3] = leaf1;
        values[4] = leaf2;
        values[5] = leaf3;
        values[6] = leaf4;
    

        
        updateMerkleTree(myindexes, values);


        bytes32 leaf = leaf1;
        bytes32[] memory validProof = new bytes32[](2);   
        validProof[0] = leaf2;
        validProof[1] = internal2;

        bytes32[] memory invalidProof = new bytes32[](2);   
        invalidProof[0] = leaf3;
        invalidProof[1] = internal2;

        bool verificationStatus = verifyLeaf(leaf, validProof);
        console.log("verification status for vaild proof: ", verificationStatus);

        verificationStatus = verifyLeaf(leaf, invalidProof);
        console.log("verification status for invaild proof: ", verificationStatus);

    }

    }


   



   

