package controllers

import (
	"net/http"
	"strconv"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/sanshirookazaki/echo-mvc/models"
)

func UserIndex(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	var ui models.Userinterface
	ui = models.NewUser()
	var t models.Task
	tasks := t.GetTaskAll(userid.(int))
	username, password := ui.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", user)
}

func UserAddTask(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	var u models.User
	var t models.Task
	tasks := t.GetTaskAll(userid.(int))
	username, password := u.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "add", user)
}

func UserAddTaskPost(c echo.Context) error {
	task := c.FormValue("task")
	session := session.Default(c)
	userid := session.Get("userid")
	var u models.User
	var t models.Task
	t.AddTask(userid.(int), task)
	tasks := t.GetTaskAll(userid.(int))
	username, password := u.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", user)
}

func UserDetailTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var u models.User
	var t models.Task
	tasks := t.GetTask(id)
	session := session.Default(c)
	userid := session.Get("userid")
	username, password := u.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "task", user)
}

func UserDeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	var u models.User
	var t models.Task
	t.DeleteTask(id)
	session := session.Default(c)
	userid := session.Get("userid")
	tasks := t.GetTaskAll(userid.(int))
	username, password := u.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", user)
}

func UserFinishTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	var u models.User
	var t models.Task
	t.FinishTask(id)
	session := session.Default(c)
	userid := session.Get("userid")
	tasks := t.GetTaskAll(userid.(int))
	username, password := u.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", user)
}

func UserHistory(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	var u models.User
	var t models.Task
	tasks := t.History(userid.(int))
	username, password := u.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "history", user)
}

func UserEditTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	session := session.Default(c)
	userid := session.Get("userid")
	var u models.User
	var t models.Task
	tasks := t.GetTask(id)
	username, password := u.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "edit", user)
}

func UserEditTaskPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task := c.FormValue("task")
	var u models.User
	var t models.Task
	t.EditTask(task, id)
	session := session.Default(c)
	userid := session.Get("userid")
	tasks := t.GetTaskAll(userid.(int))
	username, password := u.GetUsernamepass(userid.(int))
	user := models.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", user)
}
