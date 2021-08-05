package models

type TraceData struct {
	TraceRef                 string
	TraceDepth               string
	NeTypeList               string
	EventList                string
	CollectionEntityIpv4Addr string
	CollectionEntityIpv6Addr string
	InterfaceList            string
}

var TraceDataList = TraceData{
	TraceRef:                 "952112-eA6f9B",
	TraceDepth:               "MEDIUM",
	NeTypeList:               "E6F8e0ce85fCE05b15da9DeeE8F8B338082",
	EventList:                "fbC56a9BF7CA",
	CollectionEntityIpv4Addr: IpAddressData.Ipv4Addr,
	CollectionEntityIpv6Addr: IpAddressData.Ipv6Addr,
	InterfaceList:            "2D5efCca97b3d5faA3fb18105EdfaC9dA1af39CEB37c7fEb4ee93e8123509c4bf2f79D",
}
