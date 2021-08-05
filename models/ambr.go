package models

// structure for ambr data as defined in specification 29.571 release 15
type Ambr struct {
	Uplink   string
	Downlink string
}

// example ambr data
var AmbrData = Ambr{
	Uplink:   "20 Mbps",
	Downlink: "20 Mbps",
}
