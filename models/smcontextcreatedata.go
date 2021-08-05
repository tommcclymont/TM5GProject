package models

type SmContextCreateData struct {
	Supi                string
	UnauthenticatedSupi bool
	Pei                 string
	Gpsi                string
	PduSessionId        uint
	Dnn                 string
	Snssai              Snssai
	ServingNfId         string
	Guami               Guami
	ServiceName         string
	ServingNetwork      PlmnId
	RequestType         string
	N1SmMsg             string
	AnType              string
	RatType             RatType
	UeLocation          UserLocation
	UeTimezone          string
	SmContextStatusUri  string
	PcfId               string
	SupportedFeatures   string
	SelMode             string
	BackupAmfInfo       []BackupAmfInfo
	TraceData           TraceData
	UdmGroupId          string
	RoutingIndicator    string
}

var SmContextCreateDataList = SmContextCreateData{
	Supi:                "892881507",
	UnauthenticatedSupi: false,
	Pei:                 "imei-857390213674531",
	Gpsi:                "msisdn-61738596071",
	PduSessionId:        65,
	Dnn:                 DNNInfo1.Dnn,
	Snssai:              SnssaiData,
	ServingNfId:         "Amf123",
	Guami:               GuamiData,
	ServiceName:         "dummyamf",
	ServingNetwork:      PlmnIdData,
	RequestType:         "INITAL_REQUEST",
	N1SmMsg:             "N1SmMsg",
	AnType:              "3GPP_ACCESS",
	RatType:             RatTypeData,
	UeLocation:          UserlocationData,
	UeTimezone:          "00:00+0",
	SmContextStatusUri:  "https://127.0.0.78:9090/",
	PcfId:               "4e2621a0-6b91-49dd-9b13-2d67476a6b91",
	SupportedFeatures:   "001",
	SelMode:             "VERIFIED",
	BackupAmfInfo:       []BackupAmfInfo{BackupAmfInfoData},
	TraceData:           TraceDataList,
	UdmGroupId:          "dummyudm",
	RoutingIndicator:    "1",
}
