package auth

import (
	"ecommerce/delivery/helper"
	"ecommerce/usecase/auth"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	authUseCase auth.AuthUseCaseInterface
}

func NewAuthHandler(auth auth.AuthUseCaseInterface) *AuthHandler {
	return &AuthHandler{
		authUseCase: auth,
	}
}

func (ah *AuthHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		type loginData struct {
			Identifier string `json:"identifier"`
			Password   string `json:"password"`
		}
		var login loginData

		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error bind data"))
		}
		token, errorLogin := ah.authUseCase.Login(login.Identifier, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success login", responseToken))
	}
}
