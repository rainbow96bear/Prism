// main.go
package main

import (
	"log"
	"net/http"
	Mysql "prism_back/DataBase/MySQL"
	Admin "prism_back/admin"
	Kakao "prism_back/kakao"
	Session "prism_back/session"
	User "prism_back/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)



func main() {
	port := 8080
	r := mux.NewRouter()

	Mysql.SetupDB()
	Session.SetupStore()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)	
	
	r.Use(corsMiddleware)

	Kakao.RegisterHandlers(r.PathPrefix("/kakao").Subrouter())
	User.RegisterHandlers(r.PathPrefix("/userInfo").Subrouter())
	Admin.RegisterHandlers(r.PathPrefix("/admin").Subrouter())

	log.Println("Prism Server Starting on Port :", port)
	// 라우터에 CORS 미들웨어 추가
	http.Handle("/", corsMiddleware(r))

	// 서버 시작
	http.ListenAndServe(":8080", nil)
}

