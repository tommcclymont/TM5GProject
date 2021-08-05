package models

// structure for UE context in SMF data as defined in specification 29.503 release 15
type PgwInfo struct {
	Dnn     string
	PgwFqdn string
	PlmnId  PlmnId
}

var PgwInfoData = PgwInfo{
	Dnn:     DNNInfo1.Dnn,
	PgwFqdn: "pgwfqdn1",
	PlmnId:  PlmnIdData,
}
