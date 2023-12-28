package admin

import (
	"prism_back/pkg/handlers/admin/auth"
	"prism_back/pkg/middleware/adminaccess"
	Tech "prism_back/pkg/models/tech"

	"github.com/gorilla/mux"
)



func RegisterHandlers(r *mux.Router) {
	auth.RegisterHandlers(r.PathPrefix("/user").Subrouter())
	adminRouter := r.PathPrefix("/access").Subrouter()
	adminRouter.Use(adminaccess.AdminMiddleware)
	AccessRequest(adminRouter)   
}

func AccessRequest(r *mux.Router) {
    r.HandleFunc("/tech", Tech.Get_tech_list).Methods("GET")
    r.HandleFunc("/tech", Tech.Post_tech_list).Methods("POST")
    r.HandleFunc("/tech", Tech.Put_tech_list).Methods("PUT")
}