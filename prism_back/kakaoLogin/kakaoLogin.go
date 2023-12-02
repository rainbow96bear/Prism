package KakaoLogin

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)
var SECRET_KEY string = os.Getenv("SECRET_KEY")
var store = sessions.NewCookieStore([]byte(SECRET_KEY))

func GetAuthorize(res http.ResponseWriter, req *http.Request) {
	err := godotenv.Load(".env")
    
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	REST_API_KEY := os.Getenv("REST_API_KEY")
	REDIRECT_URI := os.Getenv("REDIRECT_URI")
	fmt.Println("REST_API_KEY",REST_API_KEY)
	fmt.Println("REDIRECT_URI",REDIRECT_URI)
	redirectURL := fmt.Sprintf("https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s",
	REST_API_KEY, REDIRECT_URI)
	fmt.Println("redirectURL",redirectURL)

	http.Redirect(res, req, redirectURL, http.StatusFound)
}

func GetToken(res http.ResponseWriter, req *http.Request) {

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
		return
	}

	// 요청 헤더 설정
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// HTTP 클라이언트를 생성하고 요청 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status Code:", resp.Status)

	// HTTP 응답 본문을 문자열로 읽어오기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// HTTP 응답 본문 출력
	fmt.Println("Response Body:", string(body))
	// 응답 처리 및 세션에 토큰 저장
	// ...

	// 클라이언트 리디렉션
	http.Redirect(res, req, "http://localhost:3000/home", http.StatusFound)
}

// func GetUserInfo(res http.ResponseWriter, req *http.Request) {

// }