package adminaccess

import (
	"net/http"
	"prism_back/internal/session"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		session, err := session.Store.Get(req, "admin_login")
		Admin_id := session.Values["admin_login"]
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		if Admin_id != "" {
			next.ServeHTTP(res, req)
		} else {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
		}
	})
}