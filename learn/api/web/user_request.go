package web

type UserCreateRequest struct {
	Name     string
	Email    string
	Password string
	IsAdmin  bool
}

type UserUpdateRequest struct {
	ID       int
	Name     string
	Email    string
	Password string
	IsAdmin  bool
}
