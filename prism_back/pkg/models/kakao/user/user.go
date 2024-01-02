package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"prism_back/internal/Database/mysql"
	"prism_back/internal/session"
	"prism_back/pkg/interface/basetoken"
	"prism_back/pkg/models/kakao/token"
)

type KakaoUser struct {
	User_id     string `json:"sub"`
	Nickname    string `json:"nickname,omitempty"`
	Profile_img string `json:"picture,omitempty"`
}

// I_Login 인터페이스 메서드
func (k *KakaoUser) Login(res http.ResponseWriter, req *http.Request) {
	// OAuth 로그인을 위한 Redirect URL 생성
	URL, err := getRedirectURL()
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	// Redirect URL로 이동
	http.Redirect(res, req, URL, http.StatusFound)
}

// OAuth의 Redirect URL에 대한 처리
func (k *KakaoUser)GetUserInfo(res http.ResponseWriter, req *http.Request) {
	kakao_Token := &token.Token{}
	// OAuth로 받은 code로 Token 얻기
	baseToken, err := basetoken.GetToken(kakao_Token, res, req)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// 얻은 Token으로 kakaoUser 정보 얻기
	kakaoUser, err := getUserInfo(baseToken.(*token.Token).Access_token)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	if !isSavedID(kakaoUser.User_id) {
		query := "INSERT INTO user_info (User_id, Nickname, Profile_img) VALUES (?, ?, ?)"
		_, err := mysql.DB.Exec(query, kakaoUser.User_id, kakaoUser.Nickname, kakaoUser.Profile_img)
		if err != nil {
			log.Println("사용자 정보 저장 실패 : ", err)
		}
	}
	// kakaoUser 정보로 session 만들기
	err = createSession(kakaoUser, res, req)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(res, req, fmt.Sprintf("%s/home", os.Getenv("DOMAIN")), http.StatusFound)
}

func (k *KakaoUser)Logout(res http.ResponseWriter, req *http.Request) {
	err := kakaoLogout(res, req)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// OAuth 로그인을 위할 Redirection
func getRedirectURL() (redirect_uri string, err error) {
	REST_API_KEY := os.Getenv("REST_API_KEY")
	REDIRECT_URI := os.Getenv("REDIRECT_URI")
	redirectURL := fmt.Sprintf("https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s",
	REST_API_KEY, REDIRECT_URI)
	
	// redirectURL로 redirect
	return redirectURL, nil
}

// User 정보 가져오기
func getUserInfo(AccessToken string) (kakaoUser KakaoUser, err error){
	var user KakaoUser
	requestURL := "https://kapi.kakao.com/v1/oidc/userinfo"

	// 요청 생성
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return user, fmt.Errorf("Error creating request: %v\n", err)
	}

	// 요청의 Header에 AccessToken을 추가
	req.Header.Set("Authorization", "Bearer "+AccessToken)
	client := &http.Client{}

	// 요청 실행
	resp, err := client.Do(req)
	if err != nil {
		return user, fmt.Errorf("UserInfo 얻기 오류: %v\n", err)
	}

	// 응답에 대한 body 저장
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	// body의 내용을 Unmarshal
	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, fmt.Errorf("Unmarshal 오류 : %v\n", err)
	}

	return user, nil
}


// 세션과 쿠키 생성
func createSession(user KakaoUser, res http.ResponseWriter, req *http.Request) (err error) {
	session, err := session.Store.Get(req, "user_login")
	if err != nil {
		return fmt.Errorf("세션을 가져오는데 문제 발생 : %e", err)
	}
	session.Values["User_ID"] = user.User_id
	session.Values["User_ProfileImg"] = user.Profile_img
	err = session.Save(req, res)

	if err != nil {
		return fmt.Errorf("세션을 저장하는데 문제 발생 : %e", err)
	}
	return nil
}


// 로그아웃
func kakaoLogout(res http.ResponseWriter, req *http.Request) (err error) {
	session, err := session.Store.Get(req, "user_login")
	if err != nil {
		return fmt.Errorf("세션 불러오기 실패 : %e", err)
	}

	session.Values["User_ID"] = nil
	session.Values["User_ProfileImg"] = nil
	err = session.Save(req, res)

	if err != nil {
		return fmt.Errorf("세션 저장 실패 : %e", err)
	}

	// 브라우저에 저장된 쿠키를 만료시켜 제거
	http.SetCookie(res, &http.Cookie{
		Name:   "user_login",
		Value:  "",
		MaxAge: -1,
		Domain: os.Getenv("COOKIE_DOMAIN"),
		Path:   "/",
	})

	return nil
}

func isSavedID(user_id string) (isSaved bool){
	query := "SELECT User_id FROM user_info WHERE User_id = ?"
	userID := ""
	err := mysql.DB.QueryRow(query, user_id).Scan(&userID)
	if err != nil{
		return false
	}
	return true
}