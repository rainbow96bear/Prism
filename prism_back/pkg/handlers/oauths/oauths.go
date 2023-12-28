package oauths

import (
	"prism_back/pkg/handlers/oauths/kakao"

	"github.com/gorilla/mux"
)


func RegisterHandlers(r *mux.Router) {
    
	kakao.RegisterHandlers(r.PathPrefix("/kakao").Subrouter())
}

