import numpy as np
from matplotlib import pyplot as plt


def plot_fpr_vs_bfsize_and_witUpdates(entries):
    bf_size={}
    wit_updates={}
    fpr=[0.01,0.02,0.03,0.04,0.05, 0.1]
    mt_level = 6
    for p in fpr:
        size = 0
        updates = 0
        count = 0
        for entry in entries:
            if entry.setting.falsePositiveRate==p and entry.setting.totalVCs==1000 and entry.setting.mtLevelInDLT==mt_level and entry.setting.revokedVCs==100:
                size = size + entry.setting.bloomFilterSize
                updates = updates + entry.result.numberOfVCsRetrievedWitnessFromIssuer
                count = count + 1


        size = int(size / count)
        updates = int(updates / count)
        index = int(p*100)
        bf_size[index]=size
        wit_updates[index]=updates


    print("bf size: ",bf_size.values())
    print("number of witness updates ",wit_updates)

    xpoints = np.array(fpr)
    y1points = np.array(list(bf_size.values()))
    y2points = np.array(list(wit_updates.values()))

    print(xpoints)
    print(y1points)
    print(y2points)

    range = np.linspace(start=0, stop=max(bf_size.values())+50, num=len(fpr))
    font = {'fontname': 'Times New Roman', 'size': 18, 'weight': 'bold'}
    fig, ax1 = plt.subplots()
    color = '#3182d3'

    ax1.set_xlabel('false positive rate', font)
    ax1.set_ylabel('bloom filter size (in bits)', font)
    ax1.plot(xpoints, y1points, marker = 'd', color=color, label="1000 vcs, 6 - mt_level_in_dlt")
    ax1.tick_params(axis='y')
    ax1.set_xticks(fpr)

    ax2 = ax1.twinx()  # instantiate a second axes that shares the same x-axis

    color = '#614415'
    ax2.set_ylabel('witness updates fetched from issuer', font)  # we already handled the x-label with ax1

    ax2.plot(xpoints,y2points, marker = 'o', color=color, label="1000 vcs, 6 - mt_level_in_dlt")
    ax2.tick_params(axis='y')
    ax2.set_xticks(fpr)
    fig.tight_layout()  # otherwise the right y-label is slightly clipped

    lines, labels = ax1.get_legend_handles_labels()
    lines2, labels2 = ax2.get_legend_handles_labels()
    ax2.legend(lines + lines2, labels + labels2, loc=9)

    plt.savefig("graphs/impact_of_fpr.png")

