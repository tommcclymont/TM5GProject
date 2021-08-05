package models

type UserLocation struct {
	EutraLocation EutraLocation
	NrLocation    NrLocation
	N3gaLocation  N3gaLocation
}

var UserlocationData = UserLocation{
	EutraLocation: EutraLocationData,
	NrLocation:    NrLocationData,
	N3gaLocation:  N3gaLocationData,
}
