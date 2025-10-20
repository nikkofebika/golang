package users

import "learn/fiber/app/modules/users/dto"

type UserServiceInterface interface {
	FindAll() ([]dto.UserResponse, error)
	FindById(id uint) (dto.UserResponse, error)
	Create(request dto.UserCreateRequest) error
	Update(user *User) error
	Delete(id uint) error
}

type userService struct {
	repository UserRepositoryInterface
}

func NewUserService(repository UserRepositoryInterface) UserServiceInterface {
	return &userService{repository}
}

func (service *userService) FindAll() ([]dto.UserResponse, error) {
	users, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *userService) FindById(id uint) (dto.UserResponse, error) {
	user, err := service.repository.FindById(id)

	return user, err
}

func (service *userService) Create(request dto.UserCreateRequest) error {
	user := User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		IsAdmin:  request.IsAdmin,
	}

	return service.repository.Create(&user)
}

func (service *userService) Update(user *User) error {
	return service.repository.Update(user)
}

func (service *userService) Delete(id uint) error {
	return service.repository.Delete(id)
}
