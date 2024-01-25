package images

import (
	"prism_back/pkg/handlers/assets/images/profiles"

	"github.com/gorilla/mux"
)


func RegisterHandlers(r *mux.Router) {
	profiles.RegisterHandlers(r.PathPrefix("/profiles").Subrouter())
}

