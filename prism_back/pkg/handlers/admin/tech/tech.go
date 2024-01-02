package tech

import (
	"prism_back/pkg/models/tech"

	"github.com/gorilla/mux"
)

func AccessRequest(r *mux.Router) {
	r.HandleFunc("/tech", tech.GetTechList).Methods("GET")
	r.HandleFunc("/tech", tech.PostTechList).Methods("POST")
	r.HandleFunc("/tech", tech.PutTechList).Methods("PUT")
}