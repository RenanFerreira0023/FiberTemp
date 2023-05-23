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
	portDB := os.Getenv("PORT_DB")
	hostDB := os.Getenv("HOST_DB") // 3306

	userDB := os.Getenv("USER_DB")      // root
	passwordDB := os.Getenv("PASSW_DB") // 1234

	db, err := sql.Open("mysql", userDB+":"+passwordDB+"@tcp("+hostDB+":"+portDB+")/"+nameDB+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
