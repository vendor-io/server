package controllers

import (
	"net/http"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"github.com/labstack/echo/v4"
)

func GetAllCategories(c echo.Context) error {
	var categories []models.Category
	Db.Find(&categories)

	return c.JSON(http.StatusOK, categories)
}

func AddNewCategory(c echo.Context) error {
	var cat = new(dto.CategoryDTO)

	if Err = c.Bind(cat); Err != nil {
		return Err
	}

	category := models.Category{
		Name:        cat.Name,
		ItemsAmount: 0,
	}

	result := Db.Where("Name = ?", category.Name).First(&models.Category{})

	if result.Error != nil {
		Db.Create(&category)
		return c.JSON(http.StatusCreated, category)
	}

	return c.JSON(http.StatusConflict, "Category already exists!")

}
