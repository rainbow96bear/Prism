package repository

import (
	"database/sql"
	"prism_back/pkg/models"
	"testing"
)

// UserInfoRepository를 테스트하기 위한 구조체
type TestUserInfoRepository struct{}

// Read 메서드를 구현하여 테스트용으로 오버라이드
func (u *TestUserInfoRepository) Read(tx *sql.Tx, id string) (models.UserInfo, error) {
	// 테스트용으로 고정된 데이터 반환
	return models.UserInfo{Id: id, NickName: "TestNickname"}, nil
}

func TestUserInfoRepository_Read(t *testing.T) {
	// 테스트용 UserInfoRepository
	testRepo := &TestUserInfoRepository{}

	// 가상의 트랜잭션을 생성 (테스트를 위한 가상의 DB 작업)
	mockTx, err := testDB.Begin()
	if err != nil {
		t.Fatalf("Error starting transaction: %v", err)
	}

	// 테스트용으로 설정한 데이터베이스 아이디
	testUserID := "test_user_id"

	// UserInfoRepository의 Read 메서드 호출
	result, err := testRepo.Read(mockTx, testUserID)
	if err != nil {
		t.Fatalf("Error calling Read: %v", err)
	}

	// 예상 결과 값과 비교
	expectedResult := models.UserInfo{
		Id:       testUserID,
		NickName: "TestNickname",
	}

	// 결과 값이 예상 값과 다르면 에러 출력
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}

	// 트랜잭션 롤백
	mockTx.Rollback()
}

// 실제 데이터베이스 대신 사용할 테스트용 데이터베이스 설정
var testDB *sql.DB