package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file  ", err.Error())
	}

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	dbPort := os.Getenv("MYSQL_PORT")
	dbHost := os.Getenv("MYSQL_HOST")

	fmt.Println("    |   dbHost    ", dbHost)
	fmt.Println("    |   dbPassword    ", dbPassword)
	fmt.Println("    |   dbName    ", dbName)

	fmt.Println("    |   dbPort    ", dbPort)
	fmt.Println("    |   dbUser    ", dbUser)

	db, err := sql.Open("mysql", ""+dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//	CreateTables(db)
	return db
}

// isso aqui precisa ir para o repository
func CreateTables(db *sql.DB) {

	createTableAgent := "CREATE TABLE IF NOT EXISTS users_agent (  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,  first_name VARCHAR(200),  second_name VARCHAR(200),  email VARCHAR(150),  password_agent VARCHAR(64), dt_create_account DATETIME,  dt_expired_account DATETIME,  account_valid BOOLEAN,  quantity_alerts INT(100),  quantity_account_copy INT(100))"
	_, errTbAgent := db.Exec(createTableAgent)
	if errTbAgent != nil {
		panic(errTbAgent)
	}

	createTableReceptor := "CREATE TABLE users_receptor (  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,  first_name VARCHAR(200),  second_name VARCHAR(200),  email VARCHAR(150),  dt_create_account DATETIME,  dt_expired_account DATETIME,  agent_id INT,    FOREIGN KEY (agent_id) REFERENCES users_agent(id)  )"
	_, errTbReceptor := db.Exec(createTableReceptor)
	if errTbReceptor != nil {
		panic(errTbReceptor)
	}

	createTableChannel := "CREATE TABLE IF NOT EXISTS channels (  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,  users_agent_id INT,  channel_name VARCHAR(200),  dt_create_channel DATETIME,  FOREIGN KEY (users_agent_id) REFERENCES users_agent(id))"
	_, errTbChannel := db.Exec(createTableChannel)
	if errTbChannel != nil {
		panic(errTbChannel)
	}

	createTablePermission := "CREATE TABLE IF NOT EXISTS permission (  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,  user_receptor_id INT,  channel_id INT,  FOREIGN KEY (user_receptor_id) REFERENCES users_receptor(id),  FOREIGN KEY (channel_id) REFERENCES channels(id))"
	_, errTbPermission := db.Exec(createTablePermission)
	if errTbPermission != nil {
		panic(errTbPermission)
	}

	createTableAllCopy := "CREATE TABLE IF NOT EXISTS all_copy (  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,  symbol VARCHAR(30),   action_type VARCHAR(255),  ticket BIGINT(30),  lot FLOAT(5),  target_pedding FLOAT(25),  takeprofit FLOAT(25),  stoploss FLOAT(25),  dt_send_order DATETIME,  user_agent_id INT,  channel_id INT,  FOREIGN KEY (user_agent_id) REFERENCES users_agent(id),  FOREIGN KEY (channel_id) REFERENCES channels(id))"
	_, errTbAllCopy := db.Exec(createTableAllCopy)
	if errTbAllCopy != nil {
		panic(errTbAllCopy)
	}

	createTableReqCopy := "CREATE TABLE IF NOT EXISTS req_copy (  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,  dt_send_copy DATETIME,  channel_id INT,  users_receptor_id INT,  all_copy_id INT,  FOREIGN KEY (channel_id) REFERENCES channels(id),  FOREIGN KEY (users_receptor_id) REFERENCES users_receptor(id),  FOREIGN KEY (all_copy_id) REFERENCES all_copy(id))"
	_, errTbReqCopy := db.Exec(createTableReqCopy)
	if errTbReqCopy != nil {
		panic(errTbReqCopy)
	}

	dtCreateAccount := time.Now()
	createInsertAdminAgent := ("INSERT INTO  users_agent (first_name, second_name, email, dt_create_account, dt_expired_account, account_valid, quantity_alerts, quantity_account_copy) VALUES ('admin', 'adminSubName', 'admin@rdstrader.com', ?, '2223-06-16 10:00:00', '1', '99', '100')")
	_, errInsertAdminAgent := db.Exec(createInsertAdminAgent, dtCreateAccount)
	if errInsertAdminAgent != nil {
		fmt.Println("errInsertAdminAgent    ", errInsertAdminAgent.Error())
		panic(errInsertAdminAgent)
	}

}
