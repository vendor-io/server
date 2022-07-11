package controllers

import (
	"net/http"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"github.com/labstack/echo/v4"
)

func GetAllProducts(c echo.Context) error {
	var products []models.Product
	Db.Find(&products)

	return c.JSON(http.StatusOK, products)
}

func AddNewProduct(c echo.Context) error {
	var p = new(dto.ProductDTO)

	if Err = c.Bind(p); Err != nil {
		return Err
	}

	var foundCategory models.Category
	Db.Model(&models.Category{}).Preload("Name").Find(p.Category).First(&foundCategory)

	product := models.Product{
		Name:        p.Name,
		EAN:         p.EAN,
		Price:       p.Price,
		Description: p.Description,
		CategoryID:  foundCategory.ID,
		Category:    foundCategory,
	}

	productStock := models.ProductStock{
		ProductID: product.ID,
		Product:   product,
		Stock:     0,
	}

	Db.Create(&product)
	Db.Create(&productStock)

	return c.JSON(http.StatusCreated, product)
}
