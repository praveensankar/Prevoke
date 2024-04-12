import math
import os

import numpy as np
from matplotlib import pyplot as plt

from entry import Entry, parse_entry

def calculate_bbs_time():


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


    bbsProofGenTime = 0
    bbsProofVerTime = 0


    for entry in res:
        if bbsProofGenTime==0:
            bbsProofGenTime = entry.result.bbsProofGenerationTime
        else:
            bbsProofGenTime = (bbsProofGenTime + entry.result.bbsProofGenerationTime)/2

        if bbsProofVerTime==0:
            bbsProofVerTime = entry.result.bbsVerificationTime
        else:
            bbsProofVerTime = (bbsProofVerTime + entry.result.bbsVerificationTime)/2

    print("bbs proof generation time (in seconds): ", bbsProofGenTime)
    print("bbs proof verification time (in seconds): ", bbsProofVerTime)



