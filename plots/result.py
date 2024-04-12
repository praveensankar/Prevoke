import json

from setting import Setting


class Result:
    def __init__(self, bloomFilterSize, bloomFilterIndexesPerEntry, merkTreeSizeTotal,
                 merkTreeSizeInDLT, merkleTreeNodesCountInDLT, numberOfActualFalsePositives, numberOfVCsRetrievedWitnessFromIssuer,
                 numberOfVCsRetrievedWitnessFromDLT, avgRevocationTimePerVC, revocationTimePerVCRawData, verificationTimePerValidVC, verificationTimePerValidVCRawData,
                 verificationTimePerFalsePositiveOrRevokedVC, verificationTimePerFalsePositiveOrRevokedVCRawData,
                 avgTimeToFetchWitnessFromIssuer, avgTimeToFetchWitnessFromIssuerRawData, avgTimeToFetchWitnessFromDLT,
                 avgTimeToFetchWitnessFromDLTRawData, bbsProofGenerationTime, bbsVerificationTime, contractDeploymentCost,
                 bulkIssuanceCost, avgRevocationCostInGas, avgRevocationCostInGasRawData):

        self.bloomFilterSize = bloomFilterSize
        self.bloomFilterIndexesPerEntry = bloomFilterIndexesPerEntry
        self.merkTreeSizeTotal = merkTreeSizeTotal
        self.merkTreeSizeInDLT = merkTreeSizeInDLT
        self.merkleTreeNodesCountInDLT = merkleTreeNodesCountInDLT

        self.numberOfActualFalsePositives = numberOfActualFalsePositives
        self.numberOfVCsRetrievedWitnessFromIssuer = numberOfVCsRetrievedWitnessFromIssuer
        self.numberOfVCsRetrievedWitnessFromDLT = numberOfVCsRetrievedWitnessFromDLT
        self.avgRevocationTimePerVC = avgRevocationTimePerVC
        self.revocationTimePerVCRawData = revocationTimePerVCRawData
        self.verificationTimePerValidVC = verificationTimePerValidVC
        self.verificationTimePerValidVCRawData = verificationTimePerValidVCRawData
        self.verificationTimePerFalsePositiveOrRevokedVC = verificationTimePerFalsePositiveOrRevokedVC
        self.verificationTimePerFalsePositiveOrRevokedVCRawData = verificationTimePerFalsePositiveOrRevokedVCRawData
        self.avgTimeToFetchWitnessFromIssuer = avgTimeToFetchWitnessFromIssuer
        self.avgTimeToFetchWitnessFromIssuerRawData = avgTimeToFetchWitnessFromIssuerRawData
        self.avgTimeToFetchWitnessFromDLT = avgTimeToFetchWitnessFromDLT
        self.avgTimeToFetchWitnessFromDLTRawData = avgTimeToFetchWitnessFromDLTRawData
        self.bbsProofGenerationTime = bbsProofGenerationTime
        self.bbsVerificationTime = bbsVerificationTime

        self.contractDeploymentCost = contractDeploymentCost
        self.bulkIssuanceCost = bulkIssuanceCost
        self.avgRevocationCostInGas = avgRevocationCostInGas
        self.avgRevocationCostInGasRawData = avgRevocationCostInGasRawData



    def __str__(self):
        output = "bloom filter size (in bytes): " + str(self.bloomFilterSize)
        output = "\n bloomFilter Indexes Per Entry: "+str(self.bloomFilterIndexesPerEntry)
        output += "\n merkleTree Size Total (in bytes): "+str(self.merkTreeSizeTotal)
        output += "\n merkTree Size In DLT: "+str(self.merkTreeSizeInDLT)
        output += "\n merkleTree Nodes Count In DLT: "+ str(self.merkleTreeNodesCountInDLT)
        output += "\n number Of False Positives: " + str(self.numberOfActualFalsePositives)
        output += "\n number Of VCs Retrieved Witness From Issuer: " + str(self.numberOfVCsRetrievedWitnessFromIssuer)
        output += "\n number Of VCs Retrieved Witness From DLT: " + str(self.numberOfVCsRetrievedWitnessFromDLT)
        output += "\n avg Revocation Time Per VC: " + str(self.avgRevocationTimePerVC)
        output += "\n avg verification TimePer Valid VC: " + str(self.verificationTimePerValidVC)
        output += "\n avg verification TimePer False Positive Or Revoked VC: " + str(self.verificationTimePerFalsePositiveOrRevokedVC)
        output += "\n avg Time To Fetch Witness From Issuer: " + str(self.avgTimeToFetchWitnessFromIssuer)
        output += "\n avg Time To Fetch Witness From DLT: " + str(self.avgTimeToFetchWitnessFromDLT)
        output += "\n avg bbs Proof Generation Time: " + str(self.bbsProofGenerationTime)
        output += "\n avg bbs Proof verification Time: " + str(self.bbsVerificationTime)
        output += "\n contract Deployment Cost (in gas): " + str(self.contractDeploymentCost)
        output += "\n bulk issuance cost (in gas): " + str(self.bulkIssuanceCost)
        output += "\n avg revocation cost (in gas): " + str(self.avgRevocationCostInGas)

        return output


