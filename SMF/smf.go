package main

// creates SMF server and register services

import (
	"TM5GProject/SMF/pdusession"
	"TM5GProject/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "smf", log.LstdFlags)

	// create handlers for SMF services
	smctxh := pdusession.SmContextCreate(l)
	smctxrelh := pdusession.SmContextRelease(l)

	// create serve mux for SMF
	sm := mux.NewRouter()

	// register the handlers
	postDataRouter := sm.Methods(http.MethodPost).Subrouter()
	postDataRouter.HandleFunc("/nsmf-pdusession/v2/sm-contexts", smctxh.CreateSmContext)
	postDataRouter.HandleFunc("/nsmf-pdusession/v2/sm-contexts/{smcontextref}/release", smctxrelh.ReleaseSmContext)

	// handler for logging
	handler := logger.LogHandler(sm, "SMF/smflog")

	// create SMF server
	s := &http.Server{
		Addr:    "127.0.0.82:9090",
		Handler: handler,
	}

	// start server
	err := s.ListenAndServeTLS("TLS/smf.crt", "TLS/smf.key")
	if err != nil {
		l.Fatal(err)
	}
}
