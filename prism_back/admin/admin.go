package Admin

import (
	Middleware "prism_back/admin/middleware"
	Tech "prism_back/admin/tech"
	User "prism_back/admin/user"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {

    r.HandleFunc("/user/check", User.AdminCheck).Methods("GET")
    r.HandleFunc("/user/logout", User.AdminLogout).Methods("GET")
    r.HandleFunc("/user/login", User.AdminLogin).Methods("POST")

    adminRouter := r.PathPrefix("/access").Subrouter()
	adminRouter.Use(Middleware.AdminMiddleware)
	AccessRequest(adminRouter)   
}

func AccessRequest(r *mux.Router) {
    r.HandleFunc("/tech", Tech.Get_tech_list).Methods("GET")
    r.HandleFunc("/tech", Tech.Post_tech_list).Methods("POST")
    r.HandleFunc("/tech", Tech.Put_tech_list).Methods("PUT")
}

