package Mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load(".env")
var DB *sql.DB

func SetupDB() {
	DB, err = sql.Open("mysql", "root:"+os.Getenv("MYSQL_PW")+"@tcp(localhost:3306)/prism")
	if err != nil {
		fmt.Println("Failed to open DB")
	}
}