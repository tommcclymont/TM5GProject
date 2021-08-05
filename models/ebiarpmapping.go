package models

type EbiArpmapping struct {
	EpsBearerId int
	Arp         Arp
}

var EbiArpmappingData = EbiArpmapping{
	EpsBearerId: 10,
	Arp:         ArpData,
}
