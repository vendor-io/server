package controllers

import (
	"keyboardify-server/models"
	"keyboardify-server/models/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetOrdersForUser(c echo.Context) error {
	uid := c.Param("uid")
	foundUser := FindUserViaUID(uid)

	var userOrders []models.Order
	Db.Where("user_id = ?", foundUser.ID).Find(&userOrders)

	var userOrdersDto []dto.OrderItemDTO
	for i := range userOrders {
		userOrderDto := OrderDTOResolver(userOrders[i], foundUser)

		userOrdersDto = append(userOrdersDto, userOrderDto)
	}

	for i, j := 0, len(userOrdersDto)-1; i < j; i, j = i+1, j-1 {
		userOrdersDto[i], userOrdersDto[j] = userOrdersDto[j], userOrdersDto[i]
	}

	return c.JSON(http.StatusOK, userOrdersDto)
}

func GetOrderForUserById(c echo.Context) error {
	oid := c.Param("oid")

	var foundOrder models.Order
	Db.Where("id = ?", oid).First(&foundOrder)

	var foundUser models.User
	Db.Where("id = ?", foundOrder.UserID).First(&foundUser)

	userOrderDto := OrderDTOResolver(foundOrder, foundUser)

	return c.JSON(http.StatusOK, userOrderDto)
}

func CreateOrderForUser(c echo.Context) error {
	var newOrder = new(dto.OrderDTO)

	if Err = c.Bind(newOrder); Err != nil {
		return Err
	}

	var lastOrder models.Order
	Db.Last(&lastOrder)

	foundUser := FindUserViaUID(newOrder.UserID)

	var products []models.Product

	for i := range newOrder.ProductsInOrder {
		var foundProduct models.Product
		Db.Where("id = ?", newOrder.ProductsInOrder[i].ID).First(&foundProduct)

		products = append(products, foundProduct)
	}

	var productsInOrder []models.ProductInOrder
	for i := range products {
		var productInOrderToCreate = models.ProductInOrder{
			ProductID: products[i].ID,
			Product:   products[i],
			Amount:    uint64(newOrder.ProductsInOrder[i].Amount),
			OrderID:   lastOrder.ID + 1,
		}

		productsInOrder = append(productsInOrder, productInOrderToCreate)
	}

	var foundAddress models.Address
	Db.Where("id = ?", newOrder.AddressID).First(&foundAddress)

	var newAddressInOrder = models.AddressInOrder{
		AddressID: foundAddress.ID,
		OrderID:   lastOrder.ID + 1,
	}
	Db.Create(&newAddressInOrder)

	var order = models.Order{
		UserID:          foundUser.ID,
		User:            foundUser,
		AddressInOrder:  newAddressInOrder,
		ProductsInOrder: productsInOrder,
		TotalPrice:      newOrder.TotalPrice,
		OrderStatus:     models.IN_PROGRESS,
		IsPaid:          true,
	}

	Db.Create(&order)

	var cartToRemove models.Cart
	Db.Where("user_id = ?", foundUser.ID).First(&cartToRemove)

	Db.Where("cart_id = ?", cartToRemove.ID).Delete(&models.CartProduct{})

	return c.JSON(http.StatusCreated, order)
}
