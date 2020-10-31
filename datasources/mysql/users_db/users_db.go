package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlUsername = "mysql_username"
	mysqlPassword = "mysql_password"
	mysqlDatabase = "mysql_database"
	mysqlHost     = "mysql_host"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysqlUsername)
	password = os.Getenv(mysqlPassword)
	database = os.Getenv(mysqlDatabase)
	host     = os.Getenv(mysqlHost)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		username, password, host, database,
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Printf("database successfully configured")
}
