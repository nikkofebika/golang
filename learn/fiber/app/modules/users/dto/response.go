package dto

import "learn/fiber/app/modules/users"

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	// CreatedAt string `json:"created_at"`
	// UpdatedAt string `json:"updated_at"`
}

func ToUserResponse(user users.User) UserResponse {
	return UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		// CreatedAt: user.CreatedAt.GoString(),
	}
}
