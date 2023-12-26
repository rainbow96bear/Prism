package User

import (
	"prism_back/Structs/Users/Admin"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	admin := &Admin.Admin{}
    r.HandleFunc("/check", admin.CheckRightAdmin).Methods("GET")
    r.HandleFunc("/logout", admin.Logout).Methods("GET")
    r.HandleFunc("/login", admin.Login).Methods("POST")
}
