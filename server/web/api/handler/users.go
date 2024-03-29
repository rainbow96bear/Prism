package handler

import (
	"prism_back/api/middleware"
	"prism_back/service"

	"github.com/gorilla/mux"
)

type UsersHandler struct {
	service.Profile
	service.Techs
	service.User
	middleware.ProfilesMiddleware
}

// users handler
func (u *UsersHandler)RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/info", u.User.GetInfo).Methods("GET")
	r.HandleFunc("/profiles/{id}/personaldatas", u.Profile.GetUserProfile).Methods("GET")
	r.HandleFunc("/profiles/{id}/personaldatas", u.ProfilesMiddleware.CheckAccess(u.Profile.UpdateUserProfile)).Methods("PUT")
	r.HandleFunc("/profiles/techlist",u.Techs.GetTechList).Methods("GET")
	r.HandleFunc("/profiles/{id}/techs", u.Techs.GetUserTechList).Methods("GET")
	r.HandleFunc("/profiles/{id}/techs", u.ProfilesMiddleware.CheckAccess(u.Techs.UpdateUserTechList)).Methods("PUT")

	// 프로필에 기술 스택 추가 시 기술 스택 목록을 위한 기술스택 이름 배열
}

