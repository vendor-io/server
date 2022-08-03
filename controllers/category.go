package controllers

import (
	"net/http"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"
	"keyboardify-server/util"

	"github.com/labstack/echo/v4"
)

func GetAllCategories(c echo.Context) error {
	var categories []models.Category
	Db.Find(&categories)

	return c.JSON(http.StatusOK, categories)
}

func GetCategoryBySlug(c echo.Context) error {
	slug := c.Param("slug")
	category := new(models.Category)
	Db.Where("Slug = ?", slug).First(&category)

	return c.JSON(http.StatusOK, category)
}

func AddNewCategory(c echo.Context) error {
	var cat = new(dto.CategoryDTO)

	if Err = c.Bind(cat); Err != nil {
		return Err
	}

	category := models.Category{
		Name:        cat.Name,
		Slug:        util.ToKebabCase(cat.Name),
		ItemsAmount: 0,
	}

	result := Db.Where("Name = ?", category.Name).First(&models.Category{})

	if result.Error != nil {
		Db.Create(&category)
		return c.JSON(http.StatusCreated, category)
	}

	return c.JSON(http.StatusConflict, "Category already exists!")

}
