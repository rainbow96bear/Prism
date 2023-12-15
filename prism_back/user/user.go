package User

import (
	LightInfo "prism_back/user/lightInfo"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/light_info", LightInfo.GetUserInfo_from_Session).Methods("GET")
}

