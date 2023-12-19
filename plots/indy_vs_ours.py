import numpy as np
from matplotlib import pyplot as plt
from matplotlib.ticker import MultipleLocator, LogLocator


def plot_witness_updates_vc_indy(entries, totalVCs, totalRevocations):
    mtlevelsfor1000 = {}
    indy = 0
    witnessUpdateFromIssuer = {}
    falsePositives = {}
    witnessesFromDLT = {}
    WitnessUpdatesOfIndy = {}
    levels =[4,5,6,7,8,9]


    for level in levels:

        updates = 0
        count = 0
        fp = 0
        for entry in entries:

            if entry.setting.totalVCs == totalVCs and entry.setting.falsePositiveRate == 0.1 and entry.setting.mtLevelInDLT==level:
                updates = updates + entry.result.numberOfVCsRetrievedWitnessFromIssuer
                fp = fp + entry.result.numberOfActualFalsePositives
                count = count + 1


        updates=int(updates/count)
        fp = int(fp/count)
        witnessUpdateFromIssuer[level]=updates
        falsePositives[level]=fp
        WitnessUpdatesOfIndy[level] = (totalVCs-1)

    for level in levels:
        witnessesFromDLT[level] = falsePositives[level]-witnessUpdateFromIssuer[level]


    print(falsePositives)
    print(witnessesFromDLT)
    print(witnessUpdateFromIssuer)
    print(WitnessUpdatesOfIndy)

    x1points = np.array(levels)
    print(x1points)
    x = np.arange(len(x1points))  # the label locations
    fig, ax = plt.subplots(layout='constrained')
    # ax = fig.add_axes(x1points)
    font = {'fontname': 'Times New Roman', 'size': 18, 'weight': 'bold'}

    limit = totalVCs

    if totalVCs==5000:
        limit=10000

    my_base = 2
    yticks = []
    i =1;
    while i<=totalVCs:
        yticks.append(i)
        i = i * my_base

    yticks.append(i)
    print(yticks)

    ax.set_yscale("log", base=my_base)

    y_major = LogLocator(base=my_base)
    y_minor = LogLocator(base=my_base, subs=[10, 20, 30, 40, 50])

    ax.yaxis.set_major_locator(y_major)
    ax.yaxis.set_minor_locator(y_minor)
    ax.bar(x+0.00, list(witnessesFromDLT.values()), color = '#614415', label="two phase: witnesses from dlt")
    ax.bar(x + 0.00, list(witnessUpdateFromIssuer.values()), color='#2b4261',bottom=list(witnessesFromDLT.values()), label="two phase: witnesses from issuer")
    ax.bar(x + 0.25,list(WitnessUpdatesOfIndy.values()), color = '#3182d3', width = 0.25, label="technique in indy")
    title = "indy V two phase: "+str(totalVCs)+" vcs, "+str(totalRevocations)+" revocations"
    ax.set_title(title,font)
    ax.set_ylabel('no of vcs requried to update witnesses', font)
    ax.set_xlabel('no of merkle tree accumulator levels in DLT', font)
    ax.set_xticks(range(len(x1points)), x1points)
    ax.set_yticks(yticks)
    ax.legend()
    filename ="graphs/witness_updates_indy_vs_ours_2_"+str(totalVCs)+".png"
    plt.savefig(filename)


