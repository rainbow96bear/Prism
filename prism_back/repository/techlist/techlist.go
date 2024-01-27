package techlist

import (
	"database/sql"
)

type Tech struct {
	Id       string
	TechName string
	Count    int   
	Create 			func(tx *sql.Tx, name string)
	ReadAll 		func(tx *sql.Tx) ([]Tech, error) 
	Read 			func(tx *sql.Tx, tech_name string)
	Update 			func(tx *sql.Tx, preData, newData Tech)
	UpdateCount 	func(tx *sql.Tx, tech Tech)
	GetTotalCount 	func(tx *sql.Tx, tech Tech)
}

func Create(tx *sql.Tx, name string) (error) {
	query := "INSERT INTO tech_list(Tech_name) VALUE (?)"
	_, err := tx.Exec(query, name)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func ReadAll(tx *sql.Tx) ([]Tech, error) {
	var techlist []Tech
	query := "SELECT `tech_name`, `count` FROM tech_list"
	rows, err := tx.Query(query)
	if err != nil {
		tx.Rollback()
		return []Tech{}, err
	}

	for rows.Next() {
		var tech Tech
		if err := rows.Scan(&tech.TechName); err != nil {
			tx.Rollback()
			return []Tech{}, err
		}
		techlist = append(techlist, tech)
	}
	return techlist, nil
}

func Read(tx *sql.Tx, tech_name string) (Tech, error) {
	var tech Tech
	query := "SELECT `tech_name`, `count` FROM tech_list WHERE tech_code = ? "
	err := tx.QueryRow(query, tech_name).Scan(&tech.TechName, &tech.Count)
	if err != nil {
		tx.Rollback()
		return Tech{}, err
	}
	return tech, nil
}

func Update(tx *sql.Tx, preData, newData Tech) (Tech, error) {
	query := "UPDATE tech_list SET tech_name = ? WHERE tech_name = ?"
	_, err := tx.Exec(query, preData.TechName)
	if err != nil {
		tx.Rollback()
		return preData, err
	}
	return newData, nil
}

func UpdateCount(tx *sql.Tx, tech Tech) (error){
	count, err := GetTotalCount(tx, tech)
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
func GetTotalCount(tx *sql.Tx, tech Tech) (count int, err error){
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