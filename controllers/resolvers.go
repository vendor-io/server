package controllers

import (
	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"gorm.io/gorm"
)

func FindUserViaUID(uid any) (user models.User) {
	Db.Where("uid = ?", uid).First(&user)

	return user
}

func CartDTOResolver(ids []uint, cartId uint) dto.CartWithTotalPriceDTO {
	var totalPrice uint = 0
	for i := range ids {
		var temporaryProductFromCart models.Product
		var productToExtractAmount models.CartProduct

		Db.Where("cart_id = ? AND product_id = ?", cartId, ids[i]).First(&productToExtractAmount)
		Db.Where("id = ?", ids[i]).First(&temporaryProductFromCart)

		totalPrice += temporaryProductFromCart.Price * productToExtractAmount.Amount
	}

	var productsToDTO []models.Product
	Db.Where(ids).Find(&productsToDTO)

	var productsDTO []dto.ProductInCartDTO
	for i := range productsToDTO {
		var categoryToDTO models.Category
		Db.Where("id = ?", productsToDTO[i].CategoryID).First(&categoryToDTO)

		var productToExtractAmount models.CartProduct
		Db.Where("cart_id = ? AND product_id = ?", cartId, productsToDTO[i].ID).First(&productToExtractAmount)

		var newProductToCartDTO = dto.ProductInCartDTO{
			ID:           productsToDTO[i].ID,
			Name:         productsToDTO[i].Name,
			MainImage:    productsToDTO[i].MainImage,
			Price:        productsToDTO[i].Price,
			Amount:       productToExtractAmount.Amount,
			TotalPrice:   productToExtractAmount.Amount * productsToDTO[i].Price,
			CategoryName: categoryToDTO.Name,
			CategorySlug: categoryToDTO.Slug,
		}

		productsDTO = append(productsDTO, newProductToCartDTO)
	}

	cart := dto.CartWithTotalPriceDTO{
		Products:   productsDTO,
		TotalPrice: totalPrice,
	}

	return cart
}

func ControllerDetailsResolver(cartDto dto.CartProductDTO) (models.User, models.Product, models.Cart, *gorm.DB) {
	var user models.User
	Db.Where("uid = ?", cartDto.UserID).First(&user)

	var product models.Product
	Db.Where("id = ?", cartDto.ProductID).First(&product)

	var cart models.Cart
	result := Db.Where("user_id = ?", user.ID).First(&cart)

	return user, product, cart, result
}

func OrderDTOResolver(order models.Order, user models.User) dto.OrderItemDTO {
	var tempUserOrderAddress models.AddressInOrder
	Db.Where("order_id = ?", order.AddressInOrder.AddressID).First(&tempUserOrderAddress)

	var orderAddress models.Address
	Db.Where("id = ?", tempUserOrderAddress.AddressID).First(&orderAddress)

	var orderAddressDto = dto.PlainAddressDTO{
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
	Db.Where("order_id = ?", order.ID).Find(&productsInOrder)

	for i := range productsInOrder {
		var tempProduct models.Product
		Db.Where("id = ?", productsInOrder[i].ProductID).First(&tempProduct)

		var tempProductCategory models.Category
		Db.Where("id = ?", tempProduct.CategoryID).First(&tempProductCategory)

		tempProduct.Category = tempProductCategory
		productsInOrder[i].Product = tempProduct
	}

	var orderDto = dto.OrderItemDTO{
		ID:              order.ID,
		CreatedAt:       order.CreatedAt,
		UserID:          user.ID,
		UID:             user.UID,
		Address:         orderAddressDto,
		ProductsInOrder: productsInOrder,
		TotalPrice:      order.TotalPrice,
		OrderStatus:     string(order.OrderStatus),
		IsPaid:          order.IsPaid,
	}

	return orderDto
}
