package profile

import (
	"prism_back/pkg/models/profile/personaldata"

	"github.com/gorilla/mux"
)

var PersonalData = &personaldata.PersonalData{}

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/personaldata", PersonalData.GetPersonalData).Methods("GET")
}