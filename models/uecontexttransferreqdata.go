package models

type UeContextTransferReqData struct {
	Reason            string
	AccessType        string
	PlmnId            PlmnId
	RegRequest        N1MessageContainer
	SupportedFeatures string
}

var UeContextTransferReqDataList = UeContextTransferReqData{
	Reason:            "INIT_REG",
	AccessType:        "3GPP_ACCCESS",
	PlmnId:            PlmnIdData,
	RegRequest:        N1MessageContainerData,
	SupportedFeatures: "D6e683bAfB2bbdCa",
}
