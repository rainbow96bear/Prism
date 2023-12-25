package I_Login

import "net/http"

// Interface를 활용한 Login
func Login(u I_User, res http.ResponseWriter, req *http.Request) {
	u.Login(res, req)
}

// User 인터페이스
type I_User interface {
	// 로그인 메서드
	Login(res http.ResponseWriter, req *http.Request)
}