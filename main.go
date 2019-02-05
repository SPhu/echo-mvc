package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	c "github.com/sanshirookazaki/echo-mvc/controllers"
	"github.com/sanshirookazaki/echo-mvc/util"
)

func main() {

	t := util.NewTemplate("views/*.html")

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Renderer = t

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.Render(http.StatusNotFound, "404", nil)
			} else {
				c.Render(http.StatusInternalServerError, "500", nil)
			}
		}
	}

	e.Static("/static", "static")
	store := session.NewCookieStore([]byte("secret-key"))
	store.MaxAge(86400)
	e.Use(session.Sessions("ESESSION", store))
	e.GET("/", c.Login)
	e.GET("/login", c.Login)
	e.POST("/login", c.LoginCheck)
	e.POST("/logout", c.Logout)
	e.GET("/login/new", c.LoginNewuser)
	e.POST("/login/new", c.LoginAdduser)
	e.GET("/:username/index", c.UserIndex)
	e.GET("/:username/task/:id", c.UserDetailTask)
	e.GET("/:username/task/add", c.UserAddTask)
	e.POST("/:username/task/add", c.UserAddTaskPost)
	e.POST("/:username/task/delete", c.UserDeleteTask)
	e.GET("/:username/task/:id/edit", c.UserEditTask)
	e.POST("/:username/task/:id/edit", c.UserEditTaskPost)
	e.POST("/:username/task/finish", c.UserFinishTask)
	e.GET("/:username/task/history", c.UserHistory)
	e.Logger.Fatal(e.Start(":1323"))
}
