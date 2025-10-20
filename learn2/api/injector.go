//go:build wireinject
// +build wireinject

package main

import (
	"learn/api/app"
	"learn/api/controllers"
	"learn/api/middlewares"
	"learn/api/repositories"
	"learn/api/services"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

func InitializeRouter() *http.Server {
	wire.Build(
		app.NewDb,
		ProvideValidator,

		NewServer,
		middlewares.NewAuthMiddleware,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		controllers.NewUserController,
		services.NewUserService,
		repositories.NewUserRepository,
	)
	return nil
}
