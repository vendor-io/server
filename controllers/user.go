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
	Db.Where("id = ?", id).First(&user)
	return c.JSON(http.StatusOK, user)
}

func GetUserByUid(c echo.Context) error {
	uid := c.Param("uid")
	user := new(models.User)
	Db.Where("uid = ?", uid).First(&user)
	return c.JSON(http.StatusOK, user)
}

func CheckIfUserIsSuperUser(c echo.Context) error {
	uid := c.Param("uid")
	user := new(models.User)
	result := Db.Where("uid = ? AND is_super_user = ?", uid, true).First(&user)

	if result.Error != nil {
		return c.JSON(http.StatusOK, "false")
	}

	return c.JSON(http.StatusOK, "true")
}

func CreateUser(c echo.Context) error {
	u := new(dto.NewUserDTO)
	if Err = c.Bind(u); Err != nil {
		return Err
	}

	result := Db.Where("uid = ? AND email = ?", u.UID, u.Email).Find(&models.User{})

	if result.Error != nil {
		return c.String(http.StatusCreated, "User is already created.")
	}

	user := models.User{
		Email:       u.Email,
		UID:         u.UID,
		IsSuperUser: false,
		Cart:        models.Cart{},
		Orders:      []models.Order{},
		Addresses:   []models.Address{},
	}
	Db.Create(&user)

	return c.JSON(http.StatusCreated, user)
}
