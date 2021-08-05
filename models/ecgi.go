package models

type Ecgi struct {
	PlmnId      PlmnId
	EutraCellId string
}

var EcgiData = Ecgi{
	PlmnId:      PlmnIdData,
	EutraCellId: "a123b4",
}
