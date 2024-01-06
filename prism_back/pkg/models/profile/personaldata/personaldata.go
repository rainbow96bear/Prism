package personaldata

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"prism_back/internal/Database/mysql"
	"prism_back/internal/session"
)

type PersonalData struct {
	Nickname string	`json:"nickname"`
	Profile_img string `json:"profile_img,omitempty"`
	One_line_introduce string `json:"one_line_introduce,omitempty"`
	HashTag []string `json:"hashtag,omitempty"`
}
func (p *PersonalData)GetPersonalData(res http.ResponseWriter, req *http.Request) {
	personaldata, err := getPersonalData(res, req)
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

func (p *PersonalData)SetPersonalData(res http.ResponseWriter, req *http.Request) {
	// personalData, err := getPersonalDataFromReq(res, req)
	// if err != nil {
	// 	http.Error(res, "Missing 'id' parameter", http.StatusBadRequest)
	// }
}

func getPersonalData(res http.ResponseWriter, req *http.Request) (PersonalData, error) {
	queryValues := req.URL.Query()

	// "id" 쿼리 매개변수 값 얻기
	id := queryValues.Get("id")
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

func SetPersonalDataToDB(personalData PersonalData, req *http.Request) error {
    id, err := session.GetUserID(req)
    if err != nil {
        log.Println("id 조회 실패", err)
        return err
    }

    // profile 테이블 업데이트
    profileQuery := `UPDATE profile SET One_line_introduce = ? WHERE user_info_User_id = ?`
    _, err = mysql.DB.Exec(profileQuery, personalData.One_line_introduce, id)
    if err != nil {
        log.Println("한 줄 소개 업데이트 실패: ", err)
        return err
    }

    // user_info 테이블 업데이트
    userInfoQuery := `UPDATE user_info SET Nickname = ?, Profile_img = ? WHERE User_id = ?`
    _, err = mysql.DB.Exec(userInfoQuery, personalData.Nickname, personalData.Profile_img, id)
    if err != nil {
        log.Println("닉네임 및 프로필 이미지 업데이트 실패: ", err)
        return err
    }

    return nil
}
