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


def handle_duplicates(entries):


    values = []
    keys = set()
    for entry in entries:
        if entry.setting.mtLevelInDLT in keys:
            continue
        values.append(entry)
        keys.add(entry.setting.mtLevelInDLT)

    return values


def parse_entry():
    with open("results.json") as f:
        json_data = json.load(f)

    entries = []
    for entry in json_data:
        totalVCs = entry['total_issued_vcs']
        revokedVCs = entry['total_revoked_vcs']
        falsePositiveRate = entry['false_positive_rate']
        mtLevelInDLT = entry['mt_level_in_dlt']
        bloomFilterSize = entry['bloom_filter_size']
        bloomFilterIndexesPerEntry = entry['bloom_filter_indexes_per_entry']
        mtAccumulatorPerUpdateCost = entry['mt_accumulator_per_update_cost_in_gwei']
        numberOfActualFalsePositives = entry['number_of_false_positives']
        numberOfVCsRetrievedWitnessFromIssuer = entry['number_of_vcs_retrieved_witness_from_issuer']
        numberOfVCsAffectedByMTAccumulator = entry['number_of_affected_vcs_by_MT_accumulator']

        setting = Setting(totalVCs=totalVCs, revokedVCs=revokedVCs, falsePositiveRate=falsePositiveRate,
                          mtLevelInDLT= mtLevelInDLT, bloomFilterSize=bloomFilterSize,
                          bloomFilterIndexesPerEntry=bloomFilterIndexesPerEntry)
        result = Result(mtAccumulatorPerUpdateCost=mtAccumulatorPerUpdateCost, numberOfActualFalsePositives=numberOfActualFalsePositives,
                        numberOfVCsRetrievedWitnessFromIssuer=numberOfVCsRetrievedWitnessFromIssuer, numberOfVCsAffectedByMTAccumulator=numberOfVCsAffectedByMTAccumulator)
        entry = Entry(setting=setting, result=result)
        entries.append(entry)

        entries_without_duplicates = handle_duplicates(entries)

    return entries_without_duplicates

