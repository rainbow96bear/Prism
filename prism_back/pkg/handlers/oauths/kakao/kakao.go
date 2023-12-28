package kakao

import (
	"prism_back/pkg/models/user/kakaoUser"

	"github.com/gorilla/mux"
)

var user = &kakaoUser.KakaoUser{}

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/login", user.Login).Methods("GET")

    r.HandleFunc("/with_token", user.AfterProcess).Methods("GET")

    r.HandleFunc("/logout", user.Logout).Methods("GET")
}