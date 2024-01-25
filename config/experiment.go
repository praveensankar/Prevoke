package config

type Experiment struct {
	TotalVCs int `json:"totalVCs"`
	RevokedVCs int `json:"revokedVCs"`
	FalsePositiveRate float64 `json:"falsePositiveRate"`
	MtLevelInDLT int `json:"mtLevelInDLT"`
	MtDepth int `json:"mtDepth"`
	MtHeight int `json:"mtHeight"`

}