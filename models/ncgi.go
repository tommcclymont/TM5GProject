package models

type Ncgi struct {
	PlmnId   PlmnId
	NrCellId string
}

var NcgiData = Ncgi{
	PlmnId:   PlmnIdData,
	NrCellId: "a123b4",
}
