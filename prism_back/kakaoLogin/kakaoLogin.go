package KakaoLogin

import (
	"bytes"
	"encoding/json"
	"time"

	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

// oauth token 정보


var SECRET_KEY string = os.Getenv("SECRET_KEY")
var store = sessions.NewCookieStore([]byte(SECRET_KEY))

var err = godotenv.Load(".env")

//OAuthLogin 시작
func OAuthLogin(res http.ResponseWriter, req *http.Request) {
	REST_API_KEY := os.Getenv("REST_API_KEY")
	REDIRECT_URI := os.Getenv("REDIRECT_URI")
	var redirectURL string
	redirectURL = fmt.Sprintf("https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s",
		REST_API_KEY, REDIRECT_URI)

	http.Redirect(res, req, redirectURL, http.StatusFound)
}
//user정보 처리
func OAuthLoginAfterProcess(res http.ResponseWriter, req *http.Request){
	//
	token, err := GetToken(res, req)
	if err != nil {
		fmt.Println("token 획득 실패 : ", err)
	}
	// Access_token을 이용한 user 정보 받기
	user, err := GetUserInfo(token.Access_token)
	if err != nil {
		fmt.Println("정보 획득 실패: ", err)
	}
	fmt.Println("user정보 확인", user)
	// json.NewEncoder(res).Encode(user)
	MakeCookie(res, user)
	http.Redirect(res, req, "http://localhost:3000/home", http.StatusFound)
}

// 토큰 가져오기
func GetToken(res http.ResponseWriter, req *http.Request) (Token, error) {
	// 반환할 token
	var token Token
	// 사용할 클라이언트 ID, 리다이렉트 URI, 클라이언트 시크릿, 토큰 요청 URI 등을 설정
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

	// HTTP POST 요청을 만들기
	req, err := http.NewRequest("POST", tokenURI, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return token, err
	}

	// 요청 헤더 설정
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// HTTP 클라이언트를 생성하고 요청 실행
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

// user정보 가져오기
func GetUserInfo(Access_token string) (User, error){
	var user User
	// requestURL := "https://kapi.kakao.com/v2/user/me"
	requestURL := "https://kapi.kakao.com/v1/oidc/userinfo"

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return user, fmt.Errorf("Error creating request: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+Access_token)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return user, fmt.Errorf("UserInfo 얻기 오류: %v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, fmt.Errorf("Unmarshal 오류 : %v\n", err)
	}

	return user, nil
}

// cookie 만들기
func MakeCookie(res http.ResponseWriter, user User) {
	cookie := http.Cookie{
		Name : "kakaoLogin",
		Value : fmt.Sprintf("%s",user.ID),
		Expires: time.Now().Add(30 * 24 * time.Hour),
		Path:    "/",
		HttpOnly: true,
	}
	http.SetCookie(res, &cookie)
}