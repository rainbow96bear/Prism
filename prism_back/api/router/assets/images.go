package router

import (
	handler "prism_back/api/handler/images/profileimgs"

	"github.com/gorilla/mux"
)

type Images struct {
	handler.Profileimgs
}

func (i *Images)Router(r *mux.Router) {
	i.Profileimgs.RegisterHandlers(r.PathPrefix("/profiles").Subrouter())
}