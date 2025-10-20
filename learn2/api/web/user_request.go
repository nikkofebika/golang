package web

type UserCreateRequest struct {
	Name    string `validate:"required,min=1,max=100"`
	IsAdmin bool   `validate:"required,boolean"`
}

type UserUpdateRequest struct {
	Id      int    `validate:"required"`
	Name    string `validate:"required,min=1,max=100"`
	IsAdmin bool   `validate:"required,boolean"`
}
