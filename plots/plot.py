import numpy as np
import pandas as pd
import json

from matplotlib import pyplot as plt

import revocation_cost
from bbs_time import calculate_bbs_time
from fpr_impact import *
from fpr_vs_bf_size_and_witness_updates import plot_fpr_vs_bfsize_and_witUpdates
from indy_vs_ours import plot_witness_updates_vc_indy
from entry import Entry, parse_entry
from result import Result
from scale_revocation import *
from setting import Setting
from verification_time import *


def main():


    # revocation_cost.plot_revocation_cost()
    # plot_verification_time()
    # calculate_bbs_time()
    # plot_impact_of_fpr("10K",1000, "random")
    # plot_impact_of_fpr("10K",1000, "oldest")
    # plot_impact_of_fpr("10K",500, "random")
    # plot_impact_of_fpr("10K", 500, "oldest")
    # plot_impact_of_fpr("100K",10000, "random")
    # plot_impact_of_fpr("100K",10000, "oldest")
    # plot_impact_of_fpr("100K",5000, "random")
    # plot_impact_of_fpr("100K",5000, "oldest")
    # plot_impact_of_fpr_with_false_positive("100K",10000, "oldest", 0.1)
    # plot_impact_of_fpr_with_false_positive("100K",10000, "random", 0.1)
    # plot_impact_of_fpr_with_false_positive("100K",10000, "oldest", 0.001)
    # plot_impact_of_fpr_with_false_positive("100K",10000, "random", 0.001)
    # plot_impact_of_fpr_with_false_positive_random_vs_oldest("100K", 10000, 0.001)
    # scale_revocations("100K")
    scale_revocations_fixed_bloomFilter("100K")
    scale_revocations_fixed_bloomFilter_linear_scale_plot("100K")




def plot_witness_update_saves(entries):
    mtlevelsfor1000 = []
    mtlevelsfor5000 = []
    witnessUpdateSavesfor1000 = []
    witnessUpdateSavesfor5000 = []

    for entry in entries:
        if entry.setting.totalVCs==5000:
            mtlevelsfor5000.append(entry.setting.mtLevelInDLT)
            savedWitnessUpdates = entry.result.numberOfActualFalsePositives - entry.result.numberOfVCsRetrievedWitnessFromIssuer
            witnessUpdateSavesfor5000.append(savedWitnessUpdates)

        if entry.setting.totalVCs==1000 and entry.setting.falsePositiveRate==0.1:
            mtlevelsfor1000.append(entry.setting.mtLevelInDLT)
            savedWitnessUpdates = entry.result.numberOfActualFalsePositives - entry.result.numberOfVCsRetrievedWitnessFromIssuer
            witnessUpdateSavesfor1000.append(savedWitnessUpdates)


    x1points = np.array(mtlevelsfor1000)
    y1points = np.array(witnessUpdateSavesfor1000)
    x2points = np.array(mtlevelsfor5000)
    y2points = np.array(witnessUpdateSavesfor5000)
    print(x1points)
    print(y1points)
    print(x2points)
    print(y2points)

    wusRange=np.linspace(start=0, stop= max(max(witnessUpdateSavesfor1000), max(witnessUpdateSavesfor5000)), num=max(len(x1points), len(x2points)))


    print(wusRange)
    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 18, 'weight': 'bold'}
    plt.plot(x1points, y1points, marker = 'o', label="1000 vcs, 0.1 false positive rate")
    plt.plot(x2points, y2points, marker='d', label="5000 vcs, 0.1 false positive rate")
    plt.yticks(wusRange)
    plt.title('number of vcs benefitted from two phase approach', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('no. of witness updates saved', font)
    plt.legend()
    plt.savefig("graphs/witness_updates_saved.png")

def plot_false_positives(entries):
    mtlevelsfor1000 = []
    mtlevelsfor5000 = []
    falsePositiviesfor1000 = []
    falsePositiviesfor5000 = []

    for entry in entries:
        if entry.setting.totalVCs == 5000 and entry.setting.falsePositiveRate==0.1:
            mtlevelsfor5000.append(entry.setting.mtLevelInDLT)
            falsePositiviesfor5000.append(entry.result.numberOfActualFalsePositives)

        if entry.setting.totalVCs == 1000 and entry.setting.falsePositiveRate == 0.1:
            mtlevelsfor1000.append(entry.setting.mtLevelInDLT)
            falsePositiviesfor1000.append(entry.result.numberOfActualFalsePositives)

    x1points = np.array(mtlevelsfor1000)
    y1points = np.array(falsePositiviesfor1000)
    x2points = np.array(mtlevelsfor5000)
    y2points = np.array(falsePositiviesfor5000)
    print(x1points)
    print(y1points)
    print(x2points)
    print(y2points)

    wusRange = np.linspace(start=0, stop=max(max(falsePositiviesfor1000), max(falsePositiviesfor5000)),
                           num=max(len(x1points), len(x2points)))

    print(wusRange)
    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 18, 'weight': 'bold'}
    plt.plot(x1points, y1points, marker='o', label="1000 vcs, 0.1 false positive rate")
    plt.plot(x2points, y2points, marker='d', label="5000 vcs, 0.1 false positive rate")
    plt.yticks(wusRange)
    plt.title('number of occured false positives', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('false positivies', font)
    plt.legend()
    plt.savefig("graphs/false_positivies.png")


def plot_merkle_tree_accumulator_cost(entries):
    costsfor1000 = {}
    costsfor5000 = {}
    for entry in entries:
        if entry.setting.totalVCs==1000:
            costsfor1000[entry.setting.mtLevelInDLT]=entry.result.mtAccumulatorPerUpdateCost
            # mtlevels.append(entry.setting.mtLevelInDLT)
            # costs.append(entry.result.mtAccumulatorPerUpdateCost)
            # entry.result.mtAccumulatorPerUpdateCost / 1000000

        if entry.setting.totalVCs==5000:
                    costsfor5000[entry.setting.mtLevelInDLT]=entry.result.mtAccumulatorPerUpdateCost


    costsfor1000 = dict(sorted(costsfor1000.items()))
    x1points = np.array(list(costsfor1000.keys()))
    y1points = np.array(list(costsfor1000.values()))
    costsfor5000 = dict(sorted(costsfor5000.items()))
    x2points = np.array(list(costsfor5000.keys()))
    y2points = np.array(list(costsfor5000.values()))
    print(x1points)
    print(y1points)
    print(x2points)
    print(y2points)

    wusRange = np.linspace(start=0, stop=max(costsfor1000.values()),
                           num=max(len(x1points), len(x2points)))

    print(wusRange)
    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 18, 'weight': 'bold'}
    plt.plot(x1points, y1points, marker='o', label="1000 vcs, storage cost of mt accumulator")
    plt.plot(x2points, y2points, marker='d', label="5000 vcs, storage cost of mt accumulator")
    plt.yticks(wusRange)
    plt.title('storage cost of mt accumulator', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('gwei', font)
    plt.legend()
    plt.savefig("graphs/cost_mt_accumulator.png")





def plot_witness_updates(entries):
    mtlevelsfor1000 = []
    mtlevelsfor5000 = []
    witnessUpdateSavesfor1000 = []
    witnessUpdateSavesfor5000 = []

    for entry in entries:
        if entry.setting.totalVCs==5000:
            mtlevelsfor5000.append(entry.setting.mtLevelInDLT)
            savedWitnessUpdates = entry.result.numberOfVCsRetrievedWitnessFromIssuer
            witnessUpdateSavesfor5000.append(savedWitnessUpdates)

        if entry.setting.totalVCs==1000 and entry.setting.falsePositiveRate==0.1:
            mtlevelsfor1000.append(entry.setting.mtLevelInDLT)
            savedWitnessUpdates = entry.result.numberOfVCsRetrievedWitnessFromIssuer
            witnessUpdateSavesfor1000.append(savedWitnessUpdates)


    x1points = np.array(mtlevelsfor1000)
    y1points = np.array(witnessUpdateSavesfor1000)
    x2points = np.array(mtlevelsfor5000)
    y2points = np.array(witnessUpdateSavesfor5000)
    print(x1points)
    print(y1points)
    print(x2points)
    print(y2points)

    wusRange=np.linspace(start=0, stop= max(max(witnessUpdateSavesfor1000), max(witnessUpdateSavesfor5000)), num=max(len(x1points), len(x2points)))


    print(wusRange)
    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 18, 'weight': 'bold'}
    plt.plot(x1points, y1points, marker = 'o', label="1000 vcs")
    plt.plot(x2points, y2points, marker='d', label="5000 vcs")
    plt.yticks(wusRange)
    plt.title('number of vcs benefitted from two phase approach', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('no. of vcs retrieved witnesses from issuer', font)
    plt.legend()
    plt.savefig("graphs/witness_updates.png")




def plot_witness_update_saved_for_different_false_positives(entries):
    mtlevelsfor1 = []
    mtlevelsfor2 = []
    mtlevelsfor3 = []
    mtlevelsfor4 = []

    witnessUpdateSavesfor1 = []
    witnessUpdateSavesfor2 = []
    witnessUpdateSavesfor3 = []
    witnessUpdateSavesfor4 = []

    for entry in entries:
        if entry.setting.totalVCs==1000:
            if entry.setting.falsePositiveRate==0.1:
                mtlevelsfor1.append(entry.setting.mtLevelInDLT)
                savedWitnessUpdates = entry.result.numberOfActualFalsePositives - entry.result.numberOfVCsRetrievedWitnessFromIssuer
                witnessUpdateSavesfor1.append(savedWitnessUpdates)

            if entry.setting.falsePositiveRate==0.2:
                mtlevelsfor2.append(entry.setting.mtLevelInDLT)
                savedWitnessUpdates = entry.result.numberOfActualFalsePositives - entry.result.numberOfVCsRetrievedWitnessFromIssuer
                witnessUpdateSavesfor2.append(savedWitnessUpdates)


            if entry.setting.falsePositiveRate==0.3:
                mtlevelsfor3.append(entry.setting.mtLevelInDLT)
                savedWitnessUpdates = entry.result.numberOfActualFalsePositives - entry.result.numberOfVCsRetrievedWitnessFromIssuer
                witnessUpdateSavesfor3.append(savedWitnessUpdates)


            if entry.setting.falsePositiveRate==0.4:
                mtlevelsfor4.append(entry.setting.mtLevelInDLT)
                savedWitnessUpdates = entry.result.numberOfActualFalsePositives - entry.result.numberOfVCsRetrievedWitnessFromIssuer
                witnessUpdateSavesfor4.append(savedWitnessUpdates)


    x1points = np.array(mtlevelsfor1)
    y1points = np.array(witnessUpdateSavesfor1)
    x2points = np.array(mtlevelsfor2)
    y2points = np.array(witnessUpdateSavesfor2)

    x3points = np.array(mtlevelsfor3)
    y3points = np.array(witnessUpdateSavesfor3)
    x4points = np.array(mtlevelsfor4)
    y4points = np.array(witnessUpdateSavesfor4)
    print(x1points)
    print(y1points)
    print(x2points)
    print(y2points)

    maxwitnessUpdatesaves = max(max(witnessUpdateSavesfor1), max(witnessUpdateSavesfor2), max(witnessUpdateSavesfor3), max(witnessUpdateSavesfor4))

    numberOfEntries = max(len(x1points), len(x2points), len(x3points), len(x4points))

    wusRange=np.linspace(start=0, stop= maxwitnessUpdatesaves, num=numberOfEntries)


    print(wusRange)
    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname':'Times New Roman', 'size': 18, 'weight':'bold'}
    plt.plot(x1points, y1points, marker = 'o', label="0.1 false positive rate")
    plt.plot(x2points, y2points, marker='d', label="0.2 false positive rate")
    plt.plot(x3points, y3points, marker = '*', label="0.3 false positive rate")
    plt.plot(x4points, y4points, marker='s', label="0.4 false positive rate")
    plt.yticks(wusRange)
    plt.title('False positives vs witness updates (1000 vcs)', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('no. of witness updates saved', font)
    plt.legend()
    plt.savefig("graphs/witness_updates_saved_1000_for_different_fpr.png")




def plot_witness_update_saved_due_to_levels_in_dlt(entries):
    mtlevelsfor1000 = []
    mtlevelsfor5000 = []
    NoOfAffectedVCs1000 = []
    NoOfAffectedVCs5000 = []

    for entry in entries:
        if entry.setting.totalVCs == 5000 and entry.setting.falsePositiveRate==0.1:
            mtlevelsfor5000.append(entry.setting.mtLevelInDLT)
            NoOfAffectedVCs5000.append(entry.result.numberOfVCsAffectedByMTAccumulator/1000)

        if entry.setting.totalVCs == 1000 and entry.setting.falsePositiveRate==0.1:
            mtlevelsfor1000.append(entry.setting.mtLevelInDLT)
            NoOfAffectedVCs1000.append(entry.result.numberOfVCsAffectedByMTAccumulator/1000)


    x1points = np.array(mtlevelsfor1000)
    y1points = np.array(NoOfAffectedVCs1000)
    x2points = np.array(mtlevelsfor5000)
    y2points = np.array(NoOfAffectedVCs5000)
    print(x1points)
    print(y1points)
    print(x2points)
    print(y2points)

    wusRange = np.linspace(start=0, stop=max(max(NoOfAffectedVCs1000), max(NoOfAffectedVCs5000)),
                           num=max(len(x1points), len(x2points)))

    print(wusRange)
    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 18, 'weight': 'bold'}
    plt.plot(x1points, y1points, marker='o', label="1000 vcs")
    plt.plot(x2points, y2points, marker='d', label="5000 vcs")
    plt.yticks(wusRange)
    plt.title('MT Acc vs witness updates', font)
    plt.xlabel('level of MT Acc stored in DLT', font)
    plt.ylabel('total no. of witness updates (in 1000s)', font)
    plt.legend()
    plt.savefig("graphs/mt_acc_impact.png")





if __name__=="__main__":
    main()


