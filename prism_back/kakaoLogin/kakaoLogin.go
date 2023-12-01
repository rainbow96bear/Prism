package KakaoLogin

import (
	"net/http"
)

func KakaoLogin (w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://kauth.kakao.com/oauth/authorize")
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
}