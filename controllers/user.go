package controllers

import (
	"net/http"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	if Err != nil {
		panic("failed to connect database")
	}

	var users []models.User
	Db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	Db.Where("ID = ?", id).First(&user)
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	u := new(dto.UserDTO)
	if Err = c.Bind(u); Err != nil {
		return Err
	}

	user := models.User{
		Login:     u.Login,
		Email:     u.Email,
		Password:  u.Password,
		Cart:      models.Cart{},
		Orders:    []models.Order{},
		Addresses: []models.Address{},
	}
	Db.Create(&user)

	return c.JSON(http.StatusCreated, user)
}
