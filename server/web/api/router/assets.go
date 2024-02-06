package router

import (
	router "prism/web/api/router/assets"

	"github.com/gorilla/mux"
)

type Assets struct {
	router.Images
}
func (a *Assets)Router(r *mux.Router) {
	a.Images.Router(r.PathPrefix("/images").Subrouter())
}