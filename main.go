package main

import (
	"ecommerce/config"
	"ecommerce/delivery/middlewares"
	"ecommerce/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	configs := config.GetConfig()
	utils.InitDB(configs)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middlewares.CustomLogger())

	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))

}
