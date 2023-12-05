package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	KakaoLogin "prism_back/kakaoLogin"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
var DB *sql.DB

func main() {


	port := 8080
	r := mux.NewRouter()

	DB, err := sql.Open("mysql", "root:1q2w3e4r!@tcp(localhost:3333)/Books")
	if err != nil {
		fmt.Println("Failed to open DB")
	}
	defer DB.Close()

	corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Change "*" to the actual front-end server's URL in production
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

