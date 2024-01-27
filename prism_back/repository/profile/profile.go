package profile

import (
	"database/sql"
)

type Profile struct {
	Id 				   string
	One_line_introduce string
	User_id  		   string
	Create 		func(tx *sql.Tx, id string)
	Read 		func(tx *sql.Tx, id string)
	Update 		func(tx *sql.Tx, id, one_line_introduce string)
	Delete 		func(tx *sql.Tx, id string)
}

func Create(tx *sql.Tx, id string) (string, error) {
	query := "INSERT INTO profile(id, user_info_User_id) VALUE(?, ?)"
	_, err := tx.Exec(query, id, id)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return id, nil
}

func Read(tx *sql.Tx, id string) (Profile, error) {
	var profile Profile
	query := "SELECT `id`, `One_line_introduce`, `user_info_User_id` FROM profile WHERE id = ?"
	err := tx.QueryRow(query, id).Scan(&profile.Id, &profile.One_line_introduce, &profile.User_id)
	if err != nil {
		tx.Rollback()
		return Profile{}, err
	}
	return profile, err
}

func Update(tx *sql.Tx, id, one_line_introduce string) (Profile, error) {
	query := "UPDATE profile SET One_line_introduce WHERE id = ?"
	_, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return Profile{}, err
	}
	return Profile{Id : id, One_line_introduce: one_line_introduce, User_id : id}, nil
}

func Delete(tx *sql.Tx, id string) (string, error) {
	query := "DELETE FROM profile WHERE id = ?"
	_, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return id, err
	}
	return id, nil
}