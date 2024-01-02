package admin

import (
	"prism_back/pkg/handlers/admin/auth"
	adminaccess "prism_back/pkg/middleware/admin_access"
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
    r.HandleFunc("/tech", Tech.GetTechList).Methods("GET")
    r.HandleFunc("/tech", Tech.PostTechList).Methods("POST")
    r.HandleFunc("/tech", Tech.PutTechList).Methods("PUT")
}