package models

type DeregistrationData struct {
	DeregReason string
	AccessType  string
}

var DeregistrationDataList = DeregistrationData{
	DeregReason: "UE_INITIAL_REGISTRATION",
	AccessType:  "3GPP_ACCESS",
}
