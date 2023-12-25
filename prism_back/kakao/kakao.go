package Kakao

import (
	"prism_back/Structs/Users/KakaoUser"

	"github.com/gorilla/mux"
)

var kakaoUser = &KakaoUser.KakaoUser{}

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/login", kakaoUser.Login).Methods("GET")

    r.HandleFunc("/with_token", kakaoUser.AfterProcessres).Methods("GET")

    r.HandleFunc("/logout", kakaoUser.Logout).Methods("GET")
}

