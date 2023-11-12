const rs = artifacts.require("RevocationService");
// const merkleProof = artifacts.require("MerkleProof");


module.exports = function(deployer) {
  // deployer.deploy(merkleProof);
  // deployer.link(merkleProof, rs);
  deployer.deploy(rs);
};
