package route

import (
	"fmt"
	"keyboardify-server/controllers"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func InitUser(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%s Backend 0.0.1", os.Getenv("APP")))
	})

	g.GET("/su/:uid", controllers.CheckIfUserIsSuperUser)

	g.GET("/products", controllers.GetAllProducts)
	g.GET("/products/:id", controllers.GetProductById)

	g.GET("/categories", controllers.GetAllCategories)

}
