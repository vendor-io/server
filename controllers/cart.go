package controllers

import (
	"net/http"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"github.com/labstack/echo/v4"
)

func GetCartForUser(c echo.Context) error {
	id := c.Param("id")

	var foundCart models.Cart
	Db.Where("user_id = ?", id).First(&foundCart)

	var totalPrice uint64 = 0
	for i := 0; i <= len(foundCart.ProductsInCart); i++ {
		totalPrice += foundCart.ProductsInCart[i].Product.Price * foundCart.ProductsInCart[i].Amount
	}

	cart := dto.CartDTO{
		Products:   foundCart.ProductsInCart,
		TotalPrice: totalPrice,
	}

	return c.JSON(http.StatusOK, cart)
}
