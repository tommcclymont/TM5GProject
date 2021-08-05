package models

// structure for UE context in SMF data as defined in specification 29.571 release 15
type PduSession struct {
	Dnn           string
	SmfInstanceId string
	PlmnId        PlmnId
}

var PduSessionData = PduSession{
	Dnn:           DNNInfo1.Dnn,
	SmfInstanceId: "1a800ae6-706e-474f-89de-75747f19879c",
	PlmnId:        PlmnIdData,
}
