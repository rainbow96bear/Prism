package pkg

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"prism_back/errors"

	"github.com/nfnt/resize"
)
var (
	profileImgWidth uint = 200
	profileImgHeight uint = 200
	assetsFolder = os.Getenv("RELATIVE_IMAGE_DIRECTORY")

)

type Images struct {
	
}

// formdata에서 images로 전달 받은 파일을 얻기
func (i *Images) GetImageFromReq(req *http.Request) (multipart.File, *multipart.FileHeader, error) {
	file, handler, err := req.FormFile("image")
	if err != nil {
		if err ==  http.ErrMissingFile {
			return nil, nil, errors.EmptyFile
		}
		return nil, nil, err
	}

	return file, handler, nil
}


// filePath에 fileName이른을 가지는 파일을 생성
func (i *Images)CreateNewImageFile(filePath, fileName string) (*os.File, error) {
	newFileName := filepath.Join(assetsFolder, filePath, fileName)
	newFile, err := os.Create(newFileName)
	if err != nil {
		return nil, err
	}
	return newFile, nil
}

// srcFile의 내용을 dstFile에 복사
func (i *Images)CopyFile(dstFile io.Writer, srcFile io.Reader) (error) {
	_, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

// handler의 확장자를 확인하여 profile 이미지의 크기를 조정
func (i *Images)ResizingForProfile(handler *multipart.FileHeader, srcFile io.Reader) (image.Image, error){
	var (
		img image.Image
		err error
	)

	if filepath.Ext(handler.Filename) == ".png" {
		img, err = png.Decode(srcFile)
		if err != nil {
			return nil, err
		}
	}else {
		img, _, err = image.Decode(srcFile)
		if err != nil {
			return nil, err
		}
	}

	// resizedImg := resize.Resize(profileImgWidth, profileImgWidth, img, resize.Lanczos3)
	resizedImg := resize.Thumbnail(profileImgWidth, profileImgHeight, img, resize.Lanczos3)
	
	return resizedImg, nil
}

// filePath와 fileName에 img를 JPEG 형식으로 저장
func (i *Images)EncodeForJPEG(filePath, fileName string, img image.Image) (error){
	dstFile, err := i.CreateNewImageFile(filePath, fileName)
	if err != nil {
		return err
	}
	err = jpeg.Encode(dstFile, img, nil)
	if err != nil {
		return err
	}
	return nil
}

func (i *Images)DownLoadImgFromURL(url, filePath, fileName, extension string) (error) {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// HTTP 응답이 성공적인지 확인합니다.
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP 요청이 실패했습니다. 상태 코드: %d", response.StatusCode)
	}
	newFileName := fmt.Sprintf("%s%s", fileName, extension)
	file, err := i.CreateNewImageFile(filePath, newFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = i.CopyFile(file, response.Body)
	if err != nil {
		return err
	}
	log.Println("이미지 다운로드가 완료되었습니다.")
	return nil
}