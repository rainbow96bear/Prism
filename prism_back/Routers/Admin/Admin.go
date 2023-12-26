package AdminRouter

import (
	AdminAccessMiddleware "prism_back/Middleware/AdminAccess"
	"prism_back/Routers/Admin/User"
	Tech "prism_back/Structs/tech"

	"github.com/gorilla/mux"
)



func RegisterHandlers(r *mux.Router) {
	User.RegisterHandlers(r.PathPrefix("/user").Subrouter())
	adminRouter := r.PathPrefix("/access").Subrouter()
	adminRouter.Use(AdminAccessMiddleware.AdminMiddleware)
	AccessRequest(adminRouter)   
}

func AccessRequest(r *mux.Router) {
    r.HandleFunc("/tech", Tech.Get_tech_list).Methods("GET")
    r.HandleFunc("/tech", Tech.Post_tech_list).Methods("POST")
    r.HandleFunc("/tech", Tech.Put_tech_list).Methods("PUT")
}