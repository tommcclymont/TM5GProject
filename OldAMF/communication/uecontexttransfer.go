package communication

// ue context transfer service

import (
	"github.com/tommcclymont/TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// uectxs is the struct for this handler
type uectxs struct {
	l *log.Logger
}

// UeContextTransfer creates the handler
func UeContextTransfer(l *log.Logger) *uectxs {
	return &uectxs{l}
}

func (uectx uectxs) TransferUeContext(rw http.ResponseWriter, r *http.Request) {

	var transferdata models.UeContextTransferReq
	err := json.NewDecoder(r.Body).Decode(&transferdata)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// set specified headers
	rw.Header().Set("Content-Type", "application/json")

	// extract ue context id from request path
	vars := mux.Vars(r)
	uecontextid := vars["uecontextid"]
	transferdata.UeContextId = uecontextid

	// create transfer request data
	err = models.PostUeContextReq(uecontextid, transferdata)
	if err != nil {
		http.Error(rw, "Ue Context Id not found", http.StatusInternalServerError)
		return
	}

	// get context data and convert to json
	uectxlist := models.GetUeContextData(uecontextid)
	d, err := json.Marshal(uectxlist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
