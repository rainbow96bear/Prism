package admin

import (
	"database/sql"
	"encoding/json"
	"errors"
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
	Rank 		int	 	`json:"Admin_rank,omitempty"`
	Password 	string	`json:"password,omitempty"`
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
	user_id, err := session.GetUserID(req)
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
	user_id, err := session.GetUserID(req)
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

// 세션으로 관리자 정보 확인
func getAdimUserInfoFromSession(res http.ResponseWriter, req *http.Request) (admin_info Admin, err error) {
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
func getPasswordFromRequest(req *http.Request) (password string, err error) {
	type Password struct{
		Password string `json:"password"`
	}
	var requestData Password
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&requestData)
	if err != nil {
		return requestData.Password, fmt.Errorf("디코딩 실패 : %e", err)
	}
	return requestData.Password, err
}

// DB에서 관리자 정보 조회
func getAdminUserInfoFromDB(userID string) (admin_user Admin, err error){
	err = mysql.DB.QueryRow("SELECT `ID`, `Admin_Rank`, `Password` FROM admin_user WHERE Id = ?", userID).Scan(&admin_user.User_id, &admin_user.Rank, &admin_user.Password)
	if err != nil {
		return admin_user, fmt.Errorf("DB 조회 실패 : %e", err)
	}
	return admin_user, err
}


// DB에 저장된 관리자 ID인지 확인
func checkAdminByDB(userID string) (isAdmin bool) {
    row := mysql.DB.QueryRow("SELECT `Id` FROM admin_user WHERE Id = ?", userID)

    // sql.ErrNoRows와 일치하는 에러인지 확인
    if err := row.Scan(new(int)); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return false // 레코드가 없는 경우
        }
        // 다른 에러인 경우에는 로깅 등을 수행하고 false를 반환하거나 적절한 조치를 취합니다.
        log.Println("에러 발생:", err)
        return false
    }

    return true // 레코드가 있는 경우
}



// 비밀번호 비교
func comparePassword(requestPassword, DBPassword string) (correct bool) {
	err := bcrypt.CompareHashAndPassword([]byte(DBPassword), []byte(requestPassword))
	if err != nil {
		return false
	}
	return true
}


// 관리자 로그아웃
func adminLogout(res http.ResponseWriter, req *http.Request){
    // 쿠키를 무효화하고 삭제
    http.SetCookie(res, &http.Cookie{
        Name:     "admin_login",
        Value:    "",
        MaxAge:   -1,
        Domain:   os.Getenv("COOKIE_DOMAIN"),
        Path:     "/",
        HttpOnly: true,
    })

    // Cache-Control 헤더를 설정하여 쿠키를 캐시하지 않도록 함
    res.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    res.Header().Set("Pragma", "no-cache")
    res.Header().Set("Expires", "0")

    // 현재의 세션을 파기
    session, err := session.Store.Get(req, "admin_login")
    if err == nil {
        session.Options.MaxAge = -1
        session.Save(req, res)
    }
}

// 세션 생성
func createSession(admin_info Admin, res http.ResponseWriter, req *http.Request) (err error) {
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

func MakeRootAdmin() error {

	tx, err := mysql.DB.Begin()
	if err != nil {
		return err
	}

	// 에러 발생 시 롤백
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 환경 변수에서 root admin ID 및 비밀번호 가져오기
	rootAdminID := os.Getenv("ROOT_ADMIN_ID")
	rootAdminPassword := os.Getenv("ROOT_ADMIN_PASSWORD")

	// root admin 계정이 이미 존재하는지 확인
	if checkAdminByDB(rootAdminID) {
		fmt.Println("Root admin 계정이 이미 존재합니다.")
		return nil
	}

	// root admin 비밀번호를 bcrypt를 사용하여 해싱
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rootAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Root admin 비밀번호 해싱 오류:", err)
	}
fmt.Println(string(hashedPassword))

	// admin_user 테이블에 root admin 계정 삽입
	_, err = tx.Exec("INSERT INTO admin_user(ID, Admin_rank, password) VALUES (?, ?, ?)", rootAdminID, 1, string(hashedPassword))
	if err != nil {
		log.Fatal("Root admin 계정 생성 오류:", err)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("Root admin 계정이 성공적으로 생성되었습니다.")
	return nil
}