package controllers

import (
	"fmt"
	"learn/api/helpers"
	"learn/api/services"
	"learn/api/web"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerInterface interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserController struct {
	UserService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) UserControllerInterface {
	return &UserController{userService}
}

func (controller *UserController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("SILIT")
	userResponses := controller.UserService.FindAll(request.Context())

	response := helpers.JSONResponse{
		Code: http.StatusOK,
		Data: userResponses,
	}

	helpers.WriteResponseJSON(writer, response)
}

func (controller *UserController) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helpers.PanicIfError(err)

	userResponse := controller.UserService.FindByID(request.Context(), id)

	response := helpers.JSONResponse{
		Code: http.StatusOK,
		Data: userResponse,
	}

	helpers.WriteResponseJSON(writer, response)
}

func (controller *UserController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helpers.ReadRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)

	response := helpers.JSONResponse{
		Code: http.StatusCreated,
		Data: userResponse,
	}

	helpers.WriteResponseJSON(writer, response)
}
func (controller *UserController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helpers.ReadRequestBody(request, &userUpdateRequest)

	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helpers.PanicIfError(err)
	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)

	response := helpers.JSONResponse{
		Code: http.StatusOK,
		Data: userResponse,
	}

	helpers.WriteResponseJSON(writer, response)
}
func (controller *UserController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helpers.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)

	response := helpers.JSONResponse{
		Code: http.StatusOK,
	}

	helpers.WriteResponseJSON(writer, response)
}
