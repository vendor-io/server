package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/foxsaysderp/keyboardify-server/controllers"
)

func main() {
	e := echo.New()
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":8000"))
	e.GET("/api/users/:id", controllers.GetUserById)
}
