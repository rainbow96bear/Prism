package kakao

import (
	"prism_back/pkg/models/kakao/user"

	"github.com/gorilla/mux"
)

var kakaoUser = &user.KakaoUser{}

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/login", kakaoUser.Login).Methods("GET")

    r.HandleFunc("/with_token", kakaoUser.AfterProcess).Methods("GET")

    r.HandleFunc("/logout", kakaoUser.Logout).Methods("GET")
}