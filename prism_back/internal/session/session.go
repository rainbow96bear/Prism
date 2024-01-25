package session

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)


var Store *sessions.CookieStore
var err = godotenv.Load(".env")

func SetupStore() {
	Store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
}

func GetUserID(req *http.Request) (userID string, err error){
	session, err := Store.Get(req, "user_login")
	if err != nil {
		return "", fmt.Errorf("user_login 세션 조회 실패 : %e", err)
	}
	user_id, ok := session.Values["User_ID"].(string)
	if !ok {
		user_id = ""
	}
	return user_id, nil
}