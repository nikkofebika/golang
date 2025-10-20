package main

import (
	"fmt"
	"learn/fiber/app/modules/users"
	"learn/fiber/config"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadDBConfig()
	db := config.NewDB(cfg)

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Prefork:      true,
	})

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	userController := users.NewUserController(userService)

	app.Get("users", userController.Index)
	app.Get("users/:id", userController.Show)
	app.Post("users", userController.Create)

	port := ":3000"
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}

	fmt.Println("APP RUNNING ON PORT ", port)
}

// func main() {
// 	app := fiber.New(fiber.Config{
// 		IdleTimeout:  time.Second * 5,
// 		ReadTimeout:  time.Second * 5,
// 		WriteTimeout: time.Second * 5,
// 		Prefork:      true,
// 	})

// 	app.Static("/static", "./public")

// 	// app.Get("/", func(ctx *fiber.Ctx) error {
// 	// 	return ctx.SendString("HELLO WORLD")
// 	// })

// 	helloWorld := func(ctx *fiber.Ctx) error {
// 		return ctx.JSON(fiber.Map{
// 			"Name": "Nikko",
// 			"Age":  100,
// 		})
// 	}

// 	api := app.Group("api")
// 	api.Get("users", helloWorld)

// 	// app.Get("/static/samsudin.jpg")
// 	// app.Get("/static/sample.txt")

// 	err := app.Listen(":4000")
// 	if err != nil {
// 		panic(err)
// 	}
// }
