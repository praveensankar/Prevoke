import json
import math

import numpy as np
from matplotlib import pyplot as plt
from scipy.stats import sem


class RevocationCostEntry:
    def __init__(self, totalVCs, revokedVCs, falsePositiveRate, mtLevelInDLT, revocationCost, revocationCostRawData):
        self.totalVCs = totalVCs
        self.revokedVCs = revokedVCs
        self.falsePositiveRate = falsePositiveRate
        self.mtLevelInDLT = mtLevelInDLT
        self.revocationCost = revocationCost
        self.revocationCostRawData = revocationCostRawData

    def __str__(self):
        output = "total vcs: " + str(self.totalVCs)
        output += "\t revoked vcs: " + str(self.revokedVCs)
        output += "\t false positive: " + str(self.falsePositiveRate)
        output += "\t mt level in dlt: " + str(self.mtLevelInDLT)
        output += "\t revocation cost: " + str(self.revocationCost)
        return output

    def __eq__(self, another):
        return self.totalVCs == another.totalVCs and self.revokedVCs == another.revokedVCs and self.falsePositiveRate == another.falsePositiveRate and self.mtLevelInDLT == another.mtLevelInDLT

    def __hash__(self):
        return hash(str(self.totalVCs) + str(self.revokedVCs) + str(self.falsePositiveRate) + str(self.mtLevelInDLT))


def calculate_average(entries):
    values = {}
    keys = set()

    for entry in entries:
        if entry.__hash__() in keys:
            value = values[entry]
            values[entry].revocationCost = (value.revocationCost+entry.revocationCost)/2
            rawData = np.asarray(entry.revocationCostRawData)
            rawData = np.delete(rawData, [i for i in range(math.ceil(10 * rawData.size / 100))])
            existingRawData = np.asarray(value.revocationCostRawData)
            values[entry].revocationCostRawData = np.concatenate((existingRawData, rawData), dtype=int)
            continue

        values[entry]=entry
        values[entry].revocationCostRawData = np.asarray(entry.revocationCostRawData)
        values[entry].revocationCostRawData = np.delete(values[entry].revocationCostRawData,
                                          [i for i in range(math.ceil(10 * values[entry].revocationCostRawData.size / 100))])
        keys.add(entry.__hash__())

    return values.values()


def parse_revocation_cost_entry(file):
    with open(file) as f:
        json_data = json.load(f)

    entries = []
    for entry in json_data:
        revokedVCs = entry['total_revoked_vcs']
        falsePositiveRate = entry['false_positive_rate']
        mtLevelInDLT = entry['mt_level_in_dlt']
        revocationCost = entry['revocation_cost_in_wei']
        revocationCostRawData = entry['revocation_cost_raw_data']




        entry = RevocationCostEntry(totalVCs=0, revokedVCs=revokedVCs, falsePositiveRate=falsePositiveRate,
                          mtLevelInDLT= mtLevelInDLT, revocationCost=revocationCost, revocationCostRawData=revocationCostRawData)

        entries.append(entry)

    for entry in entries:
        print(entry)

    average_values = calculate_average(entries)
    for entry in average_values:
        print(entry)

    return average_values




def plot_revocation_cost():
    entries = parse_revocation_cost_entry("results_revocation_cost.json")
    costsfor01 = {}
    costsfor001 = {}
    costsfor0001 = {}
    costsfor00001 = {}
    for entry in entries:
        if entry.falsePositiveRate==0.1:
            costsfor01[entry.mtLevelInDLT]=np.mean(entry.revocationCostRawData)
            # mtlevels.append(entry.setting.mtLevelInDLT)
            # costs.append(entry.result.mtAccumulatorPerUpdateCost)
            # entry.result.mtAccumulatorPerUpdateCost / 1000000

        if entry.falsePositiveRate == 0.01:
            costsfor001[entry.mtLevelInDLT] = np.mean(entry.revocationCostRawData)

        if entry.falsePositiveRate == 0.001:
            costsfor0001[entry.mtLevelInDLT] = np.mean(entry.revocationCostRawData)

        if entry.falsePositiveRate == 0.0001:
            costsfor00001[entry.mtLevelInDLT] = np.mean(entry.revocationCostRawData)


    costsfor01 = dict(sorted(costsfor01.items()))
    x1points = np.array(list(costsfor01.keys()))
    y1points = np.array(list(costsfor01.values()))
    costsfor001 = dict(sorted(costsfor001.items()))
    x2points = np.array(list(costsfor001.keys()))
    y2points = np.array(list(costsfor001.values()))
    costsfor0001 = dict(sorted(costsfor0001.items()))
    x3points = np.array(list(costsfor0001.keys()))
    y3points = np.array(list(costsfor0001.values()))
    costsfor00001 = dict(sorted(costsfor00001.items()))
    x4points = np.array(list(costsfor00001.keys()))
    y4points = np.array(list(costsfor00001.values()))

    y1error = sem(y1points)
    y2error = sem(y2points)
    y3error = sem(y3points)
    y4error = sem(y4points)

    print(x1points)
    print(y1points)
    print(x2points)
    print(y2points)
    print(x3points)
    print(y3points)
    print(x4points)
    print(y4points)

    yRange = np.linspace(start=0, stop=max(costsfor00001.values()),
                           num=25)

    yRange =[i*10000 for i in (5,7,9,11,13,15,17,19,21,23,25,27,29,31,33)]
    ylabel = [str(i*10)+"k" for i in (5,7,9,11,13,15,17,19,21,23,25,27,29,31,33)]
    print(yRange)
    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    plt.errorbar(x1points, y1points, marker='o', color='#614415', label="fpr=0.1", yerr=y1error)
    plt.errorbar(x2points, y2points, marker='d', color='#0072b2', label="fpr=0.01", yerr=y2error)
    plt.errorbar(x3points, y3points, marker='*', color='#d55e00', label="fpr=0.001", yerr=y3error)
    plt.errorbar(x4points, y4points, marker='+', color='#009e73', label="fpr=0.0001", yerr=y4error)
    plt.yticks(yRange, ylabel)
    plt.xticks(x1points)
    # plt.title('gas cost per ', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('amount of gas', font)
    plt.legend(fontsize="13")
    plt.savefig("graphs/result_revocation_cost.png")

