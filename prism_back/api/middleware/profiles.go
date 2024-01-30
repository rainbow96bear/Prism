package middleware

import (
	"net/http"
	"os"

	"prism_back/pkg"

	"github.com/gorilla/mux"
)

var (
	userSession string = os.Getenv("USER_SESSION")
)

type ProfilesMiddleware struct {
	pkg.Session
}

func (p *ProfilesMiddleware)CheckAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		user_id, err := p.Session.GetID(userSession, req)

		vars := mux.Vars(req)
		id := vars["id"]
		
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		if user_id == id {
			next.ServeHTTP(res, req)
		} else {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
		}
	})
}