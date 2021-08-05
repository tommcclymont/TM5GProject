package models

type DnnInfo struct {
	Dnn                 string
	DefaultDnnIndicator bool
	LboRoamingAllowed   bool
	IwkEpsInd           bool
}

var DNNInfo1 = DnnInfo{
	Dnn:                 "accesspoint",
	DefaultDnnIndicator: true,
	LboRoamingAllowed:   false,
	IwkEpsInd:           true,
}
