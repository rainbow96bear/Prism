package profile

import (
	profileaccess "prism_back/pkg/middleware/profile_access"
	"prism_back/pkg/models/images"
	"prism_back/pkg/models/profile/personaldata"
	"prism_back/pkg/models/profile/techdata"
	"prism_back/pkg/models/tech"

	"github.com/gorilla/mux"
)

var PersonalData = &personaldata.PersonalData{}
var Tech = &techdata.Tech{}

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/upload/image", images.UploadImageHandler).Methods("POST")
	r.HandleFunc("/personaldata/{id}", PersonalData.GetPersonalData).Methods("GET")
	r.HandleFunc("/techs/{id}", Tech.GetTechList).Methods("GET")
	r.HandleFunc("/tech/name_list", tech.GetTechsName).Methods("GET")
	
	// profileRouter를 생성하고 미들웨어를 추가
	profileRouter := r.PathPrefix("/update").Subrouter()
	profileRouter.Use(profileaccess.ProfileMiddleware)
	
	profileRouter.HandleFunc("/{id}", PersonalData.SetPersonalData).Methods("POST")
	profileRouter.HandleFunc("/techs/{id}", Tech.SetTechList).Methods("POST")
}
