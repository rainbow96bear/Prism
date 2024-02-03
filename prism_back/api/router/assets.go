package router

import (
	router "prism_back/api/router/assets"

	"github.com/gorilla/mux"
)

type Assets struct {
	router.Images
}
func (a *Assets)Router(r *mux.Router) {
	a.Images.Router(r.PathPrefix("/images").Subrouter())
}