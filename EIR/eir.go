package main

// creates 5G-EIR server and register services

import (
	"github.com/tommcclymont/TM5GProject/EIR/equipmentidentitycheck"
	"github.com/tommcclymont/TM5GProject/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "eir", log.LstdFlags)

	// create handlers for EIR services
	eqcheckh := equipmentidentitycheck.EquipmentStatusRetrieval(l)

	// create serve mux for EIR
	sm := mux.NewRouter()

	// register the handlers
	getDataRouter := sm.Methods(http.MethodGet).Subrouter()
	getDataRouter.HandleFunc("/n5g-eir-eic/v2/equipment-status", eqcheckh.GetEquipmentStatus)

	// handler for logging
	handler := logger.LogHandler(sm, "EIR/eirlog")

	// create EIR server
	s := &http.Server{
		Addr:    "127.0.0.81:9090",
		Handler: handler,
	}

	// start server
	err := s.ListenAndServeTLS("TLS/eir.crt", "TLS/eir.key")
	if err != nil {
		l.Fatal(err)
	}
}
