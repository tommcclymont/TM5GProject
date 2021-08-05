package main

// clear all logs

import (
	"os"
)

func main() {

	if err := os.Truncate("AMF/amflog.json", 0); err != nil {
		panic(err)
	}
	if err := os.Truncate("AUSF/ausflog.json", 0); err != nil {
		panic(err)
	}
	if err := os.Truncate("EIR/eirlog.json", 0); err != nil {
		panic(err)
	}
	if err := os.Truncate("OldAMF/oldamflog.json", 0); err != nil {
		panic(err)
	}
	if err := os.Truncate("PCF/pcflog.json", 0); err != nil {
		panic(err)
	}
	if err := os.Truncate("SMF/smflog.json", 0); err != nil {
		panic(err)
	}
	if err := os.Truncate("UDM/udmlog.json", 0); err != nil {
		panic(err)
	}
}
