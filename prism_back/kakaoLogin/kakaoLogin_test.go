package KakaoLogin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetToken(t * testing.T) {

	resRecoder := httptest.NewRecorder()
	
	code := ""
	req := httptest.NewRequest("GET", "/kakao/withToken?code="+code,nil)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	token, err := GetToken(resRecoder, req)

	if err != nil {
		fmt.Println("token 획득 실패")
	}
	fmt.Println("테스트 결과 : ",token)
}

func TestGetUserInfo(t *testing.T) {
	access_token := "" 
	user, err := GetUserInfo(access_token)
	if err != nil {
		fmt.Println(err)
	}
	expectedUserID := ""
	
	if user.NickName == expectedUserID {
		fmt.Println("성공")
		fmt.Println(user)
	}
}


	
	// KakaoLogin의 작업
	//KakaoLogin이 'https://kauth.kakao.com/oauth/token'이 URL로 AUTHORIZE_CODE과REST_API_KEY를 전달 REDIRECT_URI도 포함
	// Access Token을 GET 방식으로 "https://kapi.kakao.com/v2/user/me"에 전달 
	// 	curl -v -X GET "https://kapi.kakao.com/v2/user/me" \
	//   -H "Authorization: Bearer ${ACCESS_TOKEN}"