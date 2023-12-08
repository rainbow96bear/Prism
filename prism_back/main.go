// main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	KakaoLogin "prism_back/kakaoLogin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var db *sql.DB
var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))

func main() {
	port := 8080
	r := mux.NewRouter()

	var err error
	db, err := sql.Open("mysql", "root:"+os.Getenv("MYSQL_PW")+"@tcp(localhost:3306)/prism")
	if err != nil {
		fmt.Println("Failed to open DB")
	}
	KakaoLogin.SetupDB(db)
	KakaoLogin.SetupStore(store)
	defer db.Close()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)	
	
	r.Use(corsMiddleware)
	//Cookie로 로그인 검증
	r.HandleFunc("/userinfo", GetUserInfo_from_Session).Methods("GET")

	//Login 관련
	r.HandleFunc("/kakao/login", KakaoLogin.OAuthLogin).Methods("GET")
	r.HandleFunc("/kakao/withToken", KakaoLogin.OAuthLoginAfterProcess).Methods("GET")
	r.HandleFunc("/kakao/logout", KakaoLogin.Logout).Methods("GET")

	log.Println("Prism Server Starting on Port :", port)
	// 라우터에 CORS 미들웨어 추가
	http.Handle("/", corsMiddleware(r))

	// 서버 시작
	http.ListenAndServe(":8080", nil)
}

func GetUserInfo_from_Session(res http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "user_login")
	if err != nil {
		fmt.Println("세션을 가져오는데 문제 발생:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, ok := session.Values["User_ID"].(string)
	if !ok {
		http.Error(res, "User_ID not found in session", http.StatusInternalServerError)
		return
	}
	User_ProfileImg, ok := session.Values["User_ProfileImg"].(string)
	if !ok {
		http.Error(res, "User_ProfileImg not found in session", http.StatusInternalServerError)
		return
	}

	responseData := KakaoLogin.User{
		ID:         userID,
		ProfileImg: User_ProfileImg,
	}
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSON을 응답으로 전송
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}