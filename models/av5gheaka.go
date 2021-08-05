package models

type Av5GHeAka struct {
	AvType   string
	Rand     string
	XresStar string
	Autn     string
	Kausf    string
}

var Av5GHeAkaData = Av5GHeAka{
	AvType:   "5G_HE_AKA",
	Rand:     ResynchronizationInfoData.Rand,
	XresStar: "583FF3bE883Df0FAC2254E9507FC22d1",
	Autn:     "3f250f6AF8b13619E2F64Cf0AfC7E6e4",
	Kausf:    "Ad1AD801D0CeaC80e67C5Dddd34DDeC67Eb9651Aff9F8d2fC06508bD8cCE8eDa",
}
