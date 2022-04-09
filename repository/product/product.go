package product

import (
	"ecommerce/entities"
	"fmt"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	GetAll() ([]entities.Product, error)
	GetProductById(id int) (entities.Product, error)
	GetProductByIdUser(id int) ([]entities.Product, error)
	CreateProduct(product entities.Product) error
	DeleteProduct(id, userID int) error
	UpdateProduct(id int, product entities.Product) error
}

type ProductRepository struct {
	database *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		database: db,
	}
}

func (pr *ProductRepository) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	tx := pr.database.Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil

}

func (pr *ProductRepository) GetProductById(id int) (entities.Product, error) {
	var products entities.Product
	tx := pr.database.Where("id = ?", id).First(&products)
	if tx.Error != nil {
		return products, tx.Error
	}
	return products, nil

}

func (pr *ProductRepository) GetProductByIdUser(id int) ([]entities.Product, error) {
	var products []entities.Product
	tx := pr.database.Where("user_id = ?", id).Find(&products)
	if tx.Error != nil {
		return products, tx.Error
	}
	return products, nil

}

func (pr *ProductRepository) CreateProduct(product entities.Product) error {

	tx := pr.database.Save(&product)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (pr *ProductRepository) DeleteProduct(id, userID int) error {
	var products entities.Product
	err := pr.database.Where("id =? and user_id = ?", id, userID).First(&products).Error
	if err != nil {
		return err
	}
	pr.database.Delete(&products)

	return nil

}

func (pr *ProductRepository) UpdateProduct(id int, product entities.Product) error {

	tx := pr.database.Where("id = ?", id).Updates(&product)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
