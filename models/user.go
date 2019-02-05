package models

import (
	"strconv"

	"github.com/sanshirookazaki/echo-mvc/db"
)

type User struct {
	Userid   int
	Username string
	Password string
	Tasks
}

type Userinterface interface {
	GetUserid(string, string) int
	GetUsernamepass(int) (string, string)
}

func NewUser() *User {
	return &User{}
}

// Get userid from username and password)
func (u *User) GetUserid(username string, password string) int {
	db, err := db.Conn()
	rows, err := db.Query("SELECT userid FROM users WHERE username = \"" + username + "\" and password = \"" + password + "\"")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var userid int
	for rows.Next() {
		if err := rows.Scan(&u.Userid); err != nil {
			panic(err.Error())
		}
		userid = u.Userid
	}
	return userid
}

// Get username from userid
func (u *User) GetUsernamepass(userid int) (username, password string) {
	db, err := db.Conn()
	rows, err := db.Query("SELECT username, password FROM users WHERE userid = " + strconv.Itoa(userid))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&u.Username, &u.Password); err != nil {
			panic(err.Error())
		}
		username = u.Username
		password = u.Password
	}
	return username, password
}
