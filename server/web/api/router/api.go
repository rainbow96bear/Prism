package router

import (
	"prism_back/api/handler"
	"prism_back/api/handler/oauth"

	"github.com/gorilla/mux"
)

type API struct {
	handler.UsersHandler
	handler.AdminsHandler
	oauth.KakaoHandler
}

func (a *API)Router(r *mux.Router) {
	a.KakaoHandler.RegisterHandlers(r.PathPrefix("/oauth/kakao").Subrouter())
	
	a.UsersHandler.RegisterHandlers(r.PathPrefix("/users").Subrouter())
	
	a.AdminsHandler.RegisterHandlers(r.PathPrefix("/admins").Subrouter())
}