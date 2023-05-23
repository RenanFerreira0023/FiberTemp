package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err := godotenv.Load(".env") // Carrega as variáveis do arquivo .env
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}
	nameDB := os.Getenv("NAME_DB")
	hostDB := os.Getenv("HOST_DB")
	db, err := sql.Open("mysql", "root:1234@tcp("+hostDB+")/"+nameDB+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
