package UserInfo

import (
	"prism_back/userInfo/LightInfo"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/lightInfo", LightInfo.GetUserInfo_from_Session).Methods("GET")
}

