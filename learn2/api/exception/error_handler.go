package exception

import (
	"learn/api/helpers"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		response := helpers.JSONResponse{
			Code: http.StatusInternalServerError,
			Data: exception.Error,
		}

		writer.WriteHeader(http.StatusNotFound)
		helpers.WriteResponseJSON(writer, response)

		return true
	}

	return false
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		response := helpers.JSONResponse{
			Code: http.StatusInternalServerError,
			Data: exception.Error(),
		}

		writer.WriteHeader(http.StatusUnprocessableEntity)
		helpers.WriteResponseJSON(writer, response)

		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	response := helpers.JSONResponse{
		Code: http.StatusInternalServerError,
		Data: err,
	}

	writer.WriteHeader(http.StatusInternalServerError)
	helpers.WriteResponseJSON(writer, response)
}
