package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"prism_back/dto"
	"prism_back/dto/oauth"
	"prism_back/internal/Database/mysql"
	"prism_back/pkg"
	"prism_back/pkg/models"
	"prism_back/repository"
)


var (
	CLIENT_SECRET_KEY string = os.Getenv("CLIENT_SECRET_KEY")
	REST_API_KEY string = os.Getenv("REST_API_KEY")
	REDIRECT_URI string = os.Getenv("REDIRECT_URI")
	FRONT_URL string = os.Getenv("FRONT_DOMAIN")
	userSession string = os.Getenv("USER_SESSION")
	tokenURI string = "https://kauth.kakao.com/oauth/token"
	requestURL string = "https://kapi.kakao.com/v1/oidc/userinfo"
	profilePath string = "/profiles"
	profileImgExtension string = os.Getenv("PROFILE_IMAGE_EXTENSION")
)

type KakaoOAuth struct {
	oauth.KakaoToken
	repository.UserInfoReopository
	repository.ProfileRepository
	pkg.Session
	pkg.Images
}

func (k *KakaoOAuth) Login(res http.ResponseWriter, req *http.Request) {
	token, err := getToken(req)
	if err != nil {
		log.Println("kakao OAuth token 얻기 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	userinfo, err := getUserInfo(token.Access_token)
	if err != nil {
		log.Println("kakao OAuth token 얻기 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	isSaved, err := k.isSaved(userinfo.Id)
	if err != nil {
		log.Println("kakao OAuth token 얻기 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("kakao OAuth token 얻기 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if !isSaved {
		// 사용자 정보 저장
		err = k.UserInfoReopository.Create(tx, models.UserInfo{Id : userinfo.Id, NickName: userinfo.Nickname})
		if err != nil {
			log.Println("kakao OAuth token 얻기 오류", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		// 프로필 생성
		_, err = k.ProfileRepository.Create(tx, userinfo.Id)
		if err != nil {
			log.Println("kakao OAuth token 얻기 오류", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// 프로필 이미지 저장
		err := k.Images.DownLoadImgFromURL(userinfo.ProfileImgURL, profilePath, userinfo.Id, profileImgExtension)
		if err != nil {
			log.Println("kakao OAuth token 얻기 오류", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = k.Session.CreateSession(userSession, userinfo.Id, res, req)
	if err != nil {
		log.Println("kakao OAuth token 얻기 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, fmt.Sprintf("%s/home", FRONT_URL), http.StatusFound)
}

func (k *KakaoOAuth)Logout(res http.ResponseWriter, req *http.Request) {
	err := k.Session.DeleteSession(userSession, res, req)
	if err != nil {
		log.Println("kakao OAuth token 얻기 오류", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getToken(req *http.Request) (oauth.KakaoToken, error) {
	code := req.URL.Query().Get("code")

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", REST_API_KEY)
	data.Set("redirect_uri", REDIRECT_URI)
	data.Set("client_secret", CLIENT_SECRET_KEY)
	data.Set("code", code)
	// HTTP POST 요청 만들기
	newReq, err := http.NewRequest("POST", tokenURI, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return oauth.KakaoToken{}, fmt.Errorf("Error creating request : %e", err)
	}

	newReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// HTTP 클라이언트 생성하고 요청 실행
	client := &http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		return oauth.KakaoToken{}, fmt.Errorf("요청 실행 실패 : %e", err)
	}
	defer resp.Body.Close()

	// HTTP 응답 본문을 문자열로 읽어오기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return oauth.KakaoToken{}, fmt.Errorf("응답의 Body 읽기 오류 : %e", err)
	}
	// HTTP 응답 본문 출력
	var token oauth.KakaoToken
	err = json.Unmarshal(body, &token)
	if err != nil {
		return oauth.KakaoToken{}, fmt.Errorf("JSON 파싱 오류 : %e", err)
	}
	return token, nil
}

func getUserInfo(AccessToken string) (dto.KakaoUser, error) {

	var user dto.KakaoUser
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

func (k *KakaoOAuth)isSaved(id string) (bool, error) {
	tx, err := mysql.DB.Begin()
	if err != nil {
		return false, err
	}
	_, err = k.UserInfoReopository.Read(tx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}