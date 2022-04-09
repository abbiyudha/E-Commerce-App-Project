package product

import (
	"ecommerce/entities"
	"ecommerce/repository/product"
)

type ProductUseCaseInterface interface {
	GetAll() ([]entities.Product, error)
	GetProductById(id int) (entities.Product, error)
	GetProductByIdUser(id int) ([]entities.Product, error)
	CreateProduct(Product entities.Product) error
	DeleteProduct(id, userID int) error
	UpdateProduct(id int, Product entities.Product) error
}

type ProductUseCase struct {
	ProductRepository product.ProductRepositoryInterface
}

func NewProductUseCase(productRepo product.ProductRepositoryInterface) ProductUseCaseInterface {
	return &ProductUseCase{
		ProductRepository: productRepo,
	}

}

func (puc *ProductUseCase) GetAll() ([]entities.Product, error) {
	Products, err := puc.ProductRepository.GetAll()
	return Products, err
}

func (puc *ProductUseCase) GetProductById(id int) (entities.Product, error) {
	Product, err := puc.ProductRepository.GetProductById(id)
	return Product, err
}

func (puc *ProductUseCase) GetProductByIdUser(id int) ([]entities.Product, error) {
	Products, err := puc.ProductRepository.GetProductByIdUser(id)
	return Products, err
}

func (puc *ProductUseCase) CreateProduct(Product entities.Product) error {
	err := puc.ProductRepository.CreateProduct(Product)
	return err
}

func (puc *ProductUseCase) DeleteProduct(id, userID int) error {
	err := puc.ProductRepository.DeleteProduct(id, userID)
	return err
}

func (puc *ProductUseCase) UpdateProduct(id int, Product entities.Product) error {
	err := puc.ProductRepository.UpdateProduct(id, Product)
	return err
}
