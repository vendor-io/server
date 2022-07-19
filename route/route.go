package route

import (
	"keyboardify-server/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Keyboardify Backend 0.0.1")
	})
	g.GET("/users", controllers.GetAllUsers)
	g.GET("/users/:id", controllers.GetUserById)
	g.POST("/users", controllers.CreateUser)

	g.GET("/products", controllers.GetAllProducts)
	g.POST("/products", controllers.AddNewProduct)

	g.GET("/categories", controllers.GetAllCategories)
	g.POST("/categories", controllers.AddNewCategory)
}
