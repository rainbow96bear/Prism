package handler

import (
	"prism_back/api/middleware"
	"prism_back/service"

	"github.com/gorilla/mux"
)

type AdminsHandler struct {
	service.Techs
	service.Admin
	middleware.AdminsMiddleware
}

// admin handler
func (a *AdminsHandler)RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/authorization", a.Admin.CheckAuthorization).Methods("GET")
	r.HandleFunc("/login", a.Admin.Login).Methods("POST")
	r.HandleFunc("/logout", a.Admin.Logout).Methods("POST")

	r.HandleFunc("/techs", a.Techs.GetTechList).Methods("GET")
	r.HandleFunc("/techs", a.AdminsMiddleware.CheckAccess(a.Techs.AddTechForAdmin)).Methods("POST")
	r.HandleFunc("/techs/{id}", a.AdminsMiddleware.CheckAccess(a.Techs.UpdateTech)).Methods("PUT")
}