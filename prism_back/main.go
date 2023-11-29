package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := 8080
	r := mux.NewRouter()

	log.Println("Prism Server Starting on Port :", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), r))
}