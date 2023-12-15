package Login

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
    err := godotenv.Load("./../.env")
    MYSQL_PW := os.Getenv("MYSQL_PW")
    testDB, err = sql.Open("mysql", "root:"+MYSQL_PW+"@tcp(localhost:3306)/test_prism")
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

func Test_GetGetUserInfo_from_kakao_and_Save(t *testing.T) {
    setupTestDB()
	defer teardownTestDB()
    err := godotenv.Load("./../.env")
    if err != nil {
        fmt.Println("env 오류")
    }
    TEMP_TOKEN := os.Getenv("TEMP_TOKEN")

    user, err := GetUserInfo_from_kakao(TEMP_TOKEN)
   
    Read_Result, err := R_user_info(user,testDB)
    fmt.Println("R_UserInfo 결과 : ", Read_Result)
    isSavedID, err := IsSavedID(user, testDB)
    if err != nil {
        return
    }
    if !isSavedID {
        _, err := C_user_info(user, testDB)
        if err != nil {
            t.Fatal("생성 실패")
        }
    }
}