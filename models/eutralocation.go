package models

type EutraLocation struct {
	Tai                      Tai
	Ecgi                     Ecgi
	AgeOfLocationInformation int
	UeLocationTimestamp      string
	GeographicalInformation  string
	GeodeticInformation      string
	GlobalNgenbId            GlobalRanNodeId
}

var EutraLocationData = EutraLocation{
	Tai:                      TaiData,
	Ecgi:                     EcgiData,
	AgeOfLocationInformation: 5000,
	UeLocationTimestamp:      "2021-05-11T12:00:00",
	GeographicalInformation:  "GH7489KJD610LPO7",
	GeodeticInformation:      "GH7489KJD610LPO7",
	GlobalNgenbId:            GlobalRanNodeIdData,
}
