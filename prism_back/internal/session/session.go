package session

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	Store *sessions.CookieStore
 	err = godotenv.Load(".env")
	CookieDomain string = os.Getenv("COOKIE_DOMAIN")
	UserID string = "ID"
	Rank string = "Rank"
)
// session Setup
func SetupStore() {
	Store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
}

// sessionName을 가지는 session에 사용자의 id 저장
func CreateAdminSession(sessionName, id string, rank int, res http.ResponseWriter, req *http.Request) (error) {
	session, err := Store.Get(req, sessionName)
	if err != nil {
		return err
	}
	session.Values[UserID] = id
	session.Values[Rank] = rank
	err = session.Save(req, res)
	if err != nil {
		return err
	}
	return nil
}

func CreateUserSession(sessionName, id string, res http.ResponseWriter, req *http.Request) (error) {
	session, err := Store.Get(req, sessionName)
	if err != nil {
		return err
	}
	session.Values[UserID] = id
	err = session.Save(req, res)
	if err != nil {
		return err
	}
	return nil
}

// session에서 요청을 보낸 사용자의 id 얻기
func GetID(sessionName string, req *http.Request) (string, error) {
	session, err := Store.Get(req, sessionName)
	if err != nil {
		return "", err
	}
	id, ok := session.Values[UserID].(string)
	if !ok {
		id = ""
	}
	return id, nil
}

// session에서 요청을 보낸 사용자의 id 얻기
func GetRank(sessionName string, req *http.Request) (int, error) {
	session, err := Store.Get(req, sessionName)
	if err != nil {
		return -1, err
	}
	rank, ok := session.Values[Rank].(int)
	if !ok {
		rank = -1
	}
	return rank, nil
}

// sessionName에 해당하는 session 삭제
func DeleteSession(sessionName string, res http.ResponseWriter, req *http.Request) (err error) {
	http.SetCookie(res, &http.Cookie{
		Name:   sessionName,
		Value:  "",
		MaxAge: -1,
		Domain: CookieDomain,
		Path:   "/",
		HttpOnly: true,
	})
	session, err := Store.Get(req, sessionName)
    if err == nil {
        session.Options.MaxAge = -1
        session.Save(req, res)
    }
	return nil
}