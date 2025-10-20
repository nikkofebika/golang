package repositories

import (
	"context"
	"database/sql"
	"errors"
	"learn/api/entities"
	"learn/api/helpers"
)

type UserRepositoryInterface interface {
	FindAll(ctx context.Context, tx *sql.Tx) []entities.UserEntity
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entities.UserEntity, error)
	Create(ctx context.Context, tx *sql.Tx, userEntity entities.UserEntity) entities.UserEntity
	Update(ctx context.Context, tx *sql.Tx, userEntity entities.UserEntity) entities.UserEntity
	Delete(ctx context.Context, tx *sql.Tx, id int)
}

type UserRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (repository *UserRepository) FindAll(ctx context.Context, tx *sql.Tx) []entities.UserEntity {
	query := "SELECT id, name, is_admin FROM users"

	rows, err := tx.QueryContext(ctx, query)
	helpers.PanicIfError(err)

	users := []entities.UserEntity{}

	for rows.Next() {
		user := entities.UserEntity{}
		err := rows.Scan(&user.Id, &user.Name, &user.IsAdmin)
		helpers.PanicIfError(err)

		users = append(users, user)
	}

	return users
}

func (repository *UserRepository) FindByID(ctx context.Context, tx *sql.Tx, id int) (entities.UserEntity, error) {
	query := "SELECT id, name, is_admin FROM users WHERE id ? LIMIT 1"

	rows, err := tx.QueryContext(ctx, query, id)
	helpers.PanicIfError(err)

	user := entities.UserEntity{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.IsAdmin)
		helpers.PanicIfError(err)

		return user, nil
	}

	return user, errors.New("user not found")
}

func (repository *UserRepository) Create(ctx context.Context, tx *sql.Tx, userEntity entities.UserEntity) entities.UserEntity {
	query := "INSERT INTO users (name, is_admin) VALUES (?, ?)"

	result, err := tx.ExecContext(ctx, query, userEntity.Name, userEntity.IsAdmin)
	helpers.PanicIfError(err)

	id, err := result.LastInsertId()
	helpers.PanicIfError(err)

	userEntity.Id = int(id)

	return userEntity
}

func (repository *UserRepository) Update(ctx context.Context, tx *sql.Tx, userEntity entities.UserEntity) entities.UserEntity {
	query := "UPDATE users name=?, is_admin=? WHERE id=?"

	_, err := tx.ExecContext(ctx, query, userEntity.Name, userEntity.IsAdmin, userEntity.Id)
	helpers.PanicIfError(err)

	return userEntity
}

func (repository *UserRepository) Delete(ctx context.Context, tx *sql.Tx, id int) {
	query := "DELETE FROM users WHERE id=?"

	_, err := tx.ExecContext(ctx, query, id)
	helpers.PanicIfError(err)
}
