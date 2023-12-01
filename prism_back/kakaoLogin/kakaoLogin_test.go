package KakaoLogin

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

type UserInfo struct {
	ID          int       `json:"id"`
	ConnectedAt string `json:"connected_at"`
	KakaoAccount KakaoAccount `json:"kakao_account"`
	Properties  Properties  `json:"properties"`

}

type KakaoAccount struct {
	ProfileNicknameNeedsAgreement bool `json:"profile_nickname_needs_agreement"`
	Profile struct {
		Nickname string `json:"nickname"`
	} `json:"profile"`
}

type Properties map[string]string

func TestKakaoLogin(t *testing.T) {
	requestBody := `{}`
	
	AUTHORIZE_CODE := "7Fv4W7CefMexmGN_A8kqe4MmtYEj4Jc8NoD1HSeLFuHAAmB4lk1dTHOdPKEKKwzTAAABjCUs-mq2W8wW6V7rJg"
	
	// Front에서 AUTHORIZE_CODE를 Back으로 전달
	req := httptest.NewRequest("GET",fmt.Sprintf("/kakaoLogin?code=%s",AUTHORIZE_CODE), strings.NewReader(requestBody))
	req.Header.Set("Content_type", "application/json")
	
	respoenseWriter := httptest.NewRecorder()
	
	//KakaoLogin 실행
	KakaoLogin(respoenseWriter, req)

	var userInfo UserInfo
	fmt.Println(respoenseWriter.Body.String())
	err := json.Unmarshal([]byte(respoenseWriter.Body.String()), &userInfo)
    if err != nil {
		fmt.Println("Failed to parse JSON: ", err)
		return
	}
	nickName := userInfo.KakaoAccount.Profile.Nickname
	if nickName != "" {
		fmt.Println("성공", nickName)
	}

}


	
	// KakaoLogin의 작업
	//KakaoLogin이 'https://kauth.kakao.com/oauth/token'이 URL로 AUTHORIZE_CODE과REST_API_KEY를 전달 REDIRECT_URI도 포함
	// Access Token을 GET 방식으로 "https://kapi.kakao.com/v2/user/me"에 전달 
	// 	curl -v -X GET "https://kapi.kakao.com/v2/user/me" \
	//   -H "Authorization: Bearer ${ACCESS_TOKEN}"