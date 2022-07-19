package controllers

import (
	"net/http"
	"strconv"

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
	Db.Where("Name = ?", p.Category).First(&foundCategory)

	eanUint, eanUintErr := strconv.ParseUint(p.EAN, 10, 32)
	priceUint, priceUintErr := strconv.ParseUint(p.Price, 10, 32)

	product := models.Product{
		Name:        p.Name,
		EAN:         eanUint,
		Price:       priceUint,
		Description: p.Description,
		CategoryID:  foundCategory.ID,
		Category:    foundCategory,
	}

	productStock := models.ProductStock{
		ProductID: product.ID,
		Product:   product,
		Stock:     0,
	}

	println(eanUintErr, priceUintErr)

	result := Db.Where("EAN = ?", eanUint).First(&models.Product{})

	if result.Error != nil {
		Db.Create(&product)
		Db.Create(&productStock)

		Db.Model(&models.Category{}).Where("Name = ?", foundCategory.Name).Update("ItemsAmount", foundCategory.ItemsAmount+1)

		return c.JSON(http.StatusCreated, product)
	}

	return c.JSON(http.StatusConflict, "Product already exists!")
}
