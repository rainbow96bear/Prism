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

func (t *TechRepository)Update(tx *sql.Tx, preData, newData models.Tech) (models.Tech, error) {
	query := "UPDATE tech_list SET tech_name = ? WHERE tech_name = ?"
	_, err := tx.Exec(query, preData.TechName)
	if err != nil {
		tx.Rollback()
		return preData, err
	}
	return newData, nil
}

func (t *TechRepository)UpdateCount(tx *sql.Tx, tech models.Tech) (error){
	count, err := t.GetTotalCount(tx, tech)
	if err != nil {
		tx.Rollback()
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

// 이건 profile_has_tech_list 로 이동할 예정
func (t *TechRepository)GetTotalCount(tx *sql.Tx, tech models.Tech) (count int, err error){
	query := "SELECT COUNT(*) FROM profile_has_tech_list WHERE tech_list_Id = ?"
	err = tx.QueryRow(query, tech.Id).Scan(&count)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return count, nil
}
// func (t *Tech) Delete(tx *sql.Tx) {
// 	query := ""
// }