package repository

import (
	"database/sql"
	"prism_back/pkg/models"
)

type UserInfoReopository struct {
	models.UserInfo
}

func (u *UserInfoReopository)Create(tx *sql.Tx, userinfo models.UserInfo) (err error){
	var (
		query string
	)

	if userinfo.Id != "" {
		query = "INSERT INTO user_info(User_id, Nickname) VALUE (?, ?)"
		_, err = tx.Exec(query, userinfo.Id, userinfo.NickName)
	}else {
		query = "INSERT INTO user_info(Nickname) VALUE (?)"
		_, err = tx.Exec(query, userinfo.NickName)
	}

	if err != nil {
		return err
	}

	return nil
}

// user_info 테이블에서 id에 해당하는 user_id, nickname 정보 얻기
func (u *UserInfoReopository)Read(tx *sql.Tx, id string) (models.UserInfo, error){
	var userinfo models.UserInfo
	query := "SELECT `User_id`, `Nickname` FROM user_info WHERE user_id = ?"
	err := tx.QueryRow(query, id).Scan(&userinfo.Id, &userinfo.NickName)
	if err != nil {
		tx.Rollback()
		return models.UserInfo{}, err
	}
	return userinfo, nil
}

func (u *UserInfoReopository)Update(tx *sql.Tx, userinfo models.UserInfo) (models.UserInfo, error){
	query := "UPDATE user_info SET Nickname = ? WHERE user_id = ?"
	_, err := tx.Exec(query, userinfo.NickName, userinfo.Id)
	if err != nil {
		tx.Rollback()
		return models.UserInfo{}, err
	}
	return userinfo, nil
}

// DELETE는 CASECADE 적용 후 추가 작성
// func (u *UserInfo) Delete(tx *sql.Tx) (error){
// 	return nil
// }