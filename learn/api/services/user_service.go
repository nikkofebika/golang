package services

import (
	"context"
	"database/sql"
	"learn/api/entity"
	"learn/api/helper"
	"learn/api/repositories"
	"learn/api/web"

	"github.com/go-playground/validator/v10"
)

type UserServiceInterface interface {
	FindAll(ctx context.Context) []web.UserResponse
	FindById(ctx context.Context, id int) web.UserResponse
	Create(ctx context.Context, request web.UserRequest) web.UserResponse
	Update(ctx context.Context, request web.UserRequest) web.UserResponse
	Delete(ctx context.Context, id int)
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
	DB             *sql.DB
	Validate       *validator.Validate
}

func (userService *UserService) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := userService.UserRepository.FindAll(ctx, tx)

	return helper.ToResponses(users)
}

func (userService *UserService) FindById(ctx context.Context, id int) web.UserResponse {
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := userService.UserRepository.FindById(ctx, tx, id)

	helper.PanicIfError(err)

	return helper.ToResponse(user)
}

func (userService *UserService) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userEntity := entity.UserEntity{Name: request.Name, Email: request.Email, Password: request.Password, IsAdmin: request.IsAdmin}

	user := userService.UserRepository.Create(ctx, tx, userEntity)

	return helper.ToResponse(user)
}

func (userService *UserService) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userEntity, err := userService.UserRepository.FindById(ctx, tx, request.ID)
	helper.PanicIfError(err)
	// if err != nil {
	// exception here
	// }

	userEntity.Name = request.Name
	userEntity.Email = request.Email
	userEntity.IsAdmin = request.IsAdmin
	// if request.Password != nil {

	// }

	user := userService.UserRepository.Update(ctx, tx, userEntity)

	return helper.ToResponse(user)
}

func (userService *UserService) Delete(ctx context.Context, id int) {
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userService.UserRepository.Delete(ctx, tx, id)
}
