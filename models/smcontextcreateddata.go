package models

type SmContextCreatedData struct {
	AllocatedEbiList     []EbiArpmapping
	SmfServiceInstanceId string
	RecoveryTime         string
	SupportedFeatures    string
}

var SmContextCreatedDataList = SmContextCreatedData{
	AllocatedEbiList:     []EbiArpmapping{EbiArpmappingData},
	SmfServiceInstanceId: "1a800ae6-706e-474f-89de-75747f19879c",
	RecoveryTime:         "00:00+0",
	SupportedFeatures:    "001",
}
