package controllers

import (
	"net/http"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"github.com/labstack/echo/v4"
)

func GetCartForUser(c echo.Context) error {
	if Err != nil {
		panic("failed to connect database")
	}
	var uid = new(dto.UserIdDTO)
	if Err = c.Bind(uid); Err != nil {
		return Err
	}
	var foundCart models.Cart
	Db.Model(&models.Cart{}).Preload("Products").Find(uid).First(&foundCart)

	var totalPrice uint = 0
	for i := 0; i <= len(foundCart.ProductsInCart); i++ {
		totalPrice += foundCart.ProductsInCart[i].Product.Price * foundCart.ProductsInCart[i].Amount
	}

	cart := dto.CartDTO{
		Products:   foundCart.ProductsInCart,
		TotalPrice: totalPrice,
	}

	return c.JSON(http.StatusOK, cart)
}
