package router

import (
	"prism/web/api/handler"
	"prism/web/api/handler/oauth"

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