package user

import (
	"prism_back/pkg/models/user/info"

	"github.com/gorilla/mux"
)

var Info = &info.Info{}
func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/info", Info.GetUserInfo).Methods("GET")
}
