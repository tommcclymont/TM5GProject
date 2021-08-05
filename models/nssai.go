package models

// structure for nssai data as defined in specification 29.503 release 15
// provisioningTime absent
type Nssai struct {
	SupportedFeatures   string
	DefaultSingleNssais Snssai
	SingleNssais        Snssai
}

// create nssai dummy data
var NssaiData = Nssai{
	SupportedFeatures:   "001",
	DefaultSingleNssais: SnssaiData,
	SingleNssais:        SnssaiData,
}
