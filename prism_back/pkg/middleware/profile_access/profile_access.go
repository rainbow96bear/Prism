package profileaccess

import (
	"net/http"

	"prism_back/internal/session"

	"github.com/gorilla/mux"
)

func ProfileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		session, err := session.Store.Get(req, "user_login")
		User_id := session.Values["User_ID"]

		vars := mux.Vars(req)
		id := vars["id"]
		
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		if User_id == id {
			next.ServeHTTP(res, req)
		} else {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
		}
	})
}