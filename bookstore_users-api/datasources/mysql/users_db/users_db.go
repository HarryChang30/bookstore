package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/HarryChang30/bookstore/bookstore_utils-go/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysqlUsersUsername = "DB_USERNAME"
	mysqlUsersPassword = "DB_PASSWORD"
	mysqlUsersHost     = "DB_HOST"
	mysqlUsersSchema   = "DB_SCHEMA"
)

var (
	Client *sql.DB
)

func init() {

	loadErr := godotenv.Load(".env")
	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv(mysqlUsersUsername)
	password := os.Getenv(mysqlUsersPassword)
	host     := os.Getenv(mysqlUsersHost)
	schema   := os.Getenv(mysqlUsersSchema)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	var err error
	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
