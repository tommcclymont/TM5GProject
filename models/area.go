package models

// structure for area data as defined in specification 29.571 release 15
type Area struct {
	Areacodes string
}

// create area data 1
var AreaData1 = Area{
	Areacodes: "94568",
}

// create area data 2
var AreaData2 = Area{
	Areacodes: "94569",
}
