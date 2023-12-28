package auth

import (
	"prism_back/pkg/models/user/admin"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	admin := &admin.Admin{}
    r.HandleFunc("/check", admin.CheckRightAdmin).Methods("GET")
    r.HandleFunc("/logout", admin.Logout).Methods("GET")
    r.HandleFunc("/login", admin.Login).Methods("POST")
}
