package pkg

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

type Session struct {
	store *sessions.CookieStore
}

var Store Session
var err = godotenv.Load(".env")

// session Setup
func (s *Session)SetupStore() {
	s.store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
}

// sessionName을 가지는 session에 사용자의 id 저장
func (s *Session)CreateSession(sessionName, id string, res http.ResponseWriter, req *http.Request) (error) {
	session, err := s.store.Get(req, sessionName)
	if err != nil {
		return err
	}
	session.Values["ID"] = id
	err = session.Save(req, res)
	if err != nil {
		return err
	}
	return nil
}

// session에서 요청을 보낸 사용자의 id 얻기
func (s *Session)GetID(sessionName string, req *http.Request) (string, error) {
	session, err := s.store.Get(req, sessionName)
	if err != nil {
		return "", err
	}
	id, ok := session.Values[sessionName].(string)
	if !ok {
		id = ""
	}
	return id, nil
}

// sessionName에 해당하는 session 삭제
func (s *Session)DeleteSession(sessionName string, res http.ResponseWriter, req *http.Request) (err error) {
	http.SetCookie(res, &http.Cookie{
		Name:   sessionName,
		Value:  "",
		MaxAge: -1,
		Domain: os.Getenv("COOKIE_DOMAIN"),
		Path:   "/",
		HttpOnly: true,
	})
	session, err := s.store.Get(req, sessionName)
    if err == nil {
        session.Options.MaxAge = -1
        session.Save(req, res)
    }
	return nil
}