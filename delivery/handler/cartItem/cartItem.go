package cartItem

import (
	"ecommerce/delivery/helper"
	"ecommerce/delivery/middlewares"
	"ecommerce/entities"
	"ecommerce/usecase/cartItem"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CartHandler struct {
	cartUseCase cartItem.CartUseCaseInterface
}

func NewCartHandler(cartUseCase cartItem.CartUseCaseInterface) *CartHandler {
	return &CartHandler{
		cartUseCase: cartUseCase,
	}
}

func (ch *CartHandler) CreateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var cart entities.CartItem
		c.Bind(&cart)
		id := middlewares.ExtractToken(c)
		cart.UserID = uint(id)

		err := ch.cartUseCase.CreateCart(cart)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success add product"))
	}
}

func (ch *CartHandler) GetCartByIdUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		id = middlewares.ExtractToken(c)
		carts, err := ch.cartUseCase.GetCartByIdUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get carts by id", carts))
	}
}

func (ch *CartHandler) GetCartByIdCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))

		var UserID = middlewares.ExtractToken(c)
		cart, err := ch.cartUseCase.GetCartByIdCart(id, UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get cart by id", cart))
	}
}

func (ch *CartHandler) UpdateCartByIdCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var UserID = middlewares.ExtractToken(c)
		var cart entities.CartItem
		c.Bind(&cart)

		err := ch.cartUseCase.UpdateCart(id, UserID, cart)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update cart", cart))
	}
}

func (ch *CartHandler) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var userID = middlewares.ExtractToken(c)

		err := ch.cartUseCase.DeleteCart(id, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete cart"))
	}
}
