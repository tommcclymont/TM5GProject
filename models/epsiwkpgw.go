package models

type EpsIwkPgw struct {
	PgwFqdn       string
	SmfInstanceId string
}

var EpsIwkPgwData = EpsIwkPgw{
	PgwFqdn:       PgwInfoData.PgwFqdn,
	SmfInstanceId: EmergencyInfoData.SmfInstanceId,
}

var EpsIwkPgwMap = map[string]EpsIwkPgw{
	DNNInfo1.Dnn: EpsIwkPgwData,
}
