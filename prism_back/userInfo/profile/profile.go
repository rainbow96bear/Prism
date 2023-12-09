package Profile

import (
	"net/http"

	"github.com/joho/godotenv"
)

var err = godotenv.Load(".env")

func GetUserProfile(res http.ResponseWriter, req *http.Request) {
	
}