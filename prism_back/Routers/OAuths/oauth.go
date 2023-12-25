package OAuth

import (
	"prism_back/Routers/OAuths/Kakao"
	"prism_back/Structs/Users/KakaoUser"

	"github.com/gorilla/mux"
)

var kakaoUser = &KakaoUser.KakaoUser{}

func RegisterHandlers(r *mux.Router) {
    
	Kakao.RegisterHandlers(r.PathPrefix("/kakao").Subrouter())
}

