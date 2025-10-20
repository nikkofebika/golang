package repositories

import (
	"context"
	"database/sql"
	"errors"
	"learn/api/entity"
	"learn/api/helper"
)

type UserRepositoryInterface interface {
	FindAll(ctx context.Context, tx *sql.Tx) []entity.UserEntity
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.UserEntity, error)
	Create(ctx context.Context, tx *sql.Tx, userEntity entity.UserEntity) entity.UserEntity
	Update(ctx context.Context, tx *sql.Tx, userEntity entity.UserEntity) entity.UserEntity
	Delete(ctx context.Context, tx *sql.Tx, id int)
}

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (repository *UserRepository) FindAll(ctx context.Context, tx *sql.Tx) []entity.UserEntity {
	query := "SELECT id, name, email, is_admin FROM users"

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	users := []entity.UserEntity{}
	for rows.Next() {
		user := entity.UserEntity{}
		err := rows.Scan(&user.ID, user.Name, user.Email, user.IsAdmin)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}

func (repository *UserRepository) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.UserEntity, error) {
	query := "SELECT id, name, email, is_admin FROM users"

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	user := entity.UserEntity{}
	if rows.Next() {
		err := rows.Scan(&user.ID, user.Name, user.Email, user.IsAdmin)
		helper.PanicIfError(err)

		return user, nil
	}

	return user, errors.New("user not found")
}

func (repository *UserRepository) Create(ctx context.Context, tx *sql.Tx, userEntity entity.UserEntity) entity.UserEntity {
	query := "INSERT INTO users (name, email, password, is_admin) VALUES(?, ?, ?, ?)"

	res, err := tx.ExecContext(ctx, query, userEntity.Name, userEntity.Email, userEntity.Password, userEntity.IsAdmin)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	userEntity.ID = int(id)
	return userEntity
}

func (repository *UserRepository) Update(ctx context.Context, tx *sql.Tx, userEntity entity.UserEntity) entity.UserEntity {
	query := "UPDATE users SET name=?, email=?, is_admin=? WHERE id=?"

	_, err := tx.ExecContext(ctx, query, userEntity.Name, userEntity.Email, userEntity.IsAdmin, userEntity.ID)
	helper.PanicIfError(err)

	return userEntity
}

func (repository *UserRepository) Delete(ctx context.Context, tx *sql.Tx, id int) {
	query := "DELETE FROM users WHERE id=?"

	_, err := tx.ExecContext(ctx, query, id)
	helper.PanicIfError(err)
}
