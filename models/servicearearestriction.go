package models

// structure for servicearearestriction data as defined in specification 29.571 release 15
type ServiceAreaRestriction struct {
	Restrictiontype RestrictionType
	Areas           []Area
	Maxnumoftas     int
}

var ServiceAreaRestrictionData = ServiceAreaRestriction{
	Restrictiontype: RestrictionTypeData,
	Areas:           []Area{AreaData2},
	Maxnumoftas:     1,
}
