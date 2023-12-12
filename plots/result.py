import json

from setting import Setting


class Result:
    def __init__(self, mtAccumulatorPerUpdateCost, numberOfActualFalsePositives, numberOfVCsRetrievedWitnessFromIssuer,
                 numberOfVCsAffectedByMTAccumulator):
        self.mtAccumulatorPerUpdateCost = mtAccumulatorPerUpdateCost
        self.numberOfActualFalsePositives = numberOfActualFalsePositives
        self.numberOfVCsRetrievedWitnessFromIssuer= numberOfVCsRetrievedWitnessFromIssuer
        self.numberOfVCsAffectedByMTAccumulator = numberOfVCsAffectedByMTAccumulator

    def __str__(self):
        output = "mt accumulator update cost (in gwei): "+str(self.mtAccumulatorPerUpdateCost)
        output += "\t no. of false positives: "+str(self.numberOfActualFalsePositives)
        output += "\t no. of vcs retrieved witness from issuer: "+str(self.numberOfVCsRetrievedWitnessFromIssuer)
        output += "\t no. of vcs affected by updates in mt accumulator: "+ str(self.numberOfVCsAffectedByMTAccumulator)
        return output


