package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"prism_back/dto"
	"prism_back/errors"
	"prism_back/internal/Database/mysql"
	"prism_back/pkg"
	"prism_back/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	adminSession string = os.Getenv("ADMIN_SESSION")
	userSession string = os.Getenv("USER_SESSION")
	rootAdminID string = os.Getenv("ROOT_ADMIN_ID")
	rootAdminPassword string = os.Getenv("ROOT_ADMIN_PASSWORD")
)
type Admin struct {
	pkg.Session
	repository.AdminUserRepository
}

// admin 계정 로그인
func (a *Admin) Login(res http.ResponseWriter, req *http.Request) {
	// userSession으로 현제 로그인 사용자 확인
	id, err := a.Session.GetID(userSession, req)
	if err != nil {
		log.Println("service/admin.go : session 생성 실패", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("DB 시작 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	var adminResult dto.AdminLoginResult
	adminInfo, err := a.AdminUserRepository.GetAdminUserInfo(tx, id)
	if err != nil {
		if err == errors.IsNotAdminUser {
			log.Println("관리자 계정이 아닌 사용자 접근", id)
		}else {
			log.Println("DB에서 정보 획득 실패", err)
		}
		adminResult.Result = false
		adminResult.IsAdmin = false
		adminResult.AdminInfo = dto.AdminInfo{Id : "", Rank : 0}
		jsonResponse, err := json.Marshal(adminResult)
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
		adminResult.Result = false
		adminResult.IsAdmin = false
		adminResult.AdminInfo = dto.AdminInfo{Id : "", Rank : 0}
		jsonResponse, err := json.Marshal(adminResult)
		if err != nil {
			log.Println("tx Commit 오류", err)
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

	if comparePassword(adminInfo.Password, reqPassword) {
		adminResult.Result = false
		adminResult.IsAdmin = true
		adminResult.AdminInfo = dto.AdminInfo{Id : "", Rank : 0}
		jsonResponse, err := json.Marshal(adminResult)
		if err != nil {
			log.Println("Admin 로그인 결과 Marshal 오류", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return 
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(jsonResponse)
		return 
	}

	adminResult.Result = true
	adminResult.IsAdmin = true
	adminResult.AdminInfo = dto.AdminInfo{Id : adminInfo.Id, Rank : adminInfo.Rank}
	jsonResponse, err := json.Marshal(adminResult)
	if err != nil {
		log.Println("Admin 로그인 결과 Marshal 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return 
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}

// admin 계정 로그아웃
func (a *Admin)Logout(res http.ResponseWriter, req *http.Request) {
	a.Session.DeleteSession(adminSession, res, req)
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
	if err != nil {
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

	err = adminRepo.CreateRootAdmin(tx, rootAdminID, string(hashedPassword))
	if err != nil {
		log.Fatal("RootAdmin 계정 생성 실패", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Root 관리자 계정 생성 실패", err)
		return
	}
	return 
}