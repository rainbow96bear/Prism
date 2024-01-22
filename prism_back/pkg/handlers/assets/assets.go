package assets

import (
	"prism_back/pkg/handlers/assets/images"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	images.RegisterHandlers(r.PathPrefix("/images").Subrouter())
}