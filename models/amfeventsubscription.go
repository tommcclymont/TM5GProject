package models

type AmfEventSubscription struct {
	EventList                     []AmfEvent
	NotifyUri                     string
	NotifyCorrelationId           string
	NfId                          string
	SubsChangeNotifyUri           string
	SubsChangeNotifyCorrelationId string
	Supi                          string
	GroupId                       string
	Gpsi                          string
	Pei                           string
	AnyUE                         bool
}

var AmfEventSubscriptionData = AmfEventSubscription{
	EventList:           []AmfEvent{AmfEventData},
	NotifyUri:           "https://127.0.0.79:9090/",
	NotifyCorrelationId: "testamf",
	NfId:                "cf785448-6090-47bf-b28b-7371e2a737c6",
}
