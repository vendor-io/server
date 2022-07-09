package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
	// "github.com/foxsaysderp/keyboardify-server/models"
)

// func ShowAllUsers(c echo.Context) {
// 	db, err := gorm.Open(sqlite.Open("keyboardify_gorm.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	var users []models.User
// 	db.Find(&users)

// 	return c.String(http.StatusOK, "these are users:")
// }

func GetUserById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
