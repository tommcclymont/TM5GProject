package main

// creates UDM server and register services

import (
	"TM5GProject/UDM/subscriberdatamanagement"
	"TM5GProject/UDM/ueauthenticationudm"
	"TM5GProject/UDM/uecontextmanagement"
	"TM5GProject/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "udm", log.LstdFlags)

	// create handlers for UDM services
	amsh := subscriberdatamanagement.AmsDataRetrieval(l)
	smfsubh := subscriberdatamanagement.SmfSubDataRetrieval(l)
	uectxsmfh := subscriberdatamanagement.UeCtxSmfDataRetrieval(l)
	amfregh := uecontextmanagement.AmsRegDataUpdate(l)
	genauthh := ueauthenticationudm.AuthGeneration(l)

	// create serve mux for UDM
	sm := mux.NewRouter()

	// register the handlers
	getDataRouter := sm.Methods(http.MethodGet).Subrouter()
	getDataRouter.HandleFunc("/nudm-sdm/v2/{supi}/am-data", amsh.GetAMSData)
	getDataRouter.HandleFunc("/nudm-sdm/v2/{supi}/smf-select-data", smfsubh.GetSMFSubData)
	getDataRouter.HandleFunc("/nudm-sdm/v2/{supi}/ue-context-in-smf-data", uectxsmfh.GetUeCtxSmfData)

	putDataRouter := sm.Methods(http.MethodPut).Subrouter()
	putDataRouter.HandleFunc("/nudm-uecm/v2/{ueid}/registrations/amf-3gpp-access", amfregh.UpdateAmsRegData)

	postDataRouter := sm.Methods(http.MethodPost).Subrouter()
	postDataRouter.HandleFunc("/nudm-ueau/v2/{supi}/security-information/generate-auth-data", genauthh.GenerateAuth)

	// handler for logging
	handler := logger.LogHandler(sm, "UDM/udmlog")

	// create UDM server
	s := &http.Server{
		Addr:    "127.0.0.77:9090",
		Handler: handler,
	}

	// start server
	err := s.ListenAndServeTLS("TLS/udm.crt", "TLS/udm.key")
	if err != nil {
		l.Fatal(err)
	}
}
