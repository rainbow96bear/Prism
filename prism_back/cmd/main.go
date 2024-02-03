// main.go
package main

import (
	"log"
	"net/http"
	"os"
	"prism_back/api/router"
	"prism_back/internal/Database/mysql"
	"prism_back/internal/session"
	"prism_back/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Router struct {
	// router.Assets
	// router.API
}

var (
	assets router.Assets = router.Assets{}
	api router.API = router.API{}
)

func init(){
	mysql.SetupDB()
	session.SetupStore()
	service.InitRootAdmin()
}

func main() {
	port := 8080
	r := mux.NewRouter()
	
	r.Use(corsMiddleware)

	api.Router(r.PathPrefix("/api").Subrouter())
	assets.Router(r.PathPrefix("/assets").Subrouter())

	log.Println("Prism Server Starting on Port :", port)
	
	// 라우터에 CORS 미들웨어 추가
	http.Handle("/", corsMiddleware(r))

	// 서버 시작
	http.ListenAndServe(":8080", nil)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONT_DOMAIN"))
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if req.Method == "OPTIONS" {
			res.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(res, req)
	})
}