package hashtag

import (
	"database/sql"
)

type Hashtag struct {
	Id 		string
	Hashtag string
	Create 	func(tx *sql.Tx, id, hashtag string)
	Read 	func(tx *sql.Tx, id string)
	Update 	func(tx *sql.Tx, id string, hashtagArray []string)
	Delete 	func(tx *sql.Tx, id string)
}

func Create(tx *sql.Tx, id, hashtag string) (string, error) {
	_, err := tx.Exec("INSERT INTO hashtag_list(profile_Id, hashtag) VALUES(?, ?)", id, hashtag)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return hashtag, nil
}

func Read(tx *sql.Tx, id string) (HashtagList []string, err error) {
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

func Update(tx *sql.Tx, id string, hashtagArray []string) (hashtagList []Hashtag, err error){
	// Delete 메서드로 id에 해당하는 hashtag 삭제
	err = Delete(tx, id)
	if err != nil {
		return []Hashtag{}, err
	}

	//profile_id 값이 가져온 profileID인 row를 hashtag_list 테이블에 추가
	for _, hashtag := range hashtagArray {
		_, err := tx.Exec("INSERT INTO hashtag_list(profile_Id, hashtag) VALUES(?, ?)", id, hashtag)
		if err != nil {
			tx.Rollback()
			return []Hashtag{}, err
		}
	}
	return hashtagList, nil
}

func Delete(tx *sql.Tx, id string) (err error){
	query := "DELETE FROM hashtag_list WHERE profile_Id = ?"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}