package service

import (
	"encoding/json"
	"log"
	"net/http"
	"prism_back/dto"
	"prism_back/internal/session"
)

type User struct {
}

func (u *User) GetInfo(res http.ResponseWriter, req *http.Request) {
	id, err := session.GetID(userSession, req)
	if err != nil {
		log.Println("service/user.go : session에서 id 얻기 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	userinfo := dto.User{Id: id}
	responseJSON, err := json.Marshal(userinfo)
	if err != nil {
		return 
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(responseJSON)
}