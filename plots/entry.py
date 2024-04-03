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
                              mtLevelInDLT=entry.setting.mtLevelInDLT, bloomFilterSize=entry.setting.bloomFilterSize,
                              bloomFilterIndexesPerEntry=entry.setting.bloomFilterIndexesPerEntry)


            value = values[entry.setting]
            numberOfActualFalsePositives = (value.result.numberOfActualFalsePositives +entry.result.numberOfActualFalsePositives)/2
            numberOfVCsRetrievedWitnessFromIssuer = (value.result.numberOfVCsRetrievedWitnessFromIssuer + entry.result.numberOfVCsRetrievedWitnessFromIssuer) / 2
            numberOfVCsAffectedByMTAccumulator = (value.result.numberOfVCsAffectedByMTAccumulator + entry.result.numberOfVCsAffectedByMTAccumulator) / 2
            mtAccumulatorPerUpdateCost = (value.result.mtAccumulatorPerUpdateCost + entry.result.mtAccumulatorPerUpdateCost) / 2
            result = Result(mtAccumulatorPerUpdateCost=mtAccumulatorPerUpdateCost,
                            numberOfActualFalsePositives=numberOfActualFalsePositives,
                            numberOfVCsRetrievedWitnessFromIssuer=numberOfVCsRetrievedWitnessFromIssuer,
                            numberOfVCsAffectedByMTAccumulator=numberOfVCsAffectedByMTAccumulator)

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

        if "bloom_filter_size" in entry:
            bloomFilterSize = entry['bloom_filter_size']
        else:
            bloomFilterSize = 0

        if "bloom_filter_indexes_per_entry" in entry:
            bloomFilterIndexesPerEntry = entry['bloom_filter_indexes_per_entry']
        else:
            bloomFilterIndexesPerEntry = 0



        if "revocation_cost_in_wei" in entry:
            mtAccumulatorPerUpdateCost = entry['revocation_cost_in_wei']
        else:
            mtAccumulatorPerUpdateCost = 0

        numberOfActualFalsePositives = entry['number_of_false_positives']
        numberOfVCsRetrievedWitnessFromIssuer = entry['number_of_vcs_retrieved_witness_from_issuer']

        if "number_of_affected_vcs_by_MT_accumulator" in entry:
            numberOfVCsAffectedByMTAccumulator = entry['number_of_affected_vcs_by_MT_accumulator']

        if "number_of_vcs_affected_by_revocation_due_to_storing_optimized_MT_accumulator_in_DLT" in entry:
            numberOfVCsAffectedByMTAccumulator = entry['number_of_vcs_affected_by_revocation_due_to_storing_optimized_MT_accumulator_in_DLT']

        setting = Setting(totalVCs=totalVCs, revokedVCs=revokedVCs, falsePositiveRate=falsePositiveRate,
                          mtLevelInDLT= mtLevelInDLT, bloomFilterSize=bloomFilterSize,
                          bloomFilterIndexesPerEntry=bloomFilterIndexesPerEntry)
        result = Result(mtAccumulatorPerUpdateCost=mtAccumulatorPerUpdateCost, numberOfActualFalsePositives=numberOfActualFalsePositives,
                        numberOfVCsRetrievedWitnessFromIssuer=numberOfVCsRetrievedWitnessFromIssuer, numberOfVCsAffectedByMTAccumulator=numberOfVCsAffectedByMTAccumulator)
        entry = Entry(setting=setting, result=result)
        entries.append(entry)

    entries_without_duplicates = calculate_average(entries)
    print(entries_without_duplicates)

    return entries_without_duplicates

