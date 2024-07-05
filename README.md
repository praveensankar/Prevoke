# Prevoke- Revocation technique

_**Prevoke**_ is a privacy-preserving revocation technique for Verifiable Credentials. _**Prevoke**_ addresses the following privacy requirements:

1. The association between holders and verifiers should not be learned by issuers
2. Third parties should not learn the identities of holders whose VC got revoked, just by observing the registry
3. Revocation techniques should not leak the revocation rate

Prevoke keeps track of the revocation status of VCs using a Bloom Filter, and a Merkle Tree Accumulator (MTAcc). These data structures are stored in a DLT using a Smart Contract. **Bloom Filter stores only revoked VCs, whereas MTAcc stores only valid VCs**. _Bloom Filter is updated only during revocation, whereas the MTAcc is updated during both issuance and revocation_.

The verification of a VC consists of two phases:

* In phase 1, the inclusion of the VC in the Bloom Filter is checked. If the VC is included, phase 2 is necessary.
* In phase 2, the MTAcc is checked. For most valid VCs, phase 2 is not required since the Bloom Filter never results in false negatives.

Verifying most of the valid VCs takes constant time. Moreover, phase 2 is only required for a handful of valid VCs (false positives), and for revoked VCs. Holders retrieve the updated witnesses for MTAccs only for these cases. In addition, holders selectively disclose Bloom Filter indexes and a MTAcc proof using BBS proof, and thus, the complete VC is never shared.



## Proof of concept Implementation

This repository implements the proof of concept of Prevoke in a sandboxed environment. This implementation consists of the following workflows:

* First, Issuers issue Verifiable Credentials (VCs) to Holders
* Second, Issuers revoke VCs based on the revocation strategy
* Third, Holders create Verifiable Presentations (VPs) based on VCs and send them to Verifiers
* Fourth, Verifiers verify the revocation status of the underlying VCs corresponding to VPs
* Verifiable Credentials data model

# Table of contents

# Table of contents


* [Tutorial - Verifiable Credentials](docs/tutorial-verifiable-credentials/README.md)
  * [BBS Signature](docs/tutorial-verifiable-credentials/bbs-signature.md)
  * [Complete Verifiable Credential](docs/tutorial-verifiable-credentials/complete-verifiable-credential.md)

