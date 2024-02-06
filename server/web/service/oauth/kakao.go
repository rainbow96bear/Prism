package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"prism/web/dto"
	"prism/web/dto/oauth"
	"prism/web/errors"
	"prism/web/internal/Database/mysql"
	"prism/web/internal/session"
	"prism/web/pkg"
	"prism/web/pkg/models"
	"prism/web/repository"
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
	pkg.Images
}

func (k *KakaoOAuth) Login(res http.ResponseWriter, req *http.Request) {
	token, err := getToken(req)
	if err != nil {
		log.Println("service/oauth.go : kakao OAuth token 얻기 오류", err)
		http.Redirect(res, req, FRONT_URL, http.StatusInternalServerError)
		return
	}
	userinfo, err := getUserInfo(token.Access_token)
	if err != nil {
		log.Println("service/oauth.go : Access token으로 사용자 정보 얻기 오류", err)
		http.Redirect(res, req, FRONT_URL, http.StatusInternalServerError)
		return
	}

	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("service/oauth.go : DB 시작 오류", err)
		tx.Rollback()
		http.Redirect(res, req, FRONT_URL, http.StatusInternalServerError)
		return
	}

	isSaved, err := k.isSaved(tx, userinfo.Id)
	if err != nil {
		log.Println("service/oauth.go : 등록된 사용자 확인 오류", err)
		http.Redirect(res, req, FRONT_URL, http.StatusInternalServerError)
		return
	}
	
	if !isSaved {
		
		// 사용자 정보 저장
		err = k.UserInfoReopository.Create(tx, models.UserInfo{Id : userinfo.Id, NickName: userinfo.Nickname})
		if err != nil {
			log.Println("service/oauth.go : 사용자 정보 저장 오류", err)
			tx.Rollback()
			http.Redirect(res, req, FRONT_URL, http.StatusInternalServerError)
			return
		}

		// 프로필 생성
		_, err = k.ProfileRepository.Create(tx, userinfo.Id)
		if err != nil {
			log.Println("service/oauth.go : 사용자 프로필 생성 오류", err)
			tx.Rollback()
			http.Redirect(res, req, FRONT_URL, http.StatusInternalServerError)
			return
		}
		// 프로필 이미지 저장
		err := k.Images.DownLoadImgFromURL(userinfo.ProfileImgURL, profilePath, userinfo.Id, profileImgExtension)
		if err != nil {
			log.Println("service/oauth.go : 사용자 프로필 이미지 저장 오류", err)
			http.Redirect(res, req, FRONT_URL, http.StatusInternalServerError)
			return
		}
	}

	err = session.CreateUserSession(userSession, userinfo.Id, res, req)
	if err != nil {
		log.Println("service/oauth.go : 로그인 세션 생성 오류", err)
		http.Redirect(res, req, FRONT_URL, http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("service/oauth.go :", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(res, req, fmt.Sprintf("%s/home", FRONT_URL), http.StatusFound)
}

func (k *KakaoOAuth)Logout(res http.ResponseWriter, req *http.Request) {
	err := session.DeleteSession(userSession, res, req)
	if err != nil {
		log.Println("로그인 세션 제거 오류", err)
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

	user := dto.KakaoUser{}
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

func (k *KakaoOAuth) isSaved(tx *sql.Tx , id string) (bool, error) {	
	// UserInfoReopository를 통해 사용자 정보를 읽어옴
	_, err := k.UserInfoReopository.Read(tx, id)
	// 사용자 정보가 없는 경우 (errors.IsNotSavedUser 에러)
	if err == errors.NotSavedUser {
		// 트랜잭션 롤백 후 저장되지 않은 사용자로 판단
		return false, nil
	}else if err != nil {
		// 사용자 정보를 읽어오는 도중 에러가 발생하면 트랜잭션 롤백 후 에러 반환
		return false, err
	}

	// 사용자 정보가 있는 경우, 트랜잭션을 커밋하고 저장된 사용자로 판단
	if err != nil {
		log.Println("service/oauth.go tx commit error:", err)
		return false, err
	}
	return true, nil
}
