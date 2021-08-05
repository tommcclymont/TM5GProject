package models

type UeContext struct {
	Supi                   string
	SupiUnauthInd          bool
	GpsiList               []string
	Pei                    string
	UdmGroupId             string
	AusfGroupId            string
	RoutingIndicator       string
	GroupList              []string
	DrxParameter           []byte
	SubRfsp                uint
	UsedRfsp               uint
	SubUeAmbr              Ambr
	SmsSupport             string
	SmsfId                 string
	SeafData               SeafData
	FivegMmCapability      []byte
	PcfId                  string
	PcfAmPolicyUri         string
	AmPolicyReqTriggerList []string
	PcfUePolicyUri         string
	UePolicyReqTriggerList []string
	RestrictedRatList      []RatType
	ForbiddenAreaList      []Area
	ServiceAreaRestriction ServiceAreaRestriction
	RestrictedCnList       []CoreNetworkType
	EventSubscriptionList  []AmfEventSubscription
	SessionContextList     []PduSessionContext
	TraceData              TraceData
}
