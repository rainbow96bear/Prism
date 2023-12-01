package KakaoLogin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func KakaoLogin (w http.ResponseWriter, r *http.Request) {
	// 쿼리 파라미터에서 인증 코드 추출
	AUTHORIZE_CODE := r.URL.Query().Get("code")
	REST_API_KEY := os.Getenv("REST_API_KEY")
	REDIRECT_URI := os.Getenv("REDIRECT_URI")
	
	//body에 lastURL을 받는다.
	//응답할 때 lastURL을 전달하여 기존 페이지로 이동
	fmt.Println(AUTHORIZE_CODE)
	tokenURL := "https://kauth.kakao.com/oauth/token"
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", REST_API_KEY)
	data.Set("redirect_uri", REDIRECT_URI)
	data.Set("code", AUTHORIZE_CODE)
	response, err := http.PostForm(tokenURL,data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	// curl -v -X POST "https://kauth.kakao.com/oauth/token" \
	// -H "Content-Type: application/x-www-form-urlencoded" \
	// -d "grant_type=authorization_code" \
	// -d "client_id=${REST_API_KEY}" \
	// --data-urlencode "redirect_uri=${REDIRECT_URI}" \
	// -d "code=${AUTHORIZE_CODE}"


	// JSON 마샬링
	jsonData, err := json.Marshal(response.Body)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// HTTP 응답 헤더 설정
	w.Header().Set("Content-Type", "application/json")

	// JSON 응답 전송
	w.Write(jsonData)
}
