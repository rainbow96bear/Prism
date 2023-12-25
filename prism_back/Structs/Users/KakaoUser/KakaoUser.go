package KakaoUser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"prism_back/Interfaces/I_Token"
	"prism_back/Structs/Tokens/KakaoToken"
	Session "prism_back/session"
)

type KakaoUser struct {
	User_id     string `json:"sub"`
	Nickname    string `json:"nickname,omitempty"`
	Profile_img string `json:"picture,omitempty"`
}

// I_Login 인터페이스 메서드
func (k *KakaoUser) Login(res http.ResponseWriter, req *http.Request) {
	oAuthLogin(res, req)
}

func (k *KakaoUser)AfterProcessres(res http.ResponseWriter, req *http.Request) {
	err := oAuthLoginAfterProcess(res, req)
	if err != nil {
		
	}
	http.Redirect(res, req, "http://localhost:3000/home", http.StatusFound)
}

func (k *KakaoUser)Logout(res http.ResponseWriter, req *http.Request) {
	kakaoLogout(res, req)
}

// OAuth 로그인을 위할 Redirection
func oAuthLogin(res http.ResponseWriter, req *http.Request) {
	REST_API_KEY := os.Getenv("REST_API_KEY")
	REDIRECT_URI := os.Getenv("REDIRECT_URI")
	redirectURL := fmt.Sprintf("https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s",
	REST_API_KEY, REDIRECT_URI)
	
	// redirectURL로 redirect
	http.Redirect(res, req, redirectURL, http.StatusFound)
}

func oAuthLoginAfterProcess(res http.ResponseWriter, req *http.Request) (error) {
	kakaoToken := &KakaoToken.Token{}
	// kakao 토큰 받아서 kakaoToken 변수에 저장되길 바랍니다.
	token, err := I_Token.GetToken(kakaoToken, res, req)
	if err != nil {
		return fmt.Errorf("토큰 가져오기 실패 : %e", err)
	}
	// kakaoUser 정보 받아오기
	kakaoUser, err := getUserInfo(token.(*KakaoToken.Token).Access_token)
	if err != nil {
		return fmt.Errorf("사용자 정보 얻어오기 오류 : %e", err)
	}
	err = createSession(kakaoUser, res, req)
	if err != nil {
		return fmt.Errorf("카카오 로그인 세션 생성 실패 : %e", err)
	}
	return nil
}

// User 정보 가져오기
func getUserInfo(AccessToken string) (KakaoUser, error){
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

func createSession(user KakaoUser, res http.ResponseWriter, req *http.Request) error {
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

func kakaoLogout(res http.ResponseWriter, req *http.Request) {
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