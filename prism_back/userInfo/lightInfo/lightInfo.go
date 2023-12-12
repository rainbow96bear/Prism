package LightInfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	Login "prism_back/kakao/login"
	Session "prism_back/session"
)

func GetUserInfo_from_Session(res http.ResponseWriter, req *http.Request) {
	session, err := Session.Store.Get(req, "user_login")
	if err != nil {
		fmt.Println("세션을 가져오는데 문제 발생:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	User_id, ok := session.Values["User_ID"].(string)
	if !ok {
		http.Error(res, "User_ID not found in session", http.StatusInternalServerError)
		return
	}
	Profile_img, ok := session.Values["User_ProfileImg"].(string)
	if !ok {
		http.Error(res, "User_ProfileImg not found in session", http.StatusInternalServerError)
		return
	}

	responseData := Login.User{
		User_id:         User_id,
		Profile_img: Profile_img,
	}
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSON을 응답으로 전송
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}