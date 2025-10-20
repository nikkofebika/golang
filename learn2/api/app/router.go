package app

import (
	"learn/api/controllers"
	"learn/api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controllers.UserControllerInterface) *httprouter.Router {
	router := httprouter.New()
	router.GET("/users", controller.FindAll)
	router.GET("/users/:id", controller.FindByID)
	router.POST("/users", controller.Create)
	router.PUT("/users/:id", controller.Update)
	router.DELETE("/users", controller.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
