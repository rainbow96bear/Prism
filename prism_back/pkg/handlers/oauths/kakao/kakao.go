package kakao

import (
	"prism_back/pkg/models/kakao/user"

	"github.com/gorilla/mux"
)

var kakaoUser = &user.KakaoUser{}

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/code", kakaoUser.Login).Methods("GET")

    r.HandleFunc("/userinfo", kakaoUser.GetUserInfo).Methods("GET")

    r.HandleFunc("/logout", kakaoUser.Logout).Methods("POST")
}