package oauth

import (
	service "prism_back/service/oauth"

	"github.com/gorilla/mux"
)

type KakaoHandler struct {
	service.KakaoOAuth
}

// kakao OAuth handler
func (k *KakaoHandler) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/login", k.KakaoOAuth.Login).Methods("GET")
	r.HandleFunc("/login", k.KakaoOAuth.Logout).Methods("POST")
	// r.HandleFunc("/userino",).Methods("GET")
	// r.HandleFunc("/techs", a.Techs.GetTechListForAdmin).Methods("GET")
	// r.HandleFunc("/techs", a.Techs.AddTechForAdmin).Methods("POST")
}