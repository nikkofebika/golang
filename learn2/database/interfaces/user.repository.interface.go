package interfaces

import (
	"context"
	"learn/database/entities"
)

type UserRepositoryInterface interface {
	FindAll(ctx context.Context) ([]entities.UserEntity, error)
	FindById(ctx context.Context, id int) (entities.UserEntity, error)
	Insert(ctx context.Context, user entities.UserEntity) (entities.UserEntity, error)
}
