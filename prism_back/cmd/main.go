// main.go
package main

import (
	"log"
	"net/http"
	"os"

	mysql "prism_back/internal/Database/mysql"
	"prism_back/internal/session"
	"prism_back/pkg/handlers/admin"
	"prism_back/pkg/handlers/oauths"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)



func main() {
	port := 8080
	r := mux.NewRouter()

	mysql.SetupDB()
	session.SetupStore()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{os.Getenv("FRONT_DOMAIN")}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)	
	
	r.Use(corsMiddleware)
	oauths.RegisterHandlers(r.PathPrefix("/OAuth").Subrouter())
	admin.RegisterHandlers(r.PathPrefix("/admin").Subrouter())

	log.Println("Prism Server Starting on Port :", port)
	// 라우터에 CORS 미들웨어 추가
	http.Handle("/", corsMiddleware(r))

	// 서버 시작
	http.ListenAndServe(":8080", nil)
}

