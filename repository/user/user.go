package user

import (
	"ecommerce/entities"
	"fmt"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetAll() ([]entities.User, error)
	GetUserById(id int) (entities.User, error)
	CreateUser(user entities.User) error
	DeleteUser(id int) error
	UpdateUser(id int, user entities.User) error
}

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetAll() ([]entities.User, error) {
	var users []entities.User
	tx := ur.database.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil

}

func (ur *UserRepository) GetUserById(id int) (entities.User, error) {
	var users entities.User
	tx := ur.database.Where("id = ?", id).First(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil

}

func (ur *UserRepository) CreateUser(user entities.User) error {

	tx := ur.database.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (ur *UserRepository) DeleteUser(id int) error {
	var users entities.User
	tx := ur.database.Where("id = ?", id).Delete(&users)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (ur *UserRepository) UpdateUser(id int, user entities.User) error {

	tx := ur.database.Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
