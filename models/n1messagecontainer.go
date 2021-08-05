package models

type N1MessageContainer struct {
	N1MessageClass    string
	N1MessageContent  string
	NfId              string
	ServiceInstanceId string
}

var N1MessageContainerData = N1MessageContainer{
	N1MessageClass:    "5GMM",
	N1MessageContent:  "N1Msg",
	ServiceInstanceId: "dummyamf",
}
