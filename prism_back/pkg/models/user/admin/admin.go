package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"prism_back/internal/Database/mysql"
	"prism_back/internal/session"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	User_id 	string	`json:"id"`
	Rank 		int	 	`json:"nickname,omitempty"`
	Password 	string	`json:"picture,omitempty"`
}

// 현재 사용자가 관리자 계정인지 확인
func (a *Admin)CheckRightAdmin(res http.ResponseWriter, req *http.Request) {

	// 세션으로부터 admin 정보 확인
	admin_info, err := getAdimUserInfoFromSession(res, req)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// 세션으로부터 사용자 정보 확인
	user_id, err := getUserIDFromSession(res, req)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// 현재 사용지가 관리자인지 확인
	isAdmin := false
	if user_id != "" {
		isAdmin = checkAdminByDB(user_id)
	}

	// 관리자가 아니라면 관리자 정보 초기화
	if user_id != "" && admin_info.User_id != user_id {
		admin_info.User_id = ""
		admin_info.Rank = 0
	}
	response := map[string]interface{}{
		"isAdmin": isAdmin,
		"admin_info" : map[string]interface{}{
			"id" : admin_info.User_id,
			"rank" : admin_info.Rank,
		},
    }

    jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("JSON 마샬 에러:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

    res.Header().Set("Content-Type", "application/json")
    res.Write(jsonResponse)
}

// 관리자 계정 로그인
func (a *Admin)Login(res http.ResponseWriter, req *http.Request){

	// 세션으로부터 사용자 정보 확인
	user_id, err := getUserIDFromSession(res, req)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// 요청으로 받은 비밀번호 확인
	requestPassword, err := getPasswordFromRequest(req)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// DB에 저장된 비밀번호 확인
	adminUserInfo, err := getAdminUserInfoFromDB(user_id)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// 비밀번호 비교
	isCorrect := comparePassword(requestPassword, adminUserInfo.Password)
	if !isCorrect {
		user_id = ""
	}
	var responseJSON = map[string]interface{}{
		"isAdmin": isCorrect,
			"admin_info" : map[string]interface{}{
				"id" : user_id,
				"rank" : adminUserInfo.Rank,
			},
	}
	jsonResponse, err := json.Marshal(responseJSON)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// 로그인 성공에 따른 세션 생성
	err = createSession(adminUserInfo, res, req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}


// 관리자 로그아웃
func (a *Admin)Logout(res http.ResponseWriter, req *http.Request){
	adminLogout(res, req)
}

// 세션으로 사용자 정보 확인
func getUserIDFromSession(res http.ResponseWriter, req *http.Request) (string, error) {
	session, err := session.Store.Get(req, "user_login")
	if err != nil {
		return "", fmt.Errorf("user_login 세션 조회 실패 : %e", err)
	}
	user_id, ok := session.Values["User_ID"].(string)
	if !ok {
		user_id = ""
	}
	return user_id, nil
}


// 세션으로 관리자 정보 확인
func getAdimUserInfoFromSession(res http.ResponseWriter, req *http.Request) (Admin, error) {
	session, err := session.Store.Get(req, "admin_login")
	var admin_user Admin
	if err != nil {
		return admin_user, fmt.Errorf("admin_login 세션 조회 실패 : %e", err)
	}
	User_id, ok := session.Values["admin_id"].(string)
	if !ok {
		User_id = ""
	}
	Rank, ok := session.Values["admin_rank"].(int)
	if !ok {
		Rank = 0
	}
	admin_user.User_id = User_id
	admin_user.Rank = Rank
	return admin_user, nil
}


// 요청에서 비밀번호 확인
func getPasswordFromRequest(req *http.Request) (string, error){
	type Password struct{
		Password string `json:"password"`
	}
	var requestData Password
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		return requestData.Password, fmt.Errorf("디코딩 실패 : %e", err)
	}
	return requestData.Password, err
}

// DB에서 관리자 정보 조회
func getAdminUserInfoFromDB(userID string) (Admin, error){
	var admin_user Admin
	err := mysql.DB.QueryRow("SELECT `Admin_id`, `Rank`, `Password` FROM admin_user WHERE Admin_id = ?", userID).Scan(&admin_user.User_id, &admin_user.Rank, &admin_user.Password)
	if err != nil {
		return admin_user, fmt.Errorf("DB 조회 실패 : %e", err)
	}
	return admin_user, err
}


// DB에 저장된 관리자 ID인지 확인
func checkAdminByDB(userID string) (bool){
	row := mysql.DB.QueryRow("SELECT `Admin_id` FROM admin_user WHERE Admin_id = ?", userID)
	if row.Err() != nil {
		return false
	}
	return true
}


// 비밀번호 비교
func comparePassword(requestPassword, DBPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(DBPassword), []byte(requestPassword))
	if err != nil {
		return false
	}
	return true
}


// 관리자 로그아웃
func adminLogout(res http.ResponseWriter, req *http.Request){
	http.SetCookie(res, &http.Cookie{
		Name:   "admin_login",
		Value:  "",
		MaxAge: -1,
		Domain: os.Getenv("COOKIE_DOMAIN"),
		Path:   "/",
	})
}


// 세션 생성
func createSession(admin_info Admin, res http.ResponseWriter, req *http.Request) (error) {
	session, err := session.Store.Get(req, "admin_login")
	if err != nil {
		return fmt.Errorf("세션을 가져오는데 문제 발생 : %e", err)
	}

	session.Values["admin_id"] = admin_info.User_id
	session.Values["admin_rank"] = admin_info.Rank
	err = session.Save(req, res)
	if err != nil {
		return fmt.Errorf("세션을 저장하는데 문제 발생 : %e", err)
	}
	return nil
}