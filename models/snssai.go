package models

// structure for snssai data as defined in specification 29.571 release 15
type Snssai struct {
	Sst uint
	Sd  string
}

var SnssaiData = Snssai{
	Sst: 10,
	Sd:  "EcE040",
}
