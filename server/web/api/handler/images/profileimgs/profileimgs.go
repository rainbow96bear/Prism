package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Profileimgs struct {
}

var imageDirectory = os.Getenv("RELATIVE_IMAGE_DIRECTORY")

// 정적 파일 서버 핸들러 생성
var fileServer = http.FileServer(http.Dir(imageDirectory))


func (p *Profileimgs) RegisterHandlers(r *mux.Router) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// 상대 경로를 절대 경로로 변환
	absolutePath := filepath.Join(currentDir, imageDirectory)
	r.PathPrefix("/").Handler(http.StripPrefix("/assets/images/", http.FileServer(http.Dir(absolutePath))))
}