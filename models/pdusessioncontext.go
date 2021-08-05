package models

type PduSessionContext struct {
	PduSessionId         uint
	SmContextRef         string
	SNssai               Snssai
	Dnn                  string
	AccessType           string
	AllocatedEbiList     []EbiArpmapping
	HsmfId               string
	NsInstance           string
	SmfServiceInstanceId string
}

var PduSessionContextData = PduSessionContext{
	PduSessionId:         10,
	SmContextRef:         "https://127.0.0.78:9090/",
	SNssai:               SnssaiData,
	Dnn:                  DNNInfo1.Dnn,
	AccessType:           "3GPP_ACCESS",
	AllocatedEbiList:     []EbiArpmapping{EbiArpmappingData},
	HsmfId:               "1a800ae6-706e-474f-89de-75747f19879c",
	NsInstance:           "73d753d7-9086-4515-a118-cb2b7f61e397",
	SmfServiceInstanceId: "1a800ae6-706e-474f-89de-75747f19879c",
}
