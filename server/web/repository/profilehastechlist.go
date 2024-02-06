package repository

import (
	"database/sql"
	"prism/web/dto"
	"prism/web/pkg/models"
)

type ProfileHasTechListRepository struct {
	models.Tech
}

// repository - id에 해당하는 Tech 이름과 레벨을 모두 얻어오기
func (p *ProfileHasTechListRepository) GetUserTechList(tx *sql.Tx, id string) (dto.TechList, error) {
	var userTechList dto.TechList

	query := `SELECT tech_list.Tech_name, profile_has_tech_list.level FROM profile_has_tech_list JOIN tech_list ON profile_has_tech_list.tech_list_Id = tech_list.Id WHERE profile_Id = ?`
	
	rows, err := tx.Query(query, id)
	if err != nil {
		return dto.TechList{}, err
	}

	for rows.Next() {
		var userTech dto.UserTech
		if err := rows.Scan(&userTech.Name, &userTech.Level); err != nil {
			return dto.TechList{}, err
		}
		userTechList.UserTechList = append(userTechList.UserTechList, userTech)
	}
	
	return userTechList, nil
}



// User의 기술스택 전부 제거
func (p *ProfileHasTechListRepository)Delete(tx *sql.Tx, id string) (error) {
	query := "DELETE FROM profile_has_tech_list WHERE profile_id = ?"
	_, err := tx.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProfileHasTechListRepository)Create(tx *sql.Tx, id, tech_name string, level int) (error) {
	tech_id, err := getTechId(tx, tech_name)
	if err != nil {
		return err
	}
	query := "INSERT INTO profile_has_tech_list(profile_Id, tech_list_Id, level) VALUE(?, ?, ?)"
	_, err = tx.Exec(query, id, tech_id, level)
	if err != nil {
		return err
	}
	return nil
}

// techlist의 id에 해당하는 기술 스택의 개수 출력
func getTotalCount(tx *sql.Tx, tech models.Tech) (count int, err error){
	query := "SELECT COUNT(*) FROM profile_has_tech_list WHERE tech_list_Id = ?"
	err = tx.QueryRow(query, tech.Id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}