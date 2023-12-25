package Tech

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	Mysql "prism_back/DataBase/MySQL"

	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버
	"github.com/joho/godotenv"
)

type TechData struct {
	TechCode string `json:"Tech_code"`
	TechName string `json:"Tech_name"`
	Count int `json:"Count"`
}

type PutTechData struct {
	TechData    TechData `json:"editedValues"`
	ExistingData TechData `json:"existingData"`
}

var err = godotenv.Load("./../../.env")

func Get_tech_list(res http.ResponseWriter, req *http.Request) {
	// GetTechList 함수를 통해 데이터베이스에서 기술 목록을 가져옴
	list, err := R_TechList(Mysql.DB)
	if err != nil {
		// 에러가 발생하면 에러 출력 후 HTTP 500 에러 응답
		fmt.Println("DB 조회 실패:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	// JSON으로 직렬화 
	response, err := json.Marshal(list)
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

func Post_tech_list(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var requestData TechData

	// JSON 디코딩
	err := decoder.Decode(&requestData)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := C_tech_list(requestData, Mysql.DB)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}


func Put_tech_list(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var requestData PutTechData

	// JSON 디코딩
	err := decoder.Decode(&requestData)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := U_tech_list(requestData.ExistingData, requestData.TechData, Mysql.DB)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func C_tech_list(data TechData,db *sql.DB) (TechData,error) {
	query:= "SELECT tech_code FROM tech_list WHERE tech_code = ?"
	var tech_code string
	db.QueryRow(query, data.TechCode).Scan(&tech_code)
	if data.TechCode == tech_code {
		return data, fmt.Errorf("이미 사용중인 tech_code")
	}
	query = "INSERT INTO tech_list (tech_code, tech_name, count) VALUES (?, ?, ?)"
	_, err := db.Exec(query, data.TechCode, data.TechName, 0)
	if err != nil {
		return TechData{}, fmt.Errorf("추가 실패: %e", err)
	}
	return data, nil
}

func R_TechList(db *sql.DB) ([]TechData,error) {
	query:= "SELECT `Tech_code`, `Tech_name`, `Count` FROM tech_list ORDER BY Tech_code"
	rows, err := db.Query(query)
	if err != nil {
		return []TechData{}, fmt.Errorf("DB 조회 실패 : %v", err)
	}
	list := []TechData{}

	for rows.Next() {
		var data TechData
		if err := rows.Scan(&data.TechCode, &data.TechName, &data.Count); err != nil {
			return []TechData{}, fmt.Errorf("정보 입력 실패 : %v", err)
		}
		list = append(list, data)
	}
	return list, nil
}

func U_tech_list(preData, newData TechData, db *sql.DB) (TechData,error) {
	if newData.TechCode == "" {
		newData.TechCode = preData.TechCode
	}
	if newData.TechName == "" {
		newData.TechName = preData.TechName
	}
	if newData.TechCode != preData.TechCode{
		query:= "SELECT tech_code FROM tech_list WHERE tech_code = ?"
		var tech_code string
		db.QueryRow(query, preData.TechCode).Scan(&tech_code)
		if newData.TechCode == tech_code {
			return preData, fmt.Errorf("이미 사용중인 tech_code")
		}
	}
	query := "UPDATE tech_list SET tech_code=?, tech_name=? WHERE tech_code=?"
	_, err := db.Exec(query,newData.TechCode, newData.TechName, preData.TechCode)
	if err != nil {
		return TechData{}, fmt.Errorf("수정 실패 : %e", err)
	}
	return newData, nil
}