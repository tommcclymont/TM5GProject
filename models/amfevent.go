package models

type AmfEvent struct {
	Type          string
	ImmediateFlag bool
	RefId         int
}

var AmfEventData = AmfEvent{
	Type:          "TIMEZONE_REPORT",
	ImmediateFlag: false,
	RefId:         111,
}
