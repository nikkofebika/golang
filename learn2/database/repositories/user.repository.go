package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"learn/database/entities"
	"learn/database/interfaces"
	"strconv"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepositoryInterface {
	return &userRepository{DB: db}
}

func (repository *userRepository) Insert(ctx context.Context, user entities.UserEntity) (entities.UserEntity, error) {
	query := "INSERT INTO users (name, is_admin) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, query, user.Name, user.IsAdmin)

	if err != nil {
		return user, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return user, err
	}

	user.Id = int32(id)
	return user, nil
}

func (repository *userRepository) FindAll(ctx context.Context) ([]entities.UserEntity, error) {
	query := "SELECT id, name, is_admin FROM users"
	rows, err := repository.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []entities.UserEntity{}
	for rows.Next() {
		user := entities.UserEntity{}
		rows.Scan(&user.Id, &user.Name, &user.IsAdmin)

		users = append(users, user)
	}

	return users, nil
}

func (repository *userRepository) FindById(ctx context.Context, id int) (entities.UserEntity, error) {
	query := "SELECT id, name, is_admin FROM users WHERE id=? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)

	user := entities.UserEntity{}
	if err != nil {
		return user, err
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.IsAdmin)
		if err != nil {
			return user, err
		}
		fmt.Println("user", user)
		return user, nil
	}

	return user, errors.New("user id " + strconv.Itoa(id) + " not found")
}
