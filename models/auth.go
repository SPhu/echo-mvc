package models

import (
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanshirookazaki/echo-mvc/db"
)

type LoginForm struct {
	Username string
	Password string
}

func GetPassword(username string) string {
	db, err := db.Conn()
	rows, err := db.Query("SELECT password FROM users WHERE username = \"" + username + "\"")
	if err != nil {
		panic(err.Error())
	}
	var password string
	defer db.Close()
	defer rows.Close()
	user := User{}
	for rows.Next() {
		if err := rows.Scan(&user.Password); err != nil {
			panic(err.Error())
		}
		password = user.Password
	}
	return password
}

func GetUsername(password string) string {
	db, err := db.Conn()
	rows, err := db.Query("SELECT username FROM users WHERE password = \"" + password + "\"")
	if err != nil {
		panic(err.Error())
	}
	var username string
	defer db.Close()
	defer rows.Close()
	user := User{}
	for rows.Next() {
		if err := rows.Scan(&user.Username); err != nil {
			panic(err.Error())
		}
		username = user.Username
	}
	return username
}

func UserUniqueCheck(username string) string {
	db, err := db.Conn()
	rows, err := db.Query("SELECT username FROM users WHERE username = \"" + username + "\"")
	if err != nil {
		panic(err.Error())
	}
	var alreadyuser string
	defer db.Close()
	defer rows.Close()
	user := User{}
	for rows.Next() {
		if err := rows.Scan(&user.Username); err != nil {
			panic(err.Error())
		}
		alreadyuser = user.Username
	}
	return alreadyuser
}

func UserAdd(username, password string) {
	db, err := db.Conn()
	_, err = db.Query("INSERT INTO users (username, password) VALUES ( \"" + username + "\" , \"" + password + "\" )")
	if err != nil {
		panic(err.Error())
	}
}
