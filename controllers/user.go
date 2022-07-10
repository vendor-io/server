package controllers

import (
	"net/http"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, err = gorm.Open(sqlite.Open("keyboardify_gorm.db"), &gorm.Config{})

func GetAllUsers(c echo.Context) error {
	if err != nil {
		panic("failed to connect database")
	}

	var users []models.User
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func CreateUser(c echo.Context) error {
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return err
	}

	user := dto.UserDTO{
		Login:    u.Login,
		Email:    u.Email,
		Password: u.Password,
	}
	db.Create(u)

	return c.JSON(http.StatusCreated, user)
}
