package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"prism_back/dto"
	"prism_back/errors"
	"prism_back/internal/Database/mysql"
	"prism_back/pkg"
	"prism_back/repository"

	"github.com/gorilla/mux"
)

type Profile struct {
	userinfo repository.UserInfoReopository
	profile repository.ProfileRepository
	hashtag repository.HashtagRepository
	images pkg.Images
}

// service - profile 정보 제공
func (p *Profile)GetUserProfile(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("service/profiles.go : DB Begin 오류", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	nickname, err := p.getNickname(tx, id)
	if err != nil {
		log.Println("service/profiles.go : 닉네임 얻기 오류", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	one_line_introduce, err := p.getOneLineIntroduce(tx, id)
	if err != nil {
		log.Println("service/profiles.go : 한 줄 소개 얻기 오류", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	hashtagList, err := p.getHashtag(tx, id)
	if err != nil {
		log.Println("service/profiles.go : 해시태그 얻기 오류", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	dto := dto.Personaldata{
		Id: id,
		Nickname: nickname,
		One_line_introduce: one_line_introduce,
		Hashtag: hashtagList,
	}

	responseJSON, err := json.Marshal(dto)
	if err != nil {
		log.Println("service/profiles.go : json marshal 오류", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content/type", "application/json")
	res.Write(responseJSON)
}

// service - 요청에서 얻은 data를 db에 저장
func (p *Profile)UpdateUserProfile(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	tx, err := mysql.DB.Begin()
	if err != nil {
		log.Println("service/profiles.go :", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	dto, err := p.getProfileDataFromReq(req)
	if err != nil {
		log.Println("service/profiles.go : req에서 profile 내용 얻기 오류", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = p.userinfo.Update(tx, id, dto.Nickname)
	if err != nil {
		log.Println("service/profiles.go : profile 닉네임 업데이트 오류", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	_, err = p.profile.Update(tx, id, dto.One_line_introduce)
	if err != nil {
		log.Println("service/profiles.go : profile 한 줄 소개 업데이트 오류", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	_, err = p.hashtag.Update(tx, id, dto.Hashtag)
	if err != nil {
		log.Println("service/profiles.go : profile 해시태그 업데이트 오류", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	err = tx.Commit()
	if err != nil {
		log.Println("service/profiles.go: DB 트랜잭션 커밋 오류", err)
		tx.Rollback()
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	err = p.uploadProfileImg(req, id)
	if err != nil && err != errors.EmptyFile {
		log.Println("service/profiles.go : profile 이미지 업데이트 오류", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 성공적인 응답
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Profile updated successfully"))
}

// service - repository에 요청하여 id에 해당하는 nickname 얻기
func (p *Profile)getNickname(tx *sql.Tx, id string) (string, error){
	userinfo, err := p.userinfo.Read(tx, id)
	if err != nil && err != errors.NotSavedUser{
		return "", fmt.Errorf("Nickname 얻기 실패 %e", err)
	}
	return userinfo.NickName, nil
}

// service - repository에 요청하여 id에 해당하는 한 줄 소개 얻기
func (p *Profile)getOneLineIntroduce(tx *sql.Tx, id string) (string, error) {
	userProfile, err := p.profile.Read(tx, id)
	if err != nil {
		return "", fmt.Errorf("OneLineIntroduce 얻기 실패 %e", err)
	}
	return userProfile.One_line_introduce, nil
}

// service - repository에 요청하여 id에 해당하는 hashtag list 얻기
func (p *Profile)getHashtag(tx *sql.Tx, id string) ([]string, error) {
	hashtagList, err := p.hashtag.Read(tx, id)
	if err != nil {
		return []string{}, fmt.Errorf("hashtag 얻기 실패 %e", err)
	}
	return hashtagList, nil
}

// service - 요청에서 profile nickname과 한 줄 소개 hashtag 얻기
func (p *Profile)getProfileDataFromReq(req *http.Request) (dto.Personaldata, error) {
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		return dto.Personaldata{}, fmt.Errorf("req에서 personaldata 얻기 실패 %e", err)
	}

	nickname := req.FormValue("nickname")
	oneLineIntroduce := req.FormValue("one_line_introduce")
	hashtags := req.FormValue("hashtags")

	var hashtaglist []string
	err = json.Unmarshal([]byte(hashtags), &hashtaglist)
	if err != nil {
		return dto.Personaldata{}, fmt.Errorf("hashtag Unmarshal 오류 %e", err)
	}
	dto := dto.Personaldata{Nickname: nickname, One_line_introduce: oneLineIntroduce, Hashtag: hashtaglist}
	return dto, nil
}


func (p *Profile)uploadProfileImg(req *http.Request, id string) (error) {
	fileName := fmt.Sprintf("%s%s", id, os.Getenv("PROFILE_IMAGE_EXTENSION"))
	file, handler, err := p.images.GetImageFromReq(req)
	if err != nil{
		if err == errors.EmptyFile {
			return err
		}
		return fmt.Errorf("Req에서 image 파일 얻기 오류 %e", err)
	}
	
	image, err := p.images.ResizingForProfile(handler, file)
	if err != nil {
		return fmt.Errorf("resizing 오류 %e", err)
	}

	err = p.images.EncodeForJPEG(profileFolder, fileName, image)
	if err != nil {
		return fmt.Errorf("JEPG로 Encode 오류 %e", err)
	}
	
	defer file.Close()
	return nil
}