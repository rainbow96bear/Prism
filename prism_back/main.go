// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	KakaoLogin "prism_back/kakaoLogin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	port := 8080
	r := mux.NewRouter()

	var err error
	db, err := sql.Open("mysql", "root:0000@tcp(localhost:3306)/prism")
	if err != nil {
		fmt.Println("Failed to open DB")
	}
	defer db.Close()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:3306"}), // Change "*" to the actual front-end server's URL in production
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	r.Use(corsMiddleware)
	r.HandleFunc("/kakao/login", KakaoLogin.OAuthLogin).Methods("GET")
	r.HandleFunc("/kakao/withToken", KakaoLogin.OAuthLoginAfterProcess).Methods("GET")

	log.Println("Prism Server Starting on Port :", port)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
