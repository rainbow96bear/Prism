package userinfo

import (
	"database/sql"
)

type UserInfo struct {
	Id       string
	NickName string
	Create 	 func(tx *sql.Tx, userinf UserInfo)
	Read 	 func(tx *sql.Tx, id string)
	Update 	 func(tx *sql.Tx, userinfo UserInfo)
}

func Create(tx *sql.Tx, userinfo UserInfo) (err error){
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

func Read(tx *sql.Tx, id string) (UserInfo, error){
	var userinfo UserInfo
	query := "SELECT `User_id`, `Nickname` FROM user_info WHERE user_id = ?"
	err := tx.QueryRow(query, id).Scan(&userinfo.Id, &userinfo.NickName)
	if err != nil {
		tx.Rollback()
		return UserInfo{}, err
	}
	return userinfo, nil
}

func Update(tx *sql.Tx, userinfo UserInfo) (UserInfo, error){
	query := "UPDATE user_info SET Nickname = ? WHERE user_id = ?"
	_, err := tx.Exec(query, userinfo.NickName, userinfo.Id)
	if err != nil {
		tx.Rollback()
		return UserInfo{}, err
	}
	return userinfo, nil
}

// DELETE는 CASECADE 적용 후 추가 작성
// func (u *UserInfo) Delete(tx *sql.Tx) (error){
// 	return nil
// }