package route

import (
	"keyboardify-server/controllers"

	"github.com/labstack/echo/v4"
)

func InitSuperUser(g *echo.Group) {
	g.GET("/all", controllers.GetAllUsers)
	g.GET("/:id", controllers.GetUserById)
	g.GET("/:uid", controllers.GetUserByUid)
	g.POST("/new", controllers.CreateUser)

	g.POST("/products", controllers.AddNewProduct)
	g.POST("/categories", controllers.AddNewCategory)
}
