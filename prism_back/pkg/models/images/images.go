package images

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)
type Image struct {

}
func (i *Image) UploadImageHandler(res http.ResponseWriter, req *http.Request){
	fmt.Println("요청")
	upload(res, req)
	fmt.Println("요청 끝")
}
// 업로드된 이미지를 저장할 폴더의 경로를 지정합니다.
const uploadFolder = "../assets/profile"

func upload(res http.ResponseWriter, req *http.Request) {
	// 폼 데이터에서 파일 가져오기
	file, _, err := req.FormFile("file")
	if err != nil {
		http.Error(res, "Failed to read form file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 새로운 파일명 생성 (예: timestamp)
	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("%d_%s", timestamp, getFileName(file))

	// 파일을 업로드할 경로 생성
	filePath := filepath.Join(uploadFolder, fileName)
	fmt.Println(filePath)
	// 파일 생성
	newFile, err := os.Create(filePath)
	if err != nil {
		http.Error(res, "Failed to create new file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// 파일 복사
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(res, "Failed to copy file", http.StatusInternalServerError)
		return
	}

	// 업로드된 파일의 경로 응답
	jsonResponse := map[string]string{"filePath": filePath}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(jsonResponse)
}

// 파일명을 얻기 위한 헬퍼 함수입니다.
func getFileName(file multipart.File) string {
	// 파일의 헤더 정보를 얻습니다.
	fileHeader := make([]byte, 512)
	_, err := file.Read(fileHeader)
	if err != nil {
		// 실패하면 기본적으로 timestamp로 생성합니다.
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	file.Seek(0, 0)

	// 파일의 MIME 타입을 얻습니다.
	fileType := http.DetectContentType(fileHeader)

	// 파일명 생성 (MIME 타입 기반)
	ext, err := mime.ExtensionsByType(fileType)
	if err != nil || len(ext) == 0 {
		// 실패하면 기본적으로 timestamp로 생성합니다.
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	// 확장자를 추출합니다.
	extension := ext[0]

	// 새로운 파일명 생성 (예: timestamp + 확장자)
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
}
