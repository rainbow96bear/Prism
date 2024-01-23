package techdata

import (
	"encoding/json"
	"log"
	"net/http"
	"prism_back/internal/Database/mysql"

	"github.com/gorilla/mux"
)

type Tech struct {
	Tech_name string `json:"tech_name"`
	Level     int    `json:"level,omitempty"`
}

type TechList []Tech

func (t *Tech) GetPersonalData(res http.ResponseWriter, req *http.Request) {
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


func getTechList(id string) (TechList, error) {
	tx, err := mysql.DB.Begin()
	if err != nil {
		return TechList{}, err
	}

	query := `SELECT profile_has_tech_list.level, tech_list.Tech_name FROM profile_has_tech_list JOIN tech_list ON profile_has_tech_list.tech_list_Id = tech_list.Id WHERE profile_Id = ? `

	rows, err := tx.Query(query, id)
	if err != nil {
		log.Println(err)
		return TechList{}, err
	}
	defer rows.Close()

	var tech_list TechList
	// tech를 배열로 저장
	for rows.Next() {
		var tech Tech
		if err := rows.Scan(&tech); err != nil {
			log.Println(err)
			return TechList{}, err
		}
		tech_list = append(tech_list, tech)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return TechList{}, err
	}

	return tech_list, nil
}