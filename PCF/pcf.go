package main

// creates PCF server and register services

import (
	"github.com/tommcclymont/TM5GProject/logger"
	"log"
	"net/http"
	"os"

	"github.com/tommcclymont/TM5GProject/PCF/ampolicycontrol"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "pcf", log.LstdFlags)

	// create handlers for UDM services
	polh := ampolicycontrol.PolDataCreation(l)
	poldelh := ampolicycontrol.PolDataDeletion(l)

	// create serve mux for UDM
	sm := mux.NewRouter()

	// register the handlers
	postDataRouter := sm.Methods(http.MethodPost).Subrouter()
	postDataRouter.HandleFunc("/npcf-am-policy-control/v2/policies", polh.CreatePolData)

	deleteDataRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteDataRouter.HandleFunc("/npcf-am-policy-control/v2/policies", poldelh.DeletePolData)

	// handler for logging
	handler := logger.LogHandler(sm, "PCF/pcflog")

	// create PCF server
	s := &http.Server{
		Addr:    "127.0.0.79:9090",
		Handler: handler,
	}

	// start server
	err := s.ListenAndServeTLS("TLS/pcf.crt", "TLS/pcf.key")
	if err != nil {
		l.Fatal(err)
	}
}
