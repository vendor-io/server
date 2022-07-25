package route

import (
	"fmt"
	"keyboardify-server/controllers"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	g.Static("/public", "public")

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%s Backend 0.0.1", os.Getenv("APP")))
	})

	g.GET("/users", controllers.GetAllUsers)
	g.GET("/users/:id", controllers.GetUserById)
	g.POST("/users", controllers.CreateUser)

	g.GET("/products", controllers.GetAllProducts)
	g.GET("/products/:id", controllers.GetProductById)
	g.POST("/products", controllers.AddNewProduct)

	g.GET("/categories", controllers.GetAllCategories)
	g.POST("/categories", controllers.AddNewCategory)
}
