package admin

import (
	"prism_back/pkg/handlers/admin/auth"
	"prism_back/pkg/handlers/admin/tech"
	adminaccess "prism_back/pkg/middleware/admin_access"

	"github.com/gorilla/mux"
)



func RegisterHandlers(r *mux.Router) {
	auth.RegisterHandlers(r.PathPrefix("/user").Subrouter())
	
	adminRouter := r.PathPrefix("/access").Subrouter()
	adminRouter.Use(adminaccess.AdminMiddleware)
	tech.AccessRequest(adminRouter)   
}

