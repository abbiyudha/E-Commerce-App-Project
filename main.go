package main

import (
	"ecommerce/config"
	"ecommerce/delivery/middlewares"
	"ecommerce/delivery/routes"
	"ecommerce/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	userhandler "ecommerce/delivery/handler/user"
	userrepo "ecommerce/repository/user"
	userusecase "ecommerce/usecase/user"

	authhandler "ecommerce/delivery/handler/auth"
	authrepo "ecommerce/repository/auth"
	authusecase "ecommerce/usecase/auth"

	producthandler "ecommerce/delivery/handler/product"
	productrepository "ecommerce/repository/product"
	productusecase "ecommerce/usecase/product"
)

func main() {

	configs := config.GetConfig()
	db := utils.InitDB(configs)

	userRepo := userrepo.NewUserRepository(db)
	userUseCase := userusecase.NewUserUseCase(userRepo)
	userHandler := userhandler.NewUserHandler(userUseCase)

	authRepo := authrepo.NewAuthRepository(db)
	authUseCase := authusecase.NewAuthUseCase(authRepo)
	authHandler := authhandler.NewAuthHandler(authUseCase)

	productRepo := productrepository.NewProductRepository(db)
	productUseCase := productusecase.NewProductUseCase(productRepo)
	productHandler := producthandler.NewProductHandler(productUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middlewares.CustomLogger())

	routes.RegisterAuthPath(e, authHandler)
	routes.RegisterPath(e, userHandler, productHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))

}
