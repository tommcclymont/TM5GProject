package main

// creates old AMF server and register services

import (
	"TM5GProject/OldAMF/communication"
	"TM5GProject/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "oldamf", log.LstdFlags)

	// create handlers for old AMF services
	uectxh := communication.UeContextTransfer(l)

	// create serve mux for old AMF
	sm := mux.NewRouter()

	// register the handlers
	postDataRouter := sm.Methods(http.MethodPost).Subrouter()
	postDataRouter.HandleFunc("/namf-comm/v2/ue-contexts/{uecontextid}/transfer", uectxh.TransferUeContext)

	// handler for logging
	handler := logger.LogHandler(sm, "OldAMF/oldamflog")

	// create AMF server
	s := &http.Server{
		Addr:    "127.0.0.76:9090",
		Handler: handler,
	}

	// start server
	err := s.ListenAndServeTLS("TLS/oldamf.crt", "TLS/oldamf.key")
	if err != nil {
		l.Fatal(err)
	}
}
