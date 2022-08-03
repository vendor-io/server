package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		firebaseAuth := c.Get("firebaseAuth").(*auth.Client)

		authorizationToken := c.Request().Header.Get("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
			c.String(http.StatusBadRequest, "Id token not available")
			panic("Id token missing")
		}

		token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)

		if err != nil {
			c.String(http.StatusBadRequest, "Invalid Token")
			return err
		}

		c.Set("UUID", token.UID)
		return next(c)
	}
}
