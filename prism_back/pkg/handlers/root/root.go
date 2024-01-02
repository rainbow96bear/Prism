package root

import (
	"prism_back/pkg/handlers/admin"
	"prism_back/pkg/handlers/oauths"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	oauths.RegisterHandlers(r.PathPrefix("/OAuth").Subrouter())
	admin.RegisterHandlers(r.PathPrefix("/admin").Subrouter())
}
