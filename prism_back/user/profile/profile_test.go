package Profile

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
    err := godotenv.Load("./../../.env")
    if err != nil {
        fmt.Println("Failed to open test DB:", err)
        return
    }
    testDB, err = sql.Open("mysql", "root:"+os.Getenv("MYSQL_PW")+"@tcp(localhost:3306)/test_prism")

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
// func TestGetUserProfile(t *testing.T){
	
// }

func Test_C_profile(t *testing.T) {
    setupTestDB()
	defer teardownTestDB()
    err := godotenv.Load("./../../.env")
    if err != nil {
        fmt.Println("env 오류")
    }
	user_id := "3206615698"

    result, err := C_profile(user_id, testDB)
    fmt.Println("result", result)
    if err != nil {
        t.Fatal(err)
    }
    expected_User_id := "3206615698"
    if result.User_id != expected_User_id {
        t.Errorf("예상한 프로필 이미지 URL과 다릅니다. 기대값: %s, 결과값: %s", expected_User_id, result.User_id)
    }
}