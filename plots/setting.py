class Setting:
    def __init__(self, totalVCs, revokedVCs, falsePositiveRate, mtLevelInDLT):
        self.totalVCs = totalVCs
        self.revokedVCs = revokedVCs
        self.falsePositiveRate = falsePositiveRate
        self.mtLevelInDLT = mtLevelInDLT


    def __str__(self):
        output = "total vcs: "+str(self.totalVCs)
        output += "\t revoked vcs: "+str(self.revokedVCs)
        output += "\t false positive: "+str(self.falsePositiveRate)
        output += "\t mt level in dlt: "+ str(self.mtLevelInDLT)
        return output


    def __eq__(self, another):
        return self.totalVCs == another.totalVCs and self.revokedVCs==another.revokedVCs and self.falsePositiveRate == another.falsePositiveRate and self.mtLevelInDLT == another.mtLevelInDLT


    def __hash__(self):
        return hash(str(self.totalVCs)+ str(self.revokedVCs) + str(self.falsePositiveRate) + str(self.mtLevelInDLT))