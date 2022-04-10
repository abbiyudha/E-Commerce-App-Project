package routes

import (
	"ecommerce/delivery/handler/auth"
	"ecommerce/delivery/handler/cartItem"
	"ecommerce/delivery/handler/product"
	"ecommerce/delivery/handler/user"
	"ecommerce/delivery/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *auth.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterPath(e *echo.Echo, uh *user.UserHandler, ph *product.ProductHandler, ch *cartItem.CartHandler) {

	e.GET("/users", uh.GetAllHandler(), middlewares.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserById(), middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.DELETE("/users", uh.DeleteUser(), middlewares.JWTMiddleware())
	e.PUT("/users", uh.UpdateUser(), middlewares.JWTMiddleware())

	e.GET("/products", ph.GetAllHandler())
	e.GET("/products/:id", ph.GetProductById())
	e.GET("/products/profile", ph.GetProductByIdUser(), middlewares.JWTMiddleware())
	e.POST("/products", ph.CreateProduct(), middlewares.JWTMiddleware())
	e.DELETE("/products/:id", ph.DeleteProduct(), middlewares.JWTMiddleware())
	e.PUT("/products/:id", ph.UpdateProduct(), middlewares.JWTMiddleware())

	e.POST("/carts", ch.CreateCart(), middlewares.JWTMiddleware())
	e.GET("/carts/profile", ch.GetCartByIdUser(), middlewares.JWTMiddleware())
	e.GET("/carts/:id", ch.GetCartByIdCart(), middlewares.JWTMiddleware())
	e.PUT("/carts/:id", ch.UpdateCartByIdCart(), middlewares.JWTMiddleware())
	e.DELETE("/carts/:id", ch.DeleteCart(), middlewares.JWTMiddleware())

}
