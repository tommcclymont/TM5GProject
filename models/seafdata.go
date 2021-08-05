package models

type SeafData struct {
	NgKsi                NgKsi
	KeyAmf               KeyAmf
	Nh                   string
	Ncc                  int
	KeyAmfChangeInd      bool
	KeyAmfHderivationInd bool
}

var SeafDataList = SeafData{
	NgKsi:                NgKsiData,
	KeyAmf:               KeyAmfData,
	Nh:                   "6D3d6C76fF99fE8be6FC1346a5e63039",
	Ncc:                  1,
	KeyAmfChangeInd:      false,
	KeyAmfHderivationInd: true,
}
