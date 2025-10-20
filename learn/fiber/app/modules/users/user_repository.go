package users

import (
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAll() ([]User, error)
	FindById(id uint) (User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db}
}

func (repository *userRepository) FindAll() ([]User, error) {
	var users []User
	err := repository.DB.Find(&users).Error
	return users, err
}

func (repository *userRepository) FindById(id uint) (User, error) {
	var user User
	err := repository.DB.Take(&user, id).Error
	return user, err
}

func (repository *userRepository) Create(user *User) error {
	return repository.DB.Create(user).Error
}

func (repository *userRepository) Update(user *User) error {
	return repository.DB.Save(user).Error
}

func (repository *userRepository) Delete(id uint) error {
	return repository.DB.Delete(&User{}, id).Error
}
