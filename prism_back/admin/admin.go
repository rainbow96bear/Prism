package Admin

import (
	User "prism_back/admin/user"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
    
    r.HandleFunc("/user/check", User.AdminCheck).Methods("GET")
    r.HandleFunc("/user/logout", User.AdminLogout).Methods("GET")
    r.HandleFunc("/user/login", User.AdminLogin).Methods("POST")
}

