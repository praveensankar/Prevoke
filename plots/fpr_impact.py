import json
import os

import numpy as np
from matplotlib import pyplot as plt
from matplotlib.ticker import LogLocator, ScalarFormatter
from scipy.stats import sem


class FPRImpactEntry:
    def __init__(self, totalVCs, revokedVCs, revocationMode, falsePositiveRate, mtLevelInDLT, numberOfActualFalsePositives, numberOfVCsRetrievedWitnessFromDLT):
        self.totalVCs = totalVCs
        self.revokedVCs = revokedVCs
        self.revocationMode = revocationMode
        self.falsePositiveRate = falsePositiveRate
        self.mtLevelInDLT = mtLevelInDLT
        self.numberOfActualFalsePositives = numberOfActualFalsePositives
        self.numberOfVCsRetrievedWitnessFromDLT = numberOfVCsRetrievedWitnessFromDLT


    def __str__(self):
        output = "total vcs: " + str(self.totalVCs)
        output += "\t revoked vcs: " + str(self.revokedVCs)
        output += "\t false positive: " + str(self.falsePositiveRate)
        output += "\t mt level in dlt: " + str(self.mtLevelInDLT)
        output += "\t revocation mode: " + str(self.revocationMode)
        output += "\t number of false positives: " + str(self.numberOfActualFalsePositives)
        output += "\t number of vcs retrieved witnesses from DLT: " + str(self.numberOfVCsRetrievedWitnessFromDLT)
        return output


    def __eq__(self, another):
        return self.totalVCs == another.totalVCs and self.revokedVCs == another.revokedVCs and self.falsePositiveRate == another.falsePositiveRate and self.mtLevelInDLT == another.mtLevelInDLT and self.revocationMode == another.revocationMode

    def __hash__(self):
        return hash(str(self.totalVCs) + str(self.revokedVCs) + str(self.falsePositiveRate) + str(self.mtLevelInDLT) +  str(self.revocationMode))




def parse_fpr_impact_entry(file):

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
        totalVCs=entry['total_issued_vcs']
        revokedVCs = entry['total_revoked_vcs']
        falsePositiveRate = entry['false_positive_rate']
        mtLevelInDLT = entry['mt_level_in_dlt']
        revocationMode = entry['revocation_mode']
        numberOfActualFalsePositives = entry['number_of_false_positives']
        numberOfVCsRetrievedWitnessFromDLT = entry['number_of_vcs_retrieved_witness_from_dlt']



        entry = FPRImpactEntry(totalVCs=totalVCs, revokedVCs=revokedVCs, falsePositiveRate=falsePositiveRate,
                          mtLevelInDLT= mtLevelInDLT, revocationMode=revocationMode, numberOfActualFalsePositives=numberOfActualFalsePositives,
                               numberOfVCsRetrievedWitnessFromDLT=numberOfVCsRetrievedWitnessFromDLT)

        entries.append(entry)

    # for entry in entries:
    #     print(entry)
    dir  = dir.replace('results','plots')
    os.chdir(dir)
    return entries


def plot_impact_of_fpr(totalVCs, revokedVCs, revocationMode):

    res = list()


    if totalVCs=="10K":
        res1 = parse_fpr_impact_entry("results_computed_10K_run1.json")
        res2 = parse_fpr_impact_entry("results_computed_10K_run2.json")
        res3 = parse_fpr_impact_entry("results_computed_10K_run3.json")

        for result in res1:
            res.append(result)

        for result in res2:
            res.append(result)

        for result in res3:
            res.append(result)

    if totalVCs == "100K":
        res1 = parse_fpr_impact_entry("results_computed_100K_run1.json")
        res2 = parse_fpr_impact_entry("results_computed_100K_run2.json")
        res3 = parse_fpr_impact_entry("results_computed_100K_run3.json")

        for result in res1:
            res.append(result)

        for result in res2:
            res.append(result)

        for result in res3:
            res.append(result)

    # for r in res:
    #     print(r)

    fp1 = {}
    n1 = {}

    fp2 = {}
    n2 = {}

    fp3 = {}
    n3 = {}

    fp4 ={}
    n4 = {}



    
    for entry in res:
        t = []
        t.append(entry.numberOfActualFalsePositives)
        tArray = np.array(t)
        g = []
        g.append(entry.numberOfVCsRetrievedWitnessFromDLT)
        gArray = np.array(g)
        if entry.revokedVCs== revokedVCs and entry.falsePositiveRate == 0.1 and entry.revocationMode==revocationMode:
            if entry.mtLevelInDLT in fp1.keys():
                fp1[entry.mtLevelInDLT] = np.append(fp1[entry.mtLevelInDLT], tArray)
                n1[entry.mtLevelInDLT] =  np.append(n1[entry.mtLevelInDLT],gArray)
            else:
                fp1[entry.mtLevelInDLT]= tArray
                n1[entry.mtLevelInDLT] = gArray

        if entry.revokedVCs== revokedVCs and entry.falsePositiveRate == 0.01 and entry.revocationMode == revocationMode:
            if entry.mtLevelInDLT in fp2.keys():
                fp2[entry.mtLevelInDLT] = np.append(fp2[entry.mtLevelInDLT], tArray)
                n2[entry.mtLevelInDLT] = np.append(n2[entry.mtLevelInDLT], gArray)
            else:
                fp2[entry.mtLevelInDLT] = tArray
                n2[entry.mtLevelInDLT] = gArray

        
        if entry.revokedVCs== revokedVCs and  entry.falsePositiveRate == 0.001 and entry.revocationMode == revocationMode:
            if entry.mtLevelInDLT in fp3.keys():
                fp3[entry.mtLevelInDLT] = np.append(fp3[entry.mtLevelInDLT], tArray)
                n3[entry.mtLevelInDLT] = np.append(n3[entry.mtLevelInDLT], gArray)
            else:
                fp3[entry.mtLevelInDLT] = tArray
                n3[entry.mtLevelInDLT] = gArray


        if entry.revokedVCs== revokedVCs and  entry.falsePositiveRate == 0.0001 and entry.revocationMode == revocationMode:
            if entry.mtLevelInDLT in fp4.keys():
                fp4[entry.mtLevelInDLT] = np.append(fp4[entry.mtLevelInDLT], tArray)
                n4[entry.mtLevelInDLT] = np.append(n4[entry.mtLevelInDLT], gArray)
            else:
                fp4[entry.mtLevelInDLT] = tArray
                n4[entry.mtLevelInDLT] = gArray



    falsePositives1Krandom1 = {}
    numberOfWitFromDLT1Krandom1 = {}
    errorfp1 = {}
    errorn1 = {}

    falsePositives1Krandom01 = {}
    numberOfWitFromDLT1Krandom01 = {}
    errorfp2 = {}
    errorn2 = {}

    falsePositives1Krandom001 = {}
    numberOfWitFromDLT1Krandom001 = {}
    errorfp3 = {}
    errorn3 = {}


    falsePositives1Krandom0001 = {}
    numberOfWitFromDLT1Krandom0001 = {}
    errorfp4 = {}
    errorn4 = {}


    for key, value in fp1.items():
        falsePositives1Krandom1[key] = np.mean(value)
        errorfp1[key]=sem(value)

    for key, value in fp2.items():
        falsePositives1Krandom01[key] = np.mean(value)
        errorfp2[key] = sem(value)

    for key, value in fp3.items():
        falsePositives1Krandom001[key] = np.mean(value)
        errorfp3[key] = sem(value)


    for key, value in fp4.items():
        falsePositives1Krandom0001[key] = np.mean(value)
        errorfp4[key] = sem(value)

    for key, value in n1.items():
        numberOfWitFromDLT1Krandom1[key] = np.mean(value)
        errorn1[key] = sem(value)

    for key, value in n2.items():
        numberOfWitFromDLT1Krandom01[key] = np.mean(value)
        errorn2[key] = sem(value)

    for key, value in n3.items():
        numberOfWitFromDLT1Krandom001[key] = np.mean(value)
        errorn3[key] = sem(value)

    for key, value in n4.items():
        numberOfWitFromDLT1Krandom0001[key] = np.mean(value)
        errorn4[key] = sem(value)


    falsePositives1Krandom1 = dict(sorted(falsePositives1Krandom1.items()))
    falsePositives1Krandom01 = dict(sorted(falsePositives1Krandom01.items()))
    falsePositives1Krandom001 = dict(sorted(falsePositives1Krandom001.items()))
    falsePositives1Krandom0001 = dict(sorted(falsePositives1Krandom0001.items()))
    x1points = np.array(list(falsePositives1Krandom01.keys()))
    y1points = np.array(list(falsePositives1Krandom1.values()))
    y2points = np.array(list(falsePositives1Krandom01.values()))
    y3points = np.array(list(falsePositives1Krandom001.values()))
    y4points = np.array(list(falsePositives1Krandom0001.values()))

    # print(falsePositives1Krandom1)
    # print(numberOfWitFromDLT1Krandom1)
    # print(x1points)
    # print(y1points)
    print("revoked VCs: ", revokedVCs, "\t revocation mode: ",revocationMode)
    x1points = np.delete(x1points,[0, 2, 4, 6,8,10,12])
    y1points = np.delete(y1points, [0, 2, 4, 6,8,10,12])
    y2points = np.delete(y2points, [0, 2, 4, 6,8,10,12])
    y3points = np.delete(y3points, [0, 2, 4, 6,8,10,12])
    y4points = np.delete(y4points, [0, 2, 4, 6,8,10,12])

    y1points = np.ceil(y1points)



    y2points = np.ceil(y2points)
    y3points = np.ceil(y3points)
    y4points = np.ceil(y4points)

    print("y1: ", y1points)
    print("y2: ", y2points)
    print("y3: ", y3points)
    print("y4: ", y4points)

    errorfp1 = dict(sorted(errorfp1.items()))
    errorfp2 = dict(sorted(errorfp2.items()))
    errorfp3 = dict(sorted(errorfp3.items()))
    errorfp4 = dict(sorted(errorfp4.items()))
    ey1points = np.array(list(errorfp1.values()))
    ey2points = np.array(list(errorfp2.values()))
    ey3points = np.array(list(errorfp3.values()))
    ey4points = np.array(list(errorfp4.values()))
    ey1points = np.delete(ey1points, [0, 2, 4, 6,8,10,12])
    ey2points = np.delete(ey2points, [0, 2, 4, 6,8,10,12])
    ey3points = np.delete(ey3points, [0, 2, 4, 6,8,10,12])
    ey4points = np.delete(ey4points, [0, 2, 4, 6,8,10,12])
    print("ey1: ", ey1points)
    print("ey2: ", ey2points)
    print("ey3: ", ey3points)
    print("ey4: ", ey4points)

    errorn1 = dict(sorted(errorn1.items()))
    errorn2 = dict(sorted(errorn2.items()))
    errorn3 = dict(sorted(errorn3.items()))
    errorn4 = dict(sorted(errorn4.items()))
    ez1points = np.array(list(errorn1.values()))
    ez2points = np.array(list(errorn2.values()))
    ez3points = np.array(list(errorn3.values()))
    ez4points = np.array(list(errorn4.values()))
    ez1points = np.delete(ez1points, [0, 2, 4, 6,8,10,12])
    ez2points = np.delete(ez2points, [0, 2, 4, 6,8,10,12])
    ez3points = np.delete(ez3points, [0, 2, 4, 6,8,10,12])
    ez4points = np.delete(ez4points, [0, 2, 4, 6,8,10,12])


    numberOfWitFromDLT1Krandom1 = dict(sorted(numberOfWitFromDLT1Krandom1.items()))
    numberOfWitFromDLT1Krandom01 = dict(sorted(numberOfWitFromDLT1Krandom01.items()))
    numberOfWitFromDLT1Krandom001 = dict(sorted(numberOfWitFromDLT1Krandom001.items()))
    numberOfWitFromDLT1Krandom0001 = dict(sorted(numberOfWitFromDLT1Krandom0001.items()))



    z1points = np.array(list(numberOfWitFromDLT1Krandom1.values()))

    z2points = np.array(list(numberOfWitFromDLT1Krandom01.values()))
    z3points = np.array(list(numberOfWitFromDLT1Krandom001.values()))
    z4points = np.array(list(numberOfWitFromDLT1Krandom0001.values()))

    z1points = np.delete(z1points, [0, 2, 4, 6,8,10,12])
    z2points = np.delete(z2points, [0, 2, 4, 6,8,10,12])
    z3points = np.delete(z3points, [0, 2, 4, 6,8,10,12])
    z4points = np.delete(z4points, [0, 2, 4, 6,8,10,12])

    z1points = np.ceil(z1points)
    z2points = np.ceil(z2points)
    z3points = np.ceil(z3points)
    z4points = np.ceil(z4points)

    print("z1: ", z1points)
    print("z2: ", z2points)
    print("z3: ", z3points)
    print("z4: ", z4points)
    print("ez1: ", ez1points)
    print("ez2: ", ez2points)
    print("ez3: ", ez3points)
    print("ez4: ", ez4points)


    y1points = y1points - z1points
    y2points = y2points - z2points
    print("y1-z1:", y1points)
    print("y2-z2:", y2points)

    y3points = y3points - z3points
    y4points = y4points - z4points
    print("y3-z3:", y3points)
    print("y4-z4:", y4points)

    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    fig, ax = plt.subplots(layout='constrained')
    my_base = 10


    # y_major = LogLocator(base=my_base)
    # y_minor = LogLocator(base=my_base, subs=[10,20, 30, 40, 50])
    # ax.set_yscale("log", base=my_base)
    # ax.yaxis.set_major_locator(y_major)
    # ax.yaxis.set_minor_locator(y_minor)
    # ax.set_ylim(0, 1000)


    if totalVCs=="10K":
        ax.set_yscale('log')
        ax.yaxis.set_major_formatter(ScalarFormatter())
        ax.minorticks_off()
        ax.set_yticks([0,1,2,5,10,50,100,200,500,1000])

    if totalVCs=="100K" and revocationMode=="oldest":
        ax.set_yscale('log')
        ax.yaxis.set_major_formatter(ScalarFormatter())
        ax.minorticks_off()
        ax.set_yticks([0, 10, 100, 1000, 10000])

    if totalVCs=="100K" and revocationMode=="random":
        ax.set_yscale('log')
        ax.set_ylim(1, 15000)
        ax.yaxis.set_major_formatter(ScalarFormatter())
        ax.minorticks_off()
        # ax.set_yticks(np.arange(0, 10000, step=1000))
        ax.set_yticks([10, 100, 1000, 10000])

    barWidth = 0.21
    br1 = np.arange(len(x1points))
    br2 = [x + barWidth for x in br1]
    br3 = [x + barWidth for x in br2]
    br4 = [x + barWidth for x in br3]

    yRange = [i for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]
    ylabel = [str(i) for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]

    plt.bar(br1, y1points, color='#614415', width=barWidth, hatch="//",
            edgecolor='grey', label='fpr=0.1: false positives')
    plt.errorbar(br1, y1points, color='#010C10', yerr=ey1points, fmt="o",capsize=3, capthick=1)

    plt.bar(br1, z1points, color='#555056', width=barWidth, bottom=y1points, hatch='o',
            edgecolor='grey')
    plt.errorbar(br1, z1points, color='#C44B64', yerr=ez1points, fmt="*", capsize=3, capthick=1)

    plt.bar(br2, y2points, color='#0072b2', width=barWidth, hatch="\\",
            edgecolor='grey', label='fpr=0.01: false positives')
    plt.errorbar(br2, y2points, color='#010C10', yerr=ey2points, fmt="o", capsize=3, capthick=1)


    plt.bar(br2, z2points, color='#555056', width=barWidth, bottom=y2points, hatch='o',
            edgecolor='grey')
    plt.errorbar(br2, z2points, color='#C44B64', yerr=ez2points, fmt="*", capsize=3, capthick=1)


    plt.bar(br3, y3points, color='#d55e00', width=barWidth, hatch='-',
            edgecolor='grey', label='fpr=0.001: false positives', yerr=ey3points)

    plt.errorbar(br3, y3points, color='#010C10', yerr=ey3points, fmt="o", capsize=3, capthick=1)

    plt.bar(br3, z3points, color='#555056', width=barWidth, bottom=y3points, hatch='o',
            edgecolor='grey')
    plt.errorbar(br3, z3points, color='#C44B64', yerr=ez3points, fmt="*", capsize=3, capthick=1)

    #
    #
    plt.bar(br4, y4points, color='#009e73', width=barWidth, hatch='.',
            edgecolor='grey', label='fpr=0.0001: false positives', yerr=ey4points)
    plt.errorbar(br4, y4points, color='#010C10',  yerr=ey4points, fmt="o", capsize=3, capthick=1)


    plt.bar(br4, z4points, color='#555056', width=barWidth, bottom=y4points, hatch='o',
            edgecolor='grey', label='number of witnesses from DLT')
    plt.errorbar(br4, z4points, color='#C44B64', yerr=ez4points, fmt="*", capsize=3, capthick=1)

    plt.xticks([r + barWidth for r in range(len(x1points))],
               x1points)

    title = "Total VCs:"+str(totalVCs)+" Revoked VCs:"+str(revokedVCs)+", revocation mode:"+revocationMode
    plt.title(title, font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('impact of false positive rate', font)
    plt.legend(fontsize="10", loc=6)
    filename ="graphs/result_impact_of_fpr_"+revocationMode+"_"+str(revokedVCs)+".png"
    # plt.grid(axis='y', color = 'green', linestyle = '--', linewidth = 0.5)
    plt.savefig(filename)


def plot_impact_of_fpr_with_false_positive(totalVCs, revokedVCs, revocationMode, fp):
    res = list()

    if totalVCs == "10K":
        res1 = parse_fpr_impact_entry("results_computed_10K_run1.json")
        res2 = parse_fpr_impact_entry("results_computed_10K_run2.json")
        res3 = parse_fpr_impact_entry("results_computed_10K_run3.json")

        for result in res1:
            res.append(result)

        for result in res2:
            res.append(result)

        for result in res3:
            res.append(result)

    if totalVCs == "100K":
        res1 = parse_fpr_impact_entry("results_computed_100K_run1.json")
        res2 = parse_fpr_impact_entry("results_computed_100K_run2.json")
        res3 = parse_fpr_impact_entry("results_computed_100K_run3.json")

        for result in res1:
            res.append(result)

        for result in res2:
            res.append(result)

        for result in res3:
            res.append(result)

    # for r in res:
    #     print(r)

    fp1 = {}
    n1 = {}



    for entry in res:
        t = []
        t.append(entry.numberOfActualFalsePositives)
        tArray = np.array(t)
        g = []
        g.append(entry.numberOfVCsRetrievedWitnessFromDLT)
        gArray = np.array(g)
        if entry.revokedVCs == revokedVCs and entry.falsePositiveRate == fp and entry.revocationMode == revocationMode:
            if entry.mtLevelInDLT in fp1.keys():
                fp1[entry.mtLevelInDLT] = np.append(fp1[entry.mtLevelInDLT], tArray)
                n1[entry.mtLevelInDLT] = np.append(n1[entry.mtLevelInDLT], gArray)
            else:
                fp1[entry.mtLevelInDLT] = tArray
                n1[entry.mtLevelInDLT] = gArray


    falsePositives1Krandom1 = {}
    numberOfWitFromDLT1Krandom1 = {}
    errorfp1 = {}
    errorn1 = {}


    for key, value in fp1.items():
        falsePositives1Krandom1[key] = np.mean(value)
        errorfp1[key] = sem(value)


    for key, value in n1.items():
        numberOfWitFromDLT1Krandom1[key] = np.mean(value)
        errorn1[key] = sem(value)



    falsePositives1Krandom1 = dict(sorted(falsePositives1Krandom1.items()))


    x1points = np.array(list(falsePositives1Krandom1.keys()))
    y1points = np.array(list(falsePositives1Krandom1.values()))



    # print(falsePositives1Krandom1)
    # print(numberOfWitFromDLT1Krandom1)
    # print(x1points)
    # print(y1points)
    print("revoked VCs: ", revokedVCs, "\t revocation mode: ", revocationMode)
    x1points = np.delete(x1points, [0, 2, 4, 6, 8, 10, 12])
    y1points = np.delete(y1points, [0, 2, 4, 6, 8, 10, 12])

    y1points = np.ceil(y1points)

    print("y1: ", y1points)


    errorfp1 = dict(sorted(errorfp1.items()))
    ey1points = np.array(list(errorfp1.values()))
    ey1points = np.delete(ey1points, [0, 2, 4, 6, 8, 10, 12])

    print("ey1: ", ey1points)


    errorn1 = dict(sorted(errorn1.items()))
    ez1points = np.array(list(errorn1.values()))
    ez1points = np.delete(ez1points, [0, 2, 4, 6, 8, 10, 12])

    numberOfWitFromDLT1Krandom1 = dict(sorted(numberOfWitFromDLT1Krandom1.items()))


    z1points = np.array(list(numberOfWitFromDLT1Krandom1.values()))
    z1points = np.delete(z1points, [0, 2, 4, 6, 8, 10, 12])
    z1points = np.ceil(z1points)


    print("z1: ", z1points)
    print("ez1: ", ez1points)


    y1points = y1points - z1points
    print("y1-z1:", y1points)



    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    fig, ax = plt.subplots(layout='constrained')


    # y_major = LogLocator(base=my_base)
    # y_minor = LogLocator(base=my_base, subs=[10,20, 30, 40, 50])
    # ax.set_yscale("log", base=my_base)
    # ax.yaxis.set_major_locator(y_major)
    # ax.yaxis.set_minor_locator(y_minor)
    # ax.set_ylim(0, 1000)

    # if totalVCs == "10K":
        # ax.set_yscale('log')
        # ax.yaxis.set_major_formatter(ScalarFormatter())
        # ax.minorticks_off()
        # ax.set_yticks([0, 1, 2, 5, 10, 50, 100, 200, 500, 1000])

    # if totalVCs == "100K" and revocationMode == "oldest":
        # ax.set_yscale('log')
        # ax.yaxis.set_major_formatter(ScalarFormatter())
        # ax.minorticks_off()
        # ax.set_yticks([2000, 4000, 6000, 8000, 10000])

    # if totalVCs == "100K" and revocationMode == "random":
    #     ax.set_yscale('log')
    #     ax.set_ylim(1, 15000)
    #     ax.yaxis.set_major_formatter(ScalarFormatter())
    #     ax.minorticks_off()
    #     # ax.set_yticks(np.arange(0, 10000, step=1000))
    #     ax.set_yticks([10, 100, 1000, 10000])

    barWidth = 0.7
    br1 = np.arange(len(x1points))


    yRange = [i for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]
    ylabel = [str(i) for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]

    if fp==0.1:
        colorCode = '#614415'
        errorBarColorCode = '#010C10'

    if fp==0.01:
        colorCode = '#0072b2'
        errorBarColorCode = '#010C10'

    if fp==0.001:
        colorCode = '#d55e00'
        errorBarColorCode = '#010C10'

    if fp==0.0001:
        colorCode = '#009e73'
        errorBarColorCode = '#010C10'


    plt.bar(br1, y1points, color=colorCode, width=barWidth, hatch="//",
            edgecolor='grey', label='no. of witnesses from Issuer')
    plt.errorbar(br1, y1points, color=errorBarColorCode, yerr=ey1points, fmt="o", capsize=3, capthick=1)

    plt.bar(br1, z1points, color='#555056', width=barWidth, bottom=y1points, hatch='o',
            edgecolor='grey', label='no. of witnesses from DLT')
    plt.errorbar(br1, z1points, color='#C44B64', yerr=ez1points, fmt="*", capsize=3, capthick=1)


    print(x1points)
    plt.xticks([r for r in range(len(x1points))],
               x1points)

    # title = "Total VCs:" + str(totalVCs) + " Revoked VCs:" + str(revokedVCs) + " Revocation mode:" + revocationMode+" fpr: "+str(fp)
    title = " Revocation mode:" + revocationMode
    plt.title(title, font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('impact of false positive rate', font)
    plt.legend(fontsize="12", loc=2)
    filename = "graphs/result_impact_of_fpr_" + str(fp)+"_"+revocationMode + "_" + str(revokedVCs) + ".png"
    # plt.grid(axis='y', color = 'green', linestyle = '--', linewidth = 0.5)
    plt.savefig(filename)



def plot_impact_of_fpr_10K_oldest():
    res10K1 = parse_fpr_impact_entry("results_computed_10K_run1.json")
    res10K2 = parse_fpr_impact_entry("results_computed_10K_run2.json")
    res10K3 = parse_fpr_impact_entry("results_computed_10K_run3.json")

    res = list()

    for result in res10K1:
        res.append(result)
    # for result in res10K2:
    #     res.append(result)
    # for result in res10K3:
    #     res.append(result)

    # for r in res:
    #     print(r)
    falsePositives1Krandom1 = {}
    numberOfWitFromDLT1Krandom1 = {}

    falsePositives1Krandom01 = {}
    numberOfWitFromDLT1Krandom01 = {}

    falsePositives1Krandom001 = {}
    numberOfWitFromDLT1Krandom001 = {}

    falsePositives1Krandom0001 = {}
    numberOfWitFromDLT1Krandom0001 = {}

    for entry in res:
        if entry.revokedVCs == 1000 and entry.falsePositiveRate == 0.1 and entry.revocationMode == "oldest":
            if entry.mtLevelInDLT in falsePositives1Krandom1.keys():
                falsePositives1Krandom1[entry.mtLevelInDLT] = (falsePositives1Krandom1[
                                                                   entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom1[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom1[
                                                                       entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom1[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom1[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 1000 and entry.falsePositiveRate == 0.01 and entry.revocationMode == "oldest":
            if entry.mtLevelInDLT in falsePositives1Krandom01.keys():
                falsePositives1Krandom01[entry.mtLevelInDLT] = (falsePositives1Krandom01[
                                                                    entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom01[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom01[
                                                                        entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom01[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom01[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 1000 and entry.falsePositiveRate == 0.001 and entry.revocationMode == "oldest":
            if entry.mtLevelInDLT in falsePositives1Krandom001.keys():
                falsePositives1Krandom001[entry.mtLevelInDLT] = (falsePositives1Krandom001[
                                                                     entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom001[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom001[
                                                                         entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom001[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom001[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 1000 and entry.falsePositiveRate == 0.0001 and entry.revocationMode == "oldest":
            if entry.mtLevelInDLT in falsePositives1Krandom0001.keys():
                falsePositives1Krandom0001[entry.mtLevelInDLT] = (falsePositives1Krandom0001[
                                                                      entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom0001[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom0001[
                                                                          entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom0001[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom0001[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

    falsePositives1Krandom1 = dict(sorted(falsePositives1Krandom1.items()))
    falsePositives1Krandom01 = dict(sorted(falsePositives1Krandom01.items()))
    falsePositives1Krandom001 = dict(sorted(falsePositives1Krandom001.items()))
    falsePositives1Krandom0001 = dict(sorted(falsePositives1Krandom0001.items()))
    x1points = np.array(list(falsePositives1Krandom1.keys()))
    y1points = np.array(list(falsePositives1Krandom1.values()))
    y2points = np.array(list(falsePositives1Krandom01.values()))
    y3points = np.array(list(falsePositives1Krandom001.values()))
    y4points = np.array(list(falsePositives1Krandom0001.values()))

    # print(falsePositives1Krandom1)
    # print(numberOfWitFromDLT1Krandom1)
    # print(x1points)
    # print(y1points)
    x1points = np.delete(x1points, [0, 2, 4, 6,8,10,12])
    y1points = np.delete(y1points, [0, 2, 4, 6,8,10,12])
    y2points = np.delete(y2points, [0, 2, 4, 6,8,10,12])
    y3points = np.delete(y3points, [0, 2, 4, 6,8,10,12])
    y4points = np.delete(y4points, [0, 2, 4, 6,8,10,12])

    numberOfWitFromDLT1Krandom1 = dict(sorted(numberOfWitFromDLT1Krandom1.items()))
    numberOfWitFromDLT1Krandom01 = dict(sorted(numberOfWitFromDLT1Krandom01.items()))
    numberOfWitFromDLT1Krandom001 = dict(sorted(numberOfWitFromDLT1Krandom001.items()))
    numberOfWitFromDLT1Krandom0001 = dict(sorted(numberOfWitFromDLT1Krandom0001.items()))
    z1points = np.array(list(numberOfWitFromDLT1Krandom1.values()))
    print(z1points)
    z2points = np.array(list(numberOfWitFromDLT1Krandom01.values()))
    z3points = np.array(list(numberOfWitFromDLT1Krandom001.values()))
    z4points = np.array(list(numberOfWitFromDLT1Krandom0001.values()))

    z1points = np.delete(z1points, [0, 2, 4, 6,8,10,12])
    z2points = np.delete(z2points, [0, 2, 4, 6,8,10,12])
    z3points = np.delete(z3points,[0, 2, 4, 6,8,10,12])
    z4points = np.delete(z4points, [0, 2, 4, 6,8,10,12])
    z1error = sem(z1points)
    z2error = sem(z2points)
    z3error = sem(z3points)
    z4error = sem(z4points)

    y1points = y1points - z1points
    y2points = y2points - z2points
    print("y1-z1:", y1points)
    print("y2-z2:", y2points)

    y3points = y3points - z3points
    y4points = y4points - z4points
    print("y3-z3:", y3points)
    print("y4-z4:", y4points)

    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    fig, ax = plt.subplots(layout='constrained')
    my_base = 10

    # y_major = LogLocator(base=my_base)
    # y_minor = LogLocator(base=my_base, subs=[10,20, 30, 40, 50])
    # ax.set_yscale("log", base=my_base)
    # ax.yaxis.set_major_locator(y_major)
    # ax.yaxis.set_minor_locator(y_minor)
    plt.yscale("log", base=2)
    barWidth = 0.21
    br1 = np.arange(len(x1points))
    br2 = [x + barWidth for x in br1]
    br3 = [x + barWidth for x in br2]
    br4 = [x + barWidth for x in br3]

    yRange = [i for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]
    ylabel = [str(i) for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]

    plt.bar(br1, y1points, color='#614415', width=barWidth, hatch="//",
            edgecolor='grey', label='fpr=0.1')
    plt.bar(br1, z1points, color='#555056', width=barWidth, bottom=y1points, hatch='o',
            edgecolor='grey')

    plt.bar(br2, y2points, color='#0072b2', width=barWidth, hatch="//",
            edgecolor='grey', label='fpr=0.01')
    plt.bar(br2, z2points, color='#555056', width=barWidth, bottom=y2points, hatch='o',
            edgecolor='grey')

    plt.bar(br3, y3points, color='#d55e00', width=barWidth, hatch='//',
            edgecolor='grey', label='fpr=0.001')
    plt.bar(br3, z3points, color='#555056', width=barWidth, bottom=y3points, hatch='o',
            edgecolor='grey')
    #
    #
    plt.bar(br4, y4points, color='#009e73', width=barWidth, hatch='//',
            edgecolor='grey', label='fpr=0.0001')
    plt.bar(br4, z4points, color='#555056', width=barWidth, bottom=y4points, hatch='o',
            edgecolor='grey', label='from DLT')

    plt.xticks([r + barWidth for r in range(len(x1points))],
               x1points)

    plt.title('Total VCs:10K, Revoked VCs:1K, revocation mode: oldest ', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('impact of false positive rate', font)
    plt.legend(fontsize="10", loc=1)
    plt.savefig("graphs/result_impact_of_fpr_oldest.png")


def plot_impact_of_fpr_100_random():
    res10K1 = parse_fpr_impact_entry("results_computed_10K_run1.json")
    res10K2 = parse_fpr_impact_entry("results_computed_10K_run2.json")
    res10K3 = parse_fpr_impact_entry("results_computed_10K_run3.json")

    res = list()
    #
    for result in res10K1:
        res.append(result)
    for result in res10K2:
        res.append(result)
    for result in res10K3:
        res.append(result)

    # for r in res:
    #     print(r)
    falsePositives1Krandom1 = {}
    numberOfWitFromDLT1Krandom1 = {}

    falsePositives1Krandom01 = {}
    numberOfWitFromDLT1Krandom01 = {}

    falsePositives1Krandom001 = {}
    numberOfWitFromDLT1Krandom001 = {}

    falsePositives1Krandom0001 = {}
    numberOfWitFromDLT1Krandom0001 = {}

    for entry in res:
        if entry.revokedVCs == 100 and entry.falsePositiveRate == 0.1 and entry.revocationMode == "random":
            if entry.mtLevelInDLT in falsePositives1Krandom1.keys():
                falsePositives1Krandom1[entry.mtLevelInDLT] = (falsePositives1Krandom1[
                                                                   entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom1[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom1[
                                                                       entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom1[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom1[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 100 and entry.falsePositiveRate == 0.01 and entry.revocationMode == "random":
            if entry.mtLevelInDLT in falsePositives1Krandom01.keys():
                falsePositives1Krandom01[entry.mtLevelInDLT] = (falsePositives1Krandom01[
                                                                    entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom01[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom01[
                                                                        entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom01[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom01[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 100 and entry.falsePositiveRate == 0.001 and entry.revocationMode == "random":
            if entry.mtLevelInDLT in falsePositives1Krandom001.keys():
                falsePositives1Krandom001[entry.mtLevelInDLT] = (falsePositives1Krandom001[
                                                                     entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom001[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom001[
                                                                         entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom001[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom001[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 100 and entry.falsePositiveRate == 0.0001 and entry.revocationMode == "random":
            if entry.mtLevelInDLT in falsePositives1Krandom0001.keys():
                falsePositives1Krandom0001[entry.mtLevelInDLT] = (falsePositives1Krandom0001[
                                                                      entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom0001[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom0001[
                                                                          entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom0001[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom0001[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

    falsePositives1Krandom1 = dict(sorted(falsePositives1Krandom1.items()))
    falsePositives1Krandom01 = dict(sorted(falsePositives1Krandom01.items()))
    falsePositives1Krandom001 = dict(sorted(falsePositives1Krandom001.items()))
    falsePositives1Krandom0001 = dict(sorted(falsePositives1Krandom0001.items()))
    x1points = np.array(list(falsePositives1Krandom1.keys()))
    y1points = np.array(list(falsePositives1Krandom1.values()))
    y2points = np.array(list(falsePositives1Krandom01.values()))
    y3points = np.array(list(falsePositives1Krandom001.values()))
    y4points = np.array(list(falsePositives1Krandom0001.values()))

    # print(falsePositives1Krandom1)
    # print(numberOfWitFromDLT1Krandom1)
    # print(x1points)
    # print(y1points)
    x1points = np.delete(x1points, [0, 2, 4, 6,8,10,12])
    y1points = np.delete(y1points, [0, 2, 4, 6,8,10,12])
    y2points = np.delete(y2points, [0, 2, 4, 6,8,10,12])
    y3points = np.delete(y3points, [0, 2, 4, 6,8,10,12])
    y4points = np.delete(y4points, [0, 2, 4, 6,8,10,12])

    numberOfWitFromDLT1Krandom1 = dict(sorted(numberOfWitFromDLT1Krandom1.items()))
    numberOfWitFromDLT1Krandom01 = dict(sorted(numberOfWitFromDLT1Krandom01.items()))
    numberOfWitFromDLT1Krandom001 = dict(sorted(numberOfWitFromDLT1Krandom001.items()))
    numberOfWitFromDLT1Krandom0001 = dict(sorted(numberOfWitFromDLT1Krandom0001.items()))
    z1points = np.array(list(numberOfWitFromDLT1Krandom1.values()))
    print(z1points)
    z2points = np.array(list(numberOfWitFromDLT1Krandom01.values()))
    z3points = np.array(list(numberOfWitFromDLT1Krandom001.values()))
    z4points = np.array(list(numberOfWitFromDLT1Krandom0001.values()))

    z1points = np.delete(z1points, [0, 2, 4, 6,8,10,12])
    z2points = np.delete(z2points, [0, 2, 4, 6,8,10,12])
    z3points = np.delete(z3points, [0, 2, 4, 6,8,10,12])
    z4points = np.delete(z4points, [0, 2, 4, 6,8,10,12])
    z1error = sem(z1points)
    z2error = sem(z2points)
    z3error = sem(z3points)
    z4error = sem(z4points)

    y1points = y1points - z1points
    y2points = y2points - z2points
    print("100 random: ")
    print("y1-z1:", y1points)
    print("y2-z2:", y2points)

    y3points = y3points - z3points
    y4points = y4points - z4points
    print("y3-z3:", y3points)
    print("y4-z4:", y4points)

    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    fig, ax = plt.subplots(layout='constrained')
    my_base = 10

    # y_major = LogLocator(base=my_base)
    # y_minor = LogLocator(base=my_base, subs=[10,20, 30, 40, 50])
    # ax.set_yscale("log", base=my_base)
    # ax.yaxis.set_major_locator(y_major)
    # ax.yaxis.set_minor_locator(y_minor)
    plt.yscale("log", base=2)
    barWidth = 0.21
    br1 = np.arange(len(x1points))
    br2 = [x + barWidth for x in br1]
    br3 = [x + barWidth for x in br2]
    br4 = [x + barWidth for x in br3]

    yRange = [i for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]
    ylabel = [str(i) for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]

    plt.bar(br1, y1points, color='#614415', width=barWidth, hatch="//",
            edgecolor='grey', label='fpr=0.1: false positives')
    plt.bar(br1, z1points, color='#555056', width=barWidth, bottom=y1points, hatch='o',
            edgecolor='grey')

    plt.bar(br2, y2points, color='#0072b2', width=barWidth, hatch="//",
            edgecolor='grey', label='fpr=0.01: false positives')
    plt.bar(br2, z2points, color='#555056', width=barWidth, bottom=y2points, hatch='o',
            edgecolor='grey')

    plt.bar(br3, y3points, color='#d55e00', width=barWidth, hatch='//',
            edgecolor='grey', label='fpr=0.001: false positives')
    plt.bar(br3, z3points, color='#555056', width=barWidth, bottom=y3points, hatch='o',
            edgecolor='grey')
    #
    #
    plt.bar(br4, y4points, color='#009e73', width=barWidth, hatch='//',
            edgecolor='grey', label='fpr=0.0001: false positives')
    plt.bar(br4, z4points, color='#555056', width=barWidth, bottom=y4points, hatch='o',
            edgecolor='grey', label='number of witnesses from DLT')

    plt.xticks([r + barWidth for r in range(len(x1points))],
               x1points)

    plt.title('Total VCs:10K, Revoked VCs:100, revocation mode: random ', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('impact of false positive rate', font)
    plt.legend(fontsize="10", loc=6)
    plt.savefig("graphs/result_impact_of_fpr_random_100.png")


def plot_impact_of_fpr_100_oldest():
    res10K1 = parse_fpr_impact_entry("results_computed_10K_run1.json")
    res10K2 = parse_fpr_impact_entry("results_computed_10K_run2.json")
    res10K3 = parse_fpr_impact_entry("results_computed_10K_run3.json")

    res = list()

    for result in res10K1:
        res.append(result)
    for result in res10K2:
        res.append(result)
    for result in res10K3:
        res.append(result)

    # for r in res:
    #     print(r)
    falsePositives1Krandom1 = {}
    numberOfWitFromDLT1Krandom1 = {}

    falsePositives1Krandom01 = {}
    numberOfWitFromDLT1Krandom01 = {}

    falsePositives1Krandom001 = {}
    numberOfWitFromDLT1Krandom001 = {}

    falsePositives1Krandom0001 = {}
    numberOfWitFromDLT1Krandom0001 = {}

    for entry in res:
        if entry.revokedVCs == 100 and entry.falsePositiveRate == 0.1 and entry.revocationMode == "oldest":
            if entry.mtLevelInDLT in falsePositives1Krandom1.keys():
                falsePositives1Krandom1[entry.mtLevelInDLT] = (falsePositives1Krandom1[
                                                                   entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom1[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom1[
                                                                       entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom1[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom1[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 100 and entry.falsePositiveRate == 0.01 and entry.revocationMode == "oldest":
            if entry.mtLevelInDLT in falsePositives1Krandom01.keys():
                falsePositives1Krandom01[entry.mtLevelInDLT] = (falsePositives1Krandom01[
                                                                    entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom01[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom01[
                                                                        entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom01[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom01[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 100 and entry.falsePositiveRate == 0.001 and entry.revocationMode == "oldest":
            if entry.mtLevelInDLT in falsePositives1Krandom001.keys():
                falsePositives1Krandom001[entry.mtLevelInDLT] = (falsePositives1Krandom001[
                                                                     entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom001[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom001[
                                                                         entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom001[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom001[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

        if entry.revokedVCs == 100 and entry.falsePositiveRate == 0.0001 and entry.revocationMode == "oldest":
            if entry.mtLevelInDLT in falsePositives1Krandom0001.keys():
                falsePositives1Krandom0001[entry.mtLevelInDLT] = (falsePositives1Krandom0001[
                                                                      entry.mtLevelInDLT] + entry.numberOfActualFalsePositives) / 2
                numberOfWitFromDLT1Krandom0001[entry.mtLevelInDLT] = (numberOfWitFromDLT1Krandom0001[
                                                                          entry.mtLevelInDLT] + entry.numberOfVCsRetrievedWitnessFromDLT) / 2
            else:
                falsePositives1Krandom0001[entry.mtLevelInDLT] = entry.numberOfActualFalsePositives
                numberOfWitFromDLT1Krandom0001[entry.mtLevelInDLT] = entry.numberOfVCsRetrievedWitnessFromDLT

    falsePositives1Krandom1 = dict(sorted(falsePositives1Krandom1.items()))
    falsePositives1Krandom01 = dict(sorted(falsePositives1Krandom01.items()))
    falsePositives1Krandom001 = dict(sorted(falsePositives1Krandom001.items()))
    falsePositives1Krandom0001 = dict(sorted(falsePositives1Krandom0001.items()))
    x1points = np.array(list(falsePositives1Krandom1.keys()))
    y1points = np.array(list(falsePositives1Krandom1.values()))
    y2points = np.array(list(falsePositives1Krandom01.values()))
    y3points = np.array(list(falsePositives1Krandom001.values()))
    y4points = np.array(list(falsePositives1Krandom0001.values()))

    print(falsePositives1Krandom1)
    print(numberOfWitFromDLT1Krandom1)
    print(x1points)
    print(y1points)
    x1points = np.delete(x1points, [0, 2, 4, 6,8,10,12])
    y1points = np.delete(y1points, [0, 2, 4, 6,8,10,12])
    y2points = np.delete(y2points, [0, 2, 4, 6,8,10,12])
    y3points = np.delete(y3points, [0, 2, 4, 6,8,10,12])
    y4points = np.delete(y4points, [0, 2, 4, 6,8,10,12])

    numberOfWitFromDLT1Krandom1 = dict(sorted(numberOfWitFromDLT1Krandom1.items()))
    numberOfWitFromDLT1Krandom01 = dict(sorted(numberOfWitFromDLT1Krandom01.items()))
    numberOfWitFromDLT1Krandom001 = dict(sorted(numberOfWitFromDLT1Krandom001.items()))
    numberOfWitFromDLT1Krandom0001 = dict(sorted(numberOfWitFromDLT1Krandom0001.items()))
    z1points = np.array(list(numberOfWitFromDLT1Krandom1.values()))
    print(z1points)
    z2points = np.array(list(numberOfWitFromDLT1Krandom01.values()))
    z3points = np.array(list(numberOfWitFromDLT1Krandom001.values()))
    z4points = np.array(list(numberOfWitFromDLT1Krandom0001.values()))

    z1points = np.delete(z1points, [0, 2, 4, 6,8,10,12])
    z2points = np.delete(z2points, [0, 2, 4, 6,8,10,12])
    z3points = np.delete(z3points, [0, 2, 4, 6,8,10,12])
    z4points = np.delete(z4points, [0, 2, 4, 6,8,10,12])
    z1error = sem(z1points)
    z2error = sem(z2points)
    z3error = sem(z3points)
    z4error = sem(z4points)

    y1points = y1points - z1points
    y2points = y2points - z2points
    print("y1-z1:", y1points)
    print("y2-z2:", y2points)

    y3points = y3points - z3points
    y4points = y4points - z4points
    print("y3-z3:", y3points)
    print("y4-z4:", y4points)

    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    fig, ax = plt.subplots(layout='constrained')
    my_base = 10

    # y_major = LogLocator(base=my_base)
    # y_minor = LogLocator(base=my_base, subs=[10,20, 30, 40, 50])
    # ax.set_yscale("log", base=my_base)
    # ax.yaxis.set_major_locator(y_major)
    # ax.yaxis.set_minor_locator(y_minor)
    plt.yscale("log", base=2)
    barWidth = 0.21
    br1 = np.arange(len(x1points))
    br2 = [x + barWidth for x in br1]
    br3 = [x + barWidth for x in br2]
    br4 = [x + barWidth for x in br3]

    yRange = [i for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]
    ylabel = [str(i) for i in (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000)]

    plt.bar(br1, y1points, color='#614415', width=barWidth, hatch="//",
            edgecolor='grey', label='fpr=0.1')
    plt.bar(br1, z1points, color='#555056', width=barWidth, bottom=y1points, hatch='o',
            edgecolor='grey')

    plt.bar(br2, y2points, color='#0072b2', width=barWidth, hatch="//",
            edgecolor='grey', label='fpr=0.01')
    plt.bar(br2, z2points, color='#555056', width=barWidth, bottom=y2points, hatch='o',
            edgecolor='grey')

    plt.bar(br3, y3points, color='#d55e00', width=barWidth, hatch='//',
            edgecolor='grey', label='fpr=0.001')
    plt.bar(br3, z3points, color='#555056', width=barWidth, bottom=y3points, hatch='o',
            edgecolor='grey')
    #
    #
    plt.bar(br4, y4points, color='#009e73', width=barWidth, hatch='//',
            edgecolor='grey', label='fpr=0.0001')
    plt.bar(br4, z4points, color='#555056', width=barWidth, bottom=y4points, hatch='o',
            edgecolor='grey', label='from DLT')

    plt.xticks([r + barWidth for r in range(len(x1points))],
               x1points)

    plt.title('Total VCs:10K, Revoked VCs:100, revocation mode: oldest ', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('impact of false positive rate', font)
    plt.legend(fontsize="10", loc=1)
    plt.savefig("graphs/result_impact_of_fpr_oldest_100.png")













