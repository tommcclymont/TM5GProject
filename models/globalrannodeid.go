package models

type GlobalRanNodeId struct {
	PlmnId  PlmnId
	N3IwfId string
	GNbId   string
	NgeNbId string
}

var GlobalRanNodeIdData = GlobalRanNodeId{
	PlmnId: PlmnIdData,
}
