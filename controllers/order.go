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
		var tempUserOrderAddress models.AddressInOrder
		Db.Where("order_id = ?", userOrders[i].ID).First(&tempUserOrderAddress)

		var orderAddress models.Address
		Db.Where("id = ?", tempUserOrderAddress.AddressID).First(&orderAddress)

		var orderAddressDto = dto.AddressDTO{
			UserID:      foundUser.UID,
			FirstName:   orderAddress.FirstName,
			LastName:    orderAddress.LastName,
			Street:      orderAddress.Street,
			HouseNumber: orderAddress.HouseNumber,
			PostalCode:  orderAddress.PostalCode,
			City:        orderAddress.City,
			Country:     orderAddress.Country,
			PhoneNumber: orderAddress.PhoneNumber,
		}

		var productsInOrder []models.ProductInOrder
		Db.Where("order_id = ?", userOrders[i].ID).Find(&productsInOrder)

		for j := range productsInOrder {
			var tempProduct models.Product
			Db.Where("id = ?", productsInOrder[j].ProductID).First(&tempProduct)

			var tempProductCategory models.Category
			Db.Where("id = ?", tempProduct.CategoryID).First(&tempProductCategory)

			tempProduct.Category = tempProductCategory
			productsInOrder[j].Product = tempProduct
		}

		var tempUserOrderDto = dto.OrderItemDTO{
			ID:              userOrders[i].ID,
			CreatedAt:       userOrders[i].CreatedAt,
			UserID:          foundUser.ID,
			Address:         orderAddressDto,
			ProductsInOrder: productsInOrder,
			TotalPrice:      userOrders[i].TotalPrice,
			OrderStatus:     string(userOrders[i].OrderStatus),
			IsPaid:          userOrders[i].IsPaid,
		}

		userOrdersDto = append(userOrdersDto, tempUserOrderDto)
	}

	for i, j := 0, len(userOrdersDto)-1; i < j; i, j = i+1, j-1 {
		userOrdersDto[i], userOrdersDto[j] = userOrdersDto[j], userOrdersDto[i]
	}

	return c.JSON(http.StatusOK, userOrdersDto)
}

func CreateOrderForUser(c echo.Context) error {
	var newOrder = new(dto.OrderDTO)

	if Err = c.Bind(newOrder); Err != nil {
		return Err
	}

	var order = new(models.Order)

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
			OrderID:   order.ID,
		}

		productsInOrder = append(productsInOrder, productInOrderToCreate)
	}

	var foundAddress models.Address
	Db.Where("id = ?", newOrder.AddressID).First(&foundAddress)

	var newAddressInOrder = models.AddressInOrder{
		AddressID: foundAddress.ID,
		OrderID:   order.ID,
	}
	Db.Create(&newAddressInOrder)

	order.UserID = foundUser.ID
	order.User = foundUser
	order.AddressInOrder = newAddressInOrder
	order.ProductsInOrder = productsInOrder
	order.TotalPrice = newOrder.TotalPrice
	order.OrderStatus = models.IN_PROGRESS
	order.IsPaid = true

	Db.Create(&order)

	// var cartToRemove models.Cart
	// Db.Where("user_id = ?", foundUser.ID).First(&cartToRemove)

	// Db.Where("cart_id = ?", cartToRemove.ID).Delete(&models.CartProduct{})

	return c.JSON(http.StatusCreated, order)
}
