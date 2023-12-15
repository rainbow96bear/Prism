package Kakao

import (
	Login "prism_back/kakao/login"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/login", Login.OAuthLogin).Methods("GET")

    r.HandleFunc("/with_token", Login.OAuthLoginAfterProcess).Methods("GET")

    r.HandleFunc("/logout", Login.Logout).Methods("GET")
}

