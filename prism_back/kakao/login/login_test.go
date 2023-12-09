package Login

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