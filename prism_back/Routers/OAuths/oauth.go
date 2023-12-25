package OAuth

import (
	Kakao "prism_back/Router/OAuth/kakao"
	"prism_back/Structs/Users/KakaoUser"

	"github.com/gorilla/mux"
)

var kakaoUser = &KakaoUser.KakaoUser{}

func RegisterHandlers(r *mux.Router) {
    
	Kakao.RegisterHandlers(r.PathPrefix("/kakao").Subrouter())
}

