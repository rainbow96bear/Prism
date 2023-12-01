package kakaoLogin

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestKakaoLogin(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(KakaoLogin))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expected := "Expected response from Kakao"
	if string(body) != expected {
		t.Errorf("Expected %q, got %q", expected, string(body))
	}
}