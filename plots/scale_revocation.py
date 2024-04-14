

# [1 ,2 ,3 ,4 ,5 ,6 ,7 ,8 ,9 ,10 ,20 ,30 ,40]
import numpy as np
from matplotlib import pyplot as plt
from matplotlib.ticker import ScalarFormatter
from scipy.stats import sem

from fpr_impact import parse_fpr_impact_entry


def scale_revocations(totalVCs):
    res = list()
    total_number_of_VCs = 0
    if totalVCs == "10K":
        res1 = parse_fpr_impact_entry("results_computed_10K_run1.json")
        res2 = parse_fpr_impact_entry("results_computed_10K_run2.json")
        res3 = parse_fpr_impact_entry("results_computed_10K_run3.json")
        total_number_of_VCs = 10000
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
        total_number_of_VCs = 100000
        for result in res1:
            res.append(result)

        for result in res2:
            res.append(result)

        for result in res3:
            res.append(result)

    # for r in res:
    #     print(r)

    fp1 = {}
    fp2 = {}
    fp3 = {}
    fp4 = {}

    revocationPercentages = [1 ,2 ,3 ,4 ,5 ,6 ,7 ,8 ,9 ,10 ,20 ,30 ,40]


    for entry in res:
        t = []
        t.append(entry.numberOfActualFalsePositives)
        tArray = np.array(t)
        if entry.falsePositiveRate == 0.1:
            if entry.revokedVCs in fp1.keys():
                fp1[entry.revokedVCs] = np.append(fp1[entry.revokedVCs], tArray)
            else:
                fp1[entry.revokedVCs] = tArray

        if entry.falsePositiveRate == 0.01:
            if entry.revokedVCs in fp2.keys():
                fp2[entry.revokedVCs] = np.append(fp2[entry.revokedVCs], tArray)
            else:
                fp2[entry.revokedVCs] = tArray

        if entry.falsePositiveRate == 0.001:
            if entry.revokedVCs in fp3.keys():
                fp3[entry.revokedVCs] = np.append(fp3[entry.revokedVCs], tArray)
            else:
                fp3[entry.revokedVCs] = tArray

        if entry.falsePositiveRate == 0.0001:
            if entry.revokedVCs in fp4.keys():
                fp4[entry.revokedVCs] = np.append(fp4[entry.revokedVCs], tArray)
            else:
                fp4[entry.revokedVCs] = tArray


    falsePositives1 ={}
    errorfp1 = {}


    falsePositives2 = {}
    errorfp2 = {}


    falsePositives3 = {}
    errorfp3 = {}


    falsePositives4 = {}
    errorfp4 = {}


    for key, value in fp1.items():
        falsePositives1[key] = np.mean(value)
        errorfp1[key] = sem(value)

    for key, value in fp2.items():
        falsePositives2[key] = np.mean(value)
        errorfp2[key] = sem(value)

    for key, value in fp3.items():
        falsePositives3[key] = np.mean(value)
        errorfp3[key] = sem(value)

    for key, value in fp4.items():
        falsePositives4[key] = np.mean(value)
        errorfp4[key] = sem(value)


    falsePositives1 = dict(sorted(falsePositives1.items()))
    falsePositives2 = dict(sorted(falsePositives2.items()))
    falsePositives3 = dict(sorted(falsePositives3.items()))
    falsePositives4 = dict(sorted(falsePositives4.items()))
    x1points = np.array(list(falsePositives1.keys()))
    y1points = np.array(list(falsePositives1.values()))
    y2points = np.array(list(falsePositives2.values()))
    y3points = np.array(list(falsePositives3.values()))
    y4points = np.array(list(falsePositives4.values()))

    # print(falsePositives1Krandom1)
    # print(numberOfWitFromDLT1Krandom1)
    # print(x1points)
    # print(y1points)

    #
    # y1points = np.ceil(y1points)
    # y2points = np.ceil(y2points)
    # y3points = np.ceil(y3points)
    # y4points = np.ceil(y4points)
    x1points = np.delete(x1points, [6,8,9])
    y1points = np.delete(y1points, [6,8,9])
    y2points = np.delete(y2points, [6,8,9])
    y3points = np.delete(y3points,[6,8,9])
    y4points = np.delete(y4points, [6,8,9])
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
    ey1points = np.delete(ey1points, [6,8,9])
    ey2points = np.delete(ey2points, [6,8,9])
    ey3points = np.delete(ey3points,[6,8,9])
    ey4points = np.delete(ey4points, [6,8,9])
    print("ey1: ", ey1points)
    print("ey2: ", ey2points)
    print("ey3: ", ey3points)
    print("ey4: ", ey4points)

    # yRange = np.linspace(start=0, stop=max(costsfor00001.values()),
    #                      num=25)
    #
    x1points = x1points/1000
    x1points = np.ceil(x1points)
    fig, ax = plt.subplots(layout='constrained')
    if totalVCs == "100K":
        ax.set_xscale('log')
        ax.set_xlim(1, 50)
        ax.xaxis.set_major_formatter(ScalarFormatter())
        ax.minorticks_off()
        # ax.set_yticks(np.arange(0, 10000, step=1000))
        xlabel = [str(i) + "K" for i in [1, 2, 3, 4, 5, 7, 10, 20, 30, 40]]
        ax.set_xticks([1, 2, 3, 4, 5, 7, 10, 20, 30, 40], xlabel)

        ax.set_yscale('log')
        ax.set_ylim(1, 15000)
        ax.yaxis.set_major_formatter(ScalarFormatter())
        # ax.set_yticks(np.arange(0, 10000, step=1000))
        ax.set_yticks([1, 10, 100, 1000, 10000])

    # print(yRange)

    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}

    plt.errorbar(x1points, y1points, marker='o', color='#614415', label=r'$\it{p}:0.1$', yerr=ey1points)
    plt.errorbar(x1points, y2points, marker='d', color='#0072b2', label=r'$\it{p}:0.01$', yerr=ey2points)
    plt.errorbar(x1points, y3points, marker='*', color='#d55e00', label=r'$\it{p}:0.001$', yerr=ey3points)
    plt.errorbar(x1points, y4points, marker='+', color='#009e73', label=r'$\it{p}:0.0001$', yerr=ey4points)
    # plt.yticks(yRange, ylabel)
    # plt.xticks(x1points, xlabel)
    plt.title(r'$\it{n}$: '+str(totalVCs), font)
    plt.xlabel(r'$\it{r}$- number of revoked VCs', font)
    plt.ylabel('number of false positives', font)
    plt.legend(fontsize="13")
    plt.savefig("graphs/result_scalability_revocation.png")


def scale_revocations_fixed_bloomFilter_linear_scale_plot(totalVCs):


    if totalVCs == "100K":
        res = parse_fpr_impact_entry("results_revocation_scalability_100K.json")
        # res = parse_fpr_impact_entry("results_revocation_scalability_100K.json")

    for r in res:
        print(r)

    fp1 = {}
    fp2 = {}
    fp3 = {}
    fp4 = {}

    revocationPercentages = [10 ,20 ,30 ,40 ,50 ,60 ,70 ,80 ,90 ,100]


    for entry in res:
        t = []
        t.append(entry.numberOfActualFalsePositives)
        tArray = np.array(t)
        if entry.falsePositiveRate == 0.1:
            if entry.revokedVCs in fp1.keys():
                fp1[entry.revokedVCs] = np.append(fp1[entry.revokedVCs], tArray)
            else:
                fp1[entry.revokedVCs] = tArray

        if entry.falsePositiveRate == 0.01:
            if entry.revokedVCs in fp2.keys():
                fp2[entry.revokedVCs] = np.append(fp2[entry.revokedVCs], tArray)
            else:
                fp2[entry.revokedVCs] = tArray

        if entry.falsePositiveRate == 0.001:
            if entry.revokedVCs in fp3.keys():
                fp3[entry.revokedVCs] = np.append(fp3[entry.revokedVCs], tArray)
            else:
                fp3[entry.revokedVCs] = tArray

        if entry.falsePositiveRate == 0.0001:
            if entry.revokedVCs in fp4.keys():
                fp4[entry.revokedVCs] = np.append(fp4[entry.revokedVCs], tArray)
            else:
                fp4[entry.revokedVCs] = tArray


    falsePositives1 ={}
    errorfp1 = {}


    falsePositives2 = {}
    errorfp2 = {}


    falsePositives3 = {}
    errorfp3 = {}


    falsePositives4 = {}
    errorfp4 = {}


    for key, value in fp1.items():
        falsePositives1[key] = np.mean(value)
        errorfp1[key] = sem(value)

    for key, value in fp2.items():
        falsePositives2[key] = np.mean(value)
        errorfp2[key] = sem(value)

    for key, value in fp3.items():
        falsePositives3[key] = np.mean(value)
        errorfp3[key] = sem(value)

    for key, value in fp4.items():
        falsePositives4[key] = np.mean(value)
        errorfp4[key] = sem(value)


    falsePositives1 = dict(sorted(falsePositives1.items()))
    falsePositives2 = dict(sorted(falsePositives2.items()))
    falsePositives3 = dict(sorted(falsePositives3.items()))
    falsePositives4 = dict(sorted(falsePositives4.items()))
    x1points = np.array(list(falsePositives1.keys()))
    y1points = np.array(list(falsePositives1.values()))
    y2points = np.array(list(falsePositives2.values()))
    y3points = np.array(list(falsePositives3.values()))
    y4points = np.array(list(falsePositives4.values()))

    # print(falsePositives1Krandom1)
    # print(numberOfWitFromDLT1Krandom1)
    # print(x1points)
    # print(y1points)

    #
    # y1points = np.ceil(y1points)
    # y2points = np.ceil(y2points)
    # y3points = np.ceil(y3points)
    # y4points = np.ceil(y4points)

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

    print("ey1: ", ey1points)
    print("ey2: ", ey2points)
    print("ey3: ", ey3points)
    print("ey4: ", ey4points)

    # yRange = np.linspace(start=0, stop=max(costsfor00001.values()),
    #                      num=25)
    #
    x1points = x1points/1000
    x1points = np.ceil(x1points)
    fig, ax = plt.subplots(layout='constrained')
    if totalVCs == "100K":
        ax.minorticks_off()
        # ax.set_yscale('log')
        # ax.set_ylim(1, 12000)
        # ax.yaxis.set_major_formatter(ScalarFormatter())
        # ax.set_yticks([1, 5, 10, 100, 500, 1000, 5000, 10000])
        # ax.set_yticks([1000, 3000, 5000, 7000, 9000, 11000])

        ax.xaxis.set_major_formatter(ScalarFormatter())
        xlabel = [str(i) + "K" for i in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]]
        ax.set_xticks([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], xlabel)

    # print(yRange)

    # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
    font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}

    plt.errorbar(x1points, y1points, marker='o', color='#614415', label=r'$\it{p}:0.1$', yerr=ey1points)
    # plt.errorbar(x1points, y2points, marker='d', color='#0072b2', label=r'$\it{p}:0.01$', yerr=ey2points)
    # plt.errorbar(x1points, y3points, marker='*', color='#d55e00', label=r'$\it{p}:0.001$', yerr=ey3points)
    # plt.errorbar(x1points, y4points, marker='+', color='#009e73', label=r'$\it{p}:0.0001$', yerr=ey4points)
    # plt.yticks(yRange, ylabel)
    # plt.xticks(x1points, xlabel)
    plt.title(r'$\it{n}$: '+str(totalVCs)+", BF capacity: 20K", font)
    plt.xlabel(r'$\it{r}$- number of revoked VCs', font)
    plt.ylabel('number of false positives', font)
    plt.legend(fontsize="13")
    plt.savefig("graphs/result_scalability_revocation_fixed_bloomfilter_linear_scale_plot_0.1_100K_20K.png")


def scale_revocations_fixed_bloomFilter(totalVCs):

        if totalVCs == "100K":
            res = parse_fpr_impact_entry("results_revocation_scalability_100K.json")

        for r in res:
            print(r)

        fp1 = {}
        fp2 = {}
        fp3 = {}
        fp4 = {}

        revocationPercentages = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]

        for entry in res:
            t = []
            t.append(entry.numberOfActualFalsePositives)
            tArray = np.array(t)
            if entry.falsePositiveRate == 0.1:
                if entry.revokedVCs in fp1.keys():
                    fp1[entry.revokedVCs] = np.append(fp1[entry.revokedVCs], tArray)
                else:
                    fp1[entry.revokedVCs] = tArray

            if entry.falsePositiveRate == 0.01:
                if entry.revokedVCs in fp2.keys():
                    fp2[entry.revokedVCs] = np.append(fp2[entry.revokedVCs], tArray)
                else:
                    fp2[entry.revokedVCs] = tArray

            if entry.falsePositiveRate == 0.001:
                if entry.revokedVCs in fp3.keys():
                    fp3[entry.revokedVCs] = np.append(fp3[entry.revokedVCs], tArray)
                else:
                    fp3[entry.revokedVCs] = tArray

            if entry.falsePositiveRate == 0.0001:
                if entry.revokedVCs in fp4.keys():
                    fp4[entry.revokedVCs] = np.append(fp4[entry.revokedVCs], tArray)
                else:
                    fp4[entry.revokedVCs] = tArray

        falsePositives1 = {}
        errorfp1 = {}

        falsePositives2 = {}
        errorfp2 = {}

        falsePositives3 = {}
        errorfp3 = {}

        falsePositives4 = {}
        errorfp4 = {}

        for key, value in fp1.items():
            falsePositives1[key] = np.mean(value)
            errorfp1[key] = sem(value)

        for key, value in fp2.items():
            falsePositives2[key] = np.mean(value)
            errorfp2[key] = sem(value)

        for key, value in fp3.items():
            falsePositives3[key] = np.mean(value)
            errorfp3[key] = sem(value)

        for key, value in fp4.items():
            falsePositives4[key] = np.mean(value)
            errorfp4[key] = sem(value)

        falsePositives1 = dict(sorted(falsePositives1.items()))
        falsePositives2 = dict(sorted(falsePositives2.items()))
        falsePositives3 = dict(sorted(falsePositives3.items()))
        falsePositives4 = dict(sorted(falsePositives4.items()))
        x1points = np.array(list(falsePositives1.keys()))
        y1points = np.array(list(falsePositives1.values()))
        y2points = np.array(list(falsePositives2.values()))
        y3points = np.array(list(falsePositives3.values()))
        y4points = np.array(list(falsePositives4.values()))

        # print(falsePositives1Krandom1)
        # print(numberOfWitFromDLT1Krandom1)
        # print(x1points)
        # print(y1points)

        #
        # y1points = np.ceil(y1points)
        # y2points = np.ceil(y2points)
        # y3points = np.ceil(y3points)
        # y4points = np.ceil(y4points)

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

        print("ey1: ", ey1points)
        print("ey2: ", ey2points)
        print("ey3: ", ey3points)
        print("ey4: ", ey4points)

        # yRange = np.linspace(start=0, stop=max(costsfor00001.values()),
        #                      num=25)
        #
        x1points = x1points / 1000
        x1points = np.ceil(x1points)
        fig, ax = plt.subplots(layout='constrained')
        if totalVCs == "100K":
            ax.minorticks_off()
            ax.set_yscale('log')
            ax.set_ylim(1, 12000)
            ax.yaxis.set_major_formatter(ScalarFormatter())
            ax.set_yticks([1, 5, 10, 100, 500, 1000, 5000, 10000])

            ax.xaxis.set_major_formatter(ScalarFormatter())

            # ax.set_yticks(np.arange(0, 10000, step=1000))
            xlabel = [str(i) + "K" for i in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]]
            ax.set_xticks([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], xlabel)

        # print(yRange)

        # font = {'fontname':'Times New Roman', 'color': 'darkred', 'size': 10}
        font = {'fontname': 'Times New Roman', 'size': 15, 'weight': 'bold'}

        plt.errorbar(x1points, y1points, marker='o', color='#614415', label=r'$\it{p}:0.1$', yerr=ey1points)
        plt.errorbar(x1points, y2points, marker='d', color='#0072b2', label=r'$\it{p}:0.01$', yerr=ey2points)
        plt.errorbar(x1points, y3points, marker='*', color='#d55e00', label=r'$\it{p}:0.001$', yerr=ey3points)
        plt.errorbar(x1points, y4points, marker='+', color='#009e73', label=r'$\it{p}:0.0001$', yerr=ey4points)
        # plt.yticks(yRange, ylabel)
        # plt.xticks(x1points, xlabel)
        plt.title(r'$\it{n}$: ' + str(totalVCs) + ", BF capacity: 10K", font)
        plt.xlabel(r'$\it{r}$- number of revoked VCs', font)
        plt.ylabel('number of false positives', font)
        plt.legend(fontsize="13")
        plt.savefig("graphs/result_scalability_revocation_fixed_bloomfilter.png")