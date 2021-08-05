package models

// structure for emergency info defined in specification 29.503 release 15
type EmergencyInfo struct {
	PgwFqdn       string
	PgwIpAddress  IpAddress
	SmfInstanceId string
}

var EmergencyInfoData = EmergencyInfo{
	PgwFqdn:       "pgwfqdn1",
	PgwIpAddress:  IpAddressData,
	SmfInstanceId: "1a800ae6-706e-474f-89de-75747f19879c",
}
