package models

type PlmnId struct {
	Mcc string
	Mnc string
}

var PlmnIdData = PlmnId{
	Mcc: "300",
	Mnc: "400",
}
