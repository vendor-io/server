package route

import (
	"keyboardify-server/controllers"

	"github.com/labstack/echo/v4"
)

func InitCommon(g *echo.Group) {
	g.POST("/user/new", controllers.CreateUser)
	g.POST("/superuser/assign", controllers.AssignSuperUser)
}
