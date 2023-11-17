package connector

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func GetMysqlConnection() (*sql.DB, error) {
	connString := os.Getenv("AUTH_MYSQL_DSN")

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)

	return db, nil
}
