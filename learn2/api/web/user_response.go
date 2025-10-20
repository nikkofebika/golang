package web

type UserResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
}
