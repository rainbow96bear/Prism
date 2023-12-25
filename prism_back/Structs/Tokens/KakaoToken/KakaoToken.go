package KakaoToken

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"prism_back/Interfaces/I_Token"
)

type Token struct {
	Access_token             string `json:"access_token"`
	Token_type               string `json:"token_type"`
	Refresh_token            string `json:"refresh_token"`
	Id_token                 string `json:"id_token"`
	Expires_in               int    `json:"expires_in"`
	Scope                    string `json:"scope"`
	Refresh_token_expires_in int    `json:"refresh_token_expires_in"`
}

func (t *Token)GetToken(res http.ResponseWriter, req *http.Request) (I_Token.I_Token, error){
	token, err := GetToken(res, req)
	return token, err
}

func GetToken(res http.ResponseWriter, req *http.Request) (I_Token.I_Token, error) {
	var token Token
	CLIENT_SECRET_KEY := os.Getenv("CLIENT_SECRET_KEY")
	REST_API_KEY := os.Getenv("REST_API_KEY")
	REDIRECT_URI := os.Getenv("REDIRECT_URI")
	tokenURI := "https://kauth.kakao.com/oauth/token"

	code := req.URL.Query().Get("code")

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", REST_API_KEY)
	data.Set("redirect_uri", REDIRECT_URI)
	data.Set("client_secret", CLIENT_SECRET_KEY)
	data.Set("code", code)

	// HTTP POST 요청 만들기
	req, err := http.NewRequest("POST", tokenURI, bytes.NewBufferString(data.Encode()))
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return nil, fmt.Errorf("Error creating request : %e", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// HTTP 클라이언트 생성하고 요청 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return nil, fmt.Errorf("요청 실행 실패 : %e", err)
	}
	defer resp.Body.Close()

	// HTTP 응답 본문을 문자열로 읽어오기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return nil, fmt.Errorf("응답의 Body 읽기 오류 : %e", err)
	}
	// HTTP 응답 본문 출력
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, fmt.Errorf("JSON 파싱 오류 : %e", err)
	}
	return &token, nil
}