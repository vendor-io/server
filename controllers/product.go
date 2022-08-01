package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"
	"keyboardify-server/util"

	"github.com/labstack/echo/v4"
)

func GetAllProducts(c echo.Context) error {
	var products []models.Product
	Db.Find(&products)

	for i, j := 0, len(products)-1; i < j; i, j = i+1, j-1 {
		products[i], products[j] = products[j], products[i]
	}

	return c.JSON(http.StatusOK, products)
}

func GetProductById(c echo.Context) error {
	id := c.Param("id")
	product := new(models.Product)
	Db.Where("ID = ?", id).First(&product)
	Db.Where("ID = ?", product.CategoryID).First(&product.Category)
	return c.JSON(http.StatusOK, product)
}

func AddNewProduct(c echo.Context) error {
	var p = dto.ProductDTO{
		Name:        c.FormValue("productName"),
		EAN:         c.FormValue("productEan"),
		Price:       c.FormValue("productPrice"),
		Description: c.FormValue("productDescription"),
		Category:    c.FormValue("productCategory"),
	}

	images := util.MultipleFileUpload(c, "productImages", "images")

	imagesToString := strings.Join(images, ";")

	var foundCategory models.Category
	Db.Where("Name = ?", c.FormValue("productCategory")).First(&foundCategory)

	eanUint, eanUintErr := strconv.ParseUint(p.EAN, 10, 32)
	priceUint, priceUintErr := strconv.ParseUint(p.Price, 10, 32)

	product := models.Product{
		Name:        p.Name,
		EAN:         eanUint,
		Price:       priceUint,
		Description: p.Description,
		Category:    foundCategory,
		MainImage:   images[0],
		Images:      imagesToString,
		Stock:       0,
	}

	println(eanUintErr, priceUintErr)

	result := Db.Where("EAN = ?", eanUint).First(&models.Product{})

	if result.Error != nil {
		Db.Create(&product)

		Db.Model(&models.Category{}).Where("Name = ?", foundCategory.Name).Update("ItemsAmount", foundCategory.ItemsAmount+1)

		return c.JSON(http.StatusCreated, product)
	}

	return c.JSON(http.StatusConflict, "Product already exists!")
}
