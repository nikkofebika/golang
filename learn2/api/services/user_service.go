package services

import (
	"context"
	"database/sql"
	"learn/api/entities"
	"learn/api/exception"
	"learn/api/helpers"
	"learn/api/repositories"
	"learn/api/web"

	"github.com/go-playground/validator/v10"
)

type UserServiceInterface interface {
	FindAll(ctx context.Context) []web.UserResponse
	FindByID(ctx context.Context, id int) web.UserResponse
	Create(ctx context.Context, userCreateRequest web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, userUpdateRequest web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, id int)
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
	DB             *sql.DB
	validator      *validator.Validate
}

func NewUserService(userRepository repositories.UserRepositoryInterface, db *sql.DB, validator *validator.Validate) UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
		DB:             db,
		validator:      validator,
	}
}

func (service *UserService) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return helpers.ToEntitiesResponse(users)
}

func (service *UserService) FindByID(ctx context.Context, id int) web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByID(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helpers.ToEntityResponse(user)
}

func (service *UserService) Create(ctx context.Context, userCreateRequest web.UserCreateRequest) web.UserResponse {
	err := service.validator.Struct(userCreateRequest)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	userEntity := entities.UserEntity{
		Name:    userCreateRequest.Name,
		IsAdmin: userCreateRequest.IsAdmin,
	}

	user := service.UserRepository.Create(ctx, tx, userEntity)

	return helpers.ToEntityResponse(user)
}

func (service *UserService) Update(ctx context.Context, userUpdateRequest web.UserUpdateRequest) web.UserResponse {
	err := service.validator.Struct(userUpdateRequest)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByID(ctx, tx, userUpdateRequest.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Name = userUpdateRequest.Name
	user.IsAdmin = userUpdateRequest.IsAdmin

	user = service.UserRepository.Update(ctx, tx, user)

	return helpers.ToEntityResponse(user)
}

func (service *UserService) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	_, err = service.UserRepository.FindByID(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, id)
}
