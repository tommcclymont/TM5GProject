package models

// structure for ip address defined in specification 29.503 release 15
type IpAddress struct {
	Ipv4Addr   string
	Ipv6Addr   string
	Ipv6Prefix string
}

var IpAddressData = IpAddress{
	Ipv4Addr:   "1.1.1.1",
	Ipv6Addr:   "1111:a1a1:1111:a1a1:1111:a1a1:1111:a1a1",
	Ipv6Prefix: "1111:a1a1:1111:a1a1",
}
