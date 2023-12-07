package KakaoLogin

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var testDB *sql.DB


func setupTestDB() {
    var err error
    testDB, err = sql.Open("mysql", "root:0000@tcp(localhost:3306)/test_prism")
    if err != nil {
        fmt.Println("Failed to open test DB:", err)
        return
    }

    // 데이터베이스에 연결되었는지 확인
    if err := testDB.Ping(); err != nil {
        fmt.Println("Failed to connect to test DB:", err)
        return
    }

    fmt.Println("Connected to test DB")
}


func teardownTestDB(){
	if testDB != nil {
		testDB.Close()
	}
}

// func TestR_UserInfo(t * testing.T) {
// 	setupTestDB()
// 	defer teardownTestDB()

// 	userToCreate := User{
//         ID:          "5555",
//         NickName:    "무지개곰",
//         ProfileImg:  "http://example.com/profile.jpg",
//     }

// 	result, err := R_UserInfo(userToCreate, testDB)
//     if err != nil {
//         t.Fatalf("Failed to create user: %v", err)
//     }
//     if result == nil {
//         t.Fatalf("결과가 nil")
//     }

//     fmt.Println("Read(UserInfo) 테스트 완료")
// }

// func TestIsSavedID(t *testing.T){
//     setupTestDB()
// 	defer teardownTestDB()

//     userToCreate := User{
//         ID:          "444",
//         NickName:    "무지개곰",
//         ProfileImg:  "http://example.com/profile.jpg",
//     }
//     isSavedID, err := IsSavedID(userToCreate, testDB)
//     if err != nil {
//         return
//     }
//     if !isSavedID {
//         err := C_UserInfo(userToCreate, testDB)
//         if err != nil {
//             t.Fatal("생성 실패")
//         }
//     }
// }

func TestGetUserInfo_and_Save(t *testing.T) {
    setupTestDB()
	defer teardownTestDB()
    err := godotenv.Load("./../.env")
    if err != nil {
        fmt.Println("env 오류")
    }
    TEMP_TOKEN := os.Getenv("TEMP_TOKEN")

    user, err := GetUserInfo(TEMP_TOKEN)
   
    Read_Result, err := R_UserInfo(user,testDB)
    fmt.Println("R_UserInfo 결과 : ", Read_Result)
    isSavedID, err := IsSavedID(user, testDB)
    if err != nil {
        return
    }
    if !isSavedID {
        err := C_UserInfo(user, testDB)
        if err != nil {
            t.Fatal("생성 실패")
        }
    }
}