package ueauthentication

// ue authentication service

import (
	"github.com/tommcclymont/TM5GProject/AUSF/ausfrequests"
	"github.com/tommcclymont/TM5GProject/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// ueauths is the struct for this handler
type ueauths struct {
	l *log.Logger
}

// UeAuthenticate creates the handler
func UeAuthenticate(l *log.Logger) *ueauths {
	return &ueauths{l}
}

func (ueauth ueauths) AuthenticateUe(rw http.ResponseWriter, r *http.Request) {

	var authenticatedata models.AuthenticationInfo
	err := json.NewDecoder(r.Body).Decode(&authenticatedata)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// set specified headers
	rw.Header().Set("Location", "https://127.0.0.80:9090/nausf-auth/v2/ue-authentications/dummyauthctxid")
	rw.Header().Set("Content-Type", "application/json")

	authctxid := "dummyauthctxid"
	authenticatedata.AuthCtxId = authctxid

	// create authentication request data
	err = models.PostUeAuth(authctxid, authenticatedata)
	if err != nil {
		http.Error(rw, "Authentication Context Id not found", http.StatusInternalServerError)
		return
	}

	// ausf to udm
	const (
		UEAUURL = "https://127.0.0.77:9090/nudm-ueau/v2"
	)

	var AuthenticationInfoRequestData = models.AuthenticationInfoRequest{
		ServingNetworkName:    authenticatedata.ServingNetworkName,
		ResynchronizationInfo: authenticatedata.ResynchronizationInfo,
		SupportedFeatures:     "D6e683bAfB2bbdCa",
		AusfInstanceId:        "da3a2883-7f28-45ff-83bd-820d0f48a65f",
	}

	supi := authenticatedata.SupiOrSuci

	AuthenticationInfoData, err := ausfrequests.NewClient(UEAUURL).AuthenticateUe(context.Background(), supi, AuthenticationInfoRequestData)
	if err != nil {
		panic(err)
	}

	var Av5gAkaData = models.Av5gAka{
		Rand:      AuthenticationInfoData.AuthenticationVector.Av5GHeAka.Rand,
		Autn:      models.AuthenticationVectorData.Av5GHeAka.AvType,
		HxresStar: models.AuthenticationVectorData.Av5GHeAka.XresStar,
	}

	var FivegAuthData = models.FivegAuthData{
		Av5gAka: Av5gAkaData,
	}

	var LinksData = map[string]models.LinksValueSchema{
		"5g-aka": {Link: UEAUURL},
	}

	var authctxlist = models.UeAuthenticationCtx{
		AuthType:           AuthenticationInfoData.AuthType,
		Links:              LinksData,
		FivegAuthData:      FivegAuthData,
		ServingNetworkName: "5G:AUSF",
	}

	d, err := json.Marshal(authctxlist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
