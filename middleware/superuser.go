package middleware

import (
	"keyboardify-server/controllers"
	"keyboardify-server/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SuperuserAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		uidToken := c.Request().Header.Get("UID")

		if uidToken == "" {
			c.String(http.StatusBadRequest, "Id token not available")
			panic("Id token missing")
		}

		result := controllers.Db.Where("UID = ? AND IsSuperUser = ?", uidToken, true).First(&models.User{})

		if result.Error != nil {
			return c.JSON(http.StatusUnauthorized, "User has insufficient privileges!")
		}

		return next(c)
	}
}
