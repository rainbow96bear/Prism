package repository

import (
	"database/sql"
	"prism_back/pkg/models"
)

type TechRepository struct {
	models.Tech
}

func (t *TechRepository)Create(tx *sql.Tx, name string) (error) {
	query := "INSERT INTO tech_list(Tech_name) VALUE (?)"
	_, err := tx.Exec(query, name)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// admin 페이지의 tech 관리용 정보 요청
func (t *TechRepository)ReadAll(tx *sql.Tx) ([]models.Tech, error) {
	var techlist []models.Tech
	query := "SELECT `Id`, `tech_name`, `count` FROM tech_list"
	rows, err := tx.Query(query)
	if err != nil {
		tx.Rollback()
		return []models.Tech{}, err
	}

	for rows.Next() {
		var tech models.Tech
		if err := rows.Scan(&tech.Id, &tech.TechName, &tech.Count); err != nil {
			tx.Rollback()
			return []models.Tech{}, err
		}
		techlist = append(techlist, tech)
	}
	return techlist, nil
}


// Admin에서 기술 스택의 이름 변경
func (t *TechRepository)Update(tx *sql.Tx, preData, newData models.Tech) (models.Tech, error) {
	query := "UPDATE tech_list SET tech_name = ? WHERE tech_name = ?"
	_, err := tx.Exec(query, preData.TechName)
	if err != nil {
		tx.Rollback()
		return preData, err
	}
	return newData, nil
}

// 기술 스택을 선택한 사람의 count 수정
func (t *TechRepository)UpdateCount(tx *sql.Tx, tech models.Tech) (error){
	count, err := getTotalCount(tx, tech)
	if err != nil {
		return err
	}
	query := "UPDATE tech_list SET count = ? WHERE tech_name = ?"
	_, err = tx.Exec(query, count, tech.TechName)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// 기술 스택의 이름에 해당하는 기술 스택 id 얻기
func getTechId(tx *sql.Tx, name string) (int, error) {
	var tech_id int
	query := "SELECT id FROM tech_list WHERE tech_name = ?"
	err := tx.QueryRow(query, name).Scan(&tech_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return tech_id, nil
}