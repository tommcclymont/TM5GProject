package models

type N3gaLocation struct {
	N3gppTai   Tai
	N3IwfId    string
	UeIpv4Addr string
	UeIpv6Addr string
	PortNumber uint
}

var N3gaLocationData = N3gaLocation{
	N3gppTai:   TaiData,
	N3IwfId:    "4FG67H",
	UeIpv4Addr: "1.1.1.1",
	UeIpv6Addr: "1111:a1a1:1111:a1a1:1111:a1a1:1111:a1a1",
	PortNumber: 9090,
}
