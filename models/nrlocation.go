package models

type NrLocation struct {
	Tai                      Tai
	Ncgi                     Ncgi
	AgeOfLocationInformation int
	UeLocationTimestamp      string
	GeographicalInformation  string
	GeodeticInformation      string
	GlobalGnbId              GlobalRanNodeId
}

var NrLocationData = NrLocation{
	Tai:                      TaiData,
	Ncgi:                     NcgiData,
	AgeOfLocationInformation: 5000,
	UeLocationTimestamp:      "2021-05-11T12:00:00",
	GeographicalInformation:  "GH7489KJD610LPO7",
	GeodeticInformation:      "GH7489KJD610LPO7",
	GlobalGnbId:              GlobalRanNodeIdData,
}
