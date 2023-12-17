import numpy as np
import pandas as pd
import json

from matplotlib import pyplot as plt

from entry import Entry, parse_entry
from result import Result
from setting import Setting


def main():

    entries = parse_entry()
    # plot_false_positives(entries)
    # plot_merkle_tree_accumulator_cost(entries)
    # plot_witness_updates_vc_indy(entries)
    # plot_witness_updates(entries)
    # plot_witness_update_saves(entries)
    # plot_witness_update_saved_for_different_false_positives(entries)
    # plot_witness_update_saved_due_to_levels_in_dlt(entries)
    plot_witness_updates_vc_indy(entries)

def plot_witness_updates_vc_indy(entries):
    mtlevelsfor1000 = []
    indy = 0
    witnessUpdatefor1000 = []
    WitnessUpdatesOfIndy1000 = []




    for entry in entries:

        if entry.setting.totalVCs == 1000 and entry.setting.falsePositiveRate == 0.1:
            mtlevelsfor1000.append(entry.setting.mtLevelInDLT)
            WitnessUpdates = entry.result.numberOfVCsRetrievedWitnessFromIssuer
            witnessUpdatefor1000.append(WitnessUpdates)
            WitnessUpdatesOfIndy1000.append(10000)
    font = {'fontname': 'Times New Roman', 'size': 17, 'weight': 'bold'}

    x1points = np.array(mtlevelsfor1000)
    print(x1points)
    x = np.arange(len(x1points))  # the label locations
    fig, ax = plt.subplots(layout='constrained')
    # ax = fig.add_axes(x1points)



    wuRange = np.linspace(start=50, stop=max(WitnessUpdatesOfIndy1000),
                           num=len(x1points))

    ax.bar(x+0.00, witnessUpdatefor1000, color = 'g', width = 0.25, label="two phase technique")
    ax.bar(x + 0.25,WitnessUpdatesOfIndy1000, color = 'r', width = 0.25, label="indy")
    ax.set_title('indy V two phase technique: 1000 vcs, 100 revocations',font)
    ax.set_ylabel('no of vcs requried to witness updates', font)
    ax.set_xlabel('merkle tree accumulator levels in DLT', font)
    ax.set_xticks(range(len(x1points)), x1points)
    ax.set_yticks(wuRange)
    ax.legend()
    plt.savefig("graphs/witness_updates_indy_vs_ours.png")




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
    mtlevels = []
    costs = []
    for entry in entries:
        if entry.setting.totalVCs==5000:
            mtlevels.append(entry.setting.mtLevelInDLT)
            costs.append(entry.result.mtAccumulatorPerUpdateCost/1000000)

    xpoints = np.array(mtlevels)
    ypoints = np.array(costs)
    print(xpoints)
    print(ypoints)

    costrange=np.linspace(start=min(costs), stop= max(costs), num=len(costs))


    print(costrange)
    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 18, 'weight': 'bold'}
    plt.plot(xpoints, ypoints, marker = 'o')
    plt.yticks(costrange)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('mili ether', font)

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


