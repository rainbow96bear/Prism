package service

import (
	"encoding/json"
	"log"
	"net/http"
	"prism_back/dto"
	"prism_back/errors"
	"prism_back/internal/Database/mysql"
	"prism_back/internal/session"
	"prism_back/repository"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	repository.AdminUserRepository
}

// admin 계정 로그인
func (a *Admin) Login(res http.ResponseWriter, req *http.Request) {
	// userSession으로 현제 로그인 사용자 확인
	id, err := session.GetID(userSession, req)
	if err != nil {
		log.Println("service/admin.go : session 생성 실패", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("DB 시작 오류", err)
		tx.Rollback()
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	var adminInfo dto.AdminInfo
	adminUserInfo, err := a.AdminUserRepository.GetAdminUserInfo(tx, id)
	if err != nil {
		if err == errors.IsNotAdminUser {
			log.Println("관리자 계정이 아닌 사용자 접근", id)
		}else {
			log.Println("DB에서 정보 획득 실패", err)
			tx.Rollback()
		}
		adminInfo = dto.AdminInfo{Id : "", Rank : 0}
		jsonResponse, err := json.Marshal(adminInfo)
		if err != nil {
			log.Println("Admin 로그인 결과 Marshal 오류", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return 
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(jsonResponse)
		return 
	}
	
	err = tx.Commit()
	if err != nil {
		adminInfo = dto.AdminInfo{Id : "", Rank : 0}
		jsonResponse, err := json.Marshal(adminInfo)
		if err != nil {
			log.Println("tx Commit 오류", err)
			tx.Rollback()
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return 
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(jsonResponse)
		return 
	}

	reqPassword, err := getPasswordFromRequest(req)
	if err != nil {
		log.Println("요청에서 password 얻기 실패", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return 
	}

	if !comparePassword(adminUserInfo.Password, reqPassword) {
		adminInfo = dto.AdminInfo{Id : "", Rank : 0}
		jsonResponse, err := json.Marshal(adminInfo)
		if err != nil {
			log.Println("Admin 로그인 결과 Marshal 오류", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return 
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(jsonResponse)
		return 
	}
	
	adminInfo = dto.AdminInfo{Id : adminUserInfo.Id, Rank : adminUserInfo.Rank}
	jsonResponse, err := json.Marshal(adminInfo)
	if err != nil {
		log.Println("service/admin.go : Admin 로그인 결과 Marshal 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return 
	}

	err = session.CreateAdminSession(adminSession, id, adminInfo.Rank, res, req)
	if err != nil {
		log.Println("service/admin.go : Admin session 생성 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return 
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}

// admin 계정 로그아웃
func (a *Admin)Logout(res http.ResponseWriter, req *http.Request) {
	session.DeleteSession(adminSession, res, req)
}

// admin 계정 확인
func (a *Admin) CheckAuthorization(res http.ResponseWriter, req *http.Request){
	authorization := dto.AdminAuthorization{IsAdmin: false, AdminInfo: dto.AdminInfo{Id: "", Rank: 0}}
	adminID, err := session.GetID(adminSession, req)
	rank, err := session.GetRank(adminSession, req)
	if err != nil {
		log.Println("service/admin.go : session 확인 오류", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if adminID != "" {
		authorization.IsAdmin = true
		authorization.AdminInfo.Id=adminID
		authorization.AdminInfo.Rank = rank
	}else {
		id, err := session.GetID(userSession, req)
		if err != nil {
			log.Println("service/admin.go : session 확인 오류", err)
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tx, err := mysql.DB.Begin()
		if err != nil {
			log.Println("service/admin.go : DB 시작 오류", err)
			tx.Rollback()
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		
		_, getAdminUserInfoErr := a.AdminUserRepository.GetAdminUserInfo(tx, id)
		if getAdminUserInfoErr != nil && getAdminUserInfoErr != errors.IsNotAdminUser{
			log.Println("service/techs.go : admin 정보 얻기 오류", err)
			tx.Rollback()
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		}
		
		err = tx.Commit()
		if err != nil {
			log.Println("service/techs.go : commit 오류", err)
			tx.Rollback()
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		}

		if getAdminUserInfoErr != errors.IsNotAdminUser {
			authorization.IsAdmin = true
		}
	}
	jsonResponse, err := json.Marshal(authorization)
	if err != nil {
		log.Println("service/admin.go : Admin 로그인 결과 Marshal 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return 
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)

}
// 비밀번호 비교
func comparePassword(DBPassword, requestPassword string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(DBPassword), []byte(requestPassword))
	if err != nil {
		return false
	}
	return true
}

// 요청으로 부터 password 얻기
func getPasswordFromRequest(req *http.Request) (string, error){
	type Password struct{
		Password string `json:"password"`
	}
	var requestData Password
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		return requestData.Password, err
	}
	return requestData.Password, err
}

func InitRootAdmin() {
	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Fatal("DB 시작 오류", err)
		return
	}
	adminRepo := repository.AdminUserRepository{}
	adminInfo, err := adminRepo.GetAdminUserInfo(tx, rootAdminID)
	if err != nil && err != errors.IsNotAdminUser{
		log.Fatal("DB에서 admin 정보 확인 오류", err)
		return
	}
	
	if adminInfo.Id != "" {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rootAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Root admin 비밀번호 해싱 오류", err)
		return
	}

	err = adminRepo.CreateAdmin(tx, rootAdminID, string(hashedPassword), adminRank)
	if err != nil {
		log.Fatal("RootAdmin 계정 생성 실패", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Root 관리자 계정 생성 실패", err)
		return
	}
	log.Println("계정이 생성되었습니다. 비밀번호 : ", string(hashedPassword))
	return 
}