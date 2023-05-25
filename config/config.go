package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_LOCAL")
	fmt.Println("dbPort    ", dbPort)
	fmt.Println("dbHost    ", dbHost)
	db, err := sql.Open("mysql", "root:1234@tcp("+dbHost+":"+dbPort+")/rds_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
