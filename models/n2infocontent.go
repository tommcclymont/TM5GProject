package models

type N2InfoContent struct {
	NgapMessageType uint
	NgapIeType      string
	NgapData        string
}

var N2InfoContentData = N2InfoContent{
	NgapMessageType: 1,
	NgapIeType:      "PDU_RES_SETUP_REQ",
	NgapData:        "ngapdata",
}
