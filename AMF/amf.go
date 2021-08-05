package main

// creates AMF server and register services

import (
	"TM5GProject/AMF/eventexposure"
	"TM5GProject/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "amf", log.LstdFlags)

	// create handlers for AMF services
	evexpsubh := eventexposure.EventExposeSubscription(l)

	// create serve mux for AMF
	sm := mux.NewRouter()

	// register the handlers
	postDataRouter := sm.Methods(http.MethodPost).Subrouter()
	postDataRouter.HandleFunc("/namf-evts/v2/subscriptions", evexpsubh.SubscribeEventExpose)

	// handler for logging
	handler := logger.LogHandler(sm, "AMF/amflog")

	// create AMF server
	s := &http.Server{
		Addr:    "127.0.0.78:9090",
		Handler: handler,
	}

	// start server
	err := s.ListenAndServeTLS("TLS/amf.crt", "TLS/amf.key")
	if err != nil {
		l.Fatal(err)
	}
}
