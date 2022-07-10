package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"keyboardify-server/controllers"
	"keyboardify-server/db"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db.Init()

	e.GET("/api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/api/users/all", controllers.GetAllUsers)
	e.GET("/api/users/:id", controllers.GetUserById)
	e.POST("/api/users", controllers.CreateUser)

	e.Logger.Fatal(e.Start(":8000"))
}
