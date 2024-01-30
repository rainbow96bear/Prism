package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"prism_back/dto"
	"prism_back/internal/Database/mysql"
	"prism_back/repository"

	"github.com/gorilla/mux"
)

type Techs struct {
	repository.ProfileHasTechListRepository
	repository.TechRepository
}


// 사용자 profile에 출력할 사용자의 tech 목록 얻기
func (t *Techs) GetUserTechList(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id := vars["id"]

	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	techList, err := t.ProfileHasTechListRepository.GetUserTechList(tx, id)
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}

	responseJSON, err := json.Marshal(techList)
	if err != nil {
		log.Println("service/techs.go : json marshal 오류", err)
		return
	}

	res.Header().Set("Content/type", "application/json")
	res.Write(responseJSON)
}

// Admin 페이지에서 관리할 Tech의 목록 얻기
func (t *Techs)GetTechList(res http.ResponseWriter, req *http.Request) {
	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	techList, err := t.TechRepository.ReadAll(tx)
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	dtoTechList := []dto.Tech{}
	for _, value := range techList {
		dtoTechList = append(dtoTechList, dto.Tech{Name: value.TechName, Count:value.Count})
	}

	responseJSON, err := json.Marshal(dtoTechList)
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(responseJSON)
}

// Admin이 Tech의 목록에 기술 스택 추가
func (t *Techs)AddTechForAdmin(res http.ResponseWriter, req *http.Request) {
	
	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	body, err := io.ReadAll(req.Body)
	var tech dto.Tech
	err = json.Unmarshal(body, &tech)
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	} 

	err = t.TechRepository.Create(tx, tech.Name)
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// User의 기술 스택 수정
func (t *Techs)UpdateUserTechList(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	body, err := io.ReadAll(req.Body)
	var userTechList []dto.UserTech
	err = json.Unmarshal(body, &userTechList)
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	} 
	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.ProfileHasTechListRepository.Delete(tx, id)
	if err != nil {
		log.Println("service/techs.go :", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	for _, value := range userTechList{
		err := t.ProfileHasTechListRepository.Create(tx, id, value.Name, value.Level)
		if err != nil {
			log.Println("service/techs.go :", err)
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	
}