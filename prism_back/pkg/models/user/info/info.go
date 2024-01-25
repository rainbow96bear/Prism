package info

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"prism_back/internal/Database/mysql"
	"prism_back/internal/session"
)

type Info struct {
	UserID     string `json:"user_id"`
	Nickname   string `json:"nickname"`
}

func (l *Info)GetUserInfo(res http.ResponseWriter, req *http.Request) {
	userID, err := session.GetUserID(req)
	if err != nil {
		log.Println(err)
	}
	info, err := getUserInfo(userID)
	response, err := json.Marshal(info)
	if err != nil {
		// 직렬화 에러 발생 시 에러 출력 후 HTTP 500 에러 응답
		fmt.Println("JSON 직렬화 실패:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	// JSON 형식으로 응답
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func getUserInfo(user_id string) (info Info, err error){
	query := "SELECT `User_id`, `Nickname` FROM user_info WHERE User_id = ?"
	err = mysql.DB.QueryRow(query, user_id).Scan(&info.UserID, &info.Nickname)
	if err != nil{
		return info, err
	}
	return info, nil
}
