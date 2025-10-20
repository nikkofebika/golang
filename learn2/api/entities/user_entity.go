package entities

type UserEntity struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
}
