import math
import os

import numpy as np
from matplotlib import pyplot as plt
from matplotlib.ticker import LogLocator, ScalarFormatter
from scipy.stats import sem
from entry import Entry, parse_entry

def plot_verification_time():


    results0 = parse_entry("results_1000_100_0.100000_0.json")
    results1 = parse_entry("results_1000_100_0.100000_1.json")
    results2 = parse_entry("results_1000_100_0.100000_2.json")
    results3 = parse_entry("results_1000_100_0.100000_3.json")
    results4 = parse_entry("results_1000_100_0.100000_4.json")
    results5 = parse_entry("results_1000_100_0.100000_5.json")
    results6 = parse_entry("results_1000_100_0.100000_6.json")
    results7 = parse_entry("results_1000_100_0.100000_7.json")
    results8 = parse_entry("results_1000_100_0.100000_8.json")
    results9 = parse_entry("results_1000_100_0.100000_9.json")
    results10 = parse_entry("results_1000_100_0.100000_10.json")

    res = list()

    for result in results0:
        res.append(result)
    for result in results1:
        res.append(result)
    for result in results2:
        res.append(result)
    for result in results3:
        res.append(result)
    for result in results4:
        res.append(result)

    for result in results5:
        res.append(result)
    for result in results6:
        res.append(result)
    for result in results7:
        res.append(result)
    for result in results8:
        res.append(result)
    for result in results9:
        res.append(result)
    for result in results10:
        res.append(result)


    validVCVerTime = {}
    revokedAndFPVCVerTime = {}
    holderwitnesscomputationtime = {}
    revocationTime = {}

    for entry in res:
        if entry.setting.mtLevelInDLT in validVCVerTime.keys():
            validVCVerTime[entry.setting.mtLevelInDLT] = (validVCVerTime[entry.setting.mtLevelInDLT] + np.mean(entry.result.verificationTimePerValidVCRawData))/2
            revocationTime[entry.setting.mtLevelInDLT] = (revocationTime[entry.setting.mtLevelInDLT] + np.mean(entry.result.revocationTimePerVCRawData)) / 2
            revokedAndFPVCVerTime[entry.setting.mtLevelInDLT] = (revokedAndFPVCVerTime[entry.setting.mtLevelInDLT] + np.mean(entry.result.verificationTimePerFalsePositiveOrRevokedVCRawData)) / 2
            holderwitnesscomputationtime[entry.setting.mtLevelInDLT] = (holderwitnesscomputationtime[entry.setting.mtLevelInDLT]+ np.mean(entry.result.avgTimeToFetchWitnessFromDLTRawData))/2
        else:
            validVCVerTime[entry.setting.mtLevelInDLT] = np.mean(entry.result.verificationTimePerValidVCRawData)
            revokedAndFPVCVerTime[entry.setting.mtLevelInDLT] = np.mean(entry.result.verificationTimePerFalsePositiveOrRevokedVCRawData)
            revocationTime[entry.setting.mtLevelInDLT] =  np.mean(entry.result.revocationTimePerVCRawData)
            if entry.result.avgTimeToFetchWitnessFromDLTRawData.size>0:
                holderwitnesscomputationtime[entry.setting.mtLevelInDLT] = np.mean(entry.result.avgTimeToFetchWitnessFromDLTRawData)
            else:
                holderwitnesscomputationtime[entry.setting.mtLevelInDLT] = 0
    print(validVCVerTime)
    print(revokedAndFPVCVerTime)

    validVCVerTime = dict(sorted(validVCVerTime.items()))
    x1points = np.array(list(validVCVerTime.keys()))
    y1points = np.array(list(validVCVerTime.values()))
    y1points = y1points * 1000
    revokedAndFPVCVerTime = dict(sorted(revokedAndFPVCVerTime.items()))
    x2points = np.array(list(revokedAndFPVCVerTime.keys()))
    y2points = np.array(list(revokedAndFPVCVerTime.values()))
    y3points = np.array(list(holderwitnesscomputationtime.values()))
    y2points = y2points * 1000
    y3points = y3points * 1000
    y2points = y2points - y3points
    y4points = np.array(list(revocationTime.values()))
    y4points = y4points * 1000
    y1error = sem(y1points)
    y2error = sem(y2points)
    y3error = sem(y3points)
    y4error = sem(y4points)
    print(x1points)
    print(y1points, "\t", y1error)
    print(y2points, "\t", y2error)
    print(y3points, "\t", y3error)
    print(y4points, "\t", y4error)
    yRange = np.linspace(start=0.01, stop=math.ceil(max(revokedAndFPVCVerTime.values())),
                           num=25)

    # ylabel = [str(i)+"secs" for i in range(yRange)]

    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    fig, ax = plt.subplots(layout='constrained')
    my_base = 10
    ax.set_yscale('log')
    ax.yaxis.set_major_formatter(ScalarFormatter())
    ax.minorticks_off()
    ax.set_yticks([0,10,20,30,40,50,100,200,500,1000,2000,3000])


    barWidth = 0.25
    br1 = np.arange(len(x1points))
    br2 = [x + barWidth for x in br1]
    br3 = [x + barWidth for x in br2]

    yRange =[i for i in (1,10,100,200,400,600,800,1000,1500,2000,3000)]
    ylabel = [str(i) for i in (1,10,100,200,400,600,800,1000,1500,2000,3000)]

    plt.bar(br1, y4points, color='#f0e442', width=barWidth, hatch="//",
            edgecolor='grey', label='revocation time', yerr=y4error)
    plt.bar(br2, y1points, color='#0072b2', width=barWidth,
            edgecolor='grey', label='Phase 1- time by Verifier', yerr=y1error)
    plt.bar(br3, y2points, color='#d55e00', width=barWidth, hatch='-',
            edgecolor='grey', label='Phase 2- time  by Verifier excl. MTAcc retrieval at Holder', yerr=y2error)
    # plt.errorbar(br2, y2points, yerr=y2error, fmt="o", color="#3b3b3b")
    plt.bar(br3, y3points, color='#009e73', width=barWidth, bottom=y2points, hatch='o',
           edgecolor='grey', label='Phase 2- time to retrieve MTAcc by Holder')

    plt.xticks([r + barWidth for r in range(len(x1points))],
               x1points)
    # plt.yticks(yRange)
    plt.title('Total VCs:1000, Revoked VCs:100, fpr=0.1 ', font)
    plt.xlabel('merkle tree accumulator levels stored in DLT', font)
    plt.ylabel('time (in milli seconds)', font)
    plt.legend(fontsize="12")
    plt.savefig("graphs/result_verification_time_bar.png")



def plot_verification_time2():


    results10 = parse_entry("results_1000_100_0.100000_0.json")
    results11 = parse_entry("results_1000_100_0.010000_0.json")
    results12 = parse_entry("results_1000_100_0.001000_0.json")
    results13 = parse_entry("results_1000_100_0.000100_0.json")

    results20 = parse_entry("results_1000_200_0.100000_0.json")
    results21 = parse_entry("results_1000_200_0.010000_0.json")
    results22 = parse_entry("results_1000_200_0.001000_0.json")
    results23 = parse_entry("results_1000_200_0.000100_0.json")


    results30 = parse_entry("results_1000_300_0.100000_0.json")
    results31 = parse_entry("results_1000_300_0.010000_0.json")
    results32 = parse_entry("results_1000_300_0.001000_0.json")
    results33 = parse_entry("results_1000_300_0.000100_0.json")

    results40 = parse_entry("results_1000_400_0.100000_0.json")
    results41 = parse_entry("results_1000_400_0.010000_0.json")
    results42 = parse_entry("results_1000_400_0.001000_0.json")
    results43 = parse_entry("results_1000_400_0.000100_0.json")

    results50 = parse_entry("results_1000_500_0.100000_0.json")
    results51 = parse_entry("results_1000_500_0.010000_0.json")
    results52 = parse_entry("results_1000_500_0.001000_0.json")
    results53 = parse_entry("results_1000_500_0.000100_0.json")

    res = list()

    for result in results10:
        res.append(result)
    for result in results11:
        res.append(result)
    for result in results12:
        res.append(result)
    for result in results13:
        res.append(result)
    for result in results20:
        res.append(result)
    for result in results21:
        res.append(result)
    for result in results22:
        res.append(result)
    for result in results23:
        res.append(result)

    for result in results30:
        res.append(result)
    for result in results31:
        res.append(result)
    for result in results32:
        res.append(result)
    for result in results33:
        res.append(result)
    for result in results40:
        res.append(result)
    for result in results41:
        res.append(result)
    for result in results42:
        res.append(result)
    for result in results43:
        res.append(result)
    for result in results50:
        res.append(result)
    for result in results51:
        res.append(result)
    for result in results52:
        res.append(result)
    for result in results53:
        res.append(result)


    validVCVerTime1 = {}
    validVCVerTime01 = {}
    validVCVerTime001 = {}
    validVCVerTime0001 = {}






    for entry in res:
        if entry.setting.falsePositiveRate==0.1:
            validVCVerTime1[entry.setting.revokedVCs] = entry.result.verificationTimePerValidVC


        if entry.setting.falsePositiveRate==0.01:
            validVCVerTime01[entry.setting.revokedVCs] = entry.result.verificationTimePerValidVC


        if entry.setting.falsePositiveRate==0.001:
            validVCVerTime001[entry.setting.revokedVCs] = entry.result.verificationTimePerValidVC


        if entry.setting.falsePositiveRate==0.0001:
            validVCVerTime0001[entry.setting.revokedVCs] = entry.result.verificationTimePerValidVC




    xpoints = np.array(list(validVCVerTime1.keys()))
    validVCVerTime1 = dict(sorted(validVCVerTime1.items()))
    validVCVerTime01 = dict(sorted(validVCVerTime01.items()))
    validVCVerTime001 = dict(sorted(validVCVerTime001.items()))
    validVCVerTime0001 = dict(sorted(validVCVerTime0001.items()))




    print(validVCVerTime1)
    print(validVCVerTime01)
    print(validVCVerTime001)
    print(validVCVerTime0001)



    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    fig, ax = plt.subplots(layout='constrained')

    barWidth = 0.2
    br1 = np.arange(len(validVCVerTime1))
    br2 = [x + barWidth for x in br1]
    br3 = [x + barWidth for x in br2]
    br4 = [x + barWidth for x in br3]


    plt.bar(br1, validVCVerTime1.values(), color='#0072b2', width=barWidth,
            edgecolor='grey', label='fpr=0.1')
    plt.bar(br2, validVCVerTime01.values(), color = '#d55e00', width=barWidth,
            edgecolor='grey', label='fpr=0.01')
    plt.bar(br3, validVCVerTime001.values(), color='#009e73', width=barWidth,
            edgecolor='grey', label='fpr=0.001')
    plt.bar(br4, validVCVerTime0001.values(), color = '#f0e442', width=barWidth,
            edgecolor='grey', label='fpr=0.0001')



    plt.xticks([r + barWidth for r in range(len(validVCVerTime1))],
               validVCVerTime1.keys())
    plt.title('Total VCs:1000, `Z` levels: 0', font)
    plt.xlabel('revoked VCs', font)

    plt.ylabel('phase 1 time (in seconds)', font)
    plt.legend(fontsize="13")
    plt.savefig("graphs/result_phase1_time.png")




def plot_verification_time3():


    results10 = parse_entry("results_1000_100_0.100000_0.json")
    results11 = parse_entry("results_1000_100_0.010000_0.json")
    results12 = parse_entry("results_1000_100_0.001000_0.json")
    results13 = parse_entry("results_1000_100_0.000100_0.json")

    results20 = parse_entry("results_1000_200_0.100000_0.json")
    results21 = parse_entry("results_1000_200_0.010000_0.json")
    results22 = parse_entry("results_1000_200_0.001000_0.json")
    results23 = parse_entry("results_1000_200_0.000100_0.json")


    results30 = parse_entry("results_1000_300_0.100000_0.json")
    results31 = parse_entry("results_1000_300_0.010000_0.json")
    results32 = parse_entry("results_1000_300_0.001000_0.json")
    results33 = parse_entry("results_1000_300_0.000100_0.json")

    results40 = parse_entry("results_1000_400_0.100000_0.json")
    results41 = parse_entry("results_1000_400_0.010000_0.json")
    results42 = parse_entry("results_1000_400_0.001000_0.json")
    results43 = parse_entry("results_1000_400_0.000100_0.json")

    results50 = parse_entry("results_1000_500_0.100000_0.json")
    results51 = parse_entry("results_1000_500_0.010000_0.json")
    results52 = parse_entry("results_1000_500_0.001000_0.json")
    results53 = parse_entry("results_1000_500_0.000100_0.json")

    res = list()

    for result in results10:
        res.append(result)
    for result in results11:
        res.append(result)
    for result in results12:
        res.append(result)
    for result in results13:
        res.append(result)
    for result in results20:
        res.append(result)
    for result in results21:
        res.append(result)
    for result in results22:
        res.append(result)
    for result in results23:
        res.append(result)

    for result in results30:
        res.append(result)
    for result in results31:
        res.append(result)
    for result in results32:
        res.append(result)
    for result in results33:
        res.append(result)
    for result in results40:
        res.append(result)
    for result in results41:
        res.append(result)
    for result in results42:
        res.append(result)
    for result in results43:
        res.append(result)
    for result in results50:
        res.append(result)
    for result in results51:
        res.append(result)
    for result in results52:
        res.append(result)
    for result in results53:
        res.append(result)





    revokedAndFPVCVerTime1 = {}
    revokedAndFPVCVerTime01 = {}
    revokedAndFPVCVerTime001 = {}
    revokedAndFPVCVerTime0001 = {}



    for entry in res:
        if entry.setting.falsePositiveRate==0.1:
            revokedAndFPVCVerTime1[entry.setting.revokedVCs] = entry.result.verificationTimePerFalsePositiveOrRevokedVC

        if entry.setting.falsePositiveRate==0.01:
            revokedAndFPVCVerTime01[entry.setting.revokedVCs] = entry.result.verificationTimePerFalsePositiveOrRevokedVC

        if entry.setting.falsePositiveRate==0.001:
            revokedAndFPVCVerTime001[entry.setting.revokedVCs] = entry.result.verificationTimePerFalsePositiveOrRevokedVC

        if entry.setting.falsePositiveRate==0.0001:
            revokedAndFPVCVerTime0001[entry.setting.revokedVCs] = entry.result.verificationTimePerFalsePositiveOrRevokedVC



    xpoints = np.array(list(revokedAndFPVCVerTime0001.keys()))


    revokedAndFPVCVerTime1 = dict(sorted(revokedAndFPVCVerTime1.items()))
    revokedAndFPVCVerTime01 = dict(sorted(revokedAndFPVCVerTime01.items()))
    revokedAndFPVCVerTime001 = dict(sorted(revokedAndFPVCVerTime001.items()))
    revokedAndFPVCVerTime0001 = dict(sorted(revokedAndFPVCVerTime0001.items()))


    print(revokedAndFPVCVerTime1)
    print(revokedAndFPVCVerTime01)
    print(revokedAndFPVCVerTime001)
    print(revokedAndFPVCVerTime0001)



    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}
    fig, ax = plt.subplots(layout='constrained')

    barWidth = 0.2
    br1 = np.arange(len(revokedAndFPVCVerTime0001))
    br2 = [x + barWidth for x in br1]
    br3 = [x + barWidth for x in br2]
    br4 = [x + barWidth for x in br3]


    plt.bar(br1, revokedAndFPVCVerTime1.values(), color='#0072b2', width=barWidth,
            edgecolor='grey', label='fpr=0.1')
    plt.bar(br2, revokedAndFPVCVerTime01.values(), color = '#d55e00', width=barWidth,
            edgecolor='grey', label='fpr=0.01')
    plt.bar(br3, revokedAndFPVCVerTime001.values(), color='#009e73', width=barWidth,
            edgecolor='grey', label='fpr=0.001')
    plt.bar(br4, revokedAndFPVCVerTime0001.values(), color = '#f0e442', width=barWidth,
            edgecolor='grey', label='fpr=0.0001')



    plt.xticks([r + barWidth for r in range(len(revokedAndFPVCVerTime0001))],
               revokedAndFPVCVerTime0001.keys())
    plt.title('Total VCs:1000, `Z` levels: 0', font)
    plt.xlabel('revoked VCs', font)

    plt.ylabel('phase 2 time (in seconds)', font)
    plt.legend(fontsize="13")
    plt.savefig("graphs/result_phase2_time.png")



