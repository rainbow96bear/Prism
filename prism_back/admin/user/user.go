package User

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	Mysql "prism_back/DataBase/MySQL"
	Session "prism_back/session"

	"golang.org/x/crypto/bcrypt"
)

type RequestData struct {
	Password string `json:"password"`
}

func AdminCheck(res http.ResponseWriter, req *http.Request) {
	session, err := Session.Store.Get(req, "user_login")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	User_id := session.Values["User_ID"]

	query := "SELECT Admin_id FROM admin_user WHERE Admin_id = ?"
	rows, err := Mysql.DB.Query(query, User_id)
	if err != nil {
		fmt.Println("에러 발생:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	 }
	defer rows.Close()
	// 결과가 있는지 확인
	var isAdmin bool
	if rows.Next() {
		isAdmin = true
	} else {
		isAdmin = false
	}

	session, err = Session.Store.Get(req, "admin_login")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	Admin_id := session.Values["admin_id"]
	Admin_rank := session.Values["admin_rank"]
	var correctAdmin bool
	if User_id != Admin_id {
		correctAdmin = false
		http.SetCookie(res, &http.Cookie{
			Name:   "admin_login",
			Value:  "",
			MaxAge: -1,
			Domain: os.Getenv("COOKIE_DOMAIN"),
			Path:   "/",
		})
	}else {
		correctAdmin = true
	}
	// 결과를 JSON으로 응답
	response := map[string]interface{}{
		"isAdmin": isAdmin,
		"correctAdmin" : correctAdmin,
		"admin_info" : map[string]interface{}{
			"id" : Admin_id,
			"rank" : Admin_rank,
		},
    }

    jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("JSON 마샬 에러:", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

    res.Header().Set("Content-Type", "application/json")
    res.Write(jsonResponse)
}


func AdminLogin(res http.ResponseWriter, req *http.Request) {
	
	session, err := Session.Store.Get(req, "user_login")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	User_id := session.Values["User_ID"].(string)
	// 요청에서 전송된 비밀번호
	decoder := json.NewDecoder(req.Body)

	var requestData RequestData  // MyStruct는 적절한 구조체로 대체되어야 합니다.

	// JSON 디코딩
	err = decoder.Decode(&requestData)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return
	}
	// DB에서 저장된 암호화된 비밀번호 조회
	var Password string
	var Rank int
	err = Mysql.DB.QueryRow("SELECT `Rank`, `Password` FROM admin_user WHERE Admin_id = ?", User_id).Scan(&Rank, &Password)
	if err != nil {
		fmt.Println("error : ", err)
		http.Error(res, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 비밀번호 비교
	result := ComparePassword(Password, requestData.Password)
	var response map[string]interface{}
	if result == false {
		response = map[string]interface{}{
			"isAdmin": result,
			"admin_info" : map[string]interface{}{
				"id" : "",
				"rank" : 0,
			},
		}
	}else {
		response = map[string]interface{}{
			"isAdmin": result,
			"admin_info" : map[string]interface{}{
				"id" : User_id,
				"rank" : Rank,
			},
		}
		err = CreateSession(res, req, User_id, Rank) 
		if err != nil {
			fmt.Println("session 문제 : ", err)
		}
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}

func AdminLogout(res http.ResponseWriter, req *http.Request){
	http.SetCookie(res, &http.Cookie{
		Name:   "admin_login",
		Value:  "",
		MaxAge: -1,
		Domain: os.Getenv("COOKIE_DOMAIN"),
		Path:   "/",
	})
}

func CreateSession(res http.ResponseWriter, req *http.Request, id string, rank int) error {
	session, err := Session.Store.Get(req, "admin_login")
	if err != nil {
		fmt.Println("세션을 가져오는데 문제 발생:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return err
	}
	session.Values["admin_id"] = id
	session.Values["admin_rank"] = rank
	err = session.Save(req, res)

	if err != nil {
		fmt.Println("세션을 저장하는데 문제 발생:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func sendJSONResponse(res http.ResponseWriter, data interface{}) {
    res.Header().Set("Content-Type", "application/json")
    res.WriteHeader(http.StatusOK)

    // 인터페이스를 JSON으로 변환하여 응답 보내기
    if err := json.NewEncoder(res).Encode(data); err != nil {
        http.Error(res, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func HashPassword(password string) (string, error) {
	// 비밀번호 해싱
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// 해싱된 비밀번호를 문자열로 반환
	return string(hashedPassword), nil
}

func ComparePassword(db_password, input_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(db_password), []byte(input_password))
	if err != nil {
		return false
	}
	return true
}