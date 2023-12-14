package Profile

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
)

var err = godotenv.Load(".env")

func C_profile(user_id string, db *sql.DB) (Profile, error) {
	query := "INSERT INTO profile (One_line_introduce, user_info_User_id) VALUES (?, ?)"
	_, err := db.Exec(query, "", user_id)

	if err != nil {
		return Profile{}, fmt.Errorf("프로필 생성 실패: %v", err)
	}
	return Profile{User_id: user_id}, nil
}