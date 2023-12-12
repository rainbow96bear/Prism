package Session

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var err = godotenv.Load(".env")
var Store *sessions.CookieStore

func SetupStore() {
	Store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
}