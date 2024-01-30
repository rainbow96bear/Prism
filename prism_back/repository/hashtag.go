package repository

import (
	"database/sql"
	"prism_back/pkg/models"
)

type HashtagRepository struct {
	models.Hashtag
}

func (h *HashtagRepository)Create(tx *sql.Tx, id, hashtag string) (string, error) {
	_, err := tx.Exec("INSERT INTO hashtag_list(profile_Id, hashtag) VALUES(?, ?)", id, hashtag)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return hashtag, nil
}

// hashtag_list 테이블에서 id에 해당하는 hashtag list 얻기
func (h *HashtagRepository)Read(tx *sql.Tx, id string) (HashtagList []string, err error) {
	query := "SELECT hashtag FROM hashtag_list WHERE profile_Id = ?"
	rows, err := tx.Query(query, id)
	if err != nil {
		tx.Rollback()
		return HashtagList, err
	}
	defer rows.Close()

	// hashtag를 배열로 저장
	for rows.Next() {
		var hashtag string
		if err := rows.Scan(&hashtag); err != nil {
			tx.Rollback()
			return HashtagList, err
		}
		HashtagList = append(HashtagList,hashtag)
	}
	return HashtagList, nil
}

// repository - hashtaglist 테이블에 hashtag 수정
func (h *HashtagRepository)Update(tx *sql.Tx, id string, hashtagArray []string) (hashtagList []models.Hashtag, err error){
	// Delete 메서드로 id에 해당하는 hashtag 삭제
	err = h.Delete(tx, id)
	if err != nil {
		return []models.Hashtag{}, err
	}

	//profile_id 값이 가져온 profileID인 row를 hashtag_list 테이블에 추가
	for _, hashtag := range hashtagArray {
		_, err := tx.Exec("INSERT INTO hashtag_list(profile_Id, hashtag) VALUES(?, ?)", id, hashtag)
		if err != nil {
			tx.Rollback()
			return []models.Hashtag{}, err
		}
	}
	return hashtagList, nil
}

func (h *HashtagRepository)Delete(tx *sql.Tx, id string) (err error){
	query := "DELETE FROM hashtag_list WHERE profile_Id = ?"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}