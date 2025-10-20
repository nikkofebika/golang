package middlewares

import (
	"fmt"
	"learn/api/helpers"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("API KEY", request.Header.Get("X-API-KEY"))
	if request.Header.Get("X-API-KEY") == "secret" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.WriteHeader(http.StatusUnauthorized)

		response := helpers.JSONResponse{
			Code: http.StatusUnauthorized,
		}

		helpers.WriteResponseJSON(writer, response)
	}
}
