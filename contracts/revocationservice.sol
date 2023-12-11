//SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;
import "./console.sol";
import "./MerkleProof.sol";

contract RevocationService{

    bool private constant DEBUG = true;

    // bloom filter
    uint public constant numberOfHashFunctions = 3;
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
    mapping(uint => bytes32) public merkleTree;

    bytes32 public merkleRoot; 

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
    function issueVC(uint[] memory _mtIndexes, bytes32[] memory _mtValues) public{
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
    function revokeVC(uint256[numberOfHashFunctions] memory _bfIndexes, uint[] memory _mtIndexes, bytes32[] memory _mtValues) public{
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
        merkleRoot = merkleTree[0];
        emit Issue(indexes, _values[0][0], _values[0][1], _values[0][2], _values[0][3]);

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

        bool isValid = true;
        for (uint i = 0; i < _indexes.length; i++) {
            if(bloomFilter[_indexes[i]]==true){
                isValid = false;
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



    /*
    verifies a VC by using only merkle tree accumulator.

    Inputs:
        _bfIndexes: bloom filter indexes that needs to be set

    Operations:
        Phase 1: check the revocation status in bloom filter. if bloom filter returns false, then it means that VC might be revoked.

    Returns:
        True: indicates VC is valid
        False: indicates VC is revoked
    */
    function verificationPhase2Old(bytes32 vcLeaf, bytes32[] memory proof) public view returns(bool){
        return checkRevocationStatusInMerkleTreeAccumulator(vcLeaf, proof);
    }

    function verificationPhase2TestOld(bytes32 vcLeaf, bytes32[] memory proof) public returns(bool){
        emit VerificationPhase2(merkleTree[0], vcLeaf, proof);
        return checkRevocationStatusInMerkleTreeAccumulator(vcLeaf, proof);
    }

    /*
    check whether  the leaf value of the VC is present in the merkle tree. verify whether the proof is valid is or not not.
    If proof is valid then it means that VC is present in the merkle tree. therefore, it is not revoked.
    If proof is invalid then it means that VC is not present in the merkle tree. therefore, it is revoked.

    returns
    True -  if VC is not revoked.
    False - if VC is revoked
    */
    function checkRevocationStatusInMerkleTreeAccumulator(bytes32 leaf, bytes32[] memory proof) public view returns (bool){
        
        bool status = MerkleProof.verify(proof, merkleTree[0], leaf);
        // if (DEBUG==true){
        //     console.log("verification status of merkle proof: ", status);
        // }

        return status;
    }




    // updates a non-leaf or leaf node at the specified index
    function updateNode(uint index, bytes32 value) public{
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
                console.log("index : %d \t value :",i);
                console.logBytes(abi.encodePacked(merkleTree[i]));
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


    function testRevocation() public{

        uint[numberOfHashFunctions] memory vc1BloomFilter;

        for (uint i = 1; i < numberOfHashFunctions; i++){
            vc1BloomFilter[i] = i;
        }
        // uint256[numberOfHashFunctions] memory vc2BloomFilter = [uint256(4),uint256(5),uint256(6)];
        // uint256[numberOfHashFunctions] memory vc3BloomFilter = [uint256(7),uint256(8),uint256(9)];
        // uint256[numberOfHashFunctions] memory vc4BloomFilter = [uint256(10),uint256(11),uint256(12)];



        // mapping(uint => bytes32) memory tree;
        bytes32 vc1MTAcc = keccak256(abi.encode("vc1"));
        bytes32 vc2MTAcc = keccak256(abi.encode("vc2"));
        bytes32 vc3MTAcc = keccak256(abi.encode("vc3"));
        bytes32 vc4MTAcc = keccak256(abi.encode("vc4"));

        bytes32 internal1 = keccak256(abi.encodePacked(vc1MTAcc, vc2MTAcc));
        bytes32 internal2 = keccak256(abi.encodePacked(vc3MTAcc, vc4MTAcc));

        bytes32 root = keccak256(abi.encodePacked(internal1, internal2));

        {
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
            values[3] = vc1MTAcc;
            values[4] = vc2MTAcc;
            values[5] = vc3MTAcc;
            values[6] = vc4MTAcc;
            if (DEBUG==true){
                console.log("issuing VCs: VC1, VC2, VC3, VC34");
                console.log("merkle tree accumulator is initialized with valid vcs");
            }
            issueVC(myindexes, values);
        }


        bytes32[] memory proofForVC1 = new bytes32[](2);
        proofForVC1[0] = vc2MTAcc;
        proofForVC1[1] = internal2;



        {


            if(DEBUG==true){
                console.log("now verifiying vc1 ");
            }
            // bool statusOfVC1BeforeRevoking = verifyVC(vc1BloomFilter, vc1MTAcc, proofForVC1);
            // string memory vc1StatusBeforeRevoking = statusOfVC1BeforeRevoking ? "not revoked" : "revoked";
            if (DEBUG==true){
                // console.log("the revocation status of VC1: ",vc1StatusBeforeRevoking);

                bool statusOfVC1BeforeRevoking = verificationPhase1(vc1BloomFilter);
                string memory vc1StatusBeforeRevoking = statusOfVC1BeforeRevoking ? "not revoked" : "probably revoked";
                if (DEBUG==true){
                    console.log("phase 1: revocation status of VC1: ",vc1StatusBeforeRevoking);
                }

                // if(DEBUG==true){
                //     console.log("merkle leaf :");
                //     console.logBytes(abi.encodePacked(vc1MTAcc));
                //     console.log("merkle proof: ");
                //     for (uint i=0; i<2;i++)
                //     {
                //         console.logBytes(abi.encodePacked(proofForVC1[i]));
                //     }
                // }
                statusOfVC1BeforeRevoking = verificationPhase2Old(vc1MTAcc, proofForVC1);
                vc1StatusBeforeRevoking = statusOfVC1BeforeRevoking ? "not revoked" : " revoked";
                if (DEBUG==true){
                    console.log("phase 2: revocation status of VC1: ",vc1StatusBeforeRevoking);
                }
            }
        }


        {
            if (DEBUG==true){
                console.log("now revoking VC1......");
            }
            bytes32 revokedvc1MTAcc = keccak256(abi.encode("vc 1 revoked"));
            internal1 = keccak256(abi.encodePacked(revokedvc1MTAcc, vc2MTAcc));
            root = keccak256(abi.encodePacked(internal1, internal2));
            uint[] memory mtIndexesAfterRevocation = new uint[](3);
            bytes32[] memory mtValuesAfterRevocation = new bytes32[](3);
            mtIndexesAfterRevocation[0] = uint(0);
            mtValuesAfterRevocation[0] = root;
            mtIndexesAfterRevocation[1] = uint(1);
            mtValuesAfterRevocation[1] = internal1;
            mtIndexesAfterRevocation[2] = uint(3);
            mtValuesAfterRevocation[2] = revokedvc1MTAcc;

            revokeVC(vc1BloomFilter,mtIndexesAfterRevocation, mtValuesAfterRevocation);
        }

        // bool statusOfVC1 = verifyVC(vc1BloomFilter, vc1MTAcc, proofForVC1);
        // string memory vc1Status = statusOfVC1 ? "not revoked" : "revoked";
        // if (DEBUG==true){
        //     console.log("the revocation status of VC1: ",vc1Status);
        // }


        bool statusOfVC1InPhase1 = verificationPhase1(vc1BloomFilter);
        string memory vc1Status = statusOfVC1InPhase1 ? "not revoked" : "probably revoked";
        if (DEBUG==true){
            console.log("phase 1: revocation status of VC1: ",vc1Status);
        }

        bool statusOfVC1InPhase2 = verificationPhase2Old(vc1MTAcc, proofForVC1);
        vc1Status = statusOfVC1InPhase2 ? "not revoked" : " revoked";
        if (DEBUG==true){
            console.log("phase 2: revocation status of VC1: ",vc1Status);
        }

        
        bytes32 mtRoot = verificationPhase2();
        if (DEBUG==true){
            console.log("merkle root: ", bytes32tostring(mtRoot));
        }




    }




}