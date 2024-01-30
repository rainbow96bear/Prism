package repository

import (
	"database/sql"
	"prism_back/pkg/models"
)

type ProfileRepository struct {
	models.Profile
}

// repository - profile 테이블에 user id와 profile id를 Insert
func (p *ProfileRepository)Create(tx *sql.Tx, id string) (string, error) {
	query := "INSERT INTO profile(id, user_info_User_id) VALUE(?, ?)"
	_, err := tx.Exec(query, id, id)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return id, nil
}

// repository - profile 테이블에서 id에 해당하는 Profile 정보 얻기
func (p *ProfileRepository)Read(tx *sql.Tx, id string) (models.Profile, error) {
	var profile models.Profile
	query := "SELECT `id`, `One_line_introduce`, `user_info_User_id` FROM profile WHERE id = ?"
	err := tx.QueryRow(query, id).Scan(&profile.Id, &profile.One_line_introduce, &profile.User_id)
	if err != nil {
		tx.Rollback()
		return models.Profile{}, err
	}
	return profile, err
}

// repository - profile 테이블에 한 줄 소개 수정
func (p *ProfileRepository)Update(tx *sql.Tx, id, one_line_introduce string) (models.Profile, error) {
	query := "UPDATE profile SET One_line_introduce WHERE id = ?"
	_, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return models.Profile{}, err
	}
	updatedProfile := models.Profile{Id: id, One_line_introduce: one_line_introduce, User_id: p.User_id}
	return updatedProfile, nil
}


// repository - profile 테이블에 id에 해당하는 row 삭제
func (p *ProfileRepository)Delete(tx *sql.Tx, id string) (string, error) {
	query := "DELETE FROM profile WHERE id = ?"
	_, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return id, err
	}
	return id, nil
}