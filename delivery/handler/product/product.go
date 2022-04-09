package product

import (
	"ecommerce/delivery/helper"
	"ecommerce/delivery/middlewares"
	"ecommerce/entities"
	"ecommerce/usecase/product"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productUseCase product.ProductUseCaseInterface
}

func NewProductHandler(productUseCase product.ProductUseCaseInterface) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

func (ph *ProductHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := ph.productUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all product", products))
	}
}

func (ph *ProductHandler) GetProductById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))

		products, err := ph.productUseCase.GetProductById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get product by id", products))
	}
}

func (ph *ProductHandler) GetProductByIdUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		id = middlewares.ExtractToken(c)
		products, err := ph.productUseCase.GetProductByIdUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get product by id", products))
	}
}

func (ph *ProductHandler) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entities.Product
		c.Bind(&product)
		id := middlewares.ExtractToken(c)
		product.UserID = uint(id)
		err := ph.productUseCase.CreateProduct(product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create product"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success Create product"))
	}
}

func (ph *ProductHandler) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var userID = middlewares.ExtractToken(c)

		err := ph.productUseCase.DeleteProduct(id, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete product by id"))
	}
}

func (ph *ProductHandler) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var product entities.Product
		c.Bind(&product)

		err := ph.productUseCase.UpdateProduct(id, product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update product by id", product))
	}
}
