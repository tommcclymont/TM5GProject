package ampolicycontrol

// policy association deletion service

import (
	"TM5GProject/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// poldels is the struct for this handler
type poldels struct {
	l *log.Logger
}

// PolDataDeletion creates the handler
func PolDataDeletion(l *log.Logger) *poldels {
	return &poldels{l}
}

func (poldel poldels) DeletePolData(rw http.ResponseWriter, r *http.Request) {

	// extract polassoid from request path
	vars := mux.Vars(r)
	polassoid := vars["polassoid"]

	// delete policy association data
	err := models.DeletePolAsso(polassoid)
	if err != nil {
		http.Error(rw, "Policy Association Id not found", http.StatusInternalServerError)
		return
	}

	// write no content header for response
	rw.WriteHeader(http.StatusNoContent)
}
