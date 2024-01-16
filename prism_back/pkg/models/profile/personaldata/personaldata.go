package personaldata

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"prism_back/internal/Database/mysql"

	"github.com/gorilla/mux"
)

type PersonalData struct {
	Nickname string	`json:"nickname"`
	Profile_img string `json:"profile_img,omitempty"`
	One_line_introduce string `json:"one_line_introduce,omitempty"`
	HashTag []string `json:"hashtag,omitempty"`
}
func (p *PersonalData)GetPersonalData(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	personaldata, err := getPersonalData(id)
	if err != nil {
		http.Error(res, "Missing 'id' parameter", http.StatusBadRequest)
	}
	jsonResponse, err := json.Marshal(personaldata)
	if err != nil {
		log.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}

func (p *PersonalData) SetPersonalData(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	personalData, err := getPersonalDataFromReq(res, req)
	if err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 현재 DB에 저장된 사용자 정보 가져오기
	currentData, err := getPersonalData(id)
	fmt.Println("변경 전 정보",currentData)
	if err != nil {
		http.Error(res, "Failed to retrieve current user data", http.StatusInternalServerError)
		return
	}
	setPersonalDataToDB(personalData, currentData, id)
	// 변경된 값에 대해서만 업데이트
}


func getPersonalData(id string) (PersonalData, error) {
	// "id" 쿼리 매개변수 값 얻기
	// id 값이 빈 문자열인 경우에 대한 처리

	var personaldata PersonalData
	query := `
    SELECT user_info.Nickname, user_info.Profile_img, profile.One_line_introduce
    FROM profile
    JOIN user_info ON profile.user_info_User_id = user_info.User_id
	WHERE profile.user_info_User_id = ?
`
	err := mysql.DB.QueryRow(query, id).Scan(&personaldata.Nickname, &personaldata.Profile_img, &personaldata.One_line_introduce)
	if err != nil {
		log.Println(err)
	}
	return personaldata, nil
}

func getPersonalDataFromReq(res http.ResponseWriter, req *http.Request) (PersonalData, error) {
	
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return PersonalData{}, fmt.Errorf("요청의 Body 읽기 오류 : %e", err)
	}
	var personalData PersonalData
	err = json.Unmarshal(body, &personalData)
	if err != nil {
		return PersonalData{}, fmt.Errorf("JSON 파싱 오류 : %e", err)
	}
	return personalData, nil
}

func setPersonalDataToDB(personalData, currentData PersonalData, id string) error {
	// 닉네임이 변경 된 경우
	if personalData.Nickname != "" && personalData.Nickname != currentData.Nickname {
		userInfoQuery := `UPDATE user_info SET Nickname = ? WHERE User_id = ?`
		_, err := mysql.DB.Exec(userInfoQuery, personalData.Nickname, id)
		if err != nil {
			log.Println("닉네임 업데이트 실패: ", err)
			return err
		}
	}
	// 프로필 사진이 변경된 경우
	fmt.Println(personalData.Profile_img)
	fmt.Println(currentData.Profile_img)
	if personalData.Profile_img != "" && personalData.Profile_img != currentData.Profile_img {
		userInfoQuery := `UPDATE user_info SET, Profile_img = ? WHERE User_id = ?`
		_, err := mysql.DB.Exec(userInfoQuery, personalData.Profile_img, id)
		if err != nil {
			log.Println("프로필 이미지 업데이트 실패: ", err)
			return err
		}
	}
	// 한 줄 소개가 변경된 경우
	if personalData.One_line_introduce != "" && personalData.One_line_introduce != currentData.One_line_introduce {
		profileQuery := `UPDATE profile SET One_line_introduce = ? WHERE user_info_User_id = ?`
		_, err := mysql.DB.Exec(profileQuery, personalData.One_line_introduce, id)
		if err != nil {
			log.Println("한 줄 소개 업데이트 실패: ", err)
			return err
		}
	}
	
    return nil
}
