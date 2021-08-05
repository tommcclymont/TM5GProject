package main

// creates AUSF server and register services

import (
	"github.com/tommcclymont/TM5GProject/AUSF/ueauthentication"
	"github.com/tommcclymont/TM5GProject/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "ausf", log.LstdFlags)

	// create handlers for AUSF services
	ueauthh := ueauthentication.UeAuthenticate(l)

	// create serve mux for AUSF
	sm := mux.NewRouter()

	// register the handlers
	postDataRouter := sm.Methods(http.MethodPost).Subrouter()
	postDataRouter.HandleFunc("/nausf-auth/v2/ue-authentications", ueauthh.AuthenticateUe)

	// handler for logging
	handler := logger.LogHandler(sm, "AUSF/ausflog")

	// create AUSF server
	s := &http.Server{
		Addr:    "127.0.0.80:9090",
		Handler: handler,
	}

	// start server
	err := s.ListenAndServeTLS("TLS/ausf.crt", "TLS/ausf.key")
	if err != nil {
		l.Fatal(err)
	}
}
