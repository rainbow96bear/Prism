package repository

import (
	"database/sql"
	"prism/web/errors"
	"prism/web/pkg/models"
)



type AdminUserRepository struct {
}

// repository - admin_user 테이블에서 관리자 정보 조회
func (a *AdminUserRepository) GetAdminUserInfo(tx *sql.Tx, id string) (models.AdminUser, error) {
	var admin_user models.AdminUser
	query := "SELECT `ID`, `Rank`, `Password` FROM admin_user WHERE Id = ?"
	err := tx.QueryRow(query, id).Scan(&admin_user.Id, &admin_user.Rank, &admin_user.Password)
	if err == sql.ErrNoRows {
		return admin_user, errors.IsNotAdminUser
	} else if err != nil {
		return admin_user, err
	}
	return admin_user, nil
}

func (a *AdminUserRepository)CreateAdmin(tx *sql.Tx, id, password string, rank int) (error) {
	query := "INSERT INTO admin_user(`ID`, `Rank`, `password`) VALUES (?, ?, ?)"
	_, err := tx.Exec(query, id, rank, password)
	if err != nil {
		return err
	}
	return nil
}