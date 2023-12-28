package Tech

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

func Test_R_TechList(t *testing.T){
	setupTestDB()
	defer teardownTestDB()
   
    Read_Result, err := R_TechList(testDB)
	if err != nil {
		t.Fatalf("DB 읽기 실패 : %e", err)
	}
	expected := []TechData{{"00001", "golang", 0}}
	for index, result := range expected {
		if Read_Result[index].TechCode != result.TechCode {
			t.Fatalf("기대값 %v, 결과값 %v", Read_Result[index].TechCode, result.TechCode)
		}
		if Read_Result[index].TechName != result.TechName {
			t.Fatalf("기대값 %v, 결과값 %v", Read_Result[index].TechName, result.TechName)
		}
		if Read_Result[index].Count != result.Count {
			t.Fatalf("기대값 %v, 결과값 %v", Read_Result[index].Count, result.Count)
		}
	}
}

func Test_U_tech_list(t *testing.T){
	setupTestDB()
	defer teardownTestDB()
	previousData := TechData{"00003", "C++", 0}
	newData := TechData{"00003", "C#", 0}
	result, err := U_tech_list(previousData, newData, testDB)
	if result.TechCode == previousData.TechCode {
		fmt.Println(err)
	}
	if err != nil {
		t.Fatalf("수정 오류 : %e", err)
	}
	if newData.TechCode != result.TechCode {
		t.Fatalf("기대값 %v, 결과값 %v", newData.TechCode, result.TechCode)
	}
	if newData.TechName != result.TechName {
		t.Fatalf("기대값 %v, 결과값 %v", newData.TechName, result.TechName)
	}
}