package oauth

import (
	service "prism/web/service/oauth"

	"github.com/gorilla/mux"
)

type KakaoHandler struct {
	service.KakaoOAuth
}

// kakao OAuth handler
func (k *KakaoHandler) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/login", k.KakaoOAuth.Login).Methods("GET")
	r.HandleFunc("/logout", k.KakaoOAuth.Logout).Methods("POST")
}