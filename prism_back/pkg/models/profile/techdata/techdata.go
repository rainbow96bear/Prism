package techdata

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"prism_back/internal/Database/mysql"

	"github.com/gorilla/mux"
)

type Tech struct {
	Tech_name string `json:"tech_name"`
	Level     int    `json:"level,omitempty"`
}

type UserTechList struct {
	Tech_list []Tech `json:"tech_list,omitempty"`
}
type TechList []Tech

func (t *Tech) GetTechList(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	techList, err := getTechList(id)
	jsonResponse, err := json.Marshal(techList)
	if err != nil {
		log.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}


func getTechList(id string) (UserTechList, error) {
	tx, err := mysql.DB.Begin()
	if err != nil {
		return UserTechList{}, err
	}

	query := `SELECT tech_list.Tech_name, profile_has_tech_list.level FROM profile_has_tech_list JOIN tech_list ON profile_has_tech_list.tech_list_Id = tech_list.Id WHERE profile_Id = ? `

	rows, err := tx.Query(query, id)
	if err != nil {
		log.Println(err)
		return UserTechList{}, err
	}
	defer rows.Close()

	var tech_list UserTechList
	// tech를 배열로 저장
	for rows.Next() {
		var tech Tech
		if err := rows.Scan(&tech.Tech_name, &tech.Level); err != nil {
			log.Println(err)
			return UserTechList{}, err
		}
		tech_list.Tech_list = append(tech_list.Tech_list, tech)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return UserTechList{}, err
	}

	return tech_list, nil
}

func (t *Tech) SetTechList(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	// JSON 데이터 언마샬링
	var data struct {
		UserTechList []Tech `json:"userTechList"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(res, "Failed to unmarshal JSON", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	id := vars["id"]
	err = setTechList(id, data.UserTechList)
	if err != nil {
		log.Println(err)
		http.Error(res, "Failed to set tech list", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Tech list updated successfully"))
}

func setTechList(id string, data []Tech) error {
	tx, err := mysql.DB.Begin()
	if err != nil {
		return err
	}

	// 에러 발생 시 롤백
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	

	// Hashtags 테이블에서 해당 ProfileID에 해당하는 레코드 삭제
	_, err = tx.Exec("DELETE FROM profile_has_tech_list WHERE profile_id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 새로운 데이터 삽입
	for _, value := range data {
		// Tech_name을 기반으로 해당하는 Id를 서브쿼리를 이용하여 가져옴
		row := tx.QueryRow("SELECT Id FROM tech_list WHERE Tech_name = ?", value.Tech_name)
		var techListId int
		err := row.Scan(&techListId)
		if err != nil {
			tx.Rollback()
			return err
		}

		// profile_has_tech_list 테이블에 데이터 삽입
		_, err = tx.Exec("INSERT INTO profile_has_tech_list(profile_Id, tech_list_Id, level) VALUES(?, ?, ?)", id, techListId, value.Level)
		if err != nil {
			tx.Rollback()
			return err
		}
		
		// Count 값을 업데이트
		err = updateTechListCount(tx, techListId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 트랜잭션 커밋
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Count 값을 업데이트하는 함수
func updateTechListCount(tx *sql.Tx, techListId int) error {
    // 해당 tech_list_Id의 개수를 계산
    row := tx.QueryRow("SELECT COUNT(*) FROM profile_has_tech_list WHERE tech_list_Id = ?", techListId)
    var count int
    err := row.Scan(&count)
    if err != nil {
        return err
    }

    // tech_list 테이블에서 Count 값을 업데이트
    _, err = tx.Exec("UPDATE tech_list SET Count = ? WHERE Id = ?", count, techListId)
    return err
}