package helpers

import (
	"learn/api/entities"
	"learn/api/web"
)

func ToEntityResponse(userEntity entities.UserEntity) web.UserResponse {
	return web.UserResponse{
		Id:      userEntity.Id,
		Name:    userEntity.Name,
		IsAdmin: userEntity.IsAdmin,
	}
}

func ToEntitiesResponse(userEntities []entities.UserEntity) []web.UserResponse {
	var users []web.UserResponse
	for _, user := range userEntities {
		users = append(users, web.UserResponse{
			Id:      user.Id,
			Name:    user.Name,
			IsAdmin: user.IsAdmin,
		})
	}
	return users
}
