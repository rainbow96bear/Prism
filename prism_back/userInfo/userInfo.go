package UserInfo

import (
	LightInfo "prism_back/userInfo/lightInfo"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/lightInfo", LightInfo.GetUserInfo_from_Session).Methods("GET")
}

