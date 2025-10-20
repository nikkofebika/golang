package main

import (
	"learn/api/middlewares"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(middleware *middlewares.AuthMiddleware) *http.Server {
	port := "8000"
	return &http.Server{
		Addr:    "localhost:" + port,
		Handler: middleware,
	}
}

func ProvideValidator() *validator.Validate {
	return validator.New()
}

func main() {
	// db := app.NewDb()
	// defer db.Close()

	// validat := ProvideValidator()

	// userRepository := repositories.NewUserRepository()
	// userService := services.NewUserService(userRepository, db, validat)
	// userController := controllers.NewUserController(userService)
	// router := app.NewRouter(userController)

	// middleware := middlewares.NewAuthMiddleware(router)

	// server := NewServer(middleware)
	server := InitializeRouter()
	log.Println("Server running on port 8000")
	err := server.ListenAndServe()
	log.Fatal(err)
}
