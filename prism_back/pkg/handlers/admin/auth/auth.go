package auth

import (
	"prism_back/pkg/models/admin"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	admin := &admin.Admin{}
    r.HandleFunc("/status", admin.CheckRightAdmin).Methods("GET")
    r.HandleFunc("/login", admin.Login).Methods("POST")
    r.HandleFunc("/logout", admin.Logout).Methods("POST")
}
