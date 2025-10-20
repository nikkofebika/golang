package database

import (
	"context"
	"fmt"
	"learn/database/entities"
	"learn/database/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	repository := repositories.NewUserRepository(GetConnection())

	ctx := context.Background()
	user := entities.UserEntity{
		Name:    "Nikko FE",
		IsAdmin: true,
	}
	user, err := repository.Insert(ctx, user)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, "Nikko FE", user.Name)
	assert.Equal(t, true, user.IsAdmin)
}

func TestFindById(t *testing.T) {
	repository := repositories.NewUserRepository(GetConnection())

	ctx := context.Background()
	user, err := repository.FindById(ctx, 1)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, int32(1), user.Id)
	assert.Equal(t, "Nikko FE", user.Name)
	assert.Equal(t, true, user.IsAdmin)
}

func TestFindAll(t *testing.T) {
	repository := repositories.NewUserRepository(GetConnection())

	ctx := context.Background()
	users, err := repository.FindAll(ctx)

	if err != nil {
		t.Fatal(err.Error())
	}

	for _, user := range users {
		fmt.Println("===========")
		fmt.Println("id :", user.Id)
		fmt.Println("name :", user.Name)
		fmt.Println("is_admin :", user.IsAdmin)
		fmt.Println("===========")
	}

	// assert.Equal(t, 1, user.Id)
	// assert.Equal(t, "Nikko FE", user.Name)
	// assert.Equal(t, true, user.IsAdmin)
}
