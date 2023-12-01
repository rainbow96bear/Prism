package main

import (
	"fmt"
	"log"
	"net/http"

	KakaoLogin "prism_back/kakaoLogin"

	"github.com/gorilla/mux"
)

func main() {
	port := 8080
	r := mux.NewRouter()

	r.HandleFunc("/kakaoLogin", KakaoLogin.KakaoLogin).Methods("GET")
	log.Println("Prism Server Starting on Port :", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), r))
}

