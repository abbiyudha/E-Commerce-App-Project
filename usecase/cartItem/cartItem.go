package cartItem

import (
	"ecommerce/entities"
	"ecommerce/repository/cartItem"
	"ecommerce/repository/product"
	"errors"
)

type CartUseCaseInterface interface {
	CreateCart(cart entities.CartItem) error
	GetCartByIdUser(id int) ([]entities.CartItem, error)
	GetCartByIdCart(id, UserID int) (entities.CartItem, error)
	UpdateCart(id, userID int, cart entities.CartItem) error
	DeleteCart(id, userID int) error
}

type CartUseCase struct {
	CartRepository    cartItem.CartRepositoryInterface
	ProductRepository product.ProductRepositoryInterface
}

func NewCartUseCase(cartRepo cartItem.CartRepositoryInterface, productRepo product.ProductRepositoryInterface) CartUseCaseInterface {
	return &CartUseCase{
		CartRepository:    cartRepo,
		ProductRepository: productRepo,
	}

}

func (cuc *CartUseCase) CreateCart(cart entities.CartItem) error {

	var product, err = cuc.ProductRepository.GetProductById(int(cart.ProductID))

	if err != nil {
		return errors.New("product not found")
	}
	if cart.TotalProduct > product.Stock {
		return errors.New("Out of stock")
	}

	var totalHarga = product.Price * cart.TotalProduct

	cart.TotalPrice = totalHarga

	err = cuc.CartRepository.CreateCart(cart)

	return err
}

func (cuc *CartUseCase) UpdateCart(id, userID int, cart entities.CartItem) error {

	var (
		carts    entities.CartItem
		products entities.Product
		err      error
	)

	carts, err = cuc.CartRepository.GetCartByIdCart(id, userID)
	if err != nil {
		return errors.New("cart not found")
	}

	products, err = cuc.ProductRepository.GetProductById(int(carts.ProductID))

	if cart.TotalProduct > products.Stock {
		return errors.New("Out of stock")
	}
	var totalHarga = products.Price * cart.TotalProduct
	cart.TotalPrice = totalHarga

	err = cuc.CartRepository.UpdateCart(id, userID, cart)

	return err
}

func (cuc *CartUseCase) GetCartByIdUser(id int) ([]entities.CartItem, error) {

	carts, err := cuc.CartRepository.GetCartByIdUser(id)
	return carts, err

}

func (cuc *CartUseCase) GetCartByIdCart(id, UserID int) (entities.CartItem, error) {
	cart, err := cuc.CartRepository.GetCartByIdCart(id, UserID)
	return cart, err
}

func (cuc *CartUseCase) DeleteCart(id, userID int) error {
	err := cuc.CartRepository.DeleteCart(id, userID)
	return err
}
