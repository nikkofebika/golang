package helper

import (
	"learn/api/entity"
	"learn/api/web"
)

func ToResponse(entity entity.UserEntity) web.UserResponse {
	return web.UserResponse{
		ID:      entity.ID,
		Name:    entity.Name,
		Email:   entity.Email,
		IsAdmin: entity.IsAdmin,
	}
}

func ToResponses(entities []entity.UserEntity) []web.UserResponse {
	responses := []web.UserResponse{}
	for _, entity := range entities {
		responses = append(responses, ToResponse(entity))
	}

	return responses
}
