package profile

import (
	"prism_back/pkg/models/images"
	"prism_back/pkg/models/profile/personaldata"

	"github.com/gorilla/mux"
)

var PersonalData = &personaldata.PersonalData{}
var Image = &images.Image{}

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/upload/image", Image.UploadImageHandler).Methods("POST")
	r.HandleFunc("/personaldata/{id}", PersonalData.GetPersonalData).Methods("GET")
	r.HandleFunc("/update/{id}", PersonalData.SetPersonalData).Methods("POST")
}