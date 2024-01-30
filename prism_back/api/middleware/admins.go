package middleware

import (
	"net/http"
	"os"
	"prism_back/pkg"
)

var (
	adminSession string = os.Getenv("ADMIN_SESSION")
)

type AdminsMiddleware struct {
	pkg.Session
}

func (a *AdminsMiddleware)CheckAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		admin_id, err := a.Session.GetID(adminSession, req)
		
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		if admin_id != "" {
			next.ServeHTTP(res, req)
		} else {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
		}
	})
}