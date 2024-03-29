package route

import (
	"keyboardify-server/controllers"

	"github.com/labstack/echo/v4"
)

func InitUser(g *echo.Group) {
	g.GET("/su/:uid", controllers.CheckIfUserIsSuperUser)

	g.GET("/products", controllers.GetAllProducts)
	g.GET("/products/:id", controllers.GetProductById)
	g.GET("/products/category/:slug", controllers.GetProductsByCategorySlug)

	g.GET("/categories", controllers.GetAllCategories)
	g.GET("/categories/:slug", controllers.GetCategoryBySlug)

	g.GET("/cart/:uid", controllers.GetCartForUser)
	g.POST("/cart/add", controllers.AddProductToCart)
	g.POST("/cart/remove", controllers.RemoveProductFromCart)
	g.POST("/cart/change", controllers.ChangeAmountOfProductInCart)

	g.GET("/addresses/:uid", controllers.GetAddressesForUser)
	g.POST("/addresses/new", controllers.AddNewAddress)

	g.POST("/payment/charge", controllers.PostCharge)
	g.POST("/payment/intent", controllers.PostPaymentIntent)

	g.GET("/orders/:oid", controllers.GetOrderForUserById)
	g.GET("/orders/user/:uid", controllers.GetOrdersForUser)
	g.POST("/orders/new", controllers.CreateOrderForUser)
}
