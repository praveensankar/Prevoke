import json

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
            verificationTimePerValidVCRawData = (value.result.verificationTimePerValidVCRawData + entry.result.verificationTimePerValidVCRawData)
            verificationTimePerFalsePositiveOrRevokedVC = (value.result.verificationTimePerFalsePositiveOrRevokedVC + entry.result.verificationTimePerFalsePositiveOrRevokedVC) / 2
            verificationTimePerFalsePositiveOrRevokedVCRawData = (value.result.verificationTimePerFalsePositiveOrRevokedVCRawData + entry.result.verificationTimePerFalsePositiveOrRevokedVCRawData)
            avgTimeToFetchWitnessFromIssuer = (value.result.avgTimeToFetchWitnessFromIssuer + entry.result.avgTimeToFetchWitnessFromIssuer) / 2
            avgTimeToFetchWitnessFromIssuerRawData = (value.result.avgTimeToFetchWitnessFromIssuerRawData + entry.result.avgTimeToFetchWitnessFromIssuerRawData)
            avgTimeToFetchWitnessFromDLT = (value.result.avgTimeToFetchWitnessFromDLT + entry.result.avgTimeToFetchWitnessFromDLT) / 2
            avgTimeToFetchWitnessFromDLTRawData = (value.result.avgTimeToFetchWitnessFromDLTRawData + entry.result.avgTimeToFetchWitnessFromDLTRawData)
            bbsProofGenerationTime = (value.result.bbsProofGenerationTime + entry.result.bbsProofGenerationTime) / 2
            bbsVerificationTime = (value.result.bbsVerificationTime + entry.result.bbsVerificationTime) / 2

            contractDeploymentCost = (value.result.contractDeploymentCost + entry.result.contractDeploymentCost) / 2
            bulkIssuanceCost = (value.result.bulkIssuanceCost + entry.result.bulkIssuanceCost) / 2
            avgRevocationCostInGas = (value.result.avgRevocationCostInGas + entry.result.avgRevocationCostInGas) / 2
            avgRevocationCostInGasRawData = (value.result.avgRevocationCostInGasRawData + entry.result.avgRevocationCostInGasRawData)

            result = Result(bloomFilterSize=bloomFilterSize,
                            bloomFilterIndexesPerEntry=bloomFilterIndexesPerEntry,
                            merkTreeSizeTotal=merkTreeSizeTotal,
                            merkTreeSizeInDLT=merkTreeSizeInDLT,
                            merkleTreeNodesCountInDLT=merkleTreeNodesCountInDLT,
                            numberOfActualFalsePositives=merkleTreeNodesCountInDLT,
                            numberOfVCsRetrievedWitnessFromIssuer=numberOfActualFalsePositives,
                            numberOfVCsRetrievedWitnessFromDLT=numberOfVCsRetrievedWitnessFromDLT,
                            avgRevocationTimePerVC=avgRevocationTimePerVC,
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
        verificationTimePerValidVC = entry['verification_time_per_valid_vc']
        verificationTimePerValidVCRawData = entry['verification_time_per_valid_vc_raw_data']
        verificationTimePerFalsePositiveOrRevokedVC = entry['verification_time_per_false_positive_or_revoked_vc']
        verificationTimePerFalsePositiveOrRevokedVCRawData = entry['verification_time_per_revokedor_false_positive_vc_raw_data']
        avgTimeToFetchWitnessFromIssuer = entry['avg_time_to_fetch_witness_from_issuer']
        avgTimeToFetchWitnessFromIssuerRawData = entry['avg_time_to_fetch_witness_from_issuer_raw_data']
        avgTimeToFetchWitnessFromDLT = entry['avg_time_to_fetch_witness_from_smart_contract']
        avgTimeToFetchWitnessFromDLTRawData = entry['avg_time_to_fetch_witness_from_smart_contract_raw_data']
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
        print(entry.__str__())


    return entries_avg

