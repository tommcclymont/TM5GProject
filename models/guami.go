package models

type Guami struct {
	PlmnId PlmnId
	AmfId  string
}

var GuamiData = Guami{
	PlmnId: PlmnIdData,
	AmfId:  "Amf123",
}
