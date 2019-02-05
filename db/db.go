package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/task")
	if err != nil {
		panic(err.Error)
	}
	return db, err
}
