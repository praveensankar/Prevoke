import json
import math
import os

import numpy as np

from result import Result
from setting import Setting


class Entry:
    def __init__(self, setting, result):
        self.setting = setting
        self.result = result

    def __str__(self):
        output = self.setting.__str__()
        output = output + "\n"+ self.result.__str__()
        return output



def calculate_average(entries):


    values = {}
    keys = {}
    for entry in entries:
        if entry.setting in values:
            setting = Setting(totalVCs=entry.setting.totalVCs, revokedVCs=entry.setting.revokedVCs, falsePositiveRate=entry.setting.falsePositiveRate,
                              mtLevelInDLT=entry.setting.mtLevelInDLT)


            value = values[entry.setting]
            print(value.result.avgRevocationCostInGasRawData)


            bloomFilterSize = (value.result.bloomFilterSize + entry.result.bloomFilterSize) / 2
            bloomFilterIndexesPerEntry = (value.result.bloomFilterIndexesPerEntry + entry.result.bloomFilterIndexesPerEntry) / 2
            merkTreeSizeTotal = (value.result.merkTreeSizeTotal + entry.result.merkTreeSizeTotal) / 2

            merkTreeSizeInDLT = (value.result.merkTreeSizeInDLT + entry.result.merkTreeSizeInDLT) / 2
            merkleTreeNodesCountInDLT = (value.result.merkleTreeNodesCountInDLT + entry.result.merkleTreeNodesCountInDLT) / 2

            numberOfActualFalsePositives = (value.result.numberOfActualFalsePositives +entry.result.numberOfActualFalsePositives)/2
            numberOfVCsRetrievedWitnessFromIssuer = (value.result.numberOfVCsRetrievedWitnessFromIssuer +entry.result.numberOfVCsRetrievedWitnessFromIssuer)/2
            numberOfVCsRetrievedWitnessFromDLT = (value.result.numberOfVCsRetrievedWitnessFromDLT + entry.result.numberOfVCsRetrievedWitnessFromDLT) / 2
            avgRevocationTimePerVC = (value.result.avgRevocationTimePerVC + entry.result.avgRevocationTimePerVC) / 2
            verificationTimePerValidVC = (value.result.verificationTimePerValidVC + entry.result.verificationTimePerValidVC) / 2


            verificationTimePerFalsePositiveOrRevokedVC = (value.result.verificationTimePerFalsePositiveOrRevokedVC + entry.result.verificationTimePerFalsePositiveOrRevokedVC) / 2
            avgTimeToFetchWitnessFromIssuer = (value.result.avgTimeToFetchWitnessFromIssuer + entry.result.avgTimeToFetchWitnessFromIssuer) / 2
            avgTimeToFetchWitnessFromDLT = (value.result.avgTimeToFetchWitnessFromDLT + entry.result.avgTimeToFetchWitnessFromDLT) / 2

            bbsProofGenerationTime = (value.result.bbsProofGenerationTime + entry.result.bbsProofGenerationTime) / 2
            bbsVerificationTime = (value.result.bbsVerificationTime + entry.result.bbsVerificationTime) / 2

            contractDeploymentCost = (value.result.contractDeploymentCost + entry.result.contractDeploymentCost) / 2
            bulkIssuanceCost = (value.result.bulkIssuanceCost + entry.result.bulkIssuanceCost) / 2
            avgRevocationCostInGas = (value.result.avgRevocationCostInGas + entry.result.avgRevocationCostInGas) / 2



            verTimeValidVCnpArray = np.asarray(entry.result.verificationTimePerValidVCRawData, dtype=np.float32)
            verTimeValidVCnpArray = np.delete(verTimeValidVCnpArray, [i for i in range(math.ceil(10*verTimeValidVCnpArray.size/100))])
            verificationTimePerValidVCRawData = np.concatenate((value.result.verificationTimePerValidVCRawData, verTimeValidVCnpArray), dtype=float)

            verTimeFalsePositiveOrRevokedVCnpArray = np.asarray(entry.result.verificationTimePerFalsePositiveOrRevokedVCRawData, dtype=np.float32)
            verTimeFalsePositiveOrRevokedVCnpArray = np.delete(verTimeFalsePositiveOrRevokedVCnpArray, [i for i in range(math.ceil(10 * verTimeFalsePositiveOrRevokedVCnpArray.size / 100))])
            verificationTimePerFalsePositiveOrRevokedVCRawData = np.concatenate((value.result.verificationTimePerFalsePositiveOrRevokedVCRawData, verTimeFalsePositiveOrRevokedVCnpArray), dtype=float)

            avgTimeToFetchWitnessFromDLTnpArray = np.asarray(entry.result.avgTimeToFetchWitnessFromDLTRawData, dtype=np.float32)
            avgTimeToFetchWitnessFromDLTRawData = np.concatenate((value.result.avgTimeToFetchWitnessFromDLTRawData, avgTimeToFetchWitnessFromDLTnpArray), dtype=float)

            revocationTimePerVCRawDatanpArray = np.asarray(entry.result.revocationTimePerVCRawData, dtype=np.float32)
            revocationTimePerVCRawDatanpArray = np.delete(revocationTimePerVCRawDatanpArray, [i for i in range(math.ceil(20 * revocationTimePerVCRawDatanpArray.size / 100))])
            revocationTimePerVCRawData = np.concatenate((value.result.revocationTimePerVCRawData,
                                                                 revocationTimePerVCRawDatanpArray), dtype=float)

            avgTimeToFetchWitnessFromIssuernpArray = np.asarray(entry.result.avgTimeToFetchWitnessFromIssuerRawData,
                                                                dtype=np.float32)
            avgTimeToFetchWitnessFromIssuerRawData = np.concatenate(
                (value.result.avgTimeToFetchWitnessFromIssuerRawData,
                 avgTimeToFetchWitnessFromIssuernpArray), dtype=float)


            avgRevocationCostInGasRawDatanpArray = np.asarray(entry.result.avgRevocationCostInGasRawData, dtype=int)
            avgRevocationCostInGasRawDatanpArray = np.delete(avgRevocationCostInGasRawDatanpArray,
                                              [i for i in range(math.ceil(20 * avgRevocationCostInGasRawDatanpArray.size / 100))])
            avgRevocationCostInGasRawData = np.concatenate((value.result.avgRevocationCostInGasRawData,
                                                                 avgRevocationCostInGasRawDatanpArray), dtype=int)


            result = Result(bloomFilterSize=bloomFilterSize,
                            bloomFilterIndexesPerEntry=bloomFilterIndexesPerEntry,
                            merkTreeSizeTotal=merkTreeSizeTotal,
                            merkTreeSizeInDLT=merkTreeSizeInDLT,
                            merkleTreeNodesCountInDLT=merkleTreeNodesCountInDLT,
                            numberOfActualFalsePositives=merkleTreeNodesCountInDLT,
                            numberOfVCsRetrievedWitnessFromIssuer=numberOfActualFalsePositives,
                            numberOfVCsRetrievedWitnessFromDLT=numberOfVCsRetrievedWitnessFromDLT,
                            avgRevocationTimePerVC=avgRevocationTimePerVC,
                            revocationTimePerVCRawData=revocationTimePerVCRawData,
                            verificationTimePerValidVC=verificationTimePerValidVC,
                            verificationTimePerValidVCRawData=verificationTimePerValidVCRawData,
                            verificationTimePerFalsePositiveOrRevokedVC=verificationTimePerFalsePositiveOrRevokedVC,
                            verificationTimePerFalsePositiveOrRevokedVCRawData=verificationTimePerFalsePositiveOrRevokedVCRawData,
                            avgTimeToFetchWitnessFromIssuer=avgTimeToFetchWitnessFromIssuer,
                            avgTimeToFetchWitnessFromIssuerRawData=avgTimeToFetchWitnessFromIssuerRawData,
                            avgTimeToFetchWitnessFromDLT=avgTimeToFetchWitnessFromDLT,
                            avgTimeToFetchWitnessFromDLTRawData=avgTimeToFetchWitnessFromDLTRawData,
                            bbsProofGenerationTime=bbsProofGenerationTime,
                            bbsVerificationTime=bbsVerificationTime,
                            contractDeploymentCost=contractDeploymentCost,
                            bulkIssuanceCost=bulkIssuanceCost,
                            avgRevocationCostInGas=avgRevocationCostInGas,
                            avgRevocationCostInGasRawData=avgRevocationCostInGasRawData)



            entry1 = Entry(setting=setting, result=result)
            values[entry.setting]=entry1
            continue

        values[entry.setting]=entry
        keys[entry.setting]=entry.setting



    return values.values()


def handle_duplicates(entries):
    values = []
    keys = set()
    totalVCs = set()

    for entry in entries:
        if entry.setting.mtLevelInDLT in keys:
            continue
        values.append(entry)
        keys.add(entry.setting.mtLevelInDLT)

    return values



def parse_entry(file):

    # gives the path of demo.py
    path = os.path.realpath(__file__)

    # gives the directory where demo.py
    # exists
    dir = os.path.dirname(path)
    dir  = dir.replace('plots', 'results')
    os.chdir(dir)



    with open(file) as f:
        json_data = json.load(f)

    entries = []
    for entry in json_data:
        totalVCs = entry['total_issued_vcs']
        revokedVCs = entry['total_revoked_vcs']
        falsePositiveRate = entry['false_positive_rate']
        mtLevelInDLT = entry['mt_level_in_dlt']

        setting = Setting(totalVCs=totalVCs, revokedVCs=revokedVCs, falsePositiveRate=falsePositiveRate,
                          mtLevelInDLT= mtLevelInDLT)

        bloomFilterSize = entry['bloom_filter_size']
        bloomFilterIndexesPerEntry = entry['bloom_filter_indexes_per_entry']
        merkTreeSizeTotal = entry['merkle_tree_size_total']
        merkTreeSizeInDLT = entry['merkle_tree_size_in_dlt']
        merkleTreeNodesCountInDLT = entry['merkle_tree_nodes_count_in_dlt']


        numberOfActualFalsePositives = entry['number_of_false_positives']
        numberOfVCsRetrievedWitnessFromIssuer = entry['number_of_vcs_retrieved_witness_from_issuer']
        numberOfVCsRetrievedWitnessFromDLT = entry['number_of_vcs_retrieved_witness_from_dlt']
        avgRevocationTimePerVC = entry['revocation_timeper_vc']
        revocationTimePerVCRawData = entry['revocation_time_raw_data']
        verificationTimePerValidVC = entry['verification_time_per_valid_vc']
        verificationTimePerValidVCRawData = entry['verification_time_per_valid_vc_raw_data']
        verificationTimePerFalsePositiveOrRevokedVC = entry['verification_time_per_false_positive_or_revoked_vc']
        verificationTimePerFalsePositiveOrRevokedVCRawData = entry['verification_time_per_revokedor_false_positive_vc_raw_data']
        avgTimeToFetchWitnessFromIssuer = entry['avg_time_to_fetch_witness_from_issuer']
        avgTimeToFetchWitnessFromIssuerRawData = entry['avg_time_to_fetch_witness_from_issuer_raw_data']
        avgTimeToFetchWitnessFromDLT = entry['avg_time_to_fetch_witness_from_smart_contract']
        avgTimeToFetchWitnessFromDLTRawData = entry['avg_time_to_fetch_witness_from_smart_contract_raw_data']
        if avgTimeToFetchWitnessFromDLTRawData is None:
            avgTimeToFetchWitnessFromDLTRawData = np.array([])
        print("raw data:", avgTimeToFetchWitnessFromDLTRawData)
        bbsProofGenerationTime = entry['bbs_proof_generation_time']
        bbsVerificationTime = entry['bbs_verification_time']

        contractDeploymentCost = entry['contract_deployment_cost']
        bulkIssuanceCost = entry['bulk_issuance_cost']
        avgRevocationCostInGas = entry['revocation_cost_in_wei']
        avgRevocationCostInGasRawData = entry['revocation_cost_raw_data']




        result = Result(bloomFilterSize=bloomFilterSize,
                        bloomFilterIndexesPerEntry = bloomFilterIndexesPerEntry,
                        merkTreeSizeTotal =merkTreeSizeTotal,
                        merkTreeSizeInDLT = merkTreeSizeInDLT,
                        merkleTreeNodesCountInDLT = merkleTreeNodesCountInDLT,
                        numberOfActualFalsePositives = merkleTreeNodesCountInDLT,
                        numberOfVCsRetrievedWitnessFromIssuer = numberOfActualFalsePositives,
                        numberOfVCsRetrievedWitnessFromDLT = numberOfVCsRetrievedWitnessFromDLT,
                        avgRevocationTimePerVC = avgRevocationTimePerVC,
                        revocationTimePerVCRawData=revocationTimePerVCRawData,
                        verificationTimePerValidVC = verificationTimePerValidVC,
                        verificationTimePerValidVCRawData= verificationTimePerValidVCRawData,
                        verificationTimePerFalsePositiveOrRevokedVC= verificationTimePerFalsePositiveOrRevokedVC,
                        verificationTimePerFalsePositiveOrRevokedVCRawData=verificationTimePerFalsePositiveOrRevokedVCRawData,
                        avgTimeToFetchWitnessFromIssuer=avgTimeToFetchWitnessFromIssuer,
                        avgTimeToFetchWitnessFromIssuerRawData=avgTimeToFetchWitnessFromIssuerRawData,
                        avgTimeToFetchWitnessFromDLT= avgTimeToFetchWitnessFromDLT,
                        avgTimeToFetchWitnessFromDLTRawData=avgTimeToFetchWitnessFromDLTRawData,
                        bbsProofGenerationTime=bbsProofGenerationTime,
                        bbsVerificationTime=bbsVerificationTime,
                        contractDeploymentCost=contractDeploymentCost,
                        bulkIssuanceCost=bulkIssuanceCost,
                        avgRevocationCostInGas=avgRevocationCostInGas,
                        avgRevocationCostInGasRawData=avgRevocationCostInGasRawData)

        entry = Entry(setting=setting, result=result)
        entries.append(entry)

    entries_avg = calculate_average(entries)

    for entry in entries_avg:
        print(entry.setting)



    dir  = dir.replace('results','plots')
    os.chdir(dir)
    return entries_avg

