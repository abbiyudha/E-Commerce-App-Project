package cartItem

import (
	"ecommerce/entities"
	"gorm.io/gorm"
)

type CartRepositoryInterface interface {
	CreateCart(cart entities.CartItem) error
	GetCartByIdUser(id int) ([]entities.CartItem, error)
	GetCartByIdCart(id, UserID int) (entities.CartItem, error)
	UpdateCart(id, UserID int, cart entities.CartItem) error
	DeleteCart(id, userID int) error
}

type CartRepository struct {
	database *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		database: db,
	}
}

func (cr *CartRepository) CreateCart(cart entities.CartItem) error {

	tx := cr.database.Create(&cart)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (cr *CartRepository) GetCartByIdUser(id int) ([]entities.CartItem, error) {
	var carts []entities.CartItem
	tx := cr.database.Where("user_id = ?", id).Find(&carts)

	if tx.Error != nil {
		return carts, tx.Error
	}
	return carts, nil

}

func (cr *CartRepository) GetCartByIdCart(id, UserID int) (entities.CartItem, error) {
	var carts entities.CartItem
	tx := cr.database.Where("id = ? and user_id = ?", id, UserID).First(&carts)

	if tx.Error != nil {
		return carts, tx.Error
	}
	return carts, nil

}

func (cr *CartRepository) UpdateCart(id, UserID int, cart entities.CartItem) error {

	tx := cr.database.Where("id = ? and user_id = ?", id, UserID).Updates(&cart)

	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (cr *CartRepository) DeleteCart(id, userID int) error {
	var carts entities.CartItem
	err := cr.database.Where("id =? and user_id = ?", id, userID).First(&carts).Error
	if err != nil {
		return err
	}
	cr.database.Delete(&carts)

	return nil

}
