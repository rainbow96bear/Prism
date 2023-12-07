package KakaoLogin

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")
var store = sessions.NewCookieStore([]byte(SECRET_KEY))

var err = godotenv.Load(".env")
var MainDB *sql.DB

func setupDB() {
    var err error
    MainDB, err = sql.Open("mysql", "root:0000@tcp(localhost:3306)/prism")
    if err != nil {
        fmt.Println("Failed to open test DB:", err)
        return
    }

    // 데이터베이스에 연결되었는지 확인
    if err := MainDB.Ping(); err != nil {
        fmt.Println("Failed to connect to test DB:", err)
        return
    }

    fmt.Println("Connected to test DB")
}

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
	user, err := GetUserInfo(token.Access_token)
	if err != nil {
		fmt.Println("정보 획득 실패: ", err)
	}
	setupDB()
	fmt.Println(MainDB)
	isSavedID, err := IsSavedID(user, MainDB)
	if err != nil {
		fmt.Println(err)
	}
	if !isSavedID && err == nil {
		err = C_UserInfo(user, MainDB)
		if err != nil {
			fmt.Println("User 정보 저장 실패")
		}
	}
	 
	MakeCookie(res, user)
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

	return token, nil
}

// User 정보 가져오기
func GetUserInfo(AccessToken string) (User, error) {
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

// Cookie 만들기
func MakeCookie(res http.ResponseWriter, user User) {
	encodedNickName := url.QueryEscape(user.NickName)
	cookie := http.Cookie{
		Name:    "kakaoUser",
		Value:   fmt.Sprintf("ID=%s,NickName=%s,Img=%s", user.ID, encodedNickName, user.ProfileImg),
		Expires: time.Now().Add(30 * 24 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(res, &cookie)
}

//DB에 기록된 ID인지 확인하기
func IsSavedID(user User, db *sql.DB) (bool, error) {
	list, err := R_UserInfo(user, db)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if len(list) != 0 {
		return true, nil
	}
	return false, nil
}

// User 정보 읽기
func R_UserInfo(user User, db *sql.DB) ([]User, error) {
    query := "SELECT * FROM userinfo WHERE id = ?"
	id, _ := strconv.Atoi(user.ID)
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("User Id 값 읽기 실패")
	}
	result := []User{}
	for rows.Next() {
		var data User
		if err := rows.Scan(&data.ID, &data.NickName, &data.ProfileImg); err != nil {
			return nil, err
		}
		result = append(result, data)  // 이 부분이 누락되어 있었습니다.
	}
    return result, nil
}

func C_UserInfo(user User, db *sql.DB) error {
	query := "INSERT INTO userinfo (ID, NickName, ProfileImg) VALUES (?, ?, ?)"
	result, err := db.Exec(query, user.ID, user.NickName, user.ProfileImg)
	fmt.Println("C_UesrInfo 결과 : ", result)
	if err != nil {
		return fmt.Errorf("사용자 정보 저장 실패")
	}
	return nil
}
