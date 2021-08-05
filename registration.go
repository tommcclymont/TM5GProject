package main

// registration procedure

import (
	"github.com/tommcclymont/TM5GProject/AMF/requests"
	db "github.com/tommcclymont/TM5GProject/DB"

	"github.com/tommcclymont/TM5GProject/OldAMF/oldamfrequests"
	"github.com/tommcclymont/TM5GProject/PCF/pcfrequests"
	"github.com/tommcclymont/TM5GProject/models"
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// generate ue context for registration
	UeContext := GenerateUeContext()

	// ue context transfer step
	const (
		COMMURL = "https://127.0.0.76:9090/namf-comm/v2"
	)

	// prepare data for request
	UeContextTransferReqList := models.UeContextTransferReq{
		UeContextId:              UeContext.UeContextId,
		UeContextTransferReqData: models.UeContextTransferReqDataList,
		BinaryDataN1Message:      nil,
	}

	fmt.Printf("UE Context Transfer Request data: \n %+v\n\n", UeContextTransferReqList)

	// (AMF -> OldAMF) get ue context data
	UeCtxData, err := requests.NewClient(COMMURL).TransferUeContext(context.Background(), UeContext.UeContextId, UeContextTransferReqList)
	if err != nil {
		panic(err)
	}

	fmt.Printf("UE Context data: \n %+v\n\n", UeCtxData)

	// ue authentication step
	const (
		AUTHURL = "https://127.0.0.80:9090/nausf-auth/v2"
	)

	// prepare info to authenticate ue
	AuthenticationInfoData := models.AuthenticationInfo{
		SupiOrSuci:            UeCtxData.UeContextTransferRspData.UeContext.Supi,
		ServingNetworkName:    "5G:AUSF",
		ResynchronizationInfo: models.ResynchronizationInfoData,
		Pei:                   UeCtxData.UeContextTransferRspData.UeContext.Pei,
		TraceData:             UeCtxData.UeContextTransferRspData.UeContext.TraceData,
	}

	fmt.Printf("Authentication Info data: \n %+v\n\n", AuthenticationInfoData)

	// (AMF -> AUSF) authenticate ue
	ueAuthenticationData, err := requests.NewClient(AUTHURL).AuthenticateUe(context.Background(), AuthenticationInfoData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("UE Authentication data: \n %+v\n\n", ueAuthenticationData)

	// equipment status check step
	const (
		EICURL = "https://127.0.0.81:9090/n5g-eir-eic/v2"
	)

	// initialize optional parameters
	ueEIROptParams := requests.EIROptParams{
		Pei: UeCtxData.UeContextTransferRspData.UeContext.Pei,
	}

	// (AMF -> EIR) get equipment status
	eirData, err := requests.NewClient(EICURL).GetEquipmentStatus(context.Background(), &ueEIROptParams)
	if err != nil {
		panic(err)
	}

	if eirData.Status == "WHITELISTED" {
		fmt.Printf("UE is whitelisted \n\n")
	} else {
		fmt.Printf("UE is blacklisted; registration stopped")
		return
	}

	// amf access registration step
	const (
		UECMURL = "https://127.0.0.77:9090/nudm-uecm/v2"
	)

	// prepare amf access registration data
	amfRegPutData := models.Amf3gppAccessRegistration{
		AmfInstanceId:               "cf785448-6090-47bf-b28b-7371e2a737c6",
		DeregCallbackUri:            "https://127.0.0.78:9090/",
		Guami:                       models.GuamiData,
		RatType:                     models.RatTypeData,
		SupportedFeatures:           UeCtxData.UeContextTransferRspData.SupportedFeatures,
		PurgeFlag:                   false,
		Pei:                         UeCtxData.UeContextTransferRspData.UeContext.Pei,
		ImsVoPs:                     "HOMOGENEOUS_SUPPORT",
		AmfServiceNameDereg:         "testamf",
		PcscfRestorationCallbackUri: "https://127.0.0.78:9090/",
		AmfServiceNamePcscfRest:     "testamf",
		InitialRegistrationInd:      true,
		BackupAmfInfo:               models.BackupAmfInfoData,
		DrFlag:                      true,
		EpsInterworkingInfo:         models.EpsIwkPgwMap,
	}

	fmt.Printf("AMF Registration data: \n %+v\n\n", amfRegPutData)

	// (AMF -> UDM) update amf access registration data
	err = requests.NewClient(UECMURL).UpdateAmfRegData(context.Background(), UeCtxData.UeContextTransferRspData.UeContext.Supi, amfRegPutData)
	if err != nil {
		panic(err)
	}

	// access and mobility subscription data retrieval step
	const (
		SDMDataURL = "https://127.0.0.77:9090/nudm-sdm/v2"
	)

	// initialize optional parameters
	ueAMOptParams := requests.AMOptParams{}

	// (AMF -> UDM) get access and mobility subscription data
	accessAndMobilityData, err := requests.NewClient(SDMDataURL).GetAMData(context.Background(), UeCtxData.UeContextTransferRspData.UeContext.Supi, &ueAMOptParams)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Access and Mobility Subscription data: \n %+v\n\n", accessAndMobilityData)

	// old amf release sm context step
	const (
		PDUSessionURL = "https://127.0.0.82:9090/nsmf-pdusession/v2"
	)

	// prepare release data
	SmContextReleaseList := models.SmContextRelease{
		SmContextReleaseData:      models.SmContextReleaseDataList,
		BinaryDataN2SmInformation: nil,
	}

	fmt.Printf("SM Context Release data: \n %+v\n\n", SmContextReleaseList)

	// (OldAMF -> SMF) release SM context
	err = oldamfrequests.NewClient(PDUSessionURL).ReleaseSmContext(context.Background(), "dummysmcontextref", SmContextReleaseList)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Old AMF SM Context released")

	// smf selection subscription data retrieval step
	// initialize optional parameters
	ueSMFOptParams := requests.SMFOptParams{}

	// (AMF -> SMF) get smf selection subscription data
	smfSelectionSubscriptionData, err := requests.NewClient(SDMDataURL).GetSMFSelSubData(context.Background(), UeCtxData.UeContextTransferRspData.UeContext.Supi, &ueSMFOptParams)
	if err != nil {
		panic(err)
	}

	fmt.Printf("SMF Selection Subscription data: \n %+v\n\n", smfSelectionSubscriptionData)

	// ue context in smf data retrieval step
	// initialize optional parameters
	uectxSMFOptParams := requests.UeCtxSmfOptParams{}

	// (AMF -> SMF) get ue context in smf data
	ueContextInSmfData, err := requests.NewClient(SDMDataURL).GetUeSmfData(context.Background(), UeCtxData.UeContextTransferRspData.UeContext.Supi, &uectxSMFOptParams)
	if err != nil {
		panic(err)
	}

	fmt.Printf("UE Context in SMF data: \n %+v\n\n", ueContextInSmfData)

	// policy association data creation step
	const (
		NPCFURL = "https://127.0.0.79:9090/npcf-am-policy-control/v2"
	)

	// prepare request data
	PolicyAssociationRequestData := models.PolicyAssociationRequest{
		NotificationUri:  "https://127.0.0.78:9090/",
		AltNotifIpv4Addr: "1.1.1.1",
		AltNotifIpv6Addr: "1111:a1a1:1111:a1a1:1111:a1a1:1111:a1a1",
		Supi:             UeCtxData.UeContextTransferRspData.UeContext.Supi,
		Gpsi:             UeCtxData.UeContextTransferRspData.UeContext.GpsiList[0],
		AccessType:       "3GPP_ACCESS",
		Pei:              UeCtxData.UeContextTransferRspData.UeContext.Pei,
		UserLoc:          models.UserlocationData,
		TimeZone:         "00:00+0",
		ServingPlmn:      models.PlmnIdData,
		RatType:          models.RatTypeData,
		GroupIds:         []string{"12345678-111-22-3", "12345678-444-55-6"},
		ServAreaRes:      UeCtxData.UeContextTransferRspData.UeContext.ServiceAreaRestriction,
		Rfsp:             1,
		Guami:            models.GuamiData,
		ServiceName:      "testamf",
		SuppFeat:         UeCtxData.UeContextTransferRspData.SupportedFeatures,
	}

	fmt.Printf("Policy Association Request data: \n %+v\n\n", PolicyAssociationRequestData)

	// (AMF -> PCF) get policy association data
	policyAssociationData, err := requests.NewClient(NPCFURL).CreatePolAssoData(context.Background(), PolicyAssociationRequestData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Policy Association data: \n %+v\n\n", policyAssociationData)

	// event expose subscription step
	const (
		EVTSURL = "https://127.0.0.78:9090/namf-evts/v2"
	)

	AmfCreateEventSubscriptionData := models.AmfCreateEventSubscription{
		Subscription:      models.AmfEventSubscriptionData,
		SupportedFeatures: UeCtxData.UeContextTransferRspData.SupportedFeatures,
	}
	AmfCreateEventSubscriptionData.Subscription.Supi = UeCtxData.UeContextTransferRspData.UeContext.Supi

	fmt.Printf("AMF Event Subscription data: \n %+v\n\n", AmfCreateEventSubscriptionData)

	// (PCF -> AMF) create event exposure subscription
	eventExposureCreatedData, err := pcfrequests.NewClient(EVTSURL).CreateEventExposeSub(context.Background(), AmfCreateEventSubscriptionData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Event Exposure Created data: \n %+v\n\n", eventExposureCreatedData)

	// old amf policy deletion step
	// (OldAMF -> PCF) delete policy association data
	err = oldamfrequests.NewClient(NPCFURL).DeletePolAsso(context.Background(), "dummypolassoid")
	if err != nil {
		panic(err)
	}

	// sm context step
	// prepare sm context data
	SmContextCreateList := models.SmContextCreate{
		SmContextCreateData:          models.SmContextCreateDataList,
		BinaryDataN1SmMessage:        nil,
		BinaryDataN2SmInformation:    nil,
		BinaryDataN2SmInformationEx1: nil,
	}
	SmContextCreateList.SmContextCreateData.Supi = UeCtxData.UeContextTransferRspData.UeContext.Supi

	fmt.Printf("SM Context data: \n %+v\n\n", SmContextCreateList)

	// (AMF -> SMF) create sm context
	smContextCreatedData, err := requests.NewClient(PDUSessionURL).CreateSmContext(context.Background(), SmContextCreateList)
	if err != nil {
		panic(err)
	}

	fmt.Printf("SM Context Created data: \n %+v\n\n", smContextCreatedData)

	// registration complete
	fmt.Printf("Registration Completed!")
}

// function to generate ue context data to use for registration
func GenerateUeContext() models.UeContextTransferRsp {

	UeContext := models.UeContext{
		Supi:             RandomString("Numbers", 15),
		SupiUnauthInd:    false,
		GpsiList:         []string{"msisdn-" + RandomString("Numbers", 15)},
		Pei:              "imei-" + RandomString("Numbers", 15),
		UdmGroupId:       "dummyudm",
		AusfGroupId:      "dummyausf",
		RoutingIndicator: "1",
		DrxParameter:     nil,
		SubRfsp:          1,
		UsedRfsp:         1,
		SubUeAmbr: models.Ambr{Uplink: RandomString("Numbers", 2) + " Mbps",
			Downlink: RandomString("Numbers", 2) + " Mbps"},
		SmsSupport: "3GPP",
		SmsfId:     "dummysmf",
		SeafData: models.SeafData{NgKsi: models.NgKsiData,
			KeyAmf:               models.KeyAmfData,
			Nh:                   RandomString("Mixed", 32),
			Ncc:                  1,
			KeyAmfChangeInd:      false,
			KeyAmfHderivationInd: true},
		FivegMmCapability:      nil,
		PcfId:                  "dummypcf",
		PcfAmPolicyUri:         "https://127.0.0.79:9090",
		AmPolicyReqTriggerList: []string{"LOCATION_CHANGE"},
		PcfUePolicyUri:         "https://127.0.0.79:9090",
		UePolicyReqTriggerList: []string{"LOCATION_CHANGE"},
		RestrictedRatList:      []models.RatType{models.RatTypeData},
		ForbiddenAreaList:      []models.Area{{Areacodes: RandomString("Numbers", 5)}, {Areacodes: RandomString("Numbers", 5)}},
		ServiceAreaRestriction: models.ServiceAreaRestriction{Restrictiontype: models.RestrictionTypeData,
			Areas: []models.Area{{Areacodes: RandomString("Numbers", 5)}}},
		RestrictedCnList:      []models.CoreNetworkType{models.CoreNetworkTypeData},
		EventSubscriptionList: []models.AmfEventSubscription{models.AmfEventSubscriptionData},
		SessionContextList: []models.PduSessionContext{{PduSessionId: uint(rand.Intn(99)),
			SmContextRef:         "https://127.0.0.78:9090/",
			SNssai:               models.Snssai{Sst: uint(rand.Intn(99)), Sd: RandomString("Mixed", 6)},
			Dnn:                  RandomString("Lowercase", 11),
			AccessType:           "3GPP_ACCESS",
			AllocatedEbiList:     []models.EbiArpmapping{models.EbiArpmappingData},
			HsmfId:               "1a800ae6-706e-474f-89de-75747f19879c",
			NsInstance:           "73d753d7-9086-4515-a118-cb2b7f61e397",
			SmfServiceInstanceId: "1a800ae6-706e-474f-89de-75747f19879c"}},
		TraceData: models.TraceData{TraceRef: RandomString("Numbers", 6) + "-" + RandomString("Mixed", 6),
			TraceDepth:               "MEDIUM",
			NeTypeList:               RandomString("Mixed", 35),
			EventList:                RandomString("Mixed", 12),
			CollectionEntityIpv4Addr: RandomString("Numbers", 3) + "." + RandomString("Numbers", 3) + "." + RandomString("Numbers", 1) + "." + RandomString("Numbers", 1),
			CollectionEntityIpv6Addr: RandomString("Numbers", 4) + ":" + RandomString("Numbers", 4) + ":" + RandomString("Numbers", 4) + ":" + RandomString("Numbers", 4) + ":" + RandomString("Numbers", 4) + ":" + RandomString("Numbers", 4) + ":" + RandomString("Numbers", 4) + ":" + RandomString("Numbers", 4),
			InterfaceList:            RandomString("Mixed", 70)},
	}

	UeContextRspData := models.UeContextTransferRspData{
		UeContext:         UeContext,
		SupportedFeatures: RandomString("Mixed", 16),
		UeRadioCapability: models.N2InfoContentData,
	}

	UeContextRsp := models.UeContextTransferRsp{
		UeContextId:                 UeContext.Supi,
		UeContextTransferRspData:    UeContextRspData,
		BinaryDataN2Information:     nil,
		BinaryDataN2InformationExt1: nil,
	}

	// insert ue context data to database
	collection := db.NewDBClient().Database("Repository").Collection("uecontexttransferrsp")
	filter := bson.D{primitive.E{Key: "UeContextId", Value: UeContextRsp.UeContextId}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, UeContextRsp, opts)
	if err != nil {
		log.Fatal(err)
	}

	ams := models.ExAMSData
	ams.Supi = UeContext.Supi
	DbInsert("accessandmobilitysubscriptiondata", "Supi", ams.Supi, ams)

	amfreg := models.ExAmf3gppData
	amfreg.UeId = UeContext.Supi
	DbInsert("amf3gppaccessregistration", "Supi", amfreg.UeId, amfreg)

	evsub := models.ExEventSub
	DbInsert("amfcreatedeventsubscription", "SubscriptionId", evsub.SubscriptionId, evsub)

	eir := models.ExEirData
	eir.Pei = UeContext.Pei
	DbInsert("eirresponsedata", "Pei", eir.Pei, eir)

	polasso := models.ExPolAssoData
	DbInsert("policyassociation", "PolAssoId", polasso.PolAssoId, polasso)

	smctx := models.ExSmCtxCreateData
	DbInsert("smcontextcreate", "SmContextRef", smctx.SmContextRef, smctx)

	selsub := models.ExSelSubData
	selsub.Supi = UeContext.Supi
	DbInsert("smfselectionsubscriptiondata", "Supi", selsub.Supi, selsub)

	uesmfctx := models.ExUeSmfCtx
	uesmfctx.Supi = UeContext.Supi
	DbInsert("uecontextinsmfdata", "Supi", uesmfctx.Supi, uesmfctx)

	authinfo := models.ExAuthInfo
	authinfo.Supi = UeContext.Supi
	DbInsert("authenticationinforesult", "Supi", authinfo.Supi, authinfo)

	smctxd := models.ExSmCtx
	DbInsert("smcontextcreated", "SmContextRef", smctxd.SmContextRef, smctxd)

	return UeContextRsp
}

var nums = []rune("0123456789")
var lcase = []rune("abcdefghijklmnopqrstuvwxyz")
var mix = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVXYZ0123456789")

// function to create a random string of numbers of letters to be used in ue context generation
func RandomString(t string, n int) string {

	rand.Seed(time.Now().UnixNano())

	str := make([]rune, n)
	for i := range str {
		if t == "Numbers" {
			str[i] = nums[rand.Intn(len(nums))]
		} else if t == "Lowercase" {
			str[i] = lcase[rand.Intn(len(lcase))]
		} else if t == "Mixed" {
			str[i] = mix[rand.Intn(len(mix))]
		}
	}
	return string(str)
}

func DbInsert(coll string, key string, keyvalue string, d interface{}) {

	collection := db.NewDBClient().Database("Repository").Collection(coll)
	filter := bson.D{primitive.E{Key: key, Value: keyvalue}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, d, opts)
	if err != nil {
		panic(err)
	}
}
