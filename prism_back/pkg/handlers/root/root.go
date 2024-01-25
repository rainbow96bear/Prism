package root

import (
	"prism_back/pkg/handlers/admin"
	"prism_back/pkg/handlers/oauths"
	"prism_back/pkg/handlers/profile"
	"prism_back/pkg/handlers/user"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	oauths.RegisterHandlers(r.PathPrefix("/OAuth").Subrouter())
	admin.RegisterHandlers(r.PathPrefix("/admin").Subrouter())
	user.RegisterHandlers(r.PathPrefix("/user").Subrouter())
	profile.RegisterHandlers(r.PathPrefix("/profile").Subrouter())
}
