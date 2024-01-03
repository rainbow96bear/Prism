package personaldata

import (
	"encoding/json"
	"log"
	"net/http"
	"prism_back/internal/Database/mysql"
)

type PersonalData struct {
	Nickname string	`json:"nickname"`
	Profile_img string `json:"profile_img,omitempty"`
	One_line_introduce string `json:"one_line_introduce,omitempty"`
	HashTag []string `json:"hashtag,omitempty"`
}
func (p *PersonalData)GetPersonalData(res http.ResponseWriter, req *http.Request) {
	getPersonalData(res, req)
}

func getPersonalData(res http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()

	// "id" 쿼리 매개변수 값 얻기
	id := queryValues.Get("id")
	// id 값이 빈 문자열인 경우에 대한 처리
	if id == "" {
		http.Error(res, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	var personaldata PersonalData
	query := `
    SELECT user_info.Nickname, user_info.Profile_img, profile.One_line_introduce
    FROM profile
    JOIN user_info ON profile.user_info_User_id = user_info.User_id
`
	err := mysql.DB.QueryRow(query).Scan(&personaldata.Nickname, &personaldata.Profile_img, &personaldata.One_line_introduce)
	if err != nil {
		log.Println(err)
	}
	jsonResponse, err := json.Marshal(personaldata)
	if err != nil {
		log.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}