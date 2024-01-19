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

	"github.com/gorilla/mux"
)

// 업로드된 이미지를 저장할 폴더의 경로를 지정합니다.
var imageFolder = fmt.Sprintf("%s%s",  os.Getenv("RELATIVE_IMAGE_DIRECTORY"), "/profiles/")

func DownloadImageFromKakao(url, id string) error {
	err := downloadImage(url, id)
	if err != nil {
		return err
	}
	return nil
}

func UploadImageHandler(res http.ResponseWriter, req *http.Request){
	upload(res, req)
}

func GetImageHandler(res http.ResponseWriter, req *http.Request){
	getImageHandler(res, req)
}

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
	filePath := filepath.Join(imageFolder, fileName, ".jpg")
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

func downloadImage(url, id string) error {
	// HTTP GET 요청을 보냅니다.
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// HTTP 응답이 성공적인지 확인합니다.
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP 요청이 실패했습니다. 상태 코드: %d", response.StatusCode)
	}

	extension := ".jpg"
	
	// 로컬 파일을 생성합니다. 파일 이름에 id 값을 포함합니다.
	fileName := fmt.Sprintf("%s%s", id, extension)
	ilePath := filepath.Join(imageFolder, fileName)
	file, err := os.Create(ilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 이미지 데이터를 로컬 파일에 복사합니다.
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("이미지 다운로드가 완료되었습니다. 파일 경로: %s\n", fileName)
	return nil
}

func getImageHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	// 이미지 파일의 실제 경로를 생성
	filePath := "../assets/profile/" + id + ".jpg"
	fmt.Println(id)
	// 이미지 파일 서빙
	http.ServeFile(res, req, filePath)
}