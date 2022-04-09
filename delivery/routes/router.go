package routes

import (
	"ecommerce/delivery/handler/auth"
	"ecommerce/delivery/handler/user"
	"ecommerce/delivery/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *auth.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterPath(e *echo.Echo, uh *user.UserHandler) {

	e.GET("/users", uh.GetAllHandler(), middlewares.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserById(), middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.DELETE("/users", uh.DeleteUser(), middlewares.JWTMiddleware())
	e.PUT("/users", uh.UpdateUser(), middlewares.JWTMiddleware())

}
