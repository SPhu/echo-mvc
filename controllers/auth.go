package controllers

import (
	"html"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/sanshirookazaki/echo-mvc/models"
)

type LoginForm struct {
	Username string
	Password string
}

func Login(c echo.Context) error {
	//session := session.Default(c)
	//loginid := session.Get("login")
	return c.Render(http.StatusOK, "login", "ログインしてください")
}

func Logout(c echo.Context) error {
	session := session.Default(c)
	session.Clear()
	return c.Render(http.StatusOK, "login", "ログインしてください")
}

func LoginCheck(c echo.Context) error {
	var u models.User
	loginform := LoginForm{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	username := html.EscapeString(loginform.Username)
	password := html.EscapeString(loginform.Password)

	if username == "" && password == "" {
		return c.Render(http.StatusOK, "login", "ユーザー名とパスワードを入力して下さい。")
	} else if username == "" {
		return c.Render(http.StatusOK, "login", "ユーザー名を入力して下さい。")
	} else if password == "" {
		return c.Render(http.StatusOK, "login", "パスワードを入力して下さい。")
	}

	session := session.Default(c)
	loginpass := models.GetPassword(username)
	if password == loginpass {
		userid := u.GetUserid(username, password)
		//tasks := models.GetTaskAll(userid)
		session.Set("userid", userid)
		session.Save()
		return c.Redirect(http.StatusFound, "/"+username+"/index")
	} else {
		return c.Redirect(http.StatusFound, "login")
	}
}

func LoginNewuser(c echo.Context) error {
	return c.Render(http.StatusFound, "loginnew", "ユーザー名とパスワードを入力して下さい。")
}

func LoginAdduser(c echo.Context) error {
	var u models.User
	loginform := LoginForm{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	username := html.EscapeString(loginform.Username)
	password := html.EscapeString(loginform.Password)

	session := session.Default(c)
	loginname := models.UserUniqueCheck(username)
	if username == loginname {
		return c.Render(http.StatusFound, "loginnew", "このユーザー名は既に使われてます。")
	} else if username == "" || password == "" {
		return c.Render(http.StatusFound, "loginnew", "ユーザー名とパスワードに空白は使えません。")
	} else {
		models.UserAdd(username, password)
		userid := u.GetUserid(username, password)
		session.Set("userid", userid)
		session.Save()
		return c.Redirect(http.StatusFound, "/"+username+"/index")
	}
}
