// kakaoLogin.go
package Login

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	Mysql "prism_back/DataBase/MySQL"
	Session "prism_back/session"
	Profile "prism_back/user/profile"

	"github.com/joho/godotenv"
)

var err = godotenv.Load(".env")
var myToken string = ""

// OAuthLogin 시작
func OAuthLogin(res http.ResponseWriter, req *http.Request) {
	REST_API_KEY := os.Getenv("REST_API_KEY")
	REDIRECT_URI := os.Getenv("REDIRECT_URI")
	var redirectURL string
	redirectURL = fmt.Sprintf("https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s",
		REST_API_KEY, REDIRECT_URI)

	http.Redirect(res, req, redirectURL, http.StatusFound)
}

// User 정보 처리
func OAuthLoginAfterProcess(res http.ResponseWriter, req *http.Request) {
	token, err := GetToken(res, req)

	if err != nil {
		fmt.Println("Token 획득 실패: ", err)
	}
	// Access_token을 이용한 user 정보 받기
	user, err := GetUserInfo_from_kakao(token.Access_token)
	if err != nil {
		fmt.Println("정보 획득 실패: ", err)
	}

	isSavedID, err := IsSavedID(user, Mysql.DB)
	if err != nil {
		fmt.Println(err)
	}
	if !isSavedID && err == nil {
		_, err = C_user_info(user, Mysql.DB)
		if err != nil {
			fmt.Println(err)
		}
		_, err = Profile.C_profile(user.User_id, Mysql.DB)
		if err != nil {
			fmt.Println(err)
		}
	}
	 
	err = CreateSession(res, req, user) 
	if err != nil {
		fmt.Println("session 문제 : ", err)
	}
	http.Redirect(res, req, "http://localhost:3000/home", http.StatusFound)
}

// 토큰 가져오기
func GetToken(res http.ResponseWriter, req *http.Request) (Token, error) {
	var token Token

	CLIENT_SECRET_KEY := os.Getenv("CLIENT_SECRET_KEY")
	REST_API_KEY := os.Getenv("REST_API_KEY")
	REDIRECT_URI := os.Getenv("REDIRECT_URI")
	tokenURI := "https://kauth.kakao.com/oauth/token"

	// 요청에서 Authorization Code 가져오기
	code := req.URL.Query().Get("code")
	fmt.Println("임시 테스트를위한 code : ",code)
	// 토큰 요청에 필요한 매개변수 구성
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", REST_API_KEY)
	data.Set("redirect_uri", REDIRECT_URI)
	data.Set("client_secret", CLIENT_SECRET_KEY)
	data.Set("code", code)

	// HTTP POST 요청 만들기
	req, err = http.NewRequest("POST", tokenURI, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return token, err
	}

	// 요청 헤더 설정
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// HTTP 클라이언트 생성하고 요청 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return token, err
	}
	defer resp.Body.Close()

	// HTTP 응답 본문을 문자열로 읽어오기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return token, err
	}

	// HTTP 응답 본문 출력
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("JSON 파싱 오류:", err)
		return token, err
	}
	myToken = token.Access_token
	return token, nil
}

// User 정보 가져오기
func GetUserInfo_from_kakao(AccessToken string) (User, error) {
	var user User
	requestURL := "https://kapi.kakao.com/v1/oidc/userinfo"

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return user, fmt.Errorf("Error creating request: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+AccessToken)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return user, fmt.Errorf("UserInfo 얻기 오류: %v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, fmt.Errorf("Unmarshal 오류 : %v\n", err)
	}

	return user, nil
}

// DB에 기록된 ID인지 확인하기
func IsSavedID(user User, db *sql.DB) (bool, error) {
	user_from_DB, err := R_user_info(user, db)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if user_from_DB.User_id != "" {
		return true, nil
	}
	return false, nil
}

// User 정보 읽기
func R_user_info(user User, db *sql.DB) (User, error) {
	query := "SELECT User_id FROM user_info WHERE User_id = ?"
	rows, err := db.Query(query, user.User_id)
	if err != nil {
		return User{}, fmt.Errorf("User Id 값 읽기 실패")
	}
	result := []User{}
	for rows.Next() {
		var data User
		if err := rows.Scan(&data.User_id); err != nil {
			return User{}, err
		}
		result = append(result, data)
	}
	if len(result) > 0 {
		return result[0], nil
	}
	return User{}, nil
}

func C_user_info(user User, db *sql.DB)  (User, error) {
	query := "INSERT INTO user_info (User_id, Nickname, Profile_img) VALUES (?, ?, ?)"
	_, err := db.Exec(query, user.User_id, user.Nickname, user.Profile_img)
	if err != nil {
		return User{}, fmt.Errorf("사용자 정보 저장 실패")
	}
	return user, nil
}

func C_profile(user User, db *sql.DB) (string, error) {
	query := "INSERT INTO profile (user_info_User_id) VALUES (?)"
	result, err := db.Exec(query, user.User_id)
	fmt.Println("C_UesrInfo 결과 : ", result)
	if err != nil {
		return user.User_id, fmt.Errorf("프로필 생성 실패")
	}
	return user.User_id, nil
}

func CreateSession(res http.ResponseWriter, req *http.Request, user User) error {
	session, err := Session.Store.Get(req, "user_login")
	if err != nil {
		fmt.Println("세션을 가져오는데 문제 발생:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return err
	}
	session.Values["User_ID"] = user.User_id
	session.Values["User_ProfileImg"] = user.Profile_img
	err = session.Save(req, res)

	if err != nil {
		fmt.Println("세션을 저장하는데 문제 발생:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func Logout(res http.ResponseWriter, req *http.Request) {
	session, err := Session.Store.Get(req, "user_login")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["User_ID"] = nil
	session.Values["User_ProfileImg"] = nil
	err = session.Save(req, res)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// 브라우저에 저장된 쿠키를 만료시켜 제거
	http.SetCookie(res, &http.Cookie{
		Name:   "user_login",
		Value:  "",
		MaxAge: -1,
		Domain: os.Getenv("COOKIE_DOMAIN"),
		Path:   "/",
	})

	return
}
