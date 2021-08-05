package models

type Tai struct {
	PlmnId PlmnId
	Tac    string
}

var TaiData = Tai{
	PlmnId: PlmnIdData,
	Tac:    "a123b4",
}
