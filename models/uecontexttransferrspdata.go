package models

type UeContextTransferRspData struct {
	UeContext         UeContext
	SupportedFeatures string
	UeRadioCapability N2InfoContent
}
