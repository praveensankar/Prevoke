import numpy as np
import pandas as pd
import json

from matplotlib import pyplot as plt

from entry import Entry, parse_entry
from result import Result
from setting import Setting


def main():

    entries = parse_entry()
    plot_merkle_tree_accumulator_cost(entries)


def plot_merkle_tree_accumulator_cost(entries):
    mtlevels = []
    costs = []

    for entry in entries:
        mtlevels.append(entry.setting.mtLevelInDLT)
        costs.append(entry.result.mtAccumulatorPerUpdateCost/1000000)

    xpoints = np.array(mtlevels)
    ypoints = np.array(costs)
    print(xpoints)
    print(ypoints)


    costrange = np.arange(0, 1000, 100)

    print(costrange)

    plt.plot(xpoints, ypoints)
    plt.yticks(costrange)
    plt.ylabel('per revocation cost of MT Accumuluator in milliether')
    plt.savefig("graphs/cost_mt_accumulator.png")




if __name__=="__main__":
    main()


